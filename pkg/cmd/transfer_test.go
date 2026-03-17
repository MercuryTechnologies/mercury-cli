// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
)

func TestTransferCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"transfer", "create",
			"--amount", "0.01",
			"--destination-account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--idempotency-key", "idempotencyKey",
			"--source-account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--note", "note",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"amount: 0.01\n" +
			"destinationAccountId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"idempotencyKey: idempotencyKey\n" +
			"sourceAccountId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"note: note\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"transfer", "create",
		)
	})
}
