package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitDiff() string {
	palette := theme.New(theme.Dark)

	diff := gitui.FileDiff{Path: "config.go", State: gitui.BinaryDiff}
	view := gitui.NewDiffViewer(diff, 80, 12, palette)
	return view.View()
}
