// Package cohort describes count-based cohort measurements without I/O or UI.
package cohort

import (
	"fmt"
	"math"
	"slices"
	"time"
)

type Cohort struct{ ID, Label, Short string }
type Metric struct{ ID, Label string }
type MemberKey struct{ CohortID, MemberID string }

// Sample contains interval counts, not cumulative counters or precomputed rates.
// Every configured metric must be present, including explicit zero counts.
// Metrics may overlap; the application decides whether they partition requests.
type Sample struct {
	MemberKey
	Requests uint64
	Counts   map[string]uint64
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
			if len(s.Counts) != len(metrics) {
				return Data{}, fmt.Errorf("window %d member %s: every metric needs an explicit count", i, s.MemberID)
			}
			sc := Sample{MemberKey: s.MemberKey, Requests: s.Requests, Counts: map[string]uint64{}}
			for id := range metrics {
				n, ok := s.Counts[id]
				if !ok || n > s.Requests {
					return Data{}, fmt.Errorf("window %d member %s: invalid or missing count for %s", i, s.MemberID, id)
				}
				sc.Counts[id] = n
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
	HasRate, Eligible bool
}
type Summary struct {
	// Requests and Events are floating sums to avoid overflowing integer counters.
	Requests, Events            float64
	Rate, Mean, StdDev          float64
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
	var result Summary
	var m2 float64
	minRequests = max(1, minRequests)
	for _, s := range samples {
		n, ok := s.Counts[metricID]
		if s.CohortID != cohortID || !ok || n > s.Requests {
			continue
		}
		p := MemberRate{MemberKey: s.MemberKey, Requests: s.Requests, Events: n}
		result.Members++
		result.Requests += float64(s.Requests)
		result.Events += float64(n)
		if s.Requests > 0 {
			p.HasRate = true
			p.Rate = float64(n) / float64(s.Requests)
		}
		p.Eligible = s.Requests >= minRequests
		if p.Eligible {
			result.Eligible++
			delta := p.Rate - result.Mean
			result.Mean += delta / float64(result.Eligible)
			m2 += delta * (p.Rate - result.Mean)
		} else {
			result.Excluded++
		}
		result.Points = append(result.Points, p)
	}
	if result.Requests > 0 {
		result.HasRate = true
		result.Rate = result.Events / result.Requests
	}
	if result.Eligible >= 2 {
		result.HasSpread = true
		result.StdDev = math.Sqrt(max(0, m2/float64(result.Eligible-1)))
	}
	return result
}

// Distance is a descriptive distance from a reference member distribution, not
// a hypothesis test. Zero or unavailable spread cannot produce a finite score.
func Distance(p MemberRate, reference Summary) (float64, bool) {
	if !p.Eligible || !p.HasRate || !reference.HasSpread || reference.StdDev <= 0 {
		return 0, false
	}
	return (p.Rate - reference.Mean) / reference.StdDev, true
}
