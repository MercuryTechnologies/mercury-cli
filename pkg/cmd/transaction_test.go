// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
)

func TestTransactionsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"transactions", "update",
			"--transaction-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--category-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--note", "note",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"categoryId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"note: note\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"transactions", "update",
			"--transaction-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestTransactionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"transactions", "list",
			"--max-items", "10",
			"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--card-id", "string",
			"--category-id", "categoryId",
			"--end", "end",
			"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "1",
			"--mercury-category", "mercuryCategory",
			"--order", "asc",
			"--posted-end", "postedEnd",
			"--posted-start", "postedStart",
			"--search", "search",
			"--start", "start",
			"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--start-at", "start_at",
			"--status", "pending",
		)
	})
}

func TestTransactionsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"transactions", "get",
			"--transaction-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
