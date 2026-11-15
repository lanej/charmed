package collection_test

import (
	"slices"
	"strings"
	"testing"

	"charm.land/bubbles/v2/key"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func stackFixture(t *testing.T) *collection.SectionStack {
	t.Helper()

	m, err := collection.NewSectionStack([]collection.DocumentSection{
		{ID: "guide", Title: "Guide", Summary: "2 sections", Children: []collection.DocumentSection{
			{ID: "intro", Title: "Introduction", Body: func(int) string { return "one\ntwo\nthree\nfour\nfive" }},
			{ID: "proof", Title: "Evidence", Body: func(int) string { return strings.Repeat("evidence\n", 20) }},
		}},
		{ID: "end", Title: "Conclusion", Body: func(int) string { return strings.Repeat("conclusion\n", 10) }},
	}, layout.Size{Width: 30, Height: 4}, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	return m
}

func TestSectionStackHierarchyAndIndependentFolds(t *testing.T) {
	m := stackFixture(t)
	if !slices.Equal(m.VisibleIDs(), []string{"guide", "intro", "proof", "end"}) {
		t.Fatal(m.VisibleIDs())
	}

	m.SetCollapsed("intro", true)
	m.ScrollTo("proof")

	if m.TopID() != "proof" {
		t.Fatal(m.TopID())
	}

	m.SetCollapsed("guide", true)

	if !slices.Equal(m.VisibleIDs(), []string{"guide", "end"}) || m.TopID() != "guide" {
		t.Fatalf("collapsed ancestor did not replace its children: %v %s", m.VisibleIDs(), m.TopID())
	}

	if !m.ScrollTo("intro") || m.TopID() != "guide" {
		t.Fatal("hidden child did not resolve to its summary")
	}

	m.SetCollapsed("guide", false)

	if !m.Collapsed("intro") || m.Collapsed("proof") {
		t.Fatal("ancestor fold overwrote independent child folds")
	}

	if m.SetCollapsed("absent", true) || m.ScrollTo("absent") {
		t.Fatal("unknown IDs accepted")
	}
}

func TestSectionStackUpdatesRetainAnchorAndNativeCustomization(t *testing.T) {
	m := stackFixture(t)
	m.Viewport.KeyMap.Down = key.NewBinding(key.WithKeys("n"))
	m.ScrollTo("proof")
	m.Viewport.ScrollDown(2)
	oldOffset := m.Viewport.YOffset()
	entries := m.Sections()

	entries[0].Children[0].Body = func(int) string { return strings.Repeat("extra\n", 8) }
	if err := m.SetSections(entries); err != nil {
		t.Fatal(err)
	}

	if m.TopID() != "proof" || m.Viewport.YOffset() != oldOffset+4 {
		t.Fatal("replacement lost the relative section anchor", m.TopID(), m.Viewport.YOffset(), oldOffset)
	}

	m.Resize(layout.Size{Width: 20, Height: 3})

	if m.TopID() != "proof" || !slices.Contains(m.Viewport.KeyMap.Down.Keys(), "n") {
		t.Fatal("resize lost anchor or native bindings")
	}

	if err := m.SetSections([]collection.DocumentSection{{ID: "x"}, {ID: "x"}}); err == nil || m.TopID() != "proof" {
		t.Fatal("invalid replacement was not atomic")
	}

	entries[0].Children = entries[0].Children[:1]
	if err := m.SetSections(entries); err != nil {
		t.Fatal(err)
	}

	if m.TopID() != "guide" {
		t.Fatal("removed active section did not anchor to its ancestor", m.TopID())
	}
}

func TestSectionStackCachesRenderersAndMeasuresWrappedCells(t *testing.T) {
	calls := 0
	entries := []collection.DocumentSection{{ID: "wide", Title: "界é", Body: func(int) string {
		calls++
		return "界é界é界é界é"
	}}}

	m, err := collection.NewSectionStack(entries, layout.Size{Width: 6, Height: 8}, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	m.View()
	m.View()
	m.SetCollapsed("wide", true)
	m.SetCollapsed("wide", false)
	m.Resize(layout.Size{Width: 6, Height: 4})

	if calls != 1 {
		t.Fatal("unnecessary renderer calls", calls)
	}

	if m.HeaderAt(0) != "wide" || m.IDAt(2) != "wide" {
		t.Fatal("wrapped body offsets differ from heading offsets")
	}

	for _, line := range strings.Split(ansi.Strip(m.View()), "\n") {
		if lipgloss.Width(line) != 6 {
			t.Fatal("incorrect cell width", line)
		}
	}

	m.Resize(layout.Size{Width: 8, Height: 4})
	m.Refresh()

	if calls != 3 {
		t.Fatal("width change and explicit refresh must invalidate bodies", calls)
	}
}
