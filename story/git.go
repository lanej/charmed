package story

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
)

func gitDefinitions() []Definition {
	catalog := []struct {
		id, title, level string
		composes         []string
	}{
		{"commit-id", "Commit identity", "Atoms", nil},
		{"ref", "Ref label", "Atoms", nil},
		{"head", "HEAD indicator", "Atoms", nil},
		{"file-state", "File change", "Atoms", nil},
		{"divergence", "Divergence", "Atoms", nil},
		{"diff-summary", "Diff summary", "Atoms", nil},
		{"worktree-state", "Worktree state", "Atoms", nil},
		{"staged-state", "Staged state", "Atoms", nil},
		{"unstaged-state", "Unstaged state", "Atoms", nil},
		{"untracked-state", "Untracked state", "Atoms", nil},
		{"conflict-state", "Conflict state", "Atoms", nil},
		{"staged", "Staged changes", "Molecules", []string{"git/staged-state", "git/file-row"}},
		{"unstaged", "Unstaged changes", "Molecules", []string{"git/unstaged-state", "git/file-row"}},
		{"untracked", "Untracked changes", "Molecules", []string{"git/untracked-state", "git/file-row"}},
		{"conflicts", "Conflicted changes", "Molecules", []string{"git/conflict-state", "git/file-row"}},
		{"summary", "Repository summary", "Molecules", []string{"git/head", "git/divergence", "git/file-state"}},
		{"file-row", "File change row", "Molecules", []string{"git/file-state"}},
		{"branch-row", "Branch row", "Molecules", []string{"git/ref", "git/divergence"}},
		{"commit-row", "Commit row", "Molecules", []string{"git/commit-id", "git/ref", "git/head"}},
		{"stash-row", "Stash row", "Molecules", []string{"git/commit-id"}},
		{"worktree-row", "Worktree row", "Molecules", []string{"git/head", "git/worktree-state"}},
		{"hunk", "Diff hunk", "Molecules", nil},
		{"status", "Repository status", "Organisms", []string{"git/summary", "git/staged", "git/unstaged", "git/untracked", "git/conflicts"}},
		{"branches", "Branch list", "Organisms", []string{"git/branch-row"}},
		{"history", "Branching commit history", "Organisms", []string{"git/commit-row"}},
		{"diff", "Diff viewer", "Organisms", []string{"git/hunk"}},
		{"changes", "Changes inspector", "Organisms", []string{"git/status", "git/diff"}},
		{"stashes", "Stash list", "Organisms", []string{"git/stash-row"}},
		{"worktrees", "Worktree list", "Organisms", []string{"git/worktree-row"}},
		{"browser", "Repository browser", "Screens", []string{"git/changes", "git/branches", "git/history", "git/stashes", "git/worktrees", "tui/screen"}},
	}

	definitions := make([]Definition, 0, len(catalog))
	for _, entry := range catalog {
		id := entry.id
		config := DefaultConfig()
		config.Columns, config.Rows = 120, 32

		switch id {
		case "staged-state", "unstaged-state", "untracked-state", "conflict-state":
			config.Columns, config.Rows = 40, 8
		case "staged", "unstaged", "untracked", "conflicts":
			config.Columns, config.Rows = 80, 10
		}

		definitions = append(definitions, Definition{ID: "git/" + id, Title: entry.level + "/Git/" + entry.title, Level: entry.level, Family: "git", Description: "Composable presentation of caller-supplied Git facts; no repository access or Git operations.", Source: "story/git.go", Symbol: "gitStory", Composes: entry.composes, Defaults: config, Controls: []string{"state"}, States: []string{"ready", "empty"}, Factory: func(e Environment) (runtime.Scene, error) { return gitStory(id, e) }})
	}

	return definitions
}

