package gitui

import (
	"fmt"
	"slices"
	"strings"

	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// SectionIndicator is the state atom for a file's relationship to the index.
// It is independent of FileIndicator, which describes how that file changed.
func SectionIndicator(section ChangeSection, t theme.Theme) string {
	label := string(section)
	tone := text.Neutral

	switch section {
	case Staged:
		tone = text.Success
	case Unstaged:
		tone = text.Warning
	case UntrackedFiles:
		tone = text.Info
	case Conflicts:
		tone = text.Danger
	default:
		label = "unknown section"
	}

	return portable(text.Pill(label, tone, t, text.Normal), t)
}

// ChangeGroup is a presentation molecule: a section atom, count and matching file
// rows. It owns no focus, viewport, input handling or repository access.
type ChangeGroup struct {
	section ChangeSection
	files   []FileChange
	theme   theme.Theme
}

func NewStagedChanges(files []FileChange, t theme.Theme) (ChangeGroup, error) {
	return NewChangeGroup(Staged, files, t)
}

func NewUnstagedChanges(files []FileChange, t theme.Theme) (ChangeGroup, error) {
	return NewChangeGroup(Unstaged, files, t)
}

func NewUntrackedChanges(files []FileChange, t theme.Theme) (ChangeGroup, error) {
	return NewChangeGroup(UntrackedFiles, files, t)
}

func NewConflictChanges(files []FileChange, t theme.Theme) (ChangeGroup, error) {
	return NewChangeGroup(Conflicts, files, t)
}

// NewChangeGroup selects and copies matching files from a supplied observation.
// Mixed index/worktree changes belong to both groups; conflicts belong only to
// Conflicts. Invalid sections and empty or duplicate file identities fail early.
func NewChangeGroup(section ChangeSection, files []FileChange, t theme.Theme) (ChangeGroup, error) {
	switch section {
	case Staged, Unstaged, UntrackedFiles, Conflicts:
	default:
		return ChangeGroup{}, fmt.Errorf("unknown change section %q", section)
	}

	if err := validateFiles(files); err != nil {
		return ChangeGroup{}, err
	}

	group := ChangeGroup{section: section, theme: t}
	for _, file := range files {
		if slices.Contains(fileSections(file), section) {
			group.files = append(group.files, file)
		}
	}

	return group, nil
}

func validateFiles(files []FileChange) error {
	seen := make(map[string]bool, len(files))
	for _, file := range files {
		if file.ID == "" || seen[file.ID] {
			return fmt.Errorf("empty or duplicate file ID %q", file.ID)
		}

		seen[file.ID] = true
	}

	return nil
}

func (g ChangeGroup) Section() ChangeSection { return g.section }
func (g ChangeGroup) Count() int             { return len(g.files) }
func (g ChangeGroup) Files() []FileChange    { return slices.Clone(g.files) }

// Items lets an owning interactive composition reuse the molecule's exact rows.
// File IDs remain unchanged; the owner supplies section context for selection.
func (g ChangeGroup) Items() []Item {
	items := make([]Item, len(g.files))
	for i, file := range g.files {
		items[i] = Item{ID: file.ID, Text: FileChangeRow(file, g.section, g.theme), Group: fmt.Sprintf("%s (%d)", g.section, g.Count())}
	}

	return items
}

func (g ChangeGroup) View() string {
	lines := []string{SectionIndicator(g.section, g.theme) + g.theme.Text().Render(fmt.Sprintf(" %d", g.Count()))}
	for _, item := range g.Items() {
		lines = append(lines, g.theme.Control().Render(item.Text))
	}

	if g.Count() == 0 {
		lines = append(lines, g.theme.Control().Render("No files"))
	}

	return portable(strings.Join(lines, "\n"), g.theme)
}
