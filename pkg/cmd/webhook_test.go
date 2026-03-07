// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
)

func TestWebhooksCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "webhooks", "create",
			"--api-key", "string",
			"--url", "url",
			"--event-type", "[transaction.created]",
			"--filter-path", "[transaction.amount]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"url: url\n" +
			"eventTypes:\n" +
			"  - transaction.created\n" +
			"filterPaths:\n" +
			"  - transaction.amount\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "webhooks", "create",
			"--api-key", "string",
		)
	})
}

func TestWebhooksRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "webhooks", "retrieve",
			"--api-key", "string",
			"--webhook-endpoint-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestWebhooksUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "webhooks", "update",
			"--api-key", "string",
			"--webhook-endpoint-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--event-type", "[transaction.created]",
			"--filter-path", "[transaction.amount]",
			"--status", "active",
			"--url", "url",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"eventTypes:\n" +
			"  - transaction.created\n" +
			"filterPaths:\n" +
			"  - transaction.amount\n" +
			"status: active\n" +
			"url: url\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "webhooks", "update",
			"--api-key", "string",
			"--webhook-endpoint-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestWebhooksList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "webhooks", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "1",
			"--order", "asc",
			"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--status", "active",
		)
	})
}

func TestWebhooksDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "webhooks", "delete",
			"--api-key", "string",
			"--webhook-endpoint-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestWebhooksVerify(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "webhooks", "verify",
			"--api-key", "string",
			"--webhook-endpoint-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--event-type", "transaction.created",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("eventType: transaction.created")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "webhooks", "verify",
			"--api-key", "string",
			"--webhook-endpoint-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
