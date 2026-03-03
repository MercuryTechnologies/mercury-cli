// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
)

func TestRecipientsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"recipients", "create",
		"--email", "string",
		"--name", "name",
		"--address", "{address1: address1, city: city, postalCode: postalCode, address2: address2, state: AL}",
		"--check-info", "{address: {address1: address1, city: city, country: country, postalCode: postalCode, region: region, address2: address2}}",
		"--contact-email", "contactEmail",
		"--domestic-wire-routing-info", "{accountNumber: accountNumber, address: {address1: address1, city: city, country: country, postalCode: postalCode, region: region, address2: address2}, routingNumber: routingNumber, defaultForBenefitOf: defaultForBenefitOf}",
		"--electronic-routing-info", "{accountNumber: accountNumber, address: {address1: address1, city: city, country: country, postalCode: postalCode, region: region, address2: address2}, electronicAccountType: businessChecking, routingNumber: routingNumber}",
		"--nickname", "nickname",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(recipientsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"recipients", "create",
		"--email", "string",
		"--name", "name",
		"--address.address1", "address1",
		"--address.city", "city",
		"--address.postal-code", "postalCode",
		"--address.address2", "address2",
		"--address.state", "AL",
		"--check-info.address", "{address1: address1, city: city, country: country, postalCode: postalCode, region: region, address2: address2}",
		"--contact-email", "contactEmail",
		"--domestic-wire-routing-info.account-number", "accountNumber",
		"--domestic-wire-routing-info.address", "{address1: address1, city: city, country: country, postalCode: postalCode, region: region, address2: address2}",
		"--domestic-wire-routing-info.routing-number", "routingNumber",
		"--domestic-wire-routing-info.default-for-benefit-of", "defaultForBenefitOf",
		"--electronic-routing-info.account-number", "accountNumber",
		"--electronic-routing-info.address", "{address1: address1, city: city, country: country, postalCode: postalCode, region: region, address2: address2}",
		"--electronic-routing-info.electronic-account-type", "businessChecking",
		"--electronic-routing-info.routing-number", "routingNumber",
		"--nickname", "nickname",
	)
}

func TestRecipientsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"recipients", "retrieve",
		"--recipient-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestRecipientsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"recipients", "update",
		"--recipient-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--address", "{address1: address1, city: city, postalCode: postalCode, address2: address2, state: AL}",
		"--check-info", "{address: {address1: address1, city: city, country: country, postalCode: postalCode, region: region, address2: address2}}",
		"--contact-email", "contactEmail",
		"--domestic-wire-routing-info", "{accountNumber: accountNumber, address: {address1: address1, city: city, country: country, postalCode: postalCode, region: region, address2: address2}, routingNumber: routingNumber, defaultForBenefitOf: defaultForBenefitOf}",
		"--electronic-routing-info", "{accountNumber: accountNumber, address: {address1: address1, city: city, country: country, postalCode: postalCode, region: region, address2: address2}, electronicAccountType: businessChecking, routingNumber: routingNumber}",
		"--email", "string",
		"--name", "name",
		"--nickname", "nickname",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(recipientsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"recipients", "update",
		"--recipient-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--address.address1", "address1",
		"--address.city", "city",
		"--address.postal-code", "postalCode",
		"--address.address2", "address2",
		"--address.state", "AL",
		"--check-info.address", "{address1: address1, city: city, country: country, postalCode: postalCode, region: region, address2: address2}",
		"--contact-email", "contactEmail",
		"--domestic-wire-routing-info.account-number", "accountNumber",
		"--domestic-wire-routing-info.address", "{address1: address1, city: city, country: country, postalCode: postalCode, region: region, address2: address2}",
		"--domestic-wire-routing-info.routing-number", "routingNumber",
		"--domestic-wire-routing-info.default-for-benefit-of", "defaultForBenefitOf",
		"--electronic-routing-info.account-number", "accountNumber",
		"--electronic-routing-info.address", "{address1: address1, city: city, country: country, postalCode: postalCode, region: region, address2: address2}",
		"--electronic-routing-info.electronic-account-type", "businessChecking",
		"--electronic-routing-info.routing-number", "routingNumber",
		"--email", "string",
		"--name", "name",
		"--nickname", "nickname",
	)
}

func TestRecipientsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"recipients", "list",
		"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--limit", "1",
		"--order", "asc",
		"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestRecipientsListAttachments(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"recipients", "list-attachments",
		"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--limit", "1",
		"--order", "asc",
		"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
}

func TestRecipientsUploadAttachment(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"recipients", "upload-attachment",
		"--recipient-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--file", "",
	)
}
