package gitui

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/alecthomas/chroma/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func TestSyntaxLexerFailuresFallBackWithoutPartialTokens(t *testing.T) {
	for _, tokenize := range []func(string) (chroma.Iterator, error){
		func(string) (chroma.Iterator, error) { return nil, errors.New("lexer unavailable") },
		func(string) (chroma.Iterator, error) { panic("lexer failed") },
		func(string) (chroma.Iterator, error) {
			return func() chroma.Token { panic("iterator failed") }, nil
		},
	} {
		if tokens, ok := syntaxTokens("source", tokenize); ok || len(tokens) != 0 {
			t.Fatal("lexer failure returned partial syntax")
		}
	}
}

func TestDiffTabsKeepLexicalMeaningAndEscapeOnlyForDisplay(t *testing.T) {
	for _, language := range []string{"go", "python"} {
		hunk := Hunk{Lines: []DiffLine{{Kind: AddedLine, Text: "\treturn value"}}}

		lines := hunkSyntax(hunk, language, false)
		if lines[0].text != `\treturn value` || !tokenAt(lines[0], 2).InCategory(chroma.Keyword) {
			t.Fatalf("tab indentation corrupted %s keyword: %+v", language, lines[0])
		}
	}

	if got := diffLexerText(`return "\t"`); got != `return "\t"` {
		t.Fatalf("literal source escape became a tab: %q", got)
	}
}

func tokenAt(line syntaxLine, offset int) chroma.TokenType {
	for _, span := range line.spans {
		if span.end > offset {
			return span.kind
		}
	}

	return chroma.Text
}

func TestSyntaxRenderingIgnoresEquivalentTokenPartitions(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		palette := theme.New(mode)
		whole := syntaxLine{text: "    value", spans: []syntaxSpan{{4, chroma.Text}, {9, chroma.Name}}}
		partitioned := syntaxLine{text: whole.text, spans: []syntaxSpan{{1, chroma.TextWhitespace}, {4, chroma.Text}, {6, chroma.NameVariable}, {9, chroma.NameFunction}}}

		for _, kind := range []LineKind{ContextLine, AddedLine, DeletedLine} {
			edits := []editSpan{{6, 9}}
			if a, b := renderSyntaxLine(whole, kind, edits, palette), renderSyntaxLine(partitioned, kind, edits, palette); a != b {
				t.Fatalf("%s %s: equivalent syntax partitions changed ANSI output: %q != %q", mode, kind, a, b)
			}
		}
	}
}

func TestDiffSyntaxKeepsOldAndNewLexerStateSeparate(t *testing.T) {
	hunk := Hunk{Lines: []DiffLine{
		{Kind: ContextLine, Text: "/* comment"},
		{Kind: DeletedLine, Text: "*/"},
		{Kind: AddedLine, Text: "still a comment"},
		{Kind: ContextLine, Text: "const count = 12"},
		{Kind: AddedLine, Text: "*/"},
	}}

	old, current := hunkSyntax(hunk, "go", true), hunkSyntax(hunk, "go", false)
	if !tokenAt(old[3], 0).InCategory(chroma.Keyword) || !tokenAt(current[3], 0).InCategory(chroma.Comment) {
		t.Fatal("deleted comment terminator contaminated new-side lexical state")
	}

	independent := hunkSyntax(Hunk{Lines: []DiffLine{{Kind: AddedLine, Text: "const count = 12"}}}, "go", false)
	if !tokenAt(independent[0], 0).InCategory(chroma.Keyword) {
		t.Fatal("lexer state leaked across hunks")
	}
}

