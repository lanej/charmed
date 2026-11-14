package chart_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

var chartFrame string

func BenchmarkChartScale(b *testing.B) {
	for _, count := range []int{100, 1000, 10000} {
		b.Run(fmt.Sprint(count), func(b *testing.B) {
			points := make([]timeserieslinechart.TimePoint, count)
			for i := range points {
				points[i] = timeserieslinechart.TimePoint{Time: time.Unix(int64(i), 0), Value: math.Sin(float64(i))}
			}

			series := []chart.TimeSeriesData{{ID: "observed", Points: points}}
			plot := chart.NewTimeSeries(series, layout.Size{Width: 80, Height: 24}, theme.New(theme.Dark))

			for _, operation := range []string{"height", "width", "replace", "warm"} {
				b.Run(operation, func(b *testing.B) {
					b.ReportAllocs()

					i := 0

					for b.Loop() {
						switch operation {
						case "height":
							plot.Resize(layout.Size{Width: 80, Height: 20 + i%2*4})
						case "width":
							plot.Resize(layout.Size{Width: 60 + i%2*20, Height: 24})
						case "replace":
							plot.SetData(series)
						case "warm":
						}

						chartFrame = plot.View()
						i++
					}
				})
			}
		})
	}
}
