package text_test

import (
	"testing"

	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestPreparedMarkdownAcrossWidths(t *testing.T) {
	source := "# Heading\n\nA paragraph that must wrap across several terminal cells.\n\n```go\nvar answer = 42\n```\n\n| Key | Value |\n| --- | --- |\n| a | b |\n"

	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Mono} {
		palette := theme.New(mode)

		document, err := text.PrepareMarkdown(source, palette)
		if err != nil {
			t.Fatal(err)
		}

		for _, width := range []int{120, 40, 80, 40, 120} {
			want, err := text.Markdown(source, width, palette)
			if err != nil {
				t.Fatal(err)
			}

			got, err := document.Render(width)
			if err != nil || got != want {
				t.Fatalf("width %d: %v\ngot %q\nwant %q", width, err, got, want)
			}
		}
	}
}
