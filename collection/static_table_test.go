package collection_test

import (
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func staticThemes() map[string]theme.Theme {
	ascii := theme.New(theme.Dark)
	ascii.ASCII = true

	return map[string]theme.Theme{
		"dark": theme.New(theme.Dark), "light": theme.New(theme.Light), "mono": theme.New(theme.Mono), "ascii": ascii,
	}
}

func staticFixture() collection.StaticTable {
	omitted := text.Omission{Kind: text.SourceTruncated, Count: 12, KnownCount: true}

	return collection.StaticTable{
		Columns: []collection.TableColumn{
			{Title: "HOST", Fixed: true}, {Title: "VERSION", Min: 10}, {Title: "STATE"}, {Title: "NOTE", Min: 12},
		},
		Rows: [][]collection.TableCell{
			{
				collection.TableText("web1west"), collection.TableText("9f2c4e1a7b3d5c60918273645a4b3c2d1e0f"),
				collection.TableEvidence(feedback.EvidenceHealthy, ""), collection.TableText("serving traffic from the west pool"),
			},
			{
				collection.TableText("web2west"), collection.TableMissing(),
				collection.TableEvidence(feedback.EvidenceUnknown, ""),
				{Text: "no report", Omission: &omitted},
			},
			{collection.TableText("web3west")},
		},
	}
}

func TestStaticTableNeverExceedsTheAllocation(t *testing.T) {
	for name, palette := range staticThemes() {
		for _, width := range []int{20, 40, 55, 80, 120, 200} {
			for _, line := range strings.Split(staticFixture().View(width, palette), "\n") {
				if got := ansi.StringWidth(line); got > width {
					t.Fatalf("%s/%d: %d cells: %q", name, width, got, ansi.Strip(line))
				}
			}
		}
	}
}

func TestStaticTableStacksOnlyWhenNarrowAndNotWhole(t *testing.T) {
	palette := theme.New(theme.Mono)
	table := staticFixture()

	for width, want := range map[int]bool{20: true, 40: true, 55: true, 56: false, 80: false, 120: false} {
		if got := table.Stacked(width, palette); got != want {
			t.Fatalf("width %d: stacked %v want %v", width, got, want)
		}
	}

	small := collection.StaticTable{
		Columns: []collection.TableColumn{{Title: "HOST"}, {Title: "READ"}},
		Rows:    [][]collection.TableCell{{collection.TableText("web1"), collection.TableText("ok")}},
	}
	if small.Stacked(40, palette) {
		t.Fatal("a table that fits whole must keep its columns")
	}
}

func TestStackedTableKeepsEveryValueWhole(t *testing.T) {
	const version = "9f2c4e1a7b3d5c60918273645a4b3c2d1e0f"

	for name, palette := range staticThemes() {
		view := ansi.Strip(staticFixture().View(55, palette))
		joined := strings.NewReplacer("\n", "", " ", "").Replace(view)

		if !strings.Contains(view, "VERSION:") || !strings.Contains(joined, version) {
			t.Fatalf("%s: stacked table lost the version:\n%s", name, view)
		}

		if strings.Contains(view, "HOST  ") {
			t.Fatalf("%s: stacked table kept columns:\n%s", name, view)
		}
	}
}

func TestStaticTableNeverShowsUnprovenEvidenceAsHealthy(t *testing.T) {
	gaps := []feedback.Evidence{
		feedback.EvidenceUnknown, feedback.EvidenceInsufficient, feedback.EvidenceStale,
		feedback.EvidenceTruncated, feedback.EvidenceAttributionGap, feedback.Evidence("surprise"),
	}

	for name, palette := range staticThemes() {
		for _, width := range []int{40, 80, 120} {
			table := collection.StaticTable{Columns: []collection.TableColumn{{Title: "HOST"}, {Title: "STATE"}}}
			for _, gap := range gaps {
				table.Rows = append(table.Rows, []collection.TableCell{collection.TableText("h"), collection.TableEvidence(gap, "")})
			}

			view := strings.ToLower(ansi.Strip(table.View(width, palette)))
			if strings.Contains(view, "healthy") {
				t.Fatalf("%s/%d rendered a gap as healthy:\n%s", name, width, view)
			}

			for _, word := range []string{"unknown", "insufficient", "stale", "truncated", "attribution gap"} {
				if !strings.Contains(view, word) {
					t.Fatalf("%s/%d lost %q:\n%s", name, width, word, view)
				}
			}
		}
	}
}

func TestStaticTableKeepsMissingDistinctFromEmpty(t *testing.T) {
	table := collection.StaticTable{
		Columns: []collection.TableColumn{{Title: "NAME"}, {Title: "VALUE"}},
		Rows: [][]collection.TableCell{
			{collection.TableText("empty"), collection.TableText("")},
			{collection.TableText("missing"), collection.TableMissing()},
			{collection.TableText("short")},
		},
	}
	view := ansi.Strip(table.View(80, theme.New(theme.Mono)))

	if !strings.Contains(view, `empty    ""`) || strings.Count(view, "missing") != 3 {
		t.Fatalf("missing and empty were conflated:\n%s", view)
	}
}

func TestStaticTableEscapesPlainTextAndReportsEmpty(t *testing.T) {
	palette := theme.New(theme.Mono)
	table := collection.StaticTable{
		Columns: []collection.TableColumn{{Title: "NAME"}},
		Rows:    [][]collection.TableCell{{collection.TableText("a\x1b]0;owned\x07b")}},
	}

	if strings.Contains(table.View(80, palette), "\x1b]") || strings.Contains(table.View(80, palette), "\x07") {
		t.Fatal("control sequence reached the terminal")
	}

	if got := ansi.Strip((collection.StaticTable{}).View(80, palette)); got != "no rows" {
		t.Fatal(got)
	}
}
