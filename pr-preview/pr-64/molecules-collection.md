[Molecules](https://lanej.io/charmed/molecules.html)Collections

# [Collections](https://lanej.io/charmed/molecules-collection.html#level-molecules-collection)

2 components · [All molecules](https://lanej.io/charmed/molecules.html)

- [Static table](https://lanej.io/charmed/molecules-collection.html#collection%2fstatic-table)
- [Scroll pane](https://lanej.io/charmed/molecules-collection.html#collection%2fscroll-pane)

[Molecules](https://lanej.io/charmed/molecules.html) / Collections

## [Static table](https://lanej.io/charmed/molecules-collection.html#collection%2fstatic-table)

A non-interactive comparison table. Cells wrap in their column; 55 columns and narrower stack a table that would not fit whole, never truncating a value. Missing, unknown and truncated cells keep their word and glyph.

### [Native preview](https://lanej.io/charmed/molecules-collection.html#collection%2fstatic-table-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
HOST      VERSION       STATE
web1west  9f2c4e1a7b3d   ✓ HEALTHY 
web2west  ∅ missing      ? UNKNOWN 
```

Demo · truecolor · 64 × 8 terminal cells

```
At the allocation                                               
HOST      VERSION           STATE          NOTE                 
web1west  9f2c4e1a7b3d5c60   ✓ HEALTHY     serving              
web2west  ∅ missing          ? UNKNOWN     no report received   
web3west  3b1d7e90a2c4f658   ◷ STALE  18m  last read ! source   
                                           truncated (42        
                                           omitted)             
web4west  ""                ∅ missing      ∅ missing
```

Narrow · truecolor · 40 × 16 terminal cells

```
At the allocation                       
web1west                                
VERSION: 9f2c4e1a7b3d5c60               
STATE:  ✓ HEALTHY                       
NOTE: serving                           
                                        
web2west                                
VERSION: ∅ missing                      
STATE:  ? UNKNOWN                       
NOTE: no report received                
                                        
web3west                                
VERSION: 3b1d7e90a2c4f658               
STATE:  ◷ STALE  18m                    
NOTE: last read ! source truncated (42  
      omitted)
```

256 colors · ansi256 · 64 × 8 terminal cells

```
At the allocation                                               
HOST      VERSION           STATE          NOTE                 
web1west  9f2c4e1a7b3d5c60   ✓ HEALTHY     serving              
web2west  ∅ missing          ? UNKNOWN     no report received   
web3west  3b1d7e90a2c4f658   ◷ STALE  18m  last read ! source   
                                           truncated (42        
                                           omitted)             
web4west  ""                ∅ missing      ∅ missing
```

16 colors · ansi16 · 64 × 8 terminal cells

```
At the allocation                                               
HOST      VERSION           STATE          NOTE                 
web1west  9f2c4e1a7b3d5c60   ✓ HEALTHY     serving              
web2west  ∅ missing          ? UNKNOWN     no report received   
web3west  3b1d7e90a2c4f658   ◷ STALE  18m  last read ! source   
                                           truncated (42        
                                           omitted)             
web4west  ""                ∅ missing      ∅ missing
```

Monochrome · none · 64 × 8 terminal cells

```
At the allocation                                               
HOST      VERSION           STATE          NOTE                 
web1west  9f2c4e1a7b3d5c60  [+ HEALTHY]    serving              
web2west  missing           [? UNKNOWN]    no report received   
web3west  3b1d7e90a2c4f658  [t STALE] 18m  last read source     
                                           truncated (42        
                                           omitted)             
web4west  ""                missing        missing
```

### [Go usage](https://lanej.io/charmed/molecules-collection.html#collection%2fstatic-table-code)

Imports: [`collection`](https://pkg.go.dev/github.com/lanej/charmed/collection "github.com/lanej/charmed/collection") · [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
table := collection.StaticTable{
	Columns: []collection.TableColumn{{Title: "HOST", Fixed: true}, {Title: "VERSION", Min: 10}, {Title: "STATE"}},
	Rows: [][]collection.TableCell{
		{collection.TableText("web1west"), collection.TableText("9f2c4e1a7b3d"), collection.TableEvidence(feedback.EvidenceHealthy, "")},
		{collection.TableText("web2west"), collection.TableMissing(), collection.TableEvidence(feedback.EvidenceUnknown, "")}, // Missing is not empty.
	},
}

return table.View(48, theme.New(theme.Dark))
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-collection.html#collection%2fstatic-table-integration)

Retain the native model and any command returned by Update, and route keyboard input through the parent's active focus region. Replace displayed content through the native API. Recalculate width and height from the available terminal cells while retaining selection where the native model supports it.

### [Run this example](https://lanej.io/charmed/molecules-collection.html#collection%2fstatic-table-run)

```sh
go run ./examples/catalog collection/static-table
```

[Example source](https://github.com/lanej/charmed/blob/d022f37d13ba3e076a70ceeccf337a2d99010417/docs/snippets/collection_static_table.go) · [Implementation](https://github.com/lanej/charmed/blob/d022f37d13ba3e076a70ceeccf337a2d99010417/collection/static_table.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [feedback/evidence](https://lanej.io/charmed/atoms-feedback.html#feedback/evidence) [text/value-state](https://lanej.io/charmed/atoms-text.html#text/value-state) [text/omission](https://lanej.io/charmed/atoms-text.html#text/omission) [foundations/layout](https://lanej.io/charmed/foundations-layout.html#foundations/layout)

[Molecules](https://lanej.io/charmed/molecules.html) / Collections

## [Scroll pane](https://lanej.io/charmed/molecules-collection.html#collection%2fscroll-pane)

Bind an existing viewport to a proportional scrollbar without replacing its native model. Use this when content scrolling belongs to a focused pane inside a larger composition.

### [Native preview](https://lanej.io/charmed/molecules-collection.html#collection%2fscroll-pane-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Request observations remain scrolla│
ble.                               ┃
Request observations remain scrolla│
ble.                               │
Request observations remain scrolla│
ble.                               │
```

Demo · truecolor · 100 × 30 terminal cells

```
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            │
Request observations remain scrollable.                                                            │
Request observations remain scrollable.                                                            │
Request observations remain scrollable.                                                            │
Request observations remain scrollable.                                                            │
Request observations remain scrollable.                                                            │
Request observations remain scrollable.                                                            │
Request observations remain scrollable.                                                            │
```

Narrow · truecolor · 40 × 16 terminal cells

```
Request observations remain scrollable.┃
Request observations remain scrollable.┃
Request observations remain scrollable.┃
Request observations remain scrollable.┃
Request observations remain scrollable.┃
Request observations remain scrollable.┃
Request observations remain scrollable.│
Request observations remain scrollable.│
Request observations remain scrollable.│
Request observations remain scrollable.│
Request observations remain scrollable.│
Request observations remain scrollable.│
Request observations remain scrollable.│
Request observations remain scrollable.│
Request observations remain scrollable.│
Request observations remain scrollable.│
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            │
Request observations remain scrollable.                                                            │
Request observations remain scrollable.                                                            │
Request observations remain scrollable.                                                            │
Request observations remain scrollable.                                                            │
Request observations remain scrollable.                                                            │
Request observations remain scrollable.                                                            │
Request observations remain scrollable.                                                            │
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            ┃
Request observations remain scrollable.                                                            │
Request observations remain scrollable.                                                            │
Request observations remain scrollable.                                                            │
Request observations remain scrollable.                                                            │
Request observations remain scrollable.                                                            │
Request observations remain scrollable.                                                            │
Request observations remain scrollable.                                                            │
Request observations remain scrollable.                                                            │
```

Monochrome · none · 100 × 30 terminal cells

```
Request observations remain scrollable.                                                            #
Request observations remain scrollable.                                                            #
Request observations remain scrollable.                                                            #
Request observations remain scrollable.                                                            #
Request observations remain scrollable.                                                            #
Request observations remain scrollable.                                                            #
Request observations remain scrollable.                                                            #
Request observations remain scrollable.                                                            #
Request observations remain scrollable.                                                            #
Request observations remain scrollable.                                                            #
Request observations remain scrollable.                                                            #
Request observations remain scrollable.                                                            #
Request observations remain scrollable.                                                            #
Request observations remain scrollable.                                                            #
Request observations remain scrollable.                                                            #
Request observations remain scrollable.                                                            #
Request observations remain scrollable.                                                            #
Request observations remain scrollable.                                                            #
Request observations remain scrollable.                                                            #
Request observations remain scrollable.                                                            #
Request observations remain scrollable.                                                            #
Request observations remain scrollable.                                                            #
Request observations remain scrollable.                                                            |
Request observations remain scrollable.                                                            |
Request observations remain scrollable.                                                            |
Request observations remain scrollable.                                                            |
Request observations remain scrollable.                                                            |
Request observations remain scrollable.                                                            |
Request observations remain scrollable.                                                            |
Request observations remain scrollable.                                                            |
```

### [Go usage](https://lanej.io/charmed/molecules-collection.html#collection%2fscroll-pane-code)

Imports: [`strings`](https://pkg.go.dev/strings "strings") · [`collection`](https://pkg.go.dev/github.com/lanej/charmed/collection "github.com/lanej/charmed/collection") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
native := collection.NewViewport("", 36, 6)
native.SoftWrap = true
pane := collection.NewScrollPane(&native, layout.Size{Width: 36, Height: 6}, theme.New(theme.Dark))
pane.SetContent(strings.Repeat("Request observations remain scrollable.\n", 12))
native.ScrollDown(2)
// Route Update to the pane while active; Resize retains native customization.
return pane.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-collection.html#collection%2fscroll-pane-integration)

Pass a non-nil native viewport pointer and resize the pane from its full allocation. Configure native styles, wrapping and key bindings directly. SetContent preserves and clamps scrolling; call GotoTop explicitly for a new selection. Forward Update commands while active. Resize after changing ShowScrollbar; content and dimensions are managed by the pane.

### [Run this example](https://lanej.io/charmed/molecules-collection.html#collection%2fscroll-pane-run)

```sh
go run ./examples/catalog collection/scroll-pane
```

[Example source](https://github.com/lanej/charmed/blob/d022f37d13ba3e076a70ceeccf337a2d99010417/docs/snippets/collection_scroll_pane.go) · [Implementation](https://github.com/lanej/charmed/blob/d022f37d13ba3e076a70ceeccf337a2d99010417/collection/scroll_pane.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

### Reusable component

[Scroll pane dependency graph](https://lanej.io/charmed/dependencies.html#component-collection%2Fscroll-pane) · Partial map · [Source declaration](https://github.com/lanej/charmed/blob/d022f37d13ba3e076a70ceeccf337a2d99010417/collection/scroll_pane.go#L14)

- [Viewport](https://lanej.io/charmed/dependencies.html#component-collection%2Fviewport) · composes · content viewport
- [Scrollbar](https://lanej.io/charmed/dependencies.html#component-feedback%2Fscrollbar) · composes · scroll position

Consumer usage uses

- [Scroll pane](https://lanej.io/charmed/dependencies.html#component-collection%2Fscroll-pane) · primary demonstration

[Example source](https://github.com/lanej/charmed/blob/d022f37d13ba3e076a70ceeccf337a2d99010417/docs/snippets/collection_scroll_pane.go#L11)
