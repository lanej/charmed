package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitSummary() string {
	palette := theme.New(theme.Dark)

	return gitui.RepositorySummary(gitui.Repository{Name: "charmed", Head: gitui.Head{Kind: gitui.Attached, Ref: "refs/heads/main"}}, palette)
}
