[Molecules](https://lanej.io/charmed/molecules.html)Logs

# [Logs](https://lanej.io/charmed/molecules-logs.html#level-molecules-logs)

2 components · [All molecules](https://lanej.io/charmed/molecules.html)

- [Structured log tree](https://lanej.io/charmed/molecules-logs.html#collection%2flog-tree)
- [HTTP log](https://lanej.io/charmed/molecules-logs.html#collection%2fhttp-log)

[Molecules](https://lanej.io/charmed/molecules.html) / Logs

## [Structured log tree](https://lanej.io/charmed/molecules-logs.html#collection%2flog-tree)

Trace and span hierarchy with expandable ordered fields, stable identities and retained follow state.

### [Native preview](https://lanej.io/charmed/molecules-logs.html#collection%2flog-tree-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
   ▾  Logs · 1 retained                                         
 └── ▾  ⋈ Trace trace-c7                                        
    └── ▾  ─ Span carrier                                       
→      └── ▸  ≡ WARN  Carrier response delayed  
```

Demo · truecolor · 120 × 38 terminal cells

```
   ▾  Logs · 4 retained                                                                                                 
 └── ▾  ⋈ Trace trace-c7                                                                                                
    ├── ▾  ─ Span http-list                                                                                             
    │  └── ▾  12:04:00.000 ≡ INFO  Listed parcels                                                                       
    │     ├── ▾  time: 2026-11-08T12:04:00Z                                                                             
    │     └── ▾  count: 0                                                                                               
    └── ▾  ─ Span http-label                                                                                            
       ├── ▾  ─ Span carrier                                                                                            
       │  ├── ▾  12:04:01.200 ≡ WARN  Carrier response delayed                                                          
       │  │  ├── ▾  time: 2026-11-08T12:04:01.2Z                                                                        
       │  │  └── ▸  upstream {2}                                                                                        
       │  └── ▾  12:04:01.842 ≡ ERROR Upstream timeout                                                                  
       │     ├── ▾  time: 2026-11-08T12:04:01.842Z                                                                      
       │     ├── ▾  retryable: true                                                                                     
       │     └── ▸  causes [2]                                                                                          
→      └── ▾  12:04:01.000 ≡ INFO  Create label                                                                         
          ├── ▾  time: 2026-11-08T12:04:01Z                                                                             
          ├── ▸  parcel {2}                                                                                             
          └── ▾  attempt: 2                                                                                             
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Enter expand · g flat/trace · End follow · Ctrl+N append
```

Narrow · truecolor · 40 × 16 terminal cells

```
       │  │  │            2Z           │
       │  │  └── ▸  upstream {2}       │
       │  └── ▾  ≡ ERROR Upstream      │
       │                 timeout       │
       │     ├── ▾  time: 2026-11-     │
       │     │            08T12:04:01. │
       │     │            842Z         │
       │     ├── ▾  retryable: true    ┃
       │     └── ▸  causes [2]         ┃
→      └── ▾  ≡ INFO  Create label     ┃
          ├── ▾  time: 2026-11-        ┃
          │            08T12:04:01Z    ┃
          ├── ▸  parcel {2}            ┃
          └── ▾  attempt: 2            ┃
Enter expand · g flat/trace · End follow
```

256 colors · ansi256 · 120 × 38 terminal cells

```
   ▾  Logs · 4 retained                                                                                                 
 └── ▾  ⋈ Trace trace-c7                                                                                                
    ├── ▾  ─ Span http-list                                                                                             
    │  └── ▾  12:04:00.000 ≡ INFO  Listed parcels                                                                       
    │     ├── ▾  time: 2026-11-08T12:04:00Z                                                                             
    │     └── ▾  count: 0                                                                                               
    └── ▾  ─ Span http-label                                                                                            
       ├── ▾  ─ Span carrier                                                                                            
       │  ├── ▾  12:04:01.200 ≡ WARN  Carrier response delayed                                                          
       │  │  ├── ▾  time: 2026-11-08T12:04:01.2Z                                                                        
       │  │  └── ▸  upstream {2}                                                                                        
       │  └── ▾  12:04:01.842 ≡ ERROR Upstream timeout                                                                  
       │     ├── ▾  time: 2026-11-08T12:04:01.842Z                                                                      
       │     ├── ▾  retryable: true                                                                                     
       │     └── ▸  causes [2]                                                                                          
→      └── ▾  12:04:01.000 ≡ INFO  Create label                                                                         
          ├── ▾  time: 2026-11-08T12:04:01Z                                                                             
          ├── ▸  parcel {2}                                                                                             
          └── ▾  attempt: 2                                                                                             
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Enter expand · g flat/trace · End follow · Ctrl+N append
```

16 colors · ansi16 · 120 × 38 terminal cells

```
   ▾  Logs · 4 retained                                                                                                 
 └── ▾  ⋈ Trace trace-c7                                                                                                
    ├── ▾  ─ Span http-list                                                                                             
    │  └── ▾  12:04:00.000 ≡ INFO  Listed parcels                                                                       
    │     ├── ▾  time: 2026-11-08T12:04:00Z                                                                             
    │     └── ▾  count: 0                                                                                               
    └── ▾  ─ Span http-label                                                                                            
       ├── ▾  ─ Span carrier                                                                                            
       │  ├── ▾  12:04:01.200 ≡ WARN  Carrier response delayed                                                          
       │  │  ├── ▾  time: 2026-11-08T12:04:01.2Z                                                                        
       │  │  └── ▸  upstream {2}                                                                                        
       │  └── ▾  12:04:01.842 ≡ ERROR Upstream timeout                                                                  
       │     ├── ▾  time: 2026-11-08T12:04:01.842Z                                                                      
       │     ├── ▾  retryable: true                                                                                     
       │     └── ▸  causes [2]                                                                                          
→      └── ▾  12:04:01.000 ≡ INFO  Create label                                                                         
          ├── ▾  time: 2026-11-08T12:04:01Z                                                                             
          ├── ▸  parcel {2}                                                                                             
          └── ▾  attempt: 2                                                                                             
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Enter expand · g flat/trace · End follow · Ctrl+N append
```

Monochrome · none · 120 × 38 terminal cells

```
   v  Logs * 4 retained                                                                                                 
 +-- v  T Trace trace-c7                                                                                                
    +-- v  S Span http-list                                                                                             
    |  +-- v  12:04:00.000 L INFO  Listed parcels                                                                       
    |     +-- v  time: 2026-11-08T12:04:00Z                                                                             
    |     +-- v  count: 0                                                                                               
    +-- v  S Span http-label                                                                                            
       +-- v  S Span carrier                                                                                            
       |  +-- v  12:04:01.200 L WARN  Carrier response delayed                                                          
       |  |  +-- v  time: 2026-11-08T12:04:01.2Z                                                                        
       |  |  +-- >  upstream {2}                                                                                        
       |  +-- v  12:04:01.842 L ERROR Upstream timeout                                                                  
       |     +-- v  time: 2026-11-08T12:04:01.842Z                                                                      
       |     +-- v  retryable: true                                                                                     
       |     +-- >  causes [2]                                                                                          
>      +-- v  12:04:01.000 L INFO  Create label                                                                         
          +-- v  time: 2026-11-08T12:04:01Z                                                                             
          +-- >  parcel {2}                                                                                             
          +-- v  attempt: 2                                                                                             
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Enter expand * g flat/trace * End follow * Ctrl+N append
```

light · truecolor · 120 × 38 terminal cells

```
   ▾  Logs · 4 retained                                                                                                 
 └── ▾  ⋈ Trace trace-c7                                                                                                
    ├── ▾  ─ Span http-list                                                                                             
    │  └── ▾  12:04:00.000 ≡ INFO  Listed parcels                                                                       
    │     ├── ▾  time: 2026-11-08T12:04:00Z                                                                             
    │     └── ▾  count: 0                                                                                               
    └── ▾  ─ Span http-label                                                                                            
       ├── ▾  ─ Span carrier                                                                                            
       │  ├── ▾  12:04:01.200 ≡ WARN  Carrier response delayed                                                          
       │  │  ├── ▾  time: 2026-11-08T12:04:01.2Z                                                                        
       │  │  └── ▸  upstream {2}                                                                                        
       │  └── ▾  12:04:01.842 ≡ ERROR Upstream timeout                                                                  
       │     ├── ▾  time: 2026-11-08T12:04:01.842Z                                                                      
       │     ├── ▾  retryable: true                                                                                     
       │     └── ▸  causes [2]                                                                                          
→      └── ▾  12:04:01.000 ≡ INFO  Create label                                                                         
          ├── ▾  time: 2026-11-08T12:04:01Z                                                                             
          ├── ▸  parcel {2}                                                                                             
          └── ▾  attempt: 2                                                                                             
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Enter expand · g flat/trace · End follow · Ctrl+N append
```

contrast · truecolor · 120 × 38 terminal cells

```
   ▾  Logs · 4 retained                                                                                                 
 └── ▾  ⋈ Trace trace-c7                                                                                                
    ├── ▾  ─ Span http-list                                                                                             
    │  └── ▾  12:04:00.000 ≡ INFO  Listed parcels                                                                       
    │     ├── ▾  time: 2026-11-08T12:04:00Z                                                                             
    │     └── ▾  count: 0                                                                                               
    └── ▾  ─ Span http-label                                                                                            
       ├── ▾  ─ Span carrier                                                                                            
       │  ├── ▾  12:04:01.200 ≡ WARN  Carrier response delayed                                                          
       │  │  ├── ▾  time: 2026-11-08T12:04:01.2Z                                                                        
       │  │  └── ▸  upstream {2}                                                                                        
       │  └── ▾  12:04:01.842 ≡ ERROR Upstream timeout                                                                  
       │     ├── ▾  time: 2026-11-08T12:04:01.842Z                                                                      
       │     ├── ▾  retryable: true                                                                                     
       │     └── ▸  causes [2]                                                                                          
→      └── ▾  12:04:01.000 ≡ INFO  Create label                                                                         
          ├── ▾  time: 2026-11-08T12:04:01Z                                                                             
          ├── ▸  parcel {2}                                                                                             
          └── ▾  attempt: 2                                                                                             
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Enter expand · g flat/trace · End follow · Ctrl+N append
```

Empty · truecolor · 120 × 38 terminal cells

```
→  ▾  Logs · 0 retained                                                                                                 
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Enter expand · g flat/trace · End follow · Ctrl+N append
```

Error · truecolor · 120 × 38 terminal cells

```
   ▾  Logs · 4 retained                                                                                                 
 └── ▾  ⋈ Trace trace-c7                                                                                                
    ├── ▾  ─ Span http-list                                                                                             
    │  └── ▾  12:04:00.000 ≡ INFO  Listed parcels                                                                       
    │     ├── ▾  time: 2026-11-08T12:04:00Z                                                                             
    │     └── ▾  count: 0                                                                                               
    ├── ▾  ─ Span http-label [unresolved parent]                                                                        
→   │  └── ▾  12:04:01.000 ≡ INFO  Create label                                                                         
    │     ├── ▾  time: 2026-11-08T12:04:01Z                                                                             
    │     ├── ▸  parcel {2}                                                                                             
    │     └── ▾  attempt: 2                                                                                             
    └── ▾  ─ Span carrier [unresolved parent]                                                                           
       ├── ▾  12:04:01.200 ≡ WARN  Carrier response delayed                                                             
       │  ├── ▾  time: 2026-11-08T12:04:01.2Z                                                                           
       │  └── ▸  upstream {2}                                                                                           
       └── ▾  12:04:01.842 ≡ ERROR Upstream timeout                                                                     
          ├── ▾  time: 2026-11-08T12:04:01.842Z                                                                         
          ├── ▾  retryable: true                                                                                        
          └── ▸  causes [2]                                                                                             
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Enter expand · g flat/trace · End follow · Ctrl+N append
```

Loading · truecolor · 120 × 38 terminal cells

```
   ▾  Logs · 4 retained                                                                                                 
 └── ▾  ⋈ Trace trace-c7                                                                                                
    ├── ▾  ─ Span http-list                                                                                             
    │  └── ▾  12:04:00.000 ≡ INFO  Listed parcels                                                                       
    │     ├── ▾  time: 2026-11-08T12:04:00Z                                                                             
    │     └── ▾  count: 0                                                                                               
    └── ▾  ─ Span http-label                                                                                            
       ├── ▾  ─ Span carrier                                                                                            
       │  ├── ▾  12:04:01.200 ≡ WARN  Carrier response delayed                                                          
       │  │  ├── ▾  time: 2026-11-08T12:04:01.2Z                                                                        
       │  │  └── ▸  upstream {2}                                                                                        
       │  └── ▾  12:04:01.842 ≡ ERROR Upstream timeout                                                                  
       │     ├── ▾  time: 2026-11-08T12:04:01.842Z                                                                      
       │     ├── ▾  retryable: true                                                                                     
       │     └── ▸  causes [2]                                                                                          
→      └── ▾  12:04:01.000 ≡ INFO  Create label                                                                         
          ├── ▾  time: 2026-11-08T12:04:01Z                                                                             
          ├── ▸  parcel {2}                                                                                             
          └── ▾  attempt: 2                                                                                             
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Enter expand · g flat/trace · End follow · Ctrl+N append
```

Glyph only · truecolor · 120 × 38 terminal cells

```
   ▾  Logs · 4 retained                                                                                                 
 └── ▾  ⋈ trace-c7                                                                                                      
    ├── ▾  ─ http-list                                                                                                  
    │  └── ▾  12:04:00.000 ≡ INFO  Listed parcels                                                                       
    │     ├── ▾  time: 2026-11-08T12:04:00Z                                                                             
    │     └── ▾  count: 0                                                                                               
    └── ▾  ─ http-label                                                                                                 
       ├── ▾  ─ carrier                                                                                                 
       │  ├── ▾  12:04:01.200 ≡ WARN  Carrier response delayed                                                          
       │  │  ├── ▾  time: 2026-11-08T12:04:01.2Z                                                                        
       │  │  └── ▸  upstream {2}                                                                                        
       │  └── ▾  12:04:01.842 ≡ ERROR Upstream timeout                                                                  
       │     ├── ▾  time: 2026-11-08T12:04:01.842Z                                                                      
       │     ├── ▾  retryable: true                                                                                     
       │     └── ▸  causes [2]                                                                                          
→      └── ▾  12:04:01.000 ≡ INFO  Create label                                                                         
          ├── ▾  time: 2026-11-08T12:04:01Z                                                                             
          ├── ▸  parcel {2}                                                                                             
          └── ▾  attempt: 2                                                                                             
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Enter expand · g flat/trace · End follow · Ctrl+N append
```

Label only · truecolor · 120 × 38 terminal cells

```
   ▾  Logs · 4 retained                                                                                                 
 └── ▾  Trace trace-c7                                                                                                  
    ├── ▾  Span http-list                                                                                               
    │  └── ▾  12:04:00.000  INFO  Listed parcels                                                                        
    │     ├── ▾  time: 2026-11-08T12:04:00Z                                                                             
    │     └── ▾  count: 0                                                                                               
    └── ▾  Span http-label                                                                                              
       ├── ▾  Span carrier                                                                                              
       │  ├── ▾  12:04:01.200  WARN  Carrier response delayed                                                           
       │  │  ├── ▾  time: 2026-11-08T12:04:01.2Z                                                                        
       │  │  └── ▸  upstream {2}                                                                                        
       │  └── ▾  12:04:01.842  ERROR Upstream timeout                                                                   
       │     ├── ▾  time: 2026-11-08T12:04:01.842Z                                                                      
       │     ├── ▾  retryable: true                                                                                     
       │     └── ▸  causes [2]                                                                                          
→      └── ▾  12:04:01.000  INFO  Create label                                                                          
          ├── ▾  time: 2026-11-08T12:04:01Z                                                                             
          ├── ▸  parcel {2}                                                                                             
          └── ▾  attempt: 2                                                                                             
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Enter expand · g flat/trace · End follow · Ctrl+N append
```

Focused leaf · truecolor · 120 × 38 terminal cells

```
  ▾  Logs · 4 retained                                                                                                  
 └── ▾  ⋈ Trace trace-c7                                                                                                
    ├── ▾  ─ Span http-list                                                                                             
    │  └── ▾  12:04:00.000 ≡ INFO  Listed parcels                                                                       
    └── ▾  ─ Span http-label                                                                                            
       ├── ▾  ─ Span carrier                                                                                            
       │  ├── ▾  12:04:01.200 ≡ WARN  Carrier response delayed                                                          
       │  └── ▾  12:04:01.842 ≡ ERROR Upstream timeout                                                                  
┌ Log · log-2 ────────────────────────────────────────────────────────────────────────────────────────────────────────┐ 
│ 12:04:01.000 ≡ INFO  Create label                                                                                   │ 
│ ├──time: 2026-11-08T12:04:01Z                                                                                       │ 
│ ├──parcel {2}                                                                                                       │ 
│ │  ├──id: "p_42"                                                                                                    │ 
│ │  └──service: "express"                                                                                            │ 
│ └──attempt: 2                                                                                                       │ 
└─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘ 
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
D details · g group · End follow
```

No folds · truecolor · 120 × 38 terminal cells

```
    Logs · 4 retained                                                                                                   
 └──  ⋈ Trace trace-c7                                                                                                  
    ├──  ─ Span http-list                                                                                               
    │  └──  12:04:00.000 ≡ INFO  Listed parcels                                                                         
    │     ├──  time: 2026-11-08T12:04:00Z                                                                               
    │     └──  count: 0                                                                                                 
    └──  ─ Span http-label                                                                                              
       ├──  ─ Span carrier                                                                                              
       │  ├──  12:04:01.200 ≡ WARN  Carrier response delayed                                                            
       │  │  ├──  time: 2026-11-08T12:04:01.2Z                                                                          
       │  │  └──  upstream {2}                                                                                          
       │  │     ├──  host: "carrier.example"                                                                            
       │  │     └──  timeout_ms: 800                                                                                    
       │  └──  12:04:01.842 ≡ ERROR Upstream timeout                                                                    
       │     ├──  time: 2026-11-08T12:04:01.842Z                                                                        
       │     ├──  retryable: true                                                                                       
       │     └──  causes [2]                                                                                            
       │        ├──  [0]: "connect timeout"                                                                             
       │        └──  [1]: "retry exhausted"                                                                             
→      └──  12:04:01.000 ≡ INFO  Create label                                                                           
          ├──  time: 2026-11-08T12:04:01Z                                                                               
          ├──  parcel {2}                                                                                               
          │  ├──  id: "p_42"                                                                                            
          │  └──  service: "express"                                                                                    
          └──  attempt: 2                                                                                               
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
g flat/trace · End follow · Ctrl+N append
```

### [Go usage](https://lanej.io/charmed/molecules-logs.html#collection%2flog-tree-code)

Imports: [`time`](https://pkg.go.dev/time "time") · [`log`](https://pkg.go.dev/charm.land/log/v2 "charm.land/log/v2") · [`collection`](https://pkg.go.dev/github.com/lanej/charmed/collection "github.com/lanej/charmed/collection") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
at := time.Date(2026, 11, 8, 12, 4, 0, 0, time.UTC)
records := []collection.LogRecord{{
	ID: "log-1", RequestID: "req-82", TraceID: "trace-c7", SpanID: "carrier",
	LogEntry: collection.LogEntry{
		Time: at, Level: log.WarnLevel,
		Message: "Carrier response delayed", Fields: []any{"timeout_ms", 800},
	},
}}

logs := collection.NewLogTree(64, 12, theme.New(theme.Dark))
if err := logs.SetEntries(records); err != nil {
	return "", err
}

logs.Select("log-1")

return logs.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-logs.html#collection%2flog-tree-integration)

Submit observations on the owning Bubble Tea update loop, using stable IDs to preserve selection. Route keys to the focused model and forward returned commands. Resize the existing model when space changes; retain native styles and key maps, and distinguish missing observations from empty captured values.

### [Run this example](https://lanej.io/charmed/molecules-logs.html#collection%2flog-tree-run)

```sh
go run ./examples/catalog collection/log-tree
```

[Example source](https://github.com/lanej/charmed/blob/e7987b1035f1f0309d8f643ca46001269083830a/docs/snippets/collection_log_tree.go) · [Implementation](https://github.com/lanej/charmed/blob/e7987b1035f1f0309d8f643ca46001269083830a/story/logs.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [collection/tree](https://lanej.io/charmed/atoms-collection.html#collection/tree) [text/observation](https://lanej.io/charmed/atoms-logs.html#text/observation) [text/time](https://lanej.io/charmed/atoms-logs.html#text/time) [text/disclosure](https://lanej.io/charmed/atoms-text.html#text/disclosure) [feedback/scrollbar](https://lanej.io/charmed/atoms-feedback.html#feedback/scrollbar)

[Molecules](https://lanej.io/charmed/molecules.html) / Logs

## [HTTP log](https://lanej.io/charmed/molecules-logs.html#collection%2fhttp-log)

Request summaries and captured headers/bodies with syntax highlighting, pretty/raw views and explicit capture states.

### [Native preview](https://lanej.io/charmed/molecules-logs.html#collection%2fhttp-log-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
→   POST      502       842ms /v1/labels                        
                                                                
                                                                
                                                                
                                                                
                                                                
                                                                
                                                                
POST /v1/labels                                                 
502 Bad Gateway · 842ms                                         
2026-11-08T12:04:00Z                                            
                                                                
Request ID req-82                                               
Trace ID   trace-c7                                             
                                                                
▾ Request                  ▾ Response                           
No captured headers        No captured headers                  
Body not captured          {                                    
                             "error": "upstream_timeout",       
                             "retryable": true                  
                           }
```

Demo · truecolor · 120 × 38 terminal cells

```
    GET       200        12ms /v1/parcels                                                                               
→   POST      502       842ms /v1/labels                                                                                
    PATCH     204         8ms /v1/parcels/p_42                                                                          
    GET        ◷          n/a /health                                                                                   
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
                                                                                                                        
Request ID req-82                                                                                                       
Trace ID   trace-c7                                                                                                     
                                                                                                                        
▾ Parameters                             ▾ Response                                                                     
Query                                    Content-Type: application/problem+json; charset=utf-8                          
expand: "rate"                           Retry-After: 1                                                                 
expand: "tracking"                       {                                                                              
locale: "en_US"                            "error": "upstream_timeout",                                                 
                                           "retryable": true,                                                           
▸ Request                                  "attempt": 2                                                                 
                                         }                                                                              
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Tab list/details · r raw · 1/2/3 sections
```

Narrow · truecolor · 40 × 16 terminal cells

```
    GET       200  /v1/parcels          
→   POST      502  /v1/labels           
    PATCH     204  /v1/parcels/p_42     
    GET        ◷   /health              
POST /v1/labels                        ┃
502 Bad Gateway · 842ms                ┃
2026-11-08T12:04:01Z                   ┃
                                       │
Request ID req-82                      │
Trace ID   trace-c7                    │
                                       │
▾ Parameters                           │
                                        
                                        
Tab list/details · r raw · 1/2/3 section
```

256 colors · ansi256 · 120 × 38 terminal cells

```
    GET       200        12ms /v1/parcels                                                                               
→   POST      502       842ms /v1/labels                                                                                
    PATCH     204         8ms /v1/parcels/p_42                                                                          
    GET        ◷          n/a /health                                                                                   
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
                                                                                                                        
Request ID req-82                                                                                                       
Trace ID   trace-c7                                                                                                     
                                                                                                                        
▾ Parameters                             ▾ Response                                                                     
Query                                    Content-Type: application/problem+json; charset=utf-8                          
expand: "rate"                           Retry-After: 1                                                                 
expand: "tracking"                       {                                                                              
locale: "en_US"                            "error": "upstream_timeout",                                                 
                                           "retryable": true,                                                           
▸ Request                                  "attempt": 2                                                                 
                                         }                                                                              
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Tab list/details · r raw · 1/2/3 sections
```

16 colors · ansi16 · 120 × 38 terminal cells

```
    GET       200        12ms /v1/parcels                                                                               
→   POST      502       842ms /v1/labels                                                                                
    PATCH     204         8ms /v1/parcels/p_42                                                                          
    GET        ◷          n/a /health                                                                                   
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
                                                                                                                        
Request ID req-82                                                                                                       
Trace ID   trace-c7                                                                                                     
                                                                                                                        
▾ Parameters                             ▾ Response                                                                     
Query                                    Content-Type: application/problem+json; charset=utf-8                          
expand: "rate"                           Retry-After: 1                                                                 
expand: "tracking"                       {                                                                              
locale: "en_US"                            "error": "upstream_timeout",                                                 
                                           "retryable": true,                                                           
▸ Request                                  "attempt": 2                                                                 
                                         }                                                                              
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Tab list/details · r raw · 1/2/3 sections
```

Monochrome · none · 120 × 38 terminal cells

```
    GET       200        12ms /v1/parcels                                                                               
>   POST      502       842ms /v1/labels                                                                                
    PATCH     204         8ms /v1/parcels/p_42                                                                          
    GET        p          n/a /health                                                                                   
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
POST /v1/labels                                                                                                         
502 Bad Gateway * 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
                                                                                                                        
Request ID req-82                                                                                                       
Trace ID   trace-c7                                                                                                     
                                                                                                                        
v Parameters                             v Response                                                                     
Query                                    Content-Type: application/problem+json; charset=utf-8                          
expand: "rate"                           Retry-After: 1                                                                 
expand: "tracking"                       {                                                                              
locale: "en_US"                            "error": "upstream_timeout",                                                 
                                           "retryable": true,                                                           
> Request                                  "attempt": 2                                                                 
                                         }                                                                              
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Tab list/details * r raw * 1/2/3 sections
```

light · truecolor · 120 × 38 terminal cells

```
    GET       200        12ms /v1/parcels                                                                               
→   POST      502       842ms /v1/labels                                                                                
    PATCH     204         8ms /v1/parcels/p_42                                                                          
    GET        ◷          n/a /health                                                                                   
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
                                                                                                                        
Request ID req-82                                                                                                       
Trace ID   trace-c7                                                                                                     
                                                                                                                        
▾ Parameters                             ▾ Response                                                                     
Query                                    Content-Type: application/problem+json; charset=utf-8                          
expand: "rate"                           Retry-After: 1                                                                 
expand: "tracking"                       {                                                                              
locale: "en_US"                            "error": "upstream_timeout",                                                 
                                           "retryable": true,                                                           
▸ Request                                  "attempt": 2                                                                 
                                         }                                                                              
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Tab list/details · r raw · 1/2/3 sections
```

contrast · truecolor · 120 × 38 terminal cells

```
    GET       200        12ms /v1/parcels                                                                               
→   POST      502       842ms /v1/labels                                                                                
    PATCH     204         8ms /v1/parcels/p_42                                                                          
    GET        ◷          n/a /health                                                                                   
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
                                                                                                                        
Request ID req-82                                                                                                       
Trace ID   trace-c7                                                                                                     
                                                                                                                        
▾ Parameters                             ▾ Response                                                                     
Query                                    Content-Type: application/problem+json; charset=utf-8                          
expand: "rate"                           Retry-After: 1                                                                 
expand: "tracking"                       {                                                                              
locale: "en_US"                            "error": "upstream_timeout",                                                 
                                           "retryable": true,                                                           
▸ Request                                  "attempt": 2                                                                 
                                         }                                                                              
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Tab list/details · r raw · 1/2/3 sections
```

Empty · truecolor · 120 × 38 terminal cells

```
No HTTP requests                                                                                                        
Select an HTTP request                                                                                                  
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Tab list/details · 1/2/3 sections
```

Error · truecolor · 120 × 38 terminal cells

```
    GET       200        12ms /v1/parcels                                                                               
→   POST      502       842ms /v1/labels                                                                                
    PATCH     204         8ms /v1/parcels/p_42                                                                          
    GET        ◷          n/a /health                                                                                   
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
                                                                                                                        
Request ID req-82                                                                                                       
Trace ID   trace-c7                                                                                                     
                                                                                                                        
▾ Parameters                             ▾ Response                                                                     
Query                                    Content-Type: application/problem+json; charset=utf-8                          
expand: "rate"                           Retry-After: 1                                                                 
expand: "tracking"                       {"error":                                                                      
locale: "en_US"                          [invalid JSON: showing captured source]                                        
                                                                                                                        
▸ Request                                                                                                               
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Tab list/details · r raw · 1/2/3 sections
```

Loading · truecolor · 120 × 38 terminal cells

```
    GET       200        12ms /v1/parcels                                                                               
    POST      502       842ms /v1/labels                                                                                
    PATCH     204         8ms /v1/parcels/p_42                                                                          
→   GET        ◷          n/a /health                                                                                   
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
GET /health                                                                                                             
◷ PENDING  · elapsed unavailable                                                                                        
2026-11-08T12:04:03Z                                                                                                    
                                                                                                                        
                                                                                                                        
▸ Request                                                                                                               
                                                                                                                        
▾ Response                                                                                                              
No captured headers                                                                                                     
Body not captured                                                                                                       
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Tab list/details · 1/2/3 sections
```

Glyph only · truecolor · 120 × 38 terminal cells

```
    GET       200        12ms /v1/parcels                                                                               
→   POST      502       842ms /v1/labels                                                                                
    PATCH     204         8ms /v1/parcels/p_42                                                                          
    GET        ◷          n/a /health                                                                                   
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
                                                                                                                        
Request ID req-82                                                                                                       
Trace ID   trace-c7                                                                                                     
                                                                                                                        
▾ Parameters                             ▾ Response                                                                     
Query                                    Content-Type: application/problem+json; charset=utf-8                          
expand: "rate"                           Retry-After: 1                                                                 
expand: "tracking"                       {                                                                              
locale: "en_US"                            "error": "upstream_timeout",                                                 
                                           "retryable": true,                                                           
▸ Request                                  "attempt": 2                                                                 
                                         }                                                                              
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Tab list/details · r raw · 1/2/3 sections
```

Label only · truecolor · 120 × 38 terminal cells

```
    GET       200        12ms /v1/parcels                                                                               
→   POST      502       842ms /v1/labels                                                                                
    PATCH     204         8ms /v1/parcels/p_42                                                                          
    GET        ◷          n/a /health                                                                                   
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
                                                                                                                        
Request ID req-82                                                                                                       
Trace ID   trace-c7                                                                                                     
                                                                                                                        
 Parameters                               Response                                                                      
Query                                    Content-Type: application/problem+json; charset=utf-8                          
expand: "rate"                           Retry-After: 1                                                                 
expand: "tracking"                       {                                                                              
locale: "en_US"                            "error": "upstream_timeout",                                                 
                                           "retryable": true,                                                           
 Request                                   "attempt": 2                                                                 
                                         }                                                                              
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Tab list/details · r raw · 1/2/3 sections
```

Timeline · truecolor · 120 × 38 terminal cells

```
12:04:00.000Z                           0               +1s             +2s             +3s           +4s               
  GET /v1/parcels                       ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤            ┊               ┊   200     2.2s
> POST /v1/labels                       ┊    ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤             ┊   502     2.8s
  PATCH /v1/shipments/p_42              ┊             ├━━━━━━━━━┤       ┊               ┊               ┊   204    600ms
  GET /health                           ┊               ┊        ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━▷   ◷       2.4s
  QUERY /v1/rates                       ┊               ┊               ┊       ├━━━━━━━━━━━━━┤         ┊   200    900ms
  POST /events                          ┊               ┊           ├━━━━━━━━━┤         ┊               ┊   202    600ms
├ start  ┤ stop  ▷ pending  ? unknown                                                                                   
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
POST /v1/labels                                                                                                         
502 Bad Gateway · 2.8s                                                                                                  
2026-11-08T12:04:00.35Z                                                                                                 
                                                                                                                        
Request ID req-82                                                                                                       
Trace ID   trace-c7                                                                                                     
                                                                                                                        
▾ Parameters                             ▾ Response                                                                     
Query                                    Content-Type: application/problem+json; charset=utf-8                          
expand: "rate"                           Retry-After: 1                                                                 
expand: "tracking"                       {                                                                              
locale: "en_US"                            "error": "upstream_timeout",                                                 
                                           "retryable": true,                                                           
▸ Request                                  "attempt": 2                                                                 
                                         }                                                                              
                                                                                                                        
                                                                                                                        
                                                                                                                        
Tab list/details · r raw · 1/2/3 sections
```

Timeline grouped · truecolor · 120 × 38 terminal cells

```
12:04:00.000Z                           0               +1s             +2s             +3s           +4s               
  api · api.example.test:443 ───────────────────────────────────────────────────────────────────────────────────────────
  GET /v1/parcels                       ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤            ┊               ┊   200     2.2s
> POST /v1/labels                       ┊    ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤             ┊   502     2.8s
  GET /health                           ┊               ┊        ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━▷   ◷       2.4s
  carrier · carrier.example:8443 ───────────────────────────────────────────────────────────────────────────────────────
  PATCH /v1/shipments/p_42              ┊             ├━━━━━━━━━┤       ┊               ┊               ┊   204    600ms
  QUERY /v1/rates                       ┊               ┊               ┊       ├━━━━━━━━━━━━━┤         ┊   200    900ms
  events · events.internal:9090 ────────────────────────────────────────────────────────────────────────────────────────
  POST /events                          ┊               ┊           ├━━━━━━━━━┤         ┊               ┊   202    600ms
├ start  ┤ stop  ▷ pending  ? unknown                                                                                   
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
POST /v1/labels                                                                                                         
502 Bad Gateway · 2.8s                                                                                                  
2026-11-08T12:04:00.35Z                                                                                                 
                                                                                                                        
Request ID req-82                                                                                                       
Trace ID   trace-c7                                                                                                     
                                                                                                                        
▾ Parameters                             ▾ Response                                                                     
Query                                    Content-Type: application/problem+json; charset=utf-8                          
expand: "rate"                           Retry-After: 1                                                                 
expand: "tracking"                       {                                                                              
locale: "en_US"                            "error": "upstream_timeout",                                                 
                                           "retryable": true,                                                           
▸ Request                                  "attempt": 2                                                                 
                                         }                                                                              
                                                                                                                        
                                                                                                                        
                                                                                                                        
Tab list/details · r raw · 1/2/3 sections
```

URL full · truecolor · 120 × 38 terminal cells

```
    GET       200        12ms https://api.example.test/v1/parcels                                                       
→   POST      502       842ms https://api.example.test/v1/labels?expand=rate&expand=tracking&locale=en_US               
    PATCH     204         8ms https://api.example.test/v1/parcels/p_42                                                  
    GET        ◷          n/a https://api.example.test/health                                                           
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
POST https://api.example.test/v1/labels?expand=rate&expand=tracking&locale=en_US                                        
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
                                                                                                                        
Request ID req-82                                                                                                       
Trace ID   trace-c7                                                                                                     
                                                                                                                        
▾ Parameters                             ▾ Response                                                                     
Query                                    Content-Type: application/problem+json; charset=utf-8                          
expand: "rate"                           Retry-After: 1                                                                 
expand: "tracking"                       {                                                                              
locale: "en_US"                            "error": "upstream_timeout",                                                 
                                           "retryable": true,                                                           
▸ Request                                  "attempt": 2                                                                 
                                         }                                                                              
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Tab list/details · r raw · 1/2/3 sections
```

URL host-path · truecolor · 120 × 38 terminal cells

```
    GET       200        12ms api.example.test/v1/parcels                                                               
→   POST      502       842ms api.example.test/v1/labels?expand=rate&expand=tracking&locale=en_US                       
    PATCH     204         8ms api.example.test/v1/parcels/p_42                                                          
    GET        ◷          n/a api.example.test/health                                                                   
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
POST api.example.test/v1/labels?expand=rate&expand=tracking&locale=en_US                                                
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
                                                                                                                        
Request ID req-82                                                                                                       
Trace ID   trace-c7                                                                                                     
                                                                                                                        
▾ Parameters                             ▾ Response                                                                     
Query                                    Content-Type: application/problem+json; charset=utf-8                          
expand: "rate"                           Retry-After: 1                                                                 
expand: "tracking"                       {                                                                              
locale: "en_US"                            "error": "upstream_timeout",                                                 
                                           "retryable": true,                                                           
▸ Request                                  "attempt": 2                                                                 
                                         }                                                                              
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Tab list/details · r raw · 1/2/3 sections
```

URL path · truecolor · 120 × 38 terminal cells

```
    GET       200        12ms /v1/parcels                                                                               
→   POST      502       842ms /v1/labels                                                                                
    PATCH     204         8ms /v1/parcels/p_42                                                                          
    GET        ◷          n/a /health                                                                                   
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
                                                                                                                        
Request ID req-82                                                                                                       
Trace ID   trace-c7                                                                                                     
                                                                                                                        
▾ Parameters                             ▾ Response                                                                     
Query                                    Content-Type: application/problem+json; charset=utf-8                          
expand: "rate"                           Retry-After: 1                                                                 
expand: "tracking"                       {                                                                              
locale: "en_US"                            "error": "upstream_timeout",                                                 
                                           "retryable": true,                                                           
▸ Request                                  "attempt": 2                                                                 
                                         }                                                                              
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Tab list/details · r raw · 1/2/3 sections
```

Labeled summary · truecolor · 120 × 38 terminal cells

```
    Method   Code     Latency Target                                                                                    
    GET       200        12ms /v1/parcels                                                                               
→   POST      502       842ms /v1/labels                                                                                
    PATCH     204         8ms /v1/parcels/p_42                                                                          
    GET        ◷          n/a /health                                                                                   
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Method     POST                                                                                                         
Target     /v1/labels                                                                                                   
Time       2026-11-08T12:04:01Z                                                                                         
Latency    842ms                                                                                                        
Status     502 Bad Gateway                                                                                              
                                                                                                                        
Request ID req-82                                                                                                       
Trace ID   trace-c7                                                                                                     
                                                                                                                        
▾ Parameters                             ▾ Response                                                                     
Query                                    Content-Type: application/problem+json; charset=utf-8                          
expand: "rate"                           Retry-After: 1                                                                 
expand: "tracking"                       {                                                                              
locale: "en_US"                            "error": "upstream_timeout",                                                 
                                           "retryable": true,                                                           
▸ Request                                  "attempt": 2                                                                 
                                         }                                                                              
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Tab list/details · r raw · 1/2/3 sections
```

No folds · truecolor · 120 × 38 terminal cells

```
    GET       200        12ms /v1/parcels                                                                               
→   POST      502       842ms /v1/labels                                                                                
    PATCH     204         8ms /v1/parcels/p_42                                                                          
    GET        ◷          n/a /health                                                                                   
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
                                                                                                                        
Request ID req-82                                                                                                       
Trace ID   trace-c7                                                                                                     
                                                                                                                        
Parameters                               Response                                                                       
Query                                    Content-Type: application/problem+json; charset=utf-8                          
expand: "rate"                           Retry-After: 1                                                                 
expand: "tracking"                       {                                                                              
locale: "en_US"                            "error": "upstream_timeout",                                                 
                                           "retryable": true,                                                           
Request                                    "attempt": 2                                                                 
Content-Type: application/json           }                                                                              
X-Request-Id: req-82                                                                                                    
{                                                                                                                       
  "parcel_id": "p_42",                                                                                                  
  "service": "express"                                                                                                  
}                                                                                                                       
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Tab list/details · r raw
```

Raw body · truecolor · 120 × 38 terminal cells

```
    GET       200        12ms /v1/parcels                                                                               
→   POST      502       842ms /v1/labels                                                                                
    PATCH     204         8ms /v1/parcels/p_42                                                                          
    GET        ◷          n/a /health                                                                                   
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
                                                                                                                        
Request ID req-82                                                                                                       
Trace ID   trace-c7                                                                                                     
                                                                                                                        
▾ Parameters                             ▾ Response                                                                     
Query                                    Content-Type: application/problem+json; charset=utf-8                          
expand: "rate"                           Retry-After: 1                                                                 
expand: "tracking"                       {"error":"upstream_timeout","retryable":true,"attempt":2}                      
locale: "en_US"                                                                                                         
                                                                                                                        
▸ Request                                                                                                               
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
                                                                                                                        
Tab list/details · r raw · 1/2/3 sections
```

### [Go usage](https://lanej.io/charmed/molecules-logs.html#collection%2fhttp-log-code)

Imports: [`time`](https://pkg.go.dev/time "time") · [`collection`](https://pkg.go.dev/github.com/lanej/charmed/collection "github.com/lanej/charmed/collection") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
at := time.Date(2026, 11, 8, 12, 4, 0, 0, time.UTC)

requests := []collection.HTTPExchange{{
	ID: "http-2", RequestID: "req-82", TraceID: "trace-c7", Time: at,
	Method: "POST", URL: "https://api.example.test/v1/labels",
	StatusCode: 502, Completed: true, Duration: 842 * time.Millisecond, DurationKnown: true,
	ResponseBody: collection.HTTPBody{
		State: collection.Captured, ContentType: "application/json",
		Bytes: []byte(`{"error":"upstream_timeout","retryable":true}`),
	},
}}

logView := collection.NewHTTPLog(64, 24, theme.New(theme.Dark))
logView.SetURLFormat(text.PathAndQuery)

if err := logView.SetEntries(requests); err != nil {
	return "", err
}

logView.Select("http-2")

return logView.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-logs.html#collection%2fhttp-log-integration)

Submit observations on the owning Bubble Tea update loop, using stable IDs to preserve selection. Route keys to the focused model and forward returned commands. Resize the existing model when space changes; retain native styles and key maps, and distinguish missing observations from empty captured values.

### [Run this example](https://lanej.io/charmed/molecules-logs.html#collection%2fhttp-log-run)

```sh
go run ./examples/catalog collection/http-log
```

[Example source](https://github.com/lanej/charmed/blob/e7987b1035f1f0309d8f643ca46001269083830a/docs/snippets/collection_http_log.go) · [Implementation](https://github.com/lanej/charmed/blob/e7987b1035f1f0309d8f643ca46001269083830a/story/logs.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [collection/table](https://lanej.io/charmed/atoms-collection.html#collection/table) [text/http-method](https://lanej.io/charmed/atoms-http.html#text/http-method) [text/url](https://lanej.io/charmed/atoms-http.html#text/url) [text/http-status](https://lanej.io/charmed/atoms-http.html#text/http-status) [text/pending](https://lanej.io/charmed/atoms-logs.html#text/pending) [text/time](https://lanej.io/charmed/atoms-logs.html#text/time) [text/disclosure](https://lanej.io/charmed/atoms-text.html#text/disclosure) [feedback/scrollbar](https://lanej.io/charmed/atoms-feedback.html#feedback/scrollbar) [text/syntax](https://lanej.io/charmed/atoms-text.html#text/syntax) [collection/viewport](https://lanej.io/charmed/atoms-collection.html#collection/viewport)
