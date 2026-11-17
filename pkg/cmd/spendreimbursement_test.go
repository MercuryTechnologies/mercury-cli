// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
)

func TestSpendReimbursementsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend:reimbursements", "create",
			"--amount", "0",
			"--currency", "currency",
			"--date-paid-for", "'2016-07-22'",
			"--merchant-name", "x",
			"--category-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--notes", "x",
			"--idempotency-key", "Idempotency-Key",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"amount: 0\n" +
			"currency: currency\n" +
			"datePaidFor: '2016-07-22'\n" +
			"merchantName: x\n" +
			"categoryId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"notes: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"spend:reimbursements", "create",
			"--idempotency-key", "Idempotency-Key",
		)
	})
}

func TestSpendReimbursementsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend:reimbursements", "update",
			"--reimbursement-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--amount", "0",
			"--currency", "currency",
			"--date-paid-for", "'2016-07-22'",
			"--merchant-name", "x",
			"--category-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--notes", "x",
			"--idempotency-key", "Idempotency-Key",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"amount: 0\n" +
			"currency: currency\n" +
			"datePaidFor: '2016-07-22'\n" +
			"merchantName: x\n" +
			"categoryId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"notes: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"spend:reimbursements", "update",
			"--reimbursement-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--idempotency-key", "Idempotency-Key",
		)
	})
}

func TestSpendReimbursementsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend:reimbursements", "list",
			"--max-items", "10",
			"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "1",
			"--order", "asc",
			"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--start-at", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestSpendReimbursementsApprove(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend:reimbursements", "approve",
			"--reimbursement-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--idempotency-key", "Idempotency-Key",
		)
	})
}

func TestSpendReimbursementsCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend:reimbursements", "cancel",
			"--reimbursement-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--idempotency-key", "Idempotency-Key",
		)
	})
}

func TestSpendReimbursementsDeny(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend:reimbursements", "deny",
			"--reimbursement-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--reason", "x",
			"--idempotency-key", "Idempotency-Key",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("reason: x")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"spend:reimbursements", "deny",
			"--reimbursement-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--idempotency-key", "Idempotency-Key",
		)
	})
}

func TestSpendReimbursementsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend:reimbursements", "get",
			"--reimbursement-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestSpendReimbursementsMarkPaid(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"spend:reimbursements", "mark-paid",
			"--reimbursement-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--note", "x",
			"--idempotency-key", "Idempotency-Key",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("note: x")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"spend:reimbursements", "mark-paid",
			"--reimbursement-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--idempotency-key", "Idempotency-Key",
		)
	})
}
