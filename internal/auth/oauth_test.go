package auth

import (
	"bufio"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// RFC 6749 §5.1 defines refresh_token as OPTIONAL in the token response, and
// §6 requires the client to keep using the previously issued refresh token
// when the response omits one. A non-rotating server must not cause us to
// silently lose refresh capability.
func TestRefreshToken_PreservesRefreshTokenWhenServerOmitsIt(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"access_token": "new-access-token",
			"token_type": "Bearer",
			"expires_in": 3600
		}`))
	}))
	defer server.Close()

	config := &OAuthConfig{
		ClientID: "test-client",
		TokenURL: server.URL,
	}

	const originalRefresh = "original-refresh-token"
	newTokens, err := RefreshToken(config, originalRefresh)
	require.NoError(t, err)

	assert.Equal(t, "new-access-token", newTokens.AccessToken)
	assert.Equal(t, originalRefresh, newTokens.RefreshToken,
		"when the server omits refresh_token, the client must reuse the original (RFC 6749 §6)")
}

func TestFriendlyOAuthError(t *testing.T) {
	t.Parallel()

	scopeDetail := "Employee accounts can't use the Mercury API. Ask an admin on your team to change your role to Admin or Custom User in Mercury's Team Settings, then try again."
	scopeLink := &errorLink{URL: "https://app.mercury.com/settings/users", Text: "Open Team Settings"}

	cases := []struct {
		name       string
		errCode    string
		desc       string
		wantTitle  string
		wantDetail string
		wantLink   *errorLink
	}{
		{
			name:       "scope not allowed maps to API access message",
			errCode:    "login_request_denied",
			desc:       "The requested scope is not allowed",
			wantTitle:  "Your Mercury account doesn't have API access",
			wantDetail: scopeDetail,
			wantLink:   scopeLink,
		},
		{
			name:       "scope not allowed is case-insensitive",
			errCode:    "login_request_denied",
			desc:       "the requested SCOPE IS NOT ALLOWED here",
			wantTitle:  "Your Mercury account doesn't have API access",
			wantDetail: scopeDetail,
			wantLink:   scopeLink,
		},
		{
			name:       "unknown error with description falls through",
			errCode:    "access_denied",
			desc:       "User denied consent.",
			wantTitle:  "Sign-in didn't complete",
			wantDetail: "User denied consent.",
			wantLink:   nil,
		},
		{
			name:       "unknown error with no description shows code",
			errCode:    "server_error",
			desc:       "",
			wantTitle:  "Sign-in didn't complete",
			wantDetail: `Mercury returned "server_error". Please try again from your terminal.`,
			wantLink:   nil,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			title, detail, link := friendlyOAuthError(tc.errCode, tc.desc)
			assert.Equal(t, tc.wantTitle, title)
			assert.Equal(t, tc.wantDetail, detail)
			assert.Equal(t, tc.wantLink, link)
		})
	}
}

// The callback server must be configured with all four timeout knobs the
// net/http package documents. Before this guard, http.Server{} was constructed
// with only Handler set, so a co-resident process could hold connections open
// indefinitely against the OAuth callback receiver.
func TestNewCallbackServer_ConfiguresAllTimeouts(t *testing.T) {
	t.Parallel()

	s := newCallbackServer(http.NewServeMux())

	assert.Equal(t, callbackReadHeaderTimeout, s.ReadHeaderTimeout, "ReadHeaderTimeout must be set")
	assert.Equal(t, callbackReadTimeout, s.ReadTimeout, "ReadTimeout must be set")
	assert.Equal(t, callbackWriteTimeout, s.WriteTimeout, "WriteTimeout must be set")
	assert.Equal(t, callbackIdleTimeout, s.IdleTimeout, "IdleTimeout must be set")

	// Sanity: timeouts must be non-zero. A zero value silently disables the
	// guard, which is the bug this PR fixes.
	assert.Greater(t, s.ReadHeaderTimeout, time.Duration(0))
	assert.Greater(t, s.ReadTimeout, time.Duration(0))
	assert.Greater(t, s.WriteTimeout, time.Duration(0))
	assert.Greater(t, s.IdleTimeout, time.Duration(0))
}

// End-to-end check: a slow-header (slowloris) client must be disconnected
// within ReadHeaderTimeout instead of holding the listener until Login's
// 2-minute context deadline fires. Skipped in -short mode because it waits
// for the timeout to elapse.
func TestCallbackServer_SlowlorisIsClosedWithinReadHeaderTimeout(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip("waits for ReadHeaderTimeout to fire; skipping in -short mode")
	}

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	server := newCallbackServer(http.NewServeMux())
	go func() { _ = server.Serve(listener) }()
	t.Cleanup(func() { _ = server.Close() })

	conn, err := net.Dial("tcp", listener.Addr().String())
	require.NoError(t, err)
	t.Cleanup(func() { _ = conn.Close() })

	// Send the request line and one header byte, then stall. Without
	// ReadHeaderTimeout, the server would wait forever for the rest.
	_, err = conn.Write([]byte("GET /callback HTTP/1.1\r\nHost: 127.0.0.1\r\n"))
	require.NoError(t, err)

	// The server should respond with 408 Request Timeout (or just close the
	// connection) within ReadHeaderTimeout + slack. We assert we observe EOF
	// or a response well before Login's 2-minute context would have fired.
	deadline := time.Now().Add(callbackReadHeaderTimeout + 3*time.Second)
	require.NoError(t, conn.SetReadDeadline(deadline))

	start := time.Now()
	reader := bufio.NewReader(conn)
	_, _ = reader.ReadByte() // returns when the server responds or closes
	elapsed := time.Since(start)

	assert.Less(t, elapsed, callbackReadHeaderTimeout+2*time.Second,
		"slowloris client must be disconnected within ReadHeaderTimeout; elapsed=%s", elapsed)
}

// When the server does rotate the refresh token, we should adopt the new one.
func TestRefreshToken_AdoptsRotatedRefreshToken(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"access_token": "new-access-token",
			"refresh_token": "rotated-refresh-token",
			"token_type": "Bearer",
			"expires_in": 3600
		}`))
	}))
	defer server.Close()

	config := &OAuthConfig{
		ClientID: "test-client",
		TokenURL: server.URL,
	}

	newTokens, err := RefreshToken(config, "original-refresh-token")
	require.NoError(t, err)

	assert.Equal(t, "rotated-refresh-token", newTokens.RefreshToken)
}
