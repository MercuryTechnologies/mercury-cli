// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"strings"
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/mocktest"
)

func TestTransactionsAttachmentsAttach(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"transactions", "attachments", "attach",
			"--transaction-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--file", mocktest.TestFile(t, "Example data"),
			"--attachment-type", "receipt",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		testFile := mocktest.TestFile(t, "Example data")
		// Test piping YAML data over stdin
		pipeDataStr := "" +
			"file: Example data\n" +
			"attachmentType: receipt\n"
		pipeDataStr = strings.ReplaceAll(pipeDataStr, "Example data", testFile)
		pipeData := []byte(pipeDataStr)
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"transactions", "attachments", "attach",
			"--transaction-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
