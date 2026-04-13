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
			&requestflag.Flag[string]{
				Name:    "api-key",
				Usage:   "Bearer token authentication for Mercury API.\n\nUse your API token in the Authorization header:\n`Authorization: Bearer TOKEN`\n\nExample:\n`Authorization: Bearer secret-token:mercury_<TOKEN>`\n\nYour Mercury API token should include the 'secret-token:' prefix.\nTokens can be generated from your Mercury dashboard settings.\n",
				Sources: cli.EnvVars("MERCURY_API_KEY"),
			},
			&cli.StringFlag{
				Name:  "environment",
				Usage: "Set the environment for API requests",
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "customers",
				Category: "API RESOURCE",
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
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&invoicesCreate,
					&invoicesUpdate,
					&invoicesList,
					&invoicesCancel,
					&invoicesDownload,
					&invoicesGet,
					&invoicesListAttachments,
				},
			},
			{
				Name:     "accounts-receivable:attachments",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&accountsReceivableAttachmentsGet,
				},
			},
			{
				Name:     "cards",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&cardsList,
				},
			},
			{
				Name:     "categories",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&categoriesList,
				},
			},
			{
				Name:     "credit",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&creditList,
				},
			},
			{
				Name:     "events",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&eventsList,
					&eventsGet,
				},
			},
			{
				Name:     "org",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&orgGet,
				},
			},
			{
				Name:     "payments",
				Category: "API RESOURCE",
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
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&safesList,
					&safesDownload,
					&safesGet,
				},
			},
			{
				Name:     "statements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&statementsDownload,
				},
			},
			{
				Name:     "statements:accounts",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&statementsAccountsList,
				},
			},
			{
				Name:     "statements:treasury",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&statementsTreasuryList,
				},
			},
			{
				Name:     "treasury",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&treasuryList,
					&treasuryRetrieveTransactions,
				},
			},
			{
				Name:     "users",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&usersList,
					&usersGet,
				},
			},
			{
				Name:     "webhooks",
				Category: "API RESOURCE",
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
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&accountsList,
					&accountsGet,
				},
			},
			{
				Name:     "recipients",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&recipientsCreate,
					&recipientsUpdate,
					&recipientsList,
					&recipientsAttach,
					&recipientsGet,
					&recipientsListAttachments,
				},
			},
			{
				Name:     "transactions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&transactionsUpdate,
					&transactionsList,
					&transactionsAttach,
					&transactionsGet,
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
