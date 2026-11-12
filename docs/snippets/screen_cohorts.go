package snippets

import (
	"time"

	"github.com/lanej/charmed/cohort"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

func ExampleScreenCohorts() (string, error) {
	start := time.Date(2026, 11, 12, 12, 0, 0, 0, time.UTC)
	data := cohort.Data{
		Cohorts: []cohort.Cohort{{ID: "north", Label: "North", Short: "N"}, {ID: "south", Label: "South", Short: "S"}, {ID: "west", Label: "West", Short: "W"}},
		Metrics: []cohort.Metric{{ID: "temp", Label: "Temperature", Kind: cohort.Scalar, Unit: "C"}},
		Windows: []cohort.Window{{Start: start, End: start.Add(time.Minute), Samples: []cohort.Sample{
			{MemberKey: cohort.MemberKey{CohortID: "north", MemberID: "sensor-1"}, Measurements: map[string]cohort.Measurement{"temp": {Value: -5}}},
			{MemberKey: cohort.MemberKey{CohortID: "south", MemberID: "sensor-2"}, Measurements: map[string]cohort.Measurement{"temp": {Value: 12}}},
			{MemberKey: cohort.MemberKey{CohortID: "west", MemberID: "sensor-3"}, Measurements: map[string]cohort.Measurement{"temp": {Value: 18}}},
		}}},
	}

	comparison, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Title: "Regional measurements", Data: data, ReferenceID: "north", Theme: theme.New(theme.Dark)})
	if err != nil {
		return "", err
	}

	comparison.Resize(layout.Size{Width: 100, Height: 30})

	return comparison.View(), nil
}
