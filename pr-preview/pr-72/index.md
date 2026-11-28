# Terminal interfaces, composed in Go.

Charts, inputs, collections and status atoms. One set of native components, from the smallest indicator to a complete working screen. Shared themes and terminal-cell layout keep data, focus and resizing consistent as you compose larger views from ordinary native Go models.

[Build your first app](https://lanej.io/charmed/getting-started.html) · [Explore the components](https://lanej.io/charmed/atoms.html) · [See complete examples](https://lanej.io/charmed/screens.html)

## One model. Every piece connected.

[Explore the composition ↗](https://lanej.io/charmed/screens-pattern.html#screen/services)

Filter services, follow throughput, inspect a selection. A native Go screen assembled from the same components below.

```
Service workbench  5 services · 0 samples                                                                               
┌──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│ > Filter services…                                                                                                   │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
┌────────────────────────────────────────┐ ┌ API gateway / throughput ─────────────────────────────────────────────────┐
│  API gateway                  ✓ OK     │ │ 1358│                                                                     │
│  Event stream                 ✓ OK     │ │ 1326│                                      ⢀⣀⡀                            │
│  Job workers                  ! WARN   │ │     │                                  ⢀⠤⠒⠉⠁ ⠈⠉⠑⠢⡀                        │
│  Storage                      ✓ OK     │ │ 1261│                                ⡠⠊⠁         ⠈⠢⡀                      │
│  Scheduler                    ✓ OK     │ │     │                              ⢀⠜              ⠈⠢⡀                    │
└────────────────────────────────────────┘ │ 1197│⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣰⣁⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣘⣄⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀ │
┌────────────────────────────────────────┐ │     │⢄                          ⢀⠎                    ⠈⠢⡀                 │
│ API gateway   ✓ OK                     │ │ 1132│⠈⢢                        ⡰⠁                       ⠑⢄                │
│                                        │ │     │  ⠑⡄                    ⢀⠜                           ⠣⡀              │
│ Throughput   1240 req/min              │ │ 1068│   ⠈⠢⡀                 ⡠⠊                             ⠑⢄             │
│ Latency p95  78 ms                     │ │     │     ⠑⡄              ⢀⠔⠁                               ⠈⠢⡀           │
│                                        │ │ 1004│      ⠈⠢⡀          ⢀⠔⠁                                   ⠈⠢⡀         │
│                                        │ │     │        ⠈⠒⠤⣀⡀   ⣀⠤⠒⠁                                       ⠈⠑⠤⣀    ⣀ │
│                                        │ │  939│            ⠈⠉⠉⠉                                               ⠉⠉⠉⠉  │
│                                        │ │     │                                                                     │
│                                        │ │  875└──────────────────────────────────────────────────────────────────── │
│                                        │ │     12:00:00 12:04:06 12:08:12 12:12:18 12:16:24 12:20:30 12:24:37        │
└────────────────────────────────────────┘ └───────────────────────────────────────────────────────────────────────────┘
Tab focus   arrows navigate   Ctrl+N step   Ctrl+C quit
```

```
Service workbench  5 services · 0 sample
┌──────────────────────────────────────┐
│ > Filter services…                   │
└──────────────────────────────────────┘
┌──────────────────────────────────────┐
│  API gateway                ✓ OK     │
│  Event stream               ✓ OK     │
│  Job workers                ! WARN   │
│  Storage                    ✓ OK     │
│  Scheduler                  ✓ OK     │
└──────────────────────────────────────┘
┌ API gateway / throughput ────────────┐
│ 1358│                                │
│     │                ⡠⠔⠢⡀            │
│ 1277│               ⡜   ⠈⢆           │
│     │              ⡰⠁    ⠈⢆          │
│ 1197│⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣰⣁⣀⣀⣀⣀⣀⣀⣈⣆⣀⣀⣀⣀⣀⣀⣀⣀ │
│     │⢆           ⡰⠁        ⠘⡄        │
│ 1116│⠈⡆         ⢠⠃          ⠸⡀       │
│     │ ⠘⡄       ⢀⠇            ⠱⡀      │
│ 1036│  ⠱⡀     ⢀⠎              ⠱⡀     │
│     │   ⠘⢄   ⢠⠊                ⠑⡄    │
│  956│     ⠑⠤⠒⠁                  ⠈⠑⠢⠊ │
│     │                                │
│  875└─────────────────────────────── │
│     12:00:00 12:09:00 12:18:00       │
└──────────────────────────────────────┘
Tab focus   arrows navigate   Ctrl+N ste
```

Saved native output · run locally for keyboard interaction.

## Small pieces. Shared meaning.

[Status atoms ↗](https://lanej.io/charmed/atoms-feedback.html#feedback/status)

```
 ✓ OK  ! WARN  × ERROR  ? UNKNOWN  ⏻ OFF  ◌ INACTIVE  ◷ PENDING  ▶ RUNNING  ‖ PAUSED  ⊘ DISABLED  ● CONNECTED  ○ DISCONNECTED 
```

Semantic color and an explicit label, shared by the service list and inspector.

## [Time series ↗](https://lanej.io/charmed/atoms-chart.html#chart/timeseries)

Fine-grained dot traces over shared time windows.

```
Chart / timeseries                                              
                                                                
105│                                                            
   │⠉⠉⠉⠑⠢⢄                              ⣀⣀⣀⣀⣀⡀     ⢀⣀⡀          
 90│      ⠑⠢⢄⡀                      ⢀⠤⠊⠉     ⠈⠉⠉⠉⠉⠉⠁ ⣈⠭⠕⠢⣄      
   │ ⢀⣀⡠⠤⢄⣀⣀⣀⠬⠭⠒⠒⠛⠛⠭⠭⣒⢄           ⡠⠒⠁            ⢀⠤⠒⠉    ⠈⠫⡢⣀   
 75│⠉⠁                ⠉⠓⢤⣀      ⢀⠔⠁          ⢀⣀⠤⠊⠁         ⠈⠒⢕⡤⡀
   │                     ⠑⠫⡒⠢⠤⠔⠊⠁   ⢀⣀⠤⠤⠒⠒⠉⠉⠉⠁                ⠈⠙
 59│                       ⠈⠑⠤⢄⣀⣀⡠⠔⠊⠁                           
   │                                                            
 44└────────────────────────────────────────────────────────────
   12:00:00 12:04:39 12:09:18 12:13:57 12:18:36 12:23:15        
1 Series 1  2 Series 2                                          
                                                                
Ctrl+N advances fixture data / arrows pan supported plots
```

```
Chart / timeseries                      
                                        
105│                                    
   │⠉⠉⠒⡄                 ⢀⣀⣀⣀   ⢀⡀      
 90│   ⠈⠢⡀             ⢀⠎⠁   ⠉⠉⠉⠁⣈⠭⢢⡀   
   │ ⣀⠤⢄⣀⡨⠕⠒⠛⠭⣒⡄      ⡠⠃       ⡠⠊   ⢻⢄  
 75│⠉          ⠙⣄   ⢀⠔⠁      ⡠⠔⠁     ⠑⢕⡄
   │            ⠈⢗⠢⠤⠊  ⢀⡠⠔⠒⠉⠉          ⠙
 59│              ⠑⠤⣀⡠⠔⠁                
   │                                    
 44└────────────────────────────────────
   12:00:00 12:07:45 12:15:30 12:23:15  
1 Series 1  2 Series 2                  
                                        
Ctrl+N advances fixture data / arrows pa
```

## [Bars ↗](https://lanej.io/charmed/atoms-chart.html#chart/bars)

Compare categories with native terminal geometry.

```
Chart / bars                                                    
                                                                
                          ████████████                          
                          ████████████ ▁▁▁▁▁▁▁▁▁▁▁▁             
             ████████████ ████████████ ████████████             
             ████████████ ████████████ ▄▄▄▄▄▄▄▄▄▄▄▄ ▃▃▃▃▃▃▃▃▃▃▃▃
████████████ ████████████ ▄▄▄▄▄▄▄▄▄▄▄▄ ████████████ ████████████
████████████ ▄▄▄▄▄▄▄▄▄▄▄▄ ████████████ ████████████ ████████████
▄▄▄▄▄▄▄▄▄▄▄▄ ████████████ ████████████ ████████████ ▅▅▅▅▅▅▅▅▅▅▅▅
████████████ ████████████ ████████████ ████████████ ████████████
────────────────────────────────────────────────────────────────
API          Events       Jobs         Storage      Schedule    
                                                                
Ctrl+N advances fixture data / arrows pan supported plots
```

```
Chart / bars                            
                                        
                ███████                 
                ███████ ▁▁▁▁▁▁▁         
        ███████ ███████ ███████         
        ███████ ███████ ▄▄▄▄▄▄▄ ▃▃▃▃▃▃▃ 
███████ ███████ ▄▄▄▄▄▄▄ ███████ ███████ 
███████ ▄▄▄▄▄▄▄ ███████ ███████ ███████ 
▄▄▄▄▄▄▄ ███████ ███████ ███████ ▅▅▅▅▅▅▅ 
███████ ███████ ███████ ███████ ███████ 
────────────────────────────────────────
API     Events  Jobs    Storage Schedul 
                                        
Ctrl+N advances fixture data / arrows pa
```

## [Distributions ↗](https://lanej.io/charmed/molecules-chart.html#chart/distributions)

Cohort means, spread and selected observations.

```
Reference mean 20% s=3pp n=64                                   
··········░░░▒▒▒▒▓▓▓|▓▓▓▒▒▒W░░E·································
Candidate mean 30% s=8pp n=32                                   
····░░░░░░░░░▒▒▒▒▒▒▒▒▒▓▓▓▓▓W▓▓E▓|▓▓▓▓▓▓▓▓▓▒X▒▒▒▒▒▒▒░░░░░░░░░····
2.64%                                                     57.36%
| Reference mean  W warning boundary  E error boundary          
| Candidate mean  X selected observation
```

```
Reference mean 20% s=3pp n=64           
······░░▒▒▓▓|▓▓▓▒W░E····················
Candidate mean 30% s=8pp n=32           
··░░░░░░▒▒▒▒▒▒▓▓▓W▓E|▓▓▓▓▓▒X▒▒▒▒░░░░░░··
2.64%                             57.36%
| Reference mean  W warning boundary    
E error boundary  | Candidate mean      
X selected observation
```

## [Signed values ↗](https://lanej.io/charmed/atoms-text.html#text/signed-number)

Direction, precision and semantic emphasis.

```
Signed number / normal                                          
                                                                
Positive means improvement   +0.18pp                            
Higher server-error rate     +0.18pp                            
Zero                        +0pp                                
Small nonzero               -1e-06pp                            
Unavailable                 n/a
```

```
Signed number / normal                  
                                        
Positive means improvement   +0.18pp    
Higher server-error rate     +0.18pp    
Zero                        +0pp        
Small nonzero               -1e-06pp    
Unavailable                 n/a
```

## Bring the pieces into your terminal.

```sh
go run ./cmd/workbench run screen/services
```

[Browse the atoms ↗](https://lanej.io/charmed/atoms.html) · [Compare a deployment ↗](https://lanej.io/charmed/deployment.html) · [Build with Charmed ↗](https://lanej.io/charmed/development.html)
