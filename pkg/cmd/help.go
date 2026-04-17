package cmd

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
	"github.com/urfave/cli/v3"
)

var (
	mercuryBlue = lipgloss.NewStyle().Foreground(colorBlue)

	helpTitle = lipgloss.NewStyle().Foreground(colorBlue).Bold(true)
	helpDim   = lipgloss.NewStyle().Foreground(colorDim)
	helpCmd   = lipgloss.NewStyle().Foreground(colorBlue)
	helpFlag  = lipgloss.NewStyle().Foreground(colorBlue)
	helpDesc  = lipgloss.NewStyle().Foreground(colorLight)
	helpFrame = lipgloss.NewStyle().Foreground(colorFrame)

	logoLines = []string{
		`⠀⠀⠀⠀⠀⠀⠀⣀⡤⠶⠒⠛⠛⠉⠛⠛⢶⡶⠤⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
		`⠀⠀⠀⠀⣠⠔⠋⠁⠀⣀⡤⢤⡴⠶⠦⡄⠀⠹⡄⠀⠉⠳⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
		`⠀⠀⢠⠞⠁⠀⣠⠖⢻⡁⠀⢾⠀⠀⠀⣹⠀⠀⡿⠓⢤⡀⠀⠱⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
		`⠀⣰⠋⠀⢠⠞⠁⠀⠈⢧⡀⠈⠓⠲⠶⠧⢄⣰⠃⠀⠀⠙⢆⠀⠈⢧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
		`⢠⣇⡴⠒⠛⠛⠓⠲⡴⠋⠙⢦⣤⣤⣄⡀⠀⠈⢳⠴⠒⠛⠙⢧⠀⠈⣇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⣶⣶⡆⠀⠀⠀⠀⣶⣶⣶⠀⠀⠀⠀⢰⣶⣶⣶⣶⣶⣶⠀⠀⠀⠀⢰⣶⣶⣶⣶⣶⣄⠀⠀⠀⠀⠀⢀⣤⣶⡶⠶⢶⣶⣄⠀⠀⠀⠀⣶⡆⠀⠀⠀⠀⣶⡆⠀⠀⠀⠀⢰⣶⣶⣶⣶⣦⡄⠀⠀⠀⠰⣶⣄⠀⠀⠀⣠⣶⠆`,
		`⣾⠋⠀⣠⠴⠶⢤⡼⠁⠀⡴⠋⠀⠀⠀⠉⢳⣴⠃⠀⣠⠴⠶⢼⡆⠀⢹⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⠹⣿⡄⠀⠀⣼⡿⢻⣿⠀⠀⠀⠀⢸⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⡇⠀⠀⢈⣿⡇⠀⠀⠀⢠⣿⡟⠁⠀⠀⠀⠈⠀⠀⠀⠀⠀⣿⡇⠀⠀⠀⠀⣿⡇⠀⠀⠀⠀⢸⣿⠀⠀⠀⣹⣿⠀⠀⠀⠀⠘⢿⣦⠀⣴⡿⠃⠀`,
		`⣿⠀⠀⡇⠀⠀⢀⡇⠀⢸⡇⠀⠀⠀⠀⠀⠀⡇⠀⢰⡇⠀⠀⢈⡇⠀⢸⠆⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⠀⢻⣷⠀⢰⣿⠁⢸⣿⠀⠀⠀⠀⢸⣿⡷⠶⠶⠶⠷⠀⠀⠀⠀⢸⣿⡷⢶⣶⡿⠛⠀⠀⠀⠀⢸⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⡇⠀⠀⠀⠀⣿⡇⠀⠀⠀⠀⢸⣿⠶⣶⣾⠟⠋⠀⠀⠀⠀⠀⠈⢻⣿⡟⠁⠀⠀`,
		`⢿⠀⠀⣟⠒⠒⠋⠀⢀⡼⠳⣄⠀⠀⠀⢀⡼⠁⠀⡼⠙⠒⠒⠋⠀⢀⣾⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⠀⠀⢿⣷⣿⠃⠀⢸⣿⠀⠀⠀⠀⢸⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⡇⠀⠻⣷⡄⠀⠀⠀⠀⠘⣿⣧⡀⠀⠀⠀⢀⡀⠀⠀⠀⠀⣿⣇⠀⠀⠀⢠⣿⠇⠀⠀⠀⠀⢸⣿⠀⠈⢿⣧⡀⠀⠀⠀⠀⠀⠀⢸⣿⡇⠀⠀⠀`,
		`⠘⡆⠀⠸⣦⣤⡤⠴⠻⣄⠀⠈⠉⠛⠛⠻⢤⣀⠞⠳⢤⣤⣤⡤⠖⢋⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⠿⠀⠀⠈⠿⠏⠀⠀⠸⠿⠀⠀⠀⠀⠸⠿⠿⠿⠿⠿⠿⠆⠀⠀⠀⠸⠿⠇⠀⠀⠹⠿⠆⠀⠀⠀⠀⠈⠛⠿⢷⣶⡾⠿⠋⠀⠀⠀⠀⠈⠻⠷⣶⡾⠿⠋⠀⠀⠀⠀⠀⠼⠿⠀⠀⠈⠿⠷⠀⠀⠀⠀⠀⠀⠸⠿⠇⠀⠀⠀`,
		`⠀⠹⣄⠀⠘⢦⡀⠀⢀⡞⠙⢒⡶⠶⢤⡀⠀⠹⡄⠀⠀⣠⠎⠀⢀⡞⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
		`⠀⠀⠘⢦⡀⠀⠙⠦⣼⡁⠀⢼⠀⠀⠀⣹⠀⠀⣷⡤⠞⠁⠀⡰⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
		`⠀⠀⠀⠀⠙⠢⣄⡀⠈⢧⠀⠈⠓⠶⠶⠛⠚⠋⠁⠀⣀⠴⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
		`⠀⠀⠀⠀⠀⠀⠀⠉⠓⠶⠷⣦⣤⣀⣠⣤⠤⠴⠒⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
	}
)

