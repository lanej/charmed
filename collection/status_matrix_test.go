package collection_test

import (
	"slices"
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func matrixFixture(t *testing.T) (*collection.StatusMatrix, []collection.MatrixAxis, []collection.MatrixEntry) {
	t.Helper()

	axes := []collection.MatrixAxis{{ID: "a", Label: "Alpha"}, {ID: "b", Label: "Beta"}, {ID: "c", Label: "Gamma"}}
	entries := []collection.MatrixEntry{{RowID: "a", ColumnID: "a", Cell: feedback.StatusCell{Status: feedback.OK, Label: "v2.8"}}, {RowID: "c", ColumnID: "c", Cell: feedback.StatusCell{Status: feedback.Unknown, Label: "no probe"}}}

	m, err := collection.NewStatusMatrix(axes, axes, entries, layout.Size{Width: 70, Height: 7}, theme.New(theme.Mono))
	if err != nil {
		t.Fatal(err)
	}

	return m, axes, entries
}

func TestStatusMatrixMappingSelectionAndAbsence(t *testing.T) {
	m, _, _ := matrixFixture(t)

	selected, present := m.Selected()
	if !present || selected.Cell.Label != "v2.8" {
		t.Fatalf("initial mapping: %+v / %v", selected, present)
	}

	if !m.Select("a", "b") {
		t.Fatal("cannot select an absent pair")
	}

	selected, present = m.Selected()
	if present || selected.RowID != "a" || selected.ColumnID != "b" {
		t.Fatalf("absent pair lost identity: %+v / %v", selected, present)
	}

	if m.Select("missing", "c") {
		t.Fatal("selected unknown axis")
	}

	after, _ := m.Selected()
	if after != selected {
		t.Fatal("invalid selection moved cursor")
	}

	if !m.Select("c", "c") {
		t.Fatal("cannot select unknown-health relationship")
	}

	selected, present = m.Selected()
	if !present || selected.Cell.Status != feedback.Unknown {
		t.Fatal("unknown health conflated with no mapping")
	}
}

func TestStatusMatrixHelpIsOptionalAndAbsentPairsUseOneGlyph(t *testing.T) {
	m, _, _ := matrixFixture(t)
	m.Resize(layout.Size{Width: 70, Height: 6})
	frame := ansi.Strip(m.View())

	if strings.Contains(frame, "no mapping") || strings.Contains(frame, "h j k l") {
		t.Fatal("default matrix contains repetitive labels or instructions")
	}

	if !m.Select("a", "b") {
		t.Fatal("cannot select absent pair")
	}

	m.Focus()
	frame = ansi.Strip(m.View())

	if !strings.Contains(frame, "[-             ]") {
		t.Fatalf("absent selection must retain one glyph and focus brackets: %q", frame)
	}

	m.ShowHelp = true
	frame = ansi.Strip(m.View())

	if !strings.Contains(frame, "^ v < > / h j k l") || !strings.Contains(frame, "Rows 1-3/3") || lipgloss.Height(frame) != 6 {
		t.Fatalf("optional help displaced rows or escaped allocation: %q", frame)
	}

	m.Resize(layout.Size{Width: 70, Height: 4})
	frame = ansi.Strip(m.View())

	if !strings.Contains(frame, "Rows 1-1/3") || !strings.Contains(frame, "h j k l") {
		t.Fatal("small matrix failed to reserve optional help row")
	}
}

func TestStatusMatrixMissingGlyphAndFocusFollowTerminalPolicy(t *testing.T) {
	m, _, _ := matrixFixture(t)

	m.Theme = theme.New(theme.Dark)
	if !m.Select("a", "b") {
		t.Fatal("cannot select absent mapping")
	}

	m.Focus()

	frame := ansi.Strip(m.View())
	if !strings.Contains(frame, "▏—             ▕") || strings.Contains(frame, "no mapping") {
		t.Fatalf("default missing mapping did not use Unicode:\n%s", frame)
	}

	m.Theme.ASCII = true

	frame = ansi.Strip(m.View())
	if !strings.Contains(frame, "[-             ]") || strings.ContainsAny(frame, "▏▕—") {
		t.Fatalf("missing mapping did not respect ASCII fallback:\n%s", frame)
	}
}

func TestStatusMatrixKeyboardScrollsBothAxesAndPinsHeaders(t *testing.T) {
	m, _, _ := matrixFixture(t)
	m.Resize(layout.Size{Width: 36, Height: 3})
	m, _ = m.Update(tea.KeyPressMsg{Code: 'l', Text: "l"})

	selected, _ := m.Selected()
	if selected.ColumnID != "a" {
		t.Fatal("blurred matrix consumed navigation")
	}

	m.Focus()

	for _, key := range []rune{'l', 'l', 'l', 'j', 'j', 'j'} {
		m, _ = m.Update(tea.KeyPressMsg{Code: key, Text: string(key)})
	}

	selected, present := m.Selected()

	frame := ansi.Strip(m.View())

	if !present || selected.RowID != "c" || selected.ColumnID != "c" || !strings.Contains(frame, "? no probe") || strings.Count(frame, "Gamma") != 2 || !strings.Contains(frame, "Rows 3-3/3") {
		t.Fatalf("scroll lost selected cell, headers, or range: %+v\n%s", selected, frame)
	}

	m.Resize(layout.Size{Width: 70, Height: 7})

	if frame = ansi.Strip(m.View()); !strings.Contains(frame, "Alpha") || !strings.Contains(frame, "v2.8") {
		t.Fatal("resize did not recover earlier rows and columns")
	}

	m.Blur()

	if strings.Contains(ansi.Strip(m.View()), "[? no probe") {
		t.Fatal("blur retained focus brackets")
	}
}

func TestStatusMatrixUpdatesCopyDataAndKeepIdentity(t *testing.T) {
	m, axes, entries := matrixFixture(t)
	if !m.Select("c", "c") {
		t.Fatal("selection failed")
	}

	slices.Reverse(axes)

	if err := m.SetData(axes, axes, entries); err != nil {
		t.Fatal(err)
	}

	axes[0].ID, axes[0].Label, entries[1].Cell.Label = "mutated", "mutated", "mutated"

	selected, present := m.Selected()
	if !present || selected.RowID != "c" || selected.ColumnID != "c" || selected.Cell.Label != "no probe" {
		t.Fatalf("data reorder or caller mutation corrupted selection: %+v", selected)
	}

	if err := m.SetData([]collection.MatrixAxis{{ID: "a"}}, []collection.MatrixAxis{{ID: "a"}}, nil); err != nil {
		t.Fatal(err)
	}

	selected, _ = m.Selected()
	if selected.RowID != "a" || selected.ColumnID != "a" {
		t.Fatal("removed IDs did not reset to first axes")
	}
}

func TestStatusMatrixRejectsInvalidUpdatesAtomically(t *testing.T) {
	m, axes, entries := matrixFixture(t)

	before := m.View()
	for _, invalid := range []struct {
		rows, columns []collection.MatrixAxis
		entries       []collection.MatrixEntry
	}{
		{[]collection.MatrixAxis{{ID: ""}}, axes, nil},
		{axes, []collection.MatrixAxis{{ID: "a"}, {ID: "a"}}, nil},
		{axes, axes, []collection.MatrixEntry{{RowID: "a", ColumnID: "missing"}}},
		{axes, axes, []collection.MatrixEntry{{RowID: "missing", ColumnID: "a"}}},
		{axes, axes, []collection.MatrixEntry{entries[0], entries[0]}},
	} {
		if err := m.SetData(invalid.rows, invalid.columns, invalid.entries); err == nil {
			t.Fatal("invalid matrix update accepted")
		}

		if m.View() != before {
			t.Fatal("failed update changed existing view")
		}
	}
}

func TestStatusMatrixAlwaysHonorsAllocationAndEmptyState(t *testing.T) {
	m, _, _ := matrixFixture(t)

	for width := 0; width <= 75; width++ {
		for height := 0; height <= 8; height++ {
			m.Resize(layout.Size{Width: width, Height: height})

			frame := m.View()
			if width == 0 || height == 0 {
				if frame != "" {
					t.Fatal("zero allocation rendered content")
				}

				continue
			}

			if lipgloss.Width(frame) != width || lipgloss.Height(frame) != height {
				t.Fatalf("matrix escaped %dx%d", width, height)
			}
		}
	}

	if err := m.SetData(nil, nil, nil); err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(m.View(), "No mappings") {
		t.Fatal("empty matrix lacks empty state")
	}
}

func legendMatrix(t *testing.T, mode theme.Mode, ascii bool, statuses ...feedback.Status) *collection.StatusMatrix {
	t.Helper()

	axes := []collection.MatrixAxis{{ID: "a", Label: "Alpha"}}
	columns := make([]collection.MatrixAxis, len(statuses))
	entries := make([]collection.MatrixEntry, len(statuses))

	for i, status := range statuses {
		columns[i] = collection.MatrixAxis{ID: string(rune('a' + i)), Label: string(status)}
		entries[i] = collection.MatrixEntry{RowID: "a", ColumnID: columns[i].ID, Cell: feedback.StatusCell{Status: status}}
	}

	palette := theme.New(mode)
	palette.ASCII = ascii

	m, err := collection.NewStatusMatrix(axes, columns, entries, layout.Size{Width: 80, Height: 8}, palette)
	if err != nil {
		t.Fatal(err)
	}

	return m
}

func TestStatusMatrixLegendListsPresentStatusesInEveryTheme(t *testing.T) {
	for name, setup := range map[string]struct {
		mode  theme.Mode
		ascii bool
	}{"dark": {theme.Dark, false}, "light": {theme.Light, false}, "mono": {theme.Mono, false}, "ascii": {theme.Dark, true}, "mono ascii": {theme.Mono, true}} {
		m := legendMatrix(t, setup.mode, setup.ascii, feedback.Error, feedback.OK, feedback.Unknown, feedback.OK)
		legend := ansi.Strip(m.Legend(80))

		want := ansi.Strip(feedback.OK.View(text.Normal, m.Theme)) + "  " + ansi.Strip(feedback.Error.View(text.Normal, m.Theme)) + "  " + ansi.Strip(feedback.Unknown.View(text.Normal, m.Theme))
		if legend != want {
			t.Fatalf("%s legend %q, want %q", name, legend, want)
		}

		if strings.Contains(legend, "WARN") || strings.Count(legend, "OK") != 1 {
			t.Fatalf("%s legend lists absent or repeated statuses: %q", name, legend)
		}
	}
}

func TestStatusMatrixLegendKeepsUnrecognizedStatusesApartFromHealthy(t *testing.T) {
	for _, ascii := range []bool{false, true} {
		for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Mono} {
			m := legendMatrix(t, mode, ascii, feedback.Status("HEALTHY"), feedback.Status(""), feedback.Status("ok"))
			legend := ansi.Strip(m.Legend(80))
			unknown := ansi.Strip(feedback.Unknown.View(text.Normal, m.Theme))

			if strings.TrimSpace(legend) != strings.TrimSpace(unknown) {
				t.Fatalf("unrecognized statuses must collapse to UNKNOWN, got %q", legend)
			}

			healthy := ansi.Strip(feedback.OK.View(text.Normal, m.Theme))
			if strings.Contains(legend, strings.TrimSpace(healthy)) || strings.Contains(legend, "HEALTHY") {
				t.Fatalf("unknown shares the healthy glyph or word: %q", legend)
			}
		}
	}
}

