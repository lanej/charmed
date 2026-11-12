package collection_test

import (
	"slices"
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func plainGrid(t *testing.T) (*collection.DataGrid[int], []collection.GridAxis) {
	t.Helper()

	axes := []collection.GridAxis{{ID: "a", Label: "Alpha"}, {ID: "b", Label: "Beta"}, {ID: "c", Label: "Gamma"}}

	grid, err := collection.NewDataGrid(axes, axes, []collection.GridEntry[int]{{RowID: "a", ColumnID: "a", Value: 0}, {RowID: "c", ColumnID: "c", Value: 42}}, layout.Size{Width: 70, Height: 6}, theme.New(theme.Mono))
	if err != nil {
		t.Fatal(err)
	}

	return grid, axes
}

func TestDataGridPlainValuesDistinguishZeroAndAbsence(t *testing.T) {
	grid, _ := plainGrid(t)
	grid.Focus()

	entry, present := grid.Selected()
	if !present || entry.Value != 0 || !strings.Contains(ansi.Strip(grid.View()), "[0             ]") {
		t.Fatal("zero value was treated as missing")
	}

	if !grid.Select("a", "b") {
		t.Fatal("absent coordinate could not be selected")
	}

	entry, present = grid.Selected()

	frame := ansi.Strip(grid.View())
	if present || entry.RowID != "a" || entry.ColumnID != "b" || !strings.Contains(frame, "[              ]") || strings.Contains(frame, "UNKNOWN") || strings.Contains(frame, "no mapping") {
		t.Fatalf("plain grid imposed mapping semantics: %+v / %v\n%s", entry, present, frame)
	}
}

func TestDataGridOwnsNavigationAndPinnedHeaders(t *testing.T) {
	grid, _ := plainGrid(t)
	grid.Resize(layout.Size{Width: 36, Height: 3})
	grid, _ = grid.Update(tea.KeyPressMsg{Code: tea.KeyRight})

	entry, _ := grid.Selected()
	if entry.ColumnID != "a" {
		t.Fatal("blurred grid moved")
	}

	grid.Focus()

	for _, code := range []rune{tea.KeyRight, tea.KeyRight, tea.KeyDown, tea.KeyDown, tea.KeyTab} {
		grid, _ = grid.Update(tea.KeyPressMsg{Code: code})
	}

	entry, present := grid.Selected()

	frame := ansi.Strip(grid.View())
	if !present || entry.Value != 42 || strings.Count(frame, "Gamma") != 2 || !strings.Contains(frame, "[42            ]") || !strings.Contains(frame, "Columns 3-3/3") {
		t.Fatalf("grid lost pinned axes or selected value:\n%s", frame)
	}

	grid, _ = grid.Update(tea.KeyPressMsg{Code: tea.KeyLeft})
	grid, _ = grid.Update(tea.KeyPressMsg{Code: tea.KeyUp})

	entry, _ = grid.Selected()
	if entry.RowID != "b" || entry.ColumnID != "b" {
		t.Fatal("reverse navigation failed")
	}
}

func TestDataGridReplacementIsAtomicAndRetainsCustomization(t *testing.T) {
	grid, axes := plainGrid(t)
	if !grid.Select("c", "c") {
		t.Fatal("selection failed")
	}

	grid.Corner, grid.CellWidth, grid.ShowHelp = "Custom", 12, true
	grid.RenderCell = func(cell collection.GridCell[int]) string { return "custom" }

	slices.Reverse(axes)

	entries := []collection.GridEntry[int]{{RowID: "c", ColumnID: "c", Value: 7}}
	if err := grid.SetData(axes, axes, entries); err != nil {
		t.Fatal(err)
	}

	axes[0].ID, entries[0].Value = "mutated", 99

	entry, present := grid.Selected()
	if !present || entry.RowID != "c" || entry.Value != 7 || grid.CellWidth != 12 || !grid.ShowHelp || !strings.Contains(grid.View(), "custom") {
		t.Fatal("update lost identity, copied data or customizations")
	}

	before := grid.View()
	if err := grid.SetData([]collection.GridAxis{{ID: "a"}, {ID: "a"}}, nil, nil); err == nil {
		t.Fatal("duplicate axes accepted")
	}

	if err := grid.SetData([]collection.GridAxis{{ID: "a"}}, []collection.GridAxis{{ID: "a"}}, []collection.GridEntry[int]{{RowID: "a", ColumnID: "missing"}}); err == nil {
		t.Fatal("unknown coordinate accepted")
	}

	if grid.View() != before {
		t.Fatal("failed replacement changed grid")
	}
}

func TestDataGridRendererReceivesTypedContextAndCannotEscapeGeometry(t *testing.T) {
	type value struct{ Count int }

	axes := []collection.GridAxis{{ID: "r"}}

	grid, err := collection.NewDataGrid(axes, axes, []collection.GridEntry[value]{{RowID: "r", ColumnID: "r", Value: value{Count: 17}}}, layout.Size{Width: 36, Height: 3}, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	grid.Focus()

	calls := 0
	grid.RenderCell = func(cell collection.GridCell[value]) string {
		calls++

		if !cell.Present || !cell.Selected || cell.Entry.Value.Count != 17 || cell.Entry.RowID != "r" || cell.Width != 16 || cell.Theme.Mode != theme.Dark {
			t.Fatalf("incomplete renderer context: %+v", cell)
		}

		return cell.Theme.Heading().Render(strings.Repeat("界", 50)) + "\nextra row"
	}

	frame := grid.View()
	if calls != 1 || lipgloss.Width(frame) != 36 || lipgloss.Height(frame) != 3 || strings.Contains(frame, "extra row") {
		t.Fatalf("renderer escaped allocation: %q", frame)
	}
}

func TestDataGridUsesUnicodeWithExplicitASCIIFallback(t *testing.T) {
	grid, axes := plainGrid(t)
	grid.Theme = theme.New(theme.Dark)

	axes[0].Label = "A deliberately long axis label"
	if err := grid.SetData(axes, axes, []collection.GridEntry[int]{{RowID: "a", ColumnID: "a", Value: 0}}); err != nil {
		t.Fatal(err)
	}

	grid.Focus()
	grid.ShowHelp = true

	frame := ansi.Strip(grid.View())
	if !strings.Contains(frame, "▏0             ▕") || !strings.Contains(frame, "…") || !strings.Contains(frame, "↑ ↓ ← →") || strings.Contains(frame, ".") {
		t.Fatalf("Unicode grid fell back to ASCII decoration:\n%s", frame)
	}

	grid.Theme.ASCII = true

	frame = ansi.Strip(grid.View())
	if !strings.Contains(frame, "[0             ]") || !strings.Contains(frame, ".") || !strings.Contains(frame, "^ v < >") || strings.ContainsAny(frame, "▏▕…↑↓←→") {
		t.Fatalf("explicit ASCII mode retained Unicode decoration:\n%s", frame)
	}
}
