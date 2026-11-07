package fixture_test

import (
	"math"
	"testing"

	"github.com/lanej/charmed/cohort"
	"github.com/lanej/charmed/examples/deployment/fixture"
)

func TestScenariosExposeDifferentFailureShapes(t *testing.T) {
	for _, scenario := range fixture.Scenarios {
		t.Run(scenario, func(t *testing.T) {
			d, err := fixture.Data(scenario)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := d.Clone(); err != nil {
				t.Fatal(err)
			}
			for _, w := range d.Windows {
				for _, s := range w.Samples {
					var sum uint64
					for _, n := range s.Counts {
						sum += n
					}
					if sum != s.Requests {
						t.Fatal("HTTP classes do not partition responses")
					}
				}
			}
			last := d.Windows[len(d.Windows)-1]
			reference, candidate := cohort.Analyze(last.Samples, fixture.ReferenceVersionID, "5xx", 1000), cohort.Analyze(last.Samples, fixture.CandidateVersionID, "5xx", 1000)
			switch scenario {
			case "uniform-regression":
				if candidate.Mean-reference.Mean < 0.0035 || candidate.StdDev > 0.0005 {
					t.Fatal("fixture is not a uniform shift")
				}
			case "isolated-hosts":
				bad := 0
				for _, p := range candidate.Points {
					if z, ok := cohort.Distance(p, reference); ok && z > 3 {
						bad++
					}
				}
				if bad != 5 || candidate.StdDev < reference.StdDev*10 {
					t.Fatalf("expected five isolated regressions, got %d", bad)
				}
			case "unequal-traffic":
				if math.Abs(reference.Mean-candidate.Mean) > 0.0002 || candidate.Rate-reference.Rate < 0.005 {
					t.Fatal("fixture does not distinguish member mean from traffic-weighted impact")
				}
			case "insufficient-samples":
				if reference.Eligible != 1 || candidate.Eligible != 0 || reference.HasSpread || candidate.HasSpread {
					t.Fatal("sparse hosts accidentally establish spread")
				}
			}
		})
	}
}

func TestVersionIdentifiersAreIndependentOfDeploymentNames(t *testing.T) {
	versions := fixture.Versions{Reference: "release/2026.11.07+build.42", Candidate: "sha256:731a9e4"}
	d, err := fixture.DataWithVersions("isolated-hosts", versions)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := d.Clone(); err != nil {
		t.Fatal(err)
	}
	for _, w := range d.Windows {
		for _, s := range w.Samples {
			if s.CohortID != versions.Reference && s.CohortID != versions.Candidate {
				t.Fatalf("sample keyed by something other than version ID: %s", s.CohortID)
			}
		}
	}
	if d.Cohorts[0].ID != versions.Reference || d.Cohorts[1].ID != versions.Candidate {
		t.Fatal("comparison roles replaced actual identifiers")
	}
	if d.Windows[0].Samples[0].MemberID != "host-001" {
		t.Fatal("host identity depends on version or pool")
	}
	for _, v := range []fixture.Versions{{Reference: "", Candidate: "b"}, {Reference: "a", Candidate: "a"}, {Reference: "a\nb", Candidate: "c"}} {
		if _, err := fixture.DataWithVersions("isolated-hosts", v); err == nil {
			t.Fatal("invalid version identifiers accepted")
		}
	}
}
