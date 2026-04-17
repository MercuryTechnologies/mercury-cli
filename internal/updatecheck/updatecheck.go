// Package updatecheck implements an unobtrusive "new version available" notifier
// for the Mercury CLI, modelled after gh/flyctl/deno.
//
// The check runs in a background goroutine kicked off at the start of a CLI
// invocation. At the end of the invocation the caller joins the goroutine with
// a short deadline via Notify; if a strictly-newer release is known, a short
// banner is written to the provided writer (typically os.Stderr). A local cache
// keeps most invocations off the network entirely.
package updatecheck

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

const (
	defaultVersionURL = "https://cli.mercury.com/VERSION"
	defaultCacheTTL   = 24 * time.Hour
	httpTimeout       = 2 * time.Second
	joinTimeout       = 200 * time.Millisecond
	releaseURLPrefix  = "https://github.com/MercuryTechnologies/mercury-cli/releases/tag/v"
	disableEnvVar     = "MERCURY_NO_UPDATE_CHECK"
)

// suppressedArgs lists tokens that, when present anywhere in the CLI argv,
// mean we should skip the update check for this invocation.
var suppressedArgs = map[string]struct{}{
	"--version":   {},
	"-v":          {},
	"version":     {},
	"help":        {},
	"--help":      {},
	"-h":          {},
	"@manpages":   {},
	"@completion": {},
	"__complete":  {},
	"upgrade":     {},
}

// Options configures a Checker. Fields left at their zero value pick reasonable
// defaults; tests can override any of them.
type Options struct {
	CurrentVersion string
	Args           []string
	VersionURL     string
	CachePath      string
	CacheTTL       time.Duration
	HTTPClient     *http.Client
	Now            func() time.Time
	// LookupEnv overrides os.LookupEnv, for tests.
	LookupEnv func(string) (string, bool)
}

// Checker is the handle returned by Start. It is safe to call Notify at most
// once per Checker.
type Checker struct {
	result  chan string // latest version string (or "" if unknown / suppressed)
	current string
}

// Start kicks off a background update check and returns immediately.
// The returned *Checker is never nil; Notify on a suppressed Checker is a no-op.
func Start(ctx context.Context, opts Options) *Checker {
	opts.applyDefaults()

	c := &Checker{
		result:  make(chan string, 1),
		current: opts.CurrentVersion,
	}

	if suppress(opts) {
		c.result <- ""
		return c
	}

	go func() {
		defer func() {
			// Never let a panic in the background check take down the CLI.
			if r := recover(); r != nil {
				select {
				case c.result <- "":
				default:
				}
			}
		}()
		c.result <- resolveLatest(ctx, opts)
	}()

	return c
}

// Notify waits up to ~200ms for the background check to finish and, if a
// strictly-newer release is known, writes a short banner to w. It never
// returns an error — update notifications must never disrupt the command.
func (c *Checker) Notify(w io.Writer) {
	if c == nil {
		return
	}

	var latest string
	select {
	case latest = <-c.result:
	case <-time.After(joinTimeout):
		return
	}
	if latest == "" {
		return
	}

	cur, ok := parseSemver(c.current)
	if !ok {
		return
	}
	next, ok := parseSemver(latest)
	if !ok {
		return
	}
	if !cur.less(next) {
		return
	}

	writeBanner(w, c.current, latest)
}

func (o *Options) applyDefaults() {
	if o.VersionURL == "" {
		o.VersionURL = defaultVersionURL
	}
	if o.CacheTTL == 0 {
		o.CacheTTL = defaultCacheTTL
	}
	if o.HTTPClient == nil {
		o.HTTPClient = &http.Client{Timeout: httpTimeout}
	}
	if o.Now == nil {
		o.Now = time.Now
	}
	if o.LookupEnv == nil {
		o.LookupEnv = os.LookupEnv
	}
	if o.CachePath == "" {
		if p, err := defaultCachePath(); err == nil {
			o.CachePath = p
		}
	}
}

func suppress(opts Options) bool {
	if v, ok := opts.LookupEnv(disableEnvVar); ok && v != "" && v != "0" && !strings.EqualFold(v, "false") {
		return true
	}
	if _, ok := parseSemver(opts.CurrentVersion); !ok {
		return true
	}
	for _, a := range opts.Args {
		if _, hit := suppressedArgs[a]; hit {
			return true
		}
	}
	return false
}

// resolveLatest returns the latest known upstream version, or "" on any error.
// It consults the local cache first; only refetches when the cache is stale.
func resolveLatest(ctx context.Context, opts Options) string {
	if entry, ok := readCache(opts.CachePath); ok {
		if opts.Now().Sub(entry.CheckedAt) < opts.CacheTTL {
			return entry.LatestVersion
		}
	}

	latest, err := fetchLatest(ctx, opts)
	if err != nil || latest == "" {
		// Fall back to whatever's cached even if stale.
		if entry, ok := readCache(opts.CachePath); ok {
			return entry.LatestVersion
		}
		return ""
	}

	if opts.CachePath != "" {
		_ = writeCache(opts.CachePath, cacheEntry{
			CheckedAt:     opts.Now(),
			LatestVersion: latest,
		})
	}
	return latest
}

func fetchLatest(ctx context.Context, opts Options) (string, error) {
	reqCtx, cancel := context.WithTimeout(ctx, httpTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(reqCtx, http.MethodGet, opts.VersionURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "mercury-cli/"+opts.CurrentVersion)

	resp, err := opts.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("update check: unexpected status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(io.LimitReader(resp.Body, 64))
	if err != nil {
		return "", err
	}
	v := strings.TrimSpace(string(body))
	if _, ok := parseSemver(v); !ok {
		return "", fmt.Errorf("update check: invalid version %q", v)
	}
	return v, nil
}

// FetchLatest is an on-demand, uncached version lookup. Used by `mercury upgrade`
// to decide whether there's anything to install.
func FetchLatest(ctx context.Context) (string, error) {
	return fetchLatest(ctx, Options{
		VersionURL: defaultVersionURL,
		HTTPClient: &http.Client{Timeout: httpTimeout},
	})
}

func writeBanner(w io.Writer, current, latest string) {
	dim := lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	light := lipgloss.NewStyle().Foreground(lipgloss.Color("252")).Bold(true)

	// Disable ANSI styling when the writer isn't a TTY so we don't leak
	// escape codes into files, pipes, or CI logs.
	if !writerIsTerminal(w) {
		dim = lipgloss.NewStyle()
		light = lipgloss.NewStyle()
	}

	fmt.Fprintln(w)
	fmt.Fprintf(w, "%s %s %s %s\n",
		dim.Render("A new release of mercury is available:"),
		light.Render(current),
		dim.Render("→"),
		light.Render(latest),
	)
	fmt.Fprintf(w, "%s %s\n",
		dim.Render("To upgrade, run:"),
		light.Render("mercury upgrade"),
	)
	fmt.Fprintln(w, dim.Render(releaseURLPrefix+latest))
}

func writerIsTerminal(w io.Writer) bool {
	f, ok := w.(*os.File)
	if !ok {
		return false
	}
	return term.IsTerminal(f.Fd())
}
