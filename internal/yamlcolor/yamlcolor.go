// Adds ANSI syntax highlighting to YAML output.
package yamlcolor

import (
	"github.com/goccy/go-yaml/lexer"
	"github.com/goccy/go-yaml/printer"
)

// A palette that roughly mirrors tidwall/pretty's TerminalStyle, 
// so colored YAML and colored JSON look consistent side-by-side.
// (a caveat: tidwall/pretty has special printing for null values,
// but goccy/go-yaml does not.)
const (
	ansiReset     = "\x1b[0m"
	ansiBoldBlue  = "\x1b[1m\x1b[94m"
	ansiGreen     = "\x1b[32m"
	ansiYellow    = "\x1b[33m"
	ansiCyan      = "\x1b[36m"
)

func prop(prefix string) *printer.Property {
	return &printer.Property{Prefix: prefix, Suffix: ansiReset}
}

// Returns src with ANSI escape codes wrapping YAML map keys and scalar
// values. Colorizes unconditionally, it's the caller's job to decide
// whether the terminal actually supports color.
func Color(src []byte) []byte {
	p := printer.Printer{
		MapKey: func() *printer.Property { return prop(ansiBoldBlue) },
		String: func() *printer.Property { return prop(ansiGreen) },
		Number: func() *printer.Property { return prop(ansiYellow) },
		Bool:   func() *printer.Property { return prop(ansiCyan) },
	}
	out := []byte(p.PrintTokens(lexer.Tokenize(string(src))))
	// goccy's printer drops the trailing newline of its input.
	// We'd prefer to preserve it, since callers concatenate multiple
	// YAML docs and rely on newline separation between them.
	if len(src) > 0 && src[len(src)-1] == '\n' && (len(out) == 0 || out[len(out)-1] != '\n') {
		out = append(out, '\n')
	}
	return out
}
