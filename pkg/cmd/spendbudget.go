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

var spendBudgetsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a budget as a draft. Add members and publish it from the Mercury\ndashboard.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "frequency",
			Usage:    `Allowed values: "oneTime", "daily", "weekly", "monthly", "yearly".`,
			Required: true,
			BodyPath: "frequency",
		},
		&requestflag.Flag[*float64]{
			Name:     "amount-per-user-per-period",
			Usage:    "A non-negative dollar amount",
			BodyPath: "amountPerUserPerPeriod",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[any]{
			Name:     "end-date",
			BodyPath: "endDate",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "start-date",
			BodyPath: "startDate",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			HeaderPath: "Idempotency-Key",
		},
	},
	Action:          handleSpendBudgetsCreate,
	HideHelpCommand: true,
}

var spendBudgetsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a budget's fields. A draft is edited in place; a published budget gets a\nnew version. Omitted or null fields keep their current value. Frequency cannot\nchange after publication. Effective changes to amountPerUserPerPeriod,\nfrequency, startDate, or endDate return 409 when the budget has an active\nAgentic Card link; name and description updates remain allowed.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "budget-id",
			Usage:     "ID for the budget",
			Required:  true,
			PathParam: "budgetId",
		},
		&requestflag.Flag[*float64]{
			Name:     "amount-per-user-per-period",
			Usage:    "A non-negative dollar amount",
			BodyPath: "amountPerUserPerPeriod",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[any]{
			Name:     "end-date",
			BodyPath: "endDate",
		},
		&requestflag.Flag[*string]{
			Name:     "frequency",
			Usage:    `Allowed values: "oneTime", "daily", "weekly", "monthly", "yearly".`,
			BodyPath: "frequency",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "start-date",
			BodyPath: "startDate",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			HeaderPath: "Idempotency-Key",
		},
	},
	Action:          handleSpendBudgetsUpdate,
	HideHelpCommand: true,
}

var spendBudgetsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve the organization's non-deleted budgets ordered by creation time,\nascending (oldest first) by default. Supports cursor-based pagination with\nlimit, order, start_at, start_after, and end_before query parameters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "end-before",
			Usage:     "The ID of the budget to end the page before (exclusive). When provided, results will end just before this ID and work backwards. Use this for reverse pagination or to retrieve previous pages. Cannot be combined with start_after.",
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
			Usage:     "The ID of the budget to start the page after (exclusive). When provided, results will begin with the budget immediately following this ID. Use this for standard forward pagination to get the next page of results. Cannot be combined with end_before.",
			QueryPath: "start_after",
		},
		&requestflag.Flag[string]{
			Name:      "start-at",
			Usage:     "The ID of the budget to start the page at (inclusive). Cannot be combined with start_after or end_before.",
			QueryPath: "start_at",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleSpendBudgetsList,
	HideHelpCommand: true,
}

var spendBudgetsArchive = cli.Command{
	Name:    "archive",
	Usage:   "Archive a budget (a published budget becomes archived; a draft becomes deleted).\nReturns 409 when the budget has an active Agentic Card link.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "budget-id",
			Usage:     "ID for the budget",
			Required:  true,
			PathParam: "budgetId",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			HeaderPath: "Idempotency-Key",
		},
	},
	Action:          handleSpendBudgetsArchive,
	HideHelpCommand: true,
}

var spendBudgetsAssignTransaction = cli.Command{
	Name:    "assign-transaction",
	Usage:   "Assign a transaction to a budget so its spend counts against the cycle\ncontaining the transaction's date. If the assignment would overspend the budget,\nset allowOverspend to true; omitted and false values enforce the spend limit.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "budget-id",
			Usage:     "ID for the budget",
			Required:  true,
			PathParam: "budgetId",
		},
		&requestflag.Flag[string]{
			Name:     "transaction-id",
			Usage:    "ID for this transaction",
			Required: true,
			BodyPath: "transactionId",
		},
		&requestflag.Flag[*bool]{
			Name:     "allow-overspend",
			Usage:    " Whether to allow this assignment to exceed the budget. Defaults to false.",
			BodyPath: "allowOverspend",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			HeaderPath: "Idempotency-Key",
		},
	},
	Action:          handleSpendBudgetsAssignTransaction,
	HideHelpCommand: true,
}

var spendBudgetsGet = cli.Command{
	Name:    "get",
	Usage:   "Retrieve a budget by id, including its lifecycle status, per-user limit,\ncadence, and effective window.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "budget-id",
			Usage:     "ID for the budget",
			Required:  true,
			PathParam: "budgetId",
		},
	},
	Action:          handleSpendBudgetsGet,
	HideHelpCommand: true,
}

