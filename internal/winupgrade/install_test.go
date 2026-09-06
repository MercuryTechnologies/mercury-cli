package winupgrade

import (
	"archive/zip"
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func zipBytes(t *testing.T, names ...string) []byte {
	t.Helper()
	var b bytes.Buffer
	w := zip.NewWriter(&b)
	for _, name := range names {
		f, err := w.Create(name)
		if err != nil {
			t.Fatal(err)
		}
		if strings.HasSuffix(name, "/") {
			continue
		}
		if _, err = f.Write([]byte("new executable")); err != nil {
			t.Fatal(err)
		}
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	return b.Bytes()
}

func TestInstall(t *testing.T) {
	for _, tc := range []struct {
		name, checksum string
		status         int
		wantErr        bool
	}{
		{"valid", "valid", 200, false},
		{"corrupt", "bad", 200, true},
		{"missing checksum", "missing", 200, true},
		{"duplicate checksum", "duplicate", 200, true},
		{"HTTP failure", "valid", 404, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			archive := zipBytes(t, "mercury.exe", "../unwanted")
			name := "mercury_1.2.3_windows_amd64.zip"
			line := fmt.Sprintf("%x  %s\n", sha256.Sum256(archive), name)
			switch tc.checksum {
			case "bad":
				line = strings.Repeat("0", 64) + "  " + name
			case "missing":
				line = ""
			case "duplicate":
				line += line
			}
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch r.URL.Path {
				case "/v1.2.3/mercury_1.2.3_checksums.txt":
					fmt.Fprint(w, line)
				case "/v1.2.3/" + name:
					w.WriteHeader(tc.status)
					_, _ = w.Write(archive)
				default:
					t.Errorf("unexpected path %s", r.URL.Path)
					http.NotFound(w, r)
				}
			}))
			defer s.Close()
			dir := t.TempDir()
			dest := filepath.Join(dir, "with spaces & symbols", "mercury.exe")
			if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(dest, []byte("old"), 0o755); err != nil {
				t.Fatal(err)
			}
			_, err := install(context.Background(), s.Client(), s.URL, "v1.2.3", "amd64", dest)
			if (err != nil) != tc.wantErr {
				t.Fatalf("error = %v", err)
			}
			got, err := os.ReadFile(dest)
			if err != nil {
				t.Fatal(err)
			}
			want := "new executable"
			if tc.wantErr {
				want = "old"
			}
			if string(got) != want {
				t.Fatalf("got %q, want %q", got, want)
			}
			if _, err := os.Stat(filepath.Join(dir, "unwanted")); !os.IsNotExist(err) {
				t.Fatal("extracted unexpected file")
			}
		})
	}
}

func TestInvalidInputs(t *testing.T) {
	for _, v := range []string{"../1.2.3", "1.2.3/other", "", "1.2.3?x"} {
		if _, err := install(context.Background(), nil, "", v, "amd64", ""); err == nil {
			t.Errorf("accepted %q", v)
		}
	}
	if _, err := install(context.Background(), nil, "", "1.2.3", "mips", ""); err == nil {
		t.Fatal("accepted unsupported architecture")
	}
	for _, names := range [][]string{{"../mercury.exe"}, {"mercury.exe", "mercury.exe"}, {"mercury.exe/"}} {
		if _, err := extractBinary(zipBytes(t, names...)); err == nil {
			t.Fatalf("accepted %v", names)
		}
	}
	for _, mode := range []os.FileMode{os.ModeDir | 0o755, os.ModeSymlink | 0o777} {
		var archive bytes.Buffer
		w := zip.NewWriter(&archive)
		header := &zip.FileHeader{Name: "mercury.exe"}
		header.SetMode(mode)
		if _, err := w.CreateHeader(header); err != nil {
			t.Fatal(err)
		}
		if err := w.Close(); err != nil {
			t.Fatal(err)
		}
		if _, err := extractBinary(archive.Bytes()); err == nil {
			t.Fatalf("accepted non-regular entry with mode %v", mode)
		}
	}
}

func TestRollback(t *testing.T) {
	for _, failRollback := range []bool{false, true} {
		t.Run(fmt.Sprint(failRollback), func(t *testing.T) {
			dest := filepath.Join(t.TempDir(), "mercury.exe")
			if err := os.WriteFile(dest, []byte("old"), 0o755); err != nil {
				t.Fatal(err)
			}
			calls := 0
			backup, err := replaceBinary(dest, []byte("new"), func(a, b string) error {
				calls++
				if calls == 2 || (calls == 3 && failRollback) {
					return errors.New("injected rename failure")
				}
				return os.Rename(a, b)
			})
			if err == nil {
				t.Fatal("expected failure")
			}
			recoverPath := dest
			if failRollback {
				recoverPath = backup
			}
			got, readErr := os.ReadFile(recoverPath)
			if readErr != nil || string(got) != "old" {
				t.Fatalf("lost previous executable: %q, %v", got, readErr)
			}
		})
	}
}

func TestDownloadLimitsAndCancellation(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "12345")
	}))
	defer s.Close()
	if _, err := download(context.Background(), s.Client(), s.URL, 4); err == nil {
		t.Fatal("accepted oversized response")
	}
	if got, err := download(context.Background(), s.Client(), s.URL, 5); err != nil || string(got) != "12345" {
		t.Fatalf("exact limit: %q, %v", got, err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := download(ctx, s.Client(), s.URL, 5); !errors.Is(err, context.Canceled) {
		t.Fatalf("canceled download: %v", err)
	}
}

func TestNewDestinationAndMoveFailure(t *testing.T) {
	dest := filepath.Join(t.TempDir(), "new directory", "mercury.exe")
	if backup, err := replaceBinary(dest, []byte("old"), os.Rename); err != nil || backup != "" {
		t.Fatalf("new destination: backup=%q, error=%v", backup, err)
	}
	if _, err := replaceBinary(dest, []byte("new"), func(string, string) error {
		return errors.New("access denied")
	}); err == nil {
		t.Fatal("expected move failure")
	}
	if got, err := os.ReadFile(dest); err != nil || string(got) != "old" {
		t.Fatalf("previous executable: %q, %v", got, err)
	}
}

func TestReplaceRunningExecutable(t *testing.T) {
	if os.Getenv("MERCURY_TEST_SELF_REPLACE") == "1" {
		dest, err := os.Executable()
		if err != nil {
			t.Fatal(err)
		}
		backup, err := replaceBinary(dest, []byte("replacement"), os.Rename)
		if err != nil {
			t.Fatal(err)
		}
		if runtime.GOOS == "windows" && backup == "" {
			t.Fatal("running executable backup should be retained")
		}
		got, err := os.ReadFile(dest)
		if err != nil || string(got) != "replacement" {
			t.Fatalf("replacement: %q %v", got, err)
		}
		return
	}
	self, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	b, err := os.ReadFile(self)
	if err != nil {
		t.Fatal(err)
	}
	dest := filepath.Join(t.TempDir(), "mercury test & upgrade.exe")
	if err := os.WriteFile(dest, b, 0o755); err != nil {
		t.Fatal(err)
	}
	c := exec.Command(dest, "-test.run=^TestReplaceRunningExecutable$")
	c.Env = append(os.Environ(), "MERCURY_TEST_SELF_REPLACE=1")
	if out, err := c.CombinedOutput(); err != nil {
		t.Fatalf("self replacement: %v\n%s", err, out)
	}
}
