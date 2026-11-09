package text

import (
	"strings"

	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/lanej/charmed/theme"
)

// InlineCode styles literal text without interpreting Markdown delimiters.
func InlineCode(value string, t theme.Theme) string {
	if t.Mode == theme.Mono {
		return "`" + value + "`"
	}

	return t.Text().Foreground(t.Accent).Background(t.Surface).Padding(0, theme.InlinePadding).Render(value)
}

// CodeBlock renders literal source with the same highlighting as Markdown fences.
// Language is a lexer name such as "go" or "bash"; unknown names use plain text.
// Width controls the document allocation; long source lines can overflow it.
func CodeBlock(source, language string, width int, t theme.Theme) (string, error) {
	// WHY: A longer fence keeps embedded Markdown fences inside the source.
	fence := "~~~"
	for strings.Contains(source, fence) {
		fence += "~"
	}

	fields := strings.Fields(language)

	language = ""
	if len(fields) > 0 {
		language = strings.ReplaceAll(fields[0], "~", "")
	}

	if lexers.Get(language) == nil {
		language = "text"
	}

	if !strings.HasSuffix(source, "\n") {
		source += "\n"
	}

	return Markdown(fence+language+"\n"+source+fence+"\n", width, t)
}
