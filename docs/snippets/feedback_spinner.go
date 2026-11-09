package snippets

import (
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/theme"
)

func ExampleFeedbackSpinner() string {
	spinner := feedback.NewSpinner(theme.New(theme.Dark))
	// Return spinner.Tick from Init; retain the model and command from spinner.Update(msg).
	return spinner.View() + "  Processing requests"
}
