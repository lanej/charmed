// Command inspection demonstrates an application-owned service operation list.
package main

import (
	"fmt"
	"os"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/examples/inspection/fixture"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

type model struct {
	inspector      *pattern.ListInspector
	size           layout.Size
	quit, refresh  key.Binding
	attempts       int
	header, footer string
	status         string
	statusTone     text.Tone
}

func newModel() (*model, error) {
	t := theme.New(theme.Dark)
	size := layout.Size{Width: 112, Height: 30}

	inspector, err := pattern.NewListInspector(fixture.Records("ready", t), size, t)
	if err != nil {
		return nil, err
	}

	inspector.ListTitle, inspector.DetailTitle = "Operations", "Evidence"
	m := &model{
		inspector: inspector, size: size, status: "Synthetic observations / 09:42 UTC",
		quit:    key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
		refresh: key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "refresh")),
	}
	m.resize(size)

	return m, nil
}

func (m *model) Init() tea.Cmd { return nil }

func (m *model) resize(size layout.Size) {
	m.size = size
	t := m.inspector.Theme
	m.header = (layout.TextRow{Gap: 2, Items: []layout.TextRowItem{
		{Text: t.Heading().Render("Service operations")},
		{Flex: true},
		{Text: text.ToneStyle(m.statusTone, text.Normal, t).Render(m.status)},
	}}).View(max(1, size.Width))
	navigation := "select"
	toggleKey, toggleLabel := "enter", "inspect"

	if m.inspector.Inspecting() {
		navigation = "scroll"
		toggleKey, toggleLabel = "esc", "back"
	}

	m.footer = (layout.TextRow{Gap: 2, Items: []layout.TextRowItem{
		{Text: text.KeyHint("j/k", navigation, t)},
		{Text: text.KeyHint(toggleKey, toggleLabel, t)},
		{Text: text.KeyHint("[/]", "scroll", t)},
		{Text: text.KeyHint("r", "refresh", t)},
		{Text: text.KeyHint("q", "quit", t)},
	}}).View(max(1, size.Width))
	m.inspector.Resize(layout.Size{Width: size.Width, Height: max(0, size.Height-lipgloss.Height(m.header)-lipgloss.Height(m.footer))})
}

func (m *model) replaceObservation() {
	m.attempts++
	if m.attempts == 1 {
		m.status, m.statusTone = "Refresh failed / last observation retained", text.Warning
		m.resize(m.size)

		return
	}

	if err := m.inspector.SetRecords(fixture.Records("confirmed", m.inspector.Theme)); err != nil {
		m.status, m.statusTone = "Replacement rejected: "+err.Error(), text.Danger
	} else {
		m.status, m.statusTone = "Synthetic route observation / 09:43 UTC", text.Neutral
	}

	m.resize(m.size)
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if press, ok := msg.(tea.KeyPressMsg); ok {
		switch {
		case key.Matches(press, m.quit):
			return m, tea.Quit
		case key.Matches(press, m.refresh):
			m.replaceObservation()
			return m, nil
		}
	}

	if size, ok := msg.(tea.WindowSizeMsg); ok {
		m.resize(layout.Size{Width: size.Width, Height: size.Height})
		return m, nil
	}

	wasInspecting := m.inspector.Inspecting()
	cmd := m.inspector.Update(msg)

	if m.inspector.Inspecting() != wasInspecting {
		m.resize(m.size)
	}

	return m, cmd
}

func (m *model) View() tea.View {
	view := tea.NewView(layout.Fit(layout.Column(m.header, m.inspector.View(), m.footer), m.size))
	view.AltScreen = true
	view.BackgroundColor = m.inspector.Theme.Background
	view.ForegroundColor = m.inspector.Theme.Foreground

	return view
}

func main() {
	m, err := newModel()
	if err == nil {
		_, err = tea.NewProgram(m).Run()
	}

	if err != nil {
		if _, writeErr := fmt.Fprintln(os.Stderr, err); writeErr != nil {
			os.Exit(1)
		}

		os.Exit(1)
	}
}
