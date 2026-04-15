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

var customersCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a new customer for the organization",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "email",
			Usage:    " The email address for the customer.",
			Required: true,
			BodyPath: "email",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    " The name of the customer.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "address",
			Usage:    " Address input for creating or updating customers",
			BodyPath: "address",
		},
	},
	Action:          handleCustomersCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"address": {
		&requestflag.InnerFlag[string]{
			Name:       "address.address1",
			Usage:      " Primary street address line.",
			InnerField: "address1",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.city",
			Usage:      " City name.",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.country",
			Usage:      " Two-letter country code (ISO 3166-1 alpha-2).",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.name",
			Usage:      " The mailing name of the address.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.postal-code",
			Usage:      " Postal or ZIP code.",
			InnerField: "postalCode",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.region",
			Usage:      ` Either a two-letter US state code i.e. "CA" for California or a free-form identification of a particular region worldwide.`,
			InnerField: "region",
		},
		&requestflag.InnerFlag[any]{
			Name:       "address.address2",
			Usage:      " Secondary street address line (optional).",
			InnerField: "address2",
		},
	},
})

var customersUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update an existing customer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "customer-id",
			Usage:    "The customer who will receive the invoice. Use the /api/v1/ar/customers endpoint to list your customers and find the corresponding id, or create a new customer first.",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "email",
			Usage:    " The email address for the customer.",
			Required: true,
			BodyPath: "email",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    " The name of the customer.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[bool]{
			Name:     "resend-open-invoices",
			Usage:    " Open invoices for the customer will be resent with updated data\n when this is true.",
			Required: true,
			BodyPath: "resendOpenInvoices",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "address",
			Usage:    " Address input for creating or updating customers",
			BodyPath: "address",
		},
	},
	Action:          handleCustomersUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"address": {
		&requestflag.InnerFlag[string]{
			Name:       "address.address1",
			Usage:      " Primary street address line.",
			InnerField: "address1",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.city",
			Usage:      " City name.",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.country",
			Usage:      " Two-letter country code (ISO 3166-1 alpha-2).",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.name",
			Usage:      " The mailing name of the address.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.postal-code",
			Usage:      " Postal or ZIP code.",
			InnerField: "postalCode",
		},
		&requestflag.InnerFlag[string]{
			Name:       "address.region",
			Usage:      ` Either a two-letter US state code i.e. "CA" for California or a free-form identification of a particular region worldwide.`,
			InnerField: "region",
		},
		&requestflag.InnerFlag[any]{
			Name:       "address.address2",
			Usage:      " Secondary street address line (optional).",
			InnerField: "address2",
		},
	},
})

var customersList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a paginated list of customers. Supports cursor-based pagination with\nlimit, order, start_after, and end_before query parameters.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "end-before",
			Usage:     "The ID of the customer to end the page before (exclusive). When provided, results will end just before this ID and work backwards. Use this for reverse pagination or to retrieve previous pages. Cannot be combined with start_after.",
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
			Usage:     "The ID of the customer to start the page after (exclusive). When provided, results will begin with the customer immediately following this ID. Use this for standard forward pagination to get the next page of results. Cannot be combined with end_before.",
			QueryPath: "start_after",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleCustomersList,
	HideHelpCommand: true,
}

var customersDelete = cli.Command{
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
	Action:          handleCustomersDelete,
	HideHelpCommand: true,
}

var customersGet = cli.Command{
	Name:    "get",
	Usage:   "Retrieve details of a specific customer by their ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "customer-id",
			Usage:    "The customer who will receive the invoice. Use the /api/v1/ar/customers endpoint to list your customers and find the corresponding id, or create a new customer first.",
			Required: true,
		},
	},
	Action:          handleCustomersGet,
	HideHelpCommand: true,
}

func handleCustomersCreate(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.CustomerNewParams{}

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
	_, err = client.Customers.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "customers create", obj, format, explicitFormat, transform)
}

func handleCustomersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("customer-id") && len(unusedArgs) > 0 {
		cmd.Set("customer-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.CustomerUpdateParams{}

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
	_, err = client.Customers.Update(
		ctx,
		cmd.Value("customer-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "customers update", obj, format, explicitFormat, transform)
}

func handleCustomersList(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.CustomerListParams{}

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
		_, err = client.Customers.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, os.Stderr, "customers list", obj, format, explicitFormat, transform)
	} else {
		iter := client.Customers.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, os.Stderr, "customers list", iter, format, explicitFormat, transform, maxItems)
	}
}

func handleCustomersDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.Customers.Delete(ctx, cmd.Value("customer-id").(string), options...)
}

func handleCustomersGet(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Customers.Get(ctx, cmd.Value("customer-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "customers get", obj, format, explicitFormat, transform)
}
