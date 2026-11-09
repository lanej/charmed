package text

import (
	"bytes"
	"fmt"
	"image/color"
	"mime"
	"strings"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/charmbracelet/x/ansi"
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

	lexer := lexers.Get(language)
	if lexer == nil {
		return sourceText(source, t)
	}

	iterator, err := lexer.Tokenise(nil, source)
	if err != nil {
		return sourceText(source, t)
	}

	builder := chroma.NewStyleBuilder("charmed")

	for _, role := range []struct {
		token chroma.TokenType
		color color.Color
	}{
		{chroma.Text, t.Foreground},
		{chroma.Keyword, t.Syntax.Keyword},
		{chroma.LiteralString, t.Syntax.String},
		{chroma.LiteralNumber, t.Syntax.Number},
		{chroma.Name, t.Syntax.Name},
		{chroma.Comment, t.Syntax.Comment},
	} {
		if role.color != nil {
			r, g, b, _ := role.color.RGBA()
			builder.Add(role.token, fmt.Sprintf("#%02x%02x%02x", r>>8, g>>8, b>>8))
		}
	}

	style, err := builder.Build()
	if err != nil {
		return sourceText(source, t)
	}

	var output bytes.Buffer
	if err := formatters.TTY16m.Format(&output, style, iterator); err != nil {
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
	keyTheme, valueTheme := t, t
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
			parts[i] = sourceText(key, keyTheme) + t.Text().Render("=") + sourceText(value, valueTheme)
		} else {
			parts[i] = sourceText(part, t)
		}
	}

	return strings.Join(parts, t.Text().Render("&"))
}
