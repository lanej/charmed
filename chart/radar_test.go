package chart

import (
	"math"
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func TestSmoothRadarKeepsAxisValuesAndNeverOvershoots(t *testing.T) {
	a, b := RadarValue{Value: -2, Valid: true}, RadarValue{Value: 3, Valid: true}
	points := radarSegment(a, b, math.Pi/2, -math.Pi/2, -4, 4, true, 100)
	for i, p := range points {
		r := math.Hypot(p.X, p.Y)
		if r < .25-1e-12 || r > .875+1e-12 {
			t.Fatalf("smooth connection overshot between observations at %d: %g", i, r)
		}
	}
	if math.Abs(points[0].Y-.25) > 1e-12 || math.Abs(points[100].X-.875) > 1e-12 || math.Abs(points[100].Y) > 1e-12 {
		t.Fatal("smoothing moved an observed axis value")
	}
	straight := radarSegment(a, b, math.Pi/2, -math.Pi/2, -4, 4, false, 100)
	if points[50] == straight[50] {
		t.Fatal("smooth mode did not change the connecting geometry")
	}
	// The closed seam has the same tangent from either side of the first axis.
	first := radarSegment(a, b, math.Pi/2, -math.Pi/2, -4, 4, true, 10000)
	last := radarSegment(b, a, math.Pi, -math.Pi/2, -4, 4, true, 10000)
	leftSlope := (last[len(last)-1].Y - last[len(last)-2].Y) / (last[len(last)-1].X - last[len(last)-2].X)
	rightSlope := (first[1].Y - first[0].Y) / (first[1].X - first[0].X)
	if math.Abs(leftSlope-rightSlope) > .01 {
		t.Fatal("closed profile has a visible corner at the seam")
	}
	if len(radarSegment(a, RadarValue{}, 0, 1, -4, 4, true, 10)) != 0 {
		t.Fatal("missing observation acquired an interpolated connection")
	}
}

func TestRadarRejectsInvalidGeometryAndPreservesMissingValues(t *testing.T) {
	o := RadarOptions{Axes: []string{"a", "b", "c"}, Min: -4, Max: 4, Reference: 0, Smooth: true,
		Series: []RadarSeries{{Marker: 'R', Values: []RadarValue{{Value: 1, Valid: true}, {}, {Value: 2, Valid: true}}}}}
	frame := ansi.Strip(Radar(o, layout.Size{Width: 70, Height: 20}, theme.New(theme.Dark)))
	if strings.Contains(frame, "Invalid") || !strings.Contains(frame, "R") {
		t.Fatal("valid incomplete profile failed to render")
	}
	frame = ansi.Strip(Radar(o, layout.Size{Width: 70, Height: 10}, theme.New(theme.Mono)))
	if !strings.Contains(frame, "b / R n/a") {
		t.Fatal("text fallback invented a missing observation")
	}
	o.Series[0].Values[0].Value = math.NaN()
	if !strings.Contains(Radar(o, layout.Size{Width: 70, Height: 20}, theme.New(theme.Dark)), "Invalid radar geometry") {
		t.Fatal("non-finite observation rendered as a profile")
	}
}
