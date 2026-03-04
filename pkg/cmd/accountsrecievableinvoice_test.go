// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
)

func TestAccountsRecievableInvoicesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts-recievable:invoices", "create",
		"--api-key", "string",
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

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(accountsRecievableInvoicesCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts-recievable:invoices", "create",
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
}

func TestAccountsRecievableInvoicesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts-recievable:invoices", "retrieve",
		"--api-key", "string",
		"--invoice-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestAccountsRecievableInvoicesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts-recievable:invoices", "update",
		"--api-key", "string",
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

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(accountsRecievableInvoicesUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts-recievable:invoices", "update",
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
}

func TestAccountsRecievableInvoicesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts-recievable:invoices", "list",
		"--api-key", "string",
		"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--limit", "1",
		"--order", "asc",
		"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestAccountsRecievableInvoicesCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts-recievable:invoices", "cancel",
		"--api-key", "string",
		"--invoice-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestAccountsRecievableInvoicesDownloadPdf(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts-recievable:invoices", "download-pdf",
		"--api-key", "string",
		"--invoice-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--output", "/dev/null",
	)
}

func TestAccountsRecievableInvoicesListAttachments(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts-recievable:invoices", "list-attachments",
		"--api-key", "string",
		"--invoice-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}
