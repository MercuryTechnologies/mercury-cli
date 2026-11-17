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

var requestTransferRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a specific transfer approval request by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "request-id",
			Usage:     "ID for the transfer money approval request",
			Required:  true,
			PathParam: "requestId",
		},
	},
	Action:          handleRequestTransferRetrieve,
	HideHelpCommand: true,
}

var requestTransferRequestTransfer = cli.Command{
	Name:    "request-transfer",
	Usage:   "Create a transfer between two accounts in your organization that will require\napproval based on your organization's approval policies.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[float64]{
			Name:     "amount",
			Usage:    "A positive dollar amount with at least 1 cent.",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "destination-account-id",
			Usage:    "ID for a Mercury account.",
			Required: true,
			BodyPath: "destinationAccountId",
		},
		&requestflag.Flag[string]{
			Name:     "idempotency-key",
			Required: true,
			BodyPath: "idempotencyKey",
		},
		&requestflag.Flag[string]{
			Name:     "source-account-id",
			Usage:    "ID for a Mercury account.",
			Required: true,
			BodyPath: "sourceAccountId",
		},
		&requestflag.Flag[*string]{
			Name:     "note",
			BodyPath: "note",
		},
	},
	Action:          handleRequestTransferRequestTransfer,
	HideHelpCommand: true,
}

var requestTransferRetrieveRequestTransfer = cli.Command{
	Name:    "retrieve-request-transfer",
	Usage:   "Retrieve a paginated list of transfer approval requests for the authenticated\norganization. Supports filtering by source account, destination account, and\nstatus.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "destination-account-id",
			Usage:     "ID for a Mercury account.",
			QueryPath: "destinationAccountId",
		},
		&requestflag.Flag[string]{
			Name:      "end-before",
			Usage:     "The ID of the transfer approval request to end the page before (exclusive). When provided, results will end just before this ID and work backwards. Use this for reverse pagination or to retrieve previous pages. Cannot be combined with start_after.",
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
			Name:      "source-account-id",
			Usage:     "ID for a Mercury account.",
			QueryPath: "sourceAccountId",
		},
		&requestflag.Flag[string]{
			Name:      "start-after",
			Usage:     "The ID of the transfer approval request to start the page after (exclusive). When provided, results will begin with the transfer approval request immediately following this ID. Use this for standard forward pagination to get the next page of results. Cannot be combined with end_before.",
			QueryPath: "start_after",
		},
		&requestflag.Flag[string]{
			Name:      "start-at",
			Usage:     "The ID of the transfer approval request to start the page at (inclusive). Cannot be combined with start_after or end_before.",
			QueryPath: "start_at",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     `Allowed values: "pendingApproval", "approved", "rejected", "cancelled".`,
			QueryPath: "status",
		},
	},
	Action:          handleRequestTransferRetrieveRequestTransfer,
	HideHelpCommand: true,
}

func handleRequestTransferRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("request-id") && len(unusedArgs) > 0 {
		cmd.Set("request-id", unusedArgs[0])
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
	_, err = client.RequestTransfer.Get(ctx, cmd.Value("request-id").(string), options...)
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
		Title:          "request-transfer retrieve",
		Transform:      transform,
	})
}

func handleRequestTransferRequestTransfer(ctx context.Context, cmd *cli.Command) error {
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

	params := mercury.RequestTransferRequestTransferParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.RequestTransfer.RequestTransfer(ctx, params, options...)
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
		Title:          "request-transfer request-transfer",
		Transform:      transform,
	})
}

func handleRequestTransferRetrieveRequestTransfer(ctx context.Context, cmd *cli.Command) error {
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

	params := mercury.RequestTransferGetRequestTransferParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.RequestTransfer.GetRequestTransfer(ctx, params, options...)
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
		Title:          "request-transfer retrieve-request-transfer",
		Transform:      transform,
	})
}
