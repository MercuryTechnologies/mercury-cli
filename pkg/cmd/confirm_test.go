package cmd

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormatCurrency(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    float64
		expected string
	}{
		{0.01, "$0.01"},
		{0.10, "$0.10"},
		{1.00, "$1.00"},
		{100.5, "$100.50"},
		{999.99, "$999.99"},
		{1000.00, "$1,000.00"},
		{1234.56, "$1,234.56"},
		{1234567.89, "$1,234,567.89"},
		{0, "$0.00"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, formatCurrency(tt.input))
		})
	}
}

func TestConfirmActionIO_Confirm(t *testing.T) {
	t.Parallel()

	details := []ConfirmDetail{
		{Label: "Account", Value: "acc-123"},
		{Label: "Amount", Value: "$1,000.00"},
	}

	t.Run("empty input defaults to yes", func(t *testing.T) {
		t.Parallel()
		reader := strings.NewReader("\n")
		var output bytes.Buffer
		err := confirmActionIO(reader, &output, "Send Money", details)
		require.NoError(t, err)
	})

	t.Run("y confirms", func(t *testing.T) {
		t.Parallel()
		reader := strings.NewReader("y\n")
		var output bytes.Buffer
		err := confirmActionIO(reader, &output, "Send Money", details)
		require.NoError(t, err)
	})

	t.Run("yes confirms", func(t *testing.T) {
		t.Parallel()
		reader := strings.NewReader("yes\n")
		var output bytes.Buffer
		err := confirmActionIO(reader, &output, "Send Money", details)
		require.NoError(t, err)
	})

	t.Run("Y confirms", func(t *testing.T) {
		t.Parallel()
		reader := strings.NewReader("Y\n")
		var output bytes.Buffer
		err := confirmActionIO(reader, &output, "Send Money", details)
		require.NoError(t, err)
	})
}

func TestConfirmActionIO_Cancel(t *testing.T) {
	t.Parallel()

	details := []ConfirmDetail{
		{Label: "Account", Value: "acc-123"},
	}

	t.Run("n cancels", func(t *testing.T) {
		t.Parallel()
		reader := strings.NewReader("n\n")
		var output bytes.Buffer
		err := confirmActionIO(reader, &output, "Send Money", details)
		require.Error(t, err)
		assert.True(t, errors.Is(err, ErrCancelled))
	})

	t.Run("N cancels", func(t *testing.T) {
		t.Parallel()
		reader := strings.NewReader("N\n")
		var output bytes.Buffer
		err := confirmActionIO(reader, &output, "Send Money", details)
		require.Error(t, err)
		assert.True(t, errors.Is(err, ErrCancelled))
	})

	t.Run("no cancels", func(t *testing.T) {
		t.Parallel()
		reader := strings.NewReader("no\n")
		var output bytes.Buffer
		err := confirmActionIO(reader, &output, "Send Money", details)
		require.Error(t, err)
		assert.True(t, errors.Is(err, ErrCancelled))
	})

	t.Run("EOF cancels", func(t *testing.T) {
		t.Parallel()
		reader := strings.NewReader("")
		var output bytes.Buffer
		err := confirmActionIO(reader, &output, "Send Money", details)
		require.Error(t, err)
		assert.True(t, errors.Is(err, ErrCancelled))
	})
}

func TestConfirmActionIO_Output(t *testing.T) {
	t.Parallel()

	details := []ConfirmDetail{
		{Label: "Account", Value: "acc-123"},
		{Label: "Recipient", Value: "rec-456"},
		{Label: "Amount", Value: "$500.00"},
		{Label: "Payment Method", Value: "ach"},
	}

	reader := strings.NewReader("y\n")
	var output bytes.Buffer
	err := confirmActionIO(reader, &output, "Send Money", details)
	require.NoError(t, err)

	rendered := output.String()
	assert.Contains(t, rendered, "Send Money")
	assert.Contains(t, rendered, "Account")
	assert.Contains(t, rendered, "acc-123")
	assert.Contains(t, rendered, "Recipient")
	assert.Contains(t, rendered, "rec-456")
	assert.Contains(t, rendered, "$500.00")
	assert.Contains(t, rendered, "ach")
	assert.Contains(t, rendered, "Proceed?")
	assert.Contains(t, rendered, "[Y/n]")
}

func TestConfirmActionIO_DestructiveOutput(t *testing.T) {
	t.Parallel()

	t.Run("delete customer", func(t *testing.T) {
		t.Parallel()
		details := []ConfirmDetail{
			{Label: "Name", Value: "Acme Corp"},
			{Label: "Email", Value: "billing@acme.com"},
			{Label: "ID", Value: "cust-789"},
		}

		reader := strings.NewReader("y\n")
		var output bytes.Buffer
		err := confirmActionIO(reader, &output, "Delete Customer", details)
		require.NoError(t, err)

		rendered := output.String()
		assert.Contains(t, rendered, "Delete Customer")
		assert.Contains(t, rendered, "Acme Corp")
		assert.Contains(t, rendered, "billing@acme.com")
		assert.Contains(t, rendered, "cust-789")
		assert.Contains(t, rendered, "[Y/n]")
	})

	t.Run("cancel invoice", func(t *testing.T) {
		t.Parallel()
		details := []ConfirmDetail{
			{Label: "Invoice", Value: "INV-001"},
			{Label: "Amount", Value: "$5,000.00"},
			{Label: "Status", Value: "Unpaid"},
			{Label: "Customer", Value: "cust-789"},
			{Label: "Due Date", Value: "2026-05-01"},
		}

		reader := strings.NewReader("n\n")
		var output bytes.Buffer
		err := confirmActionIO(reader, &output, "Cancel Invoice", details)
		require.Error(t, err)
		assert.True(t, errors.Is(err, ErrCancelled))

		rendered := output.String()
		assert.Contains(t, rendered, "Cancel Invoice")
		assert.Contains(t, rendered, "INV-001")
		assert.Contains(t, rendered, "$5,000.00")
		assert.Contains(t, rendered, "Unpaid")
	})

	t.Run("delete webhook", func(t *testing.T) {
		t.Parallel()
		details := []ConfirmDetail{
			{Label: "URL", Value: "https://example.com/hooks"},
			{Label: "Events", Value: "transaction.created, transaction.updated"},
			{Label: "Status", Value: "active"},
		}

		reader := strings.NewReader("y\n")
		var output bytes.Buffer
		err := confirmActionIO(reader, &output, "Delete Webhook", details)
		require.NoError(t, err)

		rendered := output.String()
		assert.Contains(t, rendered, "Delete Webhook")
		assert.Contains(t, rendered, "https://example.com/hooks")
		assert.Contains(t, rendered, "transaction.created, transaction.updated")
		assert.Contains(t, rendered, "active")
	})
}
