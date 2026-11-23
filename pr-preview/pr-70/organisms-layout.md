[Organisms](https://lanej.io/charmed/organisms.html)Layout

# [Layout](https://lanej.io/charmed/organisms-layout.html#level-organisms-layout)

1 component · [All organisms](https://lanej.io/charmed/organisms.html)

- [Detail document](https://lanej.io/charmed/organisms-layout.html#layout%2fdetail-document)

[Organisms](https://lanej.io/charmed/organisms.html) / Layout

## [Detail document](https://lanej.io/charmed/organisms-layout.html#layout%2fdetail-document)

Group identity, observations, conditions and event history into a readable inspection document. Related sections share columns when space permits, while long narrative or history sections can occupy the full width.

### [Native preview](https://lanej.io/charmed/organisms-layout.html#layout%2fdetail-document-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Operation route-0042                                                                                

Gateway route update                                                                                
 Outcome unknown                                                                                    
Next: Read the remote route before changing traffic.                                                

Observation  ───────────────────────────────────    Verification  ──────────────────────────────────
Source: Controller journal                          Submission: Recorded                            
Observed: 09:42 UTC                                 Remote route: Not observed                      
Receipt: ∅ missing                                    No independent read follows the request.      

Activity  ──────────────────────────────────────────────────────────────────────────────────────────
◷ Request recorded                                                                                  
  09:41 UTC  Controller                                                                             

◷ Response timed out                                                                                
  09:42 UTC  Controller                                                                             
  A lost response does not establish the remote result.
```

Demo · truecolor · 112 × 34 terminal cells

```
Operation route-2026-11-23-0042                                                                                 
                                                                                                                
Gateway route update                                                                                            
 Outcome unknown                                                                                                
Next: Inspect the saved route request before changing traffic.                                                  
Observed  09:42 UTC                                                                                             
                                                                                                                
Observation  ────────────────────────────────────────    Verification  ──────────────────────────────────────── 
Domain: Network                                          Submission: Recorded                                   
Source: Controller journal                                 The request reached the controller journal.          
Observed: 09:42 UTC                                      Remote route: Not observed                             
Request receipt: ∅ missing                                 No read follows the submitted request.               
                                                         Attribution: Unknown                                   
                                                           The lost response cannot establish the effect.       
                                                                                                                
Observed evidence  ──────────────────────────────────────────────────────────────────────────────────────────── 
The controller recorded a request to switch the gateway route. Its response timed out after submission. The     
remote route has not been observed since the request.                                                           
                                                                                                                
Activity  ───────────────────────────────────────────────────────────────────────────────────────────────────── 
◷ Route request recorded                                                                                        
  09:41 UTC  Controller                                                                                         
  Switch the gateway route to the prepared target.                                                              
                                                                                                                
◷ Response timed out                                                                                            
  09:42 UTC  Controller                                                                                         
  Submission is recorded; the external result is still unknown.                                                 
                                                                                                                
Interpretation limits  ──────────────────────────────    Inspection notes  ──────────────────────────────────── 
A recorded request establishes intent. A caller's        Compare the saved identity with the observer's         
report establishes what the caller said. An              evidence. Keep the last observation when a refresh     
independent observation may confirm the external         fails. Check the observation time before taking        
state; it does not automatically prove request           another action. Record any remaining attribution gap.  
attribution, service health or incident resolution.
```

Narrow · truecolor · 40 × 16 terminal cells

```
Operation route-2026-11-23-0042        ┃
                                       ┃
Gateway route update                   ┃
 Outcome unknown                       ┃
Next: Inspect the saved route request  │
before changing traffic.               │
Observed: 09:42 UTC                    │
                                       │
Observation  ──────────────────────────│
Domain: Network                        │
Source: Controller journal             │
Observed: 09:42 UTC                    │
Request receipt: ∅ missing             │
                                       │
Verification  ─────────────────────────│
Submission: Recorded                   │
```

256 colors · ansi256 · 112 × 34 terminal cells

```
Operation route-2026-11-23-0042                                                                                 
                                                                                                                
Gateway route update                                                                                            
 Outcome unknown                                                                                                
Next: Inspect the saved route request before changing traffic.                                                  
Observed  09:42 UTC                                                                                             
                                                                                                                
Observation  ────────────────────────────────────────    Verification  ──────────────────────────────────────── 
Domain: Network                                          Submission: Recorded                                   
Source: Controller journal                                 The request reached the controller journal.          
Observed: 09:42 UTC                                      Remote route: Not observed                             
Request receipt: ∅ missing                                 No read follows the submitted request.               
                                                         Attribution: Unknown                                   
                                                           The lost response cannot establish the effect.       
                                                                                                                
Observed evidence  ──────────────────────────────────────────────────────────────────────────────────────────── 
The controller recorded a request to switch the gateway route. Its response timed out after submission. The     
remote route has not been observed since the request.                                                           
                                                                                                                
Activity  ───────────────────────────────────────────────────────────────────────────────────────────────────── 
◷ Route request recorded                                                                                        
  09:41 UTC  Controller                                                                                         
  Switch the gateway route to the prepared target.                                                              
                                                                                                                
◷ Response timed out                                                                                            
  09:42 UTC  Controller                                                                                         
  Submission is recorded; the external result is still unknown.                                                 
                                                                                                                
Interpretation limits  ──────────────────────────────    Inspection notes  ──────────────────────────────────── 
A recorded request establishes intent. A caller's        Compare the saved identity with the observer's         
report establishes what the caller said. An              evidence. Keep the last observation when a refresh     
independent observation may confirm the external         fails. Check the observation time before taking        
state; it does not automatically prove request           another action. Record any remaining attribution gap.  
attribution, service health or incident resolution.
```

16 colors · ansi16 · 112 × 34 terminal cells

```
Operation route-2026-11-23-0042                                                                                 
                                                                                                                
Gateway route update                                                                                            
 Outcome unknown                                                                                                
Next: Inspect the saved route request before changing traffic.                                                  
Observed  09:42 UTC                                                                                             
                                                                                                                
Observation  ────────────────────────────────────────    Verification  ──────────────────────────────────────── 
Domain: Network                                          Submission: Recorded                                   
Source: Controller journal                                 The request reached the controller journal.          
Observed: 09:42 UTC                                      Remote route: Not observed                             
Request receipt: ∅ missing                                 No read follows the submitted request.               
                                                         Attribution: Unknown                                   
                                                           The lost response cannot establish the effect.       
                                                                                                                
Observed evidence  ──────────────────────────────────────────────────────────────────────────────────────────── 
The controller recorded a request to switch the gateway route. Its response timed out after submission. The     
remote route has not been observed since the request.                                                           
                                                                                                                
Activity  ───────────────────────────────────────────────────────────────────────────────────────────────────── 
◷ Route request recorded                                                                                        
  09:41 UTC  Controller                                                                                         
  Switch the gateway route to the prepared target.                                                              
                                                                                                                
◷ Response timed out                                                                                            
  09:42 UTC  Controller                                                                                         
  Submission is recorded; the external result is still unknown.                                                 
                                                                                                                
Interpretation limits  ──────────────────────────────    Inspection notes  ──────────────────────────────────── 
A recorded request establishes intent. A caller's        Compare the saved identity with the observer's         
report establishes what the caller said. An              evidence. Keep the last observation when a refresh     
independent observation may confirm the external         fails. Check the observation time before taking        
state; it does not automatically prove request           another action. Record any remaining attribution gap.  
attribution, service health or incident resolution.
```

Monochrome · none · 112 × 34 terminal cells

```
Operation route-2026-11-23-0042                                                                                 
                                                                                                                
Gateway route update                                                                                            
[Outcome unknown]                                                                                               
Next: Inspect the saved route request before changing traffic.                                                  
Observed  09:42 UTC                                                                                             
                                                                                                                
Observation  ----------------------------------------    Verification  ---------------------------------------- 
Domain: Network                                          Submission: Recorded                                   
Source: Controller journal                                 The request reached the controller journal.          
Observed: 09:42 UTC                                      Remote route: Not observed                             
Request receipt: missing                                   No read follows the submitted request.               
                                                         Attribution: Unknown                                   
                                                           The lost response cannot establish the effect.       
                                                                                                                
Observed evidence  -------------------------------------------------------------------------------------------- 
The controller recorded a request to switch the gateway route. Its response timed out after submission. The     
remote route has not been observed since the request.                                                           
                                                                                                                
Activity  ----------------------------------------------------------------------------------------------------- 
t Route request recorded                                                                                        
  09:41 UTC  Controller                                                                                         
  Switch the gateway route to the prepared target.                                                              
                                                                                                                
t Response timed out                                                                                            
  09:42 UTC  Controller                                                                                         
  Submission is recorded; the external result is still unknown.                                                 
                                                                                                                
Interpretation limits  ------------------------------    Inspection notes  ------------------------------------ 
A recorded request establishes intent. A caller's        Compare the saved identity with the observer's         
report establishes what the caller said. An              evidence. Keep the last observation when a refresh     
independent observation may confirm the external         fails. Check the observation time before taking        
state; it does not automatically prove request           another action. Record any remaining attribution gap.  
attribution, service health or incident resolution.
```

ASCII · truecolor · 112 × 34 terminal cells

```
Operation route-2026-11-23-0042                                                                                 
                                                                                                                
Gateway route update                                                                                            
[Outcome unknown]                                                                                               
Next: Inspect the saved route request before changing traffic.                                                  
Observed  09:42 UTC                                                                                             
                                                                                                                
Observation  ----------------------------------------    Verification  ---------------------------------------- 
Domain: Network                                          Submission: Recorded                                   
Source: Controller journal                                 The request reached the controller journal.          
Observed: 09:42 UTC                                      Remote route: Not observed                             
Request receipt: missing                                   No read follows the submitted request.               
                                                         Attribution: Unknown                                   
                                                           The lost response cannot establish the effect.       
                                                                                                                
Observed evidence  -------------------------------------------------------------------------------------------- 
The controller recorded a request to switch the gateway route. Its response timed out after submission. The     
remote route has not been observed since the request.                                                           
                                                                                                                
Activity  ----------------------------------------------------------------------------------------------------- 
t Route request recorded                                                                                        
  09:41 UTC  Controller                                                                                         
  Switch the gateway route to the prepared target.                                                              
                                                                                                                
t Response timed out                                                                                            
  09:42 UTC  Controller                                                                                         
  Submission is recorded; the external result is still unknown.                                                 
                                                                                                                
Interpretation limits  ------------------------------    Inspection notes  ------------------------------------ 
A recorded request establishes intent. A caller's        Compare the saved identity with the observer's         
report establishes what the caller said. An              evidence. Keep the last observation when a refresh     
independent observation may confirm the external         fails. Check the observation time before taking        
state; it does not automatically prove request           another action. Record any remaining attribution gap.  
attribution, service health or incident resolution.
```

Unicode monochrome · none · 112 × 34 terminal cells

```
Operation route-2026-11-23-0042                                                                                 
                                                                                                                
Gateway route update                                                                                            
[Outcome unknown]                                                                                               
Next: Inspect the saved route request before changing traffic.                                                  
Observed  09:42 UTC                                                                                             
                                                                                                                
Observation  ----------------------------------------    Verification  ---------------------------------------- 
Domain: Network                                          Submission: Recorded                                   
Source: Controller journal                                 The request reached the controller journal.          
Observed: 09:42 UTC                                      Remote route: Not observed                             
Request receipt: missing                                   No read follows the submitted request.               
                                                         Attribution: Unknown                                   
                                                           The lost response cannot establish the effect.       
                                                                                                                
Observed evidence  -------------------------------------------------------------------------------------------- 
The controller recorded a request to switch the gateway route. Its response timed out after submission. The     
remote route has not been observed since the request.                                                           
                                                                                                                
Activity  ----------------------------------------------------------------------------------------------------- 
t Route request recorded                                                                                        
  09:41 UTC  Controller                                                                                         
  Switch the gateway route to the prepared target.                                                              
                                                                                                                
t Response timed out                                                                                            
  09:42 UTC  Controller                                                                                         
  Submission is recorded; the external result is still unknown.                                                 
                                                                                                                
Interpretation limits  ------------------------------    Inspection notes  ------------------------------------ 
A recorded request establishes intent. A caller's        Compare the saved identity with the observer's         
report establishes what the caller said. An              evidence. Keep the last observation when a refresh     
independent observation may confirm the external         fails. Check the observation time before taking        
state; it does not automatically prove request           another action. Record any remaining attribution gap.  
attribution, service health or incident resolution.
```

### [Go usage](https://lanej.io/charmed/organisms-layout.html#layout%2fdetail-document-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
document := layout.DetailDocument{
	Header: layout.IdentityHeader{Kind: "Operation", ID: "route-0042"},
	Summary: layout.RecordSummary{
		Title: "Gateway route update", Status: text.Pill("Outcome unknown", text.Warning, palette, text.Normal),
		Next: &layout.Pair{Label: "Next", Value: "Read the remote route before changing traffic."},
	},
	Sections: []layout.DetailSection{
		{ID: "observation", Title: "Observation", Fields: []layout.Pair{
			{Label: "Source", Value: "Controller journal"}, {Label: "Observed", Value: "09:42 UTC"}, {Label: "Receipt", State: text.ValueMissing},
		}},
		{ID: "verification", Title: "Verification", Statuses: []layout.StatusItem{
			{Label: "Submission", Value: "Recorded"}, {Label: "Remote route", Value: "Not observed", Detail: "No independent read follows the request."},
		}},
		{ID: "history", Title: "Activity", FullWidth: true, Timeline: []layout.TimelineEvent{
			{When: "09:41 UTC", Actor: "Controller", Title: "Request recorded"},
			{When: "09:42 UTC", Actor: "Controller", Title: "Response timed out", Detail: "A lost response does not establish the remote result."},
		}},
	},
}

return document.View(100, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-layout.html#layout%2fdetail-document-integration)

Supply stable section IDs, grouped Fields, Statuses, Timeline and Paragraphs. Sections pair above MinSectionWidth, defaulting to 40 cells per column; FullWidth breaks pairing. StatusList and Timeline also render independently. Titles, prose and event data are escaped; status values accept trusted styled text. Body appends a pure width-aware custom renderer. View returns the full content; Render also returns section ranges. A native viewport or ListInspector owns scrolling. Re-render after data, theme or width changes. Run go run ./cmd/workbench run layout/detail-document for scrollable inspection.

### [Run this example](https://lanej.io/charmed/organisms-layout.html#layout%2fdetail-document-run)

```sh
go run ./examples/catalog layout/detail-document
```

[Example source](https://github.com/lanej/charmed/blob/49f755da5bd25a044043cf98bc303185c2ebce6d/docs/snippets/layout_detail_document.go) · [Implementation](https://github.com/lanej/charmed/blob/49f755da5bd25a044043cf98bc303185c2ebce6d/layout/detail_document.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [layout/identity-header](https://lanej.io/charmed/molecules-layout.html#layout/identity-header) [layout/record-summary](https://lanej.io/charmed/molecules-layout.html#layout/record-summary) [layout/key-value](https://lanej.io/charmed/molecules-layout.html#layout/key-value)
