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
			blue, green := cohort.Analyze(last.Samples, "blue", "5xx", 1000), cohort.Analyze(last.Samples, "green", "5xx", 1000)
			switch scenario {
			case "uniform-regression":
				if green.Mean-blue.Mean < 0.0035 || green.StdDev > 0.0005 {
					t.Fatal("fixture is not a uniform shift")
				}
			case "isolated-hosts":
				bad := 0
				for _, p := range green.Points {
					if z, ok := cohort.Distance(p, blue); ok && z > 3 {
						bad++
					}
				}
				if bad != 5 || green.StdDev < blue.StdDev*10 {
					t.Fatalf("expected five isolated regressions, got %d", bad)
				}
			case "unequal-traffic":
				if math.Abs(blue.Mean-green.Mean) > 0.0002 || green.Rate-blue.Rate < 0.005 {
					t.Fatal("fixture does not distinguish member mean from traffic-weighted impact")
				}
			case "insufficient-samples":
				if blue.Eligible != 1 || green.Eligible != 0 || blue.HasSpread || green.HasSpread {
					t.Fatal("sparse hosts accidentally establish spread")
				}
			}
		})
	}
}
