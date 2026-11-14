// Package runtime adapts a composed scene to one Bubble Tea terminal session.
package runtime

import (
	"errors"
	"image/color"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/layout"
)

// Scene is an application boundary, not a required widget interface. Updates
// remain native and ordered. An optional PresentationRevision() uint64 lets a
// scene identify no-op updates; other scenes conservatively request a frame.
type Scene interface {
	Init() tea.Cmd
	Update(tea.Msg) tea.Cmd
	Resize(layout.Size)
	View() string
}

type (
	StepMsg  struct{}
	frameMsg struct {
		owner      *Root
		generation uint64
	}
)

// Root batches presentation at FPS (default 60), with no idle frame command.
// All methods run in the model's serialized update loop. Size is the latest
// requested geometry; View retains the last completed frame until presentation.
type Root struct {
	Scene                    Scene
	Background               color.Color
	Size                     layout.Size
	AltScreen                bool
	ASCII                    bool
	FPS                      int
	frame                    tea.View
	err                      error
	pending, closed, resized bool
	quitting                 bool
	arranged                 layout.Size
	geometryValid            bool
	generation, frames       uint64
	lastFrame                time.Time
	now                      func() time.Time
	after                    func(time.Duration, func(time.Time) tea.Msg) tea.Cmd
}

func (m *Root) Err() error { return m.err }

// FrameRevision advances only when a new frame has been built, not on cached View calls.
func (m *Root) FrameRevision() uint64 { return m.frames }

// PresentationPending reports whether one on-demand frame command is outstanding.
func (m *Root) PresentationPending() bool { return m.pending }

func (m *Root) Close() error {
	if m.closed {
		return nil
	}

	m.closed, m.pending = true, false

	m.generation++
	if c, ok := m.Scene.(interface{ Close() error }); ok {
		return c.Close()
	}

	return nil
}

func (m *Root) Init() tea.Cmd {
	if m.closed {
		return nil
	}

	if m.Scene == nil {
		m.fail(errors.New("presentation scene is nil"))
		return tea.Quit
	}

	m.resized = true
	m.applyGeometry()

	command := m.Scene.Init()
	if err := m.PresentNow(); err != nil {
		return tea.Quit
	}

	return command
}

// Invalidate requests a frame after direct native customization. Invalidate the
// affected component's own cache first, if it has one, then return this command.
func (m *Root) Invalidate() tea.Cmd {
	if m.pending || m.closed || m.quitting {
		return nil
	}

	if m.FPS < 0 {
		m.fail(errors.New("FPS must be positive or zero for the default"))
		return nil
	}

	fps := m.FPS
	if fps == 0 {
		fps = 60
	}

	interval := max(time.Nanosecond, time.Second/time.Duration(fps))

	delay := time.Duration(0)
	if !m.lastFrame.IsZero() {
		delay = max(time.Duration(0), m.lastFrame.Add(interval).Sub(m.clock()))
	}

	m.pending = true
	generation := m.generation

	after := m.after
	if after == nil {
		after = tea.Tick
	}

	return after(delay, func(time.Time) tea.Msg { return frameMsg{owner: m, generation: generation} })
}

func (m *Root) applyGeometry() {
	if m.resized || !m.geometryValid || m.arranged != m.Size {
		m.Scene.Resize(m.Size)
		m.arranged, m.geometryValid = m.Size, true
		m.resized = false
	}
}

func (m *Root) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.closed || m.quitting {
		return m, nil
	}

	if tick, ok := msg.(frameMsg); ok {
		if tick.owner != m || tick.generation != m.generation || !m.pending {
			return m, nil
		}

		if err := m.PresentNow(); err != nil {
			return m, nil
		}

		return m, nil
	}

	if key, ok := msg.(tea.KeyPressMsg); ok {
		if key.String() == "ctrl+c" {
			m.pending = false
			m.generation++
			m.quitting = true

			return m, tea.Quit
		}

		if key.String() == "ctrl+n" {
			msg = StepMsg{}
		}
	}

	if size, ok := msg.(tea.WindowSizeMsg); ok {
		next := layout.Size{Width: max(1, size.Width), Height: max(1, size.Height)}
		if next == m.Size {
			return m, nil
		}

		m.Size, m.resized = next, true

		return m, m.Invalidate()
	}

	if m.Scene == nil {
		m.fail(errors.New("presentation scene is nil"))
		return m, nil
	}
	// Geometry precedes subsequent input: paging and selection must see its bounds.
	m.applyGeometry()
	revision, tracked := m.Scene.(interface{ PresentationRevision() uint64 })

	var before uint64
	if tracked {
		before = revision.PresentationRevision()
	}

	command := m.Scene.Update(msg)
	if tracked && before == revision.PresentationRevision() {
		return m, command
	}

	frame := m.Invalidate()
	if command == nil {
		return m, frame
	}

	if frame == nil {
		return m, command
	}

	return m, tea.Batch(command, frame)
}

func (m *Root) fail(err error) {
	m.err = err
	m.frame = m.makeView(layout.Fit("Presentation unavailable: "+err.Error(), m.Size))
}

// PresentNow synchronously builds the latest frame for snapshots and initial
// presentation. Pending frame messages become stale; no new timer is started.
func (m *Root) PresentNow() error {
	if m.closed {
		return errors.New("presentation runtime is closed")
	}

	m.pending = false

	m.generation++
	if m.Scene == nil {
		err := errors.New("presentation scene is nil")
		m.fail(err)

		return err
	}

	if m.FPS < 0 {
		err := errors.New("FPS must be positive or zero for the default")
		m.fail(err)

		return err
	}

	m.applyGeometry()

	content := m.Scene.View()
	if m.ASCII {
		content = layout.ASCII(content)
	}

	m.frame = m.makeView(layout.Fit(content, m.Size))
	m.frames++
	m.lastFrame = m.clock()
	m.err = nil

	return nil
}

func (m *Root) makeView(content string) tea.View {
	view := tea.NewView(strings.Clone(content))
	view.AltScreen = m.AltScreen
	view.BackgroundColor = m.Background

	return view
}

// View returns the completed frame without preparing, composing or fitting it.
func (m *Root) View() tea.View { return m.frame }

// Render is the synchronous snapshot path. Failures are retained in Err and shown.
func (m *Root) Render() tea.View {
	if err := m.PresentNow(); err != nil {
		return m.frame
	}

	return m.frame
}

func (m *Root) clock() time.Time {
	if m.now != nil {
		return m.now()
	}

	return time.Now()
}
