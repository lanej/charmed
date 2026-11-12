package snippets

import (
	"time"

	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleChartOhlc() (string, error) {
	start := time.Date(2026, 11, 8, 12, 0, 0, 0, time.UTC)
	data := []chart.Candle{
		{Time: start, Open: 68, High: 77, Low: 61, Close: 72},
		{Time: start.Add(time.Minute), Open: 72, High: 81, Low: 65, Close: 76},
		{Time: start.Add(2 * time.Minute), Open: 76, High: 79, Low: 64, Close: 67},
		{Time: start.Add(3 * time.Minute), Open: 67, High: 74, Low: 62, Close: 71},
	}

	plot, err := chart.NewCandles(data, layout.Size{Width: 64, Height: 14}, theme.New(theme.Dark))
	if err != nil {
		return "", err
	}

	return plot.View(), nil
}
