package story

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/table"
	"charm.land/bubbles/v2/tree"
	tea "charm.land/bubbletea/v2"
	"charm.land/huh/v2"
	"charm.land/log/v2"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/form"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/text"
)

func organismsDefinitions() []Definition {
	d := []Definition{
		dataGridDefinition(),
		statusMatrixDefinition(),
		widget("collection/table", "Organisms/Collections/Table", "Organisms", "collection", "Native row navigation with themed headers, cells and selection.", "tableStory", []string{"state"}, tableStory),
		widget("collection/list", "Organisms/Collections/List", "Organisms", "collection", "A native list with filtering, pagination and selection.", "listStory", nil, listStory),
		widget("collection/tree", "Organisms/Collections/Tree", "Organisms", "collection", "Expand and collapse nested data without losing native tree capabilities.", "treeStory", nil, treeStory),
		widget("collection/files", "Organisms/Collections/File picker", "Organisms", "collection", "A file picker confined to a disposable fixture directory.", "fileStory", nil, fileStory),
		widget("text/markdown", "Organisms/Text/Markdown", "Organisms", "text", "Glamour renders Markdown within a native scrollable viewport.", "markdownStory", nil, markdownStory),
		widget("form/settings", "Organisms/Forms/Settings", "Organisms", "form", "A Huh form owns its nested focus and validation.", "formStory", nil, formStory),
	}

	d = append(d,
		Definition{ID: "collection/logs", Title: "Organisms/Collections/Log viewer", Level: "Organisms", Family: "collection", Description: "Structured Charm Log entries in a bounded, scrollable viewport.", Source: "story/organisms.go", Symbol: "logStory", Upstream: "https://github.com/charmbracelet/log", Composes: []string{"collection/viewport", "feedback/log-level"}, Step: true, Defaults: DefaultConfig(), Factory: logStory},
		Definition{ID: "chart/comparison-panel", Title: "Organisms/Charts/Comparison panel", Level: "Organisms", Family: "chart", Description: "A distribution comparison and native time-series history compose one reusable metric panel.", Source: "story/organisms.go", Symbol: "comparisonPanelStory", Upstream: "https://github.com/NimbleMarkets/ntcharts", Controls: []string{"sigma"}, Composes: []string{"chart/distributions", "chart/timeseries", "foundations/layout"}, Defaults: DefaultConfig(), Factory: comparisonPanelStory},
	)
	for i := range d {
		switch d[i].ID {
		case "collection/tree", "collection/table":
			d[i].Level = "Atoms"
			d[i].Title = strings.Replace(d[i].Title, "Organisms/", "Atoms/", 1)
		case "text/markdown":
			d[i].Upstream = "https://github.com/charmbracelet/glamour"
			d[i].Composes = []string{"collection/viewport"}
		case "form/settings":
			d[i].Upstream = "https://github.com/charmbracelet/huh"
		}
	}

	return d
}

func tableStory(e Environment) (runtime.Scene, error) {
	rows := []table.Row{{"API gateway", "Healthy", "1,240"}, {"Event stream", "Healthy", "860"}, {"Job workers", "Degraded", "420"}, {"Storage", "Healthy", "640"}}
	if e.Config.State == "empty" {
		rows = nil
	}

	m := collection.NewTable([]table.Column{{Title: "Service", Width: 25}, {Title: "State", Width: 15}, {Title: "Req/min", Width: 12}}, rows, e.Config.Columns, e.Config.Rows-5, e.Config.Palette())

	return &scene{resize: func(s layout.Size) { m.SetWidth(s.Width); m.SetHeight(max(1, s.Height-5)) }, update: func(msg tea.Msg) tea.Cmd {
		var cmd tea.Cmd

		m, cmd = m.Update(msg)
		if k, ok := msg.(tea.KeyPressMsg); ok {
			emit(e, map[string]any{"type": "row", "cursor": m.Cursor(), "key": k.String()})
		}

		return cmd
	}, view: func(layout.Size) string {
		if stateMessage(e) != "" {
			return title("Table", e) + stateMessage(e)
		}

		return title("Table", e) + m.View() + footer("Up / down selects a row", e)
	}}, nil
}

type fixtureItem struct{ name, description string }

func (i fixtureItem) Title() string { return i.name }

func (i fixtureItem) Description() string { return i.description }

