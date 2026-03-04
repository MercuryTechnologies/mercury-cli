// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/MercuryTechnologies/mercury-cli/internal/apiquery"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
	"github.com/stainless-sdks/mercury-go"
	"github.com/stainless-sdks/mercury-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var accountsRecievableInvoicesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a new invoice for the organization",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[bool]{
			Name:     "ach-debit-enabled",
			Usage:    " Whether or not the invoice can be paid via ACH debit.",
			Required: true,
			BodyPath: "achDebitEnabled",
		},
		&requestflag.Flag[[]string]{
			Name:     "cc-email",
			Usage:    " Emails to be CCed on invoice notifications/reminders.",
			Required: true,
			BodyPath: "ccEmails",
		},
		&requestflag.Flag[bool]{
			Name:     "credit-card-enabled",
			Usage:    " Whether or not the invoice can be paid via credit card. Requires Stripe to be setup for the Mercury account.",
			Required: true,
			BodyPath: "creditCardEnabled",
		},
		&requestflag.Flag[string]{
			Name:     "customer-id",
			Usage:    "The customer who will receive the invoice. Use the /api/v1/ar/customers endpoint to list your customers and find the corresponding id, or create a new customer first.",
			Required: true,
			BodyPath: "customerId",
		},
		&requestflag.Flag[string]{
			Name:     "destination-account-id",
			Usage:    "ID for a Mercury account.",
			Required: true,
			BodyPath: "destinationAccountId",
		},
		&requestflag.Flag[any]{
			Name:     "due-date",
			Usage:    " The due date the invoice should be paid by. YYYY-MM-DD",
			Required: true,
			BodyPath: "dueDate",
		},
		&requestflag.Flag[any]{
			Name:     "invoice-date",
			Usage:    " The date of the invoice, set by the invoice creator and likely to be context specific to the type of transaction. For example, it could be a date a service was performed. YYYY-MM-DD",
			Required: true,
			BodyPath: "invoiceDate",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "line-item",
			Usage:    " The line items for the invoice",
			Required: true,
			BodyPath: "lineItems",
		},
		&requestflag.Flag[bool]{
			Name:     "use-real-account-number",
			Usage:    " Whether or not the invoice payment instructions will show the real account and routing number for the destination account or use virtual account numbers instead. Virtual accounts are safer and are preferred in most cases.",
			Required: true,
			BodyPath: "useRealAccountNumber",
		},
		&requestflag.Flag[any]{
			Name:     "internal-note",
			Usage:    " Internal note for the invoice, visible by users in the organization but not visible to payers.",
			BodyPath: "internalNote",
		},
		&requestflag.Flag[any]{
			Name:     "invoice-number",
			Usage:    " The payer facing invoice number/identifier.",
			BodyPath: "invoiceNumber",
		},
		&requestflag.Flag[any]{
			Name:     "payer-memo",
			Usage:    " Memo for the payer of the invoice.",
			BodyPath: "payerMemo",
		},
		&requestflag.Flag[any]{
			Name:     "po-number",
			Usage:    " Purchase order number for the invoice, if applicable.",
			BodyPath: "poNumber",
		},
		&requestflag.Flag[any]{
			Name:     "send-email-option",
			Usage:    ` Rules for emailing the new invoice to payers. Can be "DontSend" to skip sending or "SendNow" to send immediately. If omitted, defaults to sending immediately.`,
			BodyPath: "sendEmailOption",
		},
		&requestflag.Flag[any]{
			Name:     "service-period-end-date",
			Usage:    " The end date for the service period this invoice covers, if applicable. YYYY-MM-DD",
			BodyPath: "servicePeriodEndDate",
		},
		&requestflag.Flag[any]{
			Name:     "service-period-start-date",
			Usage:    " The start date for the service period this invoice covers, if applicable. YYYY-MM-DD",
			BodyPath: "servicePeriodStartDate",
		},
	},
	Action:          handleAccountsRecievableInvoicesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"line-item": {
		&requestflag.InnerFlag[string]{
			Name:       "line-item.name",
			Usage:      " the name of the line item",
			InnerField: "name",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "line-item.quantity",
			Usage:      " the quantity of this item",
			InnerField: "quantity",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "line-item.unit-price",
			Usage:      "A dollar amount",
			InnerField: "unitPrice",
		},
		&requestflag.InnerFlag[any]{
			Name:       "line-item.sales-tax-rate",
			Usage:      " the sales tax applied to this item",
			InnerField: "salesTaxRate",
		},
	},
})

var accountsRecievableInvoicesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve details of an invoice by its ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "invoice-id",
			Usage:    "ID for the invoice.",
			Required: true,
		},
	},
	Action:          handleAccountsRecievableInvoicesRetrieve,
	HideHelpCommand: true,
}

var accountsRecievableInvoicesUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update an existing invoice",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "invoice-id",
			Usage:    "ID for the invoice.",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:     "ach-debit-enabled",
			Usage:    " Whether or not the invoice can be paid via ACH debit.",
			Required: true,
			BodyPath: "achDebitEnabled",
		},
		&requestflag.Flag[[]string]{
			Name:     "cc-email",
			Usage:    " List of emails to be CCed on notifications/reminders.",
			Required: true,
			BodyPath: "ccEmails",
		},
		&requestflag.Flag[bool]{
			Name:     "credit-card-enabled",
			Usage:    " Whether or not the invoice can be paid via credit card. Requires Stripe to be setup for the Mercury account.",
			Required: true,
			BodyPath: "creditCardEnabled",
		},
		&requestflag.Flag[any]{
			Name:     "due-date",
			Usage:    " The date the invoice should be paid by. YYYY-MM-DD",
			Required: true,
			BodyPath: "dueDate",
		},
		&requestflag.Flag[any]{
			Name:     "invoice-date",
			Usage:    " The date of the invoice, set by the invoice creator. Does not have to be the day the invoice was created. It can be business specific i.e. service/sale date. YYYY-MM-DD",
			Required: true,
			BodyPath: "invoiceDate",
		},
		&requestflag.Flag[string]{
			Name:     "invoice-number",
			Usage:    " The invoice number.",
			Required: true,
			BodyPath: "invoiceNumber",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "line-item",
			Usage:    " The line items for the invoice",
			Required: true,
			BodyPath: "lineItems",
		},
		&requestflag.Flag[bool]{
			Name:     "use-real-account-number",
			Usage:    " Whether or not the invoice payment instructions will show the real account and routing number for the destination account or use virtual account numbers instead.",
			Required: true,
			BodyPath: "useRealAccountNumber",
		},
		&requestflag.Flag[any]{
			Name:     "internal-note",
			Usage:    " Internal note for the invoice, visible by users in the organization but not visible to payers.",
			BodyPath: "internalNote",
		},
		&requestflag.Flag[any]{
			Name:     "payer-memo",
			Usage:    " Memo for the payer of the invoice.",
			BodyPath: "payerMemo",
		},
		&requestflag.Flag[any]{
			Name:     "po-number",
			Usage:    " The purchase order number for the invoice if applicable.",
			BodyPath: "poNumber",
		},
		&requestflag.Flag[any]{
			Name:     "service-period-end-date",
			Usage:    " The end date for the service period this invoice covers, if applicable. YYYY-MM-DD",
			BodyPath: "servicePeriodEndDate",
		},
		&requestflag.Flag[any]{
			Name:     "service-period-start-date",
			Usage:    " The start date for the service period this invoice covers, if applicable. YYYY-MM-DD",
			BodyPath: "servicePeriodStartDate",
		},
	},
	Action:          handleAccountsRecievableInvoicesUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"line-item": {
		&requestflag.InnerFlag[string]{
			Name:       "line-item.name",
			Usage:      " the name of the line item",
			InnerField: "name",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "line-item.quantity",
			Usage:      " the quantity of this item",
			InnerField: "quantity",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "line-item.unit-price",
			Usage:      "A dollar amount",
			InnerField: "unitPrice",
		},
		&requestflag.InnerFlag[any]{
			Name:       "line-item.sales-tax-rate",
			Usage:      " the sales tax applied to this item",
			InnerField: "salesTaxRate",
		},
	},
})