func gitStory(id string, e Environment) (runtime.Scene, error) {
	t := e.Config.Palette()
	repo, history, diffs := gitRepository(), gitHistory(), gitDiffs()
	branches, stashes, worktrees := gitBranches(), gitStashes(), gitWorktrees()

	if e.Config.State == "empty" {
		repo.Files = nil
		history.Commits = nil
		history.Refs = nil
		diffs = nil
		branches = nil
		stashes = nil
		worktrees = nil
	}

	w, h := e.Config.Columns, e.Config.Rows

	switch id {
	case "staged", "unstaged", "untracked", "conflicts":
		group, err := gitui.NewChangeGroup(gitui.ChangeSection(id), repo.Files, t)
		if err != nil {
			return nil, err
		}

		return static(func(layout.Size) string { return group.View() }), nil
	case "status":
		return gitui.NewStatus(repo, w, h, t)
	case "branches":
		return gitui.NewBranchList(branches, w, h, t)
	case "history":
		return gitui.NewCommitHistory(history, w, h, t)
	case "diff":
		diff := gitui.FileDiff{State: gitui.UnavailableDiff, Reason: "no supplied diff"}
		if len(diffs) > 0 {
			diff = diffs[1]
		}

		return gitui.NewDiffViewer(diff, w, h, t), nil
	case "changes":
		m, err := gitui.NewChangesInspector(repo, diffs, w, h, t)
		if err != nil {
			return nil, err
		}

		focusDiff := false

		return &scene{resize: m.Resize, view: func(layout.Size) string { return m.View() }, update: func(msg tea.Msg) tea.Cmd {
			if key, ok := msg.(tea.KeyPressMsg); ok && key.String() == "tab" {
				focusDiff = !focusDiff
				m.SetFocus(focusDiff)

				return nil
			}

			return m.Update(msg)
		}}, nil
	case "stashes":
		return gitui.NewStashList(stashes, w, h, t)
	case "worktrees":
		return gitui.NewWorktreeList(worktrees, w, h, t)
	case "browser":
		return gitBrowser(e)
	}

	return static(func(layout.Size) string {
		switch id {
		case "staged-state":
			return gitui.SectionIndicator(gitui.Staged, t)
		case "unstaged-state":
			return gitui.SectionIndicator(gitui.Unstaged, t)
		case "untracked-state":
			return gitui.SectionIndicator(gitui.UntrackedFiles, t)
		case "conflict-state":
			return gitui.SectionIndicator(gitui.Conflicts, t)
		case "commit-id":
			return gitui.CommitIdentity("a31f2c7819be822", t)
		case "ref":
			var labels []string
			for _, ref := range gitHistory().Refs {
				labels = append(labels, gitui.RefLabel(ref, t))
			}

			return strings.Join(labels, "\n\n")
		case "head":
			return gitui.HeadIndicator(repo.Head, t) + "\n\n" + gitui.HeadIndicator(gitui.Head{Kind: gitui.Detached, CommitID: "28c341aad0"}, t) + "\n\n" + gitui.HeadIndicator(gitui.Head{Kind: gitui.Unborn, Ref: "refs/heads/new"}, t)
		case "file-state":
			var labels []string
			for _, kind := range []gitui.ChangeKind{gitui.Added, gitui.Modified, gitui.Deleted, gitui.Renamed, gitui.Copied, gitui.TypeChanged, gitui.Untracked, gitui.Conflicted} {
				labels = append(labels, gitui.FileIndicator(kind, t))
			}

			return strings.Join(labels, "\n")
		case "divergence":
			return gitui.DivergenceIndicator(repo.Divergence, t) + "\n\n" + gitui.DivergenceIndicator(gitui.Divergence{}, t)
		case "diff-summary":
			return gitui.DiffSummary(gitui.DiffStat{State: gitui.TextDiff, Added: 12, Deleted: 3}, t) + "\n\n" + gitui.DiffSummary(gitui.DiffStat{State: gitui.BinaryDiff}, t) + "\n\n" + gitui.DiffSummary(gitui.DiffStat{State: gitui.UnavailableDiff}, t)
		case "worktree-state":
			return gitui.WorktreeIndicators(gitWorktrees()[1], t) + "\n\n" + gitui.WorktreeIndicators(gitWorktrees()[2], t)
		case "summary":
			return gitui.RepositorySummary(repo, t)
		case "file-row":
			return gitui.FileChangeRow(gitRepository().Files[0], gitui.Staged, t) + "\n\n" + gitui.FileChangeRow(gitRepository().Files[0], gitui.Unstaged, t)
		case "branch-row":
			return gitui.BranchRow(gitBranches()[0], t)
		case "commit-row":
			return gitui.CommitRow(gitHistory().Commits[0], gitHistory().Refs, repo.Head, t)
		case "stash-row":
			return gitui.StashRow(gitStashes()[0], t)
		case "worktree-row":
			return gitui.WorktreeRow(gitWorktrees()[1], t)
		case "hunk":
			return gitui.DiffHunk(gitDiffs()[1].Hunks[0], t, gitui.DiffOptions{Language: "go"})
		}

		return ""
	}), nil
}
