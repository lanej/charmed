package collection

import (
	"fmt"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/interaction"
	"github.com/lanej/charmed/layout"
)

type PaletteMode string

const (
	PaletteSearch   PaletteMode = "search"
	PaletteNavigate PaletteMode = "navigate"
)

func (m *CommandPalette) Mode() PaletteMode                { return m.mode }
func (m *CommandPalette) KeyLayout() interaction.KeyLayout { return m.keyLayout }

// SetKeyLayout applies a preset before native key-map overrides. An open
// palette returns to search without losing its query; forward the focus command.
func (m *CommandPalette) SetKeyLayout(preset interaction.KeyLayout) (tea.Cmd, error) {
	nav, ok := interaction.Navigation(preset)
	if !ok {
		return nil, fmt.Errorf("unknown keyboard layout %q", preset)
	}

	m.keyLayout = preset
	m.Results.KeyMap.CursorUp, m.Results.KeyMap.CursorDown = nav.Up, nav.Down
	m.Results.KeyMap.PrevPage, m.Results.KeyMap.NextPage = nav.PageUp, nav.PageDown
	m.Results.KeyMap.GoToStart, m.Results.KeyMap.GoToEnd = nav.Top, nav.Bottom

	m.mode = PaletteSearch
	if m.opened {
		return m.Search.Focus(), nil
	}

	return nil, nil
}

func (m *CommandPalette) SetBackdropMode(mode layout.BackdropMode) bool {
	if !mode.Valid() {
		return false
	}

	m.backdrop = mode

	return true
}

func (m *CommandPalette) BackdropMode() layout.BackdropMode {
	if m.backdrop != layout.DefaultBackdrop {
		return m.backdrop
	}

	if !m.DimBackdrop {
		return layout.BackdropNone
	}

	return layout.BackdropDim
}

func (m *CommandPalette) modalKey(pressed tea.KeyPressMsg) (tea.Cmd, bool) {
	if m.keyLayout != interaction.VimKeys {
		return nil, false
	}

	switch m.mode {
	case PaletteSearch:
		if key.Matches(pressed, m.Keys.Navigate) {
			m.mode = PaletteNavigate
			m.Search.Blur()

			return nil, true
		}
	case PaletteNavigate:
		if key.Matches(pressed, m.Keys.Search) {
			m.mode = PaletteSearch
			return m.Search.Focus(), true
		}
	default:
		return nil, true
	}

	return nil, false
}
