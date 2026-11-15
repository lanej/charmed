package chart_test

import (
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func TestCompactBandsShareTheMetricScaleAndKeepDirectLabels(t *testing.T) {
	selected := 44.0
	d := chart.DistributionComparison{Compact: true, Sigma: 3, Min: 0, Max: 100, Unit: "%", Series: []chart.DistributionSeries{
		{Label: "release-42", Mean: 20, StdDev: 2, HasSpread: true, Members: 10},
		{Label: "release-43", Mean: 40, StdDev: 2, HasSpread: true, Members: 10, Selected: &selected},
	}}
	lines := strings.Split(ansi.Strip(d.View(80, theme.New(theme.Mono))), "\n")
	if len(lines) != 3 || !strings.Contains(lines[0], "release-42") || !strings.Contains(lines[1], "release-43") ||
		strings.Index(lines[0], "|") >= strings.Index(lines[1], "|") || !strings.Contains(lines[1], "X") {
		t.Fatalf("compact bands lost shared scale, direct labels or selected host: %v", lines)
	}
	matrix := chart.BandMatrix{Title: "Version bands", Rows: []chart.BandRow{{Title: "Success", Distribution: d}}}
	frame := ansi.Strip(matrix.View(layout.Size{Width: 82, Height: 12}, theme.New(theme.Mono), false))
	if !strings.Contains(frame, "Success") || strings.Contains(frame, "legend") || strings.Contains(frame, "History") {
		t.Fatal("matrix did not compose the independently labeled row")
	}
	d.Series[0].Members = 0
	if !strings.Contains(ansi.Strip(d.View(80, theme.New(theme.Mono))), "release-42 n/a") {
		t.Fatal("missing samples rendered as a zero mean")
	}
}

func TestBandMatrixSharesFallbackColumnsAndMeasuresRenderedHeight(t *testing.T) {
	short := chart.DistributionComparison{Min: 0, Max: 100, Unit: "%", Series: []chart.DistributionSeries{{Label: "North", Mean: 20, Members: 2}, {Label: "South", Mean: 40, Members: 2}}}
	long := short
	long.Unit = strings.Repeat("界", 45)
	matrix := chart.BandMatrix{Title: "Shared geometry", Rows: []chart.BandRow{{Title: "Short unit", Distribution: short}, {Title: "Wide unit", Distribution: long}}}
	palette := theme.New(theme.Mono)
	frame := ansi.Strip(matrix.View(layout.Size{Width: 82, Height: matrix.Height(82, palette)}, palette, false))
	tracks := 0

	for _, line := range strings.Split(frame, "\n") {
		if !strings.Contains(line, "...") {
			continue
		}

		if column := ansi.StringWidth(line[:strings.Index(line, ".")]); column != 2 {
			t.Fatalf("shared plot start moved to cell %d", column)
		}

		tracks++
	}

	if tracks != 4 || !strings.Contains(frame, "Wide unit") {
		t.Fatal("shared fallback or measured height lost distributions")
	}
}

func TestBandMatrixAlignsBothTrackEdgesAcrossRowTypes(t *testing.T) {
	matrix := chart.BandMatrix{Rows: []chart.BandRow{
		{Title: "Temperature", Comparisons: []chart.DeltaRow{{Label: "South vs North", Property: chart.DeltaProperty{Value: 8, Extent: 50, Available: true, Unit: "C"}}}, Distribution: chart.DistributionComparison{Min: 0, Max: 100, Unit: "C", Series: []chart.DistributionSeries{{Label: "North", Mean: 20, Members: 2}, {Label: "South", Mean: 40, Members: 2}}}},
		{Title: "Energy", Comparisons: []chart.DeltaRow{{Label: "West vs North", Property: chart.DeltaProperty{Value: 20, Extent: 50, Available: true, Unit: "kWh"}}}, Distribution: chart.DistributionComparison{Min: 0, Max: 100, Unit: "kWh", Series: []chart.DistributionSeries{{Label: "West", Mean: 30, Members: 2}}}},
	}}

	palette := theme.New(theme.Mono)
	for _, width := range []int{82, 120} {
		frame := ansi.Strip(matrix.View(layout.Size{Width: width, Height: matrix.Height(width, palette)}, palette, false))
		start, end, tracks := -1, -1, 0

		for _, line := range strings.Split(frame, "\n") {
			left, right := strings.Index(line, "."), strings.LastIndex(line, ".")
			if !strings.Contains(line, "...") {
				continue
			}

			if start < 0 {
				start, end = left, right
			}

			if left != start || right != end {
				t.Fatalf("track edges differ at width %d: [%d,%d] vs [%d,%d]\n%s", width, start, end, left, right, frame)
			}

			tracks++
		}

		if tracks != 5 {
			t.Fatalf("got %d tracks, want both deltas and all three cohorts\n%s", tracks, frame)
		}
	}
}

func TestBandMatrixPreservesRequestedMarkerLegend(t *testing.T) {
	d := chart.DistributionComparison{Min: 0, Max: 100, ShowLegend: true, Series: []chart.DistributionSeries{{Label: "North", Mean: 20, Members: 2}}, Markers: []chart.BandMarker{{Value: 30, Glyph: 'W', Label: "Warning boundary"}}}
	matrix := chart.BandMatrix{Rows: []chart.BandRow{{Title: "Metric", Distribution: d}}}
	palette := theme.New(theme.Dark)

	frame := ansi.Strip(matrix.View(layout.Size{Width: 82, Height: matrix.Height(82, palette)}, palette, false))
	for _, label := range []string{"North mean", "Warning boundary"} {
		if !strings.Contains(frame, label) {
			t.Fatalf("matrix lost legend %q:\n%s", label, frame)
		}
	}
}
