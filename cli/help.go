// Package cli formats command-line help from native Cobra commands and flags.
package cli

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/colorprofile"
	"github.com/charmbracelet/x/ansi"
	"github.com/charmbracelet/x/term"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/lanej/charmed/theme"
)

const defaultWidth = 80

// HelpFormatter renders help from the command tree at the time of the request.
// An unset Width uses the destination terminal width, or 80 cells for Format
// and redirected output. Format defaults to the dark theme. Install selects
// the terminal's light or dark palette when its background can be queried.
type HelpFormatter struct {
	Width int
	Theme theme.Theme
}

// Install sets the command's inherited Cobra help function. Children with
// their own help function retain it. Execute the root with Cobra; fang.Execute
// installs its own root help function. Write errors go to Cobra's error stream.
func (f HelpFormatter) Install(cmd *cobra.Command) {
	cmd.SetHelpFunc(func(target *cobra.Command, _ []string) {
		output := target.OutOrStdout()
		format := f

		writer := colorprofile.NewWriter(output, os.Environ())
		if os.Getenv("NO_COLOR") != "" {
			writer.Profile = colorprofile.NoTTY
		}

		if terminal, ok := output.(term.File); ok && term.IsTerminal(terminal.Fd()) {
			if format.Width <= 0 {
				if width, _, err := term.GetSize(terminal.Fd()); err == nil {
					format.Width = width
				}
			}

			input, ok := target.InOrStdin().(term.File)
			if ok && !format.hasTheme() && writer.Profile > colorprofile.ASCII && !lipgloss.HasDarkBackground(input, terminal) {
				format.Theme = theme.New(theme.Light)
			}
		}

		if _, err := fmt.Fprintln(writer, format.Format(target)); err != nil {
			target.PrintErrln(err)
		}
	})
}

// Format returns the full help text without executing the command or writing
// to its streams. Use theme.Mono for plain text, or Install for terminal output.
func (f HelpFormatter) Format(cmd *cobra.Command) string {
	if f.Width <= 0 {
		f.Width = defaultWidth
	}

	if !f.hasTheme() {
		f.Theme = theme.New(theme.Dark)
	}

	var blocks []string

	if description := strings.TrimSpace(cmp.Or(cmd.Long, cmd.Short)); description != "" {
		blocks = append(blocks, f.Theme.Text().Render(ansi.Wrap(description, f.Width, "")))
	}

	var usage []string
	if cmd.Runnable() || !cmd.HasAvailableSubCommands() {
		usage = append(usage, cmd.UseLine())
	}

	if cmd.HasAvailableSubCommands() {
		usage = append(usage, cmd.CommandPath()+" [command]")
	}

	blocks = append(blocks, f.section("Usage", f.indented(f.Theme.Text().Bold(true).Render(strings.Join(usage, "\n")), theme.SpaceMedium)))
	if len(cmd.Aliases) > 0 {
		blocks = append(blocks, f.section("Aliases", f.indented(cmd.NameAndAliases(), theme.SpaceMedium)))
	}

	if cmd.Example != "" {
		blocks = append(blocks, f.section("Examples", f.indented(strings.Trim(cmd.Example, "\n"), theme.SpaceMedium)))
	}

	groups := make(map[string][]helpRow)

	var topics []helpRow

	for _, child := range cmd.Commands() {
		switch {
		case child.Hidden || child.Deprecated != "":
			continue
		case child.IsAdditionalHelpTopicCommand():
			topics = append(topics, helpRow{label: child.Name(), description: child.Short})
		default:
			groups[child.GroupID] = append(groups[child.GroupID], helpRow{label: child.Use, description: child.Short})
		}
	}

	for _, group := range cmd.Groups() {
		blocks = f.addRows(blocks, strings.TrimSuffix(group.Title, ":"), groups[group.ID])
	}

	blocks = f.addRows(blocks, "Commands", groups[""])
	blocks = f.addRows(blocks, "Flags", flagRows(cmd.LocalFlags()))
	blocks = f.addRows(blocks, "Global flags", flagRows(cmd.InheritedFlags()))
	blocks = f.addRows(blocks, "Help topics", topics)
	result := strings.Join(blocks, "\n\n")

	if f.Theme.Mode == theme.Mono {
		return ansi.Strip(result)
	}

	return result
}

