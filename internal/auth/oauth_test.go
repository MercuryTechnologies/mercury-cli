package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

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
