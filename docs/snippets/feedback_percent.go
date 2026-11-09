package snippets

import (
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleFeedbackPercent() string {
	palette := theme.New(theme.Dark)
	measured := feedback.PercentBar(68, true, 30, palette)
	unavailable := feedback.PercentBar(0, false, 30, palette)

	return layout.Column(measured, unavailable)
}
