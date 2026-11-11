package story_test

import (
	"image/color"
	"strings"
	"testing"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/story"
	"github.com/lanej/charmed/theme"
)

func TestInputFramesInitializeFocusWithoutRunningCommands(t *testing.T) {
	config := story.DefaultConfig()
	config.Columns, config.Rows = 40, 16
	palette := config.Palette()

	frame, err := story.Frame("input/text", config)
	if err != nil {
		t.Fatal(err)
	}

	cursor := false

	for _, line := range uv.NewStyledString(frame).Lines(ansi.GraphemeWidth) {
		for _, cell := range line {
			if cell.Style.Attrs&uv.AttrReverse != 0 && cell.Style.Fg != nil && color.NRGBAModel.Convert(cell.Style.Fg) == color.NRGBAModel.Convert(palette.Accent) {
				cursor = true
			}
		}
	}

	if !cursor {
		t.Fatal("saved focused input did not render its virtual cursor")
	}

	config.State = "disabled"

	disabled, err := story.Frame("input/text", config)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(ansi.Strip(disabled), palette.Glyph(theme.GlyphDisabled)+" production") {
		t.Fatal("saved disabled input lost its glyph")
	}
}

func TestBandLabelsStayAtAllocatedEndpoints(t *testing.T) {
	for _, width := range []int{20, 40, 80, 120} {
		config := story.DefaultConfig()
		config.Columns = width

		frame, err := story.Frame("chart/interval-band", config)
		if err != nil {
			t.Fatal(err)
		}

		found := false

		for _, line := range strings.Split(ansi.Strip(frame), "\n") {
			if strings.HasPrefix(line, "0") && strings.HasSuffix(line, "100") && ansi.StringWidth(line) == width {
				found = true
			}
		}

		if !found {
			t.Fatalf("band endpoints do not fill %d cells", width)
		}
	}
}

func TestNarrowPermissionsKeepUnknownAndSpecialBitsVisible(t *testing.T) {
	config := story.DefaultConfig()
	config.Columns = 40

	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		config.Theme = mode

		frame, err := story.Frame("foundations/permissions", config)
		if err != nil {
			t.Fatal(err)
		}

		for _, expected := range []string{"uninspected", "?????????", "setuid", "rwsr-xr-x", "setgid", "rw-r-Sr--", "sticky", "rwxrwxrwt", "not effective access"} {
			if !strings.Contains(ansi.Strip(frame), expected) {
				t.Fatalf("%s narrow permissions clipped %q", mode, expected)
			}
		}
	}
}

func TestNarrowLinkStoryRetainsPlainDestination(t *testing.T) {
	config := story.DefaultConfig()
	config.Columns, config.Rows = 40, 16

	frame, err := story.Frame("text/link", config)
	if err != nil {
		t.Fatal(err)
	}

	content := strings.Join(strings.Fields(ansi.Strip(frame)), "")
	if !strings.Contains(content, "https://github.com/lanej/charmed/pull/42") {
		t.Fatal("narrow link fixture clipped its explicit plain destination")
	}
}
