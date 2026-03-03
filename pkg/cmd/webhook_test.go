// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
)

func TestWebhooksCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"webhooks", "create",
		"--url", "url",
		"--event-type", "[transaction.created]",
		"--filter-path", "[transaction.amount]",
	)
}

func TestWebhooksRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"webhooks", "retrieve",
		"--webhook-endpoint-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestWebhooksUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"webhooks", "update",
		"--webhook-endpoint-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--event-type", "[transaction.created]",
		"--filter-path", "[transaction.amount]",
		"--status", "active",
		"--url", "url",
	)
}

func TestWebhooksList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"webhooks", "list",
		"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--limit", "1",
		"--order", "asc",
		"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--status", "active",
	)
}

func TestWebhooksDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"webhooks", "delete",
		"--webhook-endpoint-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestWebhooksVerify(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"webhooks", "verify",
		"--webhook-endpoint-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--event-type", "transaction.created",
	)
}
