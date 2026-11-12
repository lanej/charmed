package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitStashes() (string, error) {
	palette := theme.New(theme.Dark)

	view, err := gitui.NewStashList([]gitui.Stash{{ID: "stash-object", Ordinal: 0, Subject: "WIP: search"}}, 80, 12, palette)
	if err != nil {
		return "", err
	}
	return view.View(), nil
}
