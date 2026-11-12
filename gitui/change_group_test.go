package gitui

import (
	"reflect"
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/theme"
)

func TestChangeGroupMolecules(t *testing.T) {
	files := []FileChange{
		{ID: "both", Path: "mixed.go", Index: Added, Worktree: Modified},
		{ID: "rename", OriginalPath: "old.go", Path: "new.go", Index: Renamed},
		{ID: "new", Path: "notes.txt", Worktree: Untracked},
		{ID: "conflict", Path: "merge.go", Index: Conflicted, Worktree: Modified, Conflict: "both modified"},
		{ID: "clean", Path: "clean.go"},
	}

	for _, tc := range []struct {
		section     ChangeSection
		constructor func([]FileChange, theme.Theme) (ChangeGroup, error)
		ids         []string
	}{
		{Staged, NewStagedChanges, []string{"both", "rename"}},
		{Unstaged, NewUnstagedChanges, []string{"both"}},
		{UntrackedFiles, NewUntrackedChanges, []string{"new"}},
		{Conflicts, NewConflictChanges, []string{"conflict"}},
	} {
		t.Run(string(tc.section), func(t *testing.T) {
			group, err := tc.constructor(files, theme.New(theme.Mono))
			if err != nil {
				t.Fatal(err)
			}

			if group.Section() != tc.section || group.Count() != len(tc.ids) {
				t.Fatal("incorrect section or count")
			}

			var ids []string
			for _, item := range group.Items() {
				ids = append(ids, item.ID)
			}

			if !reflect.DeepEqual(ids, tc.ids) {
				t.Fatalf("got %v want %v", ids, tc.ids)
			}

			view := ansi.Strip(group.View())
			if !strings.Contains(view, "["+string(tc.section)+"]") {
				t.Fatal("section atom missing")
			}

			if tc.section == Conflicts && !strings.Contains(view, "both modified") {
				t.Fatal("conflict reason lost")
			}

			if tc.section == Staged && !strings.Contains(view, "old.go > new.go") {
				t.Fatal("rename paths lost")
			}

			for _, r := range view {
				if r > 127 {
					t.Fatal("ASCII molecule leaked Unicode")
				}
			}

			empty, err := tc.constructor(nil, theme.New(theme.Dark))
			if err != nil {
				t.Fatal(err)
			}

			if empty.Count() != 0 || !strings.Contains(empty.View(), "No files") {
				t.Fatal("empty molecule lost its label")
			}
		})
	}
}

func TestChangeGroupsCopyAndValidateInputs(t *testing.T) {
	files := []FileChange{{ID: "f", Path: "original.go", Index: Added}}

	group, err := NewStagedChanges(files, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	files[0].Path = "mutated.go"
	copied := group.Files()
	copied[0].Path = "also-mutated.go"
	items := group.Items()
	items[0].Text = "mutated row"

	if view := group.View(); !strings.Contains(view, "original.go") || strings.Contains(view, "mutated") {
		t.Fatal("molecule shared caller data")
	}

	for _, invalid := range [][]FileChange{{{Path: "missing.go", Index: Added}}, {{ID: "same", Index: Added}, {ID: "same", Worktree: Modified}}} {
		if _, err := NewStagedChanges(invalid, theme.New(theme.Dark)); err == nil {
			t.Fatal("invalid file identity accepted")
		}
	}

	if _, err := NewChangeGroup(ChangeSection("invalid"), files, theme.New(theme.Dark)); err == nil {
		t.Fatal("unknown section accepted")
	}
}
