package chart_test

import (
	"math"
	"strings"
	"testing"
	"time"

	"github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func TestLargeConstantHistoryKeepsItsActualRange(t *testing.T) {
	for _, value := range []float64{1e-200, 1e200, math.MaxFloat64} {
		model := chart.NewTimeSeries([]chart.TimeSeriesData{{ID: "value", Points: []timeserieslinechart.TimePoint{{Time: time.Unix(0, 0), Value: value}}}}, layout.Size{Width: 50, Height: 9}, theme.New(theme.Dark))

		low, high := model.Bubble.ViewMinY(), model.Bubble.ViewMaxY()
		if math.IsInf(low, 0) || math.IsInf(high, 0) || !(low < high) || low > value || high < value {
			t.Fatalf("%g replaced by range %g..%g", value, low, high)
		}
	}
}

func TestUnrepresentableHistoryRangeIsExplicit(t *testing.T) {
	model := chart.NewTimeSeries([]chart.TimeSeriesData{{ID: "value", Points: []timeserieslinechart.TimePoint{{Time: time.Unix(0, 0), Value: -math.MaxFloat64}, {Time: time.Unix(1, 0), Value: math.MaxFloat64}}}}, layout.Size{Width: 50, Height: 9}, theme.New(theme.Dark))
	if !strings.Contains(model.View(), "History range unavailable") {
		t.Fatal("unrepresentable range used a misleading native fallback")
	}
}

func TestOffsetRangePreservesSmallVariation(t *testing.T) {
	low, high, ok := chart.PaddedRange(1e9, 1e9+1, 0.15)
	if !ok || math.Abs((high-low)-1.3) > 1e-6 {
		t.Fatalf("one-unit variation flattened by range %g..%g", low, high)
	}
}

func TestNativeViewportRecoversAfterUnrepresentableFullRange(t *testing.T) {
	model := chart.NewTimeSeries([]chart.TimeSeriesData{{ID: "value", Points: []timeserieslinechart.TimePoint{{Time: time.Unix(0, 0), Value: -math.MaxFloat64}, {Time: time.Unix(1, 0), Value: math.MaxFloat64}}}}, layout.Size{Width: 50, Height: 10}, theme.New(theme.Dark))
	model.Bubble.YLabelFormatter = func(_ int, _ float64) string { return "VIEWPORT" }
	model.Bubble.SetViewYRange(-10, 10)

	model.Bubble.DrawBrailleAll()

	if !strings.Contains(ansi.Strip(model.View()), "VIEWPORT") {
		t.Fatalf("full-span warning prevented native viewport rendering: %s", model.View())
	}
}
