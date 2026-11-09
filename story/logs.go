package story

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/log/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/text"
)

func logDefinitions() []Definition {
	config := DefaultConfig()
	config.Columns, config.Rows = 120, 38

	return []Definition{
		{ID: "text/presentation", Title: "Atoms/Text/Presentation", Level: "Atoms", Family: "text", Description: "Glyph and label, glyph only, or label only; shared policy for semantic atoms.", Source: "text/presentation.go", Symbol: "Format", Defaults: DefaultConfig(), Factory: presentationStory},
		{ID: "text/url", Title: "Atoms/HTTP/URL", Level: "Atoms", Family: "http", Description: "Path, path with query, host with path, or exact captured URL display; no network requests or mutation.", Source: "text/url.go", Symbol: "FormatURL", Defaults: DefaultConfig(), Factory: urlStory},
		{ID: "feedback/follow", Title: "Atoms/Feedback/Follow", Level: "Atoms", Family: "feedback", Description: "Follow mode and paused browsing reuse lifecycle status colors, glyphs and insets.", Source: "feedback/controls.go", Symbol: "Follow", Composes: []string{"feedback/status"}, Defaults: DefaultConfig(), Factory: followStory},
		{ID: "feedback/scrollbar", Title: "Atoms/Feedback/Scrollbar", Level: "Atoms", Family: "feedback", Description: "Proportional vertical line scrollbar with bounded offsets, ASCII fallback and a stable empty gutter.", Source: "feedback/controls.go", Symbol: "View", Defaults: DefaultConfig(), Factory: scrollbarStory},
		{ID: "text/disclosure", Title: "Atoms/Text/Disclosure", Level: "Atoms", Family: "text", Description: "Expanded and collapsed fold labels with shared glyphs and escaped text.", Source: "text/disclosure.go", Symbol: "Disclosure", Defaults: DefaultConfig(), Factory: disclosureStory},
		{ID: "text/key-hint", Title: "Atoms/Text/Key hint", Level: "Atoms", Family: "text", Description: "Shared keyboard-action atom used by inspector controls.", Source: "text/text.go", Symbol: "KeyHint", Defaults: DefaultConfig(), Factory: keyHintStory},
		{ID: "text/pending", Title: "Atoms/Logs/Pending", Level: "Atoms", Family: "logs", Description: "Compact pending glyph and fixed-width caller-driven ellipsis frames.", Source: "text/observability.go", Symbol: "Pending", Step: true, Defaults: DefaultConfig(), Factory: pendingStory},
		{ID: "text/observation", Title: "Atoms/Logs/Observation", Level: "Atoms", Family: "logs", Description: "Shared glyphs and labels for logs, traces, spans, and service-relative direction.", Source: "text/observability.go", Symbol: "Log", Defaults: DefaultConfig(), Factory: observationStory},
		{ID: "text/time", Title: "Atoms/Logs/Time", Level: "Atoms", Family: "logs", Description: "Millisecond clock labels and rounded duration abbreviations with exact observations retained by callers.", Source: "text/observability.go", Symbol: "ShortDuration", Defaults: DefaultConfig(), Factory: abbreviatedTimeStory},
		{ID: "text/http-method", Title: "Atoms/HTTP/Method", Level: "Atoms", Family: "http", Description: "Categorical HTTP verbs with explicit extension-method fallback.", Source: "text/http.go", Symbol: "HTTPMethod", Defaults: DefaultConfig(), Factory: httpMethodStory},
		{ID: "text/http-status", Title: "Atoms/HTTP/Status code", Level: "Atoms", Family: "http", Description: "Response classes retain exact codes, including unknown and invalid observations.", Source: "text/http.go", Symbol: "HTTPStatusCode", Defaults: DefaultConfig(), Factory: httpStatusStory},
		{ID: "text/syntax", Title: "Atoms/Text/Syntax", Level: "Atoms", Family: "text", Description: "Content-type aware lexical highlighting; terminal controls stay visible as escaped data.", Source: "text/syntax.go", Symbol: "Syntax", Defaults: DefaultConfig(), Factory: syntaxStory},
		{ID: "collection/log-tree", ExampleSource: "docs/snippets/log-tree.go.txt", Title: "Molecules/Logs/Structured log tree", Level: "Molecules", Family: "logs", Description: "Trace and span hierarchy with expandable ordered fields, stable identities and retained follow state.", Source: "story/logs.go", Symbol: "logTreeStory", Composes: []string{"collection/tree", "text/observation", "text/time", "text/disclosure", "feedback/scrollbar"}, Controls: []string{"state"}, Step: true, Defaults: config, Factory: logTreeStory},
		{ID: "collection/http-log", ExampleSource: "docs/snippets/http-log.go.txt", Title: "Molecules/Logs/HTTP log", Level: "Molecules", Family: "logs", Description: "Request summaries and captured headers/bodies with syntax highlighting, pretty/raw views and explicit capture states.", Source: "story/logs.go", Symbol: "httpLogStory", Composes: []string{"collection/table", "text/http-method", "text/url", "text/http-status", "text/pending", "text/time", "text/disclosure", "feedback/scrollbar", "text/syntax", "collection/viewport"}, Controls: []string{"state"}, Step: true, Defaults: config, Factory: httpLogStory},
		{ID: "pattern/log-inspector", ExampleSource: "docs/snippets/log-inspector.go.txt", Title: "Organisms/Logs/Correlated inspector", Level: "Organisms", Family: "logs", Description: "Select HTTP requests to inspect exact request logs, widen to trace scope, or browse switchable narrow views.", Source: "story/logs.go", Symbol: "logInspectorStory", Composes: []string{"collection/http-log", "collection/log-tree", "feedback/follow", "text/presentation", "text/key-hint", "foundations/layout"}, Controls: []string{"state"}, Step: true, Defaults: config, Factory: logInspectorStory},
	}
}

