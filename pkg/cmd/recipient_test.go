// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"strings"
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
)

func TestRecipientsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(recipientsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"emails:\n" +
			"  - string\n" +
			"name: name\n" +
			"address:\n" +
			"  address1: address1\n" +
			"  city: city\n" +
			"  postalCode: postalCode\n" +
			"  address2: address2\n" +
			"  state: AL\n" +
			"checkInfo:\n" +
			"  address:\n" +
			"    address1: address1\n" +
			"    city: city\n" +
			"    country: country\n" +
			"    postalCode: postalCode\n" +
			"    region: region\n" +
			"    address2: address2\n" +
			"contactEmail: contactEmail\n" +
			"domesticWireRoutingInfo:\n" +
			"  accountNumber: accountNumber\n" +
			"  address:\n" +
			"    address1: address1\n" +
			"    city: city\n" +
			"    country: country\n" +
			"    postalCode: postalCode\n" +
			"    region: region\n" +
			"    address2: address2\n" +
			"  routingNumber: routingNumber\n" +
			"  defaultForBenefitOf: defaultForBenefitOf\n" +
			"electronicRoutingInfo:\n" +
			"  accountNumber: accountNumber\n" +
			"  address:\n" +
			"    address1: address1\n" +
			"    city: city\n" +
			"    country: country\n" +
			"    postalCode: postalCode\n" +
			"    region: region\n" +
			"    address2: address2\n" +
			"  electronicAccountType: businessChecking\n" +
			"  routingNumber: routingNumber\n" +
			"nickname: nickname\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"recipients", "create",
		)
	})
}

func TestRecipientsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(recipientsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"address:\n" +
			"  address1: address1\n" +
			"  city: city\n" +
			"  postalCode: postalCode\n" +
			"  address2: address2\n" +
			"  state: AL\n" +
			"checkInfo:\n" +
			"  address:\n" +
			"    address1: address1\n" +
			"    city: city\n" +
			"    country: country\n" +
			"    postalCode: postalCode\n" +
			"    region: region\n" +
			"    address2: address2\n" +
			"contactEmail: contactEmail\n" +
			"domesticWireRoutingInfo:\n" +
			"  accountNumber: accountNumber\n" +
			"  address:\n" +
			"    address1: address1\n" +
			"    city: city\n" +
			"    country: country\n" +
			"    postalCode: postalCode\n" +
			"    region: region\n" +
			"    address2: address2\n" +
			"  routingNumber: routingNumber\n" +
			"  defaultForBenefitOf: defaultForBenefitOf\n" +
			"electronicRoutingInfo:\n" +
			"  accountNumber: accountNumber\n" +
			"  address:\n" +
			"    address1: address1\n" +
			"    city: city\n" +
			"    country: country\n" +
			"    postalCode: postalCode\n" +
			"    region: region\n" +
			"    address2: address2\n" +
			"  electronicAccountType: businessChecking\n" +
			"  routingNumber: routingNumber\n" +
			"emails:\n" +
			"  - string\n" +
			"name: name\n" +
			"nickname: nickname\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"recipients", "update",
			"--recipient-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestRecipientsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"recipients", "list",
			"--max-items", "10",
			"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "1",
			"--order", "asc",
			"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestRecipientsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"recipients", "get",
			"--recipient-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestRecipientsListAttachments(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"recipients", "list-attachments",
			"--max-items", "10",
			"--end-before", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "1",
			"--order", "asc",
			"--start-after", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestRecipientsUploadAttachment(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"recipients", "upload-attachment",
			"--recipient-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--file", mocktest.TestFile(t, "Example data"),
		)
	})

	t.Run("piping data", func(t *testing.T) {
		testFile := mocktest.TestFile(t, "Example data")
		// Test piping YAML data over stdin
		pipeDataStr := "file: Example data"
		pipeDataStr = strings.ReplaceAll(pipeDataStr, "Example data", testFile)
		pipeData := []byte(pipeDataStr)
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"recipients", "upload-attachment",
			"--recipient-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
