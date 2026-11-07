package pattern

import (
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
)

// bandVariantView presents the same comparison with disconnected bands. It
// shares distributions, title properties, histories, selection and member tables.
func (m *CohortComparison) bandVariantView() string {
	matrix := chart.BandMatrix{Title: "Disconnected bands"}
	for i := range m.data.Metrics {
		title := m.comparisonTitle(i, m.gridSize.Width-2)
		if i == m.metric {
			title = "> " + title
		}
		matrix.Rows = append(matrix.Rows, chart.BandRow{Title: title, Distribution: m.distribution(i)})
	}
	height := max(6, m.gridSize.Height-23) + 2
	detail := chart.ComparisonPanel{Title: m.comparisonTitle(m.metric, m.gridSize.Width), History: m.plots[m.metric]}
	return layout.Column(
		matrix.View(layout.Size{Width: m.gridSize.Width, Height: max(21, m.gridSize.Height-height)}, m.theme, m.focus.Current() == "metrics"),
		detail.View(layout.Size{Width: m.gridSize.Width, Height: height}, m.theme, false),
	)
}
