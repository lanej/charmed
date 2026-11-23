// Package fixture supplies synthetic operation evidence without service access.
package fixture

import (
	"strings"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

type observation struct {
	id, group, title, condition, next, source, detail string
	tone                                              text.Tone
	context                                           []string
}

// Records keeps source meaning in the consumer. Charmed receives labels and
// renderers; it neither checks the service nor decides whether an effect succeeded.
func Records(state string, t theme.Theme) []pattern.InspectorRecord {
	if state == "empty" || state == "loading" {
		return nil
	}

	observations := []observation{
		{
			id: "route-2026-11-23-0042", group: "Network", title: "Gateway route update", condition: "Outcome unknown", tone: text.Warning,
			next: "Inspect the saved route request before changing traffic.", source: "Controller journal / last observed 09:42 UTC",
			detail:  "The controller recorded a request to switch the gateway route. Its response timed out after submission. The remote route has not been observed since the request.",
			context: []string{"Response lost after submission", "Next: inspect remote evidence"},
		},
		{
			id: "cache-2026-11-23-0018", group: "Storage", title: "Cache warm-up", condition: "In progress", tone: text.Info,
			next: "Wait for the final cache sample.", source: "Cache observer / sampled 09:42 UTC",
			detail: "The observer reports 18 of 24 partitions warm. This is a partial sample, not a completed warm-up.",
		},
		{
			id: "pool-2026-11-23-0007", group: "Compute", title: "Worker capacity", condition: "Outcome confirmed", tone: text.Success,
			next: "Compare queue latency with the service objective.", source: "Pool observer / sampled 09:41 UTC",
			detail:  "The pool observer independently read the requested capacity of twelve workers. Capacity confirmation does not establish healthy workers or acceptable queue latency.",
			context: []string{"12 workers observed"},
		},
		{
			id: "index-2026-11-23-0031", group: "Search", title: "Indexer restart", condition: "Success reported", tone: text.Neutral,
			next: "Read the process observer before declaring recovery.", source: "Operator report / recorded 09:40 UTC",
			detail:  "The operator reported a successful restart. No independent process observation accompanies that report; the service's recovery remains unverified.",
			context: []string{"Caller report only", "No independent process sample"},
		},
	}

	if state == "confirmed" {
		observations[0].condition, observations[0].tone = "Outcome confirmed", text.Success
		observations[0].source = "Route observer / sampled 09:43 UTC"
		observations[0].next = "Observe request errors before closing the incident."
		observations[0].detail = "A later read found the requested gateway route active. The route outcome is confirmed; the lost response still prevents attributing the change to this exact request."
		observations[0].context = []string{"Requested route independently observed", "Request attribution remains unknown"}
	}

	records := make([]pattern.InspectorRecord, len(observations))
	for i, observed := range observations {
		records[i] = pattern.InspectorRecord{
			ID: observed.id, Heading: observed.group,
			Row: func(width int, selected bool) string {
				style := t.Text()
				if selected {
					style = t.Heading()
				}

				lines := []string{style.Render(observed.title), text.Pill(observed.condition, observed.tone, t, text.Normal)}
				for _, context := range observed.context {
					lines = append(lines, t.Subtle().Render(context))
				}

				return ansi.Hardwrap(strings.Join(lines, "\n"), max(1, width), true)
			},
			Detail: func(width int) string {
				parts := []string{
					t.Heading().Render(observed.title) + "\n" + text.Pill(observed.condition, observed.tone, t, text.Normal),
					layout.KeyValue([]layout.Pair{{Label: "Operation", Value: observed.id}, {Label: "Source", Value: observed.source}}, width, t),
					t.Heading().Render("Observed evidence") + "\n" + observed.detail,
					t.Heading().Render("Next step") + "\n" + observed.next,
					t.Heading().Render("Interpretation limits") + "\nA recorded request establishes intent. A caller's report establishes what the caller said. An independent observation may confirm the external state; it does not automatically prove request attribution, service health or incident resolution.",
					t.Heading().Render("Inspection notes") + "\nCompare the saved identity with the observer's evidence.\nKeep the last observation when a refresh fails.\nCheck the observation time before taking another action.\nRecord any remaining attribution gap.\nChoose the next action in the application.",
				}
				if state == "error" {
					parts = append([]string{text.Pill("Refresh failed", text.Warning, t, text.Normal) + " Last observation retained."}, parts...)
				}

				return ansi.Wrap(strings.Join(parts, "\n\n"), max(1, width), " ")
			},
		}
	}

	return records
}
