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

var spendReimbursementsCreate = cli.Command{
	Name:    "create",
	Usage:   "Submit a new reimbursement as the token's user. Only that user can edit or\ncancel it. Returns the new reimbursement id.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[float64]{
			Name:     "amount",
			Usage:    " The positive amount spent.",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "currency",
			Usage:    ` The currency the money was spent in (e.g. "USD").`,
			Required: true,
			BodyPath: "currency",
		},
		&requestflag.Flag[any]{
			Name:     "date-paid-for",
			Usage:    " When the money was actually spent.",
			Required: true,
			BodyPath: "datePaidFor",
		},
		&requestflag.Flag[string]{
			Name:     "merchant-name",
			Usage:    ` Where the money was spent, e.g. "Starbucks".`,
			Required: true,
			BodyPath: "merchantName",
		},
		&requestflag.Flag[*string]{
			Name:     "category-id",
			Usage:    "ID for the category",
			BodyPath: "categoryId",
		},
		&requestflag.Flag[*string]{
			Name:     "notes",
			BodyPath: "notes",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			HeaderPath: "Idempotency-Key",
		},
	},
	Action:          handleSpendReimbursementsCreate,
	HideHelpCommand: true,
}

var spendReimbursementsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Replace the fields of an existing reimbursement. Only reimbursements still in\nreview can be edited.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "reimbursement-id",
			Usage:     "ID for the expense",
			Required:  true,
			PathParam: "reimbursementId",
		},
		&requestflag.Flag[float64]{
			Name:     "amount",
			Usage:    " The positive amount spent.",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "currency",
			Usage:    ` The currency the money was spent in (e.g. "USD").`,
			Required: true,
			BodyPath: "currency",
		},
		&requestflag.Flag[any]{
			Name:     "date-paid-for",
			Usage:    " When the money was actually spent.",
			Required: true,
			BodyPath: "datePaidFor",
		},
		&requestflag.Flag[string]{
			Name:     "merchant-name",
			Usage:    ` Where the money was spent, e.g. "Starbucks".`,
			Required: true,
			BodyPath: "merchantName",
		},
		&requestflag.Flag[*string]{
			Name:     "category-id",
			Usage:    "ID for the category",
			BodyPath: "categoryId",
		},
		&requestflag.Flag[*string]{
			Name:     "notes",
			BodyPath: "notes",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			HeaderPath: "Idempotency-Key",
		},
	},
	Action:          handleSpendReimbursementsUpdate,
	HideHelpCommand: true,
}

var spendReimbursementsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a paginated list of reimbursements for the organization ordered by\ncreation time, ascending (oldest first) by default. Each reimbursement includes\nits attachments, with an id, a file name, and a presigned download URL. The URL\nexpires after 24 hours. Request the reimbursement again for a new URL. Supports\ncursor-based pagination with limit, order, start_at, start_after, and end_before\nquery parameters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "end-before",
			Usage:     "The ID of the reimbursement to end the page before (exclusive). When provided, results will end just before this ID and work backwards. Use this for reverse pagination or to retrieve previous pages. Cannot be combined with start_after.",
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
			Usage:     "The ID of the reimbursement to start the page after (exclusive). When provided, results will begin with the reimbursement immediately following this ID. Use this for standard forward pagination to get the next page of results. Cannot be combined with end_before.",
			QueryPath: "start_after",
		},
		&requestflag.Flag[string]{
			Name:      "start-at",
			Usage:     "The ID of the reimbursement to start the page at (inclusive). Cannot be combined with start_after or end_before.",
			QueryPath: "start_at",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleSpendReimbursementsList,
	HideHelpCommand: true,
}

var spendReimbursementsApprove = cli.Command{
	Name:    "approve",
	Usage:   "Approve a submitted reimbursement. The token's user needs the \"Reimbursement\napproval and submission\" or \"Policy management and approval\" permission, and\naccess to the organization's reimbursement account. The token needs the\n`reimbursements:manage` scope or Read and Write access. Self-approval works only\nwhen the organization allows it or has only one admin.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "reimbursement-id",
			Usage:     "ID for the expense",
			Required:  true,
			PathParam: "reimbursementId",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			HeaderPath: "Idempotency-Key",
		},
	},
	Action:          handleSpendReimbursementsApprove,
	HideHelpCommand: true,
}

var spendReimbursementsCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancel a reimbursement that is still in review.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "reimbursement-id",
			Usage:     "ID for the expense",
			Required:  true,
			PathParam: "reimbursementId",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			HeaderPath: "Idempotency-Key",
		},
	},
	Action:          handleSpendReimbursementsCancel,
	HideHelpCommand: true,
}

