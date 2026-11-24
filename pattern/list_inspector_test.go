package pattern_test

import (
	"math"
	"strings"
	"testing"

	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

func inspectorRecord(id, body string) pattern.InspectorRecord {
	return pattern.InspectorRecord{
		ID:  id,
		Row: func(int, bool) string { return body },
		Detail: func(int) string {
			return "Details for " + id + "\n" + strings.Repeat("retained evidence\n", 40) + "end of " + id
		},
	}
}

func newInspector(t *testing.T, records []pattern.InspectorRecord, size layout.Size) *pattern.ListInspector {
	t.Helper()

	m, err := pattern.NewListInspector(records, size, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	return m
}

func TestListInspectorReplacementIsAtomicAndRetainsIdentity(t *testing.T) {
	t.Parallel()

	records := []pattern.InspectorRecord{inspectorRecord("a", "row a"), inspectorRecord("b", "row b")}
	m := newInspector(t, records, layout.Size{Width: 100, Height: 12})
	m.Select("b")
	m.Inspect()
	m.Detail.ScrollDown(3)
	before, offset := m.View(), m.Detail.YOffset()

	for _, invalid := range [][]pattern.InspectorRecord{
		{inspectorRecord("a", "new"), inspectorRecord("a", "duplicate")},
		{inspectorRecord("", "empty identity")},
		{inspectorRecord(" \t", "blank identity")},
		{inspectorRecord("unsafe\nidentity", "control")},
		{{ID: "missing row"}},
	} {
		if err := m.SetRecords(invalid); err == nil {
			t.Fatal("invalid replacement was accepted")
		}

		if id, present := m.SelectedID(); !present || id != "b" || !m.Inspecting() || m.View() != before {
			t.Fatal("rejected replacement mutated selection, disclosure or content")
		}
	}

	records[1].ID = "caller changed slice"

	if id, _ := m.SelectedID(); id != "b" {
		t.Fatal("caller slice mutation changed retained identity")
	}

	reordered := []pattern.InspectorRecord{inspectorRecord("b", "updated row b"), inspectorRecord("a", "row a")}
	if err := m.SetRecords(reordered); err != nil {
		t.Fatal(err)
	}

	if id, _ := m.SelectedID(); id != "b" || !m.Inspecting() || m.Detail.YOffset() != offset {
		t.Fatal("reordering lost the inspected identity or detail position")
	}

	if m.Select("absent") || m.Detail.YOffset() != offset {
		t.Fatal("selecting an absent ID changed state")
	}

	if err := m.SetRecords([]pattern.InspectorRecord{inspectorRecord("a", "remaining row")}); err != nil {
		t.Fatal(err)
	}

	if id, _ := m.SelectedID(); id != "a" || m.Inspecting() || m.Detail.YOffset() != 0 {
		t.Fatal("removed selection silently became a different inspected record")
	}
}

func TestListInspectorEmptyAndUnavailableDetails(t *testing.T) {
	t.Parallel()

	m := newInspector(t, nil, layout.Size{Width: 80, Height: 12})
	m.Empty = "No captured incidents\nWaiting for source \x1b[31m"
	m.Refresh()
	m.Move(1)
	m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})

	if _, present := m.SelectedID(); present || m.Inspect() || !strings.Contains(ansi.Strip(m.View()), "No captured incidents") {
		t.Fatal("empty data invented selection or hid its supplied explanation")
	}

	lines := strings.Split(ansi.Strip(m.List.GetContent()), "\n")
	if len(lines) != 2 || strings.TrimSpace(lines[0]) != "No captured incidents" || strings.TrimSpace(lines[1]) != "Waiting for source \\u001b[31m" {
		t.Fatal("empty explanation lost its paragraphs or allowed terminal controls")
	}

	record := inspectorRecord("one", "summary only")
	record.Detail = nil

	if err := m.SetRecords([]pattern.InspectorRecord{record}); err != nil {
		t.Fatal(err)
	}

	if m.Inspect() || m.Inspecting() {
		t.Fatal("record without details entered inspection")
	}

	record.Detail = func(int) string { return "now captured" }
	if err := m.SetRecords([]pattern.InspectorRecord{record}); err != nil {
		t.Fatal(err)
	}

	if !m.Inspect() {
		t.Fatal("newly supplied details were unavailable")
	}

	record.Detail = nil
	if err := m.SetRecords([]pattern.InspectorRecord{record}); err != nil {
		t.Fatal(err)
	}

	if m.Inspecting() {
		t.Fatal("removed details retained a stale inspection")
	}

	if err := m.SetRecords(nil); err != nil {
		t.Fatal(err)
	}

	if _, present := m.SelectedID(); present || m.Inspecting() || !strings.Contains(ansi.Strip(m.View()), "No captured incidents") {
		t.Fatal("clearing records retained a stale selection or inspector")
	}
}

