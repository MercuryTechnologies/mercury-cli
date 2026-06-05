package cmd

import _ "embed"

// applyTemplate is a blank, annotated apply.yaml template printed by
// `mercury apply --template`. It is embedded so the binary output and the
// committed examples/apply.template.yaml stay a single source of truth.
//
//go:embed apply.template.yaml
var applyTemplate string
