package gitui_test

import (
	"slices"
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func reviewFiles() []gitui.ReviewFile {
	return []gitui.ReviewFile{
		{ID: "staged-config", Diff: gitui.FileDiff{FileID: "config", Section: gitui.Staged, Path: "src/config.go", State: gitui.TextDiff, Hunks: []gitui.Hunk{
			{Header: "@@ -1 +1 @@", Lines: []gitui.DiffLine{{Kind: gitui.DeletedLine, OldLine: 1, Text: "before"}, {Kind: gitui.AddedLine, NewLine: 1, Text: "after"}}},
		}}, Stat: &gitui.DiffStat{State: gitui.TextDiff, Added: 1, Deleted: 1}},
		{ID: "unstaged-config", Diff: gitui.FileDiff{FileID: "config", Section: gitui.Unstaged, Path: "src/config.go", State: gitui.UnavailableDiff, Reason: "not captured"}},
		{ID: "binary", Diff: gitui.FileDiff{FileID: "asset", Path: "logo.png", State: gitui.BinaryDiff}},
	}
}

func TestDiffReviewMarksFoldAndContentRefresh(t *testing.T) {
	files := reviewFiles()

	m, err := gitui.NewFileDiffReview(files, layout.Size{Width: 120, Height: 24}, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	if !m.SetReviewed("staged-config", true) || !m.Collapsed("staged-config") {
		t.Fatal("reviewing did not fold the file")
	}

	m.SetCollapsed("staged-config", false)
	m.SetReviewed("staged-config", true)

	if !m.Reviewed("staged-config") || m.Collapsed("staged-config") {
		t.Fatal("reopening or repeating a mark lost the independent fold state")
	}

	files[0].Stat.Added = 7
	files[0].Diff.Path = "src/renamed.go"

	entries, err := gitui.FileReviewEntries(files)
	if err != nil {
		t.Fatal(err)
	}

	if err = m.SetEntries(entries); err != nil {
		t.Fatal(err)
	}

	if !m.Reviewed("staged-config") || m.Collapsed("staged-config") {
		t.Fatal("metadata-only update reset review/fold")
	}

	m.SetCollapsed("staged-config", true)

	files[0].Diff.Hunks[0].Lines[1].Text = "new revision"

	entries, err = gitui.FileReviewEntries(files)
	if err != nil {
		t.Fatal(err)
	}

	if err = m.SetEntries(entries); err != nil {
		t.Fatal(err)
	}

	if m.Reviewed("staged-config") || m.Collapsed("staged-config") {
		t.Fatal("changed diff was left reviewed or folded")
	}

	if !strings.Contains(ansi.Strip(m.View()), "new revision") {
		t.Fatal("refreshed content was not presented")
	}
}

func TestDiffReviewGroupSummaryAndKeyboard(t *testing.T) {
	m, err := gitui.NewFileDiffReview(reviewFiles(), layout.Size{Width: 120, Height: 24}, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	m.Select("staged-config")
	m.Update(tea.KeyPressMsg{Code: 'r'})

	if !slices.Equal(m.ReviewedIDs(), []string{"staged-config"}) {
		t.Fatal(m.ReviewedIDs())
	}

	group := m.Document.Sections()[0].ID
	m.SetCollapsed(group, true)

	view := ansi.Strip(m.View())
	if !strings.Contains(view, "2 files · 1 reviewed · stats unavailable") || strings.Contains(view, "not captured") {
		t.Fatal("collapsed subtree must be one summary with explicit unknown statistics", view)
	}

	m.SetCollapsed(group, false)

	if !m.Collapsed("staged-config") || !strings.Contains(ansi.Strip(m.View()), "Binary content changed") {
		t.Fatal("expansion lost child folds or binary rendering")
	}

	if m.SetReviewed(group, true) || m.SetReviewed("absent", true) {
		t.Fatal("reviewing groups or missing files must fail")
	}
}

func TestDiffReviewCopiesDataAndRejectsInvalidUpdates(t *testing.T) {
	files := reviewFiles()
	files[0].Reviewed = true

	m, err := gitui.NewFileDiffReview(files, layout.Size{Width: 120, Height: 24}, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	files[0].Diff.Hunks[0].Lines[1].Text = "caller mutation"

	m.SetCollapsed("staged-config", false)

	if strings.Contains(ansi.Strip(m.View()), "caller mutation") {
		t.Fatal("retained diff aliases caller data")
	}

	if err := m.SetEntries([]gitui.ReviewEntry{{ID: "same"}, {ID: "same"}}); err == nil {
		t.Fatal("duplicate identities accepted")
	}

	if !m.Reviewed("staged-config") {
		t.Fatal("failed update changed review state")
	}

	if _, err := gitui.FileReviewEntries([]gitui.ReviewFile{{ID: "x"}}); err == nil {
		t.Fatal("empty helper path accepted")
	}

	if _, err := gitui.FileReviewEntries(append(files, files[0])); err == nil {
		t.Fatal("duplicate helper identity accepted")
	}
}

func TestDiffReviewFlatCustomLabelsAndExplicitASCII(t *testing.T) {
	palette := theme.New(theme.Dark)
	palette.ASCII = true
	diff := gitui.FileDiff{State: gitui.TextDiff}

	m, err := gitui.NewDiffReview([]gitui.ReviewEntry{{ID: "query", Label: "Generated query", Diff: &diff, Reviewed: true}}, layout.Size{Width: 80, Height: 12}, palette)
	if err != nil {
		t.Fatal(err)
	}

	view := ansi.Strip(m.View())
	if !strings.Contains(view, "Generated query") || !strings.Contains(view, "reviewed") {
		t.Fatal("custom labels or initial state were not retained")
	}

	for _, r := range view {
		if r > 127 {
			t.Fatal("explicit ASCII presentation contains Unicode", r)
		}
	}

	if !m.SetReviewed("query", false) || m.Collapsed("query") {
		t.Fatal("clearing reviewed must open the body")
	}
}

func TestDiffReviewTargetsSelectionWhenViewportCannotAlignItsHeader(t *testing.T) {
	files := reviewFiles()

	m, err := gitui.NewFileDiffReview(files, layout.Size{Width: 120, Height: 40}, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	m.Select("binary")
	m.SetFocus(true)
	m.Update(tea.KeyPressMsg{Code: 'r'})

	if !m.Reviewed("binary") || m.Reviewed("staged-config") {
		t.Fatal("review key acted on viewport top instead of selected item")
	}

	m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})

	if m.Collapsed("binary") || !m.Reviewed("binary") {
		t.Fatal("fold key acted on viewport top instead of selected item")
	}
}

func TestDiffReviewPreservesRenameAndAcceptsBinaryContentRevisions(t *testing.T) {
	files := reviewFiles()
	files[0].Diff.OriginalPath = "legacy/route.go"
	files[2].ContentRevision = "blob-v1"

	m, err := gitui.NewFileDiffReview(files, layout.Size{Width: 120, Height: 30}, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(ansi.Strip(m.View()), "legacy/route.go → src/config.go") {
		t.Fatal("rename evidence was discarded")
	}

	m.SetReviewed("binary", true)

	files[2].ContentRevision = "blob-v2"

	entries, err := gitui.FileReviewEntries(files)
	if err != nil {
		t.Fatal(err)
	}

	if err = m.SetEntries(entries); err != nil {
		t.Fatal(err)
	}

	if m.Reviewed("binary") || m.Collapsed("binary") {
		t.Fatal("a changed content revision did not invalidate binary review")
	}
}
