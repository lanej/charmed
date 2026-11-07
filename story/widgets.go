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
	"charm.land/bubbles/v2/spinner"
	"charm.land/bubbles/v2/table"
	"charm.land/bubbles/v2/tree"
	tea "charm.land/bubbletea/v2"
	"charm.land/huh/v2"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/form"
	"github.com/lanej/charmed/input"
	"github.com/lanej/charmed/interaction"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/text"
)

func themeStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		t := e.Config.Palette()
		return title("Semantic tokens", e) + strings.Join([]string{t.Text().Render("Foreground / primary content"), t.Subtle().Render("Muted / supporting information"), t.Heading().Render("Accent / interactive emphasis"), t.Focus().Render("Focus / inverse + emphasis"), t.Text().Foreground(t.Success).Render("[OK] Success"), t.Text().Foreground(t.Warning).Render("[!] Warning"), t.Text().Foreground(t.Danger).Render("[x] Danger")}, "\n\n") + footer("Change theme, terminal size, or ASCII in Controls.", e)
	}), nil
}
func layoutStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		t := e.Config.Palette()
		left, right, split := layout.Split(size.Width, 1, 20, 30, .4)
		height := max(3, size.Height-4)
		a := layout.Panel("Navigation", "Cell-based dimensions\nBorders included in size", layout.Size{Width: left, Height: height}, t, true)
		b := layout.Panel("Content", "The right pane adapts to available space.\nNarrow terminals stack the same panels.", layout.Size{Width: right, Height: height}, t, false)
		if split {
			return title("Responsive layout", e) + layout.Row(a, " ", b)
		}
		height = max(3, height/2)
		return title("Stacked layout", e) + layout.Column(layout.Panel("Navigation", "Same content", layout.Size{Width: left, Height: height}, t, true), layout.Panel("Content", "Same state", layout.Size{Width: right, Height: height}, t, false))
	}), nil
}
func textStory(e Environment) (runtime.Scene, error) {
	return static(func(layout.Size) string {
		t := e.Config.Palette()
		return title("Text primitives", e) + text.Label(e.Config.Value, t) + "\n\n" + text.Badge("READY", t) + " " + text.Badge("3 SELECTED", t) + "\n\n" + text.KeyHint("Enter", "select", t) + "   " + text.KeyHint("Esc", "dismiss", t)
	}), nil
}

func inputStory(e Environment) (runtime.Scene, error) {
	m := input.NewText(input.TextOptions{Placeholder: "Type a value", Value: e.Config.Value, Width: e.Config.Columns - 4, Disabled: e.Config.State == "disabled", Theme: e.Config.Palette()})
	return &scene{init: m.Focus, resize: func(s layout.Size) { m.Resize(s.Width - 4) }, update: func(msg tea.Msg) tea.Cmd {
		var cmd tea.Cmd
		m, cmd = m.Update(msg)
		if k, ok := msg.(tea.KeyPressMsg); ok && k.String() == "enter" {
			emit(e, interaction.Submitted{Component: "input", Value: m.Value()})
		}
		return cmd
	}, view: func(s layout.Size) string {
		return title("Text input", e) + layout.Panel("Value", m.View(), layout.Size{Width: s.Width, Height: 3}, e.Config.Palette(), m.Focused()) + footer("Type to edit / Enter emits Submitted / Ctrl+C exits", e)
	}}, nil
}
func fieldStory(e Environment) (runtime.Scene, error) {
	f := input.Field{Label: "Environment name", Help: "Use at least three characters.", Theme: e.Config.Palette(), Input: input.NewText(input.TextOptions{Theme: e.Config.Palette(), Value: e.Config.Value, Width: e.Config.Columns - 2, Validate: func(s string) error {
		if len(s) < 3 {
			return errors.New("enter at least three characters")
		}
		return nil
	}})}
	if e.Config.State == "error" {
		f.Input.SetValue("x")
		f.Input.Err = errors.New("enter at least three characters")
	}
	return &scene{init: f.Input.Focus, resize: func(s layout.Size) { f.Input.Resize(s.Width) }, update: func(msg tea.Msg) tea.Cmd { var cmd tea.Cmd; f.Input, cmd = f.Input.Update(msg); return cmd }, view: func(layout.Size) string { return title("Labeled field", e) + f.View() }}, nil
}
func areaStory(e Environment) (runtime.Scene, error) {
	m := input.NewArea(e.Config.Columns, e.Config.Rows-5, e.Config.Palette())
	m.Placeholder = "Write a note"
	m.SetValue(e.Config.Value)
	return &scene{init: m.Focus, resize: func(s layout.Size) { m.SetWidth(s.Width); m.SetHeight(max(1, s.Height-5)) }, update: func(msg tea.Msg) tea.Cmd { var cmd tea.Cmd; m, cmd = m.Update(msg); return cmd }, view: func(layout.Size) string { return title("Multiline input", e) + m.View() }}, nil
}
func spinnerStory(e Environment) (runtime.Scene, error) {
	m := feedback.NewSpinner(e.Config.Palette())
	return &scene{init: func() tea.Cmd {
		if e.Config.Live && !e.Config.Palette().ReducedMotion {
			return m.Tick
		}
		return nil
	}, update: func(msg tea.Msg) tea.Cmd {
		if _, ok := msg.(runtime.StepMsg); ok {
			msg = spinner.TickMsg{ID: m.ID()}
		}
		var cmd tea.Cmd
		m, cmd = m.Update(msg)
		if !e.Config.Live {
			return nil
		}
		return cmd
	}, view: func(layout.Size) string {
		return title("Loading indicator", e) + m.View() + "  Processing fixture data" + footer("Live runs the native clock / Ctrl+N advances a frame", e)
	}}, nil
}
func progressStory(e Environment) (runtime.Scene, error) {
	amount := e.Config.Amount
	m := feedback.NewProgress(e.Config.Columns-10, e.Config.Palette())
	return &scene{resize: func(s layout.Size) { m.SetWidth(max(1, s.Width-10)) }, update: func(msg tea.Msg) tea.Cmd {
		if _, ok := msg.(runtime.StepMsg); ok {
			amount = min(100, amount+5)
			emit(e, map[string]any{"type": "progress", "amount": amount})
		}
		return nil
	}, view: func(layout.Size) string {
		return title("Progress", e) + m.ViewAs(amount/100) + fmt.Sprintf("  %.0f%%", amount) + footer("Ctrl+N advances by five percent.", e)
	}}, nil
}
func paginatorStory(e Environment) (runtime.Scene, error) {
	m := feedback.NewPaginator(50, 10, e.Config.Palette())
	return &scene{update: func(msg tea.Msg) tea.Cmd { var cmd tea.Cmd; m, cmd = m.Update(msg); return cmd }, view: func(layout.Size) string {
		return title("Pagination", e) + m.View() + fmt.Sprintf("\n\nPage %d of %d", m.Page+1, m.TotalPages) + footer("Left / right changes page", e)
	}}, nil
}
func timeStory(e Environment) (runtime.Scene, error) {
	timer := feedback.NewTimer(time.Minute)
	watch := feedback.NewStopwatch()
	return &scene{init: func() tea.Cmd {
		if e.Config.Live {
			return tea.Batch(timer.Init(), watch.Init())
		}
		return nil
	}, update: func(msg tea.Msg) tea.Cmd {
		var a, b tea.Cmd
		timer, a = timer.Update(msg)
		watch, b = watch.Update(msg)
		if !e.Config.Live {
			return nil
		}
		return tea.Batch(a, b)
	}, view: func(layout.Size) string {
		return title("Time components", e) + "Countdown  " + timer.View() + "\nStopwatch  " + watch.View() + footer("Enable Live to run the native timers.", e)
	}}, nil
}

