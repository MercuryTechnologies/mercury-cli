package cmd

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
	"github.com/urfave/cli/v3"
)

var (
	starChars = []string{"✦", "·", "⋆", "✧", "∘", "⊹", "˚", "°"}

	// Dutch still life palette
	mercuryBlue    = lipgloss.NewStyle().Foreground(lipgloss.Color("#395AFF"))
	mercuryBlueDim = lipgloss.NewStyle().Foreground(lipgloss.Color("#1E3299"))
	mercuryBlueMid = lipgloss.NewStyle().Foreground(lipgloss.Color("#2B46CC"))
	warmWhite  = lipgloss.NewStyle().Foreground(lipgloss.Color("223"))
	creamDim   = lipgloss.NewStyle().Foreground(lipgloss.Color("#1E3299"))
	lavender   = lipgloss.NewStyle().Foreground(lipgloss.Color("139"))

	helpTitle = lipgloss.NewStyle().Foreground(lipgloss.Color("#395AFF")).Bold(true)
	helpDim   = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	helpCmd   = lipgloss.NewStyle().Foreground(lipgloss.Color("#395AFF"))
	helpFlag  = lipgloss.NewStyle().Foreground(lipgloss.Color("#5B7AFF"))
	helpDesc  = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	helpFrame = lipgloss.NewStyle().Foreground(lipgloss.Color("#1E3299"))
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

func randomStar() string {
	s := starChars[rand.Intn(len(starChars))]
	switch rand.Intn(7) {
	case 0:
		return mercuryBlue.Render(s)
	case 1:
		return warmWhite.Render(s)
	case 2, 3:
		return mercuryBlueMid.Render(s)
	default:
		return mercuryBlueDim.Render(s)
	}
}

// padTo pads a plain string to exactly w display columns with trailing spaces.
func padTo(s string, w int) string {
	sw := lipgloss.Width(s)
	if sw >= w {
		return s
	}
	return s + strings.Repeat(" ", w-sw)
}

// framedLine renders ║<content padded to innerW>║
func framedLine(content string, innerW int) string {
	return helpFrame.Render("║") + padTo(content, innerW) + helpFrame.Render("║")
}

func starField(width, density int) string {
	parts := make([]string, width)
	for i := range parts {
		parts[i] = " "
	}
	count := density + rand.Intn(density/2+1)
	for i := 0; i < count; i++ {
		parts[rand.Intn(width)] = randomStar()
	}
	return strings.Join(parts, "")
}

func banner(width int) string {
	innerW := width - 2

	// Roman columns
	colL := []string{
		" ╭┈┈┈╮ ",
		" ┆⌇⌇⌇┆ ",
		" │║║║│ ",
		" │║║║│ ",
		" │║║║│ ",
		" │║║║│ ",
		" │║║║│ ",
		" │║║║│ ",
		" ╰───╯ ",
	}
	colR := []string{
		" ╭┈┈┈╮ ",
		" ┆⌇⌇⌇┆ ",
		" │║║║│ ",
		" │║║║│ ",
		" │║║║│ ",
		" │║║║│ ",
		" │║║║│ ",
		" │║║║│ ",
		" ╰───╯ ",
	}

	mercuryText := []string{
		` ███╗   ███╗███████╗██████╗  ██████╗██╗   ██╗██████╗ ██╗   ██╗`,
		` ████╗ ████║██╔════╝██╔══██╗██╔════╝██║   ██║██╔══██╗╚██╗ ██╔╝`,
		` ██╔████╔██║█████╗  ██████╔╝██║     ██║   ██║██████╔╝ ╚████╔╝ `,
		` ██║╚██╔╝██║██╔══╝  ██╔══██╗██║     ██║   ██║██╔══██╗  ╚██╔╝  `,
		` ██║ ╚═╝ ██║███████╗██║  ██║╚██████╗╚██████╔╝██║  ██║   ██║   `,
		` ╚═╝     ╚═╝╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚═════╝ ╚═╝  ╚═╝   ╚═╝   `,
	}

	colW := lipgloss.Width(colL[0])
	centerW := innerW - colW*2

	// Layout: 3 star rows above, 6 text rows, 3 star rows below = 12 interior rows
	// Columns span all 12 rows: capital at top, base at bottom, shaft fills middle
	starsAbove := 3
	textLines := len(mercuryText)
	starsBelow := 3
	totalRows := starsAbove + textLines + starsBelow

	// Build column strings for each row dynamically
	colShaftL := " │║║║│ "
	colShaftR := " │║║║│ "
	getColL := func(row int) string {
		if row == 0 {
			return colL[0] // capital top
		}
		if row == 1 {
			return colL[1] // capital bottom
		}
		if row == totalRows-2 {
			return colL[1] // base top (reuse capital detail)
		}
		if row == totalRows-1 {
			return colL[len(colL)-1] // base
		}
		return colShaftL
	}
	getColR := func(row int) string {
		if row == 0 {
			return colR[0]
		}
		if row == 1 {
			return colR[1]
		}
		if row == totalRows-2 {
			return colR[1]
		}
		if row == totalRows-1 {
			return colR[len(colR)-1]
		}
		return colShaftR
	}

	var sb strings.Builder

	// Frame top
	sb.WriteString(helpFrame.Render("╔"+strings.Repeat("═", innerW)+"╗") + "\n")

	starDensities := []int{20, 15, 12}

	for row := 0; row < totalRows; row++ {
		left := getColL(row)
		right := getColR(row)

		var center string
		ti := row - starsAbove
		if ti >= 0 && ti < textLines {
			center = mercuryText[ti]
		}

		// Center the text in the middle zone, fill empty space with stars
		cw := lipgloss.Width(center)
		lpad := (centerW - cw) / 2
		if lpad < 0 {
			lpad = 0
		}
		rpad := centerW - cw - lpad
		if rpad < 0 {
			rpad = 0
		}

		// Vary star density: denser at edges, sparser near text
		density := 12
		if row < starsAbove {
			density = starDensities[row]
		} else if row >= starsAbove+textLines {
			density = starDensities[totalRows-1-row]
		}
		_ = density

		leftStars := starField(lpad, lpad/7+1)
		rightStars := starField(rpad, rpad/7+1)
		centeredText := leftStars + mercuryBlue.Render(center) + rightStars

		line := creamDim.Render(left) + centeredText + creamDim.Render(right)
		sb.WriteString(framedLine(line, innerW) + "\n")
	}

	// Frame bottom
	sb.WriteString(helpFrame.Render("╚"+strings.Repeat("═", innerW)+"╝") + "\n")

	return sb.String()
}

func boxSection(title string, content string, width int) string {
	border := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#1E3299")).
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
		if pf, ok := f.(interface{ IsHidden() bool }); ok && pf.IsHidden() {
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

		padded := fmt.Sprintf("  %-30s", flagStr)
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

	usage := fmt.Sprintf("  %s <command> [OPTIONS]\n  %s <resource> <subcommand> [OPTIONS]",
		helpCmd.Render(cmd.Name), helpCmd.Render(cmd.Name))
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

	for _, sub := range cmd.Commands {
		if sub.Hidden {
			continue
		}
		cat := sub.Category
		if cat == "" {
			cat = "Commands"
		}
		padded := fmt.Sprintf("  %-30s", sub.Name)
		entry := helpCmd.Render(padded) + helpDesc.Render(sub.Usage)
		if idx, ok := groupMap[cat]; ok {
			groups[idx].commands = append(groups[idx].commands, entry)
		} else {
			groupMap[cat] = len(groups)
			groups = append(groups, group{title: cat, commands: []string{entry}})
		}
	}

	for _, g := range groups {
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
