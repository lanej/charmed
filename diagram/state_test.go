package diagram_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/diagram"
	"github.com/lanej/charmed/theme"
)

func TestStateDiagramSimpleDirectedTransition(t *testing.T) {
	d := diagram.StateDiagram{States: []diagram.State{{ID: "A", Initial: true}, {ID: "B", Final: true}}, Transitions: []diagram.Transition{{From: "A", To: "B", Label: "go"}}}

	out, err := d.View(40, theme.New(theme.Mono))
	if err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(out, "\n")
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], " ")
	}

	want := []string{
		"+---------+   +-------+",
		"| A       |   | B     |",
		"| initial |   | final |",
		"+-+-------+   +-+-----+",
		"  |             ^",
		"  |             |",
		"  +-------------+        go",
	}
	if !reflect.DeepEqual(lines, want) {
		t.Fatalf("directed transition differs:\n%s", out)
	}
}

func TestStateDiagramRetainsCyclesParallelEdgesAndSelfLoops(t *testing.T) {
	d := diagram.StateDiagram{
		States:      []diagram.State{{ID: "idle", Label: "Idle", Initial: true}, {ID: "run", Label: "Running"}, {ID: "done", Label: "Done", Final: true}},
		Transitions: []diagram.Transition{{From: "idle", To: "run", Label: "start"}, {From: "run", To: "run", Label: "retry"}, {From: "run", To: "done", Label: "finish"}, {From: "run", To: "idle", Label: "cancel"}, {From: "idle", To: "run", Label: "resume"}},
		Active:      "run",
	}
	before := append([]diagram.Transition(nil), d.Transitions...)

	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		palette := theme.New(mode)

		a, err := d.View(100, palette)
		if err != nil {
			t.Fatal(err)
		}

		b, err := d.View(100, palette)
		if err != nil {
			t.Fatal(err)
		}

		if a != b {
			t.Fatal("layout is nondeterministic")
		}

		plain := ansi.Strip(a)
		for _, label := range []string{"Idle", "Running", "Done", "start", "retry", "finish", "cancel", "resume", "active", "initial", "final", "Crossings do not join transitions."} {
			if strings.Count(plain, label) != 1 {
				t.Fatalf("lost or duplicated %q in %q", label, plain)
			}
		}

		arrow := "↑"
		if mode == theme.Mono {
			arrow = "^"
		}

		if strings.Count(plain, arrow) != len(d.Transitions) {
			t.Fatal("a transition lost its destination arrow")
		}

		for _, line := range strings.Split(a, "\n") {
			if ansi.StringWidth(line) > 100 {
				t.Fatal("diagram escaped its allocation")
			}
		}

		if mode == theme.Mono && a != plain {
			t.Fatal("monochrome diagram contains ANSI")
		}
	}

	if !reflect.DeepEqual(d.Transitions, before) {
		t.Fatal("renderer mutated caller transitions")
	}
}

func TestStateDiagramRejectsInvalidGraphsWithoutPartialOutput(t *testing.T) {
	for _, d := range []diagram.StateDiagram{
		{States: []diagram.State{{ID: "a"}, {ID: "a"}}},
		{States: []diagram.State{{ID: ""}}},
		{States: []diagram.State{{ID: "a", Initial: true}, {ID: "b", Initial: true}}},
		{States: []diagram.State{{ID: "a"}}, Active: "missing"},
		{States: []diagram.State{{ID: "a"}}, Transitions: []diagram.Transition{{From: "a", To: "missing"}}},
		{States: []diagram.State{{ID: "a", Final: true}}, Transitions: []diagram.Transition{{From: "a", To: "a"}}},
		{States: []diagram.State{{ID: "a", Label: "a\nb"}}},
		{States: []diagram.State{{ID: "a", Label: "a\u2028b"}}},
		{States: []diagram.State{{ID: "a"}}, Transitions: []diagram.Transition{{From: "a", To: "a", Label: "\x1b[31munsafe"}}},
	} {
		out, err := d.View(100, theme.New(theme.Dark))
		if err == nil || out != "" {
			t.Fatalf("invalid graph produced output %q, error %v", out, err)
		}
	}

	d := diagram.StateDiagram{States: []diagram.State{{ID: "a"}, {ID: "b"}}, Transitions: []diagram.Transition{{From: "a", To: "b", Label: "transition"}}}
	for _, width := range []int{-1, 0, 5, 20} {
		out, err := d.View(width, theme.New(theme.Dark))
		if err == nil || out != "" || !strings.Contains(err.Error(), "columns") {
			t.Fatalf("narrow allocation lost topology: %q, %v", out, err)
		}
	}
}

func TestStateDiagramUnicodeLabelsAndIsolatedStates(t *testing.T) {
	d := diagram.StateDiagram{States: []diagram.State{{ID: "one", Label: "待機"}, {ID: "two", Label: "cafe\u0301"}}}

	out, err := d.View(40, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(ansi.Strip(out), "待機") || !strings.Contains(ansi.Strip(out), "cafe\u0301") {
		t.Fatal("Unicode label was damaged")
	}

	lines := strings.Split(out, "\n")

	width := ansi.StringWidth(lines[0])
	for _, line := range lines {
		if ansi.StringWidth(line) != width {
			t.Fatal("Unicode labels misaligned state boxes")
		}
	}

	empty, err := (diagram.StateDiagram{}).View(40, theme.New(theme.Dark))
	if err != nil || empty != "" {
		t.Fatal("empty diagram must render empty")
	}
}

func TestStateDiagramLoopAndReverseRouting(t *testing.T) {
	d := diagram.StateDiagram{
		States:      []diagram.State{{ID: "A"}, {ID: "B"}},
		Transitions: []diagram.Transition{{From: "A", To: "A", Label: "loop"}, {From: "A", To: "B", Label: "go"}, {From: "B", To: "A", Label: "back"}},
	}

	out, err := d.View(40, theme.New(theme.Mono))
	if err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(out, "\n")[4:11]
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i][:21], " ")
	}

	want := []string{
		"  | ^ | ^       ^ |",
		"  | | | |       | |",
		"  +-+ | |       | |",
		"      | |       | |",
		"      +-+-------+ |",
		"        |         |",
		"        +---------+",
	}
	if !reflect.DeepEqual(lines, want) {
		t.Fatalf("loop or reverse transition was rewired:\n%s", out)
	}

	d.Active = "A"

	a, err := d.View(60, theme.New(theme.Mono))
	if err != nil {
		t.Fatal(err)
	}

	d.Active = "B"

	b, err := d.View(60, theme.New(theme.Mono))
	if err != nil {
		t.Fatal(err)
	}

	if strings.Split(a, "\n")[0] != strings.Split(b, "\n")[0] {
		t.Fatal("active marker moved the state boxes")
	}
}