type helpKeys struct{ next, selectKey key.Binding }

func (k helpKeys) ShortHelp() []key.Binding  { return []key.Binding{k.next, k.selectKey} }
func (k helpKeys) FullHelp() [][]key.Binding { return [][]key.Binding{k.ShortHelp()} }
func helpStory(e Environment) (runtime.Scene, error) {
	m := feedback.NewHelp(e.Config.Columns, e.Config.Palette())
	keys := helpKeys{key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "next field")), key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "select"))}
	return &scene{resize: func(s layout.Size) { m.SetWidth(s.Width) }, view: func(layout.Size) string { return title("Help from key bindings", e) + m.View(keys) }}, nil
}
func motionStory(e Environment) (runtime.Scene, error) {
	m := feedback.NewMotion(e.Config.Palette())
	return &scene{update: func(msg tea.Msg) tea.Cmd {
		if _, ok := msg.(runtime.StepMsg); ok {
			m.Step(e.Config.Amount)
			emit(e, map[string]any{"type": "motion", "position": m.Position})
		}
		return nil
	}, view: func(layout.Size) string {
		return title("Spring motion", e) + fmt.Sprintf("Position %.2f  /  Target %.0f\n\nVelocity %.2f", m.Position, e.Config.Amount, m.Velocity) + footer("Ctrl+N steps the spring. Reduced motion resolves immediately.", e)
	}}, nil
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

func (i fixtureItem) Title() string       { return i.name }
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
	root := os.Getenv("CHARMED_FIXTURE_DIR")
	var err error
	if root == "" {
		root, err = os.MkdirTemp("", "charmed-files-")
	} else {
		root = filepath.Join(root, "files")
		err = os.MkdirAll(root, 0700)
	}
	if err != nil {
		return nil, err
	}
	cleanup := func() error { return os.RemoveAll(root) }
	if err = os.Mkdir(filepath.Join(root, "components"), 0700); err != nil {
		cleanup()
		return nil, err
	}
	for _, name := range []string{"README.md", "theme.go", "components/input.go"} {
		if err = os.WriteFile(filepath.Join(root, name), []byte("// Charmed fixture\n"), 0600); err != nil {
			cleanup()
			return nil, err
		}
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
func viewportStory(e Environment) (runtime.Scene, error) {
	var lines []string
	for i := range 80 {
		lines = append(lines, fmt.Sprintf("%02d  A scrollable line of fixture content.", i+1))
	}
	m := collection.NewViewport(strings.Join(lines, "\n"), e.Config.Columns, e.Config.Rows-4)
	return &scene{resize: func(s layout.Size) { m.SetWidth(s.Width); m.SetHeight(max(1, s.Height-4)) }, update: func(msg tea.Msg) tea.Cmd { var cmd tea.Cmd; m, cmd = m.Update(msg); return cmd }, view: func(layout.Size) string { return title("Viewport", e) + m.View() }}, nil
}
func markdownStory(e Environment) (runtime.Scene, error) {
	content := "# Terminal documentation\n\nGlamour renders **Markdown** inside a Bubbles viewport.\n\n- Theme follows the terminal palette\n- Content wraps to its allocated width\n- Navigation belongs to this region\n\n```go\nfield := input.NewText(options)\ncmd := field.Focus()\n```\n\n## Composition\n\nKeep content, focus and layout independent.\n"
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
	return &scene{init: f.Init, resize: func(s layout.Size) { f.WithWidth(s.Width).WithHeight(s.Height) }, update: func(msg tea.Msg) tea.Cmd {
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
