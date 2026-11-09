package snippets

import (
	"charm.land/log/v2"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleFeedbackLogLevel() string {
	palette := theme.New(theme.Dark)

	return layout.Column(
		feedback.LogLevel(log.InfoLevel, text.Normal, palette)+" Request complete",
		feedback.LogLevel(log.WarnLevel, text.Normal, palette)+" Retry scheduled",
		feedback.LogLevel(log.ErrorLevel, text.Bright, palette)+" Upstream timeout",
	)
}
