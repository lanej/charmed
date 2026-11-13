package story

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/interaction"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/navigation"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/tui"
)

type gitBrowserScene struct {
	primarySize layout.Size
	detailSize  layout.Size
	screen      *tui.Screen
	tabs        navigation.Tabs
	layoutMode  int
	showRegions bool
	layoutErr   error
	environment Environment
	status      *gitui.Status
	history     *gitui.CommitHistory
	branches    *gitui.BranchList
	stashes     *gitui.StashList
	worktrees   *gitui.WorktreeList
	patch       *gitui.PatchBrowser
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

	screen, err := tui.NewScreen("header", "tabs", "context", "main", "details", "footer")
	if err != nil {
		return nil, err
	}

	tabs := navigation.NewTabs([]navigation.Tab{{ID: "Status", Label: "Status"}, {ID: "History", Label: "History"}, {ID: "Branches", Label: "Branches"}, {ID: "Stashes", Label: "Stashes"}, {ID: "Worktrees", Label: "Worktrees"}}, t)
	m := &gitBrowserScene{screen: screen, tabs: tabs, environment: e, status: status, history: log, branches: branchList, stashes: stashList, worktrees: worktreeList, detail: gitui.NewDiffViewer(gitui.FileDiff{State: gitui.UnavailableDiff}, 60, 24, t), context: "~/src/charmed"}

	patch, err := gitui.NewPatchBrowser(gitui.PatchObservation{}, e.Config.Size(), t)
	if err != nil {
		return nil, err
	}

	m.patch = patch
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

	m.layoutErr = m.screen.Arrange(gitBrowserLayout(m.focusDetail, m.layoutMode), size)
	if m.layoutErr != nil {
		return
	}

	m.tabs.SetWidth(size.Width)

	if bounds, visible := m.screen.Region("main"); visible {
		child := layout.PanelContentSize(layout.Size{Width: bounds.Dx(), Height: bounds.Dy()})
		if child != m.primarySize {
			m.status.Resize(child)

			for _, list := range []*gitui.List{m.history.List, m.branches.List, m.stashes.List, m.worktrees.List} {
				list.Resize(child)
			}

			m.primarySize = child
		}
	}

	if bounds, visible := m.screen.Region("details"); visible {
		child := layout.PanelContentSize(layout.Size{Width: bounds.Dx(), Height: bounds.Dy()})
		if child != m.detailSize {
			m.detail.Resize(child)
			m.patch.Resize(child)
			m.detailSize = child
		}
	}
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
		case "ctrl+tab":
			if m.page == 3 && m.focusDetail {
				m.patch.FocusDetail(!m.patch.DetailFocused())
				return nil
			}
		case "alt+l":
			m.layoutMode = (m.layoutMode + 1) % 3
		case "ctrl+l":
			m.showRegions = !m.showRegions
			return nil
		default:
			if m.focusDetail {
				if m.page == 3 {
					return m.patch.Update(msg)
				}

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
			if m.page == 3 {
				cmd = m.patch.Update(msg)
			} else {
				cmd = m.detail.Update(msg)
			}
		} else if m.page == 0 {
			cmd = m.status.Update(msg)
		} else {
			cmd = m.active().Update(msg)
		}

		m.syncDetail()

		return cmd
	}

	m.patch.FocusDetail(false)
	m.Resize(m.size)

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
				if err := m.patch.SetPatch(gitui.PatchObservation{Available: true, Files: stash.Diffs}); err != nil {
					m.detail.Viewport.SetContent(err.Error())
				}

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
	if m.layoutErr != nil {
		return layout.Fit(t.Text().Foreground(t.Danger).Render(m.layoutErr.Error()), m.size)
	}

	if m.showRegions {
		return m.screen.RegionMap(t)
	}

	names := []string{"Status", "History", "Branches", "Stashes", "Worktrees"}
	m.tabs.Select(names[m.page])

	views := map[string]string{
		"header": layout.Header{Title: "Repository browser · " + m.context}.View(m.size.Width, t),
		"tabs":   m.tabs.View(),
		"footer": layout.Footer{Help: "Ctrl+Left/Right section · Tab focus · Alt+L layout · Ctrl+L map · Ctrl+Tab patch focus", Status: m.screen.Plan().Variant()}.View(m.size.Width, t),
	}
	if bounds, visible := m.screen.Region("main"); visible {
		content := m.active().View()
		if m.page == 0 {
			content = m.status.View()
		}

		views["main"] = layout.Panel(names[m.page], content, layout.Size{Width: bounds.Dx(), Height: bounds.Dy()}, t, !m.focusDetail)
	}

	if bounds, visible := m.screen.Region("details"); visible {
		detail := m.detail.View()
		if m.page == 3 {
			detail = m.patch.View()
		}

		views["details"] = layout.Panel("Details", detail, layout.Size{Width: bounds.Dx(), Height: bounds.Dy()}, t, m.focusDetail)
	}

	if bounds, visible := m.screen.Region("context"); visible {
		size := layout.Size{Width: bounds.Dx(), Height: bounds.Dy()}
		content := layout.KeyValue([]layout.Pair{{Label: "Repository", Value: "charmed"}, {Label: "Worktree", Value: m.context}, {Label: "Section", Value: names[m.page]}, {Label: "Layout", Value: m.screen.Plan().Variant()}}, layout.PanelContentSize(size).Width, t)
		views["context"] = layout.Panel("Context", content, size, t, false)
	}

	frame, err := m.screen.Render(views)
	if err != nil {
		return layout.Fit(t.Text().Foreground(t.Danger).Render(err.Error()), m.size)
	}

	return frame.Content
}
