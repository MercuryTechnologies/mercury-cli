// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
)

func TestSpendReceiptsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend", "receipts", "list",
			"--max-items", "10",
			"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "1",
			"--order", "asc",
			"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--start-at", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestSpendReceiptsCreateUploadURL(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend", "receipts", "create-upload-url",
			"--file-name", "fileName",
			"--idempotency-key", "Idempotency-Key",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("fileName: fileName")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"spend", "receipts", "create-upload-url",
			"--idempotency-key", "Idempotency-Key",
		)
	})
}

func TestSpendReceiptsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend", "receipts", "get",
			"--receipt-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestSpendReceiptsMatch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend", "receipts", "match",
			"--receipt-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--idempotency-key", "Idempotency-Key",
		)
	})
}
