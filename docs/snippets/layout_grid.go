package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
	"github.com/lanej/charmed/tui"
)

func ExampleLayoutGrid() (string, error) {
	screen, err := tui.NewScreen("main", "tools", "activity")
	if err != nil {
		return "", err
	}

	root := &layout.Grid{
		Columns: []layout.Track{layout.Grow(20), layout.Grow(12)}, Rows: []layout.Track{layout.Grow(3), layout.Grow(3)}, Gap: layout.Size{Width: 1, Height: 1},
		Cells: []layout.Cell{{Node: &layout.Region{ID: "main"}}, {Node: &layout.Region{ID: "tools"}, Column: 1}, {Node: &layout.Region{ID: "activity"}, Row: 1, ColumnSpan: 2}},
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
