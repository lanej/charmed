package layout_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func TestTextRowSpacingAndFlexibleSpacers(t *testing.T) {
	for _, tc := range []struct {
		name  string
		row   layout.TextRow
		width int
		want  string
	}{
		{"fixed gap", layout.TextRow{Gap: 2, Items: []layout.TextRowItem{{Text: "A"}, {Text: "B"}}}, 8, "A  B    "},
		{"right aligned", layout.TextRow{Gap: 2, Items: []layout.TextRowItem{{Text: "A"}, {Flex: true}, {Text: "B"}}}, 8, "A      B"},
		{"equal spacers", layout.TextRow{Gap: 1, Items: []layout.TextRowItem{{Text: "A"}, {Flex: true}, {Text: "B"}, {Flex: true}, {Text: "C"}}}, 11, "A    B    C"},
		{"trailing spacer", layout.TextRow{Gap: 1, Items: []layout.TextRowItem{{Text: "A", Flex: true}, {Text: "B", Flex: true}}}, 11, "A     B    "},
		{"optional text", layout.TextRow{Gap: 2, Items: []layout.TextRowItem{{Text: "A"}, {}, {Flex: true}, {Text: "B"}}}, 8, "A      B"},
		{"wrapped trailing group", layout.TextRow{Gap: 2, Items: []layout.TextRowItem{{Text: "Title"}, {Flex: true}, {Text: "Status"}}}, 8, "Title   \n  Status"},
		{"stack", layout.TextRow{Stack: true, Gap: 2, Items: []layout.TextRowItem{{Text: "A"}, {Flex: true}, {Text: "B"}}}, 8, "A       \nB       "},
		{"single impossible grapheme", layout.TextRow{Items: []layout.TextRowItem{{Text: "界"}}}, 1, "!"},
		{"mixed single cells", layout.TextRow{Items: []layout.TextRowItem{{Text: "A界B"}}}, 1, "A\n!\nB"},
		{"empty", layout.TextRow{Items: []layout.TextRowItem{{}, {Flex: true}}}, 8, ""},
		{"zero width", layout.TextRow{Items: []layout.TextRowItem{{Text: "A"}}}, 0, ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			before := append([]layout.TextRowItem(nil), tc.row.Items...)
			if got := tc.row.View(tc.width); got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}

			if !reflect.DeepEqual(before, tc.row.Items) {
				t.Fatal("rendering changed caller-owned items")
			}
		})
	}
}

func TestTextRowStyledWrappingAndCellGeometry(t *testing.T) {
	palette := theme.New(theme.Dark)
	row := layout.TextRow{Gap: 2, Items: []layout.TextRowItem{
		{Text: palette.Heading().Render("界 é")},
		{Text: "A long observation subtitle with exact 3500 / 5000 counts"},
		{Flex: true},
		{Text: palette.Subtle().Render("Observed 00:16 UTC")},
	}}

	for width := 1; width <= 100; width++ {
		view := row.View(width)
		for _, line := range strings.Split(view, "\n") {
			if ansi.StringWidth(line) != width {
				t.Fatalf("width %d escaped: %q", width, line)
			}
		}

		compact := strings.Join(strings.Fields(ansi.Strip(view)), "")
		for _, part := range []string{"é", "Alongobservationsubtitlewithexact3500/5000counts", "Observed00:16UTC"} {
			if !strings.Contains(compact, part) {
				t.Fatalf("width %d dropped text %q: %q", width, part, view)
			}
		}

		if width == 1 && !strings.Contains(view, "!") {
			t.Fatal("unrenderable wide grapheme had no allocation marker")
		}

		if width > 1 && !strings.Contains(view, "界") {
			t.Fatalf("width %d dropped a wide grapheme", width)
		}
	}

	if !strings.Contains(row.View(100), "\x1b[") {
		t.Fatal("row discarded native styling")
	}
}
