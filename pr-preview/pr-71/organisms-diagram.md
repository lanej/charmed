[Organisms](https://lanej.io/charmed/organisms.html)Diagrams

# [Diagrams](https://lanej.io/charmed/organisms-diagram.html#level-organisms-diagram)

1 component · [All organisms](https://lanej.io/charmed/organisms.html)

- [Trace replay](https://lanej.io/charmed/organisms-diagram.html#diagram%2ftrace-replay)

[Organisms](https://lanej.io/charmed/organisms.html) / Diagrams

## [Trace replay](https://lanej.io/charmed/organisms-diagram.html#diagram%2ftrace-replay)

Inspect a complete supplied transition history one occurrence at a time. Repeated edges remain distinguishable without executing transitions or introducing a playback clock.

### [Native preview](https://lanej.io/charmed/organisms-diagram.html#diagram%2ftrace-replay-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
╭─────────────────╮            ┏━━━━━━━━━┓            ╔═══════════════╗
│ idle            ├━━ start ━━▶┫ running ┣── finish ─▶╢ done          ║
│ initial         ├◀─ cancel ──┫ active  ┃            ║ final         ║
╰─────────────────╯            ┗━━━━━━━━━┛            ╚═══════════════╝
```

Demo · truecolor · 120 × 40 terminal cells

```
State transitions                                                                                                       
                                                                                                                        
╭─────────────────╮            ┌─────────────┐            ╔═════════════════╗   ╔═══════════════╗                       
│ Idle            ├── start ──▶┤ Running     ├━━ finish ━▶╢ Done            ║   ║ Failed        ║                       
│ initial         ├◀─ cancel ──┤             │            ║ success, active ║   ║ error         ║                       
╰─────────────────╯            └┬─┬─────────┬┘            ╚═════════════════╝   ╚╤══════════════╝                       
                                │ │         ▲                                    ▲                                      
                                │ └─ retry ─┘                                    │                                      
                                │                                                │                                      
                                └──────────────── fail ──────────────────────────┘                                      
  ◉  Last transition                                                                                                    
  ○  Last three                                                                                                         
  ○  Complete trace                                                                                                     
                                                                                                                        
Left/Right · Home/End · Tab trace window
```

Narrow · truecolor · 40 × 16 terminal cells

```
trace replay needs 31 rows; available 8
```

256 colors · ansi256 · 120 × 40 terminal cells

```
State transitions                                                                                                       
                                                                                                                        
╭─────────────────╮            ┌─────────────┐            ╔═════════════════╗   ╔═══════════════╗                       
│ Idle            ├── start ──▶┤ Running     ├━━ finish ━▶╢ Done            ║   ║ Failed        ║                       
│ initial         ├◀─ cancel ──┤             │            ║ success, active ║   ║ error         ║                       
╰─────────────────╯            └┬─┬─────────┬┘            ╚═════════════════╝   ╚╤══════════════╝                       
                                │ │         ▲                                    ▲                                      
                                │ └─ retry ─┘                                    │                                      
                                │                                                │                                      
                                └──────────────── fail ──────────────────────────┘                                      
  ◉  Last transition                                                                                                    
  ○  Last three                                                                                                         
  ○  Complete trace                                                                                                     
                                                                                                                        
Left/Right · Home/End · Tab trace window
```

16 colors · ansi16 · 120 × 40 terminal cells

```
State transitions                                                                                                       
                                                                                                                        
╭─────────────────╮            ┌─────────────┐            ╔═════════════════╗   ╔═══════════════╗                       
│ Idle            ├── start ──▶┤ Running     ├━━ finish ━▶╢ Done            ║   ║ Failed        ║                       
│ initial         ├◀─ cancel ──┤             │            ║ success, active ║   ║ error         ║                       
╰─────────────────╯            └┬─┬─────────┬┘            ╚═════════════════╝   ╚╤══════════════╝                       
                                │ │         ▲                                    ▲                                      
                                │ └─ retry ─┘                                    │                                      
                                │                                                │                                      
                                └──────────────── fail ──────────────────────────┘                                      
  ◉  Last transition                                                                                                    
  ○  Last three                                                                                                         
  ○  Complete trace                                                                                                     
                                                                                                                        
Left/Right · Home/End · Tab trace window
```

Monochrome · none · 120 × 40 terminal cells

```
State transitions                                                                                                       
                                                                                                                        
(-----------------)            +-------------+            #=================#   #===============#                       
| Idle            +-- start -->+ Running     +== finish =># Done            |   | Failed        |                       
| initial         +<- cancel --+             |            | success, active |   | error         |                       
(-----------------)            ++-+---------++            #=================#   ##==============#                       
                                | |         ^                                    ^                                      
                                | +- retry -+                                    |                                      
                                |                                                |                                      
                                +---------------- fail --------------------------+                                      
 (*) Last transition                                                                                                    
 ( ) Last three                                                                                                         
 ( ) Complete trace                                                                                                     
                                                                                                                        
Left/Right * Home/End * Tab trace window
```

### [Go usage](https://lanej.io/charmed/organisms-diagram.html#diagram%2ftrace-replay-code)

Imports: [`diagram`](https://pkg.go.dev/github.com/lanej/charmed/diagram "github.com/lanej/charmed/diagram") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
graph := diagram.StateDiagram{
	States:      []diagram.State{{ID: "idle", Initial: true}, {ID: "running"}, {ID: "done", Final: true}},
	Transitions: []diagram.Transition{{From: "idle", To: "running", Label: "start"}, {From: "running", To: "idle", Label: "cancel"}, {From: "running", To: "done", Label: "finish"}},
	Trace:       diagram.TransitionTrace{Steps: []int{0, 1, 0, 2}},
}

replay, err := diagram.NewTraceReplay(graph, layout.Size{Width: 80, Height: 24}, theme.New(theme.Dark))
if err != nil {
	return "", err
}

replay.SetPosition(3)
// Repeated edge zero is selected by occurrence three, independently of edge index.
// Left/Right stop at either end. Home/End jump; no clock runs.
return replay.View()
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-diagram.html#diagram%2ftrace-replay-integration)

Construct from a validated diagram with its complete trace, then forward navigation messages only while active. SetPosition clamps to the supplied history; zero precedes the first transition. Left and Right stop at either end, while Home and End jump. Set Last or All for visible history. Resize preserves inspection position; View returns width and height errors instead of silently clipping topology.

### [Run this example](https://lanej.io/charmed/organisms-diagram.html#diagram%2ftrace-replay-run)

```sh
go run ./examples/catalog diagram/trace-replay
```

[Example source](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/docs/snippets/diagram_trace_replay.go) · [Implementation](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/diagram/trace_replay.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

### Reusable component

[Trace replay dependency graph](https://lanej.io/charmed/dependencies.html#component-diagram%2Ftrace-replay) · Partial map · [Source declaration](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/diagram/trace_replay.go#L21)

- [State diagram](https://lanej.io/charmed/dependencies.html#component-diagram%2Fstate-diagram) · composes · diagram

Consumer usage uses

- [Trace replay](https://lanej.io/charmed/dependencies.html#component-diagram%2Ftrace-replay) · primary demonstration
- [State diagram](https://lanej.io/charmed/dependencies.html#component-diagram%2Fstate-diagram) · graph data

[Example source](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/docs/snippets/diagram_trace_replay.go#L9)

Interactive fixture uses

- [State diagram](https://lanej.io/charmed/dependencies.html#component-diagram%2Fstate-diagram) · graph data
- [Trace replay](https://lanej.io/charmed/dependencies.html#component-diagram%2Ftrace-replay) · trace inspection
- [Radio group](https://lanej.io/charmed/dependencies.html#component-input%2Fradio-group) · history-window controls

[Example source](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/story/diagrams.go#L12)