func (i fixtureItem) FilterValue() string { return i.name }

func listStory(e Environment) (runtime.Scene, error) {
	items := []list.Item{fixtureItem{"Production", "Primary service environment"}, fixtureItem{"Staging", "Integration environment"}, fixtureItem{"Preview", "Disposable review environment"}}
	m := collection.NewList(items, e.Config.Columns, e.Config.Rows, e.Config.Palette())
	m.Title = "Environments"

	return &scene{resize: func(s layout.Size) { m.SetSize(s.Width, s.Height) }, update: func(msg tea.Msg) tea.Cmd { var cmd tea.Cmd; m, cmd = m.Update(msg); return cmd }, view: func(layout.Size) string { return m.View() }}, nil
}

func treeStory(e Environment) (runtime.Scene, error) {
	root := tree.Root("charmed").Child(tree.Root("components").Child("input", "collection", "chart"), tree.Root("workbench").Child("native", "browser", "MDX"))
	m := collection.NewTree(root, e.Config.Columns, e.Config.Rows, e.Config.Palette())

	return &scene{resize: func(s layout.Size) { m.SetSize(s.Width, s.Height) }, update: func(msg tea.Msg) tea.Cmd { var cmd tea.Cmd; m, cmd = m.Update(msg); return cmd }, view: func(layout.Size) string { return m.View() }}, nil
}

func fileStory(e Environment) (runtime.Scene, error) {
	root, err := os.MkdirTemp(os.Getenv("CHARMED_FIXTURE_DIR"), "charmed-files-")
	if err != nil {
		return nil, err
	}

	cleanup := func() error { return os.RemoveAll(root) }
	if err = writeFixtureFiles(root); err != nil {
		return nil, errors.Join(err, cleanup())
	}

	m := collection.NewFilePicker(root, e.Config.Palette())
	m.ShowPermissions = false
	m.ShowSize = false
	m, _ = m.Update(m.Init()())

	return &scene{close: cleanup, resize: func(s layout.Size) { m.SetHeight(max(1, s.Height-6)) }, update: func(msg tea.Msg) tea.Cmd {
		if k, ok := msg.(tea.KeyPressMsg); ok && key.Matches(k, m.KeyMap.Back) && m.CurrentDirectory == root {
			return nil
		}

		var cmd tea.Cmd

		m, cmd = m.Update(msg)
		if ok, path := m.DidSelectFile(msg); ok {
			emit(e, map[string]string{"type": "file", "path": strings.TrimPrefix(path, root+string(filepath.Separator))})
		}

		return cmd
	}, view: func(layout.Size) string {
		return title("Fixture files", e) + m.View() + footer("Arrows browse / Enter selects / fixture directory only", e)
	}}, nil
}

func markdownStory(e Environment) (runtime.Scene, error) {
	content := "# Terminal documentation\n\nUse `field.Focus()` with **Markdown**, *emphasis*, and [Glamour](https://github.com/charmbracelet/glamour).\n\n- Theme follows the terminal palette\n- Prose wraps to its allocated width\n\n```go\n// Focus a native input\nfield := input.NewText(options)\ncmd := field.Focus()\n```\n\n> Keep content, focus and layout independent.\n\n| Content | Rendering |\n| --- | --- |\n| Inline code | Literal spans |\n| Fenced source | Syntax highlighting |\n"

	rendered, err := text.Markdown(content, e.Config.Columns, e.Config.Palette())
	if err != nil {
		return nil, err
	}

	m := collection.NewViewport(rendered, e.Config.Columns, e.Config.Rows)

	return &scene{resize: func(s layout.Size) {
		m.SetWidth(s.Width)
		m.SetHeight(s.Height)

		if out, err := text.Markdown(content, s.Width, e.Config.Palette()); err == nil {
			m.SetContent(out)
		}
	}, update: func(msg tea.Msg) tea.Cmd { var cmd tea.Cmd; m, cmd = m.Update(msg); return cmd }, view: func(layout.Size) string { return m.View() }}, nil
}

