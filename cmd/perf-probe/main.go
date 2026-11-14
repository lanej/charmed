// Command perf-probe measures native update-to-view time using the real runtime.
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/story"
)

type probe struct {
	root     *runtime.Root
	trace    *json.Encoder
	started  time.Time
	kind     string
	traceErr error
}

func (m *probe) Init() tea.Cmd { return m.root.Init() }
func (m *probe) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch event := message.(type) {
	case tea.KeyPressMsg:
		if event.String() != "ctrl+c" {
			m.started = time.Now()
			m.kind = "key"
		}
	case tea.WindowSizeMsg:
		m.started = time.Now()
		m.kind = "resize"
	default:
		m.started = time.Now()
		m.kind = "message"
	}

	_, command := m.root.Update(message)

	return m, command
}

func (m *probe) View() tea.View {
	view := m.root.View()
	if !m.started.IsZero() {
		if err := m.trace.Encode(struct {
			Kind     string `json:"kind"`
			Duration int64  `json:"duration_ns"`
		}{m.kind, time.Since(m.started).Nanoseconds()}); err != nil {
			m.traceErr = err
		}

		m.started = time.Time{}
	}

	return view
}

func run() (err error) {
	id := flag.String("story", "git/browser", "Native fixture")
	path := flag.String("trace", "", "Trace JSONL destination")
	live := flag.Bool("live", false, "Enable fixture animation")

	flag.Parse()

	if *path == "" {
		return errors.New("--trace is required")
	}

	file, err := os.Create(*path)
	if err != nil {
		return err
	}

	defer func() { err = errors.Join(err, file.Close()) }()

	definition, err := story.Find(*id)
	if err != nil {
		return err
	}

	config := definition.Defaults
	config.Live = *live

	model, err := story.New(*id, config, nil)
	if err != nil {
		return err
	}

	defer func() { err = errors.Join(err, model.Close()) }()

	measured := &probe{root: model, trace: json.NewEncoder(file)}
	if _, err = tea.NewProgram(measured).Run(); err != nil {
		return err
	}

	return measured.traceErr
}

func main() {
	if err := run(); err != nil {
		if _, writeErr := fmt.Fprintln(os.Stderr, err); writeErr != nil {
			os.Exit(2)
		}

		os.Exit(1)
	}
}
