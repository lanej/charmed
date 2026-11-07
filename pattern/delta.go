package pattern

import (
	"math"

	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/cohort"
)

func (m *CohortComparison) delta(index int) (float64, bool) {
	s := m.summaries[index]
	return (s[1].Rate - s[0].Rate) * 100, s[0].HasRate && s[1].HasRate
}
func (m *CohortComparison) updateDeltaExtent() {
	magnitude := 0.0
	for _, window := range m.data.Windows {
		for _, metric := range m.data.Metrics {
			r := cohort.Analyze(window.Samples, m.data.Cohorts[0].ID, metric.ID, m.minRequests)
			c := cohort.Analyze(window.Samples, m.data.Cohorts[1].ID, metric.ID, m.minRequests)
			if r.HasRate && c.HasRate {
				magnitude = max(magnitude, math.Abs((c.Rate-r.Rate)*100))
			}
		}
	}
	for _, p := range m.boundaries {
		for _, limit := range p.Fleet.Thresholds() {
			magnitude = max(magnitude, math.Abs(limit.Value))
		}
	}
	m.deltaExtent = chart.SymmetricExtent(min(100, magnitude))
}
