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

var accountTransactionsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a paginated list of transactions for a specific account. Supports\nfiltering by date range, status, and search terms.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-id",
			Usage:    "ID for a Mercury account.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "category-id",
			Usage:     "UUID of a custom category. Can be returned from /categories endpoint.",
			QueryPath: "categoryId",
		},
		&requestflag.Flag[string]{
			Name:      "end",
			Usage:     "Latest date to filter transactions. If not provided, defaults to the current date. Format: YYYY-MM-DD or ISO 8601 string",
			QueryPath: "end",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000",
			Default:   1000,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "mercury-category",
			Usage:     "Name of mercuryCategory you want to filter on. Merchant Type in the UI.",
			QueryPath: "mercuryCategory",
		},
		&requestflag.Flag[int64]{
			Name:      "offset",
			Usage:     "Number of results to skip for pagination",
			QueryPath: "offset",
		},
		&requestflag.Flag[string]{
			Name:      "order",
			Usage:     "Sort order. Can be 'asc' or 'desc'. Defaults to 'desc'",
			Default:   "desc",
			QueryPath: "order",
		},
		&requestflag.Flag[string]{
			Name:      "request-id",
			Usage:     "ID returned from /account/:id/request-send-money",
			QueryPath: "requestId",
		},
		&requestflag.Flag[string]{
			Name:      "search",
			Usage:     "Search term to filter transactions by description or counterparty name",
			QueryPath: "search",
		},
		&requestflag.Flag[string]{
			Name:      "start",
			Usage:     "Earliest date to filter transactions. If not provided, defaults to 30 days before the current date. Format: YYYY-MM-DD or ISO 8601 string",
			QueryPath: "start",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     `Allowed values: "pending", "sent", "cancelled", "failed", "reversed", "blocked".`,
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAccountTransactionsList,
	HideHelpCommand: true,
}

var accountTransactionsSend = requestflag.WithInnerFlags(cli.Command{
	Name:    "send",
	Usage:   "Send money from an account to a recipient. Creates a transaction that will be\nprocessed immediately or may require approval.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-id",
			Usage:    "ID for a Mercury account.",
			Required: true,
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
	Action:          handleAccountTransactionsSend,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"purpose": {
		&requestflag.InnerFlag[any]{
			Name:       "purpose.simple",
			InnerField: "simple",
		},
	},
})

func handleAccountTransactionsList(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.AccountTransactionListParams{}

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
		_, err = client.Account.Transactions.List(
			ctx,
			cmd.Value("account-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "account:transactions list", obj, format, transform)
	} else {
		iter := client.Account.Transactions.ListAutoPaging(
			ctx,
			cmd.Value("account-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "account:transactions list", iter, format, transform, maxItems)
	}
}

func handleAccountTransactionsSend(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.AccountTransactionSendParams{}

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
	_, err = client.Account.Transactions.Send(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "account:transactions send", obj, format, transform)
}
