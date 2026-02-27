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

var webhooksDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a webhook endpoint",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "webhook-endpoint-id",
			Usage:    "ID for the webhook",
			Required: true,
		},
	},
	Action:          handleWebhooksDelete,
	HideHelpCommand: true,
}

var webhooksVerify = cli.Command{
	Name:    "verify",
	Usage:   "Send a test event to verify the webhook endpoint is properly configured and\nreachable. The request body accepts an optional 'eventType' field to specify\nwhich event type to test (e.g., 'transaction.created', 'transaction.updated').\nIf omitted from the request body, defaults to 'transaction.created'.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "webhook-endpoint-id",
			Usage:    "ID for the webhook",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "event-type",
			Usage:    " Optional event type to test. If not specified, defaults to transaction.created.",
			BodyPath: "eventType",
		},
	},
	Action:          handleWebhooksVerify,
	HideHelpCommand: true,
}

func handleWebhooksDelete(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("webhook-endpoint-id") && len(unusedArgs) > 0 {
		cmd.Set("webhook-endpoint-id", unusedArgs[0])
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

	return client.Webhooks.Delete(ctx, cmd.Value("webhook-endpoint-id").(string), options...)
}

func handleWebhooksVerify(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("webhook-endpoint-id") && len(unusedArgs) > 0 {
		cmd.Set("webhook-endpoint-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.WebhookVerifyParams{}

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

	return client.Webhooks.Verify(
		ctx,
		cmd.Value("webhook-endpoint-id").(string),
		params,
		options...,
	)
}
