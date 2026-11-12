package gitui

import (
	"reflect"
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/interaction"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func press(value string) tea.KeyPressMsg {
	switch value {
	case "enter":
		return tea.KeyPressMsg{Code: tea.KeyEnter}
	case "esc":
		return tea.KeyPressMsg{Code: tea.KeyEscape}
	default:
		return tea.KeyPressMsg{Code: rune(value[0]), Text: value}
	}
}

func TestGraphTopology(t *testing.T) {
	palette := theme.New(theme.Mono)

	for _, tc := range []struct {
		name     string
		history  History
		prefixes []string
		joins    []string
	}{
		{"linear", History{Commits: []Commit{{ID: "c", ParentIDs: []string{"b"}}, {ID: "b", ParentIDs: []string{"a"}}, {ID: "a"}}}, []string{"* ", "* ", "* "}, nil},
		{"merge", History{Commits: []Commit{{ID: "merge", ParentIDs: []string{"main", "topic"}}, {ID: "topic", ParentIDs: []string{"base"}}, {ID: "main", ParentIDs: []string{"base"}}, {ID: "base"}}}, []string{"* ", "| * ", "* | ", "  * "}, []string{"+-+", "+-+"}},
		{"octopus", History{Commits: []Commit{{ID: "merge", ParentIDs: []string{"a", "b", "c"}}, {ID: "a"}, {ID: "b"}, {ID: "c"}}}, []string{"* ", "* | | ", "  * | ", "    * "}, nil},
		{"disconnected", History{Commits: []Commit{{ID: "a"}, {ID: "b"}}}, []string{"* ", "* "}, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			rows, _, err := graph(tc.history, palette)
			if err != nil {
				t.Fatal(err)
			}

			var joins []string

			for i, row := range rows {
				if got := ansi.Strip(row.prefix); got != tc.prefixes[i] {
					t.Fatalf("node %d: got %q want %q", i, got, tc.prefixes[i])
				}

				for _, line := range row.after {
					if strings.Contains(line, "-") {
						joins = append(joins, strings.TrimSpace(ansi.Strip(line)))
					}
				}
			}

			if tc.joins != nil && !reflect.DeepEqual(joins, tc.joins) {
				t.Fatalf("joins %q want %q", joins, tc.joins)
			}
		})
	}
}

