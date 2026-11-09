package collection_test

import (
	"bytes"
	"image/color"
	"net/http"
	"strings"
	"testing"

	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestBodyCaptureStatesContentTypesAndRawPreservation(t *testing.T) {
	palette := theme.New(theme.Dark)
	original := []byte(`{"n":1,"ok":true}`)
	body := collection.HTTPBody{State: collection.Captured, ContentType: "application/problem+json; charset=utf-8", Bytes: bytes.Clone(original)}
	pretty := ansi.Strip(body.View(false, 4096, palette))

	raw := ansi.Strip(body.View(true, 4096, palette))
	if !strings.Contains(pretty, "\n  \"n\": 1") || !strings.Contains(raw, string(original)) || !bytes.Equal(body.Bytes, original) {
		t.Fatal("pretty/raw lost captured bytes", pretty, raw)
	}

	for _, example := range []struct {
		body collection.HTTPBody
		want string
	}{
		{collection.HTTPBody{}, "not captured"},
		{collection.HTTPBody{State: collection.Captured}, "Empty body"},
		{collection.HTTPBody{State: collection.Redacted, Bytes: original}, "redacted"},
		{collection.HTTPBody{State: collection.Truncated}, "capture truncated"},
		{collection.HTTPBody{State: collection.Captured, ContentType: "application/json", Bytes: []byte("{bad")}, "invalid JSON"},
		{collection.HTTPBody{State: collection.Captured, ContentType: "image/png", Bytes: []byte{0, 255}}, "binary"},
		{collection.HTTPBody{State: collection.Captured, Encoding: "gzip", Bytes: original}, "encoding gzip"},
		{collection.HTTPBody{State: collection.Captured, ContentType: "text/plain; charset=latin1", Bytes: original}, "unsupported charset"},
	} {
		if !strings.Contains(ansi.Strip(example.body.View(false, 4096, palette)), example.want) {
			t.Fatal(example.want)
		}
	}

	if !strings.Contains(ansi.Strip(body.View(false, 4, palette)), "preview truncated") {
		t.Fatal("prefix silently omitted")
	}

	binary := collection.HTTPBody{State: collection.Truncated, ContentType: "image/png", Bytes: []byte{0, 255}}
	if got := ansi.Strip(binary.View(true, 4096, palette)); !strings.Contains(got, "00 ff") || !strings.Contains(got, "capture truncated") {
		t.Fatal(got)
	}
}

func TestHTTPUpdatesPreserveSelectionNativeCustomizationAndOwnedBytes(t *testing.T) {
	m := collection.NewHTTPLog(60, 20, theme.New(theme.Dark))
	m.Limit = 2

	records := []collection.HTTPExchange{{ID: "a", Method: "GET", URL: "/a", RequestHeaders: http.Header{"X-Test": {"first", "second"}}, ResponseBody: collection.HTTPBody{State: collection.Captured, ContentType: "application/json", Bytes: []byte(`{"x":1}`)}}, {ID: "b", Method: "POST", URL: "/b"}}
	if err := m.SetEntries(records); err != nil {
		t.Fatal(err)
	}

	m.Follow = false
	m.Select("a")
	m.Table.KeyMap.LineUp.SetKeys("x")

	records[0].ResponseBody.Bytes[0] = '!'
	records[0].RequestHeaders["X-Test"][0] = "changed"

	selected, _ := m.Selected()
	if selected.ResponseBody.Bytes[0] != '{' || selected.RequestHeaders["X-Test"][0] != "first" {
		t.Fatal("capture storage aliases caller buffers")
	}

	if err := m.Upsert(collection.HTTPExchange{ID: "b", Method: "PUT", URL: "/b2"}); err != nil {
		t.Fatal(err)
	}

	m.Resize(40, 16)

	if m.SelectedID() != "a" || m.Table.KeyMap.LineUp.Keys()[0] != "x" {
		t.Fatal("update changed selection or native keymap")
	}

	if err := m.Upsert(collection.HTTPExchange{ID: "c"}); err != nil {
		t.Fatal(err)
	}

	if m.SelectedID() != "b" || len(m.Entries()) != 2 {
		t.Fatal("evicted selection not repaired deterministically")
	}

	before := m.Entries()
	if err := m.SetEntries([]collection.HTTPExchange{{ID: "bad", ResponseBody: collection.HTTPBody{State: "invalid"}}}); err == nil {
		t.Fatal("invalid capture accepted")
	}

	if len(m.Entries()) != len(before) || m.SelectedID() != "b" {
		t.Fatal("validation partially mutated model")
	}

	m.ResumeFollow()

	if m.SelectedID() != "c" {
		t.Fatal("follow did not resume")
	}

	m.UpdateRequests(tea.KeyPressMsg{Code: 'x'})

	if m.Follow {
		t.Fatal("browsing did not pause follow")
	}
}

func TestLogTreeHierarchyFieldsAndStableExpansion(t *testing.T) {
	m := collection.NewLogTree(100, 30, theme.New(theme.Dark))

	records := []collection.LogRecord{
		{ID: "a", TraceID: "trace", SpanID: "child", ParentSpanID: "parent", LogEntry: collection.LogEntry{Message: "child entry", Fields: []any{"nested", map[string]any{"b": 2}}}},
		{ID: "b", TraceID: "trace", SpanID: "parent", LogEntry: collection.LogEntry{Message: "parent entry", Fields: []any{"value", true}}},
	}
	if err := m.SetEntries(records); err != nil {
		t.Fatal(err)
	}

	m.Follow = false
	if !m.Select("b") {
		t.Fatal("entry not selectable")
	}

	m.Model.ToggleCurrentNode()

	if !strings.Contains(ansi.Strip(m.View()), "value") {
		t.Fatal("entry fields did not expand")
	}

	if err := m.Upsert(collection.LogRecord{ID: "c", LogEntry: collection.LogEntry{Message: "unrelated"}}); err != nil {
		t.Fatal(err)
	}

	if m.Selected() != "b" || !strings.Contains(ansi.Strip(m.View()), "value") {
		t.Fatalf("incoming entry lost stable selection/expansion selected=%s view=%s", m.Selected(), ansi.Strip(m.View()))
	}

	m.SetGrouped(false)

	if strings.Contains(ansi.Strip(m.View()), "Trace trace") {
		t.Fatal("flat view retained trace scaffolding")
	}

	for _, parents := range [][2]string{{"missing", ""}, {"second", "first"}} {
		r := []collection.LogRecord{{ID: "x", TraceID: "t", SpanID: "first", ParentSpanID: parents[0], LogEntry: collection.LogEntry{Message: "first"}}, {ID: "y", TraceID: "t", SpanID: "second", ParentSpanID: parents[1], LogEntry: collection.LogEntry{Message: "second"}}}

		m.SetGrouped(true)

		if err := m.SetEntries(r); err != nil {
			t.Fatal(err)
		}

		if !strings.Contains(ansi.Strip(m.View()), "unresolved parent") {
			t.Fatalf("broken hierarchy hidden: %s", ansi.Strip(m.View()))
		}
	}
}

func TestLogFieldCyclesAreBoundedAndUnknownIDsFailAtomically(t *testing.T) {
	cyclic := map[string]any{}
	cyclic["self"] = cyclic
	m := collection.NewLogTree(80, 20, theme.New(theme.Dark))

	m.FieldDepth = 2
	if err := m.SetEntries([]collection.LogRecord{{ID: "one", LogEntry: collection.LogEntry{Fields: []any{"cycle", cyclic, "odd"}}}}); err != nil {
		t.Fatal(err)
	}

	if err := m.Upsert(collection.LogRecord{ID: ""}); err == nil {
		t.Fatal("missing stable identity accepted")
	}

	if len(m.Entries()) != 1 {
		t.Fatal("invalid update changed history")
	}
}

func TestLogFieldSelectionSurvivesTopLevelInsertion(t *testing.T) {
	m := collection.NewLogTree(80, 30, theme.New(theme.Dark))

	record := collection.LogRecord{ID: "entry", LogEntry: collection.LogEntry{Fields: []any{"details", map[string]any{"inner": true}, "tag", "a"}}}
	if err := m.SetEntries([]collection.LogRecord{record}); err != nil {
		t.Fatal(err)
	}

	m.Follow = false
	m.Select("entry")
	m.Model.ToggleCurrentNode()

	for _, node := range m.Model.AllNodes() {
		if strings.Contains(ansi.Strip(node.Value()), "details") {
			m.Model.SetYOffset(node.YOffset())
			m.Model.ToggleCurrentNode()

			break
		}
	}

	record.Fields = []any{"inserted", 0, "tag", "a", "details", map[string]any{"inner": true}}
	if err := m.Upsert(record); err != nil {
		t.Fatal(err)
	}

	found := false

	for _, node := range m.Model.AllNodes() {
		if node.IsSelected() {
			found = strings.Contains(ansi.Strip(node.Value()), "details")
		}
	}

	if !found || !strings.Contains(ansi.Strip(m.View()), "inner") {
		t.Fatal("field insertion transferred UI state to another field")
	}
}

func TestJSONFormattingBudgetIsBoundedBeforeIndentation(t *testing.T) {
	source := []byte(strings.Repeat("[", 200) + "0" + strings.Repeat("]", 200))
	body := collection.HTTPBody{State: collection.Captured, ContentType: "application/json", Bytes: source}

	view := ansi.Strip(body.View(false, 1024, theme.New(theme.Dark)))
	if len(view) > 2048 || !strings.Contains(view, "formatting depth limit") || !bytes.Equal(body.Bytes, source) {
		t.Fatal("deep JSON escaped render budget")
	}
}

func TestCustomNativeNavigationPausesFollow(t *testing.T) {
	httpLog := collection.NewHTTPLog(40, 16, theme.New(theme.Dark))
	if err := httpLog.SetEntries([]collection.HTTPExchange{{ID: "a"}, {ID: "b"}}); err != nil {
		t.Fatal(err)
	}

	httpLog.Table.KeyMap.LineUp.SetKeys("x")
	httpLog.UpdateRequests(tea.KeyPressMsg{Code: 'x'})

	if httpLog.Follow || httpLog.SelectedID() != "a" {
		t.Fatal("custom HTTP navigation failed to pause follow")
	}

	if err := httpLog.Upsert(collection.HTTPExchange{ID: "c"}); err != nil {
		t.Fatal(err)
	}

	if httpLog.SelectedID() != "a" {
		t.Fatal("append stole custom-key selection")
	}

	logs := collection.NewLogTree(40, 16, theme.New(theme.Dark))
	if err := logs.SetEntries([]collection.LogRecord{{ID: "a"}, {ID: "b"}}); err != nil {
		t.Fatal(err)
	}

	logs.Model.KeyMap.Up.SetKeys("x")
	logs.Update(tea.KeyPressMsg{Code: 'x'})

	if logs.Follow || logs.Selected() != "a" {
		t.Fatal("custom log navigation failed to pause follow")
	}

	if err := logs.Upsert(collection.LogRecord{ID: "c"}); err != nil {
		t.Fatal(err)
	}

	if logs.Selected() != "a" {
		t.Fatal("append stole custom log selection")
	}
}

func TestRawFieldNamesDoNotCollideAfterEscaping(t *testing.T) {
	m := collection.NewLogTree(100, 30, theme.New(theme.Dark))

	record := collection.LogRecord{ID: "entry", LogEntry: collection.LogEntry{Fields: []any{"a\nb", map[string]any{"only-first": true}, `a\nb`, map[string]any{"only-second": true}}}}
	if err := m.SetEntries([]collection.LogRecord{record}); err != nil {
		t.Fatal(err)
	}

	m.Follow = false
	m.Select("entry")
	m.Model.ToggleCurrentNode()

	for _, node := range m.Model.AllNodes() {
		if strings.Contains(ansi.Strip(node.Value()), `a\nb`) {
			m.Model.SetYOffset(node.YOffset())
			m.Model.ToggleCurrentNode()

			break
		}
	}

	record.Fields = []any{`a\nb`, map[string]any{"only-second": true}, "a\nb", map[string]any{"only-first": true}}
	if err := m.Upsert(record); err != nil {
		t.Fatal(err)
	}

	view := ansi.Strip(m.View())
	if !strings.Contains(view, "only-first") || strings.Contains(view, "only-second") {
		t.Fatal("escaped field names transferred expansion")
	}
}

func TestNarrowTreeWrapsMessagesInsteadOfHidingThem(t *testing.T) {
	m := collection.NewLogTree(36, 30, theme.New(theme.Dark))

	records := []collection.LogRecord{{ID: "entry", TraceID: "trace", SpanID: "child", ParentSpanID: "parent", LogEntry: collection.LogEntry{Message: "Carrier response delayed"}}, {ID: "parent", TraceID: "trace", SpanID: "parent"}}
	if err := m.SetEntries(records); err != nil {
		t.Fatal(err)
	}

	m.Follow = false
	m.Select("entry")

	view := strings.ReplaceAll(ansi.Strip(m.View()), " ", "")
	if !strings.Contains(view, "delayed") {
		t.Fatal("narrow summary hid the message tail")
	}

	m.Resize(100, 30)

	if m.Selected() != "entry" {
		t.Fatal("reflow lost selected entry")
	}
}

func TestCollapsedAncestorsDoNotBreakSelectionOrNarrowWrapping(t *testing.T) {
	m := collection.NewLogTree(36, 30, theme.New(theme.Dark))

	records := []collection.LogRecord{{ID: "entry", TraceID: "trace", SpanID: "child", ParentSpanID: "parent", LogEntry: collection.LogEntry{Message: "Carrier response delayed"}}, {ID: "parent", TraceID: "trace", SpanID: "parent"}}
	if err := m.SetEntries(records); err != nil {
		t.Fatal(err)
	}

	m.Follow = false
	m.Select("entry")

	for _, node := range m.Model.AllNodes() {
		if strings.Contains(ansi.Strip(node.Value()), "Trace trace") {
			m.Model.SetYOffset(node.YOffset())
			m.Model.ToggleCurrentNode()

			break
		}
	}

	if err := m.Upsert(collection.LogRecord{ID: "unrelated"}); err != nil {
		t.Fatal(err)
	}

	if !m.Select("entry") || m.Selected() != "entry" {
		t.Fatal("programmatic selection failed to reveal collapsed ancestors")
	}

	if !strings.Contains(ansi.Strip(m.View()), "delayed") {
		t.Fatal("hidden structure lost narrow message width")
	}

	for _, node := range m.Model.AllNodes() {
		if strings.Contains(ansi.Strip(node.Value()), "Trace trace") {
			m.Model.SetYOffset(node.YOffset())
			m.Model.ToggleCurrentNode()

			break
		}
	}

	m.ResumeFollow()

	if err := m.Upsert(collection.LogRecord{ID: "latest", TraceID: "trace", SpanID: "child", ParentSpanID: "parent", LogEntry: collection.LogEntry{Message: "latest observation delayed"}}); err != nil {
		t.Fatal(err)
	}

	if m.Selected() != "latest" {
		t.Fatal("following did not reveal a new record under collapsed ancestors")
	}
}

func TestHTTPColumnsAlignAcrossMethodsAndStatuses(t *testing.T) {
	for _, width := range []int{40, 120} {
		m := collection.NewHTTPLog(width, 20, theme.New(theme.Dark))

		err := m.SetEntries([]collection.HTTPExchange{
			{ID: "a", Method: "GET", URL: "/alpha", StatusCode: 200, Completed: true},
			{ID: "b", Method: "PATCH", URL: "/bravo", StatusCode: 502, Completed: true},
			{ID: "c", Method: "CONNECT", URL: "/charlie"},
		})
		if err != nil {
			t.Fatal(err)
		}

		positions := []int{}

		for _, line := range strings.Split(ansi.Strip(m.Table.View()), "\n") {
			if index := strings.Index(line, "/"); index >= 0 {
				positions = append(positions, ansi.StringWidth(line[:index]))
			}
		}

		if len(positions) != 3 {
			t.Fatalf("width %d: missing request targets: %q", width, m.Table.View())
		}

		if positions[0] != positions[1] || positions[1] != positions[2] {
			t.Fatalf("width %d: target columns drift: %v", width, positions)
		}
	}
}

func TestHTTPCompactListPreservesExactStatus(t *testing.T) {
	m := collection.NewHTTPLog(16, 12, theme.New(theme.Dark))
	if err := m.SetEntries([]collection.HTTPExchange{{ID: "a", Method: "PATCH", URL: "/test", StatusCode: 502, Completed: true}}); err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(ansi.Strip(m.Table.View()), "502") {
		t.Fatal("compact list clipped status", m.Table.View())
	}
}

func TestLogRolesRetainOrdinaryCustomForeground(t *testing.T) {
	palette := theme.Theme{Foreground: color.RGBA{R: 220, G: 80, B: 60, A: 255}}

	httpLog := collection.NewHTTPLog(80, 30, palette)
	if err := httpLog.SetEntries([]collection.HTTPExchange{{ID: "request", RequestHeaders: http.Header{"X-Custom": {"yes"}}}}); err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(httpLog.DetailView(), palette.Text().Render("X-Custom")) {
		t.Fatal("header lost ordinary foreground", httpLog.DetailView())
	}

	logs := collection.NewLogTree(80, 20, palette)
	if err := logs.SetEntries([]collection.LogRecord{{ID: "entry", LogEntry: collection.LogEntry{Message: "message", Fields: []any{"field", "value"}}}}); err != nil {
		t.Fatal(err)
	}

	found := false

	for _, node := range logs.Model.AllNodes() {
		node.Open()
	}

	logs.Model.SetNodes(logs.Model.Root())

	for _, node := range logs.Model.AllNodes() {
		if strings.Contains(ansi.Strip(node.Value()), "field:") {
			found = true

			if !strings.Contains(node.Value(), palette.Text().Render("field")) || !strings.Contains(node.Value(), palette.Text().Render(`"value"`)) {
				t.Fatal("field lost ordinary foreground", node.Value())
			}
		}
	}

	if !found {
		t.Fatal("field regression did not inspect field")
	}
}

func TestOptionalFoldsKeepContentOpenAcrossUpdates(t *testing.T) {
	palette := theme.New(theme.Mono)
	record := collection.LogRecord{ID: "entry", LogEntry: collection.LogEntry{Message: "message", Fields: []any{"nested", map[string]any{"inner": "deep value"}}}}

	logs := collection.NewLogTree(80, 20, palette)
	if err := logs.SetEntries([]collection.LogRecord{record}); err != nil {
		t.Fatal(err)
	}

	logs.SetFolding(false)

	for _, key := range []tea.KeyPressMsg{{Code: tea.KeyEnter}, {Code: 'h'}, {Code: ' '}} {
		logs.Update(key)
	}

	if err := logs.Upsert(record); err != nil {
		t.Fatal(err)
	}

	logs.Resize(78, 20)

	if frame := ansi.Strip(logs.View()); !strings.Contains(frame, "deep value") || strings.Contains(frame, "v ") || strings.Contains(frame, "> ") {
		t.Fatal("foldless tree hid fields or showed fold controls", frame)
	}

	httpLog := collection.NewHTTPLog(80, 30, palette)
	if err := httpLog.SetEntries([]collection.HTTPExchange{{ID: "request"}}); err != nil {
		t.Fatal(err)
	}

	httpLog.SetFolding(false)
	httpLog.UpdateDetails(tea.KeyPressMsg{Code: '1'})
	httpLog.UpdateDetails(tea.KeyPressMsg{Code: '2'})

	if !httpLog.RequestOpen || !httpLog.ResponseOpen || strings.Contains(ansi.Strip(httpLog.DetailView()), "v Request") {
		t.Fatal("foldless HTTP hid a section or retained disclosure controls")
	}

	httpLog.SetFolding(true)
	httpLog.UpdateDetails(tea.KeyPressMsg{Code: '1'})

	if httpLog.RequestOpen {
		t.Fatal("enabling folds did not restore controls")
	}
}

func TestHTTPScrollbarUsesNativeViewportOffset(t *testing.T) {
	palette := theme.New(theme.Mono)

	httpLog := collection.NewHTTPLog(40, 20, palette)
	if err := httpLog.SetEntries([]collection.HTTPExchange{{ID: "request", ResponseBody: collection.HTTPBody{State: collection.Captured, ContentType: "text/plain", Bytes: []byte(strings.Repeat("line\n", 30))}}}); err != nil {
		t.Fatal(err)
	}

	httpLog.ResizeDetails(40, 8)
	httpLog.Details.GotoBottom()

	rows := strings.Split(ansi.Strip(httpLog.DetailView()), "\n")
	if len(rows) != 8 || !strings.HasSuffix(rows[7], "#") {
		t.Fatal("scroll thumb did not reach bottom", rows)
	}

	for _, row := range rows {
		if ansi.StringWidth(row) != 40 {
			t.Fatal("scroll gutter exceeded allocation", row)
		}
	}
}

func TestNoFoldsPreservesCollidingNativeNavigation(t *testing.T) {
	logs := collection.NewLogTree(80, 20, theme.New(theme.Mono))
	if err := logs.SetEntries([]collection.LogRecord{{ID: "a", LogEntry: collection.LogEntry{Message: "first"}}, {ID: "b", LogEntry: collection.LogEntry{Message: "second"}}}); err != nil {
		t.Fatal(err)
	}

	logs.SetFolding(false)
	logs.Follow = false
	logs.Select("a")
	logs.Model.KeyMap.Down.SetKeys("h")
	logs.Update(tea.KeyPressMsg{Code: 'h'})

	if logs.Selected() != "b" || !logs.Model.KeyMap.Close.Enabled() {
		t.Fatal("foldless mode swallowed navigation or mutated keymap", logs.Selected())
	}
}

func TestHTTPMinimalSummaryAndURLModesKeepCapturedSource(t *testing.T) {
	m := collection.NewHTTPLog(100, 30, theme.New(theme.Dark))

	captured := "https://api.example.test/v1/parcels%2F42?tag=a%2Bb"
	if err := m.SetEntries([]collection.HTTPExchange{{ID: "r", Method: "PATCH", URL: captured, Completed: true, StatusCode: 204, DurationKnown: true}}); err != nil {
		t.Fatal(err)
	}

	frame := ansi.Strip(m.DetailView())
	if !strings.Contains(frame, "PATCH /v1/parcels%2F42?tag=a%2Bb") || strings.Contains(frame, "Method") || strings.Contains(frame, "Target") || strings.Contains(frame, "Latency") {
		t.Fatal("HTTP summary retained redundant labels", frame)
	}

	m.SetURLFormat(text.FullURL)

	if !strings.Contains(ansi.Strip(m.DetailView()), captured) {
		t.Fatal("full URL not shown")
	}

	selected, _ := m.Selected()
	if selected.URL != captured {
		t.Fatal("formatting mutated capture")
	}

	m.SetURLFormatter(func(string) string { return "\x1b[31mcustom" })

	if !strings.Contains(ansi.Strip(m.DetailView()), `\u001b[31mcustom`) {
		t.Fatal("custom formatter escaped styling boundary")
	}
}

func TestHeaderlessListAccountsForCustomNativeHeaderHeight(t *testing.T) {
	m := collection.NewHTTPLog(40, 20, theme.New(theme.Mono))
	styles := table.DefaultStyles()
	styles.Header = styles.Header.PaddingBottom(2)
	m.Table.SetStyles(styles)

	if err := m.SetEntries([]collection.HTTPExchange{{ID: "a", Method: "GET", URL: "/first"}, {ID: "b", Method: "GET", URL: "/second"}, {ID: "c", Method: "GET", URL: "/third"}}); err != nil {
		t.Fatal(err)
	}

	m.ResizeList(40, 3)

	lines := strings.Split(ansi.Strip(m.ListView()), "\n")
	if len(lines) != 3 || !strings.Contains(lines[0], "/first") || !strings.Contains(lines[2], "/third") {
		t.Fatal("custom header leaked or consumed rows", lines)
	}
}

func TestPendingDetailsRespectPresentationChoice(t *testing.T) {
	palette := theme.New(theme.Dark)

	m := collection.NewHTTPLog(80, 30, palette)
	if err := m.SetEntries([]collection.HTTPExchange{{ID: "pending", Method: "GET", URL: "/health"}}); err != nil {
		t.Fatal(err)
	}

	for _, mode := range []text.Presentation{text.GlyphAndLabel, text.GlyphOnly, text.LabelOnly} {
		m.SetSymbols(mode)
		frame := ansi.Strip(m.DetailView())

		glyph, label := strings.Contains(frame, palette.Glyph(theme.GlyphPending)), strings.Contains(frame, "PENDING")
		if glyph != (mode != text.LabelOnly) || label != (mode != text.GlyphOnly) {
			t.Fatal("pending detail ignored symbol mode", mode, frame)
		}
	}
}

func TestLogTreePresentationNeverDropsCorrelationIdentities(t *testing.T) {
	palette := theme.New(theme.Dark)

	m := collection.NewLogTree(80, 20, palette)
	if err := m.SetEntries([]collection.LogRecord{{ID: "entry", TraceID: "trace-c7", SpanID: "http-label", LogEntry: collection.LogEntry{Message: "message"}}}); err != nil {
		t.Fatal(err)
	}

	for _, mode := range []text.Presentation{text.GlyphAndLabel, text.GlyphOnly, text.LabelOnly} {
		m.SetSymbols(mode)

		frame := ansi.Strip(m.View())
		if !strings.Contains(frame, "trace-c7") || !strings.Contains(frame, "http-label") {
			t.Fatal("presentation removed essential identity", mode, frame)
		}

		if mode == text.GlyphOnly && (strings.Contains(frame, "Trace ") || strings.Contains(frame, "Span ")) {
			t.Fatal("glyph-only retained descriptive labels", frame)
		}

		if mode == text.LabelOnly && strings.Contains(frame, palette.Glyph(theme.GlyphTrace)) {
			t.Fatal("label-only retained semantic glyph", frame)
		}
	}
}
