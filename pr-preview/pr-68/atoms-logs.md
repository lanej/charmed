[Atoms](https://lanej.io/charmed/atoms.html)Logs

# [Logs](https://lanej.io/charmed/atoms-logs.html#level-atoms-logs)

3 components · [All atoms](https://lanej.io/charmed/atoms.html)

- [Pending](https://lanej.io/charmed/atoms-logs.html#text%2fpending)
- [Observation](https://lanej.io/charmed/atoms-logs.html#text%2fobservation)
- [Time](https://lanej.io/charmed/atoms-logs.html#text%2ftime)

[Atoms](https://lanej.io/charmed/atoms.html) / Logs

## [Pending](https://lanej.io/charmed/atoms-logs.html#text%2fpending)

Compact pending glyph and fixed-width caller-driven ellipsis frames.

### [Native preview](https://lanej.io/charmed/atoms-logs.html#text%2fpending-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
  ◷        
 ◷ PENDING 
 ...
```

Demo · truecolor · 100 × 30 terminal cells

```
Pending / stable default                                                                            
                                                                                                    
  ◷   Pending                                                                                       
                                                                                                    
Caller-driven ellipsis / Ctrl+N advances                                                            
                                                                                                    
  .   Pending                                                                                       
                                                                                                    
Three equal-width frames                                                                            
                                                                                                    
  .   |  ..   |  ...
```

Narrow · truecolor · 40 × 16 terminal cells

```
Pending / stable default                
                                        
  ◷   Pending                           
                                        
Caller-driven ellipsis / Ctrl+N advances
                                        
  .   Pending                           
                                        
Three equal-width frames                
                                        
  .   |  ..   |  ...
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Pending / stable default                                                                            
                                                                                                    
  ◷   Pending                                                                                       
                                                                                                    
Caller-driven ellipsis / Ctrl+N advances                                                            
                                                                                                    
  .   Pending                                                                                       
                                                                                                    
Three equal-width frames                                                                            
                                                                                                    
  .   |  ..   |  ...
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Pending / stable default                                                                            
                                                                                                    
  ◷   Pending                                                                                       
                                                                                                    
Caller-driven ellipsis / Ctrl+N advances                                                            
                                                                                                    
  .   Pending                                                                                       
                                                                                                    
Three equal-width frames                                                                            
                                                                                                    
  .   |  ..   |  ...
```

Monochrome · none · 100 × 30 terminal cells

```
Pending / stable default                                                                            
                                                                                                    
  p   Pending                                                                                       
                                                                                                    
Caller-driven ellipsis / Ctrl+N advances                                                            
                                                                                                    
  p   Pending                                                                                       
                                                                                                    
Three equal-width frames                                                                            
                                                                                                    
  p   |   p   |   p
```

light · truecolor · 100 × 30 terminal cells

```
Pending / stable default                                                                            
                                                                                                    
  ◷   Pending                                                                                       
                                                                                                    
Caller-driven ellipsis / Ctrl+N advances                                                            
                                                                                                    
  .   Pending                                                                                       
                                                                                                    
Three equal-width frames                                                                            
                                                                                                    
  .   |  ..   |  ...
```

contrast · truecolor · 100 × 30 terminal cells

```
Pending / stable default                                                                            
                                                                                                    
  ◷   Pending                                                                                       
                                                                                                    
Caller-driven ellipsis / Ctrl+N advances                                                            
                                                                                                    
  .   Pending                                                                                       
                                                                                                    
Three equal-width frames                                                                            
                                                                                                    
  .   |  ..   |  ...
```

### [Go usage](https://lanej.io/charmed/atoms-logs.html#text%2fpending-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return layout.Column(
	text.Pending(palette),
	text.PendingWithPresentation(text.GlyphAndLabel, palette),
	text.Pending(palette, 2), // Optional animation frame supplied by your clock.
)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-logs.html#text%2fpending-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-logs.html#text%2fpending-run)

```sh
go run ./examples/catalog text/pending
```

[Example source](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/docs/snippets/text_pending.go) · [Implementation](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/text/observability.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Logs

## [Observation](https://lanej.io/charmed/atoms-logs.html#text%2fobservation)

Shared glyphs and labels for logs, traces, spans, and service-relative direction.

### [Native preview](https://lanej.io/charmed/atoms-logs.html#text%2fobservation-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
≡ Log entry      
⋈ trace-c7       
─ carrier        
↑ carrier.example
↓ API consumer
```

Demo · truecolor · 100 × 30 terminal cells

```
Observation identities                                                                              
                                                                                                    
≡ Log entry                                                                                         
⋈ Trace trace-c7                                                                                    
─ Span carrier                                                                                      
                                                                                                    
Service-relative direction                                                                          
                                                                                                    
↑ Upstream carrier.example                                                                          
↓ Downstream API consumer                                                                           
                                                                                                    
ASCII fallbacks: L / T / S / ^ / v
```

Narrow · truecolor · 40 × 16 terminal cells

```
Observation identities                  
                                        
≡ Log entry                             
⋈ Trace trace-c7                        
─ Span carrier                          
                                        
Service-relative direction              
                                        
↑ Upstream carrier.example              
↓ Downstream API consumer               
                                        
ASCII fallbacks: L / T / S / ^ / v
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Observation identities                                                                              
                                                                                                    
≡ Log entry                                                                                         
⋈ Trace trace-c7                                                                                    
─ Span carrier                                                                                      
                                                                                                    
Service-relative direction                                                                          
                                                                                                    
↑ Upstream carrier.example                                                                          
↓ Downstream API consumer                                                                           
                                                                                                    
ASCII fallbacks: L / T / S / ^ / v
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Observation identities                                                                              
                                                                                                    
≡ Log entry                                                                                         
⋈ Trace trace-c7                                                                                    
─ Span carrier                                                                                      
                                                                                                    
Service-relative direction                                                                          
                                                                                                    
↑ Upstream carrier.example                                                                          
↓ Downstream API consumer                                                                           
                                                                                                    
ASCII fallbacks: L / T / S / ^ / v
```

Monochrome · none · 100 × 30 terminal cells

```
Observation identities                                                                              
                                                                                                    
L Log entry                                                                                         
T Trace trace-c7                                                                                    
S Span carrier                                                                                      
                                                                                                    
Service-relative direction                                                                          
                                                                                                    
^ Upstream carrier.example                                                                          
v Downstream API consumer                                                                           
                                                                                                    
ASCII fallbacks: L / T / S / ^ / v
```

light · truecolor · 100 × 30 terminal cells

```
Observation identities                                                                              
                                                                                                    
≡ Log entry                                                                                         
⋈ Trace trace-c7                                                                                    
─ Span carrier                                                                                      
                                                                                                    
Service-relative direction                                                                          
                                                                                                    
↑ Upstream carrier.example                                                                          
↓ Downstream API consumer                                                                           
                                                                                                    
ASCII fallbacks: L / T / S / ^ / v
```

contrast · truecolor · 100 × 30 terminal cells

```
Observation identities                                                                              
                                                                                                    
≡ Log entry                                                                                         
⋈ Trace trace-c7                                                                                    
─ Span carrier                                                                                      
                                                                                                    
Service-relative direction                                                                          
                                                                                                    
↑ Upstream carrier.example                                                                          
↓ Downstream API consumer                                                                           
                                                                                                    
ASCII fallbacks: L / T / S / ^ / v
```

### [Go usage](https://lanej.io/charmed/atoms-logs.html#text%2fobservation-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return layout.Column(
	text.Log("Log entry", palette),
	text.Trace("trace-c7", palette),
	text.Span("carrier", palette),
	text.Upstream("carrier.example", palette),
	text.Downstream("API consumer", palette),
)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-logs.html#text%2fobservation-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-logs.html#text%2fobservation-run)

```sh
go run ./examples/catalog text/observation
```

[Example source](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/docs/snippets/text_observation.go) · [Implementation](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/text/observability.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Logs

## [Time](https://lanej.io/charmed/atoms-logs.html#text%2ftime)

Millisecond clock labels and rounded duration abbreviations with exact observations retained by callers.

### [Native preview](https://lanej.io/charmed/atoms-logs.html#text%2ftime-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
12:04:00.123  842ms
```

Demo · truecolor · 100 × 30 terminal cells

```
Abbreviated time / exact source retained                                                            
                                                                                                    
12:04:01.234  from 2026-11-08T12:04:01.234567Z                                                      
--:--:--.---  unavailable                                                                           
                                                                                                    
Latency / compact rounded display                                                                   
                                                                                                    
   420ns  from 420ns                                                                                
   842µs  from 842µs                                                                                
   842ms  from 842ms                                                                                
   1.23s  from 1.23456789s                                                                          
    1m2s  from 1m2s                                                                                 
  -1.23s  from -1.23456789s
```

Narrow · truecolor · 40 × 16 terminal cells

```
Abbreviated time / exact source retained
                                        
12:04:01.234  from 2026-11-08T12:04:01.2
--:--:--.---  unavailable               
                                        
Latency / compact rounded display       
                                        
   420ns  from 420ns                    
   842µs  from 842µs                    
   842ms  from 842ms                    
   1.23s  from 1.23456789s              
    1m2s  from 1m2s                     
  -1.23s  from -1.23456789s
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Abbreviated time / exact source retained                                                            
                                                                                                    
12:04:01.234  from 2026-11-08T12:04:01.234567Z                                                      
--:--:--.---  unavailable                                                                           
                                                                                                    
Latency / compact rounded display                                                                   
                                                                                                    
   420ns  from 420ns                                                                                
   842µs  from 842µs                                                                                
   842ms  from 842ms                                                                                
   1.23s  from 1.23456789s                                                                          
    1m2s  from 1m2s                                                                                 
  -1.23s  from -1.23456789s
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Abbreviated time / exact source retained                                                            
                                                                                                    
12:04:01.234  from 2026-11-08T12:04:01.234567Z                                                      
--:--:--.---  unavailable                                                                           
                                                                                                    
Latency / compact rounded display                                                                   
                                                                                                    
   420ns  from 420ns                                                                                
   842µs  from 842µs                                                                                
   842ms  from 842ms                                                                                
   1.23s  from 1.23456789s                                                                          
    1m2s  from 1m2s                                                                                 
  -1.23s  from -1.23456789s
```

Monochrome · none · 100 × 30 terminal cells

```
Abbreviated time / exact source retained                                                            
                                                                                                    
12:04:01.234  from 2026-11-08T12:04:01.234567Z                                                      
--:--:--.---  unavailable                                                                           
                                                                                                    
Latency / compact rounded display                                                                   
                                                                                                    
   420ns  from 420ns                                                                                
   842?s  from 842?s                                                                                
   842ms  from 842ms                                                                                
   1.23s  from 1.23456789s                                                                          
    1m2s  from 1m2s                                                                                 
  -1.23s  from -1.23456789s
```

light · truecolor · 100 × 30 terminal cells

```
Abbreviated time / exact source retained                                                            
                                                                                                    
12:04:01.234  from 2026-11-08T12:04:01.234567Z                                                      
--:--:--.---  unavailable                                                                           
                                                                                                    
Latency / compact rounded display                                                                   
                                                                                                    
   420ns  from 420ns                                                                                
   842µs  from 842µs                                                                                
   842ms  from 842ms                                                                                
   1.23s  from 1.23456789s                                                                          
    1m2s  from 1m2s                                                                                 
  -1.23s  from -1.23456789s
```

contrast · truecolor · 100 × 30 terminal cells

```
Abbreviated time / exact source retained                                                            
                                                                                                    
12:04:01.234  from 2026-11-08T12:04:01.234567Z                                                      
--:--:--.---  unavailable                                                                           
                                                                                                    
Latency / compact rounded display                                                                   
                                                                                                    
   420ns  from 420ns                                                                                
   842µs  from 842µs                                                                                
   842ms  from 842ms                                                                                
   1.23s  from 1.23456789s                                                                          
    1m2s  from 1m2s                                                                                 
  -1.23s  from -1.23456789s
```

### [Go usage](https://lanej.io/charmed/atoms-logs.html#text%2ftime-code)

Imports: [`time`](https://pkg.go.dev/time "time") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text")

```go
at := time.Date(2026, 11, 8, 12, 4, 0, 123000000, time.UTC)
return text.TimeOfDay(at) + "  " + text.ShortDuration(842*time.Millisecond)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-logs.html#text%2ftime-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-logs.html#text%2ftime-run)

```sh
go run ./examples/catalog text/time
```

[Example source](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/docs/snippets/text_time.go) · [Implementation](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/text/observability.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.
