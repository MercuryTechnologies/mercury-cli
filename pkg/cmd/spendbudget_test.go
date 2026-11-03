// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
)

func TestSpendBudgetsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend", "budgets", "create",
			"--frequency", "oneTime",
			"--amount-per-user-per-period", "0",
			"--description", "description",
			"--end-date", "'2016-07-22'",
			"--name", "name",
			"--start-date", "'2016-07-22'",
			"--idempotency-key", "Idempotency-Key",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"frequency: oneTime\n" +
			"amountPerUserPerPeriod: 0\n" +
			"description: description\n" +
			"endDate: '2016-07-22'\n" +
			"name: name\n" +
			"startDate: '2016-07-22'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"spend", "budgets", "create",
			"--idempotency-key", "Idempotency-Key",
		)
	})
}

func TestSpendBudgetsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend", "budgets", "update",
			"--budget-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--amount-per-user-per-period", "0",
			"--description", "description",
			"--end-date", "'2016-07-22'",
			"--frequency", "oneTime",
			"--name", "name",
			"--start-date", "'2016-07-22'",
			"--idempotency-key", "Idempotency-Key",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"amountPerUserPerPeriod: 0\n" +
			"description: description\n" +
			"endDate: '2016-07-22'\n" +
			"frequency: oneTime\n" +
			"name: name\n" +
			"startDate: '2016-07-22'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"spend", "budgets", "update",
			"--budget-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--idempotency-key", "Idempotency-Key",
		)
	})
}

func TestSpendBudgetsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend", "budgets", "list",
			"--max-items", "10",
			"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "1",
			"--order", "asc",
			"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--start-at", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestSpendBudgetsArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend", "budgets", "archive",
			"--budget-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--idempotency-key", "Idempotency-Key",
		)
	})
}

func TestSpendBudgetsAssignTransaction(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend", "budgets", "assign-transaction",
			"--budget-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--transaction-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--allow-overspend=true",
			"--idempotency-key", "Idempotency-Key",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"transactionId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"allowOverspend: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"spend", "budgets", "assign-transaction",
			"--budget-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--idempotency-key", "Idempotency-Key",
		)
	})
}

func TestSpendBudgetsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend", "budgets", "get",
			"--budget-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestSpendBudgetsListExpenditures(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend", "budgets", "list-expenditures",
			"--max-items", "10",
			"--budget-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "1",
			"--order", "asc",
			"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--start-at", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestSpendBudgetsListMembers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend", "budgets", "list-members",
			"--max-items", "10",
			"--budget-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "1",
			"--order", "asc",
			"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--start-at", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
