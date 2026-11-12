package cohort

import "github.com/lanej/charmed/attention"

// ComparisonPolicy uses host z(ref) boundaries and request-weighted fleet delta
// boundaries in percentage points. These are descriptive attention policies.
// Either scope, severity or side may be omitted. The zero value is unconfigured.
type (
	ComparisonPolicy    struct{ Hosts, Fleet attention.Rule }
	ComparisonAttention struct {
		Fleet                     attention.Result
		Members                   map[MemberKey]attention.Result
		Hosts, Level              attention.Level
		Warnings, Errors, Unknown int
	}
)

func (p ComparisonPolicy) Validate() error {
	if err := p.Hosts.Validate(); err != nil {
		return err
	}
	return p.Fleet.Validate()
}

func (p ComparisonPolicy) Clone() ComparisonPolicy {
	return ComparisonPolicy{Hosts: p.Hosts.Clone(), Fleet: p.Fleet.Clone()}
}
func (p ComparisonPolicy) Configured() bool { return p.Hosts.Configured() || p.Fleet.Configured() }

func Assess(p ComparisonPolicy, reference, candidate Summary, minRequests uint64) ComparisonAttention {
	return AssessMetric(p, reference, candidate, Metric{}, minRequests)
}

// AssessMetric evaluates member distances and aggregate deltas in the metric's
// display unit (percentage points for ratios). A missing reference yields Unknown
// for configured policies. Hosts/Fleet retain their legacy names for compatibility.
func AssessMetric(p ComparisonPolicy, reference, candidate Summary, metric Metric, minSupport uint64) ComparisonAttention {
	a := ComparisonAttention{Members: map[MemberKey]attention.Result{}, Hosts: attention.Off}

	available := reference.HasRate && candidate.HasRate
	if metric.Kind == Ratio {
		available = available && reference.Requests >= float64(max(1, minSupport)) && candidate.Requests >= float64(max(1, minSupport))
	}

	a.Fleet = p.Fleet.Evaluate((candidate.Rate-reference.Rate)*metric.Scale(), available)
	for _, summary := range []Summary{reference, candidate} {
		for _, member := range summary.Points {
			z, ok := Distance(member, reference)
			r := p.Hosts.Evaluate(z, ok)
			a.Members[member.MemberKey] = r
			a.Hosts = attention.Worst(a.Hosts, r.Level)
			switch r.Level {
			case attention.Warn:
				a.Warnings++
			case attention.Error:
				a.Errors++
			case attention.Unknown:
				a.Unknown++
			}
		}
	}
	if len(a.Members) == 0 && p.Hosts.Configured() {
		a.Hosts = attention.Unknown
	}
	a.Level = attention.Worst(a.Fleet.Level, a.Hosts)
	return a
}
