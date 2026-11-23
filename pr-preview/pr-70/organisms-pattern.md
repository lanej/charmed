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

[Example source](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/docs/snippets/pattern_measurement_history.go) · [Implementation](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/pattern/measurement_history.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

### Reusable component

[Measurement history dependency graph](https://lanej.io/charmed/dependencies.html#component-pattern%2Fmeasurement-history) · Partial map · [Source declaration](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/pattern/measurement_history.go#L36)

- [Grid inspector](https://lanej.io/charmed/dependencies.html#component-collection%2Fgrid-inspector) · composes · inspection

Consumer usage uses

- [Measurement history](https://lanej.io/charmed/dependencies.html#component-pattern%2Fmeasurement-history) · primary demonstration

[Example source](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/docs/snippets/pattern_measurement_history.go#L11)

[Organisms](https://lanej.io/charmed/organisms.html) / Patterns

## [List inspector](https://lanej.io/charmed/organisms-pattern.html#pattern%2flist-inspector)

Browse records with multiline summaries and inspect the evidence behind a selection. Wide views preview details beside the list; narrow views retain the selected record and open complete inspection with Enter.

### [Native preview](https://lanej.io/charmed/organisms-pattern.html#pattern%2flist-inspector-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
┌ Operations ──────────────────────────┐  ┌ Evidence ──────────────────────────────────────────────────┐
│ Network                              │  │ Operation  route-0042                                      │
│ → Gateway route update               │  │ Condition  Outcome unknown                                 │
│   Outcome unknown                    │  │ Evidence   Submission recorded; response lost. Remote rout │
│   Next: inspect remote evidence      │  │            e has not been observed.                        │
│                                      │  │ Next       Read the saved request's remote state before ch │
│ Compute                              │  │            anging traffic.                                 │
│   Worker capacity                    │  │                                                            │
│   Outcome confirmed                  │  │                                                            │
│   Next: check queue latency          │  │                                                            │
│                                      │  │                                                            │
│   Indexer restart                    │  │                                                            │
│   Success reported                   │  │                                                            │
│   Next: read process observer        │  │                                                            │
│                                      │  │                                                            │
│                                      │  │                                                            │
│                                      │  │                                                            │
│                                      │  │                                                            │
└──────────────────────────────────────┘  └────────────────────────────────────────────────────────────┘
j/k select  enter inspect  esc back  [/] scroll
```

Demo · truecolor · 112 × 28 terminal cells

```
┌ Operations ──────────────────────────┐  ┌ Evidence ──────────────────────────────────────────────────────────┐
│ Network                              │  │ Gateway route update                                               │
│ → Gateway route update               │  │  Outcome unknown                                                   │
│    Outcome unknown                   │  │                                                                    │
│   Response lost after submission     │  │ Operation  route-2026-11-23-0042                                   │
│   Next: inspect remote evidence      │  │ Source     Controller journal / last observed 09:42 UTC            │
│                                      │  │                                                                    │
│ Storage                              │  │ Observed evidence                                                  │
│   Cache warm-up                      │  │ The controller recorded a request to switch the gateway route. Its │
│    In progress                       │  │ response timed out after submission. The remote route has not been │
│                                      │  │ observed since the request.                                        │
│ Compute                              │  │                                                                    │
│   Worker capacity                    │  │ Next step                                                          │
│    Outcome confirmed                 │  │ Inspect the saved route request before changing traffic.           │
│   12 workers observed                │  │                                                                    │
│                                      │  │ Interpretation limits                                              │
│ Search                               │  │ A recorded request establishes intent. A caller's report           │
│   Indexer restart                    │  │ establishes what the caller said. An independent observation may   │
│    Success reported                  │  │ confirm the external state; it does not automatically prove        │
│   Caller report only                 │  │ request attribution, service health or incident resolution.        │
│   No independent process sample      │  │                                                                    │
│                                      │  │ Inspection notes                                                   │
│                                      │  │ Compare the saved identity with the observer's evidence.           │
│                                      │  │ Keep the last observation when a refresh fails.                    │
│                                      │  │ Check the observation time before taking another action.           │
│                                      │  │ Record any remaining attribution gap.                              │
└──────────────────────────────────────┘  └────────────────────────────────────────────────────────────────────┘
j/k select  enter inspect  [/] scroll
```

Narrow · truecolor · 40 × 16 terminal cells

```
┌ Operations ──────────────────────────┐
│ Network                              │
│ → Gateway route update               │
│    Outcome unknown                   │
│   Response lost after submission     │
│   Next: inspect remote evidence      │
│                                      │
│ Storage                              │
│   Cache warm-up                      │
│    In progress                       │
│                                      │
│ Compute                              │
│   Worker capacity                    │
│    Outcome confirmed                 │
└──────────────────────────────────────┘
j/k select  enter inspect  [/] scroll
```

256 colors · ansi256 · 112 × 28 terminal cells

```
┌ Operations ──────────────────────────┐  ┌ Evidence ──────────────────────────────────────────────────────────┐
│ Network                              │  │ Gateway route update                                               │
│ → Gateway route update               │  │  Outcome unknown                                                   │
│    Outcome unknown                   │  │                                                                    │
│   Response lost after submission     │  │ Operation  route-2026-11-23-0042                                   │
│   Next: inspect remote evidence      │  │ Source     Controller journal / last observed 09:42 UTC            │
│                                      │  │                                                                    │
│ Storage                              │  │ Observed evidence                                                  │
│   Cache warm-up                      │  │ The controller recorded a request to switch the gateway route. Its │
│    In progress                       │  │ response timed out after submission. The remote route has not been │
│                                      │  │ observed since the request.                                        │
│ Compute                              │  │                                                                    │
│   Worker capacity                    │  │ Next step                                                          │
│    Outcome confirmed                 │  │ Inspect the saved route request before changing traffic.           │
│   12 workers observed                │  │                                                                    │
│                                      │  │ Interpretation limits                                              │
│ Search                               │  │ A recorded request establishes intent. A caller's report           │
│   Indexer restart                    │  │ establishes what the caller said. An independent observation may   │
│    Success reported                  │  │ confirm the external state; it does not automatically prove        │
│   Caller report only                 │  │ request attribution, service health or incident resolution.        │
│   No independent process sample      │  │                                                                    │
│                                      │  │ Inspection notes                                                   │
│                                      │  │ Compare the saved identity with the observer's evidence.           │
│                                      │  │ Keep the last observation when a refresh fails.                    │
│                                      │  │ Check the observation time before taking another action.           │
│                                      │  │ Record any remaining attribution gap.                              │
└──────────────────────────────────────┘  └────────────────────────────────────────────────────────────────────┘
j/k select  enter inspect  [/] scroll
```

16 colors · ansi16 · 112 × 28 terminal cells

```
┌ Operations ──────────────────────────┐  ┌ Evidence ──────────────────────────────────────────────────────────┐
│ Network                              │  │ Gateway route update                                               │
│ → Gateway route update               │  │  Outcome unknown                                                   │
│    Outcome unknown                   │  │                                                                    │
│   Response lost after submission     │  │ Operation  route-2026-11-23-0042                                   │
│   Next: inspect remote evidence      │  │ Source     Controller journal / last observed 09:42 UTC            │
│                                      │  │                                                                    │
│ Storage                              │  │ Observed evidence                                                  │
│   Cache warm-up                      │  │ The controller recorded a request to switch the gateway route. Its │
│    In progress                       │  │ response timed out after submission. The remote route has not been │
│                                      │  │ observed since the request.                                        │
│ Compute                              │  │                                                                    │
│   Worker capacity                    │  │ Next step                                                          │
│    Outcome confirmed                 │  │ Inspect the saved route request before changing traffic.           │
│   12 workers observed                │  │                                                                    │
│                                      │  │ Interpretation limits                                              │
│ Search                               │  │ A recorded request establishes intent. A caller's report           │
│   Indexer restart                    │  │ establishes what the caller said. An independent observation may   │
│    Success reported                  │  │ confirm the external state; it does not automatically prove        │
│   Caller report only                 │  │ request attribution, service health or incident resolution.        │
│   No independent process sample      │  │                                                                    │
│                                      │  │ Inspection notes                                                   │
│                                      │  │ Compare the saved identity with the observer's evidence.           │
│                                      │  │ Keep the last observation when a refresh fails.                    │
│                                      │  │ Check the observation time before taking another action.           │
│                                      │  │ Record any remaining attribution gap.                              │
└──────────────────────────────────────┘  └────────────────────────────────────────────────────────────────────┘
j/k select  enter inspect  [/] scroll
```

Monochrome · none · 112 × 28 terminal cells

```
+ Operations --------------------------+  + Evidence ----------------------------------------------------------+
| Network                              |  | Gateway route update                                               |
| > Gateway route update               |  | [Outcome unknown]                                                  |
|   [Outcome unknown]                  |  |                                                                    |
|   Response lost after submission     |  | Operation  route-2026-11-23-0042                                   |
|   Next: inspect remote evidence      |  | Source     Controller journal / last observed 09:42 UTC            |
|                                      |  |                                                                    |
| Storage                              |  | Observed evidence                                                  |
|   Cache warm-up                      |  | The controller recorded a request to switch the gateway route. Its |
|   [In progress]                      |  | response timed out after submission. The remote route has not been |
|                                      |  | observed since the request.                                        |
| Compute                              |  |                                                                    |
|   Worker capacity                    |  | Next step                                                          |
|   [Outcome confirmed]                |  | Inspect the saved route request before changing traffic.           |
|   12 workers observed                |  |                                                                    |
|                                      |  | Interpretation limits                                              |
| Search                               |  | A recorded request establishes intent. A caller's report           |
|   Indexer restart                    |  | establishes what the caller said. An independent observation may   |
|   [Success reported]                 |  | confirm the external state; it does not automatically prove        |
|   Caller report only                 |  | request attribution, service health or incident resolution.        |
|   No independent process sample      |  |                                                                    |
|                                      |  | Inspection notes                                                   |
|                                      |  | Compare the saved identity with the observer's evidence.           |
|                                      |  | Keep the last observation when a refresh fails.                    |
|                                      |  | Check the observation time before taking another action.           |
|                                      |  | Record any remaining attribution gap.                              |
+--------------------------------------+  +--------------------------------------------------------------------+
j/k select  enter inspect  [/] scroll
```

Inspecting · truecolor · 112 × 28 terminal cells

```
┌ Evidence ────────────────────────────────────────────────────────────────────────────────────────────────────┐
│ Gateway route update                                                                                         │
│  Outcome unknown                                                                                             │
│                                                                                                              │
│ Operation  route-2026-11-23-0042                                                                             │
│ Source     Controller journal / last observed 09:42 UTC                                                      │
│                                                                                                              │
│ Observed evidence                                                                                            │
│ The controller recorded a request to switch the gateway route. Its response timed out after submission. The  │
│ remote route has not been observed since the request.                                                        │
│                                                                                                              │
│ Next step                                                                                                    │
│ Inspect the saved route request before changing traffic.                                                     │
│                                                                                                              │
│ Interpretation limits                                                                                        │
│ A recorded request establishes intent. A caller's report establishes what the caller said. An independent    │
│ observation may confirm the external state; it does not automatically prove request attribution, service     │
│ health or incident resolution.                                                                               │
│                                                                                                              │
│ Inspection notes                                                                                             │
│ Compare the saved identity with the observer's evidence.                                                     │
│ Keep the last observation when a refresh fails.                                                              │
│ Check the observation time before taking another action.                                                     │
│ Record any remaining attribution gap.                                                                        │
│ Choose the next action in the application.                                                                   │
│                                                                                                              │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
j/k scroll  esc back  [/] scroll
```

Empty · truecolor · 112 × 28 terminal cells

```
┌ Operations ──────────────────────────┐  ┌ Evidence ──────────────────────────────────────────────────────────┐
│ No operations observed               │  │ No details supplied                                                │
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
j/k select  enter inspect  [/] scroll
```

Loading · truecolor · 112 × 28 terminal cells

```
┌ Operations ──────────────────────────┐  ┌ Evidence ──────────────────────────────────────────────────────────┐
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
j/k select  enter inspect  [/] scroll
```

Error · truecolor · 112 × 28 terminal cells

```
┌ Operations ──────────────────────────┐  ┌ Evidence ──────────────────────────────────────────────────────────┐
│ Network                              │  │  Refresh failed  Last observation retained.                        │
│ → Gateway route update               │  │                                                                    │
│    Outcome unknown                   │  │ Gateway route update                                               │
│   Response lost after submission     │  │  Outcome unknown                                                   │
│   Next: inspect remote evidence      │  │                                                                    │
│                                      │  │ Operation  route-2026-11-23-0042                                   │
│ Storage                              │  │ Source     Controller journal / last observed 09:42 UTC            │
│   Cache warm-up                      │  │                                                                    │
│    In progress                       │  │ Observed evidence                                                  │
│                                      │  │ The controller recorded a request to switch the gateway route. Its │
│ Compute                              │  │ response timed out after submission. The remote route has not been │
│   Worker capacity                    │  │ observed since the request.                                        │
│    Outcome confirmed                 │  │                                                                    │
│   12 workers observed                │  │ Next step                                                          │
│                                      │  │ Inspect the saved route request before changing traffic.           │
│ Search                               │  │                                                                    │
│   Indexer restart                    │  │ Interpretation limits                                              │
│    Success reported                  │  │ A recorded request establishes intent. A caller's report           │
│   Caller report only                 │  │ establishes what the caller said. An independent observation may   │
│   No independent process sample      │  │ confirm the external state; it does not automatically prove        │
│                                      │  │ request attribution, service health or incident resolution.        │
│                                      │  │                                                                    │
│                                      │  │ Inspection notes                                                   │
│                                      │  │ Compare the saved identity with the observer's evidence.           │
│                                      │  │ Keep the last observation when a refresh fails.                    │
└──────────────────────────────────────┘  └────────────────────────────────────────────────────────────────────┘
j/k select  enter inspect  [/] scroll
```

Confirmed · truecolor · 112 × 28 terminal cells

```
┌ Operations ──────────────────────────┐  ┌ Evidence ──────────────────────────────────────────────────────────┐
│ Network                              │  │ Gateway route update                                               │
│ → Gateway route update               │  │  Outcome confirmed                                                 │
│    Outcome confirmed                 │  │                                                                    │
│   Requested route independently obse │  │ Operation  route-2026-11-23-0042                                   │
│   rved                               │  │ Source     Route observer / sampled 09:43 UTC                      │
│   Request attribution remains unknow │  │                                                                    │
│   n                                  │  │ Observed evidence                                                  │
│                                      │  │ A later read found the requested gateway route active. The route   │
│ Storage                              │  │ outcome is confirmed; the lost response still prevents attributing │
│   Cache warm-up                      │  │ the change to this exact request.                                  │
│    In progress                       │  │                                                                    │
│                                      │  │ Next step                                                          │
│ Compute                              │  │ Observe request errors before closing the incident.                │
│   Worker capacity                    │  │                                                                    │
│    Outcome confirmed                 │  │ Interpretation limits                                              │
│   12 workers observed                │  │ A recorded request establishes intent. A caller's report           │
│                                      │  │ establishes what the caller said. An independent observation may   │
│ Search                               │  │ confirm the external state; it does not automatically prove        │
│   Indexer restart                    │  │ request attribution, service health or incident resolution.        │
│    Success reported                  │  │                                                                    │
│   Caller report only                 │  │ Inspection notes                                                   │
│   No independent process sample      │  │ Compare the saved identity with the observer's evidence.           │
│                                      │  │ Keep the last observation when a refresh fails.                    │
│                                      │  │ Check the observation time before taking another action.           │
│                                      │  │ Record any remaining attribution gap.                              │
└──────────────────────────────────────┘  └────────────────────────────────────────────────────────────────────┘
j/k select  enter inspect  [/] scroll
```

ASCII · truecolor · 112 × 28 terminal cells

```
+ Operations --------------------------+  + Evidence ----------------------------------------------------------+
| Network                              |  | Gateway route update                                               |
| > Gateway route update               |  | [Outcome unknown]                                                  |
|   [Outcome unknown]                  |  |                                                                    |
|   Response lost after submission     |  | Operation  route-2026-11-23-0042                                   |
|   Next: inspect remote evidence      |  | Source     Controller journal / last observed 09:42 UTC            |
|                                      |  |                                                                    |
| Storage                              |  | Observed evidence                                                  |
|   Cache warm-up                      |  | The controller recorded a request to switch the gateway route. Its |
|   [In progress]                      |  | response timed out after submission. The remote route has not been |
|                                      |  | observed since the request.                                        |
| Compute                              |  |                                                                    |
|   Worker capacity                    |  | Next step                                                          |
|   [Outcome confirmed]                |  | Inspect the saved route request before changing traffic.           |
|   12 workers observed                |  |                                                                    |
|                                      |  | Interpretation limits                                              |
| Search                               |  | A recorded request establishes intent. A caller's report           |
|   Indexer restart                    |  | establishes what the caller said. An independent observation may   |
|   [Success reported]                 |  | confirm the external state; it does not automatically prove        |
|   Caller report only                 |  | request attribution, service health or incident resolution.        |
|   No independent process sample      |  |                                                                    |
|                                      |  | Inspection notes                                                   |
|                                      |  | Compare the saved identity with the observer's evidence.           |
|                                      |  | Keep the last observation when a refresh fails.                    |
|                                      |  | Check the observation time before taking another action.           |
|                                      |  | Record any remaining attribution gap.                              |
+--------------------------------------+  +--------------------------------------------------------------------+
j/k select  enter inspect  [/] scroll
```

Unicode monochrome · none · 112 × 28 terminal cells

```
+ Operations --------------------------+  + Evidence ----------------------------------------------------------+
| Network                              |  | Gateway route update                                               |
| > Gateway route update               |  | [Outcome unknown]                                                  |
|   [Outcome unknown]                  |  |                                                                    |
|   Response lost after submission     |  | Operation  route-2026-11-23-0042                                   |
|   Next: inspect remote evidence      |  | Source     Controller journal / last observed 09:42 UTC            |
|                                      |  |                                                                    |
| Storage                              |  | Observed evidence                                                  |
|   Cache warm-up                      |  | The controller recorded a request to switch the gateway route. Its |
|   [In progress]                      |  | response timed out after submission. The remote route has not been |
|                                      |  | observed since the request.                                        |
| Compute                              |  |                                                                    |
|   Worker capacity                    |  | Next step                                                          |
|   [Outcome confirmed]                |  | Inspect the saved route request before changing traffic.           |
|   12 workers observed                |  |                                                                    |
|                                      |  | Interpretation limits                                              |
| Search                               |  | A recorded request establishes intent. A caller's report           |
|   Indexer restart                    |  | establishes what the caller said. An independent observation may   |
|   [Success reported]                 |  | confirm the external state; it does not automatically prove        |
|   Caller report only                 |  | request attribution, service health or incident resolution.        |
|   No independent process sample      |  |                                                                    |
|                                      |  | Inspection notes                                                   |
|                                      |  | Compare the saved identity with the observer's evidence.           |
|                                      |  | Keep the last observation when a refresh fails.                    |
|                                      |  | Check the observation time before taking another action.           |
|                                      |  | Record any remaining attribution gap.                              |
+--------------------------------------+  +--------------------------------------------------------------------+
j/k select  enter inspect  [/] scroll
```

### [Go usage](https://lanej.io/charmed/organisms-pattern.html#pattern%2flist-inspector-code)

Imports: [`ansi`](https://pkg.go.dev/github.com/charmbracelet/x/ansi "github.com/charmbracelet/x/ansi") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`pattern`](https://pkg.go.dev/github.com/lanej/charmed/pattern "github.com/lanej/charmed/pattern") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
row := func(title, condition, next string) func(int, bool) string {
	return func(width int, selected bool) string {
		style := palette.Text()
		if selected {
			style = palette.Heading()
		}

		return ansi.Hardwrap(style.Render(title)+"\n"+condition+"\nNext: "+next, width, true)
	}
}
records := []pattern.InspectorRecord{
	{ID: "route-0042", Heading: "Network", Row: row("Gateway route update", "Outcome unknown", "inspect remote evidence"), Detail: func(width int) string {
		return layout.KeyValue([]layout.Pair{
			{Label: "Operation", Value: "route-0042"},
			{Label: "Condition", Value: "Outcome unknown"},
			{Label: "Evidence", Value: "Submission recorded; response lost. Remote route has not been observed."},
			{Label: "Next", Value: "Read the saved request's remote state before changing traffic."},
		}, width, palette)
	}},
	{ID: "pool-0007", Heading: "Compute", Row: row("Worker capacity", "Outcome confirmed", "check queue latency"), Detail: func(int) string {
		return "An independent observer found twelve workers.\nCapacity does not establish healthy workers or acceptable latency."
	}},
	{ID: "index-0031", Row: row("Indexer restart", "Success reported", "read process observer"), Detail: func(int) string {
		return "The operator reported success.\nNo independent process sample confirms recovery."
	}},
}

inspector, err := pattern.NewListInspector(records, layout.Size{Width: 104, Height: 19}, palette)
if err != nil {
	return "", err
}

inspector.ListTitle, inspector.DetailTitle = "Operations", "Evidence"
inspector.Refresh()

help := text.KeyHint("j/k", "select", palette) + "  " + text.KeyHint("enter", "inspect", palette) + "  " + text.KeyHint("esc", "back", palette) + "  " + text.KeyHint("[/]", "scroll", palette)

return layout.Column(inspector.View(), help), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-pattern.html#pattern%2flist-inspector-integration)

Supply unique stable IDs and pure width-aware row and detail renderers. Forward Update while focused; Enter inspects, Escape returns, and brackets scroll the active native viewport. SetRecords validates before replacement and retains surviving selection; selection changes reset detail scrolling. Resize preserves native styles and bindings. Call Refresh after changing captured data or presentation. Preview controls the wide detail pane; ListWidth and MinDetailWidth set its threshold. Applications own observations, ordering, actions, refresh and quit. Run go run ./examples/inspection for the complete consumer.

### [Run this example](https://lanej.io/charmed/organisms-pattern.html#pattern%2flist-inspector-run)

```sh
go run ./examples/catalog pattern/list-inspector
```

[Example source](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/docs/snippets/pattern_list_inspector.go) · [Implementation](https://github.com/lanej/charmed/blob/cc72980fb2707507db01071d8516e7386896097a/pattern/list_inspector.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [collection/viewport](https://lanej.io/charmed/atoms-collection.html#collection/viewport) [foundations/layout](https://lanej.io/charmed/foundations-layout.html#foundations/layout)
