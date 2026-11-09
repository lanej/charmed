package chart_test

import (
	"math"
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestWaffleAllocatesCellsAndBreaksTiesInInputOrder(t *testing.T) {
	w := chart.Waffle{Parts: []chart.WafflePart{{Count: 1}, {Count: 1}, {Count: 1}}}

	got := ansi.Strip(w.View(layout.Size{Width: 9, Height: 2}, theme.New(theme.Mono)))
	if got != "A A A A B\nB B C C C" {
		t.Fatalf("equal thirds in ten cells: %q", got)
	}

	w.Parts = []chart.WafflePart{{Count: 92}, {Count: 6}, {Count: 2}}

	got = ansi.Strip(w.View(layout.Size{Width: 19, Height: 10}, theme.New(theme.Dark)))
	if strings.Count(got, "■") != 92 || strings.Count(got, "▧") != 6 || strings.Count(got, "▨") != 2 {
		t.Fatalf("100-cell counts differ from observations: %q", got)
	}
}

func TestWaffleHeaderRetainsRareCategoriesAndDropsGridFirst(t *testing.T) {
	w := chart.Waffle{Parts: []chart.WafflePart{
		{Label: "ready", Count: 999, Tone: text.Success},
		{Label: "failed", Count: 1, Tone: text.Danger},
	}}
	theme := theme.New(theme.Mono)

	wide := ansi.Strip(w.Header("Fleet", 80, theme))
	if !strings.Contains(wide, "B failed 1") || !strings.Contains(wide, "["+strings.Repeat("A ", 19)+"A]") {
		t.Fatalf("exact rare count or compact grid missing: %q", wide)
	}

	narrow := ansi.Strip(w.Header("Fleet", 34, theme))
	if !strings.Contains(narrow, "B failed 1") || strings.Contains(narrow, "[") {
		t.Fatalf("grid crowded out exact counts: %q", narrow)
	}

	for width := 0; width <= 100; width++ {
		got := w.Header("Fleet\nstatus", width, theme)
		if lipgloss.Width(got) != width || (width > 0 && lipgloss.Height(got) != 1) {
			t.Fatalf("header escaped width %d: %q", width, got)
		}
	}
}

func TestWaffleMissingAndInvalidCountsAreExplicit(t *testing.T) {
	for _, tc := range []struct {
		name  string
		parts []chart.WafflePart
		want  string
	}{
		{name: "empty", want: "No observations"},
		{name: "zero", parts: []chart.WafflePart{{Count: 0}}, want: "No observations"},
		{name: "negative", parts: []chart.WafflePart{{Count: -1}}, want: "Invalid waffle"},
		{name: "overflow", parts: []chart.WafflePart{{Count: math.MaxInt}, {Count: 1}}, want: "Invalid waffle"},
		{name: "too many categories", parts: make([]chart.WafflePart, 27), want: "Invalid waffle"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			w := chart.Waffle{Parts: tc.parts}

			got := ansi.Strip(w.View(layout.Size{Width: 20, Height: 2}, theme.New(theme.Mono)))
			if !strings.Contains(got, tc.want) || w.Legend(theme.New(theme.Mono)) != tc.want {
				t.Fatalf("unexpected grid %q or legend", got)
			}
		})
	}
}

func TestWaffleZeroCategoriesNeverReceiveCells(t *testing.T) {
	w := chart.Waffle{Parts: []chart.WafflePart{{Count: 0}, {Count: math.MaxInt}, {Count: 0}}}
	for _, size := range []layout.Size{{Width: 1, Height: 1}, {Width: 20, Height: 1}, {Width: 100, Height: 100}} {
		got := ansi.Strip(w.View(size, theme.New(theme.Mono)))
		if strings.ContainsAny(got, "AC") || strings.Count(got, "B") != ((size.Width+1)/2)*size.Height {
			t.Fatal("zero-count category received cells or allocation lost")
		}
	}
}

func TestWaffleRejectsOversizedGeometryWithoutAllocatingIt(t *testing.T) {
	w := chart.Waffle{Parts: []chart.WafflePart{{Count: 1}}}

	for _, size := range []layout.Size{{Width: math.MaxInt, Height: 1}, {Width: 1, Height: math.MaxInt}, {Width: 101, Height: 100}} {
		if got := w.View(size, theme.New(theme.Mono)); got != "Waffle too large" {
			t.Fatalf("unexpected oversized fallback: %q", got)
		}
	}

	for _, size := range []layout.Size{{Width: 0, Height: 1}, {Width: 1, Height: -1}} {
		if got := w.View(size, theme.New(theme.Mono)); got != "" {
			t.Fatalf("nonpositive geometry produced output: %q", got)
		}
	}
}

func TestWaffleStatusHeaderKeepsExactCountsBesideMicroGrid(t *testing.T) {
	w := chart.Waffle{Parts: []chart.WafflePart{{Label: "ready", Count: 92, Tone: text.Success}, {Label: "pending", Count: 3, Tone: text.Info}, {Label: "failed", Count: 5, Tone: text.Danger}}}
	for _, mode := range []theme.Mode{theme.Dark, theme.Mono} {
		got := ansi.Strip(w.StatusHeader("PRODUCTION", layout.Size{Width: 80, Height: 2}, theme.New(mode)))
		if lipgloss.Width(got) != 80 || lipgloss.Height(got) != 2 || !strings.Contains(got, "100 total") || !strings.Contains(got, "failed 5") {
			t.Fatalf("micro-grid header lost geometry or exact summary: %q", got)
		}

		lines := strings.Split(got, "\n")
		if !strings.Contains(lines[0], "PRODUCTION") || !strings.Contains(lines[1], "ready 92") {
			t.Fatalf("title and summary do not sit beside the grid: %q", got)
		}
	}

	narrow := ansi.Strip(w.StatusHeader("PRODUCTION", layout.Size{Width: 42, Height: 2}, theme.New(theme.Mono)))
	if strings.Contains(narrow, "A A") || !strings.Contains(narrow, "failed 5") {
		t.Fatalf("narrow header should retain full summary after dropping grid: %q", narrow)
	}

	if got := w.StatusHeader("Fleet", layout.Size{Width: 80, Height: 1}, theme.New(theme.Mono)); got != w.Header("Fleet", 80, theme.New(theme.Mono)) {
		t.Fatal("one-row status header should use the inline form")
	}
}
