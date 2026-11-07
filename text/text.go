// Package text contains text atoms and a Glamour Markdown adapter.
package text

import (
	"strings"

	"charm.land/glamour/v2"
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

// Pill keeps its text label in every profile; color supplements its meaning.
// Semantic fills use contrasting theme surfaces, with brackets in ASCII mode.
func Pill(value string, tone Tone, t theme.Theme) string {
	if t.ASCII {
		return t.Text().Bold(true).Render("[" + value + "]")
	}
	fill := t.Foreground
	switch tone {
	case Success:
		fill = t.Success
	case Info:
		fill = t.Info
	case Warning:
		fill = t.Warning
	case Danger:
		fill = t.Danger
	}
	return t.Text().Foreground(t.Background).Background(fill).Bold(true).Padding(0, 1).Render(value)
}
func KeyHint(key, description string, t theme.Theme) string {
	return t.Heading().Render(key) + " " + t.Subtle().Render(description)
}

// Markdown renders rich content without probing the terminal or reading theme files.
func Markdown(content string, width int, t theme.Theme) (string, error) {
	style := "dark"
	if !t.IsDark() {
		style = "light"
	}
	if t.Mode == theme.Mono {
		style = "notty"
	}
	r, err := glamour.NewTermRenderer(glamour.WithStandardStyle(style), glamour.WithWordWrap(max(1, width)))
	if err != nil {
		return "", err
	}
	out, err := r.Render(content)
	return strings.TrimSpace(out), err
}
