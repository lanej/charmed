package snippets

import (
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/theme"
)

func ExampleFeedbackScrollbar() string {
	scrollbar := feedback.Scrollbar{Total: 100, Visible: 20, Offset: 35, Height: 8}
	return scrollbar.View(theme.New(theme.Dark))
}
