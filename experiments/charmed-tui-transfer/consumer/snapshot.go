package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func count(value *uint64) string {
	if value == nil {
		return "unknown"
	}

	return strconv.FormatUint(*value, 10)
}

func timestamp(value *time.Time) string {
	if value == nil {
		return "unknown / never supplied"
	}

	return value.UTC().Format("2006-01-02 15:04:05 UTC")
}

func section(title, body string) pattern.DashboardSection {
	return pattern.DashboardSection{Title: title, Body: func(int) string { return body }}
}

func (f fixture) snapshot(index int, t theme.Theme) pattern.DashboardSnapshot {
	s := f.States[index]
	result := pattern.DashboardSnapshot{
		Title:   f.Service,
		Context: "Synthetic / UTC / " + s.ID,
	}

	result.Details = append(result.Details, section("Observation context", f.Environment+"\nFixed reference: "+f.Clock.UTC().Format(time.RFC3339)))
	for _, name := range []string{"schedule", "concurrency_limit", "retention_days", "source", "accounting", "scope", "operator_note", "ledger_note", "replacement_note"} {
		raw, ok := f.Diagnostics[name]
		if !ok {
			continue
		}

		var value string
		if err := json.Unmarshal(raw, &value); err != nil {
			value = string(raw)
		}

		title := strings.ReplaceAll(name, "_", " ")
		if name == "retention_days" {
			value += " days"
		}

		result.Details = append(result.Details, section(title, value))
	}

	obs := s.Observation
	if obs == nil {
		result.Status = "Waiting for first observation / refresh " + s.Refresh.State
		result.Sections = []pattern.DashboardSection{section("No observation yet", "Counts, intervals and regions are unavailable. This is not a confirmed empty result.")}

		return result
	}

	age := (feedback.Freshness{ObservedAt: obs.ObservedAt}).Label(f.Clock)

	result.Status = "Observed " + obs.ObservedAt.UTC().Format("15:04 UTC") + " / " + age
	if s.Refresh.State == "failed" {
		result.Status = "Refresh failed / retained " + obs.ObservedAt.UTC().Format("15:04 UTC") + " / " + age
		result.StatusTone = text.Warning
		result.Sections = append(result.Sections, section("Refresh failure", s.Refresh.Error+"\nLast observation retained; no successful refresh confirmed."))
	}

	result.Summary = []layout.Pair{{Label: "Finished / due", Value: count(obs.CompletedChecks) + " / " + count(obs.ChecksDue) + " checks"}}
	bars := chart.BucketBars{}

	for _, item := range obs.Intervals {
		bucket := chart.CountBucket{Label: item.Start.UTC().Format("15:04") + "-" + item.End.UTC().Format("15:04")}
		switch item.Coverage {
		case "complete":
			bucket.State = chart.BucketComplete
		case "partial":
			bucket.State = chart.BucketPartial
		case "missing":
			bucket.State = chart.BucketMissing
		case "unknown":
			bucket.State = chart.BucketUnknown
		default:
			bucket.State = chart.BucketState("invalid")
		}

		if item.CompletedChecks != nil {
			bucket.Count = *item.CompletedChecks
		} else if item.Coverage == "complete" || item.Coverage == "partial" {
			bucket.State = chart.BucketUnknown
		}

		bars.Buckets = append(bars.Buckets, bucket)
	}

	result.Sections = append(result.Sections, section("Completion is not health", "Finished counts include passing and failing verification attempts."))
	if len(obs.Intervals) == 0 {
		result.Sections = append(result.Sections, section("Intervals", "Confirmed empty: no interval observations."))
	} else {
		title := "Finished checks / 15-minute interval (UTC)"
		result.Sections = append(result.Sections, pattern.DashboardSection{Title: title, Body: func(width int) string { return bars.View(width, t) }})
	}

	if len(obs.Regions) == 0 {
		result.Sections = append(result.Sections, section("Regions", "Confirmed empty: no region observations."))
	} else {
		var (
			rows    []string
			details []pattern.DashboardSection
		)

		for _, r := range obs.Regions {
			rows = append(rows, r.Name+" / "+r.Availability+" / "+r.Health)
			details = append(details, section(r.Name, "Availability: "+r.Availability+"\nHealth: "+r.Health+"\nLast verified: "+timestamp(r.LastVerified)))
		}

		result.Sections = append(result.Sections, section("Regions / availability / health", strings.Join(rows, "\n")+"\ni: exact verification times and diagnostics"))
		result.Details = append(details, result.Details...)
	}

	result.Details = append(result.Details, section("Refresh evidence", fmt.Sprintf("State: %s\nRequested: %s\nConfirmed: %s\nObserved: %s", s.Refresh.State, timestamp(s.Refresh.RequestedAt), timestamp(s.Refresh.ConfirmedAt), timestamp(&obs.ObservedAt))))

	return result
}
