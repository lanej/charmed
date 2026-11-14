package text

import (
	"bytes"
	"fmt"
	"image/color"
	"reflect"
	"strings"

	glamour "charm.land/glamour/v2/ansi"
	"charm.land/glamour/v2/styles"
	"charm.land/lipgloss/v2"
	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/theme"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	goldtext "github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

// MarkdownOptions overrides native styles for one render. Nil fields use the
// theme defaults. Syntax is a local Chroma style, never a global registration.
// CodeBlock.Theme and CodeBlock.Chroma in Styles are ignored; use Syntax instead.
// Mono always strips ANSI, including caller-supplied colors.
type MarkdownOptions struct {
	Styles *glamour.StyleConfig
	Syntax *chroma.Style
}

// MarkdownStyles returns an independent native Glamour configuration. Callers
// can change layout, decoration and colors before passing it in MarkdownOptions.
func MarkdownStyles(t theme.Theme) glamour.StyleConfig {
	style := styles.DarkStyleConfig
	if !t.IsDark() {
		style = styles.LightStyleConfig
	}

	if t.Mode == theme.Mono {
		style = styles.NoTTYStyleConfig
	}

	style.Document.Color = markdownColor(t.Foreground)
	style.Heading.Color = markdownColor(t.Accent)
	style.H1.BackgroundColor = nil
	style.H1.Color = markdownColor(t.Accent)
	style.H6.Color = markdownColor(t.Accent)
	style.Code.Color = markdownColor(t.Accent)
	style.Code.BackgroundColor = markdownColor(t.Surface)
	style.CodeBlock.Color = markdownColor(t.Foreground)
	style.CodeBlock.BackgroundColor = markdownColor(t.Surface)
	style.CodeBlock.Theme = ""
	style.CodeBlock.Chroma = nil
	style.BlockQuote.Color = markdownColor(t.Muted)
	style.HorizontalRule.Color = markdownColor(t.Border)
	style.Link.Color = markdownColor(t.Info)
	style.LinkText.Color = markdownColor(t.Accent)
	style.Image.Color = markdownColor(t.Info)
	style.ImageText.Color = markdownColor(t.Muted)
	style.Table.Color = markdownColor(t.Foreground)
	style.DefinitionTerm.Color = markdownColor(t.Accent)
	style.DefinitionDescription.Color = markdownColor(t.Foreground)
	cloneStylePointers(reflect.ValueOf(&style).Elem())

	return style
}

// WHY: Native style configs contain pointers even when copied as values. Give
// callers owned scalar pointers so editing one palette cannot mutate presets.
func cloneStylePointers(value reflect.Value) {
	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return
		}

		owned := reflect.New(value.Type().Elem())
		owned.Elem().Set(value.Elem())
		value.Set(owned)
		cloneStylePointers(owned.Elem())

		return
	}

	if value.Kind() == reflect.Struct {
		for i := 0; i < value.NumField(); i++ {
			cloneStylePointers(value.Field(i))
		}
	}
}

// MarkdownSyntax derives lexical roles from the theme: accent keywords and
// functions, information-colored types, categorical strings/numbers, muted
// comments and normal foreground text. Missing colors inherit normal text.
func MarkdownSyntax(t theme.Theme) (*chroma.Style, error) {
	// Chroma inherits unlisted token variants from these lexical parent roles.
	return chroma.NewStyle("charmed", chroma.StyleEntries{ //nolint:exhaustive // Intentional parent-role palette; Chroma supplies variant inheritance.
		chroma.Text:          syntaxColor(t.Foreground),
		chroma.Keyword:       "bold " + syntaxColor(t.Accent),
		chroma.KeywordType:   syntaxColor(t.Info),
		chroma.NameFunction:  syntaxColor(t.Accent),
		chroma.LiteralString: syntaxColor(t.SeriesColor(0)),
		chroma.LiteralNumber: syntaxColor(t.SeriesColor(1)),
		chroma.Comment:       "italic " + syntaxColor(t.Muted),
	})
}

func syntaxColor(c color.Color) string {
	if value := markdownColor(c); value != nil {
		return *value
	}

	return ""
}

// Markdown renders rich content without terminal queries or theme files. An
// optional configuration replaces native styles; width still controls prose.
func Markdown(content string, width int, t theme.Theme, options ...MarkdownOptions) (string, error) {
	document, err := PrepareMarkdown(content, t, options...)
	if err != nil {
		return "", err
	}

	return document.Render(width)
}

