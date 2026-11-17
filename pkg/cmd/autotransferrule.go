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

var autoTransferRulesRetrieveAutoTransferRules = cli.Command{
	Name:    "retrieve-auto-transfer-rules",
	Usage:   "Retrieve a paginated list of the organization's active auto-transfer rules:\nrules that distribute fixed amounts or percentages of an account's balance to\nother accounts, and rules that keep an account at a target balance. Pass\naccountId to return only the rules that manage that account's balance.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "account-id",
			Usage:     "ID for a Mercury account.",
			QueryPath: "accountId",
		},
		&requestflag.Flag[string]{
			Name:      "end-before",
			Usage:     "The ID of the auto-transfer rule to end the page before (exclusive). When provided, results will end just before this ID and work backwards. Use this for reverse pagination or to retrieve previous pages. Cannot be combined with start_after.",
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
			Usage:     "The ID of the auto-transfer rule to start the page after (exclusive). When provided, results will begin with the auto-transfer rule immediately following this ID. Use this for standard forward pagination to get the next page of results. Cannot be combined with end_before.",
			QueryPath: "start_after",
		},
	},
	Action:          handleAutoTransferRulesRetrieveAutoTransferRules,
	HideHelpCommand: true,
}

func handleAutoTransferRulesRetrieveAutoTransferRules(ctx context.Context, cmd *cli.Command) error {
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

	params := mercury.AutoTransferRuleGetAutoTransferRulesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.AutoTransferRules.GetAutoTransferRules(ctx, params, options...)
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
		Title:          "auto-transfer-rules retrieve-auto-transfer-rules",
		Transform:      transform,
	})
}
