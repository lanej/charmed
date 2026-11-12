package pattern

import (
	"fmt"
	"math"

	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/cohort"
)

func validateReference(data cohort.Data, id string) error {
	if id == "" {
		return nil
	}

	for _, c := range data.Cohorts {
		if c.ID == id {
			return nil
		}
	}

	return fmt.Errorf("unknown reference cohort %q", id)
}

func (m *CohortComparison) ReferenceID() string { return m.referenceID }

// SetReference selects a baseline by identity, or disables comparisons with "".
// Invalid references leave all state unchanged. SetData rejects removing an active
// reference; clear or change it first. Resize and cohort reordering retain it.
func (m *CohortComparison) SetReference(id string) error {
	if err := validateReference(m.data, id); err != nil {
		return err
	}

	m.referenceID = id
	m.refresh()
	m.event("reference")

	return nil
}

func (m *CohortComparison) reference(index int) cohort.Summary {
	for c, group := range m.data.Cohorts {
		if group.ID == m.referenceID {
			return m.summaries[index][c]
		}
	}

	return cohort.Summary{}
}

func (m *CohortComparison) referenceLabel() string {
	for _, c := range m.data.Cohorts {
		if c.ID == m.referenceID {
			return c.Label
		}
	}

	return ""
}

func (m *CohortComparison) cohortIDs() []string {
	ids := make([]string, len(m.data.Cohorts))
	for i, c := range m.data.Cohorts {
		ids[i] = c.ID
	}

	return ids
}

func (m *CohortComparison) distributionRange(index int) (float64, float64) {
	metric := m.data.Metrics[index]
	low, high := math.Inf(1), math.Inf(-1)

	for _, s := range m.summaries[index] {
		if s.Eligible > 0 {
			low = min(low, max(-math.MaxFloat64, s.Mean-3*s.StdDev))
			high = max(high, min(math.MaxFloat64, s.Mean+3*s.StdDev))
		}

		for _, p := range s.Points {
			if p.HasValue {
				low = min(low, p.Value)
				high = max(high, p.Value)
			}
		}
	}

	for _, marker := range m.boundaryMarkers(index) {
		low = min(low, marker.Value/metric.Scale())
		high = max(high, marker.Value/metric.Scale())
	}

	if math.IsInf(low, 1) {
		return 0, 1
	}

	low, high, ok := chart.PaddedRange(low*metric.Scale(), high*metric.Scale(), 0.1)
	if !ok {
		return 0, 0
	}

	return low, high
}

// memberSummary supplements observations with unavailable member identities for
// selection and attention, without changing measured counts or aggregate values.
func (m *CohortComparison) memberSummary(index, c int) cohort.Summary {
	summary := m.summaries[index][c]

	summary.Points = append([]cohort.MemberRate(nil), summary.Points...)
	if m.window < 0 {
		return summary
	}

	present := make(map[cohort.MemberKey]bool, len(summary.Points))
	for _, p := range summary.Points {
		present[p.MemberKey] = true
	}

	for _, sample := range m.data.Windows[m.window].Samples {
		if sample.CohortID == m.data.Cohorts[c].ID && !present[sample.MemberKey] {
			summary.Points = append(summary.Points, cohort.MemberRate{MemberKey: sample.MemberKey})
		}
	}

	return summary
}
