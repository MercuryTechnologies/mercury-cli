// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/MercuryTechnologies/mercury-cli/internal/apiquery"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
	"github.com/MercuryTechnologies/mercury-go"
	"github.com/MercuryTechnologies/mercury-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var recipientsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a new recipient for making payments",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "email",
			Required: true,
			BodyPath: "emails",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "address",
			BodyPath: "address",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "check-info",
			BodyPath: "checkInfo",
		},
		&requestflag.Flag[string]{
			Name:     "contact-email",
			Usage:    "Contact email address of the recipient",
			BodyPath: "contactEmail",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "domestic-wire-routing-info",
			BodyPath: "domesticWireRoutingInfo",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "electronic-routing-info",
			BodyPath: "electronicRoutingInfo",
		},
		&requestflag.Flag[string]{
			Name:     "nickname",
			BodyPath: "nickname",
		},
	},
	Action:          handleRecipientsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"address": {
		&requestflag.InnerFlag[string]{
			Name:       "address.address1",
			InnerField: "address1",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.city",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.postal-code",
			InnerField: "postalCode",
		},
		&requestflag.InnerFlag[any]{
			Name:       "address.address2",
			InnerField: "address2",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.state",
			Usage:      `Allowed values: "AL", "AK", "AZ", "AR", "CA", "CO", "CT", "DE", "DC", "FL", "GA", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "OH", "OK", "OR", "PA", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "WA", "WV", "WI", "WY".`,
			InnerField: "state",
		},
	},
	"check-info": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "check-info.address",
			InnerField: "address",
		},
	},
	"domestic-wire-routing-info": {
		&requestflag.InnerFlag[string]{
			Name:       "domestic-wire-routing-info.account-number",
			Usage:      " The account number of the bank account to use for domestic wire payments.",
			InnerField: "accountNumber",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "domestic-wire-routing-info.address",
			InnerField: "address",
		},
		&requestflag.InnerFlag[string]{
			Name:       "domestic-wire-routing-info.routing-number",
			Usage:      " The routing number of the bank account to use for domestic wire payments.",
			InnerField: "routingNumber",
		},
		&requestflag.InnerFlag[any]{
			Name:       "domestic-wire-routing-info.default-for-benefit-of",
			Usage:      " The name of the beneficiary of the domestic wire. This is the name of the entity that will receive the domestic wire.",
			InnerField: "defaultForBenefitOf",
		},
	},
	"electronic-routing-info": {
		&requestflag.InnerFlag[string]{
			Name:       "electronic-routing-info.account-number",
			Usage:      " The account number of the bank account to use for ACH payments.",
			InnerField: "accountNumber",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "electronic-routing-info.address",
			InnerField: "address",
		},
		&requestflag.InnerFlag[string]{
			Name:       "electronic-routing-info.electronic-account-type",
			Usage:      `Allowed values: "businessChecking", "businessSavings", "personalChecking", "personalSavings".`,
			InnerField: "electronicAccountType",
		},
		&requestflag.InnerFlag[string]{
			Name:       "electronic-routing-info.routing-number",
			Usage:      " The routing number of the bank account to use for ACH payments.",
			InnerField: "routingNumber",
		},
	},
})

var recipientsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve details of a specific recipient by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "recipient-id",
			Usage:    "ID for a Mercury account.",
			Required: true,
		},
	},
	Action:          handleRecipientsRetrieve,
	HideHelpCommand: true,
}

var recipientsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update an existing recipient's information",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "recipient-id",
			Usage:    "ID for a Mercury account.",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "address",
			BodyPath: "address",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "check-info",
			BodyPath: "checkInfo",
		},
		&requestflag.Flag[string]{
			Name:     "contact-email",
			Usage:    "Contact email address of the recipient",
			BodyPath: "contactEmail",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "domestic-wire-routing-info",
			BodyPath: "domesticWireRoutingInfo",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "electronic-routing-info",
			BodyPath: "electronicRoutingInfo",
		},
		&requestflag.Flag[[]string]{
			Name:     "email",
			BodyPath: "emails",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "nickname",
			BodyPath: "nickname",
		},
	},
	Action:          handleRecipientsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"address": {
		&requestflag.InnerFlag[string]{
			Name:       "address.address1",
			InnerField: "address1",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.city",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.postal-code",
			InnerField: "postalCode",
		},
		&requestflag.InnerFlag[any]{
			Name:       "address.address2",
			InnerField: "address2",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.state",
			Usage:      `Allowed values: "AL", "AK", "AZ", "AR", "CA", "CO", "CT", "DE", "DC", "FL", "GA", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "OH", "OK", "OR", "PA", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "WA", "WV", "WI", "WY".`,
			InnerField: "state",
		},
	},
	"check-info": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "check-info.address",
			InnerField: "address",
		},
	},
	"domestic-wire-routing-info": {
		&requestflag.InnerFlag[string]{
			Name:       "domestic-wire-routing-info.account-number",
			Usage:      " The account number of the bank account to use for domestic wire payments.",
			InnerField: "accountNumber",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "domestic-wire-routing-info.address",
			InnerField: "address",
		},
		&requestflag.InnerFlag[string]{
			Name:       "domestic-wire-routing-info.routing-number",
			Usage:      " The routing number of the bank account to use for domestic wire payments.",
			InnerField: "routingNumber",
		},
		&requestflag.InnerFlag[any]{
			Name:       "domestic-wire-routing-info.default-for-benefit-of",
			Usage:      " The name of the beneficiary of the domestic wire. This is the name of the entity that will receive the domestic wire.",
			InnerField: "defaultForBenefitOf",
		},
	},
	"electronic-routing-info": {
		&requestflag.InnerFlag[string]{
			Name:       "electronic-routing-info.account-number",
			Usage:      " The account number of the bank account to use for ACH payments.",
			InnerField: "accountNumber",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "electronic-routing-info.address",
			InnerField: "address",
		},
		&requestflag.InnerFlag[string]{
			Name:       "electronic-routing-info.electronic-account-type",
			Usage:      `Allowed values: "businessChecking", "businessSavings", "personalChecking", "personalSavings".`,
			InnerField: "electronicAccountType",
		},
		&requestflag.InnerFlag[string]{
			Name:       "electronic-routing-info.routing-number",
			Usage:      " The routing number of the bank account to use for ACH payments.",
			InnerField: "routingNumber",
		},
	},
})

var recipientsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a paginated list of all recipients. Use cursor parameters (start_after,\nend_before) for pagination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "end-before",
			Usage:     "The ID of the recipient to end the page before (exclusive). When provided, results will end just before this ID and work backwards. Use this for reverse pagination or to retrieve previous pages. Cannot be combined with start_after.",
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
			Usage:     "The ID of the recipient to start the page after (exclusive). When provided, results will begin with the recipient immediately following this ID. Use this for standard forward pagination to get the next page of results. Cannot be combined with end_before.",
			QueryPath: "start_after",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleRecipientsList,
	HideHelpCommand: true,
}

var recipientsListAttachments = cli.Command{
	Name:    "list-attachments",
	Usage:   "Retrieve a paginated list of all recipient tax form attachments across all\nrecipients in the organization. Use cursor parameters (start_after, end_before)\nfor pagination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "end-before",
			Usage:     "The ID of the recipient attachment to end the page before (exclusive). When provided, results will end just before this ID and work backwards. Use this for reverse pagination or to retrieve previous pages. Cannot be combined with start_after.",
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
			Usage:     "The ID of the recipient attachment to start the page after (exclusive). When provided, results will begin with the recipient attachment immediately following this ID. Use this for standard forward pagination to get the next page of results. Cannot be combined with end_before.",
			QueryPath: "start_after",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleRecipientsListAttachments,
	HideHelpCommand: true,
}

var recipientsUploadAttachment = cli.Command{
	Name:    "upload-attachment",
	Usage:   "Upload a tax form attachment for a recipient. The file is uploaded via\nmultipart/form-data. Supported file types include PDF, images (PNG, JPG, GIF),\nand common document formats. The attachment will be associated as a tax document\nfor the recipient.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "recipient-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "file",
			Usage:    "The file to upload (tax form document)",
			Required: true,
			BodyPath: "file",
		},
	},
	Action:          handleRecipientsUploadAttachment,
	HideHelpCommand: true,
}

func handleRecipientsCreate(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.RecipientNewParams{}

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
	_, err = client.Recipients.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "recipients create", obj, format, transform)
}

func handleRecipientsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("recipient-id") && len(unusedArgs) > 0 {
		cmd.Set("recipient-id", unusedArgs[0])
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
	_, err = client.Recipients.Get(ctx, cmd.Value("recipient-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "recipients retrieve", obj, format, transform)
}

func handleRecipientsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("recipient-id") && len(unusedArgs) > 0 {
		cmd.Set("recipient-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.RecipientUpdateParams{}

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
	_, err = client.Recipients.Update(
		ctx,
		cmd.Value("recipient-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "recipients update", obj, format, transform)
}

func handleRecipientsList(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.RecipientListParams{}

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
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Recipients.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "recipients list", obj, format, transform)
	} else {
		iter := client.Recipients.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "recipients list", iter, format, transform, maxItems)
	}
}

func handleRecipientsListAttachments(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.RecipientListAttachmentsParams{}

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
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Recipients.ListAttachments(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "recipients list-attachments", obj, format, transform)
	} else {
		iter := client.Recipients.ListAttachmentsAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "recipients list-attachments", iter, format, transform, maxItems)
	}
}

func handleRecipientsUploadAttachment(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("recipient-id") && len(unusedArgs) > 0 {
		cmd.Set("recipient-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.RecipientUploadAttachmentParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		MultipartFormEncoded,
		false,
	)
	if err != nil {
		return err
	}

	return client.Recipients.UploadAttachment(
		ctx,
		cmd.Value("recipient-id").(string),
		params,
		options...,
	)
}
