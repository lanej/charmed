package story

import (
	"errors"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/text"
)

const jsonFixture = `{"shipment":{"id":"shp_42","options":{"insurance":true,"currency":"USD"}},"rates":[{"carrier":"USPS","rate":8.500},{"carrier":"UPS","rate":12.50}],"attempt":1,"attempt":2,"sequence":9007199254740993}`

func jsonStory(e Environment) (runtime.Scene, error) {
	m := collection.NewJSONViewer(e.Config.Columns-4, e.Config.Rows-4, e.Config.Palette())
	m.Folding, m.Raw = !e.Config.NoFolds, e.Config.RawBody

	m.Symbols = text.Presentation(e.Config.SymbolStyle)
	if err := m.SetSource([]byte(jsonFixture)); err != nil {
		return nil, err
	}

	if e.Config.JSONFolded {
		m.SetCollapsed(1, true)
		m.Select(1)
	}

	if e.Config.State == "error" {
		if err := m.SetSource([]byte(`{"shipment":`)); err == nil {
			return nil, errors.New("invalid JSON fixture unexpectedly parsed")
		}
	}

	help := func(width int) string {
		if e.Config.NoHelp {
			return ""
		}

		hints := text.KeyHint("↑/↓", "scroll", e.Config.Palette())
		if m.Interactive() {
			hints = text.KeyHint("↑/↓", "fields", e.Config.Palette()) + " · " + text.KeyHint("Enter", "fold JSON", e.Config.Palette()) + " · " + text.KeyHint("←/→", "parent/expand", e.Config.Palette())
		}

		return ansi.Wrap(hints+" · "+text.KeyHint("r", "raw", e.Config.Palette()), max(1, width), "")
	}

	resize := func(size layout.Size) layout.Size {
		if footer := help(size.Width); footer != "" {
			size.Height -= lipgloss.Height(footer)
		}

		inner := layout.PanelContentSize(size)
		m.Resize(inner.Width, inner.Height)

		return size
	}

	return &scene{resize: func(size layout.Size) { resize(size) }, update: func(msg tea.Msg) tea.Cmd {
		if pressed, ok := msg.(tea.KeyPressMsg); ok && pressed.String() == "r" {
			m.Raw = !m.Raw
			m.Refresh()

			return nil
		}

		return m.Update(msg)
	}, view: func(size layout.Size) string {
		panelSize := resize(size)
		view := layout.Panel("JSON payload", m.View(), panelSize, e.Config.Palette(), true)
		footer := help(size.Width)

		if footer == "" {
			return view
		}

		return view + "\n" + footer
	}}, nil
}
