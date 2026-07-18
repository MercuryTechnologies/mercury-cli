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

var cardsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Issue a new virtual card.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "kind",
			Usage:    `Allowed values: "debit", "credit".`,
			Required: true,
			BodyPath: "kind",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    " The type of card to issue.",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "user-id",
			Usage:    "ID for the user",
			Required: true,
			BodyPath: "userId",
		},
		&requestflag.Flag[*string]{
			Name:     "account-id",
			Usage:    "ID for a Mercury account.",
			BodyPath: "accountId",
		},
		&requestflag.Flag[*string]{
			Name:     "nickname",
			Usage:    " Optional user-assigned label for the card.",
			BodyPath: "nickname",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "spend-limit",
			Usage:    " Spending controls to apply at issuance.",
			BodyPath: "spendLimit",
		},
	},
	Action:          handleCardsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"spend-limit": {
		&requestflag.InnerFlag[int64]{
			Name:       "spend-limit.amount-cents",
			Usage:      " Maximum total spend allowed per interval, in cents.",
			InnerField: "amountCents",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spend-limit.interval",
			Usage:      " Rolling window the limit applies to.",
			InnerField: "interval",
		},
	},
})

var cardsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update a card's nickname or spending limits.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "card-id",
			Usage:     "Unique identifier for a card",
			Required:  true,
			PathParam: "cardId",
		},
		&requestflag.Flag[*string]{
			Name:     "nickname",
			Usage:    "Nickname update action. Omit the field to keep the current nickname, send null or an empty/whitespace-only string to clear it, or send a string to set it.",
			Required: true,
			BodyPath: "nickname",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "spend-limit",
			Usage:    " Spending controls applied to a card",
			BodyPath: "spendLimit",
		},
	},
	Action:          handleCardsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"spend-limit": {
		&requestflag.InnerFlag[int64]{
			Name:       "spend-limit.amount-cents",
			Usage:      " Maximum total spend allowed per interval, in cents.",
			InnerField: "amountCents",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spend-limit.interval",
			Usage:      " Rolling window the limit applies to.",
			InnerField: "interval",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "spend-limit.atm-amount-cents",
			Usage:      " Maximum ATM withdrawal allowed per interval, in cents. Null for virtual cards.",
			InnerField: "atmAmountCents",
		},
	},
})

var cardsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a paginated list of cards.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:      "account-id",
			Usage:     "Filter cards by one or more account IDs.",
			QueryPath: "accountId",
		},
		&requestflag.Flag[string]{
			Name:      "end-before",
			Usage:     "The ID of the card to end the page before (exclusive). When provided, results will end just before this ID and work backwards. Use this for reverse pagination or to retrieve previous pages. Cannot be combined with start_after.",
			QueryPath: "end_before",
		},
		&requestflag.Flag[[]string]{
			Name:      "kind",
			Usage:     "Filter cards by kind (debit or credit).",
			QueryPath: "kind",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 500",
			Default:   500,
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
			Usage:     "The ID of the card to start the page after (exclusive). When provided, results will begin with the card immediately following this ID. Use this for standard forward pagination to get the next page of results. Cannot be combined with end_before.",
			QueryPath: "start_after",
		},
		&requestflag.Flag[[]string]{
			Name:      "status",
			Usage:     "Filter cards by one or more statuses.",
			QueryPath: "status",
		},
		&requestflag.Flag[[]string]{
			Name:      "type",
			Usage:     "Filter cards by type (virtual or physical).",
			QueryPath: "type",
		},
		&requestflag.Flag[string]{
			Name:      "user-id",
			Usage:     "Filter cards by the cardholder's user ID.",
			QueryPath: "userId",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleCardsList,
	HideHelpCommand: true,
}

var cardsCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Permanently cancel a card. This action cannot be undone.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "card-id",
			Usage:     "Unique identifier for a card",
			Required:  true,
			PathParam: "cardId",
		},
	},
	Action:          handleCardsCancel,
	HideHelpCommand: true,
}

var cardsFreeze = cli.Command{
	Name:    "freeze",
	Usage:   "Temporarily freeze a card. The card must be active.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "card-id",
			Usage:     "Unique identifier for a card",
			Required:  true,
			PathParam: "cardId",
		},
	},
	Action:          handleCardsFreeze,
	HideHelpCommand: true,
}

var cardsGet = cli.Command{
	Name:    "get",
	Usage:   "Retrieve details of a specific card by its ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "card-id",
			Usage:     "Unique identifier for a card",
			Required:  true,
			PathParam: "cardId",
		},
	},
	Action:          handleCardsGet,
	HideHelpCommand: true,
}

var cardsReveal = cli.Command{
	Name:    "reveal",
	Usage:   "Retrieve the full card number, expiration date, and CVC for a card. Available\nfor agentic cards only.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "card-id",
			Usage:     "Unique identifier for a card",
			Required:  true,
			PathParam: "cardId",
		},
	},
	Action:          handleCardsReveal,
	HideHelpCommand: true,
}

var cardsUnfreeze = cli.Command{
	Name:    "unfreeze",
	Usage:   "Unfreeze a previously frozen card, restoring it to active status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "card-id",
			Usage:     "Unique identifier for a card",
			Required:  true,
			PathParam: "cardId",
		},
	},
	Action:          handleCardsUnfreeze,
	HideHelpCommand: true,
}

func handleCardsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := mercury.CardNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Cards.New(ctx, params, options...)
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
		Title:          "cards create",
		Transform:      transform,
	})
}

func handleCardsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-id") && len(unusedArgs) > 0 {
		cmd.Set("card-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	params := mercury.CardUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Cards.Update(
		ctx,
		cmd.Value("card-id").(string),
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
		Title:          "cards update",
		Transform:      transform,
	})
}

func handleCardsList(ctx context.Context, cmd *cli.Command) error {
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

	params := mercury.CardListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Cards.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "cards list",
			Transform:      transform,
		})
	} else {
		iter := client.Cards.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "cards list",
			Transform:      transform,
		})
	}
}

func handleCardsCancel(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-id") && len(unusedArgs) > 0 {
		cmd.Set("card-id", unusedArgs[0])
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
	_, err = client.Cards.Cancel(ctx, cmd.Value("card-id").(string), options...)
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
		Title:          "cards cancel",
		Transform:      transform,
	})
}

func handleCardsFreeze(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-id") && len(unusedArgs) > 0 {
		cmd.Set("card-id", unusedArgs[0])
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
	_, err = client.Cards.Freeze(ctx, cmd.Value("card-id").(string), options...)
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
		Title:          "cards freeze",
		Transform:      transform,
	})
}

func handleCardsGet(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-id") && len(unusedArgs) > 0 {
		cmd.Set("card-id", unusedArgs[0])
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
	_, err = client.Cards.Get(ctx, cmd.Value("card-id").(string), options...)
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
		Title:          "cards get",
		Transform:      transform,
	})
}

func handleCardsReveal(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-id") && len(unusedArgs) > 0 {
		cmd.Set("card-id", unusedArgs[0])
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
	_, err = client.Cards.Reveal(ctx, cmd.Value("card-id").(string), options...)
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
		Title:          "cards reveal",
		Transform:      transform,
	})
}

func handleCardsUnfreeze(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("card-id") && len(unusedArgs) > 0 {
		cmd.Set("card-id", unusedArgs[0])
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
	_, err = client.Cards.Unfreeze(ctx, cmd.Value("card-id").(string), options...)
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
		Title:          "cards unfreeze",
		Transform:      transform,
	})
}
