// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/MercuryTechnologies/mercury-cli/internal/apiquery"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
	"github.com/stainless-sdks/mercury-go"
	"github.com/stainless-sdks/mercury-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var transferCreate = cli.Command{
	Name:    "create",
	Usage:   "Transfer funds between two accounts within the same organization. Supports\ntransfers between depository accounts (checking/savings), from a depository\naccount to a treasury/investment account, and from a treasury/investment account\nto a depository account. Creates paired debit and credit transactions.",
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
		&requestflag.Flag[any]{
			Name:     "note",
			BodyPath: "note",
		},
	},
	Action:          handleTransferCreate,
	HideHelpCommand: true,
}

func handleTransferCreate(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.TransferNewParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Transfer.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "transfer create", obj, format, transform)
}