func TestDiffLanguageDetectionOverrideAndFallback(t *testing.T) {
	for _, tc := range []struct{ path, language string }{{"a.go", ""}, {"no-extension", "go"}} {
		language := diffLanguage(tc.path, tc.language)

		lines := hunkSyntax(Hunk{Lines: []DiffLine{{Kind: AddedLine, Text: `const name = "value"`}}}, language, false)
		if !tokenAt(lines[0], 0).InCategory(chroma.Keyword) || !tokenAt(lines[0], 13).InSubCategory(chroma.LiteralString) {
			t.Fatalf("%+v did not select Go highlighting", tc)
		}
	}

	for _, language := range []string{"text", "not-a-real-language", diffLanguage("unknown.extension-123", "")} {
		lines := hunkSyntax(Hunk{Lines: []DiffLine{{Kind: AddedLine, Text: "const name = 1"}}}, language, false)
		if len(lines[0].spans) != 0 || lines[0].text != "const name = 1" {
			t.Fatalf("%q did not preserve plain source", language)
		}
	}
}

func TestInlineEditsPreserveSeparateWordsAndGraphemes(t *testing.T) {
	before, after := "left red middle blue right", "left RED middle BLUE right"
	removed, added := inlineEdits(before, after)

	want := []editSpan{{5, 8}, {16, 20}}
	if !reflect.DeepEqual(removed, want) || !reflect.DeepEqual(added, want) {
		t.Fatalf("separate word edits: %v %v", removed, added)
	}

	removed, added = inlineEdits("e\u0301🙂x", "e\u0300🙂x")
	if !reflect.DeepEqual(removed, []editSpan{{0, 3}}) || !reflect.DeepEqual(added, []editSpan{{0, 3}}) {
		t.Fatalf("split a combining sequence: %v %v", removed, added)
	}

	removed, added = inlineEdits("same", "same")
	if len(removed) != 0 || len(added) != 0 {
		t.Fatal("identical lines emphasized")
	}

	removed, added = inlineEdits("prefix "+strings.Repeat("a", 10000)+" suffix", "prefix "+strings.Repeat("b", 10000)+" suffix")
	if !reflect.DeepEqual(removed, []editSpan{{7, 10007}}) || !reflect.DeepEqual(added, removed) {
		t.Fatal("long replacement did not use bounded changed-middle fallback")
	}
}

func TestDiffReplacementPairingStopsAtContext(t *testing.T) {
	hunk := Hunk{Lines: []DiffLine{{Kind: DeletedLine, Text: "old"}, {Kind: ContextLine, Text: "context"}, {Kind: AddedLine, Text: "new"}}}

	edits := hunkEdits(hunk, hunkSyntax(hunk, "text", false))
	for _, spans := range edits {
		if len(spans) != 0 {
			t.Fatal("paired unrelated changes across context")
		}
	}
}

func TestDiffNumberGutterExpandsOnlyForShiftedPositions(t *testing.T) {
	for _, tc := range []struct {
		name  string
		lines []DiffLine
		want  string
	}{
		{"aligned context", []DiffLine{{Kind: ContextLine, OldLine: 10, NewLine: 10, Text: "same"}}, "   10  same"},
		{"replacement", []DiffLine{{Kind: DeletedLine, OldLine: 10, Text: "old"}, {Kind: AddedLine, NewLine: 10, Text: "new"}}, "   10 -old\n   10 +new"},
		{"addition only", []DiffLine{{Kind: AddedLine, NewLine: 10, Text: "new"}}, "   10 +new"},
		{"deletion only", []DiffLine{{Kind: DeletedLine, OldLine: 10, Text: "old"}}, "   10 -old"},
		{"shifted replacement", []DiffLine{{Kind: DeletedLine, OldLine: 10, Text: "old"}, {Kind: AddedLine, NewLine: 20, Text: "new"}}, "   10       -old\n         20 +new"},
		{"insertion shifts context", []DiffLine{{Kind: AddedLine, NewLine: 10, Text: "new"}, {Kind: ContextLine, OldLine: 10, NewLine: 11, Text: "same"}}, "         10 +new\n   10    11  same"},
		{"deletion shifts context", []DiffLine{{Kind: DeletedLine, OldLine: 10, Text: "old"}, {Kind: ContextLine, OldLine: 11, NewLine: 10, Text: "same"}}, "   10       -old\n   11    10  same"},
		{"missing numbers", []DiffLine{{Kind: ContextLine, Text: "same"}}, "       same"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			view := ansi.Strip(DiffHunk(Hunk{Header: "hunk", Lines: tc.lines}, theme.New(theme.Dark)))
			if view != "hunk\n"+tc.want {
				t.Fatalf("got %q; want %q", view, "hunk\n"+tc.want)
			}
		})
	}
}

