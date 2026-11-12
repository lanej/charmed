package snippets

import (
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextSort() string {
	indicator := text.SortIndicator{
		Direction: text.SortDescending, Priority: 2, Emphasis: text.Normal,
	}

	return "Latency " + indicator.View(theme.New(theme.Dark))
}
