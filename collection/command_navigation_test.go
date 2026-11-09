package collection_test

import (
	"strings"
	"testing"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/interaction"
	"github.com/lanej/charmed/layout"
)

func TestVimPaletteSeparatesSearchAndNavigation(t *testing.T) {
	m := paletteFixture(t)
	if _, err := m.SetKeyLayout(interaction.VimKeys); err != nil {
		t.Fatal(err)
	}

	m.Open()
	m.Update(tea.KeyPressMsg{Code: 'j', Text: "j"})

	if m.Search.Value() != "j" || m.Mode() != collection.PaletteSearch {
		t.Fatal("Vim movement stole search input")
	}

	m.SetQuery("")
	m.Update(tea.KeyPressMsg{Code: tea.KeyEscape})

	if !m.IsOpen() || m.Mode() != collection.PaletteNavigate || m.Search.Focused() {
		t.Fatal("Esc did not enter navigation")
	}

	m.Update(tea.KeyPressMsg{Code: 'j', Text: "j"})

	if m.SelectedID() != "trace" || m.Search.Value() != "" {
		t.Fatal("j did not move")
	}

	m.Update(tea.KeyPressMsg{Code: 'k', Text: "k"})

	if m.SelectedID() != "timeline" {
		t.Fatal("k did not move")
	}

	m.Update(tea.KeyPressMsg{Code: '/', Text: "/"})

	if m.Mode() != collection.PaletteSearch || !m.Search.Focused() || m.Search.Value() != "" {
		t.Fatal("search resume inserted slash or lost focus")
	}

	m.SetQuery("refresh")
	m.Update(tea.KeyPressMsg{Code: tea.KeyEscape})
	m.Resize(layout.Size{Width: 40, Height: 16})

	if m.Mode() != collection.PaletteNavigate || m.SelectedID() != "refresh" {
		t.Fatal("resize reset mode or selection")
	}

	cmd := m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})
	if cmd == nil || m.IsOpen() {
		t.Fatal("normal mode activation failed")
	}

	if action, ok := cmd().(collection.CommandSelectedMsg); !ok || action.ID != "refresh" {
		t.Fatal("wrong action")
	}

	m.Open()

	if m.Mode() != collection.PaletteSearch || m.Search.Value() != "" {
		t.Fatal("reopening did not reset search mode")
	}

	m.Update(tea.KeyPressMsg{Code: tea.KeyEscape})
	m.Update(tea.KeyPressMsg{Code: tea.KeyEscape})

	if m.IsOpen() {
		t.Fatal("second Esc did not dismiss")
	}
}

func TestPaletteNavigationOverridesAndBackdropConfiguration(t *testing.T) {
	m := paletteFixture(t)
	if m.Results.KeyMap.CursorUp.Help().Key != "↑" || m.Results.KeyMap.CursorDown.Help().Key != "↓" {
		t.Fatal("default movement hints advertise inactive aliases")
	}

	m.Open()

	if _, err := m.SetKeyLayout(interaction.VimKeys); err != nil {
		t.Fatal(err)
	}

	m.Results.KeyMap.CursorDown = key.NewBinding(key.WithKeys("ctrl+j"), key.WithHelp("C-j", "down"))
	m.Update(tea.KeyPressMsg{Code: tea.KeyEscape})

	if err := m.SetCommands([]collection.Command{{ID: "one", Label: "One"}, {ID: "two", Label: "Two"}}); err != nil {
		t.Fatal(err)
	}

	m.Resize(layout.Size{Width: 72, Height: 24})
	m.Update(tea.KeyPressMsg{Code: 'j', Mod: tea.ModCtrl})

	if m.SelectedID() != "two" || !strings.Contains(ansi.Strip(m.View()), "C-j") {
		t.Fatal("custom navigation or its hint lost")
	}

	if _, err := m.SetKeyLayout("invalid"); err == nil || m.KeyLayout() != interaction.VimKeys {
		t.Fatal("invalid layout mutated state")
	}

	if m.SetBackdropMode("invalid") || m.BackdropMode() != layout.BackdropDim {
		t.Fatal("invalid backdrop mutated state")
	}

	for _, mode := range []layout.BackdropMode{layout.BackdropDim, layout.BackdropMonochrome, layout.BackdropNone} {
		if !m.SetBackdropMode(mode) || m.BackdropMode() != mode {
			t.Fatal("backdrop option ignored")
		}
	}

	m.DimBackdrop = false
	if !m.SetBackdropMode(layout.DefaultBackdrop) || m.BackdropMode() != layout.BackdropNone {
		t.Fatal("legacy backdrop opt-out lost")
	}

	if !m.SetBackdropMode(layout.BackdropDim) || m.BackdropMode() != layout.BackdropDim {
		t.Fatal("explicit mode did not take precedence")
	}
}
