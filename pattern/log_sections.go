package pattern

import "github.com/lanej/charmed/layout"

// SetPaneCollapsed folds a pane to its heading, retaining selection, captured
// records, native customization and the hidden viewport's scroll position.
// IDs are "requests" (timeline/list), "request", "response" (also "http") and
// "logs". Invalid IDs and disabled folds do nothing.
func (m *LogInspector) SetPaneCollapsed(id string, collapsed bool) bool {
	id = inspectorPaneID(id)
	if !m.folding || id != requestPane && id != requestDataPane && id != httpPane && id != logPane {
		return false
	}

	switch id {
	case requestDataPane:
		m.HTTP.RequestOpen = !collapsed
	case httpPane:
		m.HTTP.ResponseOpen = !collapsed
	default:
		m.collapsed[id] = collapsed
	}

	m.Resize(m.size)

	return true
}

func (m *LogInspector) PaneCollapsed(id string) bool {
	switch inspectorPaneID(id) {
	case requestDataPane:
		return m.HTTP.Folding && !m.HTTP.RequestOpen
	case httpPane:
		return m.HTTP.Folding && !m.HTTP.ResponseOpen
	default:
		return m.collapsed[inspectorPaneID(id)]
	}
}

func inspectorPaneID(id string) string {
	if id == "response" {
		return httpPane
	}

	return id
}

// SetFolding applies to panes, HTTP sections and log fields. Disabling it
// expands all content; re-enabling it keeps that expanded state.
func (m *LogInspector) SetFolding(enabled bool) {
	m.folding = enabled
	if !enabled {
		clear(m.collapsed)
		clear(m.jsonStates)
	}

	m.HTTP.SetFolding(enabled)
	m.Logs.SetFolding(enabled)
	m.Resize(m.size)
}

// SetHelpVisible returns the two instruction rows to the content when hidden.
// Keyboard bindings and the command palette remain available, including ?.
func (m *LogInspector) SetHelpVisible(visible bool) {
	m.showHelp = visible
	m.Palette.SetHelpVisible(visible)
	m.Resize(m.size)
}

func (m *LogInspector) HelpVisible() bool { return m.showHelp }

func (m *LogInspector) headlineView() string {
	if !m.showHeadline {
		return ""
	}

	record, ok := m.HTTP.Selected()
	if !ok {
		return ""
	}

	return (layout.IdentityHeader{Kind: "Request", ID: record.RequestID}).View(m.size.Width, m.palette)
}

func (m *LogInspector) section(id, title, content string, size layout.Size) string {
	return (layout.Section{Title: title, Content: content, Collapsed: m.PaneCollapsed(id), NoFolds: !m.folding, Symbols: m.Symbols}).View(size, m.palette, m.focus.Current() == id)
}
