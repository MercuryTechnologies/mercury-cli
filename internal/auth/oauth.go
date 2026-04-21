package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	// OAuth client IDs per environment (public clients, no secret).
	ProductionOAuthClientID = "d6a5310d-f441-4a36-b12a-2fc09c380936"
	SandboxOAuthClientID    = "da707cac-28d3-4003-bbc6-c7384ca6557a"

	// OAuth endpoints per environment.
	ProductionAuthURL  = "https://oauth2.mercury.com/oauth2/auth"
	ProductionTokenURL = "https://oauth2.mercury.com/oauth2/token"
	SandboxAuthURL     = "https://oauth2-sandbox.mercury.com/oauth2/auth"
	SandboxTokenURL    = "https://oauth2-sandbox.mercury.com/oauth2/token"

	// OAuthScopes requested during login.
	OAuthScopes = "offline_access openid read write"
)

// OAuthConfig holds the OAuth configuration for a specific environment.
type OAuthConfig struct {
	ClientID string
	AuthURL  string
	TokenURL string
}

// DefaultOAuthConfig returns the OAuth configuration for the given environment.
func DefaultOAuthConfig(environment string) *OAuthConfig {
	switch environment {
	case "sandbox":
		return &OAuthConfig{
			ClientID: SandboxOAuthClientID,
			AuthURL:  SandboxAuthURL,
			TokenURL: SandboxTokenURL,
		}
	default:
		return &OAuthConfig{
			ClientID: ProductionOAuthClientID,
			AuthURL:  ProductionAuthURL,
			TokenURL: ProductionTokenURL,
		}
	}
}

// Login performs the OAuth Authorization Code + PKCE flow.
// It starts a localhost server, opens the browser for authorization,
// waits for the callback, and exchanges the code for tokens.
func Login(ctx context.Context, config *OAuthConfig) (*TokenSet, error) {
	// Generate PKCE code verifier and challenge.
	verifier, challenge, err := generatePKCE()
	if err != nil {
		return nil, fmt.Errorf("generating PKCE: %w", err)
	}

	// Generate random state parameter.
	state, err := generateState()
	if err != nil {
		return nil, fmt.Errorf("generating state: %w", err)
	}

	// Start localhost callback server on a random port.
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, fmt.Errorf("starting callback server: %w", err)
	}
	port := listener.Addr().(*net.TCPAddr).Port
	redirectURI := fmt.Sprintf("http://127.0.0.1:%d/callback", port)

	// Channel to receive the authorization code or error.
	type callbackResult struct {
		code string
		err  error
	}
	resultCh := make(chan callbackResult, 1)

	mux := http.NewServeMux()
	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		if errMsg := r.URL.Query().Get("error"); errMsg != "" {
			desc := r.URL.Query().Get("error_description")
			title, detail, link := friendlyOAuthError(errMsg, desc)
			renderError(w, title, detail, link)
			resultCh <- callbackResult{err: fmt.Errorf("OAuth error: %s - %s", errMsg, desc)}
			return
		}

		if returnedState := r.URL.Query().Get("state"); returnedState != state {
			renderError(w, "That link didn't look right", "For security, the sign-in link couldn't be verified. Please try again from your terminal.", nil)
			resultCh <- callbackResult{err: fmt.Errorf("state mismatch")}
			return
		}

		code := r.URL.Query().Get("code")
		if code == "" {
			renderError(w, "Something went wrong", "We didn't receive a sign-in code from Mercury. Please try again from your terminal.", nil)
			resultCh <- callbackResult{err: fmt.Errorf("no authorization code in callback")}
			return
		}

		renderSuccess(w)
		resultCh <- callbackResult{code: code}
	})

	server := &http.Server{Handler: mux}
	go server.Serve(listener)
	defer server.Close()

	// Build authorization URL.
	authURL, err := buildAuthURL(config, redirectURI, state, challenge)
	if err != nil {
		return nil, fmt.Errorf("building auth URL: %w", err)
	}

	// Open browser.
	if err := openBrowser(authURL); err != nil {
		// If browser fails, print the URL for manual copy-paste.
		fmt.Printf("\nCould not open browser. Please visit this URL to log in:\n\n  %s\n\n", authURL)
	}

	// Wait for callback with timeout.
	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	var result callbackResult
	select {
	case result = <-resultCh:
	case <-ctx.Done():
		return nil, fmt.Errorf("login timed out waiting for browser authorization")
	}

	if result.err != nil {
		return nil, result.err
	}

	// Exchange authorization code for tokens.
	return exchangeCode(config, result.code, redirectURI, verifier)
}

