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

var recipientsUploadAttachment = cli.Command{
	Name:    "upload-attachment",
	Usage:   "Upload a tax form attachment for a recipient. The file is uploaded via\nmultipart/form-data. Supported file types include PDF, images (PNG, JPG, GIF),\nand common document formats. The attachment will be associated as a tax document\nfor the recipient.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "recipient-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "file",
			Usage:    "The file to upload (tax form document)",
			Required: true,
			BodyPath: "file",
		},
	},
	Action:          handleRecipientsUploadAttachment,
	HideHelpCommand: true,
}

func handleRecipientsUploadAttachment(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("recipient-id") && len(unusedArgs) > 0 {
		cmd.Set("recipient-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.RecipientUploadAttachmentParams{}

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

	return client.Recipients.UploadAttachment(
		ctx,
		cmd.Value("recipient-id").(string),
		params,
		options...,
	)
}
