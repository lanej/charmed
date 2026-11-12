package story

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/interaction"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
)

type gitBrowserScene struct {
	environment Environment
	status      *gitui.Status
	history     *gitui.CommitHistory
	branches    *gitui.BranchList
	stashes     *gitui.StashList
	worktrees   *gitui.WorktreeList
	detail      *gitui.DiffViewer
	page        int
	focusDetail bool
	size        layout.Size
	context     string
	detailKey   string
}

func gitBrowser(e Environment) (runtime.Scene, error) {
	t := e.Config.Palette()

	repo, history, branches, stashes, worktrees := gitRepository(), gitHistory(), gitBranches(), gitStashes(), gitWorktrees()
	if e.Config.State == "empty" {
		repo.Files = nil
		history.Commits = nil
		history.Refs = nil
		branches = nil
		stashes = nil
		worktrees = nil
	}

	status, err := gitui.NewStatus(repo, 60, 24, t)
	if err != nil {
		return nil, err
	}

	log, err := gitui.NewCommitHistory(history, 60, 24, t)
	if err != nil {
		return nil, err
	}

	branchList, err := gitui.NewBranchList(branches, 60, 24, t)
	if err != nil {
		return nil, err
	}

	stashList, err := gitui.NewStashList(stashes, 60, 24, t)
	if err != nil {
		return nil, err
	}

	worktreeList, err := gitui.NewWorktreeList(worktrees, 60, 24, t)
	if err != nil {
		return nil, err
	}

	m := &gitBrowserScene{environment: e, status: status, history: log, branches: branchList, stashes: stashList, worktrees: worktreeList, detail: gitui.NewDiffViewer(gitui.FileDiff{State: gitui.UnavailableDiff}, 60, 24, t), context: "~/src/charmed"}
	m.Resize(e.Config.Size())
	m.syncDetail()

	return m, nil
}

func (m *gitBrowserScene) active() *gitui.List {
	switch m.page {
	case 0:
		return m.status.List
	case 1:
		return m.history.List
	case 2:
		return m.branches.List
	case 3:
		return m.stashes.List
	default:
		return m.worktrees.List
	}
}
func (m *gitBrowserScene) Init() tea.Cmd { return nil }
func (m *gitBrowserScene) Close() error  { return nil }
func (m *gitBrowserScene) Resize(size layout.Size) {
	m.size = size
	left, right := m.paneSizes()
	content := layout.PanelContentSize(left)
	m.status.Resize(content)

	for _, list := range []*gitui.List{m.history.List, m.branches.List, m.stashes.List, m.worktrees.List} {
		list.Resize(content)
	}

	m.detail.Resize(layout.PanelContentSize(right))
}

func (m *gitBrowserScene) paneSizes() (layout.Size, layout.Size) {
	height := max(1, m.size.Height-4)
	if m.size.Width < 100 {
		return layout.Size{Width: m.size.Width, Height: height / 2}, layout.Size{Width: m.size.Width, Height: height - height/2}
	}

	return layout.Size{Width: m.size.Width * 3 / 5, Height: height}, layout.Size{Width: m.size.Width - m.size.Width*3/5, Height: height}
}

func (m *gitBrowserScene) Update(msg tea.Msg) tea.Cmd {
	if selected, ok := msg.(interaction.Selected); ok {
		emit(m.environment, selected)

		if selected.Component == "git/worktrees" {
			for _, w := range gitWorktrees() {
				if w.ID != selected.ID {
					continue
				}

				repo := w.Status
				if !w.StatusKnown {
					repo = gitui.Repository{Name: w.Path, Head: w.Head, Unavailable: true, Reason: "worktree status was not supplied"}
				}

				if err := m.status.SetRepository(repo); err != nil {
					m.detail.Viewport.SetContent(err.Error())
					return nil
				}

				history := gitHistory()

				history.Head = w.Head
				if err := m.history.SetHistory(history); err != nil {
					m.detail.Viewport.SetContent(err.Error())
					return nil
				}

				m.history.Select(w.Head.CommitID)

				branches := gitBranches()
				for i := range branches {
					branches[i].Current = w.Head.Kind == gitui.Attached && branches[i].Ref.Name == w.Head.Ref
				}

				if err := m.branches.SetBranches(branches); err != nil {
					m.detail.Viewport.SetContent(err.Error())
					return nil
				}

				m.context = w.Path
				m.detailKey = ""
			}
		}
	}

	if key, ok := msg.(tea.KeyPressMsg); ok {
		switch key.String() {
		case "ctrl+right":
			m.page = (m.page + 1) % 5
			m.focusDetail = false
		case "ctrl+left":
			m.page = (m.page + 4) % 5
			m.focusDetail = false
		case "tab", "shift+tab":
			m.focusDetail = !m.focusDetail
		default:
			if m.focusDetail {
				return m.detail.Update(msg)
			}

			var cmd tea.Cmd
			if m.page == 0 {
				cmd = m.status.Update(msg)
			} else {
				cmd = m.active().Update(msg)
			}

			m.syncDetail()

			return cmd
		}
	}

	if _, key := msg.(tea.KeyPressMsg); !key {
		var cmd tea.Cmd
		if m.focusDetail {
			cmd = m.detail.Update(msg)
		} else if m.page == 0 {
			cmd = m.status.Update(msg)
		} else {
			cmd = m.active().Update(msg)
		}

		m.syncDetail()

		return cmd
	}

	for _, list := range []*gitui.List{m.status.List, m.history.List, m.branches.List, m.stashes.List, m.worktrees.List} {
		list.SetFocused(list == m.active() && !m.focusDetail)
	}

	m.syncDetail()

	return nil
}

