package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitRef() string {
	palette := theme.New(theme.Dark)

	return gitui.RefLabel(gitui.Ref{Name: "refs/heads/main", Kind: gitui.LocalBranch}, palette)
}
