// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
)

func TestTreasuryList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "treasury", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "1",
			"--order", "asc",
			"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestTreasuryRetrieveStatements(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "treasury", "retrieve-statements",
			"--api-key", "string",
			"--max-items", "10",
			"--treasury-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--document-type", "MonthlyStatement",
			"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "1",
			"--order", "asc",
			"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestTreasuryRetrieveTransactions(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "treasury", "retrieve-transactions",
			"--api-key", "string",
			"--treasury-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--cursor", "0",
			"--limit", "1",
			"--order", "asc",
		)
	})
}
