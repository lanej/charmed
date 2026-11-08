package chart_test

import (
	"math"
	"strconv"
	"strings"
	"testing"
	"time"

	"charm.land/lipgloss/v2"
	"github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func TestMetricTrendPrioritizesValueAndDistinguishesMissingHistory(t *testing.T) {
	t.Parallel()

	palette := theme.New(theme.Mono)

	for _, height := range []int{1, 2, 3, 8} {
		view := ansi.Strip(chart.MetricTrend("Request throughput", "0", "req/min", nil, layout.Size{Width: 24, Height: height}, palette))
		if lipgloss.Width(view) != 24 || lipgloss.Height(view) != height || !strings.Contains(view, "0 req/min") {
			t.Fatalf("%d-row metric lost its current value or allocation: %q", height, view)
		}

		if height >= 3 && !strings.Contains(view, "No history") {
			t.Fatal("missing history was rendered as observations")
		}
	}

	view := ansi.Strip(chart.MetricTrend("Throughput", "", "req/min", []float64{math.NaN(), math.Inf(1)}, layout.Size{Width: 24, Height: 8}, palette))
	if !strings.Contains(view, "- req/min") || !strings.Contains(view, "No history") || strings.Contains(view, "NaN") {
		t.Fatal("missing value and invalid history produced a false observation")
	}

	view = ansi.Strip(chart.MetricTrend("Throughput", "42", "req/min", []float64{10, 42, math.NaN()}, layout.Size{Width: 24, Height: 8}, palette))
	if !strings.Contains(view, "latest 42.0") {
		t.Fatal("invalid trailing sample hid the last usable history value")
	}
}

func TestPanelAllocatesWrappedLegendWithoutMutatingThePlot(t *testing.T) {
	t.Parallel()

	palette := theme.New(theme.Dark)
	size := layout.Size{Width: 40, Height: 18}
	base := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	series := []chart.TimeSeriesData{
		{ID: "requests", Label: "Request throughput"},
		{ID: "jobs", Label: "Background workers"},
		{ID: "storage", Label: "Storage operations"},
	}

	for i := range series {
		series[i].Points = []timeserieslinechart.TimePoint{{Time: base, Value: 10}, {Time: base.Add(time.Hour), Value: 20}}
	}

	options := chart.PanelOptions{ShowLegend: true}
	allocation := options.PlotSize(size, series, palette)

	if allocation.Height != layout.PanelContentSize(size).Height-3 {
		t.Fatal("the three legend rows were not reserved")
	}

	plot := chart.NewTimeSeries(series, allocation, palette)
	plot.Bubble.SetViewTimeRange(base.Add(10*time.Minute), base.Add(50*time.Minute))
	before := plot.View()
	view := chart.Panel("Rate", plot, series, size, palette, false, options)

	if plot.View() != before {
		t.Fatal("panel rendering mutated the native chart")
	}

	if lipgloss.Width(view) != size.Width || lipgloss.Height(view) != size.Height {
		t.Fatal("panel escaped its allocation")
	}

	for i, label := range []string{"Request throughput", "Background workers", "Storage operations"} {
		if !strings.Contains(ansi.Strip(view), label) {
			t.Errorf("legend clipped %s", label)
		}

		key := lipgloss.NewStyle().Foreground(palette.SeriesColor(i)).Render(strconv.Itoa(i + 1))
		if !strings.Contains(view, key) {
			t.Errorf("series %d legend lost its plotted color", i)
		}
	}

	if options.PlotSize(size, nil, palette) != layout.PanelContentSize(size) {
		t.Fatal("an empty legend reserved a phantom row")
	}
}

func TestLongLegendDescriptionsWrapInsteadOfDisappearing(t *testing.T) {
	t.Parallel()

	palette := theme.New(theme.Dark)
	label := "Requests served by background workers"
	legend := chart.SeriesLegend([]chart.TimeSeriesData{{ID: "requests", Label: label}}, 18, palette)

	if !strings.Contains(strings.Join(strings.Fields(ansi.Strip(legend)), " "), label) {
		t.Fatal("wrapping discarded part of the series label")
	}

	for _, line := range strings.Split(legend, "\n") {
		if ansi.StringWidth(line) > 18 {
			t.Fatal("wrapped legend exceeds its width")
		}
	}

	series := []chart.TimeSeriesData{{Label: "Requests\nBackground"}}
	for _, multiline := range []string{chart.Legend(series, palette), chart.SeriesLegend(series, 40, palette)} {
		if !strings.Contains(multiline, "Requests") || !strings.Contains(multiline, "Background") {
			t.Fatal("legend discarded an explicit label line")
		}
	}
}

func TestDistributionRowsKeepValuesUnitsAndMissingStatesAtNarrowWidths(t *testing.T) {
	t.Parallel()

	palette := theme.New(theme.Mono)
	distribution := chart.DistributionComparison{Compact: true, Min: 0, Max: 100, Unit: "%", Series: []chart.DistributionSeries{
		{Label: "Reference", Mean: 20, Members: 10},
		{Label: "Candidate", Mean: 40, Members: 10},
	}}

	for _, width := range []int{16, 20, 24, 40} {
		view := distribution.View(width, palette)
		for _, line := range strings.Split(view, "\n") {
			if ansi.StringWidth(line) > width {
				t.Fatalf("distribution escaped %d columns: %q", width, line)
			}
		}

		if !strings.Contains(view, "20.000%") || !strings.Contains(view, "40.000%") {
			t.Fatal("compact labels clipped a value or its unit")
		}
	}

	distribution.Compact = false
	view := distribution.View(24, palette)

	if !strings.Contains(strings.Join(strings.Fields(view), " "), "Reference mean 20.000% s=n/a n=10") {
		t.Fatal("expanded distribution dropped wrapped metadata")
	}

	distribution.Series = nil
	if !strings.Contains(distribution.View(24, palette), "No observations") || distribution.View(0, palette) != "" {
		t.Fatal("empty distribution has no explicit bounded state")
	}
}
