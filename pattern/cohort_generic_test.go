package pattern_test

import (
	"fmt"
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/attention"
	"github.com/lanej/charmed/cohort"
	"github.com/lanej/charmed/examples/cohorts/fixture"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

func TestOneToManyCohortsAndOptionalReference(t *testing.T) {
	for _, count := range []int{1, 2, 3, 8, 12} {
		t.Run(fmt.Sprint(count), func(t *testing.T) {
			d := fixture.Data()
			template := d.Windows[0].Samples[0]
			d.Cohorts = nil

			for c := range count {
				id := fmt.Sprintf("group-%d", c)
				d.Cohorts = append(d.Cohorts, cohort.Cohort{ID: id, Label: id, Short: id})
			}

			for w := range d.Windows {
				d.Windows[w].Samples = nil
				for _, group := range d.Cohorts {
					sample := template
					sample.CohortID = group.ID
					d.Windows[w].Samples = append(d.Windows[w].Samples, sample)
				}
			}

			var event map[string]any

			m, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Data: d, Theme: theme.New(theme.Mono), Emit: func(value any) {
				if v, ok := value.(map[string]any); ok {
					event = v
				}
			}})
			if err != nil {
				t.Fatal(err)
			}

			if m.ReferenceID() != "" || len(m.Comparisons("yield")) != 0 || len(m.Table.Rows()) != count {
				t.Fatal("invented reference or lost cohorts")
			}

			for _, variant := range []pattern.ComparisonVariant{pattern.Stacked, pattern.Bands} {
				if err := m.SetVariant(variant); err != nil {
					t.Fatal(err)
				}

				m.Resize(layout.Size{Width: 180, Height: 100})

				if strings.Contains(ansi.Strip(m.Panels.View()), " vs ") {
					t.Fatal("invented reference comparison")
				}

				if err := m.SetReference(d.Cohorts[0].ID); err != nil {
					t.Fatal(err)
				}

				if len(m.Comparisons("yield")) != count-1 {
					t.Fatal("not every peer compared")
				}

				for _, size := range []layout.Size{{Width: 40, Height: 16}, {Width: 180, Height: 100}} {
					m.Resize(size)

					if len(m.Table.Rows()) != count {
						t.Fatal("resize lost members")
					}
				}

				frame := ansi.Strip(m.Panels.View())
				if strings.Contains(frame, "NaN") || strings.Contains(frame, "Inf") {
					t.Fatal("nonfinite rendering")
				}

				if count == 3 && !strings.Contains(frame, "group-2 vs group-0") {
					t.Fatal("third cohort comparison missing")
				}

				if count == 12 {
					for _, size := range []layout.Size{{Width: 180, Height: 200}, {Width: 40, Height: 40}} {
						m.Resize(size)
						m.Panels.GotoBottom()

						if !strings.Contains(ansi.Strip(m.Panels.View()), "group-11:") {
							t.Fatal("last ASCII cohort history permanently clipped")
						}
					}
				}

				if err := m.SetReference(""); err != nil {
					t.Fatal(err)
				}
			}

			ids, ok := event["cohorts"].([]string)
			if !ok || event["reference"] != "" || len(ids) != count {
				t.Fatal("event still assumes two cohorts")
			}
		})
	}
}

func TestReferenceIdentityMissingMetricAndAtomicUpdates(t *testing.T) {
	d := fixture.Data()

	m, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Data: d, ReferenceID: "north", Theme: theme.New(theme.Mono)})
	if err != nil {
		t.Fatal(err)
	}

	before := m.View()
	if m.SetReference("missing") == nil || m.View() != before {
		t.Fatal("invalid reference mutated state")
	}

	d.Cohorts[0], d.Cohorts[2] = d.Cohorts[2], d.Cohorts[0]
	if err := m.SetData(d); err != nil {
		t.Fatal(err)
	}

	if m.ReferenceID() != "north" {
		t.Fatal("reference follows position")
	}

	if err := m.SetMetric("yield"); err != nil {
		t.Fatal(err)
	}
	// Navigate to a member whose temperature observation is absent.
	for range len(m.Table.Rows()) {
		if m.Selected() == (cohort.MemberKey{CohortID: "west", MemberID: "sensor-4"}) {
			break
		}

		m.Update(tea.KeyPressMsg{Code: tea.KeyDown})
	}

	if m.Selected() != (cohort.MemberKey{CohortID: "west", MemberID: "sensor-4"}) {
		t.Fatal("missing member")
	}

	rows := m.Inspector.Rows()
	if len(rows) != len(d.Metrics) || ansi.Strip(rows[0][1]) != "n/a" || !strings.Contains(rows[2][1], "%") {
		t.Fatal("missing metric shifted inspector rows or lost units")
	}

	if err := m.SetMetric("temp"); err != nil {
		t.Fatal(err)
	}

	if m.Selected() != (cohort.MemberKey{CohortID: "west", MemberID: "sensor-4"}) || len(m.Table.Rows()) != 12 || !strings.Contains(ansi.Strip(m.View()), "West / 4 members") {
		t.Fatal("missing metric dropped selected member")
	}

	if strings.Contains(ansi.Strip(m.Inspector.View()), "pp") {
		t.Fatal("scalar inherited percentage units")
	}
	// Removing an active reference is rejected before state mutation.
	d.Cohorts = d.Cohorts[:2]
	for w := range d.Windows {
		var samples []cohort.Sample

		for _, s := range d.Windows[w].Samples {
			if s.CohortID != "north" {
				samples = append(samples, s)
			}
		}

		d.Windows[w].Samples = samples
	}

	before = m.View()
	if m.SetData(d) == nil || m.View() != before {
		t.Fatal("reference removal was not atomic")
	}

	if err := m.SetReference(""); err != nil {
		t.Fatal(err)
	}

	if err := m.SetData(d); err != nil {
		t.Fatal(err)
	}
}

