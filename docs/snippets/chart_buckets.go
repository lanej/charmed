package snippets

import (
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/theme"
)

func ExampleChartBuckets() string {
	bars := chart.BucketBars{Buckets: []chart.CountBucket{
		{Label: "22:00 UTC", Count: 160, State: chart.BucketComplete},
		{Label: "23:00 UTC", Count: 173, State: chart.BucketComplete},
		{Label: "00:00 UTC", Count: 0, State: chart.BucketPartial},
		{Label: "01:00 UTC", State: chart.BucketMissing},
		{Label: "02:00 UTC"},
	}}

	return bars.View(64, theme.New(theme.Dark))
}
