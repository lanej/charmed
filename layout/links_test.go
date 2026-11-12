package layout_test

import (
	"slices"
	"strings"
	"testing"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestLinkGroupFlowsWholeItemsAndKeepsOrderedTargets(t *testing.T) {
	group := layout.LinkGroup{Links: []text.Link{
		{Label: "One", URL: "https://example.test/one"},
		{Label: "Two", URL: "https://example.test/two", Emphasis: text.Dim},
		{Label: "界é", URL: "terminal://session/tests"},
	}}
	original := append([]text.Link(nil), group.Links...)
	palette := theme.New(theme.Dark)

	for _, tc := range []struct {
		width int
		want  string
	}{
		{12, "↗ One  ↗ Two\n↗ 界é"},
		{11, "↗ One\n↗ Two\n↗ 界é"},
		{30, "↗ One  ↗ Two  ↗ 界é"},
	} {
		view, err := group.View(tc.width, palette)
		if err != nil {
			t.Fatal(err)
		}

		if ansi.Strip(view) != tc.want {
			t.Fatalf("width %d: got %q want %q", tc.width, ansi.Strip(view), tc.want)
		}

		var targets []string

		for _, row := range strings.Split(view, "\n") {
			if ansi.StringWidth(row) > tc.width {
				t.Fatal("group escaped its cell allocation")
			}

			line := uv.NewStyledString(row + "!").Lines(ansi.GraphemeWidth)[0]
			for _, cell := range line {
				if cell.Content == "!" && cell.Link.URL != "" {
					t.Fatal("group leaked its target onto adjacent output")
				}

				if cell.Link.URL != "" && (len(targets) == 0 || targets[len(targets)-1] != cell.Link.URL) {
					targets = append(targets, cell.Link.URL)
				}
			}
		}

		wantTargets := []string{original[0].URL, original[1].URL, original[2].URL}
		if !slices.Equal(targets, wantTargets) || !slices.Equal(group.Links, original) {
			t.Fatal("group changed target order or mutated its input")
		}
	}

	plain, err := group.Plain()
	if err != nil || plain != "One (https://example.test/one)\nTwo (https://example.test/two)\n界é (terminal://session/tests)" {
		t.Fatalf("group lost plain destinations: %q, %v", plain, err)
	}
}

func TestLinkGroupWrapsOversizedLabelsWithoutSplittingGraphemes(t *testing.T) {
	group := layout.LinkGroup{Links: []text.Link{{Label: "A界éBC", URL: "custom:destination", Presentation: text.LabelOnly}}}

	view, err := group.View(3, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	if ansi.Strip(view) != "A界\néBC" {
		t.Fatalf("oversized label lost content or graphemes: %q", ansi.Strip(view))
	}

	for _, line := range uv.NewStyledString(view).Lines(ansi.GraphemeWidth) {
		for _, cell := range line {
			if cell.Width > 0 && cell.Link.URL != "custom:destination" {
				t.Fatal("wrapped label lost its target")
			}
		}
	}

	for _, width := range []int{0, -1, 1} {
		if got, err := group.View(width, theme.New(theme.Dark)); err == nil || got != "" {
			t.Fatalf("impossible width %d silently lost content: %q, %v", width, got, err)
		}
	}
}

func TestLinkGroupErrorsAreAtomicAndEmptyGroupsAreEmpty(t *testing.T) {
	group := layout.LinkGroup{Links: []text.Link{{Label: "Valid", URL: "https://example.test"}, {Label: "Broken", URL: "/relative"}}}
	if got, err := group.View(80, theme.New(theme.Dark)); err == nil || got != "" || !strings.Contains(err.Error(), "item 1") {
		t.Fatal("invalid group returned partial output or lost the failing item")
	}

	if got, err := group.Plain(); err == nil || got != "" {
		t.Fatal("plain group returned partial output")
	}

	empty := layout.LinkGroup{}
	if got, err := empty.View(0, theme.New(theme.Dark)); err != nil || got != "" {
		t.Fatal("empty group rendered content")
	}

	if got, err := empty.Plain(); err != nil || got != "" {
		t.Fatal("empty plain group rendered content")
	}
}
