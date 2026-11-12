package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitCommitId() string {
	palette := theme.New(theme.Dark)

	return gitui.CommitIdentity("a31f2c7819be822", palette)
}
