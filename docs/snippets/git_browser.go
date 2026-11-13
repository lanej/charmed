package snippets

import (
	"fmt"

	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
	"github.com/lanej/charmed/tui"
)

func ExampleGitBrowser() (string, error) {
	palette := theme.New(theme.Dark)

	screen, err := tui.NewScreen("status", "worktrees")
	if err != nil {
		return "", err
	}

	root := &layout.SplitPane{First: &layout.Region{ID: "status", Min: layout.Size{Width: 40, Height: 9}}, Second: &layout.Region{ID: "worktrees", Min: layout.Size{Width: 30, Height: 5}}, Ratio: .6, Gap: 1}
	if arrangeErr := screen.Arrange(root, layout.Size{Width: 120, Height: 20}); arrangeErr != nil {
		return "", arrangeErr
	}

	bounds, visible := screen.Region("status")
	if !visible {
		return "", fmt.Errorf("status region is not visible")
	}

	repository := gitui.Repository{Name: "charmed", Head: gitui.Head{Kind: gitui.Attached, Ref: "refs/heads/main"}}

	status, err := gitui.NewStatus(repository, bounds.Dx(), bounds.Dy(), palette)
	if err != nil {
		return "", err
	}

	bounds, visible = screen.Region("worktrees")
	if !visible {
		return "", fmt.Errorf("worktrees region is not visible")
	}

	worktrees, err := gitui.NewWorktreeList([]gitui.Worktree{{ID: "main", Path: "~/src/charmed", Main: true, Head: repository.Head}}, bounds.Dx(), bounds.Dy(), palette)
	if err != nil {
		return "", err
	}

	frame, err := screen.Render(map[string]string{"status": status.View(), "worktrees": worktrees.View()})

	return frame.Content, err
}
