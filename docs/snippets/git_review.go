package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleGitReview() (string, error) {
	review, err := gitui.NewFileDiffReview([]gitui.ReviewFile{
		{ID: "routing", Diff: gitui.FileDiff{
			FileID: "routing", Path: "src/router.go", State: gitui.TextDiff,
			Hunks: []gitui.Hunk{{Header: "@@ -1 +1 @@", Lines: []gitui.DiffLine{
				{Kind: gitui.DeletedLine, OldLine: 1, Text: "return fastest"},
				{Kind: gitui.AddedLine, NewLine: 1, Text: "return lastHealthy"},
			}}},
		}, Stat: &gitui.DiffStat{State: gitui.TextDiff, Added: 1, Deleted: 1}},
		{ID: "guide", Diff: gitui.FileDiff{FileID: "guide", Path: "README.md", State: gitui.UnavailableDiff, Reason: "patch not supplied"}},
	}, layout.Size{Width: 100, Height: 20}, theme.New(theme.Dark))
	if err != nil {
		return "", err
	}

	review.SetReviewed("routing", true)

	return review.View(), nil
}
