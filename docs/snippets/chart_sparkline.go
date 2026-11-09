package snippets

import (
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleChartSparkline() string {
	values := []float64{32, 37, 34, 40, 44, 41, 47, 45, 52, 49}
	spark := chart.NewSparkline(values,
		layout.Size{Width: 48, Height: 5}, theme.New(theme.Dark))

	return spark.View()
}
