package pattern_test

import (
	"strings"
	"testing"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/attention"
	"github.com/lanej/charmed/cohort"
	"github.com/lanej/charmed/examples/deployment/fixture"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/theme"
)

func comparison(t *testing.T) (*pattern.CohortComparison, cohort.Data) {
	t.Helper()
	d, err := fixture.Data("isolated-hosts")
	if err != nil {
		t.Fatal(err)
	}

	m, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Data: d, ReferenceID: d.Cohorts[0].ID, Theme: theme.New(theme.Dark), MinRequests: 1000, Title: "Comparison"})
	if err != nil {
		t.Fatal(err)
	}
	m.Init()
	return m, d
}

func TestAttentionIsComposedOptionalAndIndependentOfDisplaySigma(t *testing.T) {
	d, _ := fixture.Data("isolated-hosts")
	var transitions []attention.Changed
	var current *pattern.CohortComparison

	m, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Data: d, ReferenceID: d.Cohorts[0].ID, Theme: theme.New(theme.Mono), MinRequests: 1000, Emit: func(value any) {
		if change, ok := value.(attention.Changed); ok {
			if current != nil {
				if a, ok := current.Attention("5xx"); !ok || a.Level == "" {
					t.Fatal("callback observed a partially updated comparison")
				}
			}
			transitions = append(transitions, change)
		}
	}})
	if err != nil {
		t.Fatal(err)
	}
	current = m
	m.Init()
	a, _ := m.Attention("5xx")
	if a.Level != attention.Off {
		t.Fatal("descriptive deviation generated attention without boundaries")
	}
	policies := fixture.Boundaries()
	if err := m.SetBoundaries(policies); err != nil {
		t.Fatal(err)
	}
	a, _ = m.Attention("5xx")
	if a.Level != attention.Error || a.Hosts != attention.Error || a.Fleet.Level != attention.Warn || a.Errors != 5 {
		t.Fatalf("wrong composed attention: %+v", a)
	}
	*policies["5xx"].Hosts.Error.Upper = 10000
	selected := m.Selected()
	before := len(transitions)
	_ = m.SetSigma(1)
	m.Resize(layout.Size{Width: 40, Height: 16})
	m.Resize(layout.Size{Width: 132, Height: 40})
	a, _ = m.Attention("5xx")
	if a.Level != attention.Error || m.Selected() != selected || len(transitions) != before {
		t.Fatal("rendering, sigma, resize or caller mutation changed attention")
	}
	frame := ansi.Strip(m.View())
	if !strings.Contains(frame, "ERROR") || !strings.Contains(frame, "+0.203pp") || !strings.Contains(frame, "Warn>=+0.1 Error>=+0.3") {
		t.Fatal("composed boundary state or explanation missing")
	}
	_ = m.SetWindow(0)
	a, _ = m.Attention("5xx")
	if a.Level != attention.Normal {
		t.Fatal("attention did not recover in healthy window")
	}
	recovered := false
	for _, event := range transitions {
		if event.Metric == "5xx" && event.Scope == "hosts" && event.Previous == attention.Error && event.Level == attention.Normal {
			recovered = true
		}
	}
	if !recovered {
		t.Fatal("missing typed recovery event")
	}
	beforeFrame := m.View()
	if m.SetBoundaries(map[string]cohort.ComparisonPolicy{"missing": {}}) == nil || m.View() != beforeFrame {
		t.Fatal("invalid policy update mutated screen")
	}
	if err := m.SetBoundaries(nil); err != nil {
		t.Fatal(err)
	}
	a, _ = m.Attention("5xx")
	if a.Level != attention.Off || strings.Contains(ansi.Strip(m.View()), "ATTN") {
		t.Fatal("disabled boundaries retained attention")
	}
}

func TestComparisonLinksSelectionAndRetainsStateAcrossResize(t *testing.T) {
	m, _ := comparison(t)
	if m.Selected().MemberID != "host-096" || m.Selected().CohortID != fixture.CandidateVersionID || m.WindowIndex() != 11 || m.MetricID() != "5xx" {
		t.Fatal("initial view did not select largest baseline deviation in the latest window")
	}
	before := m.Panels.View()
	m.Update(tea.KeyPressMsg{Code: tea.KeyDown})

	if m.Selected().MemberID != "host-095" || !strings.Contains(ansi.Strip(m.View()), "Cohort "+fixture.CandidateVersionID) || before == m.Panels.View() {
		t.Fatal("selection did not coordinate plots and inspector")
	}
	m.Update(tea.KeyPressMsg{Code: tea.KeyLeft})
	m.Update(press('1', "1"))
	m.Update(press(']', "]"))
	for _, size := range []layout.Size{{Width: 40, Height: 16}, {Width: 150, Height: 46}} {
		m.Resize(size)
		if m.Selected().MemberID != "host-095" || m.WindowIndex() != 10 || m.Sigma() != 1 || m.MetricID() != "2xx" {
			t.Fatal("resize or control change lost linked state")
		}
	}
	if len(m.Inspector.Rows()) != 4 || !strings.Contains(ansi.Strip(m.View()), fixture.CandidateVersionID+" mean") {
		t.Fatal("inspector table or direct version labels are missing")
	}
	m.Update(runtime.StepMsg{})
	if m.WindowIndex() != 11 || m.Selected().MemberID != "host-095" {
		t.Fatal("window advance lost selection")
	}
}

