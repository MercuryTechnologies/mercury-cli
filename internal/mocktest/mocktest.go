package mocktest

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var mockServerURL *url.URL

func init() {
	mockServerURL, _ = url.Parse("http://localhost:4010")
	if testURL := os.Getenv("TEST_API_BASE_URL"); testURL != "" {
		if parsed, err := url.Parse(testURL); err == nil {
			mockServerURL = parsed
		}
	}
}

// OnlyMockServerDialer only allows network connections to the mock server
type OnlyMockServerDialer struct{}

func (d *OnlyMockServerDialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	if address == mockServerURL.Host {
		return (&net.Dialer{}).DialContext(ctx, network, address)
	}

	return nil, fmt.Errorf("BLOCKED: connection to %s not allowed (only allowed: %s)", address, mockServerURL.Host)
}

func blockNetworkExceptMockServer() (http.RoundTripper, http.RoundTripper) {
	restricted := &http.Transport{
		DialContext: (&OnlyMockServerDialer{}).DialContext,
	}

	origClient, origDefault := http.DefaultClient.Transport, http.DefaultTransport
	http.DefaultClient.Transport, http.DefaultTransport = restricted, restricted
	return origClient, origDefault
}

func restoreNetwork(origClient, origDefault http.RoundTripper) {
	http.DefaultClient.Transport, http.DefaultTransport = origClient, origDefault
}

// TestRunMockTestWithFlags runs a test against a mock server with the provided
// CLI flags and ensures it succeeds
func TestRunMockTestWithFlags(t *testing.T, flags ...string) {
	origClient, origDefault := blockNetworkExceptMockServer()
	defer restoreNetwork(origClient, origDefault)

	// Check if mock server is running
	conn, err := net.DialTimeout("tcp", mockServerURL.Host, 2*time.Second)
	if err != nil {
		require.Fail(t, "Mock server is not running on "+mockServerURL.Host+". Please start the mock server before running tests.")
	} else {
		conn.Close()
	}

	// Get the path to the main command
	_, filename, _, ok := runtime.Caller(0)
	require.True(t, ok, "Could not get current file path")
	dirPath := filepath.Dir(filename)
	project := filepath.Join(dirPath, "..", "..", "cmd", "...")

	args := []string{"run", project, "--base-url", mockServerURL.String()}
	args = append(args, flags...)

	t.Logf("Testing command: mercury %s", strings.Join(args[4:], " "))

	cliCmd := exec.Command("go", args...)

	// Pipe the CLI tool's output into `head` so it doesn't hang when simulating
	// paginated or streamed endpoints. 100 lines of output should be enough to
	// test that the API endpoint worked, or report back a meaningful amount of
	// data if something went wrong.
	headCmd := exec.Command("head", "-n", "100")
	pipe, err := cliCmd.StdoutPipe()
	require.NoError(t, err, "Failed to create pipe for CLI command")
	headCmd.Stdin = pipe

	// Capture `head` output and CLI command stderr outputs:
	var output strings.Builder
	headCmd.Stdout = &output
	headCmd.Stderr = &output
	cliCmd.Stderr = &output

	// First start `head`, so it's ready for data to come in:
	err = headCmd.Start()
	require.NoError(t, err, "Failed to start `head` command")

	// Next start the CLI command so it can pipe data to `head` without
	// buffering any data in advance:
	err = cliCmd.Start()
	require.NoError(t, err, "Failed to start CLI command")

	// Ensure that the stdout pipe is closed as soon as `head` exits, to let the
	// CLI tool know that no more output is needed and it can stop streaming
	// test data for streaming/paginated endpoints. This needs to happen before
	// calling `cliCmd.Wait()`, otherwise there will be a deadlock.
	err = headCmd.Wait()
	pipe.Close()
	require.NoError(t, err, "`head` command finished with an error")

	// Finally, wait for the CLI tool to finish up:
	err = cliCmd.Wait()
	require.NoError(t, err, "CLI command failed\n%s", output.String())

	t.Logf("Test passed successfully\nOutput:\n%s", output.String())
}

func TestFile(t *testing.T, contents string) string {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "file.txt")
	require.NoError(t, os.WriteFile(filename, []byte(contents), 0644))
	return filename
}
