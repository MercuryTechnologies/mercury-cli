package updatecheck

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func newServerReturning(t *testing.T, status int, body string) (*httptest.Server, *int32) {
	t.Helper()
	var hits int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&hits, 1)
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(srv.Close)
	return srv, &hits
}

func runCheck(t *testing.T, opts Options) (string, *Checker) {
	t.Helper()
	c := Start(context.Background(), opts)
	var buf bytes.Buffer
	c.Notify(&buf)
	return buf.String(), c
}

func baseOpts(t *testing.T, current, url string) Options {
	t.Helper()
	return Options{
		CurrentVersion: current,
		VersionURL:     url,
		CachePath:      filepath.Join(t.TempDir(), "update-check.json"),
		Now:            func() time.Time { return time.Date(2026, 4, 16, 12, 0, 0, 0, time.UTC) },
		LookupEnv:      func(string) (string, bool) { return "", false },
	}
}

func TestNotifyShowsBannerWhenNewer(t *testing.T) {
	srv, hits := newServerReturning(t, 200, "0.4.0\n")
	out, _ := runCheck(t, baseOpts(t, "0.3.2", srv.URL))

	if *hits != 1 {
		t.Errorf("expected 1 HTTP hit, got %d", *hits)
	}
	if !strings.Contains(out, "0.3.2") || !strings.Contains(out, "0.4.0") {
		t.Errorf("expected banner with 0.3.2 → 0.4.0; got %q", out)
	}
	if !strings.Contains(out, "mercury upgrade") {
		t.Errorf("expected banner to mention 'mercury upgrade'; got %q", out)
	}
	if !strings.Contains(out, "releases/tag/v0.4.0") {
		t.Errorf("expected release-notes URL; got %q", out)
	}
}

func TestNotifySilentWhenUpToDate(t *testing.T) {
	srv, _ := newServerReturning(t, 200, "0.3.2")
	out, _ := runCheck(t, baseOpts(t, "0.3.2", srv.URL))
	if out != "" {
		t.Errorf("expected no output when versions equal; got %q", out)
	}
}

func TestNotifySilentWhenCurrentIsAhead(t *testing.T) {
	srv, _ := newServerReturning(t, 200, "0.3.2")
	out, _ := runCheck(t, baseOpts(t, "1.0.0", srv.URL))
	if out != "" {
		t.Errorf("expected no output when current > latest; got %q", out)
	}
}

func TestFreshCacheSkipsHTTP(t *testing.T) {
	srv, hits := newServerReturning(t, 200, "9.9.9") // should not be seen
	opts := baseOpts(t, "0.3.2", srv.URL)
	if err := writeCache(opts.CachePath, cacheEntry{
		CheckedAt:     opts.Now().Add(-1 * time.Hour),
		LatestVersion: "0.4.0",
	}); err != nil {
		t.Fatal(err)
	}

	out, _ := runCheck(t, opts)

	if *hits != 0 {
		t.Errorf("expected no HTTP hit with fresh cache, got %d", *hits)
	}
	if !strings.Contains(out, "0.4.0") {
		t.Errorf("expected cached 0.4.0 in banner; got %q", out)
	}
}

func TestStaleCacheTriggersHTTP(t *testing.T) {
	srv, hits := newServerReturning(t, 200, "0.5.0")
	opts := baseOpts(t, "0.3.2", srv.URL)
	if err := writeCache(opts.CachePath, cacheEntry{
		CheckedAt:     opts.Now().Add(-48 * time.Hour),
		LatestVersion: "0.4.0",
	}); err != nil {
		t.Fatal(err)
	}

	out, _ := runCheck(t, opts)

	if *hits != 1 {
		t.Errorf("expected 1 HTTP hit with stale cache, got %d", *hits)
	}
	if !strings.Contains(out, "0.5.0") {
		t.Errorf("expected refreshed 0.5.0 in banner; got %q", out)
	}
}

func TestHTTPErrorFallsBackToCache(t *testing.T) {
	srv, _ := newServerReturning(t, 500, "oops")
	opts := baseOpts(t, "0.3.2", srv.URL)
	if err := writeCache(opts.CachePath, cacheEntry{
		CheckedAt:     opts.Now().Add(-48 * time.Hour),
		LatestVersion: "0.4.0",
	}); err != nil {
		t.Fatal(err)
	}

	out, _ := runCheck(t, opts)
	if !strings.Contains(out, "0.4.0") {
		t.Errorf("expected stale-cache fallback to render 0.4.0; got %q", out)
	}
}

func TestHTTPErrorNoCacheIsSilent(t *testing.T) {
	srv, _ := newServerReturning(t, 500, "oops")
	out, _ := runCheck(t, baseOpts(t, "0.3.2", srv.URL))
	if out != "" {
		t.Errorf("expected silence when both HTTP and cache fail; got %q", out)
	}
}

func TestSuppressedByEnv(t *testing.T) {
	srv, hits := newServerReturning(t, 200, "0.4.0")
	opts := baseOpts(t, "0.3.2", srv.URL)
	opts.LookupEnv = func(k string) (string, bool) {
		if k == disableEnvVar {
			return "1", true
		}
		return "", false
	}

	out, _ := runCheck(t, opts)
	if *hits != 0 {
		t.Errorf("expected no HTTP call when env var is set, got %d", *hits)
	}
	if out != "" {
		t.Errorf("expected silence when env var is set; got %q", out)
	}
}

func TestSuppressedByDevVersion(t *testing.T) {
	srv, hits := newServerReturning(t, 200, "0.4.0")
	out, _ := runCheck(t, baseOpts(t, "dev", srv.URL))
	if *hits != 0 {
		t.Errorf("expected no HTTP call for dev version, got %d", *hits)
	}
	if out != "" {
		t.Errorf("expected silence for dev version; got %q", out)
	}
}

func TestSuppressedByIntrospectiveArg(t *testing.T) {
	srv, hits := newServerReturning(t, 200, "0.4.0")
	opts := baseOpts(t, "0.3.2", srv.URL)
	opts.Args = []string{"--version"}

	out, _ := runCheck(t, opts)
	if *hits != 0 {
		t.Errorf("expected no HTTP call when --version is in args, got %d", *hits)
	}
	if out != "" {
		t.Errorf("expected silence for --version; got %q", out)
	}
}

func TestSuppressedByUpgradeSubcommand(t *testing.T) {
	opts := baseOpts(t, "0.3.2", "http://unused")
	opts.Args = []string{"upgrade"}
	out, _ := runCheck(t, opts)
	if out != "" {
		t.Errorf("expected silence during upgrade; got %q", out)
	}
}

func TestInvalidUpstreamVersionIsSilent(t *testing.T) {
	srv, _ := newServerReturning(t, 200, "not-a-semver")
	out, _ := runCheck(t, baseOpts(t, "0.3.2", srv.URL))
	if out != "" {
		t.Errorf("expected silence on garbage upstream; got %q", out)
	}
}

func TestWriteCachePersistsAfterFetch(t *testing.T) {
	srv, _ := newServerReturning(t, 200, "0.4.0")
	opts := baseOpts(t, "0.3.2", srv.URL)
	runCheck(t, opts)

	entry, ok := readCache(opts.CachePath)
	if !ok {
		t.Fatal("expected cache to be written")
	}
	if entry.LatestVersion != "0.4.0" {
		t.Errorf("cached latest = %q; want 0.4.0", entry.LatestVersion)
	}
}
