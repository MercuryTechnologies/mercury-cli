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

var recipientsAttachmentsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a paginated list of all recipient tax form attachments across all\nrecipients in the organization. Use cursor parameters (start_after, end_before)\nfor pagination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "end-before",
			Usage:     "The ID of the recipient attachment to end the page before (exclusive). When provided, results will end just before this ID and work backwards. Use this for reverse pagination or to retrieve previous pages. Cannot be combined with start_after.",
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
			Usage:     "The ID of the recipient attachment to start the page after (exclusive). When provided, results will begin with the recipient attachment immediately following this ID. Use this for standard forward pagination to get the next page of results. Cannot be combined with end_before.",
			QueryPath: "start_after",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleRecipientsAttachmentsList,
	HideHelpCommand: true,
}

var recipientsAttachmentsAttach = cli.Command{
	Name:    "attach",
	Usage:   "Upload a tax form attachment for a recipient. The file is uploaded via\nmultipart/form-data. Supported file types include PDF, images (PNG, JPG, GIF),\nand common document formats. The attachment will be associated as a tax document\nfor the recipient.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "recipient-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "file",
			Usage:     "The file to upload (tax form document)",
			Required:  true,
			BodyPath:  "file",
			FileInput: true,
		},
	},
	Action:          handleRecipientsAttachmentsAttach,
	HideHelpCommand: true,
}

func handleRecipientsAttachmentsList(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.RecipientAttachmentListParams{}

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
		_, err = client.Recipients.Attachments.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "recipients:attachments list", obj, format, transform)
	} else {
		iter := client.Recipients.Attachments.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "recipients:attachments list", iter, format, transform, maxItems)
	}
}

func handleRecipientsAttachmentsAttach(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("recipient-id") && len(unusedArgs) > 0 {
		cmd.Set("recipient-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.RecipientAttachmentAttachParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		MultipartFormEncoded,
		false,
	)
	if err != nil {
		return err
	}

	return client.Recipients.Attachments.Attach(
		ctx,
		cmd.Value("recipient-id").(string),
		params,
		options...,
	)
}
