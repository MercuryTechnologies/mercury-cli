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

var recipientsInvitesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create an invite for a recipient to submit their payment details. Supply a\nrecipientId to invite an existing recipient; omit it to invite someone new, in\nwhich case the recipient is created when the invitee completes onboarding.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "contact-email",
			Usage:    " Contact email the invite is sent to. When 'recipientId' is present, updates the recipient's contact email to this value.",
			Required: true,
			BodyPath: "contactEmail",
		},
		&requestflag.Flag[[]string]{
			Name:     "payment-method",
			Usage:    " Payment methods the recipient may submit details for.",
			Required: true,
			BodyPath: "paymentMethods",
		},
		&requestflag.Flag[bool]{
			Name:     "require-tax-document",
			Usage:    " Whether the recipient must upload a tax document.",
			Required: true,
			BodyPath: "requireTaxDocument",
		},
		&requestflag.Flag[bool]{
			Name:     "send-email",
			Usage:    " When true, sends an Email to the invitee. When false, does not send an email to the invitee.",
			Required: true,
			BodyPath: "sendEmail",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    " Name the invite is created for. This field is required when 'recipientId' is absent.\n When 'recipientId' is present, this field is optional and updates the recipient's name to this value.",
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "notes",
			Usage:    " Optional notes shown to the recipient.",
			BodyPath: "notes",
		},
		&requestflag.Flag[*string]{
			Name:     "organization-name-on-request",
			Usage:    " Optional organization name to display on the request.",
			BodyPath: "organizationNameOnRequest",
		},
		&requestflag.Flag[*string]{
			Name:     "recipient-id",
			Usage:    "ID for a Mercury account.",
			BodyPath: "recipientId",
		},
	},
	Action:          handleRecipientsInvitesCreate,
	HideHelpCommand: true,
}

var recipientsInvitesList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a paginated list of all recipient invites for your organization.\nSupports filtering by status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "end-before",
			Usage:     "The ID of the recipient invite to end the page before (exclusive). When provided, results will end just before this ID and work backwards. Use this for reverse pagination or to retrieve previous pages. Cannot be combined with start_after.",
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
			Usage:     "The ID of the recipient invite to start the page after (exclusive). When provided, results will begin with the recipient invite immediately following this ID. Use this for standard forward pagination to get the next page of results. Cannot be combined with end_before.",
			QueryPath: "start_after",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     `Allowed values: "created", "completed", "expired".`,
			QueryPath: "status",
		},
	},
	Action:          handleRecipientsInvitesList,
	HideHelpCommand: true,
}

var recipientsInvitesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an active recipient invite.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "invite-id",
			Usage:     "ID for the invite",
			Required:  true,
			PathParam: "inviteId",
		},
	},
	Action:          handleRecipientsInvitesDelete,
	HideHelpCommand: true,
}

func handleRecipientsInvitesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := mercury.RecipientInviteNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Recipients.Invites.New(ctx, params, options...)
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
		Title:          "recipients:invites create",
		Transform:      transform,
	})
}

func handleRecipientsInvitesList(ctx context.Context, cmd *cli.Command) error {
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

	params := mercury.RecipientInviteListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Recipients.Invites.List(ctx, params, options...)
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
		Title:          "recipients:invites list",
		Transform:      transform,
	})
}

func handleRecipientsInvitesDelete(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("invite-id") && len(unusedArgs) > 0 {
		cmd.Set("invite-id", unusedArgs[0])
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

	return client.Recipients.Invites.Delete(ctx, cmd.Value("invite-id").(string), options...)
}
