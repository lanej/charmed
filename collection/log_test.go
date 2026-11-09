package collection_test

import (
	"strings"
	"testing"
	"time"

	"charm.land/log/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/theme"
)

func TestLogViewerKeepsAtomsFieldsAndBoundedHistory(t *testing.T) {
	viewer := collection.NewLogViewer(100, 2, theme.New(theme.Mono))

	viewer.Limit = 3
	for _, level := range []log.Level{log.DebugLevel, log.InfoLevel, log.WarnLevel, log.ErrorLevel, log.FatalLevel} {
		viewer.Append(collection.LogEntry{Time: time.Date(2026, 11, 9, 9, 10, 11, 0, time.UTC), Level: level, Message: "message", Fields: []any{"job", "one"}})
	}

	viewer.GotoTop()

	top := ansi.Strip(viewer.View())
	if strings.Contains(top, "DEBUG") || strings.Contains(top, "INFO") || !strings.Contains(top, "09:10:11 [WARN ] message job=one") {
		t.Fatalf("bounded structured log lost: %q", top)
	}

	viewer.GotoBottom()

	if bottom := ansi.Strip(viewer.View()); !strings.Contains(bottom, "[FATAL]") || strings.Count(bottom, "FATAL") != 1 {
		t.Fatalf("fatal duplicated or lost: %q", bottom)
	}

	viewer.Append(collection.LogEntry{Level: log.Level(19), Message: "unknown"})

	if !strings.Contains(ansi.Strip(viewer.View()), "LEVEL 19") {
		t.Fatal("unknown entry silently discarded")
	}
}
