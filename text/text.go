// Package text contains text atoms and a Glamour Markdown adapter.
package text

import (
	"fmt"
	"image/color"
	"strings"

	"charm.land/glamour/v2"
	"charm.land/glamour/v2/styles"
	"github.com/charmbracelet/x/ansi"
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

// Markdown renders rich content without probing the terminal or reading theme files.
func Markdown(content string, width int, t theme.Theme) (string, error) {
	style := styles.DarkStyleConfig
	if !t.IsDark() {
		style = styles.LightStyleConfig
	}

	if t.Mode == theme.Mono {
		style = styles.NoTTYStyleConfig
	} else {
		style.Document.Color = markdownColor(t.Foreground)
		style.Heading.Color = markdownColor(t.Accent)
		style.H1.BackgroundColor = nil
		style.H1.Color = style.Heading.Color
		style.Code.Color = markdownColor(t.Accent)
		style.Code.BackgroundColor = markdownColor(t.Surface)
		style.Link.Color = markdownColor(t.Info)
		style.LinkText.Color = markdownColor(t.Accent)
	}
	// WHY: Glamour caches custom Chroma styles globally under one name, so the
	// first rendered palette would otherwise leak into subsequent themes.
	style.CodeBlock.Chroma = nil

	style.CodeBlock.Theme = "github-dark"
	if !t.IsDark() {
		style.CodeBlock.Theme = "github"
	}

	if t.Mode == theme.Mono {
		style.CodeBlock.Theme = ""
	}

	r, err := glamour.NewTermRenderer(glamour.WithStyles(style), glamour.WithWordWrap(max(1, width)), glamour.WithChromaFormatter("terminal16m"))
	if err != nil {
		return "", err
	}

	out, err := r.Render(content)
	if t.Mode == theme.Mono {
		out = ansi.Strip(out)
	}

	return strings.Trim(out, "\n"), err
}

func markdownColor(c color.Color) *string {
	if c == nil {
		return nil
	}

	r, g, b, _ := c.RGBA()
	value := fmt.Sprintf("#%02x%02x%02x", r>>8, g>>8, b>>8)

	return &value
}