func httpMethodStory(e Environment) (runtime.Scene, error) {
	return static(func(layout.Size) string {
		rows := "HTTP methods / categorical colors\n\n"
		for _, method := range []string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "CONNECT", "TRACE", "PROPFIND", "QUERY", ""} {
			rows += text.HTTPMethod(method, e.Config.Palette()) + "\n"
		}

		return rows
	}), nil
}

func httpStatusStory(e Environment) (runtime.Scene, error) {
	return static(func(layout.Size) string {
		rows := "HTTP response codes / exact observations\n\n"
		for _, code := range []int{100, 200, 201, 204, 301, 304, 400, 401, 404, 429, 500, 502, 503, 0, 799} {
			rows += text.HTTPStatusCode(code, e.Config.Palette()) + " " + http.StatusText(code) + "\n"
		}

		return rows
	}), nil
}

func syntaxStory(e Environment) (runtime.Scene, error) {
	return static(func(layout.Size) string {
		return "application/problem+json; charset=utf-8\n\n" + text.Syntax("{\n  \"error\": \"upstream_timeout\",\n  \"retryable\": true,\n  \"attempt\": 2\n}", "json", e.Config.Palette())
	}), nil
}

// LogFixture supplies deterministic observations shared by molecule and organism specimens.
func LogFixture() ([]collection.HTTPExchange, []collection.LogRecord) {
	at := time.Date(2026, 11, 8, 12, 4, 0, 0, time.UTC)
	requests := []collection.HTTPExchange{
		{ID: "http-1", RequestID: "req-81", TraceID: "trace-c7", Time: at, Method: "GET", URL: "https://api.example.test/v1/parcels", StatusCode: 200, Completed: true, Duration: 12 * time.Millisecond, DurationKnown: true, ResponseHeaders: http.Header{"Content-Type": {"application/json"}}, ResponseBody: collection.HTTPBody{State: collection.Captured, ContentType: "application/json", Bytes: []byte(`{"parcels":[],"count":0}`)}},
		{ID: "http-2", RequestID: "req-82", TraceID: "trace-c7", Time: at.Add(time.Second), Method: "POST", URL: "https://api.example.test/v1/labels", StatusCode: 502, Completed: true, Duration: 842 * time.Millisecond, DurationKnown: true, RequestHeaders: http.Header{"Content-Type": {"application/json"}, "X-Request-Id": {"req-82"}}, ResponseHeaders: http.Header{"Content-Type": {"application/problem+json; charset=utf-8"}, "Retry-After": {"1"}}, RequestBody: collection.HTTPBody{State: collection.Captured, ContentType: "application/json", Bytes: []byte(`{"parcel_id":"p_42","service":"express"}`)}, ResponseBody: collection.HTTPBody{State: collection.Captured, ContentType: "application/problem+json; charset=utf-8", Bytes: []byte(`{"error":"upstream_timeout","retryable":true,"attempt":2}`)}},
		{ID: "http-3", RequestID: "req-83", TraceID: "trace-d1", Time: at.Add(2 * time.Second), Method: "PATCH", URL: "https://api.example.test/v1/parcels/p_42", StatusCode: 204, Completed: true, Duration: 8 * time.Millisecond, DurationKnown: true, ResponseBody: collection.HTTPBody{State: collection.Captured}},
		{ID: "http-4", Time: at.Add(3 * time.Second), Method: "GET", URL: "https://api.example.test/health", Completed: false},
	}
	logs := []collection.LogRecord{
		{ID: "log-1", RequestID: "req-81", TraceID: "trace-c7", SpanID: "http-list", LogEntry: collection.LogEntry{Time: at, Level: log.InfoLevel, Message: "Listed parcels", Fields: []any{"count", 0}}},
		{ID: "log-2", RequestID: "req-82", TraceID: "trace-c7", SpanID: "http-label", LogEntry: collection.LogEntry{Time: at.Add(time.Second), Level: log.InfoLevel, Message: "Create label", Fields: []any{"parcel", map[string]any{"id": "p_42", "service": "express"}, "attempt", 2}}},
		{ID: "log-3", RequestID: "req-82", TraceID: "trace-c7", SpanID: "carrier", ParentSpanID: "http-label", LogEntry: collection.LogEntry{Time: at.Add(1200 * time.Millisecond), Level: log.WarnLevel, Message: "Carrier response delayed", Fields: []any{"upstream", map[string]any{"host": "carrier.example", "timeout_ms": 800}}}},
		{ID: "log-4", RequestID: "req-82", TraceID: "trace-c7", SpanID: "carrier", ParentSpanID: "http-label", LogEntry: collection.LogEntry{Time: at.Add(1842 * time.Millisecond), Level: log.ErrorLevel, Message: "Upstream timeout", Fields: []any{"retryable", true, "causes", []any{"connect timeout", "retry exhausted"}}}},
	}

	return requests, logs
}