func formStory(e Environment) (runtime.Scene, error) {
	name := "production"
	region := "west"
	enabled := true
	features := []string{"metrics"}
	note := ""
	f := form.New(e.Config.Palette(), huh.NewGroup(
		huh.NewInput().Title("Environment").Value(&name).Validate(func(s string) error {
			if len(s) < 3 {
				return errors.New("use at least three characters")
			}

			return nil
		}),
		huh.NewSelect[string]().Title("Region").Options(huh.NewOption("US West", "west"), huh.NewOption("US East", "east")).Value(&region),
		huh.NewMultiSelect[string]().Title("Features").Options(huh.NewOption("Metrics", "metrics"), huh.NewOption("Logs", "logs")).Value(&features),
		huh.NewConfirm().Title("Enable environment?").Value(&enabled)), huh.NewGroup(
		huh.NewText().Title("Notes").Value(&note), huh.NewNote().Title("Ready").Description("Your fixture settings are ready to submit.")))

	return &scene{init: f.Init, resize: func(s layout.Size) {
		f.WithWidth(s.Width).WithHeight(s.Height)

		if e.capture {
			// Prepare the native group viewport with explicit geometry. Do not execute
			// focus, terminal queries, or other initialization commands during capture.
			_, _ = f.Update(tea.WindowSizeMsg{Width: s.Width, Height: s.Height})
		}
	}, update: func(msg tea.Msg) tea.Cmd {
		_, cmd := f.Update(msg)
		if f.State == huh.StateCompleted {
			emit(e, map[string]any{"type": "submitted", "name": name, "region": region, "enabled": enabled, "features": features})
		}

		return cmd
	}, view: func(layout.Size) string {
		if f.State == huh.StateCompleted {
			return title("Settings saved", e) + fmt.Sprintf("%s / %s / enabled %v", name, region, enabled)
		}

		return f.View()
	}}, nil
}

func logStory(e Environment) (runtime.Scene, error) {
	m := collection.NewLogViewer(e.Config.Columns, e.Config.Rows-4, e.Config.Palette())
	step := 0

	appendLine := func() {
		level := log.InfoLevel
		if step%5 == 4 {
			level = log.WarnLevel
		}

		m.Append(collection.LogEntry{Time: time.Date(2026, 1, 1, 12, 0, step, 0, time.UTC), Level: level, Message: "Fixture request completed", Fields: []any{"service", "api", "sequence", step}})
		step++
	}
	for range 15 {
		appendLine()
	}

	return &scene{resize: func(s layout.Size) { m.SetWidth(s.Width); m.SetHeight(max(1, s.Height-4)) }, update: func(msg tea.Msg) tea.Cmd {
		if _, ok := msg.(runtime.StepMsg); ok {
			appendLine()
			return nil
		}

		var cmd tea.Cmd

		m.Model, cmd = m.Model.Update(msg)

		return cmd
	}, view: func(layout.Size) string {
		return title("Structured logs", e) + m.View() + footer("Arrows scroll / Ctrl+N appends a deterministic log entry", e)
	}}, nil
}

func comparisonPanelStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		selected := 40.0
		distribution := chart.DistributionComparison{Series: []chart.DistributionSeries{{Label: "a13c9f2", Mean: 20, StdDev: 3, Members: 64, HasSpread: true}, {Label: "f4d7e81", Mean: 30, StdDev: 8, Members: 32, HasSpread: true, Selected: &selected}}, Sigma: e.Config.Sigma, Min: 0, Max: 100, Unit: "%", SpreadUnit: "pp"}

		data := sampleSeries(2, 0, 20)
		for i, id := range []string{"a13c9f2", "f4d7e81"} {
			data[i].ID, data[i].Label, data[i].Unit = id, id, "%"
		}

		plot := chart.NewTimeSeries(data, layout.Size{Width: max(1, layout.PanelContentSize(size).Width), Height: max(3, size.Height-8)}, e.Config.Palette())

		return (chart.ComparisonPanel{Title: "Response share", Distribution: distribution, History: plot}).View(size, e.Config.Palette(), false)
	}), nil
}

func writeFixtureFiles(root string) (err error) {
	directory, err := os.OpenRoot(root)
	if err != nil {
		return err
	}
	defer func() { err = errors.Join(err, directory.Close()) }()

	if err = directory.Mkdir("components", 0o700); err != nil {
		return err
	}

	for _, name := range []string{"README.md", "theme.go", "components/input.go"} {
		if err = directory.WriteFile(name, []byte("// Charmed fixture\n"), 0o600); err != nil {
			return err
		}
	}

	return nil
}
