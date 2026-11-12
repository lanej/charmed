package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func ExampleGitFileState() string {
	palette := theme.New(theme.Dark)

	return gitui.FileIndicator(gitui.Modified, palette)
}
