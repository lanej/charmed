package chart_test

import (
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/theme"
)

func TestBucketBarsCoverageControlsComparison(t *testing.T) {
	bars := chart.BucketBars{Buckets: []chart.CountBucket{
		{Label: "23:00", Count: 100, State: chart.BucketComplete},
		{Label: "00:00", Count: 9000, State: chart.BucketPartial},
		{Label: "01:00", Count: 123, State: chart.BucketMissing},
		{Label: "02:00", Count: 456},
		{Label: "03:00", State: chart.BucketComplete},
	}}
	view := ansi.Strip(bars.View(40, theme.New(theme.Dark)))

	lines := strings.Split(view, "\n")
	if !strings.Contains(lines[0], strings.Repeat("━", 27)) || !strings.Contains(lines[0], "100") {
		t.Fatalf("partial interval changed complete scale: %s", view)
	}

	for _, want := range []string{"9000 so far (partial)", "missing", "unknown"} {
		if !strings.Contains(view, want) {
			t.Fatalf("missing coverage label %q: %s", want, view)
		}
	}

	for _, line := range lines[1:] {
		if strings.Contains(line, "━") {
			t.Fatalf("incomplete or zero interval received a bar: %s", line)
		}
	}

	if strings.Contains(view, "123") || strings.Contains(view, "456") {
		t.Fatal("unobserved count was presented as evidence")
	}
}

func TestBucketBarsRetainsExactCountsAtNarrowWidths(t *testing.T) {
	bars := chart.BucketBars{Buckets: []chart.CountBucket{
		{Label: "界 / 00:00 UTC", Count: 18446744073709551615, State: chart.BucketComplete},
		{Label: "01:00 UTC", Count: 1, State: chart.BucketPartial},
	}}
	for _, width := range []int{2, 20, 40, 80, 132} {
		view := ansi.Strip(bars.View(width, theme.New(theme.Dark)))
		for _, line := range strings.Split(view, "\n") {
			if ansi.StringWidth(line) > width {
				t.Fatalf("width %d escaped: %q", width, line)
			}
		}

		joined := strings.ReplaceAll(view, "\n", "")
		if !strings.Contains(joined, "18446744073709551615") || !strings.Contains(joined, "partial") {
			t.Fatalf("width %d lost observations: %s", width, view)
		}
	}

	if got := bars.View(1, theme.New(theme.Dark)); got != "!" {
		t.Fatalf("impossible grapheme allocation was not signaled: %q", got)
	}

	palette := theme.New(theme.Dark)
	palette.ASCII = true

	view := bars.View(80, palette)
	if strings.Contains(view, "━") || !strings.Contains(view, "#") {
		t.Fatal("explicit ASCII fallback did not replace bar glyphs")
	}
}
