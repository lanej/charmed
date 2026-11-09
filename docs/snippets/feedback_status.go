package snippets

import (
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleFeedbackStatus() string {
	palette := theme.New(theme.Dark)

	return layout.Column(
		feedback.OK.View(text.Normal, palette),
		feedback.Warn.View(text.Bright, palette),
		feedback.Running.View(text.Normal, palette),
		feedback.Unknown.View(text.Dim, palette),
	)
}
