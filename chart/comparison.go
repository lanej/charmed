package chart

import (
	"fmt"
	"math"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
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
	Compact    bool
	ShowLegend bool
}

func (d DistributionComparison) View(width int, t theme.Theme) string { return d.view(width, t, -1) }

func (d DistributionComparison) compactPrefix(width int, t theme.Theme) int {
	prefix := min(29, max(8, width/3))

	for _, series := range d.Series {
		value := text.AbbreviatedNumber(series.Mean, 3, d.Unit, true)
		if prefix >= width || prefix < lipgloss.Width(value)+2 {
			return 0
		}
	}

	return prefix
}

func (d DistributionComparison) view(width int, t theme.Theme, sharedPrefix int) string {
	if width <= 0 {
		return ""
	}

	if !finite(d.Min) || !finite(d.Max) || !finite(d.Max-d.Min) || d.Min >= d.Max {
		return t.Subtle().Render(ansi.Wrap("Distribution range unavailable", width, ""))
	}

	if len(d.Series) == 0 {
		return t.Subtle().Render(ansi.Wrap("No observations", width, ""))
	}

	low, high := d.domain()

	var (
		lines  []string
		legend []legendEntry
	)

	prefixWidth := 0
	if d.Compact {
		prefixWidth = sharedPrefix
		if prefixWidth < 0 {
			prefixWidth = d.compactPrefix(width, t)
		}
	}

	for i, s := range d.Series {
		style := t.Text().Foreground(t.SeriesColor(i))
		if s.Members == 0 {
			if d.Compact {
				lines = append(lines, style.Render(ansi.Wrap(s.Label+" n/a", width, "")))
			} else {
				lines = append(lines, style.Render(ansi.Wrap(s.Label+" member mean n/a", width, "")), t.Subtle().Render(ansi.Wrap("Insufficient member samples", width, "")))
			}
			continue
		}
		spread := "n/a"
		if s.HasSpread {
			spread = text.AbbreviatedNumber(s.StdDev, 3, d.SpreadUnit, true)
		}
		if !d.Compact {
			label := fmt.Sprintf("%s mean %s s=%s n=%d", s.Label, text.AbbreviatedNumber(s.Mean, 3, d.Unit, true), spread, s.Members)
			lines = append(lines, style.Render(ansi.Wrap(label, width, "")))
		}

		opts := distributionOptions(d, s, i, low, high)
		prefix := ""
		if d.Compact {
			value := text.AbbreviatedNumber(s.Mean, 3, d.Unit, true)
			if prefixWidth == 0 {
				lines = append(lines, style.Render(ansi.Wrap(s.Label+" "+value, width, "")))
			} else {
				labelWidth := prefixWidth - lipgloss.Width(value) - 1
				label := ansi.Truncate(s.Label, labelWidth, "~")
				prefix = style.Render(layout.Fit(label, layout.Size{Width: labelWidth, Height: 1}) + " " + value)
			}
		}

		lines = append(lines, prefix+IntervalBand(opts, max(1, width-lipgloss.Width(prefix)), t))
		if d.ShowLegend {
			legend = append(legend, opts.legendEntries(s.Label+" ", t)...)
		}
	}
	if math.IsInf(low, 0) {
		lines = append(lines, t.Subtle().Render(ansi.Wrap("No eligible member spread", width, "")))
	} else {
		left, right := text.AbbreviatedNumber(low, 3, d.Unit, true), text.AbbreviatedNumber(high, 3, d.Unit, true)
		axisWidth := width - prefixWidth
		labelWidth := lipgloss.Width(left) + lipgloss.Width(right)

		if labelWidth+1 <= axisWidth {
			lines = append(lines, strings.Repeat(" ", prefixWidth)+left+strings.Repeat(" ", axisWidth-labelWidth)+right)
		} else {
			lines = append(lines, t.Subtle().Render(ansi.Wrap(left+"  "+right, width, "")))
		}
	}
	if d.ShowLegend {
		if key := renderLegend(legend, width, t); key != "" {
			lines = append(lines, key)
		}
	}
	return strings.Join(lines, "\n")
}

func (d DistributionComparison) domain() (float64, float64) {
	low, high := math.Inf(1), math.Inf(-1)
	for _, s := range d.Series {
		if s.Members > 0 {
			spread := 0.0
			if s.HasSpread {
				spread = 3 * s.StdDev
			}

			low = min(low, max(d.Min, s.Mean-spread))
			high = max(high, min(d.Max, s.Mean+spread))
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

	return low, high
}

func distributionOptions(d DistributionComparison, s DistributionSeries, index int, low, high float64) BandOptions {
	options := BandOptions{Min: low, Max: high, Center: s.Mean, Selected: s.Selected, Series: index, Markers: d.Markers}
	if s.HasSpread {
		for k := max(1, min(3, d.Sigma)); k >= 1; k-- {
			options.Intervals = append(options.Intervals, Interval{Low: max(low, s.Mean-float64(k)*s.StdDev), High: min(high, s.Mean+float64(k)*s.StdDev)})
		}
	}

	return options
}

// ComparisonPanel composes labels, distribution and native history as one region.
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
		lines = append(lines, p.Distribution.View(max(1, layout.PanelContentSize(size).Width), t))
	}

	lines = append(lines, p.History.View())

	return layout.Panel(p.Title, strings.Join(lines, "\n"), size, t, focused)
}
