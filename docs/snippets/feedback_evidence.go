package snippets

import (
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleFeedbackEvidence() string {
	palette := theme.New(theme.Dark)

	// Parse the state word once at the boundary; anything unrecognized is unknown, never healthy.
	state, _ := feedback.ParseEvidence("stale")

	return state.View(text.Normal, palette)
}
