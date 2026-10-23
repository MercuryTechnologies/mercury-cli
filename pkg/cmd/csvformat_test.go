package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func TestFormatCSV(t *testing.T) {
	t.Parallel()

	t.Run("SingleObject", func(t *testing.T) {
		t.Parallel()

		res := gjson.Parse(`{"id":"abc123","amount":-12.5,"note":null}`)
		formatted, err := formatJSON(res, ShowJSONOpts{Format: "csv", Stdout: os.Stdout})
		require.NoError(t, err)
		require.Equal(t, "id,amount,note\nabc123,-12.5,\n", string(formatted))
	})

	t.Run("ArrayOfObjects", func(t *testing.T) {
		t.Parallel()

		res := gjson.Parse(`[{"id":"abc","name":"first"},{"id":"def","name":"second"}]`)
		formatted, err := formatJSON(res, ShowJSONOpts{Format: "csv", Stdout: os.Stdout})
		require.NoError(t, err)
		require.Equal(t, "id,name\nabc,first\ndef,second\n", string(formatted))
	})

	t.Run("NestedObjectsFlattenToDottedColumns", func(t *testing.T) {
		t.Parallel()

		res := gjson.Parse(`{"id":"abc","details":{"address":{"city":"Portland"},"kind":"wire"}}`)
		formatted, err := formatJSON(res, ShowJSONOpts{Format: "csv", Stdout: os.Stdout})
		require.NoError(t, err)
		require.Equal(t, "id,details.address.city,details.kind\nabc,Portland,wire\n", string(formatted))
	})

	t.Run("ArraysEmittedAsRawJSON", func(t *testing.T) {
		t.Parallel()

		res := gjson.Parse(`{"id":"abc","attachments":[{"url":"x"}]}`)
		formatted, err := formatJSON(res, ShowJSONOpts{Format: "csv", Stdout: os.Stdout})
		require.NoError(t, err)
		require.Equal(t, "id,attachments\nabc,\"[{\"\"url\"\":\"\"x\"\"}]\"\n", string(formatted))
	})

	t.Run("QuotingAndEscaping", func(t *testing.T) {
		t.Parallel()

		res := gjson.Parse(`{"note":"has, comma and \"quotes\"","bank":"plain"}`)
		formatted, err := formatJSON(res, ShowJSONOpts{Format: "csv", Stdout: os.Stdout})
		require.NoError(t, err)
		require.Equal(t, "note,bank\n\"has, comma and \"\"quotes\"\"\",plain\n", string(formatted))
	})

	t.Run("MissingFieldsAlignToFirstRowColumns", func(t *testing.T) {
		t.Parallel()

		res := gjson.Parse(`[{"id":"abc","note":"n1"},{"id":"def"},{"id":"ghi","note":"n3","extra":"dropped"}]`)
		formatted, err := formatJSON(res, ShowJSONOpts{Format: "csv", Stdout: os.Stdout})
		require.NoError(t, err)
		require.Equal(t, "id,note\nabc,n1\ndef,\nghi,n3\n", string(formatted))
	})

	t.Run("ScalarBecomesValueColumn", func(t *testing.T) {
		t.Parallel()

		res := gjson.Parse(`{"id":"abc123"}`)
		formatted, err := formatJSON(res, ShowJSONOpts{Format: "csv", Stdout: os.Stdout, Transform: "id"})
		require.NoError(t, err)
		require.Equal(t, "value\nabc123\n", string(formatted))
	})

	t.Run("IteratorWritesSingleHeader", func(t *testing.T) {
		t.Parallel()

		iter := &sliceIterator[map[string]any]{items: []map[string]any{
			{"id": "abc", "name": "first"},
			{"id": "def", "name": "second"},
		}}
		captured := captureShowJSONIterator(t, iter, "csv", "", -1)
		assert.Equal(t, "id,name\nabc,first\ndef,second\n", captured)
	})
}
