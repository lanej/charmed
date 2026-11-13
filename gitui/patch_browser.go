package gitui

import (
	"fmt"

	"github.com/charmbracelet/x/ansi"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

// PatchObservation distinguishes an unavailable inventory from a known empty patch.
// Available=false requires no files; Reason explains the missing observation.
type PatchObservation struct {
	Available bool
	Reason    string
	Files     []FileDiff
}

// PatchSelection identifies one file in one change section.
type PatchSelection struct {
	FileID  string
	Section ChangeSection
}

// PatchBrowser copies supplied patches, coordinating file selection and native diff
// scrolling. Parents own Tab routing; direct List mutations require Refresh.
type PatchBrowser struct {
	ShowHelp      bool // Instructions are opt-in; resize after changing this field.
	List          *List
	Diff          *DiffViewer
	Details       *collection.ScrollPane
	files         map[string]FileDiff
	observation   PatchObservation
	selected      string
	size          layout.Size
	detailFocused bool
}

func NewPatchBrowser(observation PatchObservation, size layout.Size, t theme.Theme) (*PatchBrowser, error) {
	list, err := NewList("git/patch-browser", nil, size.Width, size.Height, t)
	if err != nil {
		return nil, err
	}

	diff := NewDiffViewer(FileDiff{State: UnavailableDiff}, size.Width, size.Height, t)

	m := &PatchBrowser{List: list, Diff: diff, Details: collection.NewScrollPane(&diff.Viewport, size, t)}
	if err := m.SetPatch(observation); err != nil {
		return nil, err
	}

	m.Resize(size)

	return m, nil
}

func patchID(id string, section ChangeSection) string {
	return fmt.Sprintf("%d:%s%s", len(id), id, section)
}

// SetPatch validates and deep-copies every diff before replacing the observation.
func (m *PatchBrowser) SetPatch(observation PatchObservation) error {
	if !observation.Available && len(observation.Files) > 0 {
		return fmt.Errorf("unavailable patch cannot contain files")
	}

	files := make(map[string]FileDiff, len(observation.Files))

	items := make([]Item, 0, len(observation.Files))
	for _, diff := range observation.Files {
		id := patchID(diff.FileID, diff.Section)
		if diff.FileID == "" {
			return fmt.Errorf("patch file needs an ID")
		}

		if _, exists := files[id]; exists {
			return fmt.Errorf("duplicate patch identity %q", id)
		}

		files[id] = cloneDiff(diff)
		items = append(items, Item{ID: id, Text: clean(diff.Path), Group: string(diff.Section)})
	}

	oldID, offset, x := m.List.SelectedID(), m.Diff.Viewport.YOffset(), m.Diff.Viewport.XOffset()
	if err := m.List.SetItems(items); err != nil {
		return err
	}

	m.files = files
	m.observation = PatchObservation{Available: observation.Available, Reason: observation.Reason}
	m.selected = ""
	m.Refresh()

	if oldID != "" && oldID == m.List.SelectedID() {
		m.Diff.Viewport.SetYOffset(offset)
		m.Diff.Viewport.SetXOffset(x)
	}

	return nil
}

func (m *PatchBrowser) Select(id string, section ChangeSection) bool {
	ok := m.List.Select(patchID(id, section))
	m.Refresh()

	return ok
}

func (m *PatchBrowser) Selected() (PatchSelection, bool) {
	diff, ok := m.files[m.List.SelectedID()]
	return PatchSelection{FileID: diff.FileID, Section: diff.Section}, ok
}

func (m *PatchBrowser) Refresh() {
	id := m.List.SelectedID()
	if diff, ok := m.files[id]; ok {
		if id != m.selected {
			m.Diff.SetDiff(diff)
		}
	} else {
		content := "No files changed"
		if len(m.files) > 0 {
			content = "No matching file selected"
		}

		if !m.observation.Available {
			content = "Patch unavailable"
			if m.observation.Reason != "" {
				content += ": " + clean(m.observation.Reason)
			}
		}

		m.Details.SetContent(content)
	}

	m.selected = id
}

func (m *PatchBrowser) FocusDetail(detail bool) {
	m.detailFocused = detail
	m.List.SetFocused(!detail)
	m.Resize(m.size)
}

func (m *PatchBrowser) DetailFocused() bool { return m.detailFocused }
func (m *PatchBrowser) Resize(size layout.Size) {
	m.size = size
	a, b := collection.InspectorSizes(size, m.detailFocused)
	m.List.ShowHelp = m.ShowHelp
	m.List.Resize(a)
	m.Details.Resize(b)
}

func (m *PatchBrowser) Update(msg tea.Msg) tea.Cmd {
	if m.detailFocused {
		return m.Details.Update(msg)
	}

	cmd := m.List.Update(msg)
	m.Refresh()

	return cmd
}

func (m *PatchBrowser) View() string {
	out := portable(collection.InspectorView(m.List.View(), m.Details.View(), m.size, m.detailFocused), m.Diff.Theme)
	if m.Diff.Theme.Mode == theme.Mono {
		return ansi.Strip(out)
	}

	return out
}
