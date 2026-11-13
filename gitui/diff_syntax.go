package gitui

import (
	"image/color"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
	"github.com/rivo/uniseg"
)

type syntaxSpan struct {
	end  int
	kind chroma.TokenType
}

type syntaxLine struct {
	text  string
	spans []syntaxSpan
}

func diffLexerText(source string) string {
	parts := strings.Split(source, "\t")
	for i, part := range parts {
		parts[i] = text.Inline(part)
	}

	return strings.Join(parts, "\t")
}

// Chroma iterators report some lexer failures by panicking during consumption.
func syntaxTokens(source string, tokenize func(string) (chroma.Iterator, error)) (tokens []chroma.Token, ok bool) {
	defer func() {
		if recover() != nil {
			tokens, ok = nil, false
		}
	}()

	iterator, err := tokenize(source)
	if err != nil {
		return nil, false
	}

	return iterator.Tokens(), true
}

func diffLanguage(path, override string) string {
	if override != "" {
		return override
	}

	if lexer := lexers.Match(path); lexer != nil {
		return lexer.Config().Name
	}

	return "text"
}

// Each side starts from the supplied hunk context; omitted source is not inferred.
func hunkSyntax(hunk Hunk, language string, old bool) []syntaxLine {
	lines := make([]syntaxLine, len(hunk.Lines))

	var sourceLines []string

	for i, line := range hunk.Lines {
		lines[i].text = text.Inline(line.Text)
		if (old && line.Kind == AddedLine) || (!old && line.Kind == DeletedLine) {
			continue
		}

		sourceLines = append(sourceLines, diffLexerText(line.Text))
	}

	lexer := lexers.Get(language)
	if language == "" || language == "text" || lexer == nil {
		return lines
	}

	source := strings.Join(sourceLines, "\n") + "\n"

	tokens, ok := syntaxTokens(source, func(value string) (chroma.Iterator, error) {
		return lexer.Tokenise(nil, value)
	})
	if !ok {
		return lines
	}

	if chroma.Stringify(tokens...) != source {
		return lines
	}

	index, offset := 0, 0

	for _, token := range tokens {
		for value := token.Value; value != ""; {
			for index < len(lines) && ((old && hunk.Lines[index].Kind == AddedLine) || (!old && hunk.Lines[index].Kind == DeletedLine)) {
				index++
			}

			if index == len(lines) {
				return lines
			}

			part, rest, newline := strings.Cut(value, "\n")

			offset += len(text.Inline(part))
			if part != "" {
				lines[index].spans = append(lines[index].spans, syntaxSpan{end: offset, kind: token.Type})
			}

			value = rest

			if newline {
				index++
				offset = 0
			}
		}
	}

	return lines
}

func syntaxColor(kind chroma.TokenType, t theme.Theme) color.Color {
	switch {
	case kind.InCategory(chroma.Keyword):
		return t.Syntax.Keyword
	case kind.InSubCategory(chroma.LiteralString):
		return t.Syntax.String
	case kind.InSubCategory(chroma.LiteralNumber):
		return t.Syntax.Number
	case kind.InCategory(chroma.Name):
		return t.Syntax.Name
	case kind.InCategory(chroma.Comment):
		return t.Syntax.Comment
	default:
		return t.Foreground
	}
}

func diffLineStyle(kind LineKind, emphasis bool, t theme.Theme) lipgloss.Style {
	style := t.Text()
	if t.Mode == theme.Mono {
		return lipgloss.NewStyle().Bold(emphasis).Underline(emphasis)
	}

	var background color.Color

	switch kind {
	case AddedLine:
		background = t.Diff.Added
		if emphasis {
			background = t.Diff.AddedEmphasis
		}
	case DeletedLine:
		background = t.Diff.Deleted
		if emphasis {
			background = t.Diff.DeletedEmphasis
		}
	case ContextLine:
	default:
	}

	if background != nil {
		style = style.Background(background)
	}

	return style
}

func renderSyntaxLine(line syntaxLine, kind LineKind, edits []editSpan, t theme.Theme) string {
	var output []string

	graphemes := uniseg.NewGraphemes(line.text)
	token, edit, start := 0, 0, 0
	previousKind, previousEmphasis := chroma.Text, false

	flush := func(end int) {
		style := diffLineStyle(kind, previousEmphasis, t)
		if foreground := syntaxColor(previousKind, t); t.Mode != theme.Mono && foreground != nil {
			style = style.Foreground(foreground)
		}

		output = append(output, style.Render(line.text[start:end]))
		start = end
	}

	for graphemes.Next() {
		from, _ := graphemes.Positions()
		for token < len(line.spans) && line.spans[token].end <= from {
			token++
		}

		for edit < len(edits) && edits[edit].end <= from {
			edit++
		}

		currentKind := chroma.Text
		if token < len(line.spans) {
			currentKind = line.spans[token].kind
		}

		emphasis := edit < len(edits) && edits[edit].start <= from
		if from > start && (currentKind != previousKind || emphasis != previousEmphasis) {
			flush(from)
		}

		previousKind, previousEmphasis = currentKind, emphasis
	}

	flush(len(line.text))

	return strings.Join(output, "")
}
