package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleFoundationsPermissions() string {
	palette := theme.New(theme.Dark)
	mode := text.Permissions{Mode: 0o755, Available: true}
	unknown := text.Permissions{}

	return layout.Column(
		mode.Compact(palette)+"  "+mode.Label(),
		unknown.Compact(palette)+"  "+unknown.Label(),
	)
}
