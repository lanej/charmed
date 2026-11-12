package gitui

import (
	"fmt"
	"strings"

	"github.com/lanej/charmed/theme"
)

func RepositorySummary(repo Repository, t theme.Theme) string {
	if repo.Unavailable {
		return portable(t.Heading().Render(clean(repo.Name))+" "+HeadIndicator(repo.Head, t)+"\n"+t.Subtle().Render("Status unavailable: "+clean(repo.Reason)), t)
	}

	counts := map[ChangeSection]int{}

	for _, f := range repo.Files {
		for _, section := range fileSections(f) {
			counts[section]++
		}
	}

	return portable(t.Heading().Render(clean(repo.Name))+" "+HeadIndicator(repo.Head, t)+"\n"+
		t.Text().Render(fmt.Sprintf("staged %d · unstaged %d · untracked %d · conflicts %d", counts[Staged], counts[Unstaged], counts[UntrackedFiles], counts[Conflicts]))+"\n"+
		t.Subtle().Render("upstream "+upstreamLabel(repo.Upstream))+" · "+DivergenceIndicator(repo.Divergence, t), t)
}

func upstreamLabel(upstream string) string {
	if upstream == "" {
		return "unavailable"
	}

	return clean(upstream)
}

func fileSections(f FileChange) []ChangeSection {
	if f.Conflict != "" || f.Index == Conflicted || f.Worktree == Conflicted {
		return []ChangeSection{Conflicts}
	}

	if f.Index == Untracked || f.Worktree == Untracked {
		return []ChangeSection{UntrackedFiles}
	}

	var sections []ChangeSection
	if f.Index != Unchanged {
		sections = append(sections, Staged)
	}

	if f.Worktree != Unchanged {
		sections = append(sections, Unstaged)
	}

	return sections
}

func FileChangeRow(f FileChange, section ChangeSection, t theme.Theme) string {
	kind := f.Worktree

	switch section {
	case Staged:
		kind = f.Index
	case Unstaged:
	case UntrackedFiles:
		kind = Untracked
	case Conflicts:
		kind = Conflicted
	default:
	}

	path := clean(f.Path)
	if f.OriginalPath != "" {
		path = clean(f.OriginalPath) + " → " + path
	}

	row := FileIndicator(kind, t) + " " + t.Text().Render(path)
	if f.Conflict != "" {
		row += " · " + t.Subtle().Render(clean(f.Conflict))
	}

	return portable(row, t)
}

func BranchRow(branch Branch, t theme.Theme) string {
	row := RefLabel(branch.Ref, t)
	if branch.Current {
		row += " [HEAD]"
	}

	row += " · upstream " + upstreamLabel(branch.Upstream) + " · " + DivergenceIndicator(branch.Divergence, t)
	if len(branch.WorktreePaths) > 0 {
		row += " · worktrees " + clean(strings.Join(branch.WorktreePaths, ", "))
	}

	return portable(row, t)
}

func CommitRow(commit Commit, refs []Ref, head Head, t theme.Theme) string {
	row := CommitIdentity(commit.ID, t)
	if head.CommitID == commit.ID {
		row += " " + HeadIndicator(head, t)
	}

	for _, ref := range refs {
		if ref.Target == commit.ID {
			row += " " + RefLabel(ref, t)
		}
	}

	row += " " + t.Text().Render(clean(commit.Subject))
	if commit.Author != "" {
		row += " · " + t.Subtle().Render(clean(commit.Author))
	}

	if !commit.Time.IsZero() {
		row += " · " + t.Subtle().Render(commit.Time.UTC().Format("2006-01-02 15:04"))
	}

	return portable(row, t)
}

func StashRow(stash Stash, t theme.Theme) string {
	return portable(t.Heading().Render(fmt.Sprintf("stash@{%d}", stash.Ordinal))+" "+clean(stash.Subject)+" · "+clean(stash.Branch)+" · "+CommitIdentity(stash.CommitID, t), t)
}

func WorktreeRow(w Worktree, t theme.Theme) string {
	return portable(t.Heading().Render(clean(w.Path))+" "+HeadIndicator(w.Head, t)+" "+WorktreeIndicators(w, t), t)
}
