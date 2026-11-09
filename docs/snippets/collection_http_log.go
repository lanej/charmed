package snippets

import (
	"time"

	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleCollectionHttpLog() (string, error) {
	at := time.Date(2026, 11, 8, 12, 4, 0, 0, time.UTC)

	requests := []collection.HTTPExchange{{
		ID: "http-2", RequestID: "req-82", TraceID: "trace-c7", Time: at,
		Method: "POST", URL: "https://api.example.test/v1/labels",
		StatusCode: 502, Completed: true, Duration: 842 * time.Millisecond, DurationKnown: true,
		ResponseBody: collection.HTTPBody{
			State: collection.Captured, ContentType: "application/json",
			Bytes: []byte(`{"error":"upstream_timeout","retryable":true}`),
		},
	}}

	logView := collection.NewHTTPLog(64, 24, theme.New(theme.Dark))
	logView.SetURLFormat(text.PathAndQuery)

	if err := logView.SetEntries(requests); err != nil {
		return "", err
	}

	logView.Select("http-2")

	return logView.View(), nil
}
