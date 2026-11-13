package pattern

import (
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
)

const inspectorDetailColumnsWidth = 95

type inspectorDetails struct {
	request, response layout.Size
	columns           bool
	summary           string
}

func (m *LogInspector) resizeDetailPanes() {
	reset := m.prepareJSONSelection()
	size := m.regions.details
	requestClosed, responseClosed := m.PaneCollapsed(requestDataPane), m.PaneCollapsed(httpPane)

	regions := inspectorDetails{}
	if !requestClosed || !responseClosed {
		regions.summary = m.HTTP.SummaryContent(size.Width)
		size.Height = max(0, size.Height-lipgloss.Height(regions.summary))
	}

	regions.request, regions.response = size, size
	switch {
	case requestClosed && responseClosed:
		regions.request.Height, regions.response.Height = 1, 1
	case requestClosed:
		regions.request.Height, regions.response.Height = 1, max(0, size.Height-1)
	case responseClosed:
		regions.request.Height, regions.response.Height = max(0, size.Height-1), 1
	case size.Width >= inspectorDetailColumnsWidth:
		regions.columns = true
		regions.request.Width = max(44, (size.Width-1)/3)
		regions.response.Width = size.Width - 1 - regions.request.Width
	default:
		contentWidth := max(1, layout.PanelContentSize(size).Width-1)
		if !m.geometryOnly || m.naturalWidth != contentWidth {
			m.naturalHeight = lipgloss.Height(m.HTTP.RequestContent(contentWidth)) + 2
			m.naturalWidth = contentWidth
		}

		natural := m.naturalHeight
		regions.request.Height = min(natural, max(3, size.Height/3))
		regions.response.Height = max(0, size.Height-regions.request.Height)
	}

	m.detailRegions = regions
	for _, pane := range []struct {
		model     *viewport.Model
		size      layout.Size
		collapsed bool
		content   func(int) string
		payload   func(int) (string, collection.HTTPBody, bool)
		viewer    *collection.JSONViewer
		state     collection.JSONState
	}{{&m.Request, regions.request, requestClosed, m.HTTP.RequestContent, m.HTTP.RequestPayload, m.RequestJSON, m.jsonStates[m.detailID].request}, {&m.Response, regions.response, responseClosed, m.HTTP.ResponseContent, m.HTTP.ResponsePayload, m.ResponseJSON, m.jsonStates[m.detailID].response}} {
		oldWidth := pane.model.Width()
		if !pane.collapsed {
			content := layout.PanelContentSize(pane.size)
			pane.model.SetWidth(max(1, content.Width-inspectorScrollbarWidth(content.Width, m.HTTP.ShowScrollbar)))
			pane.model.SetHeight(max(1, content.Height))
		}

		if m.geometryOnly && !reset && (pane.collapsed || oldWidth == pane.model.Width()) {
			pane.model.SetYOffset(pane.model.YOffset())
			continue
		}

		m.refreshJSONPayload(pane.viewer, pane.model, pane.payload, pane.content, reset, pane.state)
	}
}

func inspectorScrollbarWidth(width int, enabled bool) int {
	if enabled && width > 1 {
		return 1
	}

	return 0
}

func (m *LogInspector) detailPanesView(size layout.Size) string {
	m.syncJSONFocus()
	request := m.section(requestDataPane, "Request", m.payloadViewport(requestDataPane, &m.Request, m.detailRegions.request), m.detailRegions.request)
	response := m.section(httpPane, "Response", m.payloadViewport(httpPane, &m.Response, m.detailRegions.response), m.detailRegions.response)

	content := layout.Column(request, response)
	if m.detailRegions.columns {
		content = layout.Row(request, " ", response)
	}

	if m.detailRegions.summary != "" {
		content = m.detailRegions.summary + "\n" + content
	}

	return layout.Fit(content, size)
}

func (m *LogInspector) payloadViewport(id string, model *viewport.Model, size layout.Size) string {
	if m.PaneCollapsed(id) {
		return ""
	}

	content := layout.PanelContentSize(size)
	pane := &collection.ScrollPane{Model: model, Theme: m.palette, ShowScrollbar: m.HTTP.ShowScrollbar}
	pane.Resize(content)

	return pane.View()
}

func (m *LogInspector) updateDetailPane(msg tea.Msg) tea.Cmd {
	if pressed, ok := msg.(tea.KeyPressMsg); ok {
		switch pressed.String() {
		case "3", "r":
			return m.updateHTTPControls(msg)
		}
	}

	if viewer := m.focusedJSON(); viewer != nil && viewer.Interactive() {
		return viewer.Update(msg)
	}

	var cmd tea.Cmd
	if m.focus.Current() == requestDataPane {
		m.Request, cmd = m.Request.Update(msg)
	} else {
		m.Response, cmd = m.Response.Update(msg)
	}

	return cmd
}

func (m *LogInspector) updateHTTPControls(msg tea.Msg) tea.Cmd {
	cmd := m.HTTP.UpdateDetails(msg)
	m.Resize(m.size)

	return cmd
}
