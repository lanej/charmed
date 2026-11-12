package pattern

import (
	"math"

	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/cohort"
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
}
