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

var transactionsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update the note and/or category of an existing transaction. Use null values to\nclear existing data.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "transaction-id",
			Usage:    "ID for this transaction",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "category-id",
			Usage:    "ID for the category",
			Required: true,
			BodyPath: "categoryId",
		},
		&requestflag.Flag[any]{
			Name:     "note",
			Usage:    "Note update action. Omit field to keep current note, send null or empty string to clear note, send text to set note.",
			Required: true,
			BodyPath: "note",
		},
	},
	Action:          handleTransactionsUpdate,
	HideHelpCommand: true,
}

var transactionsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a paginated list of all transactions across all accounts. Supports\nadvanced filtering by date ranges, status, categories, and cursor-based\npagination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:      "account-id",
			QueryPath: "accountId",
		},
		&requestflag.Flag[string]{
			Name:      "category-id",
			Usage:     "UUID of a custom category. Can be returned from /categories endpoint.",
			QueryPath: "categoryId",
		},
		&requestflag.Flag[string]{
			Name:      "end",
			Usage:     "Latest createdAt date to filter for. If it’s not provided, it defaults to current day. Format: YYYY-MM-DD or an ISO 8601 string. Please note that your Mercury transactions on your Dashboard might have their postedAt date displayed, as opposed to createdAt",
			QueryPath: "end",
		},
		&requestflag.Flag[string]{
			Name:      "end-before",
			Usage:     "The ID of the transaction to end the page before (exclusive). When provided, results will end just before this ID and work backwards. Use this for reverse pagination or to retrieve previous pages. Cannot be combined with start_after.",
			QueryPath: "end_before",
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
		&requestflag.Flag[string]{
			Name:      "order",
			Usage:     "Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'",
			Default:   "asc",
			QueryPath: "order",
		},
		&requestflag.Flag[string]{
			Name:      "posted-end",
			Usage:     "Latest postedAt date to filter for. Format: YYYY-MM-DD or an ISO 8601 string",
			QueryPath: "postedEnd",
		},
		&requestflag.Flag[string]{
			Name:      "posted-start",
			Usage:     "Earliest postedAt date to filter for. Format: YYYY-MM-DD or an ISO 8601 string",
			QueryPath: "postedStart",
		},
		&requestflag.Flag[string]{
			Name:      "search",
			Usage:     "Search term to look for in transaction descriptions.",
			QueryPath: "search",
		},
		&requestflag.Flag[string]{
			Name:      "start",
			Usage:     "Earliest createdAt date to filter for. If not provided, it defaults to the date of your first transaction. Format: YYYY-MM-DD or an ISO 8601 string. Please note that your Mercury transactions on your Dashboard might have their postedAt date displayed, as opposed to createdAt",
			QueryPath: "start",
		},
		&requestflag.Flag[string]{
			Name:      "start-after",
			Usage:     "The ID of the transaction to start the page after (exclusive). When provided, results will begin with the transaction immediately following this ID. Use this for standard forward pagination to get the next page of results. Cannot be combined with end_before.",
			QueryPath: "start_after",
		},
		&requestflag.Flag[string]{
			Name:      "start-at",
			Usage:     "The ID of the resource to start the page at (inclusive). When provided, results will begin with and include the resource with this ID. Use this to retrieve a specific page when you know the exact starting point. Cannot be combined with start_after or end_before.",
			QueryPath: "start_at",
		},
		&requestflag.Flag[[]string]{
			Name:      "status",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleTransactionsList,
	HideHelpCommand: true,
}

var transactionsGet = cli.Command{
	Name:    "get",
	Usage:   "Retrieve a single transaction by its ID. Returns full transaction details\nincluding attachments, check images, and related metadata.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "transaction-id",
			Usage:    "ID for this transaction",
			Required: true,
		},
	},
	Action:          handleTransactionsGet,
	HideHelpCommand: true,
}

func handleTransactionsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("transaction-id") && len(unusedArgs) > 0 {
		cmd.Set("transaction-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.TransactionUpdateParams{}

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
	_, err = client.Transactions.Update(
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "transactions update",
		Transform:      transform,
	})
}

func handleTransactionsList(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.TransactionListParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Transactions.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "transactions list",
			Transform:      transform,
		})
	} else {
		iter := client.Transactions.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "transactions list",
			Transform:      transform,
		})
	}
}

func handleTransactionsGet(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("transaction-id") && len(unusedArgs) > 0 {
		cmd.Set("transaction-id", unusedArgs[0])
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
	_, err = client.Transactions.Get(ctx, cmd.Value("transaction-id").(string), options...)
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
		Title:          "transactions get",
		Transform:      transform,
	})
}
