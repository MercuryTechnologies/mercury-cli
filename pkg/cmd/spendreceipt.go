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

var spendReceiptsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve directly uploaded receipts whose matching result is pending, ordered by\ncreation time, ascending (oldest first) by default. Upload URL requests do not\nappear until matching starts. Requires see-all expense permission. Use GET\n/spend/receipts/{receiptId} to poll terminal outcomes. Supports cursor-based\npagination with limit, order, start_at, start_after, and end_before query\nparameters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "end-before",
			Usage:     "The ID of the receipt to end the page before (exclusive). When provided, results will end just before this ID and work backwards. Use this for reverse pagination or to retrieve previous pages. Cannot be combined with start_after.",
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
			Usage:     "The ID of the receipt to start the page after (exclusive). When provided, results will begin with the receipt immediately following this ID. Use this for standard forward pagination to get the next page of results. Cannot be combined with end_before.",
			QueryPath: "start_after",
		},
		&requestflag.Flag[string]{
			Name:      "start-at",
			Usage:     "The ID of the receipt to start the page at (inclusive). Cannot be combined with start_after or end_before.",
			QueryPath: "start_at",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleSpendReceiptsList,
	HideHelpCommand: true,
}

var spendReceiptsCreateUploadURL = cli.Command{
	Name:    "create-upload-url",
	Usage:   "Request a presigned URL to upload a receipt file. The URL expires after five\nminutes; replay the request with the same Idempotency-Key to receive a fresh URL\nfor the same receipt. Upload the file, then start matching via the match\nendpoint.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "file-name",
			Required: true,
			BodyPath: "fileName",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			HeaderPath: "Idempotency-Key",
		},
	},
	Action:          handleSpendReceiptsCreateUploadURL,
	HideHelpCommand: true,
}

var spendReceiptsGet = cli.Command{
	Name:    "get",
	Usage:   "Retrieve a directly uploaded receipt and its auto-matching status by id.\nUploaders can retrieve their own receipts; reading another uploader's receipt\nrequires see-all expense permission.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "receipt-id",
			Usage:     "ID for the receipt",
			Required:  true,
			PathParam: "receiptId",
		},
	},
	Action:          handleSpendReceiptsGet,
	HideHelpCommand: true,
}

var spendReceiptsMatch = cli.Command{
	Name:    "match",
	Usage:   "Start matching an uploaded receipt against the organization's transactions\n(OCR + matcher). Only the uploader can start matching. Poll GET\n/spend/receipts/{receiptId} for the outcome. Returns 409 when matching has\nalready started. Returns 409 when another request restarted matching, and 502\nwhen processing fails; retry both with a new Idempotency-Key.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "receipt-id",
			Usage:     "ID for the receipt",
			Required:  true,
			PathParam: "receiptId",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			HeaderPath: "Idempotency-Key",
		},
	},
	Action:          handleSpendReceiptsMatch,
	HideHelpCommand: true,
}

func handleSpendReceiptsList(ctx context.Context, cmd *cli.Command) error {
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

	params := mercury.SpendReceiptListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Spend.Receipts.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "spend:receipts list",
			Transform:      transform,
		})
	} else {
		iter := client.Spend.Receipts.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "spend:receipts list",
			Transform:      transform,
		})
	}
}

func handleSpendReceiptsCreateUploadURL(ctx context.Context, cmd *cli.Command) error {
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

	params := mercury.SpendReceiptNewUploadURLParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Spend.Receipts.NewUploadURL(ctx, params, options...)
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
		Title:          "spend:receipts create-upload-url",
		Transform:      transform,
	})
}

func handleSpendReceiptsGet(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("receipt-id") && len(unusedArgs) > 0 {
		cmd.Set("receipt-id", unusedArgs[0])
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
	_, err = client.Spend.Receipts.Get(ctx, cmd.Value("receipt-id").(string), options...)
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
		Title:          "spend:receipts get",
		Transform:      transform,
	})
}

func handleSpendReceiptsMatch(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("receipt-id") && len(unusedArgs) > 0 {
		cmd.Set("receipt-id", unusedArgs[0])
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

	params := mercury.SpendReceiptMatchParams{}

	return client.Spend.Receipts.Match(
		ctx,
		cmd.Value("receipt-id").(string),
		params,
		options...,
	)
}
