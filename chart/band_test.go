package chart_test

import (
	"math"
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/theme"
)

func TestIntervalGeometryAndObservationMarker(t *testing.T) {
	selected := 95.0
	options := chart.BandOptions{Min: 0, Max: 100, Center: 50, Selected: &selected, Intervals: []chart.Interval{{Low: 30, High: 70}, {Low: 10, High: 90}, {Low: 20, High: 80}}}
	band := []rune(ansi.Strip(chart.IntervalBand(options, 101, theme.New(theme.Dark))))
	if len(band) != 101 || band[15] != '░' || band[25] != '▒' || band[35] != '▓' || band[50] != '|' || band[95] != 'X' {
		t.Fatalf("incorrect nested geometry: %s", string(band))
	}
	if options.Intervals[0].Low != 30 {
		t.Fatal("render reordered caller-owned intervals")
	}
	options.Intervals = []chart.Interval{{Low: -50, High: 30}}
	if !strings.HasPrefix(ansi.Strip(chart.IntervalBand(options, 101, theme.New(theme.Mono))), "#") {
		t.Fatal("interval did not clip at domain boundary")
	}
	options.Max = math.NaN()
	if !strings.Contains(chart.IntervalBand(options, 30, theme.New(theme.Mono)), "Invalid interval") {
		t.Fatal("invalid scale rendered as data")
	}
}

func TestBandBoundaryMarkersComposeWithSelection(t *testing.T) {
	selected := 90.0
	options := chart.BandOptions{Min: 0, Max: 100, Center: 50, Selected: &selected, Markers: []chart.BandMarker{{Value: 20, Glyph: 'W'}, {Value: 30, Glyph: 'E'}, {Value: 200, Glyph: 'E'}}}
	band := ansi.Strip(chart.IntervalBand(options, 101, theme.New(theme.Mono)))
	if band[20] != 'W' || band[30] != 'E' || band[50] != '|' || band[90] != 'X' || strings.Count(band, "E") != 1 {
		t.Fatal("boundary geometry or selected marker was lost")
	}
}