func TestDiffOutputPreservesSourceGuttersAndSafeControls(t *testing.T) {
	palette := theme.New(theme.Dark)
	hunk := Hunk{Header: "@@ -10,2 +10,2 @@", Lines: []DiffLine{
		{Kind: ContextLine, OldLine: 10, NewLine: 10, Text: "const ready = true"},
		{Kind: DeletedLine, OldLine: 11, Text: `return "old"`},
		{Kind: AddedLine, NewLine: 11, Text: `return "new"`, NoNewline: true},
	}}
	view := DiffHunk(hunk, palette, DiffOptions{Language: "go"})
	plain := ansi.Strip(view)

	want := "@@ -10,2 +10,2 @@\n   10  const ready = true\n   11 -return \"old\"\n   11 +return \"new\"\n\\ No newline at end of file"
	if plain != want {
		t.Fatalf("source changed:\n%q\nwant:\n%q", plain, want)
	}

	if !strings.Contains(view, diffLineStyle(DeletedLine, true, palette).Foreground(palette.Syntax.String).Render("old")) || !strings.Contains(view, diffLineStyle(AddedLine, true, palette).Foreground(palette.Syntax.String).Render("new")) {
		t.Fatal("inline edit surfaces lost syntax string colors")
	}

	unsafe := Hunk{Lines: []DiffLine{{Kind: AddedLine, Text: "\x1b[31m\tfoo\nbar\r\u202e"}}}

	plain = ansi.Strip(DiffHunk(unsafe, palette))
	if !strings.Contains(plain, `\u001b[31m\tfoo\nbar\r\u202e`) || strings.Count(plain, "\n") != 1 {
		t.Fatalf("terminal controls were executed, hidden or created rows: %q", plain)
	}
}

func TestDiffViewerLanguageAndNativeCustomizationSurviveResize(t *testing.T) {
	palette := theme.New(theme.Dark)
	diff := FileDiff{Path: "a.go", State: TextDiff, Hunks: []Hunk{{Lines: []DiffLine{{Kind: AddedLine, Text: "const limit = 12"}}}}}
	viewer := NewDiffViewer(diff, 80, 12, palette)
	viewer.Viewport.KeyMap.Down.SetKeys("n")

	view := viewer.View()
	if !strings.Contains(view, diffLineStyle(AddedLine, false, palette).Foreground(palette.Syntax.Keyword).Render("const")) {
		t.Fatal("file path did not enable syntax highlighting")
	}

	diff.Language = "text"
	viewer.SetDiff(diff)
	viewer.Resize(layout.Size{Width: 40, Height: 8})

	if strings.Contains(viewer.View(), diffLineStyle(AddedLine, false, palette).Foreground(palette.Syntax.Keyword).Render("const")) || fmt.Sprint(viewer.Viewport.KeyMap.Down.Keys()) != "[n]" {
		t.Fatal("plain override or native key customization lost")
	}
}

func TestDiffMonoRetainsSignsAndEmphasisWithoutColor(t *testing.T) {
	hunk := Hunk{Lines: []DiffLine{{Kind: DeletedLine, Text: "return 12"}, {Kind: AddedLine, Text: "return 34"}}}

	view := DiffHunk(hunk, theme.New(theme.Mono), DiffOptions{Language: "go"})
	if !strings.Contains(view, "\x1b[1;4") || !strings.Contains(ansi.Strip(view), "-return 12") || !strings.Contains(ansi.Strip(view), "+return 34") {
		t.Fatalf("monochrome edit lost non-color cues: %q", view)
	}

	if strings.Contains(view, "38;") || strings.Contains(view, "48;") {
		t.Fatal("monochrome diff emitted colors")
	}
}
