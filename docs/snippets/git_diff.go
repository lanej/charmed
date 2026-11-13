package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitDiff() string {
	palette := theme.New(theme.Dark)

	diff := gitui.FileDiff{
		Path: "config.go", State: gitui.TextDiff,
		Hunks: []gitui.Hunk{{Header: "@@ -1 +1 @@", Lines: []gitui.DiffLine{
			{Kind: gitui.DeletedLine, OldLine: 1, Text: `const name = "before"`},
			{Kind: gitui.AddedLine, NewLine: 1, Text: `const name = "after"`},
		}}},
	}
	view := gitui.NewDiffViewer(diff, 80, 12, palette)

	return view.View()
}
