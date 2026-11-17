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

var invoicesAttachmentsProductsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new invoicing product (product catalog item). The combination of name\nand unit price must be unique within the organization.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    " Name of the product. Must be unique per organization in combination with the unit price.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[float64]{
			Name:     "unit-price",
			Usage:    "A dollar amount",
			Required: true,
			BodyPath: "unitPrice",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			Usage:    " Optional description of the product.",
			BodyPath: "description",
		},
	},
	Action:          handleInvoicesAttachmentsProductsCreate,
	HideHelpCommand: true,
}

var invoicesAttachmentsProductsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a paginated list of invoicing products (product catalog items).\nSupports cursor-based pagination with limit, order, start_after, and end_before\nquery parameters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "end-before",
			Usage:     "The ID of the product to end the page before (exclusive). When provided, results will end just before this ID and work backwards. Use this for reverse pagination or to retrieve previous pages. Cannot be combined with start_after.",
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
			Usage:     "The ID of the product to start the page after (exclusive). When provided, results will begin with the product immediately following this ID. Use this for standard forward pagination to get the next page of results. Cannot be combined with end_before.",
			QueryPath: "start_after",
		},
	},
	Action:          handleInvoicesAttachmentsProductsList,
	HideHelpCommand: true,
}

var invoicesAttachmentsProductsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an invoicing product (product catalog item). The product is unlinked from\nline items on open, pending, and recurring invoices.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "product-id",
			Usage:     "The ID of an invoicing product. Use the /api/v1/ar/products endpoint to list your products.",
			Required:  true,
			PathParam: "productId",
		},
	},
	Action:          handleInvoicesAttachmentsProductsDelete,
	HideHelpCommand: true,
}

func handleInvoicesAttachmentsProductsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := mercury.InvoiceAttachmentProductNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Invoices.Attachments.Products.New(ctx, params, options...)
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
		Title:          "invoices:attachments:products create",
		Transform:      transform,
	})
}

func handleInvoicesAttachmentsProductsList(ctx context.Context, cmd *cli.Command) error {
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

	params := mercury.InvoiceAttachmentProductListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Invoices.Attachments.Products.List(ctx, params, options...)
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
		Title:          "invoices:attachments:products list",
		Transform:      transform,
	})
}

func handleInvoicesAttachmentsProductsDelete(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("product-id") && len(unusedArgs) > 0 {
		cmd.Set("product-id", unusedArgs[0])
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

	return client.Invoices.Attachments.Products.Delete(ctx, cmd.Value("product-id").(string), options...)
}
