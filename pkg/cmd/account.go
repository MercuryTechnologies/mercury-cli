// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/MercuryTechnologies/mercury-cli/internal/apiquery"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
	"github.com/MercuryTechnologies/mercury-go"
	"github.com/MercuryTechnologies/mercury-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var accountsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get account by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-id",
			Usage:    "ID for a Mercury account.",
			Required: true,
		},
	},
	Action:          handleAccountsRetrieve,
	HideHelpCommand: true,
}

var accountsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a paginated list of accounts. Supports cursor-based pagination with\nlimit, order, start_after, and end_before query parameters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "end-before",
			Usage:     "The ID of the account to end the page before (exclusive). When provided, results will end just before this ID and work backwards. Use this for reverse pagination or to retrieve previous pages. Cannot be combined with start_after.",
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
			Usage:     "The ID of the account to start the page after (exclusive). When provided, results will begin with the account immediately following this ID. Use this for standard forward pagination to get the next page of results. Cannot be combined with end_before.",
			QueryPath: "start_after",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAccountsList,
	HideHelpCommand: true,
}

var accountsCreateTransaction = requestflag.WithInnerFlags(cli.Command{
	Name:    "create-transaction",
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
	Action:          handleAccountsCreateTransaction,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"purpose": {
		&requestflag.InnerFlag[any]{
			Name:       "purpose.simple",
			InnerField: "simple",
		},
	},
})

var accountsListCards = cli.Command{
	Name:    "list-cards",
	Usage:   "Retrieve all debit and credit cards associated with a specific account.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-id",
			Usage:    "ID for a Mercury account.",
			Required: true,
		},
	},
	Action:          handleAccountsListCards,
	HideHelpCommand: true,
}

var accountsListStatements = cli.Command{
	Name:    "list-statements",
	Usage:   "Retrieve a paginated list of monthly statements for a specific account. Supports\ncursor-based pagination with limit, order, start_after, and end_before query\nparameters, as well as date range filtering with start and end parameters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-id",
			Usage:    "ID for a Mercury account.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "end",
			Usage:     "Filter statements where the period start date is on or before this date. If the date is in the future, defaults to the current date. Format: YYYY-MM-DD",
			QueryPath: "end",
		},
		&requestflag.Flag[string]{
			Name:      "end-before",
			Usage:     "The ID of the statement to end the page before (exclusive). When provided, results will end just before this ID and work backwards. Use this for reverse pagination or to retrieve previous pages. Cannot be combined with start_after.",
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
			Usage:     "Sort order. Can be 'asc' or 'desc'. Defaults to 'desc'",
			Default:   "desc",
			QueryPath: "order",
		},
		&requestflag.Flag[string]{
			Name:      "start",
			Usage:     "Filter statements where the period start date is on or after this date. Format: YYYY-MM-DD",
			QueryPath: "start",
		},
		&requestflag.Flag[string]{
			Name:      "start-after",
			Usage:     "The ID of the statement to start the page after (exclusive). When provided, results will begin with the statement immediately following this ID. Use this for standard forward pagination to get the next page of results. Cannot be combined with end_before.",
			QueryPath: "start_after",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAccountsListStatements,
	HideHelpCommand: true,
}

var accountsListTransactions = cli.Command{
	Name:    "list-transactions",
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
	Action:          handleAccountsListTransactions,
	HideHelpCommand: true,
}

var accountsRequestSendMoney = cli.Command{
	Name:    "request-send-money",
	Usage:   "Create a \"request to send money\" that will require approval based on your\norganization's approval policies.",
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
	Action:          handleAccountsRequestSendMoney,
	HideHelpCommand: true,
}

var accountsRetrieveTransaction = cli.Command{
	Name:    "retrieve-transaction",
	Usage:   "Get transaction by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "account-id",
			Usage:    "ID for a Mercury account.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "transaction-id",
			Usage:    "ID for this transaction",
			Required: true,
		},
	},
	Action:          handleAccountsRetrieveTransaction,
	HideHelpCommand: true,
}

func handleAccountsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Accounts.Get(ctx, cmd.Value("account-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "accounts retrieve", obj, format, transform)
}

func handleAccountsList(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.AccountListParams{}

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
		_, err = client.Accounts.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "accounts list", obj, format, transform)
	} else {
		iter := client.Accounts.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "accounts list", iter, format, transform, maxItems)
	}
}

func handleAccountsCreateTransaction(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.AccountNewTransactionParams{}

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
	_, err = client.Accounts.NewTransaction(
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
	return ShowJSON(os.Stdout, "accounts create-transaction", obj, format, transform)
}

func handleAccountsListCards(ctx context.Context, cmd *cli.Command) error {
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
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Accounts.ListCards(ctx, cmd.Value("account-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "accounts list-cards", obj, format, transform)
}

func handleAccountsListStatements(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.AccountListStatementsParams{}

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
		_, err = client.Accounts.ListStatements(
			ctx,
			cmd.Value("account-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "accounts list-statements", obj, format, transform)
	} else {
		iter := client.Accounts.ListStatementsAutoPaging(
			ctx,
			cmd.Value("account-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "accounts list-statements", iter, format, transform, maxItems)
	}
}

func handleAccountsListTransactions(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.AccountListTransactionsParams{}

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
		_, err = client.Accounts.ListTransactions(
			ctx,
			cmd.Value("account-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "accounts list-transactions", obj, format, transform)
	} else {
		iter := client.Accounts.ListTransactionsAutoPaging(
			ctx,
			cmd.Value("account-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "accounts list-transactions", iter, format, transform, maxItems)
	}
}

func handleAccountsRequestSendMoney(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.AccountRequestSendMoneyParams{}

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
	_, err = client.Accounts.RequestSendMoney(
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
	return ShowJSON(os.Stdout, "accounts request-send-money", obj, format, transform)
}

func handleAccountsRetrieveTransaction(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("transaction-id") && len(unusedArgs) > 0 {
		cmd.Set("transaction-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.AccountGetTransactionParams{
		AccountID: cmd.Value("account-id").(string),
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
	_, err = client.Accounts.GetTransaction(
		ctx,
		cmd.Value("transaction-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "accounts retrieve-transaction", obj, format, transform)
}
