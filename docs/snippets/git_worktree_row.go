package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitWorktreeRow() string {
	palette := theme.New(theme.Dark)

	worktree := gitui.Worktree{ID: "search", Path: "~/src/charmed-search", Head: gitui.Head{Kind: gitui.Attached, Ref: "refs/heads/search"}, Locked: true, LockReason: "review"}
	return gitui.WorktreeRow(worktree, palette)
}
