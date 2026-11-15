[Templates](https://lanej.io/charmed/templates.html)Patterns

# [Patterns](https://lanej.io/charmed/templates-pattern.html#level-templates-pattern)

2 components · [All templates](https://lanej.io/charmed/templates.html)

- [Application shell](https://lanej.io/charmed/templates-pattern.html#template%2fapplication)
- [Report page](https://lanej.io/charmed/templates-pattern.html#template%2freport-page)

[Templates](https://lanej.io/charmed/templates.html) / Patterns

## [Application shell](https://lanej.io/charmed/templates-pattern.html#template%2fapplication)

A header, tab bar, active content panel and footer compose a resizable application shell. Left/Right changes the visible section.

### [Native preview](https://lanej.io/charmed/templates-pattern.html#template%2fapplication-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Service workbench                                               
Production / us-west-2                                          
  Overview  Services                                            
┌ Overview ────────────────────────────────────────────────────┐
│ 12 services / all healthy                                    │
│                                                              │
│                                                              │
│                                                              │
│                                                              │
│                                                              │
└──────────────────────────────────────────────────────────────┘
arrows tabs                                         ● CONNECTED 
```

Demo · truecolor · 64 × 12 terminal cells

```
Service workbench                                      SYNTHETIC
Production / us-west-2                                          
  Overview  Services  Activity  Settings                        
┌──────────────────────────────────────────────────────────────┐
│ 12 services / all healthy                                    │
│ Select a tab to inspect another section.                     │
│                                                              │
│                                                              │
│                                                              │
│                                                              │
└──────────────────────────────────────────────────────────────┘
arrows tabs  Ctrl+C quit                            ● CONNECTED 
```

Narrow · truecolor · 40 × 16 terminal cells

```
Service workbench              SYNTHETIC
Production / us-west-2                  
  Overview  Services  Activity         >
┌──────────────────────────────────────┐
│ 12 services / all healthy            │
│ Select a tab to inspect another sect │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
└──────────────────────────────────────┘
arrows tabs  Ctrl+C quit    ● CONNECTED 
```

256 colors · ansi256 · 64 × 12 terminal cells

```
Service workbench                                      SYNTHETIC
Production / us-west-2                                          
  Overview  Services  Activity  Settings                        
┌──────────────────────────────────────────────────────────────┐
│ 12 services / all healthy                                    │
│ Select a tab to inspect another section.                     │
│                                                              │
│                                                              │
│                                                              │
│                                                              │
└──────────────────────────────────────────────────────────────┘
arrows tabs  Ctrl+C quit                            ● CONNECTED 
```

16 colors · ansi16 · 64 × 12 terminal cells

```
Service workbench                                      SYNTHETIC
Production / us-west-2                                          
  Overview  Services  Activity  Settings                        
┌──────────────────────────────────────────────────────────────┐
│ 12 services / all healthy                                    │
│ Select a tab to inspect another section.                     │
│                                                              │
│                                                              │
│                                                              │
│                                                              │
└──────────────────────────────────────────────────────────────┘
arrows tabs  Ctrl+C quit                            ● CONNECTED 
```

Monochrome · none · 64 × 12 terminal cells

```
Service workbench                                      SYNTHETIC
Production / us-west-2                                          
  Overview  Services  Activity  Settings                        
+--------------------------------------------------------------+
| 12 services / all healthy                                    |
| Select a tab to inspect another section.                     |
|                                                              |
|                                                              |
|                                                              |
|                                                              |
+--------------------------------------------------------------+
arrows tabs  Ctrl+C quit                           [+ CONNECTED]
```

### [Go usage](https://lanej.io/charmed/templates-pattern.html#template%2fapplication-code)

Imports: [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`navigation`](https://pkg.go.dev/github.com/lanej/charmed/navigation "github.com/lanej/charmed/navigation") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
tabs := navigation.NewTabs([]navigation.Tab{
	{ID: "overview", Label: "Overview"}, {ID: "services", Label: "Services"},
}, palette)
tabs.SetWidth(64)

header := layout.Header{Title: "Service workbench", Context: "Production / us-west-2"}
footer := layout.Footer{Help: "arrows tabs", Status: feedback.Connected.View(text.Normal, palette)}
body := layout.Panel("Overview", "12 services / all healthy",
	layout.Size{Width: 64, Height: 8}, palette, false)

return layout.Column(header.View(64, palette), tabs.View(), body, footer.View(64, palette))
```

### [Updates, focus & sizing](https://lanej.io/charmed/templates-pattern.html#template%2fapplication-integration)

The parent model supplies observations, coordinates focus and owns application decisions. Forward initialization and update commands rather than executing them during rendering. Resize the existing composition from the terminal allocation so selection remains stable, and apply native customization at the exposed child models after any rebuilding lifecycle step.

### [Run this example](https://lanej.io/charmed/templates-pattern.html#template%2fapplication-run)

```sh
go run ./examples/catalog template/application
```

[Example source](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/template_application.go) · [Implementation](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/story/chrome.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [layout/header](https://lanej.io/charmed/molecules-layout.html#layout/header) [navigation/tabs](https://lanej.io/charmed/molecules-navigation.html#navigation/tabs) [layout/footer](https://lanej.io/charmed/molecules-layout.html#layout/footer) [foundations/layout](https://lanej.io/charmed/foundations-layout.html#foundations/layout) [feedback/status](https://lanej.io/charmed/atoms-feedback.html#feedback/status)

[Templates](https://lanej.io/charmed/templates.html) / Patterns

## [Report page](https://lanej.io/charmed/templates-pattern.html#template%2freport-page)

A finite report: header with the overall evidence state, a summary line, titled panels that sit side by side or stack at 55 columns and narrower, and the evidence footer. The overall state is the worst counted state; no evidence reads unknown.

### [Native preview](https://lanej.io/charmed/templates-pattern.html#template%2freport-page-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Service processes                                                       ◷ STALE 
api                                                                             
◷ 1 stale · ✓ 2 healthy
┌ Hosts ────────────────────────────────┐ ┌ Processes ─────────────────────────┐
│ web1west  observed  45s               │ │ db:db  RUNNING                     │
└───────────────────────────────────────┘ └────────────────────────────────────┘
source: supervisor · scope: api ·  ◷ STALE  14s ago
```

Demo · truecolor · 100 × 24 terminal cells

```
Service processes                                                                         ✓ HEALTHY 
api                                                                                                 
✓ 2 healthy                                                                                         
┌ Hosts ────────────────────────────────────────────┐ ┌ Processes ─────────────────────────────────┐
│ web1west  observed  live    45s                   │ │ db:db       RUNNING                        │
│ web2west  observed  cached  4m                    │ │ http:http  RUNNING                         │
│ worker1   observed  live    30s                   │ │ worker     FATAL                           │
└───────────────────────────────────────────────────┘ └────────────────────────────────────────────┘
┌ Limits ──────────────────────────────────────────────────────────────────────────────────────────┐
│ 1 host not retained                                                                              │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
source: supervisor · scope: api ·  ✓ HEALTHY  14s ago
```

Narrow · truecolor · 40 × 16 terminal cells

```
Service processes             ✓ HEALTHY 
api                                     
✓ 2 healthy                             
┌ Hosts ───────────────────────────────┐
│ web1west  observed  live    45s      │
│ web2west  observed  cached  4m       │
│ worker1   observed  live    30s      │
└──────────────────────────────────────┘
┌ Processes ───────────────────────────┐
│ db:db       RUNNING                  │
│ http:http  RUNNING                   │
│ worker     FATAL                     │
└──────────────────────────────────────┘
… +3 more                               
source: supervisor · scope: api         
 ✓ HEALTHY  14s ago
```

256 colors · ansi256 · 100 × 24 terminal cells

```
Service processes                                                                         ✓ HEALTHY 
api                                                                                                 
✓ 2 healthy                                                                                         
┌ Hosts ────────────────────────────────────────────┐ ┌ Processes ─────────────────────────────────┐
│ web1west  observed  live    45s                   │ │ db:db       RUNNING                        │
│ web2west  observed  cached  4m                    │ │ http:http  RUNNING                         │
│ worker1   observed  live    30s                   │ │ worker     FATAL                           │
└───────────────────────────────────────────────────┘ └────────────────────────────────────────────┘
┌ Limits ──────────────────────────────────────────────────────────────────────────────────────────┐
│ 1 host not retained                                                                              │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
source: supervisor · scope: api ·  ✓ HEALTHY  14s ago
```

16 colors · ansi16 · 100 × 24 terminal cells

```
Service processes                                                                         ✓ HEALTHY 
api                                                                                                 
✓ 2 healthy                                                                                         
┌ Hosts ────────────────────────────────────────────┐ ┌ Processes ─────────────────────────────────┐
│ web1west  observed  live    45s                   │ │ db:db       RUNNING                        │
│ web2west  observed  cached  4m                    │ │ http:http  RUNNING                         │
│ worker1   observed  live    30s                   │ │ worker     FATAL                           │
└───────────────────────────────────────────────────┘ └────────────────────────────────────────────┘
┌ Limits ──────────────────────────────────────────────────────────────────────────────────────────┐
│ 1 host not retained                                                                              │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
source: supervisor · scope: api ·  ✓ HEALTHY  14s ago
```

Monochrome · none · 100 × 24 terminal cells

```
Service processes                                                                        [+ HEALTHY]
api                                                                                                 
+ 2 healthy                                                                                         
+ Hosts --------------------------------------------+ + Processes ---------------------------------+
| web1west  observed  live    45s                   | | db:db       RUNNING                        |
| web2west  observed  cached  4m                    | | http:http  RUNNING                         |
| worker1   observed  live    30s                   | | worker     FATAL                           |
+---------------------------------------------------+ +--------------------------------------------+
+ Limits ------------------------------------------------------------------------------------------+
| 1 host not retained                                                                              |
+--------------------------------------------------------------------------------------------------+
source: supervisor * scope: api * [+ HEALTHY] 14s ago
```

### [Go usage](https://lanej.io/charmed/templates-pattern.html#template%2freport-page-code)

Imports: [`time`](https://pkg.go.dev/time "time") · [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`pattern`](https://pkg.go.dev/github.com/lanej/charmed/pattern "github.com/lanej/charmed/pattern") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
now := time.Date(2026, 11, 11, 12, 0, 0, 0, time.UTC)
page := pattern.ReportPage{
	Title: "Service processes", Context: "api",
	Counts: []feedback.EvidenceCount{{State: feedback.EvidenceHealthy, N: 2}, {State: feedback.EvidenceStale, N: 1}},
	Panels: []pattern.ReportPanel{
		{Title: "Hosts", Weight: 3, Body: func(int) string { return "web1west  observed  45s" }},
		{Title: "Processes", Weight: 2, Beside: true, Body: func(int) string { return "db:db  RUNNING" }},
	},
	Footer: feedback.EvidenceFooter{Source: "supervisor", Scope: "api", ObservedAt: now.Add(-14 * time.Second), Now: now},
}

return page.View(80, theme.New(theme.Dark))
```

### [Updates, focus & sizing](https://lanej.io/charmed/templates-pattern.html#template%2freport-page-integration)

The parent model supplies observations, coordinates focus and owns application decisions. Forward initialization and update commands rather than executing them during rendering. Resize the existing composition from the terminal allocation so selection remains stable, and apply native customization at the exposed child models after any rebuilding lifecycle step.

### [Run this example](https://lanej.io/charmed/templates-pattern.html#template%2freport-page-run)

```sh
go run ./examples/catalog template/report-page
```

[Example source](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/template_report_page.go) · [Implementation](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/pattern/report.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [layout/header](https://lanej.io/charmed/molecules-layout.html#layout/header) [feedback/evidence-summary](https://lanej.io/charmed/molecules-feedback.html#feedback/evidence-summary) [feedback/evidence-footer](https://lanej.io/charmed/molecules-feedback.html#feedback/evidence-footer) [foundations/layout](https://lanej.io/charmed/foundations-layout.html#foundations/layout) [feedback/evidence](https://lanej.io/charmed/atoms-feedback.html#feedback/evidence)
