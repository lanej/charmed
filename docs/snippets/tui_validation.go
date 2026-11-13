package snippets

import (
	"fmt"

	"github.com/charmbracelet/x/ansi"

	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/tui"
)

func ExampleTuiValidation() (string, error) {
	screen, err := tui.NewScreen("first", "second")
	if err != nil {
		return "", err
	}

	root := &layout.Flex{Name: "steps", Axis: layout.Vertical, Children: []layout.Slot{
		{Node: &layout.Region{ID: "first", Min: layout.Size{Width: 20}}, Track: layout.Grow(1)},
		{Node: &layout.Region{ID: "second", Min: layout.Size{Width: 20}}, Track: layout.Grow(1)},
	}}

	err = screen.Arrange(root, layout.Size{Width: 80, Height: 12})
	if err == nil {
		return "", fmt.Errorf("expected strict stacking validation to block rendering")
	}

	return ansi.Wrap(err.Error(), 80, ""), nil
}
