package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitStatus() (string, error) {
	palette := theme.New(theme.Dark)

	view, err := gitui.NewStatus(gitui.Repository{Name: "charmed", Files: []gitui.FileChange{{ID: "config", Path: "config.go", Index: gitui.Added}}}, 80, 16, palette)
	if err != nil {
		return "", err
	}

	return view.View(), nil
}
