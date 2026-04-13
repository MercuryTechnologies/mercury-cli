// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
)

func TestInvoicesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"invoices", "create",
			"--ach-debit-enabled=true",
			"--cc-email", "string",
			"--credit-card-enabled=true",
			"--customer-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--destination-account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--due-date", "'2016-07-22'",
			"--invoice-date", "'2016-07-22'",
			"--line-item", "{name: name, quantity: 0, unitPrice: 0, salesTaxRate: 0}",
			"--use-real-account-number=true",
			"--internal-note", "internalNote",
			"--invoice-number", "invoiceNumber",
			"--payer-memo", "payerMemo",
			"--po-number", "poNumber",
			"--send-email-option", "DontSend",
			"--service-period-end-date", "'2016-07-22'",
			"--service-period-start-date", "'2016-07-22'",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(invoicesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"invoices", "create",
			"--ach-debit-enabled=true",
			"--cc-email", "string",
			"--credit-card-enabled=true",
			"--customer-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--destination-account-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--due-date", "'2016-07-22'",
			"--invoice-date", "'2016-07-22'",
			"--line-item.name", "name",
			"--line-item.quantity", "0",
			"--line-item.unit-price", "0",
			"--line-item.sales-tax-rate", "0",
			"--use-real-account-number=true",
			"--internal-note", "internalNote",
			"--invoice-number", "invoiceNumber",
			"--payer-memo", "payerMemo",
			"--po-number", "poNumber",
			"--send-email-option", "DontSend",
			"--service-period-end-date", "'2016-07-22'",
			"--service-period-start-date", "'2016-07-22'",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"achDebitEnabled: true\n" +
			"ccEmails:\n" +
			"  - string\n" +
			"creditCardEnabled: true\n" +
			"customerId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"destinationAccountId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"dueDate: '2016-07-22'\n" +
			"invoiceDate: '2016-07-22'\n" +
			"lineItems:\n" +
			"  - name: name\n" +
			"    quantity: 0\n" +
			"    unitPrice: 0\n" +
			"    salesTaxRate: 0\n" +
			"useRealAccountNumber: true\n" +
			"internalNote: internalNote\n" +
			"invoiceNumber: invoiceNumber\n" +
			"payerMemo: payerMemo\n" +
			"poNumber: poNumber\n" +
			"sendEmailOption: DontSend\n" +
			"servicePeriodEndDate: '2016-07-22'\n" +
			"servicePeriodStartDate: '2016-07-22'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"invoices", "create",
		)
	})
}

func TestInvoicesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"invoices", "update",
			"--invoice-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--ach-debit-enabled=true",
			"--cc-email", "string",
			"--credit-card-enabled=true",
			"--due-date", "'2016-07-22'",
			"--invoice-date", "'2016-07-22'",
			"--invoice-number", "invoiceNumber",
			"--line-item", "{name: name, quantity: 0, unitPrice: 0, salesTaxRate: 0}",
			"--use-real-account-number=true",
			"--internal-note", "internalNote",
			"--payer-memo", "payerMemo",
			"--po-number", "poNumber",
			"--service-period-end-date", "'2016-07-22'",
			"--service-period-start-date", "'2016-07-22'",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(invoicesUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"invoices", "update",
			"--invoice-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--ach-debit-enabled=true",
			"--cc-email", "string",
			"--credit-card-enabled=true",
			"--due-date", "'2016-07-22'",
			"--invoice-date", "'2016-07-22'",
			"--invoice-number", "invoiceNumber",
			"--line-item.name", "name",
			"--line-item.quantity", "0",
			"--line-item.unit-price", "0",
			"--line-item.sales-tax-rate", "0",
			"--use-real-account-number=true",
			"--internal-note", "internalNote",
			"--payer-memo", "payerMemo",
			"--po-number", "poNumber",
			"--service-period-end-date", "'2016-07-22'",
			"--service-period-start-date", "'2016-07-22'",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"achDebitEnabled: true\n" +
			"ccEmails:\n" +
			"  - string\n" +
			"creditCardEnabled: true\n" +
			"dueDate: '2016-07-22'\n" +
			"invoiceDate: '2016-07-22'\n" +
			"invoiceNumber: invoiceNumber\n" +
			"lineItems:\n" +
			"  - name: name\n" +
			"    quantity: 0\n" +
			"    unitPrice: 0\n" +
			"    salesTaxRate: 0\n" +
			"useRealAccountNumber: true\n" +
			"internalNote: internalNote\n" +
			"payerMemo: payerMemo\n" +
			"poNumber: poNumber\n" +
			"servicePeriodEndDate: '2016-07-22'\n" +
			"servicePeriodStartDate: '2016-07-22'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"invoices", "update",
			"--invoice-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestInvoicesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"invoices", "list",
			"--max-items", "10",
			"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "1",
			"--order", "asc",
			"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestInvoicesCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"invoices", "cancel",
			"--invoice-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestInvoicesDownload(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"invoices", "download",
			"--invoice-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--output", "/dev/null",
		)
	})
}

func TestInvoicesGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"invoices", "get",
			"--invoice-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
