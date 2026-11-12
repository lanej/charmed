package snippets

import (
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextChange() string {
	palette := theme.New(theme.Dark)

	return text.Added.View(text.Normal, palette) + "  " +
		text.Modified.View(text.Normal, palette) + "  " +
		text.Removed.View(text.Normal, palette)
}
