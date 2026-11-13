package collection_test

import (
	"fmt"
	"strings"
	"testing"

	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func TestScrollPaneRetainsNativeCustomizationAndClamps(t *testing.T) {
	native := collection.NewViewport("", 12, 4)
	native.SoftWrap = true
	native.KeyMap.Down = key.NewBinding(key.WithKeys("x"))
	native.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	pane := collection.NewScrollPane(&native, layout.Size{Width: 12, Height: 4}, theme.New(theme.Mono))
	pane.SetContent(strings.Repeat("Long line that wraps into several rows.\n", 20))
	pane.Update(tea.KeyPressMsg{Code: 'x', Text: "x"})

	if native.YOffset() != 1 {
		t.Fatal("native custom key map not used")
	}

	native.SetYOffset(20)
	pane.SetContent("one")

	if pane.Model != &native || native.YOffset() != 0 || !native.SoftWrap || native.Style.GetForeground() != lipgloss.Color("2") {
		t.Fatal("native identity/customization or clamping lost")
	}

	if native.Width() != 11 || lipgloss.Width(pane.View()) != 12 {
		t.Fatal("scrollbar geometry changed when content shrank")
	}

	pane.Resize(layout.Size{Width: 1, Height: 0})

	if pane.View() != "" {
		t.Fatal("zero allocation rendered")
	}
}

func TestGridInspectorSelectionAndLayout(t *testing.T) {
	grid, _ := plainGrid(t)
	inspector := collection.NewGridInspector(grid, func(e collection.GridEntry[int], present bool, _ int) string {
		if !present {
			return "Absent coordinate"
		}

		return strings.Repeat("Details\n", 40)
	}, layout.Size{Width: 100, Height: 12}, theme.New(theme.Mono))
	inspector.FocusDetail(true)
	inspector.Details.Model.SetYOffset(5)
	inspector.Refresh()

	if inspector.Details.Model.YOffset() != 5 {
		t.Fatal("same coordinate refresh reset scroll")
	}

	inspector.Select("a", "b")

	if inspector.Details.Model.YOffset() != 0 || !strings.Contains(inspector.View(), "Absent coordinate") {
		t.Fatal("coordinate change did not replace/reset details")
	}

	for _, size := range []layout.Size{{Width: 100, Height: 12}, {Width: 50, Height: 12}, {Width: 30, Height: 3}} {
		inspector.Resize(size)

		if lipgloss.Width(inspector.View()) != size.Width || lipgloss.Height(inspector.View()) != size.Height {
			t.Fatalf("bad allocation %+v", size)
		}
	}

	inspector.FocusDetail(false)

	if !grid.Focused() {
		t.Fatal("grid focus not restored")
	}
}

func TestKeyedTableRetainsKeysAndRejectsAtomicInvalidReplacement(t *testing.T) {
	native := collection.NewTable([]table.Column{{Title: "Name", Width: 16}}, nil, 18, 4, theme.New(theme.Mono))
	rows := []collection.KeyedRow[int]{{Key: 0, Row: table.Row{"Zero"}}, {Key: 1, Row: table.Row{"One"}}}

	m, err := collection.NewKeyedTable(&native, rows)
	if err != nil {
		t.Fatal(err)
	}

	if k, ok := m.Selected(); !ok || k != 0 {
		t.Fatal("zero key treated as absent")
	}

	m.Select(1)

	rows[1].Row[0] = "Mutated caller"

	if strings.Contains(ansi.Strip(m.View()), "Mutated") {
		t.Fatal("rows were not copied")
	}

	if err := m.SetRows([]collection.KeyedRow[int]{rows[1], rows[0]}); err != nil {
		t.Fatal(err)
	}

	if k, _ := m.Selected(); k != 1 {
		t.Fatal("selection lost through reorder")
	}

	before := m.View()
	if err := m.SetData([]table.Column{{Title: "Bad", Width: 4}}, []collection.KeyedRow[int]{{Key: 2, Row: table.Row{"x", "y"}}}); err == nil || before != m.View() {
		t.Fatal("invalid shape mutated table")
	}

	if err := m.SetRows([]collection.KeyedRow[int]{rows[0], rows[0]}); err == nil {
		t.Fatal("duplicate keys accepted")
	}

	if err := m.SetRows([]collection.KeyedRow[int]{rows[0]}); err != nil {
		t.Fatal(err)
	}

	if k, _ := m.Selected(); k != 0 {
		t.Fatal("removed key did not select first")
	}
}

func TestKeyedTableRevealsSelectedRowsAcrossReplacement(t *testing.T) {
	native := collection.NewTable([]table.Column{{Title: "Name", Width: 16}}, nil, 18, 6, theme.New(theme.Mono))

	rows := make([]collection.KeyedRow[int], 20)
	for i := range rows {
		rows[i] = collection.KeyedRow[int]{Key: i, Row: table.Row{fmt.Sprintf("row-%02d", i)}}
	}

	m, err := collection.NewKeyedTable(&native, rows)
	if err != nil {
		t.Fatal(err)
	}

	m.Select(15)

	if !strings.Contains(ansi.Strip(m.View()), "row-15") {
		t.Fatal("selected row is off-screen")
	}

	rows[0], rows[15] = rows[15], rows[0]
	if err := m.SetRows(rows); err != nil {
		t.Fatal(err)
	}

	if k, _ := m.Selected(); k != 15 || !strings.Contains(ansi.Strip(m.View()), "row-15") {
		t.Fatal("restored key is off-screen")
	}
}

func TestScrollPaneUsesNativeVisibleContentHeight(t *testing.T) {
	native := collection.NewViewport("1\n2\n3\n4", 10, 4)
	native.Style = lipgloss.NewStyle().Border(lipgloss.NormalBorder())
	pane := collection.NewScrollPane(&native, layout.Size{Width: 10, Height: 4}, theme.New(theme.Dark))
	lines := strings.Split(ansi.Strip(pane.View()), "\n")
	visible := false

	for _, line := range lines {
		if strings.TrimSpace(ansi.Cut(line, 9, 10)) != "" {
			visible = true
		}
	}

	if !visible {
		t.Fatal("bordered viewport hid required scrollbar")
	}
}

func TestScrollPaneClampsHorizontalOffset(t *testing.T) {
	native := collection.NewViewport(strings.Repeat("x", 80), 10, 4)
	pane := collection.NewScrollPane(&native, layout.Size{Width: 10, Height: 4}, theme.New(theme.Mono))
	native.SetXOffset(40)
	pane.SetContent("short")

	if native.XOffset() != 0 {
		t.Fatal("short content retained invalid horizontal offset")
	}

	pane.SetContent(strings.Repeat("x", 80))
	native.SetXOffset(40)
	pane.Resize(layout.Size{Width: 100, Height: 4})

	if native.XOffset() != 0 {
		t.Fatal("resize retained invalid horizontal offset")
	}
}
