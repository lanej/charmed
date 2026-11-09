package snippets

import (
	"time"

	"charm.land/log/v2"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

func ExamplePatternLogInspector() (string, error) {
	at := time.Date(2026, 11, 8, 12, 4, 0, 0, time.UTC)
	records := []collection.LogRecord{{
		ID: "log-1", RequestID: "req-82", TraceID: "trace-c7", SpanID: "carrier",
		LogEntry: collection.LogEntry{
			Time: at, Level: log.WarnLevel,
			Message: "Carrier response delayed", Fields: []any{"timeout_ms", 800},
		},
	}}

	requests := []collection.HTTPExchange{{
		ID: "http-2", RequestID: "req-82", TraceID: "trace-c7", Time: at,
		Method: "POST", URL: "https://api.example.test/v1/labels",
		StatusCode: 502, Completed: true, Duration: 842 * time.Millisecond, DurationKnown: true,
		ResponseBody: collection.HTTPBody{
			State: collection.Captured, ContentType: "application/json",
			Bytes: []byte(`{"error":"upstream_timeout","retryable":true}`),
		},
	}}

	inspector, err := pattern.NewLogInspector(pattern.LogInspectorOptions{
		Theme: theme.New(theme.Dark), Requests: requests, Logs: records,
		NoHelp: true, DetachedLogDetails: true, // The default layout adapts to available space.
	})
	if err != nil {
		return "", err
	}

	inspector.Resize(layout.Size{Width: 100, Height: 30})
	inspector.SelectRequest("http-2")
	// Optional: inspector.SetPaneCollapsed("request", true) gives the response both columns.
	// inspector.SetHelpVisible(true) restores keyboard hints; FocusPane("response") changes focus.
	// Customize payload styles on inspector.Request and inspector.Response.

	return inspector.View(), nil
}
