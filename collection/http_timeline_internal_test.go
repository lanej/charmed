package collection

import (
	"strings"
	"testing"
	"time"

	"github.com/lanej/charmed/theme"
)

func TestTimelineInvalidStatusColor(t *testing.T) {
	t.Setenv("CLICOLOR_FORCE", "1")

	at := time.Date(2026, 11, 8, 12, 0, 0, 0, time.UTC)
	m := NewHTTPLog(80, 20, theme.New(theme.Dark))
	record := HTTPExchange{Time: at, Completed: true, DurationKnown: true, Duration: time.Second, StatusCode: 999}
	bar := m.timelineBar(record, at, at, time.Second, 10)

	want := logStyle(m.Theme, m.Theme.Unknown).Render("├" + strings.Repeat("━", 8) + "┤")
	if bar != want {
		t.Fatalf("invalid status must use unknown role: got %q want %q", bar, want)
	}
}
