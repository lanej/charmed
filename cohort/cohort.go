// Package cohort describes cohort measurements without I/O or UI.
package cohort

import (
	"fmt"
	"slices"
	"time"
)

type (
	Cohort     struct{ ID, Label, Short string }
	MetricKind string
)

const (
	Ratio  MetricKind = ""
	Scalar MetricKind = "scalar"
)

type Aggregation string

const (
	Mean Aggregation = ""
	Sum  Aggregation = "sum"
)

// Metric defines comparable measurements. Ratios aggregate as a ratio of sums
// and display as percentages; scalars aggregate as an unweighted mean or sum.
// Member spread always uses unweighted member values, independently of aggregation.
type Metric struct {
	ID, Label   string
	Kind        MetricKind
	Unit        string
	Aggregation Aggregation
}

func (m Metric) Scale() float64 {
	if m.Kind == Scalar {
		return 1
	}

	return 100
}

func (m Metric) DisplayUnit() string {
	if m.Kind == Scalar {
		return m.Unit
	}

	return "%"
}

func (m Metric) DeltaUnit() string {
	if m.Kind == Scalar {
		return m.Unit
	}

	return "pp"
}

// Measurement is either a scalar Value or a ratio Numerator/Denominator,
// according to its metric. A zero denominator is an unavailable ratio, not zero.
// Omit a metric from Measurements when no observation exists.
type Measurement struct {
	Value                  float64
	Numerator, Denominator uint64
}
type MemberKey struct{ CohortID, MemberID string }

// Sample contains member observations for one window. Measurements supports
// sparse scalar and ratio values; legacy Counts/Requests requires an explicit
// interval count for every metric. Cumulative counters must be converted by callers.
type Sample struct {
	MemberKey
	Requests uint64
	Counts   map[string]uint64
	// Measurements replaces Counts/Requests for generic inputs. Do not mix the
	// two representations in one sample. An empty non-nil map means all missing.
	Measurements map[string]Measurement
}
type Window struct {
	Start, End time.Time
	Samples    []Sample
}
type Data struct {
	Cohorts []Cohort
	Metrics []Metric
	Windows []Window
}

// Clone validates the complete dataset and takes ownership of an independent copy.
// Windows must be ordered, nonoverlapping intervals of equal duration. A member
// can occur once per cohort per window; absence never becomes a zero-rate sample.
func (d Data) Clone() (Data, error) {
	copy := Data{Cohorts: slices.Clone(d.Cohorts), Metrics: slices.Clone(d.Metrics)}
	cohorts, metrics := map[string]bool{}, map[string]bool{}
	for _, c := range d.Cohorts {
		if c.ID == "" || c.Label == "" || c.Short == "" || cohorts[c.ID] {
			return Data{}, fmt.Errorf("cohorts need unique IDs, labels and short labels")
		}
		cohorts[c.ID] = true
	}
	for _, m := range d.Metrics {
		if m.ID == "" || m.Label == "" || metrics[m.ID] {
			return Data{}, fmt.Errorf("metrics need unique IDs and labels")
		}

		if err := m.validate(); err != nil {
			return Data{}, err
		}
		metrics[m.ID] = true
	}
	if len(cohorts) == 0 || len(metrics) == 0 {
		return Data{}, fmt.Errorf("at least one cohort and metric are required")
	}
	var duration time.Duration
	for i, w := range d.Windows {
		if w.Start.IsZero() || !w.End.After(w.Start) || (i > 0 && w.Start.Before(d.Windows[i-1].End)) {
			return Data{}, fmt.Errorf("window %d: expected an ordered, nonoverlapping interval", i)
		}
		if i == 0 {
			duration = w.End.Sub(w.Start)
		} else if w.End.Sub(w.Start) != duration {
			return Data{}, fmt.Errorf("window %d: duration differs from the first window", i)
		}
		seen := map[MemberKey]bool{}
		wc := Window{Start: w.Start, End: w.End}
		for _, s := range w.Samples {
			if !cohorts[s.CohortID] || s.MemberID == "" || seen[s.MemberKey] {
				return Data{}, fmt.Errorf("window %d: unknown cohort, empty member or duplicate member", i)
			}
			seen[s.MemberKey] = true

			sc, err := cloneSample(s, d.Metrics)
			if err != nil {
				return Data{}, fmt.Errorf("window %d member %s: %w", i, s.MemberID, err)
			}
			wc.Samples = append(wc.Samples, sc)
		}
		copy.Windows = append(copy.Windows, wc)
	}
	return copy, nil
}

