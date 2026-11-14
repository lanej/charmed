package chart_test

import (
	"math"
	"testing"
	"time"

	"github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func TestApplyChartSnapshotOwnsDataAndSkipsEqualObservations(t *testing.T) {
	series := []chart.TimeSeriesData{{ID: "metric", Points: []timeserieslinechart.TimePoint{{Time: time.Unix(0, 0), Value: 1}, {Time: time.Unix(5, 0), Value: 2}}}}
	m := chart.NewTimeSeries(series, layout.Size{Width: 60, Height: 12}, theme.New(theme.Dark))
	m.Bubble.SetViewYRange(1.1, 1.4)

	before := m.Bubble.ViewMaxY()
	if m.ApplyData(series) {
		t.Fatal("equal snapshot rebuilt chart")
	}

	if m.Bubble.ViewMaxY() != before {
		t.Fatal("equal snapshot reset native customization")
	}

	series[0].Points[1].Value = 3
	if !m.ApplyData(series) {
		t.Fatal("changed point not applied")
	}

	fresh := chart.NewTimeSeries(series, layout.Size{Width: 60, Height: 12}, theme.New(theme.Dark))
	if m.View() != fresh.View() {
		t.Fatal("changed snapshot differs from full replacement")
	}

	if m.ApplyData(series) {
		t.Fatal("repeated snapshot rebuilt chart")
	}

	if allocs := testing.AllocsPerRun(10, func() { m.ApplyData(series) }); allocs != 0 {
		t.Fatalf("unchanged chart allocated %v", allocs)
	}
}

func TestApplyEmptyAndNonfiniteChartSnapshots(t *testing.T) {
	m := chart.NewTimeSeries(nil, layout.Size{Width: 60, Height: 12}, theme.New(theme.Dark))
	if m.ApplyData([]chart.TimeSeriesData{}) {
		t.Fatal("empty snapshot rebuilt chart")
	}

	values := []float64{1, math.NaN(), 2}

	spark := chart.NewSparkline(values, layout.Size{Width: 30, Height: 3}, theme.New(theme.Dark))
	if spark.ApplyData(values) {
		t.Fatal("same nonfinite observation rebuilt sparkline")
	}

	values[2] = 3
	if !spark.ApplyData(values) {
		t.Fatal("sparkline retained caller alias")
	}
}
