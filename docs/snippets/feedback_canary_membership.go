package snippets

import (
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleFeedbackCanaryMembership() string {
	palette := theme.New(theme.Dark)

	return layout.Row(
		layout.Column(
			"Pill",
			"",
			feedback.CanaryOn.View(text.Normal, palette),
			feedback.CanaryOff.View(text.Normal, palette),
			feedback.CanaryMembershipUnknown.View(text.Normal, palette),
		),
		"    ",
		layout.Column(
			"Compact",
			"",
			feedback.CanaryOn.View(text.Normal, palette, text.Compact),
			feedback.CanaryOff.View(text.Normal, palette, text.Compact),
			feedback.CanaryMembershipUnknown.View(text.Normal, palette, text.Compact),
		),
	)
}
