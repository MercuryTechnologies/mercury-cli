// Package winupgrade installs checksum-verified Windows release binaries.
package winupgrade

import (
	"archive/zip"
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const maxArchiveSize = 128 << 20
const releaseBaseURL = "https://cli.mercury.com"

var versionPattern = regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+(?:-[0-9A-Za-z.-]+)?$`)

// Install returns the path of any retained previous executable. On Windows,
// a running image can be renamed but cannot be overwritten or deleted.
func Install(ctx context.Context, version, arch, destination string) (string, error) {
	return install(ctx, releaseClient(10*time.Minute), releaseBaseURL, version, arch, destination)
}

// LatestVersion uses an interactive request timeout, independent of the short
// background update-notification deadline.
func LatestVersion(ctx context.Context) (string, error) {
	return latestVersion(ctx, versionClient(), releaseBaseURL+"/VERSION")
}

func versionClient() *http.Client {
	return releaseClient(30 * time.Second)
}

func releaseClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if req.URL.Scheme != "https" {
				return fmt.Errorf("refusing release redirect to non-HTTPS URL")
			}
			if len(via) >= 10 {
				return fmt.Errorf("stopped after 10 release redirects")
			}
			return nil
		},
	}
}

func latestVersion(ctx context.Context, client *http.Client, url string) (string, error) {
	body, err := download(ctx, client, url, 64)
	if err != nil {
		return "", err
	}
	version := strings.TrimPrefix(strings.TrimSpace(string(body)), "v")
	if !versionPattern.MatchString(version) {
		return "", fmt.Errorf("invalid release version %q", version)
	}
	return version, nil
}

func install(ctx context.Context, client *http.Client, baseURL, version, arch, destination string) (string, error) {
	version = strings.TrimPrefix(strings.TrimSpace(version), "v")
	if !versionPattern.MatchString(version) {
		return "", fmt.Errorf("invalid release version %q", version)
	}
	switch arch {
	case "386", "amd64", "arm64":
	default:
		return "", fmt.Errorf("unsupported Windows architecture %q", arch)
	}
	name := fmt.Sprintf("mercury_%s_windows_%s.zip", version, arch)
	root := baseURL + "/v" + version + "/"
	checksums, err := download(ctx, client, root+"mercury_"+version+"_checksums.txt", 1<<20)
	if err != nil {
		return "", err
	}
	archive, err := download(ctx, client, root+name, maxArchiveSize)
	if err != nil {
		return "", err
	}
	if err := verifyChecksum(archive, checksums, name); err != nil {
		return "", err
	}
	binary, err := extractBinary(archive)
	if err != nil {
		return "", err
	}
	if err := ctx.Err(); err != nil {
		return "", err
	}
	return replaceBinary(destination, binary, os.Rename)
}

func download(ctx context.Context, client *http.Client, url string, limit int64) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("download %s: HTTP %d", url, resp.StatusCode)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, limit+1))
	if err != nil {
		return nil, err
	}
	if int64(len(body)) > limit {
		return nil, fmt.Errorf("download exceeds %d bytes", limit)
	}
	return body, nil
}

func verifyChecksum(archive, checksums []byte, name string) error {
	want := ""
	for _, line := range strings.Split(string(checksums), "\n") {
		fields := strings.Fields(line)
		if len(fields) == 2 && strings.TrimPrefix(fields[1], "*") == name {
			if want != "" {
				return fmt.Errorf("duplicate checksum for %s", name)
			}
			want = fields[0]
		}
	}
	digest, err := hex.DecodeString(want)
	if err != nil || len(digest) != sha256.Size {
		return fmt.Errorf("missing or invalid SHA-256 checksum for %s", name)
	}
	got := sha256.Sum256(archive)
	if !bytes.Equal(got[:], digest) {
		return fmt.Errorf("SHA-256 checksum mismatch for %s", name)
	}
	return nil
}

func extractBinary(archive []byte) ([]byte, error) {
	z, err := zip.NewReader(bytes.NewReader(archive), int64(len(archive)))
	if err != nil {
		return nil, fmt.Errorf("open release ZIP: %w", err)
	}
	var binary *zip.File
	for _, f := range z.File {
		// Never extract arbitrary archive paths, symlinks, or directories.
		if f.Name != "mercury.exe" {
			continue
		}
		if binary != nil || !f.Mode().IsRegular() || f.UncompressedSize64 > maxArchiveSize {
			return nil, fmt.Errorf("invalid mercury.exe entry in release ZIP")
		}
		binary = f
	}
	if binary == nil {
		return nil, fmt.Errorf("release ZIP does not contain mercury.exe")
	}
	r, err := binary.Open()
	if err != nil {
		return nil, err
	}
	defer r.Close()
	body, err := io.ReadAll(io.LimitReader(r, maxArchiveSize+1))
	if err != nil {
		return nil, err
	}
	if len(body) == 0 || len(body) > maxArchiveSize {
		return nil, fmt.Errorf("invalid mercury.exe size")
	}
	return body, nil
}

func replaceBinary(destination string, binary []byte, rename func(string, string) error) (string, error) {
	if err := os.MkdirAll(filepath.Dir(destination), 0o755); err != nil {
		return "", err
	}
	stage, err := os.MkdirTemp(filepath.Dir(destination), ".mercury-upgrade-")
	if err != nil {
		return "", err
	}
	keep := false
	defer func() {
		if !keep {
			_ = os.RemoveAll(stage)
		}
	}()
	newPath := filepath.Join(stage, "mercury.exe")
	f, err := os.OpenFile(newPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o755)
	if err != nil {
		return "", err
	}
	if _, err := f.Write(binary); err != nil {
		_ = f.Close()
		return "", err
	}
	// Flush the staged executable before moving the previous one aside.
	if err := f.Sync(); err != nil {
		_ = f.Close()
		return "", err
	}
	if err := f.Close(); err != nil {
		return "", err
	}
	backup := filepath.Join(stage, "previous.exe")
	if info, err := os.Lstat(destination); err == nil {
		if !info.Mode().IsRegular() {
			return "", fmt.Errorf("install destination is not a regular file: %s", destination)
		}
		if err := rename(destination, backup); err != nil {
			return "", fmt.Errorf("move previous executable: %w", err)
		}
	} else if os.IsNotExist(err) {
		backup = ""
	} else {
		return "", err
	}
	if err := rename(newPath, destination); err != nil {
		if backup != "" {
			if rollbackErr := rename(backup, destination); rollbackErr != nil {
				keep = true
				return backup, fmt.Errorf("install failed: %v; restore %s manually: %w", err, backup, rollbackErr)
			}
		}
		return "", fmt.Errorf("install executable: %w", err)
	}
	if backup != "" {
		if err := os.Remove(backup); err != nil {
			keep = true
		} else {
			backup = ""
		}
	}
	return backup, nil
}
