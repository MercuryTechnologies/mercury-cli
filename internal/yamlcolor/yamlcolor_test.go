package yamlcolor

import (
	"regexp"
	"strings"
	"testing"
)

// goccy's printer folds leading whitespace into the colored span,
// which is invisible on a terminal but needs tolerating here.
func wrapped(prefix, content string) *regexp.Regexp {
	return regexp.MustCompile(regexp.QuoteMeta(prefix) + `\s*` + regexp.QuoteMeta(content) + regexp.QuoteMeta(ansiReset))
}

func TestColorWrapsKeysAndScalars(t *testing.T) {
	src := []byte(`---
name: Acme
count: 42
active: true
ratio: 3.14
`)

	got := string(Color(src))

	cases := []struct {
		name string
		re   *regexp.Regexp
	}{
		{"key", wrapped(ansiBoldBlue, "name")},
		{"string", wrapped(ansiGreen, "Acme")},
		{"integer", wrapped(ansiYellow, "42")},
		{"float", wrapped(ansiYellow, "3.14")},
		{"bool", wrapped(ansiCyan, "true")},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if !tc.re.MatchString(got) {
				t.Errorf("output missing match for %s\nfull output:\n%s", tc.re, got)
			}
		})
	}
}

func TestColorPreservesDocumentSeparator(t *testing.T) {
	src := []byte("---\nkey: value\n")
	got := string(Color(src))
	if !strings.HasPrefix(got, "---") {
		t.Errorf("expected output to start with ---, got:\n%s", got)
	}
}

func TestColorHandlesNested(t *testing.T) {
	src := []byte(`---
outer:
  inner: hello
  items:
    - 1
    - 2
`)
	got := string(Color(src))

	wants := []*regexp.Regexp{
		wrapped(ansiBoldBlue, "outer"),
		wrapped(ansiBoldBlue, "inner"),
		wrapped(ansiGreen, "hello"),
		wrapped(ansiYellow, "1"),
		wrapped(ansiYellow, "2"),
	}
	for _, w := range wants {
		if !w.MatchString(got) {
			t.Errorf("output missing match for %s\nfull output:\n%s", w, got)
		}
	}
}

func TestColorEmpty(t *testing.T) {
	if got := Color(nil); len(got) != 0 {
		t.Errorf("expected empty output for nil input, got %q", got)
	}
}

func TestColorPreservesTrailingNewline(t *testing.T) {
	src := []byte("---\nkey: value\n")
	got := Color(src)
	if len(got) == 0 || got[len(got)-1] != '\n' {
		t.Errorf("expected output to end with newline, got %q", got)
	}
}

func TestColorNoSpuriousTrailingNewline(t *testing.T) {
	src := []byte("---\nkey: value")
	got := Color(src)
	if len(got) > 0 && got[len(got)-1] == '\n' {
		t.Errorf("did not expect trailing newline when input had none, got %q", got)
	}
}

// goccy/go-yaml does not have special printing for nulls.
// If this changes in a future version, we'd like to know about it _before_ we ship!
func TestColorLeavesNullUnwrapped(t *testing.T) {
	got := string(Color([]byte("---\nkey: null\n")))
	for _, prefix := range []string{ansiBoldBlue, ansiGreen, ansiYellow, ansiCyan} {
		if wrapped(prefix, "null").MatchString(got) {
			t.Errorf("expected null value not to be wrapped, got %q. full output:\n%s", prefix, got)
		}
	}
}
