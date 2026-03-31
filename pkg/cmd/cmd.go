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
				Name:    "username",
				Hidden:  true,
				Usage:   "Basic authentication for Mercury API.\n\nUse your API token as the username with an empty password.\n\nExample:\nUsername: `secret-token:mercury_production_wma_24SCp4G81X3yHL4Wq8FgzuaP9ye3VKf2mgTDctXyRg5HY_yrucrem`\nPassword: (empty)\n",
				Sources: cli.EnvVars("MERCURY_USERNAME"),
			},
			&requestflag.Flag[string]{
				Name:    "password",
				Hidden:  true,
				Usage:   "Basic authentication for Mercury API.\n\nUse your API token as the username with an empty password.\n\nExample:\nUsername: `secret-token:mercury_production_wma_24SCp4G81X3yHL4Wq8FgzuaP9ye3VKf2mgTDctXyRg5HY_yrucrem`\nPassword: (empty)\n",
				Sources: cli.EnvVars("MERCURY_PASSWORD"),
			},
			&requestflag.Flag[string]{
				Name:    "api-key",
				Usage:   "Bearer token authentication for Mercury API.\n\nUse your API token in the Authorization header:\n`Authorization: Bearer TOKEN`\n\nExample:\n`Authorization: Bearer secret-token:mercury_production_wma_24SCp4G81X3yHL4Wq8FgzuaP9ye3VKf2mgTDctXyRg5HY_yrucrem`\n\nYour Mercury API token should include the 'secret-token:' prefix.\nTokens can be generated from your Mercury dashboard settings.\n",
				Sources: cli.EnvVars("MERCURY_API_KEY"),
			},
			&cli.StringFlag{
				Name:  "environment",
				Usage: "API environment: sandbox or production",
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "accounts-receivable:attachements",
        Usage:    "Get accounts receivable attachment details",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&accountsReceivableAttachementsRetrieve,
				},
			},
			{
				Name:     "accounts-receivable:customers",
				Usage:    "Create, update, and delete accounts receivable customers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&accountsReceivableCustomersCreate,
					&accountsReceivableCustomersRetrieve,
					&accountsReceivableCustomersUpdate,
					&accountsReceivableCustomersList,
					&accountsReceivableCustomersDelete,
				},
			},
			{
				Name:     "accounts-receivable:invoices",
				Usage:    "Create, update, and download accounts receivable invoices",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&accountsReceivableInvoicesCreate,
					&accountsReceivableInvoicesRetrieve,
					&accountsReceivableInvoicesUpdate,
					&accountsReceivableInvoicesList,
					&accountsReceivableInvoicesCancel,
					&accountsReceivableInvoicesDownloadPdf,
					&accountsReceivableInvoicesListAttachments,
				},
			},
			{
				Name:     "categories",
				Usage:    "List expense categories",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&categoriesList,
				},
			},
			{
				Name:     "credit",
				Usage:    "List credit accounts",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&creditList,
				},
			},
			{
				Name:     "events",
				Usage:    "List and inspect API events",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&eventsRetrieve,
					&eventsList,
				},
			},
			{
				Name:     "organization",
				Usage:    "View organization details",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&organizationRetrieve,
				},
			},
			{
				Name:     "request-send-money",
				Usage:    "View send money approval requests",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&requestSendMoneyRetrieve,
				},
			},
			{
				Name:     "safes",
				Usage:    "List and download SAFE agreements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&safesRetrieve,
					&safesList,
					&safesDownloadDocument,
				},
			},
			{
				Name:     "statements",
				Usage:    "Download account statements as PDF",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&statementsDownloadPdf,
				},
			},
			{
				Name:     "transfer",
				Usage:    "Initiate transfers between accounts",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&transferCreate,
				},
			},
			{
				Name:     "treasury",
				Usage:    "View treasury accounts, statements, and transactions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&treasuryList,
					&treasuryRetrieveStatements,
					&treasuryRetrieveTransactions,
				},
			},
			{
				Name:     "users",
				Usage:    "List and view organization team members",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&usersRetrieve,
					&usersList,
				},
			},
			{
				Name:     "webhooks",
				Usage:    "Set up and manage webhook endpoints",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&webhooksCreate,
					&webhooksRetrieve,
					&webhooksUpdate,
					&webhooksList,
					&webhooksDelete,
					&webhooksVerify,
				},
			},
			{
				Name:     "accounts",
        Usage:    "View accounts, cards, and transactions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&accountsRetrieve,
					&accountsList,
					&accountsCreateTransaction,
					&accountsListCards,
					&accountsListStatements,
					&accountsListTransactions,
					&accountsRequestSendMoney,
					&accountsRetrieveTransaction,
				},
			},
			{
				Name:     "recipients",
				Usage:    "Add, update, and manage payment recipients",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&recipientsCreate,
					&recipientsRetrieve,
					&recipientsUpdate,
					&recipientsList,
					&recipientsListAttachments,
					&recipientsUploadAttachment,
				},
			},
			{
				Name:     "transactions",
				Usage:    "Search, update, and attach files to transactions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&transactionsRetrieve,
					&transactionsUpdate,
					&transactionsList,
					&transactionsUploadAttachment,
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
