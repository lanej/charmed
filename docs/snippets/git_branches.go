package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitBranches() (string, error) {
	palette := theme.New(theme.Dark)

	view, err := gitui.NewBranchList([]gitui.Branch{{Ref: gitui.Ref{Name: "refs/heads/main", Kind: gitui.LocalBranch}, Current: true}}, 80, 12, palette)
	if err != nil {
		return "", err
	}

	return view.View(), nil
}