func logTreeStory(e Environment) (runtime.Scene, error) {
	_, records := LogFixture()
	if e.Config.State == "empty" {
		records = nil
	}

	if e.Config.State == "error" && len(records) > 1 {
		records[1].ParentSpanID = "missing-parent"
	}

	m := collection.NewLogTree(e.Config.Columns, e.Config.Rows-2, e.Config.Palette())
	if err := m.SetEntries(records); err != nil {
		return nil, err
	}

	m.Follow = false
	m.SetSymbols(text.Presentation(e.Config.SymbolStyle))
	m.SetDetachedDetails(e.Config.DetachedLogDetails)
	m.SetFolding(!e.Config.NoFolds)

	for _, node := range m.Model.AllNodes() {
		if value := node.Value(); value != "" {
			node.Open()
		}
	}

	m.Model.SetNodes(m.Model.Root())
	m.Select("log-2")

	step := 0

	return &scene{resize: func(s layout.Size) { m.Resize(s.Width, max(1, s.Height-2)) }, update: func(msg tea.Msg) tea.Cmd {
		if _, ok := msg.(runtime.StepMsg); ok {
			step++
			return logStep(m, step)
		}

		return m.Update(msg)
	}, view: func(layout.Size) string {
		t := e.Config.Palette()

		hints := []string{text.KeyHint("g", "flat/trace", t), text.KeyHint("End", "follow", t), text.KeyHint("Ctrl+N", "append", t)}

		if m.CanFocusDetails() {
			target := "details"
			if m.DetailsFocused {
				target = "tree"
			}

			hints = []string{text.KeyHint("D", target, t), text.KeyHint("g", "group", t), text.KeyHint("End", "follow", t)}
		}

		if m.Folding && !m.CanFocusDetails() {
			hints = append([]string{text.KeyHint("Enter", "expand", t)}, hints...)
		}

		return m.View() + "\n" + strings.Join(hints, " · ")
	}}, nil
}

func logStep(m *collection.LogTree, step int) tea.Cmd {
	err := m.Upsert(collection.LogRecord{ID: fmt.Sprintf("step-%d", step), LogEntry: collection.LogEntry{Time: time.Date(2026, 11, 8, 12, 5, step, 0, time.UTC), Level: log.InfoLevel, Message: "Incoming log observation"}})
	if err != nil {
		return func() tea.Msg { return err }
	}

	return nil
}

