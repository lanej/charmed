package snippets

import (
	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/theme"
)

func ExampleFoundationsTokens() string {
	t := theme.New(theme.Dark)

	// Keep shape and color separate; native Lip Gloss styles stay available.
	frame := lipgloss.NewStyle().
		Border(t.FrameBorder()).
		BorderForeground(t.BorderStyle(true).GetForeground()).
		Padding(theme.SpaceNone, theme.InlinePadding)

	selected := t.Selection().Render("Selected item")
	control := t.Control().Render(t.Glyph(theme.GlyphSearch) + " Search")

	return frame.Render(control + "\n" + selected)
}