func TestEveryPeerGetsMetricAwareAttention(t *testing.T) {
	d := fixture.Data()
	rule := attention.Rule{Warn: &attention.Bounds{Upper: attention.Limit(5)}, Error: &attention.Bounds{Upper: attention.Limit(15)}}

	m, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Data: d, ReferenceID: "north", Theme: theme.New(theme.Mono), Boundaries: map[string]cohort.ComparisonPolicy{"temp": {Fleet: rule}}})
	if err != nil {
		t.Fatal(err)
	}

	peers := m.Comparisons("temp")
	if len(peers) != 2 || peers["south"].Fleet.Level != attention.Warn || peers["west"].Fleet.Level != attention.Error {
		t.Fatalf("wrong scalar deltas: %+v", peers)
	}

	all, ok := m.Attention("temp")
	if !ok || all.Fleet.Level != attention.Error {
		t.Fatal("aggregate omitted worst peer")
	}

	peers["south"].Members[cohort.MemberKey{}] = attention.Result{Level: attention.Error}
	if _, exists := m.Comparisons("temp")["south"].Members[cohort.MemberKey{}]; exists {
		t.Fatal("assessment alias escaped")
	}

	if err := m.SetReference(""); err != nil {
		t.Fatal(err)
	}

	all, _ = m.Attention("temp")
	if all.Fleet.Level != attention.Unknown {
		t.Fatal("no reference invented policy verdict")
	}
}

func TestLongCohortLabelsRetainNumericComparisons(t *testing.T) {
	d := fixture.Data()
	d.Cohorts[0].Label = "Reference temperature sensor population"
	d.Cohorts[1].Label = "Southern industrial temperature sensors"

	m, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Data: d, ReferenceID: "north", Theme: theme.New(theme.Mono)})
	if err != nil {
		t.Fatal(err)
	}

	for _, variant := range []pattern.ComparisonVariant{pattern.Stacked, pattern.Bands} {
		if err := m.SetVariant(variant); err != nil {
			t.Fatal(err)
		}

		for _, size := range []layout.Size{{Width: 132, Height: 40}, {Width: 40, Height: 40}} {
			m.Resize(size)

			frame := ansi.Strip(m.Panels.View())
			if !strings.Contains(frame, "+8.000C") || !strings.Contains(frame, "+15.500C") {
				t.Fatalf("labels hid scalar deltas at %+v: %s", size, frame)
			}
		}
	}
}

func TestMissingMeasurementRetainsUnknownMemberAttention(t *testing.T) {
	d := fixture.Data()

	m, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Data: d, ReferenceID: "north", Theme: theme.New(theme.Mono), Boundaries: map[string]cohort.ComparisonPolicy{"temp": {Hosts: attention.Rule{Warn: &attention.Bounds{Upper: attention.Limit(3)}}}}})
	if err != nil {
		t.Fatal(err)
	}

	missing := cohort.MemberKey{CohortID: "west", MemberID: "sensor-4"}

	assessment, ok := m.Attention("temp")
	if !ok || assessment.Members[missing].Level != attention.Unknown || assessment.Unknown != 1 {
		t.Fatalf("missing member looked unconfigured: %+v", assessment)
	}

	if m.Comparisons("temp")["west"].Members[missing].Level != attention.Unknown {
		t.Fatal("peer assessment lost missing member")
	}
}

func TestWrappedCohortLabelsDoNotClipHistory(t *testing.T) {
	d := fixture.Data()
	d.Metrics = d.Metrics[:1]
	d.Cohorts[0].Label = "Reference temperature sensor population"
	d.Cohorts[1].Label = "Southern industrial temperature sensors"

	d.Cohorts[2].Label = "Western industrial temperature sensors"
	for w := range d.Windows {
		for s := range d.Windows[w].Samples {
			delete(d.Windows[w].Samples[s].Measurements, "energy")
			delete(d.Windows[w].Samples[s].Measurements, "yield")
			delete(d.Windows[w].Samples[s].Measurements, "humidity")
		}
	}

	m, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Data: d, ReferenceID: "north", Theme: theme.New(theme.Dark)})
	if err != nil {
		t.Fatal(err)
	}

	for _, variant := range []pattern.ComparisonVariant{pattern.Stacked, pattern.Bands} {
		if err := m.SetVariant(variant); err != nil {
			t.Fatal(err)
		}

		for _, size := range []layout.Size{{Width: 132, Height: 80}, {Width: 40, Height: 80}} {
			m.Resize(size)
			m.Panels.GotoBottom()

			if !strings.Contains(ansi.Strip(m.Panels.View()), "01:00") {
				t.Fatalf("wrapped labels hid time axis in %s at %+v", variant, size)
			}
		}
	}
}
