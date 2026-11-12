package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitBranchRow() string {
	palette := theme.New(theme.Dark)

	branch := gitui.Branch{Ref: gitui.Ref{Name: "refs/heads/main", Kind: gitui.LocalBranch}, Current: true}
	return gitui.BranchRow(branch, palette)
}
