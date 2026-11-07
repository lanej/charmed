package cohort_test

import (
	"math"
	"testing"
	"time"

	"github.com/lanej/charmed/cohort"
)

func near(t *testing.T, got, want float64) {
	t.Helper()
	if math.Abs(got-want) > 1e-12 {
		t.Fatalf("got %.15g, want %.15g", got, want)
	}
}
func sample(id string, requests, events uint64) cohort.Sample {
	return cohort.Sample{MemberKey: cohort.MemberKey{CohortID: "a", MemberID: id}, Requests: requests, Counts: map[string]uint64{"errors": events}}
}
func TestWeightedImpactAndMemberSpreadAreDifferent(t *testing.T) {
	samples := []cohort.Sample{sample("quiet", 100, 10), sample("busy", 1000, 900), sample("sparse", 10, 10), sample("idle", 0, 0)}
	s := cohort.Analyze(samples, "a", "errors", 100)
	near(t, s.Rate, 920.0/1110)
	near(t, s.Mean, 0.5)
	near(t, s.StdDev, math.Sqrt(0.32))
	if s.Requests != 1110 || s.Events != 920 || s.Members != 4 || s.Eligible != 2 || s.Excluded != 2 || !s.HasSpread || !s.HasRate {
		t.Fatalf("incorrect populations: %+v", s)
	}
	if s.Points[3].HasRate || s.Points[2].Eligible {
		t.Fatal("idle or low-volume host became an eligible observation")
	}
	z, ok := cohort.Distance(s.Points[1], s)
	if !ok {
		t.Fatal("eligible member has no descriptive score")
	}
	near(t, z, 0.4/math.Sqrt(0.32))
	if _, ok := cohort.Distance(s.Points[2], s); ok {
		t.Fatal("low-volume member received a score")
	}
}
func TestUnavailableAndZeroSpreadDoNotImplyConfidence(t *testing.T) {
	for _, samples := range [][]cohort.Sample{nil, {sample("a", 0, 0)}, {sample("a", 100, 0)}, {sample("a", 100, 0), sample("b", 100, 0)}} {
		s := cohort.Analyze(samples, "a", "errors", 10)
		if len(samples) < 2 && s.HasSpread {
			t.Fatal("fewer than two members yielded spread")
		}
		if _, ok := cohort.Distance(cohort.MemberRate{Rate: 0.1, HasRate: true, Eligible: true}, s); ok {
			t.Fatal("zero/unavailable spread yielded a finite score")
		}
	}
	if s := cohort.Analyze([]cohort.Sample{sample("a", 100, 10)}, "a", "missing", 1); s.HasRate || s.Members != 0 {
		t.Fatal("missing count became a measured zero")
	}
}
func data() cohort.Data {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	return cohort.Data{Cohorts: []cohort.Cohort{{ID: "a", Label: "A", Short: "A"}}, Metrics: []cohort.Metric{{ID: "errors", Label: "Errors"}},
		Windows: []cohort.Window{{Start: start, End: start.Add(time.Minute), Samples: []cohort.Sample{sample("one", 100, 10)}}}}
}
func TestDataCopiesAndRejectsAmbiguousObservations(t *testing.T) {
	original := data()
	copy, err := original.Clone()
	if err != nil {
		t.Fatal(err)
	}
	copy.Windows[0].Samples[0].Counts["errors"] = 99
	copy.Cohorts[0].Label = "Changed"
	if original.Windows[0].Samples[0].Counts["errors"] != 10 || original.Cohorts[0].Label != "A" {
		t.Fatal("data clone aliases input")
	}
	for name, change := range map[string]func(*cohort.Data){
		"duplicate member":      func(d *cohort.Data) { d.Windows[0].Samples = append(d.Windows[0].Samples, d.Windows[0].Samples[0]) },
		"missing count":         func(d *cohort.Data) { delete(d.Windows[0].Samples[0].Counts, "errors") },
		"wrong count key":       func(d *cohort.Data) { d.Windows[0].Samples[0].Counts = map[string]uint64{"other": 10} },
		"counter exceeds total": func(d *cohort.Data) { d.Windows[0].Samples[0].Counts["errors"] = 101 },
		"unknown cohort":        func(d *cohort.Data) { d.Windows[0].Samples[0].CohortID = "unknown" },
		"overlapping window":    func(d *cohort.Data) { d.Windows = append(d.Windows, d.Windows[0]) },
		"unequal duration": func(d *cohort.Data) {
			w := d.Windows[0]
			w.Start = w.End
			w.End = w.Start.Add(2 * time.Minute)
			d.Windows = append(d.Windows, w)
		},
	} {
		t.Run(name, func(t *testing.T) {
			d := data()
			change(&d)
			if _, err := d.Clone(); err == nil {
				t.Fatal("invalid data accepted")
			}
		})
	}
}
