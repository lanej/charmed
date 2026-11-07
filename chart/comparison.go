package chart

import (
	"fmt"
	"math"
	"strings"

	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

type DistributionSeries struct {
	Label        string
	Mean, StdDev float64
	Members      int
	HasSpread    bool
	Selected     *float64
}

// DistributionComparison composes mean/spread labels, interval atoms, shared
// caller-supplied boundary markers and a common axis. Values use the caller's unit.
type DistributionComparison struct {
	Series           []DistributionSeries
	Sigma            int
	Min, Max         float64
	Unit, SpreadUnit string
	Markers          []BandMarker
}

func (d DistributionComparison) View(width int, t theme.Theme) string {
	low, high := math.Inf(1), math.Inf(-1)
	for _, s := range d.Series {
		if s.Members > 0 {
			spread := 0.0
			if s.HasSpread {
				spread = 3 * s.StdDev
			}
			low = min(low, s.Mean-spread)
			high = max(high, s.Mean+spread)
		}
		if s.Selected != nil && finite(*s.Selected) {
			low = min(low, *s.Selected)
			high = max(high, *s.Selected)
		}
	}
	for _, m := range d.Markers {
		if finite(m.Value) && m.Value >= d.Min && m.Value <= d.Max {
			low = min(low, m.Value)
			high = max(high, m.Value)
		}
	}
	if !math.IsInf(low, 0) {
		pad := max((high-low)*0.07, (d.Max-d.Min)*0.00001)
		low = max(d.Min, low-pad)
		high = min(d.Max, high+pad)
	}
	var lines []string
	for i, s := range d.Series {
		style := t.Text().Foreground(t.SeriesColor(i))
		if s.Members == 0 {
			lines = append(lines, style.Render(s.Label+" host mean n/a"), "Insufficient member samples")
			continue
		}
		spread := "n/a"
		if s.HasSpread {
			spread = fmt.Sprintf("%.3f%s", s.StdDev, d.SpreadUnit)
		}
		lines = append(lines, style.Render(fmt.Sprintf("%s mean %.3f%s s=%s n=%d", s.Label, s.Mean, d.Unit, spread, s.Members)))
		opts := BandOptions{Min: low, Max: high, Center: s.Mean, Selected: s.Selected, Series: i, Markers: d.Markers}
		if s.HasSpread {
			for k := max(1, min(3, d.Sigma)); k >= 1; k-- {
				opts.Intervals = append(opts.Intervals, Interval{Low: s.Mean - float64(k)*s.StdDev, High: s.Mean + float64(k)*s.StdDev})
			}
		}
		lines = append(lines, IntervalBand(opts, width, t))
	}
	if math.IsInf(low, 0) {
		lines = append(lines, "No eligible host spread")
	} else {
		left, right := fmt.Sprintf("%.3f%s", low, d.Unit), fmt.Sprintf("%.3f%s", high, d.Unit)
		lines = append(lines, left+strings.Repeat(" ", max(1, width-len(left)-len(right)))+right)
	}
	return strings.Join(lines, "\n")
}

// ComparisonPanel is a pure molecule: labels + distribution + native history.
// It accepts already-computed values and never evaluates an attention policy.
type ComparisonPanel struct {
	Title        string
	Summary      []string
	Distribution DistributionComparison
	History      TimeSeries
	HistoryLabel string
}

func (p ComparisonPanel) View(size layout.Size, t theme.Theme, focused bool) string {
	lines := append([]string(nil), p.Summary...)
	lines = append(lines, p.Distribution.View(max(1, size.Width-2), t))
	lines = append(lines, p.HistoryLabel, p.History.View())
	return layout.Panel(p.Title, strings.Join(lines, "\n"), size, t, focused)
}
