package collection

import (
	"fmt"
	"io"
	"slices"
	"strings"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/interaction"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

type commandItem struct {
	command Command
	matches []int
}

func (c commandItem) FilterValue() string { return c.command.Label }

type commandDelegate struct {
	theme            theme.Theme
	descriptions     bool
	highlightMatches bool
}

func (d *commandDelegate) Height() int {
	if d.descriptions {
		return 2
	}

	return 1
}
func (d *commandDelegate) Spacing() int                        { return 0 }
func (d *commandDelegate) Update(tea.Msg, *list.Model) tea.Cmd { return nil }
func (d *commandDelegate) Render(w io.Writer, model list.Model, index int, item list.Item) {
	entry, ok := item.(commandItem)
	if !ok {
		return
	}

	command := entry.command
	width := max(1, model.Width())
	label := text.Inline(command.Label)

	style := d.theme.Text()
	if command.DisabledReason != "" {
		style = d.theme.Subtle()
	}

	style = style.Width(width)

	actionStyle := style.Width(0).Bold(true)
	if command.DisabledReason == "" && d.theme.Accent != nil {
		actionStyle = actionStyle.Foreground(d.theme.Accent)
	}

	if command.DisabledReason != "" {
		actionStyle = actionStyle.Bold(false)
		if d.theme.Muted != nil {
			actionStyle = actionStyle.Foreground(d.theme.Muted)
		}
	}

	secondaryStyle := style.Bold(false)
	if d.theme.Muted != nil {
		secondaryStyle = secondaryStyle.Foreground(d.theme.Muted)
	}

	marker := "  "
	if model.Index() == index {
		marker = d.theme.Glyph(theme.GlyphRight) + " "
	}

	var highlighted strings.Builder

	for i, r := range []rune(label) {
		runeStyle := actionStyle
		if d.highlightMatches && slices.Contains(entry.matches, i) {
			runeStyle = runeStyle.Underline(true)
		}

		part := runeStyle.Render(string(r))

		if _, err := highlighted.WriteString(part); err != nil {
			return
		}
	}

	label = highlighted.String()

	suffix := text.Inline(command.State)
	if command.Shortcut != "" {
		if suffix != "" {
			suffix += " · "
		}

		suffix += text.Inline(command.Shortcut)
	}

	if width < paletteDescriptionWidth && command.Shortcut != "" {
		suffix = text.Inline(command.Shortcut)
	}

	if command.DisabledReason != "" {
		suffix = "unavailable"
	}

	suffixWidth := min(ansi.StringWidth(suffix), width/3)
	labelWidth := max(1, width-suffixWidth-4)
	line := style.Width(0).Render(marker) + ansi.Truncate(label, labelWidth, actionStyle.Render("…"))
	line += secondaryStyle.Width(0).Bold(true).Render(strings.Repeat(" ", max(0, width-ansi.StringWidth(line)-suffixWidth)) + ansi.Truncate(suffix, suffixWidth, ""))

	if d.descriptions {
		description := command.Description
		if command.DisabledReason != "" {
			description = command.DisabledReason
		}

		line += "\n" + secondaryStyle.Render("  "+ansi.Truncate(text.Inline(description), max(1, width-2), "…"))
	}

	if _, err := fmt.Fprint(w, line); err != nil {
		return
	}
}

func (m *CommandPalette) View() string {
	if !m.opened || m.size.Width <= 0 || m.size.Height <= 0 {
		return ""
	}

	width := max(1, m.Results.Width())

	results := m.Results.View()
	if len(m.Results.Items()) == 0 {
		results = layout.Fit(m.Theme.Subtle().Render("No matching commands"), layout.Size{Width: width, Height: m.Results.Height()})
	}

	position := 0
	if len(m.Results.Items()) > 0 {
		position = m.Results.Index() + 1
	}

	status := fmt.Sprintf("%d/%d · %d matches", position, len(m.Results.Items()), len(m.Results.Items()))
	if item, ok := m.Results.SelectedItem().(commandItem); ok && item.command.DisabledReason != "" && !m.delegate.descriptions {
		status = item.command.DisabledReason
	}

	execute := m.Keys.Execute.Help().Desc
	movement := m.Results.KeyMap.CursorUp.Help().Key + "/" + m.Results.KeyMap.CursorDown.Help().Key

	footer := text.KeyHint(movement, "select", m.Theme) + " · " + text.KeyHint(m.Keys.Execute.Help().Key, execute, m.Theme) + " · " + text.KeyHint(m.Keys.Close.Help().Key, "cancel", m.Theme)
	if m.keyLayout == interaction.VimKeys {
		switch m.mode {
		case PaletteSearch:
			status = "SEARCH · " + status
			footer = text.KeyHint("↑↓", "select", m.Theme) + " · " + text.KeyHint(m.Keys.Navigate.Help().Key, "normal", m.Theme) + " · " + text.KeyHint(m.Keys.Execute.Help().Key, "run", m.Theme)
		case PaletteNavigate:
			status = "NORMAL · " + status
			footer = text.KeyHint(movement, "move", m.Theme) + " · " + text.KeyHint(m.Keys.Search.Help().Key, "search", m.Theme) + " · " + text.KeyHint(m.Keys.Close.Help().Key, "close", m.Theme)
		default:
		}
	}

	content := m.Search.View() + "\n\n" + results + "\n" + m.Theme.Subtle().Render(ansi.Truncate(text.Inline(status), width, "…")) + "\n" + footer
	height := min(m.size.Height, m.Results.Height()+paletteChromeHeight+2)
	panel := layout.Panel("Commands", content, layout.Size{Width: m.size.Width, Height: height}, m.Theme, true)

	return m.Theme.Text().Background(m.Theme.Surface).Render(panel)
}
