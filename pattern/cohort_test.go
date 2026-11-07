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
	m, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Data: d, Theme: theme.New(theme.Dark), MinRequests: 1000, Title: "Comparison"})
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
	m, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Data: d, Theme: theme.New(theme.Mono), MinRequests: 1000, Emit: func(value any) {
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
	if !strings.Contains(frame, "ERROR") || !strings.Contains(frame, "WARN") || !strings.Contains(frame, "W>=+2 E>=+3") {
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
	if m.Selected().MemberID != "host-095" || !strings.Contains(ansi.Strip(m.Inspector.View()), "ID: "+fixture.CandidateVersionID) || before == m.Panels.View() {
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
	if strings.Count(ansi.Strip(m.View()), "3=member") != 4 {
		t.Fatal("not all four panels include the selected member history")
	}
	m.Update(runtime.StepMsg{})
	if m.WindowIndex() != 11 || m.Selected().MemberID != "host-095" {
		t.Fatal("window advance lost selection")
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
	m, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Data: d, Theme: theme.New(theme.Mono), MinRequests: 1000})
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
