// Archive Sentinel presents only the supplied synthetic observations.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image/color"
	"os"
	"strconv"
	"strings"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

type model struct {
	fixture     fixture
	index       int
	dashboard   *pattern.OperationalDashboard
	size        layout.Size
	ascii, mono bool
}

func newModel(f fixture, id string, size layout.Size, ascii, mono bool) (*model, error) {
	m := &model{fixture: f, size: size, ascii: ascii, mono: mono}
	found := false

	for i, s := range f.States {
		if s.ID == id {
			m.index, found = i, true
		}
	}

	if !found {
		return nil, fmt.Errorf("unknown state %q", id)
	}

	m.dashboard = pattern.NewOperationalDashboard(f.snapshot(m.index, m.palette()), size, m.palette())
	m.dashboard.AppKeys = []key.Binding{
		key.NewBinding(key.WithKeys("1", "2", "3", "4", "5"), key.WithHelp("1-5/s", "state")),
		key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "refresh")),
		key.NewBinding(key.WithKeys("a", "m"), key.WithHelp("a/m", "ASCII/mono")),
		key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
	}
	m.dashboard.Resize(size)

	return m, nil
}

func (m *model) palette() theme.Theme {
	mode := theme.Dark
	if m.mono {
		mode = theme.Mono
	}

	t := theme.New(mode)
	// Glyph repertoire and color depth are independent user choices.
	t.ASCII = m.ascii

	return t
}

func (m *model) apply() {
	m.dashboard.Theme = m.palette()
	m.dashboard.SetSnapshot(m.fixture.snapshot(m.index, m.palette()))
}

func (m *model) Init() tea.Cmd { return nil }
func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if size, ok := msg.(tea.WindowSizeMsg); ok {
		m.size = layout.Size{Width: size.Width, Height: size.Height}
		m.dashboard.Resize(m.size)

		return m, nil
	}

	if press, ok := msg.(tea.KeyPressMsg); ok {
		switch press.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "1", "2", "3", "4", "5":
			index, err := strconv.Atoi(press.String())
			if err == nil && index <= len(m.fixture.States) {
				m.index = index - 1
				m.apply()
			}

			return m, nil
		case "s":
			m.index = (m.index + 1) % len(m.fixture.States)
			m.apply()

			return m, nil
		case "r":
			m.apply()
			return m, nil
		case "a":
			m.ascii = !m.ascii
			m.apply()

			return m, nil
		case "m":
			m.mono = !m.mono
			m.apply()

			return m, nil
		}
	}

	return m, m.dashboard.Update(msg)
}

func (m *model) View() tea.View {
	v := tea.NewView(m.dashboard.View())
	v.AltScreen = true
	v.ForegroundColor = m.dashboard.Theme.Foreground
	v.BackgroundColor = m.dashboard.Theme.Background

	return v
}

// Replay uses the same messages and Update path as the interactive program.
func (m *model) replay(sequence string) error {
	if sequence == "" {
		return nil
	}

	for _, token := range strings.Split(sequence, ",") {
		if strings.HasPrefix(token, "resize:") {
			var width, height int
			if _, err := fmt.Sscanf(token, "resize:%dx%d", &width, &height); err != nil {
				return fmt.Errorf("resize %q: %w", token, err)
			}

			m.Update(tea.WindowSizeMsg{Width: width, Height: height})

			continue
		}

		var code rune

		switch token {
		case "down":
			code = tea.KeyDown
		case "up":
			code = tea.KeyUp
		case "pgdown":
			code = tea.KeyPgDown
		case "pgup":
			code = tea.KeyPgUp
		case "esc":
			code = tea.KeyEscape
		default:
			runes := []rune(token)
			if len(runes) != 1 {
				return fmt.Errorf("unknown replay key %q", token)
			}

			code = runes[0]
		}

		m.Update(tea.KeyPressMsg{Code: code})
	}

	return nil
}

func rgb(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("#%02x%02x%02x", r>>8, g>>8, b>>8)
}

func (m *model) capture(prefix, initial, sequence string) error {
	view := m.View()

	frame := view.Content
	for _, item := range []struct{ suffix, content string }{{".ansi", frame}, {".txt", ansi.Strip(frame)}} {
		if err := os.WriteFile(prefix+item.suffix, []byte(item.content+"\n"), 0o600); err != nil {
			return fmt.Errorf("write frame: %w", err)
		}
	}

	metadata := map[string]any{
		"initial_state": initial, "final_state": m.fixture.States[m.index].ID,
		"width": m.size.Width, "height": m.size.Height, "input_sequence": sequence,
		"foreground": rgb(view.ForegroundColor), "background": rgb(view.BackgroundColor),
		"ascii": m.ascii, "monochrome": m.mono, "inspection": m.dashboard.Inspecting(),
		"overview_offset": m.dashboard.Overview.YOffset(), "inspection_offset": m.dashboard.Inspection.YOffset(),
		"overview_lines": m.dashboard.Overview.TotalLineCount(), "inspection_lines": m.dashboard.Inspection.TotalLineCount(),
	}

	data, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return fmt.Errorf("encode metadata: %w", err)
	}

	if err := os.WriteFile(prefix+".json", append(data, '\n'), 0o600); err != nil {
		return fmt.Errorf("write metadata: %w", err)
	}

	return nil
}

func run() error {
	path := flag.String("fixture", "fixture.json", "synthetic input")
	stateID := flag.String("state", "ready", "waiting, empty, ready, failed_refresh, replacement")
	width := flag.Int("width", 132, "initial columns")
	height := flag.Int("height", 40, "initial rows")
	ascii := flag.Bool("ascii", false, "ASCII glyphs")
	mono := flag.Bool("mono", false, "monochrome colors; retains Unicode unless -ascii")
	capture := flag.String("capture", "", "write ANSI, plain text and metadata with this prefix")
	sequence := flag.String("keys", "", "comma separated app key replay; resize:40x18 supported")

	flag.Parse()

	f, err := readFixture(*path)
	if err != nil {
		return err
	}

	m, err := newModel(f, *stateID, layout.Size{Width: *width, Height: *height}, *ascii, *mono)
	if err != nil {
		return err
	}

	if err := m.replay(*sequence); err != nil {
		return err
	}

	if *capture != "" {
		return m.capture(*capture, *stateID, *sequence)
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		return fmt.Errorf("run TUI: %w", err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		if _, writeErr := fmt.Fprintln(os.Stderr, err); writeErr != nil {
			os.Exit(1)
		}

		os.Exit(1)
	}
}
