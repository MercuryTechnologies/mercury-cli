// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/MercuryTechnologies/mercury-cli/internal/apiquery"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
	"github.com/MercuryTechnologies/mercury-go"
	"github.com/MercuryTechnologies/mercury-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var webhooksCreate = cli.Command{
	Name:    "create",
	Usage:   "Register a new webhook endpoint to receive event notifications",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    " The URL to which webhook events will be delivered",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[any]{
			Name:     "event-type",
			Usage:    " Optional array of event types to subscribe to. Nothing means subscribe to all event types.",
			BodyPath: "eventTypes",
		},
		&requestflag.Flag[any]{
			Name:     "filter-path",
			Usage:    " Optional array of resource field paths to filter events by. When specified, webhook events will only be sent when one of these fields changes. Nothing means no filtering (all events are sent).",
			BodyPath: "filterPaths",
		},
	},
	Action:          handleWebhooksCreate,
	HideHelpCommand: true,
}

var webhooksUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update the configuration of an existing webhook endpoint. A webhook that has\nbeen disabled due to consecutive delivery failures can be reactivated by setting\nits status to 'active'.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "webhook-endpoint-id",
			Usage:    "ID for the webhook",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "event-type",
			Usage:    " Event types to subscribe to. Send null to subscribe to all event types. Send an array to subscribe to specific types. Omit to leave unchanged.",
			BodyPath: "eventTypes",
		},
		&requestflag.Flag[any]{
			Name:     "filter-path",
			Usage:    " Resource field paths to filter events by. When specified, webhook events will only be sent when one of these fields changes. Send null for no filtering. Send an array to filter by specific fields. Omit to leave unchanged.",
			BodyPath: "filterPaths",
		},
		&requestflag.Flag[any]{
			Name:     "status",
			Usage:    " Webhook status. Only 'active' and 'paused' values are allowed. Omit to leave unchanged.",
			BodyPath: "status",
		},
		&requestflag.Flag[any]{
			Name:     "url",
			Usage:    " The URL to which webhook events will be delivered. Omit to leave unchanged.",
			BodyPath: "url",
		},
	},
	Action:          handleWebhooksUpdate,
	HideHelpCommand: true,
}

var webhooksList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a paginated list of all webhook endpoints for your organization.\nSupports filtering by status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "end-before",
			Usage:     "The ID of the webhook to end the page before (exclusive). When provided, results will end just before this ID and work backwards. Use this for reverse pagination or to retrieve previous pages. Cannot be combined with start_after.",
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
			Usage:     "The ID of the webhook to start the page after (exclusive). When provided, results will begin with the webhook immediately following this ID. Use this for standard forward pagination to get the next page of results. Cannot be combined with end_before.",
			QueryPath: "start_after",
		},
		&requestflag.Flag[[]string]{
			Name:      "status",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleWebhooksList,
	HideHelpCommand: true,
}

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

var webhooksGet = cli.Command{
	Name:    "get",
	Usage:   "Retrieve details of a specific webhook endpoint by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "webhook-endpoint-id",
			Usage:    "ID for the webhook",
			Required: true,
		},
	},
	Action:          handleWebhooksGet,
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

func handleWebhooksCreate(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.WebhookNewParams{}

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
	_, err = client.Webhooks.New(ctx, params, options...)
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
		Title:          "webhooks create",
		Transform:      transform,
	})
}

func handleWebhooksUpdate(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("webhook-endpoint-id") && len(unusedArgs) > 0 {
		cmd.Set("webhook-endpoint-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.WebhookUpdateParams{}

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
	_, err = client.Webhooks.Update(
		ctx,
		cmd.Value("webhook-endpoint-id").(string),
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
		Title:          "webhooks update",
		Transform:      transform,
	})
}

func handleWebhooksList(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.WebhookListParams{}

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

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Webhooks.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "webhooks list",
			Transform:      transform,
		})
	} else {
		iter := client.Webhooks.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "webhooks list",
			Transform:      transform,
		})
	}
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

	// CUSTOM: fetch webhook details and confirm before deleting
	webhookID := cmd.Value("webhook-endpoint-id").(string)
	var res []byte
	_, err = client.Webhooks.Get(ctx, webhookID, option.WithResponseBodyInto(&res))
	if err != nil {
		return err
	}
	obj := gjson.ParseBytes(res)
	events := "all"
	if eventTypes := obj.Get("eventTypes"); eventTypes.Exists() && eventTypes.IsArray() {
		var types []string
		for _, et := range eventTypes.Array() {
			types = append(types, et.String())
		}
		if len(types) > 0 {
			events = strings.Join(types, ", ")
		}
	}
	details := []ConfirmDetail{
		{Label: "URL", Value: obj.Get("url").String()},
		{Label: "Events", Value: events},
		{Label: "Status", Value: obj.Get("status").String()},
	}
	if err := confirmAction(cmd, "Delete Webhook", details); err != nil {
		return err
	}

	return client.Webhooks.Delete(ctx, webhookID, options...)
}

func handleWebhooksGet(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Webhooks.Get(ctx, cmd.Value("webhook-endpoint-id").(string), options...)
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
		Title:          "webhooks get",
		Transform:      transform,
	})
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
