package snippets

import (
	"github.com/lanej/charmed/navigation"
	"github.com/lanej/charmed/theme"
)

func ExampleNavigationTabs() string {
	tabs := navigation.NewTabs([]navigation.Tab{
		{ID: "overview", Label: "Overview"},
		{ID: "services", Label: "Services"},
		{ID: "activity", Label: "Activity"},
	}, theme.New(theme.Dark))
	tabs.SetWidth(64)
	tabs.Focus()
	// Return tabs.Update(msg)'s command from Update; tabs.Selected() gives the active ID.
	return tabs.View()
}
