package cmd

import (
	"strings"
	"testing"
	"unicode/utf8"
)

// TestGlobalFlagUsageFitsHelpBox guards against global flag descriptions that
// would wrap onto a second line in the custom help box rendered by help.go.
//
// The renderer in formatFlags() formats each flag row as:
//
//	"  " + flag_name_padded_to_24 + usage_first_line
//
// inside a box with Padding(0, 1) and Width(termWidth - 4). getTermWidth()
// clamps to [100, 120] in practice (the 40-col fallback only triggers when
// stdout isn't a tty and we have no width info — assume 100 for sizing).
//
// At 100 cols: content area = 96, prefix = 26, so the description budget is
// 96 - 26 = 70 chars before the line wraps.
func TestGlobalFlagUsageFitsHelpBox(t *testing.T) {
	const (
		termWidth     = 100
		boxOverhead   = 4  // 2 borders + 2 padding
		flagColumn    = 26 // "  " + 24-char left-justified flag column
		maxUsageChars = termWidth - boxOverhead - flagColumn
	)

	for _, f := range Command.Flags {
		names := f.Names()
		if len(names) == 0 {
			continue
		}
		if pf, ok := f.(interface{ IsVisible() bool }); ok && !pf.IsVisible() {
			continue
		}

		uf, ok := f.(interface{ GetUsage() string })
		if !ok {
			continue
		}
		usage := uf.GetUsage()
		if idx := strings.Index(usage, "\n"); idx > 0 {
			usage = usage[:idx]
		}

		if got := utf8.RuneCountInString(usage); got > maxUsageChars {
			t.Errorf("flag --%s usage is %d chars, exceeds %d-char budget for %d-col help box (would wrap)\n  usage: %q",
				names[0], got, maxUsageChars, termWidth, usage)
		}
	}
}
