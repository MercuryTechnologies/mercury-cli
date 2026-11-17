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

var merchantsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a paginated list of priority merchants that can be used for spend\ncontrols like merchant locking. Supports cursor-based pagination with limit,\norder, start_after, and end_before query parameters, plus an optional\ncase-insensitive search by merchant name.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "end-before",
			Usage:     "The ID of the merchant to end the page before (exclusive). When provided, results will end just before this ID and work backwards. Use this for reverse pagination or to retrieve previous pages. Cannot be combined with start_after.",
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
			Name:      "search",
			Usage:     "Case-insensitive search term to filter merchants by name.",
			QueryPath: "search",
		},
		&requestflag.Flag[string]{
			Name:      "start-after",
			Usage:     "The ID of the merchant to start the page after (exclusive). When provided, results will begin with the merchant immediately following this ID. Use this for standard forward pagination to get the next page of results. Cannot be combined with end_before.",
			QueryPath: "start_after",
		},
	},
	Action:          handleMerchantsList,
	HideHelpCommand: true,
}

func handleMerchantsList(ctx context.Context, cmd *cli.Command) error {
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

	params := mercury.MerchantListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Merchants.List(ctx, params, options...)
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
		Title:          "merchants list",
		Transform:      transform,
	})
}
