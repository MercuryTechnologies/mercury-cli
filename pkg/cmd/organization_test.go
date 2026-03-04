// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
)

func TestOrganizationRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"organization", "retrieve",
		"--api-key", "string",
	)
}
