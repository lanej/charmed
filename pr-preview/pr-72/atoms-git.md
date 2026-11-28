[Atoms](https://lanej.io/charmed/atoms.html)git

# [git](https://lanej.io/charmed/atoms-git.html#level-atoms-git)

11 components · [All atoms](https://lanej.io/charmed/atoms.html)

- [Commit identity](https://lanej.io/charmed/atoms-git.html#git%2fcommit-id)
- [Ref label](https://lanej.io/charmed/atoms-git.html#git%2fref)
- [HEAD indicator](https://lanej.io/charmed/atoms-git.html#git%2fhead)
- [File change](https://lanej.io/charmed/atoms-git.html#git%2ffile-state)
- [Divergence](https://lanej.io/charmed/atoms-git.html#git%2fdivergence)
- [Diff summary](https://lanej.io/charmed/atoms-git.html#git%2fdiff-summary)
- [Worktree state](https://lanej.io/charmed/atoms-git.html#git%2fworktree-state)
- [Staged state](https://lanej.io/charmed/atoms-git.html#git%2fstaged-state)
- [Unstaged state](https://lanej.io/charmed/atoms-git.html#git%2funstaged-state)
- [Untracked state](https://lanej.io/charmed/atoms-git.html#git%2funtracked-state)
- [Conflict state](https://lanej.io/charmed/atoms-git.html#git%2fconflict-state)

[Atoms](https://lanej.io/charmed/atoms.html) / git

## [Commit identity](https://lanej.io/charmed/atoms-git.html#git%2fcommit-id)

Show an abbreviated commit in a narrow row while retaining its full ID for selection and inspection.

### [Native preview](https://lanej.io/charmed/atoms-git.html#git%2fcommit-id-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
a31f2c7819
```

Demo · truecolor · 120 × 32 terminal cells

```
a31f2c7819
```

Narrow · truecolor · 40 × 16 terminal cells

```
a31f2c7819
```

256 colors · ansi256 · 120 × 32 terminal cells

```
a31f2c7819
```

16 colors · ansi16 · 120 × 32 terminal cells

```
a31f2c7819
```

Monochrome · none · 120 × 32 terminal cells

```
a31f2c7819
```

### [Go usage](https://lanej.io/charmed/atoms-git.html#git%2fcommit-id-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return gitui.CommitIdentity("a31f2c7819be822", palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-git.html#git%2fcommit-id-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-git.html#git%2fcommit-id-run)

```sh
go run ./examples/catalog git/commit-id
```

[Example source](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/docs/snippets/git_commit_id.go) · [Implementation](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / git

## [Ref label](https://lanej.io/charmed/atoms-git.html#git%2fref)

Display the name and kind of a ref beside the commit or branch it identifies.

### [Native preview](https://lanej.io/charmed/atoms-git.html#git%2fref-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 branch main 
```

Demo · truecolor · 120 × 32 terminal cells

```
 branch main                                                                                                            
                                                                                                                        
 branch search                                                                                                          
                                                                                                                        
 remote origin/main                                                                                                     
                                                                                                                        
 tag v0.1 
```

Narrow · truecolor · 40 × 16 terminal cells

```
 branch main                            
                                        
 branch search                          
                                        
 remote origin/main                     
                                        
 tag v0.1 
```

256 colors · ansi256 · 120 × 32 terminal cells

```
 branch main                                                                                                            
                                                                                                                        
 branch search                                                                                                          
                                                                                                                        
 remote origin/main                                                                                                     
                                                                                                                        
 tag v0.1 
```

16 colors · ansi16 · 120 × 32 terminal cells

```
 branch main                                                                                                            
                                                                                                                        
 branch search                                                                                                          
                                                                                                                        
 remote origin/main                                                                                                     
                                                                                                                        
 tag v0.1 
```

Monochrome · none · 120 × 32 terminal cells

```
[branch main]                                                                                                           
                                                                                                                        
[branch search]                                                                                                         
                                                                                                                        
[remote origin/main]                                                                                                    
                                                                                                                        
[tag v0.1]
```

### [Go usage](https://lanej.io/charmed/atoms-git.html#git%2fref-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return gitui.RefLabel(gitui.Ref{Name: "refs/heads/main", Kind: gitui.LocalBranch}, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-git.html#git%2fref-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-git.html#git%2fref-run)

```sh
go run ./examples/catalog git/ref
```

[Example source](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/docs/snippets/git_ref.go) · [Implementation](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / git

## [HEAD indicator](https://lanej.io/charmed/atoms-git.html#git%2fhead)

Mark the supplied HEAD independently of other refs pointing at the same commit.

### [Native preview](https://lanej.io/charmed/atoms-git.html#git%2fhead-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 HEAD unborn main 
```

Demo · truecolor · 120 × 32 terminal cells

```
 HEAD → main                                                                                                            
                                                                                                                        
 HEAD detached 28c341aad0                                                                                               
                                                                                                                        
 HEAD unborn new 
```

Narrow · truecolor · 40 × 16 terminal cells

```
 HEAD → main                            
                                        
 HEAD detached 28c341aad0               
                                        
 HEAD unborn new 
```

256 colors · ansi256 · 120 × 32 terminal cells

```
 HEAD → main                                                                                                            
                                                                                                                        
 HEAD detached 28c341aad0                                                                                               
                                                                                                                        
 HEAD unborn new 
```

16 colors · ansi16 · 120 × 32 terminal cells

```
 HEAD → main                                                                                                            
                                                                                                                        
 HEAD detached 28c341aad0                                                                                               
                                                                                                                        
 HEAD unborn new 
```

Monochrome · none · 120 × 32 terminal cells

```
[HEAD > main]                                                                                                           
                                                                                                                        
[HEAD detached 28c341aad0]                                                                                              
                                                                                                                        
[HEAD unborn new]
```

### [Go usage](https://lanej.io/charmed/atoms-git.html#git%2fhead-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return gitui.HeadIndicator(gitui.Head{Kind: gitui.Unborn, Ref: "refs/heads/main"}, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-git.html#git%2fhead-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-git.html#git%2fhead-run)

```sh
go run ./examples/catalog git/head
```

[Example source](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/docs/snippets/git_head.go) · [Implementation](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / git

## [File change](https://lanej.io/charmed/atoms-git.html#git%2ffile-state)

Identify how a file changed while keeping its staging section and path separate.

### [Native preview](https://lanej.io/charmed/atoms-git.html#git%2ffile-state-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 modified 
```

Demo · truecolor · 120 × 32 terminal cells

```
 added                                                                                                                  
 modified                                                                                                               
 deleted                                                                                                                
 renamed                                                                                                                
 copied                                                                                                                 
 type changed                                                                                                           
 untracked                                                                                                              
 conflicted 
```

Narrow · truecolor · 40 × 16 terminal cells

```
 added                                  
 modified                               
 deleted                                
 renamed                                
 copied                                 
 type changed                           
 untracked                              
 conflicted 
```

256 colors · ansi256 · 120 × 32 terminal cells

```
 added                                                                                                                  
 modified                                                                                                               
 deleted                                                                                                                
 renamed                                                                                                                
 copied                                                                                                                 
 type changed                                                                                                           
 untracked                                                                                                              
 conflicted 
```

16 colors · ansi16 · 120 × 32 terminal cells

```
 added                                                                                                                  
 modified                                                                                                               
 deleted                                                                                                                
 renamed                                                                                                                
 copied                                                                                                                 
 type changed                                                                                                           
 untracked                                                                                                              
 conflicted 
```

Monochrome · none · 120 × 32 terminal cells

```
[added]                                                                                                                 
[modified]                                                                                                              
[deleted]                                                                                                               
[renamed]                                                                                                               
[copied]                                                                                                                
[type changed]                                                                                                          
[untracked]                                                                                                             
[conflicted]
```

### [Go usage](https://lanej.io/charmed/atoms-git.html#git%2ffile-state-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return gitui.FileIndicator(gitui.Modified, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-git.html#git%2ffile-state-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-git.html#git%2ffile-state-run)

```sh
go run ./examples/catalog git/file-state
```

[Example source](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/docs/snippets/git_file_state.go) · [Implementation](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / git

## [Divergence](https://lanej.io/charmed/atoms-git.html#git%2fdivergence)

Show how far a branch is ahead or behind its upstream, including when those counts are unknown.

### [Native preview](https://lanej.io/charmed/atoms-git.html#git%2fdivergence-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
ahead 2 · behind 1
```

Demo · truecolor · 120 × 32 terminal cells

```
ahead 2 · behind 1                                                                                                      
                                                                                                                        
ahead/behind unavailable
```

Narrow · truecolor · 40 × 16 terminal cells

```
ahead 2 · behind 1                      
                                        
ahead/behind unavailable
```

256 colors · ansi256 · 120 × 32 terminal cells

```
ahead 2 · behind 1                                                                                                      
                                                                                                                        
ahead/behind unavailable
```

16 colors · ansi16 · 120 × 32 terminal cells

```
ahead 2 · behind 1                                                                                                      
                                                                                                                        
ahead/behind unavailable
```

Monochrome · none · 120 × 32 terminal cells

```
ahead 2 * behind 1                                                                                                      
                                                                                                                        
ahead/behind unavailable
```

### [Go usage](https://lanej.io/charmed/atoms-git.html#git%2fdivergence-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return gitui.DivergenceIndicator(gitui.Divergence{Known: true, Ahead: 2, Behind: 1}, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-git.html#git%2fdivergence-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-git.html#git%2fdivergence-run)

```sh
go run ./examples/catalog git/divergence
```

[Example source](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/docs/snippets/git_divergence.go) · [Implementation](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / git

## [Diff summary](https://lanej.io/charmed/atoms-git.html#git%2fdiff-summary)

Summarize supplied additions and deletions before opening the complete diff.

### [Native preview](https://lanej.io/charmed/atoms-git.html#git%2fdiff-summary-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
+12 -3
```

Demo · truecolor · 120 × 32 terminal cells

```
+12 -3                                                                                                                  
                                                                                                                        
binary                                                                                                                  
                                                                                                                        
diff unavailable
```

Narrow · truecolor · 40 × 16 terminal cells

```
+12 -3                                  
                                        
binary                                  
                                        
diff unavailable
```

256 colors · ansi256 · 120 × 32 terminal cells

```
+12 -3                                                                                                                  
                                                                                                                        
binary                                                                                                                  
                                                                                                                        
diff unavailable
```

16 colors · ansi16 · 120 × 32 terminal cells

```
+12 -3                                                                                                                  
                                                                                                                        
binary                                                                                                                  
                                                                                                                        
diff unavailable
```

Monochrome · none · 120 × 32 terminal cells

```
+12 -3                                                                                                                  
                                                                                                                        
binary                                                                                                                  
                                                                                                                        
diff unavailable
```

### [Go usage](https://lanej.io/charmed/atoms-git.html#git%2fdiff-summary-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return gitui.DiffSummary(gitui.DiffStat{State: gitui.TextDiff, Added: 12, Deleted: 3}, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-git.html#git%2fdiff-summary-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-git.html#git%2fdiff-summary-run)

```sh
go run ./examples/catalog git/diff-summary
```

[Example source](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/docs/snippets/git_diff_summary.go) · [Implementation](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / git

## [Worktree state](https://lanej.io/charmed/atoms-git.html#git%2fworktree-state)

Identify a worktree’s supplied state when comparing several checkout locations.

### [Native preview](https://lanej.io/charmed/atoms-git.html#git%2fworktree-state-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 main · current · clean 
```

Demo · truecolor · 120 × 32 terminal cells

```
 linked · locked: review in progress · clean                                                                            
                                                                                                                        
 linked · prunable: directory unavailable · status unavailable 
```

Narrow · truecolor · 40 × 16 terminal cells

```
 linked · locked: review in progress · c
                                        
 linked · prunable: directory unavailabl
```

256 colors · ansi256 · 120 × 32 terminal cells

```
 linked · locked: review in progress · clean                                                                            
                                                                                                                        
 linked · prunable: directory unavailable · status unavailable 
```

16 colors · ansi16 · 120 × 32 terminal cells

```
 linked · locked: review in progress · clean                                                                            
                                                                                                                        
 linked · prunable: directory unavailable · status unavailable 
```

Monochrome · none · 120 × 32 terminal cells

```
[linked * locked: review in progress * clean]                                                                           
                                                                                                                        
[linked * prunable: directory unavailable * status unavailable]
```

### [Go usage](https://lanej.io/charmed/atoms-git.html#git%2fworktree-state-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return gitui.WorktreeIndicators(gitui.Worktree{Main: true, Current: true, StatusKnown: true}, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-git.html#git%2fworktree-state-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-git.html#git%2fworktree-state-run)

```sh
go run ./examples/catalog git/worktree-state
```

[Example source](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/docs/snippets/git_worktree_state.go) · [Implementation](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / git

## [Staged state](https://lanej.io/charmed/atoms-git.html#git%2fstaged-state)

Mark a file or section observed in the Git index.

### [Native preview](https://lanej.io/charmed/atoms-git.html#git%2fstaged-state-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 staged 
```

Demo · truecolor · 40 × 8 terminal cells

```
 staged 
```

Narrow · truecolor · 40 × 16 terminal cells

```
 staged 
```

256 colors · ansi256 · 40 × 8 terminal cells

```
 staged 
```

16 colors · ansi16 · 40 × 8 terminal cells

```
 staged 
```

Monochrome · none · 40 × 8 terminal cells

```
[staged]
```

### [Go usage](https://lanej.io/charmed/atoms-git.html#git%2fstaged-state-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return gitui.SectionIndicator(gitui.Staged, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-git.html#git%2fstaged-state-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-git.html#git%2fstaged-state-run)

```sh
go run ./examples/catalog git/staged-state
```

[Example source](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/docs/snippets/git_staged_state.go) · [Implementation](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / git

## [Unstaged state](https://lanej.io/charmed/atoms-git.html#git%2funstaged-state)

Mark a working-tree change that has not been staged.

### [Native preview](https://lanej.io/charmed/atoms-git.html#git%2funstaged-state-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 unstaged 
```

Demo · truecolor · 40 × 8 terminal cells

```
 unstaged 
```

Narrow · truecolor · 40 × 16 terminal cells

```
 unstaged 
```

256 colors · ansi256 · 40 × 8 terminal cells

```
 unstaged 
```

16 colors · ansi16 · 40 × 8 terminal cells

```
 unstaged 
```

Monochrome · none · 40 × 8 terminal cells

```
[unstaged]
```

### [Go usage](https://lanej.io/charmed/atoms-git.html#git%2funstaged-state-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return gitui.SectionIndicator(gitui.Unstaged, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-git.html#git%2funstaged-state-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-git.html#git%2funstaged-state-run)

```sh
go run ./examples/catalog git/unstaged-state
```

[Example source](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/docs/snippets/git_unstaged_state.go) · [Implementation](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / git

## [Untracked state](https://lanej.io/charmed/atoms-git.html#git%2funtracked-state)

Identify files the supplied repository observation reports as untracked.

### [Native preview](https://lanej.io/charmed/atoms-git.html#git%2funtracked-state-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 untracked 
```

Demo · truecolor · 40 × 8 terminal cells

```
 untracked 
```

Narrow · truecolor · 40 × 16 terminal cells

```
 untracked 
```

256 colors · ansi256 · 40 × 8 terminal cells

```
 untracked 
```

16 colors · ansi16 · 40 × 8 terminal cells

```
 untracked 
```

Monochrome · none · 40 × 8 terminal cells

```
[untracked]
```

### [Go usage](https://lanej.io/charmed/atoms-git.html#git%2funtracked-state-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return gitui.SectionIndicator(gitui.UntrackedFiles, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-git.html#git%2funtracked-state-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-git.html#git%2funtracked-state-run)

```sh
go run ./examples/catalog git/untracked-state
```

[Example source](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/docs/snippets/git_untracked_state.go) · [Implementation](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / git

## [Conflict state](https://lanej.io/charmed/atoms-git.html#git%2fconflict-state)

Mark unresolved paths independently of the kind of file change.

### [Native preview](https://lanej.io/charmed/atoms-git.html#git%2fconflict-state-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 conflicts 
```

Demo · truecolor · 40 × 8 terminal cells

```
 conflicts 
```

Narrow · truecolor · 40 × 16 terminal cells

```
 conflicts 
```

256 colors · ansi256 · 40 × 8 terminal cells

```
 conflicts 
```

16 colors · ansi16 · 40 × 8 terminal cells

```
 conflicts 
```

Monochrome · none · 40 × 8 terminal cells

```
[conflicts]
```

### [Go usage](https://lanej.io/charmed/atoms-git.html#git%2fconflict-state-code)

Imports: [`gitui`](https://pkg.go.dev/github.com/lanej/charmed/gitui "github.com/lanej/charmed/gitui") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return gitui.SectionIndicator(gitui.Conflicts, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-git.html#git%2fconflict-state-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-git.html#git%2fconflict-state-run)

```sh
go run ./examples/catalog git/conflict-state
```

[Example source](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/docs/snippets/git_conflict_state.go) · [Implementation](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/story/git.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.
