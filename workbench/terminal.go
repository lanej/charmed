package workbench

import (
	"fmt"
	"html"
	"html/template"
	"image/color"
	"strings"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/theme"
)

// terminalHTML renders saved cells, preserving RGB and escaping all terminal text.
// NOTE: Links and terminal control effects are never emitted as HTML capabilities.
func terminalHTML(frame string) template.HTML {
	lines := uv.NewStyledString(strings.TrimRight(frame, " \n")).Lines(ansi.GraphemeWidth)

	var result []string

	for i, line := range lines {
		if i > 0 {
			result = append(result, "\n")
		}

		var (
			run      []string
			previous uv.Style
		)

		flush := func() {
			if len(run) == 0 {
				return
			}

			content := html.EscapeString(strings.Join(run, ""))
			if style := terminalStyle(previous); style != "" {
				result = append(result, `<span style="`+style+`">`+content+`</span>`)
			} else {
				result = append(result, content)
			}

			run = nil
		}

		for _, cell := range line {
			if cell.Width == 0 {
				continue
			}

			if !previous.Equal(&cell.Style) {
				flush()

				previous = cell.Style
			}

			run = append(run, cell.Content)
		}

		flush()
	}

	return template.HTML(strings.Join(result, "")) // #nosec G203 -- Text is HTML-escaped; styles contain only numeric colors and fixed literals.
}

func terminalStyle(style uv.Style) string {
	var css []string

	fg, bg := style.Fg, style.Bg
	if style.Attrs&uv.AttrReverse != 0 {
		defaults := theme.New(theme.Dark)
		if fg == nil {
			fg = defaults.Foreground
		}

		if bg == nil {
			bg = defaults.Background
		}

		fg, bg = bg, fg
	}

	if fg != nil {
		css = append(css, "color:"+terminalColor(fg)+";")
	}

	if bg != nil {
		css = append(css, "background-color:"+terminalColor(bg)+";")
	}

	if style.Attrs&uv.AttrBold != 0 {
		css = append(css, "font-weight:bold;")
	}

	if style.Attrs&uv.AttrFaint != 0 {
		css = append(css, "opacity:.7;")
	}

	if style.Attrs&uv.AttrItalic != 0 {
		css = append(css, "font-style:italic;")
	}

	var decorations []string
	if style.Underline != uv.UnderlineNone {
		decorations = append(decorations, "underline")
	}

	if style.Attrs&uv.AttrStrikethrough != 0 {
		decorations = append(decorations, "line-through")
	}

	if len(decorations) > 0 {
		css = append(css, "text-decoration:"+strings.Join(decorations, " ")+";")
	}

	if style.Attrs&uv.AttrConceal != 0 {
		css = append(css, "visibility:hidden;")
	}

	return strings.Join(css, "")
}

func terminalColor(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("#%02x%02x%02x", r>>8, g>>8, b>>8)
}
