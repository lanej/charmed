[Templates](https://lanej.io/charmed/templates.html)Patterns

# [Patterns](https://lanej.io/charmed/templates-pattern.html#level-templates-pattern)

3 components · [All templates](https://lanej.io/charmed/templates.html)

- [Application shell](https://lanej.io/charmed/templates-pattern.html#template%2fapplication)
- [Report page](https://lanej.io/charmed/templates-pattern.html#template%2freport-page)
- [Operational dashboard](https://lanej.io/charmed/templates-pattern.html#template%2foperations)

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

[Example source](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/docs/snippets/template_application.go) · [Implementation](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/story/chrome.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

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

[Example source](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/docs/snippets/template_report_page.go) · [Implementation](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/pattern/report.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [layout/header](https://lanej.io/charmed/molecules-layout.html#layout/header) [feedback/evidence-summary](https://lanej.io/charmed/molecules-feedback.html#feedback/evidence-summary) [feedback/evidence-footer](https://lanej.io/charmed/molecules-feedback.html#feedback/evidence-footer) [foundations/layout](https://lanej.io/charmed/foundations-layout.html#foundations/layout) [feedback/evidence](https://lanej.io/charmed/atoms-feedback.html#feedback/evidence)

[Templates](https://lanej.io/charmed/templates.html) / Patterns

## [Operational dashboard](https://lanej.io/charmed/templates-pattern.html#template%2foperations)

Start an operational dashboard with a summary and primary evidence, then open secondary diagnostics with i. Exact count/capacity meters show remaining quota. The title and subtitle share a header with right-aligned observation status; narrow allocations stack fields without dropping them.

### [Native preview](https://lanej.io/charmed/templates-pattern.html#template%2foperations-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 API operations  Synthetic observations / UTC                                   
                                    Observed 00:16 UTC / quota resets 01:00 UTC 
                                                                                
 REST remaining  ▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌░░░░░░░░  3500 / 5000 requests             
                                                                                
 REST attempts per hour                                                         
 22:00 UTC  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━       160 
 23:00 UTC  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━  173 
 00:00 UTC  0 so far (partial)                                                  
                                                                                
                                                                                
                                                                                
                                                                                
                                                                                
                                                                                
                                                                                
 1-6 of 6                                                                       
 ↓/j down  ↑/k up  i details  ? keys
```

Demo · truecolor · 100 × 26 terminal cells

```
 API operations  Synthetic / 14-15 Nov 2026                       Observed 00:16 UTC / resets 01:00 
                                                                                                    
 REST remaining     ▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌░░░░░░░░  3500 / 5000 requests                              
 GraphQL remaining  ▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌░  4800 / 5000 points                                
                                                                                                    
 REST attempts / UTC hour                                                                           
 00:00  0 so far (partial)                                                                          
 23:00  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━  173 
 22:00  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━         160 
 21:00  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━               147 
 20:00  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━                      134 
 19:00  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━                            121 
                                                                                                    
 Coverage                                                                                           
 Complete hourly counts share a scale. Current hour is partial.                                     
 History begins at coordinator start; earlier hours are unavailable.                                
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
 1-14 of 14                                                                                         
 ↓/j down  ↑/k up  i details  ? keys
```

Narrow · truecolor · 40 × 16 terminal cells

```
 API operations                         
 Synthetic / 14-15 Nov 2026             
      Observed 00:16 UTC / resets 01:00 
                                        
 REST remaining  3500 / 5000 requests   
 ▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌░░░░░░░░           
 GraphQL remaining  4800 / 5000 points  
 ▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌░           
                                        
 REST attempts / UTC hour               
 00:00  0 so far (partial)              
 23:00  ━━━━━━━━━━━━━━━━━━━━━━━━━━  173 
 22:00  ━━━━━━━━━━━━━━━━━━━━━━━━    160 
 21:00  ━━━━━━━━━━━━━━━━━━━━━━      147 
 1-10 of 18                             
 ↓/j down  ↑/k up  i details  ? keys
```

256 colors · ansi256 · 100 × 26 terminal cells

```
 API operations  Synthetic / 14-15 Nov 2026                       Observed 00:16 UTC / resets 01:00 
                                                                                                    
 REST remaining     ▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌░░░░░░░░  3500 / 5000 requests                              
 GraphQL remaining  ▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌░  4800 / 5000 points                                
                                                                                                    
 REST attempts / UTC hour                                                                           
 00:00  0 so far (partial)                                                                          
 23:00  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━  173 
 22:00  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━         160 
 21:00  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━               147 
 20:00  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━                      134 
 19:00  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━                            121 
                                                                                                    
 Coverage                                                                                           
 Complete hourly counts share a scale. Current hour is partial.                                     
 History begins at coordinator start; earlier hours are unavailable.                                
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
 1-14 of 14                                                                                         
 ↓/j down  ↑/k up  i details  ? keys
```

16 colors · ansi16 · 100 × 26 terminal cells

```
 API operations  Synthetic / 14-15 Nov 2026                       Observed 00:16 UTC / resets 01:00 
                                                                                                    
 REST remaining     ▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌░░░░░░░░  3500 / 5000 requests                              
 GraphQL remaining  ▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌░  4800 / 5000 points                                
                                                                                                    
 REST attempts / UTC hour                                                                           
 00:00  0 so far (partial)                                                                          
 23:00  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━  173 
 22:00  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━         160 
 21:00  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━               147 
 20:00  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━                      134 
 19:00  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━                            121 
                                                                                                    
 Coverage                                                                                           
 Complete hourly counts share a scale. Current hour is partial.                                     
 History begins at coordinator start; earlier hours are unavailable.                                
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
 1-14 of 14                                                                                         
 ↓/j down  ↑/k up  i details  ? keys
```

Monochrome · none · 100 × 26 terminal cells

```
 API operations  Synthetic / 14-15 Nov 2026                       Observed 00:16 UTC / resets 01:00 
                                                                                                    
 REST remaining     ####################--------  3500 / 5000 requests                              
 GraphQL remaining  ###########################-  4800 / 5000 points                                
                                                                                                    
 REST attempts / UTC hour                                                                           
 00:00  0 so far (partial)                                                                          
 23:00  ######################################################################################  173 
 22:00  ###############################################################################         160 
 21:00  #########################################################################               147 
 20:00  ##################################################################                      134 
 19:00  ############################################################                            121 
                                                                                                    
 Coverage                                                                                           
 Complete hourly counts share a scale. Current hour is partial.                                     
 History begins at coordinator start; earlier hours are unavailable.                                
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
 1-14 of 14                                                                                         
 v/j down  ^/k up  i details  ? keys
```

Inspection · truecolor · 100 × 26 terminal cells

```
 API operations / Details  Synthetic / 14-15 Nov 2026             Observed 00:16 UTC / resets 01:00 
                                                                                                    
 Scheduler                                                                                          
 Cadence   1s tick / 2s minimum request gap                                                         
 Queue     0 waiting / idle                                                                         
 Dispatch  Awaiting work                                                                            
                                                                                                    
 Accounting                                                                                         
 HTTP attempts include transport failures; Git transport is excluded.                               
 Quotas cover the shared credential, including other tools.                                         
 History resets on coordinator restart; no historical backfill.                                     
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
 1-9 of 9                                                                                           
 ↓/j down  ↑/k up  esc overview  ? keys
```

Empty · truecolor · 100 × 26 terminal cells

```
 API operations  Synthetic / 14-15 Nov 2026                                No observations received 
                                                                                                    
 REST remaining  ∅ missing                                                                          
 GraphQL remaining  ∅ missing                                                                       
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
 1-2 of 2                                                                                           
 ↓/j down  ↑/k up  ? keys
```

Error · truecolor · 100 × 26 terminal cells

```
 API operations  Synthetic / 14-15 Nov 2026                                                         
                                           Refresh failed / showing last observation from 00:16 UTC 
                                                                                                    
 REST remaining     ▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌░░░░░░░░  3500 / 5000 requests                              
 GraphQL remaining  ▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌▌░  4800 / 5000 points                                
                                                                                                    
 REST attempts / UTC hour                                                                           
 00:00  0 so far (partial)                                                                          
 23:00  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━  173 
 22:00  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━         160 
 21:00  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━               147 
 20:00  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━                      134 
 19:00  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━                            121 
                                                                                                    
 Coverage                                                                                           
 Complete hourly counts share a scale. Current hour is partial.                                     
 History begins at coordinator start; earlier hours are unavailable.                                
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
 1-14 of 14                                                                                         
 ↓/j down  ↑/k up  i details  ? keys
```

Loading · truecolor · 100 × 26 terminal cells

```
 API operations  Synthetic / 14-15 Nov 2026                           Waiting for first observation 
                                                                                                    
 REST remaining  ∅ missing                                                                          
 GraphQL remaining  ∅ missing                                                                       
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
 1-2 of 2                                                                                           
 ↓/j down  ↑/k up  ? keys
```

### [Go usage](https://lanej.io/charmed/templates-pattern.html#template%2foperations-code)

Imports: [`chart`](https://pkg.go.dev/github.com/lanej/charmed/chart "github.com/lanej/charmed/chart") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`pattern`](https://pkg.go.dev/github.com/lanej/charmed/pattern "github.com/lanej/charmed/pattern") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
bars := chart.BucketBars{Buckets: []chart.CountBucket{
	{Label: "22:00 UTC", Count: 160, State: chart.BucketComplete},
	{Label: "23:00 UTC", Count: 173, State: chart.BucketComplete},
	{Label: "00:00 UTC", Count: 0, State: chart.BucketPartial},
}}
snapshot := pattern.DashboardSnapshot{
	Title: "API operations", Context: "Synthetic observations / UTC",
	Status: "Observed 00:16 UTC / quota resets 01:00 UTC",
	Meters: []pattern.DashboardMeter{{Label: "REST remaining", Value: 3500, Total: 5000, Unit: "requests"}},
	Sections: []pattern.DashboardSection{{Title: "REST attempts per hour", Body: func(width int) string {
		return bars.View(width, palette)
	}}},
	Details: []pattern.DashboardSection{{Title: "Scheduler", Body: func(int) string {
		return "0 queued / idle / minimum request gap 2s"
	}}},
}
dashboard := pattern.NewOperationalDashboard(snapshot, layout.Size{Width: 80, Height: 18}, palette)

return dashboard.View()
```

### [Updates, focus & sizing](https://lanej.io/charmed/templates-pattern.html#template%2foperations-integration)

Forward Update only while focused and return its command. Resize and SetSnapshot retain both native viewports, bindings and styles; offsets clamp to available content. Meters preserve exact counts and presence states beside native proportional bars; zero capacity and invalid ratios have no bar. SetSnapshot replaces render callbacks on the application loop. Overview and inspection scroll independently; i opens details, Escape returns and ? expands help. The parent owns refresh, quit and observations.

### [Run this example](https://lanej.io/charmed/templates-pattern.html#template%2foperations-run)

```sh
go run ./examples/catalog template/operations
```

[Example source](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/docs/snippets/template_operations.go) · [Implementation](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/pattern/operational.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [layout/text-row](https://lanej.io/charmed/molecules-layout.html#layout/text-row) [layout/key-value](https://lanej.io/charmed/molecules-layout.html#layout/key-value) [collection/viewport](https://lanej.io/charmed/atoms-collection.html#collection/viewport) [text/key-hint](https://lanej.io/charmed/atoms-text.html#text/key-hint)
