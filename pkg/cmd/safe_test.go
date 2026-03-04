// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
)

func TestSafesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"safes", "retrieve",
		"--api-key", "string",
		"--safe-request-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestSafesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"safes", "list",
		"--api-key", "string",
	)
}

func TestSafesDownloadDocument(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"safes", "download-document",
		"--api-key", "string",
		"--safe-request-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--output", "/dev/null",
	)
}
