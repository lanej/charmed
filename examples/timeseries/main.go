package main

import (
	"fmt"
	"math"
	"time"

	"github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func main() {
	palette := theme.New(theme.Dark)
	size := layout.Size{Width: 80, Height: 18}
	series := []chart.TimeSeriesData{{ID: "requests", Label: "Requests", Unit: "req/min"}}
	start := time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)
	for i := range 32 {
		series[0].Points = append(series[0].Points, timeserieslinechart.TimePoint{
			Time:  start.Add(time.Duration(i) * time.Minute),
			Value: 800 + 120*math.Sin(float64(i)*0.3),
		})
	}
	plot := chart.NewTimeSeries(series, chart.PanelPlotSize(size), palette)
	fmt.Println(chart.Panel("Request throughput", plot, series, size, palette, false))
}
