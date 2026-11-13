package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
	"github.com/lanej/charmed/tui"
)

func ExampleLayoutSplitPane() (string, error) {
	screen, err := tui.NewScreen("requests", "details")
	if err != nil {
		return "", err
	}

	root := &layout.SplitPane{
		First:  &layout.Region{ID: "requests", Min: layout.Size{Width: 20}},
		Second: &layout.Region{ID: "details", Min: layout.Size{Width: 30}}, Ratio: .4, Gap: 1,
	}
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
