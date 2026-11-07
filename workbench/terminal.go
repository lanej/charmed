package workbench

import (
	"fmt"
	"html/template"
	"regexp"
	"strings"

	terminal "github.com/buildkite/terminal-to-html/v3"
	"github.com/microcosm-cc/bluemonday"
)

var terminalPolicy = func() *bluemonday.Policy {
	p := bluemonday.StrictPolicy()
	p.AllowElements("span")
	p.AllowAttrs("class").Matching(regexp.MustCompile(`^term-[a-z0-9]+( term-[a-z0-9]+)*$`)).OnElements("span")
	return p
}()

// terminalHTML preserves ANSI color while excluding links, images and executable markup.
func terminalHTML(frame string) template.HTML {
	rendered := terminal.Render([]byte(strings.TrimRight(frame, " \n")))
	return template.HTML(terminalPolicy.Sanitize(rendered))
}

func terminalCSS() string {
	var css strings.Builder
	base := []string{"#14201c", "#f28b82", "#86cf9a", "#f2c879", "#78b8ed", "#c4a7e7", "#67d4c0", "#eef1ee", "#607b6c", "#f28b82", "#86cf9a", "#f2c879", "#78b8ed", "#c4a7e7", "#67d4c0", "#ffffff"}
	for i := range 256 {
		color := ""
		switch {
		case i < 16:
			color = base[i]
		case i < 232:
			v := i - 16
			levels := []int{0, 95, 135, 175, 215, 255}
			color = fmt.Sprintf("#%02x%02x%02x", levels[v/36], levels[v/6%6], levels[v%6])
		default:
			level := 8 + (i-232)*10
			color = fmt.Sprintf("#%02x%02x%02x", level, level, level)
		}
		fmt.Fprintf(&css, ".terminal .term-fgx%d{color:%s}.terminal .term-bgx%d{background-color:%s}\n", i, color, i, color)
		if i < 8 {
			fmt.Fprintf(&css, ".terminal .term-fg%d{color:%s}.terminal .term-bg%d{background-color:%s}\n", i+30, color, i+40, color)
		}
		if i >= 8 && i < 16 {
			fmt.Fprintf(&css, ".terminal .term-fgi%d{color:%s}.terminal .term-bgi%d{background-color:%s}\n", i+82, color, i+92, color)
		}
	}
	css.WriteString(".terminal .term-fg1{font-weight:bold}.terminal .term-fg2{opacity:.7}.terminal .term-fg3{font-style:italic}.terminal .term-fg4{text-decoration:underline}.terminal .term-fg9{text-decoration:line-through}\n")
	return css.String()
}
