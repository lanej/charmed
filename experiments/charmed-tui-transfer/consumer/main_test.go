package main

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func testFixture(t *testing.T) fixture {
	t.Helper()

	f, err := readFixture("fixture.json")
	if err != nil {
		t.Fatal(err)
	}

	return f
}

func testModel(t *testing.T, stateID string, size layout.Size) *model {
	t.Helper()

	m, err := newModel(testFixture(t), stateID, size, false, false)
	if err != nil {
		t.Fatal(err)
	}

	return m
}

func press(m *model, code rune) { m.Update(tea.KeyPressMsg{Code: code}) }

func TestFixtureReferencesAndPatchKeepIndependentEvidence(t *testing.T) {
	f := testFixture(t)
	ready := f.States[2].Observation
	retained := f.States[3].Observation

	replacement := f.States[4].Observation
	if ready == nil || retained == nil || replacement == nil {
		t.Fatal("missing resolved observation")
	}

	if *ready.CompletedChecks != 47 || *retained.CompletedChecks != 47 || *replacement.CompletedChecks != 48 {
		t.Fatal("counts not preserved across reference/patch")
	}

	if len(ready.Regions) != 3 || len(retained.Regions) != 3 || len(replacement.Regions) != 1 {
		t.Fatal("patch did not replace regions independently")
	}

	if *replacement.ChecksDue != 60 || len(replacement.Intervals) != 6 {
		t.Fatal("patch lost unmentioned fields")
	}

	if ready.Intervals[1].CompletedChecks == nil || *ready.Intervals[1].CompletedChecks != 0 {
		t.Fatal("complete zero lost")
	}

	if ready.Intervals[2].CompletedChecks != nil || ready.Intervals[3].CompletedChecks != nil {
		t.Fatal("absent counts became zero")
	}

	if ready.Regions[1].Availability != "available" || ready.Regions[1].Health != "failing" {
		t.Fatal("availability conflated with health")
	}

	if ready.Regions[2].LastVerified != nil {
		t.Fatal("unknown verification time invented")
	}

	replacement.Regions[0].Name = "changed"
	if ready.Regions[0].Name != "us-west" {
		t.Fatal("patch aliases original")
	}
}

func TestStateMeaningAndCoverage(t *testing.T) {
	cases := []struct {
		id     string
		want   []string
		absent string
	}{
		{"waiting", []string{"Waiting for first observation", "not a confirmed empty"}, "0 / 0"},
		{"empty", []string{"0 / 0 checks", "Confirmed empty: no interval", "Confirmed empty: no region"}, "Waiting"},
		{"ready", []string{"47 / 60 checks", "09:15-09:30", "missing", "unknown", "3 so far (partial)", "eu-north / available / failing", "1m0s ago"}, "STALE"},
		{"failed_refresh", []string{"Refresh failed / retained 10:16 UTC", "47 / 60 checks", "verification collector unavailable"}, "Observed 10:17"},
		{"replacement", []string{"48 / 60 checks", "0s ago", "us-west / available / passing"}, "eu-north"},
	}
	for _, tc := range cases {
		t.Run(tc.id, func(t *testing.T) {
			m := testModel(t, tc.id, layout.Size{Width: 132, Height: 40})

			body := ansi.Strip(m.View().Content)
			for _, want := range tc.want {
				if !strings.Contains(body, want) {
					t.Errorf("missing %q in\n%s", want, body)
				}
			}

			if strings.Contains(body, tc.absent) {
				t.Errorf("unexpected %q", tc.absent)
			}
		})
	}
}

