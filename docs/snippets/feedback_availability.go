package snippets

import (
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleFeedbackAvailability() string {
	palette := theme.New(theme.Dark)

	return layout.Row(
		layout.Column(
			"Pill",
			"",
			feedback.Up.View(text.Normal, palette),
			feedback.Down.View(text.Normal, palette),
			feedback.AvailabilityUnknown.View(text.Normal, palette),
		),
		"    ",
		layout.Column(
			"Compact",
			"",
			feedback.Up.View(text.Normal, palette, text.Compact),
			feedback.Down.View(text.Normal, palette, text.Compact),
			feedback.AvailabilityUnknown.View(text.Normal, palette, text.Compact),
		),
	)
}
