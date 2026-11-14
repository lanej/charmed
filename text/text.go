// Package text contains text atoms and a Glamour Markdown adapter.
package text

import (
	"github.com/lanej/charmed/theme"
)

func Label(value string, t theme.Theme) string { return t.Text().Render(value) }
func Badge(value string, t theme.Theme) string { return t.Heading().Render("[" + value + "]") }

type Tone string

const (
	Neutral Tone = "neutral"
	Success Tone = "success"
	Info    Tone = "info"
	Warning Tone = "warning"
	Danger  Tone = "danger"
)

// Pill tint strengths are the semantic color fraction mixed into the background.
const (
	normalPillTint = 0.15
	dimPillTint    = 0.04
)

// Pill keeps its text label in every profile; color supplements its meaning.
// Semantic fills use contrasting theme surfaces, with brackets in ASCII mode.
func Pill(value string, tone Tone, t theme.Theme, emphasis ...Emphasis) string {
	level := Bright // Preserve existing filled pills for three-argument calls.
	if len(emphasis) > 0 {
		level = emphasis[0]
	}

	if t.ASCII {
		return ToneStyle(tone, level, t).Render("[" + value + "]")
	}

	fill := toneColor(tone, t)

	style := ToneStyle(tone, level, t).Padding(0, theme.InlinePadding)
	switch level {
	case Bright:
		style = style.Foreground(t.Background).Background(fill)
	case Dim:
		style = style.Background(blend(fill, t.Background, dimPillTint))
	default:
		style = style.Background(blend(fill, t.Background, normalPillTint))
	}

	return style.Render(value)
}

func KeyHint(key, description string, t theme.Theme) string {
	return t.Heading().Render(key) + " " + t.Subtle().Render(description)
}

// Indicator renders semantic glyphs and labels as pills, or bare glyphs in Compact mode.
// GlyphOnly retains pill geometry; Compact removes the label, fill and padding.
func Indicator(glyph, label string, tone Tone, t theme.Theme, emphasis Emphasis, options ...Presentation) string {
	mode := presentation(options)
	if mode == Compact {
		return ToneStyle(tone, emphasis, t).Render(mode.Format(glyph, label))
	}

	return Pill(mode.Format(glyph, label), tone, t, emphasis)
}
