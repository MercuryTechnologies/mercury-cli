package cmd

import (
	"context"
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/urfave/cli/v3"
)

const hatURL = "https://talk-to-my-agent.myshopify.com/products/talk-to-my-agent-baseball-cap?variant=52044190843163"

var (
	hatBlue = lipgloss.NewStyle().Foreground(lipgloss.Color("#395AFF"))
	hatDim  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	hatBold = lipgloss.NewStyle().Foreground(lipgloss.Color("252")).Bold(true)
)

func openHat(ctx context.Context, c *cli.Command) error {
	fmt.Println()

	hatArt := []string{
		`        ████████`,
		`     ██         ██`,
		`    ██           ████████`,
		`    █████████████████████`,
	}
	for _, line := range hatArt {
		fmt.Println(hatBlue.Render("    " + line))
	}

	fmt.Println()
	fmt.Println(hatDim.Render("    Dear Mercury CLI user,"))
	fmt.Println(hatDim.Render("    Your dedication to modern banking has not gone unnoticed."))
	fmt.Println()
	fmt.Println(hatBold.Render("    Please accept this token of our appreciation."))

	time.Sleep(3 * time.Second)

	fmt.Println()
	fmt.Println(hatDim.Render("    Opening talk-to-my-agent.myshopify.com..."))
	fmt.Println()

	time.Sleep(2 * time.Second)

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
	return nil
}
