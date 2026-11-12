package navigation_test

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/interaction"
	"github.com/lanej/charmed/navigation"
	"github.com/lanej/charmed/theme"
)

func TestTabsNavigation(t *testing.T) {
	items := []navigation.Tab{{ID: "one", Label: "One"}, {ID: "two", Label: "Two"}, {ID: "three", Label: "Three"}}
	m := navigation.NewTabs(items, theme.New(theme.Mono))
	m.SetWidth(20)

	key := func(code rune) tea.KeyPressMsg { return tea.KeyPressMsg{Code: code} }

	cmd := m.Update(key(tea.KeyRight))
	if cmd != nil {
		t.Fatal("blurred tabs handled input")
	}

	m.Focus()

	for _, step := range []struct {
		code rune
		want string
	}{{tea.KeyLeft, "three"}, {tea.KeyRight, "one"}, {tea.KeyEnd, "three"}, {tea.KeyHome, "one"}} {
		cmd = m.Update(key(step.code))

		active, ok := m.Selected()
		if !ok || active.ID != step.want || cmd == nil {
			t.Fatalf("navigation: active=%+v cmd nil=%v", active, cmd == nil)
		}

		selected, ok := cmd().(interaction.Selected)
		if !ok || selected.Component != "tabs" || selected.ID != step.want {
			t.Fatalf("selection event: %+v", selected)
		}
	}

	cmd = m.Update(key(tea.KeyTab))
	if cmd != nil {
		t.Fatal("tabs captured parent focus key")
	}

	m.Blur()

	if m.Focused() {
		t.Fatal("blur did not release focus")
	}
}

func TestTabsRetainSelectionAndOwnItems(t *testing.T) {
	items := []navigation.Tab{{ID: "one", Label: "One"}, {ID: "two", Label: "Two"}}
	m := navigation.NewTabs(items, theme.New(theme.Mono))
	items[0].Label = "Mutated"

	m.SetWidth(30)

	if strings.Contains(m.View(), "Mutated") {
		t.Fatal("constructor retained caller slice")
	}

	if !m.Select("two") || m.Select("missing") {
		t.Fatal("selection lookup")
	}

	m.Focus()
	m.SetItems([]navigation.Tab{{ID: "two", Label: "Renamed"}, {ID: "one", Label: "One"}})
	m.SetWidth(8)

	selected, ok := m.Selected()
	if !ok || selected.ID != "two" || !m.Focused() {
		t.Fatal("update lost selection or focus")
	}

	m.SetItems([]navigation.Tab{{ID: "other", Label: "Other"}})

	selected, ok = m.Selected()
	if !ok || selected.ID != "other" {
		t.Fatal("missing selection did not fall back")
	}

	m.SetItems(nil)

	if _, ok := m.Selected(); ok {
		t.Fatal("empty tabs have selection")
	}
}

func TestTabsOverflowGeometry(t *testing.T) {
	m := navigation.NewTabs([]navigation.Tab{{ID: "a", Label: "Alpha"}, {ID: "b", Label: "界界"}, {ID: "c", Label: "Last"}}, theme.New(theme.Mono))
	if !m.Select("b") {
		t.Fatal("select middle")
	}

	for width := 0; width < 40; width++ {
		m.SetWidth(width)

		view := ansi.Strip(m.View())
		if ansi.StringWidth(view) != width || strings.Contains(view, "\n") {
			t.Fatalf("width %d: %q", width, view)
		}

		if width == 8 && view != "< 界界 >" {
			t.Fatalf("selected tab or overflow lost: %q", view)
		}
	}
}

func TestTabsFirstItemWithoutPriorSelection(t *testing.T) {
	items := []navigation.Tab{{ID: "one", Label: "First"}, {ID: "", Label: "Empty ID"}}
	m := navigation.NewTabs(items, theme.New(theme.Mono))

	selected, ok := m.Selected()
	if !ok || selected.ID != "one" {
		t.Fatalf("constructor selected %+v", selected)
	}

	m.SetItems(nil)
	m.SetItems(items)

	selected, ok = m.Selected()
	if !ok || selected.ID != "one" {
		t.Fatalf("empty-to-populated selected %+v", selected)
	}

	if !m.Select("") {
		t.Fatal("empty ID is a valid selection")
	}

	m.SetItems([]navigation.Tab{{ID: "", Label: "Retained"}, {ID: "one", Label: "First"}})

	selected, ok = m.Selected()
	if !ok || selected.Label != "Retained" {
		t.Fatalf("existing empty ID lost: %+v", selected)
	}
}