func (m *gitBrowserScene) syncDetail() {
	key := fmt.Sprintf("%d:%s", m.page, m.active().SelectedID())
	if key == m.detailKey {
		return
	}

	m.detailKey = key
	m.detail.Viewport.GotoTop()

	if m.active().SelectedID() == "" {
		m.detail.Viewport.SetContent("No item selected")
		return
	}

	switch m.page {
	case 0:
		id, section := m.status.SelectedChange()
		for _, diff := range gitDiffs() {
			if diff.FileID == id && diff.Section == section {
				m.detail.SetDiff(diff)
				return
			}
		}

		m.detail.SetDiff(gitui.FileDiff{Path: id, State: gitui.UnavailableDiff, Reason: "content not supplied"})
	case 1:
		for _, commit := range gitHistory().Commits {
			if commit.ID == m.history.SelectedID() {
				m.detail.Viewport.SetContent("Commit " + commit.ID + "\n" + commit.Subject + "\nAuthor: " + commit.Author + "\nParents: " + strings.Join(commit.ParentIDs, ", ") + "\n\nCommit diff was not supplied.")
				return
			}
		}
	case 2:
		for _, branch := range gitBranches() {
			if branch.Ref.Name == m.branches.SelectedID() {
				m.detail.Viewport.SetContent(branch.Ref.Name + "\nTarget: " + branch.Ref.Target + "\nWorktrees: " + strings.Join(branch.WorktreePaths, ", "))
				return
			}
		}
	case 3:
		for _, stash := range gitStashes() {
			if stash.ID == m.stashes.SelectedID() {
				m.detail.SetDiff(stash.Diffs[0])
				return
			}
		}
	case 4:
		for _, w := range gitWorktrees() {
			if w.ID == m.worktrees.SelectedID() {
				m.detail.Viewport.SetContent(w.Path + "\n" + gitui.HeadIndicator(w.Head, m.environment.Config.Palette()) + "\n" + gitui.WorktreeIndicators(w, m.environment.Config.Palette()) + "\n\nEnter: inspect this worktree's supplied context.")
				return
			}
		}
	}
}

func (m *gitBrowserScene) View() string {
	t := m.environment.Config.Palette()
	names := []string{"Status", "History", "Branches", "Stashes", "Worktrees"}

	tabs := make([]string, len(names))
	for i, name := range names {
		if i == m.page {
			tabs[i] = t.Focus().Render(name)
		} else {
			tabs[i] = t.Control().Render(name)
		}
	}

	left, right := m.paneSizes()

	content := m.active().View()
	if m.page == 0 {
		content = m.status.View()
	}

	a := layout.Panel(names[m.page], content, left, t, !m.focusDetail)
	b := layout.Panel("Details", m.detail.View(), right, t, m.focusDetail)

	panes := lipgloss.JoinHorizontal(lipgloss.Top, a, b)
	if m.size.Width < 100 {
		panes = a + "\n" + b
	}

	return layout.Fit(t.Heading().Render("Repository browser · "+m.context)+"\n"+strings.Join(tabs, " ")+"\n"+panes+"\nCtrl+Left/Right section · Tab focus · Enter inspect", m.size)
}
