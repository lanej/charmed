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
