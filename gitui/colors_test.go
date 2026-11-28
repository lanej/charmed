package gitui_test

import (
	"image/color"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"
)

func TestGitTokenColorsInNativeRows(t *testing.T) {
	palette := theme.New(theme.Dark)
	palette.Git.Renamed = lipgloss.Color("#c6a8de")
	palette.Git.AddedLines = lipgloss.Color("#a3d2b8")
	palette.Git.DiffAdd = lipgloss.Color("#cfedd0")

	repo := gitui.Repository{Files: []gitui.FileChange{{Path: "new.go", Worktree: gitui.Untracked}}}
	hunk := gitui.Hunk{Lines: []gitui.DiffLine{{Kind: gitui.AddedLine, NewLine: 1, Text: "added"}}}

	for _, test := range []struct {
		view, label string
		color       color.Color
	}{
		{gitui.FileIndicator(gitui.Renamed, palette), "renamed", palette.Git.Renamed},
		{gitui.FileIndicator(gitui.TypeChanged, palette), "type changed", palette.Git.TypeChanged},
		{gitui.FileIndicator(gitui.Untracked, palette), "untracked", palette.Git.Untracked},
		{gitui.SectionIndicator(gitui.Staged, palette), "staged", palette.Success},
		{gitui.SectionIndicator(gitui.Conflicts, palette), "conflicts", palette.Danger},
		{gitui.HeadIndicator(gitui.Head{Kind: gitui.Attached, Ref: "main"}, palette), "main", palette.Git.Branch},
		{gitui.RefLabel(gitui.Ref{Kind: gitui.Tag, Name: "v1"}, palette), "tag v1", palette.Info},
		{gitui.RepositorySummary(repo, palette), "untracked 1", palette.Git.UntrackedFiles},
		{gitui.DivergenceIndicator(gitui.Divergence{Known: true, Behind: 2}, palette), "behind 2", palette.Git.Behind},
		{gitui.DiffSummary(gitui.DiffStat{State: gitui.TextDiff, Added: 3}, palette), "+3", palette.Git.AddedLines},
		{gitui.DiffHunk(hunk, palette), "1 +", palette.Git.DiffAdd},
	} {
		want := palette.Text().Foreground(test.color).Render(test.label)
		if !containsDiffANSI(test.view, want) {
			t.Errorf("%q did not use its Git foreground: %q", test.label, test.view)
		}
	}
}

func TestMissingGitColorUsesCallerTone(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast} {
		palette := theme.New(mode)
		palette.Warning = lipgloss.Color("#cc0000")
		want := palette.Text().Foreground(palette.Warning).Render("modified")

		if !containsDiffANSI(gitui.FileIndicator(gitui.Modified, palette), want) {
			t.Fatalf("%s file without a Git-specific color lost the caller's warning color", mode)
		}
	}
}
