package text

import (
	"regexp"
	"strings"

	"github.com/lanej/charmed/theme"
)

// WHY: Chroma's backtracking lexer can change token colors when its idle
// timeout clock restarts. JSON needs only a linear-time lexical scan; never
// decode values or normalize their spelling to choose a color.
var jsonLexemes = regexp.MustCompile(`"(?:\\[^\n]|[^"\\\n])*"|-?(?:0|[1-9][0-9]*)(?:\.[0-9]+)?(?:[eE][+-]?[0-9]+)?|\b(?:true|false|null)\b`)

func jsonSyntax(source string, t theme.Theme) string {
	parts := []string{}
	position := 0

	for _, span := range jsonLexemes.FindAllStringIndex(source, -1) {
		if span[0] > position {
			parts = append(parts, sourceText(source[position:span[0]], t))
		}

		color := t.Syntax.Number

		switch source[span[0]] {
		case '"':
			color = t.Syntax.String
			if strings.HasPrefix(strings.TrimLeft(source[span[1]:], " \t\r\n"), ":") {
				color = t.Syntax.Name
			}
		case 't', 'f', 'n':
			color = t.Syntax.Keyword
		}

		palette := t
		if color != nil {
			palette.Foreground = color
		}

		parts = append(parts, sourceText(source[span[0]:span[1]], palette))
		position = span[1]
	}

	if position < len(source) {
		parts = append(parts, sourceText(source[position:], t))
	}

	return strings.Join(parts, "")
}
