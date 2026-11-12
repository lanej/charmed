package snippets

import (
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleFeedbackStatusCell() string {
	palette := theme.New(theme.Dark)
	healthy := feedback.StatusCell{Status: feedback.OK, Label: "v2.8.1"}
	draining := feedback.StatusCell{Status: feedback.Warn, Label: "draining"}
	failed := feedback.StatusCell{Status: feedback.Error, Label: "unhealthy"}
	unknown := feedback.StatusCell{Status: feedback.Unknown, Label: "no probe"}

	return layout.Column(
		healthy.View(18, palette, false),
		draining.View(18, palette, true),
		failed.View(18, palette, false),
		unknown.View(18, palette, false),
	)
}
