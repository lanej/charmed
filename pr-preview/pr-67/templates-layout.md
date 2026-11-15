[Templates](https://lanej.io/charmed/templates.html)Layout

# [Layout](https://lanej.io/charmed/templates-layout.html#level-templates-layout)

2 components · [All templates](https://lanej.io/charmed/templates.html)

- [Nested columns](https://lanej.io/charmed/templates-layout.html#template%2fnested-columns)
- [Grid inside a column](https://lanej.io/charmed/templates-layout.html#template%2fnested-grid)

[Templates](https://lanej.io/charmed/templates.html) / Layout

## [Nested columns](https://lanej.io/charmed/templates-layout.html#template%2fnested-columns)

Compose a host list beside a column of summary and evidence sections. Inspect each child allocation independently, then compare wide columns, vertically stacked panes and a compact inspection-only recipe.

### [Native preview](https://lanej.io/charmed/templates-layout.html#template%2fnested-columns-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
┌ hosts ───────────────┐ ┌ summary ────────────────────────────────────────────┐
│ host-01              │ │ host-02 · v2.8.1                                    │
│ host-02              │ │ Public pool                                         │
│ host-03              │ │                                                     │
│                      │ └─────────────────────────────────────────────────────┘
│                      │                                                        
│                      │ ┌ details ────────────────────────────────────────────┐
│                      │ │ Probe: HTTP 200                                     │
│                      │ │ Latency: 12 ms                                      │
│                      │ │ Retries: 0                                          │
│                      │ │                                                     │
│                      │ │                                                     │
│                      │ │                                                     │
│                      │ │                                                     │
│                      │ │                                                     │
└──────────────────────┘ └─────────────────────────────────────────────────────┘
```

Demo · truecolor · 100 × 20 terminal cells

```
┌ Hosts ─────────────────┐ ┌ Selected host ────────────────────────────────────────────────────────┐
│ ▤ host-01              │ │ ▤ host-02 · v2.8.1                                                    │
│ ▤ host-02              │ │  ✓ OK                                                                 │
│ ▤ host-03              │ │                                                                       │
│                        │ └───────────────────────────────────────────────────────────────────────┘
│                        │                                                                          
│                        │ ┌ Probe evidence ───────────────────────────────────────────────────────┐
│                        │ │ Pool: public                                                          │
│                        │ │ Probe: HTTP 200                                                       │
│                        │ │ Latency: 12 ms                                                        │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
└────────────────────────┘ └───────────────────────────────────────────────────────────────────────┘
```

Narrow · truecolor · 40 × 16 terminal cells

```
┌ Selected host ───────────────────────┐
│ ▤ host-02 · v2.8.1                   │
│  ✓ OK                                │
│                                      │
└──────────────────────────────────────┘
                                        
┌ Probe evidence ──────────────────────┐
│ Pool: public                         │
│ Probe: HTTP 200                      │
│ Latency: 12 ms                       │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
└──────────────────────────────────────┘
```

256 colors · ansi256 · 100 × 20 terminal cells

```
┌ Hosts ─────────────────┐ ┌ Selected host ────────────────────────────────────────────────────────┐
│ ▤ host-01              │ │ ▤ host-02 · v2.8.1                                                    │
│ ▤ host-02              │ │  ✓ OK                                                                 │
│ ▤ host-03              │ │                                                                       │
│                        │ └───────────────────────────────────────────────────────────────────────┘
│                        │                                                                          
│                        │ ┌ Probe evidence ───────────────────────────────────────────────────────┐
│                        │ │ Pool: public                                                          │
│                        │ │ Probe: HTTP 200                                                       │
│                        │ │ Latency: 12 ms                                                        │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
└────────────────────────┘ └───────────────────────────────────────────────────────────────────────┘
```

16 colors · ansi16 · 100 × 20 terminal cells

```
┌ Hosts ─────────────────┐ ┌ Selected host ────────────────────────────────────────────────────────┐
│ ▤ host-01              │ │ ▤ host-02 · v2.8.1                                                    │
│ ▤ host-02              │ │  ✓ OK                                                                 │
│ ▤ host-03              │ │                                                                       │
│                        │ └───────────────────────────────────────────────────────────────────────┘
│                        │                                                                          
│                        │ ┌ Probe evidence ───────────────────────────────────────────────────────┐
│                        │ │ Pool: public                                                          │
│                        │ │ Probe: HTTP 200                                                       │
│                        │ │ Latency: 12 ms                                                        │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
└────────────────────────┘ └───────────────────────────────────────────────────────────────────────┘
```

Monochrome · none · 100 × 20 terminal cells

```
+ Hosts -----------------+ + Selected host --------------------------------------------------------+
| H host-01              | | H host-02 * v2.8.1                                                    |
| H host-02              | | [+ OK]                                                                |
| H host-03              | |                                                                       |
|                        | +-----------------------------------------------------------------------+
|                        |                                                                          
|                        | + Probe evidence -------------------------------------------------------+
|                        | | Pool: public                                                          |
|                        | | Probe: HTTP 200                                                       |
|                        | | Latency: 12 ms                                                        |
|                        | |                                                                       |
|                        | |                                                                       |
|                        | |                                                                       |
|                        | |                                                                       |
|                        | |                                                                       |
|                        | |                                                                       |
|                        | |                                                                       |
|                        | |                                                                       |
|                        | |                                                                       |
+------------------------+ +-----------------------------------------------------------------------+
```

Stacked · truecolor · 40 × 24 terminal cells

```
┌ Hosts ───────────────────────────────┐
│ ▤ host-01                            │
│ ▤ host-02                            │
│ ▤ host-03                            │
│                                      │
└──────────────────────────────────────┘
                                        
┌ Selected host ───────────────────────┐
│ ▤ host-02 · v2.8.1                   │
│  ✓ OK                                │
│                                      │
└──────────────────────────────────────┘
                                        
┌ Probe evidence ──────────────────────┐
│ Pool: public                         │
│ Probe: HTTP 200                      │
│ Latency: 12 ms                       │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
└──────────────────────────────────────┘
```

### [Go usage](https://lanej.io/charmed/templates-layout.html#template%2fnested-columns-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme") · [`tui`](https://pkg.go.dev/github.com/lanej/charmed/tui "github.com/lanej/charmed/tui")

```go
screen, err := tui.NewScreen("hosts", "summary", "details")
if err != nil {
	return "", err
}

summary, details := &layout.Region{ID: "summary"}, &layout.Region{ID: "details"}
inspection := &layout.Flex{
	Name: "inspection", Axis: layout.Vertical, Gap: 1,
	Exceptions: []layout.Exception{{Rule: layout.StackOpportunity, Reason: "The summary introduces the evidence below it."}},
	Children:   []layout.Slot{{Node: summary, Track: layout.Fixed(5)}, {Node: details, Track: layout.Grow(5)}},
}

root := &layout.Flex{Gap: 1, Children: []layout.Slot{
	{Node: &layout.Region{ID: "hosts"}, Track: layout.Fixed(24)},
	{Node: inspection, Track: layout.Grow(30)},
}}
if arrangeErr := screen.Arrange(root, layout.Size{Width: 80, Height: 16}); arrangeErr != nil {
	return "", arrangeErr
}

palette := theme.New(theme.Dark)
content := map[string]string{"hosts": "host-01\nhost-02\nhost-03", "summary": "host-02 · v2.8.1\nPublic pool", "details": "Probe: HTTP 200\nLatency: 12 ms\nRetries: 0"}
views := map[string]string{}

for _, region := range screen.Plan().Regions() {
	size := layout.Size{Width: region.Bounds.Dx(), Height: region.Bounds.Dy()}
	views[region.ID] = layout.Panel(region.ID, content[region.ID], size, palette, false)
}
// Resize native children from each leaf's PanelContentSize before rendering.
// The parent owns focus and forwards commands to its active child.
frame, err := screen.Render(views)

return frame.Content, err
```

### [Updates, focus & sizing](https://lanej.io/charmed/templates-layout.html#template%2fnested-columns-integration)

Nest Flex, Grid and Region nodes directly; only leaf region IDs are registered with tui.Screen. Arrange on resize, allocate native children from each leaf rectangle and render exactly the visible views. In the interactive workbench, Tab moves focus and Ctrl+L toggles region bounds; the catalog command prints a static consumer frame. Wide, stacked and compact recipes share the same leaves. Compact mode explicitly omits the host list; a scoped stacking exception keeps summary and evidence together.

### [Run this example](https://lanej.io/charmed/templates-layout.html#template%2fnested-columns-run)

```sh
go run ./examples/catalog template/nested-columns
```

[Example source](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/template_nested_columns.go) · [Implementation](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/story/nested_layouts.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [tui/screen](https://lanej.io/charmed/templates-tui.html#tui/screen) [layout/flex](https://lanej.io/charmed/molecules-layout.html#layout/flex) [layout/responsive](https://lanej.io/charmed/molecules-layout.html#layout/responsive) [feedback/status](https://lanej.io/charmed/atoms-feedback.html#feedback/status)

[Templates](https://lanej.io/charmed/templates.html) / Layout

## [Grid inside a column](https://lanej.io/charmed/templates-layout.html#template%2fnested-grid)

Place a metrics grid inside an inspection column, between the selected host summary and probe evidence. Use this example when a screen needs rows within columns and columns within those rows.

### [Native preview](https://lanej.io/charmed/templates-layout.html#template%2fnested-grid-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
┌ hosts ───────────────┐ ┌ cpu ────────────────────┐ ┌ memory ─────────────────┐
│ host-01              │ │ 43%                     │ │ 61%                     │
│ host-02              │ │ 4 cores                 │ │ 16 GiB                  │
│ host-03              │ │                         │ │                         │
│                      │ └─────────────────────────┘ └─────────────────────────┘
│                      │                                                        
│                      │ ┌ details ────────────────────────────────────────────┐
│                      │ │ host-02 · v2.8.1                                    │
│                      │ │ Probe: HTTP 200                                     │
│                      │ │ Latency: 12 ms                                      │
│                      │ │                                                     │
│                      │ │                                                     │
│                      │ │                                                     │
│                      │ │                                                     │
│                      │ │                                                     │
└──────────────────────┘ └─────────────────────────────────────────────────────┘
```

Demo · truecolor · 100 × 20 terminal cells

```
┌ Hosts ─────────────────┐ ┌ Selected host ────────────────────────────────────────────────────────┐
│ ▤ host-01              │ │ ▤ host-02 · v2.8.1                                                    │
│ ▤ host-02              │ │  ✓ OK                                                                 │
│ ▤ host-03              │ └───────────────────────────────────────────────────────────────────────┘
│                        │                                                                          
│                        │ ┌ CPU ─────────────────────────────┐ ┌ Memory ──────────────────────────┐
│                        │ │ 43%                              │ │ 61%                              │
│                        │ │ 4 cores                          │ │ 16 GiB                           │
│                        │ │                                  │ │                                  │
│                        │ └──────────────────────────────────┘ └──────────────────────────────────┘
│                        │                                                                          
│                        │ ┌ Probe evidence ───────────────────────────────────────────────────────┐
│                        │ │ Pool: public                                                          │
│                        │ │ Probe: HTTP 200                                                       │
│                        │ │ Latency: 12 ms                                                        │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
└────────────────────────┘ └───────────────────────────────────────────────────────────────────────┘
```

Narrow · truecolor · 40 × 16 terminal cells

```
┌ Selected host ───────────────────────┐
│ ▤ host-02 · v2.8.1                   │
│  ✓ OK                                │
└──────────────────────────────────────┘
                                        
┌ CPU ─────────────┐ ┌ Memory ─────────┐
│ 43%              │ │ 61%             │
│ 4 cores          │ │ 16 GiB          │
│                  │ │                 │
└──────────────────┘ └─────────────────┘
                                        
┌ Probe evidence ──────────────────────┐
│ Pool: public                         │
│ Probe: HTTP 200                      │
│ Latency: 12 ms                       │
└──────────────────────────────────────┘
```

256 colors · ansi256 · 100 × 20 terminal cells

```
┌ Hosts ─────────────────┐ ┌ Selected host ────────────────────────────────────────────────────────┐
│ ▤ host-01              │ │ ▤ host-02 · v2.8.1                                                    │
│ ▤ host-02              │ │  ✓ OK                                                                 │
│ ▤ host-03              │ └───────────────────────────────────────────────────────────────────────┘
│                        │                                                                          
│                        │ ┌ CPU ─────────────────────────────┐ ┌ Memory ──────────────────────────┐
│                        │ │ 43%                              │ │ 61%                              │
│                        │ │ 4 cores                          │ │ 16 GiB                           │
│                        │ │                                  │ │                                  │
│                        │ └──────────────────────────────────┘ └──────────────────────────────────┘
│                        │                                                                          
│                        │ ┌ Probe evidence ───────────────────────────────────────────────────────┐
│                        │ │ Pool: public                                                          │
│                        │ │ Probe: HTTP 200                                                       │
│                        │ │ Latency: 12 ms                                                        │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
└────────────────────────┘ └───────────────────────────────────────────────────────────────────────┘
```

16 colors · ansi16 · 100 × 20 terminal cells

```
┌ Hosts ─────────────────┐ ┌ Selected host ────────────────────────────────────────────────────────┐
│ ▤ host-01              │ │ ▤ host-02 · v2.8.1                                                    │
│ ▤ host-02              │ │  ✓ OK                                                                 │
│ ▤ host-03              │ └───────────────────────────────────────────────────────────────────────┘
│                        │                                                                          
│                        │ ┌ CPU ─────────────────────────────┐ ┌ Memory ──────────────────────────┐
│                        │ │ 43%                              │ │ 61%                              │
│                        │ │ 4 cores                          │ │ 16 GiB                           │
│                        │ │                                  │ │                                  │
│                        │ └──────────────────────────────────┘ └──────────────────────────────────┘
│                        │                                                                          
│                        │ ┌ Probe evidence ───────────────────────────────────────────────────────┐
│                        │ │ Pool: public                                                          │
│                        │ │ Probe: HTTP 200                                                       │
│                        │ │ Latency: 12 ms                                                        │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
│                        │ │                                                                       │
└────────────────────────┘ └───────────────────────────────────────────────────────────────────────┘
```

Monochrome · none · 100 × 20 terminal cells

```
+ Hosts -----------------+ + Selected host --------------------------------------------------------+
| H host-01              | | H host-02 * v2.8.1                                                    |
| H host-02              | | [+ OK]                                                                |
| H host-03              | +-----------------------------------------------------------------------+
|                        |                                                                          
|                        | + CPU -----------------------------+ + Memory --------------------------+
|                        | | 43%                              | | 61%                              |
|                        | | 4 cores                          | | 16 GiB                           |
|                        | |                                  | |                                  |
|                        | +----------------------------------+ +----------------------------------+
|                        |                                                                          
|                        | + Probe evidence -------------------------------------------------------+
|                        | | Pool: public                                                          |
|                        | | Probe: HTTP 200                                                       |
|                        | | Latency: 12 ms                                                        |
|                        | |                                                                       |
|                        | |                                                                       |
|                        | |                                                                       |
|                        | |                                                                       |
+------------------------+ +-----------------------------------------------------------------------+
```

Stacked · truecolor · 40 × 24 terminal cells

```
┌ Hosts ───────────────────────────────┐
│ ▤ host-01                            │
│ ▤ host-02                            │
│ ▤ host-03                            │
│                                      │
└──────────────────────────────────────┘
                                        
┌ Selected host ───────────────────────┐
│ ▤ host-02 · v2.8.1                   │
│  ✓ OK                                │
└──────────────────────────────────────┘
                                        
┌ CPU ─────────────┐ ┌ Memory ─────────┐
│ 43%              │ │ 61%             │
│ 4 cores          │ │ 16 GiB          │
│                  │ │                 │
└──────────────────┘ └─────────────────┘
                                        
┌ Probe evidence ──────────────────────┐
│ Pool: public                         │
│ Probe: HTTP 200                      │
│ Latency: 12 ms                       │
│                                      │
└──────────────────────────────────────┘
```

### [Go usage](https://lanej.io/charmed/templates-layout.html#template%2fnested-grid-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme") · [`tui`](https://pkg.go.dev/github.com/lanej/charmed/tui "github.com/lanej/charmed/tui")

```go
screen, err := tui.NewScreen("hosts", "cpu", "memory", "details")
if err != nil {
	return "", err
}

metrics := &layout.Grid{Columns: []layout.Track{layout.Grow(14), layout.Grow(14)}, Rows: []layout.Track{layout.Grow(5)}, Gap: layout.Size{Width: 1}, Cells: []layout.Cell{
	{Node: &layout.Region{ID: "cpu"}}, {Node: &layout.Region{ID: "memory"}, Column: 1},
}}
inspection := &layout.Flex{
	Name: "inspection", Axis: layout.Vertical, Gap: 1,
	Exceptions: []layout.Exception{{Rule: layout.StackOpportunity, Reason: "Metrics precede their supporting evidence in this column."}},
	Children:   []layout.Slot{{Node: metrics, Track: layout.Fixed(5)}, {Node: &layout.Region{ID: "details"}, Track: layout.Grow(5)}},
}
columns := &layout.Flex{Gap: 1, Children: []layout.Slot{{Node: &layout.Region{ID: "hosts"}, Track: layout.Fixed(24)}, {Node: inspection, Track: layout.Grow(30)}}}

root := &layout.Responsive{Name: "workspace", Variants: []layout.Variant{
	{Name: "columns", Min: layout.Size{Width: 80}, Root: columns},
	{Name: "inspection", Root: inspection, Hidden: []layout.Omission{{ID: "hosts", Reason: "Keep the selected host's metrics visible in narrow terminals."}}},
}}
if arrangeErr := screen.Arrange(root, layout.Size{Width: 80, Height: 16}); arrangeErr != nil {
	return "", arrangeErr
}

content := map[string]string{"hosts": "host-01\nhost-02\nhost-03", "cpu": "43%\n4 cores", "memory": "61%\n16 GiB", "details": "host-02 · v2.8.1\nProbe: HTTP 200\nLatency: 12 ms"}
views := map[string]string{}

for _, region := range screen.Plan().Regions() {
	size := layout.Size{Width: region.Bounds.Dx(), Height: region.Bounds.Dy()}
	views[region.ID] = layout.Panel(region.ID, content[region.ID], size, theme.New(theme.Dark), false)
}
// Every leaf is registered once; the compact recipe explicitly hides hosts.
// Call Arrange again on resize; retained native models stay with the parent.
frame, err := screen.Render(views)

return frame.Content, err
```

### [Updates, focus & sizing](https://lanej.io/charmed/templates-layout.html#template%2fnested-grid-integration)

Nest Flex, Grid and Region nodes directly; only leaf region IDs are registered with tui.Screen. Arrange on resize, allocate native children from each leaf rectangle and render exactly the visible views. In the interactive workbench, Tab moves focus and Ctrl+L toggles region bounds; the catalog command prints a static consumer frame. Wide, stacked and compact recipes share the same leaves. Compact mode explicitly omits the host list; a scoped stacking exception keeps summary and evidence together.

### [Run this example](https://lanej.io/charmed/templates-layout.html#template%2fnested-grid-run)

```sh
go run ./examples/catalog template/nested-grid
```

[Example source](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/template_nested_grid.go) · [Implementation](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/story/nested_layouts.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [tui/screen](https://lanej.io/charmed/templates-tui.html#tui/screen) [layout/flex](https://lanej.io/charmed/molecules-layout.html#layout/flex) [layout/grid](https://lanej.io/charmed/molecules-layout.html#layout/grid) [layout/responsive](https://lanej.io/charmed/molecules-layout.html#layout/responsive) [feedback/status](https://lanej.io/charmed/atoms-feedback.html#feedback/status)
