package diagram_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/diagram"
	"github.com/lanej/charmed/theme"
)

func tracedGraph() diagram.StateDiagram {
	return diagram.StateDiagram{
		States:      []diagram.State{{ID: "A", Initial: true}, {ID: "B"}, {ID: "C", Final: true}},
		Transitions: []diagram.Transition{{From: "A", To: "B", Label: "start"}, {From: "B", To: "B", Label: "retry"}, {From: "B", To: "A", Label: "back"}, {From: "A", To: "B", Label: "resume"}, {From: "B", To: "C", Label: "finish"}},
		Active:      "C", Trace: diagram.TransitionTrace{Steps: []int{0, 1, 2, 3, 4}},
	}
}

func TestTransitionTraceWindow(t *testing.T) {
	d := tracedGraph()
	palette := theme.New(theme.Dark)

	for _, tc := range []struct {
		name  string
		last  int
		all   bool
		shown []string
	}{
		{"default", 0, false, []string{"finish"}},
		{"last", 1, false, []string{"finish"}},
		{"three", 3, false, []string{"back", "resume", "finish"}},
		{"all", 0, true, []string{"start", "retry", "back", "resume", "finish"}},
		{"oversized", 99, false, []string{"start", "retry", "back", "resume", "finish"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			d.Trace.Last, d.Trace.All = tc.last, tc.all

			out, err := d.View(120, palette)
			if err != nil {
				t.Fatal(err)
			}

			for _, label := range []string{"start", "retry", "back", "resume", "finish"} {
				expected := palette.DimForeground

				for _, shown := range tc.shown {
					if label == shown {
						expected = palette.Foreground
					}
				}

				latest := label == "finish"
				if latest {
					expected = palette.Accent
				}
				// WHY: Label styles test edge identity independently of the route geometry.
				if !strings.Contains(ansi.Strip(out), label) {
					t.Fatalf("lost label %s", label)
				}

				at := strings.Index(out, label)

				styles := regexp.MustCompile(`\x1b\[[0-9;]*m`).FindAllString(out[:at], -1)
				if len(styles) == 0 {
					t.Fatal("label has no terminal style")
				}

				r, g, b, _ := expected.RGBA()

				prefix := fmt.Sprintf("\x1b[38;2;%d;%d;%d", r>>8, g>>8, b>>8)
				if !strings.HasPrefix(styles[len(styles)-1], prefix) {
					t.Fatalf("wrong history window for %s: %s", label, out)
				}
			}
		})
	}
}

func TestTransitionTraceRejectsInvalidHistory(t *testing.T) {
	for _, trace := range []diagram.TransitionTrace{
		{Steps: []int{-1}},
		{Steps: []int{5}},
		{Steps: []int{0, 0, 4}},
		{Steps: []int{0, 1}},
		{Steps: []int{4}, Last: -1},
		{Steps: []int{4}, All: true, Last: 1},
	} {
		d := tracedGraph()
		d.Trace = trace

		out, err := d.View(120, theme.New(theme.Dark))
		if err == nil || out != "" {
			t.Fatalf("invalid history rendered: %s, %v", out, err)
		}
	}
}

func TestTransitionTracePreservesGeometryAndCallerHistory(t *testing.T) {
	d := tracedGraph()
	d.Active = "B"
	d.Trace.Steps = []int{0, 1, 1}
	d.Trace.All = true

	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		palette := theme.New(mode)
		for _, width := range []int{40, 120} {
			with, err := d.View(width, palette)
			if err != nil {
				t.Fatal(err)
			}

			without := d
			without.Trace = diagram.TransitionTrace{}

			plain, err := without.View(width, palette)
			if err != nil {
				t.Fatal(err)
			}

			a, b := strings.Split(with, "\n"), strings.Split(plain, "\n")
			if len(a) != len(b) {
				t.Fatal("trace changed height")
			}

			for i := range a {
				if ansi.StringWidth(a[i]) != ansi.StringWidth(b[i]) {
					t.Fatal("trace changed width")
				}
			}

			if with == plain {
				t.Fatalf("trace invisible in %s at %d", mode, width)
			}

			if mode == theme.Mono && with != ansi.Strip(with) {
				t.Fatal("mono trace contains ANSI")
			}
		}
	}

	if len(d.Trace.Steps) != 3 || d.Trace.Steps[1] != 1 || d.Trace.Steps[2] != 1 {
		t.Fatal("history mutated")
	}
}

func TestTransitionTraceCanBeEmptyOrOmitActive(t *testing.T) {
	d := tracedGraph()
	d.Active = ""

	d.Trace.Steps = []int{0}
	if _, err := d.View(120, theme.New(theme.Dark)); err != nil {
		t.Fatal(err)
	}

	d.Trace = diagram.TransitionTrace{}
	if _, err := d.View(120, theme.New(theme.Dark)); err != nil {
		t.Fatal(err)
	}

	d = diagram.StateDiagram{Trace: diagram.TransitionTrace{Steps: []int{0}}}
	if out, err := d.View(120, theme.New(theme.Dark)); err == nil || out != "" {
		t.Fatal("empty graph accepted nonexistent history")
	}
}
