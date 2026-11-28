package gitui

import (
	"cmp"
	"fmt"
	"image/color"
	"strings"
	"unicode"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// clean keeps caller text from injecting terminal controls or extra rows.
func clean(value string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsControl(r) {
			return ' '
		}

		return r
	}, ansi.Strip(value))
}

func CommitIdentity(id string, t theme.Theme) string {
	return portable(t.Heading().Render(ansi.Truncate(clean(id), 10, "")), t)
}

func RefLabel(ref Ref, t theme.Theme) string {
	kind := "unknown ref"
	ink := t.Info

	switch ref.Kind {
	case LocalBranch:
		kind = "branch"
		ink = cmp.Or(t.Git.Branch, ink)
	case RemoteBranch:
		kind = "remote"
		ink = cmp.Or(t.Git.Branch, ink)
	case Tag:
		kind = "tag"
	default:
	}

	return gitPill(kind+" "+shortRef(ref.Name), ink, t)
}

func HeadIndicator(head Head, t theme.Theme) string {
	switch head.Kind {
	case Attached:
		return gitPill("HEAD → "+shortRef(head.Ref), cmp.Or(t.Git.Branch, t.Info), t)
	case Detached:
		return portable(text.Pill("HEAD detached "+ansi.Truncate(clean(head.CommitID), 10, ""), text.Warning, t, text.Normal), t)
	case Unborn:
		return portable(text.Pill("HEAD unborn "+shortRef(head.Ref), text.Neutral, t, text.Normal), t)
	default:
		return portable(text.Pill("HEAD unavailable", text.Neutral, t, text.Normal), t)
	}
}

func FileIndicator(kind ChangeKind, t theme.Theme) string {
	ink := t.Foreground

	label := string(kind)
	switch kind {
	case Added, Copied:
		ink = cmp.Or(t.Git.DiffAdd, t.Success)
	case Modified:
		ink = cmp.Or(t.Git.Modified, t.Warning)
	case Renamed:
		ink = cmp.Or(t.Git.Renamed, t.Warning)
	case TypeChanged:
		ink = cmp.Or(t.Git.TypeChanged, t.Warning)
	case Deleted:
		ink = cmp.Or(t.Git.Deleted, t.Danger)
	case Conflicted:
		ink = cmp.Or(t.Git.Conflicted, t.Danger)
	case Untracked:
		ink = cmp.Or(t.Git.Untracked, t.Info)
	case Unchanged:
		label = "unchanged"
	default:
		label = "unknown change"
	}

	return gitPill(label, ink, t)
}

func gitPill(label string, ink color.Color, t theme.Theme) string {
	if !t.ASCII {
		t.Foreground = ink
	}

	return portable(text.Pill(label, text.Neutral, t, text.Normal), t)
}

func DivergenceIndicator(d Divergence, t theme.Theme) string {
	if !d.Known {
		return portable(t.Subtle().Render("ahead/behind unavailable"), t)
	}

	ahead := t.Text().Foreground(cmp.Or(t.Git.Ahead, t.Foreground)).Render(fmt.Sprintf("ahead %d", d.Ahead))
	behind := t.Text().Foreground(cmp.Or(t.Git.Behind, t.Foreground)).Render(fmt.Sprintf("behind %d", d.Behind))

	return portable(ahead+" · "+behind, t)
}

func DiffSummary(stat DiffStat, t theme.Theme) string {
	switch stat.State {
	case TextDiff:
		return portable(t.Text().Foreground(cmp.Or(t.Git.AddedLines, t.Success)).Render(fmt.Sprintf("+%d", stat.Added))+" "+t.Text().Foreground(cmp.Or(t.Git.RemovedLines, t.Danger)).Render(fmt.Sprintf("-%d", stat.Deleted)), t)
	case BinaryDiff:
		return portable(t.Subtle().Render("binary"), t)
	case UnavailableDiff:
		return portable(t.Subtle().Render("diff unavailable"), t)
	default:
		return portable(t.Subtle().Render("diff unavailable"), t)
	}
}

func WorktreeIndicators(w Worktree, t theme.Theme) string {
	var labels []string
	if w.Main {
		labels = append(labels, "main")
	} else {
		labels = append(labels, "linked")
	}

	if w.Current {
		labels = append(labels, "current")
	}

	if w.Locked {
		labels = append(labels, "locked: "+clean(w.LockReason))
	}

	if w.Prunable {
		labels = append(labels, "prunable: "+clean(w.PruneReason))
	}

	if !w.StatusKnown || w.Status.Unavailable {
		labels = append(labels, "status unavailable")
	} else if len(w.Status.Files) == 0 {
		labels = append(labels, "clean")
	} else {
		labels = append(labels, fmt.Sprintf("%d changed files", len(w.Status.Files)))
	}

	return portable(text.Pill(strings.Join(labels, " · "), text.Neutral, t, text.Normal), t)
}

func portable(value string, t theme.Theme) string {
	if t.ASCII {
		return layout.ASCII(value)
	}

	return value
}

func shortRef(name string) string {
	for _, prefix := range []string{"refs/heads/", "refs/remotes/", "refs/tags/"} {
		if strings.HasPrefix(name, prefix) {
			return clean(strings.TrimPrefix(name, prefix))
		}
	}

	return clean(name)
}