func TestActualKeyRoutingRetainsIndependentPositions(t *testing.T) {
	m := testModel(t, "ready", layout.Size{Width: 40, Height: 18})
	for range 4 {
		press(m, tea.KeyDown)
	}

	overview := m.dashboard.Overview.YOffset()
	if overview == 0 {
		t.Fatal("overview did not overflow and scroll")
	}

	press(m, 'i')

	for range 8 {
		press(m, tea.KeyDown)
	}

	inspection := m.dashboard.Inspection.YOffset()
	if inspection == 0 || !m.dashboard.Inspecting() {
		t.Fatal("inspection did not scroll")
	}

	for _, code := range []rune{'r', 'a', 'm'} {
		press(m, code)

		if m.dashboard.Overview.YOffset() != overview || m.dashboard.Inspection.YOffset() != inspection {
			t.Fatalf("%c changed independent positions", code)
		}
	}

	press(m, tea.KeyEscape)

	if m.dashboard.Inspecting() || m.dashboard.Overview.YOffset() != overview {
		t.Fatal("back lost overview")
	}

	press(m, 'i')

	if m.dashboard.Inspection.YOffset() != inspection {
		t.Fatal("reopen lost inspection")
	}

	m.Update(tea.WindowSizeMsg{Width: 40, Height: 19})

	if m.dashboard.Overview.YOffset() != overview || m.dashboard.Inspection.YOffset() != inspection {
		t.Fatal("small resize lost place")
	}

	press(m, '5')

	if m.fixture.States[m.index].ID != "replacement" || !m.dashboard.Inspecting() {
		t.Fatal("snapshot replacement lost state")
	}

	if m.dashboard.Inspection.YOffset() != inspection {
		t.Fatal("replacement reset position despite remaining overflow")
	}

	press(m, tea.KeyEscape)

	for range 100 {
		press(m, tea.KeyDown)
	}

	if !strings.Contains(ansi.Strip(m.View().Content), "exact verification times") {
		t.Fatal("overview bottom unreachable")
	}

	press(m, 'i')

	for range 100 {
		press(m, tea.KeyDown)
	}

	if !strings.Contains(ansi.Strip(m.View().Content), "Confirmed:") {
		t.Fatal("inspection bottom unreachable")
	}

	m.Update(tea.WindowSizeMsg{Width: 132, Height: 40})

	if m.dashboard.Overview.YOffset() != 0 {
		t.Fatal("overview did not clamp after content fits")
	}
}

func TestGeometryAndFallbacksThroughApp(t *testing.T) {
	for _, size := range []layout.Size{{Width: 132, Height: 40}, {Width: 80, Height: 24}, {Width: 40, Height: 18}} {
		for _, id := range []string{"waiting", "empty", "ready", "failed_refresh", "replacement"} {
			m := testModel(t, id, size)
			for _, sequence := range []string{"", "?", "?,i", "a", "m", "esc"} {
				if err := m.replay(sequence); err != nil {
					t.Fatal(err)
				}

				view := m.View()

				lines := strings.Split(view.Content, "\n")
				if len(lines) != size.Height {
					t.Fatalf("%s %v height %d", id, size, len(lines))
				}

				for _, line := range lines {
					if ansi.StringWidth(line) != size.Width {
						t.Fatalf("%s %v escaped width: %q", id, size, line)
					}
				}

				if view.ForegroundColor == nil || view.BackgroundColor == nil {
					t.Fatal("frame defaults are unspecified")
				}
			}
		}
	}

	m := testModel(t, "ready", layout.Size{Width: 80, Height: 24})
	press(m, 'a')

	for _, r := range ansi.Strip(m.View().Content) {
		if r > 127 {
			t.Fatalf("ASCII frame contains %q", r)
		}
	}

	press(m, 'a')
	press(m, 'm')

	if m.dashboard.Theme.Mode != theme.Mono || m.dashboard.Theme.ASCII {
		t.Fatal("monochrome unnecessarily changed glyph repertoire")
	}
}

func TestInvalidInputsFail(t *testing.T) {
	f := testFixture(t)
	if _, err := newModel(f, "absent", layout.Size{Width: 80, Height: 24}, false, false); err == nil {
		t.Fatal("unknown state accepted")
	}

	if _, err := f.resolve("absent", map[string]bool{}); err == nil {
		t.Fatal("unknown reference accepted")
	}

	f.States[0].ObservationRef = "waiting"
	if _, err := f.resolve("waiting", map[string]bool{}); err == nil {
		t.Fatal("cycle accepted")
	}

	m := testModel(t, "ready", layout.Size{Width: 80, Height: 24})
	if err := m.replay("invalid"); err == nil {
		t.Fatal("invalid replay accepted")
	}
}
