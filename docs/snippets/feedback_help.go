package snippets

import (
	"charm.land/bubbles/v2/key"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/theme"
)

func ExampleFeedbackHelp() string {
	palette := theme.New(theme.Dark)
	next := key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "next field"))
	selectItem := key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "select"))
	help := feedback.NewHelp(40, palette)

	return help.ShortHelpView([]key.Binding{next, selectItem})
	// Use these same bindings in Update: key.Matches(msg, next).
}
