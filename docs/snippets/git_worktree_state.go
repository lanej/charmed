package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitWorktreeState() string {
	palette := theme.New(theme.Dark)

	return gitui.WorktreeIndicators(gitui.Worktree{Main: true, Current: true, StatusKnown: true}, palette)
}
