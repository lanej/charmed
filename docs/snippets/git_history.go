package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitHistory() (string, error) {
	palette := theme.New(theme.Dark)

	view, err := gitui.NewCommitHistory(gitui.History{Commits: []gitui.Commit{
		{ID: "merge", ParentIDs: []string{"main", "topic"}, Subject: "Merge topic"},
		{ID: "topic", ParentIDs: []string{"base"}, Subject: "Add search"},
		{ID: "main", ParentIDs: []string{"base"}, Subject: "Update guide"},
		{ID: "base", Subject: "Initial components"},
	}}, 80, 16, palette)
	if err != nil {
		return "", err
	}

	return view.View(), nil
}
