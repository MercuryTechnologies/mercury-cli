// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
)

func TestAccountsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts", "list",
			"--max-items", "10",
			"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "1",
			"--order", "asc",
			"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAccountsCreateTransaction(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts", "create-transaction",
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
		requestflag.CheckInnerFlags(accountsCreateTransaction)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts", "create-transaction",
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
			t, pipeData,
			"--api-key", "string",
			"accounts", "create-transaction",
			"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAccountsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts", "get",
			"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAccountsRequestSendMoney(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"accounts", "request-send-money",
			"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--amount", "0.01",
			"--idempotency-key", "idempotencyKey",
			"--payment-method", "ach",
			"--recipient-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--external-memo", "externalMemo",
			"--note", "note",
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
			"note: note\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"accounts", "request-send-money",
			"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
