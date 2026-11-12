package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitDivergence() string {
	palette := theme.New(theme.Dark)

	return gitui.DivergenceIndicator(gitui.Divergence{Known: true, Ahead: 2, Behind: 1}, palette)
}
