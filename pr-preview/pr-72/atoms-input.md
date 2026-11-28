[Atoms](https://lanej.io/charmed/atoms.html)Input

# [Input](https://lanej.io/charmed/atoms-input.html#level-atoms-input)

6 components · [All atoms](https://lanej.io/charmed/atoms.html)

- [Text](https://lanej.io/charmed/atoms-input.html#input%2ftext)
- [Textarea](https://lanej.io/charmed/atoms-input.html#input%2ftextarea)
- [Checkbox](https://lanej.io/charmed/atoms-input.html#input%2fcheckbox)
- [Radio](https://lanej.io/charmed/atoms-input.html#input%2fradio)
- [Toggle](https://lanej.io/charmed/atoms-input.html#input%2ftoggle)
- [Action button](https://lanej.io/charmed/atoms-input.html#input%2fbutton)

[Atoms](https://lanej.io/charmed/atoms.html) / Input

## [Text](https://lanej.io/charmed/atoms-input.html#input%2ftext)

Collect a single editable value such as an environment name. Use a labeled field when the value also needs persistent help or validation feedback.

### [Native preview](https://lanej.io/charmed/atoms-input.html#input%2ftext-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
┌ Value ───────────────────────────────┐
│ > production                         │
└──────────────────────────────────────┘
```

Demo · truecolor · 100 × 30 terminal cells

```
Text input                                                                                          
                                                                                                    
┌ Value ───────────────────────────────────────────────────────────────────────────────────────────┐
│ > production                                                                                     │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
                                                                                                    
Type to edit / Enter emits Submitted / Ctrl+C exits
```

Narrow · truecolor · 40 × 16 terminal cells

```
Text input                              
                                        
┌ Value ───────────────────────────────┐
│ > production                         │
└──────────────────────────────────────┘
                                        
Type to edit / Enter emits Submitted / C
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Text input                                                                                          
                                                                                                    
┌ Value ───────────────────────────────────────────────────────────────────────────────────────────┐
│ > production                                                                                     │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
                                                                                                    
Type to edit / Enter emits Submitted / Ctrl+C exits
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Text input                                                                                          
                                                                                                    
┌ Value ───────────────────────────────────────────────────────────────────────────────────────────┐
│ > production                                                                                     │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
                                                                                                    
Type to edit / Enter emits Submitted / Ctrl+C exits
```

Monochrome · none · 100 × 30 terminal cells

```
Text input                                                                                          
                                                                                                    
+ Value -------------------------------------------------------------------------------------------+
| > production                                                                                     |
+--------------------------------------------------------------------------------------------------+
                                                                                                    
Type to edit / Enter emits Submitted / Ctrl+C exits
```

Empty · truecolor · 100 × 30 terminal cells

```
Text input                                                                                          
                                                                                                    
┌ Value ───────────────────────────────────────────────────────────────────────────────────────────┐
│ > Type a value                                                                                   │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
                                                                                                    
Type to edit / Enter emits Submitted / Ctrl+C exits
```

Disabled · truecolor · 100 × 30 terminal cells

```
Text input                                                                                          
                                                                                                    
┌ Value ───────────────────────────────────────────────────────────────────────────────────────────┐
│ ⊘ production                                                                                     │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
                                                                                                    
Type to edit / Enter emits Submitted / Ctrl+C exits
```

### [Go usage](https://lanej.io/charmed/atoms-input.html#input%2ftext-code)

Imports: [`input`](https://pkg.go.dev/github.com/lanej/charmed/input "github.com/lanej/charmed/input") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
field := input.NewText(input.TextOptions{
	Placeholder: "Type a value", Value: "production",
	Width: 36, Theme: palette,
})
// In Init, return field.Focus(); in Update, retain field.Update(msg)'s model and command.
return layout.Panel("Value", field.View(),
	layout.Size{Width: 40, Height: 3}, palette, false)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-input.html#input%2ftext-integration)

Return Focus's command when the input becomes active. In Update, retain both the returned model and command; blur inactive controls before routing keys. Call Resize with the allocated column count. SetTheme reapplies shared styles, so apply native style overrides after a theme change.

### [Run this example](https://lanej.io/charmed/atoms-input.html#input%2ftext-run)

```sh
go run ./examples/catalog input/text
```

[Example source](https://github.com/lanej/charmed/blob/e7987b1035f1f0309d8f643ca46001269083830a/docs/snippets/input_text.go) · [Implementation](https://github.com/lanej/charmed/blob/e7987b1035f1f0309d8f643ca46001269083830a/story/atoms.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Input

## [Textarea](https://lanej.io/charmed/atoms-input.html#input%2ftextarea)

Multiline input with native editing and bounded dimensions.

### [Native preview](https://lanej.io/charmed/atoms-input.html#input%2ftextarea-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
┃   1 Deployment notes                          
┃   2 Ready for production.                     
┃                                               
┃                                               
┃                                               
┃                                               
┃                                               
┃                                           
```

Demo · truecolor · 100 × 30 terminal cells

```
Multiline input                                                                                     
                                                                                                    
┃   1 production                                                                                    
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                               
```

Narrow · truecolor · 40 × 16 terminal cells

```
Multiline input                         
                                        
┃   1 production                        
┃                                       
┃                                       
┃                                       
┃                                       
┃                                       
┃                                       
┃                                       
┃                                       
┃                                       
┃                                   
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Multiline input                                                                                     
                                                                                                    
┃   1 production                                                                                    
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                               
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Multiline input                                                                                     
                                                                                                    
┃   1 production                                                                                    
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                                   
┃                                                                                               
```

Monochrome · none · 100 × 30 terminal cells

```
Multiline input                                                                                     
                                                                                                    
|   1 production                                                                                    
|                                                                                                   
|                                                                                                   
|                                                                                                   
|                                                                                                   
|                                                                                                   
|                                                                                                   
|                                                                                                   
|                                                                                                   
|                                                                                                   
|                                                                                                   
|                                                                                                   
|                                                                                                   
|                                                                                                   
|                                                                                                   
|                                                                                                   
|                                                                                                   
|                                                                                                   
|                                                                                                   
|                                                                                                   
|                                                                                                   
|                                                                                                   
|                                                                                                   
|                                                                                                   
|                                                                                               
```

### [Go usage](https://lanej.io/charmed/atoms-input.html#input%2ftextarea-code)

Imports: [`input`](https://pkg.go.dev/github.com/lanej/charmed/input "github.com/lanej/charmed/input") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
editor := input.NewArea(48, 8, theme.New(theme.Dark))
editor.Placeholder = "Write a note"
editor.SetValue("Deployment notes\nReady for production.")
// Return editor.Focus() from Init and forward messages through editor.Update(msg).
return editor.View()
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-input.html#input%2ftextarea-integration)

Focus the native textarea when its region becomes active, and retain the model and command returned by Update. SetWidth and SetHeight from the parent allocation on resize. Configure the exposed native styles and key map directly; the wrapper does not own application validation or submission.

### [Run this example](https://lanej.io/charmed/atoms-input.html#input%2ftextarea-run)

```sh
go run ./examples/catalog input/textarea
```

[Example source](https://github.com/lanej/charmed/blob/e7987b1035f1f0309d8f643ca46001269083830a/docs/snippets/input_textarea.go) · [Implementation](https://github.com/lanej/charmed/blob/e7987b1035f1f0309d8f643ca46001269083830a/story/atoms.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Input

## [Checkbox](https://lanej.io/charmed/atoms-input.html#input%2fcheckbox)

Unchecked, checked and mixed selection with focus, native activation bindings and typed change messages.

### [Native preview](https://lanej.io/charmed/atoms-input.html#input%2fcheckbox-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
  ▣  All services 
```

Demo · truecolor · 80 × 18 terminal cells

```
input/checkbox                                                                  
                                                                                
  ▣  All services                                                               
                                                                                
  ☐  Services                                                                   
  ☑  Services                                                                   
  ▣  Services                                                                   
  ☑  Restricted ⊘ disabled Read-only                                            
                                                                                
Space / Enter activates
```

Narrow · truecolor · 40 × 16 terminal cells

```
input/checkbox                          
                                        
  ▣  All services                       
                                        
  ☐  Services                           
  ☑  Services                           
  ▣  Services                           
  ☑  Restricted ⊘ disabled Read-only    
                                        
Space / Enter activates
```

256 colors · ansi256 · 80 × 18 terminal cells

```
input/checkbox                                                                  
                                                                                
  ▣  All services                                                               
                                                                                
  ☐  Services                                                                   
  ☑  Services                                                                   
  ▣  Services                                                                   
  ☑  Restricted ⊘ disabled Read-only                                            
                                                                                
Space / Enter activates
```

16 colors · ansi16 · 80 × 18 terminal cells

```
input/checkbox                                                                  
                                                                                
  ▣  All services                                                               
                                                                                
  ☐  Services                                                                   
  ☑  Services                                                                   
  ▣  Services                                                                   
  ☑  Restricted ⊘ disabled Read-only                                            
                                                                                
Space / Enter activates
```

Monochrome · none · 80 × 18 terminal cells

```
input/checkbox                                                                  
                                                                                
 [-] All services                                                               
                                                                                
 [ ] Services                                                                   
 [x] Services                                                                   
 [-] Services                                                                   
 [x] Restricted [DISABLED] Read-only                                            
                                                                                
Space / Enter activates
```

### [Go usage](https://lanej.io/charmed/atoms-input.html#input%2fcheckbox-code)

Imports: [`input`](https://pkg.go.dev/github.com/lanej/charmed/input "github.com/lanej/charmed/input") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
checkbox := input.NewCheckbox("services", "All services", input.Mixed, palette)
checkbox.Focus()
// Retain checkbox.Update(msg)'s model and command; handle input.CheckChanged in the parent.
return checkbox.View()
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-input.html#input%2fcheckbox-integration)

The parent routes keys to the focused control and retains the model and command returned by Update. Use the stable control ID when handling its activation or change message. The control measures its own glyph and label; reserve that space in the layout.

### [Run this example](https://lanej.io/charmed/atoms-input.html#input%2fcheckbox-run)

```sh
go run ./examples/catalog input/checkbox
```

[Example source](https://github.com/lanej/charmed/blob/e7987b1035f1f0309d8f643ca46001269083830a/docs/snippets/input_checkbox.go) · [Implementation](https://github.com/lanej/charmed/blob/e7987b1035f1f0309d8f643ca46001269083830a/input/controls.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Input

## [Radio](https://lanej.io/charmed/atoms-input.html#input%2fradio)

Single-choice selection; the owning parent maintains group exclusivity.

### [Native preview](https://lanej.io/charmed/atoms-input.html#input%2fradio-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
  ◉  Production     ○  Staging
```

Demo · truecolor · 80 × 18 terminal cells

```
input/radio                                                                     
                                                                                
  ○  Production                                                                 
  ◉  Staging                                                                    
  ○  Development ⊘ disabled Unavailable                                         
                                                                                
Space / Enter activates
```

Narrow · truecolor · 40 × 16 terminal cells

```
input/radio                             
                                        
  ○  Production                         
  ◉  Staging                            
  ○  Development ⊘ disabled Unavailable 
                                        
Space / Enter activates
```

256 colors · ansi256 · 80 × 18 terminal cells

```
input/radio                                                                     
                                                                                
  ○  Production                                                                 
  ◉  Staging                                                                    
  ○  Development ⊘ disabled Unavailable                                         
                                                                                
Space / Enter activates
```

16 colors · ansi16 · 80 × 18 terminal cells

```
input/radio                                                                     
                                                                                
  ○  Production                                                                 
  ◉  Staging                                                                    
  ○  Development ⊘ disabled Unavailable                                         
                                                                                
Space / Enter activates
```

Monochrome · none · 80 × 18 terminal cells

```
input/radio                                                                     
                                                                                
 ( ) Production                                                                 
 (*) Staging                                                                    
 ( ) Development [DISABLED] Unavailable                                         
                                                                                
Space / Enter activates
```

### [Go usage](https://lanej.io/charmed/atoms-input.html#input%2fradio-code)

Imports: [`input`](https://pkg.go.dev/github.com/lanej/charmed/input "github.com/lanej/charmed/input") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
production := input.NewRadio("production", "Production", true, palette)
staging := input.NewRadio("staging", "Staging", false, palette)
// On input.RadioSelected, the parent deselects every other radio in this group.
return production.View() + "  " + staging.View()
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-input.html#input%2fradio-integration)

The parent routes keys to the focused control and retains the model and command returned by Update. Use the stable control ID when handling its activation or change message. The control measures its own glyph and label; reserve that space in the layout.

### [Run this example](https://lanej.io/charmed/atoms-input.html#input%2fradio-run)

```sh
go run ./examples/catalog input/radio
```

[Example source](https://github.com/lanej/charmed/blob/e7987b1035f1f0309d8f643ca46001269083830a/docs/snippets/input_radio.go) · [Implementation](https://github.com/lanej/charmed/blob/e7987b1035f1f0309d8f643ca46001269083830a/input/controls.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

### Reusable component

[Radio dependency graph](https://lanej.io/charmed/dependencies.html#component-input%2Fradio) · Partial map · [Source declaration](https://github.com/lanej/charmed/blob/e7987b1035f1f0309d8f643ca46001269083830a/input/controls.go#L157)

Consumer usage uses

- [Radio](https://lanej.io/charmed/dependencies.html#component-input%2Fradio) · primary demonstration

[Example source](https://github.com/lanej/charmed/blob/e7987b1035f1f0309d8f643ca46001269083830a/docs/snippets/input_radio.go#L8)

[Atoms](https://lanej.io/charmed/atoms.html) / Input

## [Toggle](https://lanej.io/charmed/atoms-input.html#input%2ftoggle)

Explicit on/off setting with typed changes, disabled reasons and stable state geometry.

### [Native preview](https://lanej.io/charmed/atoms-input.html#input%2ftoggle-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 ━━● ON  Live updates 
```

Demo · truecolor · 80 × 18 terminal cells

```
input/toggle                                                                    
                                                                                
 ○── OFF Live updates                                                           
 ━━● ON  Enabled setting                                                        
 ○── OFF Locked ⊘ disabled Read-only                                            
                                                                                
Space / Enter activates
```

Narrow · truecolor · 40 × 16 terminal cells

```
input/toggle                            
                                        
 ○── OFF Live updates                   
 ━━● ON  Enabled setting                
 ○── OFF Locked ⊘ disabled Read-only    
                                        
Space / Enter activates
```

256 colors · ansi256 · 80 × 18 terminal cells

```
input/toggle                                                                    
                                                                                
 ○── OFF Live updates                                                           
 ━━● ON  Enabled setting                                                        
 ○── OFF Locked ⊘ disabled Read-only                                            
                                                                                
Space / Enter activates
```

16 colors · ansi16 · 80 × 18 terminal cells

```
input/toggle                                                                    
                                                                                
 ○── OFF Live updates                                                           
 ━━● ON  Enabled setting                                                        
 ○── OFF Locked ⊘ disabled Read-only                                            
                                                                                
Space / Enter activates
```

Monochrome · none · 80 × 18 terminal cells

```
input/toggle                                                                    
                                                                                
 [OFF] Live updates                                                             
 [ON ] Enabled setting                                                          
 [OFF] Locked [DISABLED] Read-only                                              
                                                                                
Space / Enter activates
```

### [Go usage](https://lanej.io/charmed/atoms-input.html#input%2ftoggle-code)

Imports: [`input`](https://pkg.go.dev/github.com/lanej/charmed/input "github.com/lanej/charmed/input") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
toggle := input.NewToggle("live", "Live updates", true, theme.New(theme.Dark))
toggle.Focus()
// input.Toggled carries the setting ID and new On value.
return toggle.View()
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-input.html#input%2ftoggle-integration)

The parent routes keys to the focused control and retains the model and command returned by Update. Use the stable control ID when handling its activation or change message. The control measures its own glyph and label; reserve that space in the layout.

### [Run this example](https://lanej.io/charmed/atoms-input.html#input%2ftoggle-run)

```sh
go run ./examples/catalog input/toggle
```

[Example source](https://github.com/lanej/charmed/blob/e7987b1035f1f0309d8f643ca46001269083830a/docs/snippets/input_toggle.go) · [Implementation](https://github.com/lanej/charmed/blob/e7987b1035f1f0309d8f643ca46001269083830a/input/controls.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Input

## [Action button](https://lanej.io/charmed/atoms-input.html#input%2fbutton)

Application-owned actions with focus, disabled explanations and an explicit busy state.

### [Native preview](https://lanej.io/charmed/atoms-input.html#input%2fbutton-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 ⟨ ↻ Retry ⟩ ⟳ BUSY 
```

Demo · truecolor · 80 × 18 terminal cells

```
input/button                                                                    
                                                                                
 ⟨ ↻ Retry ⟩                                                                    
 ⟨ ↻ Refreshing ⟩ ⟳ BUSY                                                        
 ⟨ ⌫ Delete ⟩ ⊘ disabled Read-only                                              
                                                                                
Space / Enter activates
```

Narrow · truecolor · 40 × 16 terminal cells

```
input/button                            
                                        
 ⟨ ↻ Retry ⟩                            
 ⟨ ↻ Refreshing ⟩ ⟳ BUSY                
 ⟨ ⌫ Delete ⟩ ⊘ disabled Read-only      
                                        
Space / Enter activates
```

256 colors · ansi256 · 80 × 18 terminal cells

```
input/button                                                                    
                                                                                
 ⟨ ↻ Retry ⟩                                                                    
 ⟨ ↻ Refreshing ⟩ ⟳ BUSY                                                        
 ⟨ ⌫ Delete ⟩ ⊘ disabled Read-only                                              
                                                                                
Space / Enter activates
```

16 colors · ansi16 · 80 × 18 terminal cells

```
input/button                                                                    
                                                                                
 ⟨ ↻ Retry ⟩                                                                    
 ⟨ ↻ Refreshing ⟩ ⟳ BUSY                                                        
 ⟨ ⌫ Delete ⟩ ⊘ disabled Read-only                                              
                                                                                
Space / Enter activates
```

Monochrome · none · 80 × 18 terminal cells

```
input/button                                                                    
                                                                                
 [ Retry ]                                                                      
 [ Refreshing ] [BUSY]                                                          
 [ Delete ] [DISABLED] Read-only                                                
                                                                                
Space / Enter activates
```

### [Go usage](https://lanej.io/charmed/atoms-input.html#input%2fbutton-code)

Imports: [`input`](https://pkg.go.dev/github.com/lanej/charmed/input "github.com/lanej/charmed/input") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
button := input.NewButton("retry", "Retry", theme.New(theme.Dark))
button.Icon = theme.GlyphRefresh
button.Focus()
// Handle input.Activated in the application; set Busy while its operation is outstanding.
button.Busy = true

return button.View()
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-input.html#input%2fbutton-integration)

The parent routes keys to the focused control and retains the model and command returned by Update. Use the stable control ID when handling its activation or change message. The control measures its own glyph and label; reserve that space in the layout.

### [Run this example](https://lanej.io/charmed/atoms-input.html#input%2fbutton-run)

```sh
go run ./examples/catalog input/button
```

[Example source](https://github.com/lanej/charmed/blob/e7987b1035f1f0309d8f643ca46001269083830a/docs/snippets/input_button.go) · [Implementation](https://github.com/lanej/charmed/blob/e7987b1035f1f0309d8f643ca46001269083830a/input/controls.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.
