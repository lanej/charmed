package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleLayoutKeyValue() string {
	pairs := []layout.Pair{
		{Label: "profile", Value: "dev"},
		{Label: "service", Value: "v1.42.0 (contract 7)"},
		{Label: "compat", State: text.ValueMissing}, // Missing is not an empty string.
	}

	return layout.KeyValue(pairs, 80, theme.New(theme.Dark))
}
