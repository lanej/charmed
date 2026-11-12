package pattern

import (
	"fmt"
	"maps"

	"charm.land/bubbles/v2/table"
	"github.com/lanej/charmed/attention"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/cohort"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
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

// Comparisons returns independent per-cohort assessments against the selected
// reference. The reference itself is excluded; without a reference the map is empty.
func (m *CohortComparison) Comparisons(id string) map[string]cohort.ComparisonAttention {
	result := map[string]cohort.ComparisonAttention{}

	for i, metric := range m.data.Metrics {
		if metric.ID != id {
			continue
		}

		for id, value := range m.comparisons[i] {
			value.Members = maps.Clone(value.Members)
			result[id] = value
		}
	}

	return result
}

func (m *CohortComparison) assessAttention() []attention.Changed {
	m.assessments = make([]cohort.ComparisonAttention, len(m.data.Metrics))
	m.comparisons = make([]map[string]cohort.ComparisonAttention, len(m.data.Metrics))
	var changes []attention.Changed
	next := map[string]attention.Level{}
	for i, metric := range m.data.Metrics {
		policy := m.boundaries[metric.ID]
		reference := m.reference(i)
		// Evaluate member attention once, so reference members are not counted once
		// for every peer cohort. Fleet attention is the worst peer comparison.
		all := cohort.Summary{}

		for c := range m.data.Cohorts {
			summary := m.memberSummary(i, c)
			all.Points = append(all.Points, summary.Points...)
		}

		reference.Points = nil
		for _, point := range all.Points {
			if point.CohortID == m.referenceID {
				reference.Points = append(reference.Points, point)
			}
		}
		// The member pass includes every cohort against the chosen reference.
		memberPolicy := policy
		memberPolicy.Fleet = attention.Rule{}
		a := cohort.AssessMetric(memberPolicy, referenceWithoutPoints(reference), all, metric, m.minRequests)
		a.Fleet = policy.Fleet.Evaluate(0, false)

		m.comparisons[i] = map[string]cohort.ComparisonAttention{}
		if m.referenceID != "" {
			first := true

			for c, group := range m.data.Cohorts {
				if group.ID == m.referenceID {
					continue
				}

				pair := cohort.AssessMetric(policy, reference, m.memberSummary(i, c), metric, m.minRequests)

				m.comparisons[i][group.ID] = pair
				if first || attention.Priority(pair.Fleet.Level) > attention.Priority(a.Fleet.Level) {
					a.Fleet = pair.Fleet
				}

				first = false

				for _, state := range []struct {
					scope string
					level attention.Level
				}{{"hosts", pair.Hosts}, {"fleet", pair.Fleet.Level}} {
					key := m.referenceID + "\x00" + group.ID + "\x00" + metric.ID + "\x00" + state.scope
					if previous, ok := m.attentionLevels[key]; ok && m.initialized && previous != state.level && m.emit != nil {
						changes = append(changes, attention.Changed{Type: "attention", Metric: metric.ID, Scope: state.scope, Reference: m.referenceID, Candidate: group.ID, Previous: previous, Level: state.level})
					}
					next[key] = state.level
				}
			}
		}

		a.Level = attention.Worst(a.Hosts, a.Fleet.Level)
		m.assessments[i] = a
	}
	m.attentionLevels = next
	return changes
}

func referenceWithoutPoints(summary cohort.Summary) cohort.Summary {
	summary.Points = nil
	return summary
}

func (m *CohortComparison) tableColumns() {
	width := m.sideSize.Width
	if width == 0 {
		width = 40
	}

	width = layout.PanelContentSize(layout.Size{Width: width}).Width

	const versionWidth, distanceWidth = 7, 6

	columns := []table.Column{{Title: "MEMBER", Width: 1}, {Title: "COHORT", Width: versionWidth}, {Title: "z(ref)", Width: distanceWidth}}
	if m.boundaries[m.MetricID()].Hosts.Configured() {
		columns = append(columns, table.Column{Title: "ATTN", Width: feedback.MaxAttentionWidth(m.theme)})
	}

	occupied := collection.TablePaddingWidth(len(columns))
	for _, column := range columns[1:] {
		occupied += column.Width
	}

	columns[0].Width = max(1, width-occupied)

	m.Table.SetColumns(columns)
}

func (m *CohortComparison) boundaryMarkers(index int) []chart.BandMarker {
	reference := m.reference(index)
	if !reference.HasSpread || reference.StdDev <= 0 {
		return nil
	}
	var markers []chart.BandMarker
	for _, t := range m.boundaries[m.data.Metrics[index].ID].Hosts.Thresholds() {
		glyph := 'W'
		if t.Level == attention.Error {
			glyph = 'E'
		}

		markers = append(markers, chart.BandMarker{Value: (reference.Mean + t.Value*reference.StdDev) * m.data.Metrics[index].Scale(), Glyph: glyph, Style: feedback.AttentionStyle(t.Level, m.theme)})
	}
	return markers
}
