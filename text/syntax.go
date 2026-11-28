package text

import (
	"bytes"
	"mime"
	"strings"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/internal/lexical"
	"github.com/lanej/charmed/theme"
)

// ContentLanguage uses declared media types, never payload sniffing.
func ContentLanguage(contentType string) string {
	media, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return ""
	}

	media = strings.ToLower(media)
	switch {
	case media == "application/json" || strings.HasSuffix(media, "+json"):
		return "json"
	case media == "application/xml" || media == "text/xml" || strings.HasSuffix(media, "+xml"):
		return "xml"
	case media == "text/html":
		return "html"
	case media == "text/css":
		return "css"
	case media == "text/javascript" || media == "application/javascript":
		return "javascript"
	case media == "application/x-www-form-urlencoded" || strings.HasPrefix(media, "text/"):
		if media == "application/x-www-form-urlencoded" {
			return "form"
		}

		return "text"
	default:
		return ""
	}
}

// Syntax highlights escaped source without changing whitespace or executing content.
func Syntax(source, language string, t theme.Theme) string {
	source = Escape(source)
	if language == "form" {
		return formSyntax(source, t)
	}

	if t.Mode == theme.Mono || language == "text" || language == "" {
		return sourceText(source, t)
	}

	if language == "json" {
		return jsonSyntax(source, t)
	}

	lexer := lexers.Get(language)
	if lexer == nil {
		return sourceText(source, t)
	}

	iterator, err := lexer.Tokenise(nil, source)
	if err != nil {
		return sourceText(source, t)
	}

	entries := chroma.StyleEntries{ //nolint:exhaustive // Chroma inherits unlisted lexical variants.
		chroma.Text: syntaxColor(t.Foreground), chroma.Keyword: syntaxColor(t.Syntax.Keyword),
		chroma.LiteralString: syntaxColor(t.Syntax.String), chroma.LiteralNumber: syntaxColor(t.Syntax.Number),
		chroma.Name: syntaxColor(t.Syntax.Name), chroma.Comment: syntaxColor(t.Syntax.Comment),
	}
	lexical.Entries(entries, t)

	style, err := chroma.NewStyle("charmed", entries)
	if err != nil {
		return sourceText(source, t)
	}

	var output bytes.Buffer
	if err := formatters.TTY16m.Format(&output, style, lexical.Iterator(iterator, style, language)); err != nil {
		return sourceText(source, t)
	}

	if ansi.Strip(output.String()) != source {
		return sourceText(source, t)
	}

	return output.String()
}

func sourceText(source string, t theme.Theme) string {
	lines := strings.Split(source, "\n")
	for i, line := range lines {
		lines[i] = t.Text().TabWidth(-1).Render(line)
	}

	return strings.Join(lines, "\n")
}

func formSyntax(source string, t theme.Theme) string {
	keyTheme, valueTheme, operatorTheme := t, t, t
	if t.Lexical.Operator != nil {
		operatorTheme.Foreground = t.Lexical.Operator
	}

	if t.Syntax.Name != nil {
		keyTheme.Foreground = t.Syntax.Name
	}

	if t.Syntax.String != nil {
		valueTheme.Foreground = t.Syntax.String
	}

	parts := strings.Split(source, "&")
	for i, part := range parts {
		key, value, found := strings.Cut(part, "=")
		if found && t.Mode != theme.Mono {
			parts[i] = sourceText(key, keyTheme) + operatorTheme.Text().Render("=") + sourceText(value, valueTheme)
		} else {
			parts[i] = sourceText(part, t)
		}
	}

	return strings.Join(parts, operatorTheme.Text().Render("&"))
}
