// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/MercuryTechnologies/mercury-cli/internal/apiquery"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
	"github.com/MercuryTechnologies/mercury-go"
	"github.com/urfave/cli/v3"
)

var statementsDownloadPdf = cli.Command{
	Name:    "download-pdf",
	Usage:   "Downloads a PDF file for the specified account statement. The response includes\na Content-Disposition header for proper file download handling. Returns binary\nPDF data.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "statement-id",
			Usage:    "ID for the account statement",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "The file where the response contents will be stored. Use the value '-' to force output to stdout.",
		},
	},
	Action:          handleStatementsDownloadPdf,
	HideHelpCommand: true,
}

func handleStatementsDownloadPdf(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("statement-id") && len(unusedArgs) > 0 {
		cmd.Set("statement-id", unusedArgs[0])
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

	response, err := client.Statements.DownloadPdf(ctx, cmd.Value("statement-id").(string), options...)
	if err != nil {
		return err
	}
	message, err := writeBinaryResponse(response, os.Stdout, cmd.String("output"))
	if message != "" {
		fmt.Println(message)
	}
	return err
}
