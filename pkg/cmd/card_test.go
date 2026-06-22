// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
)

func TestCardsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cards", "create",
			"--kind", "debit",
			"--type", "virtual",
			"--user-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--nickname", "nickname",
			"--spend-limit", "{amountCents: 0, interval: daily}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(cardsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cards", "create",
			"--kind", "debit",
			"--type", "virtual",
			"--user-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--nickname", "nickname",
			"--spend-limit.amount-cents", "0",
			"--spend-limit.interval", "daily",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"kind: debit\n" +
			"type: virtual\n" +
			"userId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"accountId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"nickname: nickname\n" +
			"spendLimit:\n" +
			"  amountCents: 0\n" +
			"  interval: daily\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"cards", "create",
		)
	})
}

func TestCardsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cards", "update",
			"--card-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--nickname", "nickname",
			"--spend-limit", "{amountCents: 0, interval: daily, atmAmountCents: 0}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(cardsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cards", "update",
			"--card-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--nickname", "nickname",
			"--spend-limit.amount-cents", "0",
			"--spend-limit.interval", "daily",
			"--spend-limit.atm-amount-cents", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"nickname: nickname\n" +
			"spendLimit:\n" +
			"  amountCents: 0\n" +
			"  interval: daily\n" +
			"  atmAmountCents: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"cards", "update",
			"--card-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestCardsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cards", "list",
			"--max-items", "10",
			"--account-id", "string",
			"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--kind", "debit",
			"--limit", "1",
			"--order", "asc",
			"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--status", "active",
			"--type", "virtual",
			"--user-id", "userId",
		)
	})
}

func TestCardsCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cards", "cancel",
			"--card-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestCardsFreeze(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cards", "freeze",
			"--card-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestCardsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cards", "get",
			"--card-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestCardsUnfreeze(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"cards", "unfreeze",
			"--card-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
