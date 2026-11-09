package snippets

import (
	"time"

	"github.com/lanej/charmed/feedback"
)

func ExampleFeedbackTime() string {
	timer := feedback.NewTimer(time.Minute)
	stopwatch := feedback.NewStopwatch()
	// Return their Init commands with tea.Batch and forward tick messages to Update.
	return "Countdown  " + timer.View() + "\nStopwatch  " + stopwatch.View()
}