// PrepareMarkdown parses and highlights once, independently of viewport size.
func PrepareMarkdown(content string, t theme.Theme, options ...MarkdownOptions) (*MarkdownDocument, error) {
	style := MarkdownStyles(t)

	syntax, err := MarkdownSyntax(t)
	if err != nil {
		return nil, err
	}

	for _, option := range options {
		if option.Styles != nil {
			style = *option.Styles
			cloneStylePointers(reflect.ValueOf(&style).Elem())
		}

		if option.Syntax != nil {
			syntax = option.Syntax
		}
	}
	// WHY: Highlight locally; Glamour caches custom Chroma styles globally under
	// one name, which lets the first theme contaminate subsequent renders.
	style.CodeBlock.Theme = ""
	style.CodeBlock.Chroma = nil
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM, extension.DefinitionList),
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
	)

	source := []byte(content)
	document := md.Parser().Parse(goldtext.NewReader(source))

	background := ""
	if style.CodeBlock.BackgroundColor != nil {
		background = (ansi.Style{}).BackgroundColor(lipgloss.Color(*style.CodeBlock.BackgroundColor)).String()
	}

	source, err = highlightMarkdownSource(document, source, syntax, t.Mode == theme.Mono, background)
	if err != nil {
		return nil, err
	}

	return &MarkdownDocument{source: source, document: document, style: style, mono: t.Mode == theme.Mono}, nil
}

// MarkdownDocument owns prepared source and the latest width result. Render calls
// must be serialized. Preparing a new document replaces content or presentation.
type MarkdownDocument struct {
	source   []byte
	document ast.Node
	style    glamour.StyleConfig
	mono     bool
	width    int
	rendered string
	valid    bool
}

func (d *MarkdownDocument) Render(width int) (string, error) {
	width = max(1, width)
	if d.valid && d.width == width {
		return d.rendered, nil
	}

	r := renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(
		glamour.NewRenderer(glamour.Options{Styles: d.style, WordWrap: width}), 1000)))

	var out bytes.Buffer
	if err := r.Render(&out, d.source, d.document); err != nil {
		return "", err
	}

	rendered := out.String()
	if d.mono {
		rendered = ansi.Strip(rendered)
	}

	d.width, d.rendered, d.valid = width, strings.Trim(rendered, "\n"), true

	return d.rendered, nil
}

func highlightMarkdownSource(document ast.Node, source []byte, syntax *chroma.Style, mono bool, background string) ([]byte, error) {
	err := ast.Walk(document, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if span, ok := node.(*ast.CodeSpan); ok {
			// WHY: Glamour decodes HTML entities in code spans; Markdown treats them as literal source.
			for child := span.FirstChild(); child != nil; child = child.NextSibling() {
				part, ok := child.(*ast.Text)
				if !ok {
					continue
				}

				value := strings.ReplaceAll(string(part.Segment.Value(source)), "&", "&amp;")
				start := len(source)
				source = append(source, value...)
				part.Segment = goldtext.NewSegment(start, len(source))
			}

			return ast.WalkSkipChildren, nil
		}

		language := "text"

		switch code := node.(type) {
		case *ast.FencedCodeBlock:
			language = string(code.Language(source))
		case *ast.CodeBlock:
		default:
			return ast.WalkContinue, nil
		}

		var code []byte

		for i := 0; i < node.Lines().Len(); i++ {
			line := node.Lines().At(i)
			code = append(code, line.Value(source)...)
		}

		value := string(code)

		if !mono {
			lexer := lexers.Get(language)
			if lexer == nil {
				lexer = lexers.Fallback
			}

			tokens, err := chroma.Coalesce(lexer).Tokenise(nil, value)
			if err != nil {
				return ast.WalkStop, fmt.Errorf("highlight %s: %w", language, err)
			}

			var highlighted bytes.Buffer
			if err = formatters.TTY16m.Format(&highlighted, syntax, tokens); err != nil {
				return ast.WalkStop, err
			}

			value = strings.ReplaceAll(highlighted.String(), "\x1b[0m", "\x1b[0m"+background)
		}
		// WHY: Glamour's plain code renderer unescapes Markdown backslashes. The AST
		// is already parsed; doubling them here preserves literal source exactly.
		value = strings.ReplaceAll(value, "\\", "\\\\")
		start := len(source)
		source = append(source, value...)

		node.Lines().Clear()
		node.Lines().Append(goldtext.NewSegment(start, len(source)))

		return ast.WalkSkipChildren, nil
	})

	return source, err
}

func markdownColor(c color.Color) *string {
	if c == nil {
		return nil
	}

	r, g, b, _ := c.RGBA()
	value := fmt.Sprintf("#%02x%02x%02x", r>>8, g>>8, b>>8)

	return &value
}
