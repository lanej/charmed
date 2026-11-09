package snippets

import (
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextKeyHint() string {
	palette := theme.New(theme.Dark)
	return text.KeyHint("Enter", "select", palette) + "  " + text.KeyHint("Esc", "dismiss", palette)
}
