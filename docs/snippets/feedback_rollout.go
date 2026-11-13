package snippets

import (
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleFeedbackRollout() string {
	palette := theme.New(theme.Dark)

	return layout.Row(
		layout.Column(
			"Pill",
			"",
			feedback.Canary.View(text.Normal, palette),
			feedback.Live.View(text.Normal, palette),
			feedback.RolloutUnknown.View(text.Normal, palette),
		),
		"    ",
		layout.Column(
			"Compact",
			"",
			feedback.Canary.View(text.Normal, palette, text.Compact),
			feedback.Live.View(text.Normal, palette, text.Compact),
			feedback.RolloutUnknown.View(text.Normal, palette, text.Compact),
		),
	)
}
