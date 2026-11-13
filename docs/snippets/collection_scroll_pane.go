package snippets

import (
	"strings"

	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleCollectionScrollPane() (string, error) {
	native := collection.NewViewport("", 36, 6)
	native.SoftWrap = true
	pane := collection.NewScrollPane(&native, layout.Size{Width: 36, Height: 6}, theme.New(theme.Dark))
	pane.SetContent(strings.Repeat("Request observations remain scrollable.\n", 12))
	native.ScrollDown(2)
	// Route Update to the pane while active; Resize retains native customization.
	return pane.View(), nil
}
