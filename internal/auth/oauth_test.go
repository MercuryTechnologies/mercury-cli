package auth

import (
	"context"
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
	newTokens, err := RefreshToken(context.Background(), config, originalRefresh)
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

	newTokens, err := RefreshToken(context.Background(), config, "original-refresh-token")
	require.NoError(t, err)

	assert.Equal(t, "rotated-refresh-token", newTokens.RefreshToken)
}

// A stalled token endpoint must not hang the CLI: the caller's context
// deadline must propagate down to the HTTP request. Before this guard was
// added, doTokenRequest used http.Post with http.DefaultClient (no Timeout
// and no context), so login + refresh would block until the kernel TCP
// timeout — minutes, not seconds.
func TestRefreshToken_RespectsCallerContextDeadline(t *testing.T) {
	t.Parallel()

	// Server that holds the connection open until the test tears it down.
	stop := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		select {
		case <-stop:
		case <-r.Context().Done():
		}
	}))
	t.Cleanup(func() {
		close(stop)
		server.Close()
	})

	config := &OAuthConfig{ClientID: "test-client", TokenURL: server.URL}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	start := time.Now()
	_, err := RefreshToken(ctx, config, "any-refresh-token")
	elapsed := time.Since(start)

	require.Error(t, err)
	assert.Less(t, elapsed, 2*time.Second,
		"RefreshToken must abort within the caller's context deadline; elapsed=%s", elapsed)
}

// Even when the caller passes context.Background() (no deadline), the
// per-request cap inside doTokenRequest must fire so a degraded token
// endpoint can't block indefinitely. Skipped in -short mode because it
// burns the full tokenRequestTimeout window.
func TestRefreshToken_PerRequestTimeoutAppliesWithoutCallerDeadline(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip("relies on the per-request cap; skipping in -short mode")
	}

	stop := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		select {
		case <-stop:
		case <-r.Context().Done():
		}
	}))
	t.Cleanup(func() {
		close(stop)
		server.Close()
	})

	config := &OAuthConfig{ClientID: "test-client", TokenURL: server.URL}

	start := time.Now()
	_, err := RefreshToken(context.Background(), config, "any-refresh-token")
	elapsed := time.Since(start)

	require.Error(t, err)
	// The cap is tokenRequestTimeout (10s). Allow generous slack for slow CI
	// but assert we're nowhere near the kernel TCP timeout (which is minutes).
	assert.Less(t, elapsed, tokenRequestTimeout+5*time.Second,
		"per-request timeout must bound RefreshToken; elapsed=%s", elapsed)
}
