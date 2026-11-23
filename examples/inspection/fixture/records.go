// Package fixture supplies synthetic operation evidence without service access.
package fixture

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

type observation struct {
	id, group, title, condition, next, source, observed, detail string
	tone                                                        text.Tone
	verification                                                []layout.StatusItem
	history                                                     []layout.TimelineEvent
}

// Records supplies structured presentation while keeping the interpretation of
// observations and the choice of next action in the application.
func Records(state string, t theme.Theme) []pattern.InspectorRecord {
	if state == "empty" || state == "loading" {
		return nil
	}

	observations := []observation{
		{
			id: "route-2026-11-23-0042", group: "Network", title: "Gateway route update", condition: "Outcome unknown", tone: text.Warning,
			next: "Inspect the saved route request before changing traffic.", source: "Controller journal", observed: "09:42 UTC",
			detail: "The controller recorded a request to switch the gateway route. Its response timed out after submission. The remote route has not been observed since the request.",
			verification: []layout.StatusItem{
				{Label: "Submission", Value: "Recorded", Detail: "The request reached the controller journal."},
				{Label: "Remote route", Value: "Not observed", Detail: "No read follows the submitted request."},
				{Label: "Attribution", Value: "Unknown", Detail: "The lost response cannot establish the effect."},
			},
			history: []layout.TimelineEvent{
				{When: "09:41 UTC", Actor: "Controller", Title: "Route request recorded", Detail: "Switch the gateway route to the prepared target."},
				{When: "09:42 UTC", Actor: "Controller", Title: "Response timed out", Detail: "Submission is recorded; the external result is still unknown."},
			},
		},
		{
			id: "cache-2026-11-23-0018", group: "Storage", title: "Cache warm-up", condition: "In progress", tone: text.Info,
			next: "Wait for the final cache sample.", source: "Cache observer", observed: "09:42 UTC",
			detail: "The observer reports 18 of 24 partitions warm. This is a partial sample, not a completed warm-up.",
			verification: []layout.StatusItem{
				{Label: "Partitions", Value: "18 of 24 warm", Detail: "The observer's sample covers the current warm-up."},
				{Label: "Completion", Value: "Not confirmed", Detail: "A final sample has not been captured."},
			},
			history: []layout.TimelineEvent{{When: "09:42 UTC", Actor: "Cache observer", Title: "Partial sample captured", Detail: "Six partitions remain outside the confirmed warm set."}},
		},
		{
			id: "pool-2026-11-23-0007", group: "Compute", title: "Worker capacity", condition: "Outcome confirmed", tone: text.Success,
			next: "Compare queue latency with the service objective.", source: "Pool observer", observed: "09:41 UTC",
			detail: "The pool observer independently read the requested capacity of twelve workers. Capacity confirmation does not establish healthy workers or acceptable queue latency.",
			verification: []layout.StatusItem{
				{Label: "Capacity", Value: "12 workers observed", Detail: "The observed count matches the requested capacity."},
				{Label: "Queue latency", Value: "Not measured", Detail: "Capacity alone cannot establish the service objective."},
			},
			history: []layout.TimelineEvent{{When: "09:41 UTC", Actor: "Pool observer", Title: "Capacity independently confirmed", Detail: "Twelve workers are present in the observed pool."}},
		},
		{
			id: "index-2026-11-23-0031", group: "Search", title: "Indexer restart", condition: "Success reported", tone: text.Neutral,
			next: "Read the process observer before declaring recovery.", source: "Operator report", observed: "09:40 UTC",
			detail: "The operator reported a successful restart. No independent process observation accompanies that report; the service's recovery remains unverified.",
			verification: []layout.StatusItem{
				{Label: "Caller report", Value: "Success reported", Detail: "This records what the operator said."},
				{Label: "Recovery", Value: "Not confirmed", Detail: "No independent process sample was supplied."},
			},
			history: []layout.TimelineEvent{{When: "09:40 UTC", Actor: "Operator", Title: "Restart reported successful", Detail: "The report includes no independent process observation."}},
		},
	}

	if state == "confirmed" {
		observations[0].condition, observations[0].tone = "Outcome confirmed", text.Success
		observations[0].source, observations[0].observed = "Route observer", "09:43 UTC"
		observations[0].next = "Observe request errors before closing the incident."
		observations[0].detail = "A later read found the requested gateway route active. The route outcome is confirmed; the lost response still prevents attributing the change to this exact request."
		observations[0].verification[1] = layout.StatusItem{Label: "Remote route", Value: "Requested route active", Detail: "Independently observed at 09:43 UTC."}
		observations[0].history = append(observations[0].history, layout.TimelineEvent{When: "09:43 UTC", Actor: "Route observer", Title: "Requested route observed", Detail: "The external state is confirmed; attribution to this request remains unknown."})
	}

	records := make([]pattern.InspectorRecord, len(observations))
	for i, observed := range observations {
		summary := layout.RecordSummary{
			Title: observed.title, Status: text.Pill(observed.condition, observed.tone, t, text.Normal),
			Fields: []layout.Pair{{Label: "Observed", Value: observed.observed}},
			Next:   &layout.Pair{Label: "Next", Value: observed.next},
		}
		document := layout.DetailDocument{
			Header:  layout.IdentityHeader{Kind: "Operation", ID: observed.id},
			Summary: summary,
			Sections: []layout.DetailSection{
				{ID: "observation", Title: "Observation", Fields: []layout.Pair{
					{Label: "Domain", Value: observed.group},
					{Label: "Source", Value: observed.source},
					{Label: "Observed", Value: observed.observed},
					{Label: "Request receipt", State: text.ValueMissing},
				}},
				{ID: "verification", Title: "Verification", Statuses: observed.verification},
				{ID: "evidence", Title: "Observed evidence", Paragraphs: []string{observed.detail}, FullWidth: true},
				{ID: "history", Title: "Activity", Timeline: observed.history, FullWidth: true},
				{ID: "limits", Title: "Interpretation limits", Paragraphs: []string{
					"A recorded request establishes intent. A caller's report establishes what the caller said. An independent observation may confirm the external state; it does not automatically prove request attribution, service health or incident resolution.",
				}},
				{ID: "notes", Title: "Inspection notes", Paragraphs: []string{
					"Compare the saved identity with the observer's evidence. Keep the last observation when a refresh fails. Check the observation time before taking another action. Record any remaining attribution gap.",
				}},
			},
		}

		if state == "error" {
			document.Summary.Subtitle = text.Pill("Refresh failed", text.Warning, t, text.Normal) + " Last observation retained."
		}

		records[i] = pattern.InspectorRecord{ID: observed.id, Heading: observed.group, Summary: &summary, Document: &document}
	}

	return records
}