func TestStatusMatrixLegendWrapsWholeEntriesWithinWidth(t *testing.T) {
	m := legendMatrix(t, theme.Dark, false, feedback.OK, feedback.Warn, feedback.Error, feedback.Unknown, feedback.Disconnected)

	for width := 1; width <= 90; width++ {
		legend := m.Legend(width)
		for _, line := range strings.Split(legend, "\n") {
			if got := lipgloss.Width(line); got > width {
				t.Fatalf("width %d line is %d cells: %q", width, got, ansi.Strip(line))
			}
		}
	}

	if lines := strings.Split(ansi.Strip(m.Legend(30)), "\n"); len(lines) < 2 || strings.Contains(lines[0], "…") {
		t.Fatalf("narrow legend did not wrap whole entries: %q", lines)
	}

	if m.Legend(0) != "" {
		t.Fatal("zero width legend is not empty")
	}
}

func TestStatusMatrixLegendIsOptionalAndTakesRowsFromTheAllocation(t *testing.T) {
	m := legendMatrix(t, theme.Dark, false, feedback.OK, feedback.Error)
	legend := ansi.Strip(feedback.OK.View(text.Normal, m.Theme)) + "  " + ansi.Strip(feedback.Error.View(text.Normal, m.Theme))

	if plain := ansi.Strip(m.View()); strings.Contains(plain, legend) || plain != ansi.Strip(m.DataGrid.View()) {
		t.Fatalf("legend appeared by default:\n%s", plain)
	}

	m.ShowLegend = true
	frame := ansi.Strip(m.View())

	if lipgloss.Height(frame) != 8 || !strings.Contains(frame, "Rows 1-1/1") || !strings.Contains(frame, legend) {
		t.Fatalf("legend missing or escaped allocation:\n%s", frame)
	}

	empty, err := collection.NewStatusMatrix(nil, nil, nil, layout.Size{Width: 40, Height: 4}, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	empty.ShowLegend = true

	if empty.Legend(40) != "" || strings.Contains(ansi.Strip(empty.View()), "OK") {
		t.Fatal("empty matrix has a legend")
	}
}
