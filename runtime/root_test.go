package runtime

import (
	"strings"
	"testing"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/layout"
)

type testScene struct {
	updates, prepares, views int
	size                     layout.Size
	revision                 uint64
	change                   bool
	initialized              bool
	values                   []int
	resizes, closes          int
	command                  tea.Cmd
}

func (s *testScene) Init() tea.Cmd { s.initialized = true; return nil }
func (s *testScene) Update(msg tea.Msg) tea.Cmd {
	if value, ok := msg.(int); ok {
		s.values = append(s.values, value)
	}

	s.updates++
	if s.change {
		s.revision++
	}

	return s.command
}
func (s *testScene) Resize(size layout.Size)      { s.size = size; s.resizes++ }
func (s *testScene) PresentationRevision() uint64 { return s.revision }
func (s *testScene) View() string                 { s.views++; s.prepares++; return strings.Repeat("x", s.size.Width) }

func testRoot(t *testing.T) (*Root, *testScene, *int) {
	t.Helper()

	s := &testScene{change: true}
	calls := new(int)
	root := &Root{Scene: s, Size: layout.Size{Width: 20, Height: 4}}
	root.now = func() time.Time { return time.Unix(100, 0) }
	root.after = func(d time.Duration, f func(time.Time) tea.Msg) tea.Cmd {
		*calls++

		if d != time.Second/60 {
			t.Fatalf("unexpected interval %s", d)
		}

		return func() tea.Msg { return f(time.Time{}) }
	}
	root.Init()

	return root, s, calls
}

func TestBurstCoalescesPresentationWithoutDroppingUpdates(t *testing.T) {
	root, scene, timers := testRoot(t)
	initial := root.View().Content

	var frame tea.Cmd

	for i := range 100 {
		_, command := root.Update(i)
		if i == 0 {
			frame = command
		} else if command != nil {
			t.Fatal("more than one frame pending")
		}

		if root.View().Content != initial {
			t.Fatal("View presented unprepared state")
		}
	}

	for i, value := range scene.values {
		if value != i {
			t.Fatal("updates reordered")
		}
	}

	if len(scene.values) != 100 {
		t.Fatal("updates lost")
	}

	if scene.updates != 100 || scene.prepares != 1 || *timers != 1 {
		t.Fatalf("updates=%d preparations=%d timers=%d", scene.updates, scene.prepares, *timers)
	}

	root.Update(frame())

	if scene.prepares != 2 {
		t.Fatal("burst was not presented once")
	}

	for range 100 {
		root.View()
	}

	if scene.prepares != 2 || scene.views != 2 {
		t.Fatal("cached View performed scene work")
	}

	if root.pending {
		t.Fatal("idle runtime still pending")
	}
}

func TestResizeUsesFinalGeometryAndCommittedFrame(t *testing.T) {
	root, scene, timers := testRoot(t)
	before := root.View().Content
	_, frame := root.Update(tea.WindowSizeMsg{Width: 30, Height: 8})
	root.Update(tea.WindowSizeMsg{Width: 40, Height: 10})

	if scene.resizes != 1 {
		t.Fatal("resize burst repeatedly arranged scene")
	}

	if root.View().Content != before {
		t.Fatal("pending geometry changed committed frame")
	}

	root.Update(frame())

	if scene.size != (layout.Size{Width: 40, Height: 10}) || !strings.Contains(root.View().Content, strings.Repeat("x", 40)) {
		t.Fatal("final geometry lost")
	}

	if *timers != 1 {
		t.Fatal("resize burst scheduled multiple frames")
	}

	if _, cmd := root.Update(tea.WindowSizeMsg{Width: 40, Height: 10}); cmd != nil {
		t.Fatal("unchanged size scheduled a frame")
	}
}