func httpLogStory(e Environment) (runtime.Scene, error) {
	records, _ := LogFixture()
	if e.Config.Timeline {
		records = TimelineFixture()
	}

	if e.Config.State == "empty" {
		records = nil
	}

	if e.Config.State == "error" {
		records[1].ResponseBody.Bytes = []byte(`{"error":`)
	}

	m := collection.NewHTTPLog(e.Config.Columns, e.Config.Rows-2, e.Config.Palette())
	m.Timeline = e.Config.Timeline

	m.SetTimelineGrouping(timelineGroups(e.Config.TimelineGroups)...)
	m.TimelineLegend = e.Config.Legend
	m.TimelineNow = time.Date(2026, 11, 8, 12, 4, 4, 0, time.UTC)
	m.Symbols = text.Presentation(e.Config.SymbolStyle)
	m.URLFormat = text.URLFormat(e.Config.URLStyle)
	m.VerboseDetails = e.Config.VerboseDetails

	m.ShowColumnLabels = e.Config.ColumnLabels
	if err := m.SetEntries(records); err != nil {
		return nil, err
	}

	m.Follow = false
	m.Select("http-2")

	if e.Config.State == "loading" {
		m.Select("http-4")
	}

	m.RequestOpen = false
	m.SetFolding(!e.Config.NoFolds)
	m.Raw = e.Config.RawBody
	m.Resize(e.Config.Columns, e.Config.Rows-2)

	step := 0

	return &scene{resize: func(s layout.Size) { m.Resize(s.Width, max(1, s.Height-2)) }, update: func(msg tea.Msg) tea.Cmd {
		if _, ok := msg.(runtime.StepMsg); ok {
			step++

			record := stepExchange(step)
			if err := m.Upsert(record); err != nil {
				return func() tea.Msg { return err }
			}

			return nil
		}

		return m.Update(msg)
	}, view: func(layout.Size) string {
		t := e.Config.Palette()

		hints := []string{text.KeyHint("Tab", "list/details", t)}
		if m.CanToggleRaw() {
			hints = append(hints, text.KeyHint("r", "raw", t))
		}

		if m.Folding {
			hints = append(hints, text.KeyHint("1/2", "sections", t))
		}

		return m.View() + "\n" + strings.Join(hints, " · ")
	}}, nil
}

func logInspectorStory(e Environment) (runtime.Scene, error) {
	requests, logs := LogFixture()
	if e.Config.InspectorLayout == "timeline" {
		requests = TimelineFixture()
	}

	if e.Config.State == "empty" {
		requests = nil
		logs = nil
	}

	m, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Theme: e.Config.Palette(), Requests: requests, Logs: logs, Symbols: text.Presentation(e.Config.SymbolStyle), URLFormat: text.URLFormat(e.Config.URLStyle), VerboseDetails: e.Config.VerboseDetails, ColumnLabels: e.Config.ColumnLabels, DetachedLogDetails: e.Config.DetachedLogDetails, Layout: pattern.InspectorLayout(e.Config.InspectorLayout), TimelineGroups: timelineGroups(e.Config.TimelineGroups), TimelineNow: time.Date(2026, 11, 8, 12, 4, 4, 0, time.UTC)})
	if err != nil {
		return nil, err
	}

	m.SelectRequest("http-2")

	if e.Config.State == "loading" {
		m.SelectRequest("http-4")
	}

	if e.Config.State == "error" {
		requests[1].ResponseBody.Bytes = []byte(`{"error":`)
		if err := m.UpsertRequests(requests[1]); err != nil {
			return nil, err
		}
	}

	m.HTTP.RequestOpen = false
	m.HTTP.SetFolding(!e.Config.NoFolds)
	m.Logs.SetFolding(!e.Config.NoFolds)

	m.HTTP.Raw = e.Config.RawBody
	if e.Config.TraceLogs {
		m.SetScope(pattern.TraceScope)
	}

	if e.Config.LogPane != "" {
		m.FocusPane(e.Config.LogPane)
	}

	m.HTTP.ResizeDetails(70, 18)

	m.Logs.Follow = false
	m.Logs.Select("log-2")

	if e.Config.DetachedLogDetails {
		m.Logs.Select("log-3")
	}

	if m.Logs.Folding {
		m.Logs.Model.ToggleCurrentNode()
	}

	step := 0

	return &scene{init: m.Init, resize: m.Resize, view: func(layout.Size) string { return m.View() }, update: func(msg tea.Msg) tea.Cmd {
		if _, ok := msg.(runtime.StepMsg); ok {
			step++
			if err := m.UpsertRequests(stepExchange(step)); err != nil {
				return func() tea.Msg { return err }
			}

			if err := m.UpsertLogs(collection.LogRecord{ID: fmt.Sprintf("incoming-%d", step), RequestID: "req-82", TraceID: "trace-c7", SpanID: "http-label", LogEntry: collection.LogEntry{Time: time.Date(2026, 11, 8, 12, 5, step, 0, time.UTC), Level: log.InfoLevel, Message: "Retry observation", Fields: []any{"attempt", step + 2}}}); err != nil {
				return func() tea.Msg { return err }
			}

			return nil
		}

		return m.Update(msg)
	}}, nil
}

