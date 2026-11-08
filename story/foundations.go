package story

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func foundationsDefinitions() []Definition {
	d := []Definition{
		widget("foundations/theme", "Foundations/Theme", "Foundations", "theme", "Semantic roles, terminal palettes, contrast and no-color output.", "themeStory", nil, themeStory),
		widget("foundations/layout", "Foundations/Layout", "Foundations", "layout", "Cell-based panels and responsive allocation.", "layoutStory", nil, layoutStory),
		widget("foundations/glyphs", "Foundations/Glyphs", "Foundations", "theme", "Named semantic glyphs, explicit ASCII equivalents and adjacent labels.", "glyphStory", nil, glyphStory),
		widget("foundations/focus", "Foundations/Focus and spacing", "Foundations", "theme", "One cell on both sides of controls and panel content. Focus and emphasis retain the same allocation.", "focusStory", nil, focusStory),
		widget("feedback/motion", "Foundations/Motion", "Foundations", "motion", "A Harmonica spring stepped explicitly by the owner; reduced motion resolves immediately.", "motionStory", []string{"amount"}, motionStory),
	}

	for i := range d {
		if d[i].ID == "feedback/motion" {
			d[i].Upstream = "https://github.com/charmbracelet/harmonica"
			d[i].Step = true
		}
	}
	return d
}

func themeStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		t := e.Config.Palette()

		rows := []string{"Semantic tokens / normal and dim", "", t.Text().Render("Primary content"), t.Subtle().Render("Supporting information"), t.Heading().Render("Accent / interaction"), ""}
		for _, tone := range []text.Tone{text.Neutral, text.Success, text.Info, text.Warning, text.Danger} {
			rows = append(rows, text.ToneStyle(tone, text.Normal, t).Render(fmt.Sprintf("%-9s", tone))+"  "+text.ToneStyle(tone, text.Dim, t).Render("dim "+string(tone)))
		}

		rows = append(rows,
			t.Text().Foreground(t.Unknown).Render("unknown  ")+"  "+t.Text().Foreground(t.DimUnknown).Render("dim unknown"),
			t.Text().Foreground(t.Inactive).Render("off      ")+"  "+t.Text().Foreground(t.DimInactive).Render("dim off"))

		rows = append(rows, "", "Categorical data / independent palette")

		var series []string
		for i, color := range t.Series {
			series = append(series, t.Text().Foreground(color).Render(fmt.Sprintf("%d ━━", i+1)))
		}

		rows = append(rows, strings.Join(series, "  "))

		return strings.Join(rows, "\n")
	}), nil
}

func glyphStory(e Environment) (runtime.Scene, error) {
	return static(func(layout.Size) string {
		t := e.Config.Palette()
		fallback := t
		fallback.ASCII = true
		rows := []string{"Meaning     U A     Meaning     U A", ""}

		var cells []string
		for _, glyph := range []theme.Glyph{theme.GlyphOK, theme.GlyphWarning, theme.GlyphError, theme.GlyphUnknown, theme.GlyphOff, theme.GlyphInfo, theme.GlyphSelected, theme.GlyphUnselected, theme.GlyphExpanded, theme.GlyphCollapsed, theme.GlyphUp, theme.GlyphDown, theme.GlyphLeft, theme.GlyphRight, theme.GlyphLoading, theme.GlyphUnavailable} {
			cells = append(cells, layout.Fit(fmt.Sprintf("%-11s %s %s", glyph, t.Glyph(glyph), fallback.Glyph(glyph)), layout.Size{Width: 18, Height: 1}))
		}

		for i := 0; i < len(cells); i += 2 {
			rows = append(rows, cells[i]+"  "+cells[i+1])
		}

		return strings.Join(rows, "\n")
	}), nil
}

func focusStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		t := e.Config.Palette()

		rows := []string{"Stable control insets / 1 cell per side", "", "Normal  |" + t.Control().Render("Value") + "|", "Focused |" + t.Focus().Render("Value") + "|", "Row     |" + t.Selection().Render("Value") + "|", ""}
		for _, emphasis := range []text.Emphasis{text.Bright, text.Normal, text.Dim} {
			rows = append(rows, fmt.Sprintf("%-7s |", emphasis)+text.Pill("Value", text.Info, t, emphasis)+"|")
		}

		rows = append(rows, "", layout.Panel("Content gutter", "Text starts one cell from the border.", layout.Size{Width: size.Width, Height: 3}, t, true))

		return strings.Join(rows, "\n")
	}), nil
}

func layoutStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		t := e.Config.Palette()
		left, right, split := layout.Split(size.Width, 1, 20, 30, .4)
		height := max(3, size.Height-4)
		a := layout.Panel("Navigation", "Cell-based dimensions\nBorders included in size", layout.Size{Width: left, Height: height}, t, true)
		b := layout.Panel("Content", "The right pane adapts to available space.\nNarrow terminals stack the same panels.", layout.Size{Width: right, Height: height}, t, false)
		if split {
			return title("Responsive layout", e) + layout.Row(a, " ", b)
		}
		height = max(3, height/2)
		return title("Stacked layout", e) + layout.Column(layout.Panel("Navigation", "Same content", layout.Size{Width: left, Height: height}, t, true), layout.Panel("Content", "Same state", layout.Size{Width: right, Height: height}, t, false))
	}), nil
}

func motionStory(e Environment) (runtime.Scene, error) {
	m := feedback.NewMotion(e.Config.Palette())
	return &scene{update: func(msg tea.Msg) tea.Cmd {
		if _, ok := msg.(runtime.StepMsg); ok {
			m.Step(e.Config.Amount)
			emit(e, map[string]any{"type": "motion", "position": m.Position})
		}
		return nil
	}, view: func(layout.Size) string {
		return title("Spring motion", e) + fmt.Sprintf("Position %.2f  /  Target %.0f\n\nVelocity %.2f", m.Position, e.Config.Amount, m.Velocity) + footer("Ctrl+N steps the spring. Reduced motion resolves immediately.", e)
	}}, nil
}
