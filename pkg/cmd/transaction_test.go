// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
)

func TestTransactionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"transactions", "retrieve",
		"--transaction-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestTransactionsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"transactions", "update",
		"--transaction-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--category-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--note", "note",
	)
}

func TestTransactionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"transactions", "list",
		"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--category-id", "categoryId",
		"--end", "end",
		"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--limit", "1",
		"--mercury-category", "mercuryCategory",
		"--order", "asc",
		"--posted-end", "postedEnd",
		"--posted-start", "postedStart",
		"--search", "search",
		"--start", "start",
		"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--start-at", "start_at",
		"--status", "pending",
	)
}

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
