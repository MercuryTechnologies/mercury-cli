// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
)

func TestAccountTransactionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "account:transactions", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--category-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--end", "end",
			"--limit", "1",
			"--mercury-category", "mercuryCategory",
			"--offset", "0",
			"--order", "asc",
			"--request-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--search", "search",
			"--start", "start",
			"--status", "pending",
		)
	})
}

func TestAccountTransactionsSend(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "account:transactions", "send",
			"--api-key", "string",
			"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--amount", "0.01",
			"--idempotency-key", "idempotencyKey",
			"--payment-method", "ach",
			"--recipient-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--external-memo", "externalMemo",
			"--note", "note",
			"--purpose", "{simple: {category: Employee, additionalInfo: additionalInfo}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(accountTransactionsSend)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "account:transactions", "send",
			"--api-key", "string",
			"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--amount", "0.01",
			"--idempotency-key", "idempotencyKey",
			"--payment-method", "ach",
			"--recipient-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--external-memo", "externalMemo",
			"--note", "note",
			"--purpose.simple", "{category: Employee, additionalInfo: additionalInfo}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"amount: 0.01\n" +
			"idempotencyKey: idempotencyKey\n" +
			"paymentMethod: ach\n" +
			"recipientId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"externalMemo: externalMemo\n" +
			"note: note\n" +
			"purpose:\n" +
			"  simple:\n" +
			"    category: Employee\n" +
			"    additionalInfo: additionalInfo\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "account:transactions", "send",
			"--api-key", "string",
			"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
