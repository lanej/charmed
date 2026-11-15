package pattern

import (
	"math"

	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/cohort"
	"github.com/lanej/charmed/text"
)

func (m *CohortComparison) cohortDelta(index, candidate int) (float64, bool) {
	r, c := m.reference(index), m.summaries[index][candidate]
	return (c.Value - r.Value) * m.data.Metrics[index].Scale(), m.referenceID != "" && r.HasValue && c.HasValue
}

// Each metric keeps a stable domain across windows and cohorts, in its own unit.
func (m *CohortComparison) metricDeltaExtent(index int) float64 {
	metric := m.data.Metrics[index]
	magnitude := 0.0
	for _, window := range m.data.Windows {
		r := cohort.AnalyzeMetric(window.Samples, m.referenceID, metric, m.minRequests)
		for _, group := range m.data.Cohorts {
			if group.ID == m.referenceID {
				continue
			}

			c := cohort.AnalyzeMetric(window.Samples, group.ID, metric, m.minRequests)
			if r.HasValue && c.HasValue {
				magnitude = max(magnitude, math.Abs((c.Value-r.Value)*metric.Scale()))
			}
		}
	}

	for _, limit := range m.boundaries[metric.ID].Fleet.Thresholds() {
		magnitude = max(magnitude, math.Abs(limit.Value))
	}

	return chart.SymmetricExtent(magnitude)
}

func (m *CohortComparison) updateDeltaExtent() {
	m.deltaExtent = m.metricDeltaExtent(m.metric)
	m.comparisonLabelWidth, m.deltaLabelWidth = m.measureComparisonColumns()
}

// Column widths span all metrics and windows; selection and window changes cannot
// move delta baselines. Numeric domains remain specific to each metric.
func (m *CohortComparison) measureComparisonColumns() (int, int) {
	labelWidth, numberWidth := 0, 3

	for _, group := range m.data.Cohorts {
		if group.ID != m.referenceID {
			labelWidth = max(labelWidth, lipgloss.Width(group.Label+" vs "+m.referenceLabel()))
		}
	}

	for _, metric := range m.data.Metrics {
		for _, window := range m.data.Windows {
			reference := cohort.AnalyzeMetric(window.Samples, m.referenceID, metric, m.minRequests)
			for _, group := range m.data.Cohorts {
				candidate := cohort.AnalyzeMetric(window.Samples, group.ID, metric, m.minRequests)
				mean := text.AbbreviatedNumber(candidate.Mean*metric.Scale(), 3, metric.DisplayUnit(), candidate.Eligible > 0)
				numberWidth = max(numberWidth, lipgloss.Width(mean))

				if group.ID == m.referenceID {
					continue
				}

				label := (chart.DeltaProperty{Value: (candidate.Value - reference.Value) * metric.Scale(), Extent: 1, Available: reference.HasValue && candidate.HasValue, Unit: metric.DeltaUnit(), Unitless: metric.DeltaUnit() == ""}).Label()
				numberWidth = max(numberWidth, lipgloss.Width(label))
			}
		}
	}

	return labelWidth, numberWidth
}

func (m *CohortComparison) comparisonColumns(width int) (int, int) {
	return min(m.comparisonLabelWidth, max(0, width-m.deltaLabelWidth-5)), m.deltaLabelWidth
}
