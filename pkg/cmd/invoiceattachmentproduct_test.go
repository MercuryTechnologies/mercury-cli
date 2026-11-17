// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
)

func TestInvoicesAttachmentsProductsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"invoices:attachments:products", "create",
			"--name", "x",
			"--unit-price", "0",
			"--description", "description",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: x\n" +
			"unitPrice: 0\n" +
			"description: description\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"invoices:attachments:products", "create",
		)
	})
}

func TestInvoicesAttachmentsProductsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"invoices:attachments:products", "list",
			"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "1",
			"--order", "asc",
			"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestInvoicesAttachmentsProductsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"invoices:attachments:products", "delete",
			"--product-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
