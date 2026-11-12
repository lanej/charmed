package snippets

import (
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleChartMetric() string {
	history := []float64{900, 950, 920, 1040, 1100, 1080, 1240}

	return chart.MetricTrend("Throughput", "1,240", "req/min", history,
		layout.Size{Width: 32, Height: 6}, theme.New(theme.Dark))
}
