[Screens](https://lanej.io/charmed/screens.html)git

# [git](https://lanej.io/charmed/screens-git.html#level-screens-git)

1 component · [All screens](https://lanej.io/charmed/screens.html)

- [Repository browser](https://lanej.io/charmed/screens-git.html#git%2fbrowser)

[Screens](https://lanej.io/charmed/screens.html) / git

## [Repository browser](https://lanej.io/charmed/screens-git.html#git%2fbrowser)

Explore status, changes, history, branches, stashes and worktrees from one keyboard-driven example.

### [Native preview](https://lanej.io/charmed/screens-git.html#git%2fbrowser-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
charmed  HEAD → main                                                    ▸ ~/src/charmed  HEAD → main   main · status un…
staged 0 · unstaged 0 · untracked 0 · conflicts 0                                                                       
upstream unavailable · ahead/behind unavailable                                                                         
                                                                                                                        
Working tree clean                                                                                                      
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
↑/↓ j/k · / search · Enter select                                       ↑/↓ j/k · / search · Enter select
```

Demo · truecolor · 120 × 32 terminal cells

```
Repository browser · ~/src/charmed                                                                                      
  Status  History  Branches  Stashes  Worktrees                                                                         
┌ Context ─────────────────┐ ┌ Status ────────────────────────────────────────────────┐ ┌ Details ─────────────────────┐
│ Repository: charmed      │ │ charmed  HEAD → main                                   │ │ gitui/status.go              │
│ Worktree: ~/src/charmed  │ │ staged 2 · unstaged 1 · untracked 1 · conflicts 1      │ │ @@ -0,0 +1,6 @@              │
│ Section: Status          │ │ upstream refs/remotes/origin/main · ahead 2 · behind 1 │ │ 1 +package gitui             │
│ Layout: three-column     │ │                                                        │ │ 2 +                          │
│                          │ │ staged (2)                                             │ │ 3 +// Status belongs to the  │
│                          │ │ ▸  added  gitui/status.go                              │ │ 4 +func label() string {     │
│                          │ │    renamed  docs/git.md → docs/git-components.md       │ │ 5 +    return "ready"        │
│                          │ │ unstaged (1)                                           │ │ 6 +}                         │
│                          │ │    modified  gitui/status.go                           │ │                              │
│                          │ │ untracked (1)                                          │ │                              │
│                          │ │    untracked  notes.txt                                │ │                              │
│                          │ │ conflicts (1)                                          │ │                              │
│                          │ │    conflicted  theme/theme.go · both modified          │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │ ↑/↓ j/k · / search · Enter select                      │ │                              │
│                          │ │                                                        │ │                              │
└──────────────────────────┘ └────────────────────────────────────────────────────────┘ └──────────────────────────────┘
Ctrl+Left/Right section · Tab focus · Alt+L layout · Ctrl+L map · Ctrl+Tab patch focus                      three-column
```

Narrow · truecolor · 40 × 16 terminal cells

```
Repository browser · ~/src/charmed      
  Status  History  Branches  Stashes   >
┌ Status ──────────────────────────────┐
│ charmed  HEAD → main                 │
│ staged 2 · unstaged 1 · untracked 1  │
│ upstream refs/remotes/origin/main ·  │
│                                      │
│ staged (2)                           │
│ ▸  added  gitui/status.go            │
│    renamed  docs/git.md → docs/git-… │
│ unstaged (1)                         │
│    modified  gitui/status.go         │
│ ↑/↓ j/k · / search · Enter select    │
│                                      │
└──────────────────────────────────────┘
Ctrl+Left/Right section · Tab focus · Al
```

256 colors · ansi256 · 120 × 32 terminal cells

```
Repository browser · ~/src/charmed                                                                                      
  Status  History  Branches  Stashes  Worktrees                                                                         
┌ Context ─────────────────┐ ┌ Status ────────────────────────────────────────────────┐ ┌ Details ─────────────────────┐
│ Repository: charmed      │ │ charmed  HEAD → main                                   │ │ gitui/status.go              │
│ Worktree: ~/src/charmed  │ │ staged 2 · unstaged 1 · untracked 1 · conflicts 1      │ │ @@ -0,0 +1,6 @@              │
│ Section: Status          │ │ upstream refs/remotes/origin/main · ahead 2 · behind 1 │ │ 1 +package gitui             │
│ Layout: three-column     │ │                                                        │ │ 2 +                          │
│                          │ │ staged (2)                                             │ │ 3 +// Status belongs to the  │
│                          │ │ ▸  added  gitui/status.go                              │ │ 4 +func label() string {     │
│                          │ │    renamed  docs/git.md → docs/git-components.md       │ │ 5 +    return "ready"        │
│                          │ │ unstaged (1)                                           │ │ 6 +}                         │
│                          │ │    modified  gitui/status.go                           │ │                              │
│                          │ │ untracked (1)                                          │ │                              │
│                          │ │    untracked  notes.txt                                │ │                              │
│                          │ │ conflicts (1)                                          │ │                              │
│                          │ │    conflicted  theme/theme.go · both modified          │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │ ↑/↓ j/k · / search · Enter select                      │ │                              │
│                          │ │                                                        │ │                              │
└──────────────────────────┘ └────────────────────────────────────────────────────────┘ └──────────────────────────────┘
Ctrl+Left/Right section · Tab focus · Alt+L layout · Ctrl+L map · Ctrl+Tab patch focus                      three-column
```

16 colors · ansi16 · 120 × 32 terminal cells

```
Repository browser · ~/src/charmed                                                                                      
  Status  History  Branches  Stashes  Worktrees                                                                         
┌ Context ─────────────────┐ ┌ Status ────────────────────────────────────────────────┐ ┌ Details ─────────────────────┐
│ Repository: charmed      │ │ charmed  HEAD → main                                   │ │ gitui/status.go              │
│ Worktree: ~/src/charmed  │ │ staged 2 · unstaged 1 · untracked 1 · conflicts 1      │ │ @@ -0,0 +1,6 @@              │
│ Section: Status          │ │ upstream refs/remotes/origin/main · ahead 2 · behind 1 │ │ 1 +package gitui             │
│ Layout: three-column     │ │                                                        │ │ 2 +                          │
│                          │ │ staged (2)                                             │ │ 3 +// Status belongs to the  │
│                          │ │ ▸  added  gitui/status.go                              │ │ 4 +func label() string {     │
│                          │ │    renamed  docs/git.md → docs/git-components.md       │ │ 5 +    return "ready"        │
│                          │ │ unstaged (1)                                           │ │ 6 +}                         │
│                          │ │    modified  gitui/status.go                           │ │                              │
│                          │ │ untracked (1)                                          │ │                              │
│                          │ │    untracked  notes.txt                                │ │                              │
│                          │ │ conflicts (1)                                          │ │                              │
│                          │ │    conflicted  theme/theme.go · both modified          │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │                                                        │ │                              │
│                          │ │ ↑/↓ j/k · / search · Enter select                      │ │                              │
│                          │ │                                                        │ │                              │
└──────────────────────────┘ └────────────────────────────────────────────────────────┘ └──────────────────────────────┘
Ctrl+Left/Right section · Tab focus · Alt+L layout · Ctrl+L map · Ctrl+Tab patch focus                      three-column
```

Monochrome · none · 120 × 32 terminal cells

```
Repository browser * ~/src/charmed                                                                                      
  Status  History  Branches  Stashes  Worktrees                                                                         
+ Context -----------------+ + Status ------------------------------------------------+ + Details ---------------------+
| Repository: charmed      | | charmed [HEAD > main]                                  | | gitui/status.go              |
| Worktree: ~/src/charmed  | | staged 2 * unstaged 1 * untracked 1 * conflicts 1      | | @@ -0,0 +1,6 @@              |
| Section: Status          | | upstream refs/remotes/origin/main * ahead 2 * behind 1 | | 1 +package gitui             |
| Layout: three-column     | |                                                        | | 2 +                          |
|                          | | staged (2)                                             | | 3 +// Status belongs to the  |
|                          | | > [added] gitui/status.go                              | | 4 +func label() string {     |
|                          | |   [renamed] docs/git.md > docs/git-components.md       | | 5 +    return "ready"        |
|                          | | unstaged (1)                                           | | 6 +}                         |
|                          | |   [modified] gitui/status.go                           | |                              |
|                          | | untracked (1)                                          | |                              |
|                          | |   [untracked] notes.txt                                | |                              |
|                          | | conflicts (1)                                          | |                              |
|                          | |   [conflicted] theme/theme.go * both modified          | |                              |
|                          | |                                                        | |                              |
|                          | |                                                        | |                              |
|                          | |                                                        | |                              |
|                          | |                                                        | |                              |
|                          | |                                                        | |                              |
|                          | |                                                        | |                              |
|                          | |                                                        | |                              |
|                          | |                                                        | |                              |
|                          | |                                                        | |                              |
|                          | |                                                        | |                              |
|                          | |                                                        | |                              |
|                          | |                                                        | |                              |
|                          | | up/down j/k * / search * Enter select                  | |                              |
|                          | |                                                        | |                              |
+--------------------------+ +--------------------------------------------------------+ +------------------------------+
Ctrl+Left/Right section * Tab focus * Alt+L layout * Ctrl+L map * Ctrl+Tab patch focus                      three-column
```

### [Go usage](https://lanej.io/charmed/screens-git.html#git%2fbrowser-code)

Imports: [`fmt`](https://pkg.go.dev/fmt "fmt") · [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme") · [`tui`](https://pkg.go.dev/github.com/lanej/charmed/tui "github.com/lanej/charmed/tui")

```go
palette := theme.New(theme.Dark)

screen, err := tui.NewScreen("status", "worktrees")
if err != nil {
	return "", err
}

root := &layout.SplitPane{First: &layout.Region{ID: "status", Min: layout.Size{Width: 40, Height: 9}}, Second: &layout.Region{ID: "worktrees", Min: layout.Size{Width: 30, Height: 5}}, Ratio: .6, Gap: 1}
if arrangeErr := screen.Arrange(root, layout.Size{Width: 120, Height: 20}); arrangeErr != nil {
	return "", arrangeErr
}

bounds, visible := screen.Region("status")
if !visible {
	return "", fmt.Errorf("status region is not visible")
}

repository := gitui.Repository{Name: "charmed", Head: gitui.Head{Kind: gitui.Attached, Ref: "refs/heads/main"}}

status, err := gitui.NewStatus(repository, bounds.Dx(), bounds.Dy(), palette)
if err != nil {
	return "", err
}

bounds, visible = screen.Region("worktrees")
if !visible {
	return "", fmt.Errorf("worktrees region is not visible")
}

worktrees, err := gitui.NewWorktreeList([]gitui.Worktree{{ID: "main", Path: "~/src/charmed", Main: true, Head: repository.Head}}, bounds.Dx(), bounds.Dy(), palette)
if err != nil {
	return "", err
}

frame, err := screen.Render(map[string]string{"status": status.View(), "worktrees": worktrees.View()})

return frame.Content, err
```

### [Updates, focus & sizing](https://lanej.io/charmed/screens-git.html#git%2fbrowser-integration)

The parent model supplies observations, coordinates focus and owns application decisions. Forward initialization and update commands rather than executing them during rendering. Resize the existing composition from the terminal allocation so selection remains stable, and apply native customization at the exposed child models after any rebuilding lifecycle step.

### [Run this example](https://lanej.io/charmed/screens-git.html#git%2fbrowser-run)

```sh
go run ./examples/catalog git/browser
```

[Example source](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/docs/snippets/git_browser.go) · [Implementation](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [git/changes](https://lanej.io/charmed/organisms-git.html#git/changes) [git/branches](https://lanej.io/charmed/organisms-git.html#git/branches) [git/history](https://lanej.io/charmed/organisms-git.html#git/history) [git/stashes](https://lanej.io/charmed/organisms-git.html#git/stashes) [git/worktrees](https://lanej.io/charmed/organisms-git.html#git/worktrees) [tui/screen](https://lanej.io/charmed/templates-tui.html#tui/screen)
