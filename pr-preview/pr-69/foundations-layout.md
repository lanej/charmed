[Foundations](https://lanej.io/charmed/foundations.html)Layout

# [Layout](https://lanej.io/charmed/foundations-layout.html#level-foundations-layout)

1 component · [All foundations](https://lanej.io/charmed/foundations.html)

- [Layout](https://lanej.io/charmed/foundations-layout.html#foundations%2flayout)

[Foundations](https://lanej.io/charmed/foundations.html) / Layout

## [Layout](https://lanej.io/charmed/foundations-layout.html#foundations%2flayout)

Allocate terminal-cell rectangles before rendering a composition. Panels reserve borders and gutters so child components receive the space they can actually use.

### [Native preview](https://lanej.io/charmed/foundations-layout.html#foundations%2flayout-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
┌ Summary ────────────────────┐ ┌ Details ─────────────────────────────────────┐
│ 12 services / all healthy   │ │ Region: us-west-2                            │
│                             │ │                                              │
│                             │ │                                              │
│                             │ │                                              │
└─────────────────────────────┘ └──────────────────────────────────────────────┘
```

Demo · truecolor · 100 × 30 terminal cells

```
Responsive layout                                                                                   
                                                                                                    
┌ Navigation ─────────────────────────┐ ┌ Content ─────────────────────────────────────────────────┐
│ Cell-based dimensions               │ │ The right pane adapts to available space.                │
│ Borders included in size            │ │ Narrow terminals stack the same panels.                  │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
└─────────────────────────────────────┘ └──────────────────────────────────────────────────────────┘
```

Narrow · truecolor · 40 × 16 terminal cells

```
Stacked layout                          
                                        
┌ Navigation ──────────────────────────┐
│ Same content                         │
│                                      │
│                                      │
│                                      │
└──────────────────────────────────────┘
┌ Content ─────────────────────────────┐
│ Same state                           │
│                                      │
│                                      │
│                                      │
└──────────────────────────────────────┘
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Responsive layout                                                                                   
                                                                                                    
┌ Navigation ─────────────────────────┐ ┌ Content ─────────────────────────────────────────────────┐
│ Cell-based dimensions               │ │ The right pane adapts to available space.                │
│ Borders included in size            │ │ Narrow terminals stack the same panels.                  │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
└─────────────────────────────────────┘ └──────────────────────────────────────────────────────────┘
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Responsive layout                                                                                   
                                                                                                    
┌ Navigation ─────────────────────────┐ ┌ Content ─────────────────────────────────────────────────┐
│ Cell-based dimensions               │ │ The right pane adapts to available space.                │
│ Borders included in size            │ │ Narrow terminals stack the same panels.                  │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
│                                     │ │                                                          │
└─────────────────────────────────────┘ └──────────────────────────────────────────────────────────┘
```

Monochrome · none · 100 × 30 terminal cells

```
Responsive layout                                                                                   
                                                                                                    
+ Navigation -------------------------+ + Content -------------------------------------------------+
| Cell-based dimensions               | | The right pane adapts to available space.                |
| Borders included in size            | | Narrow terminals stack the same panels.                  |
|                                     | |                                                          |
|                                     | |                                                          |
|                                     | |                                                          |
|                                     | |                                                          |
|                                     | |                                                          |
|                                     | |                                                          |
|                                     | |                                                          |
|                                     | |                                                          |
|                                     | |                                                          |
|                                     | |                                                          |
|                                     | |                                                          |
|                                     | |                                                          |
|                                     | |                                                          |
|                                     | |                                                          |
|                                     | |                                                          |
|                                     | |                                                          |
|                                     | |                                                          |
|                                     | |                                                          |
|                                     | |                                                          |
|                                     | |                                                          |
|                                     | |                                                          |
|                                     | |                                                          |
+-------------------------------------+ +----------------------------------------------------------+
```

### [Go usage](https://lanej.io/charmed/foundations-layout.html#foundations%2flayout-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
left, right, sideBySide := layout.Split(80, 1, 24, 32, .4)
summary := layout.Panel("Summary", "12 services / all healthy",
	layout.Size{Width: left, Height: 6}, palette, true)

details := layout.Panel("Details", "Region: us-west-2",
	layout.Size{Width: right, Height: 6}, palette, false)
if sideBySide {
	return layout.Row(summary, " ", details)
}

return layout.Column(summary, details)
```

### [Updates, focus & sizing](https://lanej.io/charmed/foundations-layout.html#foundations%2flayout-integration)

The parent owns state, focus and input routing; layout helpers render the values supplied to them. Measure headings, borders and gutters before allocating child rectangles. Re-render from the new terminal-cell dimensions on resize, and reserve the actual rendered height when text may wrap.

### [Run this example](https://lanej.io/charmed/foundations-layout.html#foundations%2flayout-run)

```sh
go run ./examples/catalog foundations/layout
```

[Example source](https://github.com/lanej/charmed/blob/ebf9018696944bf7fff7b999487964b17085a58d/docs/snippets/foundations_layout.go) · [Implementation](https://github.com/lanej/charmed/blob/ebf9018696944bf7fff7b999487964b17085a58d/story/foundations.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.
