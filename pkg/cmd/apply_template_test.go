package cmd

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"

	"github.com/MercuryTechnologies/mercury-cli/internal/requestflag"
	"github.com/goccy/go-yaml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestApplyTemplateIsValidYAML verifies the embedded template parses as YAML
// into a non-empty top-level map. This guards against the committed template
// drifting into something that wouldn't round-trip through `mercury apply`.
func TestApplyTemplateIsValidYAML(t *testing.T) {
	t.Parallel()

	require.NotEmpty(t, applyTemplate, "embedded applyTemplate must not be empty")

	var m map[string]any
	err := yaml.Unmarshal([]byte(applyTemplate), &m)
	require.NoError(t, err, "embedded applyTemplate must parse as YAML")
	require.NotEmpty(t, m, "parsed applyTemplate must have at least one top-level key")
}

// TestApplyTemplateMatchesExamplesCopy verifies the embedded
// pkg/cmd/apply.template.yaml is byte-identical to examples/apply.template.yaml.
// The two are intentionally kept as a single source of truth: the binary prints
// the embedded copy, while the examples copy exists for discoverability. The
// test working directory is the package dir (pkg/cmd), so the examples file is
// reached via a relative path two levels up.
func TestApplyTemplateMatchesExamplesCopy(t *testing.T) {
	t.Parallel()

	examplesCopy, err := os.ReadFile("../../examples/apply.template.yaml")
	require.NoError(t, err, "examples/apply.template.yaml must be readable")

	assert.Equal(t, applyTemplate, string(examplesCopy),
		"embedded pkg/cmd/apply.template.yaml must be byte-identical to examples/apply.template.yaml")
}

// TestApplyTemplateKeysAreValidBodyParams is a drift guard: every top-level key
// in the template must correspond to a real body parameter of the apply command.
// The valid set is derived from onboardingApply.Flags — for each flag
// implementing requestflag.InRequest, its (non-empty) GetBodyPath() is the
// canonical API name the template uses (e.g. "beneficialOwners", "about",
// "formationDetails"). If the schema changes and a flag/body path is renamed or
// removed, this catches the template referencing a key the API no longer accepts.
func TestApplyTemplateKeysAreValidBodyParams(t *testing.T) {
	t.Parallel()

	validBodyPaths := map[string]bool{}
	for _, f := range onboardingApply.Flags {
		if inReq, ok := f.(requestflag.InRequest); ok {
			if bodyPath := inReq.GetBodyPath(); bodyPath != "" {
				validBodyPaths[bodyPath] = true
			}
		}
	}
	require.NotEmpty(t, validBodyPaths,
		"expected onboardingApply to expose at least one body-path flag")

	var m map[string]any
	require.NoError(t, yaml.Unmarshal([]byte(applyTemplate), &m))

	for key := range m {
		assert.Truef(t, validBodyPaths[key],
			"template key %q is not a valid apply body parameter (valid: %v)", key, keys(validBodyPaths))
	}
}

// TestApplyTemplateFlagPrintsAndMakesNoCall verifies that running the apply
// command with --template writes the template to stdout and returns nil without
// any auth or network access. onboardingApplyOverride short-circuits on the
// template flag before constructing a client, so this runs fully in-process.
func TestApplyTemplateFlagPrintsAndMakesNoCall(t *testing.T) {
	// Not parallel: this test swaps os.Stdout, which is process-global.
	out := captureStdout(t, func() {
		err := onboardingApply.Run(context.Background(), []string{"apply", "--template"})
		require.NoError(t, err, "running apply --template must not error")
	})

	require.NotEmpty(t, out, "apply --template must write the template to stdout")
	assert.Equal(t, applyTemplate, out,
		"apply --template output must match the embedded template exactly")
}

// captureStdout redirects os.Stdout for the duration of fn and returns whatever
// was written. It restores the original stdout before returning.
func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	r, w, err := os.Pipe()
	require.NoError(t, err)

	orig := os.Stdout
	os.Stdout = w
	defer func() { os.Stdout = orig }()

	done := make(chan string, 1)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		done <- buf.String()
	}()

	fn()

	require.NoError(t, w.Close())
	return <-done
}

// keys returns the keys of a set, for readable assertion failure messages.
func keys(set map[string]bool) []string {
	out := make([]string, 0, len(set))
	for k := range set {
		out = append(out, k)
	}
	return out
}
