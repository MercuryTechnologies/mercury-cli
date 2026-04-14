package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
	"github.com/urfave/cli/v3"
)

// ErrCancelled is returned when the user cancels a confirmation prompt.
var ErrCancelled = errors.New("operation cancelled")

// ConfirmDetail represents a single row in the confirmation summary.
type ConfirmDetail struct {
	Label string
	Value string
}

// formatCurrency formats a float64 as $1,234.56.
func formatCurrency(amount float64) string {
	negative := amount < 0
	if negative {
		amount = -amount
	}

	s := fmt.Sprintf("%.2f", amount)
	parts := strings.SplitN(s, ".", 2)
	intPart := parts[0]
	decPart := parts[1]

	// Insert commas every 3 digits from the right
	if len(intPart) > 3 {
		var b strings.Builder
		remainder := len(intPart) % 3
		if remainder > 0 {
			b.WriteString(intPart[:remainder])
		}
		for i := remainder; i < len(intPart); i += 3 {
			if b.Len() > 0 {
				b.WriteByte(',')
			}
			b.WriteString(intPart[i : i+3])
		}
		intPart = b.String()
	}

	if negative {
		return "-$" + intPart + "." + decPart
	}
	return "$" + intPart + "." + decPart
}

// confirmMoneyMovement prompts the user to confirm a money movement command.
// If --yes is set or stdin is not a terminal, it either skips or errors accordingly.
func confirmMoneyMovement(cmd *cli.Command, action string, details []ConfirmDetail) error {
	if cmd.Bool("yes") {
		return nil
	}

	if !term.IsTerminal(os.Stdin.Fd()) {
		return fmt.Errorf("refusing to execute without confirmation in non-interactive mode\nUse --yes to skip confirmation prompts")
	}

	return confirmMoneyMovementIO(os.Stdin, os.Stderr, action, details)
}

// confirmMoneyMovementIO is the testable core of the confirmation prompt.
// It renders a styled summary to writer and reads Y/n from reader.
func confirmMoneyMovementIO(reader io.Reader, writer io.Writer, action string, details []ConfirmDetail) error {
	titleStyle := lipgloss.NewStyle().Foreground(colorBlue).Bold(true)
	labelStyle := lipgloss.NewStyle().Foreground(colorDim)
	valueStyle := lipgloss.NewStyle().Foreground(colorLight)
	ruleStyle := lipgloss.NewStyle().Foreground(colorFrame)
	promptStyle := lipgloss.NewStyle().Foreground(colorLight).Bold(true)

	// Find max label width for alignment
	maxLabel := 0
	for _, d := range details {
		if len(d.Label) > maxLabel {
			maxLabel = len(d.Label)
		}
	}

	fmt.Fprintln(writer)
	fmt.Fprintln(writer, ruleStyle.Render("  ──── ")+titleStyle.Render(action)+ruleStyle.Render(" ────"))
	fmt.Fprintln(writer)

	for _, d := range details {
		padding := strings.Repeat(" ", maxLabel-len(d.Label)+2)
		fmt.Fprintln(writer, "  "+labelStyle.Render("  "+d.Label)+padding+valueStyle.Render(d.Value))
	}

	fmt.Fprintln(writer)
	fmt.Fprint(writer, "  "+promptStyle.Render("Proceed?")+labelStyle.Render(" [Y/n] "))

	scanner := bufio.NewScanner(reader)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return fmt.Errorf("failed to read input: %w", err)
		}
		return ErrCancelled
	}

	answer := strings.TrimSpace(strings.ToLower(scanner.Text()))
	switch answer {
	case "", "y", "yes":
		return nil
	default:
		return ErrCancelled
	}
}

// buildPaymentConfirmDetails extracts confirmation details for payments create and request.
func buildPaymentConfirmDetails(cmd *cli.Command) []ConfirmDetail {
	details := []ConfirmDetail{
		{Label: "Account", Value: cmd.Value("account-id").(string)},
		{Label: "Recipient", Value: cmd.Value("recipient-id").(string)},
		{Label: "Amount", Value: formatCurrency(cmd.Value("amount").(float64))},
		{Label: "Payment Method", Value: cmd.Value("payment-method").(string)},
	}

	if cmd.IsSet("external-memo") {
		details = append(details, ConfirmDetail{Label: "External Memo", Value: cmd.Value("external-memo").(string)})
	}
	if cmd.IsSet("note") {
		details = append(details, ConfirmDetail{Label: "Note", Value: cmd.Value("note").(string)})
	}

	return details
}

// buildTransferConfirmDetails extracts confirmation details for payments transfer.
func buildTransferConfirmDetails(cmd *cli.Command) []ConfirmDetail {
	details := []ConfirmDetail{
		{Label: "From", Value: cmd.Value("source-account-id").(string)},
		{Label: "To", Value: cmd.Value("destination-account-id").(string)},
		{Label: "Amount", Value: formatCurrency(cmd.Value("amount").(float64))},
	}

	if cmd.IsSet("note") {
		details = append(details, ConfirmDetail{Label: "Note", Value: fmt.Sprintf("%v", cmd.Value("note"))})
	}

	return details
}
