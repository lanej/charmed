package snippets

import (
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextIdentifier() string {
	identifier := text.Identifier{Value: "trace-a17c9dd285e14766-92ef", Emphasis: text.Normal}
	// Use identifier.Value for keys and actions; Label is only a display abbreviation.
	return identifier.View(16, theme.New(theme.Dark))
}
