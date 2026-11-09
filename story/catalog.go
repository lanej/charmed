package story

import (
	"slices"
	"strings"

	"github.com/lanej/charmed/runtime"
)

func definitions() []Definition {
	var d []Definition
	for _, level := range [][]Definition{foundationsDefinitions(), atomsDefinitions(), moleculesDefinitions(), organismsDefinitions(), screensDefinitions(), logDefinitions()} {
		d = append(d, level...)
	}
	for i := range d {
		if len(d[i].States) == 0 && slices.Contains(d[i].Controls, "state") {
			d[i].States = []string{"ready", "empty", "loading", "error"}
		}
	}
	return d
}

func widget(id, title, level, family, description, symbol string, controls []string, factory func(Environment) (runtime.Scene, error)) Definition {
	return Definition{ID: id, Title: title, Level: level, Family: family, Description: description, Source: "story/" + strings.ToLower(level) + ".go", Symbol: symbol, Upstream: "https://github.com/charmbracelet/bubbles", Controls: controls, Defaults: DefaultConfig(), Factory: factory}
}
