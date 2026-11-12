package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitHead() string {
	palette := theme.New(theme.Dark)

	return gitui.HeadIndicator(gitui.Head{Kind: gitui.Unborn, Ref: "refs/heads/main"}, palette)
}
