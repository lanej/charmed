package collection

import (
	"fmt"
	"slices"
	"strings"

	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/input"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

const (
	PaletteMaxWidth         = 72
	PaletteMaxResults       = 8
	paletteDescriptionWidth = 56
	paletteChromeHeight     = 4
)

// Command describes an action; its execution belongs to the application.
type Command struct {
	ID, Label, Description, Shortcut, State, DisabledReason string
	Keywords                                                []string
}

type CommandSelectedMsg struct{ ID string }

type PaletteKeyMap struct{ Open, Close, Execute key.Binding }

// CommandPalette retains native search and list customization on resize and
// command updates. Results/items and their geometry are owned by the wrapper.
type CommandPalette struct {
	Search      input.Text
	Results     list.Model
	Keys        PaletteKeyMap
	Theme       theme.Theme
	DimBackdrop bool
	commands    []Command
	size        layout.Size
	opened      bool
	delegate    *commandDelegate
}

func NewCommandPalette(commands []Command, width, height int, t theme.Theme) (*CommandPalette, error) {
	d := &commandDelegate{theme: t}
	m := &CommandPalette{
		Theme: t, DimBackdrop: true, delegate: d,
		Search:  input.NewText(input.TextOptions{Placeholder: "Search commands…", Width: width, Theme: t}),
		Results: list.New(nil, d, width, height),
		Keys:    PaletteKeyMap{Open: key.NewBinding(key.WithKeys("ctrl+k"), key.WithHelp("Ctrl+K", "commands")), Close: key.NewBinding(key.WithKeys("esc"), key.WithHelp("Esc", "close")), Execute: key.NewBinding(key.WithKeys("enter"), key.WithHelp("Enter", "run"))},
	}
	m.Results.SetShowTitle(false)
	m.Results.SetShowFilter(false)
	m.Results.SetFilteringEnabled(false)
	m.Results.SetShowStatusBar(false)
	m.Results.SetShowHelp(false)
	m.Results.SetShowPagination(false)
	m.Results.DisableQuitKeybindings()
	m.Results.KeyMap.CursorUp.SetKeys("up")
	m.Results.KeyMap.CursorDown.SetKeys("down")
	m.Results.KeyMap.PrevPage.SetKeys("pgup")
	m.Results.KeyMap.NextPage.SetKeys("pgdown")
	m.Results.KeyMap.GoToStart.SetKeys("home")
	m.Results.KeyMap.GoToEnd.SetKeys("end")

	if err := m.SetCommands(commands); err != nil {
		return nil, err
	}

	m.Resize(layout.Size{Width: width, Height: height})

	return m, nil
}

func (m *CommandPalette) IsOpen() bool { return m.opened }
func (m *CommandPalette) SelectedID() string {
	if item, ok := m.Results.SelectedItem().(commandItem); ok {
		return item.command.ID
	}

	return ""
}

// SetCommands copies caller data and preserves the selected ID when still
// matched. Invalid replacement is atomic; filtering never changes the source.
func (m *CommandPalette) SetCommands(commands []Command) error {
	seen := make(map[string]bool, len(commands))
	for _, command := range commands {
		if strings.TrimSpace(command.ID) == "" || strings.TrimSpace(command.Label) == "" || seen[command.ID] {
			return fmt.Errorf("invalid or duplicate command %q", command.ID)
		}

		seen[command.ID] = true
	}

	m.commands = slices.Clone(commands)
	for i := range m.commands {
		m.commands[i].Keywords = slices.Clone(commands[i].Keywords)
	}

	m.filter(true)

	return nil
}

func (m *CommandPalette) Open() tea.Cmd {
	m.opened = true
	m.Search.SetValue("")
	m.filter(false)

	return m.Search.Focus()
}

func (m *CommandPalette) Close() {
	m.opened = false
	m.Search.Blur()
}

func (m *CommandPalette) SetQuery(query string) {
	m.Search.SetValue(query)
	m.filter(false)
}

func (m *CommandPalette) filter(preserve bool) {
	selected := m.SelectedID()

	targets := make([]string, len(m.commands))
	for i, command := range m.commands {
		targets[i] = text.Inline(command.Label + " " + strings.Join(command.Keywords, " "))
	}

	ranks := make([]list.Rank, len(m.commands))
	for i := range ranks {
		ranks[i].Index = i
	}

	if query := strings.TrimSpace(m.Search.Value()); query != "" {
		ranks = list.DefaultFilter(query, targets)
	}

	items := make([]list.Item, len(ranks))
	index := 0

	for i, rank := range ranks {
		command := m.commands[rank.Index]

		items[i] = commandItem{command: command, matches: rank.MatchedIndexes}
		if preserve && command.ID == selected {
			index = i
		}
	}

	m.Results.SetItems(items)
	m.Results.Select(index)
	m.Resize(m.size)
}

func (m *CommandPalette) Resize(size layout.Size) {
	m.size = layout.Size{Width: min(PaletteMaxWidth, max(0, size.Width)), Height: max(0, size.Height)}
	inner := layout.PanelContentSize(m.size)
	m.Search.Resize(inner.Width)
	m.delegate.descriptions = inner.Width >= paletteDescriptionWidth && inner.Height >= paletteChromeHeight+2
	rows := min(PaletteMaxResults, max(1, len(m.Results.Items())))
	height := min(max(1, inner.Height-paletteChromeHeight), rows*m.delegate.Height())
	m.Results.SetSize(max(1, inner.Width), height)
}

func (m *CommandPalette) Update(msg tea.Msg) tea.Cmd {
	if pressed, ok := msg.(tea.KeyPressMsg); ok {
		if key.Matches(pressed, m.Keys.Open) {
			if m.opened {
				m.Close()
				return nil
			}

			return m.Open()
		}

		if !m.opened {
			return nil
		}

		if key.Matches(pressed, m.Keys.Close) {
			m.Close()
			return nil
		}

		if key.Matches(pressed, m.Keys.Execute) {
			if item, found := m.Results.SelectedItem().(commandItem); found && item.command.DisabledReason == "" {
				id := item.command.ID

				m.Close()

				return func() tea.Msg { return CommandSelectedMsg{ID: id} }
			}

			return nil
		}

		keys := m.Results.KeyMap
		if key.Matches(pressed, keys.CursorUp, keys.CursorDown, keys.NextPage, keys.PrevPage, keys.GoToStart, keys.GoToEnd) {
			var cmd tea.Cmd

			m.Results, cmd = m.Results.Update(msg)

			return cmd
		}
	}

	if !m.opened {
		return nil
	}

	before := m.Search.Value()

	var cmd tea.Cmd

	m.Search, cmd = m.Search.Update(msg)
	if m.Search.Value() != before {
		m.filter(false)
	}

	return cmd
}
