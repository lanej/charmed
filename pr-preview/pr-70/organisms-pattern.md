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

[Example source](https://github.com/lanej/charmed/blob/49f755da5bd25a044043cf98bc303185c2ebce6d/docs/snippets/pattern_measurement_history.go) · [Implementation](https://github.com/lanej/charmed/blob/49f755da5bd25a044043cf98bc303185c2ebce6d/pattern/measurement_history.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

### Reusable component

[Measurement history dependency graph](https://lanej.io/charmed/dependencies.html#component-pattern%2Fmeasurement-history) · Partial map · [Source declaration](https://github.com/lanej/charmed/blob/49f755da5bd25a044043cf98bc303185c2ebce6d/pattern/measurement_history.go#L36)

- [Grid inspector](https://lanej.io/charmed/dependencies.html#component-collection%2Fgrid-inspector) · composes · inspection

Consumer usage uses

- [Measurement history](https://lanej.io/charmed/dependencies.html#component-pattern%2Fmeasurement-history) · primary demonstration

[Example source](https://github.com/lanej/charmed/blob/49f755da5bd25a044043cf98bc303185c2ebce6d/docs/snippets/pattern_measurement_history.go#L11)

[Organisms](https://lanej.io/charmed/organisms.html) / Patterns

## [List inspector](https://lanej.io/charmed/organisms-pattern.html#pattern%2flist-inspector)

Browse structured summaries and inspect the selected record’s grouped facts, conditions and history. Wide views preview details beside the list; narrow views retain selection and open complete inspection with Enter.

### [Native preview](https://lanej.io/charmed/organisms-pattern.html#pattern%2flist-inspector-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
┌ Operations ──────────────────────────┐  ┌ Evidence ──────────────────────────────────────────────────┐
│ Network                              │  │ Operation route-0042                                       │
│ → Gateway route update               │  │                                                            │
│    Outcome unknown                   │  │ Gateway route update                                       │
│   Next: Inspect the saved route      │  │  Outcome unknown                                           │
│   request.                           │  │ Next: Inspect the saved route request.                     │
│   Observed: 09:42 UTC                │  │ Observed  09:42 UTC                                        │
│                                      │  │                                                            │
│ Compute                              │  │ Observation  ───────────────────────────────────────────── │
│   Worker capacity                    │  │ Source   Controller journal                                │
│    Outcome confirmed                 │  │ Receipt  ∅ missing                                         │
│   Next: Check queue latency.         │  │                                                            │
│   Capacity: 12 workers observed      │  │ Verification  ──────────────────────────────────────────── │
│                                      │  │ Submission    Recorded                                     │
│                                      │  │ Remote route  Not observed                                 │
│                                      │  │                                                            │
│                                      │  │ Activity  ──────────────────────────────────────────────── │
│                                      │  │ ◷ Response timed out                                       │
│                                      │  │   09:42 UTC  Controller                                    │
└──────────────────────────────────────┘  └────────────────────────────────────────────────────────────┘
j/k select  enter inspect  esc back
```

Demo · truecolor · 112 × 28 terminal cells

```
┌ Operations ──────────────────────────┐  ┌ Evidence ──────────────────────────────────────────────────────────┐
│ Network                              │  │ Operation route-2026-11-23-0042                                    │
│ → Gateway route update               │  │                                                                    │
│    Outcome unknown                   │  │ Gateway route update                                               │
│   Next: Inspect the saved route      │  │  Outcome unknown                                                   │
│   request before changing traffic.   │  │ Next: Inspect the saved route request before changing traffic.     │
│   Observed: 09:42 UTC                │  │ Observed  09:42 UTC                                                │
│                                      │  │                                                                    │
│ Storage                              │  │ Observation  ───────────────────────────────────────────────────── │
│   Cache warm-up                      │  │ Domain           Network                                           │
│    In progress                       │  │ Source           Controller journal                                │
│   Next: Wait for the final cache     │  │ Observed         09:42 UTC                                         │
│   sample.                            │  │ Request receipt  ∅ missing                                         │
│   Observed: 09:42 UTC                │  │                                                                    │
│                                      │  │ Verification  ──────────────────────────────────────────────────── │
│ Compute                              │  │ Submission    Recorded                                             │
│   Worker capacity                    │  │   The request reached the controller journal.                      │
│    Outcome confirmed                 │  │ Remote route  Not observed                                         │
│   Next: Compare queue latency with   │  │   No read follows the submitted request.                           │
│   the service objective.             │  │ Attribution   Unknown                                              │
│   Observed: 09:41 UTC                │  │   The lost response cannot establish the effect.                   │
│                                      │  │                                                                    │
│ Search                               │  │ Observed evidence  ─────────────────────────────────────────────── │
│   Indexer restart                    │  │ The controller recorded a request to switch the gateway route. Its │
│    Success reported                  │  │ response timed out after submission. The remote route has not been │
│   Next: Read the process observer    │  │ observed since the request.                                        │
└──────────────────────────────────────┘  └────────────────────────────────────────────────────────────────────┘
j/k select  enter inspect  [/] scroll
```

Narrow · truecolor · 40 × 16 terminal cells

```
┌ Operations ──────────────────────────┐
│ Network                              │
│ → Gateway route update               │
│    Outcome unknown                   │
│   Next: Inspect the saved route      │
│   request before changing traffic.   │
│   Observed: 09:42 UTC                │
│                                      │
│ Storage                              │
│   Cache warm-up                      │
│    In progress                       │
│   Next: Wait for the final cache     │
│   sample.                            │
│   Observed: 09:42 UTC                │
└──────────────────────────────────────┘
j/k select  enter inspect  [/] scroll
```

256 colors · ansi256 · 112 × 28 terminal cells

```
┌ Operations ──────────────────────────┐  ┌ Evidence ──────────────────────────────────────────────────────────┐
│ Network                              │  │ Operation route-2026-11-23-0042                                    │
│ → Gateway route update               │  │                                                                    │
│    Outcome unknown                   │  │ Gateway route update                                               │
│   Next: Inspect the saved route      │  │  Outcome unknown                                                   │
│   request before changing traffic.   │  │ Next: Inspect the saved route request before changing traffic.     │
│   Observed: 09:42 UTC                │  │ Observed  09:42 UTC                                                │
│                                      │  │                                                                    │
│ Storage                              │  │ Observation  ───────────────────────────────────────────────────── │
│   Cache warm-up                      │  │ Domain           Network                                           │
│    In progress                       │  │ Source           Controller journal                                │
│   Next: Wait for the final cache     │  │ Observed         09:42 UTC                                         │
│   sample.                            │  │ Request receipt  ∅ missing                                         │
│   Observed: 09:42 UTC                │  │                                                                    │
│                                      │  │ Verification  ──────────────────────────────────────────────────── │
│ Compute                              │  │ Submission    Recorded                                             │
│   Worker capacity                    │  │   The request reached the controller journal.                      │
│    Outcome confirmed                 │  │ Remote route  Not observed                                         │
│   Next: Compare queue latency with   │  │   No read follows the submitted request.                           │
│   the service objective.             │  │ Attribution   Unknown                                              │
│   Observed: 09:41 UTC                │  │   The lost response cannot establish the effect.                   │
│                                      │  │                                                                    │
│ Search                               │  │ Observed evidence  ─────────────────────────────────────────────── │
│   Indexer restart                    │  │ The controller recorded a request to switch the gateway route. Its │
│    Success reported                  │  │ response timed out after submission. The remote route has not been │
│   Next: Read the process observer    │  │ observed since the request.                                        │
└──────────────────────────────────────┘  └────────────────────────────────────────────────────────────────────┘
j/k select  enter inspect  [/] scroll
```

16 colors · ansi16 · 112 × 28 terminal cells

```
┌ Operations ──────────────────────────┐  ┌ Evidence ──────────────────────────────────────────────────────────┐
│ Network                              │  │ Operation route-2026-11-23-0042                                    │
│ → Gateway route update               │  │                                                                    │
│    Outcome unknown                   │  │ Gateway route update                                               │
│   Next: Inspect the saved route      │  │  Outcome unknown                                                   │
│   request before changing traffic.   │  │ Next: Inspect the saved route request before changing traffic.     │
│   Observed: 09:42 UTC                │  │ Observed  09:42 UTC                                                │
│                                      │  │                                                                    │
│ Storage                              │  │ Observation  ───────────────────────────────────────────────────── │
│   Cache warm-up                      │  │ Domain           Network                                           │
│    In progress                       │  │ Source           Controller journal                                │
│   Next: Wait for the final cache     │  │ Observed         09:42 UTC                                         │
│   sample.                            │  │ Request receipt  ∅ missing                                         │
│   Observed: 09:42 UTC                │  │                                                                    │
│                                      │  │ Verification  ──────────────────────────────────────────────────── │
│ Compute                              │  │ Submission    Recorded                                             │
│   Worker capacity                    │  │   The request reached the controller journal.                      │
│    Outcome confirmed                 │  │ Remote route  Not observed                                         │
│   Next: Compare queue latency with   │  │   No read follows the submitted request.                           │
│   the service objective.             │  │ Attribution   Unknown                                              │
│   Observed: 09:41 UTC                │  │   The lost response cannot establish the effect.                   │
│                                      │  │                                                                    │
│ Search                               │  │ Observed evidence  ─────────────────────────────────────────────── │
│   Indexer restart                    │  │ The controller recorded a request to switch the gateway route. Its │
│    Success reported                  │  │ response timed out after submission. The remote route has not been │
│   Next: Read the process observer    │  │ observed since the request.                                        │
└──────────────────────────────────────┘  └────────────────────────────────────────────────────────────────────┘
j/k select  enter inspect  [/] scroll
```

Monochrome · none · 112 × 28 terminal cells

```
+ Operations --------------------------+  + Evidence ----------------------------------------------------------+
| Network                              |  | Operation route-2026-11-23-0042                                    |
| > Gateway route update               |  |                                                                    |
|   [Outcome unknown]                  |  | Gateway route update                                               |
|   Next: Inspect the saved route      |  | [Outcome unknown]                                                  |
|   request before changing traffic.   |  | Next: Inspect the saved route request before changing traffic.     |
|   Observed: 09:42 UTC                |  | Observed  09:42 UTC                                                |
|                                      |  |                                                                    |
| Storage                              |  | Observation  ----------------------------------------------------- |
|   Cache warm-up                      |  | Domain           Network                                           |
|   [In progress]                      |  | Source           Controller journal                                |
|   Next: Wait for the final cache     |  | Observed         09:42 UTC                                         |
|   sample.                            |  | Request receipt  missing                                           |
|   Observed: 09:42 UTC                |  |                                                                    |
|                                      |  | Verification  ---------------------------------------------------- |
| Compute                              |  | Submission    Recorded                                             |
|   Worker capacity                    |  |   The request reached the controller journal.                      |
|   [Outcome confirmed]                |  | Remote route  Not observed                                         |
|   Next: Compare queue latency with   |  |   No read follows the submitted request.                           |
|   the service objective.             |  | Attribution   Unknown                                              |
|   Observed: 09:41 UTC                |  |   The lost response cannot establish the effect.                   |
|                                      |  |                                                                    |
| Search                               |  | Observed evidence  ----------------------------------------------- |
|   Indexer restart                    |  | The controller recorded a request to switch the gateway route. Its |
|   [Success reported]                 |  | response timed out after submission. The remote route has not been |
|   Next: Read the process observer    |  | observed since the request.                                        |
+--------------------------------------+  +--------------------------------------------------------------------+
j/k select  enter inspect  [/] scroll
```

Inspecting · truecolor · 112 × 28 terminal cells

```
┌ Evidence ────────────────────────────────────────────────────────────────────────────────────────────────────┐
│ Operation route-2026-11-23-0042                                                                              │
│                                                                                                              │
│ Gateway route update                                                                                         │
│  Outcome unknown                                                                                             │
│ Next: Inspect the saved route request before changing traffic.                                               │
│ Observed  09:42 UTC                                                                                          │
│                                                                                                              │
│ Observation  ───────────────────────────────────────    Verification  ────────────────────────────────────── │
│ Domain: Network                                         Submission: Recorded                                 │
│ Source: Controller journal                                The request reached the controller journal.        │
│ Observed: 09:42 UTC                                     Remote route: Not observed                           │
│ Request receipt: ∅ missing                                No read follows the submitted request.             │
│                                                         Attribution: Unknown                                 │
│                                                           The lost response cannot establish the effect.     │
│                                                                                                              │
│ Observed evidence  ───────────────────────────────────────────────────────────────────────────────────────── │
│ The controller recorded a request to switch the gateway route. Its response timed out after submission. The  │
│ remote route has not been observed since the request.                                                        │
│                                                                                                              │
│ Activity  ────────────────────────────────────────────────────────────────────────────────────────────────── │
│ ◷ Route request recorded                                                                                     │
│   09:41 UTC  Controller                                                                                      │
│   Switch the gateway route to the prepared target.                                                           │
│                                                                                                              │
│ ◷ Response timed out                                                                                         │
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
│ Network                              │  │ Operation route-2026-11-23-0042                                    │
│ → Gateway route update               │  │                                                                    │
│    Outcome unknown                   │  │ Gateway route update                                               │
│   Next: Inspect the saved route      │  │  Refresh failed  Last observation retained.                        │
│   request before changing traffic.   │  │  Outcome unknown                                                   │
│   Observed: 09:42 UTC                │  │ Next: Inspect the saved route request before changing traffic.     │
│                                      │  │ Observed  09:42 UTC                                                │
│ Storage                              │  │                                                                    │
│   Cache warm-up                      │  │ Observation  ───────────────────────────────────────────────────── │
│    In progress                       │  │ Domain           Network                                           │
│   Next: Wait for the final cache     │  │ Source           Controller journal                                │
│   sample.                            │  │ Observed         09:42 UTC                                         │
│   Observed: 09:42 UTC                │  │ Request receipt  ∅ missing                                         │
│                                      │  │                                                                    │
│ Compute                              │  │ Verification  ──────────────────────────────────────────────────── │
│   Worker capacity                    │  │ Submission    Recorded                                             │
│    Outcome confirmed                 │  │   The request reached the controller journal.                      │
│   Next: Compare queue latency with   │  │ Remote route  Not observed                                         │
│   the service objective.             │  │   No read follows the submitted request.                           │
│   Observed: 09:41 UTC                │  │ Attribution   Unknown                                              │
│                                      │  │   The lost response cannot establish the effect.                   │
│ Search                               │  │                                                                    │
│   Indexer restart                    │  │ Observed evidence  ─────────────────────────────────────────────── │
│    Success reported                  │  │ The controller recorded a request to switch the gateway route. Its │
│   Next: Read the process observer    │  │ response timed out after submission. The remote route has not been │
└──────────────────────────────────────┘  └────────────────────────────────────────────────────────────────────┘
j/k select  enter inspect  [/] scroll
```

Confirmed · truecolor · 112 × 28 terminal cells

```
┌ Operations ──────────────────────────┐  ┌ Evidence ──────────────────────────────────────────────────────────┐
│ Network                              │  │ Operation route-2026-11-23-0042                                    │
│ → Gateway route update               │  │                                                                    │
│    Outcome confirmed                 │  │ Gateway route update                                               │
│   Next: Observe request errors       │  │  Outcome confirmed                                                 │
│   before closing the incident.       │  │ Next: Observe request errors before closing the incident.          │
│   Observed: 09:43 UTC                │  │ Observed  09:43 UTC                                                │
│                                      │  │                                                                    │
│ Storage                              │  │ Observation  ───────────────────────────────────────────────────── │
│   Cache warm-up                      │  │ Domain           Network                                           │
│    In progress                       │  │ Source           Route observer                                    │
│   Next: Wait for the final cache     │  │ Observed         09:43 UTC                                         │
│   sample.                            │  │ Request receipt  ∅ missing                                         │
│   Observed: 09:42 UTC                │  │                                                                    │
│                                      │  │ Verification  ──────────────────────────────────────────────────── │
│ Compute                              │  │ Submission    Recorded                                             │
│   Worker capacity                    │  │   The request reached the controller journal.                      │
│    Outcome confirmed                 │  │ Remote route  Requested route active                               │
│   Next: Compare queue latency with   │  │   Independently observed at 09:43 UTC.                             │
│   the service objective.             │  │ Attribution   Unknown                                              │
│   Observed: 09:41 UTC                │  │   The lost response cannot establish the effect.                   │
│                                      │  │                                                                    │
│ Search                               │  │ Observed evidence  ─────────────────────────────────────────────── │
│   Indexer restart                    │  │ A later read found the requested gateway route active. The route   │
│    Success reported                  │  │ outcome is confirmed; the lost response still prevents attributing │
│   Next: Read the process observer    │  │ the change to this exact request.                                  │
└──────────────────────────────────────┘  └────────────────────────────────────────────────────────────────────┘
j/k select  enter inspect  [/] scroll
```

ASCII · truecolor · 112 × 28 terminal cells

```
+ Operations --------------------------+  + Evidence ----------------------------------------------------------+
| Network                              |  | Operation route-2026-11-23-0042                                    |
| > Gateway route update               |  |                                                                    |
|   [Outcome unknown]                  |  | Gateway route update                                               |
|   Next: Inspect the saved route      |  | [Outcome unknown]                                                  |
|   request before changing traffic.   |  | Next: Inspect the saved route request before changing traffic.     |
|   Observed: 09:42 UTC                |  | Observed  09:42 UTC                                                |
|                                      |  |                                                                    |
| Storage                              |  | Observation  ----------------------------------------------------- |
|   Cache warm-up                      |  | Domain           Network                                           |
|   [In progress]                      |  | Source           Controller journal                                |
|   Next: Wait for the final cache     |  | Observed         09:42 UTC                                         |
|   sample.                            |  | Request receipt  missing                                           |
|   Observed: 09:42 UTC                |  |                                                                    |
|                                      |  | Verification  ---------------------------------------------------- |
| Compute                              |  | Submission    Recorded                                             |
|   Worker capacity                    |  |   The request reached the controller journal.                      |
|   [Outcome confirmed]                |  | Remote route  Not observed                                         |
|   Next: Compare queue latency with   |  |   No read follows the submitted request.                           |
|   the service objective.             |  | Attribution   Unknown                                              |
|   Observed: 09:41 UTC                |  |   The lost response cannot establish the effect.                   |
|                                      |  |                                                                    |
| Search                               |  | Observed evidence  ----------------------------------------------- |
|   Indexer restart                    |  | The controller recorded a request to switch the gateway route. Its |
|   [Success reported]                 |  | response timed out after submission. The remote route has not been |
|   Next: Read the process observer    |  | observed since the request.                                        |
+--------------------------------------+  +--------------------------------------------------------------------+
j/k select  enter inspect  [/] scroll
```

Unicode monochrome · none · 112 × 28 terminal cells

```
+ Operations --------------------------+  + Evidence ----------------------------------------------------------+
| Network                              |  | Operation route-2026-11-23-0042                                    |
| > Gateway route update               |  |                                                                    |
|   [Outcome unknown]                  |  | Gateway route update                                               |
|   Next: Inspect the saved route      |  | [Outcome unknown]                                                  |
|   request before changing traffic.   |  | Next: Inspect the saved route request before changing traffic.     |
|   Observed: 09:42 UTC                |  | Observed  09:42 UTC                                                |
|                                      |  |                                                                    |
| Storage                              |  | Observation  ----------------------------------------------------- |
|   Cache warm-up                      |  | Domain           Network                                           |
|   [In progress]                      |  | Source           Controller journal                                |
|   Next: Wait for the final cache     |  | Observed         09:42 UTC                                         |
|   sample.                            |  | Request receipt  missing                                           |
|   Observed: 09:42 UTC                |  |                                                                    |
|                                      |  | Verification  ---------------------------------------------------- |
| Compute                              |  | Submission    Recorded                                             |
|   Worker capacity                    |  |   The request reached the controller journal.                      |
|   [Outcome confirmed]                |  | Remote route  Not observed                                         |
|   Next: Compare queue latency with   |  |   No read follows the submitted request.                           |
|   the service objective.             |  | Attribution   Unknown                                              |
|   Observed: 09:41 UTC                |  |   The lost response cannot establish the effect.                   |
|                                      |  |                                                                    |
| Search                               |  | Observed evidence  ----------------------------------------------- |
|   Indexer restart                    |  | The controller recorded a request to switch the gateway route. Its |
|   [Success reported]                 |  | response timed out after submission. The remote route has not been |
|   Next: Read the process observer    |  | observed since the request.                                        |
+--------------------------------------+  +--------------------------------------------------------------------+
j/k select  enter inspect  [/] scroll
```

### [Go usage](https://lanej.io/charmed/organisms-pattern.html#pattern%2flist-inspector-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`pattern`](https://pkg.go.dev/github.com/lanej/charmed/pattern "github.com/lanej/charmed/pattern") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
summary := layout.RecordSummary{
	Title: "Gateway route update", Status: text.Pill("Outcome unknown", text.Warning, palette, text.Normal),
	Fields: []layout.Pair{{Label: "Observed", Value: "09:42 UTC"}},
	Next:   &layout.Pair{Label: "Next", Value: "Inspect the saved route request."},
}
document := layout.DetailDocument{
	Header: layout.IdentityHeader{Kind: "Operation", ID: "route-0042"}, Summary: summary,
	Sections: []layout.DetailSection{
		{ID: "source", Title: "Observation", Fields: []layout.Pair{{Label: "Source", Value: "Controller journal"}, {Label: "Receipt", State: text.ValueMissing}}},
		{ID: "checks", Title: "Verification", Statuses: []layout.StatusItem{{Label: "Submission", Value: "Recorded"}, {Label: "Remote route", Value: "Not observed"}}},
		{ID: "history", Title: "Activity", FullWidth: true, Timeline: []layout.TimelineEvent{
			{When: "09:42 UTC", Actor: "Controller", Title: "Response timed out", Detail: "Submission is recorded; the external result remains unknown."},
		}},
	},
}
records := []pattern.InspectorRecord{
	{ID: "route-0042", Heading: "Network", Summary: &summary, Document: &document},
	{ID: "pool-0007", Heading: "Compute", Summary: &layout.RecordSummary{
		Title: "Worker capacity", Status: text.Pill("Outcome confirmed", text.Success, palette, text.Normal),
		Fields: []layout.Pair{{Label: "Capacity", Value: "12 workers observed"}},
		Next:   &layout.Pair{Label: "Next", Value: "Check queue latency."},
	}},
}

inspector, err := pattern.NewListInspector(records, layout.Size{Width: 104, Height: 20}, palette)
if err != nil {
	return "", err
}

inspector.ListTitle, inspector.DetailTitle = "Operations", "Evidence"
inspector.Refresh()

return layout.Column(inspector.View(), text.KeyHint("j/k", "select", palette)+"  "+text.KeyHint("enter", "inspect", palette)+"  "+text.KeyHint("esc", "back", palette)), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-pattern.html#pattern%2flist-inspector-integration)

Supply stable record IDs with Summary and Document; Row and Detail callbacks override the typed presentation when needed. SetRecords validates atomically and copies typed snapshots. Structured details retain section-relative scrolling; selection changes reset details. Enter inspects, Escape returns and brackets scroll. Resize preserves native styles and bindings; Refresh reflows current content. Rebuild records to change pre-styled values. Preview controls the wide detail pane. Applications own observations, ordering, actions, refresh and quit. Run go run ./examples/inspection for the complete consumer.

### [Run this example](https://lanej.io/charmed/organisms-pattern.html#pattern%2flist-inspector-run)

```sh
go run ./examples/catalog pattern/list-inspector
```

[Example source](https://github.com/lanej/charmed/blob/49f755da5bd25a044043cf98bc303185c2ebce6d/docs/snippets/pattern_list_inspector.go) · [Implementation](https://github.com/lanej/charmed/blob/49f755da5bd25a044043cf98bc303185c2ebce6d/pattern/list_inspector.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [collection/viewport](https://lanej.io/charmed/atoms-collection.html#collection/viewport) [layout/record-summary](https://lanej.io/charmed/molecules-layout.html#layout/record-summary) [layout/detail-document](https://lanej.io/charmed/organisms-layout.html#layout/detail-document) [foundations/layout](https://lanej.io/charmed/foundations-layout.html#foundations/layout)
