package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/MercuryTechnologies/mercury-cli/internal/updatecheck"
	"github.com/MercuryTechnologies/mercury-cli/internal/winupgrade"
	"github.com/urfave/cli/v3"
)

const installScriptURL = "https://cli.mercury.com/install.sh"

var upgrade = cli.Command{
	Name:     "upgrade",
	Usage:    "Upgrade mercury to the latest release",
	Category: "Utility",
	Description: "Downloads and installs the latest mercury CLI release (using the official install script on macOS/Linux).\n\n" +
		"Honors the env vars MERCURY_INSTALL_DIR (to override the install location) and\n" +
		"MERCURY_VERSION (to pin a specific version). The --version flag sets MERCURY_VERSION.",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "force",
			Aliases: []string{"f"},
			Usage:   "Re-install even if already on the latest version",
		},
		&cli.StringFlag{
			Name:  "version",
			Usage: "Install a specific version (e.g. 0.3.1) instead of the latest",
		},
	},
	Action: runUpgrade,
}

func runUpgrade(ctx context.Context, c *cli.Command) error {
	if runtime.GOOS == "windows" {
		return runWindowsUpgrade(ctx, c)
	}

	pinned := c.String("version")
	force := c.Bool("force")

	if pinned == "" && !force {
		latest, err := updatecheck.FetchLatest(ctx)
		if err == nil && latest == Version {
			fmt.Fprintf(os.Stdout, "Already at the latest version (%s). Use --force to reinstall.\n", Version)
			return nil
		}
	}

	if _, err := exec.LookPath("sh"); err != nil {
		return cli.Exit("mercury upgrade requires `sh` on PATH", 1)
	}
	if _, err := exec.LookPath("curl"); err != nil {
		return cli.Exit("mercury upgrade requires `curl` on PATH", 1)
	}

	cmd := exec.CommandContext(ctx, "sh", "-c", fmt.Sprintf("curl -fsSL %s | sh", installScriptURL))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Env = os.Environ()
	if pinned != "" {
		cmd.Env = append(cmd.Env, "MERCURY_VERSION="+pinned)
	}

	if err := cmd.Run(); err != nil {
		return cli.Exit(fmt.Sprintf("upgrade failed: %v", err), 1)
	}
	return nil
}

func runWindowsUpgrade(ctx context.Context, c *cli.Command) error {
	version := c.String("version")
	if version == "" {
		version = os.Getenv("MERCURY_VERSION")
	}
	if version == "" {
		latest, err := updatecheck.FetchLatest(ctx)
		if err != nil {
			return fmt.Errorf("fetch latest release: %w", err)
		}
		if latest == Version && !c.Bool("force") {
			fmt.Fprintf(os.Stdout, "Already at the latest version (%s). Use --force to reinstall.\n", Version)
			return nil
		}
		version = latest
	}
	destination, err := os.Executable()
	if err != nil {
		return err
	}
	if dir := os.Getenv("MERCURY_INSTALL_DIR"); dir != "" {
		destination, err = filepath.Abs(filepath.Join(dir, "mercury.exe"))
		if err != nil {
			return err
		}
	} else {
		destination, err = filepath.EvalSymlinks(destination)
		if err != nil {
			return err
		}
	}
	fmt.Fprintf(os.Stdout, "Installing mercury %s to %s...\n", version, destination)
	backup, err := winupgrade.Install(ctx, version, runtime.GOARCH, destination)
	if err != nil {
		return fmt.Errorf("upgrade failed: %w", err)
	}
	fmt.Fprintf(os.Stdout, "Mercury %s installed to %s.\n", version, destination)
	if backup != "" {
		fmt.Fprintf(os.Stdout, "Previous executable retained at %s; you can remove it after this command exits.\n", backup)
	}
	return nil
}
