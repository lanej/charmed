package snippets

import (
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextCode() (string, error) {
	palette := theme.New(theme.Dark)
	inline := text.InlineCode("field.Focus()", palette)

	block, err := text.CodeBlock("field := input.NewText(options)", "go", 48, palette)
	if err != nil {
		return "", err
	}

	return inline + "\n\n" + block, nil
}
