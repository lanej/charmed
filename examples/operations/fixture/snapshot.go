// Package fixture supplies deterministic operational observations, with no API access.
package fixture

import (
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func Snapshot(state string, t theme.Theme) pattern.DashboardSnapshot {
	buckets := chart.BucketBars{Buckets: []chart.CountBucket{
		{Label: "00:00", Count: 0, State: chart.BucketPartial},
		{Label: "23:00", Count: 173, State: chart.BucketComplete},
		{Label: "22:00", Count: 160, State: chart.BucketComplete},
		{Label: "21:00", Count: 147, State: chart.BucketComplete},
		{Label: "20:00", Count: 134, State: chart.BucketComplete},
		{Label: "19:00", Count: 121, State: chart.BucketComplete},
	}}
	snapshot := pattern.DashboardSnapshot{
		Title: "API operations", Context: "Synthetic / 14-15 Nov 2026",
		Status: "Observed 00:16 UTC / resets 01:00",
		Meters: []pattern.DashboardMeter{
			{Label: "REST remaining", Value: 3500, Total: 5000, Unit: "requests"},
			{Label: "GraphQL remaining", Value: 4800, Total: 5000, Unit: "points"},
		},
		Sections: []pattern.DashboardSection{
			{Title: "REST attempts / UTC hour", Body: func(width int) string { return buckets.View(width, t) }},
			{Title: "Coverage", Body: func(int) string {
				return "Complete hourly counts share a scale. Current hour is partial.\nHistory begins at coordinator start; earlier hours are unavailable."
			}},
		},
		Details: []pattern.DashboardSection{
			{Title: "Scheduler", Body: func(width int) string {
				return layout.KeyValue([]layout.Pair{
					{Label: "Cadence", Value: "1s tick / 2s minimum request gap"},
					{Label: "Queue", Value: "0 waiting / idle"},
					{Label: "Dispatch", Value: "Awaiting work"},
				}, width, t)
			}},
			{Title: "Accounting", Body: func(int) string {
				return "HTTP attempts include transport failures; Git transport is excluded.\nQuotas cover the shared credential, including other tools.\nHistory resets on coordinator restart; no historical backfill."
			}},
		},
	}

	switch state {
	case "empty", "loading":
		snapshot.Status = "No observations received"
		if state == "loading" {
			snapshot.Status = "Waiting for first observation"
		}

		for i := range snapshot.Meters {
			snapshot.Meters[i].State = text.ValueMissing
		}

		snapshot.Sections, snapshot.Details = nil, nil
	case "error":
		snapshot.Status = "Refresh failed / showing last observation from 00:16 UTC"
		snapshot.StatusTone = text.Warning
	}

	return snapshot
}
