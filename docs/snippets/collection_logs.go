package snippets

import (
	"time"

	"charm.land/log/v2"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/theme"
)

func ExampleCollectionLogs() string {
	logs := collection.NewLogViewer(64, 8, theme.New(theme.Dark))
	logs.Append(collection.LogEntry{
		Time:  time.Date(2026, 11, 8, 12, 4, 0, 0, time.UTC),
		Level: log.InfoLevel, Message: "Request completed",
		Fields: []any{"service", "api", "status", 200},
	})
	logs.Append(collection.LogEntry{Level: log.WarnLevel, Message: "Retry scheduled"})

	return logs.View()
}
