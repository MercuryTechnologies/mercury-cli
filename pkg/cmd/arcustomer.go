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

var arCustomersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a customer. This action cannot be undone.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "customer-id",
			Usage:    "The customer who will receive the invoice. Use the /api/v1/ar/customers endpoint to list your customers and find the corresponding id, or create a new customer first.",
			Required: true,
		},
	},
	Action:          handleArCustomersDelete,
	HideHelpCommand: true,
}

func handleArCustomersDelete(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("customer-id") && len(unusedArgs) > 0 {
		cmd.Set("customer-id", unusedArgs[0])
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

	return client.Ar.Customers.Delete(ctx, cmd.Value("customer-id").(string), options...)
}
