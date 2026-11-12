package snippets

import (
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextMarkdown() (string, error) {
	palette := theme.New(theme.Dark)
	content := "# Help\n\nCall `field.Focus()` to **focus** the input.\n\n```go\ncmd := field.Focus()\n```"
	styles := text.MarkdownStyles(palette)
	styles.Document.Margin = new(uint(0))

	rendered, err := text.Markdown(content, 48, palette, text.MarkdownOptions{Styles: &styles})
	if err != nil {
		return "", err
	}

	viewport := collection.NewViewport(rendered, 48, 12)

	return viewport.View(), nil
}
