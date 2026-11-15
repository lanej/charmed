[Molecules](https://lanej.io/charmed/molecules.html)Diagrams

# [Diagrams](https://lanej.io/charmed/molecules-diagram.html#level-molecules-diagram)

1 component · [All molecules](https://lanej.io/charmed/molecules.html)

- [State transitions](https://lanej.io/charmed/molecules-diagram.html#diagram%2fstate-transitions)

[Molecules](https://lanej.io/charmed/molecules.html) / Diagrams

## [State transitions](https://lanej.io/charmed/molecules-diagram.html#diagram%2fstate-transitions)

Labeled arrows connect state boxes directly, with distinct initial and terminal borders, optional success/error outcomes, and vertical layout for narrow regions.

### [Native preview](https://lanej.io/charmed/molecules-diagram.html#diagram%2fstate-transitions-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

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

Narrow · truecolor · 40 × 40 terminal cells

```
State transitions                       
                                        
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
  ◉  Last transition                    
  ○  Last three                         
  ○  Complete trace                     
                                        
Left/Right · Home/End · Tab trace window
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

Last 3 transitions · truecolor · 120 × 40 terminal cells

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
  ○  Last transition                                                                                                    
  ◉  Last three                                                                                                         
  ○  Complete trace                                                                                                     
                                                                                                                        
Left/Right · Home/End · Tab trace window
```

Whole path · truecolor · 120 × 40 terminal cells

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
  ○  Last transition                                                                                                    
  ○  Last three                                                                                                         
  ◉  Complete trace                                                                                                     
                                                                                                                        
Left/Right · Home/End · Tab trace window
```

### [Go usage](https://lanej.io/charmed/molecules-diagram.html#diagram%2fstate-transitions-code)

Imports: [`diagram`](https://pkg.go.dev/github.com/lanej/charmed/diagram "github.com/lanej/charmed/diagram") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
graph := diagram.StateDiagram{
	States: []diagram.State{
		{ID: "idle", Label: "Idle", Initial: true},
		{ID: "running", Label: "Running"},
		{ID: "done", Label: "Done", Final: true, Outcome: diagram.OutcomeSuccess},
		{ID: "failed", Label: "Failed", Final: true, Outcome: diagram.OutcomeError},
	},
	Transitions: []diagram.Transition{
		{From: "idle", To: "running", Label: "start"},
		{From: "running", To: "running", Label: "retry"},
		{From: "running", To: "done", Label: "finish"},
		{From: "running", To: "idle", Label: "cancel"},
		{From: "running", To: "failed", Label: "fail"},
	},
	Active: "done",
	Trace: diagram.TransitionTrace{
		Steps: []int{0, 1, 3, 0, 2}, // start, retry, cancel, start, finish.
		Last:  1,                    // Last N steps; zero defaults to one. All: true shows the whole path.
	},
}

rendered, err := graph.View(64, palette)
if err != nil {
	return "", err
}
// Update Active and Trace.Steps when application state changes, then render again.
return rendered, nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-diagram.html#diagram%2fstate-transitions-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/molecules-diagram.html#diagram%2fstate-transitions-run)

```sh
go run ./examples/catalog diagram/state-transitions
```

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/diagram_state_transitions.go) · [Implementation](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/story/diagrams.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

### Reusable component

[State diagram dependency graph](https://lanej.io/charmed/dependencies.html#component-diagram%2Fstate-diagram) · Partial map · [Source declaration](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/diagram/state.go#L36)

Consumer usage uses

- [State diagram](https://lanej.io/charmed/dependencies.html#component-diagram%2Fstate-diagram) · primary demonstration

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/docs/snippets/diagram_state_transitions.go#L8)

Interactive fixture uses

- [State diagram](https://lanej.io/charmed/dependencies.html#component-diagram%2Fstate-diagram) · graph data
- [Trace replay](https://lanej.io/charmed/dependencies.html#component-diagram%2Ftrace-replay) · trace inspection
- [Radio group](https://lanej.io/charmed/dependencies.html#component-input%2Fradio-group) · history-window controls

[Example source](https://github.com/lanej/charmed/blob/651bab6833695213690f9cc2258ed202ceac56c9/story/diagrams.go#L12)
