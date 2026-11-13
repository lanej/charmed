package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
	"github.com/lanej/charmed/tui"
)

func ExampleLayoutDock() (string, error) {
	screen, err := tui.NewScreen("header", "main", "tools", "footer")
	if err != nil {
		return "", err
	}

	root := &layout.Dock{
		Top: &layout.Region{ID: "header", Role: layout.ChromeRegion}, TopHeight: 1,
		Left: &layout.Region{ID: "tools", Role: layout.SupportingRegion}, LeftWidth: 16,
		Body:   &layout.Region{ID: "main", Min: layout.Size{Width: 20, Height: 3}},
		Bottom: &layout.Region{ID: "footer", Role: layout.ChromeRegion}, BottomHeight: 1, Gap: 1,
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
