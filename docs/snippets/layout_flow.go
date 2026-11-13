package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
	"github.com/lanej/charmed/tui"
)

func ExampleLayoutFlow() (string, error) {
	screen, err := tui.NewScreen("build", "test", "publish")
	if err != nil {
		return "", err
	}

	root := &layout.Flow{Gap: layout.Size{Width: 1, Height: 1}, Children: []layout.Box{
		{Node: &layout.Region{ID: "build"}, Size: layout.Size{Width: 30, Height: 3}},
		{Node: &layout.Region{ID: "test"}, Size: layout.Size{Width: 30, Height: 3}},
		{Node: &layout.Region{ID: "publish"}, Size: layout.Size{Width: 30, Height: 3}},
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
