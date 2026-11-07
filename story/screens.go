package story

import (
	"github.com/lanej/charmed/cohort"
	"github.com/lanej/charmed/examples/deployment/fixture"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/runtime"
)

func screensDefinitions() []Definition {
	d := []Definition{}
	d = append(d,
		Definition{ID: "screen/services", Title: "Screens/Service workbench", Level: "Screens", Family: "pattern", Description: "Filter a service list and inspect linked throughput and details. Resize preserves the selected service.", Source: "pattern/dashboard.go", Symbol: "NewDashboard", Upstream: "https://github.com/charmbracelet/bubbletea", Controls: []string{"state"}, States: []string{"ready", "empty"}, Composes: []string{"input/text", "collection/table", "chart/timeseries", "collection/viewport", "foundations/layout"}, Clock: true, Step: true, Defaults: func() Config { c := DefaultConfig(); c.Columns = 120; c.Rows = 34; return c }(), Factory: servicesStory},
		Definition{ID: "screen/deployment", Title: "Screens/Deployment comparison", Level: "Screens", Family: "pattern", Description: "Compare 2xx, 3xx, 4xx and 5xx by version identifier during a blue/green deployment. Reference/candidate roles, request-weighted histories, 1/2/3s host-spread bands and linked inspection use synthetic interval counts.", Source: "examples/deployment/main.go", Upstream: "https://github.com/NimbleMarkets/ntcharts", Controls: []string{"state", "scenario", "sigma", "referenceID", "candidateID", "boundaries"}, States: []string{"ready", "empty"}, Scenarios: []string{"isolated-hosts", "uniform-regression", "unequal-traffic", "insufficient-samples"}, Composes: []string{"chart/comparison-panel", "chart/delta", "collection/table", "feedback/percent", "text/primitives", "feedback/attention", "collection/viewport"}, Step: true, Defaults: func() Config { c := DefaultConfig(); c.Columns = 132; c.Rows = 40; c.Boundaries = true; return c }(), Factory: deploymentStory},
	)
	return d
}

func deploymentStory(e Environment) (runtime.Scene, error) {
	data, err := fixture.DataWithVersions(e.Config.Scenario, fixture.Versions{Reference: e.Config.ReferenceID, Candidate: e.Config.CandidateID})
	if err != nil {
		return nil, err
	}
	if e.Config.State == "empty" {
		data.Windows = nil
	}
	t := e.Config.Palette()
	var boundaries map[string]cohort.ComparisonPolicy
	if e.Config.Boundaries {
		boundaries = fixture.Boundaries()
	}
	return pattern.NewCohortComparison(pattern.CohortComparisonOptions{Title: "Deployment comparison", Scope: "method=blue/green / SYNTHETIC / " + e.Config.Scenario,
		Data: data, Theme: t, MinRequests: 1000, Sigma: e.Config.Sigma, Emit: e.Emit, MetricTones: fixture.StatusTones(), Boundaries: boundaries})
}

func servicesStory(e Environment) (runtime.Scene, error) {
	services := pattern.SampleServices()
	if e.Config.State == "empty" {
		services = nil
	}
	return pattern.NewDashboard(pattern.DashboardOptions{Theme: e.Config.Palette(), Services: services, Live: e.Config.Live, Emit: e.Emit}), nil
}
