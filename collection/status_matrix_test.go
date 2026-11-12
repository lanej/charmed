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
