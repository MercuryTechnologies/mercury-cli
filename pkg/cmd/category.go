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

var categoriesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new custom expense category for the organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    " Name of the category",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[bool]{
			Name:     "visible-for-card-spend",
			Usage:    " Whether this category is applicable to card transactions",
			Required: true,
			BodyPath: "visibleForCardSpend",
		},
		&requestflag.Flag[bool]{
			Name:     "visible-for-other",
			Usage:    " Whether this category is applicable to all other transaction kinds",
			Required: true,
			BodyPath: "visibleForOther",
		},
		&requestflag.Flag[bool]{
			Name:     "visible-for-reimbursements",
			Usage:    " Whether this category is applicable to expense reimbursement transactions",
			Required: true,
			BodyPath: "visibleForReimbursements",
		},
	},
	Action:          handleCategoriesCreate,
	HideHelpCommand: true,
}

var categoriesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update an existing custom expense category for the organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "expense-category-id",
			Usage:     "ID for the category",
			Required:  true,
			PathParam: "expenseCategoryId",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    " New name for the category",
			BodyPath: "name",
		},
		&requestflag.Flag[*bool]{
			Name:     "visible-for-card-spend",
			Usage:    " Whether this category is applicable to card transactions",
			BodyPath: "visibleForCardSpend",
		},
		&requestflag.Flag[*bool]{
			Name:     "visible-for-other",
			Usage:    " Whether this category is applicable to all other transaction kinds",
			BodyPath: "visibleForOther",
		},
		&requestflag.Flag[*bool]{
			Name:     "visible-for-reimbursements",
			Usage:    " Whether this category is applicable to expense reimbursement transactions",
			BodyPath: "visibleForReimbursements",
		},
	},
	Action:          handleCategoriesUpdate,
	HideHelpCommand: true,
}

var categoriesList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a paginated list of all available custom expense categories for the\norganization. Supports cursor-based pagination with limit, order, start_after,\nand end_before query parameters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "end-before",
			Usage:     "The ID of the category to end the page before (exclusive). When provided, results will end just before this ID and work backwards. Use this for reverse pagination or to retrieve previous pages. Cannot be combined with start_after.",
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
			Usage:     "The ID of the category to start the page after (exclusive). When provided, results will begin with the category immediately following this ID. Use this for standard forward pagination to get the next page of results. Cannot be combined with end_before.",
			QueryPath: "start_after",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleCategoriesList,
	HideHelpCommand: true,
}

var categoriesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a custom expense category for the organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "expense-category-id",
			Usage:     "ID for the category",
			Required:  true,
			PathParam: "expenseCategoryId",
		},
	},
	Action:          handleCategoriesDelete,
	HideHelpCommand: true,
}

func handleCategoriesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := mercury.CategoryNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Categories.New(ctx, params, options...)
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
		Title:          "categories create",
		Transform:      transform,
	})
}

func handleCategoriesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("expense-category-id") && len(unusedArgs) > 0 {
		cmd.Set("expense-category-id", unusedArgs[0])
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

	params := mercury.CategoryUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Categories.Update(
		ctx,
		cmd.Value("expense-category-id").(string),
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
		Title:          "categories update",
		Transform:      transform,
	})
}

func handleCategoriesList(ctx context.Context, cmd *cli.Command) error {
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

	params := mercury.CategoryListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Categories.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "categories list",
			Transform:      transform,
		})
	} else {
		iter := client.Categories.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "categories list",
			Transform:      transform,
		})
	}
}

func handleCategoriesDelete(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("expense-category-id") && len(unusedArgs) > 0 {
		cmd.Set("expense-category-id", unusedArgs[0])
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

	return client.Categories.Delete(ctx, cmd.Value("expense-category-id").(string), options...)
}
