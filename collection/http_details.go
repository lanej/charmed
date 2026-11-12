package collection

import (
	"net/http"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
)

const (
	httpParameterMinWidth = 24
	httpBodyMinWidth      = 35
	httpDetailColumnGap   = 3
)

// SummaryContent renders the selected exchange's metadata without bodies or
// keyboard instructions. It does not change the native combined Details view.
func (m *HTTPLog) SummaryContent(width int) string {
	record, ok := m.Selected()
	if !ok {
		return "Select an HTTP request"
	}

	rows := m.summaryRows(record, max(1, width))
	if record.TraceID != "" {
		rows = append(rows, detailLine("Trace ID", emptyID(record.TraceID), max(1, width)))
	}

	return ansi.Wrap(strings.Join(rows, "\n"), max(1, width), "")
}

// RequestContent and ResponseContent expose the selected directional payloads
// for independent native viewports. Raw/capture/parameter policies remain owned
// by HTTPLog; callers own pane folding, sizing, scrolling and focus.
func (m *HTTPLog) RequestContent(width int) string {
	prefix, body, ok := m.RequestPayload(width)
	if !ok {
		return "Select an HTTP request"
	}

	return prefix + "\n" + ansi.Wrap(body.View(m.Raw, m.BodyPreviewBytes, m.Theme), max(1, width), "")
}

// RequestPayload and ResponsePayload expose a styled header/parameter prefix
// and copied body for interactive payload renderers. Capture metadata and the
// content-type fallback are identical to the standalone content views.
func (m *HTTPLog) RequestPayload(width int) (string, HTTPBody, bool) {
	record, ok := m.Selected()
	if !ok {
		return "", HTTPBody{}, false
	}

	parts := []string{}
	if parameters := m.parameterView(record, max(1, width)); parameters != "" {
		parts = append(parts, parameters)
	}

	parts = append(parts, headerView(record.RequestHeaders, m.Theme, max(1, width)))

	return strings.Join(parts, "\n\n"), bodyWithContentType(record.RequestBody, record.RequestHeaders), true
}

func (m *HTTPLog) ResponseContent(width int) string {
	prefix, body, ok := m.ResponsePayload(width)
	if !ok {
		return "Select an HTTP request"
	}

	return prefix + "\n" + ansi.Wrap(body.View(m.Raw, m.BodyPreviewBytes, m.Theme), max(1, width), "")
}

func (m *HTTPLog) ResponsePayload(width int) (string, HTTPBody, bool) {
	record, ok := m.Selected()
	if !ok {
		return "", HTTPBody{}, false
	}

	return headerView(record.ResponseHeaders, m.Theme, max(1, width)), bodyWithContentType(record.ResponseBody, record.ResponseHeaders), true
}

func (m *HTTPLog) detailSections(record HTTPExchange, width int) string {
	left := m.requestSections(record, width)
	right := m.exchangeSection("Response", m.ResponseOpen, record.ResponseHeaders, record.ResponseBody, width)
	leftMinimum := max(httpParameterMinWidth, lipgloss.Width(left))
	rightMinimum := max(httpBodyMinWidth, lipgloss.Width(right))
	// WHY: A fixed breakpoint can split UUIDs and payloads even when a stacked
	// view would display the captured lines intact. Closed columns need no width.
	if width < leftMinimum+rightMinimum+httpDetailColumnGap || m.Folding && (!m.ResponseOpen || !m.RequestOpen && (!m.HasParameters() || !m.ParametersOpen)) {
		return left + "\n\n" + right
	}

	leftWidth := min(width-httpDetailColumnGap-rightMinimum, max(leftMinimum, (width-httpDetailColumnGap)/3))
	rightWidth := width - leftWidth - httpDetailColumnGap
	leftContent := m.requestSections(record, leftWidth)
	rightContent := m.exchangeSection("Response", m.ResponseOpen, record.ResponseHeaders, record.ResponseBody, rightWidth)
	leftContent = layout.Fit(leftContent, layout.Size{Width: leftWidth, Height: lipgloss.Height(leftContent)})

	return layout.Row(leftContent, strings.Repeat(" ", httpDetailColumnGap), rightContent)
}

func (m *HTTPLog) requestSections(record HTTPExchange, width int) string {
	left := []string{}
	if parameters := m.parameterView(record, width); parameters != "" {
		left = append(left, parameters)
	}

	left = append(left, m.exchangeSection("Request", m.RequestOpen, record.RequestHeaders, record.RequestBody, width))

	return ansi.Wrap(strings.Join(left, "\n\n"), width, "")
}

func (m *HTTPLog) exchangeSection(name string, open bool, headers http.Header, body HTTPBody, width int) string {
	heading := m.Theme.Heading().Render(name)
	if m.Folding {
		if marker := text.Disclosure("", open, m.Theme, m.Symbols); marker != "" {
			heading = marker + " " + heading
		}
	}

	if m.Folding && !open {
		return heading
	}

	content := strings.Join([]string{heading, headerView(headers, m.Theme, width), bodyWithContentType(body, headers).View(m.Raw, m.BodyPreviewBytes, m.Theme)}, "\n")

	return ansi.Wrap(content, width, "")
}
