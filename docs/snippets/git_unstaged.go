package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitUnstaged() (string, error) {
	palette := theme.New(theme.Dark)

	files := []gitui.FileChange{
		{ID: "config", Path: "config.go", Index: gitui.Added, Worktree: gitui.Modified},
		{ID: "notes", Path: "notes.txt", Worktree: gitui.Untracked},
		{ID: "merge", Path: "merge.go", Index: gitui.Conflicted, Conflict: "both modified"},
	}

	group, err := gitui.NewUnstagedChanges(files, palette)
	if err != nil {
		return "", err
	}

	return group.View(), nil
}
