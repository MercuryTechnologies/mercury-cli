package cmd

import (
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v3"
)

func TestNestColonGroups(t *testing.T) {
	t.Parallel()

	group := func(name string, leaves ...string) *cli.Command {
		cmd := &cli.Command{Name: name}
		for _, leaf := range leaves {
			cmd.Commands = append(cmd.Commands, &cli.Command{Name: leaf})
		}
		return cmd
	}

	tests := []struct {
		name  string
		input []*cli.Command
		want  []string
	}{
		{
			name:  "reuses an existing parent",
			input: []*cli.Command{group("invoices", "list"), group("invoices:attachments", "get")},
			want:  []string{"invoices", "invoices list", "invoices attachments", "invoices attachments get"},
		},
		{
			name:  "creates a missing parent once",
			input: []*cli.Command{group("accounts", "list"), group("spend:budgets", "list"), group("spend:receipts", "upload")},
			want:  []string{"accounts", "accounts list", "spend", "spend budgets", "spend budgets list", "spend receipts", "spend receipts upload"},
		},
		{
			name:  "nests a:b before a:b:c regardless of input order",
			input: []*cli.Command{group("a:b:c", "x"), group("a:b", "y")},
			want:  []string{"a", "a b", "a b y", "a b c", "a b c x"},
		},
		{
			name:  "creates every missing ancestor",
			input: []*cli.Command{group("a:b:c", "x")},
			want:  []string{"a", "a b", "a b c", "a b c x"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			root := &cli.Command{Name: "mercury", Commands: tt.input}
			nestColonGroups(root)
			assert.Equal(t, tt.want, commandPaths(root, ""))
		})
	}
}

func TestCommandTreeHasNoColonGroups(t *testing.T) {
	t.Parallel()

	for _, path := range commandPaths(Command, "") {
		assert.NotContains(t, path, ":")
	}
	spend := Command.Command("spend")
	require.NotNil(t, spend)
	assert.Equal(t, resourcesCategory, spend.Category)
	assert.NotNil(t, spend.Command("budgets"))

	receipts := spend.Command("receipts")
	require.NotNil(t, receipts)
	assert.Empty(t, receipts.Category)
	assert.NotEmpty(t, receipts.Usage)
}

func TestRootOverrides(t *testing.T) {
	t.Parallel()

	for _, sub := range Command.Commands {
		assert.NotEqual(t, "API RESOURCE", sub.Category, sub.Name)
	}
	cards := Command.Command("cards")
	require.NotNil(t, cards)
	assert.Equal(t, groupUsage["cards"], cards.Usage)

	var flags []string
	for _, f := range Command.Flags {
		flags = append(flags, f.Names()[0])
	}
	debug := slices.Index(flags, "debug")
	require.GreaterOrEqual(t, debug, 0)
	assert.Equal(t, "yes", flags[debug+1])

	environment := slices.Index(flags, "environment")
	require.GreaterOrEqual(t, environment, 0)
	envFlag, ok := Command.Flags[environment].(*cli.StringFlag)
	require.True(t, ok, "--environment is no longer a *cli.StringFlag, so its validator override is skipped")
	require.NotNil(t, envFlag.Validator)
	assert.NoError(t, envFlag.Validator("sandbox"))
	assert.Error(t, envFlag.Validator("bogus"))
}

// commandPaths lists every command under cmd depth-first, in order, as
// space-separated paths.
func commandPaths(cmd *cli.Command, prefix string) []string {
	var paths []string
	for _, sub := range cmd.Commands {
		path := strings.TrimSpace(prefix + " " + sub.Name)
		paths = append(paths, path)
		paths = append(paths, commandPaths(sub, path)...)
	}
	return paths
}
