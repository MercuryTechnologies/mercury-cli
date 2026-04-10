package cmd

import (
	"context"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/urfave/cli/v3"
)

const hatURL = "https://talk-to-my-agent.myshopify.com/products/talk-to-my-agent-baseball-cap?variant=52044190843163"

func openHat(ctx context.Context, c *cli.Command) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", hatURL)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", hatURL)
	default:
		cmd = exec.Command("xdg-open", hatURL)
	}
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to open browser: %w", err)
	}
	fmt.Println("Opening store in your browser...")
	return nil
}
