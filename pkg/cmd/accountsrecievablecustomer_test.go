// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
)

func TestAccountsRecievableCustomersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts-recievable:customers", "create",
		"--api-key", "string",
		"--email", "email",
		"--name", "name",
		"--address", "{address1: address1, city: city, country: country, name: name, postalCode: postalCode, region: region, address2: address2}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(accountsRecievableCustomersCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts-recievable:customers", "create",
		"--email", "email",
		"--name", "name",
		"--address.address1", "address1",
		"--address.city", "city",
		"--address.country", "country",
		"--address.name", "name",
		"--address.postal-code", "postalCode",
		"--address.region", "region",
		"--address.address2", "address2",
	)
}

func TestAccountsRecievableCustomersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts-recievable:customers", "retrieve",
		"--api-key", "string",
		"--customer-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestAccountsRecievableCustomersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts-recievable:customers", "update",
		"--api-key", "string",
		"--customer-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--email", "email",
		"--name", "name",
		"--resend-open-invoices=true",
		"--address", "{address1: address1, city: city, country: country, name: name, postalCode: postalCode, region: region, address2: address2}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(accountsRecievableCustomersUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts-recievable:customers", "update",
		"--customer-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--email", "email",
		"--name", "name",
		"--resend-open-invoices=true",
		"--address.address1", "address1",
		"--address.city", "city",
		"--address.country", "country",
		"--address.name", "name",
		"--address.postal-code", "postalCode",
		"--address.region", "region",
		"--address.address2", "address2",
	)
}

func TestAccountsRecievableCustomersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts-recievable:customers", "list",
		"--api-key", "string",
		"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--limit", "1",
		"--order", "asc",
		"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestAccountsRecievableCustomersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"accounts-recievable:customers", "delete",
		"--api-key", "string",
		"--customer-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}
