package story

import (
	"fmt"
	"math"
	"time"

	"github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart"
	"github.com/lanej/charmed/chart"
)

func sampleValues(count, step int, amount float64) []float64 {
	values := make([]float64, count)
	for i := range values {
		values[i] = max(1, amount+14*math.Sin(float64(i+step)*.3)+5*math.Cos(float64(i)*.7))
	}
	return values
}

func sampleSeries(count, step int, amount float64) []chart.TimeSeriesData {
	base := time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)
	data := make([]chart.TimeSeriesData, count)
	for s := range count {
		data[s] = chart.TimeSeriesData{ID: fmt.Sprintf("series-%d", s), Label: fmt.Sprintf("Series %d", s+1), Unit: "req/min"}
		for i, v := range sampleValues(32, step+s*4, amount+float64(s*12)) {
			data[s].Points = append(data[s].Points, timeserieslinechart.TimePoint{Time: base.Add(time.Duration(i) * time.Minute), Value: v})
		}
	}
	return data
}
