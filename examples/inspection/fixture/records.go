// Package fixture supplies illustrative delivery records without service access.
package fixture

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

type delivery struct {
	id, group, title, condition, observed, next, reason, revision, owner string
	tone                                                                 text.Tone
	verification, feedback                                               []layout.StatusItem
	activity                                                             []layout.TimelineEvent
}

// Records exposes a deterministic queue for the catalog and interactive example.
// Only a later observation changes the selected pull request's merge outcome.
func Records(state string, t theme.Theme) []pattern.InspectorRecord {
	if state == "empty" || state == "loading" {
		return nil
	}

	deliveries := []delivery{
		routeUpdate(state, t),
		{
			id: "example/service#43", group: "Needs attention", title: "PR #43 · Cache limits",
			condition: "Feedback open", tone: text.Info, observed: "09:40 UTC",
			next: "Review the requested change", reason: "Worker 1 acknowledged the review and has not recorded a disposition.",
			revision: "bf8e219", owner: "Worker 1",
			feedback: []layout.StatusItem{
				{Label: "Transport", Value: "Delivered"}, {Label: "Receipt", Value: "Acknowledged"}, {Label: "Disposition", Value: "Open"},
			},
			activity: []layout.TimelineEvent{{When: "09:40 UTC", Actor: "Worker 1", Title: "Review acknowledged"}},
		},
		{
			id: "example/service#44", group: "In progress", title: "PR #44 · Retry interval",
			condition: "Tests running", tone: text.Info, observed: "09:39 UTC",
			next: "Wait for the integration suite", reason: "The unit suite passed. Integration tests are still running for this revision.",
			revision: "21bef30", owner: "Worker 3",
			verification: []layout.StatusItem{
				{Label: "Unit tests", Value: "Passed"}, {Label: "Integration", Value: "Running"},
			},
			activity: []layout.TimelineEvent{{When: "09:39 UTC", Actor: "CI", Title: "Integration suite started"}},
		},
		{
			id: "example/service#41", group: "Completed", title: "PR #41 · Request parsing",
			condition: "Merged", tone: text.Success, observed: "09:38 UTC", revision: "ce82d70", owner: "Worker 1",
			verification: []layout.StatusItem{{Label: "Merge", Value: "Confirmed"}},
			activity:     []layout.TimelineEvent{{When: "09:38 UTC", Actor: "GitHub", Title: "Merge confirmed"}},
		},
	}

	records := make([]pattern.InspectorRecord, 0, len(deliveries))
	for _, pr := range deliveries {
		records = append(records, pr.record(state, t))
	}

	return records
}

func routeUpdate(state string, t theme.Theme) delivery {
	label := func(value string, glyph theme.Glyph, tone text.Tone) string {
		return text.ToneStyle(tone, text.Normal, t).Render(t.Glyph(glyph) + " " + value)
	}

	pr := delivery{
		id: "example/service#42", group: "Needs attention", title: "PR #42 · Route update",
		condition: "Outcome uncertain", tone: text.Warning, observed: "09:42 UTC",
		next: "Reconcile the saved merge operation", reason: "The request was recorded, but the remote response was unavailable.",
		revision: "a81f39d", owner: "Worker 2",
		verification: []layout.StatusItem{
			{Label: "Tests", Value: label("Passed", theme.GlyphOK, text.Success)},
			{Label: "Review", Value: label("Complete", theme.GlyphOK, text.Success)},
			{Label: "Merge", Value: label("Unverified", theme.GlyphUnknown, text.Warning)},
		},
		feedback: []layout.StatusItem{
			{Label: "Source", Value: label("Recorded", theme.GlyphOK, text.Info), Detail: "Review comment"},
			{Label: "Transport", Value: label("Delivered", theme.GlyphOK, text.Success), Detail: "To Worker 2"},
			{Label: "Receipt", Value: label("Unconfirmed", theme.GlyphUnknown, text.Warning), Detail: "No acknowledgement"},
			{Label: "Disposition", Value: label("Open", theme.GlyphPending, text.Info), Detail: "Awaiting decision"},
		},
		activity: []layout.TimelineEvent{
			{When: "09:41 UTC", Title: "Merge requested", Detail: "Request persisted"},
			{When: "09:42 UTC", Title: "Remote response unavailable", Detail: "External outcome unknown"},
		},
	}

	if state == "confirmed" {
		pr.group, pr.condition, pr.tone = "Completed", "Merge confirmed", text.Success
		pr.observed, pr.next, pr.reason = "09:43 UTC", "", ""
		pr.verification[2].Value = label("Confirmed", theme.GlyphOK, text.Success)
		pr.activity = append(pr.activity, layout.TimelineEvent{When: "09:43 UTC", Actor: "GitHub", Title: "Merge confirmed", Detail: "The pull request is merged at a81f39d."})
	}

	return pr
}

func (pr delivery) record(state string, t theme.Theme) pattern.InspectorRecord {
	summary := layout.RecordSummary{
		Title: pr.title, Status: text.Pill(pr.condition, pr.tone, t, text.Normal),
		Fields: []layout.Pair{{Label: "Observed", Value: pr.observed}},
	}
	if pr.next != "" {
		summary.Next = &layout.Pair{Label: "Next action", Value: pr.next}
	}

	document := layout.DetailDocument{
		MinSectionWidth: 32,
		Header:          layout.IdentityHeader{ID: "DELIVERY INSPECTION · ILLUSTRATIVE"}, Summary: summary, ActionDetail: pr.reason,
		Sections: []layout.DetailSection{{ID: "identity", Title: "Identity", Fields: []layout.Pair{
			{Label: "Repository", Value: "example/service"}, {Label: "Revision", Value: pr.revision}, {Label: "Owner", Value: pr.owner},
		}}},
	}
	if len(pr.verification) > 0 {
		document.Sections = append(document.Sections, layout.DetailSection{ID: "verification", Title: "Verification", Statuses: pr.verification})
	}

	if len(pr.feedback) > 0 {
		document.Sections = append(document.Sections, layout.DetailSection{ID: "feedback", Title: "Feedback delivery", Stages: pr.feedback, FullWidth: true})
	}

	document.Sections = append(document.Sections, layout.DetailSection{ID: "activity", Title: "Activity", Timeline: pr.activity, FullWidth: true})
	if state == "error" {
		document.Summary.Subtitle = text.Pill("Refresh failed", text.Warning, t, text.Normal) + " Last observation retained."
	}

	return pattern.InspectorRecord{ID: pr.id, Heading: pr.group, Summary: &summary, Document: &document}
}