func (f HelpFormatter) hasTheme() bool {
	return f.Theme.Mode != "" || f.Theme.Foreground != nil || f.Theme.Accent != nil || f.Theme.Muted != nil
}

type helpRow struct{ label, argument, description string }

func (row helpRow) labelWidth() int {
	return ansi.StringWidth(row.label + row.argument)
}

func flagRows(flags *pflag.FlagSet) []helpRow {
	var rows []helpRow

	flags.VisitAll(func(flag *pflag.Flag) {
		if flag.Hidden {
			return
		}

		label := "--" + flag.Name
		if flag.Shorthand != "" && flag.ShorthandDeprecated == "" {
			label = "-" + flag.Shorthand + ", " + label
		}

		metavar, description := pflag.UnquoteUsage(flag)

		var argument string
		if metavar != "" {
			argument = " " + metavar
		}

		kind := flag.Value.Type()
		boolean := kind == "bool" || kind == "boolfunc"
		optional := flag.NoOptDefVal

		implicit := boolean && optional == "true" || kind == "count" && optional == "+1"
		if optional != "" && !implicit {
			argument += "[=" + flagValue(kind, optional) + "]"
		}

		if slices.Contains(flag.Annotations[cobra.BashCompOneRequiredFlag], "true") {
			description += " (required)"
		}

		if value := flag.DefValue; value != "" && (!boolean || value != "false") {
			description += " (default " + flagValue(kind, value) + ")"
		}

		if flag.Deprecated != "" {
			description += " (deprecated: " + flag.Deprecated + ")"
		}

		rows = append(rows, helpRow{label: label, argument: argument, description: strings.TrimSpace(description)})
	})

	return rows
}

func flagValue(kind, value string) string {
	if kind == "string" {
		return strconv.Quote(value)
	}

	return value
}

func (f HelpFormatter) section(title, body string) string {
	return f.Theme.Heading().Render(ansi.Wrap(title, f.Width, "")) + "\n" + body
}

func (f HelpFormatter) indented(s string, indent int) string {
	// Leave room for a wide grapheme when the whole terminal is only a few cells.
	indent = min(indent, max(0, f.Width-2))
	prefix := strings.Repeat(" ", indent)
	wrapped := ansi.Wrap(s, f.Width-indent, "")

	return prefix + strings.ReplaceAll(wrapped, "\n", "\n"+prefix)
}

func (f HelpFormatter) addRows(blocks []string, title string, rows []helpRow) []string {
	if len(rows) == 0 {
		return blocks
	}

	const indent, gap, minDescription = theme.SpaceMedium, theme.SpaceMedium, 22

	labelBudget := f.Width - indent - gap - minDescription

	var widest int

	for _, row := range rows {
		if width := row.labelWidth(); row.description != "" && width <= labelBudget {
			widest = max(widest, width)
		}
	}

	column := indent + widest + gap
	lines := make([]string, 0, len(rows))

	for _, row := range rows {
		label := f.Theme.Text().Bold(true).Render(row.label)
		if row.argument != "" {
			label += f.Theme.Text().Render(row.argument)
		}

		width := row.labelWidth()
		if width > labelBudget || row.description == "" {
			lines = append(lines, f.indented(label, indent))
			if row.description != "" {
				lines = append(lines, f.indented(f.Theme.Subtle().Render(row.description), indent+gap))
			}

			continue
		}

		prefix := strings.Repeat(" ", indent) + label + strings.Repeat(" ", widest-width+gap)
		description := ansi.Wrap(f.Theme.Subtle().Render(row.description), f.Width-column, "")
		lines = append(lines, prefix+strings.ReplaceAll(description, "\n", "\n"+strings.Repeat(" ", column)))
	}

	return append(blocks, f.section(title, strings.Join(lines, "\n")))
}
