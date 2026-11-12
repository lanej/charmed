package cohort

import (
	"fmt"
	"maps"
	"math"
)

func (m Metric) validate() error {
	switch m.Kind {
	case Ratio:
		if m.Unit != "" {
			return fmt.Errorf("ratio %s uses percentage units", m.ID)
		}
	case Scalar:
	default:
		return fmt.Errorf("unknown metric kind %q", m.Kind)
	}

	switch m.Aggregation {
	case Mean:
	case Sum:
		if m.Kind != Scalar {
			return fmt.Errorf("sum requires a scalar metric")
		}
	default:
		return fmt.Errorf("unknown aggregation %q", m.Aggregation)
	}

	return nil
}

func cloneSample(s Sample, metrics []Metric) (Sample, error) {
	result := s
	if s.Measurements != nil {
		if s.Counts != nil || s.Requests != 0 {
			return Sample{}, fmt.Errorf("cannot mix measurements with counts/requests")
		}

		known := make(map[string]bool, len(metrics))
		for _, m := range metrics {
			known[m.ID] = true

			v, ok := s.Measurements[m.ID]
			if !ok {
				continue
			}

			if math.IsNaN(v.Value) || math.IsInf(v.Value, 0) {
				return Sample{}, fmt.Errorf("nonfinite measurement %s", m.ID)
			}

			if m.Kind == Scalar && (v.Numerator != 0 || v.Denominator != 0) {
				return Sample{}, fmt.Errorf("scalar %s contains ratio fields", m.ID)
			}

			if m.Kind == Ratio && (v.Value != 0 || (v.Denominator == 0 && v.Numerator != 0)) {
				return Sample{}, fmt.Errorf("invalid ratio %s", m.ID)
			}
		}

		for id := range s.Measurements {
			if !known[id] {
				return Sample{}, fmt.Errorf("unknown metric %s", id)
			}
		}

		result.Measurements = maps.Clone(s.Measurements)

		return result, nil
	}

	if len(s.Counts) != len(metrics) {
		return Sample{}, fmt.Errorf("every metric needs an explicit count")
	}

	for _, m := range metrics {
		n, ok := s.Counts[m.ID]
		if m.Kind != Ratio || !ok || n > s.Requests {
			return Sample{}, fmt.Errorf("invalid or missing count for %s", m.ID)
		}
	}

	result.Counts = maps.Clone(s.Counts)

	return result, nil
}

func measurementFor(s Sample, m Metric) (Measurement, bool) {
	if s.Measurements != nil {
		v, ok := s.Measurements[m.ID]
		return v, ok
	}

	n, ok := s.Counts[m.ID]

	return Measurement{Numerator: n, Denominator: s.Requests}, ok && m.Kind == Ratio && n <= s.Requests
}