func getTermWidth() int {
	w, _, err := term.GetSize(os.Stdout.Fd())
	if err != nil || w < 40 {
		return 100
	}
	if w > 120 {
		return 120
	}
	return w
}

func banner(width int) string {
	var sb strings.Builder
	sb.WriteString("\n")
	for _, line := range logoLines {
		lw := lipgloss.Width(line)
		pad := (width - lw) / 2
		if pad < 0 {
			pad = 0
		}
		sb.WriteString(strings.Repeat(" ", pad) + mercuryBlue.Render(line) + "\n")
	}
	return sb.String()
}

func boxSection(title string, content string, width int) string {
	border := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorFrame).
		Padding(0, 1).
		Width(width - 4)

	header := helpTitle.Render("─── " + title + " ")
	remaining := width - lipgloss.Width(header)
	if remaining > 0 {
		header += helpDim.Render(strings.Repeat("─", remaining))
	}

	return header + "\n" + border.Render(content) + "\n"
}

func formatFlags(flags []cli.Flag) string {
	var lines []string
	for _, f := range flags {
		names := f.Names()
		if len(names) == 0 {
			continue
		}
		if pf, ok := f.(interface{ IsVisible() bool }); ok && !pf.IsVisible() {
			continue
		}

		var flagStr string
		if len(names) > 1 {
			flagStr = fmt.Sprintf("-%s, --%s", names[1], names[0])
		} else {
			flagStr = fmt.Sprintf("    --%s", names[0])
		}

		usage := ""
		if uf, ok := f.(interface{ GetUsage() string }); ok {
			usage = uf.GetUsage()
			if idx := strings.Index(usage, "\n"); idx > 0 {
				usage = usage[:idx]
			}
		}

		padded := fmt.Sprintf("  %-36s", flagStr)
		lines = append(lines, helpFlag.Render(padded)+helpDesc.Render(usage))
	}
	return strings.Join(lines, "\n")
}

func renderCustomHelp(w io.Writer, cmd *cli.Command) {
	width := getTermWidth()

	var sb strings.Builder

	sb.WriteString("\n")
	sb.WriteString(banner(width))
	sb.WriteString("\n")

	usage := fmt.Sprintf("  %s <resource> <subcommand> [OPTIONS]",
		helpCmd.Render(cmd.Name))
	sb.WriteString(boxSection("Usage", usage, width))
	sb.WriteString("\n")

	if len(cmd.Flags) > 0 {
		sb.WriteString(boxSection("Options", formatFlags(cmd.Flags), width))
		sb.WriteString("\n")
	}

	type group struct {
		title    string
		commands []string
	}
	groups := []group{}
	groupMap := map[string]int{}
	for _, cat := range []string{"Auth", "Resources", "Utility"} {
		groupMap[cat] = len(groups)
		groups = append(groups, group{title: cat})
	}

	sorted := make([]*cli.Command, len(cmd.Commands))
	copy(sorted, cmd.Commands)
	slices.SortFunc(sorted, func(a, b *cli.Command) int {
		return strings.Compare(a.Name, b.Name)
	})

	for _, sub := range sorted {
		if sub.Hidden {
			continue
		}
		cat := sub.Category
		if cat == "" {
			cat = "Commands"
		}
		padded := fmt.Sprintf("  %-36s", sub.Name)
		entry := helpCmd.Render(padded) + helpDesc.Render(sub.Usage)
		if idx, ok := groupMap[cat]; ok {
			groups[idx].commands = append(groups[idx].commands, entry)
		} else {
			groupMap[cat] = len(groups)
			groups = append(groups, group{title: cat, commands: []string{entry}})
		}
	}

	for _, g := range groups {
		if len(g.commands) == 0 {
			continue
		}
		sb.WriteString(boxSection(g.title, strings.Join(g.commands, "\n"), width))
		sb.WriteString("\n")
	}

	if cmd.Version != "" {
		sb.WriteString(helpDim.Render(fmt.Sprintf("  v%s", cmd.Version)) + "\n\n")
	}

	fmt.Fprint(w, sb.String())
}

func init() {
	origPrinter := cli.HelpPrinter
	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
		if cmd, ok := data.(*cli.Command); ok && cmd.Root() == cmd {
			renderCustomHelp(w, cmd)
			return
		}
		origPrinter(w, templ, data)
	}
}
