package cohort

import "math"

func finite(value float64) bool { return !math.IsNaN(value) && !math.IsInf(value, 0) }

// Scaling before summing and squaring avoids overflow in intermediate member
// statistics. Truly unrepresentable sums or spreads are explicitly unavailable.
func (s *Summary) memberMoments() float64 {
	scale := 0.0

	for _, p := range s.Points {
		if p.Eligible {
			scale = max(scale, math.Abs(p.Value))
		}
	}

	if s.Eligible == 0 {
		return 0
	}

	if scale == 0 {
		s.HasSpread = s.Eligible >= 2
		return 0
	}

	sum := 0.0

	for _, p := range s.Points {
		if p.Eligible {
			sum += p.Value / scale
		}
	}

	mean := sum / float64(s.Eligible)

	s.Mean = mean * scale
	if s.Eligible >= 2 {
		squares := 0.0

		for _, p := range s.Points {
			if p.Eligible {
				delta := p.Value/scale - mean
				squares += delta * delta
			}
		}

		spread := math.Sqrt(squares/float64(s.Eligible-1)) * scale
		if finite(spread) {
			s.StdDev, s.HasSpread = spread, true
		}
	}

	return sum * scale
}
