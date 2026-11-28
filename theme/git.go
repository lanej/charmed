package theme

import "image/color"

// GitColors gives Git-specific meanings to badges, summary counts and diff signs.
// Unset fields inherit the component's ordinary text or semantic status color.
type GitColors struct {
	Staged, Modified, Deleted, Renamed, TypeChanged, Untracked, Conflicted color.Color
	Branch, Ahead, Behind, Stashed                                         color.Color
	SummaryBranch, ChangedFiles, UntrackedFiles, AddedLines, RemovedLines  color.Color
	DiffAdd, DiffDelete                                                    color.Color
}

func sageGitColors() GitColors {
	return GitColors{
		Staged:         overrideColor(roleGitFileStaged, roleStatusSuccess),
		Modified:       overrideColor(roleGitFileModified, roleStatusWarning),
		Deleted:        overrideColor(roleGitFileDeleted, roleStatusError),
		Renamed:        overrideColor(roleGitFileRenamed, roleStatusWarning),
		TypeChanged:    overrideColor(roleGitFileTypechanged, roleStatusWarning),
		Untracked:      overrideColor(roleGitFileUntracked, roleStatusInfo),
		Conflicted:     overrideColor(roleGitFileConflicted, roleStatusError),
		Branch:         overrideColor(roleGitFileBranch, roleStatusInfo),
		Ahead:          overrideColor(roleGitFileAhead, roleText),
		Behind:         overrideColor(roleGitFileBehind, roleText),
		Stashed:        overrideColor(roleGitFileStashed, roleHeading),
		SummaryBranch:  overrideColor(roleGitSummaryBranch, roleTextMuted),
		ChangedFiles:   overrideColor(roleGitSummaryChangedFiles, roleText),
		UntrackedFiles: hex(roleGitSummaryUntrackedFiles),
		AddedLines:     overrideColor(roleGitSummaryAddedLines, roleStatusSuccess),
		RemovedLines:   overrideColor(roleGitSummaryRemovedLines, roleStatusError),
		DiffAdd:        overrideColor(roleDiffAdd, roleStatusSuccess),
		DiffDelete:     overrideColor(roleDiffDelete, roleStatusError),
	}
}
