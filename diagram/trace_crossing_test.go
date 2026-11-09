package diagram_test

import (
	"testing"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/lanej/charmed/diagram"
	"github.com/lanej/charmed/theme"
)

func TestLatestVerticalTraceKeepsCrossingOwnership(t *testing.T) {
	for _, tc := range []struct {
		name   string
		width  int
		steps  []int
		active string
	}{
		{"horizontal layout", 120, []int{0, 1}, "A"},
		{"vertical layout", 30, []int{1, 0}, "C"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			palette := theme.New(theme.Dark)
			d := diagram.StateDiagram{
				States:      []diagram.State{{ID: "A"}, {ID: "B"}, {ID: "C"}},
				Transitions: []diagram.Transition{{From: "A", To: "C", Label: "out"}, {From: "C", To: "A", Label: "back"}},
				Active:      tc.active,
				Trace:       diagram.TransitionTrace{Steps: tc.steps},
			}

			out, err := d.View(tc.width, palette)
			if err != nil {
				t.Fatal(err)
			}

			styled := uv.NewStyledString(out)
			screen := uv.NewScreenBuffer(styled.UnicodeWidth(), styled.Height())
			styled.Draw(screen, uv.Rect(0, 0, screen.Width(), screen.Height()))

			crossings := 0

			for y := 0; y < screen.Height(); y++ {
				for x := 0; x < screen.Width(); x++ {
					cell := screen.CellAt(x, y)
					if cell == nil || cell.Content != "┊" {
						continue
					}

					if !cell.Style.Equal(&uv.Style{Fg: palette.Accent, Attrs: uv.AttrBold}) {
						t.Fatalf("latest underpass at (%d,%d) lost accent/bold: %s", x, y, out)
					}

					for _, neighbor := range []int{y - 1, y + 1} {
						if neighbor < 0 || neighbor >= screen.Height() {
							continue
						}

						crossing := screen.CellAt(x, neighbor)
						if crossing == nil || crossing.Content != "─" {
							continue
						}

						if !crossing.Style.Equal(&uv.Style{Fg: palette.DimForeground}) {
							t.Fatalf("possible horizontal crossing at (%d,%d) inherited trace styling: %s", x, neighbor, out)
						}

						crossings++
					}
				}
			}

			if crossings == 0 {
				t.Fatalf("fixture did not expose a traced vertical underpass: %s", out)
			}
		})
	}
}
