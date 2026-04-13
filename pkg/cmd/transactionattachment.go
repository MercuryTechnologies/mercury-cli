// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/MercuryTechnologies/mercury-cli/internal/apiquery"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
	"github.com/MercuryTechnologies/mercury-go"
	"github.com/urfave/cli/v3"
)

var transactionsAttachmentsAttach = cli.Command{
	Name:    "attach",
	Usage:   "Upload a file attachment to a transaction. The file is uploaded via\nmultipart/form-data. Supported file types include PDF, images (PNG, JPG, GIF),\nand common document formats.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "transaction-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "file",
			Usage:     "The file to upload",
			Required:  true,
			BodyPath:  "file",
			FileInput: true,
		},
		&requestflag.Flag[string]{
			Name:     "attachment-type",
			Usage:    "Type of attachment: 'receipt', 'bill', or 'other'. Defaults to 'other'.",
			BodyPath: "attachmentType",
		},
	},
	Action:          handleTransactionsAttachmentsAttach,
	HideHelpCommand: true,
}

func handleTransactionsAttachmentsAttach(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("transaction-id") && len(unusedArgs) > 0 {
		cmd.Set("transaction-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.TransactionAttachmentAttachParams{}

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

	return client.Transactions.Attachments.Attach(
		ctx,
		cmd.Value("transaction-id").(string),
		params,
		options...,
	)
}