var spendBudgetsListExpenditures = cli.Command{
	Name:    "list-expenditures",
	Usage:   "List the non-deleted expenditures assigned to a budget, ordered by expenditure\ncreation time, ascending (oldest first) by default. Supports cursor-based\npagination with limit, order, start_at, start_after, and end_before query\nparameters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "budget-id",
			Usage:     "ID for the budget",
			Required:  true,
			PathParam: "budgetId",
		},
		&requestflag.Flag[string]{
			Name:      "end-before",
			Usage:     "The ID of the budget expenditure to end the page before (exclusive). When provided, results will end just before this ID and work backwards. Use this for reverse pagination or to retrieve previous pages. Cannot be combined with start_after.",
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
			Usage:     "The ID of the budget expenditure to start the page after (exclusive). When provided, results will begin with the budget expenditure immediately following this ID. Use this for standard forward pagination to get the next page of results. Cannot be combined with end_before.",
			QueryPath: "start_after",
		},
		&requestflag.Flag[string]{
			Name:      "start-at",
			Usage:     "The ID of the budget expenditure to start the page at (inclusive). Cannot be combined with start_after or end_before.",
			QueryPath: "start_at",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleSpendBudgetsListExpenditures,
	HideHelpCommand: true,
}

var spendBudgetsListMembers = cli.Command{
	Name:    "list-members",
	Usage:   "List the members of a budget and their current or final cycle spend, ordered by\nthe creation time of each budget member's assignment record, ascending (oldest\nfirst) by default. Supports cursor-based pagination with limit, order, start_at,\nstart_after, and end_before query parameters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "budget-id",
			Usage:     "ID for the budget",
			Required:  true,
			PathParam: "budgetId",
		},
		&requestflag.Flag[string]{
			Name:      "end-before",
			Usage:     "The ID of the organization user to end the page before (exclusive). When provided, results will end just before this ID and work backwards. Use this for reverse pagination or to retrieve previous pages. Cannot be combined with start_after.",
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
			Usage:     "The ID of the organization user to start the page after (exclusive). When provided, results will begin with the organization user immediately following this ID. Use this for standard forward pagination to get the next page of results. Cannot be combined with end_before.",
			QueryPath: "start_after",
		},
		&requestflag.Flag[string]{
			Name:      "start-at",
			Usage:     "The ID of the organization user to start the page at (inclusive). Cannot be combined with start_after or end_before.",
			QueryPath: "start_at",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleSpendBudgetsListMembers,
	HideHelpCommand: true,
}

func handleSpendBudgetsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := mercury.SpendBudgetNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Spend.Budgets.New(ctx, params, options...)
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
		Title:          "spend:budgets create",
		Transform:      transform,
	})
}

func handleSpendBudgetsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("budget-id") && len(unusedArgs) > 0 {
		cmd.Set("budget-id", unusedArgs[0])
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

	params := mercury.SpendBudgetUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Spend.Budgets.Update(
		ctx,
		cmd.Value("budget-id").(string),
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
		Title:          "spend:budgets update",
		Transform:      transform,
	})
}

func handleSpendBudgetsList(ctx context.Context, cmd *cli.Command) error {
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

	params := mercury.SpendBudgetListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Spend.Budgets.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "spend:budgets list",
			Transform:      transform,
		})
	} else {
		iter := client.Spend.Budgets.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "spend:budgets list",
			Transform:      transform,
		})
	}
}

func handleSpendBudgetsArchive(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("budget-id") && len(unusedArgs) > 0 {
		cmd.Set("budget-id", unusedArgs[0])
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

	params := mercury.SpendBudgetArchiveParams{}

	return client.Spend.Budgets.Archive(
		ctx,
		cmd.Value("budget-id").(string),
		params,
		options...,
	)
}

func handleSpendBudgetsAssignTransaction(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("budget-id") && len(unusedArgs) > 0 {
		cmd.Set("budget-id", unusedArgs[0])
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

	params := mercury.SpendBudgetAssignTransactionParams{}

	return client.Spend.Budgets.AssignTransaction(
		ctx,
		cmd.Value("budget-id").(string),
		params,
		options...,
	)
}

func handleSpendBudgetsGet(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("budget-id") && len(unusedArgs) > 0 {
		cmd.Set("budget-id", unusedArgs[0])
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
	_, err = client.Spend.Budgets.Get(ctx, cmd.Value("budget-id").(string), options...)
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
		Title:          "spend:budgets get",
		Transform:      transform,
	})
}

func handleSpendBudgetsListExpenditures(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("budget-id") && len(unusedArgs) > 0 {
		cmd.Set("budget-id", unusedArgs[0])
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

	params := mercury.SpendBudgetListExpendituresParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Spend.Budgets.ListExpenditures(
			ctx,
			cmd.Value("budget-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "spend:budgets list-expenditures",
			Transform:      transform,
		})
	} else {
		iter := client.Spend.Budgets.ListExpendituresAutoPaging(
			ctx,
			cmd.Value("budget-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "spend:budgets list-expenditures",
			Transform:      transform,
		})
	}
}

func handleSpendBudgetsListMembers(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("budget-id") && len(unusedArgs) > 0 {
		cmd.Set("budget-id", unusedArgs[0])
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

	params := mercury.SpendBudgetListMembersParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Spend.Budgets.ListMembers(
			ctx,
			cmd.Value("budget-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "spend:budgets list-members",
			Transform:      transform,
		})
	} else {
		iter := client.Spend.Budgets.ListMembersAutoPaging(
			ctx,
			cmd.Value("budget-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "spend:budgets list-members",
			Transform:      transform,
		})
	}
}
