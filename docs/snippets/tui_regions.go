package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
	"github.com/lanej/charmed/tui"
)

func ExampleTuiRegions() (string, error) {
	screen, err := tui.NewScreen("header", "body")
	if err != nil {
		return "", err
	}

	root := &layout.Dock{Top: &layout.Region{ID: "header", Role: layout.ChromeRegion}, TopHeight: 1, Body: &layout.Region{ID: "body"}}
	if arrangeErr := screen.Arrange(root, layout.Size{Width: 80, Height: 12}); arrangeErr != nil {
		return "", arrangeErr
	}

	return screen.RegionMap(theme.New(theme.Dark)), nil
}
