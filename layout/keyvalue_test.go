package layout_test

import (
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestWidthClassSnapsDown(t *testing.T) {
	for width, want := range map[int]layout.WidthClass{
		0: layout.Width40, 54: layout.Width40, 55: layout.Width55, 79: layout.Width55, 80: layout.Width80,
		119: layout.Width100, 120: layout.Width120, 159: layout.Width120, 400: layout.Width160,
	} {
		if got := layout.WidthClassOf(width); got != want {
			t.Fatalf("width %d: got %d want %d", width, got, want)
		}
	}

	if !layout.Width55.Stacked() || layout.Width80.Stacked() {
		t.Fatal("stacking threshold moved")
	}
}

func TestKeyValueAlignsStacksAndKeepsMissingDistinct(t *testing.T) {
	palette := theme.New(theme.Mono)
	pairs := []layout.Pair{
		{Label: "profile", Value: "dev"},
		{Label: "empty", Value: ""},
		{Label: "compat", State: text.ValueMissing},
	}

	if got, want := ansi.Strip(layout.KeyValue(pairs, 80, palette)), "profile  dev\nempty    \"\"\ncompat   missing"; got != want {
		t.Fatalf("got %q want %q", got, want)
	}

	if got, want := ansi.Strip(layout.KeyValue(pairs, 50, palette)), "profile: dev\nempty: \"\"\ncompat: missing"; got != want {
		t.Fatalf("got %q want %q", got, want)
	}

	// WidthClassOf snaps 56 to 79 columns down to Width55, which is not "55 columns and narrower".
	for width, want := range map[int]string{55: "profile: dev", 56: "profile  dev", 79: "profile  dev"} {
		if got := ansi.Strip(layout.KeyValue(pairs[:1], width, palette)); got != want {
			t.Fatalf("width %d: got %q want %q", width, got, want)
		}
	}

	long := []layout.Pair{{Label: "note", Value: strings.Repeat("word ", 12)}}
	for _, width := range []int{20, 60, 90} {
		for _, line := range strings.Split(layout.KeyValue(long, width, palette), "\n") {
			if ansi.StringWidth(line) > width {
				t.Fatalf("width %d: %q", width, line)
			}
		}
	}
}

func TestKeyValueNeverExceedsTheAllocation(t *testing.T) {
	ascii := theme.New(theme.Dark)
	ascii.ASCII = true

	themes := map[string]theme.Theme{"dark": theme.New(theme.Dark), "mono": theme.New(theme.Mono), "ascii": ascii}
	pairs := []layout.Pair{
		{Label: strings.Repeat("l", 20), Value: "x"},
		{Label: strings.Repeat("label ", 8), Value: strings.Repeat("v", 70)},
		{Label: "short", State: text.ValueMissing},
	}

	for name, palette := range themes {
		for _, width := range []int{20, 40, 55, 80, 120} {
			view := layout.KeyValue(pairs, width, palette)

			for _, line := range strings.Split(view, "\n") {
				if ansi.StringWidth(line) > width {
					t.Fatalf("%s/%d: %d cells: %q", name, width, ansi.StringWidth(line), ansi.Strip(line))
				}
			}

			if plain := ansi.Strip(view); !strings.Contains(plain, "x") || !strings.Contains(plain, "missing") {
				t.Fatalf("%s/%d lost a value: %q", name, width, plain)
			}
		}
	}
}
