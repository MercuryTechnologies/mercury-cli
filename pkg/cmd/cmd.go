// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/MercuryTechnologies/mercury-cli/internal/autocomplete"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
)

var (
	Command            *cli.Command
	CommandErrorBuffer bytes.Buffer
)

func init() {
	Command = &cli.Command{
		Name:      "mercury",
		Usage:     "CLI for the mercury API",
		Suggest:   true,
		Version:   Version,
		ErrWriter: &CommandErrorBuffer,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			// CUSTOM: confirmation prompt skip flag
			&cli.BoolFlag{
				Name:    "yes",
				Aliases: []string{"y"},
				Usage:   "Skip confirmation prompts",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
				Validator: func(baseURL string) error {
					return ValidateBaseURL(baseURL, "--base-url")
				},
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "The format for displaying response data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "format-error",
				Usage: "The format for displaying error data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "transform",
				Usage: "The GJSON transformation for data output.",
			},
			&cli.StringFlag{
				Name:  "transform-error",
				Usage: "The GJSON transformation for errors.",
			},
			&cli.BoolFlag{
				Name:    "raw-output",
				Aliases: []string{"r"},
				Usage:   "If the result is a string, print it without JSON quotes. This can be useful for making output transforms talk to non-JSON-based systems.",
			},
			&requestflag.Flag[string]{
				Name:    "api-key",
				Usage:   "Bearer token authentication for Mercury API.\n\nUse your API token in the Authorization header:\n`Authorization: Bearer TOKEN`\n\nExample:\n`Authorization: Bearer secret-token:mercury_<TOKEN>`\n\nYour Mercury API token should include the 'secret-token:' prefix.\nTokens can be generated from your Mercury dashboard settings.\n",
				Sources: cli.EnvVars("MERCURY_API_KEY"),
			},
			&cli.StringFlag{
				Name:  "environment",
				Usage: "API environment: sandbox or production",
				Validator: func(environment string) error {
					return ValidateEnvironment(environment, "--environment")
				},
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "customers",
				Usage:    "Create, update, and manage customers",
				Category: "Resources",
				Suggest:  true,
				Commands: []*cli.Command{
					&customersCreate,
					&customersUpdate,
					&customersList,
					&customersDelete,
					&customersGet,
				},
			},
			{
				Name:     "invoices",
				Usage:    "Create, update, and manage invoices",
				Category: "Resources",
				Suggest:  true,
				Commands: []*cli.Command{
					&invoicesCreate,
					&invoicesUpdate,
					&invoicesList,
					&invoicesCancel,
					&invoicesDownload,
					&invoicesGet,
					{
						Name:    "attachments",
						Usage:   "Get accounts receivable attachment details",
						Suggest: true,
						Commands: []*cli.Command{
							&invoicesAttachmentsList,
							&invoicesAttachmentsGet,
						},
					},
				},
			},
			{
				Name:     "cards",
				Usage:    "List debit and credit cards for an account",
				Category: "Resources",
				Suggest:  true,
				Commands: []*cli.Command{
					&cardsList,
				},
			},
			{
				Name:     "categories",
				Usage:    "List expense categories",
				Category: "Resources",
				Suggest:  true,
				Commands: []*cli.Command{
					&categoriesList,
				},
			},
			{
				Name:     "credit",
				Usage:    "List credit accounts",
				Category: "Resources",
				Suggest:  true,
				Commands: []*cli.Command{
					&creditList,
				},
			},
			{
				Name:     "events",
				Usage:    "List and inspect API events",
				Category: "Resources",
				Suggest:  true,
				Commands: []*cli.Command{
					&eventsList,
					&eventsGet,
				},
			},
			{
				Name:     "org",
				Usage:    "View organization details",
				Category: "Resources",
				Suggest:  true,
				Commands: []*cli.Command{
					&orgGet,
				},
			},
			{
				Name:     "payments",
				Usage:    "Send money, request approvals, and transfer between accounts",
				Category: "Resources",
				Suggest:  true,
				Commands: []*cli.Command{
					&paymentsCreate,
					&paymentsList,
					&paymentsGet,
					&paymentsRequest,
					&paymentsTransfer,
				},
			},
			{
				Name:     "safes",
				Usage:    "List and download SAFE agreements",
				Category: "Resources",
				Suggest:  true,
				Commands: []*cli.Command{
					&safesList,
					&safesDownload,
					&safesGet,
				},
			},
			{
				Name:     "statements",
				Usage:    "Download account statements as PDF",
				Category: "Resources",
				Suggest:  true,
				Commands: []*cli.Command{
					&statementsDownload,
					{
						Name:    "accounts",
						Usage:   "List monthly account statements",
						Suggest: true,
						Commands: []*cli.Command{
							&statementsAccountsList,
						},
					},
					{
						Name:    "treasury",
						Usage:   "List treasury account statements",
						Suggest: true,
						Commands: []*cli.Command{
							&statementsTreasuryList,
						},
					},
				},
			},
			{
				Name:     "treasury",
				Usage:    "View treasury accounts, statements, and transactions",
				Category: "Resources",
				Suggest:  true,
				Commands: []*cli.Command{
					&treasuryList,
					&treasuryTransactions,
				},
			},
			{
				Name:     "users",
				Usage:    "List and view organization team members",
				Category: "Resources",
				Suggest:  true,
				Commands: []*cli.Command{
					&usersList,
					&usersGet,
				},
			},
			{
				Name:     "webhooks",
				Usage:    "Set up and manage webhook endpoints",
				Category: "Resources",
				Suggest:  true,
				Commands: []*cli.Command{
					&webhooksCreate,
					&webhooksUpdate,
					&webhooksList,
					&webhooksDelete,
					&webhooksGet,
					&webhooksVerify,
				},
			},
			{
				Name:     "accounts",
				Usage:    "View accounts, cards, and transactions",
				Category: "Resources",
				Suggest:  true,
				Commands: []*cli.Command{
					&accountsList,
					&accountsGet,
				},
			},
			{
				Name:     "recipients",
				Usage:    "Add, update, and manage payment recipients",
				Category: "Resources",
				Suggest:  true,
				Commands: []*cli.Command{
					&recipientsCreate,
					&recipientsUpdate,
					&recipientsList,
					&recipientsGet,
					{
						Name:    "attachments",
						Usage:   "List and upload recipient tax form attachments",
						Suggest: true,
						Commands: []*cli.Command{
							&recipientsAttachmentsList,
							&recipientsAttachmentsAttach,
						},
					},
				},
			},
			{
				Name:     "transactions",
				Usage:    "Search, update, and attach files to transactions",
				Category: "Resources",
				Suggest:  true,
				Commands: []*cli.Command{
					&transactionsUpdate,
					&transactionsList,
					&transactionsGet,
					{
						Name:    "attachments",
						Usage:   "Upload file attachments to transactions",
						Suggest: true,
						Commands: []*cli.Command{
							&transactionsAttachmentsAttach,
						},
					},
				},
			},
			// CUSTOM: OAuth auth commands
			&authLogin,
			&authLogout,
			&authStatus,
			// CUSTOM: self-upgrade command
			&upgrade,
			{
				Name:     "onboarding",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&onboardingSubmit1,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "mercury @manpages [-o mercury.1] [--gzip]",
				Hidden:          true,
				Action:          generateManpages,
				HideHelpCommand: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "write manpages to the given folder",
						Value:   "man",
					},
					&cli.BoolFlag{
						Name:    "gzip",
						Aliases: []string{"z"},
						Usage:   "output gzipped manpage files to .gz",
						Value:   true,
					},
					&cli.BoolFlag{
						Name:    "text",
						Aliases: []string{"z"},
						Usage:   "output uncompressed text files",
						Value:   false,
					},
				},
			},
			{
				Name:            "__complete",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.ExecuteShellCompletion,
			},
			{
				Name:            "@completion",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.OutputCompletionScript,
			},
			{
				Name:     "hat",
				Usage:    "The most important command",
				Category: "Resources",
				Action:   openHat,
			},
		},
		HideHelpCommand: true,
	}
}

func generateManpages(ctx context.Context, c *cli.Command) error {
	manpage, err := docs.ToManWithSection(Command, 1)
	if err != nil {
		return err
	}
	dir := c.String("output")
	err = os.MkdirAll(filepath.Join(dir, "man1"), 0755)
	if err != nil {
		// handle error
	}
	if c.Bool("text") {
		file, err := os.Create(filepath.Join(dir, "man1", "mercury.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "mercury.1.gz"))
		if err != nil {
			return err
		}
		defer file.Close()
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		_, err = gzWriter.Write([]byte(manpage))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote manpages to %s\n", dir)
	return nil
}
