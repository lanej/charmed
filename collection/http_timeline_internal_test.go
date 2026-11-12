package collection

import (
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/charmbracelet/x/ansi"
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

func TestTimelineStreamingKeepsObservedStatus(t *testing.T) {
	at := time.Date(2026, 11, 8, 12, 0, 0, 0, time.UTC)
	m := NewHTTPLog(100, 30, theme.New(theme.Dark))

	for _, code := range []int{0, 200, 503} {
		record := HTTPExchange{Time: at, StatusCode: code}
		meta := ansi.Strip(m.timelineMetadata(record, at.Add(2*time.Second)))

		want := strconv.Itoa(code)
		if code == 0 {
			want = m.Theme.Glyph(theme.GlyphPending)
		}

		if strings.TrimSpace(ansi.Cut(meta, 0, 5)) != want || !strings.HasSuffix(meta, "2s") {
			t.Fatalf("streaming status%d lost: %q", code, meta)
		}
	}
}