func TestNoopCloseAndStaleFrames(t *testing.T) {
	root, scene, timers := testRoot(t)
	scene.change = false

	if _, cmd := root.Update("no-op"); cmd != nil || *timers != 0 {
		t.Fatal("no-op scheduled work")
	}

	frame := root.Invalidate()
	if err := root.PresentNow(); err != nil {
		t.Fatal(err)
	}

	count := scene.prepares

	root.Update(frame())

	if scene.prepares != count {
		t.Fatal("stale frame prepared")
	}

	frame = root.Invalidate()
	if err := root.Close(); err != nil {
		t.Fatal(err)
	}

	root.Update(frame())
	root.Update("after close")

	if scene.prepares != count || scene.updates != 1 {
		t.Fatal("closed runtime processed work")
	}
}

func TestInvalidConfigurationIsObservable(t *testing.T) {
	root, _, _ := testRoot(t)

	root.FPS = -1
	if err := root.PresentNow(); err == nil || root.Err() == nil || root.View().Content == "" {
		t.Fatal("invalid configuration hidden")
	}

	if root.pending {
		t.Fatal("invalid configuration loops")
	}

	root.FPS = 60
	if err := root.PresentNow(); err != nil {
		t.Fatal(err)
	}
}

func TestForeignFrameCannotTriggerPresentation(t *testing.T) {
	root, scene, _ := testRoot(t)
	other, _, _ := testRoot(t)
	frame := other.Invalidate()
	root.Invalidate()
	root.Update(frame())

	if scene.prepares != 1 || !root.pending {
		t.Fatal("foreign frame consumed pending work")
	}
}

func TestInitialPresentationFollowsInitialization(t *testing.T) {
	root, scene, _ := testRoot(t)
	if !scene.initialized || root.FrameRevision() != 1 {
		t.Fatal("initialization did not precede first frame")
	}
}

func (s *testScene) Close() error { s.closes++; return nil }

type initializedScene struct {
	testScene
	initializedAt layout.Size
}

func (s *initializedScene) Init() tea.Cmd { s.initializedAt = s.size; s.initialized = true; return nil }

func (s *initializedScene) View() string {
	if s.initialized {
		return "initialized"
	}

	return "cold"
}

func TestInitSeesGeometryAndItsChangesReachFirstFrame(t *testing.T) {
	scene := &initializedScene{}
	root := &Root{Scene: scene, Size: layout.Size{Width: 20, Height: 4}}
	root.Init()

	if scene.initializedAt != root.Size || !strings.Contains(root.View().Content, "initialized") {
		t.Fatal("initial geometry, initialization and presentation were ordered incorrectly")
	}
}

func TestNativeCommandsAndQuitCleanupArePreserved(t *testing.T) {
	root, scene, _ := testRoot(t)
	scene.command = func() tea.Msg { return "native result" }
	_, command := root.Update(1)

	batch, ok := command().(tea.BatchMsg)
	if !ok || len(batch) != 2 || batch[0]() != "native result" {
		t.Fatal("native command lost while scheduling frame")
	}

	root.Update(tea.KeyPressMsg{Code: 'c', Mod: tea.ModCtrl})

	if err := root.Close(); err != nil {
		t.Fatal(err)
	}

	if scene.closes != 1 {
		t.Fatal("quit skipped scene cleanup")
	}

	if err := root.Close(); err != nil {
		t.Fatal(err)
	}

	if scene.closes != 1 {
		t.Fatal("close ran twice")
	}
}

func TestSparseUpdatesUseOnlyRemainingFrameBudget(t *testing.T) {
	root, _, _ := testRoot(t)
	now := time.Unix(100, 0)
	root.now = func() time.Time { return now }

	var delay time.Duration

	root.after = func(d time.Duration, f func(time.Time) tea.Msg) tea.Cmd {
		delay = d
		return func() tea.Msg { return f(now) }
	}
	now = now.Add(time.Second)
	command := root.Invalidate()

	if delay != 0 {
		t.Fatal("sparse input unnecessarily waited a full frame")
	}

	root.Update(command())

	now = now.Add(time.Millisecond)
	command = root.Invalidate()

	if delay != time.Second/60-time.Millisecond {
		t.Fatalf("remaining frame budget %s", delay)
	}

	root.Update(command())

	if root.PresentationPending() {
		t.Fatal("presentation scheduled idle work")
	}
}
