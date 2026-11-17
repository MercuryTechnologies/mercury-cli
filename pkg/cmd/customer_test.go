// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
)

func TestCustomersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"customers", "create",
			"--email", "email",
			"--name", "x",
			"--address", "{address1: x, city: x, country: country, name: name, postalCode: x, region: region, address2: x}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(customersCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"customers", "create",
			"--email", "email",
			"--name", "x",
			"--address.address1", "x",
			"--address.city", "x",
			"--address.country", "country",
			"--address.name", "name",
			"--address.postal-code", "x",
			"--address.region", "region",
			"--address.address2", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"email: email\n" +
			"name: x\n" +
			"address:\n" +
			"  address1: x\n" +
			"  city: x\n" +
			"  country: country\n" +
			"  name: name\n" +
			"  postalCode: x\n" +
			"  region: region\n" +
			"  address2: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"customers", "create",
		)
	})
}

func TestCustomersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"customers", "update",
			"--customer-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--email", "email",
			"--name", "x",
			"--resend-open-invoices=true",
			"--address", "{address1: x, city: x, country: country, name: name, postalCode: x, region: region, address2: x}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(customersUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"customers", "update",
			"--customer-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--email", "email",
			"--name", "x",
			"--resend-open-invoices=true",
			"--address.address1", "x",
			"--address.city", "x",
			"--address.country", "country",
			"--address.name", "name",
			"--address.postal-code", "x",
			"--address.region", "region",
			"--address.address2", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"email: email\n" +
			"name: x\n" +
			"resendOpenInvoices: true\n" +
			"address:\n" +
			"  address1: x\n" +
			"  city: x\n" +
			"  country: country\n" +
			"  name: name\n" +
			"  postalCode: x\n" +
			"  region: region\n" +
			"  address2: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"customers", "update",
			"--customer-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestCustomersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"customers", "list",
			"--max-items", "10",
			"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "1",
			"--order", "asc",
			"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestCustomersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"customers", "delete",
			"--customer-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestCustomersGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"customers", "get",
			"--customer-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
