# Component dependencies

Explore how reusable components are built and where they are reused. Component relationships remain separate from consumer examples and interactive fixtures, so demo controls do not become dependencies of the component they demonstrate.

This first audited slice covers the grid, status-cell, semantic state, scroll-pane and replay families. Maps are explicitly partial: an empty relationship list does not prove that a component is a leaf. Atomic levels describe classification rather than graph depth.

**20 reusable components**20 of 162 catalog entries mapped · partial audited slice[Download graph JSON ↗](https://lanej.io/charmed/component-graph.json)

## Canary membership

Partial map

API

`feedback.CanaryMembership`

Classification

Atoms

Evidence

[feedback/canary\_membership.go · CanaryMembership](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/feedback/canary_membership.go#L11)

### Native example

[Go usage ↗](https://lanej.io/charmed/atoms-feedback.html#feedback%2Fcanary-membership)

Saved output of the compiled consumer example

```
Pill                 Compact
                            
 🐤🟢 Canary on      ◉      
 🐤⚪️ Canary off     ◌      
 ? Unknown           ?
```

### Built from

- [Indicator](https://lanej.io/charmed/dependencies.html#component-text%2Findicator)composes · pool membership presentation[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/feedback/canary_membership.go#L21)

### Used by

- No incoming component relationships recorded in this slice.

### Direct example usage

- [Atoms/Feedback/Canary membership · Consumer usage](https://lanej.io/charmed/atoms-feedback.html#feedback%2Fcanary-membership)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/feedback_canary_membership.go#L10)
  
  - [Canary membership](https://lanej.io/charmed/dependencies.html#component-feedback%2Fcanary-membership)primary demonstration

## Indicator

Partial map

API

`text.Indicator`

Classification

Atoms

Evidence

[text/text.go · Indicator](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/text/text.go#L60)

### Native example

[Go usage ↗](https://lanej.io/charmed/atoms-text.html#text%2Findicator)

Saved output of the compiled consumer example

```
 ↑ Up 
 ↑    
 Up   
↑
```

### Built from

- [Color pill](https://lanej.io/charmed/dependencies.html#component-text%2Fpill)composes · labeled pill presentation[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/text/text.go#L60)

### Used by

- [Canary membership](https://lanej.io/charmed/dependencies.html#component-feedback%2Fcanary-membership)composes · pool membership presentation[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/feedback/canary_membership.go#L21)
- [Status](https://lanej.io/charmed/dependencies.html#component-feedback%2Fstatus)composes · state presentation[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/feedback/status.go#L45)
- [Availability](https://lanej.io/charmed/dependencies.html#component-feedback%2Favailability)composes · state presentation[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/feedback/availability.go#L18)
- [Health](https://lanej.io/charmed/dependencies.html#component-feedback%2Fhealth)composes · state presentation[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/feedback/health.go#L18)
- [Readiness](https://lanej.io/charmed/dependencies.html#component-feedback%2Freadiness)composes · state presentation[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/feedback/readiness.go#L18)
- [Rollout](https://lanej.io/charmed/dependencies.html#component-feedback%2Frollout)composes · state presentation[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/feedback/rollout.go#L19)

### Direct example usage

- [Atoms/Text/Indicator · Consumer usage](https://lanej.io/charmed/atoms-text.html#text%2Findicator)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/text_indicator.go#L9)
  
  - [Indicator](https://lanej.io/charmed/dependencies.html#component-text%2Findicator)primary demonstration

## Color pill

Partial map

API

`text.Pill`

Classification

Atoms

Evidence

[text/text.go · Pill](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/text/text.go#L29)

### Native example

[Go usage ↗](https://lanej.io/charmed/atoms-text.html#text%2Fpill)

Saved output of the compiled consumer example

```
 Ready       
 Warning     
 Unavailable 
```

### Built from

- No relationships recorded in this partial map.

### Used by

- [Indicator](https://lanej.io/charmed/dependencies.html#component-text%2Findicator)composes · labeled pill presentation[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/text/text.go#L60)

### Direct example usage

- [Atoms/Text/Color pill · Consumer usage](https://lanej.io/charmed/atoms-text.html#text%2Fpill)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/text_pill.go#L9)
  
  - [Color pill](https://lanej.io/charmed/dependencies.html#component-text%2Fpill)primary demonstration

## Availability

Partial map

API

`feedback.Availability`

Classification

Atoms

Evidence

[feedback/availability.go · Availability](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/feedback/availability.go#L9)

### Native example

[Go usage ↗](https://lanej.io/charmed/atoms-feedback.html#feedback%2Favailability)

Saved output of the compiled consumer example

```
Pill           Compact
                      
 ↑ Up          ↑      
 ↓ Down        ↓      
 ? Unknown     ?
```

### Built from

- [Indicator](https://lanej.io/charmed/dependencies.html#component-text%2Findicator)composes · state presentation[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/feedback/availability.go#L18)

### Used by

- No incoming component relationships recorded in this slice.

### Direct example usage

- [Atoms/Feedback/Availability · Consumer usage](https://lanej.io/charmed/atoms-feedback.html#feedback%2Favailability)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/feedback_availability.go#L10)
  
  - [Availability](https://lanej.io/charmed/dependencies.html#component-feedback%2Favailability)primary demonstration

## Health

Partial map

API

`feedback.Health`

Classification

Atoms

Evidence

[feedback/health.go · Health](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/feedback/health.go#L9)

### Native example

[Go usage ↗](https://lanej.io/charmed/atoms-feedback.html#feedback%2Fhealth)

Saved output of the compiled consumer example

```
Pill             Compact
                        
 ♥ Healthy       ♥      
 × Unhealthy     ×      
 ? Unknown       ?
```

### Built from

- [Indicator](https://lanej.io/charmed/dependencies.html#component-text%2Findicator)composes · state presentation[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/feedback/health.go#L18)

### Used by

- No incoming component relationships recorded in this slice.

### Direct example usage

- [Atoms/Feedback/Health · Consumer usage](https://lanej.io/charmed/atoms-feedback.html#feedback%2Fhealth)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/feedback_health.go#L10)
  
  - [Health](https://lanej.io/charmed/dependencies.html#component-feedback%2Fhealth)primary demonstration

## Readiness

Partial map

API

`feedback.Readiness`

Classification

Atoms

Evidence

[feedback/readiness.go · Readiness](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/feedback/readiness.go#L9)

### Native example

[Go usage ↗](https://lanej.io/charmed/atoms-feedback.html#feedback%2Freadiness)

Saved output of the compiled consumer example

```
Pill             Compact
                        
 ✓ Ready         ✓      
 ◇ Not ready     ◇      
 ? Unknown       ?
```

### Built from

- [Indicator](https://lanej.io/charmed/dependencies.html#component-text%2Findicator)composes · state presentation[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/feedback/readiness.go#L18)

### Used by

- No incoming component relationships recorded in this slice.

### Direct example usage

- [Atoms/Feedback/Readiness · Consumer usage](https://lanej.io/charmed/atoms-feedback.html#feedback%2Freadiness)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/feedback_readiness.go#L10)
  
  - [Readiness](https://lanej.io/charmed/dependencies.html#component-feedback%2Freadiness)primary demonstration

## Rollout

Partial map

API

`feedback.Rollout`

Classification

Atoms

Evidence

[feedback/rollout.go · Rollout](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/feedback/rollout.go#L10)

### Native example

[Go usage ↗](https://lanej.io/charmed/atoms-feedback.html#feedback%2Frollout)

Saved output of the compiled consumer example

```
Pill           Compact
                      
 🐤 Canary     ◔      
 ● Live        ●      
 ? Unknown     ?
```

### Built from

- [Indicator](https://lanej.io/charmed/dependencies.html#component-text%2Findicator)composes · state presentation[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/feedback/rollout.go#L19)

### Used by

- No incoming component relationships recorded in this slice.

### Direct example usage

- [Atoms/Feedback/Rollout · Consumer usage](https://lanej.io/charmed/atoms-feedback.html#feedback%2Frollout)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/feedback_rollout.go#L10)
  
  - [Rollout](https://lanej.io/charmed/dependencies.html#component-feedback%2Frollout)primary demonstration

## Data grid

Partial map

API

`collection.DataGrid`

Classification

Organisms

Evidence

[collection/data\_grid.go · DataGrid](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/collection/data_grid.go#L44)

### Native example

[Go usage ↗](https://lanej.io/charmed/organisms-collection.html#collection%2Fdata-grid)

Saved output of the compiled consumer example

```
 Service / Region   US East          US West          
 API               ▏24            ▕  0                
 Workers            48                                
Rows 1-2/2  Columns 1-2/2
```

### Built from

- No relationships recorded in this partial map.

### Used by

- [Status matrix](https://lanej.io/charmed/dependencies.html#component-collection%2Fstatus-matrix)specializes · grid base[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/collection/status_matrix.go#L23)
- [Grid inspector](https://lanej.io/charmed/dependencies.html#component-collection%2Fgrid-inspector)composes · grid[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/collection/grid_inspector.go#L51)

### Direct example usage

- [Organisms/Collections/Data grid · Consumer usage](https://lanej.io/charmed/organisms-collection.html#collection%2Fdata-grid)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/collection_data_grid.go#L9)
  
  - [Data grid](https://lanej.io/charmed/dependencies.html#component-collection%2Fdata-grid)primary demonstration

## Status matrix

Partial map

API

`collection.StatusMatrix`

Classification

Organisms

Evidence

[collection/status\_matrix.go · StatusMatrix](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/collection/status_matrix.go#L23)

### Native example

[Go usage ↗](https://lanej.io/charmed/organisms-collection.html#collection%2Fstatus-matrix)

Saved output of the compiled consumer example

```
 Balancer / Host    ▤ host-01        ▤ host-02        ▤ host-03       
 ⇄ lb-public       ▏✓ v2.8.1      ▕  ! draining       ✓ v2.8.0        
 ⇄ lb-internal      —                × unhealthy      ? no probe      
Rows 1-2/2  Columns 1-3/3
```

### Built from

- [Data grid](https://lanej.io/charmed/dependencies.html#component-collection%2Fdata-grid)specializes · grid base[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/collection/status_matrix.go#L23)
- [Status cell](https://lanej.io/charmed/dependencies.html#component-feedback%2Fstatus-cell)composes · cell presentation[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/collection/status_matrix.go#L65)

### Used by

- No incoming component relationships recorded in this slice.

### Direct example usage

- [Organisms/Collections/Status matrix · Consumer usage](https://lanej.io/charmed/organisms-collection.html#collection%2Fstatus-matrix)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/collection_status_matrix.go#L10)
  
  - [Status matrix](https://lanej.io/charmed/dependencies.html#component-collection%2Fstatus-matrix)primary demonstration

## Status cell

Partial map

API

`feedback.StatusCell`

Classification

Molecules

Evidence

[feedback/status\_cell.go · StatusCell](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/feedback/status_cell.go#L11)

### Native example

[Go usage ↗](https://lanej.io/charmed/molecules-feedback.html#feedback%2Fstatus-cell)

Saved output of the compiled consumer example

```
 ✓ v2.8.1         
▏! draining      ▕
 × unhealthy      
 ? no probe       
```

### Built from

- [Status](https://lanej.io/charmed/dependencies.html#component-feedback%2Fstatus)composes · health presentation[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/feedback/status_cell.go#L18)

### Used by

- [Status matrix](https://lanej.io/charmed/dependencies.html#component-collection%2Fstatus-matrix)composes · cell presentation[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/collection/status_matrix.go#L65)

### Direct example usage

- [Molecules/Feedback/Status cell · Consumer usage](https://lanej.io/charmed/molecules-feedback.html#feedback%2Fstatus-cell)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/feedback_status_cell.go#L9)
  
  - [Status cell](https://lanej.io/charmed/dependencies.html#component-feedback%2Fstatus-cell)primary demonstration

## Status

Partial map

API

`feedback.Status`

Classification

Atoms

Evidence

[feedback/status.go · Status](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/feedback/status.go#L12)

### Native example

[Go usage ↗](https://lanej.io/charmed/atoms-feedback.html#feedback%2Fstatus)

Saved output of the compiled consumer example

```
 ✓ OK      
 ! WARN    
 ▶ RUNNING 
 ? UNKNOWN 
```

### Built from

- [Indicator](https://lanej.io/charmed/dependencies.html#component-text%2Findicator)composes · state presentation[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/feedback/status.go#L45)

### Used by

- [Status cell](https://lanej.io/charmed/dependencies.html#component-feedback%2Fstatus-cell)composes · health presentation[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/feedback/status_cell.go#L18)

### Direct example usage

- [Atoms/Feedback/Status · Consumer usage](https://lanej.io/charmed/atoms-feedback.html#feedback%2Fstatus)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/feedback_status.go#L10)
  
  - [Status](https://lanej.io/charmed/dependencies.html#component-feedback%2Fstatus)primary demonstration

## Grid inspector

Partial map

API

`collection.GridInspector`

Classification

Organisms

Evidence

[collection/grid\_inspector.go · GridInspector](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/collection/grid_inspector.go#L50)

### Native example

[Go usage ↗](https://lanej.io/charmed/organisms-collection.html#collection%2Fgrid-inspector)

Saved output of the compiled consumer example

```
 Grid               API              Workers                          
 API               ▏24            ▕                                   
 Workers                                                              
Rows 1-2/2  Columns 1-2/2                                             
                                                                      
                                                                      
api / api                                                             
Observed: true                                                        
Replicas: 24
```

### Built from

- [Data grid](https://lanej.io/charmed/dependencies.html#component-collection%2Fdata-grid)composes · grid[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/collection/grid_inspector.go#L51)
- [Scroll pane](https://lanej.io/charmed/dependencies.html#component-collection%2Fscroll-pane)composes · details[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/collection/grid_inspector.go#L52)

### Used by

- [Measurement history](https://lanej.io/charmed/dependencies.html#component-pattern%2Fmeasurement-history)composes · inspection[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/pattern/measurement_history.go#L37)

### Direct example usage

- [Organisms/Collections/Grid inspector · Consumer usage](https://lanej.io/charmed/organisms-collection.html#collection%2Fgrid-inspector)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/collection_grid_inspector.go#L11)
  
  - [Grid inspector](https://lanej.io/charmed/dependencies.html#component-collection%2Fgrid-inspector)primary demonstration

## Scroll pane

Partial map

API

`collection.ScrollPane`

Classification

Molecules

Evidence

[collection/scroll\_pane.go · ScrollPane](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/collection/scroll_pane.go#L14)

### Native example

[Go usage ↗](https://lanej.io/charmed/molecules-collection.html#collection%2Fscroll-pane)

Saved output of the compiled consumer example

```
Request observations remain scrolla│
ble.                               ┃
Request observations remain scrolla│
ble.                               │
Request observations remain scrolla│
ble.                               │
```

### Built from

- [Viewport](https://lanej.io/charmed/dependencies.html#component-collection%2Fviewport)composes · content viewport[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/collection/scroll_pane.go#L15)
- [Scrollbar](https://lanej.io/charmed/dependencies.html#component-feedback%2Fscrollbar)composes · scroll position[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/collection/scroll_pane.go#L59)

### Used by

- [Grid inspector](https://lanej.io/charmed/dependencies.html#component-collection%2Fgrid-inspector)composes · details[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/collection/grid_inspector.go#L52)

### Direct example usage

- [Molecules/Collections/Scroll pane · Consumer usage](https://lanej.io/charmed/molecules-collection.html#collection%2Fscroll-pane)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/collection_scroll_pane.go#L11)
  
  - [Scroll pane](https://lanej.io/charmed/dependencies.html#component-collection%2Fscroll-pane)primary demonstration

## Viewport

Partial map

API

`collection.NewViewport`

Classification

Atoms

Evidence

[collection/collection.go · NewViewport](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/collection/collection.go#L36)

### Native example

[Go usage ↗](https://lanej.io/charmed/atoms-collection.html#collection%2Fviewport)

Saved output of the compiled consumer example

```
API gateway: healthy                            
Event stream: healthy                           
Job workers: degraded
```

### Built from

- No relationships recorded in this partial map.

### Used by

- [Scroll pane](https://lanej.io/charmed/dependencies.html#component-collection%2Fscroll-pane)composes · content viewport[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/collection/scroll_pane.go#L15)

### Direct example usage

- [Atoms/Collections/Viewport · Consumer usage](https://lanej.io/charmed/atoms-collection.html#collection%2Fviewport)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/collection_viewport.go#L7)
  
  - [Viewport](https://lanej.io/charmed/dependencies.html#component-collection%2Fviewport)primary demonstration

## Scrollbar

Partial map

API

`feedback.Scrollbar`

Classification

Atoms

Evidence

[feedback/controls.go · Scrollbar](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/feedback/controls.go#L22)

### Native example

[Go usage ↗](https://lanej.io/charmed/atoms-feedback.html#feedback%2Fscrollbar)

Saved output of the compiled consumer example

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

### Built from

- No relationships recorded in this partial map.

### Used by

- [Scroll pane](https://lanej.io/charmed/dependencies.html#component-collection%2Fscroll-pane)composes · scroll position[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/collection/scroll_pane.go#L59)

### Direct example usage

- [Atoms/Feedback/Scrollbar · Consumer usage](https://lanej.io/charmed/atoms-feedback.html#feedback%2Fscrollbar)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/feedback_scrollbar.go#L8)
  
  - [Scrollbar](https://lanej.io/charmed/dependencies.html#component-feedback%2Fscrollbar)primary demonstration

## Measurement history

Partial map

API

`pattern.MeasurementHistory`

Classification

Organisms

Evidence

[pattern/measurement\_history.go · MeasurementHistory](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/pattern/measurement_history.go#L36)

### Native example

[Go usage ↗](https://lanej.io/charmed/organisms-pattern.html#pattern%2Fmeasurement-history)

Saved output of the compiled consumer example

```
 Metric / Window      02:00:00       03:00:00       04:00:00       05:00:00     
 Temperature         Metric absent  Metric absent  Metric absent   Metric absent
 Energy              25kWh          26kWh          27kWh          28kWh         
 Yield               85%            86%            87%            88%           
 Humidity            39%            40%            41%            42%           
Rows 1-4/4  Columns 3-6/6                                                       
                                                                                
                                                                                
west / sensor-4                                                                 
Temperature                                                                     
2026-01-01T05:00:00Z – 2026-01-01T06:00:00Z                                     
Metric absent
```

### Built from

- [Grid inspector](https://lanej.io/charmed/dependencies.html#component-collection%2Fgrid-inspector)composes · inspection[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/pattern/measurement_history.go#L37)

### Used by

- No incoming component relationships recorded in this slice.

### Direct example usage

- [Organisms/Patterns/Measurement history · Consumer usage](https://lanej.io/charmed/organisms-pattern.html#pattern%2Fmeasurement-history)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/pattern_measurement_history.go#L11)
  
  - [Measurement history](https://lanej.io/charmed/dependencies.html#component-pattern%2Fmeasurement-history)primary demonstration

## State diagram

Partial map

API

`diagram.StateDiagram`

Classification

Molecules

Evidence

[diagram/state.go · StateDiagram](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/diagram/state.go#L36)

### Native example

[Go usage ↗](https://lanej.io/charmed/molecules-diagram.html#diagram%2Fstate-transitions)

Saved output of the compiled consumer example

```
╭───────────────────╮                 
│ Idle              │                 
│ initial           │                 
╰┬────────┬─────────╯                 
 │        ▲                           
 │ start  │ cancel                    
 ▼        │                           
┌┴────────┴─────────┐                 
│ Running           ├───────────┐     
│                   │           │     
│                   ├─┐         │     
│                   │ │ retry   │     
│                   │ │         │     
│                   │ │         │     
│                   ├◀┘         │     
└┬──────────────────┘           │     
 ┃                              │     
 ┃ finish                       │ fail
 ▼                              │     
╔╧══════════════════╗           │     
║ Done              ║           │     
║ success, active   ║           │     
╚═══════════════════╝           │     
                                │     
                                │     
                                │     
╔═══════════════════╗           │     
║ Failed            ╟◀──────────┘     
║ error             ║                 
║                   ║                 
╚═══════════════════╝
```

### Built from

- No relationships recorded in this partial map.

### Used by

- [Trace replay](https://lanej.io/charmed/dependencies.html#component-diagram%2Ftrace-replay)composes · diagram[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/diagram/trace_replay.go#L25)

### Direct example usage

- [Molecules/Diagrams/State transitions · Consumer usage](https://lanej.io/charmed/molecules-diagram.html#diagram%2Fstate-transitions)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/diagram_state_transitions.go#L8)
  
  - [State diagram](https://lanej.io/charmed/dependencies.html#component-diagram%2Fstate-diagram)primary demonstration
- [Organisms/Diagrams/Trace replay · Consumer usage](https://lanej.io/charmed/organisms-diagram.html#diagram%2Ftrace-replay)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/diagram_trace_replay.go#L9)
  
  - [Trace replay](https://lanej.io/charmed/dependencies.html#component-diagram%2Ftrace-replay)primary demonstration
  - [State diagram](https://lanej.io/charmed/dependencies.html#component-diagram%2Fstate-diagram)graph data
- [Organisms/Diagrams/Trace replay · Interactive fixture](https://lanej.io/charmed/organisms-diagram.html#panel-diagram%2Ftrace-replay-0)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/story/diagrams.go#L12)
  
  - [State diagram](https://lanej.io/charmed/dependencies.html#component-diagram%2Fstate-diagram)graph data
  - [Trace replay](https://lanej.io/charmed/dependencies.html#component-diagram%2Ftrace-replay)trace inspection
  - [Radio group](https://lanej.io/charmed/dependencies.html#component-input%2Fradio-group)history-window controls
- [Molecules/Diagrams/State transitions · Interactive fixture](https://lanej.io/charmed/molecules-diagram.html#panel-diagram%2Fstate-transitions-0)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/story/diagrams.go#L12)
  
  - [State diagram](https://lanej.io/charmed/dependencies.html#component-diagram%2Fstate-diagram)graph data
  - [Trace replay](https://lanej.io/charmed/dependencies.html#component-diagram%2Ftrace-replay)trace inspection
  - [Radio group](https://lanej.io/charmed/dependencies.html#component-input%2Fradio-group)history-window controls

## Trace replay

Partial map

API

`diagram.TraceReplay`

Classification

Organisms

Evidence

[diagram/trace\_replay.go · TraceReplay](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/diagram/trace_replay.go#L21)

### Native example

[Go usage ↗](https://lanej.io/charmed/organisms-diagram.html#diagram%2Ftrace-replay)

Saved output of the compiled consumer example

```
╭─────────────────╮            ┏━━━━━━━━━┓            ╔═══════════════╗
│ idle            ├━━ start ━━▶┫ running ┣── finish ─▶╢ done          ║
│ initial         ├◀─ cancel ──┫ active  ┃            ║ final         ║
╰─────────────────╯            ┗━━━━━━━━━┛            ╚═══════════════╝
```

### Built from

- [State diagram](https://lanej.io/charmed/dependencies.html#component-diagram%2Fstate-diagram)composes · diagram[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/diagram/trace_replay.go#L25)

### Used by

- No incoming component relationships recorded in this slice.

### Direct example usage

- [Organisms/Diagrams/Trace replay · Consumer usage](https://lanej.io/charmed/organisms-diagram.html#diagram%2Ftrace-replay)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/diagram_trace_replay.go#L9)
  
  - [Trace replay](https://lanej.io/charmed/dependencies.html#component-diagram%2Ftrace-replay)primary demonstration
  - [State diagram](https://lanej.io/charmed/dependencies.html#component-diagram%2Fstate-diagram)graph data
- [Organisms/Diagrams/Trace replay · Interactive fixture](https://lanej.io/charmed/organisms-diagram.html#panel-diagram%2Ftrace-replay-0)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/story/diagrams.go#L12)
  
  - [State diagram](https://lanej.io/charmed/dependencies.html#component-diagram%2Fstate-diagram)graph data
  - [Trace replay](https://lanej.io/charmed/dependencies.html#component-diagram%2Ftrace-replay)trace inspection
  - [Radio group](https://lanej.io/charmed/dependencies.html#component-input%2Fradio-group)history-window controls
- [Molecules/Diagrams/State transitions · Interactive fixture](https://lanej.io/charmed/molecules-diagram.html#panel-diagram%2Fstate-transitions-0)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/story/diagrams.go#L12)
  
  - [State diagram](https://lanej.io/charmed/dependencies.html#component-diagram%2Fstate-diagram)graph data
  - [Trace replay](https://lanej.io/charmed/dependencies.html#component-diagram%2Ftrace-replay)trace inspection
  - [Radio group](https://lanej.io/charmed/dependencies.html#component-input%2Fradio-group)history-window controls

## Radio group

Partial map

API

`input.RadioGroup`

Classification

Molecules

Evidence

[input/radio\_group.go · RadioGroup](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/input/radio_group.go#L20)

### Native example

[Go usage ↗](https://lanej.io/charmed/molecules-input.html#input%2Fradio-group)

Saved output of the compiled consumer example

```
  ◉  Last transition 
  ○  Last three transitions 
  ○  Complete trace 
  ○  Live execution ⊘ disabled Replay is inspection only
```

### Built from

- [Radio](https://lanej.io/charmed/dependencies.html#component-input%2Fradio)composes · options[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/input/radio_group.go#L21)

### Used by

- No incoming component relationships recorded in this slice.

### Direct example usage

- [Molecules/Input/Radio group · Consumer usage](https://lanej.io/charmed/molecules-input.html#input%2Fradio-group)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/input_radio_group.go#L8)
  
  - [Radio group](https://lanej.io/charmed/dependencies.html#component-input%2Fradio-group)primary demonstration
- [Organisms/Diagrams/Trace replay · Interactive fixture](https://lanej.io/charmed/organisms-diagram.html#panel-diagram%2Ftrace-replay-0)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/story/diagrams.go#L12)
  
  - [State diagram](https://lanej.io/charmed/dependencies.html#component-diagram%2Fstate-diagram)graph data
  - [Trace replay](https://lanej.io/charmed/dependencies.html#component-diagram%2Ftrace-replay)trace inspection
  - [Radio group](https://lanej.io/charmed/dependencies.html#component-input%2Fradio-group)history-window controls
- [Molecules/Diagrams/State transitions · Interactive fixture](https://lanej.io/charmed/molecules-diagram.html#panel-diagram%2Fstate-transitions-0)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/story/diagrams.go#L12)
  
  - [State diagram](https://lanej.io/charmed/dependencies.html#component-diagram%2Fstate-diagram)graph data
  - [Trace replay](https://lanej.io/charmed/dependencies.html#component-diagram%2Ftrace-replay)trace inspection
  - [Radio group](https://lanej.io/charmed/dependencies.html#component-input%2Fradio-group)history-window controls

## Radio

Partial map

API

`input.Radio`

Classification

Atoms

Evidence

[input/controls.go · Radio](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/input/controls.go#L157)

### Native example

[Go usage ↗](https://lanej.io/charmed/atoms-input.html#input%2Fradio)

Saved output of the compiled consumer example

```
  ◉  Production     ○  Staging
```

### Built from

- No relationships recorded in this partial map.

### Used by

- [Radio group](https://lanej.io/charmed/dependencies.html#component-input%2Fradio-group)composes · options[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/input/radio_group.go#L21)

### Direct example usage

- [Atoms/Input/Radio · Consumer usage](https://lanej.io/charmed/atoms-input.html#input%2Fradio)[Source ↗](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/input_radio.go#L8)
  
  - [Radio](https://lanej.io/charmed/dependencies.html#component-input%2Fradio)primary demonstration