func TestListInspectorPreservesIndependentNativeScrollPositions(t *testing.T) {
	t.Parallel()

	records := []pattern.InspectorRecord{}
	for _, id := range []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"} {
		records = append(records, inspectorRecord(id, "row "+id))
	}

	m := newInspector(t, records, layout.Size{Width: 80, Height: 8})
	m.Update(tea.KeyPressMsg{Code: ']'})
	listOffset := m.List.YOffset()

	if listOffset == 0 {
		t.Fatal("native list page binding did not scroll")
	}

	if id, _ := m.SelectedID(); id != "a" {
		t.Fatal("native scrolling changed record selection")
	}

	m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})
	m.Update(tea.KeyPressMsg{Code: ']'})
	detailOffset := m.Detail.YOffset()
	m.Update(tea.KeyPressMsg{Code: tea.KeyEscape})

	if detailOffset == 0 || m.Inspecting() || m.List.YOffset() != listOffset {
		t.Fatal("inspection/back lost the retained list position or detail scrolling")
	}

	m.Inspect()
	m.Update(tea.KeyPressMsg{Code: tea.KeyDown})

	if id, _ := m.SelectedID(); id != "a" || m.Detail.YOffset() != detailOffset+1 {
		t.Fatal("inspection navigation changed records instead of using the native detail viewport")
	}

	m.Back()
	m.Move(1)

	if id, _ := m.SelectedID(); id != "b" || m.Detail.YOffset() != 0 {
		t.Fatal("selection change did not reset the old record's detail offset")
	}

	m.Move(math.MaxInt)

	if id, _ := m.SelectedID(); id != "j" {
		t.Fatal("large forward movement did not clamp to the last supplied record")
	}

	m.Move(math.MinInt)

	if id, _ := m.SelectedID(); id != "a" {
		t.Fatal("large backward movement did not clamp to the first supplied record")
	}
}

func TestListInspectorMeasuresMultilineRowsAndKeepsTallRecordsReachable(t *testing.T) {
	t.Parallel()

	first := inspectorRecord("short", "first row")
	first.Heading = "Earlier observations"
	wide := inspectorRecord("unicode", "\x1b[31m"+strings.Repeat("界", 14)+"\x1b[0m\nunicode tail")
	selected := inspectorRecord("selected", "selected start\nmiddle one\nmiddle two\nselected end")
	tall := inspectorRecord("tall", "tall start\n"+strings.Repeat("tall middle\n", 12)+"tall end")
	m := newInspector(t, []pattern.InspectorRecord{first, wide, selected, tall}, layout.Size{Width: 24, Height: 8})
	m.Select("selected")

	view := ansi.Strip(m.View())
	if !strings.Contains(view, "selected start") || !strings.Contains(view, "selected end") || !strings.Contains(view, "→") {
		t.Fatal("variable row heights hid the selected record or selection marker")
	}

	m.Select("unicode")

	if !strings.Contains(ansi.Strip(m.View()), "unicode tail") {
		t.Fatal("ANSI or wide cells corrupted the row's visible span")
	}

	m.Select("tall")

	if !strings.Contains(ansi.Strip(m.View()), "tall start") {
		t.Fatal("oversized record did not reveal its beginning")
	}

	for range 3 {
		m.Update(tea.KeyPressMsg{Code: ']'})
	}

	if !strings.Contains(ansi.Strip(m.View()), "tall end") {
		t.Fatal("native paging cannot reach the rest of an oversized record")
	}

	offset := m.List.YOffset()
	m.Refresh()

	if m.List.YOffset() != offset {
		t.Fatal("refresh snapped an oversized record back to its heading")
	}

	m.Select("tall")

	if !strings.Contains(ansi.Strip(m.View()), "tall start") {
		t.Fatal("reselecting an oversized record cannot return to its start")
	}
}

