package snippets

import (
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleLayoutFooter() string {
	palette := theme.New(theme.Dark)
	footer := layout.Footer{
		Help: text.KeyHint("?", "help", palette) + "  " +
			text.KeyHint("q", "quit", palette),
		Status: feedback.Connected.View(text.Normal, palette),
	}

	return footer.View(64, palette)
}
