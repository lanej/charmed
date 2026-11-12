package diagram_test

import (
	"fmt"
	"image/color"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/diagram"
	"github.com/lanej/charmed/theme"
)

func TestStateDiagramSimpleDirectedTransition(t *testing.T) {
	d := diagram.StateDiagram{States: []diagram.State{{ID: "A", Initial: true}, {ID: "B", Final: true}}, Transitions: []diagram.Transition{{From: "A", To: "B", Label: "go"}}}

	out, err := d.View(80, theme.New(theme.Mono))
	if err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(out, "\n")
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], " ")
	}

	want := []string{
		"(-----------------)        #===============#",
		"| A               +-- go -># B             |",
		"| initial         |        | final         |",
		"(-----------------)        #===============#",
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
		for _, label := range []string{"Idle", "Running", "Done", "start", "retry", "finish", "cancel", "resume", "active", "initial", "final"} {
			if strings.Count(plain, label) != 1 {
				t.Fatalf("lost or duplicated %q in %q", label, plain)
			}
		}

		arrows := strings.Count(plain, "▶") + strings.Count(plain, "◀") + strings.Count(plain, "▲")
		if mode == theme.Mono {
			arrows = strings.Count(plain, ">") + strings.Count(plain, "<") + strings.Count(plain, "^")
		}

		if arrows != len(d.Transitions) {
			t.Fatalf("transition lost its arrow: %q", plain)
		}

		if strings.Contains(plain, "Crossings") {
			t.Fatal("diagram requires a legend")
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
		{States: []diagram.State{{ID: "a", Outcome: diagram.StateOutcome(99)}}},
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
	for _, width := range []int{-1, 0, 5} {
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

	lines := strings.Split(out, "\n")
	if !strings.Contains(lines[1], "A") || !strings.Contains(lines[1], "go") || !strings.Contains(lines[1], "B") {
		t.Fatalf("forward label is detached from its nodes: %q", out)
	}

	if !strings.Contains(lines[2], "<- back") {
		t.Fatalf("reverse edge lacks its inline label and arrow: %q", out)
	}

	if !strings.Contains(lines[len(lines)-1], "loop") || !strings.Contains(lines[len(lines)-2], "^") {
		t.Fatalf("self-loop label or destination lost: %q", out)
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

func TestStateDiagramNarrowReflowKeepsLabelsOnEdges(t *testing.T) {
	d := diagram.StateDiagram{States: []diagram.State{{ID: "idle", Label: "Idle", Initial: true}, {ID: "running", Label: "Running"}, {ID: "done", Label: "Done", Final: true}}, Transitions: []diagram.Transition{{From: "idle", To: "running", Label: "start"}, {From: "running", To: "running", Label: "retry"}, {From: "running", To: "done", Label: "finish"}, {From: "running", To: "idle", Label: "cancel"}}, Active: "running"}

	out, err := d.View(40, theme.New(theme.Mono))
	if err != nil {
		t.Fatal(err)
	}

	for _, line := range strings.Split(out, "\n") {
		if ansi.StringWidth(line) > 40 {
			t.Fatal("vertical diagram escaped width")
		}
	}

	for _, label := range []string{"Idle", "Running", "Done", "start", "retry", "finish", "cancel"} {
		if strings.Count(out, label) != 1 {
			t.Fatalf("vertical diagram lost or duplicated %q: %q", label, out)
		}
	}

	if strings.Contains(out, "Crossings") {
		t.Fatal("diagram has a legend")
	}

	if len(regexp.MustCompile(`(?m)^[ |]*v[ |]*$`).FindAllString(out, -1)) != 2 {
		t.Fatalf("vertical edges lost arrowheads: %q", out)
	}

	if !strings.Contains(out, "| start") || !strings.Contains(out, "| finish") || !strings.Contains(out, "| retry") {
		t.Fatalf("vertical labels detached from paths: %q", out)
	}
}

func TestStateDiagramWidthGuidanceIsSufficient(t *testing.T) {
	d := diagram.StateDiagram{States: []diagram.State{{ID: "A"}, {ID: "B"}}, Transitions: []diagram.Transition{{From: "A", To: "B"}, {From: "B", To: "A"}}}

	_, err := d.View(1, theme.New(theme.Mono))
	if err == nil {
		t.Fatal("expected insufficient width")
	}

	var required, available int
	if _, scanErr := fmt.Sscanf(err.Error(), "state diagram needs %d columns; available %d", &required, &available); scanErr != nil {
		t.Fatal(scanErr)
	}

	if _, renderErr := d.View(required, theme.New(theme.Mono)); renderErr != nil {
		t.Fatalf("reported minimum is insufficient: %v", renderErr)
	}
}

func TestStateDiagramLifecycleFitsEightyColumns(t *testing.T) {
	d := diagram.StateDiagram{States: []diagram.State{{ID: "disconnected", Label: "Disconnected", Initial: true}, {ID: "connecting", Label: "Connecting"}, {ID: "connected", Label: "Connected"}, {ID: "retrying", Label: "Retrying"}, {ID: "failed", Label: "Failed"}, {ID: "closed", Label: "Closed", Final: true}}, Transitions: []diagram.Transition{{From: "disconnected", To: "connecting", Label: "connect"}, {From: "connecting", To: "connected", Label: "success"}, {From: "connecting", To: "retrying", Label: "timeout"}, {From: "retrying", To: "connecting", Label: "retry"}, {From: "connecting", To: "failed", Label: "failure"}, {From: "connected", To: "closed", Label: "close"}}, Active: "connected"}

	out, err := d.View(80, theme.New(theme.Mono))
	if err != nil {
		t.Fatal(err)
	}

	for _, state := range d.States {
		if !strings.Contains(out, state.Label) {
			t.Fatalf("lost state %s", state.ID)
		}
	}

	for _, line := range strings.Split(out, "\n") {
		if ansi.StringWidth(line) > 80 {
			t.Fatal("lifecycle diagram exceeded 80 columns")
		}
	}
}

func TestStateDiagramComplexEdgesSurviveBothLayouts(t *testing.T) {
	d := diagram.StateDiagram{States: []diagram.State{{ID: "a", Label: "Alpha", Initial: true}, {ID: "b", Label: "Beta"}, {ID: "c", Label: "Gamma"}, {ID: "d", Label: "Delta", Final: true}}, Transitions: []diagram.Transition{{From: "a", To: "c", Label: "skip"}, {From: "b", To: "d", Label: "ship"}, {From: "c", To: "a", Label: "reset"}, {From: "a", To: "c", Label: "resend"}, {From: "b", To: "b", Label: "retry"}, {From: "c", To: "c", Label: "again"}, {From: "a", To: "b", Label: "start"}, {From: "b", To: "c", Label: "advance"}, {From: "c", To: "d", Label: "finish"}}, Active: "b"}
	for _, width := range []int{80, 100, 120} {
		out, err := d.View(width, theme.New(theme.Dark))
		if err != nil {
			t.Fatal(err)
		}

		plain := ansi.Strip(out)
		for _, state := range d.States {
			if strings.Count(plain, state.Label) != 1 {
				t.Fatalf("lost node %s", state.ID)
			}
		}

		for _, edge := range d.Transitions {
			if strings.Count(plain, edge.Label) != 1 {
				t.Fatalf("lost or overpainted edge label %q at width %d:\n%s", edge.Label, width, plain)
			}
		}

		count := 0
		for _, arrow := range []string{"▶", "◀", "▲", "▼"} {
			count += strings.Count(plain, arrow)
		}

		if count != len(d.Transitions) {
			t.Fatalf("lost destination arrow at width %d: %d arrows for %d edges\n%s", width, count, len(d.Transitions), plain)
		}

		if strings.Contains(plain, "Crossings") {
			t.Fatal("complex graph needs a legend")
		}
	}
}

func TestStateKindsAndOutcomes(t *testing.T) {
	d := diagram.StateDiagram{States: []diagram.State{
		{ID: "Start", Initial: true},
		{ID: "Neutral", Final: true},
		{ID: "Accepted", Final: true, Outcome: diagram.OutcomeSuccess},
		{ID: "Rejected", Final: true, Outcome: diagram.OutcomeError},
	}}

	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		for _, width := range []int{40, 200} {
			palette := theme.New(mode)

			out, err := d.View(width, palette)
			if err != nil {
				t.Fatal(err)
			}

			plain := ansi.Strip(out)
			for _, label := range []string{"Start", "Neutral", "Accepted", "Rejected", "initial", "success", "error"} {
				if !strings.Contains(plain, label) {
					t.Fatalf("lost %q in %s", label, plain)
				}
			}

			if mode == theme.Mono {
				if out != plain || !strings.Contains(plain, "(---") || !strings.Contains(plain, "#===") {
					t.Fatalf("missing ASCII state shapes: %s", out)
				}
			} else if !strings.Contains(plain, "╭─") || !strings.Contains(plain, "╔═") {
				t.Fatalf("missing Unicode state shapes: %s", out)
			}
		}
	}
}

func TestOutcomesAreIndependentOfFinalAndActive(t *testing.T) {
	d := diagram.StateDiagram{States: []diagram.State{{ID: "A", Initial: true, Outcome: diagram.OutcomeError}, {ID: "B", Final: true, Outcome: diagram.OutcomeSuccess}}, Transitions: []diagram.Transition{{From: "A", To: "B", Label: "recover"}}, Active: "B"}
	palette := theme.New(theme.Dark)
	palette.Success = color.RGBA{R: 19, G: 83, B: 147, A: 255}

	out, err := d.View(100, palette)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(ansi.Strip(out), "success, active") || !strings.Contains(out, "\x1b[38;2;19;83;147;1mB") {
		t.Fatalf("active state lost custom success color or marker: %s", out)
	}

	d.Active = ""

	inactive, err := d.View(100, palette)
	if err != nil {
		t.Fatal(err)
	}

	if strings.Count(out, "\n") != strings.Count(inactive, "\n") || ansi.StringWidth(strings.Split(out, "\n")[0]) != ansi.StringWidth(strings.Split(inactive, "\n")[0]) {
		t.Fatal("active changed geometry")
	}
}

func TestPossibleAndActiveStateEmphasis(t *testing.T) {
	for _, tc := range []struct {
		outcome       diagram.StateOutcome
		brightR, dimR uint8
	}{
		{diagram.OutcomeNone, 23, 11}, {diagram.OutcomeSuccess, 29, 13}, {diagram.OutcomeError, 31, 17},
	} {
		outcome := tc.outcome
		palette := theme.New(theme.Dark)
		palette.Accent = color.RGBA{R: 23, G: 89, B: 151, A: 255}
		palette.Success = color.RGBA{R: 29, G: 89, B: 151, A: 255}
		palette.Danger = color.RGBA{R: 31, G: 89, B: 151, A: 255}
		palette.DimForeground = color.RGBA{R: 11, G: 37, B: 61, A: 255}
		palette.DimSuccess = color.RGBA{R: 13, G: 37, B: 61, A: 255}
		palette.DimDanger = color.RGBA{R: 17, G: 37, B: 61, A: 255}
		d := diagram.StateDiagram{States: []diagram.State{{ID: "Current", Outcome: outcome}, {ID: "Possible", Outcome: outcome}}, Active: "Current"}

		out, err := d.View(100, palette)
		if err != nil {
			t.Fatal(err)
		}

		if !strings.Contains(out, fmt.Sprintf("\x1b[38;2;%d;89;151;1m", tc.brightR)) || !strings.Contains(out, fmt.Sprintf("\x1b[38;2;%d;37;61m", tc.dimR)) {
			t.Fatalf("bright and dim semantic colors missing: %s", out)
		}

		plain := ansi.Strip(out)
		if !strings.Contains(plain, "┏━") || !strings.Contains(plain, "┌─") || !strings.Contains(plain, "active") {
			t.Fatalf("current and possible states lack distinct borders: %s", out)
		}

		d.Active = "Possible"

		moved, err := d.View(100, palette)
		if err != nil {
			t.Fatal(err)
		}

		if moved == out || !strings.Contains(ansi.Strip(moved), "┏━") {
			t.Fatal("active emphasis did not move with current state")
		}
	}
}
