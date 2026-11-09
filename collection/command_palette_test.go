package collection_test

import (
	"strings"
	"testing"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func paletteFixture(t *testing.T) *collection.CommandPalette {
	t.Helper()

	m, err := collection.NewCommandPalette([]collection.Command{
		{ID: "timeline", Label: "Layout: timeline", Keywords: []string{"gantt"}},
		{ID: "trace", Label: "Scope: trace", DisabledReason: "No trace selected"},
		{ID: "refresh", Label: "Refresh", Shortcut: "F5"},
	}, 100, 30, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	return m
}

func TestCommandPaletteSearchActivationCancellationAndAvailability(t *testing.T) {
	m := paletteFixture(t)
	if m.IsOpen() || m.View() != "" {
		t.Fatal("palette starts open")
	}

	m.Open()
	m.SetQuery("laytm")

	if m.SelectedID() != "timeline" || len(m.Results.Items()) != 1 {
		t.Fatal("fuzzy abbreviation failed")
	}

	m.SetQuery("gantt")

	if m.SelectedID() != "timeline" {
		t.Fatal("keywords not searchable")
	}

	cmd := m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})
	if cmd == nil || m.IsOpen() {
		t.Fatal("activation did not close")
	}

	if action, ok := cmd().(collection.CommandSelectedMsg); !ok || action.ID != "timeline" {
		t.Fatal("wrong activation")
	}

	if m.Update(tea.KeyPressMsg{Code: tea.KeyEnter}) != nil {
		t.Fatal("closed palette activated twice")
	}

	m.Open()

	if m.Search.Value() != "" {
		t.Fatal("opening retained old query")
	}

	m.SetQuery("trace")

	if m.Update(tea.KeyPressMsg{Code: tea.KeyEnter}) != nil || !m.IsOpen() || !strings.Contains(ansi.Strip(m.View()), "No trace selected") {
		t.Fatal("disabled command activated or hid reason")
	}

	m.SetQuery("zzzz")

	if m.SelectedID() != "" || !strings.Contains(ansi.Strip(m.View()), "No matching commands") || m.Update(tea.KeyPressMsg{Code: tea.KeyEnter}) != nil {
		t.Fatal("empty result failed")
	}

	m.Update(tea.KeyPressMsg{Code: tea.KeyEscape})

	if m.IsOpen() || m.Search.Focused() {
		t.Fatal("cancel retained focus")
	}
}

func TestCommandPaletteUpdatesOwnDataAndRetainNativeCustomization(t *testing.T) {
	m := paletteFixture(t)
	m.Open()
	m.Results.Select(2)
	m.Results.KeyMap.CursorDown = key.NewBinding(key.WithKeys("ctrl+j"))
	m.Search.Prompt = "find> "

	commands := []collection.Command{{ID: "refresh", Label: "Refresh", Keywords: []string{"fetch"}}, {ID: "second", Label: "Second"}}
	if err := m.SetCommands(commands); err != nil {
		t.Fatal(err)
	}

	commands[0].Label, commands[0].Keywords[0] = "changed", "changed"

	if m.SelectedID() != "refresh" {
		t.Fatal("replacement lost selected ID")
	}

	m.SetQuery("fetch")

	if m.SelectedID() != "refresh" {
		t.Fatal("keywords alias caller")
	}

	m.SetQuery("")
	m.Resize(layout.Size{Width: 40, Height: 16})
	m.Update(tea.KeyPressMsg{Code: 'j', Mod: tea.ModCtrl})

	if m.SelectedID() != "second" || m.Search.Prompt != "find> " {
		t.Fatal("resize replaced native customization")
	}

	for _, invalid := range [][]collection.Command{{{ID: "", Label: "bad"}}, {{ID: "x", Label: " "}}, {{ID: "x", Label: "A"}, {ID: "x", Label: "B"}}} {
		if err := m.SetCommands(invalid); err == nil || m.SelectedID() != "second" {
			t.Fatal("invalid update was not atomic")
		}
	}

	m.SetQuery("")
	m.Update(tea.KeyPressMsg{Code: 'j', Text: "j"})

	if m.Search.Value() != "j" {
		t.Fatal("list swallowed printable search input")
	}

	m.SetQuery("")
	m.Update(tea.PasteMsg{Content: "fetch"})

	if m.Search.Value() != "fetch" || m.SelectedID() != "refresh" {
		t.Fatal("paste did not filter commands")
	}
}

func TestCommandPaletteFitsTinyAndUnicodeViews(t *testing.T) {
	m := paletteFixture(t)
	m.Open()

	if err := m.SetCommands([]collection.Command{{ID: "unicode", Label: "Inspect 界界 café", Description: "description\x1b[31m", State: "current"}}); err != nil {
		t.Fatal(err)
	}

	for _, size := range []layout.Size{{Width: 100, Height: 30}, {Width: 40, Height: 16}, {Width: 12, Height: 7}, {Width: 1, Height: 1}, {}} {
		m.Resize(size)

		view := m.View()
		if size.Width == 0 {
			if view != "" {
				t.Fatal("zero allocation rendered")
			}

			continue
		}

		lines := strings.Split(ansi.Strip(view), "\n")
		if len(lines) > size.Height {
			t.Fatal("palette height overflow", size)
		}

		for _, line := range lines {
			if ansi.StringWidth(line) > min(size.Width, collection.PaletteMaxWidth) {
				t.Fatal("palette width overflow", size, line)
			}
		}
	}
}
