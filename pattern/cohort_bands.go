package pattern

import (
	"fmt"

	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
)

func (m *CohortComparison) bandOverviewView() string {
	matrix := chart.BandMatrix{Title: fmt.Sprintf("Version bands / host mean +/- %ds", m.sigma)}
	for i := range m.data.Metrics {
		distribution := m.distribution(i)
		for c := range distribution.Series {
			distribution.Series[c].Label = m.data.Cohorts[c].ID
		}
		title := m.metricLabel(i, false) + " " + feedback.Attention(m.assessments[i].Level, m.theme)
		if i == m.metric {
			title = "> " + title
		}
		matrix.Rows = append(matrix.Rows, chart.BandRow{Title: title, Distribution: distribution})
	}
	// The same composed summary/history panel used by detailed comparisons,
	// with distributions already visible in the matrix above it.
	summaries := m.summaries[m.metric]
	delta := "n/a"
	if summaries[0].HasRate && summaries[1].HasRate {
		delta = fmt.Sprintf("%+.3fpp", (summaries[1].Rate-summaries[0].Rate)*100)
	}
	detail := chart.ComparisonPanel{
		Title: m.metricLabel(m.metric, false),
		Summary: []string{
			fmt.Sprintf("Fleet %s %s  %s %s", m.data.Cohorts[0].ID, rate(summaries[0].Rate, summaries[0].HasRate), m.data.Cohorts[1].ID, rate(summaries[1].Rate, summaries[1].HasRate)),
			"Delta " + delta + " weighted " + feedback.Attention(m.assessments[m.metric].Fleet.Level, m.theme),
		},
		History: m.plots[m.metric],
	}
	return layout.Column(
		matrix.View(layout.Size{Width: m.gridSize.Width, Height: max(21, m.gridSize.Height-12)}, m.theme, m.focus.Current() == "metrics"),
		detail.View(layout.Size{Width: m.gridSize.Width, Height: 12}, m.theme, false),
	)
}
