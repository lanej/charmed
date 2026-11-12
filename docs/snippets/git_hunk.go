package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitHunk() string {
	palette := theme.New(theme.Dark)

	hunk := gitui.Hunk{Header: "@@ -0,0 +1 @@", Lines: []gitui.DiffLine{{Kind: gitui.AddedLine, NewLine: 1, Text: "package gitui"}}}

	return gitui.DiffHunk(hunk, palette)
}
