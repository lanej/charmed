package pattern

import (
	"fmt"

	"charm.land/bubbles/v2/table"
	"github.com/lanej/charmed/attention"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/cohort"
	"github.com/lanej/charmed/feedback"
)

func copyPolicies(policies map[string]cohort.ComparisonPolicy, metrics []cohort.Metric) (map[string]cohort.ComparisonPolicy, error) {
	known := map[string]bool{}
	for _, m := range metrics {
		known[m.ID] = true
	}
	copy := map[string]cohort.ComparisonPolicy{}
	for id, p := range policies {
		if !known[id] {
			return nil, fmt.Errorf("boundaries reference unknown metric %q", id)
		}
		if err := p.Validate(); err != nil {
			return nil, fmt.Errorf("%s: %w", id, err)
		}
		copy[id] = p.Clone()
	}
	return copy, nil
}

// SetBoundaries validates and copies policies before updating any UI state.
// An empty map leaves deviations descriptive and emits no warning/error claims.
func (m *CohortComparison) SetBoundaries(policies map[string]cohort.ComparisonPolicy) error {
	copy, err := copyPolicies(policies, m.data.Metrics)
	if err != nil {
		return err
	}
	m.boundaries = copy
	m.refresh()
	return nil
}

// Attention returns an independent assessment for a metric in the selected window.
func (m *CohortComparison) Attention(id string) (cohort.ComparisonAttention, bool) {
	for i, metric := range m.data.Metrics {
		if metric.ID == id {
			a := m.assessments[i]
			a.Members = make(map[cohort.MemberKey]attention.Result, len(m.assessments[i].Members))
			for k, v := range m.assessments[i].Members {
				a.Members[k] = v
			}
			return a, true
		}
	}
	return cohort.ComparisonAttention{}, false
}
func (m *CohortComparison) assessAttention() []attention.Changed {
	m.assessments = make([]cohort.ComparisonAttention, len(m.data.Metrics))
	var changes []attention.Changed
	next := map[string]attention.Level{}
	for i, metric := range m.data.Metrics {
		a := cohort.Assess(m.boundaries[metric.ID], m.summaries[i][0], m.summaries[i][1], m.minRequests)
		m.assessments[i] = a
		for _, state := range []struct {
			scope string
			level attention.Level
		}{{"hosts", a.Hosts}, {"fleet", a.Fleet.Level}} {
			key := m.data.Cohorts[0].ID + "\x00" + m.data.Cohorts[1].ID + "\x00" + metric.ID + "\x00" + state.scope
			if previous, ok := m.attentionLevels[key]; ok && m.initialized && previous != state.level && m.emit != nil {
				changes = append(changes, attention.Changed{Type: "attention", Metric: metric.ID, Scope: state.scope, Reference: m.data.Cohorts[0].ID, Candidate: m.data.Cohorts[1].ID, Previous: previous, Level: state.level})
			}
			next[key] = state.level
		}
	}
	m.attentionLevels = next
	return changes
}
func (m *CohortComparison) tableColumns() {
	width := m.sideSize.Width
	if width == 0 {
		width = 40
	}
	memberWidth := max(1, width-18)
	if m.boundaries[m.MetricID()].Hosts.Configured() {
		memberWidth = max(1, width-27)
	}
	columns := []table.Column{{Title: "MEMBER", Width: memberWidth}, {Title: "ROLE", Width: 4}, {Title: "z(ref)", Width: 6}}
	if m.boundaries[m.MetricID()].Hosts.Configured() {
		columns = append(columns, table.Column{Title: "ATTN", Width: 7})
	}
	m.Table.SetColumns(columns)
}
func (m *CohortComparison) boundaryMarkers(index int) []chart.BandMarker {
	reference := m.summaries[index][0]
	if !reference.HasSpread || reference.StdDev <= 0 {
		return nil
	}
	var markers []chart.BandMarker
	for _, t := range m.boundaries[m.data.Metrics[index].ID].Hosts.Thresholds() {
		glyph := 'W'
		if t.Level == attention.Error {
			glyph = 'E'
		}
		markers = append(markers, chart.BandMarker{Value: (reference.Mean + t.Value*reference.StdDev) * 100, Glyph: glyph, Style: feedback.AttentionStyle(t.Level, m.theme)})
	}
	return markers
}
