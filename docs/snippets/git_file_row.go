package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitFileRow() string {
	palette := theme.New(theme.Dark)

	file := gitui.FileChange{ID: "config", Path: "config.go", Index: gitui.Added, Worktree: gitui.Modified}

	return gitui.FileChangeRow(file, gitui.Staged, palette)
}
