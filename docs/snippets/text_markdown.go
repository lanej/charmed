package snippets

import (
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextMarkdown() (string, error) {
	content := "# Terminal documentation\n\nUse **Charmed** to compose your terminal UI.\n\n- Shared theme\n- Bounded layout"

	rendered, err := text.Markdown(content, 48, theme.New(theme.Dark))
	if err != nil {
		return "", err
	}

	viewport := collection.NewViewport(rendered, 48, 12)

	return viewport.View(), nil
}
