[Templates](https://lanej.io/charmed/templates.html)tui

# [tui](https://lanej.io/charmed/templates-tui.html#level-templates-tui)

4 components · [All templates](https://lanej.io/charmed/templates.html)

- [Declared screen](https://lanej.io/charmed/templates-tui.html#tui%2fscreen)
- [Region inspector](https://lanej.io/charmed/templates-tui.html#tui%2fregions)
- [Strict layout validation](https://lanej.io/charmed/templates-tui.html#tui%2fvalidation)
- [Static declared render](https://lanej.io/charmed/templates-tui.html#tui%2fstatic)

[Templates](https://lanej.io/charmed/templates.html) / tui

## [Declared screen](https://lanej.io/charmed/templates-tui.html#tui%2fscreen)

Compose a registered set of terminal regions with strict allocation and content validation.

### [Native preview](https://lanej.io/charmed/templates-tui.html#tui%2fscreen-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Service overview                                                                
Native component view
```

Demo · truecolor · 100 × 20 terminal cells

```
┌ primary ──────────────────────┐ ┌ secondary ──────────────────────┐ ┌ tools ─────────────────────┐
│ 33×20                         │ │ 35×20                           │ │ 30×20                      │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
└───────────────────────────────┘ └─────────────────────────────────┘ └────────────────────────────┘
```

Narrow · truecolor · 40 × 16 terminal cells

```
┌ primary ─────────────────────────────┐
│ 40×16                                │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
└──────────────────────────────────────┘
```

256 colors · ansi256 · 100 × 20 terminal cells

```
┌ primary ──────────────────────┐ ┌ secondary ──────────────────────┐ ┌ tools ─────────────────────┐
│ 33×20                         │ │ 35×20                           │ │ 30×20                      │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
└───────────────────────────────┘ └─────────────────────────────────┘ └────────────────────────────┘
```

16 colors · ansi16 · 100 × 20 terminal cells

```
┌ primary ──────────────────────┐ ┌ secondary ──────────────────────┐ ┌ tools ─────────────────────┐
│ 33×20                         │ │ 35×20                           │ │ 30×20                      │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
└───────────────────────────────┘ └─────────────────────────────────┘ └────────────────────────────┘
```

Monochrome · none · 100 × 20 terminal cells

```
+ primary ----------------------+ + secondary ----------------------+ + tools ---------------------+
| 33x20                         | | 35x20                           | | 30x20                      |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
+-------------------------------+ +---------------------------------+ +----------------------------+
```

### [Go usage](https://lanej.io/charmed/templates-tui.html#tui%2fscreen-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`tui`](https://pkg.go.dev/github.com/lanej/charmed/tui "github.com/lanej/charmed/tui")

```go
screen, err := tui.NewScreen("header", "body")
if err != nil {
	return "", err
}

root := &layout.Dock{Top: &layout.Region{ID: "header", Role: layout.ChromeRegion}, TopHeight: 1, Body: &layout.Region{ID: "body"}}
if arrangeErr := screen.Arrange(root, layout.Size{Width: 80, Height: 12}); arrangeErr != nil {
	return "", arrangeErr
}

frame, err := screen.Render(map[string]string{"header": "Service overview", "body": "Native component view"})

return frame.Content, err
```

### [Updates, focus & sizing](https://lanej.io/charmed/templates-tui.html#tui%2fscreen-integration)

The screen owns allocation validation and composition, while the application owns Bubble Tea messages, commands, focus and component state. Arrange must succeed before rendering; a failed resize invalidates the current arrangement. Pass only visible views and inspect accepted exceptions through diagnostics. Validate representative terminal sizes with CheckLayouts, which requires every declared responsive branch to be exercised.

### [Run this example](https://lanej.io/charmed/templates-tui.html#tui%2fscreen-run)

```sh
go run ./examples/catalog tui/screen
```

[Example source](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/docs/snippets/tui_screen.go) · [Implementation](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/story/layouts.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [layout/responsive](https://lanej.io/charmed/molecules-layout.html#layout/responsive) [layout/dock](https://lanej.io/charmed/molecules-layout.html#layout/dock)

[Templates](https://lanej.io/charmed/templates.html) / tui

## [Region inspector](https://lanej.io/charmed/templates-tui.html#tui%2fregions)

Inspect the exact cell rectangles used to resize and compose native views.

### [Native preview](https://lanej.io/charmed/templates-tui.html#tui%2fregions-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
header (0,0) 80x1                                                               
┌ body ────────────────────────────────────────────────────────────────────────┐
│ body (0,1) 80x11                                                             │
│                                                                              │
│                                                                              │
│                                                                              │
│                                                                              │
│                                                                              │
│                                                                              │
│                                                                              │
│                                                                              │
└──────────────────────────────────────────────────────────────────────────────┘
```

Demo · truecolor · 100 × 20 terminal cells

```
┌ primary ──────────────────────┐ ┌ secondary ──────────────────────┐ ┌ tools ─────────────────────┐
│ primary (0,0) 33x20           │ │ secondary (34,0) 35x20          │ │ tools (70,0) 30x20         │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
└───────────────────────────────┘ └─────────────────────────────────┘ └────────────────────────────┘
```

Narrow · truecolor · 40 × 16 terminal cells

```
┌ primary ─────────────────────────────┐
│ primary (0,0) 40x16                  │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
└──────────────────────────────────────┘
```

256 colors · ansi256 · 100 × 20 terminal cells

```
┌ primary ──────────────────────┐ ┌ secondary ──────────────────────┐ ┌ tools ─────────────────────┐
│ primary (0,0) 33x20           │ │ secondary (34,0) 35x20          │ │ tools (70,0) 30x20         │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
└───────────────────────────────┘ └─────────────────────────────────┘ └────────────────────────────┘
```

16 colors · ansi16 · 100 × 20 terminal cells

```
┌ primary ──────────────────────┐ ┌ secondary ──────────────────────┐ ┌ tools ─────────────────────┐
│ primary (0,0) 33x20           │ │ secondary (34,0) 35x20          │ │ tools (70,0) 30x20         │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
│                               │ │                                 │ │                            │
└───────────────────────────────┘ └─────────────────────────────────┘ └────────────────────────────┘
```

Monochrome · none · 100 × 20 terminal cells

```
+ primary ----------------------+ + secondary ----------------------+ + tools ---------------------+
| primary (0,0) 33x20           | | secondary (34,0) 35x20          | | tools (70,0) 30x20         |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
|                               | |                                 | |                            |
+-------------------------------+ +---------------------------------+ +----------------------------+
```

### [Go usage](https://lanej.io/charmed/templates-tui.html#tui%2fregions-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme") · [`tui`](https://pkg.go.dev/github.com/lanej/charmed/tui "github.com/lanej/charmed/tui")

```go
screen, err := tui.NewScreen("header", "body")
if err != nil {
	return "", err
}

root := &layout.Dock{Top: &layout.Region{ID: "header", Role: layout.ChromeRegion}, TopHeight: 1, Body: &layout.Region{ID: "body"}}
if arrangeErr := screen.Arrange(root, layout.Size{Width: 80, Height: 12}); arrangeErr != nil {
	return "", arrangeErr
}

return screen.RegionMap(theme.New(theme.Dark)), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/templates-tui.html#tui%2fregions-integration)

The screen owns allocation validation and composition, while the application owns Bubble Tea messages, commands, focus and component state. Arrange must succeed before rendering; a failed resize invalidates the current arrangement. Pass only visible views and inspect accepted exceptions through diagnostics. Validate representative terminal sizes with CheckLayouts, which requires every declared responsive branch to be exercised.

### [Run this example](https://lanej.io/charmed/templates-tui.html#tui%2fregions-run)

```sh
go run ./examples/catalog tui/regions
```

[Example source](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/docs/snippets/tui_regions.go) · [Implementation](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/story/layouts.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [tui/screen](https://lanej.io/charmed/templates-tui.html#tui/screen)

[Templates](https://lanej.io/charmed/templates.html) / tui

## [Strict layout validation](https://lanej.io/charmed/templates-tui.html#tui%2fvalidation)

See strict layout rejection and a scoped exception with a reviewable rationale.

### [Native preview](https://lanej.io/charmed/templates-tui.html#tui%2fvalidation-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
layout rule stack-opportunity at steps: 2 content groups are stacked at 80
columns although their effective minimum widths fit a row; consider columns or
declare a scoped exception
```

Demo · truecolor · 100 × 20 terminal cells

```
BLOCKED: layout rule stack-opportunity at steps: 2 content groups are stacked at 100 columns        
although their effective minimum widths fit a row; consider columns or declare a scoped exception   
                                                                                                    
Accepted local exception                                                                            
Chronological steps are read top to bottom.
```

Narrow · truecolor · 40 × 16 terminal cells

```
BLOCKED: layout rule stack-opportunity  
at steps: 2 content groups are stacked  
at 40 columns although their effective  
minimum widths fit a row; consider      
columns or declare a scoped exception   
                                        
Accepted local exception                
Chronological steps are read top to     
bottom.
```

256 colors · ansi256 · 100 × 20 terminal cells

```
BLOCKED: layout rule stack-opportunity at steps: 2 content groups are stacked at 100 columns        
although their effective minimum widths fit a row; consider columns or declare a scoped exception   
                                                                                                    
Accepted local exception                                                                            
Chronological steps are read top to bottom.
```

16 colors · ansi16 · 100 × 20 terminal cells

```
BLOCKED: layout rule stack-opportunity at steps: 2 content groups are stacked at 100 columns        
although their effective minimum widths fit a row; consider columns or declare a scoped exception   
                                                                                                    
Accepted local exception                                                                            
Chronological steps are read top to bottom.
```

Monochrome · none · 100 × 20 terminal cells

```
BLOCKED: layout rule stack-opportunity at steps: 2 content groups are stacked at 100 columns        
although their effective minimum widths fit a row; consider columns or declare a scoped exception   
                                                                                                    
Accepted local exception                                                                            
Chronological steps are read top to bottom.
```

### [Go usage](https://lanej.io/charmed/templates-tui.html#tui%2fvalidation-code)

Imports: [`fmt`](https://pkg.go.dev/fmt "fmt") · [`ansi`](https://pkg.go.dev/github.com/charmbracelet/x/ansi "github.com/charmbracelet/x/ansi") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`tui`](https://pkg.go.dev/github.com/lanej/charmed/tui "github.com/lanej/charmed/tui")

```go
screen, err := tui.NewScreen("first", "second")
if err != nil {
	return "", err
}

root := &layout.Flex{Name: "steps", Axis: layout.Vertical, Children: []layout.Slot{
	{Node: &layout.Region{ID: "first", Min: layout.Size{Width: 20}}, Track: layout.Grow(1)},
	{Node: &layout.Region{ID: "second", Min: layout.Size{Width: 20}}, Track: layout.Grow(1)},
}}

err = screen.Arrange(root, layout.Size{Width: 80, Height: 12})
if err == nil {
	return "", fmt.Errorf("expected strict stacking validation to block rendering")
}

return ansi.Wrap(err.Error(), 80, ""), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/templates-tui.html#tui%2fvalidation-integration)

The screen owns allocation validation and composition, while the application owns Bubble Tea messages, commands, focus and component state. Arrange must succeed before rendering; a failed resize invalidates the current arrangement. Pass only visible views and inspect accepted exceptions through diagnostics. Validate representative terminal sizes with CheckLayouts, which requires every declared responsive branch to be exercised.

### [Run this example](https://lanej.io/charmed/templates-tui.html#tui%2fvalidation-run)

```sh
go run ./examples/catalog tui/validation
```

[Example source](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/docs/snippets/tui_validation.go) · [Implementation](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/story/layouts.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [layout/flex](https://lanej.io/charmed/molecules-layout.html#layout/flex) [tui/screen](https://lanej.io/charmed/templates-tui.html#tui/screen)

[Templates](https://lanej.io/charmed/templates.html) / tui

## [Static declared render](https://lanej.io/charmed/templates-tui.html#tui%2fstatic)

Compose a declared layout for printed output whose height is measured from its content instead of a terminal.

### [Native preview](https://lanej.io/charmed/templates-tui.html#tui%2fstatic-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Totals for the period, grouped by    Each finding names its evidence and
outcome.                             its owner.
```

Demo · truecolor · 100 × 20 terminal cells

```
Summary                                            Detail                                           
Totals for the period, grouped by outcome. The     Each finding names its evidence and its owner.   
summary keeps its place beside the detail when     When the page narrows the detail is stacked below
the page is wide enough.                           the summary instead of being cut off.
```

Narrow · truecolor · 40 × 16 terminal cells

```
Summary                                 
Totals for the period, grouped by       
outcome. The summary keeps its place    
beside the detail when the page is wide 
enough.                                 
Detail                                  
Each finding names its evidence and its 
owner. When the page narrows the detail 
is stacked below the summary instead of 
being cut off.
```

256 colors · ansi256 · 100 × 20 terminal cells

```
Summary                                            Detail                                           
Totals for the period, grouped by outcome. The     Each finding names its evidence and its owner.   
summary keeps its place beside the detail when     When the page narrows the detail is stacked below
the page is wide enough.                           the summary instead of being cut off.
```

16 colors · ansi16 · 100 × 20 terminal cells

```
Summary                                            Detail                                           
Totals for the period, grouped by outcome. The     Each finding names its evidence and its owner.   
summary keeps its place beside the detail when     When the page narrows the detail is stacked below
the page is wide enough.                           the summary instead of being cut off.
```

Monochrome · none · 100 × 20 terminal cells

```
Summary                                            Detail                                           
Totals for the period, grouped by outcome. The     Each finding names its evidence and its owner.   
summary keeps its place beside the detail when     When the page narrows the detail is stacked below
the page is wide enough.                           the summary instead of being cut off.
```

### [Go usage](https://lanej.io/charmed/templates-tui.html#tui%2fstatic-code)

Imports: [`ansi`](https://pkg.go.dev/github.com/charmbracelet/x/ansi "github.com/charmbracelet/x/ansi") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`tui`](https://pkg.go.dev/github.com/lanej/charmed/tui "github.com/lanej/charmed/tui")

```go
declare := func(_ layout.Size, heights tui.StaticHeights) layout.Node {
	summary := &layout.Region{ID: "summary", Min: layout.Size{Width: 20, Height: heights.Height("summary")}}
	detail := &layout.Region{ID: "detail", Min: layout.Size{Width: 20, Height: heights.Height("detail")}}

	return &layout.Responsive{Name: "report", Variants: []layout.Variant{
		{Name: "side-by-side", Min: layout.Size{Width: 60, Height: 1}, Root: &layout.SplitPane{First: summary, Second: detail, Gap: 2}},
		{Name: "stacked", Root: &layout.Flex{Name: "stack", Axis: layout.Vertical, Exceptions: []layout.Exception{{Rule: layout.StackOpportunity, Reason: "Printed reports read top to bottom."}}, Children: []layout.Slot{
			{Node: summary, Track: layout.Fixed(heights.Height("summary"))},
			{Node: detail, Track: layout.Fixed(heights.Height("detail"))},
		}}},
	}}
}

text := map[string]string{
	"summary": "Totals for the period, grouped by outcome.",
	"detail":  "Each finding names its evidence and its owner.",
}

frame, err := tui.RenderStatic([]string{"summary", "detail"}, 72, declare, func(id string, width int) string {
	return ansi.Wrap(text[id], width, "")
})

return frame.Content, err
```

### [Updates, focus & sizing](https://lanej.io/charmed/templates-tui.html#tui%2fstatic-integration)

Pass the registered region identities, a width, a declaration that uses the measured heights and a render function that wraps each region to the width it is given. RenderStatic holds no state: call it again when the content or width changes. It rejects any arrangement, design-rule or oversize failure, so choose a stacked fallback in the caller. Hidden regions are not rendered and keep their reasons in the returned plan.

### [Run this example](https://lanej.io/charmed/templates-tui.html#tui%2fstatic-run)

```sh
go run ./examples/catalog tui/static
```

[Example source](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/docs/snippets/tui_static.go) · [Implementation](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/story/layouts.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [layout/responsive](https://lanej.io/charmed/molecules-layout.html#layout/responsive) [tui/screen](https://lanej.io/charmed/templates-tui.html#tui/screen)
