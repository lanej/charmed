package pattern

import (
	"strings"

	"charm.land/lipgloss/v2"

	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
)

// bandVariantView presents the same comparison with disconnected bands. It
// shares distributions, title properties, histories, selection and member tables.
func (m *CohortComparison) bandVariantView() string {
	const selectionPrefix = "> "

	matrix := chart.BandMatrix{Title: "Disconnected bands"}

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

		matrix.Rows = append(matrix.Rows, chart.BandRow{Title: title, Distribution: m.distribution(i)})
	}

	height := m.historyHeight() + 2
	detail := chart.ComparisonPanel{Title: m.comparisonTitle(m.metric, layout.PanelContentSize(m.gridSize).Width), History: m.plots[m.metric]}

	return layout.Column(
		matrix.View(layout.Size{Width: m.gridSize.Width, Height: max(21, len(m.data.Metrics)*(2*len(m.data.Cohorts)+4))}, m.theme, m.focus.Current() == "metrics"),
		detail.View(layout.Size{Width: m.gridSize.Width, Height: height}, m.theme, false),
	)
}
