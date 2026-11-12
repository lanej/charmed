package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitStashRow() string {
	palette := theme.New(theme.Dark)

	stash := gitui.Stash{ID: "stash-object", Ordinal: 0, Subject: "WIP: search", Branch: "main"}

	return gitui.StashRow(stash, palette)
}
