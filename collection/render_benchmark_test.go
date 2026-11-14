package collection_test

import (
	"strconv"
	"testing"

	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func renderAxes(count int) []collection.GridAxis {
	axes := make([]collection.GridAxis, count)
	for i := range axes {
		axes[i] = collection.GridAxis{ID: strconv.Itoa(i)}
	}

	return axes
}

var gridFrame string

func BenchmarkGridWindow(b *testing.B) {
	for _, rows := range []int{100, 10000} {
		b.Run(strconv.Itoa(rows), func(b *testing.B) {
			grid, err := collection.NewDataGrid[int](renderAxes(rows), renderAxes(100), nil, layout.Size{Width: 120, Height: 40}, theme.New(theme.Dark))
			if err != nil {
				b.Fatal(err)
			}

			b.ReportAllocs()

			for b.Loop() {
				gridFrame = grid.View()
			}
		})
	}
}

func TestLargeGridRendersOnlyItsVisibleWindow(t *testing.T) {
	rows, columns := renderAxes(10000), renderAxes(100)

	grid, err := collection.NewDataGrid[int](rows, columns, nil, layout.Size{Width: 120, Height: 40}, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	rendered := 0
	grid.RenderCell = func(collection.GridCell[int]) string {
		rendered++
		return "cell"
	}

	for _, size := range []layout.Size{{Width: 120, Height: 40}, {Width: 40, Height: 12}} {
		grid.Resize(size)

		for _, position := range [][2]string{{rows[0].ID, columns[0].ID}, {rows[len(rows)-1].ID, columns[len(columns)-1].ID}} {
			if !grid.Select(position[0], position[1]) {
				t.Fatal("large grid lost a keyed coordinate")
			}

			rendered = 0

			grid.View()

			if rendered == 0 || rendered > size.Width*size.Height {
				t.Fatalf("rendered %d cells for a %v allocation", rendered, size)
			}
		}
	}
}
