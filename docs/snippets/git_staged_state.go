package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitStagedState() string {
	palette := theme.New(theme.Dark)

	return gitui.SectionIndicator(gitui.Staged, palette)
}
