package gitui

import (
	"slices"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/interaction"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

type BranchList struct{ *List }

func NewBranchList(branches []Branch, width, height int, t theme.Theme) (*BranchList, error) {
	list, err := NewList("git/branches", nil, width, height, t)
	if err != nil {
		return nil, err
	}

	m := &BranchList{List: list}
	if err := m.SetBranches(branches); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *BranchList) SetBranches(branches []Branch) error {
	items := make([]Item, len(branches))
	for i, branch := range branches {
		items[i] = Item{ID: branch.Ref.Name, Text: BranchRow(branch, m.Theme)}
	}

	return m.SetItems(items)
}

type StashList struct{ *List }

func NewStashList(stashes []Stash, width, height int, t theme.Theme) (*StashList, error) {
	list, err := NewList("git/stashes", nil, width, height, t)
	if err != nil {
		return nil, err
	}

	m := &StashList{List: list}
	if err := m.SetStashes(stashes); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *StashList) SetStashes(stashes []Stash) error {
	items := make([]Item, len(stashes))
	for i, stash := range stashes {
		items[i] = Item{ID: stash.ID, Text: StashRow(stash, m.Theme)}
	}

	return m.SetItems(items)
}

type WorktreeList struct{ *List }

func NewWorktreeList(worktrees []Worktree, width, height int, t theme.Theme) (*WorktreeList, error) {
	list, err := NewList("git/worktrees", nil, width, height, t)
	if err != nil {
		return nil, err
	}

	m := &WorktreeList{List: list}
	if err := m.SetWorktrees(worktrees); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *WorktreeList) SetWorktrees(worktrees []Worktree) error {
	items := make([]Item, len(worktrees))
	for i, worktree := range worktrees {
		items[i] = Item{ID: worktree.ID, Text: WorktreeRow(worktree, m.Theme)}
	}

	return m.SetItems(items)
}

// CommitHistory contains native collection controls without exposing routing state.
type CommitHistory struct{ *List }

func NewCommitHistory(history History, width, height int, t theme.Theme) (*CommitHistory, error) {
	list, err := NewList("git/history", nil, width, height, t)
	if err != nil {
		return nil, err
	}

	list.preserveRows = true

	m := &CommitHistory{List: list}
	if err := m.SetHistory(history); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *CommitHistory) SetHistory(history History) error {
	rows, width, err := graph(history, m.Theme)
	if err != nil {
		return err
	}

	items := make([]Item, len(history.Commits))
	for i, commit := range history.Commits {
		items[i] = Item{ID: commit.ID, Text: rows[i].prefix + " " + CommitRow(commit, history.Refs, history.Head, m.Theme), After: rows[i].after}
	}

	m.minWidth = width + 14

	return m.SetItems(items)
}

type changeIdentity struct {
	id      string
	section ChangeSection
}

// Status composes the repository summary and independently selectable change sections.
type Status struct {
	List    *List
	repo    Repository
	changes map[string]changeIdentity
	size    layout.Size
}

func NewStatus(repo Repository, width, height int, t theme.Theme) (*Status, error) {
	list, err := NewList("git/status", nil, width, height, t)
	if err != nil {
		return nil, err
	}

	m := &Status{List: list}
	if err := m.SetRepository(repo); err != nil {
		return nil, err
	}

	m.Resize(layout.Size{Width: width, Height: height})

	return m, nil
}

func (m *Status) SetRepository(repo Repository) error {
	if err := validateFiles(repo.Files); err != nil {
		return err
	}

	if repo.Unavailable {
		repo.Files = nil
	}

	var items []Item

	changes := map[string]changeIdentity{}

	for _, section := range []ChangeSection{Staged, Unstaged, UntrackedFiles, Conflicts} {
		group, err := NewChangeGroup(section, repo.Files, m.List.Theme)
		if err != nil {
			return err
		}

		for _, item := range group.Items() {
			fileID := item.ID
			item.ID = string(section) + ":" + fileID
			items = append(items, item)
			changes[item.ID] = changeIdentity{id: fileID, section: section}
		}
	}

	if err := m.List.SetItems(items); err != nil {
		return err
	}

	repo.Files = slices.Clone(repo.Files)
	m.repo, m.changes = repo, changes

	m.List.EmptyText = "Working tree clean"
	if repo.Unavailable {
		m.List.EmptyText = "No status observation"
	}

	m.List.render()

	return nil
}

func (m *Status) SelectedChange() (string, ChangeSection) {
	selected := m.changes[m.List.SelectedID()]
	return selected.id, selected.section
}

func (m *Status) Update(msg tea.Msg) tea.Cmd {
	cmd := m.List.Update(msg)
	if cmd == nil {
		return nil
	}

	identity := m.changes[m.List.SelectedID()]

	return func() tea.Msg {
		result := cmd()
		if _, ok := result.(interaction.Selected); ok {
			return interaction.Selected{Component: "git/status/" + string(identity.section), ID: identity.id}
		}

		return result
	}
}

func (m *Status) Init() tea.Cmd { return nil }
func (m *Status) Close() error  { return nil }
func (m *Status) Resize(size layout.Size) {
	m.size = size
	m.List.Resize(layout.Size{Width: size.Width, Height: max(1, size.Height-4)})
}

func (m *Status) View() string {
	return layout.Fit(RepositorySummary(m.repo, m.List.Theme)+"\n\n"+m.List.View(), m.size)
}
