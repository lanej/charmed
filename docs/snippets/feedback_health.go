package snippets

import (
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleFeedbackHealth() string {
	palette := theme.New(theme.Dark)

	return layout.Row(
		layout.Column(
			"Pill",
			"",
			feedback.Healthy.View(text.Normal, palette),
			feedback.Unhealthy.View(text.Normal, palette),
			feedback.HealthUnknown.View(text.Normal, palette),
		),
		"    ",
		layout.Column(
			"Compact",
			"",
			feedback.Healthy.View(text.Normal, palette, text.Compact),
			feedback.Unhealthy.View(text.Normal, palette, text.Compact),
			feedback.HealthUnknown.View(text.Normal, palette, text.Compact),
		),
	)
}
