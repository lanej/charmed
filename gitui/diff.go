package gitui

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

// DiffGutter selects measured line-number columns.
type DiffGutter string

const (
	SingleGutter DiffGutter = ""
	DualGutter   DiffGutter = "dual"
	AutoGutter   DiffGutter = "auto"
)

// DiffOptions defaults to a compact single gutter and path-aware syntax.
// Language overrides detection; DisableSyntax keeps diff surfaces and signs.
type DiffOptions struct {
	Language      string
	Gutter        DiffGutter
	DisableSyntax bool
}

func diffOptions(options []DiffOptions) DiffOptions {
	var config DiffOptions
	for _, option := range options {
		config = option
	}

	return config
}

func diffHeading(t theme.Theme) lipgloss.Style {
	if t.Mode == theme.Mono {
		return lipgloss.NewStyle().Bold(true)
	}

	return t.Heading()
}

func DiffHunk(hunk Hunk, t theme.Theme, options ...DiffOptions) string {
	config := diffOptions(options)

	language := config.Language
	if language == "" || config.DisableSyntax {
		language = "text"
	}

	return renderDiffHunk(hunk, language, language, t, config)
}

func renderDiffHunk(hunk Hunk, oldLanguage, newLanguage string, t theme.Theme, options ...DiffOptions) string {
	config := diffOptions(options)
	if config.DisableSyntax {
		oldLanguage, newLanguage = "text", "text"
	}

	if t.Mode == theme.Mono {
		oldLanguage, newLanguage = "text", "text"
	}

	oldLines, newLines := hunkSyntax(hunk, oldLanguage, true), hunkSyntax(hunk, newLanguage, false)
	edits := hunkEdits(hunk, newLines)
	dualNumbers := false

	switch config.Gutter {
	case SingleGutter:
	case DualGutter:
		dualNumbers = true
	case AutoGutter:
		dualNumbers = shiftedLineNumbers(hunk)
	default:
	}

	oldWidth, newWidth, singleWidth := 1, 1, 1
	for _, line := range hunk.Lines {
		oldWidth = max(oldWidth, len(fmt.Sprint(max(0, line.OldLine))))
		newWidth = max(newWidth, len(fmt.Sprint(max(0, line.NewLine))))

		number := line.NewLine
		if line.Kind == DeletedLine || number == 0 {
			number = line.OldLine
		}

		singleWidth = max(singleWidth, len(fmt.Sprint(max(0, number))))
	}

	lines := []string{diffHeading(t).Render(clean(hunk.Header))}
	for i, line := range hunk.Lines {
		prefix := " "
		style := diffLineStyle(line.Kind, false, t)
		syntax := newLines[i]

		switch line.Kind {
		case ContextLine:
		case AddedLine:
			prefix = "+"

			if t.Mode != theme.Mono {
				style = style.Foreground(t.Success)
			}
		case DeletedLine:
			prefix = "-"

			if t.Mode != theme.Mono {
				style = style.Foreground(t.Danger)
			}

			syntax = oldLines[i]
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

		gutter := fmt.Sprintf("%*s %*s %s", oldWidth, oldLine, newWidth, newLine, prefix)
		if !dualNumbers {
			number := newLine
			if line.Kind == DeletedLine || number == "" {
				number = oldLine
			}

			gutter = fmt.Sprintf("%*s %s", singleWidth, number, prefix)
		}

		gutter = style.Render(gutter)

		lines = append(lines, gutter+renderSyntaxLine(syntax, line.Kind, edits[i], t))
		if line.NoNewline {
			marker := t.Subtle()
			if t.Mode == theme.Mono {
				marker = lipgloss.NewStyle().Faint(true)
			}

			lines = append(lines, marker.Render("\\ No newline at end of file"))
		}
	}

	return portable(strings.Join(lines, "\n"), t)
}

func diffContent(diff FileDiff, t theme.Theme, options ...DiffOptions) string {
	config := diffOptions(options)
	if config.Language != "" {
		diff.Language = config.Language
	}

	header := clean(diff.Path)
	if diff.OriginalPath != "" {
		header = clean(diff.OriginalPath) + " → " + header
	}

	header = diffHeading(t).Render(header) + "\n"

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

	oldPath := diff.OriginalPath
	if oldPath == "" {
		oldPath = diff.Path
	}

	oldLanguage, newLanguage := diffLanguage(oldPath, diff.Language), diffLanguage(diff.Path, diff.Language)
	for _, hunk := range diff.Hunks {
		hunks = append(hunks, renderDiffHunk(hunk, oldLanguage, newLanguage, t, config))
	}

	return header + strings.Join(hunks, "\n")
}

// DiffViewer manages content and dimensions; native key bindings and styles survive SetDiff.
type DiffViewer struct {
	Viewport viewport.Model
	Theme    theme.Theme
	Options  DiffOptions
}

func NewDiffViewer(diff FileDiff, width, height int, t theme.Theme, options ...DiffOptions) *DiffViewer {
	m := &DiffViewer{Viewport: collection.NewViewport("", width, height), Theme: t, Options: diffOptions(options)}
	m.SetDiff(diff)

	return m
}

func (m *DiffViewer) SetDiff(diff FileDiff) {
	m.Viewport.SetContent(diffContent(diff, m.Theme, m.Options))
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