func TestListInspectorAnchorsScrolledContentAcrossInsertionAndResize(t *testing.T) {
	t.Parallel()

	records := []pattern.InspectorRecord{}
	for _, id := range []string{"a", "b", "c", "d", "e", "f", "g", "h"} {
		records = append(records, inspectorRecord(id, "row "+id))
	}

	m := newInspector(t, records, layout.Size{Width: 80, Height: 8})
	m.List.SetYOffset(4)

	replacement := append([]pattern.InspectorRecord{inspectorRecord("new", "new observation")}, records...)

	if err := m.SetRecords(replacement); err != nil {
		t.Fatal(err)
	}

	m.Resize(layout.Size{Width: 60, Height: 8})
	top := strings.TrimSpace(ansi.Strip(strings.Split(m.List.View(), "\n")[0]))

	if top != "row c" {
		t.Fatalf("insertion or reflow displaced the visible record: %q", top)
	}

	if id, _ := m.SelectedID(); id != "a" {
		t.Fatal("anchoring changed the selected identity")
	}

	m.Move(1)

	if !strings.Contains(ansi.Strip(m.View()), "row b") {
		t.Fatal("record navigation did not reveal selection after manual scrolling")
	}
}

func TestListInspectorRevealsSelectionChangedWhileListWasHidden(t *testing.T) {
	t.Parallel()

	records := []pattern.InspectorRecord{}
	for _, id := range []string{"a", "b", "c", "d", "e", "f", "g", "h"} {
		records = append(records, inspectorRecord(id, "record "+id))
	}

	m := newInspector(t, records, layout.Size{Width: 80, Height: 8})
	m.Inspect()
	m.Select("h")
	m.Back()

	if id, _ := m.SelectedID(); id != "h" || !strings.Contains(ansi.Strip(m.View()), "record h") {
		t.Fatal("selection changed behind inspection but was not revealed on return")
	}
}

func TestListInspectorBudgetsNativeGuttersBeforeRendering(t *testing.T) {
	t.Parallel()

	record := pattern.InspectorRecord{
		ID: "gutter",
		Row: func(width int, _ bool) string {
			return strings.Repeat("r", max(0, width-3)) + "END"
		},
		Detail: func(width int) string {
			return strings.Repeat("d", max(0, width-3)) + "END"
		},
	}
	m := newInspector(t, []pattern.InspectorRecord{record}, layout.Size{Width: 100, Height: 12})
	gutter := func(viewport.GutterContext) string { return "|  " }
	m.List.LeftGutterFunc, m.Detail.LeftGutterFunc = gutter, gutter
	m.List.Style = m.List.Style.Padding(0, 1)
	m.Detail.Style = m.Detail.Style.Padding(0, 1)
	m.Refresh()

	if !strings.Contains(ansi.Strip(m.List.View()), "END") || !strings.Contains(ansi.Strip(m.Detail.View()), "END") {
		t.Fatal("native gutters or padding truncated the callback's supplied content width")
	}

	m.Inspect()

	if !strings.Contains(ansi.Strip(m.View()), "END") {
		t.Fatal("inspection did not retain the native gutter's width budget")
	}
}

func TestListInspectorNativeCustomizationAndFocusSurviveReplacement(t *testing.T) {
	t.Parallel()

	records := []pattern.InspectorRecord{inspectorRecord("a", "row a"), inspectorRecord("b", "row b")}
	m := newInspector(t, records, layout.Size{Width: 100, Height: 12})
	m.List.Style = m.List.Style.Bold(true).Padding(1, 2)
	m.Detail.Style = m.Detail.Style.Italic(true)
	m.Detail.KeyMap.Down = key.NewBinding(key.WithKeys("n"), key.WithHelp("n", "next line"))
	m.List.KeyMap.PageDown = key.NewBinding(key.WithKeys("v"), key.WithHelp("v", "page down"))
	m.Keys.Next = key.NewBinding(key.WithKeys("x"), key.WithHelp("x", "next record"))
	m.List.SoftWrap, m.Detail.SoftWrap = true, true

	if err := m.SetRecords(records); err != nil {
		t.Fatal(err)
	}

	m.Resize(layout.Size{Width: 132, Height: 18})

	if !m.List.Style.GetBold() || !m.Detail.Style.GetItalic() || m.List.KeyMap.PageDown.Help().Key != "v" || m.Detail.KeyMap.Down.Help().Key != "n" {
		t.Fatal("replacement or resize discarded native customization")
	}

	if m.List.SoftWrap || m.Detail.SoftWrap {
		t.Fatal("native wrapping overrode the wrapper's measured row projection")
	}

	m.Blur()
	m.Update(tea.KeyPressMsg{Code: 'x'})
	m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})

	if id, _ := m.SelectedID(); id != "a" || m.Inspecting() {
		t.Fatal("blurred inspector handled keyboard input")
	}

	m.Focus()
	m.Update(tea.KeyPressMsg{Code: 'x'})
	m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})
	m.Update(tea.KeyPressMsg{Code: 'n'})

	if id, _ := m.SelectedID(); id != "b" || !m.Inspecting() || m.Detail.YOffset() != 1 {
		t.Fatal("custom record or native detail binding was lost")
	}
}

