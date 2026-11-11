package snippets

import (
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleFeedbackOutcome() string {
	palette := theme.New(theme.Dark)
	// Timeout describes a reported result; it does not establish rollback or absence of effects.
	return feedback.TimedOut.View(text.Normal, palette) + "  " +
		feedback.Canceled.View(text.Normal, palette)
}
