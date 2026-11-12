package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitCommitRow() (string, error) {
	palette := theme.New(theme.Dark)

	commit := gitui.Commit{ID: "a31f2c7819be822", Subject: "Add search", Author: "Josh"}
	return gitui.CommitRow(commit, nil, gitui.Head{}, palette)
}
