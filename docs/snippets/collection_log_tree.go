package snippets

import (
	"time"

	"charm.land/log/v2"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/theme"
)

func ExampleCollectionLogTree() (string, error) {
	at := time.Date(2026, 11, 8, 12, 4, 0, 0, time.UTC)
	records := []collection.LogRecord{{
		ID: "log-1", RequestID: "req-82", TraceID: "trace-c7", SpanID: "carrier",
		LogEntry: collection.LogEntry{
			Time: at, Level: log.WarnLevel,
			Message: "Carrier response delayed", Fields: []any{"timeout_ms", 800},
		},
	}}

	logs := collection.NewLogTree(64, 12, theme.New(theme.Dark))
	if err := logs.SetEntries(records); err != nil {
		return "", err
	}

	logs.Select("log-1")

	return logs.View(), nil
}