// RefreshToken exchanges a refresh token for new tokens.
func RefreshToken(config *OAuthConfig, refreshToken string) (*TokenSet, error) {
	data := url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {refreshToken},
		"client_id":     {config.ClientID},
	}

	return doTokenRequest(config.TokenURL, data, refreshToken)
}

// friendlyOAuthError maps OAuth error codes and descriptions to user-facing
// copy for the callback error page. Unknown errors fall back to the provider's
// description.
func friendlyOAuthError(errCode, desc string) (title, detail string, link *errorLink) {
	if strings.Contains(strings.ToLower(desc), "scope is not allowed") {
		return "Your Mercury account doesn't have API access",
			"Employee accounts can't use the Mercury API. Ask an admin on your team to change your role to Admin or Custom User in Mercury's Team Settings, then try again.",
			&errorLink{URL: "https://app.mercury.com/settings/users", Text: "Open Team Settings"}
	}
	if desc != "" {
		return "Sign-in didn't complete", desc, nil
	}
	return "Sign-in didn't complete", fmt.Sprintf("Mercury returned %q. Please try again from your terminal.", errCode), nil
}

// generatePKCE creates a PKCE code verifier and its S256 challenge.
func generatePKCE() (verifier, challenge string, err error) {
	buf := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return "", "", err
	}
	verifier = base64.RawURLEncoding.EncodeToString(buf)
	h := sha256.Sum256([]byte(verifier))
	challenge = base64.RawURLEncoding.EncodeToString(h[:])
	return verifier, challenge, nil
}

// generateState creates a random state parameter.
func generateState() (string, error) {
	buf := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}

// buildAuthURL constructs the full authorization URL.
func buildAuthURL(config *OAuthConfig, redirectURI, state, challenge string) (string, error) {
	u, err := url.Parse(config.AuthURL)
	if err != nil {
		return "", err
	}
	q := u.Query()
	q.Set("response_type", "code")
	q.Set("client_id", config.ClientID)
	q.Set("redirect_uri", redirectURI)
	q.Set("scope", OAuthScopes)
	q.Set("state", state)
	q.Set("code_challenge", challenge)
	q.Set("code_challenge_method", "S256")
	u.RawQuery = q.Encode()
	return u.String(), nil
}

// exchangeCode exchanges an authorization code for tokens.
func exchangeCode(config *OAuthConfig, code, redirectURI, verifier string) (*TokenSet, error) {
	data := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {redirectURI},
		"client_id":     {config.ClientID},
		"code_verifier": {verifier},
	}

	return doTokenRequest(config.TokenURL, data, "")
}

// tokenResponse is the raw response from the token endpoint.
type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	Error        string `json:"error"`
	ErrorDesc    string `json:"error_description"`
}

// doTokenRequest makes a POST to the token endpoint and parses the response.
func doTokenRequest(tokenURL string, data url.Values, fallbackRefreshToken string) (*TokenSet, error) {
	resp, err := http.Post(tokenURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("token request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading token response: %w", err)
	}

	var tok tokenResponse
	if err := json.Unmarshal(body, &tok); err != nil {
		return nil, fmt.Errorf("parsing token response: %w", err)
	}

	if tok.Error != "" {
		return nil, fmt.Errorf("token error: %s - %s", tok.Error, tok.ErrorDesc)
	}

	if tok.AccessToken == "" {
		return nil, fmt.Errorf("no access token in response")
	}

	expiry := time.Now().Add(time.Duration(tok.ExpiresIn) * time.Second)

	refreshToken := tok.RefreshToken
	if refreshToken == "" {
		refreshToken = fallbackRefreshToken
	}

	return &TokenSet{
		AccessToken:  tok.AccessToken,
		RefreshToken: refreshToken,
		TokenType:    tok.TokenType,
		Expiry:       expiry,
	}, nil
}

// openBrowser opens the given URL in the user's default browser.
func openBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	return cmd.Start()
}