var spendReimbursementsDeny = cli.Command{
	Name:    "deny",
	Usage:   "Deny a submitted reimbursement. The reason is recorded as a reviewer comment.\nThe token's user needs the \"Reimbursement approval and submission\" or \"Policy\nmanagement and approval\" permission, and access to the organization's\nreimbursement account. The token needs the `reimbursements:manage` scope or Read\nand Write access.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "reimbursement-id",
			Usage:     "ID for the expense",
			Required:  true,
			PathParam: "reimbursementId",
		},
		&requestflag.Flag[string]{
			Name:     "reason",
			Required: true,
			BodyPath: "reason",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			HeaderPath: "Idempotency-Key",
		},
	},
	Action:          handleSpendReimbursementsDeny,
	HideHelpCommand: true,
}

var spendReimbursementsGet = cli.Command{
	Name:    "get",
	Usage:   "Retrieve details of a specific reimbursement by its id, including its review and\npayment status. Attachments include an id, a file name, and a presigned download\nURL. The URL expires after 24 hours. Request the reimbursement again for a new\nURL.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "reimbursement-id",
			Usage:     "ID for the expense",
			Required:  true,
			PathParam: "reimbursementId",
		},
	},
	Action:          handleSpendReimbursementsGet,
	HideHelpCommand: true,
}

var spendReimbursementsMarkPaid = cli.Command{
	Name:    "mark-paid",
	Usage:   "Record that the submitter was paid outside Mercury, for example by check or\npayroll. Mercury won’t pay the reimbursement, and its payment status reads as\npaidOutsideMercury. Denied or canceled reimbursements, and ones with a Mercury\npayment in progress or completed, can’t be marked. The optional note is recorded\nas a reviewer comment. Approving or denying the reimbursement afterwards undoes\nthe mark.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "reimbursement-id",
			Usage:     "ID for the expense",
			Required:  true,
			PathParam: "reimbursementId",
		},
		&requestflag.Flag[*string]{
			Name:     "note",
			Usage:    " Recorded as a reviewer comment, e.g. how and when the submitter was paid.",
			BodyPath: "note",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			HeaderPath: "Idempotency-Key",
		},
	},
	Action:          handleSpendReimbursementsMarkPaid,
	HideHelpCommand: true,
}

func handleSpendReimbursementsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := mercury.SpendReimbursementNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Spend.Reimbursements.New(ctx, params, options...)
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
		Title:          "spend:reimbursements create",
		Transform:      transform,
	})
}

func handleSpendReimbursementsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("reimbursement-id") && len(unusedArgs) > 0 {
		cmd.Set("reimbursement-id", unusedArgs[0])
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

	params := mercury.SpendReimbursementUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Spend.Reimbursements.Update(
		ctx,
		cmd.Value("reimbursement-id").(string),
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
		Title:          "spend:reimbursements update",
		Transform:      transform,
	})
}

func handleSpendReimbursementsList(ctx context.Context, cmd *cli.Command) error {
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

	params := mercury.SpendReimbursementListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Spend.Reimbursements.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "spend:reimbursements list",
			Transform:      transform,
		})
	} else {
		iter := client.Spend.Reimbursements.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "spend:reimbursements list",
			Transform:      transform,
		})
	}
}

func handleSpendReimbursementsApprove(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("reimbursement-id") && len(unusedArgs) > 0 {
		cmd.Set("reimbursement-id", unusedArgs[0])
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

	params := mercury.SpendReimbursementApproveParams{}

	return client.Spend.Reimbursements.Approve(
		ctx,
		cmd.Value("reimbursement-id").(string),
		params,
		options...,
	)
}

func handleSpendReimbursementsCancel(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("reimbursement-id") && len(unusedArgs) > 0 {
		cmd.Set("reimbursement-id", unusedArgs[0])
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

	params := mercury.SpendReimbursementCancelParams{}

	return client.Spend.Reimbursements.Cancel(
		ctx,
		cmd.Value("reimbursement-id").(string),
		params,
		options...,
	)
}

func handleSpendReimbursementsDeny(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("reimbursement-id") && len(unusedArgs) > 0 {
		cmd.Set("reimbursement-id", unusedArgs[0])
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

	params := mercury.SpendReimbursementDenyParams{}

	return client.Spend.Reimbursements.Deny(
		ctx,
		cmd.Value("reimbursement-id").(string),
		params,
		options...,
	)
}

func handleSpendReimbursementsGet(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("reimbursement-id") && len(unusedArgs) > 0 {
		cmd.Set("reimbursement-id", unusedArgs[0])
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
	_, err = client.Spend.Reimbursements.Get(ctx, cmd.Value("reimbursement-id").(string), options...)
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
		Title:          "spend:reimbursements get",
		Transform:      transform,
	})
}

func handleSpendReimbursementsMarkPaid(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("reimbursement-id") && len(unusedArgs) > 0 {
		cmd.Set("reimbursement-id", unusedArgs[0])
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

	params := mercury.SpendReimbursementMarkPaidParams{}

	return client.Spend.Reimbursements.MarkPaid(
		ctx,
		cmd.Value("reimbursement-id").(string),
		params,
		options...,
	)
}
