package gitui

import (
	"fmt"
	"slices"
	"strings"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

// ReviewEntry is either a group (Children) or a leaf (Diff). A nil leaf Diff
// renders an explicit unavailable state. IDs are caller identities, independent
// of paths: staged and unstaged versions can appear together. Stat is optional;
// missing statistics are never counted as zero. Reviewed initializes new IDs.
type ReviewEntry struct {
	ID, Label string
	// ContentRevision identifies changes absent from the patch, such as binary bytes.
	ContentRevision string
	Diff            *FileDiff
	Stat            *DiffStat
	Children        []ReviewEntry
	Reviewed        bool
}

// ReviewFile is the input to the optional slash-separated file-path helper.
// ID identifies the review item; FileDiff.FileID and Section retain Git meaning.
type ReviewFile struct {
	ID              string
	ContentRevision string
	Diff            FileDiff
	Stat            *DiffStat
	Reviewed        bool
}

// DiffReview composes a generic outline/document with Git diff presentation.
// It owns local review state; applications persist ReviewedIDs themselves.
// SetEntries clears review only when a retained leaf's actual diff changes.
// Native outline/viewport customization follows OutlineDocument's lifecycle.
type DiffReview struct {
	*pattern.OutlineDocument
	ReviewKey key.Binding
	Options   DiffOptions
	entries   []ReviewEntry
	leaves    map[string]ReviewEntry
	reviewed  map[string]bool
	Theme     theme.Theme
}

func NewDiffReview(entries []ReviewEntry, size layout.Size, t theme.Theme, options ...DiffOptions) (*DiffReview, error) {
	document, err := pattern.NewOutlineDocument(nil, size, t)
	if err != nil {
		return nil, err
	}

	m := &DiffReview{OutlineDocument: document, Theme: t, Options: diffOptions(options), ReviewKey: key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "reviewed"))}

	m.AppKeys = []key.Binding{m.ReviewKey}
	if err := m.SetEntries(entries); err != nil {
		return nil, err
	}

	return m, nil
}

// NewFileDiffReview groups slash-separated paths without reading a filesystem.
func NewFileDiffReview(files []ReviewFile, size layout.Size, t theme.Theme, options ...DiffOptions) (*DiffReview, error) {
	entries, err := FileReviewEntries(files)
	if err != nil {
		return nil, err
	}

	return NewDiffReview(entries, size, t, options...)
}

// SetDiffOptions refreshes presentation while retaining selection, review marks
// and folds. Direct Options changes take effect on the next SetEntries call.
func (m *DiffReview) SetDiffOptions(options DiffOptions) {
	m.Options = options
	m.Document.Refresh()
}

// SetEntries copies all nested data before replacement. Invalid or duplicate
// identities and entries containing both children and a diff fail atomically.
// Metadata changes retain review and folding. Changed content opens the leaf;
// a collapsed ancestor still controls whether that leaf is visible.
func (m *DiffReview) SetEntries(entries []ReviewEntry) error {
	seen := make(map[string]bool)
	leaves := make(map[string]ReviewEntry)
	reviewed := make(map[string]bool)

	var changed []string

	var clone func([]ReviewEntry) ([]ReviewEntry, error)

	clone = func(source []ReviewEntry) ([]ReviewEntry, error) {
		result := make([]ReviewEntry, len(source))
		for i, entry := range source {
			if entry.ID == "" || seen[entry.ID] {
				return nil, fmt.Errorf("empty or duplicate review ID %q", entry.ID)
			}

			if len(entry.Children) > 0 && entry.Diff != nil {
				return nil, fmt.Errorf("review group %q also has a diff", entry.ID)
			}

			seen[entry.ID] = true
			if entry.Diff != nil {
				diff := cloneDiff(*entry.Diff)
				entry.Diff = &diff
			}

			if entry.Stat != nil {
				stat := *entry.Stat
				entry.Stat = &stat
			}

			children, err := clone(entry.Children)
			if err != nil {
				return nil, err
			}

			entry.Children = children
			if len(children) == 0 {
				reviewed[entry.ID] = entry.Reviewed
				if old, exists := m.leaves[entry.ID]; exists {
					reviewed[entry.ID] = m.reviewed[entry.ID]
					if old.ContentRevision != entry.ContentRevision || !sameReviewContent(old.Diff, entry.Diff) {
						reviewed[entry.ID] = false
						changed = append(changed, entry.ID)
					}
				}

				leaves[entry.ID] = entry
			}

			result[i] = entry
		}

		return result, nil
	}

	cloned, err := clone(entries)
	if err != nil {
		return err
	}

	if err := m.OutlineDocument.SetEntries(m.outlineEntries(cloned, reviewed)); err != nil {
		return err
	}

	m.entries, m.leaves, m.reviewed = cloned, leaves, reviewed
	for _, id := range changed {
		m.SetCollapsed(id, false)
	}

	return nil
}

func sameReviewContent(a, b *FileDiff) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	return a.State == b.State && a.Reason == b.Reason && slices.EqualFunc(a.Hunks, b.Hunks, func(a, b Hunk) bool {
		return a.Header == b.Header && slices.Equal(a.Lines, b.Lines)
	})
}

func (m *DiffReview) Reviewed(id string) bool { return m.reviewed[id] }

// ReviewedIDs returns marks in document order, including hidden descendants.
func (m *DiffReview) ReviewedIDs() []string {
	var ids []string

	var visit func([]ReviewEntry)

	visit = func(entries []ReviewEntry) {
		for _, entry := range entries {
			if m.reviewed[entry.ID] {
				ids = append(ids, entry.ID)
			}

			visit(entry.Children)
		}
	}

	visit(m.entries)

	return ids
}

