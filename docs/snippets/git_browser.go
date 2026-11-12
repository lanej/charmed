package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleGitBrowser() (string, error) {
	palette := theme.New(theme.Dark)

	repository := gitui.Repository{Name: "charmed", Head: gitui.Head{Kind: gitui.Attached, Ref: "refs/heads/main"}}
	status, err := gitui.NewStatus(repository, 60, 16, palette)
	if err != nil {
		return "", err
	}
	worktrees, err := gitui.NewWorktreeList([]gitui.Worktree{{ID: "main", Path: "~/src/charmed", Main: true, Head: repository.Head}}, 60, 16, palette)
	if err != nil {
		return "", err
	}
	return layout.Row(status.View(), worktrees.View()), nil
}
