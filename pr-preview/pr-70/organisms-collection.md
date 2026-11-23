[Organisms](https://lanej.io/charmed/organisms.html)Collections

# [Collections](https://lanej.io/charmed/organisms-collection.html#level-organisms-collection)

9 components · [All organisms](https://lanej.io/charmed/organisms.html)

- [Data grid](https://lanej.io/charmed/organisms-collection.html#collection%2fdata-grid)
- [Status matrix](https://lanej.io/charmed/organisms-collection.html#collection%2fstatus-matrix)
- [List](https://lanej.io/charmed/organisms-collection.html#collection%2flist)
- [File picker](https://lanej.io/charmed/organisms-collection.html#collection%2ffiles)
- [Log viewer](https://lanej.io/charmed/organisms-collection.html#collection%2flogs)
- [Grid inspector](https://lanej.io/charmed/organisms-collection.html#collection%2fgrid-inspector)
- [Keyed table](https://lanej.io/charmed/organisms-collection.html#collection%2fkeyed-table)
- [Section stack](https://lanej.io/charmed/organisms-collection.html#collection%2fsection-stack)
- [Outline document](https://lanej.io/charmed/organisms-collection.html#pattern%2foutline-document)

[Organisms](https://lanej.io/charmed/organisms.html) / Collections

## [Data grid](https://lanej.io/charmed/organisms-collection.html#collection%2fdata-grid)

Typed values, stable axes, pinned headers and two-axis keyboard navigation. Plain values render directly; caller renderers own cell meaning and styling. Instructions are opt-in.

### [Native preview](https://lanej.io/charmed/organisms-collection.html#collection%2fdata-grid-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 Service / Region   US East          US West          
 API               ▏24            ▕  0                
 Workers            48                                
Rows 1-2/2  Columns 1-2/2
```

Demo · truecolor · 100 × 30 terminal cells

```
Replica counts by region                                                                            
                                                                                                    
 Service / Region   US East          US West          Europe           Asia Pacific                 
 API               ▏24            ▕  16               12               8                            
 Workers            48               32               16                                            
 Cache              6                4                0                                             
Rows 1-3/3  Columns 1-4/4                                                                           
api / east: 24 replicas
```

Narrow · truecolor · 40 × 16 terminal cells

```
Replica counts by region                
                                        
 Service / Region   US East             
 API               ▏24            ▕     
 Workers            48                  
 Cache              6                   
Rows 1-3/3  Columns 1-1/4               
api / east: 24 replicas
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Replica counts by region                                                                            
                                                                                                    
 Service / Region   US East          US West          Europe           Asia Pacific                 
 API               ▏24            ▕  16               12               8                            
 Workers            48               32               16                                            
 Cache              6                4                0                                             
Rows 1-3/3  Columns 1-4/4                                                                           
api / east: 24 replicas
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Replica counts by region                                                                            
                                                                                                    
 Service / Region   US East          US West          Europe           Asia Pacific                 
 API               ▏24            ▕  16               12               8                            
 Workers            48               32               16                                            
 Cache              6                4                0                                             
Rows 1-3/3  Columns 1-4/4                                                                           
api / east: 24 replicas
```

Monochrome · none · 100 × 30 terminal cells

```
Replica counts by region                                                                            
                                                                                                    
 Service / Region   US East          US West          Europe           Asia Pacific                 
 API               [24            ]  16               12               8                            
 Workers            48               32               16                                            
 Cache              6                4                0                                             
Rows 1-3/3  Columns 1-4/4                                                                           
api / east: 24 replicas
```

Static · truecolor · 100 × 30 terminal cells

```
Replica counts by region                                                                            
                                                                                                    
 Service / Region   US East          US West          Europe           Asia Pacific                 
 API               ▏24            ▕  16               12               8                            
 Workers            48               32               16                                            
 Cache              6                4                0                                             
api / east: 24 replicas
```

### [Go usage](https://lanej.io/charmed/organisms-collection.html#collection%2fdata-grid-code)

Imports: [`collection`](https://pkg.go.dev/github.com/lanej/charmed/collection "github.com/lanej/charmed/collection") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
rows := []collection.GridAxis{{ID: "api", Label: "API"}, {ID: "worker", Label: "Workers"}}
regions := []collection.GridAxis{{ID: "east", Label: "US East"}, {ID: "west", Label: "US West"}}
entries := []collection.GridEntry[int]{
	{RowID: "api", ColumnID: "east", Value: 24},
	{RowID: "api", ColumnID: "west", Value: 0},
	{RowID: "worker", ColumnID: "east", Value: 48},
}

grid, err := collection.NewDataGrid(rows, regions, entries, layout.Size{Width: 54, Height: 4}, theme.New(theme.Dark))
if err != nil {
	return "", err
}

grid.Corner = "Service / Region"
grid.Focus()
// The parent forwards Update(msg), returns its command and calls Resize(size).
// Read Selected() for stable IDs and presence; zero is a present value.
// Set ShowHelp to opt into instructions, HidePosition for a complete static render,
// or RenderCell for custom styling.
return grid.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-collection.html#collection%2fdata-grid-integration)

Construct the grid with ordered axes and keyed entries, then focus it when its region becomes active. Forward Update messages and return its command; the parent owns Tab routing and Resize allocations. Replace observations through SetData and inspect Selected for stable IDs and presence. Theme, geometry, renderer and optional help survive updates.

### [Run this example](https://lanej.io/charmed/organisms-collection.html#collection%2fdata-grid-run)

```sh
go run ./examples/catalog collection/data-grid
```

[Example source](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/docs/snippets/collection_data_grid.go) · [Implementation](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/collection/data_grid.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

### Reusable component

[Data grid dependency graph](https://lanej.io/charmed/dependencies.html#component-collection%2Fdata-grid) · Partial map · [Source declaration](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/collection/data_grid.go#L44)

Consumer usage uses

- [Data grid](https://lanej.io/charmed/dependencies.html#component-collection%2Fdata-grid) · primary demonstration

[Example source](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/docs/snippets/collection_data_grid.go#L9)

[Organisms](https://lanej.io/charmed/organisms.html) / Collections

## [Status matrix](https://lanej.io/charmed/organisms-collection.html#collection%2fstatus-matrix)

Stable row/column identities and sparse status cells compose a keyboard-navigable mapping grid with pinned headers. Arrows or h/j/k/l reveal off-screen hosts; absent mappings differ from unknown health.

### [Native preview](https://lanej.io/charmed/organisms-collection.html#collection%2fstatus-matrix-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 Balancer / Host    ▤ host-01        ▤ host-02        ▤ host-03       
 ⇄ lb-public       ▏✓ v2.8.1      ▕  ! draining       ✓ v2.8.0        
 ⇄ lb-internal      —                × unhealthy      ? no probe      
Rows 1-2/2  Columns 1-3/3
```

Demo · truecolor · 100 × 30 terminal cells

```
Load balancer / host mapping                                                                        
                                                                                                    
 Balancer / Host    ▤ host-01        ▤ host-02        ▤ host-03        ▤ host-04                    
 ⇄ lb-public       ▏✓ v2.8.1      ▕  ✓ v2.8.1         ! draining       ✓ v2.8.0                     
 ⇄ lb-internal      —                × unhealthy      ✓ v2.8.0         —                            
 ⇄ lb-canary        —                —                —                ▶ warming                    
 ⇄ lb-batch         —                —                —                —                            
Rows 1-4/4  Columns 1-4/6                                                                           
public / host-01: OK | v2.8.1
```

Narrow · truecolor · 40 × 16 terminal cells

```
Load balancer / host mapping            
                                        
 Balancer / Host    ▤ host-01           
 ⇄ lb-public       ▏✓ v2.8.1      ▕     
 ⇄ lb-internal      —                   
 ⇄ lb-canary        —                   
 ⇄ lb-batch         —                   
Rows 1-4/4  Columns 1-1/6               
public / host-01: OK | v2.8.1
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Load balancer / host mapping                                                                        
                                                                                                    
 Balancer / Host    ▤ host-01        ▤ host-02        ▤ host-03        ▤ host-04                    
 ⇄ lb-public       ▏✓ v2.8.1      ▕  ✓ v2.8.1         ! draining       ✓ v2.8.0                     
 ⇄ lb-internal      —                × unhealthy      ✓ v2.8.0         —                            
 ⇄ lb-canary        —                —                —                ▶ warming                    
 ⇄ lb-batch         —                —                —                —                            
Rows 1-4/4  Columns 1-4/6                                                                           
public / host-01: OK | v2.8.1
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Load balancer / host mapping                                                                        
                                                                                                    
 Balancer / Host    ▤ host-01        ▤ host-02        ▤ host-03        ▤ host-04                    
 ⇄ lb-public       ▏✓ v2.8.1      ▕  ✓ v2.8.1         ! draining       ✓ v2.8.0                     
 ⇄ lb-internal      —                × unhealthy      ✓ v2.8.0         —                            
 ⇄ lb-canary        —                —                —                ▶ warming                    
 ⇄ lb-batch         —                —                —                —                            
Rows 1-4/4  Columns 1-4/6                                                                           
public / host-01: OK | v2.8.1
```

Monochrome · none · 100 × 30 terminal cells

```
Load balancer / host mapping                                                                        
                                                                                                    
 Balancer / Host    H host-01        H host-02        H host-03        H host-04                    
 L lb-public       [+ v2.8.1      ]  + v2.8.1         ! draining       + v2.8.0                     
 L lb-internal      -                x unhealthy      + v2.8.0         -                            
 L lb-canary        -                -                -                > warming                    
 L lb-batch         -                -                -                -                            
Rows 1-4/4  Columns 1-4/6                                                                           
public / host-01: OK | v2.8.1
```

### [Go usage](https://lanej.io/charmed/organisms-collection.html#collection%2fstatus-matrix-code)

Imports: [`collection`](https://pkg.go.dev/github.com/lanej/charmed/collection "github.com/lanej/charmed/collection") · [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
balancer, host := palette.Glyph(theme.GlyphLoadBalancer), palette.Glyph(theme.GlyphHost)
rows := []collection.MatrixAxis{
	{ID: "public", Label: balancer + " lb-public"},
	{ID: "internal", Label: balancer + " lb-internal"},
}
hosts := []collection.MatrixAxis{
	{ID: "host-01", Label: host + " host-01"},
	{ID: "host-02", Label: host + " host-02"},
	{ID: "host-03", Label: host + " host-03"},
}
entries := []collection.MatrixEntry{
	{RowID: "public", ColumnID: "host-01", Cell: feedback.StatusCell{Status: feedback.OK, Label: "v2.8.1"}},
	{RowID: "public", ColumnID: "host-02", Cell: feedback.StatusCell{Status: feedback.Warn, Label: "draining"}},
	{RowID: "public", ColumnID: "host-03", Cell: feedback.StatusCell{Status: feedback.OK, Label: "v2.8.0"}},
	{RowID: "internal", ColumnID: "host-02", Cell: feedback.StatusCell{Status: feedback.Error, Label: "unhealthy"}},
	{RowID: "internal", ColumnID: "host-03", Cell: feedback.StatusCell{Status: feedback.Unknown, Label: "no probe"}},
}

matrix, err := collection.NewStatusMatrix(rows, hosts, entries, layout.Size{Width: 70, Height: 4}, palette)
if err != nil {
	return "", err
}

matrix.Corner = "Balancer / Host"
matrix.Focus()
// Forward Update(msg) while focused; the parent owns Tab and Resize(size).
// SetData retains selection by ID. Selected distinguishes absent mappings.
// Set ShowHelp to opt into instructions; the embedded DataGrid stays accessible.
return matrix.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-collection.html#collection%2fstatus-matrix-integration)

Construct the grid with ordered axes and keyed entries, then focus it when its region becomes active. Forward Update messages and return its command; the parent owns Tab routing and Resize allocations. Replace observations through SetData and inspect Selected for stable IDs and presence. Theme, geometry, renderer and optional help survive updates.

### [Run this example](https://lanej.io/charmed/organisms-collection.html#collection%2fstatus-matrix-run)

```sh
go run ./examples/catalog collection/status-matrix
```

[Example source](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/docs/snippets/collection_status_matrix.go) · [Implementation](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/collection/status_matrix.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

### Reusable component

[Status matrix dependency graph](https://lanej.io/charmed/dependencies.html#component-collection%2Fstatus-matrix) · Partial map · [Source declaration](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/collection/status_matrix.go#L23)

- [Data grid](https://lanej.io/charmed/dependencies.html#component-collection%2Fdata-grid) · specializes · grid base
- [Status cell](https://lanej.io/charmed/dependencies.html#component-feedback%2Fstatus-cell) · composes · cell presentation

Consumer usage uses

- [Status matrix](https://lanej.io/charmed/dependencies.html#component-collection%2Fstatus-matrix) · primary demonstration

[Example source](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/docs/snippets/collection_status_matrix.go#L10)

[Organisms](https://lanej.io/charmed/organisms.html) / Collections

## [List](https://lanej.io/charmed/organisms-collection.html#collection%2flist)

A native list with filtering, pagination and selection.

### [Native preview](https://lanej.io/charmed/organisms-collection.html#collection%2flist-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
  Environments                         
                                       
  3 items                              
                                       
 Production                            
 Primary service environment           
                                       
 Staging                               
 Integration environment               
                                       
                                       
  ◉○                                   
                                       
  ↑/k up • ↓/j down • / filter • ? more
```

Demo · truecolor · 100 × 30 terminal cells

```
  Environments                                                                                      
                                                                                                    
  3 items                                                                                           
                                                                                                    
 Production                                                                                         
 Primary service environment                                                                        
                                                                                                    
 Staging                                                                                            
 Integration environment                                                                            
                                                                                                    
 Preview                                                                                            
 Disposable review environment                                                                      
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
  ↑/k up • ↓/j down • / filter • ? more
```

Narrow · truecolor · 40 × 16 terminal cells

```
  Environments                          
                                        
  3 items                               
                                        
 Production                             
 Primary service environment            
                                        
 Staging                                
 Integration environment                
                                        
 Preview                                
 Disposable review environment          
                                        
                                        
                                        
  ↑/k up • ↓/j down • / filter • ? more
```

256 colors · ansi256 · 100 × 30 terminal cells

```
  Environments                                                                                      
                                                                                                    
  3 items                                                                                           
                                                                                                    
 Production                                                                                         
 Primary service environment                                                                        
                                                                                                    
 Staging                                                                                            
 Integration environment                                                                            
                                                                                                    
 Preview                                                                                            
 Disposable review environment                                                                      
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
  ↑/k up • ↓/j down • / filter • ? more
```

16 colors · ansi16 · 100 × 30 terminal cells

```
  Environments                                                                                      
                                                                                                    
  3 items                                                                                           
                                                                                                    
 Production                                                                                         
 Primary service environment                                                                        
                                                                                                    
 Staging                                                                                            
 Integration environment                                                                            
                                                                                                    
 Preview                                                                                            
 Disposable review environment                                                                      
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
  ↑/k up • ↓/j down • / filter • ? more
```

Monochrome · none · 100 × 30 terminal cells

```
  Environments                                                                                      
                                                                                                    
  3 items                                                                                           
                                                                                                    
 Production                                                                                         
 Primary service environment                                                                        
                                                                                                    
 Staging                                                                                            
 Integration environment                                                                            
                                                                                                    
 Preview                                                                                            
 Disposable review environment                                                                      
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
  ^/k up * v/j down * / filter * ? more
```

### [Go usage](https://lanej.io/charmed/organisms-collection.html#collection%2flist-code)

Imports: [`list`](https://pkg.go.dev/charm.land/bubbles/v2/list "charm.land/bubbles/v2/list") · [`collection`](https://pkg.go.dev/github.com/lanej/charmed/collection "github.com/lanej/charmed/collection") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
type environmentItem struct{ name, description string }

func (i environmentItem) Title() string { return i.name }

func (i environmentItem) Description() string { return i.description }

func (i environmentItem) FilterValue() string { return i.name }

items := []list.Item{
	environmentItem{"Production", "Primary service environment"},
	environmentItem{"Staging", "Integration environment"},
	environmentItem{"Preview", "Disposable review environment"},
}
environments := collection.NewList(items, 48, 14, theme.New(theme.Dark))
environments.Title = "Environments"

return environments.View()
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-collection.html#collection%2flist-integration)

The parent model supplies observations, coordinates focus and owns application decisions. Forward initialization and update commands rather than executing them during rendering. Resize the existing composition from the terminal allocation so selection remains stable, and apply native customization at the exposed child models after any rebuilding lifecycle step.

### [Run this example](https://lanej.io/charmed/organisms-collection.html#collection%2flist-run)

```sh
go run ./examples/catalog collection/list
```

[Example source](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/docs/snippets/collection_list.go) · [Implementation](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/story/organisms.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Organisms](https://lanej.io/charmed/organisms.html) / Collections

## [File picker](https://lanej.io/charmed/organisms-collection.html#collection%2ffiles)

Let someone browse a directory and select an allowed file type. Initialize the native picker before capturing its view so the example contains actual files.

### [Native preview](https://lanej.io/charmed/organisms-collection.html#collection%2ffiles-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
>  README.md 
   main.go 
   theme.go
```

Demo · truecolor · 100 × 30 terminal cells

```
Fixture files                                                                                       
                                                                                                    
>  components                                                                                       
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
Arrows browse / Enter selects / fixture directory only
```

Narrow · truecolor · 40 × 16 terminal cells

```
Fixture files                           
                                        
>  components                           
                                        
                                        
                                        
                                        
                                        
                                        
                                        
                                        
                                        
                                        
                                        
Arrows browse / Enter selects / fixture 
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Fixture files                                                                                       
                                                                                                    
>  components                                                                                       
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
Arrows browse / Enter selects / fixture directory only
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Fixture files                                                                                       
                                                                                                    
>  components                                                                                       
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
Arrows browse / Enter selects / fixture directory only
```

Monochrome · none · 100 × 30 terminal cells

```
Fixture files                                                                                       
                                                                                                    
>  components                                                                                       
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
Arrows browse / Enter selects / fixture directory only
```

### [Go usage](https://lanej.io/charmed/organisms-collection.html#collection%2ffiles-code)

Imports: [`errors`](https://pkg.go.dev/errors "errors") · [`os`](https://pkg.go.dev/os "os") · [`filepath`](https://pkg.go.dev/path/filepath "path/filepath") · [`collection`](https://pkg.go.dev/github.com/lanej/charmed/collection "github.com/lanej/charmed/collection") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
directory, err := os.MkdirTemp("", "charmed-example-")
if err != nil {
	return "", err
}
defer func() { err = errors.Join(err, os.RemoveAll(directory)) }()

for _, name := range []string{"main.go", "README.md", "theme.go"} {
	if err := os.WriteFile(filepath.Join(directory, name), []byte("// Sample file\n"), 0o600); err != nil {
		return "", err
	}
}

picker := collection.NewFilePicker(directory, theme.New(theme.Dark))
picker.AllowedTypes = []string{".go", ".md"}
picker.ShowPermissions, picker.ShowSize = false, false
picker.SetHeight(6)
// A running Bubble Tea program executes this command and sends its message.
picker, _ = picker.Update(picker.Init()())

return picker.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-collection.html#collection%2ffiles-integration)

In a running application, return picker.Init from Init and retain picker.Update's model and command. Inspect DidSelectFile on incoming messages to handle selection. This snapshot executes the initial directory-read command synchronously; a temporary directory makes its output reproducible. SetHeight when the parent allocation changes.

### [Run this example](https://lanej.io/charmed/organisms-collection.html#collection%2ffiles-run)

```sh
go run ./examples/catalog collection/files
```

[Example source](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/docs/snippets/collection_files.go) · [Implementation](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/story/organisms.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Organisms](https://lanej.io/charmed/organisms.html) / Collections

## [Log viewer](https://lanej.io/charmed/organisms-collection.html#collection%2flogs)

Structured Charm Log entries in a bounded, scrollable viewport.

### [Native preview](https://lanej.io/charmed/organisms-collection.html#collection%2flogs-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
12:04:00  INFO   Request completed service=api status=200       
00:00:00  WARN   Retry scheduled
```

Demo · truecolor · 100 × 30 terminal cells

```
Structured logs                                                                                     
                                                                                                    
12:00:00  INFO   Fixture request completed service=api sequence=0                                   
12:00:01  INFO   Fixture request completed service=api sequence=1                                   
12:00:02  INFO   Fixture request completed service=api sequence=2                                   
12:00:03  INFO   Fixture request completed service=api sequence=3                                   
12:00:04  WARN   Fixture request completed service=api sequence=4                                   
12:00:05  INFO   Fixture request completed service=api sequence=5                                   
12:00:06  INFO   Fixture request completed service=api sequence=6                                   
12:00:07  INFO   Fixture request completed service=api sequence=7                                   
12:00:08  INFO   Fixture request completed service=api sequence=8                                   
12:00:09  WARN   Fixture request completed service=api sequence=9                                   
12:00:10  INFO   Fixture request completed service=api sequence=10                                  
12:00:11  INFO   Fixture request completed service=api sequence=11                                  
12:00:12  INFO   Fixture request completed service=api sequence=12                                  
12:00:13  INFO   Fixture request completed service=api sequence=13                                  
12:00:14  WARN   Fixture request completed service=api sequence=14                                  
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
Arrows scroll / Ctrl+N appends a deterministic log entry
```

Narrow · truecolor · 40 × 16 terminal cells

```
Structured logs                         
                                        
12:00:03  INFO   Fixture request complet
12:00:04  WARN   Fixture request complet
12:00:05  INFO   Fixture request complet
12:00:06  INFO   Fixture request complet
12:00:07  INFO   Fixture request complet
12:00:08  INFO   Fixture request complet
12:00:09  WARN   Fixture request complet
12:00:10  INFO   Fixture request complet
12:00:11  INFO   Fixture request complet
12:00:12  INFO   Fixture request complet
12:00:13  INFO   Fixture request complet
12:00:14  WARN   Fixture request complet
                                        
Arrows scroll / Ctrl+N appends a determi
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Structured logs                                                                                     
                                                                                                    
12:00:00  INFO   Fixture request completed service=api sequence=0                                   
12:00:01  INFO   Fixture request completed service=api sequence=1                                   
12:00:02  INFO   Fixture request completed service=api sequence=2                                   
12:00:03  INFO   Fixture request completed service=api sequence=3                                   
12:00:04  WARN   Fixture request completed service=api sequence=4                                   
12:00:05  INFO   Fixture request completed service=api sequence=5                                   
12:00:06  INFO   Fixture request completed service=api sequence=6                                   
12:00:07  INFO   Fixture request completed service=api sequence=7                                   
12:00:08  INFO   Fixture request completed service=api sequence=8                                   
12:00:09  WARN   Fixture request completed service=api sequence=9                                   
12:00:10  INFO   Fixture request completed service=api sequence=10                                  
12:00:11  INFO   Fixture request completed service=api sequence=11                                  
12:00:12  INFO   Fixture request completed service=api sequence=12                                  
12:00:13  INFO   Fixture request completed service=api sequence=13                                  
12:00:14  WARN   Fixture request completed service=api sequence=14                                  
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
Arrows scroll / Ctrl+N appends a deterministic log entry
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Structured logs                                                                                     
                                                                                                    
12:00:00  INFO   Fixture request completed service=api sequence=0                                   
12:00:01  INFO   Fixture request completed service=api sequence=1                                   
12:00:02  INFO   Fixture request completed service=api sequence=2                                   
12:00:03  INFO   Fixture request completed service=api sequence=3                                   
12:00:04  WARN   Fixture request completed service=api sequence=4                                   
12:00:05  INFO   Fixture request completed service=api sequence=5                                   
12:00:06  INFO   Fixture request completed service=api sequence=6                                   
12:00:07  INFO   Fixture request completed service=api sequence=7                                   
12:00:08  INFO   Fixture request completed service=api sequence=8                                   
12:00:09  WARN   Fixture request completed service=api sequence=9                                   
12:00:10  INFO   Fixture request completed service=api sequence=10                                  
12:00:11  INFO   Fixture request completed service=api sequence=11                                  
12:00:12  INFO   Fixture request completed service=api sequence=12                                  
12:00:13  INFO   Fixture request completed service=api sequence=13                                  
12:00:14  WARN   Fixture request completed service=api sequence=14                                  
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
Arrows scroll / Ctrl+N appends a deterministic log entry
```

Monochrome · none · 100 × 30 terminal cells

```
Structured logs                                                                                     
                                                                                                    
12:00:00 [INFO ] Fixture request completed service=api sequence=0                                   
12:00:01 [INFO ] Fixture request completed service=api sequence=1                                   
12:00:02 [INFO ] Fixture request completed service=api sequence=2                                   
12:00:03 [INFO ] Fixture request completed service=api sequence=3                                   
12:00:04 [WARN ] Fixture request completed service=api sequence=4                                   
12:00:05 [INFO ] Fixture request completed service=api sequence=5                                   
12:00:06 [INFO ] Fixture request completed service=api sequence=6                                   
12:00:07 [INFO ] Fixture request completed service=api sequence=7                                   
12:00:08 [INFO ] Fixture request completed service=api sequence=8                                   
12:00:09 [WARN ] Fixture request completed service=api sequence=9                                   
12:00:10 [INFO ] Fixture request completed service=api sequence=10                                  
12:00:11 [INFO ] Fixture request completed service=api sequence=11                                  
12:00:12 [INFO ] Fixture request completed service=api sequence=12                                  
12:00:13 [INFO ] Fixture request completed service=api sequence=13                                  
12:00:14 [WARN ] Fixture request completed service=api sequence=14                                  
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
Arrows scroll / Ctrl+N appends a deterministic log entry
```

### [Go usage](https://lanej.io/charmed/organisms-collection.html#collection%2flogs-code)

Imports: [`time`](https://pkg.go.dev/time "time") · [`log`](https://pkg.go.dev/charm.land/log/v2 "charm.land/log/v2") · [`collection`](https://pkg.go.dev/github.com/lanej/charmed/collection "github.com/lanej/charmed/collection") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
logs := collection.NewLogViewer(64, 8, theme.New(theme.Dark))
logs.Append(collection.LogEntry{
	Time:  time.Date(2026, 11, 8, 12, 4, 0, 0, time.UTC),
	Level: log.InfoLevel, Message: "Request completed",
	Fields: []any{"service", "api", "status", 200},
})
logs.Append(collection.LogEntry{Level: log.WarnLevel, Message: "Retry scheduled"})

return logs.View()
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-collection.html#collection%2flogs-integration)

The parent model supplies observations, coordinates focus and owns application decisions. Forward initialization and update commands rather than executing them during rendering. Resize the existing composition from the terminal allocation so selection remains stable, and apply native customization at the exposed child models after any rebuilding lifecycle step.

### [Run this example](https://lanej.io/charmed/organisms-collection.html#collection%2flogs-run)

```sh
go run ./examples/catalog collection/logs
```

[Example source](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/docs/snippets/collection_logs.go) · [Implementation](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/story/organisms.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [collection/viewport](https://lanej.io/charmed/atoms-collection.html#collection/viewport) [feedback/log-level](https://lanej.io/charmed/atoms-feedback.html#feedback/log-level)

[Organisms](https://lanej.io/charmed/organisms.html) / Collections

## [Grid inspector](https://lanej.io/charmed/organisms-collection.html#collection%2fgrid-inspector)

Inspect exact values behind a selected grid coordinate. Keep sparse coordinates visible while rendering application-specific details in an independent scrollable pane.

### [Native preview](https://lanej.io/charmed/organisms-collection.html#collection%2fgrid-inspector-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 Grid               API              Workers                          
 API               ▏24            ▕                                   
 Workers                                                              
Rows 1-2/2  Columns 1-2/2                                             
                                                                      
                                                                      
api / api                                                             
Observed: true                                                        
Replicas: 24
```

Demo · truecolor · 100 × 30 terminal cells

```
 Grid               API              Workers                api / api                               
 API               ▏24            ▕                         Observed: true                          
 Workers                                                    Replicas: 24                            
Rows 1-2/2  Columns 1-2/2
```

Narrow · truecolor · 40 × 16 terminal cells

```
 Grid               API                 
 API               ▏24            ▕     
 Workers                                
Rows 1-2/2  Columns 1-1/2               
                                        
                                        
                                        
                                        
api / api                               
Observed: true                          
Replicas: 24
```

256 colors · ansi256 · 100 × 30 terminal cells

```
 Grid               API              Workers                api / api                               
 API               ▏24            ▕                         Observed: true                          
 Workers                                                    Replicas: 24                            
Rows 1-2/2  Columns 1-2/2
```

16 colors · ansi16 · 100 × 30 terminal cells

```
 Grid               API              Workers                api / api                               
 API               ▏24            ▕                         Observed: true                          
 Workers                                                    Replicas: 24                            
Rows 1-2/2  Columns 1-2/2
```

Monochrome · none · 100 × 30 terminal cells

```
 Grid               API              Workers                api / api                               
 API               [24            ]                         Observed: true                          
 Workers                                                    Replicas: 24                            
Rows 1-2/2  Columns 1-2/2
```

### [Go usage](https://lanej.io/charmed/organisms-collection.html#collection%2fgrid-inspector-code)

Imports: [`fmt`](https://pkg.go.dev/fmt "fmt") · [`collection`](https://pkg.go.dev/github.com/lanej/charmed/collection "github.com/lanej/charmed/collection") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
size := layout.Size{Width: 70, Height: 12}
axes := []collection.GridAxis{{ID: "api", Label: "API"}, {ID: "worker", Label: "Workers"}}

grid, err := collection.NewDataGrid(axes, axes, []collection.GridEntry[int]{{RowID: "api", ColumnID: "api", Value: 24}}, size, theme.New(theme.Dark))
if err != nil {
	return "", err
}

inspector := collection.NewGridInspector(grid, func(e collection.GridEntry[int], present bool, _ int) string {
	return fmt.Sprintf("%s / %s\nObserved: %t\nReplicas: %d", e.RowID, e.ColumnID, present, e.Value)
}, size, theme.New(theme.Dark))
// Parent routes Tab using FocusDetail. Refresh after direct grid mutations.
return inspector.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-collection.html#collection%2fgrid-inspector-integration)

Supply an existing grid and a detail renderer that handles present=false. Route Tab in the parent with FocusDetail and return Update commands. SetData and Select refresh details automatically; call Refresh after direct grid mutations. Coordinate changes reset detail scrolling, while refreshes preserve it. Resize stacks panes below ninety columns and shows only the focused pane at very short heights.

### [Run this example](https://lanej.io/charmed/organisms-collection.html#collection%2fgrid-inspector-run)

```sh
go run ./examples/catalog collection/grid-inspector
```

[Example source](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/docs/snippets/collection_grid_inspector.go) · [Implementation](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/collection/grid_inspector.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

### Reusable component

[Grid inspector dependency graph](https://lanej.io/charmed/dependencies.html#component-collection%2Fgrid-inspector) · Partial map · [Source declaration](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/collection/grid_inspector.go#L50)

- [Data grid](https://lanej.io/charmed/dependencies.html#component-collection%2Fdata-grid) · composes · grid
- [Scroll pane](https://lanej.io/charmed/dependencies.html#component-collection%2Fscroll-pane) · composes · details

Consumer usage uses

- [Grid inspector](https://lanej.io/charmed/dependencies.html#component-collection%2Fgrid-inspector) · primary demonstration

[Example source](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/docs/snippets/collection_grid_inspector.go#L11)

[Organisms](https://lanej.io/charmed/organisms.html) / Collections

## [Keyed table](https://lanej.io/charmed/organisms-collection.html#collection%2fkeyed-table)

Keep a selected table row recognizable across filtering and reordering. Caller-owned stable keys remain independent of formatted native table cells and may include their zero value.

### [Native preview](https://lanej.io/charmed/organisms-collection.html#collection%2fkeyed-table-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 Service                   Replicas     
 Workers                   12           
```

Demo · truecolor · 100 × 30 terminal cells

```
 Service                   Replicas                                                                 
 API gateway               24                                                                       
 Workers                   12          
```

Narrow · truecolor · 40 × 16 terminal cells

```
 Service                   Replicas     
 API gateway               24           
 Workers                   12          
```

256 colors · ansi256 · 100 × 30 terminal cells

```
 Service                   Replicas                                                                 
 API gateway               24                                                                       
 Workers                   12          
```

16 colors · ansi16 · 100 × 30 terminal cells

```
 Service                   Replicas                                                                 
 API gateway               24                                                                       
 Workers                   12          
```

Monochrome · none · 100 × 30 terminal cells

```
 Service                   Replicas                                                                 
 API gateway               24                                                                       
 Workers                   12          
```

### [Go usage](https://lanej.io/charmed/organisms-collection.html#collection%2fkeyed-table-code)

Imports: [`table`](https://pkg.go.dev/charm.land/bubbles/v2/table "charm.land/bubbles/v2/table") · [`collection`](https://pkg.go.dev/github.com/lanej/charmed/collection "github.com/lanej/charmed/collection") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
columns := []table.Column{{Title: "Service", Width: 24}, {Title: "Replicas", Width: 12}}
native := collection.NewTable(columns, nil, 40, 5, theme.New(theme.Dark))
rows := []collection.KeyedRow[string]{{Key: "api", Row: table.Row{"API gateway", "24"}}, {Key: "worker", Row: table.Row{"Workers", "12"}}}

keyed, err := collection.NewKeyedTable(&native, rows)
if err != nil {
	return "", err
}

keyed.Select("worker")
// Reordering and filtering retain the selected key when it survives.
if err := keyed.SetRows([]collection.KeyedRow[string]{rows[1]}); err != nil {
	return "", err
}
// Forward Update commands; configure styles and focus on keyed.Model.
return keyed.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-collection.html#collection%2fkeyed-table-integration)

Bind the wrapper to an existing native table pointer. Replace rows through SetRows or columns and rows together through SetData; duplicate keys and mismatched shapes fail before mutation. Surviving keys retain selection and removed keys select the first remaining row. Configure native focus, dimensions, styles and bindings directly; forward Update commands and avoid replacing rows behind the wrapper.

### [Run this example](https://lanej.io/charmed/organisms-collection.html#collection%2fkeyed-table-run)

```sh
go run ./examples/catalog collection/keyed-table
```

[Example source](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/docs/snippets/collection_keyed_table.go) · [Implementation](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/collection/keyed_table.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [collection/table](https://lanej.io/charmed/atoms-collection.html#collection/table)

[Organisms](https://lanej.io/charmed/organisms.html) / Collections

## [Section stack](https://lanej.io/charmed/organisms-collection.html#collection%2fsection-stack)

Compose a single scrolling document from ordered sections with independent folds and stable identities. Bodies can describe any content; files are optional.

### [Native preview](https://lanej.io/charmed/organisms-collection.html#collection%2fsection-stack-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
▾ Intent                                                                
Retain the last healthy route while observations are incomplete.        
                                                                        
▾ Evidence  3 observations                                              
West: 42 ms                                                             
Central: 61 ms                                                          
East: not captured
```

Demo · truecolor · 120 × 32 terminal cells

```
▾ Routing proposal  2 sections                                                                                          
▾ Intent  Decision context                                                                                              
Choose a regional route using supplied latency observations.                                                            
Keep the last healthy route when new evidence is incomplete.                                                            
                                                                                                                        
▾ Evidence  3 observations                                                                                              
Region       p95       Samples                                                                                          
West         42 ms     12,400                                                                                           
Central      61 ms      8,100                                                                                           
East         unknown   not captured                                                                                     
                                                                                                                        
▾ Decision  Awaiting review                                                                                             
Retain the current route until the East observation arrives.                                                            
This document is application-supplied content, not a file tree.
```

Narrow · truecolor · 40 × 16 terminal cells

```
▾ Routing proposal  2 sections          
▾ Intent  Decision context              
Choose a regional route using supplied l
atency observations.                    
Keep the last healthy route when new evi
dence is incomplete.                    
                                        
▾ Evidence  3 observations              
Region       p95       Samples          
West         42 ms     12,400           
Central      61 ms      8,100           
East         unknown   not captured     
                                        
▾ Decision  Awaiting review             
Retain the current route until the East 
observation arrives.
```

256 colors · ansi256 · 120 × 32 terminal cells

```
▾ Routing proposal  2 sections                                                                                          
▾ Intent  Decision context                                                                                              
Choose a regional route using supplied latency observations.                                                            
Keep the last healthy route when new evidence is incomplete.                                                            
                                                                                                                        
▾ Evidence  3 observations                                                                                              
Region       p95       Samples                                                                                          
West         42 ms     12,400                                                                                           
Central      61 ms      8,100                                                                                           
East         unknown   not captured                                                                                     
                                                                                                                        
▾ Decision  Awaiting review                                                                                             
Retain the current route until the East observation arrives.                                                            
This document is application-supplied content, not a file tree.
```

16 colors · ansi16 · 120 × 32 terminal cells

```
▾ Routing proposal  2 sections                                                                                          
▾ Intent  Decision context                                                                                              
Choose a regional route using supplied latency observations.                                                            
Keep the last healthy route when new evidence is incomplete.                                                            
                                                                                                                        
▾ Evidence  3 observations                                                                                              
Region       p95       Samples                                                                                          
West         42 ms     12,400                                                                                           
Central      61 ms      8,100                                                                                           
East         unknown   not captured                                                                                     
                                                                                                                        
▾ Decision  Awaiting review                                                                                             
Retain the current route until the East observation arrives.                                                            
This document is application-supplied content, not a file tree.
```

Monochrome · none · 120 × 32 terminal cells

```
v Routing proposal  2 sections                                                                                          
v Intent  Decision context                                                                                              
Choose a regional route using supplied latency observations.                                                            
Keep the last healthy route when new evidence is incomplete.                                                            
                                                                                                                        
v Evidence  3 observations                                                                                              
Region       p95       Samples                                                                                          
West         42 ms     12,400                                                                                           
Central      61 ms      8,100                                                                                           
East         unknown   not captured                                                                                     
                                                                                                                        
v Decision  Awaiting review                                                                                             
Retain the current route until the East observation arrives.                                                            
This document is application-supplied content, not a file tree.
```

ASCII · truecolor · 120 × 32 terminal cells

```
v Routing proposal  2 sections                                                                                          
v Intent  Decision context                                                                                              
Choose a regional route using supplied latency observations.                                                            
Keep the last healthy route when new evidence is incomplete.                                                            
                                                                                                                        
v Evidence  3 observations                                                                                              
Region       p95       Samples                                                                                          
West         42 ms     12,400                                                                                           
Central      61 ms      8,100                                                                                           
East         unknown   not captured                                                                                     
                                                                                                                        
v Decision  Awaiting review                                                                                             
Retain the current route until the East observation arrives.                                                            
This document is application-supplied content, not a file tree.
```

Unicode monochrome · none · 120 × 32 terminal cells

```
v Routing proposal  2 sections                                                                                          
v Intent  Decision context                                                                                              
Choose a regional route using supplied latency observations.                                                            
Keep the last healthy route when new evidence is incomplete.                                                            
                                                                                                                        
v Evidence  3 observations                                                                                              
Region       p95       Samples                                                                                          
West         42 ms     12,400                                                                                           
Central      61 ms      8,100                                                                                           
East         unknown   not captured                                                                                     
                                                                                                                        
v Decision  Awaiting review                                                                                             
Retain the current route until the East observation arrives.                                                            
This document is application-supplied content, not a file tree.
```

### [Go usage](https://lanej.io/charmed/organisms-collection.html#collection%2fsection-stack-code)

Imports: [`collection`](https://pkg.go.dev/github.com/lanej/charmed/collection "github.com/lanej/charmed/collection") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
stack, err := collection.NewSectionStack([]collection.DocumentSection{
	{ID: "intent", Title: "Intent", Body: func(int) string { return "Retain the last healthy route while observations are incomplete." }},
	{ID: "evidence", Title: "Evidence", Summary: "3 observations", Body: func(int) string { return "West: 42 ms\nCentral: 61 ms\nEast: not captured" }},
}, layout.Size{Width: 72, Height: 12}, theme.New(theme.Dark))
if err != nil {
	return "", err
}

return stack.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-collection.html#collection%2fsection-stack-integration)

Supply ordered sections and width-aware body renderers, forward Update commands, and resize the existing stack. Native viewport styles and bindings survive; content, wrapping and scroll anchoring belong to the stack. Call Refresh when renderer-captured presentation changes. See diff-review.html for lifecycle and composition.

### [Run this example](https://lanej.io/charmed/organisms-collection.html#collection%2fsection-stack-run)

```sh
go run ./examples/catalog collection/section-stack
```

[Example source](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/docs/snippets/collection_section_stack.go) · [Implementation](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/collection/section_stack.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [collection/viewport](https://lanej.io/charmed/atoms-collection.html#collection/viewport) [text/disclosure](https://lanej.io/charmed/atoms-text.html#text/disclosure)

[Organisms](https://lanej.io/charmed/organisms.html) / Collections

## [Outline document](https://lanej.io/charmed/organisms-collection.html#pattern%2foutline-document)

Navigate arbitrary hierarchical content with a native tree linked to stacked sections. Selection, scrolling and subtree folding stay synchronized.

### [Native preview](https://lanej.io/charmed/organisms-collection.html#pattern%2foutline-document-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
┌ Outline ──────────────────────┐┌ Document ───────────────────────────────────────────────────────┐
│    ▾  Contents                ││ ▾ Routing proposal                                              │
│  ├ ▾  Routing proposal        ││ ▾ Intent                                                        │
│  │├  ▾ Intent                 ││ Keep the last healthy route.                                    │
│ →│└  ▾ Evidence               ││                                                                 │
│  └  ▾ Decision                ││ ▾ Evidence                                                      │
│                               ││ West: 42 ms                                                     │
│                               ││ East: observation missing                                       │
│                               ││                                                                 │
│                               ││ ▾ Decision                                                      │
│                               ││ Wait for complete evidence.                                     │
│                               ││                                                                 │
│                               ││                                                                 │
│                               ││                                                                 │
│                               ││                                                                 │
│                               ││                                                                 │
│                               ││                                                                 │
│                               ││                                                                 │
└───────────────────────────────┘└─────────────────────────────────────────────────────────────────┘
↑/↓ navigate · Enter fold · Tab pane
```

Demo · truecolor · 120 × 32 terminal cells

```
┌ Outline ─────────────────────────┐┌ Document ────────────────────────────────────────────────────────────────────────┐
│    ▾  Contents                   ││ ▾ Routing proposal  2 sections                                                   │
│ →├ ▾  Routing proposal  2 sect…  ││ ▾ Intent  Decision context                                                       │
│  │├  ▾ Intent  Decision cont…    ││ Choose a regional route using supplied latency observations.                     │
│  │└  ▾ Evidence  3 observati…    ││ Keep the last healthy route when new evidence is incomplete.                     │
│  └  ▾ Decision  Awaiting rev…    ││                                                                                  │
│                                  ││ ▾ Evidence  3 observations                                                       │
│                                  ││ Region       p95       Samples                                                   │
│                                  ││ West         42 ms     12,400                                                    │
│                                  ││ Central      61 ms      8,100                                                    │
│                                  ││ East         unknown   not captured                                              │
│                                  ││                                                                                  │
│                                  ││ ▾ Decision  Awaiting review                                                      │
│                                  ││ Retain the current route until the East observation arrives.                     │
│                                  ││ This document is application-supplied content, not a file tree.                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
└──────────────────────────────────┘└──────────────────────────────────────────────────────────────────────────────────┘
↑/↓ navigate · Enter fold · Tab pane
```

Narrow · truecolor · 40 × 16 terminal cells

```
┌ Outline ─────────────────────────────┐
│    ▾  Contents                       │
│ →├ ▾  Routing proposal  2 sections   │
│  │├  ▾ Intent  Decision context      │
└──────────────────────────────────────┘
┌ Document ────────────────────────────┐
│ ▾ Routing proposal  2 sections       │
│ ▾ Intent  Decision context           │
│ Choose a regional route using suppli │
│ ed latency observations.             │
│ Keep the last healthy route when new │
│  evidence is incomplete.             │
│                                      │
│ ▾ Evidence  3 observations           │
└──────────────────────────────────────┘
↑/↓ · Enter fold · Tab
```

256 colors · ansi256 · 120 × 32 terminal cells

```
┌ Outline ─────────────────────────┐┌ Document ────────────────────────────────────────────────────────────────────────┐
│    ▾  Contents                   ││ ▾ Routing proposal  2 sections                                                   │
│ →├ ▾  Routing proposal  2 sect…  ││ ▾ Intent  Decision context                                                       │
│  │├  ▾ Intent  Decision cont…    ││ Choose a regional route using supplied latency observations.                     │
│  │└  ▾ Evidence  3 observati…    ││ Keep the last healthy route when new evidence is incomplete.                     │
│  └  ▾ Decision  Awaiting rev…    ││                                                                                  │
│                                  ││ ▾ Evidence  3 observations                                                       │
│                                  ││ Region       p95       Samples                                                   │
│                                  ││ West         42 ms     12,400                                                    │
│                                  ││ Central      61 ms      8,100                                                    │
│                                  ││ East         unknown   not captured                                              │
│                                  ││                                                                                  │
│                                  ││ ▾ Decision  Awaiting review                                                      │
│                                  ││ Retain the current route until the East observation arrives.                     │
│                                  ││ This document is application-supplied content, not a file tree.                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
└──────────────────────────────────┘└──────────────────────────────────────────────────────────────────────────────────┘
↑/↓ navigate · Enter fold · Tab pane
```

16 colors · ansi16 · 120 × 32 terminal cells

```
┌ Outline ─────────────────────────┐┌ Document ────────────────────────────────────────────────────────────────────────┐
│    ▾  Contents                   ││ ▾ Routing proposal  2 sections                                                   │
│ →├ ▾  Routing proposal  2 sect…  ││ ▾ Intent  Decision context                                                       │
│  │├  ▾ Intent  Decision cont…    ││ Choose a regional route using supplied latency observations.                     │
│  │└  ▾ Evidence  3 observati…    ││ Keep the last healthy route when new evidence is incomplete.                     │
│  └  ▾ Decision  Awaiting rev…    ││                                                                                  │
│                                  ││ ▾ Evidence  3 observations                                                       │
│                                  ││ Region       p95       Samples                                                   │
│                                  ││ West         42 ms     12,400                                                    │
│                                  ││ Central      61 ms      8,100                                                    │
│                                  ││ East         unknown   not captured                                              │
│                                  ││                                                                                  │
│                                  ││ ▾ Decision  Awaiting review                                                      │
│                                  ││ Retain the current route until the East observation arrives.                     │
│                                  ││ This document is application-supplied content, not a file tree.                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
└──────────────────────────────────┘└──────────────────────────────────────────────────────────────────────────────────┘
↑/↓ navigate · Enter fold · Tab pane
```

Monochrome · none · 120 × 32 terminal cells

```
+ Outline -------------------------++ Document ------------------------------------------------------------------------+
|    v  Contents                   || v Routing proposal  2 sections                                                   |
| >+ v  Routing proposal  2 sect.  || v Intent  Decision context                                                       |
|  |+  v Intent  Decision cont.    || Choose a regional route using supplied latency observations.                     |
|  |+  v Evidence  3 observati.    || Keep the last healthy route when new evidence is incomplete.                     |
|  +  v Decision  Awaiting rev.    ||                                                                                  |
|                                  || v Evidence  3 observations                                                       |
|                                  || Region       p95       Samples                                                   |
|                                  || West         42 ms     12,400                                                    |
|                                  || Central      61 ms      8,100                                                    |
|                                  || East         unknown   not captured                                              |
|                                  ||                                                                                  |
|                                  || v Decision  Awaiting review                                                      |
|                                  || Retain the current route until the East observation arrives.                     |
|                                  || This document is application-supplied content, not a file tree.                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
+----------------------------------++----------------------------------------------------------------------------------+
^/v navigate * Enter fold * Tab pane
```

Folded · truecolor · 120 × 32 terminal cells

```
┌ Outline ─────────────────────────┐┌ Document ────────────────────────────────────────────────────────────────────────┐
│    ▾  Contents                   ││ ▸ Routing proposal  2 sections                                                   │
│ →├ ▸  Routing proposal  2 sect…  ││ ▾ Decision  Awaiting review                                                      │
│  └  ▾ Decision  Awaiting rev…    ││ Retain the current route until the East observation arrives.                     │
│                                  ││ This document is application-supplied content, not a file tree.                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
└──────────────────────────────────┘└──────────────────────────────────────────────────────────────────────────────────┘
↑/↓ navigate · Enter fold · Tab pane
```

Empty · truecolor · 120 × 32 terminal cells

```
┌ Outline ─────────────────────────┐┌ Document ────────────────────────────────────────────────────────────────────────┐
│ →  ▾  Contents                   ││ No sections                                                                      │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
│                                  ││                                                                                  │
└──────────────────────────────────┘└──────────────────────────────────────────────────────────────────────────────────┘
↑/↓ navigate · Enter fold · Tab pane
```

ASCII · truecolor · 120 × 32 terminal cells

```
+ Outline -------------------------++ Document ------------------------------------------------------------------------+
|    v  Contents                   || v Routing proposal  2 sections                                                   |
| >+ v  Routing proposal  2 sect.  || v Intent  Decision context                                                       |
|  |+  v Intent  Decision cont.    || Choose a regional route using supplied latency observations.                     |
|  |+  v Evidence  3 observati.    || Keep the last healthy route when new evidence is incomplete.                     |
|  +  v Decision  Awaiting rev.    ||                                                                                  |
|                                  || v Evidence  3 observations                                                       |
|                                  || Region       p95       Samples                                                   |
|                                  || West         42 ms     12,400                                                    |
|                                  || Central      61 ms      8,100                                                    |
|                                  || East         unknown   not captured                                              |
|                                  ||                                                                                  |
|                                  || v Decision  Awaiting review                                                      |
|                                  || Retain the current route until the East observation arrives.                     |
|                                  || This document is application-supplied content, not a file tree.                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
+----------------------------------++----------------------------------------------------------------------------------+
^/v navigate * Enter fold * Tab pane
```

Unicode monochrome · none · 120 × 32 terminal cells

```
+ Outline -------------------------++ Document ------------------------------------------------------------------------+
|    v  Contents                   || v Routing proposal  2 sections                                                   |
| >+ v  Routing proposal  2 sect.  || v Intent  Decision context                                                       |
|  |+  v Intent  Decision cont.    || Choose a regional route using supplied latency observations.                     |
|  |+  v Evidence  3 observati.    || Keep the last healthy route when new evidence is incomplete.                     |
|  +  v Decision  Awaiting rev.    ||                                                                                  |
|                                  || v Evidence  3 observations                                                       |
|                                  || Region       p95       Samples                                                   |
|                                  || West         42 ms     12,400                                                    |
|                                  || Central      61 ms      8,100                                                    |
|                                  || East         unknown   not captured                                              |
|                                  ||                                                                                  |
|                                  || v Decision  Awaiting review                                                      |
|                                  || Retain the current route until the East observation arrives.                     |
|                                  || This document is application-supplied content, not a file tree.                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
|                                  ||                                                                                  |
+----------------------------------++----------------------------------------------------------------------------------+
^/v navigate * Enter fold * Tab pane
```

### [Go usage](https://lanej.io/charmed/organisms-collection.html#pattern%2foutline-document-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`pattern`](https://pkg.go.dev/github.com/lanej/charmed/pattern "github.com/lanej/charmed/pattern") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
document, err := pattern.NewOutlineDocument([]pattern.OutlineEntry{
	{ID: "proposal", Title: "Routing proposal", Children: []pattern.OutlineEntry{
		{ID: "intent", Title: "Intent", Body: func(int) string { return "Keep the last healthy route." }},
		{ID: "evidence", Title: "Evidence", Body: func(int) string { return "West: 42 ms\nEast: observation missing" }},
	}},
	{ID: "decision", Title: "Decision", Body: func(int) string { return "Wait for complete evidence." }},
}, layout.Size{Width: 100, Height: 20}, theme.New(theme.Dark))
if err != nil {
	return "", err
}

document.Select("evidence")

return document.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-collection.html#pattern%2foutline-document-integration)

Supply stable unique IDs and width-aware body renderers. Return Update commands, route Tab in the parent through SetFocus, and call SetOrigin before forwarding screen-coordinate mouse events. Native tree styles and bindings survive replacement and resizing; nodes are rebuilt. The stack owns viewport content and vertical anchors. Run go run ./examples/review -document for a complete non-file consumer, and see diff-review.html for lifecycle details.

### [Run this example](https://lanej.io/charmed/organisms-collection.html#pattern%2foutline-document-run)

```sh
go run ./examples/catalog pattern/outline-document
```

[Example source](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/docs/snippets/pattern_outline_document.go) · [Implementation](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/pattern/outline_document.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [collection/tree](https://lanej.io/charmed/atoms-collection.html#collection/tree) [collection/section-stack](https://lanej.io/charmed/organisms-collection.html#collection/section-stack)
