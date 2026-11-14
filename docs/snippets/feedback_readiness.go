package snippets

import (
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleFeedbackReadiness() string {
	palette := theme.New(theme.Dark)

	return layout.Row(
		layout.Column(
			"Pill",
			"",
			feedback.Ready.View(text.Normal, palette),
			feedback.NotReady.View(text.Normal, palette),
			feedback.ReadinessUnknown.View(text.Normal, palette),
		),
		"    ",
		layout.Column(
			"Compact",
			"",
			feedback.Ready.View(text.Normal, palette, text.Compact),
			feedback.NotReady.View(text.Normal, palette, text.Compact),
			feedback.ReadinessUnknown.View(text.Normal, palette, text.Compact),
		),
	)
}
