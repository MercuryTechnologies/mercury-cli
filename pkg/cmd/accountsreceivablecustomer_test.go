// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
)

func TestAccountsReceivableCustomersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts-receivable:customers", "create",
			"--email", "email",
			"--name", "name",
			"--address", "{address1: address1, city: city, country: country, name: name, postalCode: postalCode, region: region, address2: address2}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(accountsReceivableCustomersCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts-receivable:customers", "create",
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
			"accounts-receivable:customers", "create",
		)
	})
}

func TestAccountsReceivableCustomersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts-receivable:customers", "retrieve",
			"--customer-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAccountsReceivableCustomersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts-receivable:customers", "update",
			"--customer-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--email", "email",
			"--name", "name",
			"--resend-open-invoices=true",
			"--address", "{address1: address1, city: city, country: country, name: name, postalCode: postalCode, region: region, address2: address2}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(accountsReceivableCustomersUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts-receivable:customers", "update",
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
			"accounts-receivable:customers", "update",
			"--customer-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAccountsReceivableCustomersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts-receivable:customers", "list",
			"--max-items", "10",
			"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "1",
			"--order", "asc",
			"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAccountsReceivableCustomersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts-receivable:customers", "delete",
			"--customer-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
