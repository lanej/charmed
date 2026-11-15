[Organisms](https://lanej.io/charmed/organisms.html)git

# [git](https://lanej.io/charmed/organisms-git.html#level-organisms-git)

8 components · [All organisms](https://lanej.io/charmed/organisms.html)

- [Repository status](https://lanej.io/charmed/organisms-git.html#git%2fstatus)
- [Branch list](https://lanej.io/charmed/organisms-git.html#git%2fbranches)
- [Branching commit history](https://lanej.io/charmed/organisms-git.html#git%2fhistory)
- [Diff viewer](https://lanej.io/charmed/organisms-git.html#git%2fdiff)
- [Changes inspector](https://lanej.io/charmed/organisms-git.html#git%2fchanges)
- [Stash list](https://lanej.io/charmed/organisms-git.html#git%2fstashes)
- [Worktree list](https://lanej.io/charmed/organisms-git.html#git%2fworktrees)
- [Patch browser](https://lanej.io/charmed/organisms-git.html#git%2fpatch-browser)

[Organisms](https://lanej.io/charmed/organisms.html) / git

## [Repository status](https://lanej.io/charmed/organisms-git.html#git%2fstatus)

Inspect staged, unstaged, untracked and conflicting files within the same repository observation.

### [Native preview](https://lanej.io/charmed/organisms-git.html#git%2fstatus-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
charmed  HEAD unavailable                                                       
staged 1 · unstaged 0 · untracked 0 · conflicts 0                               
upstream unavailable · ahead/behind unavailable                                 
                                                                                
staged (1)                                                                      
▸  added  config.go                                                             
                                                                                
                                                                                
                                                                                
                                                                                
                                                                                
                                                                                
                                                                                
                                                                                
↑/↓ j/k · / search · Enter select
```

Demo · truecolor · 120 × 32 terminal cells

```
charmed  HEAD → main                                                                                                    
staged 2 · unstaged 1 · untracked 1 · conflicts 1                                                                       
upstream refs/remotes/origin/main · ahead 2 · behind 1                                                                  
                                                                                                                        
staged (2)                                                                                                              
▸  added  gitui/status.go                                                                                               
   renamed  docs/git.md → docs/git-components.md                                                                        
unstaged (1)                                                                                                            
   modified  gitui/status.go                                                                                            
untracked (1)                                                                                                           
   untracked  notes.txt                                                                                                 
conflicts (1)                                                                                                           
   conflicted  theme/theme.go · both modified                                                                           
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
↑/↓ j/k · / search · Enter select
```

Narrow · truecolor · 40 × 16 terminal cells

```
charmed  HEAD → main                    
staged 2 · unstaged 1 · untracked 1 · co
upstream refs/remotes/origin/main · ahea
                                        
staged (2)                              
▸  added  gitui/status.go               
   renamed  docs/git.md → docs/git-comp…
unstaged (1)                            
   modified  gitui/status.go            
untracked (1)                           
   untracked  notes.txt                 
conflicts (1)                           
   conflicted  theme/theme.go · both mo…
                                        
↑/↓ j/k · / search · Enter select
```

256 colors · ansi256 · 120 × 32 terminal cells

```
charmed  HEAD → main                                                                                                    
staged 2 · unstaged 1 · untracked 1 · conflicts 1                                                                       
upstream refs/remotes/origin/main · ahead 2 · behind 1                                                                  
                                                                                                                        
staged (2)                                                                                                              
▸  added  gitui/status.go                                                                                               
   renamed  docs/git.md → docs/git-components.md                                                                        
unstaged (1)                                                                                                            
   modified  gitui/status.go                                                                                            
untracked (1)                                                                                                           
   untracked  notes.txt                                                                                                 
conflicts (1)                                                                                                           
   conflicted  theme/theme.go · both modified                                                                           
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
↑/↓ j/k · / search · Enter select
```

16 colors · ansi16 · 120 × 32 terminal cells

```
charmed  HEAD → main                                                                                                    
staged 2 · unstaged 1 · untracked 1 · conflicts 1                                                                       
upstream refs/remotes/origin/main · ahead 2 · behind 1                                                                  
                                                                                                                        
staged (2)                                                                                                              
▸  added  gitui/status.go                                                                                               
   renamed  docs/git.md → docs/git-components.md                                                                        
unstaged (1)                                                                                                            
   modified  gitui/status.go                                                                                            
untracked (1)                                                                                                           
   untracked  notes.txt                                                                                                 
conflicts (1)                                                                                                           
   conflicted  theme/theme.go · both modified                                                                           
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
↑/↓ j/k · / search · Enter select
```

Monochrome · none · 120 × 32 terminal cells

```
charmed [HEAD > main]                                                                                                   
staged 2 * unstaged 1 * untracked 1 * conflicts 1                                                                       
upstream refs/remotes/origin/main * ahead 2 * behind 1                                                                  
                                                                                                                        
staged (2)                                                                                                              
> [added] gitui/status.go                                                                                               
  [renamed] docs/git.md > docs/git-components.md                                                                        
unstaged (1)                                                                                                            
  [modified] gitui/status.go                                                                                            
untracked (1)                                                                                                           
  [untracked] notes.txt                                                                                                 
conflicts (1)                                                                                                           
  [conflicted] theme/theme.go * both modified                                                                           
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
up/down j/k * / search * Enter select
```

### [Go usage](https://lanej.io/charmed/organisms-git.html#git%2fstatus-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

view, err := gitui.NewStatus(gitui.Repository{Name: "charmed", Files: []gitui.FileChange{{ID: "config", Path: "config.go", Index: gitui.Added}}}, 80, 16, palette)
if err != nil {
	return "", err
}

return view.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-git.html#git%2fstatus-integration)

Supply observations from the application and resize the model to its allocated terminal cells. Route keyboard input to the focused view and return its command. Refresh through the public data setter and use stable identities when acting on a selection.

### [Run this example](https://lanej.io/charmed/organisms-git.html#git%2fstatus-run)

```sh
go run ./examples/catalog git/status
```

[Example source](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/docs/snippets/git_status.go) · [Implementation](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [git/summary](https://lanej.io/charmed/molecules-git.html#git/summary) [git/staged](https://lanej.io/charmed/molecules-git.html#git/staged) [git/unstaged](https://lanej.io/charmed/molecules-git.html#git/unstaged) [git/untracked](https://lanej.io/charmed/molecules-git.html#git/untracked) [git/conflicts](https://lanej.io/charmed/molecules-git.html#git/conflicts)

[Organisms](https://lanej.io/charmed/organisms.html) / git

## [Branch list](https://lanej.io/charmed/organisms-git.html#git%2fbranches)

Search and select a branch while retaining its full ref identity across refreshes.

### [Native preview](https://lanej.io/charmed/organisms-git.html#git%2fbranches-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
▸  branch main  [HEAD] · upstream unavailable · ahead/behind unavailable        
                                                                                
                                                                                
                                                                                
                                                                                
                                                                                
                                                                                
                                                                                
                                                                                
                                                                                
↑/↓ j/k · / search · Enter select
```

Demo · truecolor · 120 × 32 terminal cells

```
▸  branch main  [HEAD] · upstream refs/remotes/origin/main · ahead 2 · behind 1 · worktrees ~/src/charmed               
   branch search  · upstream unavailable · ahead/behind unavailable · worktrees ~/src/charmed-search                    
   remote origin/main  · upstream unavailable · ahead/behind unavailable                                                
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
↑/↓ j/k · / search · Enter select
```

Narrow · truecolor · 40 × 16 terminal cells

```
▸  branch main  [HEAD] · upstream refs/…
   branch search  · upstream unavailabl…
   remote origin/main  · upstream unava…
                                        
                                        
                                        
                                        
                                        
                                        
                                        
                                        
                                        
                                        
                                        
↑/↓ j/k · / search · Enter select
```

256 colors · ansi256 · 120 × 32 terminal cells

```
▸  branch main  [HEAD] · upstream refs/remotes/origin/main · ahead 2 · behind 1 · worktrees ~/src/charmed               
   branch search  · upstream unavailable · ahead/behind unavailable · worktrees ~/src/charmed-search                    
   remote origin/main  · upstream unavailable · ahead/behind unavailable                                                
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
↑/↓ j/k · / search · Enter select
```

16 colors · ansi16 · 120 × 32 terminal cells

```
▸  branch main  [HEAD] · upstream refs/remotes/origin/main · ahead 2 · behind 1 · worktrees ~/src/charmed               
   branch search  · upstream unavailable · ahead/behind unavailable · worktrees ~/src/charmed-search                    
   remote origin/main  · upstream unavailable · ahead/behind unavailable                                                
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
↑/↓ j/k · / search · Enter select
```

Monochrome · none · 120 × 32 terminal cells

```
> [branch main] [HEAD] * upstream refs/remotes/origin/main * ahead 2 * behind 1 * worktrees ~/src/charmed               
  [branch search] * upstream unavailable * ahead/behind unavailable * worktrees ~/src/charmed-search                    
  [remote origin/main] * upstream unavailable * ahead/behind unavailable                                                
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
up/down j/k * / search * Enter select
```

### [Go usage](https://lanej.io/charmed/organisms-git.html#git%2fbranches-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

view, err := gitui.NewBranchList([]gitui.Branch{{Ref: gitui.Ref{Name: "refs/heads/main", Kind: gitui.LocalBranch}, Current: true}}, 80, 12, palette)
if err != nil {
	return "", err
}

return view.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-git.html#git%2fbranches-integration)

Supply observations from the application and resize the model to its allocated terminal cells. Route keyboard input to the focused view and return its command. Refresh through the public data setter and use stable identities when acting on a selection.

### [Run this example](https://lanej.io/charmed/organisms-git.html#git%2fbranches-run)

```sh
go run ./examples/catalog git/branches
```

[Example source](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/docs/snippets/git_branches.go) · [Implementation](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [git/branch-row](https://lanej.io/charmed/molecules-git.html#git/branch-row)

[Organisms](https://lanej.io/charmed/organisms.html) / git

## [Branching commit history](https://lanej.io/charmed/organisms-git.html#git%2fhistory)

Follow parent lanes through a supplied commit graph, search commits and scroll continuation rows.

### [Native preview](https://lanej.io/charmed/organisms-git.html#git%2fhistory-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
▸ ●  merge Merge topic                                                          
  │                                                                             
  ├─┐                                                                           
  │ │                                                                           
  │ ●  topic Add search                                                         
  ● │  main Update guide                                                        
  │ │                                                                           
  └─┤                                                                           
    │                                                                           
    ●  base Initial components                                                  
                                                                                
                                                                                
                                                                                
                                                                                
↑/↓ j/k · / search · Enter select
```

Demo · truecolor · 120 × 32 terminal cells

```
▸ ●  a31f2c7819  HEAD → main   branch main  Merge search · Josh · 2026-11-11 12:00                                      
  │                                                                                                                     
  ├─┐                                                                                                                   
  │ │                                                                                                                   
  │ ●  92bc1820ae  branch search  Add search input · Josh                                                               
  │ ●  870ada0912 Filter results · Josh                                                                                 
  ● │  51dd20c991  remote origin/main  Update guide · Josh                                                              
  │ │                                                                                                                   
  └─┤                                                                                                                   
    │                                                                                                                   
    ●  28c341aad0  tag v0.1  Initial components · Josh                                                                  
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
↑/↓ j/k · / search · Enter select
```

Narrow · truecolor · 40 × 16 terminal cells

```
▸ ●  a31f2c7819  HEAD → main   branch m…
  │                                     
  ├─┐                                   
  │ │                                   
  │ ●  92bc1820ae  branch search  Add s…
  │ ●  870ada0912 Filter results · Josh 
  ● │  51dd20c991  remote origin/main  …
  │ │                                   
  └─┤                                   
    │                                   
    ●  28c341aad0  tag v0.1  Initial co…
                                        
                                        
                                        
↑/↓ j/k · / search · Enter select
```

256 colors · ansi256 · 120 × 32 terminal cells

```
▸ ●  a31f2c7819  HEAD → main   branch main  Merge search · Josh · 2026-11-11 12:00                                      
  │                                                                                                                     
  ├─┐                                                                                                                   
  │ │                                                                                                                   
  │ ●  92bc1820ae  branch search  Add search input · Josh                                                               
  │ ●  870ada0912 Filter results · Josh                                                                                 
  ● │  51dd20c991  remote origin/main  Update guide · Josh                                                              
  │ │                                                                                                                   
  └─┤                                                                                                                   
    │                                                                                                                   
    ●  28c341aad0  tag v0.1  Initial components · Josh                                                                  
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
↑/↓ j/k · / search · Enter select
```

16 colors · ansi16 · 120 × 32 terminal cells

```
▸ ●  a31f2c7819  HEAD → main   branch main  Merge search · Josh · 2026-11-11 12:00                                      
  │                                                                                                                     
  ├─┐                                                                                                                   
  │ │                                                                                                                   
  │ ●  92bc1820ae  branch search  Add search input · Josh                                                               
  │ ●  870ada0912 Filter results · Josh                                                                                 
  ● │  51dd20c991  remote origin/main  Update guide · Josh                                                              
  │ │                                                                                                                   
  └─┤                                                                                                                   
    │                                                                                                                   
    ●  28c341aad0  tag v0.1  Initial components · Josh                                                                  
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
↑/↓ j/k · / search · Enter select
```

Monochrome · none · 120 × 32 terminal cells

```
> *  a31f2c7819 [HEAD > main] [branch main] Merge search * Josh * 2026-11-11 12:00                                      
  |                                                                                                                     
  +-+                                                                                                                   
  | |                                                                                                                   
  | *  92bc1820ae [branch search] Add search input * Josh                                                               
  | *  870ada0912 Filter results * Josh                                                                                 
  * |  51dd20c991 [remote origin/main] Update guide * Josh                                                              
  | |                                                                                                                   
  +-+                                                                                                                   
    |                                                                                                                   
    *  28c341aad0 [tag v0.1] Initial components * Josh                                                                  
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
up/down j/k * / search * Enter select
```

### [Go usage](https://lanej.io/charmed/organisms-git.html#git%2fhistory-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

view, err := gitui.NewCommitHistory(gitui.History{Commits: []gitui.Commit{
	{ID: "merge", ParentIDs: []string{"main", "topic"}, Subject: "Merge topic"},
	{ID: "topic", ParentIDs: []string{"base"}, Subject: "Add search"},
	{ID: "main", ParentIDs: []string{"base"}, Subject: "Update guide"},
	{ID: "base", Subject: "Initial components"},
}}, 80, 16, palette)
if err != nil {
	return "", err
}

return view.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-git.html#git%2fhistory-integration)

Supply observations from the application and resize the model to its allocated terminal cells. Route keyboard input to the focused view and return its command. Refresh through the public data setter and use stable identities when acting on a selection.

### [Run this example](https://lanej.io/charmed/organisms-git.html#git%2fhistory-run)

```sh
go run ./examples/catalog git/history
```

[Example source](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/docs/snippets/git_history.go) · [Implementation](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [git/commit-row](https://lanej.io/charmed/molecules-git.html#git/commit-row)

[Organisms](https://lanej.io/charmed/organisms.html) / git

## [Diff viewer](https://lanej.io/charmed/organisms-git.html#git%2fdiff)

Inspect syntax-aware file diffs with change backgrounds and inline emphasis; binary and unavailable observations stay explicit.

### [Native preview](https://lanej.io/charmed/organisms-git.html#git%2fdiff-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
config.go                                                                       
@@ -1 +1 @@                                                                     
    1 -const name = "before"                                                    
    1 +const name = "after"
```

Demo · truecolor · 120 × 32 terminal cells

```
gitui/status.go                                                                                                         
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
gitui/status.go                         
@@ -3,5 +3,5 @@                         
    3  // Status belongs to the supplied
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
gitui/status.go                                                                                                         
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
gitui/status.go                                                                                                         
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
gitui/status.go                                                                                                         
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

### [Go usage](https://lanej.io/charmed/organisms-git.html#git%2fdiff-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

diff := gitui.FileDiff{
	Path: "config.go", State: gitui.TextDiff,
	Hunks: []gitui.Hunk{{Header: "@@ -1 +1 @@", Lines: []gitui.DiffLine{
		{Kind: gitui.DeletedLine, OldLine: 1, Text: `const name = "before"`},
		{Kind: gitui.AddedLine, NewLine: 1, Text: `const name = "after"`},
	}}},
}
view := gitui.NewDiffViewer(diff, 80, 12, palette)

return view.View()
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-git.html#git%2fdiff-integration)

Supply FileDiff with ordered hunk lines and paths. Language overrides path detection; text or an unknown language renders plain syntax. Old and new sides lex separately within each supplied hunk. Replacement lines pair by position inside contiguous change blocks; inline edits preserve graphemes. Theme.Syntax owns foregrounds and Theme.Diff owns change backgrounds. SetDiff rebuilds content and resets scroll while preserving native bindings and styles. Resize only updates viewport dimensions. Binary and unavailable states bypass syntax rendering.

### [Run this example](https://lanej.io/charmed/organisms-git.html#git%2fdiff-run)

```sh
go run ./examples/catalog git/diff
```

[Example source](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/docs/snippets/git_diff.go) · [Implementation](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [git/hunk](https://lanej.io/charmed/molecules-git.html#git/hunk)

[Organisms](https://lanej.io/charmed/organisms.html) / git

## [Changes inspector](https://lanej.io/charmed/organisms-git.html#git%2fchanges)

Coordinate file selection with the associated diff through separately focused status and diff panes.

### [Native preview](https://lanej.io/charmed/organisms-git.html#git%2fchanges-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
┌ Status ────────────────────────────────────────┐┌ Diff ──────────────────────────────────────────┐
│   HEAD unavailable                             ││ config.go                                      │
│ staged 1 · unstaged 0 · untracked 0 · conflict ││ No textual changes                             │
│ upstream unavailable · ahead/behind unavailabl ││                                                │
│                                                ││                                                │
│ staged (1)                                     ││                                                │
│ ▸  added  config.go                            ││                                                │
│                                                ││                                                │
│                                                ││                                                │
│                                                ││                                                │
│                                                ││                                                │
│                                                ││                                                │
│                                                ││                                                │
│                                                ││                                                │
│                                                ││                                                │
│                                                ││                                                │
│                                                ││                                                │
│ ↑/↓ j/k · / search · Enter select              ││                                                │
│                                                ││                                                │
└────────────────────────────────────────────────┘└────────────────────────────────────────────────┘
```

Demo · truecolor · 120 × 32 terminal cells

```
┌ Status ──────────────────────────────────────────────────┐┌ Diff ────────────────────────────────────────────────────┐
│ charmed  HEAD → main                                     ││ gitui/status.go                                          │
│ staged 2 · unstaged 1 · untracked 1 · conflicts 1        ││ @@ -0,0 +1,6 @@                                          │
│ upstream refs/remotes/origin/main · ahead 2 · behind 1   ││     1 +package gitui                                     │
│                                                          ││     2 +                                                  │
│ staged (2)                                               ││     3 +// Status belongs to the supplied worktree.       │
│ ▸  added  gitui/status.go                                ││     4 +func label() string {                             │
│    renamed  docs/git.md → docs/git-components.md         ││     5 +    return "ready"                                │
│ unstaged (1)                                             ││     6 +}                                                 │
│    modified  gitui/status.go                             ││                                                          │
│ untracked (1)                                            ││                                                          │
│    untracked  notes.txt                                  ││                                                          │
│ conflicts (1)                                            ││                                                          │
│    conflicted  theme/theme.go · both modified            ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│ ↑/↓ j/k · / search · Enter select                        ││                                                          │
│                                                          ││                                                          │
└──────────────────────────────────────────────────────────┘└──────────────────────────────────────────────────────────┘
```

Narrow · truecolor · 40 × 16 terminal cells

```
┌ Status ──────────────────────────────┐
│ charmed  HEAD → main                 │
│ staged 2 · unstaged 1 · untracked 1  │
│ upstream refs/remotes/origin/main ·  │
│                                      │
│ ▸  added  gitui/status.go            │
│ ↑/↓ j/k · / search · Enter select    │
└──────────────────────────────────────┘
┌ Diff ────────────────────────────────┐
│ gitui/status.go                      │
│ @@ -0,0 +1,6 @@                      │
│     1 +package gitui                 │
│     2 +                              │
│     3 +// Status belongs to the supp │
│     4 +func label() string {         │
└──────────────────────────────────────┘
```

256 colors · ansi256 · 120 × 32 terminal cells

```
┌ Status ──────────────────────────────────────────────────┐┌ Diff ────────────────────────────────────────────────────┐
│ charmed  HEAD → main                                     ││ gitui/status.go                                          │
│ staged 2 · unstaged 1 · untracked 1 · conflicts 1        ││ @@ -0,0 +1,6 @@                                          │
│ upstream refs/remotes/origin/main · ahead 2 · behind 1   ││     1 +package gitui                                     │
│                                                          ││     2 +                                                  │
│ staged (2)                                               ││     3 +// Status belongs to the supplied worktree.       │
│ ▸  added  gitui/status.go                                ││     4 +func label() string {                             │
│    renamed  docs/git.md → docs/git-components.md         ││     5 +    return "ready"                                │
│ unstaged (1)                                             ││     6 +}                                                 │
│    modified  gitui/status.go                             ││                                                          │
│ untracked (1)                                            ││                                                          │
│    untracked  notes.txt                                  ││                                                          │
│ conflicts (1)                                            ││                                                          │
│    conflicted  theme/theme.go · both modified            ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│ ↑/↓ j/k · / search · Enter select                        ││                                                          │
│                                                          ││                                                          │
└──────────────────────────────────────────────────────────┘└──────────────────────────────────────────────────────────┘
```

16 colors · ansi16 · 120 × 32 terminal cells

```
┌ Status ──────────────────────────────────────────────────┐┌ Diff ────────────────────────────────────────────────────┐
│ charmed  HEAD → main                                     ││ gitui/status.go                                          │
│ staged 2 · unstaged 1 · untracked 1 · conflicts 1        ││ @@ -0,0 +1,6 @@                                          │
│ upstream refs/remotes/origin/main · ahead 2 · behind 1   ││     1 +package gitui                                     │
│                                                          ││     2 +                                                  │
│ staged (2)                                               ││     3 +// Status belongs to the supplied worktree.       │
│ ▸  added  gitui/status.go                                ││     4 +func label() string {                             │
│    renamed  docs/git.md → docs/git-components.md         ││     5 +    return "ready"                                │
│ unstaged (1)                                             ││     6 +}                                                 │
│    modified  gitui/status.go                             ││                                                          │
│ untracked (1)                                            ││                                                          │
│    untracked  notes.txt                                  ││                                                          │
│ conflicts (1)                                            ││                                                          │
│    conflicted  theme/theme.go · both modified            ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│                                                          ││                                                          │
│ ↑/↓ j/k · / search · Enter select                        ││                                                          │
│                                                          ││                                                          │
└──────────────────────────────────────────────────────────┘└──────────────────────────────────────────────────────────┘
```

Monochrome · none · 120 × 32 terminal cells

```
+ Status --------------------------------------------------++ Diff ----------------------------------------------------+
| charmed [HEAD > main]                                    || gitui/status.go                                          |
| staged 2 * unstaged 1 * untracked 1 * conflicts 1        || @@ -0,0 +1,6 @@                                          |
| upstream refs/remotes/origin/main * ahead 2 * behind 1   ||     1 +package gitui                                     |
|                                                          ||     2 +                                                  |
| staged (2)                                               ||     3 +// Status belongs to the supplied worktree.       |
| > [added] gitui/status.go                                ||     4 +func label() string {                             |
|   [renamed] docs/git.md > docs/git-components.md         ||     5 +    return "ready"                                |
| unstaged (1)                                             ||     6 +}                                                 |
|   [modified] gitui/status.go                             ||                                                          |
| untracked (1)                                            ||                                                          |
|   [untracked] notes.txt                                  ||                                                          |
| conflicts (1)                                            ||                                                          |
|   [conflicted] theme/theme.go * both modified            ||                                                          |
|                                                          ||                                                          |
|                                                          ||                                                          |
|                                                          ||                                                          |
|                                                          ||                                                          |
|                                                          ||                                                          |
|                                                          ||                                                          |
|                                                          ||                                                          |
|                                                          ||                                                          |
|                                                          ||                                                          |
|                                                          ||                                                          |
|                                                          ||                                                          |
|                                                          ||                                                          |
|                                                          ||                                                          |
|                                                          ||                                                          |
|                                                          ||                                                          |
| up/down j/k * / search * Enter select                    ||                                                          |
|                                                          ||                                                          |
+----------------------------------------------------------++----------------------------------------------------------+
```

### [Go usage](https://lanej.io/charmed/organisms-git.html#git%2fchanges-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

view, err := gitui.NewChangesInspector(
	gitui.Repository{Files: []gitui.FileChange{{ID: "config", Path: "config.go", Index: gitui.Added}}},
	[]gitui.FileDiff{{FileID: "config", Section: gitui.Staged, Path: "config.go", State: gitui.TextDiff}},
	100, 20, palette)
if err != nil {
	return "", err
}

return view.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-git.html#git%2fchanges-integration)

Supply observations from the application and resize the model to its allocated terminal cells. Route keyboard input to the focused view and return its command. Refresh through the public data setter and use stable identities when acting on a selection.

### [Run this example](https://lanej.io/charmed/organisms-git.html#git%2fchanges-run)

```sh
go run ./examples/catalog git/changes
```

[Example source](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/docs/snippets/git_changes.go) · [Implementation](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [git/status](https://lanej.io/charmed/organisms-git.html#git/status) [git/diff](https://lanej.io/charmed/organisms-git.html#git/diff)

[Organisms](https://lanej.io/charmed/organisms.html) / git

## [Stash list](https://lanej.io/charmed/organisms-git.html#git%2fstashes)

Browse supplied stash entries and return the stable identity of the selected stash.

### [Native preview](https://lanej.io/charmed/organisms-git.html#git%2fstashes-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
▸ stash@{0} WIP: search ·  ·                                                    
                                                                                
                                                                                
                                                                                
                                                                                
                                                                                
                                                                                
                                                                                
                                                                                
                                                                                
↑/↓ j/k · / search · Enter select
```

Demo · truecolor · 120 × 32 terminal cells

```
▸ stash@{0} WIP: status presentation · refs/heads/main · a31f2c7819                                                     
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
↑/↓ j/k · / search · Enter select
```

Narrow · truecolor · 40 × 16 terminal cells

```
▸ stash@{0} WIP: status presentation · …
                                        
                                        
                                        
                                        
                                        
                                        
                                        
                                        
                                        
                                        
                                        
                                        
                                        
↑/↓ j/k · / search · Enter select
```

256 colors · ansi256 · 120 × 32 terminal cells

```
▸ stash@{0} WIP: status presentation · refs/heads/main · a31f2c7819                                                     
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
↑/↓ j/k · / search · Enter select
```

16 colors · ansi16 · 120 × 32 terminal cells

```
▸ stash@{0} WIP: status presentation · refs/heads/main · a31f2c7819                                                     
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
↑/↓ j/k · / search · Enter select
```

Monochrome · none · 120 × 32 terminal cells

```
> stash@{0} WIP: status presentation * refs/heads/main * a31f2c7819                                                     
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
up/down j/k * / search * Enter select
```

### [Go usage](https://lanej.io/charmed/organisms-git.html#git%2fstashes-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

view, err := gitui.NewStashList([]gitui.Stash{{ID: "stash-object", Ordinal: 0, Subject: "WIP: search"}}, 80, 12, palette)
if err != nil {
	return "", err
}

return view.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-git.html#git%2fstashes-integration)

Supply observations from the application and resize the model to its allocated terminal cells. Route keyboard input to the focused view and return its command. Refresh through the public data setter and use stable identities when acting on a selection.

### [Run this example](https://lanej.io/charmed/organisms-git.html#git%2fstashes-run)

```sh
go run ./examples/catalog git/stashes
```

[Example source](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/docs/snippets/git_stashes.go) · [Implementation](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [git/stash-row](https://lanej.io/charmed/molecules-git.html#git/stash-row)

[Organisms](https://lanej.io/charmed/organisms.html) / git

## [Worktree list](https://lanej.io/charmed/organisms-git.html#git%2fworktrees)

Compare checked-out worktrees and select one by its stable identity.

### [Native preview](https://lanej.io/charmed/organisms-git.html#git%2fworktrees-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
▸ ~/src/charmed-search  HEAD unavailable   linked · locked: review · status una…
                                                                                
                                                                                
                                                                                
                                                                                
                                                                                
                                                                                
                                                                                
                                                                                
                                                                                
↑/↓ j/k · / search · Enter select
```

Demo · truecolor · 120 × 32 terminal cells

```
▸ ~/src/charmed  HEAD → main   main · current · 4 changed files                                                         
  ~/src/charmed-search  HEAD → search   linked · locked: review in progress · clean                                     
  ~/src/charmed-old  HEAD detached 28c341aad0   linked · prunable: directory unavailable · status unavailable           
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
↑/↓ j/k · / search · Enter select
```

Narrow · truecolor · 40 × 16 terminal cells

```
▸ ~/src/charmed  HEAD → main   main · c…
  ~/src/charmed-search  HEAD → search  …
  ~/src/charmed-old  HEAD detached 28c3…
                                        
                                        
                                        
                                        
                                        
                                        
                                        
                                        
                                        
                                        
                                        
↑/↓ j/k · / search · Enter select
```

256 colors · ansi256 · 120 × 32 terminal cells

```
▸ ~/src/charmed  HEAD → main   main · current · 4 changed files                                                         
  ~/src/charmed-search  HEAD → search   linked · locked: review in progress · clean                                     
  ~/src/charmed-old  HEAD detached 28c341aad0   linked · prunable: directory unavailable · status unavailable           
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
↑/↓ j/k · / search · Enter select
```

16 colors · ansi16 · 120 × 32 terminal cells

```
▸ ~/src/charmed  HEAD → main   main · current · 4 changed files                                                         
  ~/src/charmed-search  HEAD → search   linked · locked: review in progress · clean                                     
  ~/src/charmed-old  HEAD detached 28c341aad0   linked · prunable: directory unavailable · status unavailable           
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
↑/↓ j/k · / search · Enter select
```

Monochrome · none · 120 × 32 terminal cells

```
> ~/src/charmed [HEAD > main] [main * current * 4 changed files]                                                        
  ~/src/charmed-search [HEAD > search] [linked * locked: review in progress * clean]                                    
  ~/src/charmed-old [HEAD detached 28c341aad0] [linked * prunable: directory unavailable * status unavailable]          
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
up/down j/k * / search * Enter select
```

### [Go usage](https://lanej.io/charmed/organisms-git.html#git%2fworktrees-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

view, err := gitui.NewWorktreeList([]gitui.Worktree{{ID: "search", Path: "~/src/charmed-search", Locked: true, LockReason: "review"}}, 80, 12, palette)
if err != nil {
	return "", err
}

return view.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-git.html#git%2fworktrees-integration)

Supply observations from the application and resize the model to its allocated terminal cells. Route keyboard input to the focused view and return its command. Refresh through the public data setter and use stable identities when acting on a selection.

### [Run this example](https://lanej.io/charmed/organisms-git.html#git%2fworktrees-run)

```sh
go run ./examples/catalog git/worktrees
```

[Example source](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/docs/snippets/git_worktrees.go) · [Implementation](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [git/worktree-row](https://lanej.io/charmed/molecules-git.html#git/worktree-row)

[Organisms](https://lanej.io/charmed/organisms.html) / git

## [Patch browser](https://lanej.io/charmed/organisms-git.html#git%2fpatch-browser)

Browse every file in an observed patch, including staged and unstaged versions of the same file. Inventory availability remains distinct from a known empty patch.

### [Native preview](https://lanej.io/charmed/organisms-git.html#git%2fpatch-browser-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
staged                                                                          
▸ config.yaml                                                                   
unstaged                                                                        
  logo.png                                                                      
                                                                                
                                                                                
                                                                                
config.yaml                                                                     
@@ -1 +1 @@                                                                     
    1 -replicas: 2                                                              
    1 +replicas: 3
```

Demo · truecolor · 100 × 30 terminal cells

```
staged                                                      gitui/status.go                         
▸ gitui/status.go                                           @@ -0,0 +1,6 @@                         
unstaged                                                        1 +package gitui                    
  gitui/status.go                                               2 +                                 
staged                                                          3 +// Status belongs to the supplie 
  docs/git-components.md                                        4 +func label() string {            
untracked                                                       5 +    return "ready"               
  notes.txt                                                     6 +}                                
conflicts                                                                                           
  theme/theme.go
```

Narrow · truecolor · 40 × 16 terminal cells

```
staged                                  
▸ gitui/status.go                       
unstaged                                
  gitui/status.go                       
staged                                  
  docs/git-components.md                
untracked                               
  notes.txt                             
gitui/status.go                         
@@ -0,0 +1,6 @@                         
    1 +package gitui                    
    2 +                                 
    3 +// Status belongs to the supplie 
    4 +func label() string {            
    5 +    return "ready"               
    6 +}
```

256 colors · ansi256 · 100 × 30 terminal cells

```
staged                                                      gitui/status.go                         
▸ gitui/status.go                                           @@ -0,0 +1,6 @@                         
unstaged                                                        1 +package gitui                    
  gitui/status.go                                               2 +                                 
staged                                                          3 +// Status belongs to the supplie 
  docs/git-components.md                                        4 +func label() string {            
untracked                                                       5 +    return "ready"               
  notes.txt                                                     6 +}                                
conflicts                                                                                           
  theme/theme.go
```

16 colors · ansi16 · 100 × 30 terminal cells

```
staged                                                      gitui/status.go                         
▸ gitui/status.go                                           @@ -0,0 +1,6 @@                         
unstaged                                                        1 +package gitui                    
  gitui/status.go                                               2 +                                 
staged                                                          3 +// Status belongs to the supplie 
  docs/git-components.md                                        4 +func label() string {            
untracked                                                       5 +    return "ready"               
  notes.txt                                                     6 +}                                
conflicts                                                                                           
  theme/theme.go
```

Monochrome · none · 100 × 30 terminal cells

```
staged                                                      gitui/status.go                         
> gitui/status.go                                           @@ -0,0 +1,6 @@                         
unstaged                                                        1 +package gitui                    
  gitui/status.go                                               2 +                                 
staged                                                          3 +// Status belongs to the supplie 
  docs/git-components.md                                        4 +func label() string {            
untracked                                                       5 +    return "ready"               
  notes.txt                                                     6 +}                                
conflicts                                                                                           
  theme/theme.go
```

### [Go usage](https://lanej.io/charmed/organisms-git.html#git%2fpatch-browser-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
patch := gitui.PatchObservation{Available: true, Files: []gitui.FileDiff{
	{FileID: "config", Section: gitui.Staged, Path: "config.yaml", State: gitui.TextDiff, Hunks: []gitui.Hunk{
		{Header: "@@ -1 +1 @@", Lines: []gitui.DiffLine{{Kind: gitui.DeletedLine, OldLine: 1, Text: "replicas: 2"}, {Kind: gitui.AddedLine, NewLine: 1, Text: "replicas: 3"}}},
	}},
	{FileID: "logo", Section: gitui.Unstaged, Path: "logo.png", State: gitui.BinaryDiff},
}}

browser, err := gitui.NewPatchBrowser(patch, layout.Size{Width: 80, Height: 14}, theme.New(theme.Dark))
if err != nil {
	return "", err
}
// FocusDetail controls routing; SetPatch deep-copies and preserves surviving identity.
// Available=true with no files is empty; the zero observation is unavailable.
return browser.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-git.html#git%2fpatch-browser-integration)

Supply a PatchObservation with explicit availability and ordered diffs; SetPatch validates identities and deep-copies nested hunk data. Select by file ID and change section. Route Tab through FocusDetail and forward Update commands to the active pane. Resize stacks narrow panes. List and diff native customization remain accessible; content, dimensions and selection scrolling are managed by the browser. Call Refresh after direct list selection changes. Instructions are opt-in through ShowHelp; resize after changing it.

### [Run this example](https://lanej.io/charmed/organisms-git.html#git%2fpatch-browser-run)

```sh
go run ./examples/catalog git/patch-browser
```

[Example source](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/docs/snippets/git_patch_browser.go) · [Implementation](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/gitui/patch_browser.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [git/diff](https://lanej.io/charmed/organisms-git.html#git/diff) [collection/scroll-pane](https://lanej.io/charmed/molecules-collection.html#collection/scroll-pane)
