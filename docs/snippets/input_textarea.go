package snippets

import (
	"github.com/lanej/charmed/input"
	"github.com/lanej/charmed/theme"
)

func ExampleInputTextarea() string {
	editor := input.NewArea(48, 8, theme.New(theme.Dark))
	editor.Placeholder = "Write a note"
	editor.SetValue("Deployment notes\nReady for production.")
	// Return editor.Focus() from Init and forward messages through editor.Update(msg).
	return editor.View()
}
