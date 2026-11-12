package story

import (
	"time"

	"github.com/lanej/charmed/cohort"
	"github.com/lanej/charmed/examples/deployment/fixture"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/runtime"
)

func screensDefinitions() []Definition {
	d := []Definition{}
	d = append(d,
		Definition{ID: "screen/services", Title: "Screens/Service workbench", Level: "Screens", Family: "pattern", Description: "Filter a service list and inspect linked throughput and details. Resize preserves the selected service.", Source: "pattern/dashboard.go", Symbol: "NewDashboard", Upstream: "https://github.com/charmbracelet/bubbletea", Controls: []string{"state", "legend"}, States: []string{"ready", "empty"}, Composes: []string{"input/text", "collection/table", "feedback/status", "chart/timeseries", "collection/viewport", "foundations/layout"}, Clock: true, Step: true, Defaults: func() Config { c := DefaultConfig(); c.Columns = 120; c.Rows = 34; return c }(), Factory: servicesStory},
		Definition{ID: "screen/deployment", Title: "Screens/Deployment comparison", Level: "Screens", Family: "pattern", Description: "Compare 2xx, 3xx, 4xx and 5xx by version identifier during a blue/green deployment. Reference/candidate roles, request-weighted histories, 1/2/3s host-spread bands and linked inspection use synthetic interval counts.", Source: "examples/deployment/main.go", Upstream: "https://github.com/NimbleMarkets/ntcharts", Controls: []string{"state", "scenario", "sigma", "referenceID", "candidateID", "boundaries", "comparisonVariant"}, States: []string{"ready", "empty"}, Scenarios: []string{"isolated-hosts", "uniform-regression", "unequal-traffic", "insufficient-samples"}, Composes: []string{"chart/comparison-panel", "chart/band-matrix", "chart/delta", "collection/table", "feedback/percent", "text/pill", "feedback/attention", "collection/viewport"}, Step: true, Defaults: func() Config { c := DefaultConfig(); c.Columns = 132; c.Rows = 40; c.Boundaries = true; return c }(), Factory: deploymentStory},
		Definition{ID: "template/report-page", Title: "Templates/Report page", Level: "Templates", Family: "pattern", Description: "A finite report: header with the overall evidence state, a summary line, titled panels that sit side by side or stack at 55 columns and narrower, and the evidence footer. The overall state is the worst counted state; no evidence reads unknown.", Source: "pattern/report.go", Symbol: "ReportPage", States: []string{"ready", "error", "empty"}, Composes: []string{"layout/header", "feedback/evidence-summary", "feedback/evidence-footer", "foundations/layout", "feedback/evidence"}, Defaults: func() Config { c := DefaultConfig(); c.Columns = 100; c.Rows = 24; return c }(), Factory: reportPageStory},
	)
	return d
}

func reportPageStory(e Environment) (runtime.Scene, error) {
	now := time.Date(2026, 11, 11, 12, 0, 0, 0, time.UTC)
	counts := []feedback.EvidenceCount{{State: feedback.EvidenceHealthy, N: 2}}

	switch e.Config.State {
	case "error":
		counts = append(counts, feedback.EvidenceCount{State: feedback.EvidenceStale, N: 1}, feedback.EvidenceCount{State: feedback.EvidenceTruncated, N: 1})
	case "empty":
		counts = nil
	}

	page := pattern.ReportPage{
		Title: "Service processes", Context: "api", Counts: counts,
		Panels: []pattern.ReportPanel{
			{Title: "Hosts", Weight: 3, Body: func(int) string {
				return "web1west  observed  live    45s\nweb2west  observed  cached  4m\nworker1   observed  live    30s"
			}},
			{Title: "Processes", Weight: 2, Beside: true, Body: func(int) string { return "db:db       RUNNING\nhttp:http  RUNNING\nworker     FATAL" }},
			{Title: "Limits", Body: func(int) string { return "1 host not retained" }},
		},
		Footer: feedback.EvidenceFooter{Source: "supervisor", Scope: "api", ObservedAt: now.Add(-14 * time.Second), Now: now},
	}

	return static(func(size layout.Size) string { return page.ViewSize(size, e.Config.Palette()) }), nil
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
		Data: data, Theme: t, MinRequests: 1000, Sigma: e.Config.Sigma, Emit: e.Emit, MetricTones: fixture.StatusTones(), Boundaries: boundaries,
		Variant: pattern.ComparisonVariant(e.Config.ComparisonVariant), MetricSigns: fixture.StatusSigns()})
}

func servicesStory(e Environment) (runtime.Scene, error) {
	services := pattern.SampleServices()
	if e.Config.State == "empty" {
		services = nil
	}
	return pattern.NewDashboard(pattern.DashboardOptions{Theme: e.Config.Palette(), Services: services, Live: e.Config.Live, ShowLegend: e.Config.Legend, Emit: e.Emit}), nil
}
