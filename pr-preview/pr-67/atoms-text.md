[Atoms](https://lanej.io/charmed/atoms.html)Text

# [Text](https://lanej.io/charmed/atoms-text.html#level-atoms-text)

17 components · [All atoms](https://lanej.io/charmed/atoms.html)

- [Code](https://lanej.io/charmed/atoms-text.html#text%2fcode)
- [Link](https://lanej.io/charmed/atoms-text.html#text%2flink)
- [Primitives](https://lanej.io/charmed/atoms-text.html#text%2fprimitives)
- [Signed number](https://lanej.io/charmed/atoms-text.html#text%2fsigned-number)
- [Color pill](https://lanej.io/charmed/atoms-text.html#text%2fpill)
- [Presentation](https://lanej.io/charmed/atoms-text.html#text%2fpresentation)
- [Disclosure](https://lanej.io/charmed/atoms-text.html#text%2fdisclosure)
- [Key hint](https://lanej.io/charmed/atoms-text.html#text%2fkey-hint)
- [Syntax](https://lanej.io/charmed/atoms-text.html#text%2fsyntax)
- [Value state](https://lanej.io/charmed/atoms-text.html#text%2fvalue-state)
- [Quantity](https://lanej.io/charmed/atoms-text.html#text%2fquantity)
- [Identifier](https://lanej.io/charmed/atoms-text.html#text%2fidentifier)
- [Match highlighting](https://lanej.io/charmed/atoms-text.html#text%2fmatch)
- [Sort indicator](https://lanej.io/charmed/atoms-text.html#text%2fsort)
- [Change marker](https://lanej.io/charmed/atoms-text.html#text%2fchange)
- [Omission marker](https://lanej.io/charmed/atoms-text.html#text%2fomission)
- [Indicator](https://lanej.io/charmed/atoms-text.html#text%2findicator)

[Atoms](https://lanej.io/charmed/atoms.html) / Text

## [Code](https://lanej.io/charmed/atoms-text.html#text%2fcode)

Render literal inline code and syntax-highlighted source with the same palette as Markdown.

### [Native preview](https://lanej.io/charmed/atoms-text.html#text%2fcode-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 field.Focus() 

                                              
    field := input.NewText(options)           
    
```

Demo · truecolor · 100 × 30 terminal cells

```
Code                                                                                                
                                                                                                    
Call  field.Focus()  to focus the input.                                                            
                                                                                                    
                                                                                                    
    // Shared terminal defaults                                                                     
    options := input.TextOptions{Value: "production", Width: 32}                                    
    field := input.NewText(options)                                                                 
    cmd := field.Focus()                                                                            
    
```

Narrow · truecolor · 40 × 16 terminal cells

```
Code                                    
                                        
Call  field.Focus()  to focus the input.
                                        
                                        
    // Shared terminal defaults         
    options := input.                   
  TextOptions{Value: "production",      
  Width: 32}                            
    field := input.NewText(options)     
    cmd := field.Focus()                
    
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Code                                                                                                
                                                                                                    
Call  field.Focus()  to focus the input.                                                            
                                                                                                    
                                                                                                    
    // Shared terminal defaults                                                                     
    options := input.TextOptions{Value: "production", Width: 32}                                    
    field := input.NewText(options)                                                                 
    cmd := field.Focus()                                                                            
    
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Code                                                                                                
                                                                                                    
Call  field.Focus()  to focus the input.                                                            
                                                                                                    
                                                                                                    
    // Shared terminal defaults                                                                     
    options := input.TextOptions{Value: "production", Width: 32}                                    
    field := input.NewText(options)                                                                 
    cmd := field.Focus()                                                                            
    
```

Monochrome · none · 100 × 30 terminal cells

```
Code                                                                                                
                                                                                                    
Call `field.Focus()` to focus the input.                                                            
                                                                                                    
                                                                                                    
    // Shared terminal defaults                                                                     
    options := input.TextOptions{Value: "production", Width: 32}                                    
    field := input.NewText(options)                                                                 
    cmd := field.Focus()
```

### [Go usage](https://lanej.io/charmed/atoms-text.html#text%2fcode-code)

Imports: [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
inline := text.InlineCode("field.Focus()", palette)

block, err := text.CodeBlock("field := input.NewText(options)", "go", 48, palette)
if err != nil {
	return "", err
}

return inline + "\n\n" + block, nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-text.html#text%2fcode-integration)

Pass literal source, a language and the available width to CodeBlock. Re-render when the palette or width changes, and preserve long lines in a viewport when needed. InlineCode can be composed directly with surrounding terminal text.

### [Run this example](https://lanej.io/charmed/atoms-text.html#text%2fcode-run)

```sh
go run ./examples/catalog text/code
```

[Example source](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/docs/snippets/text_code.go) · [Implementation](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/text/code.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Text

## [Link](https://lanej.io/charmed/atoms-text.html#text%2flink)

Compact underlined labels carry exact absolute URIs through OSC 8; NewWebLink adds HTTP/HTTPS validation. The terminal owns hover and opening; Plain retains destinations without terminal controls. Saved previews show styling only.

### [Native preview](https://lanej.io/charmed/atoms-text.html#text%2flink-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
↗ PR #42
Worker session
PR #42 (https://github.com/lanej/charmed/pull/42)
```

Demo · truecolor · 80 × 18 terminal cells

```
Link / prominence                                                               
                                                                                
normal  ↗ PR #42                                                                
bright  ↗ PR #42                                                                
dim     ↗ PR #42                                                                
                                                                                
Presentation                                                                    
both    ↗ PR #42                                                                
label   PR #42                                                                  
glyph   ↗                                                                       
                                                                                
Explicit plain output                                                           
PR #42 (https://github.com/lanej/charmed/pull/42)
```

Narrow · truecolor · 40 × 16 terminal cells

```
Link / prominence                       
                                        
normal  ↗ PR #42                        
bright  ↗ PR #42                        
dim     ↗ PR #42                        
                                        
Presentation                            
both    ↗ PR #42                        
label   PR #42                          
glyph   ↗                               
                                        
Explicit plain output                   
PR #42                                  
(https://github.com/lanej/charmed/pull/4
2)
```

256 colors · ansi256 · 80 × 18 terminal cells

```
Link / prominence                                                               
                                                                                
normal  ↗ PR #42                                                                
bright  ↗ PR #42                                                                
dim     ↗ PR #42                                                                
                                                                                
Presentation                                                                    
both    ↗ PR #42                                                                
label   PR #42                                                                  
glyph   ↗                                                                       
                                                                                
Explicit plain output                                                           
PR #42 (https://github.com/lanej/charmed/pull/42)
```

16 colors · ansi16 · 80 × 18 terminal cells

```
Link / prominence                                                               
                                                                                
normal  ↗ PR #42                                                                
bright  ↗ PR #42                                                                
dim     ↗ PR #42                                                                
                                                                                
Presentation                                                                    
both    ↗ PR #42                                                                
label   PR #42                                                                  
glyph   ↗                                                                       
                                                                                
Explicit plain output                                                           
PR #42 (https://github.com/lanej/charmed/pull/42)
```

Monochrome · none · 80 × 18 terminal cells

```
Link / prominence                                                               
                                                                                
normal  @ PR #42                                                                
bright  @ PR #42                                                                
dim     @ PR #42                                                                
                                                                                
Presentation                                                                    
both    @ PR #42                                                                
label   PR #42                                                                  
glyph   @                                                                       
                                                                                
Explicit plain output                                                           
PR #42 (https://github.com/lanej/charmed/pull/42)
```

light · truecolor · 80 × 18 terminal cells

```
Link / prominence                                                               
                                                                                
normal  ↗ PR #42                                                                
bright  ↗ PR #42                                                                
dim     ↗ PR #42                                                                
                                                                                
Presentation                                                                    
both    ↗ PR #42                                                                
label   PR #42                                                                  
glyph   ↗                                                                       
                                                                                
Explicit plain output                                                           
PR #42 (https://github.com/lanej/charmed/pull/42)
```

contrast · truecolor · 80 × 18 terminal cells

```
Link / prominence                                                               
                                                                                
normal  ↗ PR #42                                                                
bright  ↗ PR #42                                                                
dim     ↗ PR #42                                                                
                                                                                
Presentation                                                                    
both    ↗ PR #42                                                                
label   PR #42                                                                  
glyph   ↗                                                                       
                                                                                
Explicit plain output                                                           
PR #42 (https://github.com/lanej/charmed/pull/42)
```

### [Go usage](https://lanej.io/charmed/atoms-text.html#text%2flink-code)

Imports: [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
link, err := text.NewWebLink("PR #42", "https://github.com/lanej/charmed/pull/42")
if err != nil {
	return "", err
}

compact, err := link.View(theme.New(theme.Dark))
if err != nil {
	return "", err
}

plain, err := link.Plain()
if err != nil {
	return "", err
}

session := text.Link{Label: "Worker session", URL: "terminal://session/build/pane/tests", Presentation: text.LabelOnly, Emphasis: text.Dim}

sessionView, err := session.View(theme.New(theme.Dark))
if err != nil {
	return "", err
}

return compact + "\n" + sessionView + "\n" + plain, nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-text.html#text%2flink-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-text.html#text%2flink-run)

```sh
go run ./examples/catalog text/link
```

[Example source](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/docs/snippets/text_link.go) · [Implementation](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/text/link.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Text

## [Primitives](https://lanej.io/charmed/atoms-text.html#text%2fprimitives)

Labels, badges and keyboard hints share semantic roles.

### [Native preview](https://lanej.io/charmed/atoms-text.html#text%2fprimitives-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
production          
[READY] [3 SELECTED]
 Success            
Enter select
```

Demo · truecolor · 100 × 30 terminal cells

```
Text primitives                                                                                     
                                                                                                    
production                                                                                          
                                                                                                    
[READY] [3 SELECTED]                                                                                
                                                                                                    
 Success   Info                                                                                     
 Warning   Danger                                                                                   
                                                                                                    
Enter select   Esc dismiss
```

Narrow · truecolor · 40 × 16 terminal cells

```
Text primitives                         
                                        
production                              
                                        
[READY] [3 SELECTED]                    
                                        
 Success   Info                         
 Warning   Danger                       
                                        
Enter select   Esc dismiss
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Text primitives                                                                                     
                                                                                                    
production                                                                                          
                                                                                                    
[READY] [3 SELECTED]                                                                                
                                                                                                    
 Success   Info                                                                                     
 Warning   Danger                                                                                   
                                                                                                    
Enter select   Esc dismiss
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Text primitives                                                                                     
                                                                                                    
production                                                                                          
                                                                                                    
[READY] [3 SELECTED]                                                                                
                                                                                                    
 Success   Info                                                                                     
 Warning   Danger                                                                                   
                                                                                                    
Enter select   Esc dismiss
```

Monochrome · none · 100 × 30 terminal cells

```
Text primitives                                                                                     
                                                                                                    
production                                                                                          
                                                                                                    
[READY] [3 SELECTED]                                                                                
                                                                                                    
[Success] [Info]                                                                                    
[Warning] [Danger]                                                                                  
                                                                                                    
Enter select   Esc dismiss
```

### [Go usage](https://lanej.io/charmed/atoms-text.html#text%2fprimitives-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return layout.Column(
	text.Label("production", palette),
	text.Badge("READY", palette)+" "+text.Badge("3 SELECTED", palette),
	text.Pill("Success", text.Success, palette),
	text.KeyHint("Enter", "select", palette),
)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-text.html#text%2fprimitives-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-text.html#text%2fprimitives-run)

```sh
go run ./examples/catalog text/primitives
```

[Example source](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/docs/snippets/text_primitives.go) · [Implementation](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/story/atoms.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Text

## [Signed number](https://lanej.io/charmed/atoms-text.html#text%2fsigned-number)

Exact signed values, configurable positive/negative tones, neutral zero, tiny-value preservation and explicit unavailable state.

### [Native preview](https://lanej.io/charmed/atoms-text.html#text%2fsigned-number-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
+0.180pp
```

Demo · truecolor · 80 × 12 terminal cells

```
Signed number / normal                                                          
                                                                                
Positive means improvement   +0.180pp                                           
Higher server-error rate     +0.180pp                                           
Zero                        +0.000pp                                            
Small nonzero               -1e-06pp                                            
Unavailable                 n/a
```

Narrow · truecolor · 40 × 16 terminal cells

```
Signed number / normal                  
                                        
Positive means improvement   +0.180pp   
Higher server-error rate     +0.180pp   
Zero                        +0.000pp    
Small nonzero               -1e-06pp    
Unavailable                 n/a
```

256 colors · ansi256 · 80 × 12 terminal cells

```
Signed number / normal                                                          
                                                                                
Positive means improvement   +0.180pp                                           
Higher server-error rate     +0.180pp                                           
Zero                        +0.000pp                                            
Small nonzero               -1e-06pp                                            
Unavailable                 n/a
```

16 colors · ansi16 · 80 × 12 terminal cells

```
Signed number / normal                                                          
                                                                                
Positive means improvement   +0.180pp                                           
Higher server-error rate     +0.180pp                                           
Zero                        +0.000pp                                            
Small nonzero               -1e-06pp                                            
Unavailable                 n/a
```

Monochrome · none · 80 × 12 terminal cells

```
Signed number / normal                                                          
                                                                                
Positive means improvement   +0.180pp                                           
Higher server-error rate     +0.180pp                                           
Zero                        +0.000pp                                            
Small nonzero               -1e-06pp                                            
Unavailable                 n/a
```

Bright · truecolor · 80 × 12 terminal cells

```
Signed number / bright                                                          
                                                                                
Positive means improvement   +0.180pp                                           
Higher server-error rate     +0.180pp                                           
Zero                        +0.000pp                                            
Small nonzero               -1e-06pp                                            
Unavailable                 n/a
```

Normal · truecolor · 80 × 12 terminal cells

```
Signed number / normal                                                          
                                                                                
Positive means improvement   +0.180pp                                           
Higher server-error rate     +0.180pp                                           
Zero                        +0.000pp                                            
Small nonzero               -1e-06pp                                            
Unavailable                 n/a
```

Dim · truecolor · 80 × 12 terminal cells

```
Signed number / dim                                                             
                                                                                
Positive means improvement   +0.180pp                                           
Higher server-error rate     +0.180pp                                           
Zero                        +0.000pp                                            
Small nonzero               -1e-06pp                                            
Unavailable                 n/a
```

### [Go usage](https://lanej.io/charmed/atoms-text.html#text%2fsigned-number-code)

Imports: [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
change := text.SignedNumber{
	Value: .18, Available: true, Precision: 3, Unit: "pp", Emphasis: text.Bright,
	Tones: text.SignTones{Positive: text.Danger, Negative: text.Success},
}

return change.View(theme.New(theme.Dark))
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-text.html#text%2fsigned-number-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-text.html#text%2fsigned-number-run)

```sh
go run ./examples/catalog text/signed-number
```

[Example source](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/docs/snippets/text_signed_number.go) · [Implementation](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/text/signed.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Text

## [Color pill](https://lanej.io/charmed/atoms-text.html#text%2fpill)

Semantic tone and bright, normal or dim emphasis are independent. ASCII output retains the label.

### [Native preview](https://lanej.io/charmed/atoms-text.html#text%2fpill-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 Ready       
 Warning     
 Unavailable 
```

Demo · truecolor · 80 × 12 terminal cells

```
Pills / tone and emphasis                                                       
                                                                                
bright   neutral    success    info    warning    danger                        
normal   neutral    success    info    warning    danger                        
dim      neutral    success    info    warning    danger                        
                                                                                
Selected emphasis   production 
```

Narrow · truecolor · 40 × 16 terminal cells

```
Pills / tone and emphasis               
                                        
bright                                  
 neutral    success    info             
 warning    danger                      
normal                                  
 neutral    success    info             
 warning    danger                      
dim                                     
 neutral    success    info             
 warning    danger                      
                                        
Selected emphasis   production 
```

256 colors · ansi256 · 80 × 12 terminal cells

```
Pills / tone and emphasis                                                       
                                                                                
bright   neutral    success    info    warning    danger                        
normal   neutral    success    info    warning    danger                        
dim      neutral    success    info    warning    danger                        
                                                                                
Selected emphasis   production 
```

16 colors · ansi16 · 80 × 12 terminal cells

```
Pills / tone and emphasis                                                       
                                                                                
bright   neutral    success    info    warning    danger                        
normal   neutral    success    info    warning    danger                        
dim      neutral    success    info    warning    danger                        
                                                                                
Selected emphasis   production 
```

Monochrome · none · 80 × 12 terminal cells

```
Pills / tone and emphasis                                                       
                                                                                
bright  [neutral]  [success]  [info]  [warning]  [danger]                       
normal  [neutral]  [success]  [info]  [warning]  [danger]                       
dim     [neutral]  [success]  [info]  [warning]  [danger]                       
                                                                                
Selected emphasis  [production]
```

Bright · truecolor · 80 × 12 terminal cells

```
Pills / tone and emphasis                                                       
                                                                                
bright   neutral    success    info    warning    danger                        
normal   neutral    success    info    warning    danger                        
dim      neutral    success    info    warning    danger                        
                                                                                
Selected emphasis   production 
```

Normal · truecolor · 80 × 12 terminal cells

```
Pills / tone and emphasis                                                       
                                                                                
bright   neutral    success    info    warning    danger                        
normal   neutral    success    info    warning    danger                        
dim      neutral    success    info    warning    danger                        
                                                                                
Selected emphasis   production 
```

Dim · truecolor · 80 × 12 terminal cells

```
Pills / tone and emphasis                                                       
                                                                                
bright   neutral    success    info    warning    danger                        
normal   neutral    success    info    warning    danger                        
dim      neutral    success    info    warning    danger                        
                                                                                
Selected emphasis   production 
```

### [Go usage](https://lanej.io/charmed/atoms-text.html#text%2fpill-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return layout.Column(
	text.Pill("Ready", text.Success, palette, text.Bright),
	text.Pill("Warning", text.Warning, palette, text.Normal),
	text.Pill("Unavailable", text.Neutral, palette, text.Dim),
)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-text.html#text%2fpill-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-text.html#text%2fpill-run)

```sh
go run ./examples/catalog text/pill
```

[Example source](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/docs/snippets/text_pill.go) · [Implementation](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/text/text.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

### Reusable component

[Color pill dependency graph](https://lanej.io/charmed/dependencies.html#component-text%2Fpill) · Partial map · [Source declaration](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/text/text.go#L29)

Consumer usage uses

- [Color pill](https://lanej.io/charmed/dependencies.html#component-text%2Fpill) · primary demonstration

[Example source](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/docs/snippets/text_pill.go#L9)

[Atoms](https://lanej.io/charmed/atoms.html) / Text

## [Presentation](https://lanej.io/charmed/atoms-text.html#text%2fpresentation)

Glyph and label, glyph only, or label only; shared policy for semantic atoms.

### [Native preview](https://lanej.io/charmed/atoms-text.html#text%2fpresentation-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 ● CONNECTED 
 ●           
 CONNECTED   
●
```

Demo · truecolor · 100 × 30 terminal cells

```
Semantic atom presentation                                                                          
                                                                                                    
both                    glyph                   label                                               
                                                                                                    
 ‖ PAUSED                ‖                       PAUSED                                             
⋈ Trace                 ⋈                       Trace                                               
▾ Response              ▾                       Response
```

Narrow · truecolor · 40 × 16 terminal cells

```
both                                    
                                        
 ‖ PAUSED                               
⋈ Trace                                 
▾ Response                              
                                        
glyph                                   
                                        
 ‖                                      
⋈                                       
▾                                       
                                        
label                                   
                                        
 PAUSED                                 
Trace
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Semantic atom presentation                                                                          
                                                                                                    
both                    glyph                   label                                               
                                                                                                    
 ‖ PAUSED                ‖                       PAUSED                                             
⋈ Trace                 ⋈                       Trace                                               
▾ Response              ▾                       Response
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Semantic atom presentation                                                                          
                                                                                                    
both                    glyph                   label                                               
                                                                                                    
 ‖ PAUSED                ‖                       PAUSED                                             
⋈ Trace                 ⋈                       Trace                                               
▾ Response              ▾                       Response
```

Monochrome · none · 100 × 30 terminal cells

```
Semantic atom presentation                                                                          
                                                                                                    
both                    glyph                   label                                               
                                                                                                    
[= PAUSED]              [=]                     [PAUSED]                                            
T Trace                 T                       Trace                                               
v Response              v                       Response
```

### [Go usage](https://lanej.io/charmed/atoms-text.html#text%2fpresentation-code)

Imports: [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return layout.Column(
	feedback.Connected.View(text.Normal, palette, text.GlyphAndLabel),
	feedback.Connected.View(text.Normal, palette, text.GlyphOnly),
	feedback.Connected.View(text.Normal, palette, text.LabelOnly),
	feedback.Connected.View(text.Normal, palette, text.Compact),
)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-text.html#text%2fpresentation-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-text.html#text%2fpresentation-run)

```sh
go run ./examples/catalog text/presentation
```

[Example source](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/docs/snippets/text_presentation.go) · [Implementation](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/text/presentation.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Text

## [Disclosure](https://lanej.io/charmed/atoms-text.html#text%2fdisclosure)

Expanded and collapsed fold labels with shared glyphs and escaped text.

### [Native preview](https://lanej.io/charmed/atoms-text.html#text%2fdisclosure-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
▾ Request headers
▸ Response body
```

Demo · truecolor · 100 × 30 terminal cells

```
Optional disclosure controls                                                                        
                                                                                                    
▸ Response                                                                                          
▾ Response                                                                                          
                                                                                                    
Folding disabled: all content stays open                                                            
                                                                                                    
Response
```

Narrow · truecolor · 40 × 16 terminal cells

```
Optional disclosure controls            
                                        
▸ Response                              
▾ Response                              
                                        
Folding disabled: all content stays open
                                        
Response
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Optional disclosure controls                                                                        
                                                                                                    
▸ Response                                                                                          
▾ Response                                                                                          
                                                                                                    
Folding disabled: all content stays open                                                            
                                                                                                    
Response
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Optional disclosure controls                                                                        
                                                                                                    
▸ Response                                                                                          
▾ Response                                                                                          
                                                                                                    
Folding disabled: all content stays open                                                            
                                                                                                    
Response
```

Monochrome · none · 100 × 30 terminal cells

```
Optional disclosure controls                                                                        
                                                                                                    
> Response                                                                                          
v Response                                                                                          
                                                                                                    
Folding disabled: all content stays open                                                            
                                                                                                    
Response
```

### [Go usage](https://lanej.io/charmed/atoms-text.html#text%2fdisclosure-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return layout.Column(
	text.Disclosure("Request headers", true, palette),
	text.Disclosure("Response body", false, palette),
)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-text.html#text%2fdisclosure-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-text.html#text%2fdisclosure-run)

```sh
go run ./examples/catalog text/disclosure
```

[Example source](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/docs/snippets/text_disclosure.go) · [Implementation](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/text/disclosure.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Text

## [Key hint](https://lanej.io/charmed/atoms-text.html#text%2fkey-hint)

Shared keyboard-action atom used by inspector controls.

### [Native preview](https://lanej.io/charmed/atoms-text.html#text%2fkey-hint-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Enter select  Esc dismiss
```

Demo · truecolor · 100 × 30 terminal cells

```
Keyboard action atoms                                                                               
                                                                                                    
Tab next pane                                                                                       
F resume follow                                                                                     
Enter toggle fold                                                                                   
r raw / pretty
```

Narrow · truecolor · 40 × 16 terminal cells

```
Keyboard action atoms                   
                                        
Tab next pane                           
F resume follow                         
Enter toggle fold                       
r raw / pretty
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Keyboard action atoms                                                                               
                                                                                                    
Tab next pane                                                                                       
F resume follow                                                                                     
Enter toggle fold                                                                                   
r raw / pretty
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Keyboard action atoms                                                                               
                                                                                                    
Tab next pane                                                                                       
F resume follow                                                                                     
Enter toggle fold                                                                                   
r raw / pretty
```

Monochrome · none · 100 × 30 terminal cells

```
Keyboard action atoms                                                                               
                                                                                                    
Tab next pane                                                                                       
F resume follow                                                                                     
Enter toggle fold                                                                                   
r raw / pretty
```

### [Go usage](https://lanej.io/charmed/atoms-text.html#text%2fkey-hint-code)

Imports: [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
return text.KeyHint("Enter", "select", palette) + "  " + text.KeyHint("Esc", "dismiss", palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-text.html#text%2fkey-hint-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-text.html#text%2fkey-hint-run)

```sh
go run ./examples/catalog text/key-hint
```

[Example source](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/docs/snippets/text_key_hint.go) · [Implementation](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/text/text.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Text

## [Syntax](https://lanej.io/charmed/atoms-text.html#text%2fsyntax)

Content-type aware lexical highlighting; terminal controls stay visible as escaped data.

### [Native preview](https://lanej.io/charmed/atoms-text.html#text%2fsyntax-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
{
\t  "error": "upstream_timeout",
\t  "retryable": true,
\t  "attempt": 2
\t}
```

Demo · truecolor · 100 × 30 terminal cells

```
application/problem+json; charset=utf-8                                                             
                                                                                                    
{                                                                                                   
  "error": "upstream_timeout",                                                                      
  "retryable": true,                                                                                
  "attempt": 2                                                                                      
}
```

Narrow · truecolor · 40 × 16 terminal cells

```
application/problem+json; charset=utf-8 
                                        
{                                       
  "error": "upstream_timeout",          
  "retryable": true,                    
  "attempt": 2                          
}
```

256 colors · ansi256 · 100 × 30 terminal cells

```
application/problem+json; charset=utf-8                                                             
                                                                                                    
{                                                                                                   
  "error": "upstream_timeout",                                                                      
  "retryable": true,                                                                                
  "attempt": 2                                                                                      
}
```

16 colors · ansi16 · 100 × 30 terminal cells

```
application/problem+json; charset=utf-8                                                             
                                                                                                    
{                                                                                                   
  "error": "upstream_timeout",                                                                      
  "retryable": true,                                                                                
  "attempt": 2                                                                                      
}
```

Monochrome · none · 100 × 30 terminal cells

```
application/problem+json; charset=utf-8                                                             
                                                                                                    
{                                                                                                   
  "error": "upstream_timeout",                                                                      
  "retryable": true,                                                                                
  "attempt": 2                                                                                      
}
```

light · truecolor · 100 × 30 terminal cells

```
application/problem+json; charset=utf-8                                                             
                                                                                                    
{                                                                                                   
  "error": "upstream_timeout",                                                                      
  "retryable": true,                                                                                
  "attempt": 2                                                                                      
}
```

contrast · truecolor · 100 × 30 terminal cells

```
application/problem+json; charset=utf-8                                                             
                                                                                                    
{                                                                                                   
  "error": "upstream_timeout",                                                                      
  "retryable": true,                                                                                
  "attempt": 2                                                                                      
}
```

### [Go usage](https://lanej.io/charmed/atoms-text.html#text%2fsyntax-code)

Imports: [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
body := `{
	  "error": "upstream_timeout",
	  "retryable": true,
	  "attempt": 2
	}`

return text.Syntax(body, "json", theme.New(theme.Dark))
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-text.html#text%2fsyntax-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-text.html#text%2fsyntax-run)

```sh
go run ./examples/catalog text/syntax
```

[Example source](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/docs/snippets/text_syntax.go) · [Implementation](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/text/syntax.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Text

## [Value state](https://lanej.io/charmed/atoms-text.html#text%2fvalue-state)

Present, missing, redacted and invalid values remain distinct from empty strings, zero and false.

### [Native preview](https://lanej.io/charmed/atoms-text.html#text%2fvalue-state-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
""  ∅ missing  ▧ redacted
```

Demo · truecolor · 80 × 18 terminal cells

```
text/value-state / normal                                                       
                                                                                
""                                                                              
0                                                                               
false                                                                           
api                                                                             
∅ missing                                                                       
▧ redacted                                                                      
! invalid
```

Narrow · truecolor · 40 × 16 terminal cells

```
text/value-state / normal               
                                        
""                                      
0                                       
false                                   
api                                     
∅ missing                               
▧ redacted                              
! invalid
```

256 colors · ansi256 · 80 × 18 terminal cells

```
text/value-state / normal                                                       
                                                                                
""                                                                              
0                                                                               
false                                                                           
api                                                                             
∅ missing                                                                       
▧ redacted                                                                      
! invalid
```

16 colors · ansi16 · 80 × 18 terminal cells

```
text/value-state / normal                                                       
                                                                                
""                                                                              
0                                                                               
false                                                                           
api                                                                             
∅ missing                                                                       
▧ redacted                                                                      
! invalid
```

Monochrome · none · 80 × 18 terminal cells

```
text/value-state / normal                                                       
                                                                                
""                                                                              
0                                                                               
false                                                                           
api                                                                             
missing                                                                         
redacted                                                                        
invalid
```

Bright · truecolor · 80 × 18 terminal cells

```
text/value-state / bright                                                       
                                                                                
""                                                                              
0                                                                               
false                                                                           
api                                                                             
∅ missing                                                                       
▧ redacted                                                                      
! invalid
```

Normal · truecolor · 80 × 18 terminal cells

```
text/value-state / normal                                                       
                                                                                
""                                                                              
0                                                                               
false                                                                           
api                                                                             
∅ missing                                                                       
▧ redacted                                                                      
! invalid
```

Dim · truecolor · 80 × 18 terminal cells

```
text/value-state / dim                                                          
                                                                                
""                                                                              
0                                                                               
false                                                                           
api                                                                             
∅ missing                                                                       
▧ redacted                                                                      
! invalid
```

### [Go usage](https://lanej.io/charmed/atoms-text.html#text%2fvalue-state-code)

Imports: [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
// An empty string is present; redaction never renders the supplied secret.
return text.ValuePresent.View("", text.Normal, palette) + "  " +
	text.ValueMissing.View("", text.Normal, palette) + "  " +
	text.ValueRedacted.View("secret", text.Normal, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-text.html#text%2fvalue-state-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-text.html#text%2fvalue-state-run)

```sh
go run ./examples/catalog text/value-state
```

[Example source](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/docs/snippets/text_value_state.go) · [Implementation](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/text/value.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Text

## [Quantity](https://lanej.io/charmed/atoms-text.html#text%2fquantity)

Decimal or binary scaling, literal units, right alignment and small nonzero preservation.

### [Native preview](https://lanej.io/charmed/atoms-text.html#text%2fquantity-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
         1.2 MiB
```

Demo · truecolor · 80 × 18 terminal cells

```
text/quantity / normal                                                          
                                                                                
         1.2 MiB                                                                
       842 req/s                                                                
           12.5k                                                                
         1e-06 s                                                                
      0 requests                                                                
         missing
```

Narrow · truecolor · 40 × 16 terminal cells

```
text/quantity / normal                  
                                        
         1.2 MiB                        
       842 req/s                        
           12.5k                        
         1e-06 s                        
      0 requests                        
         missing
```

256 colors · ansi256 · 80 × 18 terminal cells

```
text/quantity / normal                                                          
                                                                                
         1.2 MiB                                                                
       842 req/s                                                                
           12.5k                                                                
         1e-06 s                                                                
      0 requests                                                                
         missing
```

16 colors · ansi16 · 80 × 18 terminal cells

```
text/quantity / normal                                                          
                                                                                
         1.2 MiB                                                                
       842 req/s                                                                
           12.5k                                                                
         1e-06 s                                                                
      0 requests                                                                
         missing
```

Monochrome · none · 80 × 18 terminal cells

```
text/quantity / normal                                                          
                                                                                
         1.2 MiB                                                                
       842 req/s                                                                
           12.5k                                                                
         1e-06 s                                                                
      0 requests                                                                
         missing
```

Bright · truecolor · 80 × 18 terminal cells

```
text/quantity / bright                                                          
                                                                                
         1.2 MiB                                                                
       842 req/s                                                                
           12.5k                                                                
         1e-06 s                                                                
      0 requests                                                                
         missing
```

Normal · truecolor · 80 × 18 terminal cells

```
text/quantity / normal                                                          
                                                                                
         1.2 MiB                                                                
       842 req/s                                                                
           12.5k                                                                
         1e-06 s                                                                
      0 requests                                                                
         missing
```

Dim · truecolor · 80 × 18 terminal cells

```
text/quantity / dim                                                             
                                                                                
         1.2 MiB                                                                
       842 req/s                                                                
           12.5k                                                                
         1e-06 s                                                                
      0 requests                                                                
         missing
```

### [Go usage](https://lanej.io/charmed/atoms-text.html#text%2fquantity-code)

Imports: [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
quantity := text.Quantity{
	Value: 1258291, Unit: "B", Scale: text.BinaryScale,
	Precision: 1, Width: 16, Emphasis: text.Normal,
}

return quantity.View(theme.New(theme.Dark))
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-text.html#text%2fquantity-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-text.html#text%2fquantity-run)

```sh
go run ./examples/catalog text/quantity
```

[Example source](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/docs/snippets/text_quantity.go) · [Implementation](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/text/quantity.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Text

## [Identifier](https://lanej.io/charmed/atoms-text.html#text%2fidentifier)

Cell-bounded middle abbreviation keeps the full underlying identifier intact.

### [Native preview](https://lanej.io/charmed/atoms-text.html#text%2fidentifier-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
trace-a...6-92ef
```

Demo · truecolor · 80 × 18 terminal cells

```
text/identifier / normal                                                        
                                                                                
trace-a17c9dd285e14766-92ef                                                     
trace-a...6-92ef                                                                
tra...ef
```

Narrow · truecolor · 40 × 16 terminal cells

```
text/identifier / normal                
                                        
trace-a17c9dd285e14766-92ef             
trace-a...6-92ef                        
tra...ef
```

256 colors · ansi256 · 80 × 18 terminal cells

```
text/identifier / normal                                                        
                                                                                
trace-a17c9dd285e14766-92ef                                                     
trace-a...6-92ef                                                                
tra...ef
```

16 colors · ansi16 · 80 × 18 terminal cells

```
text/identifier / normal                                                        
                                                                                
trace-a17c9dd285e14766-92ef                                                     
trace-a...6-92ef                                                                
tra...ef
```

Monochrome · none · 80 × 18 terminal cells

```
text/identifier / normal                                                        
                                                                                
trace-a17c9dd285e14766-92ef                                                     
trace-a...6-92ef                                                                
tra...ef
```

Bright · truecolor · 80 × 18 terminal cells

```
text/identifier / bright                                                        
                                                                                
trace-a17c9dd285e14766-92ef                                                     
trace-a...6-92ef                                                                
tra...ef
```

Normal · truecolor · 80 × 18 terminal cells

```
text/identifier / normal                                                        
                                                                                
trace-a17c9dd285e14766-92ef                                                     
trace-a...6-92ef                                                                
tra...ef
```

Dim · truecolor · 80 × 18 terminal cells

```
text/identifier / dim                                                           
                                                                                
trace-a17c9dd285e14766-92ef                                                     
trace-a...6-92ef                                                                
tra...ef
```

### [Go usage](https://lanej.io/charmed/atoms-text.html#text%2fidentifier-code)

Imports: [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
identifier := text.Identifier{Value: "trace-a17c9dd285e14766-92ef", Emphasis: text.Normal}
// Use identifier.Value for keys and actions; Label is only a display abbreviation.
return identifier.View(16, theme.New(theme.Dark))
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-text.html#text%2fidentifier-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-text.html#text%2fidentifier-run)

```sh
go run ./examples/catalog text/identifier
```

[Example source](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/docs/snippets/text_identifier.go) · [Implementation](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/text/identifier.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Text

## [Match highlighting](https://lanej.io/charmed/atoms-text.html#text%2fmatch)

Caller-supplied rune ranges highlight entire graphemes while source controls remain escaped data.

### [Native preview](https://lanej.io/charmed/atoms-text.html#text%2fmatch-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Layout: timeline
```

Demo · truecolor · 80 × 18 terminal cells

```
text/match / normal                                                             
                                                                                
Layout: timeline                                                                
界é sample
```

Narrow · truecolor · 40 × 16 terminal cells

```
text/match / normal                     
                                        
Layout: timeline                        
界é sample
```

256 colors · ansi256 · 80 × 18 terminal cells

```
text/match / normal                                                             
                                                                                
Layout: timeline                                                                
界é sample
```

16 colors · ansi16 · 80 × 18 terminal cells

```
text/match / normal                                                             
                                                                                
Layout: timeline                                                                
界é sample
```

Monochrome · none · 80 × 18 terminal cells

```
text/match / normal                                                             
                                                                                
Layout: timeline                                                                
?e? sample
```

Bright · truecolor · 80 × 18 terminal cells

```
text/match / bright                                                             
                                                                                
Layout: timeline                                                                
界é sample
```

Normal · truecolor · 80 × 18 terminal cells

```
text/match / normal                                                             
                                                                                
Layout: timeline                                                                
界é sample
```

Dim · truecolor · 80 × 18 terminal cells

```
text/match / dim                                                                
                                                                                
Layout: timeline                                                                
界é sample
```

### [Go usage](https://lanej.io/charmed/atoms-text.html#text%2fmatch-code)

Imports: [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
normal := theme.New(theme.Dark).Text()
// Ranges are half-open rune offsets into this original source string.
return text.HighlightMatches("Layout: timeline",
	[]text.MatchRange{{Start: 0, End: 3}, {Start: 8, End: 9}},
	normal, normal.Underline(true))
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-text.html#text%2fmatch-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-text.html#text%2fmatch-run)

```sh
go run ./examples/catalog text/match
```

[Example source](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/docs/snippets/text_match.go) · [Implementation](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/text/match.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Text

## [Sort indicator](https://lanej.io/charmed/atoms-text.html#text%2fsort)

Explicit unsorted, ascending and descending order with optional sort-key precedence.

### [Native preview](https://lanej.io/charmed/atoms-text.html#text%2fsort-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Latency ↓ DESC 2
```

Demo · truecolor · 80 × 18 terminal cells

```
text/sort / normal                                                              
                                                                                
UNSORTED                                                                        
↑ ASC 1                                                                         
↓ DESC 1
```

Narrow · truecolor · 40 × 16 terminal cells

```
text/sort / normal                      
                                        
UNSORTED                                
↑ ASC 1                                 
↓ DESC 1
```

256 colors · ansi256 · 80 × 18 terminal cells

```
text/sort / normal                                                              
                                                                                
UNSORTED                                                                        
↑ ASC 1                                                                         
↓ DESC 1
```

16 colors · ansi16 · 80 × 18 terminal cells

```
text/sort / normal                                                              
                                                                                
UNSORTED                                                                        
↑ ASC 1                                                                         
↓ DESC 1
```

Monochrome · none · 80 × 18 terminal cells

```
text/sort / normal                                                              
                                                                                
UNSORTED                                                                        
^ ASC 1                                                                         
v DESC 1
```

Bright · truecolor · 80 × 18 terminal cells

```
text/sort / bright                                                              
                                                                                
UNSORTED                                                                        
↑ ASC 1                                                                         
↓ DESC 1
```

Normal · truecolor · 80 × 18 terminal cells

```
text/sort / normal                                                              
                                                                                
UNSORTED                                                                        
↑ ASC 1                                                                         
↓ DESC 1
```

Dim · truecolor · 80 × 18 terminal cells

```
text/sort / dim                                                                 
                                                                                
UNSORTED                                                                        
↑ ASC 1                                                                         
↓ DESC 1
```

### [Go usage](https://lanej.io/charmed/atoms-text.html#text%2fsort-code)

Imports: [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
indicator := text.SortIndicator{
	Direction: text.SortDescending, Priority: 2, Emphasis: text.Normal,
}

return "Latency " + indicator.View(theme.New(theme.Dark))
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-text.html#text%2fsort-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-text.html#text%2fsort-run)

```sh
go run ./examples/catalog text/sort
```

[Example source](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/docs/snippets/text_sort.go) · [Implementation](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/text/markers.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Text

## [Change marker](https://lanej.io/charmed/atoms-text.html#text%2fchange)

Added, modified, removed and unchanged categories preserve meaning without color.

### [Native preview](https://lanej.io/charmed/atoms-text.html#text%2fchange-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
⊕ added  ↺ modified  ⊖ removed
```

Demo · truecolor · 80 × 18 terminal cells

```
text/change / normal                                                            
                                                                                
≡ unchanged                                                                     
⊕ added                                                                         
↺ modified                                                                      
⊖ removed
```

Narrow · truecolor · 40 × 16 terminal cells

```
text/change / normal                    
                                        
≡ unchanged                             
⊕ added                                 
↺ modified                              
⊖ removed
```

256 colors · ansi256 · 80 × 18 terminal cells

```
text/change / normal                                                            
                                                                                
≡ unchanged                                                                     
⊕ added                                                                         
↺ modified                                                                      
⊖ removed
```

16 colors · ansi16 · 80 × 18 terminal cells

```
text/change / normal                                                            
                                                                                
≡ unchanged                                                                     
⊕ added                                                                         
↺ modified                                                                      
⊖ removed
```

Monochrome · none · 80 × 18 terminal cells

```
text/change / normal                                                            
                                                                                
= unchanged                                                                     
+ added                                                                         
~ modified                                                                      
- removed
```

Bright · truecolor · 80 × 18 terminal cells

```
text/change / bright                                                            
                                                                                
≡ unchanged                                                                     
⊕ added                                                                         
↺ modified                                                                      
⊖ removed
```

Normal · truecolor · 80 × 18 terminal cells

```
text/change / normal                                                            
                                                                                
≡ unchanged                                                                     
⊕ added                                                                         
↺ modified                                                                      
⊖ removed
```

Dim · truecolor · 80 × 18 terminal cells

```
text/change / dim                                                               
                                                                                
≡ unchanged                                                                     
⊕ added                                                                         
↺ modified                                                                      
⊖ removed
```

### [Go usage](https://lanej.io/charmed/atoms-text.html#text%2fchange-code)

Imports: [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return text.Added.View(text.Normal, palette) + "  " +
	text.Modified.View(text.Normal, palette) + "  " +
	text.Removed.View(text.Normal, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-text.html#text%2fchange-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-text.html#text%2fchange-run)

```sh
go run ./examples/catalog text/change
```

[Example source](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/docs/snippets/text_change.go) · [Implementation](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/text/markers.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Text

## [Omission marker](https://lanej.io/charmed/atoms-text.html#text%2fomission)

Display-hidden and source-truncated evidence retain known, unknown and zero omitted counts.

### [Native preview](https://lanej.io/charmed/atoms-text.html#text%2fomission-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
… +24 more  ! source truncated (count unknown)
```

Demo · truecolor · 80 × 18 terminal cells

```
text/omission / normal                                                          
                                                                                
… +24 more                                                                      
… more hidden                                                                   
! source truncated (count unknown)                                              
! source truncated (42 omitted)
```

Narrow · truecolor · 40 × 16 terminal cells

```
text/omission / normal                  
                                        
… +24 more                              
… more hidden                           
! source truncated (count unknown)      
! source truncated (42 omitted)
```

256 colors · ansi256 · 80 × 18 terminal cells

```
text/omission / normal                                                          
                                                                                
… +24 more                                                                      
… more hidden                                                                   
! source truncated (count unknown)                                              
! source truncated (42 omitted)
```

16 colors · ansi16 · 80 × 18 terminal cells

```
text/omission / normal                                                          
                                                                                
… +24 more                                                                      
… more hidden                                                                   
! source truncated (count unknown)                                              
! source truncated (42 omitted)
```

Monochrome · none · 80 × 18 terminal cells

```
text/omission / normal                                                          
                                                                                
+24 more                                                                        
more hidden                                                                     
source truncated (count unknown)                                                
source truncated (42 omitted)
```

Bright · truecolor · 80 × 18 terminal cells

```
text/omission / bright                                                          
                                                                                
… +24 more                                                                      
… more hidden                                                                   
! source truncated (count unknown)                                              
! source truncated (42 omitted)
```

Normal · truecolor · 80 × 18 terminal cells

```
text/omission / normal                                                          
                                                                                
… +24 more                                                                      
… more hidden                                                                   
! source truncated (count unknown)                                              
! source truncated (42 omitted)
```

Dim · truecolor · 80 × 18 terminal cells

```
text/omission / dim                                                             
                                                                                
… +24 more                                                                      
… more hidden                                                                   
! source truncated (count unknown)                                              
! source truncated (42 omitted)
```

### [Go usage](https://lanej.io/charmed/atoms-text.html#text%2fomission-code)

Imports: [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
hidden := text.Omission{Count: 24, KnownCount: true, Emphasis: text.Normal}
truncated := text.Omission{Kind: text.SourceTruncated, Emphasis: text.Normal}

return hidden.View(palette) + "  " + truncated.View(palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-text.html#text%2fomission-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-text.html#text%2fomission-run)

```sh
go run ./examples/catalog text/omission
```

[Example source](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/docs/snippets/text_omission.go) · [Implementation](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/text/markers.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Text

## [Indicator](https://lanej.io/charmed/atoms-text.html#text%2findicator)

Render a supplied glyph, label and tone consistently. Default, GlyphOnly and LabelOnly use a pill; Compact uses a bare glyph without a label, fill or padding.

### [Native preview](https://lanej.io/charmed/atoms-text.html#text%2findicator-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 ↑ Up 
 ↑    
 Up   
↑
```

Demo · truecolor · 80 × 12 terminal cells

```
 ↑ Up                                                                           
 ↑                                                                              
 Up                                                                             
↑
```

Narrow · truecolor · 40 × 16 terminal cells

```
 ↑ Up                                   
 ↑                                      
 Up                                     
↑
```

256 colors · ansi256 · 80 × 12 terminal cells

```
 ↑ Up                                                                           
 ↑                                                                              
 Up                                                                             
↑
```

16 colors · ansi16 · 80 × 12 terminal cells

```
 ↑ Up                                                                           
 ↑                                                                              
 Up                                                                             
↑
```

Monochrome · none · 80 × 12 terminal cells

```
[^ Up]                                                                          
[^]                                                                             
[Up]                                                                            
^
```

Bright · truecolor · 80 × 12 terminal cells

```
 ↑ Up                                                                           
 ↑                                                                              
 Up                                                                             
↑
```

Normal · truecolor · 80 × 12 terminal cells

```
 ↑ Up                                                                           
 ↑                                                                              
 Up                                                                             
↑
```

Dim · truecolor · 80 × 12 terminal cells

```
 ↑ Up                                                                           
 ↑                                                                              
 Up                                                                             
↑
```

### [Go usage](https://lanej.io/charmed/atoms-text.html#text%2findicator-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
glyph := palette.Glyph(theme.GlyphUp)

return layout.Column(
	text.Indicator(glyph, "Up", text.Success, palette, text.Normal),
	text.Indicator(glyph, "Up", text.Success, palette, text.Normal, text.GlyphOnly),
	text.Indicator(glyph, "Up", text.Success, palette, text.Normal, text.LabelOnly),
	text.Indicator(glyph, "Up", text.Success, palette, text.Normal, text.Compact),
)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-text.html#text%2findicator-integration)

Supply glyphs through theme.Glyph and labels as plain text; controls are escaped. text.Compact retains tone and emphasis without pill geometry. Keep context or a legend for unlabeled marks. This renderer owns presentation only; the application owns state meaning.

### [Run this example](https://lanej.io/charmed/atoms-text.html#text%2findicator-run)

```sh
go run ./examples/catalog text/indicator
```

[Example source](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/docs/snippets/text_indicator.go) · [Implementation](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/text/text.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

### Reusable component

[Indicator dependency graph](https://lanej.io/charmed/dependencies.html#component-text%2Findicator) · Partial map · [Source declaration](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/text/text.go#L60)

- [Color pill](https://lanej.io/charmed/dependencies.html#component-text%2Fpill) · composes · labeled pill presentation

Consumer usage uses

- [Indicator](https://lanej.io/charmed/dependencies.html#component-text%2Findicator) · primary demonstration

[Example source](https://github.com/lanej/charmed/blob/4899458b504339096b00026ea0cdba306a9e5503/docs/snippets/text_indicator.go#L9)
