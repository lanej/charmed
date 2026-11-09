package story

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
)

func commandDefinitions() []Definition {
	config := DefaultConfig()
	config.Columns, config.Rows = 88, 28
	config.Value = ""

	return []Definition{{ID: "collection/command-palette", Title: "Molecules/Navigation/Command palette", Level: "Molecules", Family: "commands", Description: "Fuzzy-searchable commands with shortcuts, current settings and unavailable reasons; applications own execution.", Source: "story/commands.go", Symbol: "commandPaletteStory", ExampleSource: "docs/snippets/command-palette.go.txt", Composes: []string{"input/text", "text/key-hint", "foundations/layout"}, Controls: []string{"state", "value"}, States: []string{"ready", "empty", "disabled"}, Defaults: config, Factory: commandPaletteStory}}
}

func commandPaletteStory(e Environment) (runtime.Scene, error) {
	commands := []collection.Command{
		{ID: "timeline", Label: "Layout: timeline", Description: "Requests above details and related logs", Keywords: []string{"gantt", "view"}, State: "current"},
		{ID: "columns", Label: "Layout: columns", Description: "Requests, details and logs side by side"},
		{ID: "details", Label: "Focus: HTTP details", Description: "Inspect parameters, headers and response bodies", Shortcut: "Alt+2"},
		{ID: "follow", Label: "Follow requests", Description: "Keep arriving observations in view", Shortcut: "F"},
		{ID: "trace", Label: "Scope: trace", Description: "Correlate logs across the selected trace", DisabledReason: "Selected request has no trace ID"},
		{ID: "raw", Label: "Show raw body", Description: "Inspect captured bytes without formatting", Shortcut: "r"},
		{ID: "params", Label: "Expand parameters", Description: "Show decoded query and form values", Shortcut: "3"},
		{ID: "url", Label: "URL: full", Description: "Display the original captured URL"},
		{ID: "request", Label: "Expand request", Description: "Show captured request headers and body", Shortcut: "1"},
	}
	if e.Config.State == "empty" {
		commands = nil
	}

	m, err := collection.NewCommandPalette(commands, e.Config.Columns, e.Config.Rows, e.Config.Palette())
	if err != nil {
		return nil, err
	}

	m.Open()
	m.SetQuery(e.Config.Value)

	if e.Config.State == "disabled" {
		m.SetQuery("Scope: trace")
	}

	selected := ""

	return &scene{init: m.Search.Focus, resize: m.Resize, update: func(msg tea.Msg) tea.Cmd {
		if action, ok := msg.(collection.CommandSelectedMsg); ok {
			selected = action.ID
			return nil
		}

		return m.Update(msg)
	}, view: func(size layout.Size) string {
		background := strings.Repeat("Application content remains visible beneath the palette\n", max(1, size.Height))
		if selected != "" {
			background = "Selected command: " + selected + "\n" + background
		}

		if !m.IsOpen() {
			return layout.Fit(background, size)
		}

		return layout.CenterOverlay(background, m.View(), size, e.Config.Palette(), m.DimBackdrop)
	}}, nil
}
