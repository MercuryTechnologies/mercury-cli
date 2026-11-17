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

var creditList = cli.Command{
	Name:            "list",
	Usage:           "Retrieve a list of all credit accounts for the organization.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleCreditList,
	HideHelpCommand: true,
}

var creditRetrieveStatements = cli.Command{
	Name:    "retrieve-statements",
	Usage:   "Retrieve historical IO credit statements with cursor-based pagination. Each\nstatement includes its billing period (startDate and endDate), closing balance\n(endingBalance), and opening balance (details.accountDetails.startingBalance),\ncalculated using the same strict posting-date boundaries as the PDF. Balances\nuse signed account amounts: amounts owed are negative. Filter by statement start\ndate with start and end. Use getAccountStatements for depository accounts.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "ID for a Mercury account.",
			Required:  true,
			PathParam: "accountId",
		},
		&requestflag.Flag[string]{
			Name:      "end",
			Usage:     "Filter statements where the period start date is on or before this date. If the date is in the future, defaults to the current date. Format: YYYY-MM-DD",
			QueryPath: "end",
		},
		&requestflag.Flag[string]{
			Name:      "end-before",
			Usage:     "The ID of the statement to end the page before (exclusive). When provided, results will end just before this ID and work backwards. Use this for reverse pagination or to retrieve previous pages. Cannot be combined with start_after.",
			QueryPath: "end_before",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results to return. Allowed range: 1 to 100. Defaults to 25",
			Default:   25,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "order",
			Usage:     "Sort order. Can be 'asc' or 'desc'. Defaults to 'desc'",
			Default:   "desc",
			QueryPath: "order",
		},
		&requestflag.Flag[string]{
			Name:      "start",
			Usage:     "Filter statements where the period start date is on or after this date. Format: YYYY-MM-DD",
			QueryPath: "start",
		},
		&requestflag.Flag[string]{
			Name:      "start-after",
			Usage:     "The ID of the statement to start the page after (exclusive). When provided, results will begin with the statement immediately following this ID. Use this for standard forward pagination to get the next page of results. Cannot be combined with end_before.",
			QueryPath: "start_after",
		},
	},
	Action:          handleCreditRetrieveStatements,
	HideHelpCommand: true,
}

func handleCreditList(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Credit.List(ctx, options...)
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
		Title:          "credit list",
		Transform:      transform,
	})
}

func handleCreditRetrieveStatements(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("account-id") && len(unusedArgs) > 0 {
		cmd.Set("account-id", unusedArgs[0])
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

	params := mercury.CreditGetStatementsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Credit.GetStatements(
		ctx,
		cmd.Value("account-id").(string),
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
		Title:          "credit retrieve-statements",
		Transform:      transform,
	})
}
