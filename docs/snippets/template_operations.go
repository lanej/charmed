package snippets

import (
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

func ExampleTemplateOperations() string {
	palette := theme.New(theme.Dark)
	bars := chart.BucketBars{Buckets: []chart.CountBucket{
		{Label: "22:00 UTC", Count: 160, State: chart.BucketComplete},
		{Label: "23:00 UTC", Count: 173, State: chart.BucketComplete},
		{Label: "00:00 UTC", Count: 0, State: chart.BucketPartial},
	}}
	snapshot := pattern.DashboardSnapshot{
		Title: "API operations", Context: "Synthetic observations / UTC",
		Status:  "Observed 00:16 UTC / quota resets 01:00 UTC",
		Summary: []layout.Pair{{Label: "REST remaining", Value: "3500 / 5000 requests"}},
		Sections: []pattern.DashboardSection{{Title: "REST attempts per hour", Body: func(width int) string {
			return bars.View(width, palette)
		}}},
		Details: []pattern.DashboardSection{{Title: "Scheduler", Body: func(int) string {
			return "0 queued / idle / minimum request gap 2s"
		}}},
	}
	dashboard := pattern.NewOperationalDashboard(snapshot, layout.Size{Width: 80, Height: 18}, palette)

	return dashboard.View()
}
