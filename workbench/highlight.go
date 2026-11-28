package workbench

import (
	"go/scanner"
	"go/token"
	"html"
	"html/template"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/lanej/charmed/internal/lexical"
)

var (
	fencedCode = regexp.MustCompile("(?s)<pre><code(?: class=\"language-([^\"]+)\")?>(.*?)</code></pre>")
	inlineCode = regexp.MustCompile("<code>([^<]*)</code>")
)

// highlightMarkdown replaces escaped Goldmark code with server-rendered, semantic
// token spans. No browser JavaScript is necessary, and terminal ANSI spans remain
// separate from source highlighting.
func highlightMarkdown(markdownHTML string) string {
	rendered := fencedCode.ReplaceAllStringFunc(markdownHTML, func(match string) string {
		groups := fencedCode.FindStringSubmatch(match)
		lang := strings.ToLower(strings.TrimSpace(groups[1]))
		if lang == "" {
			lang = "text"
		}
		code := html.UnescapeString(groups[2])
		return "<pre tabindex=\"0\" class=\"code-block\"><code class=\"language-" +
			html.EscapeString(lang) + "\">" + string(highlightCode(code, lang)) + "</code></pre>"
	})
	rendered = inlineCode.ReplaceAllStringFunc(rendered, func(match string) string {
		groups := inlineCode.FindStringSubmatch(match)
		value := html.UnescapeString(groups[1])
		// Inline paths/commands should not be guessed to be Go expressions.
		if strings.ContainsAny(value, " /\\") {
			return "<code class=\"inline-code\">" + html.EscapeString(value) + "</code>"
		}
		return "<code class=\"inline-code\">" + string(highlightGo(value)) + "</code>"
	})

	return headingPermalinks(strings.ReplaceAll(rendered, "<table>", "<table tabindex=\"0\">"))
}

func highlightCode(code, language string) template.HTML {
	switch strings.ToLower(language) {
	case "go", "golang":
		return highlightGo(code)
	case "sh", "bash", "shell", "console", "zsh":
		return highlightShell(code)
	default:
		return template.HTML(html.EscapeString(code))
	}
}

// Go's own scanner recognizes keywords, literals and comments without a third-
// party runtime. Offsets preserve whitespace and punctuation exactly.
func highlightGo(source string) template.HTML {
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("sample.go", fset.Base(), len(source))
	s.Init(file, []byte(source), nil, scanner.ScanComments)
	var out strings.Builder
	last := 0
	previous := token.ILLEGAL
	for {
		pos, tok, literal := s.Scan()
		if tok == token.EOF {
			break
		}
		if tok == token.SEMICOLON && literal == "\n" {
			// Synthetic semicolon; the source newline remains in the gap.
			continue
		}
		start := file.Offset(pos)
		if start < last || start > len(source) {
			continue
		}
		out.WriteString(html.EscapeString(source[last:start]))
		width := len(literal)
		if width == 0 {
			width = len(tok.String())
		}
		end := start + width
		if end > len(source) || source[start:end] != tokenText(tok, literal) {
			_, size := utf8.DecodeRuneInString(source[start:])
			end = start + size
		}
		if end > len(source) || end <= start {
			continue
		}
		class := ""
		switch {
		case tok == token.COMMENT:
			class = "syn-comment"
		case tok == token.STRING || tok == token.CHAR:
			class = "syn-string"
		case tok == token.INT || tok == token.FLOAT || tok == token.IMAG:
			class = "syn-number"
		case tok.IsKeyword():
			class = "syn-keyword"
		case tok == token.IDENT:
			class = identifierClass(literal, previous, source[end:])
		case tok >= token.ADD && tok <= token.ELLIPSIS || tok == token.TILDE:
			class = "syn-operator"
		}
		writeSyntax(&out, class, source[start:end])

		if tok != token.COMMENT {
			previous = tok
		}

		last = end
	}
	out.WriteString(html.EscapeString(source[last:]))
	return template.HTML(out.String())
}

func tokenText(tok token.Token, literal string) string {
	if literal != "" {
		return literal
	}
	return tok.String()
}

func identifierClass(ident string, previous token.Token, rest string) string {
	if previous == token.TYPE || lexical.GoType(ident) {
		return "syn-type"
	}

	switch ident {
	case "true":
		return "syn-true"
	case "false":
		return "syn-false"
	case "nil":
		return "syn-null"
	case "iota":
		return "syn-constant"
	}

	if strings.HasPrefix(strings.TrimSpace(rest), "(") {
		return "syn-function"
	}

	if previous == token.PERIOD {
		return "syn-member"
	}

	return ""
}

// Shell commands receive lightweight lexical styling. This is intentionally not
// a shell evaluator: unknown syntax remains escaped, readable source.
func highlightShell(source string) template.HTML {
	var out strings.Builder
	firstWord := true
	for i := 0; i < len(source); {
		c := source[i]
		if c == '\n' {
			out.WriteByte('\n')
			i++
			firstWord = true
			continue
		}
		if c == ' ' || c == '\t' || c == '\r' {
			out.WriteByte(c)
			i++
			continue
		}
		if c == '#' {
			end := i
			for end < len(source) && source[end] != '\n' {
				end++
			}
			writeSyntax(&out, "syn-comment", source[i:end])
			i = end
			continue
		}
		if c == '\'' || c == '"' {
			quote := c
			end := i + 1
			for end < len(source) {
				if source[end] == '\\' && quote == '"' && end+1 < len(source) {
					end += 2
					continue
				}
				if source[end] == quote {
					end++
					break
				}
				if source[end] == '\n' {
					break
				}
				end++
			}
			writeSyntax(&out, "syn-string", source[i:end])
			i = end
			firstWord = false
			continue
		}
		if c == '$' {
			end := i + 1
			for end < len(source) && (isWordChar(rune(source[end])) || source[end] == '{' || source[end] == '}') {
				end++
			}

			writeSyntax(&out, "syn-reference", source[i:end])
			i = end
			continue
		}
		end := i
		for end < len(source) && source[end] != '\n' && source[end] != ' ' && source[end] != '\t' && source[end] != '\r' && source[end] != '\'' && source[end] != '"' {
			end++
		}
		if end == i {
			out.WriteString(html.EscapeString(source[i : i+1]))
			i++
			continue
		}
		word := source[i:end]
		class := ""
		if strings.HasPrefix(word, "-") {
			class = "syn-flag"
		} else if firstWord {
			class = "syn-command"
		}
		writeSyntax(&out, class, word)
		firstWord = false
		i = end
	}
	return template.HTML(out.String())
}

func isWordChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
}

func writeSyntax(out *strings.Builder, class, value string) {
	if class == "" {
		out.WriteString(html.EscapeString(value))
		return
	}
	out.WriteString("<span class=\"")
	out.WriteString(class)
	out.WriteString("\">")
	out.WriteString(html.EscapeString(value))
	out.WriteString("</span>")
}
