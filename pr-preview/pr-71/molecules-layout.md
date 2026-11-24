[Molecules](https://lanej.io/charmed/molecules.html)Layout

# [Layout](https://lanej.io/charmed/molecules-layout.html#level-molecules-layout)

13 components · [All molecules](https://lanej.io/charmed/molecules.html)

- [Text row](https://lanej.io/charmed/molecules-layout.html#layout%2ftext-row)
- [Link group](https://lanej.io/charmed/molecules-layout.html#layout%2flink-group)
- [Header](https://lanej.io/charmed/molecules-layout.html#layout%2fheader)
- [Identity header](https://lanej.io/charmed/molecules-layout.html#layout%2fidentity-header)
- [Foldable section](https://lanej.io/charmed/molecules-layout.html#layout%2fsection)
- [Key-value](https://lanej.io/charmed/molecules-layout.html#layout%2fkey-value)
- [Footer](https://lanej.io/charmed/molecules-layout.html#layout%2ffooter)
- [Weighted flex layout](https://lanej.io/charmed/molecules-layout.html#layout%2fflex)
- [Grid with spanning cells](https://lanej.io/charmed/molecules-layout.html#layout%2fgrid)
- [Wrapping flow layout](https://lanej.io/charmed/molecules-layout.html#layout%2fflow)
- [Docked layout](https://lanej.io/charmed/molecules-layout.html#layout%2fdock)
- [Constrained split panes](https://lanej.io/charmed/molecules-layout.html#layout%2fsplit-pane)
- [Responsive layout recipes](https://lanej.io/charmed/molecules-layout.html#layout%2fresponsive)

[Molecules](https://lanej.io/charmed/molecules.html) / Layout

## [Text row](https://lanej.io/charmed/molecules-layout.html#layout%2ftext-row)

Compose headers, operational status lines and compact summaries from styled text segments, fixed gaps and flexible space. Fields wrap in order rather than disappearing when the terminal narrows.

### [Native preview](https://lanej.io/charmed/molecules-layout.html#layout%2ftext-row-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
API operations  Production / us-east-1        Observed 00:16 UTC
```

Demo · truecolor · 64 × 8 terminal cells

```
API operations  Production / us-east-1        Observed 00:16 UTC
```

Narrow · truecolor · 40 × 16 terminal cells

```
API operations  Production / us-east-1  
                      Observed 00:16 UTC
```

256 colors · ansi256 · 64 × 8 terminal cells

```
API operations  Production / us-east-1        Observed 00:16 UTC
```

16 colors · ansi16 · 64 × 8 terminal cells

```
API operations  Production / us-east-1        Observed 00:16 UTC
```

Monochrome · none · 64 × 8 terminal cells

```
API operations  Production / us-east-1        Observed 00:16 UTC
```

### [Go usage](https://lanej.io/charmed/molecules-layout.html#layout%2ftext-row-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
row := layout.TextRow{Gap: 2, Items: []layout.TextRowItem{
	{Text: palette.Heading().Render("API operations")},
	{Text: palette.Subtle().Render("Production / us-east-1")},
	{Flex: true},
	{Text: palette.Text().Render("Observed 00:16 UTC")},
}}

return row.View(64)
// Use row.Stack = true for an explicitly vertical presentation.
// Measure the returned height when allocating a parent panel or viewport.
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-layout.html#layout%2ftext-row-integration)

Supply trusted native display strings, escaping application data with text.Inline before styling. Gap reserves cells between nonempty items; flexible spacers share remaining width. View wraps at item boundaries and within long items. Stack selects explicit vertical layout. Allocate the returned height or provide scrolling; Fit clips. Rendering is pure, preserves native styling and owns no focus or input handling.

### [Run this example](https://lanej.io/charmed/molecules-layout.html#layout%2ftext-row-run)

```sh
go run ./examples/catalog layout/text-row
```

[Example source](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/docs/snippets/layout_text_row.go) · [Implementation](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/layout/text_row.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [foundations/layout](https://lanej.io/charmed/foundations-layout.html#foundations/layout) [text/primitives](https://lanej.io/charmed/atoms-text.html#text/primitives)

[Molecules](https://lanej.io/charmed/molecules.html) / Layout

## [Link group](https://lanej.io/charmed/molecules-layout.html#layout%2flink-group)

Related destinations flow in order with shared spacing, wrap between items, and retain exact targets through long-label wrapping. Each link keeps its own presentation and emphasis.

### [Native preview](https://lanej.io/charmed/molecules-layout.html#layout%2flink-group-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
↗ Logs  ↗ Artifact
↗ Worker session
```

Demo · truecolor · 64 × 16 terminal cells

```
Build #42                                                Passed 
↗ Logs  ↗ Artifact  ↗ Worker session                            
                                                                
Reflow / 24 cells                                               
↗ Logs  ↗ Artifact                                              
↗ Worker session                                                
                                                                
Dense / label only                                              
Logs  Artifact  Worker session
```

Narrow · truecolor · 40 × 16 terminal cells

```
Build #42                        Passed 
↗ Logs  ↗ Artifact  ↗ Worker session    
                                        
Reflow / 24 cells                       
↗ Logs  ↗ Artifact                      
↗ Worker session                        
                                        
Dense / label only                      
Logs  Artifact  Worker session
```

256 colors · ansi256 · 64 × 16 terminal cells

```
Build #42                                                Passed 
↗ Logs  ↗ Artifact  ↗ Worker session                            
                                                                
Reflow / 24 cells                                               
↗ Logs  ↗ Artifact                                              
↗ Worker session                                                
                                                                
Dense / label only                                              
Logs  Artifact  Worker session
```

16 colors · ansi16 · 64 × 16 terminal cells

```
Build #42                                                Passed 
↗ Logs  ↗ Artifact  ↗ Worker session                            
                                                                
Reflow / 24 cells                                               
↗ Logs  ↗ Artifact                                              
↗ Worker session                                                
                                                                
Dense / label only                                              
Logs  Artifact  Worker session
```

Monochrome · none · 64 × 16 terminal cells

```
Build #42                                               [Passed]
@ Logs  @ Artifact  @ Worker session                            
                                                                
Reflow / 24 cells                                               
@ Logs  @ Artifact                                              
@ Worker session                                                
                                                                
Dense / label only                                              
Logs  Artifact  Worker session
```

light · truecolor · 64 × 16 terminal cells

```
Build #42                                                Passed 
↗ Logs  ↗ Artifact  ↗ Worker session                            
                                                                
Reflow / 24 cells                                               
↗ Logs  ↗ Artifact                                              
↗ Worker session                                                
                                                                
Dense / label only                                              
Logs  Artifact  Worker session
```

contrast · truecolor · 64 × 16 terminal cells

```
Build #42                                                Passed 
↗ Logs  ↗ Artifact  ↗ Worker session                            
                                                                
Reflow / 24 cells                                               
↗ Logs  ↗ Artifact                                              
↗ Worker session                                                
                                                                
Dense / label only                                              
Logs  Artifact  Worker session
```

### [Go usage](https://lanej.io/charmed/molecules-layout.html#layout%2flink-group-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
group := layout.LinkGroup{Links: []text.Link{
	{Label: "Logs", URL: "https://example.test/builds/42/logs"},
	{Label: "Artifact", URL: "https://example.test/builds/42/artifact"},
	{Label: "Worker session", URL: "terminal://session/build/pane/tests", Emphasis: text.Dim},
}}

return group.View(24, theme.New(theme.Dark))
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-layout.html#layout%2flink-group-integration)

The parent owns state, focus and input routing; layout helpers render the values supplied to them. Measure headings, borders and gutters before allocating child rectangles. Re-render from the new terminal-cell dimensions on resize, and reserve the actual rendered height when text may wrap.

### [Run this example](https://lanej.io/charmed/molecules-layout.html#layout%2flink-group-run)

```sh
go run ./examples/catalog layout/link-group
```

[Example source](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/docs/snippets/layout_link_group.go) · [Implementation](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/layout/links.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [text/link](https://lanej.io/charmed/atoms-text.html#text/link) [foundations/layout](https://lanej.io/charmed/foundations-layout.html#foundations/layout)

[Molecules](https://lanej.io/charmed/molecules.html) / Layout

## [Header](https://lanej.io/charmed/molecules-layout.html#layout%2fheader)

Title and trailing actions with an optional context row. Narrow widths preserve the title before actions.

### [Native preview](https://lanej.io/charmed/molecules-layout.html#layout%2fheader-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Service workbench                                      r refresh
Production / us-west-2
```

Demo · truecolor · 64 × 8 terminal cells

```
Service workbench                                      r refresh
Production / us-west-2
```

Narrow · truecolor · 40 × 16 terminal cells

```
Service workbench              r refresh
Production / us-west-2
```

256 colors · ansi256 · 64 × 8 terminal cells

```
Service workbench                                      r refresh
Production / us-west-2
```

16 colors · ansi16 · 64 × 8 terminal cells

```
Service workbench                                      r refresh
Production / us-west-2
```

Monochrome · none · 64 × 8 terminal cells

```
Service workbench                                      r refresh
Production / us-west-2
```

### [Go usage](https://lanej.io/charmed/molecules-layout.html#layout%2fheader-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
header := layout.Header{
	Title:   "Service workbench",
	Context: "Production / us-west-2",
	Actions: text.KeyHint("r", "refresh", palette),
}

return header.View(64, palette)
// Bind r in the application; Header only renders the hint.
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-layout.html#layout%2fheader-integration)

The parent owns state, focus and input routing; layout helpers render the values supplied to them. Measure headings, borders and gutters before allocating child rectangles. Re-render from the new terminal-cell dimensions on resize, and reserve the actual rendered height when text may wrap.

### [Run this example](https://lanej.io/charmed/molecules-layout.html#layout%2fheader-run)

```sh
go run ./examples/catalog layout/header
```

[Example source](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/docs/snippets/layout_header.go) · [Implementation](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/story/chrome.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [layout/text-row](https://lanej.io/charmed/molecules-layout.html#layout/text-row) [foundations/layout](https://lanej.io/charmed/foundations-layout.html#foundations/layout)

[Molecules](https://lanej.io/charmed/molecules.html) / Layout

## [Identity header](https://lanej.io/charmed/molecules-layout.html#layout%2fidentity-header)

A request UUID or other entity identifier is the headline. Narrow terminals move the kind above the complete identifier before wrapping.

### [Native preview](https://lanej.io/charmed/molecules-layout.html#layout%2fidentity-header-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Request 01234567-89ab-cdef-0123-456789abcdef                    
POST /v2/shipments/shp_fixture/buy · HTTP 500
```

Demo · truecolor · 64 × 8 terminal cells

```
Request 01234567-89ab-cdef-0123-456789abcdef                    
POST /v2/shipments/shp_fixture/buy · HTTP 500 · 12.5ms
```

Narrow · truecolor · 40 × 16 terminal cells

```
Request                                 
01234567-89ab-cdef-0123-456789abcdef    
POST /v2/shipments/shp_fixture/buy ·    
HTTP 500 · 12.5ms
```

256 colors · ansi256 · 64 × 8 terminal cells

```
Request 01234567-89ab-cdef-0123-456789abcdef                    
POST /v2/shipments/shp_fixture/buy · HTTP 500 · 12.5ms
```

16 colors · ansi16 · 64 × 8 terminal cells

```
Request 01234567-89ab-cdef-0123-456789abcdef                    
POST /v2/shipments/shp_fixture/buy · HTTP 500 · 12.5ms
```

Monochrome · none · 64 × 8 terminal cells

```
Request 01234567-89ab-cdef-0123-456789abcdef                    
POST /v2/shipments/shp_fixture/buy * HTTP 500 * 12.5ms
```

### [Go usage](https://lanej.io/charmed/molecules-layout.html#layout%2fidentity-header-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
headline := layout.IdentityHeader{
	Kind: "Request", ID: "01234567-89ab-cdef-0123-456789abcdef",
	Context: "POST /v2/shipments/shp_fixture/buy · HTTP 500",
}

return headline.View(64, theme.New(theme.Dark))
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-layout.html#layout%2fidentity-header-integration)

The parent owns state, focus and input routing; layout helpers render the values supplied to them. Measure headings, borders and gutters before allocating child rectangles. Re-render from the new terminal-cell dimensions on resize, and reserve the actual rendered height when text may wrap.

### [Run this example](https://lanej.io/charmed/molecules-layout.html#layout%2fidentity-header-run)

```sh
go run ./examples/catalog layout/identity-header
```

[Example source](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/docs/snippets/layout_identity_header.go) · [Implementation](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/story/chrome.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [layout/text-row](https://lanej.io/charmed/molecules-layout.html#layout/text-row) [foundations/layout](https://lanej.io/charmed/foundations-layout.html#foundations/layout)

[Molecules](https://lanej.io/charmed/molecules.html) / Layout

## [Foldable section](https://lanej.io/charmed/molecules-layout.html#layout%2fsection)

A panel folds to one heading row, returning space to sibling content. The application owns state and bindings; press Enter to toggle this specimen.

### [Native preview](https://lanej.io/charmed/molecules-layout.html#layout%2fsection-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
┌ ▾ Request timeline ──────────────────────────────────────────┐
│ POST /v2/shipments/shp_fixture/buy   500   12.5ms            │
└──────────────────────────────────────────────────────────────┘
┌ Response ────────────────────────────────────────────────────┐
│ Internal server error                                        │
│ Could not purchase shipment                                  │
│                                                              │
└──────────────────────────────────────────────────────────────┘
```

Demo · truecolor · 64 × 8 terminal cells

```
┌ ▾ Request timeline ──────────────────────────────────────────┐
│ POST /v2/shipments/shp_fixture/buy   500   12.5ms            │
└──────────────────────────────────────────────────────────────┘
┌ Response ────────────────────────────────────────────────────┐
│ Internal server error                                        │
│ Could not purchase shipment                                  │
│                                                              │
└──────────────────────────────────────────────────────────────┘
```

Narrow · truecolor · 40 × 16 terminal cells

```
┌ ▾ Request timeline ──────────────────┐
│ POST /v2/shipments/shp_fixture/buy   │
└──────────────────────────────────────┘
┌ Response ────────────────────────────┐
│ Internal server error                │
│ Could not purchase shipment          │
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

256 colors · ansi256 · 64 × 8 terminal cells

```
┌ ▾ Request timeline ──────────────────────────────────────────┐
│ POST /v2/shipments/shp_fixture/buy   500   12.5ms            │
└──────────────────────────────────────────────────────────────┘
┌ Response ────────────────────────────────────────────────────┐
│ Internal server error                                        │
│ Could not purchase shipment                                  │
│                                                              │
└──────────────────────────────────────────────────────────────┘
```

16 colors · ansi16 · 64 × 8 terminal cells

```
┌ ▾ Request timeline ──────────────────────────────────────────┐
│ POST /v2/shipments/shp_fixture/buy   500   12.5ms            │
└──────────────────────────────────────────────────────────────┘
┌ Response ────────────────────────────────────────────────────┐
│ Internal server error                                        │
│ Could not purchase shipment                                  │
│                                                              │
└──────────────────────────────────────────────────────────────┘
```

Monochrome · none · 64 × 8 terminal cells

```
+ v Request timeline ------------------------------------------+
| POST /v2/shipments/shp_fixture/buy   500   12.5ms            |
+--------------------------------------------------------------+
+ Response ----------------------------------------------------+
| Internal server error                                        |
| Could not purchase shipment                                  |
|                                                              |
+--------------------------------------------------------------+
```

### [Go usage](https://lanej.io/charmed/molecules-layout.html#layout%2fsection-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
section := layout.Section{
	Title: "Request timeline", Content: "POST /v2/shipments/shp_fixture/buy   500   12.5ms",
	Collapsed: false, // Set true to return the section's space to the response.
}
height := section.Height(3)
timeline := section.View(layout.Size{Width: 64, Height: height}, palette, true)
response := layout.Panel("Response", "Internal server error\nCould not purchase shipment",
	layout.Size{Width: 64, Height: 8 - height}, palette, false)

return layout.Column(timeline, response)
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-layout.html#layout%2fsection-integration)

The parent owns state, focus and input routing; layout helpers render the values supplied to them. Measure headings, borders and gutters before allocating child rectangles. Re-render from the new terminal-cell dimensions on resize, and reserve the actual rendered height when text may wrap.

### [Run this example](https://lanej.io/charmed/molecules-layout.html#layout%2fsection-run)

```sh
go run ./examples/catalog layout/section
```

[Example source](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/docs/snippets/layout_section.go) · [Implementation](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/story/chrome.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [foundations/layout](https://lanej.io/charmed/foundations-layout.html#foundations/layout) [text/disclosure](https://lanej.io/charmed/atoms-text.html#text/disclosure)

[Molecules](https://lanej.io/charmed/molecules.html) / Layout

## [Key-value](https://lanej.io/charmed/molecules-layout.html#layout%2fkey-value)

Aligned labels and values; 55 columns and narrower stack each pair. Missing values stay distinct from empty strings.

### [Native preview](https://lanej.io/charmed/molecules-layout.html#layout%2fkey-value-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
profile  dev
service  v1.42.0 (contract 7)
compat   ∅ missing
```

Demo · truecolor · 64 × 8 terminal cells

```
At the allocation                                               
profile  dev                                                    
service  v1.42.0 (contract 7)                                   
note     ""                                                     
compat   ∅ missing                                              
detail   a long value wraps within the width left after the labe
         l column, never past it
```

Narrow · truecolor · 40 × 16 terminal cells

```
At the allocation                       
profile: dev                            
service: v1.42.0 (contract 7)           
note: ""                                
compat: ∅ missing                       
detail: a long value wraps within the wi
dth left after the label column,        
 never past it                          
                                        
At 50 columns                           
profile: dev                            
service: v1.42.0 (contract 7)           
note: ""                                
compat: ∅ missing                       
detail: a long value wraps within the wi
fter the label column, never past it
```

256 colors · ansi256 · 64 × 8 terminal cells

```
At the allocation                                               
profile  dev                                                    
service  v1.42.0 (contract 7)                                   
note     ""                                                     
compat   ∅ missing                                              
detail   a long value wraps within the width left after the labe
         l column, never past it
```

16 colors · ansi16 · 64 × 8 terminal cells

```
At the allocation                                               
profile  dev                                                    
service  v1.42.0 (contract 7)                                   
note     ""                                                     
compat   ∅ missing                                              
detail   a long value wraps within the width left after the labe
         l column, never past it
```

Monochrome · none · 64 × 8 terminal cells

```
At the allocation                                               
profile  dev                                                    
service  v1.42.0 (contract 7)                                   
note     ""                                                     
compat   missing                                                
detail   a long value wraps within the width left after the labe
         l column, never past it
```

### [Go usage](https://lanej.io/charmed/molecules-layout.html#layout%2fkey-value-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
pairs := []layout.Pair{
	{Label: "profile", Value: "dev"},
	{Label: "service", Value: "v1.42.0 (contract 7)"},
	{Label: "compat", State: text.ValueMissing}, // Missing is not an empty string.
}

return layout.KeyValue(pairs, 80, theme.New(theme.Dark))
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-layout.html#layout%2fkey-value-integration)

The parent owns state, focus and input routing; layout helpers render the values supplied to them. Measure headings, borders and gutters before allocating child rectangles. Re-render from the new terminal-cell dimensions on resize, and reserve the actual rendered height when text may wrap.

### [Run this example](https://lanej.io/charmed/molecules-layout.html#layout%2fkey-value-run)

```sh
go run ./examples/catalog layout/key-value
```

[Example source](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/docs/snippets/layout_key_value.go) · [Implementation](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/layout/keyvalue.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [text/value-state](https://lanej.io/charmed/atoms-text.html#text/value-state) [foundations/layout](https://lanej.io/charmed/foundations-layout.html#foundations/layout)

[Molecules](https://lanej.io/charmed/molecules.html) / Layout

## [Footer](https://lanej.io/charmed/molecules-layout.html#layout%2ffooter)

Keyboard help and right-aligned status share one row. Narrow widths preserve help before status.

### [Native preview](https://lanej.io/charmed/molecules-layout.html#layout%2ffooter-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
? help  q quit                                      ● CONNECTED 
```

Demo · truecolor · 64 × 8 terminal cells

```
? help  q quit                                      ● CONNECTED 
```

Narrow · truecolor · 40 × 16 terminal cells

```
? help  q quit              ● CONNECTED 
```

256 colors · ansi256 · 64 × 8 terminal cells

```
? help  q quit                                      ● CONNECTED 
```

16 colors · ansi16 · 64 × 8 terminal cells

```
? help  q quit                                      ● CONNECTED 
```

Monochrome · none · 64 × 8 terminal cells

```
? help  q quit                                     [+ CONNECTED]
```

### [Go usage](https://lanej.io/charmed/molecules-layout.html#layout%2ffooter-code)

Imports: [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
footer := layout.Footer{
	Help: text.KeyHint("?", "help", palette) + "  " +
		text.KeyHint("q", "quit", palette),
	Status: feedback.Connected.View(text.Normal, palette),
}

return footer.View(64, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-layout.html#layout%2ffooter-integration)

The parent owns state, focus and input routing; layout helpers render the values supplied to them. Measure headings, borders and gutters before allocating child rectangles. Re-render from the new terminal-cell dimensions on resize, and reserve the actual rendered height when text may wrap.

### [Run this example](https://lanej.io/charmed/molecules-layout.html#layout%2ffooter-run)

```sh
go run ./examples/catalog layout/footer
```

[Example source](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/docs/snippets/layout_footer.go) · [Implementation](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/story/chrome.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [layout/text-row](https://lanej.io/charmed/molecules-layout.html#layout/text-row) [foundations/layout](https://lanej.io/charmed/foundations-layout.html#foundations/layout) [text/primitives](https://lanej.io/charmed/atoms-text.html#text/primitives) [feedback/status](https://lanej.io/charmed/atoms-feedback.html#feedback/status)

[Molecules](https://lanej.io/charmed/molecules.html) / Layout

## [Weighted flex layout](https://lanej.io/charmed/molecules-layout.html#layout%2fflex)

Share terminal cells between weighted regions while preserving minimums, maximums, gaps and alignment.

### [Native preview](https://lanej.io/charmed/molecules-layout.html#layout%2fflex-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
┌ main ───────────────────────────────────────────────────────┐ ┌ tools ───────┐
│ Native view                                                 │ │ Native view  │
│                                                             │ │              │
│                                                             │ │              │
│                                                             │ │              │
│                                                             │ │              │
│                                                             │ │              │
│                                                             │ │              │
│                                                             │ │              │
│                                                             │ │              │
│                                                             │ │              │
│                                                             │ │              │
│                                                             │ │              │
│                                                             │ │              │
│                                                             │ │              │
└─────────────────────────────────────────────────────────────┘ └──────────────┘
```

Demo · truecolor · 100 × 20 terminal cells

```
┌ primary ───────────────────────────────────────────┐ ┌ secondary ─────────────────────┐ ┌ tools ─┐
│ 54×20                                              │ │ 34×20                          │ │ 10×20  │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
└────────────────────────────────────────────────────┘ └────────────────────────────────┘ └────────┘
```

Narrow · truecolor · 40 × 16 terminal cells

```
┌ primary ───┐ ┌ secondary ─┐ ┌ tools ─┐
│ 14×16      │ │ 14×16      │ │ 10×16  │
│            │ │            │ │        │
│            │ │            │ │        │
│            │ │            │ │        │
│            │ │            │ │        │
│            │ │            │ │        │
│            │ │            │ │        │
│            │ │            │ │        │
│            │ │            │ │        │
│            │ │            │ │        │
│            │ │            │ │        │
│            │ │            │ │        │
│            │ │            │ │        │
│            │ │            │ │        │
└────────────┘ └────────────┘ └────────┘
```

256 colors · ansi256 · 100 × 20 terminal cells

```
┌ primary ───────────────────────────────────────────┐ ┌ secondary ─────────────────────┐ ┌ tools ─┐
│ 54×20                                              │ │ 34×20                          │ │ 10×20  │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
└────────────────────────────────────────────────────┘ └────────────────────────────────┘ └────────┘
```

16 colors · ansi16 · 100 × 20 terminal cells

```
┌ primary ───────────────────────────────────────────┐ ┌ secondary ─────────────────────┐ ┌ tools ─┐
│ 54×20                                              │ │ 34×20                          │ │ 10×20  │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
│                                                    │ │                                │ │        │
└────────────────────────────────────────────────────┘ └────────────────────────────────┘ └────────┘
```

Monochrome · none · 100 × 20 terminal cells

```
+ primary -------------------------------------------+ + secondary ---------------------+ + tools -+
| 54x20                                              | | 34x20                          | | 10x20  |
|                                                    | |                                | |        |
|                                                    | |                                | |        |
|                                                    | |                                | |        |
|                                                    | |                                | |        |
|                                                    | |                                | |        |
|                                                    | |                                | |        |
|                                                    | |                                | |        |
|                                                    | |                                | |        |
|                                                    | |                                | |        |
|                                                    | |                                | |        |
|                                                    | |                                | |        |
|                                                    | |                                | |        |
|                                                    | |                                | |        |
|                                                    | |                                | |        |
|                                                    | |                                | |        |
|                                                    | |                                | |        |
|                                                    | |                                | |        |
+----------------------------------------------------+ +--------------------------------+ +--------+
```

### [Go usage](https://lanej.io/charmed/molecules-layout.html#layout%2fflex-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme") · [`tui`](https://pkg.go.dev/github.com/lanej/charmed/tui "github.com/lanej/charmed/tui")

```go
screen, err := tui.NewScreen("main", "tools")
if err != nil {
	return "", err
}

root := &layout.Flex{Gap: 1, Children: []layout.Slot{
	{Node: &layout.Region{ID: "main", Min: layout.Size{Width: 20}}, Track: layout.Grow(20)},
	{Node: &layout.Region{ID: "tools", Min: layout.Size{Width: 12}}, Track: layout.Fixed(16)},
}}
if arrangeErr := screen.Arrange(root, layout.Size{Width: 80, Height: 16}); arrangeErr != nil {
	return "", arrangeErr
}

views := map[string]string{}

for _, region := range screen.Plan().Regions() {
	size := layout.Size{Width: region.Bounds.Dx(), Height: region.Bounds.Dy()}
	views[region.ID] = layout.Panel(region.ID, "Native view", size, theme.New(theme.Dark), false)
}

frame, err := screen.Render(views)

return frame.Content, err
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-layout.html#layout%2fflex-integration)

Declare region identities and hard minimum sizes, then arrange a registered tui.Screen at each terminal resize. Resize native child models from the resulting rectangles and render exactly the visible region views. Geometry and design rules block by default; exceptions name a specific rule and reason on its region or container. Hidden components retain caller-owned state, and CheckLayouts exercises every responsive variant.

### [Run this example](https://lanej.io/charmed/molecules-layout.html#layout%2fflex-run)

```sh
go run ./examples/catalog layout/flex
```

[Example source](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/docs/snippets/layout_flex.go) · [Implementation](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/story/layouts.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [foundations/layout](https://lanej.io/charmed/foundations-layout.html#foundations/layout)

[Molecules](https://lanej.io/charmed/molecules.html) / Layout

## [Grid with spanning cells](https://lanej.io/charmed/molecules-layout.html#layout%2fgrid)

Align regions across shared tracks and span cells without overlapping neighbors.

### [Native preview](https://lanej.io/charmed/molecules-layout.html#layout%2fgrid-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
┌ main ────────────────────────────────────┐ ┌ tools ──────────────────────────┐
│ Native view                              │ │ Native view                     │
│                                          │ │                                 │
│                                          │ │                                 │
│                                          │ │                                 │
│                                          │ │                                 │
│                                          │ │                                 │
└──────────────────────────────────────────┘ └─────────────────────────────────┘
                                                                                
┌ activity ────────────────────────────────────────────────────────────────────┐
│ Native view                                                                  │
│                                                                              │
│                                                                              │
│                                                                              │
│                                                                              │
└──────────────────────────────────────────────────────────────────────────────┘
```

Demo · truecolor · 100 × 20 terminal cells

```
┌ primary ──────────────────────────────────────┐ ┌ secondary ─────────────────────────────────────┐
│ 49×10                                         │ │ 50×10                                          │
│                                               │ │                                                │
│                                               │ │                                                │
│                                               │ │                                                │
│                                               │ │                                                │
│                                               │ │                                                │
│                                               │ │                                                │
│                                               │ │                                                │
└───────────────────────────────────────────────┘ └────────────────────────────────────────────────┘
                                                                                                    
┌ tools ───────────────────────────────────────────────────────────────────────────────────────────┐
│ 100×9                                                                                            │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
```

Narrow · truecolor · 40 × 16 terminal cells

```
┌ primary ────────┐ ┌ secondary ───────┐
│ 19×8            │ │ 20×8             │
│                 │ │                  │
│                 │ │                  │
│                 │ │                  │
│                 │ │                  │
│                 │ │                  │
└─────────────────┘ └──────────────────┘
                                        
┌ tools ───────────────────────────────┐
│ 40×7                                 │
│                                      │
│                                      │
│                                      │
│                                      │
└──────────────────────────────────────┘
```

256 colors · ansi256 · 100 × 20 terminal cells

```
┌ primary ──────────────────────────────────────┐ ┌ secondary ─────────────────────────────────────┐
│ 49×10                                         │ │ 50×10                                          │
│                                               │ │                                                │
│                                               │ │                                                │
│                                               │ │                                                │
│                                               │ │                                                │
│                                               │ │                                                │
│                                               │ │                                                │
│                                               │ │                                                │
└───────────────────────────────────────────────┘ └────────────────────────────────────────────────┘
                                                                                                    
┌ tools ───────────────────────────────────────────────────────────────────────────────────────────┐
│ 100×9                                                                                            │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
```

16 colors · ansi16 · 100 × 20 terminal cells

```
┌ primary ──────────────────────────────────────┐ ┌ secondary ─────────────────────────────────────┐
│ 49×10                                         │ │ 50×10                                          │
│                                               │ │                                                │
│                                               │ │                                                │
│                                               │ │                                                │
│                                               │ │                                                │
│                                               │ │                                                │
│                                               │ │                                                │
│                                               │ │                                                │
└───────────────────────────────────────────────┘ └────────────────────────────────────────────────┘
                                                                                                    
┌ tools ───────────────────────────────────────────────────────────────────────────────────────────┐
│ 100×9                                                                                            │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
```

Monochrome · none · 100 × 20 terminal cells

```
+ primary --------------------------------------+ + secondary -------------------------------------+
| 49x10                                         | | 50x10                                          |
|                                               | |                                                |
|                                               | |                                                |
|                                               | |                                                |
|                                               | |                                                |
|                                               | |                                                |
|                                               | |                                                |
|                                               | |                                                |
+-----------------------------------------------+ +------------------------------------------------+
                                                                                                    
+ tools -------------------------------------------------------------------------------------------+
| 100x9                                                                                            |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
+--------------------------------------------------------------------------------------------------+
```

### [Go usage](https://lanej.io/charmed/molecules-layout.html#layout%2fgrid-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme") · [`tui`](https://pkg.go.dev/github.com/lanej/charmed/tui "github.com/lanej/charmed/tui")

```go
screen, err := tui.NewScreen("main", "tools", "activity")
if err != nil {
	return "", err
}

root := &layout.Grid{
	Columns: []layout.Track{layout.Grow(20), layout.Grow(12)}, Rows: []layout.Track{layout.Grow(3), layout.Grow(3)}, Gap: layout.Size{Width: 1, Height: 1},
	Cells: []layout.Cell{{Node: &layout.Region{ID: "main"}}, {Node: &layout.Region{ID: "tools"}, Column: 1}, {Node: &layout.Region{ID: "activity"}, Row: 1, ColumnSpan: 2}},
}
if arrangeErr := screen.Arrange(root, layout.Size{Width: 80, Height: 16}); arrangeErr != nil {
	return "", arrangeErr
}

views := map[string]string{}

for _, region := range screen.Plan().Regions() {
	size := layout.Size{Width: region.Bounds.Dx(), Height: region.Bounds.Dy()}
	views[region.ID] = layout.Panel(region.ID, "Native view", size, theme.New(theme.Dark), false)
}

frame, err := screen.Render(views)

return frame.Content, err
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-layout.html#layout%2fgrid-integration)

Declare region identities and hard minimum sizes, then arrange a registered tui.Screen at each terminal resize. Resize native child models from the resulting rectangles and render exactly the visible region views. Geometry and design rules block by default; exceptions name a specific rule and reason on its region or container. Hidden components retain caller-owned state, and CheckLayouts exercises every responsive variant.

### [Run this example](https://lanej.io/charmed/molecules-layout.html#layout%2fgrid-run)

```sh
go run ./examples/catalog layout/grid
```

[Example source](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/docs/snippets/layout_grid.go) · [Implementation](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/story/layouts.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [foundations/layout](https://lanej.io/charmed/foundations-layout.html#foundations/layout)

[Molecules](https://lanej.io/charmed/molecules.html) / Layout

## [Wrapping flow layout](https://lanej.io/charmed/molecules-layout.html#layout%2fflow)

Wrap complete action or summary regions onto new rows when horizontal space runs out.

### [Native preview](https://lanej.io/charmed/molecules-layout.html#layout%2fflow-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
┌ build ─────────────────────┐ ┌ test ──────────────────────┐                   
│ Native view                │ │ Native view                │                   
└────────────────────────────┘ └────────────────────────────┘                   
                                                                                
┌ publish ───────────────────┐                                                  
│ Native view                │                                                  
└────────────────────────────┘
```

Demo · truecolor · 100 × 20 terminal cells

```
┌ primary ──────────────────────┐ ┌ secondary ────────────────────┐                                 
│ 33×3                          │ │ 33×3                          │                                 
└───────────────────────────────┘ └───────────────────────────────┘                                 
                                                                                                    
┌ tools ────────────────────────┐                                                                   
│ 33×3                          │                                                                   
└───────────────────────────────┘
```

Narrow · truecolor · 40 × 16 terminal cells

```
┌ primary ───┐ ┌ secondary ─┐           
│ 14×3       │ │ 14×3       │           
└────────────┘ └────────────┘           
                                        
┌ tools ─────┐                          
│ 14×3       │                          
└────────────┘
```

256 colors · ansi256 · 100 × 20 terminal cells

```
┌ primary ──────────────────────┐ ┌ secondary ────────────────────┐                                 
│ 33×3                          │ │ 33×3                          │                                 
└───────────────────────────────┘ └───────────────────────────────┘                                 
                                                                                                    
┌ tools ────────────────────────┐                                                                   
│ 33×3                          │                                                                   
└───────────────────────────────┘
```

16 colors · ansi16 · 100 × 20 terminal cells

```
┌ primary ──────────────────────┐ ┌ secondary ────────────────────┐                                 
│ 33×3                          │ │ 33×3                          │                                 
└───────────────────────────────┘ └───────────────────────────────┘                                 
                                                                                                    
┌ tools ────────────────────────┐                                                                   
│ 33×3                          │                                                                   
└───────────────────────────────┘
```

Monochrome · none · 100 × 20 terminal cells

```
+ primary ----------------------+ + secondary --------------------+                                 
| 33x3                          | | 33x3                          |                                 
+-------------------------------+ +-------------------------------+                                 
                                                                                                    
+ tools ------------------------+                                                                   
| 33x3                          |                                                                   
+-------------------------------+
```

### [Go usage](https://lanej.io/charmed/molecules-layout.html#layout%2fflow-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme") · [`tui`](https://pkg.go.dev/github.com/lanej/charmed/tui "github.com/lanej/charmed/tui")

```go
screen, err := tui.NewScreen("build", "test", "publish")
if err != nil {
	return "", err
}

root := &layout.Flow{Gap: layout.Size{Width: 1, Height: 1}, Children: []layout.Box{
	{Node: &layout.Region{ID: "build"}, Size: layout.Size{Width: 30, Height: 3}},
	{Node: &layout.Region{ID: "test"}, Size: layout.Size{Width: 30, Height: 3}},
	{Node: &layout.Region{ID: "publish"}, Size: layout.Size{Width: 30, Height: 3}},
}}
if arrangeErr := screen.Arrange(root, layout.Size{Width: 80, Height: 16}); arrangeErr != nil {
	return "", arrangeErr
}

views := map[string]string{}

for _, region := range screen.Plan().Regions() {
	size := layout.Size{Width: region.Bounds.Dx(), Height: region.Bounds.Dy()}
	views[region.ID] = layout.Panel(region.ID, "Native view", size, theme.New(theme.Dark), false)
}

frame, err := screen.Render(views)

return frame.Content, err
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-layout.html#layout%2fflow-integration)

Declare region identities and hard minimum sizes, then arrange a registered tui.Screen at each terminal resize. Resize native child models from the resulting rectangles and render exactly the visible region views. Geometry and design rules block by default; exceptions name a specific rule and reason on its region or container. Hidden components retain caller-owned state, and CheckLayouts exercises every responsive variant.

### [Run this example](https://lanej.io/charmed/molecules-layout.html#layout%2fflow-run)

```sh
go run ./examples/catalog layout/flow
```

[Example source](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/docs/snippets/layout_flow.go) · [Implementation](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/story/layouts.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [foundations/layout](https://lanej.io/charmed/foundations-layout.html#foundations/layout)

[Molecules](https://lanej.io/charmed/molecules.html) / Layout

## [Docked layout](https://lanej.io/charmed/molecules-layout.html#layout%2fdock)

Reserve header, footer and side regions around an expanding primary workspace.

### [Native preview](https://lanej.io/charmed/molecules-layout.html#layout%2fdock-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Native view                                                                     
                                                                                
┌ tools ───────┐ ┌ main ───────────────────────────────────────────────────────┐
│ Native view  │ │ Native view                                                 │
│              │ │                                                             │
│              │ │                                                             │
│              │ │                                                             │
│              │ │                                                             │
│              │ │                                                             │
│              │ │                                                             │
│              │ │                                                             │
│              │ │                                                             │
│              │ │                                                             │
└──────────────┘ └─────────────────────────────────────────────────────────────┘
                                                                                
Native view
```

Demo · truecolor · 100 × 20 terminal cells

```
┌ tools ───────────────────────────────────────────────────────────────────────────────────────────┐
│ 100×3                                                                                            │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
                                                                                                    
┌ secondary ───┐ ┌ primary ────────────────────────────────────────────────────────────────────────┐
│ 16×16        │ │ 83×16                                                                           │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
└──────────────┘ └─────────────────────────────────────────────────────────────────────────────────┘
```

Narrow · truecolor · 40 × 16 terminal cells

```
┌ tools ───────────────────────────────┐
│ 40×3                                 │
└──────────────────────────────────────┘
                                        
┌ secondary ───┐ ┌ primary ────────────┐
│ 16×12        │ │ 23×12               │
│              │ │                     │
│              │ │                     │
│              │ │                     │
│              │ │                     │
│              │ │                     │
│              │ │                     │
│              │ │                     │
│              │ │                     │
│              │ │                     │
└──────────────┘ └─────────────────────┘
```

256 colors · ansi256 · 100 × 20 terminal cells

```
┌ tools ───────────────────────────────────────────────────────────────────────────────────────────┐
│ 100×3                                                                                            │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
                                                                                                    
┌ secondary ───┐ ┌ primary ────────────────────────────────────────────────────────────────────────┐
│ 16×16        │ │ 83×16                                                                           │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
└──────────────┘ └─────────────────────────────────────────────────────────────────────────────────┘
```

16 colors · ansi16 · 100 × 20 terminal cells

```
┌ tools ───────────────────────────────────────────────────────────────────────────────────────────┐
│ 100×3                                                                                            │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
                                                                                                    
┌ secondary ───┐ ┌ primary ────────────────────────────────────────────────────────────────────────┐
│ 16×16        │ │ 83×16                                                                           │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
│              │ │                                                                                 │
└──────────────┘ └─────────────────────────────────────────────────────────────────────────────────┘
```

Monochrome · none · 100 × 20 terminal cells

```
+ tools -------------------------------------------------------------------------------------------+
| 100x3                                                                                            |
+--------------------------------------------------------------------------------------------------+
                                                                                                    
+ secondary ---+ + primary ------------------------------------------------------------------------+
| 16x16        | | 83x16                                                                           |
|              | |                                                                                 |
|              | |                                                                                 |
|              | |                                                                                 |
|              | |                                                                                 |
|              | |                                                                                 |
|              | |                                                                                 |
|              | |                                                                                 |
|              | |                                                                                 |
|              | |                                                                                 |
|              | |                                                                                 |
|              | |                                                                                 |
|              | |                                                                                 |
|              | |                                                                                 |
+--------------+ +---------------------------------------------------------------------------------+
```

### [Go usage](https://lanej.io/charmed/molecules-layout.html#layout%2fdock-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme") · [`tui`](https://pkg.go.dev/github.com/lanej/charmed/tui "github.com/lanej/charmed/tui")

```go
screen, err := tui.NewScreen("header", "main", "tools", "footer")
if err != nil {
	return "", err
}

root := &layout.Dock{
	Top: &layout.Region{ID: "header", Role: layout.ChromeRegion}, TopHeight: 1,
	Left: &layout.Region{ID: "tools", Role: layout.SupportingRegion}, LeftWidth: 16,
	Body:   &layout.Region{ID: "main", Min: layout.Size{Width: 20, Height: 3}},
	Bottom: &layout.Region{ID: "footer", Role: layout.ChromeRegion}, BottomHeight: 1, Gap: 1,
}
if arrangeErr := screen.Arrange(root, layout.Size{Width: 80, Height: 16}); arrangeErr != nil {
	return "", arrangeErr
}

views := map[string]string{}

for _, region := range screen.Plan().Regions() {
	size := layout.Size{Width: region.Bounds.Dx(), Height: region.Bounds.Dy()}
	views[region.ID] = layout.Panel(region.ID, "Native view", size, theme.New(theme.Dark), false)
}

frame, err := screen.Render(views)

return frame.Content, err
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-layout.html#layout%2fdock-integration)

Declare region identities and hard minimum sizes, then arrange a registered tui.Screen at each terminal resize. Resize native child models from the resulting rectangles and render exactly the visible region views. Geometry and design rules block by default; exceptions name a specific rule and reason on its region or container. Hidden components retain caller-owned state, and CheckLayouts exercises every responsive variant.

### [Run this example](https://lanej.io/charmed/molecules-layout.html#layout%2fdock-run)

```sh
go run ./examples/catalog layout/dock
```

[Example source](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/docs/snippets/layout_dock.go) · [Implementation](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/story/layouts.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [layout/flex](https://lanej.io/charmed/molecules-layout.html#layout/flex)

[Molecules](https://lanej.io/charmed/molecules.html) / Layout

## [Constrained split panes](https://lanej.io/charmed/molecules-layout.html#layout%2fsplit-pane)

Divide primary and detail panes with a stable ratio and hard minimum sizes.

### [Native preview](https://lanej.io/charmed/molecules-layout.html#layout%2fsplit-pane-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
┌ requests ───────────────────┐ ┌ details ─────────────────────────────────────┐
│ Native view                 │ │ Native view                                  │
│                             │ │                                              │
│                             │ │                                              │
│                             │ │                                              │
│                             │ │                                              │
│                             │ │                                              │
│                             │ │                                              │
│                             │ │                                              │
│                             │ │                                              │
│                             │ │                                              │
│                             │ │                                              │
│                             │ │                                              │
│                             │ │                                              │
│                             │ │                                              │
└─────────────────────────────┘ └──────────────────────────────────────────────┘
```

Demo · truecolor · 100 × 20 terminal cells

```
┌ primary ─────────────────────────────────────────────────────┐ ┌ secondary ──────────────────────┐
│ 64×20                                                        │ │ 35×20                           │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
└──────────────────────────────────────────────────────────────┘ └─────────────────────────────────┘
```

Narrow · truecolor · 40 × 16 terminal cells

```
┌ primary ──────────────┐ ┌ secondary ─┐
│ 25×16                 │ │ 14×16      │
│                       │ │            │
│                       │ │            │
│                       │ │            │
│                       │ │            │
│                       │ │            │
│                       │ │            │
│                       │ │            │
│                       │ │            │
│                       │ │            │
│                       │ │            │
│                       │ │            │
│                       │ │            │
│                       │ │            │
└───────────────────────┘ └────────────┘
```

256 colors · ansi256 · 100 × 20 terminal cells

```
┌ primary ─────────────────────────────────────────────────────┐ ┌ secondary ──────────────────────┐
│ 64×20                                                        │ │ 35×20                           │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
└──────────────────────────────────────────────────────────────┘ └─────────────────────────────────┘
```

16 colors · ansi16 · 100 × 20 terminal cells

```
┌ primary ─────────────────────────────────────────────────────┐ ┌ secondary ──────────────────────┐
│ 64×20                                                        │ │ 35×20                           │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
│                                                              │ │                                 │
└──────────────────────────────────────────────────────────────┘ └─────────────────────────────────┘
```

Monochrome · none · 100 × 20 terminal cells

```
+ primary -----------------------------------------------------+ + secondary ----------------------+
| 64x20                                                        | | 35x20                           |
|                                                              | |                                 |
|                                                              | |                                 |
|                                                              | |                                 |
|                                                              | |                                 |
|                                                              | |                                 |
|                                                              | |                                 |
|                                                              | |                                 |
|                                                              | |                                 |
|                                                              | |                                 |
|                                                              | |                                 |
|                                                              | |                                 |
|                                                              | |                                 |
|                                                              | |                                 |
|                                                              | |                                 |
|                                                              | |                                 |
|                                                              | |                                 |
|                                                              | |                                 |
+--------------------------------------------------------------+ +---------------------------------+
```

### [Go usage](https://lanej.io/charmed/molecules-layout.html#layout%2fsplit-pane-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme") · [`tui`](https://pkg.go.dev/github.com/lanej/charmed/tui "github.com/lanej/charmed/tui")

```go
screen, err := tui.NewScreen("requests", "details")
if err != nil {
	return "", err
}

root := &layout.SplitPane{
	First:  &layout.Region{ID: "requests", Min: layout.Size{Width: 20}},
	Second: &layout.Region{ID: "details", Min: layout.Size{Width: 30}}, Ratio: .4, Gap: 1,
}
if arrangeErr := screen.Arrange(root, layout.Size{Width: 80, Height: 16}); arrangeErr != nil {
	return "", arrangeErr
}

views := map[string]string{}

for _, region := range screen.Plan().Regions() {
	size := layout.Size{Width: region.Bounds.Dx(), Height: region.Bounds.Dy()}
	views[region.ID] = layout.Panel(region.ID, "Native view", size, theme.New(theme.Dark), false)
}

frame, err := screen.Render(views)

return frame.Content, err
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-layout.html#layout%2fsplit-pane-integration)

Declare region identities and hard minimum sizes, then arrange a registered tui.Screen at each terminal resize. Resize native child models from the resulting rectangles and render exactly the visible region views. Geometry and design rules block by default; exceptions name a specific rule and reason on its region or container. Hidden components retain caller-owned state, and CheckLayouts exercises every responsive variant.

### [Run this example](https://lanej.io/charmed/molecules-layout.html#layout%2fsplit-pane-run)

```sh
go run ./examples/catalog layout/split-pane
```

[Example source](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/docs/snippets/layout_split_pane.go) · [Implementation](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/story/layouts.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [layout/flex](https://lanej.io/charmed/molecules-layout.html#layout/flex)

[Molecules](https://lanej.io/charmed/molecules.html) / Layout

## [Responsive layout recipes](https://lanej.io/charmed/molecules-layout.html#layout%2fresponsive)

Select an explicit width-and-height recipe, recording why secondary regions are hidden.

### [Native preview](https://lanej.io/charmed/molecules-layout.html#layout%2fresponsive-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
┌ main ───────────────────────────────┐ ┌ details ─────────────────────────────┐
│ Native view                         │ │ Native view                          │
│                                     │ │                                      │
│                                     │ │                                      │
│                                     │ │                                      │
│                                     │ │                                      │
│                                     │ │                                      │
│                                     │ │                                      │
│                                     │ │                                      │
│                                     │ │                                      │
│                                     │ │                                      │
│                                     │ │                                      │
│                                     │ │                                      │
│                                     │ │                                      │
│                                     │ │                                      │
└─────────────────────────────────────┘ └──────────────────────────────────────┘
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

### [Go usage](https://lanej.io/charmed/molecules-layout.html#layout%2fresponsive-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme") · [`tui`](https://pkg.go.dev/github.com/lanej/charmed/tui "github.com/lanej/charmed/tui")

```go
screen, err := tui.NewScreen("main", "details")
if err != nil {
	return "", err
}

root := &layout.Responsive{Name: "workspace", Variants: []layout.Variant{
	{Name: "wide", Min: layout.Size{Width: 70, Height: 12}, Root: &layout.SplitPane{First: &layout.Region{ID: "main"}, Second: &layout.Region{ID: "details"}, Gap: 1}},
	{Name: "focused", Root: &layout.Region{ID: "main", Exceptions: []layout.Exception{{Rule: layout.SingleRegion, Reason: "Single focused pane; parent selection reveals details."}}}, Hidden: []layout.Omission{{ID: "details", Reason: "Parent focus selects this pane at narrow widths."}}},
}}
if arrangeErr := screen.Arrange(root, layout.Size{Width: 80, Height: 16}); arrangeErr != nil {
	return "", arrangeErr
}

views := map[string]string{}

for _, region := range screen.Plan().Regions() {
	size := layout.Size{Width: region.Bounds.Dx(), Height: region.Bounds.Dy()}
	views[region.ID] = layout.Panel(region.ID, "Native view", size, theme.New(theme.Dark), false)
}

frame, err := screen.Render(views)

return frame.Content, err
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-layout.html#layout%2fresponsive-integration)

Declare region identities and hard minimum sizes, then arrange a registered tui.Screen at each terminal resize. Resize native child models from the resulting rectangles and render exactly the visible region views. Geometry and design rules block by default; exceptions name a specific rule and reason on its region or container. Hidden components retain caller-owned state, and CheckLayouts exercises every responsive variant.

### [Run this example](https://lanej.io/charmed/molecules-layout.html#layout%2fresponsive-run)

```sh
go run ./examples/catalog layout/responsive
```

[Example source](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/docs/snippets/layout_responsive.go) · [Implementation](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/story/layouts.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [layout/grid](https://lanej.io/charmed/molecules-layout.html#layout/grid) [layout/split-pane](https://lanej.io/charmed/molecules-layout.html#layout/split-pane)
