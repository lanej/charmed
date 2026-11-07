package chart_test

import (
	"fmt"
	"math"
	"strings"
	"testing"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func TestResizePreservesChartViewport(t *testing.T) {
	base := time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)
	data := []chart.TimeSeriesData{{ID: "observed", Points: []timeserieslinechart.TimePoint{{Time: base, Value: 10}, {Time: base.Add(time.Hour), Value: 20}}}}
	plot := chart.NewTimeSeries(data, layout.Size{Width: 80, Height: 20}, theme.New(theme.Dark))
	plot.Bubble.SetViewTimeRange(base.Add(10*time.Minute), base.Add(40*time.Minute))
	beforeMin, beforeMax := plot.Bubble.ViewMinX(), plot.Bubble.ViewMaxX()
	plot.Resize(layout.Size{Width: 60, Height: 14})
	if plot.Bubble.ViewMinX() != beforeMin || plot.Bubble.ViewMaxX() != beforeMax {
		t.Fatal("resize reset the selected time window")
	}

	assertDots := func(stage string) {
		t.Helper()
		if !strings.ContainsFunc(ansi.Strip(plot.View()), func(r rune) bool { return r >= '\u2801' && r <= '\u28ff' }) {
			t.Fatalf("%s lost dot rendering", stage)
		}
	}
	assertDots("resize")
	plot.Focus()
	plot, _ = plot.Update(tea.KeyPressMsg{Code: tea.KeyRight})
	assertDots("pan")
	plot.SetData(data)
	assertDots("data update")
}

func TestTimeSeriesRetainsNativePrecisionOptions(t *testing.T) {
	base := time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)
	data := []chart.TimeSeriesData{{ID: "ratio", Points: []timeserieslinechart.TimePoint{{Time: base, Value: 0.12}, {Time: base.Add(time.Hour), Value: 0.24}}}}
	plot := chart.NewTimeSeries(data, layout.Size{Width: 50, Height: 10}, theme.New(theme.Dark),
		timeserieslinechart.WithYLabelFormatter(func(_ int, value float64) string { return fmt.Sprintf("%.2f", value) }))
	plot.SetData(data)
	plot.Resize(layout.Size{Width: 60, Height: 12})
	if !strings.Contains(ansi.Strip(plot.View()), "0.") {
		t.Fatal("fractional axis precision disappeared after data update or resize")
	}
}

func TestCandlesRejectInvalidFinancialGeometry(t *testing.T) {
	for _, candle := range []chart.Candle{{Open: 10, High: 9, Low: 8, Close: 10}, {Open: 10, High: 12, Low: 11, Close: 12}, {Open: math.NaN(), High: 12, Low: 8, Close: 10}} {
		if _, err := chart.NewCandles([]chart.Candle{candle}, layout.Size{Width: 40, Height: 12}, theme.New(theme.Dark)); err == nil {
			t.Fatalf("accepted invalid candle %+v", candle)
		}
	}
}
