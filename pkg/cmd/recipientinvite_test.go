// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
)

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
