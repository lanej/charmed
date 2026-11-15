[Atoms](https://lanej.io/charmed/atoms.html)Feedback

# [Feedback](https://lanej.io/charmed/atoms-feedback.html#level-atoms-feedback)

16 components · [All atoms](https://lanej.io/charmed/atoms.html)

- [Spinner](https://lanej.io/charmed/atoms-feedback.html#feedback%2fspinner)
- [Progress](https://lanej.io/charmed/atoms-feedback.html#feedback%2fprogress)
- [Timers](https://lanej.io/charmed/atoms-feedback.html#feedback%2ftime)
- [Log level](https://lanej.io/charmed/atoms-feedback.html#feedback%2flog-level)
- [Status](https://lanej.io/charmed/atoms-feedback.html#feedback%2fstatus)
- [Connection](https://lanej.io/charmed/atoms-feedback.html#feedback%2fconnection)
- [Follow](https://lanej.io/charmed/atoms-feedback.html#feedback%2ffollow)
- [Scrollbar](https://lanej.io/charmed/atoms-feedback.html#feedback%2fscrollbar)
- [Freshness](https://lanej.io/charmed/atoms-feedback.html#feedback%2ffreshness)
- [Evidence state](https://lanej.io/charmed/atoms-feedback.html#feedback%2fevidence)
- [Operation outcome](https://lanej.io/charmed/atoms-feedback.html#feedback%2foutcome)
- [Availability](https://lanej.io/charmed/atoms-feedback.html#feedback%2favailability)
- [Health](https://lanej.io/charmed/atoms-feedback.html#feedback%2fhealth)
- [Readiness](https://lanej.io/charmed/atoms-feedback.html#feedback%2freadiness)
- [Rollout](https://lanej.io/charmed/atoms-feedback.html#feedback%2frollout)
- [Canary membership](https://lanej.io/charmed/atoms-feedback.html#feedback%2fcanary-membership)

[Atoms](https://lanej.io/charmed/atoms.html) / Feedback

## [Spinner](https://lanej.io/charmed/atoms-feedback.html#feedback%2fspinner)

Native spinner; motion can run live or stay paused for evidence.

### [Native preview](https://lanej.io/charmed/atoms-feedback.html#feedback%2fspinner-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
⣾   Processing requests
```

Demo · truecolor · 100 × 30 terminal cells

```
Loading indicator                                                                                   
                                                                                                    
⣾   Processing fixture data                                                                         
                                                                                                    
Live runs the native clock / Ctrl+N advances a frame
```

Narrow · truecolor · 40 × 16 terminal cells

```
Loading indicator                       
                                        
⣾   Processing fixture data             
                                        
Live runs the native clock / Ctrl+N adva
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Loading indicator                                                                                   
                                                                                                    
⣾   Processing fixture data                                                                         
                                                                                                    
Live runs the native clock / Ctrl+N advances a frame
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Loading indicator                                                                                   
                                                                                                    
⣾   Processing fixture data                                                                         
                                                                                                    
Live runs the native clock / Ctrl+N advances a frame
```

Monochrome · none · 100 × 30 terminal cells

```
Loading indicator                                                                                   
                                                                                                    
|  Processing fixture data                                                                          
                                                                                                    
Live runs the native clock / Ctrl+N advances a frame
```

### [Go usage](https://lanej.io/charmed/atoms-feedback.html#feedback%2fspinner-code)

Imports: [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
spinner := feedback.NewSpinner(theme.New(theme.Dark))
// Return spinner.Tick from Init; retain the model and command from spinner.Update(msg).
return spinner.View() + "  Processing requests"
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-feedback.html#feedback%2fspinner-integration)

Return the native spinner tick command from initialization, then retain the model and command returned by Update so animation continues. The parent owns placement and lifecycle; a spinner needs no keyboard focus. Use a static loading indicator when reduced motion is requested.

### [Run this example](https://lanej.io/charmed/atoms-feedback.html#feedback%2fspinner-run)

```sh
go run ./examples/catalog feedback/spinner
```

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/feedback_spinner.go) · [Implementation](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/story/atoms.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Feedback

## [Progress](https://lanej.io/charmed/atoms-feedback.html#feedback%2fprogress)

A themed progress meter with a textual percentage.

### [Native preview](https://lanej.io/charmed/atoms-feedback.html#feedback%2fprogress-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌░░░░░░░░░░░░░  68%
```

Demo · truecolor · 100 × 30 terminal cells

```
Progress                                                                                            
                                                                                                    
▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌░░░░░░░░░░░░░░░░░░░░░░░░░░░░░  68%     
                                                                                                    
Ctrl+N advances by five percent.
```

Narrow · truecolor · 40 × 16 terminal cells

```
Progress                                
                                        
▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌░░░░░░░░░░  68%     
                                        
Ctrl+N advances by five percent.
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Progress                                                                                            
                                                                                                    
▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌░░░░░░░░░░░░░░░░░░░░░░░░░░░░░  68%     
                                                                                                    
Ctrl+N advances by five percent.
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Progress                                                                                            
                                                                                                    
▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌░░░░░░░░░░░░░░░░░░░░░░░░░░░░░  68%     
                                                                                                    
Ctrl+N advances by five percent.
```

Monochrome · none · 100 × 30 terminal cells

```
Progress                                                                                            
                                                                                                    
#############################################################-----------------------------  68%     
                                                                                                    
Ctrl+N advances by five percent.
```

### [Go usage](https://lanej.io/charmed/atoms-feedback.html#feedback%2fprogress-code)

Imports: [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
progress := feedback.NewProgress(40, theme.New(theme.Dark))
completed, total := 68.0, 100.0

return progress.ViewAs(completed/total) + "  68%"
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-feedback.html#feedback%2fprogress-integration)

Use ViewAs for an immediate progress fraction without starting animation. For animated transitions, forward the native percentage-update command and subsequent messages through the application loop. Progress display needs no keyboard focus; keep completion counts in the application and recalculate its width from the parent allocation.

### [Run this example](https://lanej.io/charmed/atoms-feedback.html#feedback%2fprogress-run)

```sh
go run ./examples/catalog feedback/progress
```

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/feedback_progress.go) · [Implementation](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/story/atoms.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Feedback

## [Timers](https://lanej.io/charmed/atoms-feedback.html#feedback%2ftime)

Timer and stopwatch messages continue independently of keyboard focus.

### [Native preview](https://lanej.io/charmed/atoms-feedback.html#feedback%2ftime-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Countdown  1m0s
Stopwatch  0s
```

Demo · truecolor · 100 × 30 terminal cells

```
Time components                                                                                     
                                                                                                    
Countdown  1m0s                                                                                     
Stopwatch  0s                                                                                       
                                                                                                    
Enable Live to run the native timers.
```

Narrow · truecolor · 40 × 16 terminal cells

```
Time components                         
                                        
Countdown  1m0s                         
Stopwatch  0s                           
                                        
Enable Live to run the native timers.
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Time components                                                                                     
                                                                                                    
Countdown  1m0s                                                                                     
Stopwatch  0s                                                                                       
                                                                                                    
Enable Live to run the native timers.
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Time components                                                                                     
                                                                                                    
Countdown  1m0s                                                                                     
Stopwatch  0s                                                                                       
                                                                                                    
Enable Live to run the native timers.
```

Monochrome · none · 100 × 30 terminal cells

```
Time components                                                                                     
                                                                                                    
Countdown  1m0s                                                                                     
Stopwatch  0s                                                                                       
                                                                                                    
Enable Live to run the native timers.
```

### [Go usage](https://lanej.io/charmed/atoms-feedback.html#feedback%2ftime-code)

Imports: [`time`](https://pkg.go.dev/time "time") · [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback")

```go
timer := feedback.NewTimer(time.Minute)
stopwatch := feedback.NewStopwatch()
// Return their Init commands with tea.Batch and forward tick messages to Update.
return "Countdown  " + timer.View() + "\nStopwatch  " + stopwatch.View()
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-feedback.html#feedback%2ftime-integration)

Return the timer and stopwatch Init commands with tea.Batch, then retain each native model and command as tick messages arrive. The owning application decides when to start, stop or reset timekeeping. These displays need no keyboard focus; the parent places their output within its allocation.

### [Run this example](https://lanej.io/charmed/atoms-feedback.html#feedback%2ftime-run)

```sh
go run ./examples/catalog feedback/time
```

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/feedback_time.go) · [Implementation](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/story/atoms.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Feedback

## [Log level](https://lanej.io/charmed/atoms-feedback.html#feedback%2flog-level)

Six log severities from TRACE to FATAL with distinct tones, stable label widths and independent emphasis.

### [Native preview](https://lanej.io/charmed/atoms-feedback.html#feedback%2flog-level-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 INFO   Request complete
 WARN   Retry scheduled 
 ERROR  Upstream timeout
```

Demo · truecolor · 80 × 12 terminal cells

```
Log level atoms / normal                                                        
                                                                                
 TRACE                                                                          
 DEBUG                                                                          
 INFO                                                                           
 WARN                                                                           
 ERROR                                                                          
 FATAL 
```

Narrow · truecolor · 40 × 16 terminal cells

```
Log level atoms / normal                
                                        
 TRACE                                  
 DEBUG                                  
 INFO                                   
 WARN                                   
 ERROR                                  
 FATAL 
```

256 colors · ansi256 · 80 × 12 terminal cells

```
Log level atoms / normal                                                        
                                                                                
 TRACE                                                                          
 DEBUG                                                                          
 INFO                                                                           
 WARN                                                                           
 ERROR                                                                          
 FATAL 
```

16 colors · ansi16 · 80 × 12 terminal cells

```
Log level atoms / normal                                                        
                                                                                
 TRACE                                                                          
 DEBUG                                                                          
 INFO                                                                           
 WARN                                                                           
 ERROR                                                                          
 FATAL 
```

Monochrome · none · 80 × 12 terminal cells

```
Log level atoms / normal                                                        
                                                                                
[TRACE]                                                                         
[DEBUG]                                                                         
[INFO ]                                                                         
[WARN ]                                                                         
[ERROR]                                                                         
[FATAL]
```

Bright · truecolor · 80 × 12 terminal cells

```
Log level atoms / bright                                                        
                                                                                
 TRACE                                                                          
 DEBUG                                                                          
 INFO                                                                           
 WARN                                                                           
 ERROR                                                                          
 FATAL 
```

Normal · truecolor · 80 × 12 terminal cells

```
Log level atoms / normal                                                        
                                                                                
 TRACE                                                                          
 DEBUG                                                                          
 INFO                                                                           
 WARN                                                                           
 ERROR                                                                          
 FATAL 
```

Dim · truecolor · 80 × 12 terminal cells

```
Log level atoms / dim                                                           
                                                                                
 TRACE                                                                          
 DEBUG                                                                          
 INFO                                                                           
 WARN                                                                           
 ERROR                                                                          
 FATAL 
```

### [Go usage](https://lanej.io/charmed/atoms-feedback.html#feedback%2flog-level-code)

Imports: [`log`](https://pkg.go.dev/charm.land/log/v2 "charm.land/log/v2") · [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return layout.Column(
	feedback.LogLevel(log.InfoLevel, text.Normal, palette)+" Request complete",
	feedback.LogLevel(log.WarnLevel, text.Normal, palette)+" Retry scheduled",
	feedback.LogLevel(log.ErrorLevel, text.Bright, palette)+" Upstream timeout",
)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-feedback.html#feedback%2flog-level-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-feedback.html#feedback%2flog-level-run)

```sh
go run ./examples/catalog feedback/log-level
```

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/feedback_log_level.go) · [Implementation](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/feedback/log_level.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [text/pill](https://lanej.io/charmed/atoms-text.html#text/pill)

[Atoms](https://lanej.io/charmed/atoms.html) / Feedback

## [Status](https://lanej.io/charmed/atoms-feedback.html#feedback%2fstatus)

Render general display states without treating them as a health model. For observed availability, health, readiness and production traffic, prefer the distinct typed atom families. text.Compact gives a bare glyph; text.GlyphOnly retains the pill shell.

### [Native preview](https://lanej.io/charmed/atoms-feedback.html#feedback%2fstatus-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 ✓ OK      
 ! WARN    
 ▶ RUNNING 
 ? UNKNOWN 
```

Demo · truecolor · 80 × 14 terminal cells

```
Status atoms / normal                                                           
                                                                                
 ✓ OK                                                                           
 ! WARN                                                                         
 × ERROR                                                                        
 ? UNKNOWN                                                                      
 ⏻ OFF                                                                          
 ◌ INACTIVE                                                                     
 ◷ PENDING                                                                      
 ▶ RUNNING                                                                      
 ‖ PAUSED                                                                       
 ⊘ DISABLED                                                                     
 ● CONNECTED                                                                    
 ○ DISCONNECTED 
```

Narrow · truecolor · 40 × 16 terminal cells

```
Status atoms / normal                   
                                        
 ✓ OK                                   
 ! WARN                                 
 × ERROR                                
 ? UNKNOWN                              
 ⏻ OFF                                  
 ◌ INACTIVE                             
 ◷ PENDING                              
 ▶ RUNNING                              
 ‖ PAUSED                               
 ⊘ DISABLED                             
 ● CONNECTED                            
 ○ DISCONNECTED 
```

256 colors · ansi256 · 80 × 14 terminal cells

```
Status atoms / normal                                                           
                                                                                
 ✓ OK                                                                           
 ! WARN                                                                         
 × ERROR                                                                        
 ? UNKNOWN                                                                      
 ⏻ OFF                                                                          
 ◌ INACTIVE                                                                     
 ◷ PENDING                                                                      
 ▶ RUNNING                                                                      
 ‖ PAUSED                                                                       
 ⊘ DISABLED                                                                     
 ● CONNECTED                                                                    
 ○ DISCONNECTED 
```

16 colors · ansi16 · 80 × 14 terminal cells

```
Status atoms / normal                                                           
                                                                                
 ✓ OK                                                                           
 ! WARN                                                                         
 × ERROR                                                                        
 ? UNKNOWN                                                                      
 ⏻ OFF                                                                          
 ◌ INACTIVE                                                                     
 ◷ PENDING                                                                      
 ▶ RUNNING                                                                      
 ‖ PAUSED                                                                       
 ⊘ DISABLED                                                                     
 ● CONNECTED                                                                    
 ○ DISCONNECTED 
```

Monochrome · none · 80 × 14 terminal cells

```
Status atoms / normal                                                           
                                                                                
[+ OK]                                                                          
[! WARN]                                                                        
[x ERROR]                                                                       
[? UNKNOWN]                                                                     
[- OFF]                                                                         
[- INACTIVE]                                                                    
[p PENDING]                                                                     
[> RUNNING]                                                                     
[= PAUSED]                                                                      
[/ DISABLED]                                                                    
[+ CONNECTED]                                                                   
[- DISCONNECTED]
```

Bright · truecolor · 80 × 14 terminal cells

```
Status atoms / bright                                                           
                                                                                
 ✓ OK                                                                           
 ! WARN                                                                         
 × ERROR                                                                        
 ? UNKNOWN                                                                      
 ⏻ OFF                                                                          
 ◌ INACTIVE                                                                     
 ◷ PENDING                                                                      
 ▶ RUNNING                                                                      
 ‖ PAUSED                                                                       
 ⊘ DISABLED                                                                     
 ● CONNECTED                                                                    
 ○ DISCONNECTED 
```

Normal · truecolor · 80 × 14 terminal cells

```
Status atoms / normal                                                           
                                                                                
 ✓ OK                                                                           
 ! WARN                                                                         
 × ERROR                                                                        
 ? UNKNOWN                                                                      
 ⏻ OFF                                                                          
 ◌ INACTIVE                                                                     
 ◷ PENDING                                                                      
 ▶ RUNNING                                                                      
 ‖ PAUSED                                                                       
 ⊘ DISABLED                                                                     
 ● CONNECTED                                                                    
 ○ DISCONNECTED 
```

Dim · truecolor · 80 × 14 terminal cells

```
Status atoms / dim                                                              
                                                                                
 ✓ OK                                                                           
 ! WARN                                                                         
 × ERROR                                                                        
 ? UNKNOWN                                                                      
 ⏻ OFF                                                                          
 ◌ INACTIVE                                                                     
 ◷ PENDING                                                                      
 ▶ RUNNING                                                                      
 ‖ PAUSED                                                                       
 ⊘ DISABLED                                                                     
 ● CONNECTED                                                                    
 ○ DISCONNECTED 
```

### [Go usage](https://lanej.io/charmed/atoms-feedback.html#feedback%2fstatus-code)

Imports: [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return layout.Column(
	feedback.OK.View(text.Normal, palette),
	feedback.Warn.View(text.Bright, palette),
	feedback.Running.View(text.Normal, palette),
	feedback.Unknown.View(text.Dim, palette),
)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-feedback.html#feedback%2fstatus-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-feedback.html#feedback%2fstatus-run)

```sh
go run ./examples/catalog feedback/status
```

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/feedback_status.go) · [Implementation](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/feedback/status.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

### Reusable component

[Status dependency graph](https://lanej.io/charmed/dependencies.html#component-feedback%2Fstatus) · Partial map · [Source declaration](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/feedback/status.go#L12)

- [Indicator](https://lanej.io/charmed/dependencies.html#component-text%2Findicator) · composes · state presentation

Consumer usage uses

- [Status](https://lanej.io/charmed/dependencies.html#component-feedback%2Fstatus) · primary demonstration

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/feedback_status.go#L10)

[Atoms](https://lanej.io/charmed/atoms.html) / Feedback

## [Connection](https://lanej.io/charmed/atoms-feedback.html#feedback%2fconnection)

Connected and disconnected describe connectivity independently of health, severity or unknown state. Labels and glyphs remain visible in every profile.

### [Native preview](https://lanej.io/charmed/atoms-feedback.html#feedback%2fconnection-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 ● CONNECTED    
 ○ DISCONNECTED 
```

Demo · truecolor · 64 × 8 terminal cells

```
 ● CONNECTED                                                    
 ○ DISCONNECTED 
```

Narrow · truecolor · 40 × 16 terminal cells

```
 ● CONNECTED                            
 ○ DISCONNECTED 
```

256 colors · ansi256 · 64 × 8 terminal cells

```
 ● CONNECTED                                                    
 ○ DISCONNECTED 
```

16 colors · ansi16 · 64 × 8 terminal cells

```
 ● CONNECTED                                                    
 ○ DISCONNECTED 
```

Monochrome · none · 64 × 8 terminal cells

```
[+ CONNECTED]                                                   
[- DISCONNECTED]
```

Bright · truecolor · 64 × 8 terminal cells

```
 ● CONNECTED                                                    
 ○ DISCONNECTED 
```

Normal · truecolor · 64 × 8 terminal cells

```
 ● CONNECTED                                                    
 ○ DISCONNECTED 
```

Dim · truecolor · 64 × 8 terminal cells

```
 ● CONNECTED                                                    
 ○ DISCONNECTED 
```

### [Go usage](https://lanej.io/charmed/atoms-feedback.html#feedback%2fconnection-code)

Imports: [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return layout.Column(
	feedback.Connected.View(text.Normal, palette),
	feedback.Disconnected.View(text.Normal, palette),
)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-feedback.html#feedback%2fconnection-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-feedback.html#feedback%2fconnection-run)

```sh
go run ./examples/catalog feedback/connection
```

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/feedback_connection.go) · [Implementation](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/story/chrome.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [feedback/status](https://lanej.io/charmed/atoms-feedback.html#feedback/status)

[Atoms](https://lanej.io/charmed/atoms.html) / Feedback

## [Follow](https://lanej.io/charmed/atoms-feedback.html#feedback%2ffollow)

Follow mode and paused browsing reuse lifecycle status colors, glyphs and insets.

### [Native preview](https://lanej.io/charmed/atoms-feedback.html#feedback%2ffollow-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 ▶ FOLLOW 
 ‖ PAUSED 
```

Demo · truecolor · 100 × 30 terminal cells

```
Browsing mode / shared status atoms                                                                 
                                                                                                    
 ▶ FOLLOW                                                                                           
 ‖ PAUSED                                                                                           
                                                                                                    
F resume follow
```

Narrow · truecolor · 40 × 16 terminal cells

```
Browsing mode / shared status atoms     
                                        
 ▶ FOLLOW                               
 ‖ PAUSED                               
                                        
F resume follow
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Browsing mode / shared status atoms                                                                 
                                                                                                    
 ▶ FOLLOW                                                                                           
 ‖ PAUSED                                                                                           
                                                                                                    
F resume follow
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Browsing mode / shared status atoms                                                                 
                                                                                                    
 ▶ FOLLOW                                                                                           
 ‖ PAUSED                                                                                           
                                                                                                    
F resume follow
```

Monochrome · none · 100 × 30 terminal cells

```
Browsing mode / shared status atoms                                                                 
                                                                                                    
[> FOLLOW]                                                                                          
[= PAUSED]                                                                                          
                                                                                                    
F resume follow
```

### [Go usage](https://lanej.io/charmed/atoms-feedback.html#feedback%2ffollow-code)

Imports: [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return layout.Column(
	feedback.Follow(true, text.Normal, palette),
	feedback.Follow(false, text.Normal, palette),
)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-feedback.html#feedback%2ffollow-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-feedback.html#feedback%2ffollow-run)

```sh
go run ./examples/catalog feedback/follow
```

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/feedback_follow.go) · [Implementation](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/feedback/controls.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [feedback/status](https://lanej.io/charmed/atoms-feedback.html#feedback/status)

[Atoms](https://lanej.io/charmed/atoms.html) / Feedback

## [Scrollbar](https://lanej.io/charmed/atoms-feedback.html#feedback%2fscrollbar)

Proportional vertical line scrollbar with bounded offsets, ASCII fallback and a stable empty gutter.

### [Native preview](https://lanej.io/charmed/atoms-feedback.html#feedback%2fscrollbar-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
│
│
│
┃
┃
│
│
│
```

Demo · truecolor · 100 × 30 terminal cells

```
Scrollbar / top · middle · bottom                                                                   
                                                                                                    
┃     │     │                                                                                       
│     │     │                                                                                       
│     │     │                                                                                       
│     │     │                                                                                       
│     ┃     │                                                                                       
│     │     │                                                                                       
│     │     │                                                                                       
│     │     │                                                                                       
│     │     │                                                                                       
│     │     ┃                                                                                       
                                                                                                    
Thumb = visible lines / total lines                                                                 
No overflow = reserved blank gutter
```

Narrow · truecolor · 40 × 16 terminal cells

```
Scrollbar / top · middle · bottom       
                                        
┃     │     │                           
│     │     │                           
│     │     │                           
│     │     │                           
│     ┃     │                           
│     │     │                           
│     │     │                           
│     │     │                           
│     │     │                           
│     │     ┃                           
                                        
Thumb = visible lines / total lines     
No overflow = reserved blank gutter
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Scrollbar / top · middle · bottom                                                                   
                                                                                                    
┃     │     │                                                                                       
│     │     │                                                                                       
│     │     │                                                                                       
│     │     │                                                                                       
│     ┃     │                                                                                       
│     │     │                                                                                       
│     │     │                                                                                       
│     │     │                                                                                       
│     │     │                                                                                       
│     │     ┃                                                                                       
                                                                                                    
Thumb = visible lines / total lines                                                                 
No overflow = reserved blank gutter
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Scrollbar / top · middle · bottom                                                                   
                                                                                                    
┃     │     │                                                                                       
│     │     │                                                                                       
│     │     │                                                                                       
│     │     │                                                                                       
│     ┃     │                                                                                       
│     │     │                                                                                       
│     │     │                                                                                       
│     │     │                                                                                       
│     │     │                                                                                       
│     │     ┃                                                                                       
                                                                                                    
Thumb = visible lines / total lines                                                                 
No overflow = reserved blank gutter
```

Monochrome · none · 100 × 30 terminal cells

```
Scrollbar / top * middle * bottom                                                                   
                                                                                                    
#     |     |                                                                                       
|     |     |                                                                                       
|     |     |                                                                                       
|     |     |                                                                                       
|     #     |                                                                                       
|     |     |                                                                                       
|     |     |                                                                                       
|     |     |                                                                                       
|     |     |                                                                                       
|     |     #                                                                                       
                                                                                                    
Thumb = visible lines / total lines                                                                 
No overflow = reserved blank gutter
```

### [Go usage](https://lanej.io/charmed/atoms-feedback.html#feedback%2fscrollbar-code)

Imports: [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
scrollbar := feedback.Scrollbar{Total: 100, Visible: 20, Offset: 35, Height: 8}
return scrollbar.View(theme.New(theme.Dark))
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-feedback.html#feedback%2fscrollbar-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-feedback.html#feedback%2fscrollbar-run)

```sh
go run ./examples/catalog feedback/scrollbar
```

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/feedback_scrollbar.go) · [Implementation](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/feedback/controls.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

### Reusable component

[Scrollbar dependency graph](https://lanej.io/charmed/dependencies.html#component-feedback%2Fscrollbar) · Partial map · [Source declaration](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/feedback/controls.go#L22)

Consumer usage uses

- [Scrollbar](https://lanej.io/charmed/dependencies.html#component-feedback%2Fscrollbar) · primary demonstration

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/feedback_scrollbar.go#L8)

[Atoms](https://lanej.io/charmed/atoms.html) / Feedback

## [Freshness](https://lanej.io/charmed/atoms-feedback.html#feedback%2ffreshness)

Explicit observation and current times with a caller-owned stale threshold; future timestamps stay visible.

### [Native preview](https://lanej.io/charmed/atoms-feedback.html#feedback%2ffreshness-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
◷ STALE 8m0s ago
```

Demo · truecolor · 80 × 18 terminal cells

```
feedback/freshness / normal                                                     
                                                                                
◷ 0s ago                                                                        
◷ 12s ago                                                                       
◷ STALE 8m0s ago                                                                
◷ in 1s (future)                                                                
— age unavailable
```

Narrow · truecolor · 40 × 16 terminal cells

```
feedback/freshness / normal             
                                        
◷ 0s ago                                
◷ 12s ago                               
◷ STALE 8m0s ago                        
◷ in 1s (future)                        
— age unavailable
```

256 colors · ansi256 · 80 × 18 terminal cells

```
feedback/freshness / normal                                                     
                                                                                
◷ 0s ago                                                                        
◷ 12s ago                                                                       
◷ STALE 8m0s ago                                                                
◷ in 1s (future)                                                                
— age unavailable
```

16 colors · ansi16 · 80 × 18 terminal cells

```
feedback/freshness / normal                                                     
                                                                                
◷ 0s ago                                                                        
◷ 12s ago                                                                       
◷ STALE 8m0s ago                                                                
◷ in 1s (future)                                                                
— age unavailable
```

Monochrome · none · 80 × 18 terminal cells

```
feedback/freshness / normal                                                     
                                                                                
0s ago                                                                          
12s ago                                                                         
STALE 8m0s ago                                                                  
in 1s (future)                                                                  
age unavailable
```

Bright · truecolor · 80 × 18 terminal cells

```
feedback/freshness / bright                                                     
                                                                                
◷ 0s ago                                                                        
◷ 12s ago                                                                       
◷ STALE 8m0s ago                                                                
◷ in 1s (future)                                                                
— age unavailable
```

Normal · truecolor · 80 × 18 terminal cells

```
feedback/freshness / normal                                                     
                                                                                
◷ 0s ago                                                                        
◷ 12s ago                                                                       
◷ STALE 8m0s ago                                                                
◷ in 1s (future)                                                                
— age unavailable
```

Dim · truecolor · 80 × 18 terminal cells

```
feedback/freshness / dim                                                        
                                                                                
◷ 0s ago                                                                        
◷ 12s ago                                                                       
◷ STALE 8m0s ago                                                                
◷ in 1s (future)                                                                
— age unavailable
```

### [Go usage](https://lanej.io/charmed/atoms-feedback.html#feedback%2ffreshness-code)

Imports: [`time`](https://pkg.go.dev/time "time") · [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
now := time.Date(2026, 11, 10, 12, 0, 0, 0, time.UTC)
freshness := feedback.Freshness{
	ObservedAt: now.Add(-8 * time.Minute), StaleAfter: time.Minute,
	Emphasis: text.Normal,
}

return freshness.View(now, theme.New(theme.Dark))
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-feedback.html#feedback%2ffreshness-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-feedback.html#feedback%2ffreshness-run)

```sh
go run ./examples/catalog feedback/freshness
```

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/feedback_freshness.go) · [Implementation](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/feedback/freshness.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Feedback

## [Evidence state](https://lanej.io/charmed/atoms-feedback.html#feedback%2fevidence)

Healthy, degraded, down, unknown, stale, insufficient, truncated and attribution-gap evidence; only healthy reads as OK.

### [Native preview](https://lanej.io/charmed/atoms-feedback.html#feedback%2fevidence-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 ◷ STALE 
```

Demo · truecolor · 80 × 18 terminal cells

```
feedback/evidence / normal                                                      
                                                                                
 ✓ HEALTHY  status OK                                                           
 ! DEGRADED  status WARN                                                        
 × DOWN  status ERROR                                                           
 ? UNKNOWN  status UNKNOWN                                                      
 ◷ STALE  status WARN                                                           
 ∅ INSUFFICIENT  status UNKNOWN                                                 
 … TRUNCATED  status WARN                                                       
 ≠ ATTRIBUTION GAP  status UNKNOWN
```

Narrow · truecolor · 40 × 16 terminal cells

```
feedback/evidence / normal              
                                        
 ✓ HEALTHY  status OK                   
 ! DEGRADED  status WARN                
 × DOWN  status ERROR                   
 ? UNKNOWN  status UNKNOWN              
 ◷ STALE  status WARN                   
 ∅ INSUFFICIENT  status UNKNOWN         
 … TRUNCATED  status WARN               
 ≠ ATTRIBUTION GAP  status UNKNOWN
```

256 colors · ansi256 · 80 × 18 terminal cells

```
feedback/evidence / normal                                                      
                                                                                
 ✓ HEALTHY  status OK                                                           
 ! DEGRADED  status WARN                                                        
 × DOWN  status ERROR                                                           
 ? UNKNOWN  status UNKNOWN                                                      
 ◷ STALE  status WARN                                                           
 ∅ INSUFFICIENT  status UNKNOWN                                                 
 … TRUNCATED  status WARN                                                       
 ≠ ATTRIBUTION GAP  status UNKNOWN
```

16 colors · ansi16 · 80 × 18 terminal cells

```
feedback/evidence / normal                                                      
                                                                                
 ✓ HEALTHY  status OK                                                           
 ! DEGRADED  status WARN                                                        
 × DOWN  status ERROR                                                           
 ? UNKNOWN  status UNKNOWN                                                      
 ◷ STALE  status WARN                                                           
 ∅ INSUFFICIENT  status UNKNOWN                                                 
 … TRUNCATED  status WARN                                                       
 ≠ ATTRIBUTION GAP  status UNKNOWN
```

Monochrome · none · 80 × 18 terminal cells

```
feedback/evidence / normal                                                      
                                                                                
[+ HEALTHY] status OK                                                           
[! DEGRADED] status WARN                                                        
[x DOWN] status ERROR                                                           
[? UNKNOWN] status UNKNOWN                                                      
[t STALE] status WARN                                                           
[- INSUFFICIENT] status UNKNOWN                                                 
[. TRUNCATED] status WARN                                                       
[% ATTRIBUTION GAP] status UNKNOWN
```

Bright · truecolor · 80 × 18 terminal cells

```
feedback/evidence / bright                                                      
                                                                                
 ✓ HEALTHY  status OK                                                           
 ! DEGRADED  status WARN                                                        
 × DOWN  status ERROR                                                           
 ? UNKNOWN  status UNKNOWN                                                      
 ◷ STALE  status WARN                                                           
 ∅ INSUFFICIENT  status UNKNOWN                                                 
 … TRUNCATED  status WARN                                                       
 ≠ ATTRIBUTION GAP  status UNKNOWN
```

Normal · truecolor · 80 × 18 terminal cells

```
feedback/evidence / normal                                                      
                                                                                
 ✓ HEALTHY  status OK                                                           
 ! DEGRADED  status WARN                                                        
 × DOWN  status ERROR                                                           
 ? UNKNOWN  status UNKNOWN                                                      
 ◷ STALE  status WARN                                                           
 ∅ INSUFFICIENT  status UNKNOWN                                                 
 … TRUNCATED  status WARN                                                       
 ≠ ATTRIBUTION GAP  status UNKNOWN
```

Dim · truecolor · 80 × 18 terminal cells

```
feedback/evidence / dim                                                         
                                                                                
 ✓ HEALTHY  status OK                                                           
 ! DEGRADED  status WARN                                                        
 × DOWN  status ERROR                                                           
 ? UNKNOWN  status UNKNOWN                                                      
 ◷ STALE  status WARN                                                           
 ∅ INSUFFICIENT  status UNKNOWN                                                 
 … TRUNCATED  status WARN                                                       
 ≠ ATTRIBUTION GAP  status UNKNOWN
```

### [Go usage](https://lanej.io/charmed/atoms-feedback.html#feedback%2fevidence-code)

Imports: [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

// Parse the state word once at the boundary; anything unrecognized is unknown, never healthy.
state, _ := feedback.ParseEvidence("stale")

return state.View(text.Normal, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-feedback.html#feedback%2fevidence-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-feedback.html#feedback%2fevidence-run)

```sh
go run ./examples/catalog feedback/evidence
```

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/feedback_evidence.go) · [Implementation](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/feedback/evidence.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [text/pill](https://lanej.io/charmed/atoms-text.html#text/pill) [feedback/status](https://lanej.io/charmed/atoms-feedback.html#feedback/status)

[Atoms](https://lanej.io/charmed/atoms.html) / Feedback

## [Operation outcome](https://lanej.io/charmed/atoms-feedback.html#feedback%2foutcome)

Succeeded, failed, canceled, timed out and unknown results never infer the outcome of pending work.

### [Native preview](https://lanej.io/charmed/atoms-feedback.html#feedback%2foutcome-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 ◷ TIMED OUT    ⊘ CANCELED 
```

Demo · truecolor · 80 × 18 terminal cells

```
feedback/outcome / normal                                                       
                                                                                
 ? UNKNOWN                                                                      
 ✓ SUCCEEDED                                                                    
 × FAILED                                                                       
 ⊘ CANCELED                                                                     
 ◷ TIMED OUT 
```

Narrow · truecolor · 40 × 16 terminal cells

```
feedback/outcome / normal               
                                        
 ? UNKNOWN                              
 ✓ SUCCEEDED                            
 × FAILED                               
 ⊘ CANCELED                             
 ◷ TIMED OUT 
```

256 colors · ansi256 · 80 × 18 terminal cells

```
feedback/outcome / normal                                                       
                                                                                
 ? UNKNOWN                                                                      
 ✓ SUCCEEDED                                                                    
 × FAILED                                                                       
 ⊘ CANCELED                                                                     
 ◷ TIMED OUT 
```

16 colors · ansi16 · 80 × 18 terminal cells

```
feedback/outcome / normal                                                       
                                                                                
 ? UNKNOWN                                                                      
 ✓ SUCCEEDED                                                                    
 × FAILED                                                                       
 ⊘ CANCELED                                                                     
 ◷ TIMED OUT 
```

Monochrome · none · 80 × 18 terminal cells

```
feedback/outcome / normal                                                       
                                                                                
[? UNKNOWN]                                                                     
[+ SUCCEEDED]                                                                   
[x FAILED]                                                                      
[/ CANCELED]                                                                    
[t TIMED OUT]
```

Bright · truecolor · 80 × 18 terminal cells

```
feedback/outcome / bright                                                       
                                                                                
 ? UNKNOWN                                                                      
 ✓ SUCCEEDED                                                                    
 × FAILED                                                                       
 ⊘ CANCELED                                                                     
 ◷ TIMED OUT 
```

Normal · truecolor · 80 × 18 terminal cells

```
feedback/outcome / normal                                                       
                                                                                
 ? UNKNOWN                                                                      
 ✓ SUCCEEDED                                                                    
 × FAILED                                                                       
 ⊘ CANCELED                                                                     
 ◷ TIMED OUT 
```

Dim · truecolor · 80 × 18 terminal cells

```
feedback/outcome / dim                                                          
                                                                                
 ? UNKNOWN                                                                      
 ✓ SUCCEEDED                                                                    
 × FAILED                                                                       
 ⊘ CANCELED                                                                     
 ◷ TIMED OUT 
```

### [Go usage](https://lanej.io/charmed/atoms-feedback.html#feedback%2foutcome-code)

Imports: [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
// Timeout describes a reported result; it does not establish rollback or absence of effects.
return feedback.TimedOut.View(text.Normal, palette) + "  " +
	feedback.Canceled.View(text.Normal, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-feedback.html#feedback%2foutcome-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-feedback.html#feedback%2foutcome-run)

```sh
go run ./examples/catalog feedback/outcome
```

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/feedback_outcome.go) · [Implementation](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/feedback/outcome.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [text/pill](https://lanej.io/charmed/atoms-text.html#text/pill)

[Atoms](https://lanej.io/charmed/atoms.html) / Feedback

## [Availability](https://lanej.io/charmed/atoms-feedback.html#feedback%2favailability)

Up and Down describe supplied reachability without asserting readiness, health or traffic. Normal views are labeled pills; text.Compact displays only the colored glyph.

### [Native preview](https://lanej.io/charmed/atoms-feedback.html#feedback%2favailability-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Pill           Compact
                      
 ↑ Up          ↑      
 ↓ Down        ↓      
 ? Unknown     ?
```

Demo · truecolor · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 ↑ Up               ↑                                                           
 ↓ Down             ↓                                                           
 ? Unknown          ?
```

Narrow · truecolor · 40 × 16 terminal cells

```
Pill                Compact             
                                        
 ↑ Up               ↑                   
 ↓ Down             ↓                   
 ? Unknown          ?
```

256 colors · ansi256 · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 ↑ Up               ↑                                                           
 ↓ Down             ↓                                                           
 ? Unknown          ?
```

16 colors · ansi16 · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 ↑ Up               ↑                                                           
 ↓ Down             ↓                                                           
 ? Unknown          ?
```

Monochrome · none · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
[^ Up]              ^                                                           
[v Down]            v                                                           
[? Unknown]         ?
```

Bright · truecolor · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 ↑ Up               ↑                                                           
 ↓ Down             ↓                                                           
 ? Unknown          ?
```

Normal · truecolor · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 ↑ Up               ↑                                                           
 ↓ Down             ↓                                                           
 ? Unknown          ?
```

Dim · truecolor · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 ↑ Up               ↑                                                           
 ↓ Down             ↓                                                           
 ? Unknown          ?
```

### [Go usage](https://lanej.io/charmed/atoms-feedback.html#feedback%2favailability-code)

Imports: [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return layout.Row(
	layout.Column(
		"Pill",
		"",
		feedback.Up.View(text.Normal, palette),
		feedback.Down.View(text.Normal, palette),
		feedback.AvailabilityUnknown.View(text.Normal, palette),
	),
	"    ",
	layout.Column(
		"Compact",
		"",
		feedback.Up.View(text.Normal, palette, text.Compact),
		feedback.Down.View(text.Normal, palette, text.Compact),
		feedback.AvailabilityUnknown.View(text.Normal, palette, text.Compact),
	),
)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-feedback.html#feedback%2favailability-integration)

Supply Availability values independently of the other state families. Zero and unsupported values render Unknown. Pass text.Compact for a bare glyph without fill or padding; retain a nearby label or legend when context does not identify the state. text.GlyphOnly keeps the existing padded pill shape. Rendering is static and owns no checks, traffic policy, focus or event loop.

### [Run this example](https://lanej.io/charmed/atoms-feedback.html#feedback%2favailability-run)

```sh
go run ./examples/catalog feedback/availability
```

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/feedback_availability.go) · [Implementation](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/feedback/availability.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

### Reusable component

[Availability dependency graph](https://lanej.io/charmed/dependencies.html#component-feedback%2Favailability) · Partial map · [Source declaration](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/feedback/availability.go#L9)

- [Indicator](https://lanej.io/charmed/dependencies.html#component-text%2Findicator) · composes · state presentation

Consumer usage uses

- [Availability](https://lanej.io/charmed/dependencies.html#component-feedback%2Favailability) · primary demonstration

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/feedback_availability.go#L10)

[Atoms](https://lanej.io/charmed/atoms.html) / Feedback

## [Health](https://lanej.io/charmed/atoms-feedback.html#feedback%2fhealth)

Healthy and Unhealthy describe health-check results independently of availability and rollout. Normal views are labeled pills; text.Compact displays only the colored glyph.

### [Native preview](https://lanej.io/charmed/atoms-feedback.html#feedback%2fhealth-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Pill             Compact
                        
 ♥ Healthy       ♥      
 × Unhealthy     ×      
 ? Unknown       ?
```

Demo · truecolor · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 ♥ Healthy          ♥                                                           
 × Unhealthy        ×                                                           
 ? Unknown          ?
```

Narrow · truecolor · 40 × 16 terminal cells

```
Pill                Compact             
                                        
 ♥ Healthy          ♥                   
 × Unhealthy        ×                   
 ? Unknown          ?
```

256 colors · ansi256 · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 ♥ Healthy          ♥                                                           
 × Unhealthy        ×                                                           
 ? Unknown          ?
```

16 colors · ansi16 · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 ♥ Healthy          ♥                                                           
 × Unhealthy        ×                                                           
 ? Unknown          ?
```

Monochrome · none · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
[H Healthy]         H                                                           
[x Unhealthy]       x                                                           
[? Unknown]         ?
```

Bright · truecolor · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 ♥ Healthy          ♥                                                           
 × Unhealthy        ×                                                           
 ? Unknown          ?
```

Normal · truecolor · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 ♥ Healthy          ♥                                                           
 × Unhealthy        ×                                                           
 ? Unknown          ?
```

Dim · truecolor · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 ♥ Healthy          ♥                                                           
 × Unhealthy        ×                                                           
 ? Unknown          ?
```

### [Go usage](https://lanej.io/charmed/atoms-feedback.html#feedback%2fhealth-code)

Imports: [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return layout.Row(
	layout.Column(
		"Pill",
		"",
		feedback.Healthy.View(text.Normal, palette),
		feedback.Unhealthy.View(text.Normal, palette),
		feedback.HealthUnknown.View(text.Normal, palette),
	),
	"    ",
	layout.Column(
		"Compact",
		"",
		feedback.Healthy.View(text.Normal, palette, text.Compact),
		feedback.Unhealthy.View(text.Normal, palette, text.Compact),
		feedback.HealthUnknown.View(text.Normal, palette, text.Compact),
	),
)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-feedback.html#feedback%2fhealth-integration)

Supply Health values independently of the other state families. Zero and unsupported values render Unknown. Pass text.Compact for a bare glyph without fill or padding; retain a nearby label or legend when context does not identify the state. text.GlyphOnly keeps the existing padded pill shape. Rendering is static and owns no checks, traffic policy, focus or event loop.

### [Run this example](https://lanej.io/charmed/atoms-feedback.html#feedback%2fhealth-run)

```sh
go run ./examples/catalog feedback/health
```

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/feedback_health.go) · [Implementation](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/feedback/health.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

### Reusable component

[Health dependency graph](https://lanej.io/charmed/dependencies.html#component-feedback%2Fhealth) · Partial map · [Source declaration](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/feedback/health.go#L9)

- [Indicator](https://lanej.io/charmed/dependencies.html#component-text%2Findicator) · composes · state presentation

Consumer usage uses

- [Health](https://lanej.io/charmed/dependencies.html#component-feedback%2Fhealth) · primary demonstration

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/feedback_health.go#L10)

[Atoms](https://lanej.io/charmed/atoms.html) / Feedback

## [Readiness](https://lanej.io/charmed/atoms-feedback.html#feedback%2freadiness)

Ready means the target may accept work; Not ready does not imply unhealthy or down. Normal views are labeled pills; text.Compact displays only the colored glyph.

### [Native preview](https://lanej.io/charmed/atoms-feedback.html#feedback%2freadiness-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Pill             Compact
                        
 ✓ Ready         ✓      
 ◇ Not ready     ◇      
 ? Unknown       ?
```

Demo · truecolor · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 ✓ Ready            ✓                                                           
 ◇ Not ready        ◇                                                           
 ? Unknown          ?
```

Narrow · truecolor · 40 × 16 terminal cells

```
Pill                Compact             
                                        
 ✓ Ready            ✓                   
 ◇ Not ready        ◇                   
 ? Unknown          ?
```

256 colors · ansi256 · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 ✓ Ready            ✓                                                           
 ◇ Not ready        ◇                                                           
 ? Unknown          ?
```

16 colors · ansi16 · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 ✓ Ready            ✓                                                           
 ◇ Not ready        ◇                                                           
 ? Unknown          ?
```

Monochrome · none · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
[r Ready]           r                                                           
[- Not ready]       -                                                           
[? Unknown]         ?
```

Bright · truecolor · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 ✓ Ready            ✓                                                           
 ◇ Not ready        ◇                                                           
 ? Unknown          ?
```

Normal · truecolor · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 ✓ Ready            ✓                                                           
 ◇ Not ready        ◇                                                           
 ? Unknown          ?
```

Dim · truecolor · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 ✓ Ready            ✓                                                           
 ◇ Not ready        ◇                                                           
 ? Unknown          ?
```

### [Go usage](https://lanej.io/charmed/atoms-feedback.html#feedback%2freadiness-code)

Imports: [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return layout.Row(
	layout.Column(
		"Pill",
		"",
		feedback.Ready.View(text.Normal, palette),
		feedback.NotReady.View(text.Normal, palette),
		feedback.ReadinessUnknown.View(text.Normal, palette),
	),
	"    ",
	layout.Column(
		"Compact",
		"",
		feedback.Ready.View(text.Normal, palette, text.Compact),
		feedback.NotReady.View(text.Normal, palette, text.Compact),
		feedback.ReadinessUnknown.View(text.Normal, palette, text.Compact),
	),
)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-feedback.html#feedback%2freadiness-integration)

Supply Readiness values independently of the other state families. Zero and unsupported values render Unknown. Pass text.Compact for a bare glyph without fill or padding; retain a nearby label or legend when context does not identify the state. text.GlyphOnly keeps the existing padded pill shape. Rendering is static and owns no checks, traffic policy, focus or event loop.

### [Run this example](https://lanej.io/charmed/atoms-feedback.html#feedback%2freadiness-run)

```sh
go run ./examples/catalog feedback/readiness
```

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/feedback_readiness.go) · [Implementation](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/feedback/readiness.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

### Reusable component

[Readiness dependency graph](https://lanej.io/charmed/dependencies.html#component-feedback%2Freadiness) · Partial map · [Source declaration](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/feedback/readiness.go#L9)

- [Indicator](https://lanej.io/charmed/dependencies.html#component-text%2Findicator) · composes · state presentation

Consumer usage uses

- [Readiness](https://lanej.io/charmed/dependencies.html#component-feedback%2Freadiness) · primary demonstration

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/feedback_readiness.go#L10)

[Atoms](https://lanej.io/charmed/atoms.html) / Feedback

## [Rollout](https://lanej.io/charmed/atoms-feedback.html#feedback%2frollout)

Canary identifies the limited-rollout cohort; its pool membership is shown separately by CanaryOn or CanaryOff. Live denotes production traffic. Neither asserts health or means log-follow mode. Normal views are labeled pills; text.Compact displays only the colored glyph.

### [Native preview](https://lanej.io/charmed/atoms-feedback.html#feedback%2frollout-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Pill           Compact
                      
 🐤 Canary     ◔      
 ● Live        ●      
 ? Unknown     ?
```

Demo · truecolor · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 🐤 Canary          ◔                                                           
 ● Live             ●                                                           
 ? Unknown          ?
```

Narrow · truecolor · 40 × 16 terminal cells

```
Pill                Compact             
                                        
 🐤 Canary          ◔                   
 ● Live             ●                   
 ? Unknown          ?
```

256 colors · ansi256 · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 🐤 Canary          ◔                                                           
 ● Live             ●                                                           
 ? Unknown          ?
```

16 colors · ansi16 · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 🐤 Canary          ◔                                                           
 ● Live             ●                                                           
 ? Unknown          ?
```

Monochrome · none · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
[c Canary]          c                                                           
[l Live]            l                                                           
[? Unknown]         ?
```

Bright · truecolor · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 🐤 Canary          ◔                                                           
 ● Live             ●                                                           
 ? Unknown          ?
```

Normal · truecolor · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 🐤 Canary          ◔                                                           
 ● Live             ●                                                           
 ? Unknown          ?
```

Dim · truecolor · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 🐤 Canary          ◔                                                           
 ● Live             ●                                                           
 ? Unknown          ?
```

### [Go usage](https://lanej.io/charmed/atoms-feedback.html#feedback%2frollout-code)

Imports: [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return layout.Row(
	layout.Column(
		"Pill",
		"",
		feedback.Canary.View(text.Normal, palette),
		feedback.Live.View(text.Normal, palette),
		feedback.RolloutUnknown.View(text.Normal, palette),
	),
	"    ",
	layout.Column(
		"Compact",
		"",
		feedback.Canary.View(text.Normal, palette, text.Compact),
		feedback.Live.View(text.Normal, palette, text.Compact),
		feedback.RolloutUnknown.View(text.Normal, palette, text.Compact),
	),
)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-feedback.html#feedback%2frollout-integration)

Supply Rollout values independently of the other state families. Zero and unsupported values render Unknown. Pass text.Compact for a bare glyph without fill or padding; retain a nearby label or legend when context does not identify the state. text.GlyphOnly keeps the existing padded pill shape. Rendering is static and owns no checks, traffic policy, focus or event loop. Canary uses a two-cell bird in full pills and a one-cell mark in Compact; explicit ASCII keeps c.

### [Run this example](https://lanej.io/charmed/atoms-feedback.html#feedback%2frollout-run)

```sh
go run ./examples/catalog feedback/rollout
```

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/feedback_rollout.go) · [Implementation](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/feedback/rollout.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

### Reusable component

[Rollout dependency graph](https://lanej.io/charmed/dependencies.html#component-feedback%2Frollout) · Partial map · [Source declaration](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/feedback/rollout.go#L10)

- [Indicator](https://lanej.io/charmed/dependencies.html#component-text%2Findicator) · composes · state presentation

Consumer usage uses

- [Rollout](https://lanej.io/charmed/dependencies.html#component-feedback%2Frollout) · primary demonstration

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/feedback_rollout.go#L10)

[Atoms](https://lanej.io/charmed/atoms.html) / Feedback

## [Canary membership](https://lanej.io/charmed/atoms-feedback.html#feedback%2fcanary-membership)

Canary on means the new-version target is in the serving pool; Canary off means it is out of the serving pool. Pool membership remains independent of health, readiness and reachability. Both states retain distinct compact glyphs. Full pills use 🐤🟢 for in-pool and 🐤⚪ for out-of-pool; green describes membership, independently of health.

### [Native preview](https://lanej.io/charmed/atoms-feedback.html#feedback%2fcanary-membership-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Pill                 Compact
                            
 🐤🟢 Canary on      ◉      
 🐤⚪️ Canary off     ◌      
 ? Unknown           ?
```

Demo · truecolor · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 🐤🟢 Canary on     ◉                                                           
 🐤⚪️ Canary off    ◌                                                           
 ? Unknown          ?
```

Narrow · truecolor · 40 × 16 terminal cells

```
Pill                Compact             
                                        
 🐤🟢 Canary on     ◉                   
 🐤⚪️ Canary off    ◌                   
 ? Unknown          ?
```

256 colors · ansi256 · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 🐤🟢 Canary on     ◉                                                           
 🐤⚪️ Canary off    ◌                                                           
 ? Unknown          ?
```

16 colors · ansi16 · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 🐤🟢 Canary on     ◉                                                           
 🐤⚪️ Canary off    ◌                                                           
 ? Unknown          ?
```

Monochrome · none · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
[+ Canary on]       +                                                           
[- Canary off]      -                                                           
[? Unknown]         ?
```

Bright · truecolor · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 🐤🟢 Canary on     ◉                                                           
 🐤⚪️ Canary off    ◌                                                           
 ? Unknown          ?
```

Normal · truecolor · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 🐤🟢 Canary on     ◉                                                           
 🐤⚪️ Canary off    ◌                                                           
 ? Unknown          ?
```

Dim · truecolor · 80 × 12 terminal cells

```
Pill                Compact                                                     
                                                                                
 🐤🟢 Canary on     ◉                                                           
 🐤⚪️ Canary off    ◌                                                           
 ? Unknown          ?
```

### [Go usage](https://lanej.io/charmed/atoms-feedback.html#feedback%2fcanary-membership-code)

Imports: [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return layout.Row(
	layout.Column(
		"Pill",
		"",
		feedback.CanaryOn.View(text.Normal, palette),
		feedback.CanaryOff.View(text.Normal, palette),
		feedback.CanaryMembershipUnknown.View(text.Normal, palette),
	),
	"    ",
	layout.Column(
		"Compact",
		"",
		feedback.CanaryOn.View(text.Normal, palette, text.Compact),
		feedback.CanaryOff.View(text.Normal, palette, text.Compact),
		feedback.CanaryMembershipUnknown.View(text.Normal, palette, text.Compact),
	),
)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-feedback.html#feedback%2fcanary-membership-integration)

Supply CanaryOn or CanaryOff from observed pool membership; rendering does not change routing or infer whether a target is healthy. Canary identifies the cohort even while its members are removed. Zero and unsupported values render Unknown. Pass text.Compact for a bare glyph, or keep the default labeled pill. Keep a nearby label or legend for compact marks. Full Unicode marks occupy four cells; Compact retains the single-cell marks. Explicit ASCII keeps the existing + and - fallbacks.

### [Run this example](https://lanej.io/charmed/atoms-feedback.html#feedback%2fcanary-membership-run)

```sh
go run ./examples/catalog feedback/canary-membership
```

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/feedback_canary_membership.go) · [Implementation](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/feedback/canary_membership.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

### Reusable component

[Canary membership dependency graph](https://lanej.io/charmed/dependencies.html#component-feedback%2Fcanary-membership) · Partial map · [Source declaration](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/feedback/canary_membership.go#L11)

- [Indicator](https://lanej.io/charmed/dependencies.html#component-text%2Findicator) · composes · pool membership presentation

Consumer usage uses

- [Canary membership](https://lanej.io/charmed/dependencies.html#component-feedback%2Fcanary-membership) · primary demonstration

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/feedback_canary_membership.go#L10)
