package gitui

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func DiffHunk(hunk Hunk, t theme.Theme) string {
	lines := []string{t.Heading().Render(clean(hunk.Header))}
	for _, line := range hunk.Lines {
		prefix := " "
		style := t.Text()

		switch line.Kind {
		case ContextLine:
		case AddedLine:
			prefix = "+"
			style = style.Foreground(t.Success)
		case DeletedLine:
			prefix = "-"
			style = style.Foreground(t.Danger)
		default:
			prefix = "?"
		}

		oldLine, newLine := "", ""
		if line.OldLine > 0 {
			oldLine = fmt.Sprint(line.OldLine)
		}

		if line.NewLine > 0 {
			newLine = fmt.Sprint(line.NewLine)
		}

		lines = append(lines, style.Render(fmt.Sprintf("%5s %5s %s%s", oldLine, newLine, prefix, clean(line.Text))))
		if line.NoNewline {
			lines = append(lines, t.Subtle().Render("\\ No newline at end of file"))
		}
	}

	return portable(strings.Join(lines, "\n"), t)
}

func diffContent(diff FileDiff, t theme.Theme) string {
	header := clean(diff.Path)
	if diff.OriginalPath != "" {
		header = clean(diff.OriginalPath) + " → " + header
	}

	header = t.Heading().Render(header) + "\n"

	switch diff.State {
	case BinaryDiff:
		return header + "Binary content changed"
	case UnavailableDiff:
		return header + "Diff unavailable: " + clean(diff.Reason)
	case TextDiff:
		if len(diff.Hunks) == 0 {
			return header + "No textual changes"
		}
	default:
		return header + "Diff unavailable: unknown state"
	}

	var hunks []string
	for _, hunk := range diff.Hunks {
		hunks = append(hunks, DiffHunk(hunk, t))
	}

	return header + strings.Join(hunks, "\n")
}

// DiffViewer manages content and dimensions; native key bindings and styles survive SetDiff.
type DiffViewer struct {
	Viewport viewport.Model
	Theme    theme.Theme
}

func NewDiffViewer(diff FileDiff, width, height int, t theme.Theme) *DiffViewer {
	m := &DiffViewer{Viewport: collection.NewViewport("", width, height), Theme: t}
	m.SetDiff(diff)

	return m
}

func (m *DiffViewer) SetDiff(diff FileDiff) {
	m.Viewport.SetContent(diffContent(diff, m.Theme))
	m.Viewport.GotoTop()
}

func (m *DiffViewer) Resize(size layout.Size) {
	m.Viewport.SetWidth(max(1, size.Width))
	m.Viewport.SetHeight(max(1, size.Height))
}

func (m *DiffViewer) Update(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd

	m.Viewport, cmd = m.Viewport.Update(msg)

	return cmd
}
func (m *DiffViewer) View() string  { return portable(m.Viewport.View(), m.Theme) }
func (m *DiffViewer) Init() tea.Cmd { return nil }
func (m *DiffViewer) Close() error  { return nil }
