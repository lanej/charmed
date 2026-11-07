// Package text contains text atoms and a Glamour Markdown adapter.
package text

import (
	"strings"

	"charm.land/glamour/v2"
	"github.com/lanej/charmed/theme"
)

func Label(value string, t theme.Theme) string { return t.Text().Render(value) }
func Badge(value string, t theme.Theme) string { return t.Heading().Render("[" + value + "]") }
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
