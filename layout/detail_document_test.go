package layout_test

import (
	"math"
	"reflect"
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestDetailDocumentResponsiveGroupsAndSectionRanges(t *testing.T) {
	palette := theme.New(theme.Mono)
	widths := map[string]int{}
	section := func(id, title, body string, full bool) layout.DetailSection {
		return layout.DetailSection{ID: id, Title: title, FullWidth: full, Body: func(width int) string {
			widths[id] = width
			return body
		}}
	}
	document := layout.DetailDocument{Sections: []layout.DetailSection{
		section("identity", "Identity", "owner\nhead", false),
		section("verification", "Verification", "checks\nreviews\nworkflows\nmergeability", false),
		section("activity", "Activity", "event1\nevent2", true),
		section("transport", "Transport", "request accepted", false),
		section("receipt", "Receipt", "confirmation missing", false),
	}}

	wide := document.Render(84, palette)

	if want := map[string]int{"identity": 40, "verification": 40, "activity": 84, "transport": 40, "receipt": 40}; !reflect.DeepEqual(widths, want) {
		t.Fatalf("sections did not receive their actual wide allocations: %v", widths)
	}

	if got := wide.Sections; len(got) != 5 || got[0].Start != got[1].Start || got[0].End != got[1].End || got[2].Start <= got[1].End || got[3].Start != got[4].Start || got[3].Start <= got[2].End {
		t.Fatalf("paired and full-width section ranges disagree: %+v", got)
	}

	assertDetailRanges(t, wide)
	assertRecordBounds(t, wide.Content, 84)

	narrow := document.Render(80, palette)

	for id, width := range widths {
		if width != 80 {
			t.Fatalf("narrow section %s received %d columns", id, width)
		}
	}

	assertDetailRanges(t, narrow)

	for i, span := range narrow.Sections {
		if i > 0 && span.Start <= narrow.Sections[i-1].End {
			t.Fatalf("narrow sections did not stack in declaration order: %+v", narrow.Sections)
		}
	}

	if document.View(80, palette) != narrow.Content {
		t.Fatal("View disagrees with Render")
	}

	document.MinSectionWidth = math.MaxInt
	if got := document.Render(84, palette); got.Sections[0].Start == got.Sections[1].Start {
		t.Fatal("overflowed minimum unexpectedly paired sections")
	}
}

func TestDetailDocumentRetainsCompleteEvidenceAcrossWidths(t *testing.T) {
	ascii := theme.New(theme.Dark)
	ascii.ASCII = true

	for _, palette := range []theme.Theme{theme.New(theme.Dark), theme.New(theme.Mono), ascii} {
		document := layout.DetailDocument{
			Header:  layout.IdentityHeader{Kind: "Operation", ID: "route-42", Context: "Controller journal"},
			Summary: layout.RecordSummary{Title: "Gateway route", Status: text.Pill("Outcome unknown", text.Warning, palette), Next: &layout.Pair{Label: "Next", Value: "Inspect remote evidence"}},
			Sections: []layout.DetailSection{
				{ID: "identity", Title: "Identity", Fields: []layout.Pair{{Label: "Owner", Value: "netops"}, {Label: "Receipt", State: text.ValueMissing}, {Label: "Empty", Value: ""}}},
				{ID: "verification", Title: "Verification", Statuses: []layout.StatusItem{{Label: "Request", Value: text.Pill("Accepted", text.Info, palette), Detail: "An accepted request does not confirm the remote effect."}}},
				{ID: "activity", Title: "Activity", FullWidth: true, Timeline: []layout.TimelineEvent{{When: "09:41 UTC", Actor: "controller", Title: "Request submitted", Detail: "Response unavailable"}, {When: "09:40 UTC", Actor: "operator", Title: "Intent recorded", Detail: "Source order retained"}}},
				{ID: "limits", Title: "Interpretation limits", Paragraphs: []string{"No independent observation confirms the route.", "literal\x1b[2J"}, Body: func(int) string { return "Custom evidence retained" }},
			},
		}
		for _, width := range []int{1, 2, 20, 40, 55, 80, 84, 132} {
			rendered := document.Render(width, palette)
			assertRecordBounds(t, rendered.Content, width)
			assertDetailRanges(t, rendered)
			compact := compactRecordText(rendered.Content)
			// Wrapped explanations in the right column read vertically even when
			// more facts occupy the same rows in the left column.
			verification := rendered.Sections[1]

			if verification.Start == rendered.Sections[0].Start {
				lines := strings.Split(rendered.Content, "\n")[verification.Start:verification.End]
				heading := ansi.Strip(lines[0])
				start := ansi.StringWidth(strings.SplitN(heading, "Verification", 2)[0])

				for i, line := range lines {
					lines[i] = ansi.Cut(line, start, width)
				}

				compact += compactRecordText(strings.Join(lines, "\n"))
			}

			for _, expected := range []string{"route-42", "Gatewayroute", "Outcomeunknown", "Inspectremoteevidence", "netops", "missing", `""`, "Accepted", "doesnotconfirm", "effect.", "Requestsubmitted", "Intentrecorded", "Noindependentobservation", "Customevidenceretained", `\u001b[2J`} {
				if !strings.Contains(compact, expected) {
					t.Fatalf("width %d lost %q: %q", width, expected, ansi.Strip(rendered.Content))
				}
			}

			if strings.Index(compact, "Requestsubmitted") > strings.Index(compact, "Intentrecorded") {
				t.Fatal("timeline silently reordered supplied source evidence")
			}
		}
	}
}

func TestDetailDocumentHandlesEmptyAndSingleCellHeaders(t *testing.T) {
	palette := theme.New(theme.Dark)
	document := layout.DetailDocument{Header: layout.IdentityHeader{Kind: "界", ID: "界", Context: "界"}, Sections: []layout.DetailSection{{ID: "empty"}}}
	view := document.View(1, palette)
	assertRecordBounds(t, view, 1)

	if strings.Contains(view, "界") || strings.Count(ansi.Strip(view), "!") < 3 {
		t.Fatalf("impossible header grapheme disappeared without a marker: %q", view)
	}

	if document.View(0, palette) != "" || (layout.DetailDocument{}).View(80, palette) != "" || layout.StatusList(nil, 80, palette) != "" || layout.Timeline(nil, 80, palette) != "" {
		t.Fatal("empty allocation or empty document produced content")
	}

	if got := compactRecordText(layout.StatusList([]layout.StatusItem{{Label: "Empty"}}, 80, palette)); !strings.Contains(got, `Empty""`) {
		t.Fatalf("empty status was silently omitted: %q", got)
	}
}

func assertDetailRanges(t *testing.T, rendering layout.DetailRendering) {
	t.Helper()

	rows := strings.Split(rendering.Content, "\n")
	for _, span := range rendering.Sections {
		if span.Start < 0 || span.End <= span.Start || span.End > len(rows) {
			t.Fatalf("section %+v escaped %d rendered rows", span, len(rows))
		}
	}
}

func TestDetailDocumentKeepsIndependentDeliveryStages(t *testing.T) {
	palette := theme.New(theme.Dark)
	document := layout.DetailDocument{
		Summary:      layout.RecordSummary{Title: "PR #42", Status: "Outcome uncertain", Next: &layout.Pair{Label: "Next action", Value: "Inspect the saved request"}},
		ActionDetail: "Check the receipt.",
		Sections: []layout.DetailSection{{ID: "feedback", Title: "Feedback delivery", Stages: []layout.StatusItem{
			{Label: "Transport", Value: "Recorded"},
			{Label: "Receipt", Value: "Unknown"},
			{Label: "Disposition", Value: "Open", Detail: "Owner has not responded."},
		}}},
	}

	wide := ansi.Strip(document.View(80, palette))
	lines := strings.Split(wide, "\n")

	if !strings.Contains(lines[0], "PR #42") || !strings.Contains(lines[0], "Outcome uncertain") {
		t.Fatalf("identity and outcome no longer share the wide headline: %q", lines[0])
	}

	var stages, outcomes string

	for i, line := range lines {
		if strings.Contains(line, "Transport") && i+1 < len(lines) {
			stages, outcomes = line, lines[i+1]
		}
	}

	for _, pair := range [][2]string{{"Transport", "Recorded"}, {"Receipt", "Unknown"}, {"Disposition", "Open"}} {
		stage, outcome := strings.Index(stages, pair[0]), strings.Index(outcomes, pair[1])
		if stage < 0 || outcome < 0 || ansi.StringWidth(stages[:stage]) != ansi.StringWidth(outcomes[:outcome]) {
			t.Fatalf("state %q is not under stage %q: %q", pair[1], pair[0], wide)
		}
	}

	for _, width := range []int{1, 2, 20, 32, 55, 80} {
		view := document.View(width, palette)
		assertRecordBounds(t, view, width)
		compact := compactRecordText(view)

		for _, wanted := range []string{"Recorded", "Unknown", "Open", "Ownerhasnotresponded.", "Checkthereceipt."} {
			if !strings.Contains(compact, wanted) {
				t.Fatalf("width %d lost %q: %q", width, wanted, ansi.Strip(view))
			}
		}
	}
}

func TestDetailStagesKeepWrappedHeadingsAndValuesInTheirRows(t *testing.T) {
	section := layout.DetailSection{Stages: []layout.StatusItem{
		{Label: "Transport", Value: "Recorded", Detail: "journal"},
		{Label: "Recipient acknowledgement", Value: "Awaiting confirmation", Detail: "recipient"},
		{Label: "Disposition", Value: "Open"},
	}}
	document := layout.DetailDocument{Sections: []layout.DetailSection{section}}
	view := ansi.Strip(document.View(56, theme.New(theme.Dark)))
	lines := strings.Split(view, "\n")

	if len(lines) != 5 || !strings.Contains(lines[1], "acknowledgement") ||
		!strings.Contains(lines[2], "Recorded") || !strings.Contains(lines[2], "Awaiting") || !strings.Contains(lines[2], "Open") ||
		!strings.Contains(lines[3], "confirmation") ||
		!strings.Contains(lines[4], "journal") || !strings.Contains(lines[4], "recipient") {
		t.Fatalf("stage labels, values and explanations lost their shared baselines:\n%s", view)
	}
}
