package auth

import (
	"fmt"

	"github.com/urfave/cli/v3"
)

// GetToken returns a valid OAuth access token for the given environment.
// Returns ("", nil) if no credentials are stored (no-op — the API will return 401).
// Automatically refreshes expired tokens when a refresh token is available.
func GetToken(environment string) (string, error) {
	creds, err := LoadCredentials()
	if err != nil {
		return "", fmt.Errorf("loading credentials: %w", err)
	}

	tokens, ok := creds[environment]
	if !ok || tokens == nil {
		return "", nil
	}

	// Token is still valid.
	if !tokens.IsExpired() {
		return tokens.AccessToken, nil
	}

	// Token is expired — try to refresh.
	if tokens.RefreshToken == "" {
		// No refresh token; clear stale credentials.
		_ = ClearCredentials(environment)
		return "", fmt.Errorf("session expired, please run 'mercury login' to re-authenticate")
	}

	config := DefaultOAuthConfig(environment)
	newTokens, err := RefreshToken(config, tokens.RefreshToken)
	if err != nil {
		// Refresh failed; clear stale credentials.
		_ = ClearCredentials(environment)
		return "", fmt.Errorf("session expired (refresh failed: %v), please run 'mercury login' to re-authenticate", err)
	}

	// Save the refreshed tokens.
	creds[environment] = newTokens
	if err := SaveCredentials(creds); err != nil {
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