func TestHistoryRejectsInvalidAndPreservesSelection(t *testing.T) {
	m, err := NewCommitHistory(History{Commits: []Commit{{ID: "a"}, {ID: "b"}}}, 80, 16, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	m.Select("b")

	for _, history := range []History{
		{Commits: []Commit{{ID: "a"}, {ID: "a"}}},
		{Commits: []Commit{{ID: "a", ParentIDs: []string{"a"}}}},
		{Commits: []Commit{{ID: "a"}, {ID: "b", ParentIDs: []string{"a"}}}},
		{Commits: []Commit{{ID: "a", ParentIDs: []string{"missing"}}}},
		{Commits: []Commit{{ID: "a", ParentIDs: []string{"b", "b"}}, {ID: "b"}}},
	} {
		if err := m.SetHistory(history); err == nil {
			t.Fatal("invalid history accepted")
		}

		if m.SelectedID() != "b" {
			t.Fatal("invalid replacement mutated selection")
		}
	}

	history := History{Commits: []Commit{{ID: "b", ParentIDs: []string{"outside"}}}, Truncated: true}
	if err := m.SetHistory(history); err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(ansi.Strip(m.View()), "history continues") {
		t.Fatal("missing history boundary")
	}

	m.Resize(layout.Size{Width: 12, Height: 8})

	if !strings.Contains(ansi.Strip(m.View()), "More width") {
		t.Fatal("narrow graph silently clipped")
	}

	if m.SelectedID() != "b" {
		t.Fatal("resize lost selection")
	}
}

func TestListRefreshCopyFilterAndSelection(t *testing.T) {
	items := []Item{{ID: "one", Text: "one", After: []string{"continuation"}}, {ID: "two", Text: "two"}, {ID: "three", Text: "three"}}

	m, err := NewList("test", items, 50, 10, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	items[0].After[0] = "mutated"

	if strings.Contains(m.View(), "mutated") {
		t.Fatal("shared continuation slice")
	}

	m.Select("two")

	cmd := m.Update(press("enter"))
	if cmd == nil {
		t.Fatal("no selection command")
	}

	if got := cmd(); got != (interaction.Selected{Component: "test", ID: "two"}) {
		t.Fatalf("selection %#v", got)
	}

	if err := m.SetItems([]Item{items[2], items[1], items[0]}); err != nil {
		t.Fatal(err)
	}

	if m.SelectedID() != "two" {
		t.Fatal("refresh lost identity")
	}

	if err := m.SetItems([]Item{items[2], items[0]}); err != nil {
		t.Fatal(err)
	}

	if m.SelectedID() != "one" {
		t.Fatal("removed selection did not choose nearest item")
	}

	m.Update(press("/"))
	m.Update(press("z"))
	m.Update(press("enter"))

	if !strings.Contains(m.View(), "No matching items") || m.Update(press("enter")) != nil {
		t.Fatal("empty filter emitted selection")
	}

	m.Update(press("esc"))
	m.SetFocused(false)
	before := m.SelectedID()
	m.Update(press("j"))

	if m.SelectedID() != before {
		t.Fatal("unfocused list consumed key")
	}
}

func TestStatusAndDiffInspector(t *testing.T) {
	repo := Repository{Name: "test", Head: Head{Kind: Unborn, Ref: "refs/heads/main"}, Files: []FileChange{
		{ID: "both", Path: "界.go", Index: Added, Worktree: Modified},
		{ID: "conflict", Path: "conflict.go", Index: Conflicted, Worktree: Modified, Conflict: "both modified"},
		{ID: "untracked", Path: "new.go", Worktree: Untracked},
	}}
	diffs := []FileDiff{
		{FileID: "both", Section: Staged, Path: "界.go", State: TextDiff, Hunks: []Hunk{{Header: "staged", Lines: []DiffLine{{Kind: AddedLine, NewLine: 1, Text: "original", NoNewline: true}}}}},
		{FileID: "both", Section: Unstaged, Path: "界.go", State: BinaryDiff},
	}

	m, err := NewChangesInspector(repo, diffs, 120, 28, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	diffs[0].Hunks[0].Lines[0].Text = "mutated"

	if err := m.SetDiffs([]FileDiff{{FileID: "x"}, {FileID: "x"}}); err == nil {
		t.Fatal("duplicate diffs accepted")
	}

	view := ansi.Strip(m.View())
	for _, want := range []string{"staged 1", "unstaged 1", "untracked 1", "conflicts 1", "original", "No newline"} {
		if !strings.Contains(view, want) {
			t.Fatalf("missing %q in\n%s", want, view)
		}
	}

	m.Update(press("j"))

	id, section := m.Status.SelectedChange()
	if id != "both" || section != Unstaged {
		t.Fatalf("selection %s %s", id, section)
	}

	if !strings.Contains(m.View(), "Binary content changed") {
		t.Fatal("selection did not update diff")
	}

	cmd := m.Update(press("enter"))
	if cmd == nil || cmd() != (interaction.Selected{Component: "git/status/unstaged", ID: "both"}) {
		t.Fatal("section identity lost")
	}

	m.SetFocus(true)
	m.Update(press("j"))

	if got, _ := m.Status.SelectedChange(); got != "both" {
		t.Fatal("diff focus moved status")
	}
}

func TestWorktreesRetainUnknownAndReasons(t *testing.T) {
	palette := theme.New(theme.Dark)

	view := ansi.Strip(WorktreeRow(Worktree{ID: "w", Path: "/src/w", Head: Head{Kind: Detached, CommitID: "1234567890"}, Locked: true, LockReason: "review", Prunable: true, PruneReason: "missing"}, palette))
	for _, want := range []string{"detached", "locked: review", "prunable: missing", "status unavailable"} {
		if !strings.Contains(view, want) {
			t.Fatalf("missing %q: %s", want, view)
		}
	}

	if strings.Contains(view, "clean") {
		t.Fatal("unknown status displayed clean")
	}

	if strings.Contains(ansi.Strip(DivergenceIndicator(Divergence{}, palette)), "ahead 0") {
		t.Fatal("unknown comparison displayed zero")
	}

	if got := clean("a\x1b[31mb\n\tc"); got != "ab  c" {
		t.Fatalf("unsafe text: %q", got)
	}
}

func TestCrossingsAreNotJunctions(t *testing.T) {
	history := History{Commits: []Commit{
		{ID: "A", ParentIDs: []string{"B", "C", "D"}},
		{ID: "B", ParentIDs: []string{"D"}},
		{ID: "C"},
		{ID: "D"},
	}}

	rows, _, err := graph(history, theme.New(theme.Mono))
	if err != nil {
		t.Fatal(err)
	}
	// B joins D across the unrelated C lane. ':' must preserve C's vertical path.
	got := ansi.Strip(strings.Join(rows[1].after, "\n"))
	if !strings.Contains(got, "+-:-+") {
		t.Fatalf("crossing became a junction:\n%s", got)
	}

	if ansi.Strip(rows[2].prefix) != "  * | " || ansi.Strip(rows[3].prefix) != "    * " {
		t.Fatalf("pending parents moved: %q %q", rows[2].prefix, rows[3].prefix)
	}

	history.Commits = []Commit{{ID: "A", ParentIDs: []string{"B", "D"}}, {ID: "B", ParentIDs: []string{"C", "D"}}, {ID: "C", ParentIDs: []string{"D"}}, {ID: "D"}}
	if _, _, err := graph(history, theme.New(theme.Mono)); err != nil {
		t.Fatal(err)
	}
}

func TestSelectionCommandsSnapshotIdentity(t *testing.T) {
	m, err := NewStatus(Repository{Files: []FileChange{{ID: "a", Path: "a", Index: Added}}}, 80, 12, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	cmd := m.Update(press("enter"))
	if err := m.SetRepository(Repository{}); err != nil {
		t.Fatal(err)
	}

	m.List.Component = "changed"
	if got := cmd(); got != (interaction.Selected{Component: "git/status/staged", ID: "a"}) {
		t.Fatalf("selection read mutable state: %#v", got)
	}
}

func TestNoMatchClearsInspectorAndASCIIIsPortable(t *testing.T) {
	m, err := NewChangesInspector(Repository{Files: []FileChange{{ID: "a", Path: "a", Index: Added}}}, nil, 100, 24, theme.New(theme.Mono))
	if err != nil {
		t.Fatal(err)
	}

	m.Update(press("/"))
	m.Update(press("z"))

	if id, _ := m.Status.SelectedChange(); id != "" {
		t.Fatal("hidden selection retained")
	}

	if m.Status.List.SelectedID() != "" {
		t.Fatal("list retained hidden selection")
	}

	palette := theme.New(theme.Mono)

	values := []string{HeadIndicator(Head{Kind: Attached, Ref: "main"}, palette), DivergenceIndicator(Divergence{Known: true}, palette), FileChangeRow(FileChange{Path: "new", OriginalPath: "old", Index: Renamed}, Staged, palette), RepositorySummary(Repository{}, palette), WorktreeRow(Worktree{Path: "/src", Head: Head{Kind: Unborn}}, palette)}
	for _, value := range values {
		for _, r := range ansi.Strip(value) {
			if r > 127 {
				t.Fatalf("ASCII API leaked %q", value)
			}
		}
	}

	view := RepositorySummary(Repository{Unavailable: true, Reason: "not supplied"}, palette)
	if strings.Contains(view, "staged 0") || !strings.Contains(view, "unavailable") {
		t.Fatalf("unknown status became zero counts: %s", view)
	}

	if _, err := NewStatus(Repository{Files: []FileChange{{Index: Added}}}, 80, 12, palette); err == nil {
		t.Fatal("empty file ID accepted")
	}
}

func TestTruncatedOctopusBoundaryPathsAndScroll(t *testing.T) {
	history := History{Commits: []Commit{{ID: "A", ParentIDs: []string{"B", "C", "D", "E", "F"}}}, Truncated: true}

	rows, _, err := graph(history, theme.New(theme.Mono))
	if err != nil {
		t.Fatal(err)
	}

	boundary := rows[0].after[len(rows[0].after)-5:]
	if got := ansi.Strip(boundary[0]); !strings.HasPrefix(got, "~ | | | | B") {
		t.Fatalf("boundary broke pending lanes: %q", got)
	}

	if got := ansi.Strip(boundary[4]); !strings.HasPrefix(got, "        ~ F") {
		t.Fatalf("last boundary misplaced: %q", got)
	}

	m, err := NewCommitHistory(history, 80, 8, theme.New(theme.Mono))
	if err != nil {
		t.Fatal(err)
	}

	for range 30 {
		m.Update(tea.KeyPressMsg{Code: 'e', Mod: tea.ModCtrl})
	}

	if !strings.Contains(m.View(), "F (history continues)") {
		t.Fatal("last continuation is unreachable")
	}
}

func TestPasteAndDirectOrganismASCII(t *testing.T) {
	palette := theme.New(theme.Mono)

	m, err := NewBranchList([]Branch{{Ref: Ref{Name: "refs/heads/main", Kind: LocalBranch}}}, 30, 12, palette)
	if err != nil {
		t.Fatal(err)
	}

	m.Update(press("/"))
	m.Update(tea.PasteMsg{Content: "no such branch"})

	if m.SelectedID() != "" || !strings.Contains(m.View(), "No matching items") {
		t.Fatal("paste did not filter")
	}

	history, err := NewCommitHistory(History{Commits: []Commit{{ID: "long-commit-ID", Subject: "a long subject that must truncate"}}}, 30, 8, palette)
	if err != nil {
		t.Fatal(err)
	}

	diff := NewDiffViewer(FileDiff{Path: "new", OriginalPath: "old", State: TextDiff}, 30, 8, palette)
	for _, view := range []string{m.View(), history.View(), diff.View()} {
		for _, r := range ansi.Strip(view) {
			if r > 127 {
				t.Fatalf("organism leaked Unicode: %q", view)
			}
		}
	}
}
