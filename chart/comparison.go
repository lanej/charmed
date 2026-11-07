package chart

import (
	"fmt"
	"math"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
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
	// Compact puts a direct series label and mean beside each band.
	Compact bool
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
	prefixWidth := 0
	if d.Compact {
		prefixWidth = min(29, max(8, width/3))
	}
	for i, s := range d.Series {
		style := t.Text().Foreground(t.SeriesColor(i))
		if s.Members == 0 {
			if d.Compact {
				lines = append(lines, style.Render(ansi.Truncate(s.Label, max(1, prefixWidth-5), "~")+" n/a"))
			} else {
				lines = append(lines, style.Render(s.Label+" host mean n/a"), "Insufficient member samples")
			}
			continue
		}
		spread := "n/a"
		if s.HasSpread {
			spread = fmt.Sprintf("%.3f%s", s.StdDev, d.SpreadUnit)
		}
		if !d.Compact {
			lines = append(lines, style.Render(fmt.Sprintf("%s mean %.3f%s s=%s n=%d", s.Label, s.Mean, d.Unit, spread, s.Members)))
		}
		opts := BandOptions{Min: low, Max: high, Center: s.Mean, Selected: s.Selected, Series: i, Markers: d.Markers}
		if s.HasSpread {
			for k := max(1, min(3, d.Sigma)); k >= 1; k-- {
				opts.Intervals = append(opts.Intervals, Interval{Low: s.Mean - float64(k)*s.StdDev, High: s.Mean + float64(k)*s.StdDev})
			}
		}
		prefix := ""
		if d.Compact {
			value := fmt.Sprintf("%.3f%s", s.Mean, d.Unit)
			labelWidth := max(1, prefixWidth-len(value)-2)
			label := ansi.Truncate(s.Label, labelWidth, "~")
			prefix = style.Render(layout.Fit(label+strings.Repeat(" ", max(1, labelWidth-lipgloss.Width(label)+1))+value,
				layout.Size{Width: prefixWidth, Height: 1}))
		}
		lines = append(lines, prefix+IntervalBand(opts, max(1, width-prefixWidth), t))
	}
	if math.IsInf(low, 0) {
		lines = append(lines, "No eligible host spread")
	} else {
		left, right := fmt.Sprintf("%.3f%s", low, d.Unit), fmt.Sprintf("%.3f%s", high, d.Unit)
		lines = append(lines, strings.Repeat(" ", prefixWidth)+left+strings.Repeat(" ", max(1, width-prefixWidth-len(left)-len(right)))+right)
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
}

func (p ComparisonPanel) View(size layout.Size, t theme.Theme, focused bool) string {
	lines := append([]string(nil), p.Summary...)
	if len(p.Distribution.Series) > 0 {
		lines = append(lines, p.Distribution.View(max(1, size.Width-2), t))
	}
	lines = append(lines, p.History.View())
	return layout.Panel(p.Title, strings.Join(lines, "\n"), size, t, focused)
}

type BandRow struct {
	Title        string
	Distribution DistributionComparison
}

// BandMatrix composes disconnected, directly labeled comparisons in one frame.
// Each metric owns a numeric axis shared by its versions, never a connecting
// profile across categories. The rows preserve version colors and interval atoms.
type BandMatrix struct {
	Title string
	Rows  []BandRow
}

func (b BandMatrix) View(size layout.Size, t theme.Theme, focused bool) string {
	var rows []string
	for _, row := range b.Rows {
		d := row.Distribution
		d.Compact = true
		rows = append(rows, row.Title+"\n"+d.View(max(1, size.Width-2), t))
	}
	if len(rows) == 0 {
		rows = []string{"No observations"}
	}
	return layout.Panel(b.Title, strings.Join(rows, "\n\n"), size, t, focused)
}
