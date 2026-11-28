[Organisms](https://lanej.io/charmed/organisms.html)Layout

# [Layout](https://lanej.io/charmed/organisms-layout.html#level-organisms-layout)

1 component · [All organisms](https://lanej.io/charmed/organisms.html)

- [Detail document](https://lanej.io/charmed/organisms-layout.html#layout%2fdetail-document)

[Organisms](https://lanej.io/charmed/organisms.html) / Layout

## [Detail document](https://lanej.io/charmed/organisms-layout.html#layout%2fdetail-document)

Present an action with its evidence, grouped facts, independent delivery stages and a chronological event rail. Related sections share columns when space permits, while long narrative or history sections can occupy the full width.

### [Native preview](https://lanej.io/charmed/organisms-layout.html#layout%2fdetail-document-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
DELIVERY INSPECTION · ILLUSTRATIVE                                                                  
PR #42 · Route update   Outcome uncertain                                         Observed 09:42 UTC

┃ NEXT ACTION                                                                                       
┃ Reconcile the saved merge operation                                                               
┃ The request was recorded, but the remote response was unavailable.                                

Identity  ──────────────────────────────────────    Verification  ──────────────────────────────────
Repository  example/service                         Tests   Passed                                  
Revision    a81f39d                                 Review  Complete                                
Owner       Worker 2                                Merge   Unverified                              

Feedback delivery  ─────────────────────────────────────────────────────────────────────────────────
Transport                        → Receipt                         → Disposition                    
Delivered                          Unconfirmed                       Open                           

Activity  ──────────────────────────────────────────────────────────────────────────────────────────
09:41 UTC  ○ Merge requested                                                                        
           │ Request persisted                                                                      
09:42 UTC  ○ Remote response unavailable                                                            
             External outcome unknown
```

Demo · truecolor · 112 × 34 terminal cells

```
DELIVERY INSPECTION · ILLUSTRATIVE                                                                              
PR #42 · Route update   Outcome uncertain                                                    Observed 09:42 UTC 
                                                                                                                
┃ NEXT ACTION                                                                                                   
┃ Reconcile the saved merge operation                                                                           
┃ The request was recorded, but the remote response was unavailable.                                            
                                                                                                                
Identity  ───────────────────────────────────────────    Verification  ──────────────────────────────────────── 
Repository  example/service                              Tests   ✓ Passed                                       
Revision    a81f39d                                      Review  ✓ Complete                                     
Owner       Worker 2                                     Merge   ? Unverified                                   
                                                                                                                
Feedback delivery  ──────────────────────────────────────────────────────────────────────────────────────────── 
Source                     → Transport                  → Receipt                   → Disposition               
✓ Recorded                   ✓ Delivered                  ? Unconfirmed               ◷ Open                    
Review comment               To Worker 2                  No acknowledgement          Awaiting decision         
                                                                                                                
Activity  ───────────────────────────────────────────────────────────────────────────────────────────────────── 
09:41 UTC  ○ Merge requested                                                                                    
           │ Request persisted                                                                                  
09:42 UTC  ○ Remote response unavailable                                                                        
             External outcome unknown
```

Narrow · truecolor · 40 × 16 terminal cells

```
DELIVERY INSPECTION · ILLUSTRATIVE     ┃
PR #42 · Route update                  ┃
 Outcome uncertain   Observed 09:42 UTC┃
                                       ┃
┃ NEXT ACTION                          ┃
┃ Reconcile the saved merge operation  ┃
┃ The request was recorded, but the    ┃
┃ remote response was unavailable.     │
                                       │
Identity  ─────────────────────────────│
Repository  example/service            │
Revision    a81f39d                    │
Owner       Worker 2                   │
                                       │
Verification  ─────────────────────────│
Tests   ✓ Passed                       │
```

256 colors · ansi256 · 112 × 34 terminal cells

```
DELIVERY INSPECTION · ILLUSTRATIVE                                                                              
PR #42 · Route update   Outcome uncertain                                                    Observed 09:42 UTC 
                                                                                                                
┃ NEXT ACTION                                                                                                   
┃ Reconcile the saved merge operation                                                                           
┃ The request was recorded, but the remote response was unavailable.                                            
                                                                                                                
Identity  ───────────────────────────────────────────    Verification  ──────────────────────────────────────── 
Repository  example/service                              Tests   ✓ Passed                                       
Revision    a81f39d                                      Review  ✓ Complete                                     
Owner       Worker 2                                     Merge   ? Unverified                                   
                                                                                                                
Feedback delivery  ──────────────────────────────────────────────────────────────────────────────────────────── 
Source                     → Transport                  → Receipt                   → Disposition               
✓ Recorded                   ✓ Delivered                  ? Unconfirmed               ◷ Open                    
Review comment               To Worker 2                  No acknowledgement          Awaiting decision         
                                                                                                                
Activity  ───────────────────────────────────────────────────────────────────────────────────────────────────── 
09:41 UTC  ○ Merge requested                                                                                    
           │ Request persisted                                                                                  
09:42 UTC  ○ Remote response unavailable                                                                        
             External outcome unknown
```

16 colors · ansi16 · 112 × 34 terminal cells

```
DELIVERY INSPECTION · ILLUSTRATIVE                                                                              
PR #42 · Route update   Outcome uncertain                                                    Observed 09:42 UTC 
                                                                                                                
┃ NEXT ACTION                                                                                                   
┃ Reconcile the saved merge operation                                                                           
┃ The request was recorded, but the remote response was unavailable.                                            
                                                                                                                
Identity  ───────────────────────────────────────────    Verification  ──────────────────────────────────────── 
Repository  example/service                              Tests   ✓ Passed                                       
Revision    a81f39d                                      Review  ✓ Complete                                     
Owner       Worker 2                                     Merge   ? Unverified                                   
                                                                                                                
Feedback delivery  ──────────────────────────────────────────────────────────────────────────────────────────── 
Source                     → Transport                  → Receipt                   → Disposition               
✓ Recorded                   ✓ Delivered                  ? Unconfirmed               ◷ Open                    
Review comment               To Worker 2                  No acknowledgement          Awaiting decision         
                                                                                                                
Activity  ───────────────────────────────────────────────────────────────────────────────────────────────────── 
09:41 UTC  ○ Merge requested                                                                                    
           │ Request persisted                                                                                  
09:42 UTC  ○ Remote response unavailable                                                                        
             External outcome unknown
```

Monochrome · none · 112 × 34 terminal cells

```
DELIVERY INSPECTION * ILLUSTRATIVE                                                                              
PR #42 * Route update  [Outcome uncertain]                                                   Observed 09:42 UTC 
                                                                                                                
# NEXT ACTION                                                                                                   
# Reconcile the saved merge operation                                                                           
# The request was recorded, but the remote response was unavailable.                                            
                                                                                                                
Identity  -------------------------------------------    Verification  ---------------------------------------- 
Repository  example/service                              Tests   + Passed                                       
Revision    a81f39d                                      Review  + Complete                                     
Owner       Worker 2                                     Merge   ? Unverified                                   
                                                                                                                
Feedback delivery  -------------------------------------------------------------------------------------------- 
Source                     > Transport                  > Receipt                   > Disposition               
+ Recorded                   + Delivered                  ? Unconfirmed               p Open                    
Review comment               To Worker 2                  No acknowledgement          Awaiting decision         
                                                                                                                
Activity  ----------------------------------------------------------------------------------------------------- 
09:41 UTC  o Merge requested                                                                                    
           | Request persisted                                                                                  
09:42 UTC  o Remote response unavailable                                                                        
             External outcome unknown
```

ASCII · truecolor · 112 × 34 terminal cells

```
DELIVERY INSPECTION * ILLUSTRATIVE                                                                              
PR #42 * Route update  [Outcome uncertain]                                                   Observed 09:42 UTC 
                                                                                                                
# NEXT ACTION                                                                                                   
# Reconcile the saved merge operation                                                                           
# The request was recorded, but the remote response was unavailable.                                            
                                                                                                                
Identity  -------------------------------------------    Verification  ---------------------------------------- 
Repository  example/service                              Tests   + Passed                                       
Revision    a81f39d                                      Review  + Complete                                     
Owner       Worker 2                                     Merge   ? Unverified                                   
                                                                                                                
Feedback delivery  -------------------------------------------------------------------------------------------- 
Source                     > Transport                  > Receipt                   > Disposition               
+ Recorded                   + Delivered                  ? Unconfirmed               p Open                    
Review comment               To Worker 2                  No acknowledgement          Awaiting decision         
                                                                                                                
Activity  ----------------------------------------------------------------------------------------------------- 
09:41 UTC  o Merge requested                                                                                    
           | Request persisted                                                                                  
09:42 UTC  o Remote response unavailable                                                                        
             External outcome unknown
```

Unicode monochrome · none · 112 × 34 terminal cells

```
DELIVERY INSPECTION * ILLUSTRATIVE                                                                              
PR #42 * Route update  [Outcome uncertain]                                                   Observed 09:42 UTC 
                                                                                                                
# NEXT ACTION                                                                                                   
# Reconcile the saved merge operation                                                                           
# The request was recorded, but the remote response was unavailable.                                            
                                                                                                                
Identity  -------------------------------------------    Verification  ---------------------------------------- 
Repository  example/service                              Tests   + Passed                                       
Revision    a81f39d                                      Review  + Complete                                     
Owner       Worker 2                                     Merge   ? Unverified                                   
                                                                                                                
Feedback delivery  -------------------------------------------------------------------------------------------- 
Source                     > Transport                  > Receipt                   > Disposition               
+ Recorded                   + Delivered                  ? Unconfirmed               p Open                    
Review comment               To Worker 2                  No acknowledgement          Awaiting decision         
                                                                                                                
Activity  ----------------------------------------------------------------------------------------------------- 
09:41 UTC  o Merge requested                                                                                    
           | Request persisted                                                                                  
09:42 UTC  o Remote response unavailable                                                                        
             External outcome unknown
```

### [Go usage](https://lanej.io/charmed/organisms-layout.html#layout%2fdetail-document-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
document := layout.DetailDocument{
	Header: layout.IdentityHeader{ID: "DELIVERY INSPECTION · ILLUSTRATIVE"},
	Summary: layout.RecordSummary{
		Title: "PR #42 · Route update", Status: text.Pill("Outcome uncertain", text.Warning, palette, text.Normal),
		Fields: []layout.Pair{{Label: "Observed", Value: "09:42 UTC"}},
		Next:   &layout.Pair{Label: "Next action", Value: "Reconcile the saved merge operation"},
	},
	ActionDetail: "The request was recorded, but the remote response was unavailable.",
	Sections: []layout.DetailSection{
		{ID: "identity", Title: "Identity", Fields: []layout.Pair{
			{Label: "Repository", Value: "example/service"}, {Label: "Revision", Value: "a81f39d"}, {Label: "Owner", Value: "Worker 2"},
		}},
		{ID: "checks", Title: "Verification", Statuses: []layout.StatusItem{
			{Label: "Tests", Value: "Passed"}, {Label: "Review", Value: "Complete"}, {Label: "Merge", Value: "Unverified"},
		}},
		{ID: "feedback", Title: "Feedback delivery", FullWidth: true, Stages: []layout.StatusItem{
			{Label: "Transport", Value: "Delivered"}, {Label: "Receipt", Value: "Unconfirmed"}, {Label: "Disposition", Value: "Open"},
		}},
		{ID: "activity", Title: "Activity", FullWidth: true, Timeline: []layout.TimelineEvent{
			{When: "09:41 UTC", Title: "Merge requested", Detail: "Request persisted"},
			{When: "09:42 UTC", Title: "Remote response unavailable", Detail: "External outcome unknown"},
		}},
	},
}

return document.View(100, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-layout.html#layout%2fdetail-document-integration)

Set Summary.Next and optional ActionDetail for the action region. Supply stable section IDs, grouped Fields, Statuses, Stages, Timeline and Paragraphs. Stage states are independent. Sections pair above MinSectionWidth, defaulting to 40 cells per column; FullWidth breaks pairing. StatusList and Timeline also render independently. Titles, prose and event data are escaped; status values accept trusted styled text. Body appends a pure width-aware custom renderer. View returns the full content; Render also returns section ranges. A native viewport or ListInspector owns scrolling. Re-render after data, theme or width changes. Run go run ./cmd/workbench run layout/detail-document for scrollable inspection.

### [Run this example](https://lanej.io/charmed/organisms-layout.html#layout%2fdetail-document-run)

```sh
go run ./examples/catalog layout/detail-document
```

[Example source](https://github.com/lanej/charmed/blob/cc0823efe6b1be9d1911e813477912332826207b/docs/snippets/layout_detail_document.go) · [Implementation](https://github.com/lanej/charmed/blob/cc0823efe6b1be9d1911e813477912332826207b/layout/detail_document.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [layout/identity-header](https://lanej.io/charmed/molecules-layout.html#layout/identity-header) [layout/record-summary](https://lanej.io/charmed/molecules-layout.html#layout/record-summary) [layout/key-value](https://lanej.io/charmed/molecules-layout.html#layout/key-value)
