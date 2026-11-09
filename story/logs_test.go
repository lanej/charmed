package story_test

import (
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/story"
	"github.com/lanej/charmed/theme"
)

func TestLogInspectorSavedPanesAreNativeAndDeterministic(t *testing.T) {
	for pane, expected := range map[string]string{"requests": "HTTP requests", "http": "upstream_timeout", "logs": "Carrier"} {
		for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Mono} {
			config := story.DefaultConfig()
			config.Columns, config.Rows = 40, 24
			config.LogPane = pane
			config.Theme = mode

			first, err := story.Frame("pattern/log-inspector", config)
			if err != nil {
				t.Fatal(err)
			}

			second, err := story.Frame("pattern/log-inspector", config)
			if err != nil {
				t.Fatal(err)
			}

			if first != second || !strings.Contains(ansi.Strip(first), expected) {
				t.Fatalf("%s/%s lost deterministic native content %q", pane, mode, expected)
			}
		}
	}
}

func TestLogFixtureFlagsRejectUnknownPane(t *testing.T) {
	config := story.DefaultConfig()

	config.LogPane = "missing"
	if err := config.Validate(); err == nil {
		t.Fatal("unknown inspector pane accepted")
	}
}

func TestHTTPTimelineFixtureEmpty(t *testing.T) {
	config := story.DefaultConfig()
	config.Timeline = true
	config.State = "empty"

	frame, err := story.Frame("collection/http-log", config)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(ansi.Strip(frame), "No HTTP requests") {
		t.Fatalf("timeline lost empty state: %q", frame)
	}
}
