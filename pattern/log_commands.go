package pattern

import (
	"strings"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
)

func commandState(active bool) string {
	if active {
		return "current"
	}

	return ""
}

// Commands returns current availability without executing application actions.
func (m *LogInspector) Commands() []collection.Command {
	commands := []collection.Command{}
	for _, preset := range []InspectorLayout{InspectorSidebar, InspectorTimeline, InspectorColumns, InspectorFocused} {
		commands = append(commands, collection.Command{ID: "layout/" + string(preset), Label: "Layout: " + string(preset), Description: "Arrange requests, HTTP details and related logs", Keywords: []string{"view", "arrangement"}, State: commandState(m.Layout() == preset || m.Layout() == InspectorAuto && m.ActiveLayout() == preset)})
	}

	for _, pane := range []struct{ id, label, shortcut string }{{requestPane, "HTTP requests", "Alt+1"}, {httpPane, "HTTP details", "Alt+2"}, {logPane, "Related logs", "Alt+3"}} {
		commands = append(commands, collection.Command{ID: "pane/" + pane.id, Label: "Focus: " + pane.label, Description: "Move keyboard focus to this pane", Shortcut: pane.shortcut, State: commandState(m.FocusedPane() == pane.id)})
	}

	for _, follow := range []struct {
		id, label string
		active    bool
	}{{requestPane, "requests", m.HTTP.Follow}, {logPane, "logs", m.Logs.Follow}} {
		label := "Follow " + follow.label
		if follow.active {
			label = "Pause " + follow.label
		}

		commands = append(commands, collection.Command{ID: "follow/" + follow.id, Label: label, Description: "Control automatic selection of arriving observations", Keywords: []string{"follow", "pause", follow.label}})
	}

	traceReason := ""
	if selected, ok := m.HTTP.Selected(); !ok || selected.TraceID == "" {
		traceReason = "Selected request has no trace ID"
	}

	commands = append(commands,
		collection.Command{ID: "scope/request", Label: "Scope: request", Description: "Show logs matching the selected request ID", State: commandState(m.Scope() == RequestScope)},
		collection.Command{ID: "scope/trace", Label: "Scope: trace", Description: "Show logs across the selected trace", State: commandState(m.Scope() == TraceScope), DisabledReason: traceReason},
	)

	rawReason := ""
	if !m.HTTP.CanToggleRaw() {
		rawReason = "No captured body bytes available"
	}

	bodyMode := "Show raw body"
	if m.HTTP.Raw {
		bodyMode = "Show formatted body"
	}

	commands = append(commands, collection.Command{ID: "body/raw", Label: bodyMode, Description: "Switch captured body formatting", Shortcut: "r", Keywords: []string{"pretty", "syntax", "raw", "body"}, DisabledReason: rawReason})

	for _, section := range []struct {
		id, label, shortcut string
		open                bool
	}{{"request", "request", "1", m.HTTP.RequestOpen}, {"response", "response", "2", m.HTTP.ResponseOpen}, {"parameters", "parameters", "3", m.HTTP.ParametersOpen}} {
		label := "Expand " + section.label
		if section.open {
			label = "Collapse " + section.label
		}

		reason := ""
		if !m.HTTP.Folding {
			reason = "Section folding is disabled"
		} else if _, ok := m.HTTP.Selected(); !ok {
			reason = "Select an HTTP request first"
		} else if section.id == "parameters" && !m.HTTP.HasParameters() {
			reason = "No query or form parameters"
		}

		commands = append(commands, collection.Command{ID: "section/" + section.id, Label: label, Description: "Toggle HTTP detail section", Shortcut: section.shortcut, DisabledReason: reason})
	}

	for _, format := range []text.URLFormat{text.PathOnly, text.PathAndQuery, text.HostAndPath, text.FullURL} {
		current := m.HTTP.URLFormat == format || m.HTTP.URLFormat == text.DefaultURLFormat && format == text.PathOnly
		commands = append(commands, collection.Command{ID: "url/" + string(format), Label: "URL: " + string(format), Description: "Change displayed targets; captured URLs are retained", State: commandState(current && m.HTTP.URLFormatter == nil)})
	}

	return commands
}

func (m *LogInspector) OpenCommands() tea.Cmd {
	if err := m.Palette.SetCommands(m.Commands()); err != nil {
		return func() tea.Msg { return err }
	}

	return m.Palette.Open()
}

func (m *LogInspector) runCommand(id string) tea.Cmd {
	available := false

	for _, command := range m.Commands() {
		if command.ID == id && command.DisabledReason == "" {
			available = true
			break
		}
	}

	if !available {
		return nil
	}

	category, value, _ := strings.Cut(id, "/")
	switch category {
	case "layout":
		m.SetLayout(InspectorLayout(value))
	case "pane":
		m.FocusPane(value)
	case "scope":
		if value == "request" {
			m.SetScope(RequestScope)
		} else {
			m.SetScope(TraceScope)
		}
	case "url":
		m.HTTP.SetURLFormatter(nil)
		m.HTTP.SetURLFormat(text.URLFormat(value))
	case "body":
		return m.HTTP.UpdateDetails(tea.KeyPressMsg{Code: 'r'})
	case "section":
		code := map[string]rune{"request": '1', "response": '2', "parameters": '3'}[value]
		return m.HTTP.UpdateDetails(tea.KeyPressMsg{Code: code})
	case "follow":
		if value == requestPane {
			if m.HTTP.Follow {
				m.HTTP.Follow = false
			} else {
				m.HTTP.ResumeFollow()
				m.correlate()
			}
		} else {
			if m.Logs.Follow {
				m.Logs.Follow = false
			} else {
				m.Logs.ResumeFollow()
			}
		}
	}

	return nil
}

func (m *LogInspector) Update(msg tea.Msg) tea.Cmd {
	if selected, ok := msg.(collection.CommandSelectedMsg); ok {
		return m.runCommand(selected.ID)
	}

	if pressed, ok := msg.(tea.KeyPressMsg); ok {
		if !m.Palette.IsOpen() && key.Matches(pressed, m.Palette.Keys.Open) {
			return m.OpenCommands()
		}

		if m.Palette.IsOpen() {
			if err := m.Palette.SetCommands(m.Commands()); err != nil {
				return func() tea.Msg { return err }
			}

			return m.Palette.Update(msg)
		}
	}

	if m.Palette.IsOpen() {
		if _, ok := msg.(tea.PasteMsg); ok {
			return m.Palette.Update(msg)
		}

		return tea.Batch(m.Palette.Update(msg), m.updateContent(msg))
	}

	return m.updateContent(msg)
}

func (m *LogInspector) View() string {
	base := m.contentView()
	if !m.Palette.IsOpen() {
		return base
	}

	return layout.CenterOverlayWithBackdrop(base, m.Palette.View(), m.size, m.palette, m.Palette.BackdropMode())
}
