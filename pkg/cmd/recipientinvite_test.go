// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
)

func TestRecipientsInvitesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"recipients:invites", "create",
			"--contact-email", "contactEmail",
			"--payment-method", "ach",
			"--require-tax-document=true",
			"--send-email=true",
			"--name", "name",
			"--notes", "notes",
			"--organization-name-on-request", "organizationNameOnRequest",
			"--recipient-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"contactEmail: contactEmail\n" +
			"paymentMethods:\n" +
			"  - ach\n" +
			"requireTaxDocument: true\n" +
			"sendEmail: true\n" +
			"name: name\n" +
			"notes: notes\n" +
			"organizationNameOnRequest: organizationNameOnRequest\n" +
			"recipientId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"recipients:invites", "create",
		)
	})
}

func TestRecipientsInvitesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"recipients:invites", "list",
			"--end-before", "end_before",
			"--limit", "1",
			"--order", "asc",
			"--start-after", "start_after",
			"--status", "created",
		)
	})
}
