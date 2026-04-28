package cmd

import (
	"bufio"
	"context"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/urfave/cli/v3"
)

const hatCartURL = "https://talktomyagent.shop/cart/52056132124955:1"

var (
	hatBlue = lipgloss.NewStyle().Foreground(colorBlue)
	hatDim  = lipgloss.NewStyle().Foreground(colorDim)
	hatBold = lipgloss.NewStyle().Foreground(colorLight).Bold(true)
)

func promptField(scanner *bufio.Scanner, label string, required bool) (string, error) {
	for {
		fmt.Print(hatDim.Render("  "+label) + " ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				return "", fmt.Errorf("failed to read input: %w", err)
			}
			return "", fmt.Errorf("input cancelled")
		}
		value := strings.TrimSpace(scanner.Text())
		if !required || value != "" {
			return value, nil
		}
		fmt.Println(hatDim.Render("  This field is required. Please try again."))
	}
}

func promptShippingAddress() (*shippingAddress, error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println()
	fmt.Println(hatDim.Render("  Where should we ship your hat?"))
	fmt.Println()

	email, err := promptField(scanner, "Email:", true)
	if err != nil {
		return nil, err
	}

	firstName, err := promptField(scanner, "First name:", true)
	if err != nil {
		return nil, err
	}

	lastName, err := promptField(scanner, "Last name:", true)
	if err != nil {
		return nil, err
	}

	address1, err := promptField(scanner, "Address line 1:", true)
	if err != nil {
		return nil, err
	}

	address2, err := promptField(scanner, "Address line 2 (optional):", false)
	if err != nil {
		return nil, err
	}

	city, err := promptField(scanner, "City:", true)
	if err != nil {
		return nil, err
	}

	state, err := promptField(scanner, "State:", true)
	if err != nil {
		return nil, err
	}

	zip, err := promptField(scanner, "ZIP code:", true)
	if err != nil {
		return nil, err
	}

	return &shippingAddress{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Address1:  address1,
		Address2:  address2,
		City:      city,
		State:     state,
		Zip:       zip,
	}, nil
}

type shippingAddress struct {
	Email     string
	FirstName string
	LastName  string
	Address1  string
	Address2  string
	City      string
	State     string
	Zip       string
}

func buildShopifyURL(addr *shippingAddress) string {
	params := url.Values{}
	params.Set("checkout[email]", addr.Email)
	params.Set("checkout[shipping_address][first_name]", addr.FirstName)
	params.Set("checkout[shipping_address][last_name]", addr.LastName)
	params.Set("checkout[shipping_address][address1]", addr.Address1)
	if addr.Address2 != "" {
		params.Set("checkout[shipping_address][address2]", addr.Address2)
	}
	params.Set("checkout[shipping_address][city]", addr.City)
	params.Set("checkout[shipping_address][province]", addr.State)
	params.Set("checkout[shipping_address][zip]", addr.Zip)
	params.Set("checkout[shipping_address][country]", "US")
	params.Set("utm_source", "mercury-cli")
	params.Set("utm_campaign", "agent-hat")

	return hatCartURL + "?" + params.Encode()
}

func openHat(ctx context.Context, c *cli.Command) error {
	fmt.Println()

	hatArt := []string{
		`        ████████`,
		`     ██         ██`,
		`    ██           ████████`,
		`    █████████████████████`,
	}
	for _, line := range hatArt {
		fmt.Println(hatBlue.Render("  " + line))
	}

	fmt.Println()
	fmt.Println(hatDim.Render("  Dear Mercury CLI user,"))
	fmt.Println(hatDim.Render("  Your dedication to modern banking has not gone unnoticed."))
	fmt.Println()
	fmt.Println(hatBold.Render("  Please accept this token of our appreciation."))

	time.Sleep(3 * time.Second)

	addr, err := promptShippingAddress()
	if err != nil {
		fmt.Println()
		fmt.Println(hatDim.Render("  No worries! Come back anytime."))
		fmt.Println()
		return nil
	}

	shopifyURL := buildShopifyURL(addr)

	fmt.Println()
	fmt.Println(hatDim.Render("  Opening talktomyagent.shop..."))
	fmt.Println()

	time.Sleep(2 * time.Second)

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", shopifyURL)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", shopifyURL)
	default:
		cmd = exec.Command("xdg-open", shopifyURL)
	}
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to open browser: %w", err)
	}
	return nil
}
