package cohort_test

import (
	"testing"

	"github.com/lanej/charmed/attention"
	"github.com/lanej/charmed/cohort"
)

func TestHostSigmaAndFleetPercentagePointScopes(t *testing.T) {
	member := cohort.MemberRate{MemberKey: cohort.MemberKey{CohortID: "candidate", MemberID: "host"}, Requests: 1000, Rate: 0.75, Eligible: true, HasRate: true}
	reference := cohort.Summary{Mean: 0, StdDev: 0.25, HasSpread: true, Rate: 0.25, HasRate: true, Requests: 1000}
	candidate := cohort.Summary{Rate: 0.5, HasRate: true, Requests: 1000, Points: []cohort.MemberRate{member}}
	policy := cohort.ComparisonPolicy{Hosts: attention.Rule{Warn: &attention.Bounds{Upper: attention.Limit(2)}, Error: &attention.Bounds{Upper: attention.Limit(3)}}, Fleet: attention.Rule{Warn: &attention.Bounds{Upper: attention.Limit(10)}, Error: &attention.Bounds{Upper: attention.Limit(25)}}}
	a := cohort.Assess(policy, reference, candidate, 1000)
	if a.Fleet.Level != attention.Error || a.Fleet.Value != 25 || a.Members[member.MemberKey].Value != 3 || a.Errors != 1 {
		t.Fatalf("wrong units or scope: %+v", a)
	}
	reference.StdDev = 0
	a = cohort.Assess(policy, reference, candidate, 1000)
	if a.Hosts != attention.Unknown || a.Fleet.Level != attention.Error || a.Level != attention.Error {
		t.Fatal("known fleet breach or unavailable host spread mishandled")
	}
	candidate.Requests = 20
	if cohort.Assess(policy, reference, candidate, 1000).Fleet.Level != attention.Unknown {
		t.Fatal("low-volume fleet was evaluated")
	}
	if cohort.Assess(cohort.ComparisonPolicy{}, reference, candidate, 1000).Level != attention.Off {
		t.Fatal("no boundaries produced attention")
	}
}
