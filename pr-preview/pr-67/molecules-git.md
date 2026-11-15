[Molecules](https://lanej.io/charmed/molecules.html)git

# [git](https://lanej.io/charmed/molecules-git.html#level-molecules-git)

11 components · [All molecules](https://lanej.io/charmed/molecules.html)

- [Staged changes](https://lanej.io/charmed/molecules-git.html#git%2fstaged)
- [Unstaged changes](https://lanej.io/charmed/molecules-git.html#git%2funstaged)
- [Untracked changes](https://lanej.io/charmed/molecules-git.html#git%2funtracked)
- [Conflicted changes](https://lanej.io/charmed/molecules-git.html#git%2fconflicts)
- [Repository summary](https://lanej.io/charmed/molecules-git.html#git%2fsummary)
- [File change row](https://lanej.io/charmed/molecules-git.html#git%2ffile-row)
- [Branch row](https://lanej.io/charmed/molecules-git.html#git%2fbranch-row)
- [Commit row](https://lanej.io/charmed/molecules-git.html#git%2fcommit-row)
- [Stash row](https://lanej.io/charmed/molecules-git.html#git%2fstash-row)
- [Worktree row](https://lanej.io/charmed/molecules-git.html#git%2fworktree-row)
- [Diff hunk](https://lanej.io/charmed/molecules-git.html#git%2fhunk)

[Molecules](https://lanej.io/charmed/molecules.html) / git

## [Staged changes](https://lanej.io/charmed/molecules-git.html#git%2fstaged)

Display staged entries, their count and the empty state as a standalone section.

### [Native preview](https://lanej.io/charmed/molecules-git.html#git%2fstaged-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 staged  1
  added  config.go
```

Demo · truecolor · 80 × 10 terminal cells

```
 staged  2                                                                      
  added  gitui/status.go                                                        
  renamed  docs/git.md → docs/git-components.md
```

Narrow · truecolor · 40 × 16 terminal cells

```
 staged  2                              
  added  gitui/status.go                
  renamed  docs/git.md → docs/git-compon
```

256 colors · ansi256 · 80 × 10 terminal cells

```
 staged  2                                                                      
  added  gitui/status.go                                                        
  renamed  docs/git.md → docs/git-components.md
```

16 colors · ansi16 · 80 × 10 terminal cells

```
 staged  2                                                                      
  added  gitui/status.go                                                        
  renamed  docs/git.md → docs/git-components.md
```

Monochrome · none · 80 × 10 terminal cells

```
[staged] 2                                                                      
 [added] gitui/status.go                                                        
 [renamed] docs/git.md > docs/git-components.md
```

### [Go usage](https://lanej.io/charmed/molecules-git.html#git%2fstaged-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

files := []gitui.FileChange{
	{ID: "config", Path: "config.go", Index: gitui.Added, Worktree: gitui.Modified},
	{ID: "notes", Path: "notes.txt", Worktree: gitui.Untracked},
	{ID: "merge", Path: "merge.go", Index: gitui.Conflicted, Conflict: "both modified"},
}

group, err := gitui.NewStagedChanges(files, palette)
if err != nil {
	return "", err
}

return group.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-git.html#git%2fstaged-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/molecules-git.html#git%2fstaged-run)

```sh
go run ./examples/catalog git/staged
```

[Example source](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/git_staged.go) · [Implementation](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [git/staged-state](https://lanej.io/charmed/atoms-git.html#git/staged-state) [git/file-row](https://lanej.io/charmed/molecules-git.html#git/file-row)

[Molecules](https://lanej.io/charmed/molecules.html) / git

## [Unstaged changes](https://lanej.io/charmed/molecules-git.html#git%2funstaged)

Show unstaged file changes alongside their paths and an explicit empty state.

### [Native preview](https://lanej.io/charmed/molecules-git.html#git%2funstaged-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 unstaged  1
  modified  config.go
```

Demo · truecolor · 80 × 10 terminal cells

```
 unstaged  1                                                                    
  modified  gitui/status.go
```

Narrow · truecolor · 40 × 16 terminal cells

```
 unstaged  1                            
  modified  gitui/status.go
```

256 colors · ansi256 · 80 × 10 terminal cells

```
 unstaged  1                                                                    
  modified  gitui/status.go
```

16 colors · ansi16 · 80 × 10 terminal cells

```
 unstaged  1                                                                    
  modified  gitui/status.go
```

Monochrome · none · 80 × 10 terminal cells

```
[unstaged] 1                                                                    
 [modified] gitui/status.go
```

### [Go usage](https://lanej.io/charmed/molecules-git.html#git%2funstaged-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

files := []gitui.FileChange{
	{ID: "config", Path: "config.go", Index: gitui.Added, Worktree: gitui.Modified},
	{ID: "notes", Path: "notes.txt", Worktree: gitui.Untracked},
	{ID: "merge", Path: "merge.go", Index: gitui.Conflicted, Conflict: "both modified"},
}

group, err := gitui.NewUnstagedChanges(files, palette)
if err != nil {
	return "", err
}

return group.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-git.html#git%2funstaged-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/molecules-git.html#git%2funstaged-run)

```sh
go run ./examples/catalog git/unstaged
```

[Example source](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/git_unstaged.go) · [Implementation](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [git/unstaged-state](https://lanej.io/charmed/atoms-git.html#git/unstaged-state) [git/file-row](https://lanej.io/charmed/molecules-git.html#git/file-row)

[Molecules](https://lanej.io/charmed/molecules.html) / git

## [Untracked changes](https://lanej.io/charmed/molecules-git.html#git%2funtracked)

Group untracked paths separately from tracked modifications in the same repository.

### [Native preview](https://lanej.io/charmed/molecules-git.html#git%2funtracked-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 untracked  1
  untracked  notes.txt
```

Demo · truecolor · 80 × 10 terminal cells

```
 untracked  1                                                                   
  untracked  notes.txt
```

Narrow · truecolor · 40 × 16 terminal cells

```
 untracked  1                           
  untracked  notes.txt
```

256 colors · ansi256 · 80 × 10 terminal cells

```
 untracked  1                                                                   
  untracked  notes.txt
```

16 colors · ansi16 · 80 × 10 terminal cells

```
 untracked  1                                                                   
  untracked  notes.txt
```

Monochrome · none · 80 × 10 terminal cells

```
[untracked] 1                                                                   
 [untracked] notes.txt
```

### [Go usage](https://lanej.io/charmed/molecules-git.html#git%2funtracked-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

files := []gitui.FileChange{
	{ID: "config", Path: "config.go", Index: gitui.Added, Worktree: gitui.Modified},
	{ID: "notes", Path: "notes.txt", Worktree: gitui.Untracked},
	{ID: "merge", Path: "merge.go", Index: gitui.Conflicted, Conflict: "both modified"},
}

group, err := gitui.NewUntrackedChanges(files, palette)
if err != nil {
	return "", err
}

return group.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-git.html#git%2funtracked-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/molecules-git.html#git%2funtracked-run)

```sh
go run ./examples/catalog git/untracked
```

[Example source](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/git_untracked.go) · [Implementation](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [git/untracked-state](https://lanej.io/charmed/atoms-git.html#git/untracked-state) [git/file-row](https://lanej.io/charmed/molecules-git.html#git/file-row)

[Molecules](https://lanej.io/charmed/molecules.html) / git

## [Conflicted changes](https://lanej.io/charmed/molecules-git.html#git%2fconflicts)

Group unresolved paths so readers can inspect conflicts independently of other changes.

### [Native preview](https://lanej.io/charmed/molecules-git.html#git%2fconflicts-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 conflicts  1
  conflicted  merge.go · both modified
```

Demo · truecolor · 80 × 10 terminal cells

```
 conflicts  1                                                                   
  conflicted  theme/theme.go · both modified
```

Narrow · truecolor · 40 × 16 terminal cells

```
 conflicts  1                           
  conflicted  theme/theme.go · both modi
```

256 colors · ansi256 · 80 × 10 terminal cells

```
 conflicts  1                                                                   
  conflicted  theme/theme.go · both modified
```

16 colors · ansi16 · 80 × 10 terminal cells

```
 conflicts  1                                                                   
  conflicted  theme/theme.go · both modified
```

Monochrome · none · 80 × 10 terminal cells

```
[conflicts] 1                                                                   
 [conflicted] theme/theme.go * both modified
```

### [Go usage](https://lanej.io/charmed/molecules-git.html#git%2fconflicts-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

files := []gitui.FileChange{
	{ID: "config", Path: "config.go", Index: gitui.Added, Worktree: gitui.Modified},
	{ID: "notes", Path: "notes.txt", Worktree: gitui.Untracked},
	{ID: "merge", Path: "merge.go", Index: gitui.Conflicted, Conflict: "both modified"},
}

group, err := gitui.NewConflictChanges(files, palette)
if err != nil {
	return "", err
}

return group.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-git.html#git%2fconflicts-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/molecules-git.html#git%2fconflicts-run)

```sh
go run ./examples/catalog git/conflicts
```

[Example source](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/git_conflicts.go) · [Implementation](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [git/conflict-state](https://lanej.io/charmed/atoms-git.html#git/conflict-state) [git/file-row](https://lanej.io/charmed/molecules-git.html#git/file-row)

[Molecules](https://lanej.io/charmed/molecules.html) / git

## [Repository summary](https://lanej.io/charmed/molecules-git.html#git%2fsummary)

Combine the current HEAD, divergence and file counts in a compact repository summary.

### [Native preview](https://lanej.io/charmed/molecules-git.html#git%2fsummary-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
charmed  HEAD → main 
staged 0 · unstaged 0 · untracked 0 · conflicts 0
upstream unavailable · ahead/behind unavailable
```

Demo · truecolor · 120 × 32 terminal cells

```
charmed  HEAD → main                                                                                                    
staged 2 · unstaged 1 · untracked 1 · conflicts 1                                                                       
upstream refs/remotes/origin/main · ahead 2 · behind 1
```

Narrow · truecolor · 40 × 16 terminal cells

```
charmed  HEAD → main                    
staged 2 · unstaged 1 · untracked 1 · co
upstream refs/remotes/origin/main · ahea
```

256 colors · ansi256 · 120 × 32 terminal cells

```
charmed  HEAD → main                                                                                                    
staged 2 · unstaged 1 · untracked 1 · conflicts 1                                                                       
upstream refs/remotes/origin/main · ahead 2 · behind 1
```

16 colors · ansi16 · 120 × 32 terminal cells

```
charmed  HEAD → main                                                                                                    
staged 2 · unstaged 1 · untracked 1 · conflicts 1                                                                       
upstream refs/remotes/origin/main · ahead 2 · behind 1
```

Monochrome · none · 120 × 32 terminal cells

```
charmed [HEAD > main]                                                                                                   
staged 2 * unstaged 1 * untracked 1 * conflicts 1                                                                       
upstream refs/remotes/origin/main * ahead 2 * behind 1
```

### [Go usage](https://lanej.io/charmed/molecules-git.html#git%2fsummary-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return gitui.RepositorySummary(gitui.Repository{Name: "charmed", Head: gitui.Head{Kind: gitui.Attached, Ref: "refs/heads/main"}}, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-git.html#git%2fsummary-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/molecules-git.html#git%2fsummary-run)

```sh
go run ./examples/catalog git/summary
```

[Example source](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/git_summary.go) · [Implementation](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [git/head](https://lanej.io/charmed/atoms-git.html#git/head) [git/divergence](https://lanej.io/charmed/atoms-git.html#git/divergence) [git/file-state](https://lanej.io/charmed/atoms-git.html#git/file-state)

[Molecules](https://lanej.io/charmed/molecules.html) / git

## [File change row](https://lanej.io/charmed/molecules-git.html#git%2ffile-row)

Pair a file path with its change indicator for use in a list or status section.

### [Native preview](https://lanej.io/charmed/molecules-git.html#git%2ffile-row-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 added  config.go
```

Demo · truecolor · 120 × 32 terminal cells

```
 added  gitui/status.go                                                                                                 
                                                                                                                        
 modified  gitui/status.go
```

Narrow · truecolor · 40 × 16 terminal cells

```
 added  gitui/status.go                 
                                        
 modified  gitui/status.go
```

256 colors · ansi256 · 120 × 32 terminal cells

```
 added  gitui/status.go                                                                                                 
                                                                                                                        
 modified  gitui/status.go
```

16 colors · ansi16 · 120 × 32 terminal cells

```
 added  gitui/status.go                                                                                                 
                                                                                                                        
 modified  gitui/status.go
```

Monochrome · none · 120 × 32 terminal cells

```
[added] gitui/status.go                                                                                                 
                                                                                                                        
[modified] gitui/status.go
```

### [Go usage](https://lanej.io/charmed/molecules-git.html#git%2ffile-row-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

file := gitui.FileChange{ID: "config", Path: "config.go", Index: gitui.Added, Worktree: gitui.Modified}

return gitui.FileChangeRow(file, gitui.Staged, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-git.html#git%2ffile-row-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/molecules-git.html#git%2ffile-row-run)

```sh
go run ./examples/catalog git/file-row
```

[Example source](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/git_file_row.go) · [Implementation](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [git/file-state](https://lanej.io/charmed/atoms-git.html#git/file-state)

[Molecules](https://lanej.io/charmed/molecules.html) / git

## [Branch row](https://lanej.io/charmed/molecules-git.html#git%2fbranch-row)

Compare a branch with its upstream and mark the current branch in a list.

### [Native preview](https://lanej.io/charmed/molecules-git.html#git%2fbranch-row-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 branch main  [HEAD] · upstream unavailable · ahead/behind unavailable
```

Demo · truecolor · 120 × 32 terminal cells

```
 branch main  [HEAD] · upstream refs/remotes/origin/main · ahead 2 · behind 1 · worktrees ~/src/charmed
```

Narrow · truecolor · 40 × 16 terminal cells

```
 branch main  [HEAD] · upstream refs/rem
```

256 colors · ansi256 · 120 × 32 terminal cells

```
 branch main  [HEAD] · upstream refs/remotes/origin/main · ahead 2 · behind 1 · worktrees ~/src/charmed
```

16 colors · ansi16 · 120 × 32 terminal cells

```
 branch main  [HEAD] · upstream refs/remotes/origin/main · ahead 2 · behind 1 · worktrees ~/src/charmed
```

Monochrome · none · 120 × 32 terminal cells

```
[branch main] [HEAD] * upstream refs/remotes/origin/main * ahead 2 * behind 1 * worktrees ~/src/charmed
```

### [Go usage](https://lanej.io/charmed/molecules-git.html#git%2fbranch-row-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

branch := gitui.Branch{Ref: gitui.Ref{Name: "refs/heads/main", Kind: gitui.LocalBranch}, Current: true}

return gitui.BranchRow(branch, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-git.html#git%2fbranch-row-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/molecules-git.html#git%2fbranch-row-run)

```sh
go run ./examples/catalog git/branch-row
```

[Example source](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/git_branch_row.go) · [Implementation](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [git/ref](https://lanej.io/charmed/atoms-git.html#git/ref) [git/divergence](https://lanej.io/charmed/atoms-git.html#git/divergence)

[Molecules](https://lanej.io/charmed/molecules.html) / git

## [Commit row](https://lanej.io/charmed/molecules-git.html#git%2fcommit-row)

Show a commit and its associated refs together in one history row.

### [Native preview](https://lanej.io/charmed/molecules-git.html#git%2fcommit-row-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
a31f2c7819 Add search · Josh
```

Demo · truecolor · 120 × 32 terminal cells

```
a31f2c7819  HEAD → main   branch main  Merge search · Josh · 2026-11-11 12:00
```

Narrow · truecolor · 40 × 16 terminal cells

```
a31f2c7819  HEAD → main   branch main  M
```

256 colors · ansi256 · 120 × 32 terminal cells

```
a31f2c7819  HEAD → main   branch main  Merge search · Josh · 2026-11-11 12:00
```

16 colors · ansi16 · 120 × 32 terminal cells

```
a31f2c7819  HEAD → main   branch main  Merge search · Josh · 2026-11-11 12:00
```

Monochrome · none · 120 × 32 terminal cells

```
a31f2c7819 [HEAD > main] [branch main] Merge search * Josh * 2026-11-11 12:00
```

### [Go usage](https://lanej.io/charmed/molecules-git.html#git%2fcommit-row-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

commit := gitui.Commit{ID: "a31f2c7819be822", Subject: "Add search", Author: "Josh"}

return gitui.CommitRow(commit, nil, gitui.Head{}, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-git.html#git%2fcommit-row-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/molecules-git.html#git%2fcommit-row-run)

```sh
go run ./examples/catalog git/commit-row
```

[Example source](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/git_commit_row.go) · [Implementation](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [git/commit-id](https://lanej.io/charmed/atoms-git.html#git/commit-id) [git/ref](https://lanej.io/charmed/atoms-git.html#git/ref) [git/head](https://lanej.io/charmed/atoms-git.html#git/head)

[Molecules](https://lanej.io/charmed/molecules.html) / git

## [Stash row](https://lanej.io/charmed/molecules-git.html#git%2fstash-row)

Display a stash entry while retaining its stable ID for later selection.

### [Native preview](https://lanej.io/charmed/molecules-git.html#git%2fstash-row-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
stash@{0} WIP: search · main · 
```

Demo · truecolor · 120 × 32 terminal cells

```
stash@{0} WIP: status presentation · refs/heads/main · a31f2c7819
```

Narrow · truecolor · 40 × 16 terminal cells

```
stash@{0} WIP: status presentation · ref
```

256 colors · ansi256 · 120 × 32 terminal cells

```
stash@{0} WIP: status presentation · refs/heads/main · a31f2c7819
```

16 colors · ansi16 · 120 × 32 terminal cells

```
stash@{0} WIP: status presentation · refs/heads/main · a31f2c7819
```

Monochrome · none · 120 × 32 terminal cells

```
stash@{0} WIP: status presentation * refs/heads/main * a31f2c7819
```

### [Go usage](https://lanej.io/charmed/molecules-git.html#git%2fstash-row-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

stash := gitui.Stash{ID: "stash-object", Ordinal: 0, Subject: "WIP: search", Branch: "main"}

return gitui.StashRow(stash, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-git.html#git%2fstash-row-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/molecules-git.html#git%2fstash-row-run)

```sh
go run ./examples/catalog git/stash-row
```

[Example source](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/git_stash_row.go) · [Implementation](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [git/commit-id](https://lanej.io/charmed/atoms-git.html#git/commit-id)

[Molecules](https://lanej.io/charmed/molecules.html) / git

## [Worktree row](https://lanej.io/charmed/molecules-git.html#git%2fworktree-row)

Show a checkout path, its HEAD and the supplied worktree state in one row.

### [Native preview](https://lanej.io/charmed/molecules-git.html#git%2fworktree-row-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
~/src/charmed-search  HEAD → search   linked · locked: review · status unavailable 
```

Demo · truecolor · 120 × 32 terminal cells

```
~/src/charmed-search  HEAD → search   linked · locked: review in progress · clean 
```

Narrow · truecolor · 40 × 16 terminal cells

```
~/src/charmed-search  HEAD → search   li
```

256 colors · ansi256 · 120 × 32 terminal cells

```
~/src/charmed-search  HEAD → search   linked · locked: review in progress · clean 
```

16 colors · ansi16 · 120 × 32 terminal cells

```
~/src/charmed-search  HEAD → search   linked · locked: review in progress · clean 
```

Monochrome · none · 120 × 32 terminal cells

```
~/src/charmed-search [HEAD > search] [linked * locked: review in progress * clean]
```

### [Go usage](https://lanej.io/charmed/molecules-git.html#git%2fworktree-row-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

worktree := gitui.Worktree{ID: "search", Path: "~/src/charmed-search", Head: gitui.Head{Kind: gitui.Attached, Ref: "refs/heads/search"}, Locked: true, LockReason: "review"}

return gitui.WorktreeRow(worktree, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-git.html#git%2fworktree-row-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/molecules-git.html#git%2fworktree-row-run)

```sh
go run ./examples/catalog git/worktree-row
```

[Example source](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/git_worktree_row.go) · [Implementation](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [git/head](https://lanej.io/charmed/atoms-git.html#git/head) [git/worktree-state](https://lanej.io/charmed/atoms-git.html#git/worktree-state)

[Molecules](https://lanej.io/charmed/molecules.html) / git

## [Diff hunk](https://lanej.io/charmed/molecules-git.html#git%2fhunk)

Render supplied hunk lines with syntax colors, line numbers and inline change emphasis. Select a language through DiffOptions.

### [Native preview](https://lanej.io/charmed/molecules-git.html#git%2fhunk-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
@@ -0,0 +1 @@
1 +package gitui
```

Demo · truecolor · 120 × 32 terminal cells

```
@@ -3,5 +3,5 @@                                                                                                         
3  // Status belongs to the supplied worktree.                                                                          
4  func label() string {                                                                                                
5 -    return "ready"                                                                                                   
5 +    return "changed"                                                                                                 
6  }                                                                                                                    
7 -const limit = 10                                                                                                     
7 +const limit = 25                                                                                                     
\ No newline at end of file
```

Narrow · truecolor · 40 × 16 terminal cells

```
@@ -3,5 +3,5 @@                         
3  // Status belongs to the supplied wor
4  func label() string {                
5 -    return "ready"                   
5 +    return "changed"                 
6  }                                    
7 -const limit = 10                     
7 +const limit = 25                     
\ No newline at end of file
```

256 colors · ansi256 · 120 × 32 terminal cells

```
@@ -3,5 +3,5 @@                                                                                                         
3  // Status belongs to the supplied worktree.                                                                          
4  func label() string {                                                                                                
5 -    return "ready"                                                                                                   
5 +    return "changed"                                                                                                 
6  }                                                                                                                    
7 -const limit = 10                                                                                                     
7 +const limit = 25                                                                                                     
\ No newline at end of file
```

16 colors · ansi16 · 120 × 32 terminal cells

```
@@ -3,5 +3,5 @@                                                                                                         
3  // Status belongs to the supplied worktree.                                                                          
4  func label() string {                                                                                                
5 -    return "ready"                                                                                                   
5 +    return "changed"                                                                                                 
6  }                                                                                                                    
7 -const limit = 10                                                                                                     
7 +const limit = 25                                                                                                     
\ No newline at end of file
```

Monochrome · none · 120 × 32 terminal cells

```
@@ -3,5 +3,5 @@                                                                                                         
3  // Status belongs to the supplied worktree.                                                                          
4  func label() string {                                                                                                
5 -    return "ready"                                                                                                   
5 +    return "changed"                                                                                                 
6  }                                                                                                                    
7 -const limit = 10                                                                                                     
7 +const limit = 25                                                                                                     
\ No newline at end of file
```

### [Go usage](https://lanej.io/charmed/molecules-git.html#git%2fhunk-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

hunk := gitui.Hunk{Header: "@@ -0,0 +1 @@", Lines: []gitui.DiffLine{{Kind: gitui.AddedLine, NewLine: 1, Text: "package gitui"}}}

return gitui.DiffHunk(hunk, palette, gitui.DiffOptions{Language: "go"})
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-git.html#git%2fhunk-integration)

Supply a Hunk and optional DiffOptions language; unknown languages render plain syntax. Old and new lines lex separately from the supplied context. Replacement lines pair by position inside contiguous change blocks, with stronger surfaces for changed graphemes. Rendering does not mutate the observation. Re-render after theme changes; place long lines in a viewport. Tabs and terminal controls remain visible as escapes. Monochrome retains signs and bold underlining for inline changes.

### [Run this example](https://lanej.io/charmed/molecules-git.html#git%2fhunk-run)

```sh
go run ./examples/catalog git/hunk
```

[Example source](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/git_hunk.go) · [Implementation](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.
