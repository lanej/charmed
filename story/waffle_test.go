package story_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/story"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestWaffleLegendKeepsCategoryKeysAtNarrowWidths(t *testing.T) {
	definition, err := story.Find("chart/waffle")
	if err != nil {
		t.Fatal(err)
	}

	for _, width := range []int{20, 40, 80} {
		config := definition.Defaults
		config.Columns = width
		config.Theme = theme.Mono

		frame, err := story.Frame(definition.ID, config)
		if err != nil {
			t.Fatal(err)
		}

		for _, key := range []string{"A ready 92", "B pending 3", "C failed 5"} {
			if !strings.Contains(ansi.Strip(frame), key) {
				t.Fatalf("%d-column waffle lost legend key %q", width, key)
			}
		}
	}
}

func TestWaffleTableKeepsCountsBarsAndStatusAtoms(t *testing.T) {
	definition, err := story.Find("chart/waffle-header")
	if err != nil {
		t.Fatal(err)
	}

	for _, width := range []int{40, 80} {
		config := definition.Defaults
		config.Columns = width
		config.Theme = theme.Mono

		frame, err := story.Frame(definition.ID, config)
		if err != nil {
			t.Fatal(err)
		}

		plain := ansi.Strip(frame)
		if !strings.Contains(plain, "Instance scale: 0-40") {
			t.Fatal("instance bars lost their shared scale")
		}

		bars := map[string]int{}

		for _, expected := range []struct {
			name   string
			count  int
			status feedback.Status
		}{
			{"checkout", 24, feedback.OK},
			{"labels", 28, feedback.OK},
			{"rating", 40, feedback.OK},
			{"tracking", 3, feedback.Pending},
			{"events", 5, feedback.Error},
		} {
			found := false

			for _, line := range strings.Split(plain, "\n") {
				if !strings.HasPrefix(line, expected.name+" ") {
					continue
				}

				found = true
				fields := strings.Fields(line)

				atom := ansi.Strip(expected.status.View(text.Normal, theme.New(theme.Mono)))
				if len(fields) < 2 || fields[1] != strconv.Itoa(expected.count) || !strings.Contains(line, atom) {
					t.Fatalf("%d-column row lost its count or status atom: %q", width, line)
				}

				bars[expected.name] = strings.Count(line, "#")
			}

			if !found {
				t.Fatalf("%d-column table lost %s", width, expected.name)
			}
		}

		if bars["rating"] <= bars["labels"] || bars["labels"] <= bars["checkout"] || bars["checkout"] <= bars["events"] || bars["events"] < bars["tracking"] || bars["tracking"] <= 0 {
			t.Fatalf("%d-column bar magnitudes do not follow the counts: %v", width, bars)
		}
	}
}
