package snippets

import (
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextOmission() string {
	palette := theme.New(theme.Dark)
	hidden := text.Omission{Count: 24, KnownCount: true, Emphasis: text.Normal}
	truncated := text.Omission{Kind: text.SourceTruncated, Emphasis: text.Normal}

	return hidden.View(palette) + "  " + truncated.View(palette)
}
