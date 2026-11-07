package chart_test

import (
	"math"
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func TestBandLegendUsesPlottedMarkerStyles(t *testing.T) {
	t.Parallel()
	palette := theme.New(theme.Dark)
	warn := lipgloss.NewStyle().Foreground(palette.Warning).Bold(true)
	failure := lipgloss.NewStyle().Foreground(palette.Danger).Bold(true)
	mean := lipgloss.NewStyle().Foreground(lipgloss.Color("#123456"))
	selected := 90.0
	opts := chart.BandOptions{Min: 0, Max: 100, Center: 50, Selected: &selected, Style: &mean, Markers: []chart.BandMarker{
		{Value: 20, Glyph: 'W', Label: "warning boundary", Style: warn},
		{Value: 30, Glyph: 'E', Label: "error boundary", Style: failure},
		{Value: 200, Glyph: 'Z', Label: "outside domain", Style: failure},
	}}
	band, legend := chart.IntervalBand(opts, 101, palette), opts.Legend(40, palette)
	for _, key := range []string{warn.Render("W"), failure.Render("E"), mean.Bold(true).Render("|"), palette.Text().Bold(true).Render("X")} {
		if !strings.Contains(band, key) || !strings.Contains(legend, key) {
			t.Errorf("plot and legend disagree on styled key %q", key)
		}
	}
	if strings.Contains(legend, "outside domain") {
		t.Error("legend includes an invisible marker")
	}
	for _, line := range strings.Split(legend, "\n") {
		if ansi.StringWidth(line) > 40 {
			t.Error("legend exceeds its allocation")
		}
	}
	opts.Intervals = []chart.Interval{{Low: math.NaN(), High: 10}}
	if opts.Legend(40, palette) != "" {
		t.Error("invalid geometry has a misleading legend")
	}
}

func TestDistributionLegendIsOptionalAndKeepsBothMeanColors(t *testing.T) {
	palette := theme.New(theme.Dark)
	d := chart.DistributionComparison{Min: 0, Max: 100, Sigma: 1, Series: []chart.DistributionSeries{{Label: "Reference", Mean: 20, Members: 10}, {Label: "Candidate", Mean: 40, Members: 10}}}
	without := d.View(80, palette)
	d.ShowLegend = true
	with := d.View(80, palette)
	if strings.Contains(ansi.Strip(without), "| Reference mean") {
		t.Fatal("default distribution includes a legend")
	}
	for _, label := range []string{"Reference mean", "Candidate mean"} {
		if !strings.Contains(ansi.Strip(with), label) {
			t.Errorf("missing %s", label)
		}
	}
	for i := range 2 {
		key := lipgloss.NewStyle().Foreground(palette.SeriesColor(i)).Bold(true).Render("|")
		if !strings.Contains(with, key) {
			t.Errorf("mean key lost series %d color", i)
		}
	}
}

func TestPanelLegendIsOptionalAndReservesItsRow(t *testing.T) {
	palette := theme.New(theme.Dark)
	size := layout.Size{Width: 50, Height: 12}
	series := []chart.TimeSeriesData{{ID: "observed", Label: "Observed"}}
	plot := chart.NewTimeSeries(series, chart.PanelPlotSize(size), palette)
	if strings.Contains(ansi.Strip(chart.Panel("Rate", plot, series, size, palette, false)), "1 Observed") {
		t.Fatal("default panel includes legend")
	}
	options := chart.PanelOptions{ShowLegend: true}
	with := chart.PanelPlotSize(size, options)
	if with.Height != chart.PanelPlotSize(size).Height-1 {
		t.Fatal("legend has no allocated row")
	}
	plot.Resize(with)
	if !strings.Contains(chart.Panel("Rate", plot, series, size, palette, false, options), chart.Legend(series, palette)) {
		t.Fatal("opt-in legend is absent or has different series colors")
	}
}
