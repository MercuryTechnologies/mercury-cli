// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/mercury-cli/internal/mocktest"
)

func TestArInvoicesCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ar:invoices", "cancel",
		"--invoice-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}
