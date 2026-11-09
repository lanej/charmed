package snippets

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/huh/v2"
	"github.com/lanej/charmed/form"
	"github.com/lanej/charmed/theme"
)

func ExampleFormSettings() string {
	name, region := "production", "west"
	features := []string{"metrics"}
	settings := form.New(theme.New(theme.Dark), huh.NewGroup(
		huh.NewInput().Title("Environment").Value(&name),
		huh.NewSelect[string]().Title("Region").Options(
			huh.NewOption("US West", "west"), huh.NewOption("US East", "east"),
		).Value(&region),
		huh.NewMultiSelect[string]().Title("Features").Options(
			huh.NewOption("Metrics", "metrics"), huh.NewOption("Logs", "logs"),
		).Value(&features),
	))
	settings.WithWidth(48).WithHeight(12)
	_, _ = settings.Update(tea.WindowSizeMsg{Width: 48, Height: 12})
	// For a standalone interactive form, call settings.Run() and handle its error.
	return settings.View()
}
