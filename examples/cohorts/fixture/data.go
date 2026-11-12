// Package fixture supplies synthetic regional measurements for the generic organism.
package fixture

import (
	"fmt"
	"time"

	"github.com/lanej/charmed/cohort"
)

func Data() cohort.Data {
	data := cohort.Data{
		Cohorts: []cohort.Cohort{{ID: "north", Label: "North", Short: "N"}, {ID: "south", Label: "South", Short: "S"}, {ID: "west", Label: "West", Short: "W"}},
		Metrics: []cohort.Metric{{ID: "temp", Label: "Temperature", Kind: cohort.Scalar, Unit: "C"}, {ID: "energy", Label: "Energy", Kind: cohort.Scalar, Unit: "kWh", Aggregation: cohort.Sum}, {ID: "yield", Label: "Yield"}},
	}

	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	for window := range 6 {
		w := cohort.Window{Start: start.Add(time.Duration(window) * time.Hour), End: start.Add(time.Duration(window+1) * time.Hour)}

		for c, group := range data.Cohorts {
			for member := range 4 {
				measurements := map[string]cohort.Measurement{
					"temp":   {Value: float64(c*8 + member + window - 5)},
					"energy": {Value: float64(10 + c*5 + member + window)},
					"yield":  {Numerator: uint64(70 + c*5 + member + window), Denominator: 100},
				}
				if c == 2 && member == 3 {
					delete(measurements, "temp")
				}

				w.Samples = append(w.Samples, cohort.Sample{MemberKey: cohort.MemberKey{CohortID: group.ID, MemberID: fmt.Sprintf("sensor-%d", member+1)}, Measurements: measurements})
			}
		}

		data.Windows = append(data.Windows, w)
	}

	return data
}
