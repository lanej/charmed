package chart

import "math"

// PaddedRange returns finite, distinct bounds containing both observations.
// Magnitude-aware padding survives rounding for large constants; an unrepresentable
// span is unavailable rather than replaced with a misleading default axis.
func PaddedRange(low, high, fraction float64) (float64, float64, bool) {
	if !finite(low) || !finite(high) || !finite(fraction) || fraction <= 0 || low > high || !finite(high-low) {
		return 0, 0, false
	}

	originalLow, originalHigh := low, high
	magnitude := max(math.Abs(low), math.Abs(high))

	padding := (high - low) * fraction
	if high == low {
		padding = magnitude * fraction
	}

	if !finite(padding) {
		return 0, 0, false
	}

	if padding == 0 {
		padding = math.SmallestNonzeroFloat64
		if magnitude == 0 {
			padding = 1
		}
	}

	low = max(-math.MaxFloat64, low-padding)

	high = min(math.MaxFloat64, high+padding)
	if !finite(high - low) {
		low, high = originalLow, originalHigh
	}

	return low, high, finite(low) && finite(high) && high > low && finite(high-low)
}
