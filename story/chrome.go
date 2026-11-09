package story

import (
	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/navigation"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/text"
)

func chromeDefinitions() []Definition {
	definitions := []Definition{
		{ID: "feedback/connection", Title: "Atoms/Feedback/Connection", Level: "Atoms", Family: "feedback", Description: "Connected and disconnected describe connectivity independently of health, severity or unknown state. Labels and glyphs remain visible in every profile.", Source: "story/chrome.go", Symbol: "connectionStory", Controls: []string{"emphasis"}, Composes: []string{"feedback/status"}, Defaults: DefaultConfig(), Factory: connectionStory},
		{ID: "navigation/tabs", Title: "Molecules/Navigation/Tabs", Level: "Molecules", Family: "navigation", Description: "Focused arrow-key navigation with stable selection, visible overflow and shared selection styling.", Source: "story/chrome.go", Symbol: "tabsStory", Composes: []string{"foundations/layout"}, Defaults: DefaultConfig(), Factory: tabsStory},
		{ID: "layout/header", Title: "Molecules/Layout/Header", Level: "Molecules", Family: "layout", Description: "Title and trailing actions with an optional context row. Narrow widths preserve the title before actions.", Source: "story/chrome.go", Symbol: "headerStory", Composes: []string{"foundations/layout"}, Defaults: DefaultConfig(), Factory: headerStory},
		{ID: "layout/footer", Title: "Molecules/Layout/Footer", Level: "Molecules", Family: "layout", Description: "Keyboard help and right-aligned status share one row. Narrow widths preserve help before status.", Source: "story/chrome.go", Symbol: "footerStory", Composes: []string{"foundations/layout", "text/primitives", "feedback/status"}, Defaults: DefaultConfig(), Factory: footerStory},
		{ID: "template/application", Title: "Templates/Application shell", Level: "Templates", Family: "pattern", Description: "A header, tab bar, active content panel and footer compose a resizable application shell. Left/Right changes the visible section.", Source: "story/chrome.go", Symbol: "applicationStory", Composes: []string{"layout/header", "navigation/tabs", "layout/footer", "foundations/layout", "feedback/status"}, Defaults: DefaultConfig(), Factory: applicationStory},
	}
	for i := range definitions {
		definitions[i].Defaults.Columns = 64
		definitions[i].Defaults.Rows = 8

		if definitions[i].Level == "Templates" {
			definitions[i].Defaults.Rows = 12
		}
	}

	return definitions
}

func connectionStory(e Environment) (runtime.Scene, error) {
	return static(func(layout.Size) string {
		t := e.Config.Palette()
		emphasis := text.Emphasis(e.Config.Emphasis)

		return layout.Column(feedback.Connected.View(emphasis, t), feedback.Disconnected.View(emphasis, t))
	}), nil
}

func chromeTabs(e Environment) navigation.Tabs {
	m := navigation.NewTabs([]navigation.Tab{{ID: "overview", Label: "Overview"}, {ID: "services", Label: "Services"}, {ID: "activity", Label: "Activity"}, {ID: "settings", Label: "Settings"}}, e.Config.Palette())
	m.Focus()

	return m
}

func tabsStory(e Environment) (runtime.Scene, error) {
	m := chromeTabs(e)

	return &scene{resize: func(size layout.Size) { m.SetWidth(size.Width) }, update: func(msg tea.Msg) tea.Cmd {
		return m.Update(msg)
	}, view: func(layout.Size) string {
		return title("Tabs", e) + m.View() + footer("Left / Right select / Home / End jump", e)
	}}, nil
}

func headerStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		return (layout.Header{Title: "Service workbench", Context: "Production / us-west-2", Actions: text.KeyHint("r", "refresh", e.Config.Palette())}).View(size.Width, e.Config.Palette())
	}), nil
}

func footerStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		t := e.Config.Palette()
		return (layout.Footer{Help: text.KeyHint("?", "help", t) + "  " + text.KeyHint("q", "quit", t), Status: feedback.Connected.View(text.Normal, t)}).View(size.Width, t)
	}), nil
}

func applicationStory(e Environment) (runtime.Scene, error) {
	m := chromeTabs(e)

	return &scene{resize: func(size layout.Size) { m.SetWidth(size.Width) }, update: func(msg tea.Msg) tea.Cmd {
		return m.Update(msg)
	}, view: func(size layout.Size) string {
		t := e.Config.Palette()
		selected, _ := m.Selected()
		content := map[string]string{
			"overview": "12 services / all healthy\nSelect a tab to inspect another section.",
			"services": "api          running\nworker       running\nscheduler    idle",
			"activity": "12:04  Deployment completed\n12:01  Health checks passed",
			"settings": "Environment  production\nRegion       us-west-2",
		}[selected.ID]
		header := (layout.Header{Title: "Service workbench", Context: "Production / us-west-2", Actions: "SYNTHETIC"}).View(size.Width, t)
		body := layout.Panel("", content, layout.Size{Width: size.Width, Height: max(0, size.Height-4)}, t, false)
		footer := (layout.Footer{Help: "arrows tabs  Ctrl+C quit", Status: feedback.Connected.View(text.Normal, t)}).View(size.Width, t)

		return layout.Column(header, m.View(), body, footer)
	}}, nil
}
