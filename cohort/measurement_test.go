package cohort_test

import (
	"math"
	"testing"
	"time"

	"github.com/lanej/charmed/cohort"
)

func TestGenericMeasurementsPreserveAggregationAndMissingness(t *testing.T) {
	samples := []cohort.Sample{
		{MemberKey: cohort.MemberKey{CohortID: "a", MemberID: "one"}, Measurements: map[string]cohort.Measurement{"temp": {Value: -10}, "yield": {Numerator: 9, Denominator: 10}}},
		{MemberKey: cohort.MemberKey{CohortID: "a", MemberID: "two"}, Measurements: map[string]cohort.Measurement{"temp": {Value: 0}, "yield": {Numerator: 10, Denominator: 100}}},
		{MemberKey: cohort.MemberKey{CohortID: "a", MemberID: "missing"}, Measurements: map[string]cohort.Measurement{}},
	}
	scalar := cohort.Metric{ID: "temp", Kind: cohort.Scalar, Unit: "C"}

	s := cohort.AnalyzeMetric(samples, "a", scalar, 1000)
	if s.Value != -5 || s.Mean != -5 || s.Members != 2 || s.Eligible != 2 || !s.HasSpread || math.Abs(s.StdDev-math.Sqrt(50)) > 1e-12 {
		t.Fatalf("scalar mean/spread or explicit zero lost: %+v", s)
	}

	scalar.Aggregation = cohort.Sum

	s = cohort.AnalyzeMetric(samples, "a", scalar, 1000)
	if s.Value != -10 || s.Mean != -5 {
		t.Fatalf("sum changed member mean: %+v", s)
	}

	r := cohort.AnalyzeMetric(samples, "a", cohort.Metric{ID: "yield"}, 10)
	if math.Abs(r.Value-19.0/110) > 1e-12 || r.Mean != 0.5 || r.Members != 2 {
		t.Fatalf("ratio was averaged or missing became zero: %+v", r)
	}

	r = cohort.AnalyzeMetric(samples, "a", cohort.Metric{ID: "yield"}, 100)
	if r.Value != 19.0/110 || r.Eligible != 1 || r.HasSpread {
		t.Fatalf("support threshold changed aggregate: %+v", r)
	}
}

func genericData() cohort.Data {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	return cohort.Data{Cohorts: []cohort.Cohort{{ID: "a", Label: "A", Short: "A"}}, Metrics: []cohort.Metric{{ID: "temp", Label: "Temperature", Kind: cohort.Scalar, Unit: "C"}, {ID: "yield", Label: "Yield"}}, Windows: []cohort.Window{{Start: start, End: start.Add(time.Minute), Samples: []cohort.Sample{{MemberKey: cohort.MemberKey{CohortID: "a", MemberID: "one"}, Measurements: map[string]cohort.Measurement{"temp": {Value: -10}, "yield": {Numerator: 3, Denominator: 2}}}}}}}
}

func TestGenericValidationAndOwnership(t *testing.T) {
	original := genericData()

	cloned, err := original.Clone()
	if err != nil {
		t.Fatal(err)
	}

	cloned.Windows[0].Samples[0].Measurements["temp"] = cohort.Measurement{Value: 20}
	if original.Windows[0].Samples[0].Measurements["temp"].Value != -10 {
		t.Fatal("clone aliases input")
	}

	for name, mutate := range map[string]func(*cohort.Data){
		"nonfinite": func(d *cohort.Data) {
			d.Windows[0].Samples[0].Measurements["temp"] = cohort.Measurement{Value: math.Inf(1)}
		},
		"nan": func(d *cohort.Data) {
			d.Windows[0].Samples[0].Measurements["temp"] = cohort.Measurement{Value: math.NaN()}
		},
		"unknown metric": func(d *cohort.Data) { d.Windows[0].Samples[0].Measurements["other"] = cohort.Measurement{} },
		"mixed inputs":   func(d *cohort.Data) { d.Windows[0].Samples[0].Counts = map[string]uint64{} },
		"scalar ratio": func(d *cohort.Data) {
			d.Windows[0].Samples[0].Measurements["temp"] = cohort.Measurement{Denominator: 1}
		},
		"ratio scalar":        func(d *cohort.Data) { d.Windows[0].Samples[0].Measurements["yield"] = cohort.Measurement{Value: 2} },
		"zero denominator":    func(d *cohort.Data) { d.Windows[0].Samples[0].Measurements["yield"] = cohort.Measurement{Numerator: 1} },
		"unknown kind":        func(d *cohort.Data) { d.Metrics[0].Kind = "other" },
		"unknown aggregation": func(d *cohort.Data) { d.Metrics[0].Aggregation = "other" },
		"sum ratio":           func(d *cohort.Data) { d.Metrics[1].Aggregation = cohort.Sum },
	} {
		t.Run(name, func(t *testing.T) {
			d := genericData()
			mutate(&d)

			if _, err := d.Clone(); err == nil {
				t.Fatal("invalid data accepted")
			}
		})
	}

	samples := original.Windows[0].Samples
	samples[0].Measurements["yield"] = cohort.Measurement{}

	r := cohort.AnalyzeMetric(samples, "a", original.Metrics[1], 1)
	if r.HasValue || r.Eligible != 0 || r.Members != 1 {
		t.Fatal("zero denominator invented a value")
	}
}

func TestScalarMomentsAvoidIntermediateOverflow(t *testing.T) {
	samples := []cohort.Sample{
		{MemberKey: cohort.MemberKey{CohortID: "a", MemberID: "one"}, Measurements: map[string]cohort.Measurement{"v": {Value: 1e200}}},
		{MemberKey: cohort.MemberKey{CohortID: "a", MemberID: "two"}, Measurements: map[string]cohort.Measurement{"v": {Value: -1e200}}},
	}
	metric := cohort.Metric{ID: "v", Kind: cohort.Scalar}

	s := cohort.AnalyzeMetric(samples, "a", metric, 1)
	if !s.HasValue || s.Value != 0 || !s.HasSpread || math.Abs(s.StdDev/1e200-math.Sqrt(2)) > 1e-12 {
		t.Fatalf("representable moments overflowed: %+v", s)
	}

	samples[0].Measurements["v"] = cohort.Measurement{Value: math.MaxFloat64}
	samples[1].Measurements["v"] = cohort.Measurement{Value: -math.MaxFloat64}

	s = cohort.AnalyzeMetric(samples, "a", metric, 1)
	if !s.HasValue || s.Mean != 0 || s.HasSpread {
		t.Fatalf("unrepresentable spread exposed: %+v", s)
	}

	samples[1].Measurements["v"] = cohort.Measurement{Value: math.MaxFloat64}
	metric.Aggregation = cohort.Sum

	s = cohort.AnalyzeMetric(samples, "a", metric, 1)
	if s.HasValue || s.HasRate {
		t.Fatal("overflowing sum is available")
	}

	if _, ok := cohort.Distance(cohort.MemberRate{Rate: math.MaxFloat64, HasRate: true, Eligible: true}, cohort.Summary{Mean: -math.MaxFloat64, StdDev: 1, HasSpread: true}); ok {
		t.Fatal("nonfinite distance available")
	}
}
