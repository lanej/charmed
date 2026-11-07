package fixture

import (
	"github.com/lanej/charmed/attention"
	"github.com/lanej/charmed/cohort"
)

// Boundaries are illustrative application policy, never library defaults.
// Falling success and rising server-error rates use opposite one-sided limits.
func Boundaries() map[string]cohort.ComparisonPolicy {
	return map[string]cohort.ComparisonPolicy{
		"2xx": {Hosts: attention.Rule{Warn: &attention.Bounds{Lower: attention.Limit(-2)}, Error: &attention.Bounds{Lower: attention.Limit(-3)}}, Fleet: attention.Rule{Warn: &attention.Bounds{Lower: attention.Limit(-0.1)}, Error: &attention.Bounds{Lower: attention.Limit(-0.3)}}},
		"5xx": {Hosts: attention.Rule{Warn: &attention.Bounds{Upper: attention.Limit(2)}, Error: &attention.Bounds{Upper: attention.Limit(3)}}, Fleet: attention.Rule{Warn: &attention.Bounds{Upper: attention.Limit(0.1)}, Error: &attention.Bounds{Upper: attention.Limit(0.3)}}},
	}
}
