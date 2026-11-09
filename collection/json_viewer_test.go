package collection_test

import (
	"strings"
	"testing"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/theme"
)

func TestJSONViewerPreservesSourceEvidence(t *testing.T) {
	source := []byte(`{"z":9007199254740993,"a":1.2300e+04,"z":"\u0061\u001b","empty":{},"list":[],"nil":null}`)
	m := collection.NewJSONViewer(100, 20, theme.New(theme.Mono))

	m.Folding = false
	if err := m.SetSource(source); err != nil {
		t.Fatal(err)
	}

	want := "{\n  \"z\": 9007199254740993,\n  \"a\": 1.2300e+04,\n  \"z\": \"\\u0061\\u001b\",\n  \"empty\": {},\n  \"list\": [],\n  \"nil\": null\n}"
	if got := jsonViewerText(m.View()); got != want {
		t.Fatalf("changed JSON evidence:\n%s", got)
	}

	source[2] = 'x'
	m.Raw = true
	m.Refresh()

	if got := strings.TrimSpace(ansi.Strip(m.View())); !strings.HasPrefix(got, `{"z":9007199254740993,"a":1.2300e+04,"z":`) {
		t.Fatalf("raw view changed lexemes or retained caller's mutable bytes: %s", got)
	}
}

func jsonViewerText(view string) string {
	lines := strings.Split(ansi.Strip(view), "\n")
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, " ")
	}

	return strings.TrimSpace(strings.Join(lines, "\n"))
}

func TestJSONViewerIndependentDuplicateMembersAndRoundTrips(t *testing.T) {
	source := []byte(`{"same":{"first":111},"same":{"second":222},"rates":[{"rate":333},444]}`)

	m := collection.NewJSONViewer(80, 30, theme.New(theme.Mono))
	if err := m.SetSource(source); err != nil {
		t.Fatal(err)
	}

	m.Model.KeyMap.Down = key.NewBinding(key.WithKeys("x"), key.WithHelp("x", "next"))
	m.Update(tea.KeyPressMsg{Code: 'x'})
	m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})

	view := ansi.Strip(m.View())
	if !m.Collapsed(1) || strings.Contains(view, "111") || !strings.Contains(view, "222") || !strings.Contains(view, "1 field}") {
		t.Fatalf("duplicate keys shared fold state: %s", view)
	}

	snapshot := m.Snapshot()
	m.Raw = true
	m.Refresh()

	if m.Interactive() || !strings.Contains(ansi.Strip(m.View()), "111") {
		t.Fatal("raw view lost source or retained structural routing")
	}

	m.Raw = false
	m.Resize(24, 8)
	m.Resize(80, 30)

	if !m.Collapsed(1) || m.Selected() != 1 {
		t.Fatal("raw/resize round trip lost folds or selection")
	}

	m.Update(tea.KeyPressMsg{Code: 'x'})

	if m.Selected() != 3 {
		t.Fatal("native navigation override did not skip folded descendants")
	}

	if err := m.SetSource([]byte(`{"replacement":true}`)); err != nil {
		t.Fatal(err)
	}

	if m.RestoreState(snapshot) || m.Selected() != 0 || m.Collapsed(1) {
		t.Fatal("new body inherited stale structure")
	}

	if err := m.SetSource(source); err != nil {
		t.Fatal(err)
	}

	if !m.RestoreState(snapshot) || !m.Collapsed(1) || m.Selected() != 1 {
		t.Fatal("same source could not restore its snapshot")
	}
}

func TestJSONViewerCursorRemainsVisibleAndFoldsCanBeDisabled(t *testing.T) {
	m := collection.NewJSONViewer(30, 5, theme.New(theme.Mono))

	m.Prefix = "Header\nAnother header"
	if err := m.SetSource([]byte(`{"outer":{"child":{"message":"a long wrapped string that must remain visible"}},"tail":true}`)); err != nil {
		t.Fatal(err)
	}

	m.Select(3)
	m.Update(tea.KeyPressMsg{Code: tea.KeyLeft})

	if m.Selected() != 2 {
		t.Fatal("left did not navigate from scalar to parent")
	}

	m.Update(tea.KeyPressMsg{Code: tea.KeyLeft})

	if !m.Collapsed(2) || !strings.Contains(ansi.Strip(m.View()), "child") {
		t.Fatal("collapse left cursor outside viewport")
	}

	m.SetCollapsed(1, true)

	if m.Selected() != 1 {
		t.Fatal("collapsing ancestor left selection in a hidden child")
	}

	m.Folding = false
	m.Resize(80, 20)
	m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})

	if m.Interactive() || !strings.Contains(ansi.Strip(m.View()), "message") {
		t.Fatal("disabled folds left hidden content or accepted fold input")
	}

	m.Folding = true
	m.Refresh()

	if m.Collapsed(1) || m.Collapsed(2) {
		t.Fatal("re-enabling folds restored stale collapsed state")
	}
}

