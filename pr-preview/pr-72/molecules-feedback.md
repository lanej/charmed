[Molecules](https://lanej.io/charmed/molecules.html)Feedback

# [Feedback](https://lanej.io/charmed/molecules-feedback.html#level-molecules-feedback)

7 components · [All molecules](https://lanej.io/charmed/molecules.html)

- [Status cell](https://lanej.io/charmed/molecules-feedback.html#feedback%2fstatus-cell)
- [Pagination](https://lanej.io/charmed/molecules-feedback.html#feedback%2fpaginator)
- [Help](https://lanej.io/charmed/molecules-feedback.html#feedback%2fhelp)
- [Attention](https://lanej.io/charmed/molecules-feedback.html#feedback%2fattention)
- [Percentage bar](https://lanej.io/charmed/molecules-feedback.html#feedback%2fpercent)
- [Evidence summary](https://lanej.io/charmed/molecules-feedback.html#feedback%2fevidence-summary)
- [Evidence footer](https://lanej.io/charmed/molecules-feedback.html#feedback%2fevidence-footer)

[Molecules](https://lanej.io/charmed/molecules.html) / Feedback

## [Status cell](https://lanej.io/charmed/molecules-feedback.html#feedback%2fstatus-cell)

A semantic health fill and glyph with version or refined-state text; focus edges preserve health color.

### [Native preview](https://lanej.io/charmed/molecules-feedback.html#feedback%2fstatus-cell-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 ✓ v2.8.1         
▏! draining      ▕
 × unhealthy      
 ? no probe       
```

Demo · truecolor · 100 × 30 terminal cells

```
Health + version or refined state                                                                   
                                                                                                    
 ✓ v2.8.1          ▏✓ v2.8.1        ▕ OK                                                            
 ! draining        ▏! draining      ▕ WARN                                                          
 × v2.8.1          ▏× v2.8.1        ▕ ERROR                                                         
 ? v2.8.1          ▏? v2.8.1        ▕ UNKNOWN                                                       
 ⏻ v2.8.1          ▏⏻ v2.8.1        ▕ OFF                                                           
 ◌ v2.8.1          ▏◌ v2.8.1        ▕ INACTIVE                                                      
 ◷ v2.8.1          ▏◷ v2.8.1        ▕ PENDING                                                       
 ▶ v2.8.1          ▏▶ v2.8.1        ▕ RUNNING                                                       
 ‖ v2.8.1          ▏‖ v2.8.1        ▕ PAUSED                                                        
 ⊘ v2.8.1          ▏⊘ v2.8.1        ▕ DISABLED                                                      
 ● v2.8.1          ▏● v2.8.1        ▕ CONNECTED                                                     
 ○ v2.8.1          ▏○ v2.8.1        ▕ DISCONNECTED
```

Narrow · truecolor · 40 × 16 terminal cells

```
Health + version or refined state       
                                        
 ✓ v2.8.1          ▏✓ v2.8.1        ▕ OK
 ! draining        ▏! draining      ▕ WA
 × v2.8.1          ▏× v2.8.1        ▕ ER
 ? v2.8.1          ▏? v2.8.1        ▕ UN
 ⏻ v2.8.1          ▏⏻ v2.8.1        ▕ OF
 ◌ v2.8.1          ▏◌ v2.8.1        ▕ IN
 ◷ v2.8.1          ▏◷ v2.8.1        ▕ PE
 ▶ v2.8.1          ▏▶ v2.8.1        ▕ RU
 ‖ v2.8.1          ▏‖ v2.8.1        ▕ PA
 ⊘ v2.8.1          ▏⊘ v2.8.1        ▕ DI
 ● v2.8.1          ▏● v2.8.1        ▕ CO
 ○ v2.8.1          ▏○ v2.8.1        ▕ DI
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Health + version or refined state                                                                   
                                                                                                    
 ✓ v2.8.1          ▏✓ v2.8.1        ▕ OK                                                            
 ! draining        ▏! draining      ▕ WARN                                                          
 × v2.8.1          ▏× v2.8.1        ▕ ERROR                                                         
 ? v2.8.1          ▏? v2.8.1        ▕ UNKNOWN                                                       
 ⏻ v2.8.1          ▏⏻ v2.8.1        ▕ OFF                                                           
 ◌ v2.8.1          ▏◌ v2.8.1        ▕ INACTIVE                                                      
 ◷ v2.8.1          ▏◷ v2.8.1        ▕ PENDING                                                       
 ▶ v2.8.1          ▏▶ v2.8.1        ▕ RUNNING                                                       
 ‖ v2.8.1          ▏‖ v2.8.1        ▕ PAUSED                                                        
 ⊘ v2.8.1          ▏⊘ v2.8.1        ▕ DISABLED                                                      
 ● v2.8.1          ▏● v2.8.1        ▕ CONNECTED                                                     
 ○ v2.8.1          ▏○ v2.8.1        ▕ DISCONNECTED
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Health + version or refined state                                                                   
                                                                                                    
 ✓ v2.8.1          ▏✓ v2.8.1        ▕ OK                                                            
 ! draining        ▏! draining      ▕ WARN                                                          
 × v2.8.1          ▏× v2.8.1        ▕ ERROR                                                         
 ? v2.8.1          ▏? v2.8.1        ▕ UNKNOWN                                                       
 ⏻ v2.8.1          ▏⏻ v2.8.1        ▕ OFF                                                           
 ◌ v2.8.1          ▏◌ v2.8.1        ▕ INACTIVE                                                      
 ◷ v2.8.1          ▏◷ v2.8.1        ▕ PENDING                                                       
 ▶ v2.8.1          ▏▶ v2.8.1        ▕ RUNNING                                                       
 ‖ v2.8.1          ▏‖ v2.8.1        ▕ PAUSED                                                        
 ⊘ v2.8.1          ▏⊘ v2.8.1        ▕ DISABLED                                                      
 ● v2.8.1          ▏● v2.8.1        ▕ CONNECTED                                                     
 ○ v2.8.1          ▏○ v2.8.1        ▕ DISCONNECTED
```

Monochrome · none · 100 × 30 terminal cells

```
Health + version or refined state                                                                   
                                                                                                    
 + v2.8.1          [+ v2.8.1        ] OK                                                            
 ! draining        [! draining      ] WARN                                                          
 x v2.8.1          [x v2.8.1        ] ERROR                                                         
 ? v2.8.1          [? v2.8.1        ] UNKNOWN                                                       
 - v2.8.1          [- v2.8.1        ] OFF                                                           
 - v2.8.1          [- v2.8.1        ] INACTIVE                                                      
 p v2.8.1          [p v2.8.1        ] PENDING                                                       
 > v2.8.1          [> v2.8.1        ] RUNNING                                                       
 = v2.8.1          [= v2.8.1        ] PAUSED                                                        
 / v2.8.1          [/ v2.8.1        ] DISABLED                                                      
 + v2.8.1          [+ v2.8.1        ] CONNECTED                                                     
 - v2.8.1          [- v2.8.1        ] DISCONNECTED
```

### [Go usage](https://lanej.io/charmed/molecules-feedback.html#feedback%2fstatus-cell-code)

Imports: [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
healthy := feedback.StatusCell{Status: feedback.OK, Label: "v2.8.1"}
draining := feedback.StatusCell{Status: feedback.Warn, Label: "draining"}
failed := feedback.StatusCell{Status: feedback.Error, Label: "unhealthy"}
unknown := feedback.StatusCell{Status: feedback.Unknown, Label: "no probe"}

return layout.Column(
	healthy.View(18, palette, false),
	draining.View(18, palette, true),
	failed.View(18, palette, false),
	unknown.View(18, palette, false),
)
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-feedback.html#feedback%2fstatus-cell-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/molecules-feedback.html#feedback%2fstatus-cell-run)

```sh
go run ./examples/catalog feedback/status-cell
```

[Example source](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/docs/snippets/feedback_status_cell.go) · [Implementation](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/feedback/status_cell.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

### Reusable component

[Status cell dependency graph](https://lanej.io/charmed/dependencies.html#component-feedback%2Fstatus-cell) · Partial map · [Source declaration](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/feedback/status_cell.go#L11)

- [Status](https://lanej.io/charmed/dependencies.html#component-feedback%2Fstatus) · composes · health presentation

Consumer usage uses

- [Status cell](https://lanej.io/charmed/dependencies.html#component-feedback%2Fstatus-cell) · primary demonstration

[Example source](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/docs/snippets/feedback_status_cell.go#L9)

[Molecules](https://lanej.io/charmed/molecules.html) / Feedback

## [Pagination](https://lanej.io/charmed/molecules-feedback.html#feedback%2fpaginator)

Native pagination logic with shared active and inactive styles.

### [Native preview](https://lanej.io/charmed/molecules-feedback.html#feedback%2fpaginator-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
○○◉○○
Items 21–30 of 50
```

Demo · truecolor · 100 × 30 terminal cells

```
Pagination                                                                                          
                                                                                                    
◉○○○○                                                                                               
                                                                                                    
Page 1 of 5                                                                                         
                                                                                                    
Left / right changes page
```

Narrow · truecolor · 40 × 16 terminal cells

```
Pagination                              
                                        
◉○○○○                                   
                                        
Page 1 of 5                             
                                        
Left / right changes page
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Pagination                                                                                          
                                                                                                    
◉○○○○                                                                                               
                                                                                                    
Page 1 of 5                                                                                         
                                                                                                    
Left / right changes page
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Pagination                                                                                          
                                                                                                    
◉○○○○                                                                                               
                                                                                                    
Page 1 of 5                                                                                         
                                                                                                    
Left / right changes page
```

Monochrome · none · 100 × 30 terminal cells

```
Pagination                                                                                          
                                                                                                    
*oooo                                                                                               
                                                                                                    
Page 1 of 5                                                                                         
                                                                                                    
Left / right changes page
```

### [Go usage](https://lanej.io/charmed/molecules-feedback.html#feedback%2fpaginator-code)

Imports: [`fmt`](https://pkg.go.dev/fmt "fmt") · [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
pages := feedback.NewPaginator(50, 10, theme.New(theme.Dark))
pages.Page = 2 // Zero-based: show the third page of five.
start, end := pages.GetSliceBounds(50)

return pages.View() + fmt.Sprintf("\nItems %d–%d of 50", start+1, end)
// Forward messages to pages.Update(msg); use start:end to select your visible items.
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-feedback.html#feedback%2fpaginator-integration)

Keep the native control state in the owning application model and route keyboard messages only while its region is active. Retain the model and command returned by Update. The parent supplies page counts or other configuration, then recomputes its placement when the terminal allocation changes.

### [Run this example](https://lanej.io/charmed/molecules-feedback.html#feedback%2fpaginator-run)

```sh
go run ./examples/catalog feedback/paginator
```

[Example source](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/docs/snippets/feedback_paginator.go) · [Implementation](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/story/molecules.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Molecules](https://lanej.io/charmed/molecules.html) / Feedback

## [Help](https://lanej.io/charmed/molecules-feedback.html#feedback%2fhelp)

Help is generated from the same native key bindings that handle input.

### [Native preview](https://lanej.io/charmed/molecules-feedback.html#feedback%2fhelp-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
tab next field • enter select
```

Demo · truecolor · 100 × 30 terminal cells

```
Help from key bindings                                                                              
                                                                                                    
tab next field • enter select
```

Narrow · truecolor · 40 × 16 terminal cells

```
Help from key bindings                  
                                        
tab next field • enter select
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Help from key bindings                                                                              
                                                                                                    
tab next field • enter select
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Help from key bindings                                                                              
                                                                                                    
tab next field • enter select
```

Monochrome · none · 100 × 30 terminal cells

```
Help from key bindings                                                                              
                                                                                                    
tab next field * enter select
```

### [Go usage](https://lanej.io/charmed/molecules-feedback.html#feedback%2fhelp-code)

Imports: [`key`](https://pkg.go.dev/charm.land/bubbles/v2/key "charm.land/bubbles/v2/key") · [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
next := key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "next field"))
selectItem := key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "select"))
help := feedback.NewHelp(40, palette)

return help.ShortHelpView([]key.Binding{next, selectItem})
// Use these same bindings in Update: key.Matches(msg, next).
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-feedback.html#feedback%2fhelp-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/molecules-feedback.html#feedback%2fhelp-run)

```sh
go run ./examples/catalog feedback/help
```

[Example source](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/docs/snippets/feedback_help.go) · [Implementation](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/story/molecules.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Molecules](https://lanej.io/charmed/molecules.html) / Feedback

## [Attention](https://lanej.io/charmed/molecules-feedback.html#feedback%2fattention)

Optional warn/error boundaries compose a labeled attention indicator from the shared semantic pill atom.

### [Native preview](https://lanej.io/charmed/molecules-feedback.html#feedback%2fattention-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 ! WARN 
```

Demo · truecolor · 100 × 30 terminal cells

```
Composed attention indicator                                                                        
                                                                                                    
 ! WARN                                                                                             
                                                                                                    
Value 68                                                                                            
Warn>=+60 Error>=+80 units                                                                          
                                                                                                    
Rules are optional. Missing measurements are UNKNOWN.                                               
Equality triggers; ERROR takes precedence over WARN.                                                
The indicator composes the shared semantic pill atom.
```

Narrow · truecolor · 40 × 16 terminal cells

```
Composed attention indicator            
                                        
 ! WARN                                 
                                        
Value 68                                
Warn>=+60 Error>=+80 units              
                                        
Rules are optional. Missing measurements
Equality triggers; ERROR takes precedenc
The indicator composes the shared semant
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Composed attention indicator                                                                        
                                                                                                    
 ! WARN                                                                                             
                                                                                                    
Value 68                                                                                            
Warn>=+60 Error>=+80 units                                                                          
                                                                                                    
Rules are optional. Missing measurements are UNKNOWN.                                               
Equality triggers; ERROR takes precedence over WARN.                                                
The indicator composes the shared semantic pill atom.
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Composed attention indicator                                                                        
                                                                                                    
 ! WARN                                                                                             
                                                                                                    
Value 68                                                                                            
Warn>=+60 Error>=+80 units                                                                          
                                                                                                    
Rules are optional. Missing measurements are UNKNOWN.                                               
Equality triggers; ERROR takes precedence over WARN.                                                
The indicator composes the shared semantic pill atom.
```

Monochrome · none · 100 × 30 terminal cells

```
Composed attention indicator                                                                        
                                                                                                    
[! WARN]                                                                                            
                                                                                                    
Value 68                                                                                            
Warn>=+60 Error>=+80 units                                                                          
                                                                                                    
Rules are optional. Missing measurements are UNKNOWN.                                               
Equality triggers; ERROR takes precedence over WARN.                                                
The indicator composes the shared semantic pill atom.
```

No boundaries · truecolor · 100 × 30 terminal cells

```
Composed attention indicator                                                                        
                                                                                                    
OFF / no boundaries                                                                                 
                                                                                                    
Value 68                                                                                            
off                                                                                                 
                                                                                                    
Rules are optional. Missing measurements are UNKNOWN.                                               
Equality triggers; ERROR takes precedence over WARN.                                                
The indicator composes the shared semantic pill atom.
```

### [Go usage](https://lanej.io/charmed/molecules-feedback.html#feedback%2fattention-code)

Imports: [`attention`](https://pkg.go.dev/github.com/lanej/charmed/attention "github.com/lanej/charmed/attention") · [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
rule := attention.Rule{
	Warn:  &attention.Bounds{Upper: attention.Limit(60)},
	Error: &attention.Bounds{Upper: attention.Limit(80)},
}
result := rule.Evaluate(68, true) // Value and availability are separate.

return feedback.Attention(result.Level, palette)
// With no boundaries the rule is Off; unavailable data is Unknown.
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-feedback.html#feedback%2fattention-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/molecules-feedback.html#feedback%2fattention-run)

```sh
go run ./examples/catalog feedback/attention
```

[Example source](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/docs/snippets/feedback_attention.go) · [Implementation](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/feedback/attention.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [feedback/status](https://lanej.io/charmed/atoms-feedback.html#feedback/status)

[Molecules](https://lanej.io/charmed/molecules.html) / Feedback

## [Percentage bar](https://lanej.io/charmed/molecules-feedback.html#feedback%2fpercent)

A native progress bar and exact percentage compose a fixed-width graphical table value on a 0–100% scale.

### [Native preview](https://lanej.io/charmed/molecules-feedback.html#feedback%2fpercent-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌░░░░░░░░  68%
                           n/a
```

Demo · truecolor · 100 × 30 terminal cells

```
Compact percentage bar                                                                              
                                                                                                    
▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌░░░░░░░░  68%
```

Narrow · truecolor · 40 × 16 terminal cells

```
Compact percentage bar                  
                                        
▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌░░░░░░░░  68%
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Compact percentage bar                                                                              
                                                                                                    
▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌░░░░░░░░  68%
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Compact percentage bar                                                                              
                                                                                                    
▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌░░░░░░░░  68%
```

Monochrome · none · 100 × 30 terminal cells

```
Compact percentage bar                                                                              
                                                                                                    
#################--------  68%
```

### [Go usage](https://lanej.io/charmed/molecules-feedback.html#feedback%2fpercent-code)

Imports: [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
measured := feedback.PercentBar(68, true, 30, palette)
unavailable := feedback.PercentBar(0, false, 30, palette)

return layout.Column(measured, unavailable)
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-feedback.html#feedback%2fpercent-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/molecules-feedback.html#feedback%2fpercent-run)

```sh
go run ./examples/catalog feedback/percent
```

[Example source](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/docs/snippets/feedback_percent.go) · [Implementation](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/feedback/percent.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [feedback/progress](https://lanej.io/charmed/atoms-feedback.html#feedback/progress)

[Molecules](https://lanej.io/charmed/molecules.html) / Feedback

## [Evidence summary](https://lanej.io/charmed/molecules-feedback.html#feedback%2fevidence-summary)

Counts per evidence state, worst first. No evidence is never zero healthy.

### [Native preview](https://lanej.io/charmed/molecules-feedback.html#feedback%2fevidence-summary-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
? 1 unknown · ◷ 1 stale · ✓ 2 healthy
```

Demo · truecolor · 64 × 8 terminal cells

```
mixed                                                           
? 1 unknown · ◷ 1 stale · ✓ 2 healthy                           
                                                                
failing                                                         
× 1 down · ! 2 degraded · ✓ 5 healthy                           
                                                                
gaps                                                            
∅ 3 insufficient · ≠ 1 attribution gap · … 2 truncated
```

Narrow · truecolor · 40 × 16 terminal cells

```
mixed                                   
? 1 unknown · ◷ 1 stale · ✓ 2 healthy   
                                        
failing                                 
× 1 down · ! 2 degraded · ✓ 5 healthy   
                                        
gaps                                    
∅ 3 insufficient · ≠ 1 attribution gap  
… 2 truncated                           
                                        
empty                                   
? no evidence
```

256 colors · ansi256 · 64 × 8 terminal cells

```
mixed                                                           
? 1 unknown · ◷ 1 stale · ✓ 2 healthy                           
                                                                
failing                                                         
× 1 down · ! 2 degraded · ✓ 5 healthy                           
                                                                
gaps                                                            
∅ 3 insufficient · ≠ 1 attribution gap · … 2 truncated
```

16 colors · ansi16 · 64 × 8 terminal cells

```
mixed                                                           
? 1 unknown · ◷ 1 stale · ✓ 2 healthy                           
                                                                
failing                                                         
× 1 down · ! 2 degraded · ✓ 5 healthy                           
                                                                
gaps                                                            
∅ 3 insufficient · ≠ 1 attribution gap · … 2 truncated
```

Monochrome · none · 64 × 8 terminal cells

```
mixed                                                           
? 1 unknown * t 1 stale * + 2 healthy                           
                                                                
failing                                                         
x 1 down * ! 2 degraded * + 5 healthy                           
                                                                
gaps                                                            
- 3 insufficient * % 1 attribution gap * . 2 truncated
```

### [Go usage](https://lanej.io/charmed/molecules-feedback.html#feedback%2fevidence-summary-code)

Imports: [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
counts := []feedback.EvidenceCount{
	{State: feedback.EvidenceHealthy, N: 2},
	{State: feedback.EvidenceUnknown, N: 1},
	{State: feedback.EvidenceStale, N: 1},
}

return feedback.SummaryLine(counts, 80, theme.New(theme.Dark))
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-feedback.html#feedback%2fevidence-summary-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/molecules-feedback.html#feedback%2fevidence-summary-run)

```sh
go run ./examples/catalog feedback/evidence-summary
```

[Example source](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/docs/snippets/feedback_evidence_summary.go) · [Implementation](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/feedback/evidence.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [feedback/evidence](https://lanej.io/charmed/atoms-feedback.html#feedback/evidence)

[Molecules](https://lanej.io/charmed/molecules.html) / Feedback

## [Evidence footer](https://lanej.io/charmed/molecules-feedback.html#feedback%2fevidence-footer)

Source, scope, state with observation age, and omission. Missing parts say so; aged evidence keeps its state word.

### [Native preview](https://lanej.io/charmed/molecules-feedback.html#feedback%2fevidence-footer-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
source: sentry · scope: prod ·  ◷ STALE  18m ago
```

Demo · truecolor · 64 × 8 terminal cells

```
fresh                                                           
source: connect · scope: dc1,dc2 ·  ✓ HEALTHY  14s ago          
                                                                
stale                                                           
source: sentry · scope: prod ·  ◷ STALE  18m ago                
                                                                
truncated                                                       
source: logs · scope: last 1h ·  … TRUNCATED  1s ago
```

Narrow · truecolor · 40 × 16 terminal cells

```
fresh                                   
source: connect · scope: dc1,dc2        
 ✓ HEALTHY  14s ago                     
                                        
stale                                   
source: sentry · scope: prod            
 ◷ STALE  18m ago                       
                                        
truncated                               
source: logs · scope: last 1h           
 … TRUNCATED  1s ago                    
! source truncated (4200 omitted)       
                                        
missing                                 
source: ∅ missing · scope: ∅ missing    
 ? UNKNOWN  age unavailable
```

256 colors · ansi256 · 64 × 8 terminal cells

```
fresh                                                           
source: connect · scope: dc1,dc2 ·  ✓ HEALTHY  14s ago          
                                                                
stale                                                           
source: sentry · scope: prod ·  ◷ STALE  18m ago                
                                                                
truncated                                                       
source: logs · scope: last 1h ·  … TRUNCATED  1s ago
```

16 colors · ansi16 · 64 × 8 terminal cells

```
fresh                                                           
source: connect · scope: dc1,dc2 ·  ✓ HEALTHY  14s ago          
                                                                
stale                                                           
source: sentry · scope: prod ·  ◷ STALE  18m ago                
                                                                
truncated                                                       
source: logs · scope: last 1h ·  … TRUNCATED  1s ago
```

Monochrome · none · 64 × 8 terminal cells

```
fresh                                                           
source: connect * scope: dc1,dc2 * [+ HEALTHY] 14s ago          
                                                                
stale                                                           
source: sentry * scope: prod * [t STALE] 18m ago                
                                                                
truncated                                                       
source: logs * scope: last 1h * [. TRUNCATED] 1s ago
```

### [Go usage](https://lanej.io/charmed/molecules-feedback.html#feedback%2fevidence-footer-code)

Imports: [`time`](https://pkg.go.dev/time "time") · [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
now := time.Date(2026, 11, 11, 12, 0, 0, 0, time.UTC)
footer := feedback.EvidenceFooter{
	Source: "sentry", Scope: "prod", State: feedback.EvidenceStale,
	ObservedAt: now.Add(-18 * time.Minute), Now: now, // Explicit clocks keep rendering deterministic.
}

return footer.View(80, theme.New(theme.Dark))
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-feedback.html#feedback%2fevidence-footer-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/molecules-feedback.html#feedback%2fevidence-footer-run)

```sh
go run ./examples/catalog feedback/evidence-footer
```

[Example source](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/docs/snippets/feedback_evidence_footer.go) · [Implementation](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/feedback/evidence.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [feedback/evidence](https://lanej.io/charmed/atoms-feedback.html#feedback/evidence) [text/value-state](https://lanej.io/charmed/atoms-text.html#text/value-state) [text/omission](https://lanej.io/charmed/atoms-text.html#text/omission)
