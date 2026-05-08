package jsonview

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

// formatResult is the static `--format=pretty` render path. A malicious
// string value (memo, recipient name, error message, etc.) returned by the
// Mercury API must not carry ANSI escapes, OSC sequences, BEL, or other
// control bytes through to the user's terminal. Mirrors the explorer-mode
// strip from #51 for the static render path.
func TestStaticFormatResult_StripsControlBytesFromStringValues(t *testing.T) {
	t.Parallel()

	payload := "innocent\x1b[2J\x1b[H### SPOOFED ###"
	res := gjson.Parse(`"` + escapeJSON(payload) + `"`)

	out := formatResult(res, 0, 200)

	assert.NotContains(t, out, "\x1b", "ESC must not survive formatResult")
	assert.Contains(t, out, "innocent[2J[H### SPOOFED ###",
		"plain bytes after the strip should remain visible")
}

// formatJSONObject is the path that renders object keys for `pretty` output.
// Keys can carry control bytes from a hostile API response; they must not
// reach the terminal raw.
func TestStaticFormatJSONObject_StripsControlBytesFromKeys(t *testing.T) {
	t.Parallel()

	jsonBlob := []byte(`{"key[2Jhost":"value]52;c;evil"}`)
	root := gjson.ParseBytes(jsonBlob)

	out := formatJSONObject(root, 0, 200)

	assert.NotContains(t, out, "\x1b", "ESC must not survive formatJSONObject")
	assert.NotContains(t, out, "\a", "BEL must not survive formatJSONObject")
}

// RenderJSON is the public entry the CLI uses for `--format=pretty`. End-to-end
// check that a hostile string value cannot deliver an OSC-52 clipboard-write
// payload through the static renderer.
func TestRenderJSON_StripsOSC52ClipboardPayload(t *testing.T) {
	t.Parallel()

	jsonBlob := `{"memo":"hello]52;c;Y3VybCBldmlsLmNvbQ==world"}`
	root := gjson.Parse(jsonBlob)

	out := RenderJSON("test", root)

	assert.NotContains(t, out, "\x1b]", "OSC introducer must not survive RenderJSON")
	assert.NotContains(t, out, "\a", "BEL must not survive RenderJSON")
}

// escapeJSON quotes a payload so it can be embedded as a JSON string. We
// intentionally do not pull in encoding/json here so the test reflects the
// raw bytes a server response would carry.
func escapeJSON(s string) string {
	out := make([]byte, 0, len(s)+8)
	for _, b := range []byte(s) {
		switch b {
		case '"':
			out = append(out, '\\', '"')
		case '\\':
			out = append(out, '\\', '\\')
		default:
			if b < 0x20 {
				out = append(out, []byte{'\\', 'u', '0', '0',
					hex(b >> 4), hex(b & 0xf)}...)
			} else {
				out = append(out, b)
			}
		}
	}
	return string(out)
}

func hex(b byte) byte {
	if b < 10 {
		return '0' + b
	}
	return 'a' + (b - 10)
}
