package gitui

import (
	"fmt"
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

	switch ref.Kind {
	case LocalBranch:
		kind = "branch"
	case RemoteBranch:
		kind = "remote"
	case Tag:
		kind = "tag"
	default:
	}

	return portable(text.Pill(kind+" "+shortRef(ref.Name), text.Info, t, text.Normal), t)
}

func HeadIndicator(head Head, t theme.Theme) string {
	switch head.Kind {
	case Attached:
		return portable(text.Pill("HEAD → "+shortRef(head.Ref), text.Info, t, text.Normal), t)
	case Detached:
		return portable(text.Pill("HEAD detached "+ansi.Truncate(clean(head.CommitID), 10, ""), text.Warning, t, text.Normal), t)
	case Unborn:
		return portable(text.Pill("HEAD unborn "+shortRef(head.Ref), text.Neutral, t, text.Normal), t)
	default:
		return portable(text.Pill("HEAD unavailable", text.Neutral, t, text.Normal), t)
	}
}

func FileIndicator(kind ChangeKind, t theme.Theme) string {
	tone := text.Neutral

	label := string(kind)
	switch kind {
	case Added, Copied:
		tone = text.Success
	case Modified, Renamed, TypeChanged:
		tone = text.Warning
	case Deleted, Conflicted:
		tone = text.Danger
	case Untracked:
		tone = text.Info
	case Unchanged:
		label = "unchanged"
	default:
		label = "unknown change"
	}

	return portable(text.Pill(label, tone, t, text.Normal), t)
}

func DivergenceIndicator(d Divergence, t theme.Theme) string {
	if !d.Known {
		return portable(t.Subtle().Render("ahead/behind unavailable"), t)
	}

	return portable(t.Text().Render(fmt.Sprintf("ahead %d · behind %d", d.Ahead, d.Behind)), t)
}

func DiffSummary(stat DiffStat, t theme.Theme) string {
	switch stat.State {
	case TextDiff:
		return portable(t.Text().Render(fmt.Sprintf("+%d -%d", stat.Added, stat.Deleted)), t)
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
