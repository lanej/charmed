package snippets

import (
	"fmt"

	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/theme"
)

func ExampleFeedbackPaginator() string {
	pages := feedback.NewPaginator(50, 10, theme.New(theme.Dark))
	pages.Page = 2 // Zero-based: show the third page of five.
	start, end := pages.GetSliceBounds(50)

	return pages.View() + fmt.Sprintf("\nItems %d–%d of 50", start+1, end)
	// Forward messages to pages.Update(msg); use start:end to select your visible items.
}
