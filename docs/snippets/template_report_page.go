package snippets

import (
	"time"

	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

func ExampleTemplateReportPage() string {
	now := time.Date(2026, 11, 11, 12, 0, 0, 0, time.UTC)
	page := pattern.ReportPage{
		Title: "Service processes", Context: "api",
		Counts: []feedback.EvidenceCount{{State: feedback.EvidenceHealthy, N: 2}, {State: feedback.EvidenceStale, N: 1}},
		Panels: []pattern.ReportPanel{
			{Title: "Hosts", Weight: 3, Body: func(int) string { return "web1west  observed  45s" }},
			{Title: "Processes", Weight: 2, Beside: true, Body: func(int) string { return "db:db  RUNNING" }},
		},
		Footer: feedback.EvidenceFooter{Source: "supervisor", Scope: "api", ObservedAt: now.Add(-14 * time.Second), Now: now},
	}

	return page.View(80, theme.New(theme.Dark))
}
