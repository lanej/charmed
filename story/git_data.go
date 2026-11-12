package story

import (
	"time"

	"github.com/lanej/charmed/gitui"
)

func gitRepository() gitui.Repository {
	return gitui.Repository{Name: "charmed", Head: gitui.Head{Kind: gitui.Attached, Ref: "refs/heads/main", CommitID: "a31f2c7819"}, Upstream: "refs/remotes/origin/main", Divergence: gitui.Divergence{Known: true, Ahead: 2, Behind: 1}, Files: []gitui.FileChange{
		{ID: "status", Path: "gitui/status.go", Index: gitui.Added, Worktree: gitui.Modified},
		{ID: "guide", OriginalPath: "docs/git.md", Path: "docs/git-components.md", Index: gitui.Renamed},
		{ID: "notes", Path: "notes.txt", Worktree: gitui.Untracked},
		{ID: "theme", Path: "theme/theme.go", Index: gitui.Conflicted, Worktree: gitui.Conflicted, Conflict: "both modified"},
	}}
}

func gitHistory() gitui.History {
	return gitui.History{Head: gitRepository().Head, Refs: []gitui.Ref{
		{Name: "refs/heads/main", Kind: gitui.LocalBranch, Target: "a31f2c7819"},
		{Name: "refs/heads/search", Kind: gitui.LocalBranch, Target: "92bc1820ae"},
		{Name: "refs/remotes/origin/main", Kind: gitui.RemoteBranch, Target: "51dd20c991"},
		{Name: "refs/tags/v0.1", Kind: gitui.Tag, Target: "28c341aad0"},
	}, Commits: []gitui.Commit{
		{ID: "a31f2c7819", ParentIDs: []string{"51dd20c991", "92bc1820ae"}, Subject: "Merge search", Author: "Josh", Time: time.Date(2026, 11, 11, 12, 0, 0, 0, time.UTC)},
		{ID: "92bc1820ae", ParentIDs: []string{"870ada0912"}, Subject: "Add search input", Author: "Josh"},
		{ID: "870ada0912", ParentIDs: []string{"28c341aad0"}, Subject: "Filter results", Author: "Josh"},
		{ID: "51dd20c991", ParentIDs: []string{"28c341aad0"}, Subject: "Update guide", Author: "Josh"},
		{ID: "28c341aad0", Subject: "Initial components", Author: "Josh"},
	}}
}

func gitDiffs() []gitui.FileDiff {
	return []gitui.FileDiff{
		{FileID: "status", Section: gitui.Staged, Path: "gitui/status.go", State: gitui.TextDiff, Hunks: []gitui.Hunk{{Header: "@@ -0,0 +1,2 @@", Lines: []gitui.DiffLine{
			{Kind: gitui.AddedLine, NewLine: 1, Text: "package gitui"},
			{Kind: gitui.AddedLine, NewLine: 2, Text: "// Repository status presentation."},
		}}}},
		{FileID: "status", Section: gitui.Unstaged, Path: "gitui/status.go", State: gitui.TextDiff, Hunks: []gitui.Hunk{{Header: "@@ -1,2 +1,2 @@", Lines: []gitui.DiffLine{
			{Kind: gitui.ContextLine, OldLine: 1, NewLine: 1, Text: "package gitui"},
			{Kind: gitui.DeletedLine, OldLine: 2, Text: "// Repository status presentation."},
			{Kind: gitui.AddedLine, NewLine: 2, Text: "// Status and worktree presentation.", NoNewline: true},
		}}}},
		{FileID: "guide", Section: gitui.Staged, OriginalPath: "docs/git.md", Path: "docs/git-components.md", State: gitui.TextDiff},
		{FileID: "notes", Section: gitui.UntrackedFiles, Path: "notes.txt", State: gitui.UnavailableDiff, Reason: "untracked content not supplied"},
		{FileID: "theme", Section: gitui.Conflicts, Path: "theme/theme.go", State: gitui.UnavailableDiff, Reason: "unmerged index; resolve in the consuming application"},
	}
}

func gitWorktrees() []gitui.Worktree {
	repo := gitRepository()

	return []gitui.Worktree{
		{ID: "main", Path: "~/src/charmed", Head: repo.Head, Main: true, Current: true, StatusKnown: true, Status: repo},
		{ID: "search", Path: "~/src/charmed-search", Head: gitui.Head{Kind: gitui.Attached, Ref: "refs/heads/search", CommitID: "92bc1820ae"}, Locked: true, LockReason: "review in progress", StatusKnown: true, Status: gitui.Repository{Name: "charmed-search", Head: gitui.Head{Kind: gitui.Attached, Ref: "refs/heads/search", CommitID: "92bc1820ae"}}},
		{ID: "detached", Path: "~/src/charmed-old", Head: gitui.Head{Kind: gitui.Detached, CommitID: "28c341aad0"}, Prunable: true, PruneReason: "directory unavailable"},
	}
}

func gitBranches() []gitui.Branch {
	history := gitHistory()

	return []gitui.Branch{
		{Ref: history.Refs[0], Current: true, Upstream: "refs/remotes/origin/main", Divergence: gitui.Divergence{Known: true, Ahead: 2, Behind: 1}, WorktreePaths: []string{"~/src/charmed"}},
		{Ref: history.Refs[1], WorktreePaths: []string{"~/src/charmed-search"}},
		{Ref: history.Refs[2]},
	}
}

func gitStashes() []gitui.Stash {
	return []gitui.Stash{{ID: "stash-object-b72ce1", Ordinal: 0, Subject: "WIP: status presentation", Branch: "refs/heads/main", CommitID: "a31f2c7819", Diffs: gitDiffs()}}
}
