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
	"context"
	"fmt"
	"os"
	"reflect"

	"github.com/MercuryTechnologies/mercury-cli/internal/apiquery"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
	mercury "github.com/MercuryTechnologies/mercury-go"
	"github.com/MercuryTechnologies/mercury-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

func ptr[T any](v T) *T { return &v }

// setFlagUsage overwrites the Usage text of a named flag on a command. It uses
// reflection so one helper works across all generic `requestflag.Flag[T]`
// instantiations without a per-type switch. Silently no-ops if the flag is
// missing (e.g. renamed by a Stainless update) — run `go build ./...` after
// codegen to catch any drift via the surrounding compile-checked references.
func setFlagUsage(cmd *cli.Command, name, usage string) {
	for _, f := range cmd.Flags {
		names := f.Names()
		if len(names) == 0 || names[0] != name {
			continue
		}
		v := reflect.ValueOf(f)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		u := v.FieldByName("Usage")
		if u.IsValid() && u.CanSet() && u.Kind() == reflect.String {
			u.SetString(usage)
		}
		return
	}
}

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
		case "payments":
			sub.Usage = "Send money, request approvals, and transfer between accounts"
		case "treasury":
			sub.Usage = "View treasury accounts and transactions"
		case "onboarding":
			sub.Usage = "Use the CLI to help you get started on creating a Mercury account"
			sub.Category = "Create Account"
		}
	}

	// ── Global flag overrides ────────────────────────────────────────────────
	for _, f := range Command.Flags {
		switch f.Names()[0] {
		case "api-key":
			if rf, ok := f.(*requestflag.Flag[string]); ok {
				rf.Usage = "Mercury API token (from dashboard settings)"
			}
		case "format":
			if sf, ok := f.(*cli.StringFlag); ok {
				sf.Usage = "Output format (auto|json|jsonl|pretty|raw|yaml|explore)"
			}
		case "format-error":
			if sf, ok := f.(*cli.StringFlag); ok {
				sf.Usage = "Error format (auto|json|jsonl|pretty|raw|yaml|explore)"
			}
		case "transform":
			if sf, ok := f.(*cli.StringFlag); ok {
				sf.Usage = "Transform output with a GJSON expression"
			}
		case "transform-error":
			if sf, ok := f.(*cli.StringFlag); ok {
				sf.Usage = "Transform error output with a GJSON expression"
			}
		case "raw-output":
			if bf, ok := f.(*cli.BoolFlag); ok {
				bf.Usage = "If the result is a string, print it without JSON quotes."
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

	// payments create flags
	setFlagUsage(&paymentsCreate, "account-id", "Account to send from")
	setFlagUsage(&paymentsCreate, "recipient-id", "Recipient to pay (use 'recipients list' to find)")
	setFlagUsage(&paymentsCreate, "amount", "Amount in dollars, e.g. 5000.00")
	setFlagUsage(&paymentsCreate, "payment-method", "One of: ach, check, domesticWire, internationalWire")
	setFlagUsage(&paymentsCreate, "idempotency-key", "Unique key to prevent duplicate payments")
	setFlagUsage(&paymentsCreate, "external-memo", "Memo visible to the recipient")
	setFlagUsage(&paymentsCreate, "note", "Internal note (visible to your team only)")
	setFlagUsage(&paymentsCreate, "purpose", "Wire transfer purpose (required for domesticWire)")

	// payments request flags
	setFlagUsage(&paymentsRequest, "account-id", "Account to send from")
	setFlagUsage(&paymentsRequest, "recipient-id", "Recipient to pay (use 'recipients list' to find)")
	setFlagUsage(&paymentsRequest, "amount", "Amount in dollars, e.g. 5000.00")
	setFlagUsage(&paymentsRequest, "idempotency-key", "Unique key to prevent duplicate payments")
	setFlagUsage(&paymentsRequest, "external-memo", "Memo visible to the recipient")
	setFlagUsage(&paymentsRequest, "note", "Internal note (visible to your team only)")

	// payments transfer flags
	setFlagUsage(&paymentsTransfer, "source-account-id", "Account to transfer from")
	setFlagUsage(&paymentsTransfer, "destination-account-id", "Account to transfer to")
	setFlagUsage(&paymentsTransfer, "amount", "Amount in dollars, e.g. 5000.00")
	setFlagUsage(&paymentsTransfer, "idempotency-key", "Unique key to prevent duplicate transfers")
	setFlagUsage(&paymentsTransfer, "note", "Internal note")

	// payments list flags
	setFlagUsage(&paymentsList, "account-id", "Filter by account")
	setFlagUsage(&paymentsList, "end-before", "Paginate backwards from this request ID")
	setFlagUsage(&paymentsList, "start-after", "Paginate forwards from this request ID")
	setFlagUsage(&paymentsList, "status", "Filter by status (pendingApproval, approved, rejected, cancelled)")

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

	// onboarding
	for _, f := range onboardingSubmit.Flags {
		switch rf := f.(type) {
		case *requestflag.Flag[string]:
			if rf.Name == "partner" {
				rf.Hidden = true
				rf.Required = false
				rf.Const = true
				rf.Default = "Mercury Integrations"
			}
		case *requestflag.Flag[*string]:
			switch rf.Name {
			case "application-type":
				rf.Hidden = true
				rf.Const = true
				rf.Default = ptr("DefaultApplication")
			case "webhook-url":
				rf.Hidden = true
			}
		}
	}

	// onboarding submit flags
	setFlagUsage(&onboardingSubmit, "beneficial-owner", "Beneficial owner details (name, address, ID, ownership %)")
	setFlagUsage(&onboardingSubmit, "about", "Company info (name, industry, description, website)")
	setFlagUsage(&onboardingSubmit, "business-contact-details", "Business contact address and phone number")
	setFlagUsage(&onboardingSubmit, "business-legal-address", "Registered legal address of the business")
	setFlagUsage(&onboardingSubmit, "business-physical-address", "Physical/mailing address of the business")
	setFlagUsage(&onboardingSubmit, "formation-details", "EIN, company structure, and formation documents")
	setFlagUsage(&onboardingSubmit, "invite-email", "Email address to send the application invite to")

	// Override onboarding submit to display the signup link prominently.
	onboardingSubmit.Action = onboardingSubmitOverride

	// webhooks
	webhooksList.Usage = "List webhook endpoints (supports pagination)"
	webhooksGet.Usage = "Get a webhook endpoint by ID"
	webhooksVerify.Usage = "Send a test event to verify a webhook endpoint"
	webhooksUpdate.Usage = "Update a webhook endpoint's configuration"
}

func onboardingSubmitOverride(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	if len(cmd.Args().Slice()) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", cmd.Args().Slice())
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := mercury.OnboardingSubmitParams{}
	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Onboarding.Submit(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	if link := obj.Get("signupLink"); link.Exists() {
		fmt.Fprintf(os.Stdout, "\n  Open this link to continue your Mercury application:\n\n")
		fmt.Fprintf(os.Stdout, "  \033[1;4m%s\033[0m\n\n", link.String())
		return nil
	}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "onboarding submit",
		Transform:      transform,
	})
}
