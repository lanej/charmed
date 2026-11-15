package gitui_test

import (
	"image/color"
	"strings"
	"testing"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func TestDiffGuttersKeepSideMeaningAndUseMeasuredWidths(t *testing.T) {
	hunk := gitui.Hunk{Header: "@@ -99,2 +9,2 @@", Lines: []gitui.DiffLine{
		{Kind: gitui.ContextLine, OldLine: 99, NewLine: 9, Text: "context"},
		{Kind: gitui.DeletedLine, OldLine: 100, Text: "before"},
		{Kind: gitui.AddedLine, NewLine: 10, Text: "after", NoNewline: true},
	}}
	palette := theme.New(theme.Dark)

	single := ansi.Strip(gitui.DiffHunk(hunk, palette))
	for _, want := range []string{"  9  context", "100 -before", " 10 +after", "\\ No newline at end of file"} {
		if !strings.Contains(single, want) {
			t.Fatalf("single gutter lost %q: %s", want, single)
		}
	}

	dual := ansi.Strip(gitui.DiffHunk(hunk, palette, gitui.DiffOptions{Gutter: gitui.DualGutter}))
	if !strings.Contains(dual, " 99  9  context") || !strings.Contains(dual, "100    -before") || !strings.Contains(dual, "    10 +after") {
		t.Fatal("dual gutter lost old/new line identities", dual)
	}

	if got := ansi.Strip(gitui.DiffHunk(hunk, palette, gitui.DiffOptions{Gutter: "invalid"})); got != single {
		t.Fatal("unknown gutter did not use safe single-column fallback")
	}
}

func TestDiffSyntaxKeepsOldAndNewFragmentsIndependent(t *testing.T) {
	palette := theme.New(theme.Dark)
	hunk := gitui.Hunk{Header: "@@ -1,3 +1,3 @@", Lines: []gitui.DiffLine{
		{Kind: gitui.DeletedLine, OldLine: 1, Text: "/* old comment"},
		{Kind: gitui.AddedLine, NewLine: 1, Text: "func ready() bool {"},
		{Kind: gitui.ContextLine, OldLine: 2, NewLine: 2, Text: "return true"},
		{Kind: gitui.DeletedLine, OldLine: 3, Text: "*/"},
		{Kind: gitui.AddedLine, NewLine: 3, Text: "}"},
	}}

	view := normalizedDiffANSI(gitui.DiffHunk(hunk, palette, gitui.DiffOptions{Language: "go"}))
	for _, want := range []string{
		palette.Text().Foreground(palette.Syntax.Keyword).Render("return"),
		palette.Text().Foreground(palette.Syntax.Comment).Render("/* old comment"),
		palette.Text().Foreground(palette.Danger).Render("1 -"),
		palette.Text().Foreground(palette.Success).Render("1 +"),
	} {
		if !containsDiffANSI(view, want) {
			t.Fatalf("missing independent syntax or diff-sign color %q in %q", want, view)
		}
	}

	plain := normalizedDiffANSI(gitui.DiffHunk(hunk, palette, gitui.DiffOptions{Language: "go", DisableSyntax: true}))
	if ansi.Strip(plain) != ansi.Strip(view) || containsDiffANSI(plain, palette.Text().Foreground(palette.Syntax.Keyword).Render("return")) {
		t.Fatal("disabling syntax changed source or retained lexical colors")
	}
}

func TestDiffReviewInfersLanguageAndRetainsStateAcrossPresentationChanges(t *testing.T) {
	palette := theme.New(theme.Dark)
	files := reviewFiles()
	files[0].Diff.Hunks[0].Lines[1].Text = "return 42"

	m, err := gitui.NewFileDiffReview(files, layout.Size{Width: 120, Height: 24}, palette)
	if err != nil {
		t.Fatal(err)
	}

	if !containsDiffANSI(m.View(), palette.Text().Foreground(palette.Syntax.Keyword).Render("return")) {
		t.Fatal("Go filename did not select syntax highlighting")
	}

	added, deleted := 0, 0

	for _, line := range uv.NewStyledString(m.Outline.View()).Lines(ansi.GraphemeWidth) {
		for _, cell := range line {
			switch cell.Content {
			case "+":
				added++

				if !equalDiffColor(cell.Style.Fg, palette.Success) {
					t.Fatal("outline styling overwrote the additions color")
				}
			case "-":
				deleted++

				if !equalDiffColor(cell.Style.Fg, palette.Danger) {
					t.Fatal("outline styling overwrote the deletions color")
				}
			}
		}
	}

	if added == 0 || deleted == 0 {
		t.Fatal("tree statistics were truncated before their colored signs")
	}

	m.Select("staged-config")
	m.SetReviewed("staged-config", true)
	m.SetCollapsed("staged-config", false)
	m.SetDiffOptions(gitui.DiffOptions{Gutter: gitui.DualGutter, DisableSyntax: true})

	if !m.Reviewed("staged-config") || m.Collapsed("staged-config") || m.SelectedID() != "staged-config" {
		t.Fatal("presentation change reset review, fold or selection")
	}

	if containsDiffANSI(m.View(), palette.Text().Foreground(palette.Syntax.Keyword).Render("return")) {
		t.Fatal("presentation setter did not refresh the cached syntax")
	}

	stat := gitui.DiffSummary(gitui.DiffStat{State: gitui.TextDiff, Added: 17, Deleted: 3}, palette)
	if ansi.Strip(stat) != "+17 -3" || !strings.Contains(stat, palette.Text().Foreground(palette.Success).Render("+17")) || !strings.Contains(stat, palette.Text().Foreground(palette.Danger).Render("-3")) {
		t.Fatal("diff statistics lost exact counts or separate semantic colors")
	}
}

// Both reset encodings have the same terminal meaning.
func normalizedDiffANSI(value string) string { return strings.ReplaceAll(value, "\x1b[0m", "\x1b[m") }

func equalDiffColor(actual, expected color.Color) bool {
	if actual == nil || expected == nil {
		return actual == nil && expected == nil
	}

	r, g, b, a := actual.RGBA()

	er, eg, eb, ea := expected.RGBA()

	return r == er && g == eg && b == eb && a == ea
}

func TestDiffSourceControlsRemainVisibleData(t *testing.T) {
	source := "\treturn \x1b[31mvalue\nnext"
	hunk := gitui.Hunk{Lines: []gitui.DiffLine{{Kind: gitui.AddedLine, NewLine: 1, Text: source}}}

	view := gitui.DiffHunk(hunk, theme.New(theme.Dark), gitui.DiffOptions{Language: "go"})
	if !strings.Contains(ansi.Strip(view), `\treturn \u001b[31mvalue\nnext`) || strings.Count(ansi.Strip(view), "\n") != 1 {
		t.Fatal("source controls were hidden or escaped the source row", view)
	}

	palette := theme.New(theme.Dark)
	if !containsDiffANSI(view, palette.Text().Foreground(palette.Syntax.Keyword).Render("return")) {
		t.Fatal("escaping the source tab broke keyword recognition", view)
	}
}

// Compare displayed content and foregrounds independently of diff backgrounds.
func containsDiffANSI(view, want string) bool {
	expected := uv.NewStyledString(want).Lines(ansi.GraphemeWidth)
	if len(expected) != 1 || len(expected[0]) == 0 {
		return false
	}

	for _, line := range uv.NewStyledString(view).Lines(ansi.GraphemeWidth) {
		for i := 0; i+len(expected[0]) <= len(line); i++ {
			match := true

			for j, cell := range expected[0] {
				if line[i+j].Content != cell.Content || !equalDiffColor(line[i+j].Style.Fg, cell.Style.Fg) {
					match = false
					break
				}
			}

			if match {
				return true
			}
		}
	}

	return false
}
