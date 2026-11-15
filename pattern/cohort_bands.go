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

	labelWidth, valueWidth := m.comparisonColumns(layout.PanelContentSize(m.gridSize).Width)
	matrix := chart.BandMatrix{Title: "Disconnected bands", LabelWidth: labelWidth, ValueWidth: valueWidth}

	for i := range m.data.Metrics {
		width := layout.PanelContentSize(m.gridSize).Width

		marker := "  "
		if i == m.metric {
			marker = selectionPrefix
		}

		title := marker + m.comparisonTitle(i, max(1, width-lipgloss.Width(marker)))

		distribution := m.distribution(i)
		matrix.Rows = append(matrix.Rows, chart.BandRow{Title: title, Distribution: distribution, Comparisons: m.comparisonItems(i)})
	}

	matrixHeight := matrix.Height(m.gridSize.Width, m.theme)

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