func TestListInspectorViewAndNativeScrollingDoNotRunRenderers(t *testing.T) {
	t.Parallel()

	rows, details := 0, 0
	record := pattern.InspectorRecord{
		ID: "measured",
		Row: func(int, bool) string {
			rows++
			return strings.Repeat("row observation\n", 30)
		},
		Detail: func(int) string {
			details++
			return strings.Repeat("detail observation\n", 30)
		},
	}
	m := newInspector(t, []pattern.InspectorRecord{record}, layout.Size{Width: 100, Height: 12})
	preparedRows, preparedDetails := rows, details

	m.View()
	m.Update(tea.KeyPressMsg{Code: ']'})
	m.View()

	if rows == 0 || details == 0 || rows != preparedRows || details != preparedDetails {
		t.Fatal("View or native scrolling prepared content")
	}

	m.Refresh()

	if rows <= preparedRows || details <= preparedDetails {
		t.Fatal("explicit refresh did not prepare updated caller-owned presentation")
	}

	m.Inspect()

	preparedRows, preparedDetails = rows, details

	m.View()
	m.Update(tea.KeyPressMsg{Code: tea.KeyDown})
	m.View()

	if rows != preparedRows || details != preparedDetails {
		t.Fatal("inspection View or native scrolling prepared content")
	}
}

func TestListInspectorResponsiveOptionsAndTerminalBounds(t *testing.T) {
	t.Parallel()

	records := []pattern.InspectorRecord{inspectorRecord("one", "Selected record\nSupporting evidence")}
	m := newInspector(t, records, layout.Size{Width: 86, Height: 12})

	if !m.Split() || !strings.Contains(ansi.Strip(m.View()), "Details for one") {
		t.Fatal("wide preview did not show the selected record's evidence")
	}

	m.Preview = false
	m.Refresh()

	if m.Split() || strings.Contains(ansi.Strip(m.View()), "Details for one") {
		t.Fatal("inventory option retained a preview")
	}

	m.Preview, m.ListWidth, m.MinDetailWidth = true, 48, 52
	m.Resize(layout.Size{Width: 100, Height: 12})

	if m.Split() {
		t.Fatal("preview ignored declared pane widths")
	}

	m.Resize(layout.Size{Width: 102, Height: 12})

	if !m.Split() {
		t.Fatal("preview did not use available pane allocation")
	}

	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Mono} {
		for _, ascii := range []bool{false, true} {
			for _, size := range []layout.Size{{}, {Width: 1, Height: 1}, {Width: 2, Height: 6}, {Width: 24, Height: 2}, {Width: 40, Height: 18}, {Width: 80, Height: 24}, {Width: 132, Height: 40}} {
				palette := theme.New(mode)
				palette.ASCII = ascii

				component, err := pattern.NewListInspector(records, size, palette)
				if err != nil {
					t.Fatal(err)
				}

				for _, inspect := range []bool{false, true} {
					if inspect {
						component.Inspect()
					}

					view := component.View()
					if size.Width == 0 || size.Height == 0 {
						if view != "" {
							t.Fatal("zero allocation rendered content")
						}

						continue
					}

					lines := strings.Split(view, "\n")
					if len(lines) != size.Height {
						t.Fatalf("%v: rendered height %d", size, len(lines))
					}

					for _, line := range lines {
						if ansi.StringWidth(line) != size.Width {
							t.Fatalf("%v: rendered line escaped allocation: %q", size, line)
						}
					}
				}
			}
		}
	}
}
