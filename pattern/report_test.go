package pattern_test

import (
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

func reportBody(text string) func(int) string { return func(int) string { return text } }

func samplePage(counts ...feedback.EvidenceCount) pattern.ReportPage {
	now := time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)

	return pattern.ReportPage{
		Title: "Service processes", Context: "api · 3 hosts", Counts: counts,
		Panels: []pattern.ReportPanel{
			{Title: "Hosts", Body: reportBody("web1west   observed   live   45s\nweb2west   observed   cached 4m"), Weight: 3},
			{Title: "Processes", Beside: true, Weight: 2, Body: func(width int) string {
				return "db:db RUNNING a-rather-long-process-name-that-must-wrap-" + strconv.Itoa(width)
			}},
			{Title: "Limits", Body: reportBody("none reported")},
			{Title: "Hidden", Body: reportBody("")},
		},
		Footer: feedback.EvidenceFooter{Source: "supervisor", Scope: "api", ObservedAt: now.Add(-14 * time.Second), Now: now},
	}
}

func healthy() []feedback.EvidenceCount {
	return []feedback.EvidenceCount{{State: feedback.EvidenceHealthy, N: 3}}
}

func themes() []theme.Theme {
	ascii := theme.New(theme.Mono)
	ascii.ASCII = true

	return []theme.Theme{theme.New(theme.Dark), theme.New(theme.Light), theme.New(theme.Contrast), theme.New(theme.Mono), ascii}
}

func TestReportPageFitsEveryWidthAndTheme(t *testing.T) {
	for _, th := range themes() {
		for width := 20; width <= 200; width += 5 {
			view := samplePage(healthy()...).View(width, th)
			for _, line := range strings.Split(view, "\n") {
				if ansi.StringWidth(line) > width {
					t.Fatalf("%s %d: line wider than the width: %q", th.Mode, width, line)
				}
			}

			plain := ansi.Strip(view)
			if !strings.Contains(plain, "HEALTHY") && width >= 40 {
				t.Fatalf("%s %d lost the state word:\n%s", th.Mode, width, plain)
			}

			if th.ASCII && strings.ContainsFunc(plain, func(r rune) bool { return r > 127 && r != '·' }) {
				t.Fatalf("ASCII page carries non-ASCII glyphs at %d:\n%s", width, plain)
			}
		}
	}
}

func TestReportPageSideBySideOnlyWhenItFits(t *testing.T) {
	th := theme.New(theme.Mono)

	for width, wide := range map[int]bool{40: false, 55: false, 60: false, 62: true, 80: true, 120: true} {
		plain := ansi.Strip(samplePage(healthy()...).View(width, th))
		lines := strings.Split(plain, "\n")

		side := false

		for _, line := range lines {
			if strings.Contains(line, "Hosts") && strings.Contains(line, "Processes") {
				side = true
			}
		}

		if side != wide {
			t.Fatalf("width %d: side-by-side=%v, want %v\n%s", width, side, wide, plain)
		}
	}
}

func TestReportPageStacksAtOrBelowTheStackWidth(t *testing.T) {
	page := samplePage(healthy()...)
	page.Panels[1].Min = 10
	page.Panels[0].Min = 10

	plain := ansi.Strip(page.View(layout.StackWidth, theme.New(theme.Mono)))
	for _, line := range strings.Split(plain, "\n") {
		if strings.Contains(line, "Hosts") && strings.Contains(line, "Processes") {
			t.Fatalf("panels shared a row at the stack width:\n%s", plain)
		}
	}

	wide := ansi.Strip(page.View(layout.StackWidth+1, theme.New(theme.Mono)))
	if !strings.Contains(wide, "Hosts") {
		t.Fatal(wide)
	}
}

func TestReportPageNeverReadsHealthierThanItsEvidence(t *testing.T) {
	cases := map[string][]feedback.EvidenceCount{
		"none":         nil,
		"zero":         {{State: feedback.EvidenceHealthy, N: 0}},
		"unknown":      {{State: feedback.EvidenceHealthy, N: 4}, {State: feedback.EvidenceUnknown, N: 1}},
		"insufficient": {{State: feedback.EvidenceHealthy, N: 4}, {State: feedback.EvidenceInsufficient, N: 1}},
		"stale":        {{State: feedback.EvidenceHealthy, N: 4}, {State: feedback.EvidenceStale, N: 1}},
		"truncated":    {{State: feedback.EvidenceHealthy, N: 4}, {State: feedback.EvidenceTruncated, N: 1}},
		"gap":          {{State: feedback.EvidenceHealthy, N: 4}, {State: feedback.EvidenceAttributionGap, N: 1}},
		"unrecognized": {{State: feedback.EvidenceHealthy, N: 4}, {State: feedback.Evidence("FINE"), N: 1}},
		"down":         {{State: feedback.EvidenceHealthy, N: 4}, {State: feedback.EvidenceDown, N: 1}},
	}

	for name, counts := range cases {
		page := samplePage(counts...)
		if page.Overall() == feedback.EvidenceHealthy {
			t.Fatalf("%s: overall is healthy", name)
		}

		for _, th := range themes() {
			for _, width := range []int{40, 55, 80, 120} {
				plain := ansi.Strip(page.View(width, th))
				lines := strings.Split(plain, "\n")
				header := lines[0]
				footer := plain[strings.LastIndex(plain, "source:"):]

				if strings.Contains(header, "HEALTHY") || strings.Contains(footer, "HEALTHY") {
					t.Fatalf("%s %s %d reads healthy:\n%s", name, th.Mode, width, plain)
				}
			}
		}
	}

	if got := samplePage(healthy()...).Overall(); got != feedback.EvidenceHealthy {
		t.Fatal(got)
	}
}

func TestReportPageViewSizeKeepsChromeAndSaysWhatItDropped(t *testing.T) {
	th := theme.New(theme.Mono)
	page := samplePage(healthy()...)
	page.Panels[0].Body = reportBody(strings.Repeat("row\n", 30))

	for _, size := range []layout.Size{{Width: 80, Height: 12}, {Width: 40, Height: 8}, {Width: 120, Height: 40}} {
		view := ansi.Strip(page.ViewSize(size, th))

		lines := strings.Split(view, "\n")
		if len(lines) != size.Height {
			t.Fatalf("%v: %d lines", size, len(lines))
		}

		for _, line := range lines {
			if ansi.StringWidth(line) > size.Width {
				t.Fatalf("%v: %q", size, line)
			}
		}

		if !strings.Contains(view, "HEALTHY") || !strings.Contains(view, "source:") {
			t.Fatalf("%v lost chrome:\n%s", size, view)
		}

		if size.Height < 30 && !strings.Contains(view, "more") {
			t.Fatalf("%v hid rows without saying so:\n%s", size, view)
		}
	}
}
