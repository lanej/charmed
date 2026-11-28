[Organisms](https://lanej.io/charmed/organisms.html)Patterns

# [Patterns](https://lanej.io/charmed/organisms-pattern.html#level-organisms-pattern)

2 components · [All organisms](https://lanej.io/charmed/organisms.html)

- [Measurement history](https://lanej.io/charmed/organisms-pattern.html#pattern%2fmeasurement-history)
- [List inspector](https://lanej.io/charmed/organisms-pattern.html#pattern%2flist-inspector)

[Organisms](https://lanej.io/charmed/organisms.html) / Patterns

## [Measurement history](https://lanej.io/charmed/organisms-pattern.html#pattern%2fmeasurement-history)

Inspect one member across supplied measurement windows and metric rows. Exact intervals and raw observations remain accessible independently of the compact value shown in each cell.

### [Native preview](https://lanej.io/charmed/organisms-pattern.html#pattern%2fmeasurement-history-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

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

Demo · truecolor · 100 × 30 terminal cells

```
 Metric / Window      04:00:00       05:00:00               west / sensor-4                         
 Temperature         Metric absent   Metric absent          Temperature                             
 Energy              27kWh          28kWh                   2026-01-01T05:00:00Z – 2026-01-01T06:00 
 Yield               87%            88%                     :00Z                                    
 Humidity            41%            42%                     Metric absent                           
Rows 1-4/4  Columns 5-6/6
```

Narrow · truecolor · 40 × 16 terminal cells

```
 Metric / Window      05:00:00          
 Temperature          Metric absent     
 Energy              28kWh              
 Yield               88%                
 Humidity            42%                
Rows 1-4/4  Columns 6-6/6               
                                        
                                        
west / sensor-4                         
Temperature                             
2026-01-01T05:00:00Z – 2026-01-01T06:00 
:00Z                                    
Metric absent
```

256 colors · ansi256 · 100 × 30 terminal cells

```
 Metric / Window      04:00:00       05:00:00               west / sensor-4                         
 Temperature         Metric absent   Metric absent          Temperature                             
 Energy              27kWh          28kWh                   2026-01-01T05:00:00Z – 2026-01-01T06:00 
 Yield               87%            88%                     :00Z                                    
 Humidity            41%            42%                     Metric absent                           
Rows 1-4/4  Columns 5-6/6
```

16 colors · ansi16 · 100 × 30 terminal cells

```
 Metric / Window      04:00:00       05:00:00               west / sensor-4                         
 Temperature         Metric absent   Metric absent          Temperature                             
 Energy              27kWh          28kWh                   2026-01-01T05:00:00Z – 2026-01-01T06:00 
 Yield               87%            88%                     :00Z                                    
 Humidity            41%            42%                     Metric absent                           
Rows 1-4/4  Columns 5-6/6
```

Monochrome · none · 100 × 30 terminal cells

```
 Metric / Window      04:00:00       05:00:00               west / sensor-4                         
 Temperature         Metric absent   Metric absent          Temperature                             
 Energy              27kWh          28kWh                   2026-01-01T05:00:00Z - 2026-01-01T06:00 
 Yield               87%            88%                     :00Z                                    
 Humidity            41%            42%                     Metric absent                           
Rows 1-4/4  Columns 5-6/6
```

### [Go usage](https://lanej.io/charmed/organisms-pattern.html#pattern%2fmeasurement-history-code)

Imports: [`cohort`](https://pkg.go.dev/github.com/lanej/charmed/cohort "github.com/lanej/charmed/cohort") · [`fixture`](https://pkg.go.dev/github.com/lanej/charmed/examples/cohorts/fixture "github.com/lanej/charmed/examples/cohorts/fixture") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`pattern`](https://pkg.go.dev/github.com/lanej/charmed/pattern "github.com/lanej/charmed/pattern") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
history, err := pattern.NewMeasurementHistory(fixture.Data(), cohort.MemberKey{CohortID: "west", MemberID: "sensor-4"}, layout.Size{Width: 80, Height: 16}, theme.New(theme.Dark))
if err != nil {
	return "", err
}
// Temperature is missing for this sensor; other measurements remain observable.
// SetData copies and validates. SetMember resets detail scroll, preserving the coordinate.
// Parent owns Tab through history.Inspector.FocusDetail and forwards Update commands.
return history.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-pattern.html#pattern%2fmeasurement-history-integration)

Supply cohort data and a member key; construction and SetData validate and copy the observations. Surviving metric and window identities retain selection; removing a selected window chooses the latest supplied window. SetMember resets detail scrolling. Route focus through Inspector.FocusDetail and forward Update commands. Native grid customization survives refreshes; missing members, missing metrics and unavailable ratios remain distinct from scalar zero.

### [Run this example](https://lanej.io/charmed/organisms-pattern.html#pattern%2fmeasurement-history-run)

```sh
go run ./examples/catalog pattern/measurement-history
```

[Example source](https://github.com/lanej/charmed/blob/cc0823efe6b1be9d1911e813477912332826207b/docs/snippets/pattern_measurement_history.go) · [Implementation](https://github.com/lanej/charmed/blob/cc0823efe6b1be9d1911e813477912332826207b/pattern/measurement_history.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

### Reusable component

[Measurement history dependency graph](https://lanej.io/charmed/dependencies.html#component-pattern%2Fmeasurement-history) · Partial map · [Source declaration](https://github.com/lanej/charmed/blob/cc0823efe6b1be9d1911e813477912332826207b/pattern/measurement_history.go#L36)

- [Grid inspector](https://lanej.io/charmed/dependencies.html#component-collection%2Fgrid-inspector) · composes · inspection

Consumer usage uses

- [Measurement history](https://lanej.io/charmed/dependencies.html#component-pattern%2Fmeasurement-history) · primary demonstration

[Example source](https://github.com/lanej/charmed/blob/cc0823efe6b1be9d1911e813477912332826207b/docs/snippets/pattern_measurement_history.go#L11)

[Organisms](https://lanej.io/charmed/organisms.html) / Patterns

## [List inspector](https://lanej.io/charmed/organisms-pattern.html#pattern%2flist-inspector)

Browse structured summaries and inspect the selected record’s next action, verification, delivery stages and history. Wide views preview details beside the list; narrow views retain selection and open complete inspection with Enter.

### [Native preview](https://lanej.io/charmed/organisms-pattern.html#pattern%2flist-inspector-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
┌ Deliveries ──────────────────────────┐  ┌────────────────────────────────────────────────────────────┐
│ Needs attention                      │  │ DELIVERY INSPECTION · ILLUSTRATIVE                         │
│ → PR #42 · Route update              │  │ PR #42 · Route update   Outcome uncertain                  │
│ ┃  Outcome uncertain                 │  │                                         Observed 09:42 UTC │
│ ┃ Next action: Reconcile the saved   │  │                                                            │
│ ┃ merge operation                    │  │ ┃ NEXT ACTION                                              │
│ ┃ Observed: 09:42 UTC                │  │ ┃ Reconcile the saved merge operation                      │
│                                      │  │ ┃ The request was recorded, but the remote response was    │
│ Completed                            │  │ ┃ unavailable.                                             │
│   PR #41 · Request parsing           │  │                                                            │
│    Merged                            │  │ Identity  ──────────────────────────────────────────────── │
│   Observed: 09:38 UTC                │  │ Repository  example/service                                │
│                                      │  │ Owner       Worker 2                                       │
│                                      │  │                                                            │
│                                      │  │ Verification  ──────────────────────────────────────────── │
│                                      │  │ Tests  Passed                                              │
│                                      │  │ Merge  Unverified                                          │
│                                      │  │                                                            │
│                                      │  │                                                            │
└──────────────────────────────────────┘  └────────────────────────────────────────────────────────────┘
j/k select  enter inspect  esc back
```

Demo · truecolor · 112 × 28 terminal cells

```
┌ Deliveries ──────────────────────────┐  ┌────────────────────────────────────────────────────────────────────┐
│ Needs attention                      │  │ DELIVERY INSPECTION · ILLUSTRATIVE                                 │
│ → PR #42 · Route update              │  │ PR #42 · Route update   Outcome uncertain       Observed 09:42 UTC │
│ ┃  Outcome uncertain                 │  │                                                                    │
│ ┃ Next action: Reconcile the saved   │  │ ┃ NEXT ACTION                                                      │
│ ┃ merge operation                    │  │ ┃ Reconcile the saved merge operation                              │
│ ┃ Observed: 09:42 UTC                │  │ ┃ The request was recorded, but the remote response was            │
│                                      │  │ ┃ unavailable.                                                     │
│ Needs attention                      │  │                                                                    │
│   PR #43 · Cache limits              │  │ Identity  ──────────────────────────────────────────────────────── │
│    Feedback open                     │  │ Repository  example/service                                        │
│   Next action: Review the requested  │  │ Revision    a81f39d                                                │
│   change                             │  │ Owner       Worker 2                                               │
│   Observed: 09:40 UTC                │  │                                                                    │
│                                      │  │ Verification  ──────────────────────────────────────────────────── │
│ In progress                          │  │ Tests   ✓ Passed                                                   │
│   PR #44 · Retry interval            │  │ Review  ✓ Complete                                                 │
│    Tests running                     │  │ Merge   ? Unverified                                               │
│   Next action: Wait for the          │  │                                                                    │
│   integration suite                  │  │ Feedback delivery  ─────────────────────────────────────────────── │
│   Observed: 09:39 UTC                │  │ Source       ✓ Recorded                                            │
│                                      │  │              Review comment                                        │
│ Completed                            │  │ Transport    ✓ Delivered                                           │
│   PR #41 · Request parsing           │  │              To Worker 2                                           │
│    Merged                            │  │ Receipt      ? Unconfirmed                                         │
│   Observed: 09:38 UTC                │  │              No acknowledgement                                    │
└──────────────────────────────────────┘  └────────────────────────────────────────────────────────────────────┘
j/k select  enter inspect  [/] page
```

Narrow · truecolor · 40 × 16 terminal cells

```
┌ Deliveries ──────────────────────────┐
│ Needs attention                      │
│ → PR #42 · Route update              │
│ ┃  Outcome uncertain                 │
│ ┃ Next action: Reconcile the saved   │
│ ┃ merge operation                    │
│ ┃ Observed: 09:42 UTC                │
│                                      │
│ Needs attention                      │
│   PR #43 · Cache limits              │
│    Feedback open                     │
│   Next action: Review the requested  │
│   change                             │
│   Observed: 09:40 UTC                │
└──────────────────────────────────────┘
j/k select  enter inspect  [/] page
```

256 colors · ansi256 · 112 × 28 terminal cells

```
┌ Deliveries ──────────────────────────┐  ┌────────────────────────────────────────────────────────────────────┐
│ Needs attention                      │  │ DELIVERY INSPECTION · ILLUSTRATIVE                                 │
│ → PR #42 · Route update              │  │ PR #42 · Route update   Outcome uncertain       Observed 09:42 UTC │
│ ┃  Outcome uncertain                 │  │                                                                    │
│ ┃ Next action: Reconcile the saved   │  │ ┃ NEXT ACTION                                                      │
│ ┃ merge operation                    │  │ ┃ Reconcile the saved merge operation                              │
│ ┃ Observed: 09:42 UTC                │  │ ┃ The request was recorded, but the remote response was            │
│                                      │  │ ┃ unavailable.                                                     │
│ Needs attention                      │  │                                                                    │
│   PR #43 · Cache limits              │  │ Identity  ──────────────────────────────────────────────────────── │
│    Feedback open                     │  │ Repository  example/service                                        │
│   Next action: Review the requested  │  │ Revision    a81f39d                                                │
│   change                             │  │ Owner       Worker 2                                               │
│   Observed: 09:40 UTC                │  │                                                                    │
│                                      │  │ Verification  ──────────────────────────────────────────────────── │
│ In progress                          │  │ Tests   ✓ Passed                                                   │
│   PR #44 · Retry interval            │  │ Review  ✓ Complete                                                 │
│    Tests running                     │  │ Merge   ? Unverified                                               │
│   Next action: Wait for the          │  │                                                                    │
│   integration suite                  │  │ Feedback delivery  ─────────────────────────────────────────────── │
│   Observed: 09:39 UTC                │  │ Source       ✓ Recorded                                            │
│                                      │  │              Review comment                                        │
│ Completed                            │  │ Transport    ✓ Delivered                                           │
│   PR #41 · Request parsing           │  │              To Worker 2                                           │
│    Merged                            │  │ Receipt      ? Unconfirmed                                         │
│   Observed: 09:38 UTC                │  │              No acknowledgement                                    │
└──────────────────────────────────────┘  └────────────────────────────────────────────────────────────────────┘
j/k select  enter inspect  [/] page
```

16 colors · ansi16 · 112 × 28 terminal cells

```
┌ Deliveries ──────────────────────────┐  ┌────────────────────────────────────────────────────────────────────┐
│ Needs attention                      │  │ DELIVERY INSPECTION · ILLUSTRATIVE                                 │
│ → PR #42 · Route update              │  │ PR #42 · Route update   Outcome uncertain       Observed 09:42 UTC │
│ ┃  Outcome uncertain                 │  │                                                                    │
│ ┃ Next action: Reconcile the saved   │  │ ┃ NEXT ACTION                                                      │
│ ┃ merge operation                    │  │ ┃ Reconcile the saved merge operation                              │
│ ┃ Observed: 09:42 UTC                │  │ ┃ The request was recorded, but the remote response was            │
│                                      │  │ ┃ unavailable.                                                     │
│ Needs attention                      │  │                                                                    │
│   PR #43 · Cache limits              │  │ Identity  ──────────────────────────────────────────────────────── │
│    Feedback open                     │  │ Repository  example/service                                        │
│   Next action: Review the requested  │  │ Revision    a81f39d                                                │
│   change                             │  │ Owner       Worker 2                                               │
│   Observed: 09:40 UTC                │  │                                                                    │
│                                      │  │ Verification  ──────────────────────────────────────────────────── │
│ In progress                          │  │ Tests   ✓ Passed                                                   │
│   PR #44 · Retry interval            │  │ Review  ✓ Complete                                                 │
│    Tests running                     │  │ Merge   ? Unverified                                               │
│   Next action: Wait for the          │  │                                                                    │
│   integration suite                  │  │ Feedback delivery  ─────────────────────────────────────────────── │
│   Observed: 09:39 UTC                │  │ Source       ✓ Recorded                                            │
│                                      │  │              Review comment                                        │
│ Completed                            │  │ Transport    ✓ Delivered                                           │
│   PR #41 · Request parsing           │  │              To Worker 2                                           │
│    Merged                            │  │ Receipt      ? Unconfirmed                                         │
│   Observed: 09:38 UTC                │  │              No acknowledgement                                    │
└──────────────────────────────────────┘  └────────────────────────────────────────────────────────────────────┘
j/k select  enter inspect  [/] page
```

Monochrome · none · 112 × 28 terminal cells

```
+ Deliveries --------------------------+  +--------------------------------------------------------------------+
| Needs attention                      |  | DELIVERY INSPECTION * ILLUSTRATIVE                                 |
| > PR #42 * Route update              |  | PR #42 * Route update  [Outcome uncertain]      Observed 09:42 UTC |
| # [Outcome uncertain]                |  |                                                                    |
| # Next action: Reconcile the saved   |  | # NEXT ACTION                                                      |
| # merge operation                    |  | # Reconcile the saved merge operation                              |
| # Observed: 09:42 UTC                |  | # The request was recorded, but the remote response was            |
|                                      |  | # unavailable.                                                     |
| Needs attention                      |  |                                                                    |
|   PR #43 * Cache limits              |  | Identity  -------------------------------------------------------- |
|   [Feedback open]                    |  | Repository  example/service                                        |
|   Next action: Review the requested  |  | Revision    a81f39d                                                |
|   change                             |  | Owner       Worker 2                                               |
|   Observed: 09:40 UTC                |  |                                                                    |
|                                      |  | Verification  ---------------------------------------------------- |
| In progress                          |  | Tests   + Passed                                                   |
|   PR #44 * Retry interval            |  | Review  + Complete                                                 |
|   [Tests running]                    |  | Merge   ? Unverified                                               |
|   Next action: Wait for the          |  |                                                                    |
|   integration suite                  |  | Feedback delivery  ----------------------------------------------- |
|   Observed: 09:39 UTC                |  | Source       + Recorded                                            |
|                                      |  |              Review comment                                        |
| Completed                            |  | Transport    + Delivered                                           |
|   PR #41 * Request parsing           |  |              To Worker 2                                           |
|   [Merged]                           |  | Receipt      ? Unconfirmed                                         |
|   Observed: 09:38 UTC                |  |              No acknowledgement                                    |
+--------------------------------------+  +--------------------------------------------------------------------+
j/k select  enter inspect  [/] page
```

Inspecting · truecolor · 112 × 28 terminal cells

```
┌──────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│ DELIVERY INSPECTION · ILLUSTRATIVE                                                                           │
│ PR #42 · Route update   Outcome uncertain                                                 Observed 09:42 UTC │
│                                                                                                              │
│ ┃ NEXT ACTION                                                                                                │
│ ┃ Reconcile the saved merge operation                                                                        │
│ ┃ The request was recorded, but the remote response was unavailable.                                         │
│                                                                                                              │
│ Identity  ──────────────────────────────────────────    Verification  ────────────────────────────────────── │
│ Repository  example/service                             Tests   ✓ Passed                                     │
│ Revision    a81f39d                                     Review  ✓ Complete                                   │
│ Owner       Worker 2                                    Merge   ? Unverified                                 │
│                                                                                                              │
│ Feedback delivery  ───────────────────────────────────────────────────────────────────────────────────────── │
│ Source                    → Transport                 → Receipt                   → Disposition              │
│ ✓ Recorded                  ✓ Delivered                 ? Unconfirmed               ◷ Open                   │
│ Review comment              To Worker 2                 No acknowledgement          Awaiting decision        │
│                                                                                                              │
│ Activity  ────────────────────────────────────────────────────────────────────────────────────────────────── │
│ 09:41 UTC  ○ Merge requested                                                                                 │
│            │ Request persisted                                                                               │
│ 09:42 UTC  ○ Remote response unavailable                                                                     │
│              External outcome unknown                                                                        │
│                                                                                                              │
│                                                                                                              │
│                                                                                                              │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
j/k scroll  esc back  [/] page
```

Empty · truecolor · 112 × 28 terminal cells

```
┌ Deliveries ──────────────────────────┐  ┌────────────────────────────────────────────────────────────────────┐
│ No deliveries observed               │  │ No details supplied                                                │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
└──────────────────────────────────────┘  └────────────────────────────────────────────────────────────────────┘
j/k select  enter inspect  [/] page
```

Loading · truecolor · 112 × 28 terminal cells

```
┌ Deliveries ──────────────────────────┐  ┌────────────────────────────────────────────────────────────────────┐
│ Waiting for the first observation    │  │ No details supplied                                                │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
│                                      │  │                                                                    │
└──────────────────────────────────────┘  └────────────────────────────────────────────────────────────────────┘
j/k select  enter inspect  [/] page
```

Error · truecolor · 112 × 28 terminal cells

```
┌ Deliveries ──────────────────────────┐  ┌────────────────────────────────────────────────────────────────────┐
│ Needs attention                      │  │ DELIVERY INSPECTION · ILLUSTRATIVE                                 │
│ → PR #42 · Route update              │  │ PR #42 · Route update   Outcome uncertain       Observed 09:42 UTC │
│ ┃  Outcome uncertain                 │  │  Refresh failed  Last observation retained.                        │
│ ┃ Next action: Reconcile the saved   │  │                                                                    │
│ ┃ merge operation                    │  │ ┃ NEXT ACTION                                                      │
│ ┃ Observed: 09:42 UTC                │  │ ┃ Reconcile the saved merge operation                              │
│                                      │  │ ┃ The request was recorded, but the remote response was            │
│ Needs attention                      │  │ ┃ unavailable.                                                     │
│   PR #43 · Cache limits              │  │                                                                    │
│    Feedback open                     │  │ Identity  ──────────────────────────────────────────────────────── │
│   Next action: Review the requested  │  │ Repository  example/service                                        │
│   change                             │  │ Revision    a81f39d                                                │
│   Observed: 09:40 UTC                │  │ Owner       Worker 2                                               │
│                                      │  │                                                                    │
│ In progress                          │  │ Verification  ──────────────────────────────────────────────────── │
│   PR #44 · Retry interval            │  │ Tests   ✓ Passed                                                   │
│    Tests running                     │  │ Review  ✓ Complete                                                 │
│   Next action: Wait for the          │  │ Merge   ? Unverified                                               │
│   integration suite                  │  │                                                                    │
│   Observed: 09:39 UTC                │  │ Feedback delivery  ─────────────────────────────────────────────── │
│                                      │  │ Source       ✓ Recorded                                            │
│ Completed                            │  │              Review comment                                        │
│   PR #41 · Request parsing           │  │ Transport    ✓ Delivered                                           │
│    Merged                            │  │              To Worker 2                                           │
│   Observed: 09:38 UTC                │  │ Receipt      ? Unconfirmed                                         │
└──────────────────────────────────────┘  └────────────────────────────────────────────────────────────────────┘
j/k select  enter inspect  [/] page
```

Confirmed · truecolor · 112 × 28 terminal cells

```
┌ Deliveries ──────────────────────────┐  ┌────────────────────────────────────────────────────────────────────┐
│ Completed                            │  │ DELIVERY INSPECTION · ILLUSTRATIVE                                 │
│ → PR #42 · Route update              │  │ PR #42 · Route update   Merge confirmed         Observed 09:43 UTC │
│ ┃  Merge confirmed                   │  │                                                                    │
│ ┃ Observed: 09:43 UTC                │  │ Identity  ──────────────────────────────────────────────────────── │
│                                      │  │ Repository  example/service                                        │
│ Needs attention                      │  │ Revision    a81f39d                                                │
│   PR #43 · Cache limits              │  │ Owner       Worker 2                                               │
│    Feedback open                     │  │                                                                    │
│   Next action: Review the requested  │  │ Verification  ──────────────────────────────────────────────────── │
│   change                             │  │ Tests   ✓ Passed                                                   │
│   Observed: 09:40 UTC                │  │ Review  ✓ Complete                                                 │
│                                      │  │ Merge   ✓ Confirmed                                                │
│ In progress                          │  │                                                                    │
│   PR #44 · Retry interval            │  │ Feedback delivery  ─────────────────────────────────────────────── │
│    Tests running                     │  │ Source       ✓ Recorded                                            │
│   Next action: Wait for the          │  │              Review comment                                        │
│   integration suite                  │  │ Transport    ✓ Delivered                                           │
│   Observed: 09:39 UTC                │  │              To Worker 2                                           │
│                                      │  │ Receipt      ? Unconfirmed                                         │
│ Completed                            │  │              No acknowledgement                                    │
│   PR #41 · Request parsing           │  │ Disposition  ◷ Open                                                │
│    Merged                            │  │              Awaiting decision                                     │
│   Observed: 09:38 UTC                │  │                                                                    │
│                                      │  │ Activity  ──────────────────────────────────────────────────────── │
│                                      │  │ 09:41 UTC  ○ Merge requested                                       │
└──────────────────────────────────────┘  └────────────────────────────────────────────────────────────────────┘
j/k select  enter inspect  [/] page
```

ASCII · truecolor · 112 × 28 terminal cells

```
+ Deliveries --------------------------+  +--------------------------------------------------------------------+
| Needs attention                      |  | DELIVERY INSPECTION * ILLUSTRATIVE                                 |
| > PR #42 * Route update              |  | PR #42 * Route update  [Outcome uncertain]      Observed 09:42 UTC |
| # [Outcome uncertain]                |  |                                                                    |
| # Next action: Reconcile the saved   |  | # NEXT ACTION                                                      |
| # merge operation                    |  | # Reconcile the saved merge operation                              |
| # Observed: 09:42 UTC                |  | # The request was recorded, but the remote response was            |
|                                      |  | # unavailable.                                                     |
| Needs attention                      |  |                                                                    |
|   PR #43 * Cache limits              |  | Identity  -------------------------------------------------------- |
|   [Feedback open]                    |  | Repository  example/service                                        |
|   Next action: Review the requested  |  | Revision    a81f39d                                                |
|   change                             |  | Owner       Worker 2                                               |
|   Observed: 09:40 UTC                |  |                                                                    |
|                                      |  | Verification  ---------------------------------------------------- |
| In progress                          |  | Tests   + Passed                                                   |
|   PR #44 * Retry interval            |  | Review  + Complete                                                 |
|   [Tests running]                    |  | Merge   ? Unverified                                               |
|   Next action: Wait for the          |  |                                                                    |
|   integration suite                  |  | Feedback delivery  ----------------------------------------------- |
|   Observed: 09:39 UTC                |  | Source       + Recorded                                            |
|                                      |  |              Review comment                                        |
| Completed                            |  | Transport    + Delivered                                           |
|   PR #41 * Request parsing           |  |              To Worker 2                                           |
|   [Merged]                           |  | Receipt      ? Unconfirmed                                         |
|   Observed: 09:38 UTC                |  |              No acknowledgement                                    |
+--------------------------------------+  +--------------------------------------------------------------------+
j/k select  enter inspect  [/] page
```

Unicode monochrome · none · 112 × 28 terminal cells

```
+ Deliveries --------------------------+  +--------------------------------------------------------------------+
| Needs attention                      |  | DELIVERY INSPECTION * ILLUSTRATIVE                                 |
| > PR #42 * Route update              |  | PR #42 * Route update  [Outcome uncertain]      Observed 09:42 UTC |
| # [Outcome uncertain]                |  |                                                                    |
| # Next action: Reconcile the saved   |  | # NEXT ACTION                                                      |
| # merge operation                    |  | # Reconcile the saved merge operation                              |
| # Observed: 09:42 UTC                |  | # The request was recorded, but the remote response was            |
|                                      |  | # unavailable.                                                     |
| Needs attention                      |  |                                                                    |
|   PR #43 * Cache limits              |  | Identity  -------------------------------------------------------- |
|   [Feedback open]                    |  | Repository  example/service                                        |
|   Next action: Review the requested  |  | Revision    a81f39d                                                |
|   change                             |  | Owner       Worker 2                                               |
|   Observed: 09:40 UTC                |  |                                                                    |
|                                      |  | Verification  ---------------------------------------------------- |
| In progress                          |  | Tests   + Passed                                                   |
|   PR #44 * Retry interval            |  | Review  + Complete                                                 |
|   [Tests running]                    |  | Merge   ? Unverified                                               |
|   Next action: Wait for the          |  |                                                                    |
|   integration suite                  |  | Feedback delivery  ----------------------------------------------- |
|   Observed: 09:39 UTC                |  | Source       + Recorded                                            |
|                                      |  |              Review comment                                        |
| Completed                            |  | Transport    + Delivered                                           |
|   PR #41 * Request parsing           |  |              To Worker 2                                           |
|   [Merged]                           |  | Receipt      ? Unconfirmed                                         |
|   Observed: 09:38 UTC                |  |              No acknowledgement                                    |
+--------------------------------------+  +--------------------------------------------------------------------+
j/k select  enter inspect  [/] page
```

### [Go usage](https://lanej.io/charmed/organisms-pattern.html#pattern%2flist-inspector-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`pattern`](https://pkg.go.dev/github.com/lanej/charmed/pattern "github.com/lanej/charmed/pattern") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
summary := layout.RecordSummary{
	Title: "PR #42 · Route update", Status: text.Pill("Outcome uncertain", text.Warning, palette, text.Normal),
	Fields: []layout.Pair{{Label: "Observed", Value: "09:42 UTC"}},
	Next:   &layout.Pair{Label: "Next action", Value: "Reconcile the saved merge operation"},
}
document := layout.DetailDocument{
	Header: layout.IdentityHeader{ID: "DELIVERY INSPECTION · ILLUSTRATIVE"}, Summary: summary,
	ActionDetail: "The request was recorded, but the remote response was unavailable.",
	Sections: []layout.DetailSection{
		{ID: "identity", Title: "Identity", Fields: []layout.Pair{{Label: "Repository", Value: "example/service"}, {Label: "Owner", Value: "Worker 2"}}},
		{ID: "checks", Title: "Verification", Statuses: []layout.StatusItem{{Label: "Tests", Value: "Passed"}, {Label: "Merge", Value: "Unverified"}}},
	},
}
records := []pattern.InspectorRecord{
	{ID: "example/service#42", Heading: "Needs attention", Summary: &summary, Document: &document},
	{ID: "example/service#41", Heading: "Completed", Summary: &layout.RecordSummary{
		Title: "PR #41 · Request parsing", Status: text.Pill("Merged", text.Success, palette, text.Normal),
		Fields: []layout.Pair{{Label: "Observed", Value: "09:38 UTC"}},
	}},
}

inspector, err := pattern.NewListInspector(records, layout.Size{Width: 104, Height: 20}, palette)
if err != nil {
	return "", err
}

inspector.ListTitle, inspector.DetailTitle = "Deliveries", ""
inspector.Refresh()

return layout.Column(inspector.View(), text.KeyHint("j/k", "select", palette)+"  "+text.KeyHint("enter", "inspect", palette)+"  "+text.KeyHint("esc", "back", palette)), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-pattern.html#pattern%2flist-inspector-integration)

Supply stable record IDs with Summary and Document; Row and Detail callbacks override the typed presentation when needed. SetRecords validates atomically and copies typed snapshots. Structured details retain section-relative scrolling; selection changes reset details. Enter inspects, Escape returns and brackets scroll. Resize preserves native styles and bindings; Refresh reflows current content. Rebuild records to change pre-styled values. Preview controls the wide detail pane. Applications own observations, ordering, actions, refresh and quit. Run go run ./examples/inspection for the complete consumer.

### [Run this example](https://lanej.io/charmed/organisms-pattern.html#pattern%2flist-inspector-run)

```sh
go run ./examples/catalog pattern/list-inspector
```

[Example source](https://github.com/lanej/charmed/blob/cc0823efe6b1be9d1911e813477912332826207b/docs/snippets/pattern_list_inspector.go) · [Implementation](https://github.com/lanej/charmed/blob/cc0823efe6b1be9d1911e813477912332826207b/pattern/list_inspector.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [collection/viewport](https://lanej.io/charmed/atoms-collection.html#collection/viewport) [layout/record-summary](https://lanej.io/charmed/molecules-layout.html#layout/record-summary) [layout/detail-document](https://lanej.io/charmed/organisms-layout.html#layout/detail-document) [foundations/layout](https://lanej.io/charmed/foundations-layout.html#foundations/layout)
