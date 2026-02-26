// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/mercury-cli/internal/mocktest"
)

func TestTransactionsUploadAttachment(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"transactions", "upload-attachment",
		"--transaction-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--file", "",
		"--attachment-type", "receipt",
	)
}