func stepExchange(step int) collection.HTTPExchange {
	return collection.HTTPExchange{ID: fmt.Sprintf("incoming-%d", step), RequestID: fmt.Sprintf("req-live-%d", step), Time: time.Date(2026, 11, 8, 12, 5, step, 0, time.UTC), Method: "GET", URL: "https://api.example.test/v1/events", StatusCode: 200, Completed: true, Duration: 4 * time.Millisecond, DurationKnown: true, ResponseBody: collection.HTTPBody{State: collection.Captured, ContentType: "application/json", Bytes: []byte(`{"events":[]}`)}}
}

func pendingStory(e Environment) (runtime.Scene, error) {
	var frame uint

	return &scene{update: func(msg tea.Msg) tea.Cmd {
		if _, ok := msg.(runtime.StepMsg); ok {
			frame++
		}

		return nil
	}, view: func(layout.Size) string {
		t := e.Config.Palette()
		return "Pending / stable default\n\n" + text.Pending(t) + " Pending\n\nCaller-driven ellipsis / Ctrl+N advances\n\n" + text.Pending(t, frame) + " Pending\n\nThree equal-width frames\n\n" + text.Pending(t, 0) + " | " + text.Pending(t, 1) + " | " + text.Pending(t, 2)
	}}, nil
}

func observationStory(e Environment) (runtime.Scene, error) {
	return static(func(layout.Size) string {
		t := e.Config.Palette()
		return "Observation identities\n\n" + text.Log("Log entry", t) + "\n" + text.Trace("Trace trace-c7", t) + "\n" + text.Span("Span carrier", t) + "\n\nService-relative direction\n\n" + text.Upstream("Upstream carrier.example", t) + "\n" + text.Downstream("Downstream API consumer", t) + "\n\nASCII fallbacks: L / T / S / ^ / v"
	}), nil
}

func abbreviatedTimeStory(e Environment) (runtime.Scene, error) {
	return static(func(layout.Size) string {
		t := e.Config.Palette()
		at := time.Date(2026, 11, 8, 12, 4, 1, 234567000, time.UTC)

		rows := "Abbreviated time / exact source retained\n\n" + text.TimeOfDay(at) + "  from " + at.Format(time.RFC3339Nano) + "\n" + text.TimeOfDay(time.Time{}) + "  unavailable\n\nLatency / compact rounded display\n\n"
		for _, d := range []time.Duration{420 * time.Nanosecond, 842 * time.Microsecond, 842 * time.Millisecond, 1234567890, 62 * time.Second, -1234567890} {
			rows += text.Label(fmt.Sprintf("%8s  from %s", text.ShortDuration(d), d), t) + "\n"
		}

		return rows
	}), nil
}

func followStory(e Environment) (runtime.Scene, error) {
	return static(func(layout.Size) string {
		t := e.Config.Palette()
		return "Browsing mode / shared status atoms\n\n" + feedback.Follow(true, text.Normal, t) + "\n" + feedback.Follow(false, text.Normal, t) + "\n\n" + text.KeyHint("F", "resume follow", t)
	}), nil
}

