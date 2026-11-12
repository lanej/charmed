package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitDiffSummary() string {
	palette := theme.New(theme.Dark)

	return gitui.DiffSummary(gitui.DiffStat{State: gitui.TextDiff, Added: 12, Deleted: 3}, palette)
}
