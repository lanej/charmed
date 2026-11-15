// Package fixture supplies deterministic review and non-file document examples.
package fixture

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/pattern"
)

// Document deliberately has no file or Git concepts.
func Document() []pattern.OutlineEntry {
	return []pattern.OutlineEntry{
		{ID: "proposal", Title: "Routing proposal", Summary: "2 sections", Children: []pattern.OutlineEntry{
			{ID: "intent", Title: "Intent", Summary: "Decision context", Body: func(int) string {
				return "Choose a regional route using supplied latency observations.\nKeep the last healthy route when new evidence is incomplete."
			}},
			{ID: "evidence", Title: "Evidence", Summary: "3 observations", Body: func(int) string {
				return "Region       p95       Samples\nWest         42 ms     12,400\nCentral      61 ms      8,100\nEast         unknown   not captured"
			}},
		}},
		{ID: "decision", Title: "Decision", Summary: "Awaiting review", Body: func(int) string {
			return "Retain the current route until the East observation arrives.\nThis document is application-supplied content, not a file tree."
		}},
	}
}

func Files() []gitui.ReviewFile {
	return []gitui.ReviewFile{
		{ID: "routing", Diff: gitui.FileDiff{FileID: "routing", Path: "src/router.go", State: gitui.TextDiff, Hunks: []gitui.Hunk{{Header: "@@ -18,4 +18,6 @@ func ChooseRoute", Lines: []gitui.DiffLine{
			{Kind: gitui.ContextLine, OldLine: 18, NewLine: 18, Text: "func ChooseRoute(observation Observation) Route {"},
			{Kind: gitui.DeletedLine, OldLine: 19, Text: "    return observation.Fastest()"},
			{Kind: gitui.AddedLine, NewLine: 19, Text: "    if !observation.Complete {"},
			{Kind: gitui.AddedLine, NewLine: 20, Text: "        return observation.Previous"},
			{Kind: gitui.AddedLine, NewLine: 21, Text: "    }"},
			{Kind: gitui.AddedLine, NewLine: 22, Text: "    return observation.FastestHealthy()"},
			{Kind: gitui.ContextLine, OldLine: 20, NewLine: 23, Text: "}"},
		}}}}, Stat: &gitui.DiffStat{State: gitui.TextDiff, Added: 4, Deleted: 1}},
		{ID: "tests", Diff: gitui.FileDiff{FileID: "tests", Path: "src/router_test.go", State: gitui.TextDiff, Hunks: []gitui.Hunk{{Header: "@@ -8,0 +9,4 @@", Lines: []gitui.DiffLine{
			{Kind: gitui.AddedLine, NewLine: 9, Text: "func TestIncompleteEvidenceRetainsRoute(t *testing.T) {"},
			{Kind: gitui.AddedLine, NewLine: 10, Text: "    got := ChooseRoute(incompleteObservation())"},
			{Kind: gitui.AddedLine, NewLine: 11, Text: "    require.Equal(t, previousRoute, got)"},
			{Kind: gitui.AddedLine, NewLine: 12, Text: "}"},
		}}}}, Stat: &gitui.DiffStat{State: gitui.TextDiff, Added: 4}},
		{ID: "guide", Diff: gitui.FileDiff{FileID: "guide", Path: "docs/routing.md", State: gitui.TextDiff, Hunks: []gitui.Hunk{{Header: "@@ -1 +1 @@", Lines: []gitui.DiffLine{
			{Kind: gitui.DeletedLine, OldLine: 1, Text: "Always select the fastest region."},
			{Kind: gitui.AddedLine, NewLine: 1, Text: "Require complete observations before changing routes."},
		}}}}, Stat: &gitui.DiffStat{State: gitui.TextDiff, Added: 1, Deleted: 1}},
		{ID: "image", Diff: gitui.FileDiff{FileID: "image", Path: "docs/routing.png", State: gitui.BinaryDiff}, Stat: &gitui.DiffStat{State: gitui.BinaryDiff}},
		{ID: "settings", Diff: gitui.FileDiff{FileID: "settings", Path: "config.yaml", State: gitui.UnavailableDiff, Reason: "patch was not supplied"}},
	}
}
