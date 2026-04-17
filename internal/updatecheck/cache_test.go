package updatecheck

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCacheRoundTrip(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "update-check.json")

	want := cacheEntry{
		CheckedAt:     time.Date(2026, 4, 16, 14, 20, 0, 0, time.UTC),
		LatestVersion: "0.4.0",
	}
	if err := writeCache(path, want); err != nil {
		t.Fatalf("writeCache: %v", err)
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat: %v", err)
	}
	if mode := info.Mode().Perm(); mode != 0600 {
		t.Errorf("cache file perms = %o; want 0600", mode)
	}

	got, ok := readCache(path)
	if !ok {
		t.Fatal("readCache returned ok=false")
	}
	if got.LatestVersion != want.LatestVersion {
		t.Errorf("latest = %q; want %q", got.LatestVersion, want.LatestVersion)
	}
	if !got.CheckedAt.Equal(want.CheckedAt) {
		t.Errorf("checked_at = %v; want %v", got.CheckedAt, want.CheckedAt)
	}
}

func TestReadCacheMissing(t *testing.T) {
	_, ok := readCache(filepath.Join(t.TempDir(), "missing.json"))
	if ok {
		t.Error("expected ok=false for missing file")
	}
}

func TestReadCacheCorrupt(t *testing.T) {
	path := filepath.Join(t.TempDir(), "corrupt.json")
	if err := os.WriteFile(path, []byte("{not json"), 0600); err != nil {
		t.Fatal(err)
	}
	_, ok := readCache(path)
	if ok {
		t.Error("expected ok=false for corrupt JSON")
	}
}

func TestReadCacheEmptyVersion(t *testing.T) {
	path := filepath.Join(t.TempDir(), "empty.json")
	if err := os.WriteFile(path, []byte(`{"checked_at":"2026-01-01T00:00:00Z","latest_version":""}`), 0600); err != nil {
		t.Fatal(err)
	}
	_, ok := readCache(path)
	if ok {
		t.Error("expected ok=false when latest_version is empty")
	}
}
