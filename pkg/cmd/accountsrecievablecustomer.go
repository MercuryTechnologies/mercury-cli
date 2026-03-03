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

var accountsRecievableCustomersCreate = requestflag.WithInnerFlags(cli.Command{
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
	Action:          handleAccountsRecievableCustomersCreate,
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

var accountsRecievableCustomersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve details of a specific customer by their ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "customer-id",
			Usage:    "The customer who will receive the invoice. Use the /api/v1/ar/customers endpoint to list your customers and find the corresponding id, or create a new customer first.",
			Required: true,
		},
	},
	Action:          handleAccountsRecievableCustomersRetrieve,
	HideHelpCommand: true,
}

var accountsRecievableCustomersUpdate = requestflag.WithInnerFlags(cli.Command{
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
	Action:          handleAccountsRecievableCustomersUpdate,
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

var accountsRecievableCustomersList = cli.Command{
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
	},
	Action:          handleAccountsRecievableCustomersList,
	HideHelpCommand: true,
}

var accountsRecievableCustomersDelete = cli.Command{
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
	Action:          handleAccountsRecievableCustomersDelete,
	HideHelpCommand: true,
}

func handleAccountsRecievableCustomersCreate(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.AccountsRecievableCustomerNewParams{}

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
	_, err = client.AccountsRecievable.Customers.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "accounts-recievable:customers create", obj, format, transform)
}

func handleAccountsRecievableCustomersRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.AccountsRecievable.Customers.Get(ctx, cmd.Value("customer-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "accounts-recievable:customers retrieve", obj, format, transform)
}

func handleAccountsRecievableCustomersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("customer-id") && len(unusedArgs) > 0 {
		cmd.Set("customer-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.AccountsRecievableCustomerUpdateParams{}

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
	_, err = client.AccountsRecievable.Customers.Update(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "accounts-recievable:customers update", obj, format, transform)
}

func handleAccountsRecievableCustomersList(ctx context.Context, cmd *cli.Command) error {
	client := mercury.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := mercury.AccountsRecievableCustomerListParams{}

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
		_, err = client.AccountsRecievable.Customers.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "accounts-recievable:customers list", obj, format, transform)
	} else {
		iter := client.AccountsRecievable.Customers.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "accounts-recievable:customers list", iter, format, transform)
	}
}

func handleAccountsRecievableCustomersDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.AccountsRecievable.Customers.Delete(ctx, cmd.Value("customer-id").(string), options...)
}
