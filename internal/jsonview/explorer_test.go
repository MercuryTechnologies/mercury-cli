package jsonview

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/charmbracelet/bubbles/help"
	"github.com/tidwall/gjson"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNavigateForward_EmptyRowData(t *testing.T) {
	t.Parallel()

	// An empty JSON array produces a TableView with no rows.
	emptyArray := gjson.Parse("[]")
	view, err := newTableView("", emptyArray, false)
	require.NoError(t, err)

	viewer := &JSONViewer{
		stack: []JSONView{view},
		root:  "test",
		help:  help.New(),
	}

	// Should return without panicking despite the empty data set.
	model, cmd := viewer.navigateForward()
	require.Equal(t, model, viewer, "expected same viewer model returned")
	require.Nil(t, cmd)

	// Stack should remain unchanged (no new view pushed).
	require.Equal(t, 1, len(viewer.stack), "expected stack length 1, got %d", len(viewer.stack))
}

// rawJSONItem implements HasRawJSON, returning pre-built JSON.
type rawJSONItem struct {
	raw string
}

func (r rawJSONItem) RawJSON() string { return r.raw }

func TestMarshalItemsToJSONArray_WithHasRawJSON(t *testing.T) {
	t.Parallel()

	items := []any{
		rawJSONItem{raw: `{"id":1,"name":"alice"}`},
		rawJSONItem{raw: `{"id":2,"name":"bob"}`},
	}

	got, err := marshalItemsToJSONArray(items)
	require.NoError(t, err)
	require.JSONEq(t, `[{"id":1,"name":"alice"},{"id":2,"name":"bob"}]`, string(got))
}

func TestMarshalItemsToJSONArray_WithoutHasRawJSON(t *testing.T) {
	t.Parallel()

	items := []any{
		map[string]any{"id": 1, "name": "alice"},
		map[string]any{"id": 2, "name": "bob"},
	}

	got, err := marshalItemsToJSONArray(items)
	require.NoError(t, err)
	require.JSONEq(t, `[{"id":1,"name":"alice"},{"id":2,"name":"bob"}]`, string(got))
}

func TestStripControlBytes(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		in   string
		want string
	}{
		{"plain ASCII unchanged", "Innocent memo", "Innocent memo"},
		{"preserves tab", "col1\tcol2", "col1\tcol2"},
		{"preserves newline", "line1\nline2", "line1\nline2"},
		{"preserves CRLF — note CR is stripped", "line1\r\nline2", "line1\nline2"},
		{"strips ESC (CSI clear screen)", "Hello\x1b[2JWorld", "Hello[2JWorld"},
		{"strips OSC-52 clipboard write", "x\x1b]52;c;Y3VybA==\a", "x]52;c;Y3VybA=="},
		// OSC-8 hyperlink: ESC ] 8 ;; URL ESC \ TEXT ESC ] 8 ;; ESC \
		// Stripping only the ESC bytes (the only control characters here)
		// leaves the literal ']', ';', and '\' intact — that is the goal:
		// the surviving printable bytes are inert.
		{"strips OSC-8 hyperlink open+close", "\x1b]8;;https://x\x1b\\link\x1b]8;;\x1b\\", "]8;;https://x\\link]8;;\\"},
		{"strips DEL (0x7F)", "abc\x7fdef", "abcdef"},
		{"strips C1 controls (0x80, 0x85)", "hello", "hello"},
		{"strips bell", "alert\abeep", "alertbeep"},
		{"strips NUL", "x\x00y", "xy"},
		{"preserves UTF-8 with currency + emoji", "café 🏦", "café 🏦"},
		{"empty stays empty", "", ""},
		{"all-control becomes empty", "\x00\x01\x1b\x7f", ""},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, stripControlBytes(tc.in),
				"stripControlBytes(%q) = %q; want %q", tc.in, stripControlBytes(tc.in), tc.want)
		})
	}
}

// formatValue is the gjson.String render path used by table-row cells. A
// malicious memo, recipient name, or payment note in an API response must
// not survive this function with control bytes intact, or the bytes execute
// in the user's terminal under --format=explore.
func TestFormatValue_StripsControlBytesFromStrings(t *testing.T) {
	t.Parallel()

	payload := "memo\x1b[2J\x1b[H### SPOOFED ###"
	jsonBlob := mustEncodeJSON(t, map[string]string{"memo": payload})
	memo := gjson.GetBytes(jsonBlob, "memo")

	out := formatValue(memo, false /* user-facing render path */)

	assert.NotContains(t, out, "\x1b", "ESC must not survive formatValue")
	assert.Equal(t, "memo[2J[H### SPOOFED ###", out)
}

// formatObjectKey renders both keys and values into the table-cell summary
// shown when an object is collapsed; a control byte in either side would
// reach the terminal.
func TestFormatObjectKey_StripsControlBytesFromKeyAndValue(t *testing.T) {
	t.Parallel()

	jsonBlob := mustEncodeJSON(t, map[string]string{
		"key\x1b[2Jhost": "value\x1b]52;c;evil\a",
	})
	root := gjson.ParseBytes(jsonBlob)
	out := formatValue(root, false) // root is an object, hits formatObject → formatObjectKey

	assert.NotContains(t, out, "\x1b", "ESC must not survive formatObjectKey")
	assert.NotContains(t, out, "\a", "BEL must not survive formatObjectKey")
}

// quoteString builds the breadcrumb shown in the TUI title; control bytes in
// JSON object keys must not execute when the user navigates a path that
// includes them.
func TestQuoteString_StripsControlBytes(t *testing.T) {
	t.Parallel()

	out := quoteString("key\x1b[2Jname")

	assert.NotContains(t, out, "\x1b", "ESC must not survive quoteString")
	// Style.Render wraps the value in ANSI styling sequences of its own; we
	// only care that the malicious payload was scrubbed before styling.
	assert.Contains(t, out, "key[2Jname")
}

// mustEncodeJSON marshals v to RFC 8259 JSON (escaping control bytes as
// \uXXXX), reproducing what an HTTP response from the Mercury API would
// look like on the wire.
func mustEncodeJSON(t *testing.T, v any) []byte {
	t.Helper()
	b, err := json.Marshal(v)
	require.NoError(t, err)
	// Sanity — the marshaled JSON must NOT contain a raw ESC byte; if it did,
	// the test fixture would be doing the work the SUT is meant to do.
	require.False(t, strings.ContainsRune(string(b), 0x1B), "fixture leaked raw ESC")
	return b
}
