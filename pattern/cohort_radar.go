package pattern

import (
	"fmt"
	"math"
	"strings"

	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/cohort"
	"github.com/lanej/charmed/layout"
)

// radarOptions keeps normalization anchored to reference hosts, even when the
// candidate distribution widens. Every axis uses unweighted eligible-host means.
// Full 3s extents and all eligible hosts set a stable domain for this window:
// selection, smoothing and the visible sigma limit never rescale the overview.
func (m *CohortComparison) radarOptions() (chart.RadarOptions, float64, []string) {
	o := chart.RadarOptions{Smooth: m.smooth, Reference: 0, SelectedAxis: m.metric}
	for c, group := range m.data.Cohorts {
		o.Series = append(o.Series, chart.RadarSeries{Label: group.ID, Marker: []rune{'R', 'C'}[c]})
	}
	o.Series = append(o.Series, chart.RadarSeries{Label: m.selected.MemberID, Marker: 'X'})
	span := 3.0
	var unavailable []string
	for i, metric := range m.data.Metrics {
		o.Axes = append(o.Axes, metric.ID)
		ref := m.summaries[i][0]
		valid := ref.HasSpread && ref.StdDev > 0
		if !valid {
			unavailable = append(unavailable, metric.ID)
		}
		for c, s := range m.summaries[i] {
			p := chart.RadarValue{Valid: valid && s.Eligible > 0}
			if p.Valid {
				p.Value = (s.Mean - ref.Mean) / ref.StdDev
				if s.HasSpread {
					for k := 3; k >= 1; k-- {
						interval := normalizedInterval(s, ref, k)
						span = max(span, math.Abs(interval.Low), math.Abs(interval.High))
						if k <= m.sigma {
							p.Intervals = append(p.Intervals, interval)
						}
					}
				}
				span = max(span, math.Abs(p.Value))
			}
			o.Series[c].Values = append(o.Series[c].Values, p)
		}
		host := chart.RadarValue{}
		for _, s := range m.summaries[i] {
			for _, p := range s.Points {
				if z, ok := cohort.Distance(p, ref); ok {
					span = max(span, math.Abs(z))
					if p.MemberKey == m.selected {
						host = chart.RadarValue{Valid: true, Value: z}
					}
				}
			}
		}
		o.Series[2].Values = append(o.Series[2].Values, host)
	}
	span = math.Ceil(span)
	o.Min, o.Max = -span*1.05, span*1.05
	if span <= 9 {
		o.Rings = []float64{-3, -2, -1, 1, 2, 3}
	} else {
		o.Rings = []float64{-span, -span / 2, span / 2, span}
	}
	return o, span, unavailable
}

func normalizedInterval(s, ref cohort.Summary, k int) chart.Interval {
	// Rates are physically bounded, before conversion to reference SD units.
	return chart.Interval{
		Low:  (max(0, s.Mean-float64(k)*s.StdDev) - ref.Mean) / ref.StdDev,
		High: (min(1, s.Mean+float64(k)*s.StdDev) - ref.Mean) / ref.StdDev,
	}
}

func (m *CohortComparison) radarView() string {
	height := max(17, m.gridSize.Height-13)
	o, span, unavailable := m.radarOptions()
	var pills []string
	for i := range m.data.Metrics {
		pills = append(pills, m.metricLabel(i, true))
	}
	curve := "smooth"
	if !m.smooth {
		curve = "straight"
	}
	legend := []string{}
	for i, s := range o.Series {
		legend = append(legend, m.theme.Text().Foreground(m.theme.SeriesColor(i)).Render(fmt.Sprintf("%c %s", s.Marker, []string{"reference", "candidate", "selected host"}[i])))
	}
	plot := chart.Radar(o, layout.Size{Width: max(1, m.gridSize.Width-2), Height: height - 6}, m.theme)
	if len(o.Axes) < 3 {
		plot = layout.Fit("Radar needs at least three metrics; use r for details.", layout.Size{Width: max(1, m.gridSize.Width-2), Height: height - 6})
	} else if len(unavailable) == len(o.Axes) {
		plot = layout.Fit("Reference spread unavailable: no normalized profile.\nRaw fleet rates and host observations remain in the detail view.", layout.Size{Width: max(1, m.gridSize.Width-2), Height: height - 6})
	}
	status := "[ / ] select axis; r details; s curve"
	if len(unavailable) > 0 {
		status = "n/a reference spread: " + strings.Join(unavailable, ", ")
	}
	caption := fmt.Sprintf("0 = reference mean; range +/-%.1fs(ref); whiskers +/- %ds", span*1.05, m.sigma)
	content := strings.Join([]string{strings.Join(pills, " "), strings.Join(legend, " / "), plot, caption, status}, "\n")
	summaries := m.summaries[m.metric]
	delta := "n/a"
	if summaries[0].HasRate && summaries[1].HasRate {
		delta = fmt.Sprintf("%+.3fpp", (summaries[1].Rate-summaries[0].Rate)*100)
	}
	detail := []string{
		fmt.Sprintf("Fleet %s %s  %s %s", m.data.Cohorts[0].Short, rate(summaries[0].Rate, summaries[0].HasRate), m.data.Cohorts[1].Short, rate(summaries[1].Rate, summaries[1].HasRate)),
		"Delta " + delta + " (request weighted)",
		"History %: 1=" + m.data.Cohorts[0].Short + " 2=" + m.data.Cohorts[1].Short + " 3=member",
		m.plots[m.metric].View(),
	}
	return layout.Column(layout.Panel("Version profile / "+curve, content, layout.Size{Width: m.gridSize.Width, Height: height}, m.theme, m.focus.Current() == "metrics"),
		layout.Panel(m.metricLabel(m.metric, false), strings.Join(detail, "\n"), layout.Size{Width: m.gridSize.Width, Height: 13}, m.theme, false))
}
