package pattern_test

import (
	"strings"
	"testing"
	"time"

	"github.com/lanej/charmed/cohort"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

func TestMeasurementHistoryPreservesEvidenceAndCoordinates(t *testing.T) {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	member := cohort.MemberKey{CohortID: "c", MemberID: "m"}
	data := cohort.Data{Cohorts: []cohort.Cohort{{ID: "c", Label: "C", Short: "C"}}, Metrics: []cohort.Metric{{ID: "scalar", Label: "Scalar", Kind: cohort.Scalar}, {ID: "ratio", Label: "Ratio"}}}

	for i := range 3 {
		w := cohort.Window{Start: start.Add(time.Duration(i) * time.Hour), End: start.Add(time.Duration(i+1) * time.Hour)}
		if i > 0 {
			w.Samples = []cohort.Sample{{MemberKey: member, Measurements: map[string]cohort.Measurement{"scalar": {Value: 0}, "ratio": {Numerator: 0, Denominator: 0}}}}
		}

		if i == 1 {
			delete(w.Samples[0].Measurements, "scalar")
		}

		data.Windows = append(data.Windows, w)
	}

	m, err := pattern.NewMeasurementHistory(data, member, layout.Size{Width: 80, Height: 14}, theme.New(theme.Mono))
	if err != nil {
		t.Fatal(err)
	}

	selected, ok := m.Selected()
	if !ok || selected.State != pattern.MeasurementObserved || selected.Measurement.Value != 0 || !selected.Start.Equal(data.Windows[2].Start) {
		t.Fatal("scalar zero or latest initial selection lost")
	}

	check := func(row string, i int, state pattern.MeasurementState) {
		t.Helper()

		if !m.Inspector.Select(row, data.Windows[i].Start.Format(time.RFC3339Nano)) {
			t.Fatal("coordinate missing")
		}

		v, ok := m.Selected()
		if !ok || v.State != state || !v.End.Equal(data.Windows[i].End) {
			t.Fatalf("wrong evidence %+v", v)
		}
	}
	check("scalar", 0, pattern.MeasurementMemberAbsent)
	check("scalar", 1, pattern.MeasurementMetricAbsent)

	if strings.Contains(m.Inspector.Details.Model.View(), "Raw value:") {
		t.Fatal("missing metric displayed fabricated raw zero")
	}

	check("ratio", 2, pattern.MeasurementRatioUnavailable)
	check("ratio", 1, pattern.MeasurementRatioUnavailable)

	if err := m.SetData(data); err != nil {
		t.Fatal(err)
	}

	if v, _ := m.Selected(); !v.Start.Equal(data.Windows[1].Start) {
		t.Fatal("retained window changed")
	}

	m.Inspector.Grid.RenderCell = func(c collection.GridCell[pattern.HistoryMeasurement]) string { return c.Entry.RowID }
	m.SetMember(cohort.MemberKey{CohortID: "c", MemberID: "absent"})

	if v, _ := m.Selected(); v.State != pattern.MeasurementMemberAbsent {
		t.Fatal("member replacement used old observations")
	}

	shortened := data

	shortened.Windows = []cohort.Window{data.Windows[2]}
	if err := m.SetData(shortened); err != nil {
		t.Fatal(err)
	}

	if v, _ := m.Selected(); !v.Start.Equal(data.Windows[2].Start) {
		t.Fatal("removed window did not select latest")
	}
}
