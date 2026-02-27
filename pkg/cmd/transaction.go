// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/MercuryTechnologies/mercury-cli/internal/apiquery"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
	"github.com/stainless-sdks/mercury-go"
	"github.com/urfave/cli/v3"
)

var transactionsUploadAttachment = cli.Command{
	Name:    "upload-attachment",
	Usage:   "Upload a file attachment to a transaction. The file is uploaded via\nmultipart/form-data. Supported file types include PDF, images (PNG, JPG, GIF),\nand common document formats.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "transaction-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "file",
			Usage:    "The file to upload",
			Required: true,
			BodyPath: "file",
		},
		&requestflag.Flag[string]{
			Name:     "attachment-type",
			Usage:    "Type of attachment: 'receipt', 'bill', or 'other'. Defaults to 'other'.",
			BodyPath: "attachmentType",
		},
	},
	Action:          handleTransactionsUploadAttachment,
	HideHelpCommand: true,
}

func handleTransactionsUploadAttachment(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("transaction-id") && len(unusedArgs) > 0 {
		cmd.Set("transaction-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.TransactionUploadAttachmentParams{}

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

	return client.Transactions.UploadAttachment(
		ctx,
		cmd.Value("transaction-id").(string),
		params,
		options...,
	)
}