var accountsRecievableInvoicesList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a paginated list of invoices. Supports cursor-based pagination with\nlimit, order, start_after, and end_before query parameters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "end-before",
			Usage:     "The ID of the invoice to end the page before (exclusive). When provided, results will end just before this ID and work backwards. Use this for reverse pagination or to retrieve previous pages. Cannot be combined with start_after.",
			QueryPath: "end_before",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000",
			Default:   1000,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "order",
			Usage:     "Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'",
			Default:   "asc",
			QueryPath: "order",
		},
		&requestflag.Flag[string]{
			Name:      "start-after",
			Usage:     "The ID of the invoice to start the page after (exclusive). When provided, results will begin with the invoice immediately following this ID. Use this for standard forward pagination to get the next page of results. Cannot be combined with end_before.",
			QueryPath: "start_after",
		},
	},
	Action:          handleAccountsRecievableInvoicesList,
	HideHelpCommand: true,
}

var accountsRecievableInvoicesCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancel an invoice. This action cannot be undone.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "invoice-id",
			Usage:    "ID for the invoice.",
			Required: true,
		},
	},
	Action:          handleAccountsRecievableInvoicesCancel,
	HideHelpCommand: true,
}

var accountsRecievableInvoicesDownloadPdf = cli.Command{
	Name:    "download-pdf",
	Usage:   "Downloads a PDF file for the specified invoice. The response includes a\nContent-Disposition header set to 'attachment' with the filename.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "invoice-id",
			Usage:    "ID for the invoice.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "The file where the response contents will be stored. Use the value '-' to force output to stdout.",
		},
	},
	Action:          handleAccountsRecievableInvoicesDownloadPdf,
	HideHelpCommand: true,
}

var accountsRecievableInvoicesListAttachments = cli.Command{
	Name:    "list-attachments",
	Usage:   "Retrieve a list of all attachments for a specific invoice",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "invoice-id",
			Usage:    "ID for the invoice.",
			Required: true,
		},
	},
	Action:          handleAccountsRecievableInvoicesListAttachments,
	HideHelpCommand: true,
}

func handleAccountsRecievableInvoicesCreate(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.AccountsRecievableInvoiceNewParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AccountsRecievable.Invoices.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "accounts-recievable:invoices create", obj, format, transform)
}

func handleAccountsRecievableInvoicesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("invoice-id") && len(unusedArgs) > 0 {
		cmd.Set("invoice-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AccountsRecievable.Invoices.Get(ctx, cmd.Value("invoice-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "accounts-recievable:invoices retrieve", obj, format, transform)
}

func handleAccountsRecievableInvoicesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("invoice-id") && len(unusedArgs) > 0 {
		cmd.Set("invoice-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.AccountsRecievableInvoiceUpdateParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AccountsRecievable.Invoices.Update(
		ctx,
		cmd.Value("invoice-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "accounts-recievable:invoices update", obj, format, transform)
}

func handleAccountsRecievableInvoicesList(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.AccountsRecievableInvoiceListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.AccountsRecievable.Invoices.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "accounts-recievable:invoices list", obj, format, transform)
	} else {
		iter := client.AccountsRecievable.Invoices.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "accounts-recievable:invoices list", iter, format, transform)
	}
}

func handleAccountsRecievableInvoicesCancel(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("invoice-id") && len(unusedArgs) > 0 {
		cmd.Set("invoice-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	return client.AccountsRecievable.Invoices.Cancel(ctx, cmd.Value("invoice-id").(string), options...)
}

func handleAccountsRecievableInvoicesDownloadPdf(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("invoice-id") && len(unusedArgs) > 0 {
		cmd.Set("invoice-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	response, err := client.AccountsRecievable.Invoices.DownloadPdf(ctx, cmd.Value("invoice-id").(string), options...)
	if err != nil {
		return err
	}
	message, err := writeBinaryResponse(response, cmd.String("output"))
	if message != "" {
		fmt.Println(message)
	}
	return err
}

func handleAccountsRecievableInvoicesListAttachments(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("invoice-id") && len(unusedArgs) > 0 {
		cmd.Set("invoice-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AccountsRecievable.Invoices.ListAttachments(ctx, cmd.Value("invoice-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "accounts-recievable:invoices list-attachments", obj, format, transform)
}
