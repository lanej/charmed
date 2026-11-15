[Foundations](https://lanej.io/charmed/foundations.html)Theme

# [Theme](https://lanej.io/charmed/foundations-theme.html#level-foundations-theme)

4 components · [All foundations](https://lanej.io/charmed/foundations.html)

- [Theme](https://lanej.io/charmed/foundations-theme.html#foundations%2ftheme)
- [Glyphs](https://lanej.io/charmed/foundations-theme.html#foundations%2fglyphs)
- [Spacing, borders and selection](https://lanej.io/charmed/foundations-theme.html#foundations%2ftokens)
- [Focus and spacing](https://lanej.io/charmed/foundations-theme.html#foundations%2ffocus)

[Foundations](https://lanej.io/charmed/foundations.html) / Theme

## [Theme](https://lanej.io/charmed/foundations-theme.html#foundations%2ftheme)

Semantic roles, terminal palettes, contrast and no-color output.

### [Native preview](https://lanej.io/charmed/foundations-theme.html#foundations%2ftheme-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Primary content       
Supporting information
Accent / interaction  
 Success              
 Warning 
```

Demo · truecolor · 100 × 30 terminal cells

```
Semantic tokens / normal and dim                                                                    
                                                                                                    
Primary content                                                                                     
Supporting information                                                                              
Accent / interaction                                                                                
                                                                                                    
neutral    dim neutral                                                                              
success    dim success                                                                              
info       dim info                                                                                 
warning    dim warning                                                                              
danger     dim danger                                                                               
unknown    dim unknown                                                                              
off        dim off                                                                                  
link       dim link                                                                                 
                                                                                                    
Categorical data / independent palette                                                              
1 ━━  2 ━━  3 ━━  4 ━━  5 ━━
```

Narrow · truecolor · 40 × 16 terminal cells

```
Semantic tokens / normal and dim        
                                        
Primary content                         
Supporting information                  
Accent / interaction                    
                                        
neutral    dim neutral                  
success    dim success                  
info       dim info                     
warning    dim warning                  
danger     dim danger                   
unknown    dim unknown                  
off        dim off                      
link       dim link                     
                                        
Categorical data / independent palette
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Semantic tokens / normal and dim                                                                    
                                                                                                    
Primary content                                                                                     
Supporting information                                                                              
Accent / interaction                                                                                
                                                                                                    
neutral    dim neutral                                                                              
success    dim success                                                                              
info       dim info                                                                                 
warning    dim warning                                                                              
danger     dim danger                                                                               
unknown    dim unknown                                                                              
off        dim off                                                                                  
link       dim link                                                                                 
                                                                                                    
Categorical data / independent palette                                                              
1 ━━  2 ━━  3 ━━  4 ━━  5 ━━
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Semantic tokens / normal and dim                                                                    
                                                                                                    
Primary content                                                                                     
Supporting information                                                                              
Accent / interaction                                                                                
                                                                                                    
neutral    dim neutral                                                                              
success    dim success                                                                              
info       dim info                                                                                 
warning    dim warning                                                                              
danger     dim danger                                                                               
unknown    dim unknown                                                                              
off        dim off                                                                                  
link       dim link                                                                                 
                                                                                                    
Categorical data / independent palette                                                              
1 ━━  2 ━━  3 ━━  4 ━━  5 ━━
```

Monochrome · none · 100 × 30 terminal cells

```
Semantic tokens / normal and dim                                                                    
                                                                                                    
Primary content                                                                                     
Supporting information                                                                              
Accent / interaction                                                                                
                                                                                                    
neutral    dim neutral                                                                              
success    dim success                                                                              
info       dim info                                                                                 
warning    dim warning                                                                              
danger     dim danger                                                                               
unknown    dim unknown                                                                              
off        dim off                                                                                  
link       dim link                                                                                 
                                                                                                    
Categorical data / independent palette
```

### [Go usage](https://lanej.io/charmed/foundations-theme.html#foundations%2ftheme-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return layout.Column(
	palette.Text().Render("Primary content"),
	palette.Subtle().Render("Supporting information"),
	palette.Heading().Render("Accent / interaction"),
	text.Pill("Success", text.Success, palette),
	text.Pill("Warning", text.Warning, palette),
)
```

### [Updates, focus & sizing](https://lanej.io/charmed/foundations-theme.html#foundations%2ftheme-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/foundations-theme.html#foundations%2ftheme-run)

```sh
go run ./examples/catalog foundations/theme
```

[Example source](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/docs/snippets/foundations_theme.go) · [Implementation](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/story/foundations.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Foundations](https://lanej.io/charmed/foundations.html) / Theme

## [Glyphs](https://lanej.io/charmed/foundations-theme.html#foundations%2fglyphs)

Named semantic glyphs, explicit ASCII equivalents and adjacent labels.

### [Native preview](https://lanej.io/charmed/foundations-theme.html#foundations%2fglyphs-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
⌕ Search   
▾ Expanded 
▸ Collapsed
◉ Selected
```

Demo · truecolor · 100 × 30 terminal cells

```
ok              ✓   warning         !                                                               
error           ×   unknown         ?                                                               
off             ⏻   pending         ◷                                                               
running         ▶   paused          ‖                                                               
disabled        ⊘   connected       ●                                                               
disconnected    ○   info            i                                                               
search          ⌕   filter          ▽                                                               
sort            ↕   refresh         ↻                                                               
link            ↗   expand          ⤢                                                               
selected        ◉   unselected      ○                                                               
expanded        ▾   collapsed       ▸                                                               
up              ↑   down            ↓                                                               
left            ←   right           →                                                               
loading         …   unavailable     —                                                               
log             ≡   trace           ⋈                                                               
span            ─   upstream        ↑                                                               
downstream      ↓   scroll-track    │                                                               
scroll-thumb    ┃   checked         ☑                                                               
unchecked       ☐   mixed           ▣                                                               
action          ↵   busy            ⟳                                                               
delete          ⌫   inactive        ◌                                                               
added           ⊕   modified        ↺                                                               
removed         ⊖   unchanged       ≡                                                               
missing         ∅   redacted        ▧                                                               
clock           ◷   truncated       …                                                               
attribution-gap ≠   focus-start     ▏                                                               
focus-end       ▕   load-balancer   ⇄                                                               
host            ▤   healthy         ♥                                                               
ready           ✓   not-ready       ◇                                                               
canary          ◔   canary-on       ◉
```

Narrow · truecolor · 40 × 34 terminal cells

```
Meaning     U A     Meaning     U A     
                                        
ok              ✓   warning         !   
error           ×   unknown         ?   
off             ⏻   pending         ◷   
running         ▶   paused          ‖   
disabled        ⊘   connected       ●   
disconnected    ○   info            i   
search          ⌕   filter          ▽   
sort            ↕   refresh         ↻   
link            ↗   expand          ⤢   
selected        ◉   unselected      ○   
expanded        ▾   collapsed       ▸   
up              ↑   down            ↓   
left            ←   right           →   
loading         …   unavailable     —   
log             ≡   trace           ⋈   
span            ─   upstream        ↑   
downstream      ↓   scroll-track    │   
scroll-thumb    ┃   checked         ☑   
unchecked       ☐   mixed           ▣   
action          ↵   busy            ⟳   
delete          ⌫   inactive        ◌   
added           ⊕   modified        ↺   
removed         ⊖   unchanged       ≡   
missing         ∅   redacted        ▧   
clock           ◷   truncated       …   
attribution-gap ≠   focus-start     ▏   
focus-end       ▕   load-balancer   ⇄   
host            ▤   healthy         ♥   
ready           ✓   not-ready       ◇   
canary          ◔   canary-on       ◉   
canary-off      ◌   live            ●   
bar-fill        ━
```

256 colors · ansi256 · 100 × 30 terminal cells

```
ok              ✓   warning         !                                                               
error           ×   unknown         ?                                                               
off             ⏻   pending         ◷                                                               
running         ▶   paused          ‖                                                               
disabled        ⊘   connected       ●                                                               
disconnected    ○   info            i                                                               
search          ⌕   filter          ▽                                                               
sort            ↕   refresh         ↻                                                               
link            ↗   expand          ⤢                                                               
selected        ◉   unselected      ○                                                               
expanded        ▾   collapsed       ▸                                                               
up              ↑   down            ↓                                                               
left            ←   right           →                                                               
loading         …   unavailable     —                                                               
log             ≡   trace           ⋈                                                               
span            ─   upstream        ↑                                                               
downstream      ↓   scroll-track    │                                                               
scroll-thumb    ┃   checked         ☑                                                               
unchecked       ☐   mixed           ▣                                                               
action          ↵   busy            ⟳                                                               
delete          ⌫   inactive        ◌                                                               
added           ⊕   modified        ↺                                                               
removed         ⊖   unchanged       ≡                                                               
missing         ∅   redacted        ▧                                                               
clock           ◷   truncated       …                                                               
attribution-gap ≠   focus-start     ▏                                                               
focus-end       ▕   load-balancer   ⇄                                                               
host            ▤   healthy         ♥                                                               
ready           ✓   not-ready       ◇                                                               
canary          ◔   canary-on       ◉
```

16 colors · ansi16 · 100 × 30 terminal cells

```
ok              ✓   warning         !                                                               
error           ×   unknown         ?                                                               
off             ⏻   pending         ◷                                                               
running         ▶   paused          ‖                                                               
disabled        ⊘   connected       ●                                                               
disconnected    ○   info            i                                                               
search          ⌕   filter          ▽                                                               
sort            ↕   refresh         ↻                                                               
link            ↗   expand          ⤢                                                               
selected        ◉   unselected      ○                                                               
expanded        ▾   collapsed       ▸                                                               
up              ↑   down            ↓                                                               
left            ←   right           →                                                               
loading         …   unavailable     —                                                               
log             ≡   trace           ⋈                                                               
span            ─   upstream        ↑                                                               
downstream      ↓   scroll-track    │                                                               
scroll-thumb    ┃   checked         ☑                                                               
unchecked       ☐   mixed           ▣                                                               
action          ↵   busy            ⟳                                                               
delete          ⌫   inactive        ◌                                                               
added           ⊕   modified        ↺                                                               
removed         ⊖   unchanged       ≡                                                               
missing         ∅   redacted        ▧                                                               
clock           ◷   truncated       …                                                               
attribution-gap ≠   focus-start     ▏                                                               
focus-end       ▕   load-balancer   ⇄                                                               
host            ▤   healthy         ♥                                                               
ready           ✓   not-ready       ◇                                                               
canary          ◔   canary-on       ◉
```

Monochrome · none · 100 × 30 terminal cells

```
ok              +   warning         !                                                               
error           x   unknown         ?                                                               
off             -   pending         p                                                               
running         >   paused          =                                                               
disabled        /   connected       +                                                               
disconnected    -   info            i                                                               
search          /   filter          f                                                               
sort            s   refresh         r                                                               
link            @   expand          +                                                               
selected        *   unselected      o                                                               
expanded        v   collapsed       >                                                               
up              ^   down            v                                                               
left            <   right           >                                                               
loading         .   unavailable     -                                                               
log             L   trace           T                                                               
span            S   upstream        ^                                                               
downstream      v   scroll-track    |                                                               
scroll-thumb    #   checked         x                                                               
unchecked           mixed           -                                                               
action          >   busy            ~                                                               
delete          x   inactive        -                                                               
added           +   modified        ~                                                               
removed         -   unchanged       =                                                               
missing         -   redacted        #                                                               
clock           t   truncated       .                                                               
attribution-gap %   focus-start     [                                                               
focus-end       ]   load-balancer   L                                                               
host            H   healthy         H                                                               
ready           r   not-ready       -                                                               
canary          c   canary-on       +
```

### [Go usage](https://lanej.io/charmed/foundations-theme.html#foundations%2fglyphs-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return layout.Column(
	palette.Glyph(theme.GlyphSearch)+" Search",
	palette.Glyph(theme.GlyphExpanded)+" Expanded",
	palette.Glyph(theme.GlyphCollapsed)+" Collapsed",
	palette.Glyph(theme.GlyphSelected)+" Selected",
)
// Use theme.New(theme.Mono) for the matching ASCII glyphs.
```

### [Updates, focus & sizing](https://lanej.io/charmed/foundations-theme.html#foundations%2fglyphs-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/foundations-theme.html#foundations%2fglyphs-run)

```sh
go run ./examples/catalog foundations/glyphs
```

[Example source](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/docs/snippets/foundations_glyphs.go) · [Implementation](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/story/foundations.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Foundations](https://lanej.io/charmed/foundations.html) / Theme

## [Spacing, borders and selection](https://lanej.io/charmed/foundations-theme.html#foundations%2ftokens)

Cell spacing scale, normal and focused border roles, and readable collection selection.

### [Native preview](https://lanej.io/charmed/foundations-theme.html#foundations%2ftokens-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
┌─────────────────┐
│  ⌕ Search       │
│  Selected item  │
└─────────────────┘
```

Demo · truecolor · 100 × 30 terminal cells

```
Cell spacing / horizontal examples                                                                  
                                                                                                    
none   0  ||                                                                                        
small  1  | |                                                                                       
medium 2  |  |                                                                                      
large  4  |    |                                                                                    
                                                                                                    
┌ Normal ────────┐  ┌ Focused ───────┐                                                              
│ Border         │  │ FocusBorder    │                                                              
└────────────────┘  └────────────────┘                                                              
                                                                                                    
Selection / stable insets                                                                           
rest     | Value |                                                                                  
focused  | Value |                                                                                  
selected | Value |
```

Narrow · truecolor · 40 × 16 terminal cells

```
Cell spacing / horizontal examples      
                                        
none   0  ||                            
small  1  | |                           
medium 2  |  |                          
large  4  |    |                        
                                        
┌ Normal ────────┐  ┌ Focused ───────┐  
│ Border         │  │ FocusBorder    │  
└────────────────┘  └────────────────┘  
                                        
Selection / stable insets               
rest     | Value |                      
focused  | Value |                      
selected | Value |
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Cell spacing / horizontal examples                                                                  
                                                                                                    
none   0  ||                                                                                        
small  1  | |                                                                                       
medium 2  |  |                                                                                      
large  4  |    |                                                                                    
                                                                                                    
┌ Normal ────────┐  ┌ Focused ───────┐                                                              
│ Border         │  │ FocusBorder    │                                                              
└────────────────┘  └────────────────┘                                                              
                                                                                                    
Selection / stable insets                                                                           
rest     | Value |                                                                                  
focused  | Value |                                                                                  
selected | Value |
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Cell spacing / horizontal examples                                                                  
                                                                                                    
none   0  ||                                                                                        
small  1  | |                                                                                       
medium 2  |  |                                                                                      
large  4  |    |                                                                                    
                                                                                                    
┌ Normal ────────┐  ┌ Focused ───────┐                                                              
│ Border         │  │ FocusBorder    │                                                              
└────────────────┘  └────────────────┘                                                              
                                                                                                    
Selection / stable insets                                                                           
rest     | Value |                                                                                  
focused  | Value |                                                                                  
selected | Value |
```

Monochrome · none · 100 × 30 terminal cells

```
Cell spacing / horizontal examples                                                                  
                                                                                                    
none   0  ||                                                                                        
small  1  | |                                                                                       
medium 2  |  |                                                                                      
large  4  |    |                                                                                    
                                                                                                    
+ Normal --------+  + Focused -------+                                                              
| Border         |  | FocusBorder    |                                                              
+----------------+  +----------------+                                                              
                                                                                                    
Selection / stable insets                                                                           
rest     | Value |                                                                                  
focused  | Value |                                                                                  
selected | Value |
```

### [Go usage](https://lanej.io/charmed/foundations-theme.html#foundations%2ftokens-code)

Imports: [`lipgloss`](https://pkg.go.dev/charm.land/lipgloss/v2 "charm.land/lipgloss/v2") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
t := theme.New(theme.Dark)

// Keep shape and color separate; native Lip Gloss styles stay available.
frame := lipgloss.NewStyle().
	Border(t.FrameBorder()).
	BorderForeground(t.BorderStyle(true).GetForeground()).
	Padding(theme.SpaceNone, theme.InlinePadding)

selected := t.Selection().Render("Selected item")
control := t.Control().Render(t.Glyph(theme.GlyphSearch) + " Search")

return frame.Render(control + "\n" + selected)
```

### [Updates, focus & sizing](https://lanej.io/charmed/foundations-theme.html#foundations%2ftokens-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/foundations-theme.html#foundations%2ftokens-run)

```sh
go run ./examples/catalog foundations/tokens
```

[Example source](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/docs/snippets/foundations_tokens.go) · [Implementation](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/story/foundations.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Foundations](https://lanej.io/charmed/foundations.html) / Theme

## [Focus and spacing](https://lanej.io/charmed/foundations-theme.html#foundations%2ffocus)

Route keyboard input among nested regions while keeping only the active child focused. The parent owns navigation between regions and forwards child commands.

### [Native preview](https://lanej.io/charmed/foundations-theme.html#foundations%2ffocus-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 Normal control  
 Focused control 
 Selected row 
```

Demo · truecolor · 100 × 30 terminal cells

```
Stable control insets / 1 cell per side                                                             
                                                                                                    
Normal  | Value |                                                                                   
Focused | Value |                                                                                   
Row     | Value |                                                                                   
                                                                                                    
bright  | Value |                                                                                   
normal  | Value |                                                                                   
dim     | Value |                                                                                   
                                                                                                    
┌ Content gutter ──────────────────────────────────────────────────────────────────────────────────┐
│ Text starts one cell from the border.                                                            │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
```

Narrow · truecolor · 40 × 16 terminal cells

```
Stable control insets / 1 cell per side 
                                        
Normal  | Value |                       
Focused | Value |                       
Row     | Value |                       
                                        
bright  | Value |                       
normal  | Value |                       
dim     | Value |                       
                                        
┌ Content gutter ──────────────────────┐
│ Text starts one cell from the border │
└──────────────────────────────────────┘
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Stable control insets / 1 cell per side                                                             
                                                                                                    
Normal  | Value |                                                                                   
Focused | Value |                                                                                   
Row     | Value |                                                                                   
                                                                                                    
bright  | Value |                                                                                   
normal  | Value |                                                                                   
dim     | Value |                                                                                   
                                                                                                    
┌ Content gutter ──────────────────────────────────────────────────────────────────────────────────┐
│ Text starts one cell from the border.                                                            │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Stable control insets / 1 cell per side                                                             
                                                                                                    
Normal  | Value |                                                                                   
Focused | Value |                                                                                   
Row     | Value |                                                                                   
                                                                                                    
bright  | Value |                                                                                   
normal  | Value |                                                                                   
dim     | Value |                                                                                   
                                                                                                    
┌ Content gutter ──────────────────────────────────────────────────────────────────────────────────┐
│ Text starts one cell from the border.                                                            │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
```

Monochrome · none · 100 × 30 terminal cells

```
Stable control insets / 1 cell per side                                                             
                                                                                                    
Normal  | Value |                                                                                   
Focused | Value |                                                                                   
Row     | Value |                                                                                   
                                                                                                    
bright  |[Value]|                                                                                   
normal  |[Value]|                                                                                   
dim     |[Value]|                                                                                   
                                                                                                    
+ Content gutter ----------------------------------------------------------------------------------+
| Text starts one cell from the border.                                                            |
+--------------------------------------------------------------------------------------------------+
```

### [Go usage](https://lanej.io/charmed/foundations-theme.html#foundations%2ffocus-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return layout.Column(
	palette.Control().Render("Normal control"),
	palette.Focus().Render("Focused control"),
	palette.Selection().Render("Selected row"),
)
```

### [Updates, focus & sizing](https://lanej.io/charmed/foundations-theme.html#foundations%2ffocus-integration)

The parent owns state, focus and input routing; layout helpers render the values supplied to them. Measure headings, borders and gutters before allocating child rectangles. Re-render from the new terminal-cell dimensions on resize, and reserve the actual rendered height when text may wrap.

### [Run this example](https://lanej.io/charmed/foundations-theme.html#foundations%2ffocus-run)

```sh
go run ./examples/catalog foundations/focus
```

[Example source](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/docs/snippets/foundations_focus.go) · [Implementation](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/story/foundations.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.