func TestJSONViewerRevealsSelectedFieldBelowHeaders(t *testing.T) {
	for _, prefix := range []string{"", "Header", "Header\nAnother header", "Header\n"} {
		m := collection.NewJSONViewer(40, 1, theme.New(theme.Mono))
		m.Prefix = prefix

		if err := m.SetSource([]byte(`{"first":1,"second":2}`)); err != nil {
			t.Fatal(err)
		}

		m.Select(1)

		if !strings.Contains(ansi.Strip(m.View()), `"first": 1`) {
			t.Fatalf("selected field hidden below prefix %q: %s", prefix, ansi.Strip(m.View()))
		}
	}
}

func TestJSONViewerRejectsInvalidAndOverBudgetStructure(t *testing.T) {
	for _, test := range []struct {
		name, source, notice string
		limits               collection.JSONLimits
	}{
		{name: "invalid", source: `{"missing":`, notice: "invalid JSON"},
		{name: "trailing", source: `{} {}`, notice: "invalid JSON"},
		{name: "bytes", source: `{"secret":1234567890}`, notice: "byte limit", limits: collection.JSONLimits{Bytes: 8}},
		{name: "depth", source: `{"a":{"b":{"c":1}}}`, notice: "depth limit", limits: collection.JSONLimits{Depth: 1}},
		{name: "nodes", source: `[1,2,3,4]`, notice: "node limit", limits: collection.JSONLimits{Nodes: 2}},
		{name: "output", source: `{"a":{"b":{"c":{"d":12345}}}}`, notice: "rendered-output limit", limits: collection.JSONLimits{RenderBytes: 60}},
	} {
		t.Run(test.name, func(t *testing.T) {
			m := collection.NewJSONViewer(120, 30, theme.New(theme.Mono))
			m.Limits = test.limits

			err := m.SetSource([]byte(test.source))
			if err == nil || !strings.Contains(err.Error(), test.notice) || m.Interactive() || !strings.Contains(ansi.Strip(m.View()), test.notice) {
				t.Fatalf("missing safe fallback: %v\n%s", err, ansi.Strip(m.View()))
			}
		})
	}
}

func TestJSONViewerScalarRootsAndEmptyContainers(t *testing.T) {
	for _, source := range []string{`true`, `null`, `-0.00E-9`, `"\u0061"`, `{}`, `[]`} {
		m := collection.NewJSONViewer(40, 4, theme.New(theme.Mono))

		m.Folding = false
		if err := m.SetSource([]byte(source)); err != nil {
			t.Fatal(err)
		}

		if got := strings.TrimSpace(ansi.Strip(m.View())); got != source {
			t.Fatalf("%s rendered as %s", source, got)
		}
	}
}

func TestJSONViewerNativePagingReadsLongValuesAndKeepsRawOffset(t *testing.T) {
	m := collection.NewJSONViewer(24, 4, theme.New(theme.Mono))
	if err := m.SetSource([]byte(`{"message":"` + strings.Repeat("payload ", 80) + `","tail":true}`)); err != nil {
		t.Fatal(err)
	}

	m.Select(1)
	before := m.Model.YOffset()
	m.Update(tea.KeyPressMsg{Code: tea.KeyPgDown})

	if m.Model.YOffset() <= before || m.Selected() != 1 || !strings.Contains(ansi.Strip(m.View()), "payload") {
		t.Fatal("paging skipped the middle of the wrapped scalar")
	}

	formatted := m.Model.YOffset()
	m.Raw = true
	m.Refresh()
	m.Model.SetYOffset(2)
	m.Raw = false
	m.Refresh()

	if m.Model.YOffset() != formatted {
		t.Fatal("raw round trip lost formatted scroll offset")
	}

	m.Raw = true
	m.Refresh()

	if m.Model.YOffset() != 2 {
		t.Fatal("formatted round trip lost raw scroll offset")
	}
}

func TestJSONViewerEscapesUnsafeInputAndBoundsDeepDocuments(t *testing.T) {
	m := collection.NewJSONViewer(100, 8, theme.New(theme.Mono))
	if err := m.SetSource([]byte("{\"unsafe\":\"\x1b]0;title\a\"}")); err == nil {
		t.Fatal("accepted literal terminal controls as JSON")
	}

	view := ansi.Strip(m.View())
	if strings.Contains(view, "\x1b") || !strings.Contains(view, `\u001b`) {
		t.Fatal("fallback did not escape terminal controls")
	}

	m.Limits.Depth = 10000
	if err := m.SetSource([]byte(strings.Repeat("[", 300) + "0" + strings.Repeat("]", 300))); err == nil || !strings.Contains(err.Error(), "depth limit (256)") {
		t.Fatalf("caller limits bypassed the traversal bound: %v", err)
	}
}
