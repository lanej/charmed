package pattern

import (
	"strings"

	"charm.land/lipgloss/v2"

	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

// bandVariantView presents the same comparison with disconnected bands. It
// shares distributions, title properties, histories, selection and member tables.
func (m *CohortComparison) bandVariantView() string {
	const selectionPrefix = "> "

	matrix := chart.BandMatrix{Title: "Disconnected bands"}
	matrixHeight := 2 * theme.BorderCells

	for i := range m.data.Metrics {
		width := layout.PanelContentSize(m.gridSize).Width
		if i == m.metric {
			width -= lipgloss.Width(selectionPrefix)
		}

		title := m.comparisonTitle(i, width)
		if rows := m.comparisonRows(i, width); len(rows) > 0 {
			title += "\n" + strings.Join(rows, "\n")
		}

		if i == m.metric {
			title = selectionPrefix + title
		}

		distribution := m.distribution(i)
		matrix.Rows = append(matrix.Rows, chart.BandRow{Title: title, Distribution: distribution})
		distribution.Compact = true

		if i > 0 {
			matrixHeight += 2
		}

		matrixHeight += lipgloss.Height(title + "\n" + distribution.View(max(1, layout.PanelContentSize(m.gridSize).Width), m.theme))
	}

	height := m.historyHeight() + 2

	notes := m.historyRows(m.metric, layout.PanelContentSize(m.gridSize).Width)
	if len(notes) > 0 {
		height += lipgloss.Height(strings.Join(notes, "\n")) + 1
	}

	detail := chart.ComparisonPanel{Summary: notes, Title: m.comparisonTitle(m.metric, layout.PanelContentSize(m.gridSize).Width), History: m.plots[m.metric]}

	return layout.Column(
		matrix.View(layout.Size{Width: m.gridSize.Width, Height: matrixHeight}, m.theme, m.focus.Current() == "metrics"),
		detail.View(layout.Size{Width: m.gridSize.Width, Height: height}, m.theme, false),
	)
}
