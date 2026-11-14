package chart

import (
	"math"
	"slices"

	"github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart"
)

// ApplyData replaces changed observations and reports whether preparation ran.
// Equal data leaves native customization intact. Use SetData for an explicit
// full refresh, including reapplying constructor options and theme defaults.
func (m *TimeSeries) ApplyData(series []TimeSeriesData) bool {
	equal := slices.EqualFunc(m.series, series, func(a, b TimeSeriesData) bool {
		return a.ID == b.ID && a.Label == b.Label && a.Unit == b.Unit && slices.EqualFunc(a.Points, b.Points, func(x, y timeserieslinechart.TimePoint) bool {
			return x.Time.Equal(y.Time) && math.Float64bits(x.Value) == math.Float64bits(y.Value)
		})
	})
	if equal {
		return false
	}

	m.SetData(series)

	return true
}

// ApplyData replaces changed observations. SetData remains the full-refresh path.
func (m *Sparkline) ApplyData(values []float64) bool {
	if slices.EqualFunc(m.values, values, func(a, b float64) bool { return math.Float64bits(a) == math.Float64bits(b) }) {
		return false
	}

	m.SetData(values)

	return true
}
