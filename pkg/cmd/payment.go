// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/MercuryTechnologies/mercury-cli/internal/apiquery"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
	"github.com/MercuryTechnologies/mercury-go"
	"github.com/MercuryTechnologies/mercury-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var paymentsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Send money from an account to a recipient. Creates a transaction that will be\nprocessed immediately or may require approval.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "ID for a Mercury account.",
			Required:  true,
			PathParam: "accountId",
		},
		&requestflag.Flag[float64]{
			Name:     "amount",
			Usage:    "A positive dollar amount with at least 1 cent.",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "idempotency-key",
			Usage:    "Unique string identifying the transaction",
			Required: true,
			BodyPath: "idempotencyKey",
		},
		&requestflag.Flag[string]{
			Name:     "payment-method",
			Usage:    "If domesticWire is used, then the purpose field is required.",
			Required: true,
			BodyPath: "paymentMethod",
		},
		&requestflag.Flag[string]{
			Name:     "recipient-id",
			Usage:    "ID for a Mercury account.",
			Required: true,
			BodyPath: "recipientId",
		},
		&requestflag.Flag[string]{
			Name:     "external-memo",
			Usage:    "Optional external memo",
			BodyPath: "externalMemo",
		},
		&requestflag.Flag[string]{
			Name:     "note",
			Usage:    "Optional note",
			BodyPath: "note",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "purpose",
			Usage:    " External API representation of SendMoneyPurpose.\n Only exposes the 'simple' field to decouple internal implementation from external API.",
			BodyPath: "purpose",
		},
	},
	Action:          handlePaymentsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"purpose": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "purpose.simple",
			InnerField: "simple",
		},
	},
})

var paymentsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a paginated list of send money approval requests for the authenticated\norganization. Supports filtering by account and status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "ID for a Mercury account.",
			QueryPath: "accountId",
		},
		&requestflag.Flag[string]{
			Name:      "end-before",
			Usage:     "The ID of the send money approval request to end the page before (exclusive). When provided, results will end just before this ID and work backwards. Use this for reverse pagination or to retrieve previous pages. Cannot be combined with start_after.",
			QueryPath: "end_before",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000",
			Default:   1000,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "start-after",
			Usage:     "The ID of the send money approval request to start the page after (exclusive). When provided, results will begin with the send money approval request immediately following this ID. Use this for standard forward pagination to get the next page of results. Cannot be combined with end_before.",
			QueryPath: "start_after",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     `Allowed values: "pendingApproval", "approved", "rejected", "cancelled".`,
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handlePaymentsList,
	HideHelpCommand: true,
}

var paymentsGet = cli.Command{
	Name:    "get",
	Usage:   "Get send money approval request by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "request-id",
			Usage:     "ID for the send money approval request",
			Required:  true,
			PathParam: "requestId",
		},
	},
	Action:          handlePaymentsGet,
	HideHelpCommand: true,
}

var paymentsRequest = cli.Command{
	Name:    "request",
	Usage:   "Create a \"request to send money\" that will require approval based on your\norganization's approval policies.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "ID for a Mercury account.",
			Required:  true,
			PathParam: "accountId",
		},
		&requestflag.Flag[float64]{
			Name:     "amount",
			Usage:    "A positive dollar amount with at least 1 cent.",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "idempotency-key",
			Usage:    "Unique string identifying the transaction",
			Required: true,
			BodyPath: "idempotencyKey",
		},
		&requestflag.Flag[string]{
			Name:     "payment-method",
			Usage:    `Allowed values: "ach", "check", "domesticWire", "internationalWire".`,
			Required: true,
			BodyPath: "paymentMethod",
		},
		&requestflag.Flag[string]{
			Name:     "recipient-id",
			Usage:    "ID for a Mercury account.",
			Required: true,
			BodyPath: "recipientId",
		},
		&requestflag.Flag[string]{
			Name:     "external-memo",
			Usage:    "Optional external memo",
			BodyPath: "externalMemo",
		},
		&requestflag.Flag[string]{
			Name:     "note",
			Usage:    "Optional note",
			BodyPath: "note",
		},
	},
	Action:          handlePaymentsRequest,
	HideHelpCommand: true,
}

var paymentsTransfer = cli.Command{
	Name:    "transfer",
	Usage:   "Transfer funds between two accounts within the same organization. Supports\ntransfers between depository accounts (checking/savings), from a depository\naccount to a treasury/investment account, and from a treasury/investment account\nto a depository account. Creates paired debit and credit transactions.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[float64]{
			Name:     "amount",
			Usage:    "A positive dollar amount with at least 1 cent.",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "destination-account-id",
			Usage:    "ID for a Mercury account.",
			Required: true,
			BodyPath: "destinationAccountId",
		},
		&requestflag.Flag[string]{
			Name:     "idempotency-key",
			Required: true,
			BodyPath: "idempotencyKey",
		},
		&requestflag.Flag[string]{
			Name:     "source-account-id",
			Usage:    "ID for a Mercury account.",
			Required: true,
			BodyPath: "sourceAccountId",
		},
		&requestflag.Flag[*string]{
			Name:     "note",
			BodyPath: "note",
		},
	},
	Action:          handlePaymentsTransfer,
	HideHelpCommand: true,
}

func handlePaymentsCreate(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
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

	params := mercury.PaymentNewParams{}

	// CUSTOM: confirmation prompt before sending money
	if err := confirmAction(cmd, "Send Money", buildPaymentConfirmDetails(cmd)); err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Payments.New(
		ctx,
		cmd.Value("account-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "payments create",
		Transform:      transform,
	})
}

func handlePaymentsList(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := mercury.PaymentListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Payments.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "payments list",
			Transform:      transform,
		})
	} else {
		iter := client.Payments.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "payments list",
			Transform:      transform,
		})
	}
}

func handlePaymentsGet(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("request-id") && len(unusedArgs) > 0 {
		cmd.Set("request-id", unusedArgs[0])
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
	_, err = client.Payments.Get(ctx, cmd.Value("request-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "payments get",
		Transform:      transform,
	})
}

func handlePaymentsRequest(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
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

	params := mercury.PaymentRequestParams{}

	// CUSTOM: confirmation prompt before requesting to send money
	if err := confirmAction(cmd, "Request to Send Money", buildPaymentConfirmDetails(cmd)); err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Payments.Request(
		ctx,
		cmd.Value("account-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "payments request",
		Transform:      transform,
	})
}

func handlePaymentsTransfer(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
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

	// CUSTOM: confirmation prompt before transferring funds
	if err := confirmAction(cmd, "Transfer Funds", buildTransferConfirmDetails(cmd)); err != nil {
		return err
	}

	params := mercury.PaymentTransferParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Payments.Transfer(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "payments transfer",
		Transform:      transform,
	})
}
