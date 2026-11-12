package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitChanges() (string, error) {
	palette := theme.New(theme.Dark)

	view, err := gitui.NewChangesInspector(
		gitui.Repository{Files: []gitui.FileChange{{ID: "config", Path: "config.go", Index: gitui.Added}}},
		[]gitui.FileDiff{{FileID: "config", Section: gitui.Staged, Path: "config.go", State: gitui.TextDiff}},
		100, 20, palette)
	if err != nil {
		return "", err
	}
	return view.View(), nil
}
