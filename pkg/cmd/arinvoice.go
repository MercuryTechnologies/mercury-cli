// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/mercury-cli/internal/apiquery"
	"github.com/stainless-sdks/mercury-cli/internal/requestflag"
	"github.com/stainless-sdks/mercury-go"
	"github.com/urfave/cli/v3"
)

var arInvoicesCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancel an invoice. This action cannot be undone.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "invoice-id",
			Usage:    "ID for the invoice.",
			Required: true,
		},
	},
	Action:          handleArInvoicesCancel,
	HideHelpCommand: true,
}

func handleArInvoicesCancel(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("invoice-id") && len(unusedArgs) > 0 {
		cmd.Set("invoice-id", unusedArgs[0])
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

	return client.Ar.Invoices.Cancel(ctx, cmd.Value("invoice-id").(string), options...)
}
