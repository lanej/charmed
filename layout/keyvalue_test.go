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

	long := []layout.Pair{{Label: "note", Value: strings.Repeat("word ", 12)}}
	for _, width := range []int{20, 60, 90} {
		for _, line := range strings.Split(layout.KeyValue(long, width, palette), "\n") {
			if ansi.StringWidth(line) > width {
				t.Fatalf("width %d: %q", width, line)
			}
		}
	}
}