func TestMemberInspectorTableSelectsMetricsAndHasNoLegends(t *testing.T) {
	m, _ := comparison(t)
	if len(m.Inspector.Rows()) != 4 || m.Inspector.Cursor() != 3 {
		t.Fatal("inspector did not present the four metric rows")
	}
	selected := m.Selected()
	m.Update(tea.KeyPressMsg{Code: tea.KeyTab})
	m.Update(tea.KeyPressMsg{Code: tea.KeyTab})
	if !m.Inspector.Focused() || m.Table.Focused() {
		t.Fatal("inspector table did not receive focus")
	}
	m.Update(tea.KeyPressMsg{Code: tea.KeyUp})
	if m.MetricID() != "4xx" || m.Inspector.Cursor() != 2 || m.Selected() != selected {
		t.Fatal("inspector selection did not coordinate the active metric")
	}
	m.Resize(layout.Size{Width: 40, Height: 40})
	m.Resize(layout.Size{Width: 132, Height: 40})
	if m.MetricID() != "4xx" || m.Inspector.Cursor() != 2 || m.Selected() != selected {
		t.Fatal("resize lost table selection")
	}
	frame := ansi.Strip(m.View())
	for _, legend := range []string{"History %:", "3=member", "X = this member", "W/E =", "R=reference", "C=candidate", "1/2/3 bands", "distance from reference hosts"} {
		if strings.Contains(frame, legend) {
			t.Fatalf("legend remains: %s", legend)
		}
	}
	if strings.Contains(frame, "Fleet ") || strings.Contains(frame, "Delta ") {
		t.Fatal("repeated percentage-summary labels remain")
	}

	if !strings.Contains(frame, "METRIC") || !strings.Contains(frame, "VALUE") || !strings.Contains(frame, "COHORT") {
		t.Fatal("table data headers are missing")
	}
}

func TestDeltaTitleScaleDoesNotFollowTheSelectedWindow(t *testing.T) {
	m, _ := comparison(t)
	if !strings.Contains(ansi.Strip(m.View()), "delta +/-0.5pp") {
		t.Fatal("missing shared title-bar domain")
	}
	if err := m.SetWindow(0); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(ansi.Strip(m.View()), "delta +/-0.5pp") {
		t.Fatal("window selection rescaled the comparison")
	}
}

func TestUpdatesAreAtomicAndFollowLatestOnlyWhenSelected(t *testing.T) {
	m, d := comparison(t)
	before := m.View()
	invalid := d
	invalid.Metrics = nil
	if m.SetData(invalid) == nil || m.View() != before {
		t.Fatal("invalid update mutated the screen")
	}
	w := d.Windows[len(d.Windows)-1]
	w.Start, w.End = w.End, w.End.Add(5*time.Minute)
	d.Windows = append(d.Windows, w)
	if err := m.SetData(d); err != nil {
		t.Fatal(err)
	}
	if m.WindowIndex() != 12 {
		t.Fatal("latest view did not follow appended window")
	}
	if err := m.SetWindow(2); err != nil {
		t.Fatal(err)
	}
	w.Start, w.End = w.End, w.End.Add(5*time.Minute)
	d.Windows = append(d.Windows, w)
	if err := m.SetData(d); err != nil {
		t.Fatal(err)
	}
	if m.WindowIndex() != 2 {
		t.Fatal("append interrupted historical inspection")
	}
	if m.SetWindow(-1) == nil || m.SetSigma(4) == nil || m.SetMetric("missing") == nil {
		t.Fatal("invalid control value accepted")
	}
	d.Windows[2].Samples[0].Counts["5xx"] = 999999
	if m.SetWindow(2) != nil {
		t.Fatal("caller mutation affected owned data")
	}
}

func TestSparseComparisonDoesNotInventSpread(t *testing.T) {
	d, _ := fixture.Data("insufficient-samples")

	m, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Data: d, ReferenceID: d.Cohorts[0].ID, Theme: theme.New(theme.Mono), MinRequests: 1000})
	if err != nil {
		t.Fatal(err)
	}
	frame := ansi.Strip(m.View())
	if strings.Contains(frame, "NaN") || strings.Contains(frame, "+Inf") || !strings.Contains(frame, "baseline s unavailable") {
		t.Fatal("sparse data yielded misleading spread")
	}
	d.Windows = nil
	if err := m.SetData(d); err != nil {
		t.Fatal(err)
	}
	if m.Selected().MemberID != "" || !strings.Contains(ansi.Strip(m.View()), "No windows") {
		t.Fatal("empty update retained phantom measurements")
	}
}
