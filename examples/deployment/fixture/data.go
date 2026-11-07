// Package fixture supplies deterministic synthetic deployment counts. It is an
// example data source; the UI packages have no deployment or telemetry dependency.
package fixture

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/lanej/charmed/cohort"
)

var Scenarios = []string{"isolated-hosts", "uniform-regression", "unequal-traffic", "insufficient-samples"}

const ReferenceVersionID = "a13c9f2"
const CandidateVersionID = "f4d7e81"

// Versions are opaque identifiers, independent of the deployment method,
// environment names, comparison roles and visual series aliases.
type Versions struct{ Reference, Candidate string }

func (v Versions) Validate() error {
	for _, id := range []string{v.Reference, v.Candidate} {
		if strings.TrimSpace(id) == "" || len(id) > 128 || !utf8.ValidString(id) || strings.ContainsFunc(id, unicode.IsControl) {
			return fmt.Errorf("version identifiers must be nonempty printable text of at most 128 bytes")
		}
	}
	if v.Reference == v.Candidate {
		return fmt.Errorf("reference and candidate need different version identifiers")
	}
	return nil
}

func Data(scenario string) (cohort.Data, error) {
	return DataWithVersions(scenario, Versions{Reference: ReferenceVersionID, Candidate: CandidateVersionID})
}

func DataWithVersions(scenario string, versions Versions) (cohort.Data, error) {
	if err := versions.Validate(); err != nil {
		return cohort.Data{}, err
	}
	if !slices.Contains(Scenarios, scenario) {
		return cohort.Data{}, fmt.Errorf("unknown deployment scenario %q", scenario)
	}
	d := cohort.Data{Cohorts: []cohort.Cohort{{ID: versions.Reference, Label: versions.Reference, Short: "R"}, {ID: versions.Candidate, Label: versions.Candidate, Short: "C"}},
		Metrics: []cohort.Metric{{ID: "2xx", Label: "2xx Success"}, {ID: "3xx", Label: "3xx Redirects"}, {ID: "4xx", Label: "4xx Client error"}, {ID: "5xx", Label: "5xx Server error"}}}
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
					// Host rates match, but the candidate routes most traffic to the
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
				w.Samples = append(w.Samples, cohort.Sample{MemberKey: cohort.MemberKey{CohortID: group.ID, MemberID: fmt.Sprintf("host-%03d", c*64+host+1)}, Requests: requests, Counts: counts})
			}
		}
		d.Windows = append(d.Windows, w)
	}
	return d, nil
}