// SetReviewed marks a leaf and folds it, or clears its mark and opens it.
// Repeating the current mark is a no-op, preserving deliberate reopening.
func (m *DiffReview) SetReviewed(id string, reviewed bool) bool {
	if _, exists := m.leaves[id]; !exists {
		return false
	}

	if m.reviewed[id] == reviewed {
		return true
	}

	m.reviewed[id] = reviewed
	// Identities were validated by SetEntries; this projection cannot fail.
	if err := m.OutlineDocument.SetEntries(m.outlineEntries(m.entries, m.reviewed)); err != nil {
		m.reviewed[id] = !reviewed
		return false
	}

	m.SetCollapsed(id, reviewed)

	return true
}

type reviewTotals struct {
	files, reviewed int
	added, deleted  int
	known           bool
}

func totalReview(entries []ReviewEntry, reviewed map[string]bool) reviewTotals {
	total := reviewTotals{known: true}

	for _, entry := range entries {
		if len(entry.Children) > 0 {
			child := totalReview(entry.Children, reviewed)
			total.files += child.files
			total.reviewed += child.reviewed
			total.added += child.added
			total.deleted += child.deleted
			total.known = total.known && child.known

			continue
		}

		total.files++
		if reviewed[entry.ID] {
			total.reviewed++
		}

		if entry.Stat != nil && entry.Stat.State == TextDiff && entry.Stat.Added >= 0 && entry.Stat.Deleted >= 0 {
			total.added += entry.Stat.Added
			total.deleted += entry.Stat.Deleted
		} else {
			total.known = false
		}
	}

	return total
}

func (m *DiffReview) outlineEntries(entries []ReviewEntry, reviewed map[string]bool) []pattern.OutlineEntry {
	result := make([]pattern.OutlineEntry, len(entries))
	for i, entry := range entries {
		label := clean(entry.Label)
		if label == "" {
			label = clean(entry.ID)
		}

		section := pattern.OutlineEntry{ID: entry.ID, Title: label, Children: m.outlineEntries(entry.Children, reviewed)}
		if len(entry.Children) > 0 {
			total := totalReview(entry.Children, reviewed)

			section.Summary = fmt.Sprintf("%d files · %d reviewed", total.files, total.reviewed)
			if total.known {
				section.Summary = DiffSummary(DiffStat{State: TextDiff, Added: total.added, Deleted: total.deleted}, m.Theme) + " · " + section.Summary
			} else {
				section.Summary += " · stats unavailable"
			}
		} else {
			if entry.Stat != nil && entry.Stat.Added >= 0 && entry.Stat.Deleted >= 0 {
				section.Summary = DiffSummary(*entry.Stat, m.Theme)
			}

			if reviewed[entry.ID] {
				section.Title = m.Theme.Glyph(theme.GlyphOK) + " " + section.Title
				section.Summary += " · reviewed"
				section.Collapsed = true
			}

			diff := FileDiff{State: UnavailableDiff, Reason: "content was not supplied"}
			if entry.Diff != nil {
				diff = *entry.Diff
			}

			section.Body = func(int) string {
				content := diffContent(diff, m.Theme, m.Options)
				if diff.OriginalPath != "" {
					return content
				}

				_, body, _ := strings.Cut(content, "\n")

				return body
			}
		}

		result[i] = section
	}

	return result
}

func (m *DiffReview) Update(msg tea.Msg) tea.Cmd {
	if pressed, ok := msg.(tea.KeyPressMsg); ok && key.Matches(pressed, m.ReviewKey) {
		id := m.SelectedID()

		m.SetReviewed(id, !m.Reviewed(id))

		return nil
	}

	return m.OutlineDocument.Update(msg)
}

// FileReviewEntries preserves first-seen sibling order and creates directory
// groups. Directory IDs use a reserved prefix; conflicting caller IDs fail.
// Paths are display input only and are never cleaned or resolved on disk.
func FileReviewEntries(files []ReviewFile) ([]ReviewEntry, error) {
	const directoryPrefix = "\x00directory:"

	var entries []ReviewEntry

	seen := make(map[string]bool)
	for _, file := range files {
		if file.ID == "" || seen[file.ID] || strings.HasPrefix(file.ID, directoryPrefix) {
			return nil, fmt.Errorf("invalid or duplicate review file ID %q", file.ID)
		}

		if file.Diff.Path == "" {
			return nil, fmt.Errorf("review file %q has no path", file.ID)
		}

		seen[file.ID] = true
		parts := strings.Split(file.Diff.Path, "/")
		diff := cloneDiff(file.Diff)

		entry := ReviewEntry{ID: file.ID, Label: parts[len(parts)-1], Diff: &diff, Reviewed: file.Reviewed, ContentRevision: file.ContentRevision}
		if file.Stat != nil {
			stat := *file.Stat
			entry.Stat = &stat
		}

		insertFileReview(&entries, parts[:len(parts)-1], "", entry)
	}

	return entries, nil
}

func insertFileReview(entries *[]ReviewEntry, directories []string, prefix string, leaf ReviewEntry) {
	if len(directories) == 0 {
		*entries = append(*entries, leaf)
		return
	}

	path := prefix + directories[0] + "/"
	id := "\x00directory:" + path

	i := slices.IndexFunc(*entries, func(entry ReviewEntry) bool { return entry.ID == id })
	if i < 0 {
		*entries = append(*entries, ReviewEntry{ID: id, Label: directories[0] + "/"})
		i = len(*entries) - 1
	}

	insertFileReview(&(*entries)[i].Children, directories[1:], path, leaf)
}
