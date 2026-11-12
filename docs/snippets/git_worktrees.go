package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitWorktrees() (string, error) {
	palette := theme.New(theme.Dark)

	view, err := gitui.NewWorktreeList([]gitui.Worktree{{ID: "search", Path: "~/src/charmed-search", Locked: true, LockReason: "review"}}, 80, 12, palette)
	if err != nil {
		return "", err
	}

	return view.View(), nil
}
