package snippets

import (
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/theme"
)

func ExampleFeedbackProgress() string {
	progress := feedback.NewProgress(40, theme.New(theme.Dark))
	completed, total := 68.0, 100.0

	return progress.ViewAs(completed/total) + "  68%"
}
