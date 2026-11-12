package pattern_test

import (
	"fmt"
	"math"
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

	if err := m.SetMetric("yield"); err != nil {
		t.Fatal(err)
	}

	for range len(m.Table.Rows()) {
		if m.Selected() == missing {
			break
		}

		m.Update(tea.KeyPressMsg{Code: tea.KeyDown})
	}

	if m.Selected() != missing || !strings.Contains(ansi.Strip(m.Inspector.Rows()[0][3]), "UNKNOWN") {
		t.Fatal("missing inspector measurement did not retain Unknown attention")
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

func TestInspectorPreservesTinyNonzeroMeasurements(t *testing.T) {
	d := fixture.Data()
	d.Metrics = []cohort.Metric{{ID: "ratio", Label: "Ratio"}, {ID: "tiny", Label: "Tiny", Kind: cohort.Scalar, Unit: "C"}, {ID: "negative", Label: "Negative", Kind: cohort.Scalar, Unit: "C"}, {ID: "zero", Label: "Zero", Kind: cohort.Scalar, Unit: "C"}}
	d.Cohorts = d.Cohorts[:1]
	d.Windows = d.Windows[:1]
	d.Windows[0].Samples = []cohort.Sample{{MemberKey: cohort.MemberKey{CohortID: "north", MemberID: "sensor"}, Measurements: map[string]cohort.Measurement{"ratio": {Numerator: 1, Denominator: 1000000}, "tiny": {Value: 0.000001}, "negative": {Value: -0.000001}, "zero": {Value: 0}}}}

	m, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Data: d, Theme: theme.New(theme.Mono)})
	if err != nil {
		t.Fatal(err)
	}

	for i, want := range []string{"0.0001%", "1e-06C", "-1e-06C", "0.000C"} {
		if got := ansi.Strip(m.Inspector.Rows()[i][1]); got != want {
			t.Fatalf("metric %s: got %q, want %q", d.Metrics[i].ID, got, want)
		}
	}
}

func TestHistoryExposesMissingIntervals(t *testing.T) {
	d := fixture.Data()
	d.Cohorts = d.Cohorts[:1]
	d.Metrics = []cohort.Metric{{ID: "value", Label: "Value", Kind: cohort.Scalar}}

	d.Windows = d.Windows[:3]
	for i := range d.Windows {
		d.Windows[i].Samples = []cohort.Sample{{MemberKey: cohort.MemberKey{CohortID: "north", MemberID: "sensor"}, Measurements: map[string]cohort.Measurement{"value": {Value: 42}}}}
	}

	complete, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Data: d, Theme: theme.New(theme.Dark)})
	if err != nil {
		t.Fatal(err)
	}

	delete(d.Windows[1].Samples[0].Measurements, "value")

	missing, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Data: d, Theme: theme.New(theme.Dark)})
	if err != nil {
		t.Fatal(err)
	}

	for _, variant := range []pattern.ComparisonVariant{pattern.Stacked, pattern.Bands} {
		if err := complete.SetVariant(variant); err != nil {
			t.Fatal(err)
		}

		if err := missing.SetVariant(variant); err != nil {
			t.Fatal(err)
		}

		complete.Resize(layout.Size{Width: 150, Height: 100})
		missing.Resize(layout.Size{Width: 150, Height: 100})

		frame := ansi.Strip(missing.Panels.View())
		if frame == ansi.Strip(complete.Panels.View()) || !strings.Contains(strings.Join(strings.Fields(frame), " "), "bridge missing intervals") || !strings.Contains(strings.ReplaceAll(strings.Join(strings.Fields(frame), ""), "│", ""), "2026-01-01T01:00:00Z–2026-01-01T02:00:00Z") {
			t.Fatalf("unavailable window was indistinguishable from a measured value in %s: %s", variant, frame)
		}
	}
}

func TestExtremeScalarHistoryRetainsFiniteObservationRange(t *testing.T) {
	for _, value := range []float64{1e-200, 1e200, math.MaxFloat64} {
		d := fixture.Data()
		d.Cohorts = d.Cohorts[:1]
		d.Metrics = []cohort.Metric{{ID: "value", Label: "Value", Kind: cohort.Scalar, Unit: "C"}}
		d.Windows = d.Windows[:1]
		d.Windows[0].Samples = []cohort.Sample{{MemberKey: cohort.MemberKey{CohortID: "north", MemberID: "sensor"}, Measurements: map[string]cohort.Measurement{"value": {Value: value}}}}

		m, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Data: d, Theme: theme.New(theme.Dark)})
		if err != nil {
			t.Fatal(err)
		}

		m.Resize(layout.Size{Width: 132, Height: 80})

		frame := ansi.Strip(m.Panels.View())
		for _, invalid := range []string{"Invalid interval", "+Inf", "-Inf", "NaN", "0.00C"} {
			if strings.Contains(frame, invalid) {
				t.Fatalf("value %g rendered %q", value, invalid)
			}
		}

		if !strings.Contains(frame, "e+") && value > 1 {
			t.Fatal("large value lost its magnitude")
		}

		if !strings.Contains(ansi.Strip(m.Inspector.Rows()[0][1]), "e") {
			t.Fatal("extreme inspector value lost its magnitude")
		}
	}
}
