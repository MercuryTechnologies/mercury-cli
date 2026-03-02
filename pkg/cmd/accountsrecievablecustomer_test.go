// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
)

func TestAccountsRecievableCustomersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts-recievable:customers", "delete",
		"--customer-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}
