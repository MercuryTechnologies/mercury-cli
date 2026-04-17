package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/MercuryTechnologies/mercury-cli/internal/auth"
	"github.com/charmbracelet/lipgloss"
	"github.com/urfave/cli/v3"
)

var (
	authTitle = lipgloss.NewStyle().Foreground(colorBlue).Bold(true)
	authDim   = lipgloss.NewStyle().Foreground(colorDim)
	authValue = lipgloss.NewStyle().Foreground(colorLight)
	authRule  = lipgloss.NewStyle().Foreground(colorFrame)
)

var authLogin = cli.Command{
	Name:     "login",
	Usage:    "Sign in to Mercury in your browser",
	Category: "Auth",
	Action:   handleLogin,
}

var authLogout = cli.Command{
	Name:     "logout",
	Usage:    "Sign out and delete saved tokens",
	Category: "Auth",
	Action:   handleLogout,
}

var authStatus = cli.Command{
	Name:     "status",
	Usage:    "Show Mercury sign-in status per environment",
	Category: "Auth",
	Action:   handleStatus,
}

func handleLogin(ctx context.Context, cmd *cli.Command) error {
	environment := auth.ResolveEnvironment(cmd)
	config := auth.DefaultOAuthConfig(environment)

	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, authDim.Render("  Opening browser to log in to Mercury")+authValue.Render(" ("+environment+")")+authDim.Render("..."))

	tokens, err := auth.Login(ctx, config)
	if err != nil {
		return fmt.Errorf("login failed: %w", err)
	}

	// Save credentials.
	creds, err := auth.LoadCredentials()
	if err != nil {
		creds = auth.Credentials{}
	}
	creds[environment] = tokens
	if err := auth.SaveCredentials(creds); err != nil {
		return fmt.Errorf("saving credentials: %w", err)
	}

	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, authRule.Render("  ──── ")+authTitle.Render("Logged in")+authRule.Render(" ────"))
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "  "+authDim.Render("  Environment  ")+authValue.Render(environment))
	fmt.Fprintln(os.Stderr, "  "+authDim.Render("  Expires      ")+authValue.Render(tokens.Expiry.Local().Format("Jan 02, 2006 3:04 PM")))
	fmt.Fprintln(os.Stderr)

	return nil
}

func handleLogout(ctx context.Context, cmd *cli.Command) error {
	environment := auth.ResolveEnvironment(cmd)

	if err := auth.ClearCredentials(environment); err != nil {
		return fmt.Errorf("clearing credentials: %w", err)
	}

	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, authDim.Render("  Logged out of Mercury")+authValue.Render(" ("+environment+")")+authDim.Render("."))
	fmt.Fprintln(os.Stderr)

	return nil
}

func envTokenStatus(tokens *auth.TokenSet) string {
	if tokens == nil {
		return "not logged in"
	}
	if tokens.IsExpired() {
		if tokens.RefreshToken != "" {
			return "expired (will auto-refresh)"
		}
		return "expired"
	}
	return "logged in, expires " + tokens.Expiry.Local().Format("Jan 02, 2006 3:04 PM")
}

func handleStatus(ctx context.Context, cmd *cli.Command) error {
	creds, err := auth.LoadCredentials()
	if err != nil {
		return fmt.Errorf("loading credentials: %w", err)
	}

	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, authRule.Render("  ──── ")+authTitle.Render("Authentication Status")+authRule.Render(" ────"))
	fmt.Fprintln(os.Stderr)

	if cmd.IsSet("api-key") || os.Getenv("MERCURY_API_KEY") != "" {
		fmt.Fprintln(os.Stderr, "  "+authDim.Render("  API Key      ")+authValue.Render("set (takes precedence)"))
	}

	fmt.Fprintln(os.Stderr, "  "+authDim.Render("  Production   ")+authValue.Render(envTokenStatus(creds["production"])))
	fmt.Fprintln(os.Stderr, "  "+authDim.Render("  Sandbox      ")+authValue.Render(envTokenStatus(creds["sandbox"])))
	fmt.Fprintln(os.Stderr)

	return nil
}
