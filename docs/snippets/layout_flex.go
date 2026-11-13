package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
	"github.com/lanej/charmed/tui"
)

func ExampleLayoutFlex() (string, error) {
	screen, err := tui.NewScreen("main", "tools")
	if err != nil {
		return "", err
	}

	root := &layout.Flex{Gap: 1, Children: []layout.Slot{
		{Node: &layout.Region{ID: "main", Min: layout.Size{Width: 20}}, Track: layout.Grow(20)},
		{Node: &layout.Region{ID: "tools", Min: layout.Size{Width: 12}}, Track: layout.Fixed(16)},
	}}
	if arrangeErr := screen.Arrange(root, layout.Size{Width: 80, Height: 16}); arrangeErr != nil {
		return "", arrangeErr
	}

	views := map[string]string{}

	for _, region := range screen.Plan().Regions() {
		size := layout.Size{Width: region.Bounds.Dx(), Height: region.Bounds.Dy()}
		views[region.ID] = layout.Panel(region.ID, "Native view", size, theme.New(theme.Dark), false)
	}

	frame, err := screen.Render(views)

	return frame.Content, err
}
