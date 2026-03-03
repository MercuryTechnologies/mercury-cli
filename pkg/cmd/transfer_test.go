// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
)

func TestTransferCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"transfer", "create",
		"--amount", "0.01",
		"--destination-account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--idempotency-key", "idempotencyKey",
		"--source-account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--note", "note",
	)
}