type MemberRate struct {
	MemberKey
	Requests, Events  uint64
	Rate              float64
	Value             float64
	HasValue          bool
	HasRate, Eligible bool
}
type Summary struct {
	// Requests and Events are floating sums to avoid overflowing integer counters.
	Requests, Events            float64
	Rate, Mean, StdDev          float64
	Value                       float64
	HasValue                    bool
	HasRate, HasSpread          bool
	Members, Eligible, Excluded int
	Points                      []MemberRate
}

// Analyze separates request-weighted impact from unweighted member spread.
// Rates are fractions [0,1]. Spread is the sample standard deviation (n-1) of
// member rates with at least minRequests. Low-volume members still contribute
// to fleet Rate, but not Mean or StdDev. Two eligible members are needed for s.
// Use validated Data; an unknown/missing metric is omitted rather than read as 0.
func Analyze(samples []Sample, cohortID, metricID string, minRequests uint64) Summary {
	return AnalyzeMetric(samples, cohortID, Metric{ID: metricID}, minRequests)
}

// AnalyzeMetric retains missingness, aggregate value and unweighted sample
// spread separately. minSupport applies to ratio denominators only; every
// available scalar is eligible. Legacy Rate/HasRate mirror Value/HasValue.
func AnalyzeMetric(samples []Sample, cohortID string, metric Metric, minSupport uint64) Summary {
	var result Summary

	for _, s := range samples {
		if s.CohortID != cohortID {
			continue
		}

		measurement, exists := measurementFor(s, metric)
		if !exists {
			continue
		}

		point := MemberRate{MemberKey: s.MemberKey, Requests: measurement.Denominator, Events: measurement.Numerator}
		result.Members++
		result.Requests += float64(measurement.Denominator)

		result.Events += float64(measurement.Numerator)
		if metric.Kind == Scalar {
			point.Value, point.HasValue, point.Eligible = measurement.Value, true, true
		} else if measurement.Denominator > 0 {
			point.Value = float64(measurement.Numerator) / float64(measurement.Denominator)
			point.HasValue = true
			point.Eligible = measurement.Denominator >= max(1, minSupport)
		}

		point.Rate, point.HasRate = point.Value, point.HasValue
		if point.Eligible {
			result.Eligible++
		} else {
			result.Excluded++
		}

		result.Points = append(result.Points, point)
	}

	total := result.memberMoments()
	if metric.Kind == Scalar {
		result.HasValue = result.Eligible > 0

		result.Value = result.Mean
		if metric.Aggregation == Sum {
			result.Value = total
		}
	} else if result.Requests > 0 {
		result.HasValue = true
		result.Value = result.Events / result.Requests
	}

	if !finite(result.Value) {
		result.Value, result.HasValue = 0, false
	}

	result.Rate, result.HasRate = result.Value, result.HasValue
	return result
}

// Distance is a descriptive distance from a reference member distribution, not
// a hypothesis test. Zero or unavailable spread cannot produce a finite score.
func Distance(p MemberRate, reference Summary) (float64, bool) {
	if !p.Eligible || !p.HasRate || !reference.HasSpread || reference.StdDev <= 0 || !finite(p.Rate) || !finite(reference.Mean) || !finite(reference.StdDev) {
		return 0, false
	}

	distance := (p.Rate - reference.Mean) / reference.StdDev
	if !finite(distance) {
		return 0, false
	}

	return distance, true
}
