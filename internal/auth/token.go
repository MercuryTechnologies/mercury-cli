package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/urfave/cli/v3"
)

// getTokenRefreshTimeout bounds GetToken's silent refresh attempt. doTokenRequest
// imposes a tighter per-request cap; this value just prevents wedged retries
// from holding up the calling command.
const getTokenRefreshTimeout = 15 * time.Second

// GetToken returns a valid OAuth access token for the given environment.
// Returns ("", nil) if no credentials are stored (no-op — the API will return 401).
// Automatically refreshes expired tokens when a refresh token is available.
func GetToken(environment string) (string, error) {
	tokens, err := LoadToken(environment)
	if err != nil {
		return "", fmt.Errorf("loading credentials: %w", err)
	}
	if tokens == nil {
		return "", nil
	}

	if !tokens.IsExpired() {
		return tokens.AccessToken, nil
	}

	if tokens.RefreshToken == "" {
		_ = ClearToken(environment)
		return "", fmt.Errorf("session expired, please run 'mercury login' to re-authenticate")
	}

	config := DefaultOAuthConfig(environment)
	// GetToken is invoked from getDefaultRequestOptions(cmd), which is called
	// from Stainless-generated handlers and does not propagate the request
	// context. Build a local deadline so a stalled OAuth backend can't hang
	// every CLI invocation.
	ctx, cancel := context.WithTimeout(context.Background(), getTokenRefreshTimeout)
	defer cancel()
	newTokens, err := RefreshToken(ctx, config, tokens.RefreshToken)
	if err != nil {
		_ = ClearToken(environment)
		return "", fmt.Errorf("session expired (refresh failed: %v), please run 'mercury login' to re-authenticate", err)
	}

	if _, err := SaveToken(environment, newTokens); err != nil {
		// Non-fatal — we still have the new token in memory.
		fmt.Printf("Warning: could not save refreshed credentials: %v\n", err)
	}

	return newTokens.AccessToken, nil
}

// ResolveEnvironment determines the target environment from the CLI command flags.
// Defaults to "production" if not specified.
func ResolveEnvironment(cmd *cli.Command) string {
	if env := cmd.String("environment"); env != "" {
		return env
	}
	return "production"
}
