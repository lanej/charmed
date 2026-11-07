package pattern_test

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/theme"
)

func press(code rune, text string) tea.KeyPressMsg { return tea.KeyPressMsg{Code: code, Text: text} }
func TestDashboardCoordinatesSelectionResizeAndBackgroundData(t *testing.T) {
	m := pattern.NewDashboard(pattern.DashboardOptions{Theme: theme.New(theme.Dark), Services: pattern.SampleServices()})
	m.Init()
	if !m.Filter.Focused() {
		t.Fatal("filter must own initial focus")
	}
	m.Update(tea.KeyPressMsg{Code: tea.KeyTab})
	if !m.Table.Focused() || m.Filter.Focused() {
		t.Fatal("Tab did not transfer focus to the table")
	}
	m.Update(tea.KeyPressMsg{Code: tea.KeyDown})
	if m.SelectedID != "events" {
		t.Fatalf("selection = %q", m.SelectedID)
	}
	before := m.Plot.View()
	m.Update(runtime.StepMsg{})
	if m.Step != 1 || m.Plot.View() == before {
		t.Fatal("unfocused chart did not receive its background update")
	}
	m.Resize(layout.Size{Width: 40, Height: 16})
	if m.SelectedID != "events" {
		t.Fatal("resize lost stable selection")
	}
	if !strings.Contains(ansi.Strip(m.View()), "Event stream") {
		t.Fatal("linked plot no longer reflects selection")
	}
	m.Resize(layout.Size{Width: 120, Height: 34})
	if !strings.Contains(ansi.Strip(m.Details.View()), "Event stream") {
		t.Fatal("inspector disagrees with selection")
	}
}
func TestInputQIsTextAndEmptyFiltersRemainNavigable(t *testing.T) {
	m := pattern.NewDashboard(pattern.DashboardOptions{Theme: theme.New(theme.Dark), Services: pattern.SampleServices()})
	m.Init()
	m.Update(press('q', "q"))
	if m.Filter.Value() != "q" || m.SelectedID != "" {
		t.Fatal("q should edit the filter and produce an empty result")
	}
	m.Update(tea.KeyPressMsg{Code: tea.KeyTab})
	m.Update(tea.KeyPressMsg{Code: tea.KeyDown})
	if m.SelectedID != "" {
		t.Fatal("an empty table selected a nonexistent row")
	}
	m.Update(tea.KeyPressMsg{Code: tea.KeyTab, Mod: tea.ModShift})
	m.Update(tea.KeyPressMsg{Code: 'u', Mod: tea.ModCtrl})
	if m.Filter.Value() != "" || m.SelectedID != "api" {
		t.Fatal("clearing an empty filter did not restore a valid selection")
	}
}

func TestDashboardUsesStatusAtomsWithoutRedundantHeadings(t *testing.T) {
	m := pattern.NewDashboard(pattern.DashboardOptions{Theme: theme.New(theme.Dark), Services: pattern.SampleServices()})
	frame := ansi.Strip(m.View())
	for _, heading := range []string{"SERVICE", "STATE", "Services", "Inspector"} {
		if strings.Contains(frame, heading) {
			t.Errorf("redundant heading %q remains", heading)
		}
	}
	for _, label := range []string{"API gateway", "Event stream", "Job workers", "Storage", "Scheduler", "OK", "WARN"} {
		if !strings.Contains(frame, label) {
			t.Errorf("missing service or shared status label %q", label)
		}
	}
	rows := m.Table.Rows()
	if !strings.Contains(rows[2][1], "\x1b[") {
		t.Error("warning status has no native styling")
	}
	m = pattern.NewDashboard(pattern.DashboardOptions{Theme: theme.New(theme.Mono), Services: []pattern.Service{{ID: "unknown", Name: "Unclassified", Status: "unrecognized"}}})
	if !strings.Contains(ansi.Strip(m.View()), "[UNKNOWN]") {
		t.Error("unrecognized health must use the explicit unknown status atom")
	}
}
