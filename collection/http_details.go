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

func (m *HTTPLog) detailSections(record HTTPExchange, width int) string {
	if width < httpParameterMinWidth+httpBodyMinWidth+httpDetailColumnGap {
		parts := []string{}
		if parameters := m.parameterView(record, width); parameters != "" {
			parts = append(parts, parameters)
		}

		parts = append(parts, m.exchangeSection("Request", m.RequestOpen, record.RequestHeaders, record.RequestBody, width), m.exchangeSection("Response", m.ResponseOpen, record.ResponseHeaders, record.ResponseBody, width))

		return strings.Join(parts, "\n\n")
	}

	leftWidth := max(httpParameterMinWidth, (width-httpDetailColumnGap)/3)
	rightWidth := width - leftWidth - httpDetailColumnGap

	left := []string{}
	if parameters := m.parameterView(record, leftWidth); parameters != "" {
		left = append(left, parameters)
	}

	left = append(left, m.exchangeSection("Request", m.RequestOpen, record.RequestHeaders, record.RequestBody, leftWidth))
	leftContent := ansi.Wrap(strings.Join(left, "\n\n"), leftWidth, "")
	rightContent := m.exchangeSection("Response", m.ResponseOpen, record.ResponseHeaders, record.ResponseBody, rightWidth)
	leftContent = layout.Fit(leftContent, layout.Size{Width: leftWidth, Height: lipgloss.Height(leftContent)})

	return layout.Row(leftContent, strings.Repeat(" ", httpDetailColumnGap), rightContent)
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
