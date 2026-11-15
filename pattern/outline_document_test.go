package pattern_test

import (
	"slices"
	"strings"
	"testing"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/interaction"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

func outlineFixture(t *testing.T) *pattern.OutlineDocument {
	t.Helper()

	entries := []pattern.OutlineEntry{{ID: "plan", Title: "Proposal", Children: []pattern.OutlineEntry{
		{ID: "intent", Title: "Intent", Body: func(int) string { return strings.Repeat("purpose\n", 20) }},
		{ID: "evidence", Title: "Evidence", Body: func(int) string { return strings.Repeat("observation\n", 20) }},
	}}, {ID: "decision", Title: "Decision", Body: func(int) string { return strings.Repeat("accepted\n", 30) }}}

	m, err := pattern.NewOutlineDocument(entries, layout.Size{Width: 100, Height: 20}, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	return m
}

func TestOutlineDocumentLinksArbitraryContentBothDirections(t *testing.T) {
	m := outlineFixture(t)
	m.Select("intent")
	m.Update(tea.KeyPressMsg{Code: 'j'})

	if m.SelectedID() != "evidence" || m.Document.TopID() != "evidence" {
		t.Fatal("tree movement did not scroll linked section", m.SelectedID(), m.Document.TopID())
	}

	m.SetFocus(true)
	m.Update(tea.KeyPressMsg{Code: tea.KeyEnd})

	if m.SelectedID() != "decision" || !m.DocumentFocused() {
		t.Fatal("document scrolling did not track selection without focus change")
	}

	m.SetCollapsed("plan", true)

	if !m.Select("intent") || m.SelectedID() != "plan" || m.Document.TopID() != "plan" {
		t.Fatal("selecting hidden content should reveal its visible ancestor")
	}

	m.SetFocus(false)
	m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})

	if m.Collapsed("plan") {
		t.Fatal("keyboard did not expand group")
	}
}

func TestOutlineDocumentMouseCoordinatesRespectOriginAndScroll(t *testing.T) {
	m := outlineFixture(t)
	m.SetOrigin(7, 5)
	// At 100x20 the outline content starts at (2,1). Native root is row 0,
	// proposal is row 1, intent is row 2, and evidence is row 3.
	m.Update(tea.MouseClickMsg{X: 7 + 3, Y: 5 + 1 + 3, Button: tea.MouseLeft})

	if m.SelectedID() != "evidence" || m.Document.TopID() != "evidence" {
		t.Fatal("outline click did not reveal evidence", m.SelectedID())
	}

	// The document content begins at x=35. Its top row is now Evidence.
	m.Update(tea.MouseClickMsg{X: 7 + 36, Y: 5 + 1, Button: tea.MouseLeft})

	if !m.Collapsed("evidence") || !m.DocumentFocused() {
		t.Fatal("scrolled document heading did not fold")
	}

	selected := m.SelectedID()
	m.Update(tea.MouseClickMsg{X: 0, Y: 0, Button: tea.MouseLeft})

	if m.SelectedID() != selected {
		t.Fatal("outside click changed selection")
	}

	m.Update(tea.MouseWheelMsg{X: 7 + 36, Y: 5 + 2, Button: tea.MouseWheelDown})

	if m.SelectedID() != m.Document.TopID() {
		t.Fatal("wheel scrolling did not track active section")
	}
}

func TestOutlineDocumentPreservesBindingsAndBounds(t *testing.T) {
	m := outlineFixture(t)
	if err := m.SetKeyLayout(interaction.StandardKeys); err != nil {
		t.Fatal(err)
	}

	m.Outline.KeyMap.Down = key.NewBinding(key.WithKeys("n"))
	if err := m.SetEntries(m.Document.Sections()); err != nil {
		t.Fatal(err)
	}

	for _, size := range []layout.Size{{Width: 120, Height: 30}, {Width: 60, Height: 18}, {Width: 40, Height: 10}, {Width: 3, Height: 2}, {Width: 0, Height: 0}} {
		m.Resize(size)

		view := m.View()
		if size.Width == 0 || size.Height == 0 {
			if view != "" {
				t.Fatal("zero allocation must be empty")
			}

			continue
		}

		if lipgloss.Width(view) != size.Width || lipgloss.Height(view) != size.Height {
			t.Fatalf("allocation %v rendered %dx%d", size, lipgloss.Width(view), lipgloss.Height(view))
		}
	}

	if !slices.Contains(m.Outline.KeyMap.Down.Keys(), "n") {
		t.Fatal("replacement or resize reapplied presets")
	}

	if err := m.SetKeyLayout("invalid"); err == nil {
		t.Fatal("invalid navigation preset accepted")
	}
}

func TestOutlineDocumentCapsIndentationWithoutFlatteningHierarchy(t *testing.T) {
	entry := pattern.OutlineEntry{ID: "leaf", Title: "Deep leaf", Summary: "+12 -3", Body: func(int) string { return "content" }}
	for _, id := range []string{"five", "four", "three", "two", "one"} {
		entry = pattern.OutlineEntry{ID: id, Title: id, Children: []pattern.OutlineEntry{entry}}
	}

	m, err := pattern.NewOutlineDocument([]pattern.OutlineEntry{entry}, layout.Size{Width: 100, Height: 20}, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	m.Select("leaf")

	view := m.Outline.View()
	if !strings.Contains(view, "Deep leaf") {
		t.Fatal("deep indentation consumed the leaf label", view)
	}

	if err = m.SetIndentDepth(0); err != nil {
		t.Fatal(err)
	}

	m.SetCollapsed("three", true)

	if len(m.Document.VisibleIDs()) != 3 || !m.Select("leaf") || m.SelectedID() != "three" {
		t.Fatal("indent cap changed logical ancestry or subtree folding")
	}

	m.SetCollapsed("three", false)

	if err = m.SetIndentDepth(-1); err == nil || m.IndentDepth != 0 {
		t.Fatal("invalid indentation change was not rejected atomically")
	}

	m.Resize(layout.Size{Width: 40, Height: 18})
	m.Select("leaf")

	if !strings.Contains(m.Outline.View(), "Deep leaf") {
		t.Fatal("narrow navigation lost the deep leaf label", m.Outline.View())
	}
}
