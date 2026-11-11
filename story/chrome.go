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
		{ID: "layout/link-group", Title: "Molecules/Layout/Link group", Level: "Molecules", Family: "layout", Description: "Related destinations flow in order with shared spacing, wrap between items, and retain exact targets through long-label wrapping. Each link keeps its own presentation and emphasis.", Source: "layout/links.go", Symbol: "LinkGroup", Composes: []string{"text/link", "foundations/layout"}, Defaults: DefaultConfig(), Factory: linkGroupStory},
		{ID: "feedback/connection", Title: "Atoms/Feedback/Connection", Level: "Atoms", Family: "feedback", Description: "Connected and disconnected describe connectivity independently of health, severity or unknown state. Labels and glyphs remain visible in every profile.", Source: "story/chrome.go", Symbol: "connectionStory", Controls: []string{"emphasis"}, Composes: []string{"feedback/status"}, Defaults: DefaultConfig(), Factory: connectionStory},
		{ID: "navigation/tabs", Title: "Molecules/Navigation/Tabs", Level: "Molecules", Family: "navigation", Description: "Focused arrow-key navigation with stable selection, visible overflow and shared selection styling.", Source: "story/chrome.go", Symbol: "tabsStory", Composes: []string{"foundations/layout"}, Defaults: DefaultConfig(), Factory: tabsStory},
		{ID: "layout/header", Title: "Molecules/Layout/Header", Level: "Molecules", Family: "layout", Description: "Title and trailing actions with an optional context row. Narrow widths preserve the title before actions.", Source: "story/chrome.go", Symbol: "headerStory", Composes: []string{"foundations/layout"}, Defaults: DefaultConfig(), Factory: headerStory},
		{ID: "layout/identity-header", Title: "Molecules/Layout/Identity header", Level: "Molecules", Family: "layout", Description: "A request UUID or other entity identifier is the headline. Narrow terminals move the kind above the complete identifier before wrapping.", Source: "story/chrome.go", Symbol: "identityHeaderStory", Composes: []string{"foundations/layout"}, Defaults: DefaultConfig(), Factory: identityHeaderStory},
		{ID: "layout/section", Title: "Molecules/Layout/Foldable section", Level: "Molecules", Family: "layout", Description: "A panel folds to one heading row, returning space to sibling content. The application owns state and bindings; press Enter to toggle this specimen.", Source: "story/chrome.go", Symbol: "sectionStory", Composes: []string{"foundations/layout", "text/disclosure"}, Defaults: DefaultConfig(), Factory: sectionStory},
		{ID: "layout/footer", Title: "Molecules/Layout/Footer", Level: "Molecules", Family: "layout", Description: "Keyboard help and right-aligned status share one row. Narrow widths preserve help before status.", Source: "story/chrome.go", Symbol: "footerStory", Composes: []string{"foundations/layout", "text/primitives", "feedback/status"}, Defaults: DefaultConfig(), Factory: footerStory},
		{ID: "template/application", Title: "Templates/Application shell", Level: "Templates", Family: "pattern", Description: "A header, tab bar, active content panel and footer compose a resizable application shell. Left/Right changes the visible section.", Source: "story/chrome.go", Symbol: "applicationStory", Composes: []string{"layout/header", "navigation/tabs", "layout/footer", "foundations/layout", "feedback/status"}, Defaults: DefaultConfig(), Factory: applicationStory},
	}
	for i := range definitions {
		definitions[i].Defaults.Columns = 64
		definitions[i].Defaults.Rows = 8

		if definitions[i].ID == "layout/link-group" {
			definitions[i].Defaults.Rows = 16
		}

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

func identityHeaderStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		return (layout.IdentityHeader{Kind: "Request", ID: "01234567-89ab-cdef-0123-456789abcdef", Context: "POST /v2/shipments/shp_fixture/buy · HTTP 500 · 12.5ms"}).View(size.Width, e.Config.Palette())
	}), nil
}

func sectionStory(e Environment) (runtime.Scene, error) {
	section := layout.Section{Title: "Request timeline", Content: "POST /v2/shipments/shp_fixture/buy   500   12.5ms", NoFolds: e.Config.NoFolds}

	return &scene{update: func(msg tea.Msg) tea.Cmd {
		if key, ok := msg.(tea.KeyPressMsg); ok && key.String() == "enter" && !section.NoFolds {
			section.Collapsed = !section.Collapsed
		}

		return nil
	}, view: func(size layout.Size) string {
		top := section.Height(3)
		return layout.Column(section.View(layout.Size{Width: size.Width, Height: top}, e.Config.Palette(), true), layout.Panel("Response", "Internal server error\nCould not purchase shipment", layout.Size{Width: size.Width, Height: max(0, size.Height-top)}, e.Config.Palette(), false))
	}}, nil
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
