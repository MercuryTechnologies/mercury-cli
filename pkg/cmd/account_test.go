// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
)

func TestAccountRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"account", "retrieve",
		"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestAccountList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"account", "list",
		"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--limit", "1",
		"--order", "asc",
		"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestAccountListCards(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"account", "list-cards",
		"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestAccountListStatements(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"account", "list-statements",
		"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--end", "end",
		"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--limit", "1",
		"--order", "asc",
		"--start", "start",
		"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestAccountRequestSendMoney(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"account", "request-send-money",
		"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--amount", "0.01",
		"--idempotency-key", "idempotencyKey",
		"--payment-method", "ach",
		"--recipient-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--external-memo", "externalMemo",
		"--note", "note",
	)
}

func TestAccountRetrieveTransaction(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"account", "retrieve-transaction",
		"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--transaction-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}
