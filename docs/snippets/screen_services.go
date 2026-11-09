package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

func ExampleScreenServices() string {
	dashboard := pattern.NewDashboard(pattern.DashboardOptions{
		Theme: theme.New(theme.Dark), ShowLegend: true,
		Services: []pattern.Service{
			{ID: "api", Name: "API gateway", Status: "healthy", Rate: 1240, Latency: 12},
			{ID: "events", Name: "Event stream", Status: "healthy", Rate: 860, Latency: 18},
			{ID: "jobs", Name: "Job workers", Status: "degraded", Rate: 420, Latency: 85},
		},
	})
	dashboard.Resize(layout.Size{Width: 100, Height: 30})
	// Forward Init and Update commands in your application to enable keyboard interaction.
	return dashboard.View()
}
