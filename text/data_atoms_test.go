package text_test

import (
	"math"
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestQuantityFormattingKeepsUnitsSignAndMissingDistinct(t *testing.T) {
	for _, tc := range []struct {
		q    text.Quantity
		want string
	}{
		{text.Quantity{}, "0"},
		{text.Quantity{Value: 1258291, Unit: "B", Scale: text.BinaryScale, Precision: 1}, "1.2 MiB"},
		{text.Quantity{Value: 842, Unit: "req/s"}, "842 req/s"},
		{text.Quantity{Value: 1250, Scale: text.DecimalScale, Precision: 1}, "1.2k"},
		{text.Quantity{Value: 999.96, Unit: "B", Scale: text.DecimalScale, Precision: 1}, "1.0 kB"},
		{text.Quantity{Value: -2048, Unit: "B", Scale: text.BinaryScale}, "-2 KiB"},
		{text.Quantity{Value: .0000001, Precision: 2}, "1e-07"},
		{text.Quantity{Value: .5}, "0.5"},
		{text.Quantity{Value: -.5}, "-0.5"},
		{text.Quantity{Value: math.Copysign(0, -1)}, "0"},
		{text.Quantity{Value: math.Inf(1)}, "invalid"},
		{text.Quantity{State: text.ValueMissing}, "missing"},
		{text.Quantity{State: text.ValueRedacted, Value: 42}, "redacted"},
		{text.Quantity{Scale: text.QuantityScale("bogus")}, "invalid scale"},
	} {
		if got := tc.q.Label(); got != tc.want {
			t.Fatalf("%+v: got %q want %q", tc.q, got, tc.want)
		}
	}

	if text.ValuePresent.Label("") != `""` || text.ValuePresent.Label("0") != "0" || text.ValuePresent.Label("false") != "false" {
		t.Fatal("present values conflated with missing")
	}

	if !strings.Contains(text.ValueState("other").Label(""), "unknown value state") {
		t.Fatal("invalid value state hidden")
	}
}

func TestIdentifierAbbreviationPreservesIdentityAndCellBounds(t *testing.T) {
	id := text.Identifier{Value: "trace-a17c-92ef"}
	if got := id.Label(11); got != "trac...92ef" {
		t.Fatalf("wrong abbreviation: %q", got)
	}

	if id.Value != "trace-a17c-92ef" || id.Label(100) != id.Value {
		t.Fatal("identity was abbreviated in storage")
	}

	for _, value := range []string{"界界abc界", "e\u0301abcde\u0301", "\x1b[31msecret", ""} {
		for width := range 20 {
			label := (text.Identifier{Value: value}).Label(width)
			if ansi.StringWidth(label) > width || strings.Contains(label, "\x1b") {
				t.Fatalf("unsafe/oversize ID at width %d: %q", width, label)
			}
		}
	}
}

func TestMatchRangesUseOriginalRunesAndKeepGraphemes(t *testing.T) {
	normal := lipgloss.NewStyle()
	match := normal.Underline(true)
	source := "界e\u0301x\x1b"

	got := text.HighlightMatches(source, []text.MatchRange{{Start: 2, End: 3}, {Start: -3, End: 0}, {Start: 40, End: 99}, {Start: 3, End: 1}}, normal, match)
	if ansi.Strip(got) != `界éx\u001b` {
		t.Fatalf("highlighting changed escaped text: %q", got)
	}

	cells := uv.NewStyledString(got).Lines(ansi.GraphemeWidth)
	for _, cell := range cells[0] {
		if (cell.Style.Underline != uv.UnderlineNone) != (cell.Content == "é") {
			t.Fatalf("match lost its grapheme boundary: %+v", cell)
		}
	}
}

func TestMarkersPreserveSemanticDistinctionsInEveryTheme(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		palette := theme.New(mode)

		for _, emphasis := range []text.Emphasis{text.Bright, text.Normal, text.Dim} {
			for _, kind := range []text.ChangeKind{text.Unchanged, text.Added, text.Modified, text.Removed} {
				if !strings.Contains(ansi.Strip(kind.View(emphasis, palette)), string(kind)) {
					t.Fatal("change label lost")
				}
			}

			for _, state := range []text.ValueState{text.ValueMissing, text.ValueRedacted, text.ValueInvalid} {
				if !strings.Contains(ansi.Strip(state.View("secret", emphasis, palette)), string(state)) || strings.Contains(ansi.Strip(state.View("secret", emphasis, palette)), "secret") {
					t.Fatal("value-state label leaked hidden value")
				}
			}

			asc := text.SortIndicator{Direction: text.SortAscending, Priority: 2, Emphasis: emphasis}
			if !strings.Contains(ansi.Strip(asc.View(palette)), "ASC 2") {
				t.Fatal("sort priority lost")
			}

			if !strings.Contains(ansi.Strip((text.SortIndicator{}).View(palette)), "UNSORTED") {
				t.Fatal("unsorted implied an ordering")
			}
		}
	}

	for _, tc := range []struct {
		omission text.Omission
		want     string
	}{
		{text.Omission{Count: 24, KnownCount: true}, "+24 more"},
		{text.Omission{}, "more hidden"},
		{text.Omission{KnownCount: true}, ""},
		{text.Omission{Kind: text.SourceTruncated}, "source truncated (count unknown)"},
		{text.Omission{Kind: text.SourceTruncated, Count: 42, KnownCount: true}, "source truncated (42 omitted)"},
		{text.Omission{Count: -1, KnownCount: true}, "invalid omission count"},
	} {
		if got := tc.omission.Label(); got != tc.want {
			t.Fatalf("omission %q want %q", got, tc.want)
		}
	}
}
