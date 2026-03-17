// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
)

func TestAccountsRecievableCustomersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts-recievable:customers", "create",
			"--email", "email",
			"--name", "name",
			"--address", "{address1: address1, city: city, country: country, name: name, postalCode: postalCode, region: region, address2: address2}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(accountsRecievableCustomersCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts-recievable:customers", "create",
			"--email", "email",
			"--name", "name",
			"--address.address1", "address1",
			"--address.city", "city",
			"--address.country", "country",
			"--address.name", "name",
			"--address.postal-code", "postalCode",
			"--address.region", "region",
			"--address.address2", "address2",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"email: email\n" +
			"name: name\n" +
			"address:\n" +
			"  address1: address1\n" +
			"  city: city\n" +
			"  country: country\n" +
			"  name: name\n" +
			"  postalCode: postalCode\n" +
			"  region: region\n" +
			"  address2: address2\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"accounts-recievable:customers", "create",
		)
	})
}

func TestAccountsRecievableCustomersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts-recievable:customers", "retrieve",
			"--customer-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAccountsRecievableCustomersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts-recievable:customers", "update",
			"--customer-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--email", "email",
			"--name", "name",
			"--resend-open-invoices=true",
			"--address", "{address1: address1, city: city, country: country, name: name, postalCode: postalCode, region: region, address2: address2}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(accountsRecievableCustomersUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts-recievable:customers", "update",
			"--customer-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--email", "email",
			"--name", "name",
			"--resend-open-invoices=true",
			"--address.address1", "address1",
			"--address.city", "city",
			"--address.country", "country",
			"--address.name", "name",
			"--address.postal-code", "postalCode",
			"--address.region", "region",
			"--address.address2", "address2",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"email: email\n" +
			"name: name\n" +
			"resendOpenInvoices: true\n" +
			"address:\n" +
			"  address1: address1\n" +
			"  city: city\n" +
			"  country: country\n" +
			"  name: name\n" +
			"  postalCode: postalCode\n" +
			"  region: region\n" +
			"  address2: address2\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"accounts-recievable:customers", "update",
			"--customer-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAccountsRecievableCustomersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts-recievable:customers", "list",
			"--max-items", "10",
			"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "1",
			"--order", "asc",
			"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAccountsRecievableCustomersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts-recievable:customers", "delete",
			"--customer-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
