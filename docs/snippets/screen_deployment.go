package snippets

import (
	"time"

	"github.com/lanej/charmed/cohort"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

func ExampleScreenDeployment() (string, error) {
	start := time.Date(2026, 11, 8, 12, 0, 0, 0, time.UTC)
	data := cohort.Data{
		Cohorts: []cohort.Cohort{
			{ID: "blue", Label: "Blue", Short: "blue"}, {ID: "green", Label: "Green", Short: "green"},
		},
		Metrics: []cohort.Metric{{ID: "5xx", Label: "Server errors"}},
		Windows: []cohort.Window{{Start: start, End: start.Add(time.Minute), Samples: []cohort.Sample{
			{MemberKey: cohort.MemberKey{CohortID: "blue", MemberID: "api-1"}, Requests: 2000, Counts: map[string]uint64{"5xx": 4}},
			{MemberKey: cohort.MemberKey{CohortID: "blue", MemberID: "api-2"}, Requests: 2000, Counts: map[string]uint64{"5xx": 6}},
			{MemberKey: cohort.MemberKey{CohortID: "green", MemberID: "api-3"}, Requests: 2000, Counts: map[string]uint64{"5xx": 12}},
			{MemberKey: cohort.MemberKey{CohortID: "green", MemberID: "api-4"}, Requests: 2000, Counts: map[string]uint64{"5xx": 18}},
		}}},
	}

	comparison, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{
		Title: "Deployment comparison", Data: data, ReferenceID: "blue", Theme: theme.New(theme.Dark),
		MinRequests: 1000, Sigma: 3, Variant: pattern.Bands,
	})
	if err != nil {
		return "", err
	}

	comparison.Resize(layout.Size{Width: 100, Height: 30})

	return comparison.View(), nil
}
