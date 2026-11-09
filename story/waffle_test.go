package story_test

import (
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/story"
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