func scrollbarStory(e Environment) (runtime.Scene, error) {
	return static(func(layout.Size) string {
		t := e.Config.Palette()
		bars := []string{}

		for _, offset := range []int{0, 40, 90} {
			bar := feedback.Scrollbar{Total: 100, Visible: 10, Offset: offset, Height: 10}
			bars = append(bars, bar.View(t))
		}

		return "Scrollbar / top · middle · bottom\n\n" + layout.Row(bars[0], "     ", bars[1], "     ", bars[2]) + "\n\nThumb = visible lines / total lines\nNo overflow = reserved blank gutter"
	}), nil
}

func disclosureStory(e Environment) (runtime.Scene, error) {
	return static(func(layout.Size) string {
		t := e.Config.Palette()
		return "Optional disclosure controls\n\n" + text.Disclosure("Response", false, t) + "\n" + text.Disclosure("Response", true, t) + "\n\nFolding disabled: all content stays open\n\n" + t.Heading().Render("Response")
	}), nil
}

func keyHintStory(e Environment) (runtime.Scene, error) {
	return static(func(layout.Size) string {
		t := e.Config.Palette()
		return "Keyboard action atoms\n\n" + text.KeyHint("Tab", "next pane", t) + "\n" + text.KeyHint("F", "resume follow", t) + "\n" + text.KeyHint("Enter", "toggle fold", t) + "\n" + text.KeyHint("r", "raw / pretty", t)
	}), nil
}

func presentationStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		t := e.Config.Palette()

		groups := []string{}
		for _, mode := range []text.Presentation{text.GlyphAndLabel, text.GlyphOnly, text.LabelOnly} {
			groups = append(groups, string(mode)+"\n\n"+feedback.Paused.View(text.Normal, t, mode)+"\n"+text.Trace("Trace", t, mode)+"\n"+text.Disclosure("Response", true, t, mode))
		}

		if size.Width < 70 {
			return strings.Join(groups, "\n\n")
		}

		return "Semantic atom presentation\n\n" + layout.Row(layout.Fit(groups[0], layout.Size{Width: 24, Height: 6}), layout.Fit(groups[1], layout.Size{Width: 24, Height: 6}), groups[2])
	}), nil
}

func urlStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		captured := "https://api.example.test:8443/v1/labels/p_42?limit=25&tag=a%2Fb#detail"

		rows := []string{"URL display / captured source unchanged", ""}
		for _, format := range []text.URLFormat{text.PathOnly, text.PathAndQuery, text.HostAndPath, text.FullURL} {
			rows = append(rows, string(format), text.FormatURL(captured, format), "")
		}

		return ansi.Wrap(strings.Join(rows, "\n"), max(1, size.Width), "")
	}), nil
}

// TimelineFixture includes overlapping completed and pending requests.
func TimelineFixture() []collection.HTTPExchange {
	records, _ := LogFixture()
	at := records[0].Time
	records[0].Service = "api"
	records[1].Service = "api"
	records[2].Service = "carrier"
	records[2].URL = "https://carrier.example:8443/v1/shipments/p_42"
	records[3].Service = "api"
	records[0].Duration = 2200 * time.Millisecond
	records[1].Time = at.Add(350 * time.Millisecond)
	records[1].Duration = 2800 * time.Millisecond
	records[2].Time = at.Add(900 * time.Millisecond)
	records[2].Duration = 600 * time.Millisecond
	records[3].Time = at.Add(1600 * time.Millisecond)

	return append(records, collection.HTTPExchange{ID: "http-5", Service: "carrier", Time: at.Add(2500 * time.Millisecond), Method: "QUERY", URL: "https://carrier.example:8443/v1/rates", Completed: true, StatusCode: 200, DurationKnown: true, Duration: 900 * time.Millisecond}, collection.HTTPExchange{ID: "http-6", Service: "events", Time: at.Add(1800 * time.Millisecond), Method: "POST", URL: "http://events.internal:9090/events", Completed: true, StatusCode: 202, DurationKnown: true, Duration: 600 * time.Millisecond})
}

func timelineGroups(names []string) []collection.HTTPGroup {
	var groups []collection.HTTPGroup

	for _, group := range names {
		switch group {
		case "service":
			groups = append(groups, collection.GroupByService)
		case "host":
			groups = append(groups, collection.GroupByHost)
		case "port":
			groups = append(groups, collection.GroupByPort)
		case "endpoint":
			groups = append(groups, collection.GroupByEndpoint)
		}
	}

	return groups
}
