package cmd

// overrides.go — Human-authored copy overrides for Stainless-generated commands.
//
// Stainless regenerates all other files in this package from the OpenAPI spec.
// This file is NOT generated and will survive codegen runs. It patches Usage
// strings (and global flag descriptions) to be concise, consistent, and free
// of HTTP implementation details.
//
// After a Stainless update, run `go build ./...` to confirm nothing broke.
// If Stainless renames a variable, the compiler will tell you here.

import (
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
	"github.com/urfave/cli/v3"
)

func init() {
	// ── Top-level resource descriptions (shown in `mercury --help` box) ──────
	//
	// These match by string name because the resource groups are declared inline
	// in cmd.go, not as named package-level vars. If Stainless renames a resource,
	// the match silently falls back to the generated description (no compiler
	// error). This is the same fragility as the `// CUSTOM:` approach used
	// elsewhere in cmd.go — just worth keeping in mind after codegen updates.
	for _, sub := range Command.Commands {
		switch sub.Name {
		case "accounts":
			sub.Usage = "List and view bank accounts"
		case "treasury":
			sub.Usage = "View treasury accounts and transactions"
		}
	}

	// ── Global flag overrides ────────────────────────────────────────────────
	for _, f := range Command.Flags {
		switch f.Names()[0] {
		case "api-key":
			if rf, ok := f.(*requestflag.Flag[string]); ok {
				rf.Usage = "Mercury API token (generate at dashboard.mercury.com/settings)"
			}
		case "transform":
			if sf, ok := f.(*cli.StringFlag); ok {
				sf.Usage = "Transform output using a GJSON expression"
			}
		case "transform-error":
			if sf, ok := f.(*cli.StringFlag); ok {
				sf.Usage = "Transform error output using a GJSON expression"
			}
		}
	}

	// ── Subcommand description overrides ─────────────────────────────────────
	//
	// Principles applied:
	//   - Keep first line under ~60 chars (avoids wrapping in the help box)
	//   - Imperative voice: "List", "Get", "Create" (not "Retrieve a paginated...")
	//   - No HTTP internals (Content-Disposition, multipart/form-data)
	//   - "(supports pagination)" as standard suffix for list commands
	//   - Keep genuinely useful warnings ("This action cannot be undone.")

	// accounts
	accountsList.Usage = "List all accounts in your organization (supports pagination)"
	accountsGet.Usage = "Get an account by ID"

	// cards
	cardsList.Usage = "List debit and credit cards for an account"

	// categories
	categoriesList.Usage = "List expense categories (supports pagination)"

	// credit
	creditList.Usage = "List credit accounts for your organization"

	// customers
	customersList.Usage = "List customers (supports pagination)"
	customersGet.Usage = "Get a customer by ID"

	// events
	eventsList.Usage = "List API events (supports pagination)"
	eventsGet.Usage = "Get an API event by ID"

	// invoices
	invoicesList.Usage = "List invoices (supports pagination)"
	invoicesDownload.Usage = "Download an invoice as PDF"
	invoicesGet.Usage = "Get an invoice by ID"

	// invoice attachments
	invoicesAttachmentsList.Usage = "List attachments for an invoice"
	invoicesAttachmentsGet.Usage = "Get attachment details including download URL"

	// org
	orgGet.Usage = "View your organization details"

	// payments
	paymentsCreate.Usage = "Send money from an account to a recipient"
	paymentsList.Usage = "List payment approval requests (supports pagination)"
	paymentsGet.Usage = "Get a payment approval request by ID"
	paymentsRequest.Usage = "Request approval to send money"
	paymentsTransfer.Usage = "Transfer funds between accounts in your organization"

	// recipients
	recipientsList.Usage = "List payment recipients (supports pagination)"
	recipientsGet.Usage = "Get a recipient by ID"

	// recipient attachments
	recipientsAttachmentsList.Usage = "List recipient tax form attachments (supports pagination)"
	recipientsAttachmentsAttach.Usage = "Upload a tax form attachment for a recipient"

	// safes
	safesList.Usage = "List SAFE agreements for your organization"
	safesDownload.Usage = "Download a SAFE agreement as PDF"
	safesGet.Usage = "Get a SAFE agreement by ID"

	// statements
	statementsDownload.Usage = "Download an account statement as PDF"
	statementsAccountsList.Usage = "List monthly statements for an account (supports pagination)"
	statementsTreasuryList.Usage = "List statements for a treasury account (supports pagination)"

	// transactions
	transactionsUpdate.Usage = "Update a transaction's note or category"
	transactionsList.Usage = "List transactions across all accounts (supports pagination)"
	transactionsGet.Usage = "Get a transaction by ID"

	// transaction attachments
	transactionsAttachmentsAttach.Usage = "Upload a file attachment to a transaction"

	// treasury
	treasuryList.Usage = "List treasury accounts (supports pagination)"
	treasuryTransactions.Usage = "List transactions for a treasury account (supports pagination)"

	// users
	usersList.Usage = "List organization team members (supports pagination)"
	usersGet.Usage = "Get a team member by ID"

	// webhooks
	webhooksList.Usage = "List webhook endpoints (supports pagination)"
	webhooksGet.Usage = "Get a webhook endpoint by ID"
	webhooksVerify.Usage = "Send a test event to verify a webhook endpoint"
	webhooksUpdate.Usage = "Update a webhook endpoint's configuration"
}
