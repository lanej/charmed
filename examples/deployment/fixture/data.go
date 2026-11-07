// Package fixture supplies deterministic synthetic deployment counts. It is an
// example data source; the UI packages have no deployment or telemetry dependency.
package fixture

import (
	"fmt"
	"math"
	"slices"
	"time"

	"github.com/lanej/charmed/cohort"
)

var Scenarios = []string{"isolated-hosts", "uniform-regression", "unequal-traffic", "insufficient-samples"}

func Data(scenario string) (cohort.Data, error) {
	if !slices.Contains(Scenarios, scenario) {
		return cohort.Data{}, fmt.Errorf("unknown deployment scenario %q", scenario)
	}
	d := cohort.Data{Cohorts: []cohort.Cohort{{ID: "blue", Label: "Blue v1.8", Short: "B"}, {ID: "green", Label: "Green v1.9", Short: "G"}},
		Metrics: []cohort.Metric{{ID: "2xx", Label: "2xx"}, {ID: "3xx", Label: "3xx"}, {ID: "4xx", Label: "4xx"}, {ID: "5xx", Label: "5xx"}}}
	start := time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)
	for bucket := range 12 {
		w := cohort.Window{Start: start.Add(time.Duration(bucket) * 5 * time.Minute), End: start.Add(time.Duration(bucket+1) * 5 * time.Minute)}
		for c, group := range d.Cohorts {
			hosts := 64
			if c == 1 {
				hosts = 32
			}
			for host := range hosts {
				requests := uint64(20000 + (host*7919+bucket*233)%60000)
				if scenario == "unequal-traffic" {
					// Host rates match, but green routes most traffic to the
					// naturally higher-error third of the population.
					requests = 80000
					if host%3 == 0 {
						requests = 8000
					}
					if c == 1 {
						requests = 8000
						if host%3 == 0 {
							requests = 80000
						}
					}
				}
				if scenario == "insufficient-samples" {
					requests = uint64(80 + (host*17+bucket*3)%400)
					if c == 0 && host == 0 {
						requests = 50000
					}
					if host%9 == 0 && host > 0 {
						requests = 0
					}
				}
				phase := float64(host)*0.7 + float64(bucket)*0.21
				redirect := 0.021 + 0.0015*math.Sin(phase*0.6)
				client := 0.016 + 0.001*math.Cos(phase)
				server := 0.002 + 0.00035*math.Sin(phase)
				if scenario == "unequal-traffic" {
					server = 0.001
					if host%3 == 0 {
						server = 0.012
					}
				}
				if c == 1 && bucket >= 4 {
					ramp := min(1, float64(bucket-3)/3)
					if scenario == "uniform-regression" {
						server += 0.004 * ramp
					}
					if scenario == "isolated-hosts" && host >= 27 {
						server += (0.010 + float64(host-27)*0.002) * ramp
					}
				}
				counts := map[string]uint64{"3xx": uint64(math.Round(redirect * float64(requests))), "4xx": uint64(math.Round(client * float64(requests))), "5xx": uint64(math.Round(server * float64(requests)))}
				counts["2xx"] = requests - counts["3xx"] - counts["4xx"] - counts["5xx"]
				w.Samples = append(w.Samples, cohort.Sample{MemberKey: cohort.MemberKey{CohortID: group.ID, MemberID: fmt.Sprintf("%s-%03d", group.Short, host+1)}, Requests: requests, Counts: counts})
			}
		}
		d.Windows = append(d.Windows, w)
	}
	return d, nil
}
