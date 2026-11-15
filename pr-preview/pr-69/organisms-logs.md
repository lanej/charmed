[Organisms](https://lanej.io/charmed/organisms.html)Logs

# [Logs](https://lanej.io/charmed/organisms-logs.html#level-organisms-logs)

1 component · [All organisms](https://lanej.io/charmed/organisms.html)

- [Correlated inspector](https://lanej.io/charmed/organisms-logs.html#pattern%2flog-inspector)

[Organisms](https://lanej.io/charmed/organisms.html) / Logs

## [Correlated inspector](https://lanej.io/charmed/organisms-logs.html#pattern%2flog-inspector)

Investigate one HTTP exchange alongside its correlated application logs. Request and response payloads keep separate viewports while the parent coordinates selection, focus and available space.

### [Native preview](https://lanej.io/charmed/organisms-logs.html#pattern%2flog-inspector-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Request req-82                                                                                      
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                     0          +211ms      +421ms     +632ms +842ms                │
│ > POST /v1/labels                 ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤   502    842ms │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v1/labels                                                                                     
502 Bad Gateway · 842ms                                                                             
2026-11-08T12:04:00Z                                                                                
Trace ID   trace-c7                                                                                 
┌ ▾ Request ───────────────────────────────┐ ┌ ▾ Response ─────────────────────────────────────────┐
│ No captured headers                      │ │ No captured headers                                 │
│ Body not captured                        │ │   ▾ {                                               │
│                                          │ │       "error": "upstream_timeout",                  │
│                                          │ │       "retryable": true                             │
│                                          │ │     }                                               │
│                                          │ │                                                     │
│                                          │ │                                                     │
│                                          │ │                                                     │
│                                          │ │                                                     │
│                                          │ │                                                     │
│                                          │ │                                                     │
│                                          │ │                                                     │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ▶ FOLLOW  ─────────────────────────────────────────────────────────────────────┐
│ request req-82 · 1 matching logs                                                                 │
│     └── ▾  ─ Span carrier                                                                      │ │
│ ┌ Log · log-1 ────────────────────────────────────────────────────────────────────────────────┐┃ │
│ │ 12:04:00.000 ≡ WARN  Carrier response delayed                                             ┃ │┃ │
│ └─────────────────────────────────────────────────────────────────────────────────────────────┘┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
```

Demo · truecolor · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   GET /v1/parcels                       ┤              ┊              ┊              ┊              ┊   200     12ms │
│ > POST /v1/labels                       ┊              ├━━━━━━━━━━━┤  ┊              ┊              ┊   502    842ms │
│   PATCH /v1/parcels/p_42                ┊              ┊              ┤              ┊              ┊   204      8ms │
│   GET /health                           ┊              ┊              ┊              ├━━━━━━━━━━━━━━▷   ◷         1s │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
Trace ID   trace-c7                                                                                                     
┌ ▾ Request ───────────────────────────────┐ ┌ ▾ Response ─────────────────────────────────────────────────────────────┐
│ ▾ Parameters                             │ │ Content-Type: application/problem+json; charset=utf-8                   │
│ Query                                    │ │ Retry-After: 1                                                          │
│ expand: "rate"                           │ │   ▾ {                                                                   │
│ expand: "tracking"                       │ │       "error": "upstream_timeout",                                      │
│ locale: "en_US"                          │ │       "retryable": true,                                                │
│                                          │ │       "attempt": 2                                                      │
│ Content-Type: application/json           │ │     }                                                                   │
│ X-Request-Id: req-82                     │ │                                                                         │
│   ▾ {                                    │ │                                                                         │
│       "parcel_id": "p_42",               │ │                                                                         │
│       "service": "express"               │ │                                                                         │
│     }                                    │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                                                             │ │
│ →      └── ▾  12:04:01.000 ≡ INFO  Create label                                                                    │ │
│           ├── ▾  time: 2026-11-08T12:04:01Z                                                                        ┃ │
│           ├── ▸  parcel {2}                                                                                        ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope · z fold pane
```

Narrow · truecolor · 40 × 16 terminal cells

```
Request req-82                          
┌ ▾ HTTP requests ·  ‖ PAUSED  ────────┐
│     GET       200  /v1/parcels       │
│ →   POST      502  /v1/labels        │
│     PATCH     204  /v1/parcels/p_42  │
│     GET        ◷   /health           │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
└──────────────────────────────────────┘
Tab panes · ? help · Ctrl+K commands    
↑/k/↓/j select · F follow · t scope · z 
```

256 colors · ansi256 · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   GET /v1/parcels                       ┤              ┊              ┊              ┊              ┊   200     12ms │
│ > POST /v1/labels                       ┊              ├━━━━━━━━━━━┤  ┊              ┊              ┊   502    842ms │
│   PATCH /v1/parcels/p_42                ┊              ┊              ┤              ┊              ┊   204      8ms │
│   GET /health                           ┊              ┊              ┊              ├━━━━━━━━━━━━━━▷   ◷         1s │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
Trace ID   trace-c7                                                                                                     
┌ ▾ Request ───────────────────────────────┐ ┌ ▾ Response ─────────────────────────────────────────────────────────────┐
│ ▾ Parameters                             │ │ Content-Type: application/problem+json; charset=utf-8                   │
│ Query                                    │ │ Retry-After: 1                                                          │
│ expand: "rate"                           │ │   ▾ {                                                                   │
│ expand: "tracking"                       │ │       "error": "upstream_timeout",                                      │
│ locale: "en_US"                          │ │       "retryable": true,                                                │
│                                          │ │       "attempt": 2                                                      │
│ Content-Type: application/json           │ │     }                                                                   │
│ X-Request-Id: req-82                     │ │                                                                         │
│   ▾ {                                    │ │                                                                         │
│       "parcel_id": "p_42",               │ │                                                                         │
│       "service": "express"               │ │                                                                         │
│     }                                    │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                                                             │ │
│ →      └── ▾  12:04:01.000 ≡ INFO  Create label                                                                    │ │
│           ├── ▾  time: 2026-11-08T12:04:01Z                                                                        ┃ │
│           ├── ▸  parcel {2}                                                                                        ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope · z fold pane
```

16 colors · ansi16 · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   GET /v1/parcels                       ┤              ┊              ┊              ┊              ┊   200     12ms │
│ > POST /v1/labels                       ┊              ├━━━━━━━━━━━┤  ┊              ┊              ┊   502    842ms │
│   PATCH /v1/parcels/p_42                ┊              ┊              ┤              ┊              ┊   204      8ms │
│   GET /health                           ┊              ┊              ┊              ├━━━━━━━━━━━━━━▷   ◷         1s │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
Trace ID   trace-c7                                                                                                     
┌ ▾ Request ───────────────────────────────┐ ┌ ▾ Response ─────────────────────────────────────────────────────────────┐
│ ▾ Parameters                             │ │ Content-Type: application/problem+json; charset=utf-8                   │
│ Query                                    │ │ Retry-After: 1                                                          │
│ expand: "rate"                           │ │   ▾ {                                                                   │
│ expand: "tracking"                       │ │       "error": "upstream_timeout",                                      │
│ locale: "en_US"                          │ │       "retryable": true,                                                │
│                                          │ │       "attempt": 2                                                      │
│ Content-Type: application/json           │ │     }                                                                   │
│ X-Request-Id: req-82                     │ │                                                                         │
│   ▾ {                                    │ │                                                                         │
│       "parcel_id": "p_42",               │ │                                                                         │
│       "service": "express"               │ │                                                                         │
│     }                                    │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                                                             │ │
│ →      └── ▾  12:04:01.000 ≡ INFO  Create label                                                                    │ │
│           ├── ▾  time: 2026-11-08T12:04:01Z                                                                        ┃ │
│           ├── ▸  parcel {2}                                                                                        ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope · z fold pane
```

Monochrome · none · 120 × 38 terminal cells

```
Request req-82                                                                                                          
+ v Request timeline * [= PAUSED] -------------------------------------------------------------------------------------+
| 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                |
|   GET /v1/parcels                       ]              |              |              |              |   200     12ms |
| > POST /v1/labels                       |              [===========]  |              |              |   502    842ms |
|   PATCH /v1/parcels/p_42                |              |              ]              |              |   204      8ms |
|   GET /health                           |              |              |              [==============>   p         1s |
+----------------------------------------------------------------------------------------------------------------------+
POST /v1/labels                                                                                                         
502 Bad Gateway * 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
Trace ID   trace-c7                                                                                                     
+ v Request -------------------------------+ + v Response -------------------------------------------------------------+
| v Parameters                             | | Content-Type: application/problem+json; charset=utf-8                   |
| Query                                    | | Retry-After: 1                                                          |
| expand: "rate"                           | |   v {                                                                   |
| expand: "tracking"                       | |       "error": "upstream_timeout",                                      |
| locale: "en_US"                          | |       "retryable": true,                                                |
|                                          | |       "attempt": 2                                                      |
| Content-Type: application/json           | |     }                                                                   |
| X-Request-Id: req-82                     | |                                                                         |
|   v {                                    | |                                                                         |
|       "parcel_id": "p_42",               | |                                                                         |
|       "service": "express"               | |                                                                         |
|     }                                    | |                                                                         |
|                                          | |                                                                         |
|                                          | |                                                                         |
|                                          | |                                                                         |
+------------------------------------------+ +-------------------------------------------------------------------------+
+ v Related logs * [= PAUSED] -----------------------------------------------------------------------------------------+
| request req-82 * 3 matching logs                                                                                     |
|        |  +-- >  12:04:01.842 L ERROR Upstream timeout                                                             | |
| >      +-- v  12:04:01.000 L INFO  Create label                                                                    | |
|           +-- v  time: 2026-11-08T12:04:01Z                                                                        # |
|           +-- >  parcel {2}                                                                                        # |
+----------------------------------------------------------------------------------------------------------------------+
Tab panes * L layout * Alt+1/2/3/4 views * Ctrl+K commands * ? help                                                     
^/k/v/j select * F follow * t scope * z fold pane
```

Backdrop dim · truecolor · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   api · api.example.test:443 ─────────────────────────────────────────────────────────────────────────────────────── │
│   GET /v1/parcels                       ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤           ┊              ┊   200     2.2s │
│ > POST /v1/labels                       ┊    ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤            ┊   502     2.8s │
│   GET /health                           ┊              ┊        ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━▷   ◷       2.4s │
│   carrier · carrier.example:8443 ─────────────────────────────────────────────────────────────────────────────────── │
│   PATCH /v1/shipments/p_42              ┊            ├━━━━━━━━┤       ┊              ┊              ┊   204    600ms │
│   QUERY /v1/rates                       ┊              ┊              ┊      ├━━━━━━━━━━━━━┤        ┊   200    900ms │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v1/labels                                                                                                         
502 Bad Gateway · 2.8s  ┌ Commands ────────────────────────────────────────────────────────────┐                        
2026-11-08T12:04:00.35Z │ > laytm                                                              │                        
Trace ID   trace-c7     │                                                                      │                        
┌ ▾ Request ────────────│ → Layout: timeline                                           current │───────────────────────┐
│ ▾ Parameters          │   Arrange requests, HTTP details and related logs                    │tf-8                   │
│ Query                 │   Layout: sidebar                                                    │                       │
│ expand: "rate"        │   Arrange requests, HTTP details and related logs                    │                       │
│ expand: "tracking"    │   Layout: columns                                                    │                       │
│ locale: "en_US"       │   Arrange requests, HTTP details and related logs                    │                       │
│                       │   Layout: focused                                                    │                       │
│ Content-Type: applicat│   Arrange requests, HTTP details and related logs                    │                       │
│ X-Request-Id: req-82  │ 1/4 · 4 matches                                                      │                       │
│   ▾ {                 │ ↑/↓ select · Enter run · Esc cancel                                  │                       │
│       "parcel_id": "p_└──────────────────────────────────────────────────────────────────────┘                       │
│       "service": "express"               │ │                                                                         │
│     }                                    │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│ ┌ Log · log-3 ────────────────────────────────────────────────────────────────────────────────────────────────────┐│ │
│ │ 12:04:01.200 ≡ WARN  Carrier response delayed                                                                 ┃ ││ │
│ └─────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘┃ │
│        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                                                             ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope · z fold pane                                                                       
```

Backdrop monochrome · truecolor · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   api · api.example.test:443 ─────────────────────────────────────────────────────────────────────────────────────── │
│   GET /v1/parcels                       ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤           ┊              ┊   200     2.2s │
│ > POST /v1/labels                       ┊    ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤            ┊   502     2.8s │
│   GET /health                           ┊              ┊        ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━▷   ◷       2.4s │
│   carrier · carrier.example:8443 ─────────────────────────────────────────────────────────────────────────────────── │
│   PATCH /v1/shipments/p_42              ┊            ├━━━━━━━━┤       ┊              ┊              ┊   204    600ms │
│   QUERY /v1/rates                       ┊              ┊              ┊      ├━━━━━━━━━━━━━┤        ┊   200    900ms │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v1/labels                                                                                                         
502 Bad Gateway · 2.8s  ┌ Commands ────────────────────────────────────────────────────────────┐                        
2026-11-08T12:04:00.35Z │ > laytm                                                              │                        
Trace ID   trace-c7     │                                                                      │                        
┌ ▾ Request ────────────│ → Layout: timeline                                           current │───────────────────────┐
│ ▾ Parameters          │   Arrange requests, HTTP details and related logs                    │tf-8                   │
│ Query                 │   Layout: sidebar                                                    │                       │
│ expand: "rate"        │   Arrange requests, HTTP details and related logs                    │                       │
│ expand: "tracking"    │   Layout: columns                                                    │                       │
│ locale: "en_US"       │   Arrange requests, HTTP details and related logs                    │                       │
│                       │   Layout: focused                                                    │                       │
│ Content-Type: applicat│   Arrange requests, HTTP details and related logs                    │                       │
│ X-Request-Id: req-82  │ 1/4 · 4 matches                                                      │                       │
│   ▾ {                 │ ↑/↓ select · Enter run · Esc cancel                                  │                       │
│       "parcel_id": "p_└──────────────────────────────────────────────────────────────────────┘                       │
│       "service": "express"               │ │                                                                         │
│     }                                    │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│ ┌ Log · log-3 ────────────────────────────────────────────────────────────────────────────────────────────────────┐│ │
│ │ 12:04:01.200 ≡ WARN  Carrier response delayed                                                                 ┃ ││ │
│ └─────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘┃ │
│        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                                                             ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope · z fold pane                                                                       
```

Backdrop none · truecolor · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   api · api.example.test:443 ─────────────────────────────────────────────────────────────────────────────────────── │
│   GET /v1/parcels                       ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤           ┊              ┊   200     2.2s │
│ > POST /v1/labels                       ┊    ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤            ┊   502     2.8s │
│   GET /health                           ┊              ┊        ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━▷   ◷       2.4s │
│   carrier · carrier.example:8443 ─────────────────────────────────────────────────────────────────────────────────── │
│   PATCH /v1/shipments/p_42              ┊            ├━━━━━━━━┤       ┊              ┊              ┊   204    600ms │
│   QUERY /v1/rates                       ┊              ┊              ┊      ├━━━━━━━━━━━━━┤        ┊   200    900ms │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v1/labels                                                                                                         
502 Bad Gateway · 2.8s  ┌ Commands ────────────────────────────────────────────────────────────┐                        
2026-11-08T12:04:00.35Z │ > laytm                                                              │                        
Trace ID   trace-c7     │                                                                      │                        
┌ ▾ Request ────────────│ → Layout: timeline                                           current │───────────────────────┐
│ ▾ Parameters          │   Arrange requests, HTTP details and related logs                    │tf-8                   │
│ Query                 │   Layout: sidebar                                                    │                       │
│ expand: "rate"        │   Arrange requests, HTTP details and related logs                    │                       │
│ expand: "tracking"    │   Layout: columns                                                    │                       │
│ locale: "en_US"       │   Arrange requests, HTTP details and related logs                    │                       │
│                       │   Layout: focused                                                    │                       │
│ Content-Type: applicat│   Arrange requests, HTTP details and related logs                    │                       │
│ X-Request-Id: req-82  │ 1/4 · 4 matches                                                      │                       │
│   ▾ {                 │ ↑/↓ select · Enter run · Esc cancel                                  │                       │
│       "parcel_id": "p_└──────────────────────────────────────────────────────────────────────┘                       │
│       "service": "express"               │ │                                                                         │
│     }                                    │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│ ┌ Log · log-3 ────────────────────────────────────────────────────────────────────────────────────────────────────┐│ │
│ │ 12:04:01.200 ≡ WARN  Carrier response delayed                                                                 ┃ ││ │
│ └─────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘┃ │
│        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                                                             ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope · z fold pane
```

Vim search · truecolor · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   GET /v1/parcels                       ┤              ┊              ┊              ┊              ┊   200     12ms │
│ > POST /v1/labels                       ┊              ├━━━━━━━━━━━┤  ┊              ┊              ┊   502    842ms │
│   PATCH /v1/parcels/p_42                ┊              ┊              ┤              ┊              ┊   204      8ms │
│   GET /health                           ┊              ┊              ┊              ├━━━━━━━━━━━━━━▷   ◷         1s │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
Trace ID   trace-c7                                                                                                     
┌ ▾ Request ────────────┌ Commands ────────────────────────────────────────────────────────────┐───────────────────────┐
│ ▾ Parameters          │ > laytm                                                              │tf-8                   │
│ Query                 │                                                                      │                       │
│ expand: "rate"        │ → Layout: timeline                                           current │                       │
│ expand: "tracking"    │   Arrange requests, HTTP details and related logs                    │                       │
│ locale: "en_US"       │   Layout: sidebar                                                    │                       │
│                       │   Arrange requests, HTTP details and related logs                    │                       │
│ Content-Type: applicat│   Layout: columns                                                    │                       │
│ X-Request-Id: req-82  │   Arrange requests, HTTP details and related logs                    │                       │
│   ▾ {                 │   Layout: focused                                                    │                       │
│       "parcel_id": "p_│   Arrange requests, HTTP details and related logs                    │                       │
│       "service": "expr│ SEARCH · 1/4 · 4 matches                                             │                       │
│     }                 │ ↑↓ select · Esc normal · Enter run                                   │                       │
│                       └──────────────────────────────────────────────────────────────────────┘                       │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│ ┌ Log · log-3 ────────────────────────────────────────────────────────────────────────────────────────────────────┐│ │
│ │ 12:04:01.200 ≡ WARN  Carrier response delayed                                                                 ┃ ││ │
│ └─────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘┃ │
│        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                                                             ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
k/↑/j/↓ select · F follow · t scope · z fold pane                                                                       
```

Vim normal · truecolor · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   GET /v1/parcels                       ┤              ┊              ┊              ┊              ┊   200     12ms │
│ > POST /v1/labels                       ┊              ├━━━━━━━━━━━┤  ┊              ┊              ┊   502    842ms │
│   PATCH /v1/parcels/p_42                ┊              ┊              ┤              ┊              ┊   204      8ms │
│   GET /health                           ┊              ┊              ┊              ├━━━━━━━━━━━━━━▷   ◷         1s │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
Trace ID   trace-c7                                                                                                     
┌ ▾ Request ────────────┌ Commands ────────────────────────────────────────────────────────────┐───────────────────────┐
│ ▾ Parameters          │ > laytm                                                              │tf-8                   │
│ Query                 │                                                                      │                       │
│ expand: "rate"        │ → Layout: timeline                                           current │                       │
│ expand: "tracking"    │   Arrange requests, HTTP details and related logs                    │                       │
│ locale: "en_US"       │   Layout: sidebar                                                    │                       │
│                       │   Arrange requests, HTTP details and related logs                    │                       │
│ Content-Type: applicat│   Layout: columns                                                    │                       │
│ X-Request-Id: req-82  │   Arrange requests, HTTP details and related logs                    │                       │
│   ▾ {                 │   Layout: focused                                                    │                       │
│       "parcel_id": "p_│   Arrange requests, HTTP details and related logs                    │                       │
│       "service": "expr│ NORMAL · 1/4 · 4 matches                                             │                       │
│     }                 │ k/↑/j/↓ move · i search · Esc close                                  │                       │
│                       └──────────────────────────────────────────────────────────────────────┘                       │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│ ┌ Log · log-3 ────────────────────────────────────────────────────────────────────────────────────────────────────┐│ │
│ │ 12:04:01.200 ≡ WARN  Carrier response delayed                                                                 ┃ ││ │
│ └─────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘┃ │
│        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                                                             ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
k/↑/j/↓ select · F follow · t scope · z fold pane                                                                       
```

Vim narrow · truecolor · 40 × 20 terminal cells

```
Request req-82                          
┌ ▾ HTTP requests ·  ‖ PAUSED  ────────┐
│     GET       200  /v1/parcels       │
┌ Commands ────────────────────────────┐
│ > Search commands…                   │
│                                      │
│ → Layout: auto               current │
│   Layout: sidebar                    │
│   Layout: timeline           current │
│   Layout: columns                    │
│   Layout: focused                    │
│   Focus: HTTP requests         Alt+1 │
│   Collapse pane: HTTP requests       │
│   Focus: Request               Alt+4 │
│ NORMAL · 1/27 · 27 matches           │
│ k/↑/j/↓ move · i search · Esc close  │
└──────────────────────────────────────┘
└──────────────────────────────────────┘
Tab panes · ? help · Ctrl+K commands    
k/↑/j/↓ select · F follow · t scope · z 
```

Commands · truecolor · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   api · api.example.test:443 ─────────────────────────────────────────────────────────────────────────────────────── │
│   GET /v1/parcels                       ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤           ┊              ┊   200     2.2s │
│ > POST /v1/labels                       ┊    ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤            ┊   502     2.8s │
│   GET /health                           ┊              ┊        ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━▷   ◷       2.4s │
│   carrier · carrier.example:8443 ─────────────────────────────────────────────────────────────────────────────────── │
│   PATCH /v1/shipments/┌ Commands ────────────────────────────────────────────────────────────┐      ┊   204    600ms │
│   QUERY /v1/rates     │ > Search commands…                                                   │      ┊   200    900ms │
└───────────────────────│                                                                      │───────────────────────┘
POST /v1/labels         │ → Layout: auto                                                       │                        
502 Bad Gateway · 2.8s  │   Adapt to terminal size and timeline content                        │                        
2026-11-08T12:04:00.35Z │   Layout: sidebar                                                    │                        
Trace ID   trace-c7     │   Arrange requests, HTTP details and related logs                    │                        
┌ ▾ Request ────────────│   Layout: timeline                                           current │───────────────────────┐
│ ▾ Parameters          │   Arrange requests, HTTP details and related logs                    │tf-8                   │
│ Query                 │   Layout: columns                                                    │                       │
│ expand: "rate"        │   Arrange requests, HTTP details and related logs                    │                       │
│ expand: "tracking"    │   Layout: focused                                                    │                       │
│ locale: "en_US"       │   Arrange requests, HTTP details and related logs                    │                       │
│                       │   Focus: HTTP requests                               current · Alt+1 │                       │
│ Content-Type: applicat│   Move keyboard focus to this pane                                   │                       │
│ X-Request-Id: req-82  │   Collapse pane: HTTP requests                                       │                       │
│   ▾ {                 │   Return this pane's space to expanded content                       │                       │
│       "parcel_id": "p_│   Focus: Request                                               Alt+4 │                       │
│       "service": "expr│   Move keyboard focus to this pane                                   │                       │
│     }                 │ 1/27 · 27 matches                                                    │                       │
└───────────────────────│ ↑/↓ select · Enter run · Esc cancel                                  │───────────────────────┘
┌ ▾ Related logs ·  ‖ PA└──────────────────────────────────────────────────────────────────────┘───────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│ ┌ Log · log-3 ────────────────────────────────────────────────────────────────────────────────────────────────────┐│ │
│ │ 12:04:01.200 ≡ WARN  Carrier response delayed                                                                 ┃ ││ │
│ └─────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘┃ │
│        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                                                             ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope · z fold pane                                                                       
```

Commands filtered · truecolor · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   api · api.example.test:443 ─────────────────────────────────────────────────────────────────────────────────────── │
│   GET /v1/parcels                       ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤           ┊              ┊   200     2.2s │
│ > POST /v1/labels                       ┊    ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤            ┊   502     2.8s │
│   GET /health                           ┊              ┊        ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━▷   ◷       2.4s │
│   carrier · carrier.example:8443 ─────────────────────────────────────────────────────────────────────────────────── │
│   PATCH /v1/shipments/p_42              ┊            ├━━━━━━━━┤       ┊              ┊              ┊   204    600ms │
│   QUERY /v1/rates                       ┊              ┊              ┊      ├━━━━━━━━━━━━━┤        ┊   200    900ms │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v1/labels                                                                                                         
502 Bad Gateway · 2.8s  ┌ Commands ────────────────────────────────────────────────────────────┐                        
2026-11-08T12:04:00.35Z │ > laytm                                                              │                        
Trace ID   trace-c7     │                                                                      │                        
┌ ▾ Request ────────────│ → Layout: timeline                                           current │───────────────────────┐
│ ▾ Parameters          │   Arrange requests, HTTP details and related logs                    │tf-8                   │
│ Query                 │   Layout: sidebar                                                    │                       │
│ expand: "rate"        │   Arrange requests, HTTP details and related logs                    │                       │
│ expand: "tracking"    │   Layout: columns                                                    │                       │
│ locale: "en_US"       │   Arrange requests, HTTP details and related logs                    │                       │
│                       │   Layout: focused                                                    │                       │
│ Content-Type: applicat│   Arrange requests, HTTP details and related logs                    │                       │
│ X-Request-Id: req-82  │ 1/4 · 4 matches                                                      │                       │
│   ▾ {                 │ ↑/↓ select · Enter run · Esc cancel                                  │                       │
│       "parcel_id": "p_└──────────────────────────────────────────────────────────────────────┘                       │
│       "service": "express"               │ │                                                                         │
│     }                                    │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│ ┌ Log · log-3 ────────────────────────────────────────────────────────────────────────────────────────────────────┐│ │
│ │ 12:04:01.200 ≡ WARN  Carrier response delayed                                                                 ┃ ││ │
│ └─────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘┃ │
│        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                                                             ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope · z fold pane                                                                       
```

Commands highlighted · truecolor · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   api · api.example.test:443 ─────────────────────────────────────────────────────────────────────────────────────── │
│   GET /v1/parcels                       ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤           ┊              ┊   200     2.2s │
│ > POST /v1/labels                       ┊    ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤            ┊   502     2.8s │
│   GET /health                           ┊              ┊        ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━▷   ◷       2.4s │
│   carrier · carrier.example:8443 ─────────────────────────────────────────────────────────────────────────────────── │
│   PATCH /v1/shipments/p_42              ┊            ├━━━━━━━━┤       ┊              ┊              ┊   204    600ms │
│   QUERY /v1/rates                       ┊              ┊              ┊      ├━━━━━━━━━━━━━┤        ┊   200    900ms │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v1/labels                                                                                                         
502 Bad Gateway · 2.8s  ┌ Commands ────────────────────────────────────────────────────────────┐                        
2026-11-08T12:04:00.35Z │ > laytm                                                              │                        
Trace ID   trace-c7     │                                                                      │                        
┌ ▾ Request ────────────│ → Layout: timeline                                           current │───────────────────────┐
│ ▾ Parameters          │   Arrange requests, HTTP details and related logs                    │tf-8                   │
│ Query                 │   Layout: sidebar                                                    │                       │
│ expand: "rate"        │   Arrange requests, HTTP details and related logs                    │                       │
│ expand: "tracking"    │   Layout: columns                                                    │                       │
│ locale: "en_US"       │   Arrange requests, HTTP details and related logs                    │                       │
│                       │   Layout: focused                                                    │                       │
│ Content-Type: applicat│   Arrange requests, HTTP details and related logs                    │                       │
│ X-Request-Id: req-82  │ 1/4 · 4 matches                                                      │                       │
│   ▾ {                 │ ↑/↓ select · Enter run · Esc cancel                                  │                       │
│       "parcel_id": "p_└──────────────────────────────────────────────────────────────────────┘                       │
│       "service": "express"               │ │                                                                         │
│     }                                    │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│ ┌ Log · log-3 ────────────────────────────────────────────────────────────────────────────────────────────────────┐│ │
│ │ 12:04:01.200 ≡ WARN  Carrier response delayed                                                                 ┃ ││ │
│ └─────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘┃ │
│        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                                                             ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope · z fold pane                                                                       
```

Commands unavailable · truecolor · 120 × 38 terminal cells

```
Request unavailable                                                                                                     
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   api · api.example.test:443 ─────────────────────────────────────────────────────────────────────────────────────── │
│   GET /v1/parcels                       ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤           ┊              ┊   200     2.2s │
│   POST /v1/labels                       ┊    ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤            ┊   502     2.8s │
│ > GET /health                           ┊              ┊        ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━▷   ◷       2.4s │
│   carrier · carrier.example:8443 ─────────────────────────────────────────────────────────────────────────────────── │
│   PATCH /v1/shipments/p_42              ┊            ├━━━━━━━━┤       ┊              ┊              ┊   204    600ms │
│   QUERY /v1/rates                       ┊              ┊              ┊      ├━━━━━━━━━━━━━┤        ┊   200    900ms │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
GET /health                                                                                                             
◷ PENDING  · elapsed unavailable                                                                                        
2026-11-08T12:04:01.6Z                                                                                                  
┌ ▾ Request ───────────────────────────────┐ ┌ ▾ Response ─────────────────────────────────────────────────────────────┐
│ No captured headers   ┌ Commands ────────────────────────────────────────────────────────────┐                       │
│ Body not captured     │ > Scope: trace                                                       │                       │
│                       │                                                                      │                       │
│                       │ → Scope: trace                                           unavailable │                       │
│                       │   Selected request has no trace ID                                   │                       │
│                       │ 1/1 · 1 matches                                                      │                       │
│                       │ ↑/↓ select · Enter run · Esc cancel                                  │                       │
│                       └──────────────────────────────────────────────────────────────────────┘                       │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ Correlation unavailable: no request ID                                                                               │
│ → ▸  Logs · 0 retained                                                                                               │
│                                                                                                                      │
│                                                                                                                      │
│                                                                                                                      │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · z fold pane                                                                                 
```

Commands empty · truecolor · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   api · api.example.test:443 ─────────────────────────────────────────────────────────────────────────────────────── │
│   GET /v1/parcels                       ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤           ┊              ┊   200     2.2s │
│ > POST /v1/labels                       ┊    ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤            ┊   502     2.8s │
│   GET /health                           ┊              ┊        ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━▷   ◷       2.4s │
│   carrier · carrier.example:8443 ─────────────────────────────────────────────────────────────────────────────────── │
│   PATCH /v1/shipments/p_42              ┊            ├━━━━━━━━┤       ┊              ┊              ┊   204    600ms │
│   QUERY /v1/rates                       ┊              ┊              ┊      ├━━━━━━━━━━━━━┤        ┊   200    900ms │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v1/labels                                                                                                         
502 Bad Gateway · 2.8s                                                                                                  
2026-11-08T12:04:00.35Z                                                                                                 
Trace ID   trace-c7                                                                                                     
┌ ▾ Request ────────────┌ Commands ────────────────────────────────────────────────────────────┐───────────────────────┐
│ ▾ Parameters          │ > zzzzzz                                                             │tf-8                   │
│ Query                 │                                                                      │                       │
│ expand: "rate"        │ No matching commands                                                 │                       │
│ expand: "tracking"    │                                                                      │                       │
│ locale: "en_US"       │ 0/0 · 0 matches                                                      │                       │
│                       │ ↑/↓ select · Enter run · Esc cancel                                  │                       │
│ Content-Type: applicat└──────────────────────────────────────────────────────────────────────┘                       │
│ X-Request-Id: req-82                     │ │                                                                         │
│   ▾ {                                    │ │                                                                         │
│       "parcel_id": "p_42",               │ │                                                                         │
│       "service": "express"               │ │                                                                         │
│     }                                    │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│ ┌ Log · log-3 ────────────────────────────────────────────────────────────────────────────────────────────────────┐│ │
│ │ 12:04:01.200 ≡ WARN  Carrier response delayed                                                                 ┃ ││ │
│ └─────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘┃ │
│        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                                                             ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope · z fold pane                                                                       
```

Commands narrow · truecolor · 40 × 20 terminal cells

```
Request req-82                          
┌ ▾ HTTP requests ·  ‖ PAUSED  ────────┐
│     GET       200  /v1/parcels       │
┌ Commands ────────────────────────────┐
│ > Search commands…                   │
│                                      │
│ → Layout: auto               current │
│   Layout: sidebar                    │
│   Layout: timeline           current │
│   Layout: columns                    │
│   Layout: focused                    │
│   Focus: HTTP requests         Alt+1 │
│   Collapse pane: HTTP requests       │
│   Focus: Request               Alt+4 │
│ 1/27 · 27 matches                    │
│ ↑/↓ select · Enter run · Esc cancel  │
└──────────────────────────────────────┘
└──────────────────────────────────────┘
Tab panes · ? help · Ctrl+K commands    
↑/k/↓/j select · F follow · t scope · z 
```

light · truecolor · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   GET /v1/parcels                       ┤              ┊              ┊              ┊              ┊   200     12ms │
│ > POST /v1/labels                       ┊              ├━━━━━━━━━━━┤  ┊              ┊              ┊   502    842ms │
│   PATCH /v1/parcels/p_42                ┊              ┊              ┤              ┊              ┊   204      8ms │
│   GET /health                           ┊              ┊              ┊              ├━━━━━━━━━━━━━━▷   ◷         1s │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
Trace ID   trace-c7                                                                                                     
┌ ▾ Request ───────────────────────────────┐ ┌ ▾ Response ─────────────────────────────────────────────────────────────┐
│ ▾ Parameters                             │ │ Content-Type: application/problem+json; charset=utf-8                   │
│ Query                                    │ │ Retry-After: 1                                                          │
│ expand: "rate"                           │ │   ▾ {                                                                   │
│ expand: "tracking"                       │ │       "error": "upstream_timeout",                                      │
│ locale: "en_US"                          │ │       "retryable": true,                                                │
│                                          │ │       "attempt": 2                                                      │
│ Content-Type: application/json           │ │     }                                                                   │
│ X-Request-Id: req-82                     │ │                                                                         │
│   ▾ {                                    │ │                                                                         │
│       "parcel_id": "p_42",               │ │                                                                         │
│       "service": "express"               │ │                                                                         │
│     }                                    │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                                                             │ │
│ →      └── ▾  12:04:01.000 ≡ INFO  Create label                                                                    │ │
│           ├── ▾  time: 2026-11-08T12:04:01Z                                                                        ┃ │
│           ├── ▸  parcel {2}                                                                                        ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope · z fold pane
```

contrast · truecolor · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   GET /v1/parcels                       ┤              ┊              ┊              ┊              ┊   200     12ms │
│ > POST /v1/labels                       ┊              ├━━━━━━━━━━━┤  ┊              ┊              ┊   502    842ms │
│   PATCH /v1/parcels/p_42                ┊              ┊              ┤              ┊              ┊   204      8ms │
│   GET /health                           ┊              ┊              ┊              ├━━━━━━━━━━━━━━▷   ◷         1s │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
Trace ID   trace-c7                                                                                                     
┌ ▾ Request ───────────────────────────────┐ ┌ ▾ Response ─────────────────────────────────────────────────────────────┐
│ ▾ Parameters                             │ │ Content-Type: application/problem+json; charset=utf-8                   │
│ Query                                    │ │ Retry-After: 1                                                          │
│ expand: "rate"                           │ │   ▾ {                                                                   │
│ expand: "tracking"                       │ │       "error": "upstream_timeout",                                      │
│ locale: "en_US"                          │ │       "retryable": true,                                                │
│                                          │ │       "attempt": 2                                                      │
│ Content-Type: application/json           │ │     }                                                                   │
│ X-Request-Id: req-82                     │ │                                                                         │
│   ▾ {                                    │ │                                                                         │
│       "parcel_id": "p_42",               │ │                                                                         │
│       "service": "express"               │ │                                                                         │
│     }                                    │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                                                             │ │
│ →      └── ▾  12:04:01.000 ≡ INFO  Create label                                                                    │ │
│           ├── ▾  time: 2026-11-08T12:04:01Z                                                                        ┃ │
│           ├── ▸  parcel {2}                                                                                        ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope · z fold pane
```

Empty · truecolor · 120 × 38 terminal cells

```
┌ ▾ Request timeline ·  ▶ FOLLOW  ─────────────────────────────────────────────────────────────────────────────────────┐
│ No HTTP requests                                                                                                     │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Select an HTTP request                                                                                                  
┌ ▾ Request ───────────────────────────────┐ ┌ ▾ Response ─────────────────────────────────────────────────────────────┐
│ Select an HTTP request                   │ │ Select an HTTP request                                                  │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ Select an HTTP request                                                                                               │
│ → ▸  Logs · 0 retained                                                                                               │
│                                                                                                                      │
│                                                                                                                      │
│                                                                                                                      │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · z fold pane
```

Error · truecolor · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   GET /v1/parcels                       ┤              ┊              ┊              ┊              ┊   200     12ms │
│ > POST /v1/labels                       ┊              ├━━━━━━━━━━━┤  ┊              ┊              ┊   502    842ms │
│   PATCH /v1/parcels/p_42                ┊              ┊              ┤              ┊              ┊   204      8ms │
│   GET /health                           ┊              ┊              ┊              ├━━━━━━━━━━━━━━▷   ◷         1s │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
Trace ID   trace-c7                                                                                                     
┌ ▾ Request ───────────────────────────────┐ ┌ ▾ Response ─────────────────────────────────────────────────────────────┐
│ ▾ Parameters                             │ │ Content-Type: application/problem+json; charset=utf-8                   │
│ Query                                    │ │ Retry-After: 1                                                          │
│ expand: "rate"                           │ │ {"error":                                                               │
│ expand: "tracking"                       │ │ [invalid JSON: showing captured source]                                 │
│ locale: "en_US"                          │ │                                                                         │
│                                          │ │                                                                         │
│ Content-Type: application/json           │ │                                                                         │
│ X-Request-Id: req-82                     │ │                                                                         │
│   ▾ {                                    │ │                                                                         │
│       "parcel_id": "p_42",               │ │                                                                         │
│       "service": "express"               │ │                                                                         │
│     }                                    │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                                                             │ │
│ →      └── ▾  12:04:01.000 ≡ INFO  Create label                                                                    │ │
│           ├── ▾  time: 2026-11-08T12:04:01Z                                                                        ┃ │
│           ├── ▸  parcel {2}                                                                                        ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope · z fold pane
```

Loading · truecolor · 120 × 38 terminal cells

```
Request unavailable                                                                                                     
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   GET /v1/parcels                       ┤              ┊              ┊              ┊              ┊   200     12ms │
│   POST /v1/labels                       ┊              ├━━━━━━━━━━━┤  ┊              ┊              ┊   502    842ms │
│   PATCH /v1/parcels/p_42                ┊              ┊              ┤              ┊              ┊   204      8ms │
│ > GET /health                           ┊              ┊              ┊              ├━━━━━━━━━━━━━━▷   ◷         1s │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
GET /health                                                                                                             
◷ PENDING  · elapsed unavailable                                                                                        
2026-11-08T12:04:03Z                                                                                                    
┌ ▾ Request ───────────────────────────────┐ ┌ ▾ Response ─────────────────────────────────────────────────────────────┐
│ No captured headers                      │ │ No captured headers                                                     │
│ Body not captured                        │ │ Body not captured                                                       │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ Correlation unavailable: no request ID                                                                               │
│ → ▸  Logs · 0 retained                                                                                               │
│                                                                                                                      │
│                                                                                                                      │
│                                                                                                                      │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · z fold pane
```

Glyph only · truecolor · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌ ▾ Request timeline ·  ‖  ────────────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   GET /v1/parcels                       ┤              ┊              ┊              ┊              ┊   200     12ms │
│ > POST /v1/labels                       ┊              ├━━━━━━━━━━━┤  ┊              ┊              ┊   502    842ms │
│   PATCH /v1/parcels/p_42                ┊              ┊              ┤              ┊              ┊   204      8ms │
│   GET /health                           ┊              ┊              ┊              ├━━━━━━━━━━━━━━▷   ◷         1s │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
Trace ID   trace-c7                                                                                                     
┌ ▾ Request ───────────────────────────────┐ ┌ ▾ Response ─────────────────────────────────────────────────────────────┐
│ ▾ Parameters                             │ │ Content-Type: application/problem+json; charset=utf-8                   │
│ Query                                    │ │ Retry-After: 1                                                          │
│ expand: "rate"                           │ │   ▾ {                                                                   │
│ expand: "tracking"                       │ │       "error": "upstream_timeout",                                      │
│ locale: "en_US"                          │ │       "retryable": true,                                                │
│                                          │ │       "attempt": 2                                                      │
│ Content-Type: application/json           │ │     }                                                                   │
│ X-Request-Id: req-82                     │ │                                                                         │
│   ▾ {                                    │ │                                                                         │
│       "parcel_id": "p_42",               │ │                                                                         │
│       "service": "express"               │ │                                                                         │
│     }                                    │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖  ────────────────────────────────────────────────────────────────────────────────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                                                             │ │
│ →      └── ▾  12:04:01.000 ≡ INFO  Create label                                                                    │ │
│           ├── ▾  time: 2026-11-08T12:04:01Z                                                                        ┃ │
│           ├── ▸  parcel {2}                                                                                        ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope · z fold pane
```

Label only · truecolor · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌  Request timeline ·  PAUSED  ────────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   GET /v1/parcels                       ┤              ┊              ┊              ┊              ┊   200     12ms │
│ > POST /v1/labels                       ┊              ├━━━━━━━━━━━┤  ┊              ┊              ┊   502    842ms │
│   PATCH /v1/parcels/p_42                ┊              ┊              ┤              ┊              ┊   204      8ms │
│   GET /health                           ┊              ┊              ┊              ├━━━━━━━━━━━━━━▷   ◷         1s │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
Trace ID   trace-c7                                                                                                     
┌  Request ────────────────────────────────┐ ┌  Response ──────────────────────────────────────────────────────────────┐
│  Parameters                              │ │ Content-Type: application/problem+json; charset=utf-8                   │
│ Query                                    │ │ Retry-After: 1                                                          │
│ expand: "rate"                           │ │     {                                                                   │
│ expand: "tracking"                       │ │       "error": "upstream_timeout",                                      │
│ locale: "en_US"                          │ │       "retryable": true,                                                │
│                                          │ │       "attempt": 2                                                      │
│ Content-Type: application/json           │ │     }                                                                   │
│ X-Request-Id: req-82                     │ │                                                                         │
│     {                                    │ │                                                                         │
│       "parcel_id": "p_42",               │ │                                                                         │
│       "service": "express"               │ │                                                                         │
│     }                                    │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌  Related logs ·  PAUSED  ────────────────────────────────────────────────────────────────────────────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│        │  └── ▸  12:04:01.842  ERROR Upstream timeout                                                              │ │
│ →      └── ▾  12:04:01.000  INFO  Create label                                                                     │ │
│           ├── ▾  time: 2026-11-08T12:04:01Z                                                                        ┃ │
│           ├── ▸  parcel {2}                                                                                        ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope · z fold pane
```

Focused leaf · truecolor · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   GET /v1/parcels                       ┤              ┊              ┊              ┊              ┊   200     12ms │
│ > POST /v1/labels                       ┊              ├━━━━━━━━━━━┤  ┊              ┊              ┊   502    842ms │
│   PATCH /v1/parcels/p_42                ┊              ┊              ┤              ┊              ┊   204      8ms │
│   GET /health                           ┊              ┊              ┊              ├━━━━━━━━━━━━━━▷   ◷         1s │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
Trace ID   trace-c7                                                                                                     
┌ ▾ Request ───────────────────────────────┐ ┌ ▾ Response ─────────────────────────────────────────────────────────────┐
│ ▾ Parameters                             │ │ Content-Type: application/problem+json; charset=utf-8                   │
│ Query                                    │ │ Retry-After: 1                                                          │
│ expand: "rate"                           │ │   ▾ {                                                                   │
│ expand: "tracking"                       │ │       "error": "upstream_timeout",                                      │
│ locale: "en_US"                          │ │       "retryable": true,                                                │
│                                          │ │       "attempt": 2                                                      │
│ Content-Type: application/json           │ │     }                                                                   │
│ X-Request-Id: req-82                     │ │                                                                         │
│   ▾ {                                    │ │                                                                         │
│       "parcel_id": "p_42",               │ │                                                                         │
│       "service": "express"               │ │                                                                         │
│     }                                    │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│ ┌ Log · log-3 ────────────────────────────────────────────────────────────────────────────────────────────────────┐│ │
│ │ 12:04:01.200 ≡ WARN  Carrier response delayed                                                                 ┃ ││ │
│ └─────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘┃ │
│        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                                                             ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope · z fold pane
```

URL full · truecolor · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   GET https://api.example.test/v1/par…  ┤              ┊              ┊              ┊              ┊   200     12ms │
│ > POST https://api.example.test/v1/la…  ┊              ├━━━━━━━━━━━┤  ┊              ┊              ┊   502    842ms │
│   PATCH https://api.example.test/v1/p…  ┊              ┊              ┤              ┊              ┊   204      8ms │
│   GET https://api.example.test/health   ┊              ┊              ┊              ├━━━━━━━━━━━━━━▷   ◷         1s │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST https://api.example.test/v1/labels?expand=rate&expand=tracking&locale=en_US                                        
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
Trace ID   trace-c7                                                                                                     
┌ ▾ Request ───────────────────────────────┐ ┌ ▾ Response ─────────────────────────────────────────────────────────────┐
│ ▾ Parameters                             │ │ Content-Type: application/problem+json; charset=utf-8                   │
│ Query                                    │ │ Retry-After: 1                                                          │
│ expand: "rate"                           │ │   ▾ {                                                                   │
│ expand: "tracking"                       │ │       "error": "upstream_timeout",                                      │
│ locale: "en_US"                          │ │       "retryable": true,                                                │
│                                          │ │       "attempt": 2                                                      │
│ Content-Type: application/json           │ │     }                                                                   │
│ X-Request-Id: req-82                     │ │                                                                         │
│   ▾ {                                    │ │                                                                         │
│       "parcel_id": "p_42",               │ │                                                                         │
│       "service": "express"               │ │                                                                         │
│     }                                    │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                                                             │ │
│ →      └── ▾  12:04:01.000 ≡ INFO  Create label                                                                    │ │
│           ├── ▾  time: 2026-11-08T12:04:01Z                                                                        ┃ │
│           ├── ▸  parcel {2}                                                                                        ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope · z fold pane
```

URL host-path · truecolor · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   GET api.example.test/v1/parcels       ┤              ┊              ┊              ┊              ┊   200     12ms │
│ > POST api.example.test/v1/labels?exp…  ┊              ├━━━━━━━━━━━┤  ┊              ┊              ┊   502    842ms │
│   PATCH api.example.test/v1/parcels/p…  ┊              ┊              ┤              ┊              ┊   204      8ms │
│   GET api.example.test/health           ┊              ┊              ┊              ├━━━━━━━━━━━━━━▷   ◷         1s │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST api.example.test/v1/labels?expand=rate&expand=tracking&locale=en_US                                                
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
Trace ID   trace-c7                                                                                                     
┌ ▾ Request ───────────────────────────────┐ ┌ ▾ Response ─────────────────────────────────────────────────────────────┐
│ ▾ Parameters                             │ │ Content-Type: application/problem+json; charset=utf-8                   │
│ Query                                    │ │ Retry-After: 1                                                          │
│ expand: "rate"                           │ │   ▾ {                                                                   │
│ expand: "tracking"                       │ │       "error": "upstream_timeout",                                      │
│ locale: "en_US"                          │ │       "retryable": true,                                                │
│                                          │ │       "attempt": 2                                                      │
│ Content-Type: application/json           │ │     }                                                                   │
│ X-Request-Id: req-82                     │ │                                                                         │
│   ▾ {                                    │ │                                                                         │
│       "parcel_id": "p_42",               │ │                                                                         │
│       "service": "express"               │ │                                                                         │
│     }                                    │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                                                             │ │
│ →      └── ▾  12:04:01.000 ≡ INFO  Create label                                                                    │ │
│           ├── ▾  time: 2026-11-08T12:04:01Z                                                                        ┃ │
│           ├── ▸  parcel {2}                                                                                        ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope · z fold pane
```

URL path · truecolor · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   GET /v1/parcels                       ┤              ┊              ┊              ┊              ┊   200     12ms │
│ > POST /v1/labels                       ┊              ├━━━━━━━━━━━┤  ┊              ┊              ┊   502    842ms │
│   PATCH /v1/parcels/p_42                ┊              ┊              ┤              ┊              ┊   204      8ms │
│   GET /health                           ┊              ┊              ┊              ├━━━━━━━━━━━━━━▷   ◷         1s │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
Trace ID   trace-c7                                                                                                     
┌ ▾ Request ───────────────────────────────┐ ┌ ▾ Response ─────────────────────────────────────────────────────────────┐
│ ▾ Parameters                             │ │ Content-Type: application/problem+json; charset=utf-8                   │
│ Query                                    │ │ Retry-After: 1                                                          │
│ expand: "rate"                           │ │   ▾ {                                                                   │
│ expand: "tracking"                       │ │       "error": "upstream_timeout",                                      │
│ locale: "en_US"                          │ │       "retryable": true,                                                │
│                                          │ │       "attempt": 2                                                      │
│ Content-Type: application/json           │ │     }                                                                   │
│ X-Request-Id: req-82                     │ │                                                                         │
│   ▾ {                                    │ │                                                                         │
│       "parcel_id": "p_42",               │ │                                                                         │
│       "service": "express"               │ │                                                                         │
│     }                                    │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                                                             │ │
│ →      └── ▾  12:04:01.000 ≡ INFO  Create label                                                                    │ │
│           ├── ▾  time: 2026-11-08T12:04:01Z                                                                        ┃ │
│           ├── ▸  parcel {2}                                                                                        ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope · z fold pane
```

Labeled summary · truecolor · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   GET /v1/parcels                       ┤              ┊              ┊              ┊              ┊   200     12ms │
│ > POST /v1/labels                       ┊              ├━━━━━━━━━━━┤  ┊              ┊              ┊   502    842ms │
│   PATCH /v1/parcels/p_42                ┊              ┊              ┤              ┊              ┊   204      8ms │
│   GET /health                           ┊              ┊              ┊              ├━━━━━━━━━━━━━━▷   ◷         1s │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Method     POST                                                                                                         
Target     /v1/labels                                                                                                   
Time       2026-11-08T12:04:01Z                                                                                         
Latency    842ms                                                                                                        
Status     502 Bad Gateway                                                                                              
Trace ID   trace-c7                                                                                                     
┌ ▾ Request ───────────────────────────────┐ ┌ ▾ Response ─────────────────────────────────────────────────────────────┐
│ ▾ Parameters                             │ │ Content-Type: application/problem+json; charset=utf-8                   │
│ Query                                    │ │ Retry-After: 1                                                          │
│ expand: "rate"                           │ │   ▾ {                                                                   │
│ expand: "tracking"                       │ │       "error": "upstream_timeout",                                      │
│ locale: "en_US"                          │ │       "retryable": true,                                                │
│                                          │ │       "attempt": 2                                                      │
│ Content-Type: application/json           │ │     }                                                                   │
│ X-Request-Id: req-82                     │ │                                                                         │
│   ▾ {                                    │ │                                                                         │
│       "parcel_id": "p_42",               │ │                                                                         │
│       "service": "express"               │ │                                                                         │
│     }                                    │ │                                                                         │
│                                          │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                                                             │ │
│ →      └── ▾  12:04:01.000 ≡ INFO  Create label                                                                    │ │
│           ├── ▾  time: 2026-11-08T12:04:01Z                                                                        ┃ │
│           ├── ▸  parcel {2}                                                                                        ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope · z fold pane
```

Single request · truecolor · 160 × 44 terminal cells

```
Request 01234567-89ab-cdef-0123-456789abcdef                                                                                                                    
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:01.000Z                           0                        +3ms                     +6ms                     +9ms                 +13ms                │
│ > POST /v2/shipments/shp_fixture/buy    ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤   500     13ms │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v2/shipments/shp_fixture/buy                                                                                      ┌ ▾ Related logs ·  ‖ PAUSED  ─────────┐
500 Internal Server Error · 12.5ms                                                                                      │ request 01234567-89ab-cdef-0123-4567 │
2026-11-08T12:04:01Z                                                                                                    │   ▾  Logs · 1 retained               │
┌ ▾ Request ───────────────────────────────┐ ┌ ▾ Response ────────────────────────────────────────────────────────────┐ │ ┌ Log · log-2 ────────────────────┐  │
│ Content-Type: application/json           │ │ Content-Type: application/json                                         │ │ │ ≡ INFO  Create label            │  │
│ X-Ep-Request-Uuid:                       │ │ X-Ep-Request-Uuid: 01234567-89ab-cdef-0123-456789abcdef                │ │ │ ├──time: 2026-11-08T12:04:01Z   │  │
│ 01234567-89ab-cdef-0123-456789abcdef     │ │ → ▾ {                                                                  │ │ │ ├──parcel {2}                   │  │
│   ▾ {                                    │ │   ▾   "error": {                                                       │ │ │ │  ├──id: "p_42"                │  │
│   ▾   "rate": {                          │ │         "code": "INTERNAL_SERVER_ERROR",                               │ │ │ │  └──service: "express"        │  │
│         "id": "rate_fixture"             │ │         "message": "Could not purchase shipment: the carrier returned  │ │ │ └──attempt: 2                   │  │
│       }                                  │ │      an internal server error"                                         │ │ └─────────────────────────────────┘  │
│     }                                    │ │       }                                                                │ │                                      │
│                                          │ │     }                                                                  │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
└──────────────────────────────────────────┘ └────────────────────────────────────────────────────────────────────────┘ └──────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                                                             
↑/k/↓/j fields · Enter fold JSON · r raw · z fold pane
```

Folded logs · truecolor · 160 × 44 terminal cells

```
Request 01234567-89ab-cdef-0123-456789abcdef                                                                                                                    
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:01.000Z                           0                        +3ms                     +6ms                     +9ms                 +13ms                │
│ > POST /v2/shipments/shp_fixture/buy    ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤   500     13ms │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v2/shipments/shp_fixture/buy                                                                                                                              
500 Internal Server Error · 12.5ms                                                                                                                              
2026-11-08T12:04:01Z                                                                                                                                            
┌ ▾ Request ────────────────────────────────────────┐ ┌ ▾ Response ────────────────────────────────────────────────────────────────────────────────────────────┐
│ Content-Type: application/json                    │ │ Content-Type: application/json                                                                         │
│ X-Ep-Request-Uuid:                                │ │ X-Ep-Request-Uuid: 01234567-89ab-cdef-0123-456789abcdef                                                │
│ 01234567-89ab-cdef-0123-456789abcdef              │ │ → ▾ {                                                                                                  │
│   ▾ {                                             │ │   ▾   "error": {                                                                                       │
│   ▾   "rate": {                                   │ │         "code": "INTERNAL_SERVER_ERROR",                                                               │
│         "id": "rate_fixture"                      │ │         "message": "Could not purchase shipment: the carrier returned an internal server error"        │
│       }                                           │ │       }                                                                                                │
│     }                                             │ │     }                                                                                                  │
│                                                   │ │                                                                                                        │
│                                                   │ │                                                                                                        │
│                                                   │ │                                                                                                        │
│                                                   │ │                                                                                                        │
│                                                   │ │                                                                                                        │
│                                                   │ │                                                                                                        │
│                                                   │ │                                                                                                        │
│                                                   │ │                                                                                                        │
│                                                   │ │                                                                                                        │
│                                                   │ │                                                                                                        │
│                                                   │ │                                                                                                        │
│                                                   │ │                                                                                                        │
│                                                   │ │                                                                                                        │
│                                                   │ │                                                                                                        │
│                                                   │ │                                                                                                        │
│                                                   │ │                                                                                                        │
│                                                   │ │                                                                                                        │
│                                                   │ │                                                                                                        │
│                                                   │ │                                                                                                        │
│                                                   │ │                                                                                                        │
│                                                   │ │                                                                                                        │
│                                                   │ │                                                                                                        │
│                                                   │ │                                                                                                        │
└───────────────────────────────────────────────────┘ └────────────────────────────────────────────────────────────────────────────────────────────────────────┘
▸ Related logs ·  ‖ PAUSED                                                                                                                                      
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                                                             
↑/k/↓/j fields · Enter fold JSON · r raw · z fold pane
```

Request folded · truecolor · 160 × 44 terminal cells

```
Request 01234567-89ab-cdef-0123-456789abcdef                                                                                                                    
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:01.000Z                           0                        +3ms                     +6ms                     +9ms                 +13ms                │
│ > POST /v2/shipments/shp_fixture/buy    ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤   500     13ms │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v2/shipments/shp_fixture/buy                                                                                      ┌ ▾ Related logs ·  ‖ PAUSED  ─────────┐
500 Internal Server Error · 12.5ms                                                                                      │ request 01234567-89ab-cdef-0123-4567 │
2026-11-08T12:04:01Z                                                                                                    │   ▾  Logs · 1 retained               │
▸ Request                                                                                                               │ ┌ Log · log-2 ────────────────────┐  │
┌ ▾ Response ─────────────────────────────────────────────────────────────────────────────────────────────────────────┐ │ │ ≡ INFO  Create label            │  │
│ Content-Type: application/json                                                                                      │ │ │ ├──time: 2026-11-08T12:04:01Z   │  │
│ X-Ep-Request-Uuid: 01234567-89ab-cdef-0123-456789abcdef                                                             │ │ │ ├──parcel {2}                   │  │
│ → ▾ {                                                                                                               │ │ │ │  ├──id: "p_42"                │  │
│   ▾   "error": {                                                                                                    │ │ │ │  └──service: "express"        │  │
│         "code": "INTERNAL_SERVER_ERROR",                                                                            │ │ │ └──attempt: 2                   │  │
│         "message": "Could not purchase shipment: the carrier returned an internal server error"                     │ │ └─────────────────────────────────┘  │
│       }                                                                                                             │ │                                      │
│     }                                                                                                               │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
└─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘ └──────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                                                             
↑/k/↓/j fields · Enter fold JSON · r raw · z fold pane
```

Response folded · truecolor · 160 × 44 terminal cells

```
Request 01234567-89ab-cdef-0123-456789abcdef                                                                                                                    
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:01.000Z                           0                        +3ms                     +6ms                     +9ms                 +13ms                │
│ > POST /v2/shipments/shp_fixture/buy    ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤   500     13ms │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v2/shipments/shp_fixture/buy                                                                                      ┌ ▾ Related logs ·  ‖ PAUSED  ─────────┐
500 Internal Server Error · 12.5ms                                                                                      │ request 01234567-89ab-cdef-0123-4567 │
2026-11-08T12:04:01Z                                                                                                    │   ▾  Logs · 1 retained               │
┌ ▾ Request ──────────────────────────────────────────────────────────────────────────────────────────────────────────┐ │ ┌ Log · log-2 ────────────────────┐  │
│ Content-Type: application/json                                                                                      │ │ │ ≡ INFO  Create label            │  │
│ X-Ep-Request-Uuid: 01234567-89ab-cdef-0123-456789abcdef                                                             │ │ │ ├──time: 2026-11-08T12:04:01Z   │  │
│   ▾ {                                                                                                               │ │ │ ├──parcel {2}                   │  │
│   ▾   "rate": {                                                                                                     │ │ │ │  ├──id: "p_42"                │  │
│         "id": "rate_fixture"                                                                                        │ │ │ │  └──service: "express"        │  │
│       }                                                                                                             │ │ │ └──attempt: 2                   │  │
│     }                                                                                                               │ │ └─────────────────────────────────┘  │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
│                                                                                                                     │ │                                      │
└─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘ │                                      │
▸ Response                                                                                                              └──────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                                                             
Pane folded · z expand pane
```

No instructions · truecolor · 160 × 44 terminal cells

```
Request 01234567-89ab-cdef-0123-456789abcdef                                                                                                                    
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:01.000Z                           0                        +3ms                     +6ms                     +9ms                 +13ms                │
│ > POST /v2/shipments/shp_fixture/buy    ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤   500     13ms │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v2/shipments/shp_fixture/buy                                                                                      ┌ ▾ Related logs ·  ‖ PAUSED  ─────────┐
500 Internal Server Error · 12.5ms                                                                                      │ request 01234567-89ab-cdef-0123-4567 │
2026-11-08T12:04:01Z                                                                                                    │   ▾  Logs · 1 retained               │
┌ ▾ Request ───────────────────────────────┐ ┌ ▾ Response ────────────────────────────────────────────────────────────┐ │ ┌ Log · log-2 ────────────────────┐  │
│ Content-Type: application/json           │ │ Content-Type: application/json                                         │ │ │ ≡ INFO  Create label            │  │
│ X-Ep-Request-Uuid:                       │ │ X-Ep-Request-Uuid: 01234567-89ab-cdef-0123-456789abcdef                │ │ │ ├──time: 2026-11-08T12:04:01Z   │  │
│ 01234567-89ab-cdef-0123-456789abcdef     │ │ → ▾ {                                                                  │ │ │ ├──parcel {2}                   │  │
│   ▾ {                                    │ │   ▾   "error": {                                                       │ │ │ │  ├──id: "p_42"                │  │
│   ▾   "rate": {                          │ │         "code": "INTERNAL_SERVER_ERROR",                               │ │ │ │  └──service: "express"        │  │
│         "id": "rate_fixture"             │ │         "message": "Could not purchase shipment: the carrier returned  │ │ │ └──attempt: 2                   │  │
│       }                                  │ │      an internal server error"                                         │ │ └─────────────────────────────────┘  │
│     }                                    │ │       }                                                                │ │                                      │
│                                          │ │     }                                                                  │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
└──────────────────────────────────────────┘ └────────────────────────────────────────────────────────────────────────┘ └──────────────────────────────────────┘
```

Single request narrow · truecolor · 60 × 40 terminal cells

```
Request 01234567-89ab-cdef-0123-456789abcdef                
POST /v2/shipments/shp_fixture/buy                          
500 Internal Server Error · 12.5ms                          
2026-11-08T12:04:01Z                                        
┌ ▾ Request ───────────────────────────────────────────────┐
│ Content-Type: application/json                           │
│ X-Ep-Request-Uuid: 01234567-89ab-cdef-0123-456789abcdef  │
│   ▾ {                                                    │
│   ▾   "rate": {                                          │
│         "id": "rate_fixture"                             │
│       }                                                  │
│     }                                                    │
└──────────────────────────────────────────────────────────┘
┌ ▾ Response ──────────────────────────────────────────────┐
│ Content-Type: application/json                           │
│ X-Ep-Request-Uuid: 01234567-89ab-cdef-0123-456789abcdef  │
│ → ▾ {                                                    │
│   ▾   "error": {                                         │
│         "code": "INTERNAL_SERVER_ERROR",                 │
│         "message": "Could not purchase shipment: the ca  │
│     rrier returned an internal server error"             │
│       }                                                  │
│     }                                                    │
│                                                          │
│                                                          │
│                                                          │
│                                                          │
│                                                          │
│                                                          │
│                                                          │
│                                                          │
│                                                          │
│                                                          │
│                                                          │
│                                                          │
│                                                          │
│                                                          │
└──────────────────────────────────────────────────────────┘
Tab panes · ? help · Ctrl+K commands                        
↑/k/↓/j fields · Enter fold JSON · r raw · z fold pane
```

JSON folding · truecolor · 160 × 44 terminal cells

```
Request 01234567-89ab-cdef-0123-456789abcdef                                                                                                                    
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:01.000Z                           0                        +3ms                     +6ms                     +9ms                 +13ms                │
│ > POST /v2/shipments/shp_fixture/buy    ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤   500     13ms │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v2/shipments/shp_fixture/buy                                                                                      ┌ ▾ Related logs ·  ‖ PAUSED  ─────────┐
500 Internal Server Error · 12.5ms                                                                                      │ request 01234567-89ab-cdef-0123-4567 │
2026-11-08T12:04:01Z                                                                                                    │   ▾  Logs · 1 retained               │
┌ ▾ Request ───────────────────────────────┐ ┌ ▾ Response ────────────────────────────────────────────────────────────┐ │ ┌ Log · log-2 ────────────────────┐  │
│ Content-Type: application/json           │ │ Content-Type: application/json                                         │ │ │ ≡ INFO  Create label            │  │
│ X-Ep-Request-Uuid:                       │ │ X-Ep-Request-Uuid: 01234567-89ab-cdef-0123-456789abcdef                │ │ │ ├──time: 2026-11-08T12:04:01Z   │  │
│ 01234567-89ab-cdef-0123-456789abcdef     │ │   ▾ {                                                                  │ │ │ ├──parcel {2}                   │  │
│   ▾ {                                    │ │ → ▸   "shipment": {… 2 fields},                                        │ │ │ │  ├──id: "p_42"                │  │
│   ▾   "rate": {                          │ │   ▾   "rates": [                                                       │ │ │ │  └──service: "express"        │  │
│         "id": "rate_fixture"             │ │   ▾     {                                                              │ │ │ └──attempt: 2                   │  │
│       }                                  │ │           "carrier": "USPS",                                           │ │ └─────────────────────────────────┘  │
│     }                                    │ │           "rate": 8.500                                                │ │                                      │
│                                          │ │         },                                                             │ │                                      │
│                                          │ │   ▾     {                                                              │ │                                      │
│                                          │ │           "carrier": "UPS",                                            │ │                                      │
│                                          │ │           "rate": 12.50                                                │ │                                      │
│                                          │ │         }                                                              │ │                                      │
│                                          │ │       ],                                                               │ │                                      │
│                                          │ │       "attempt": 1,                                                    │ │                                      │
│                                          │ │       "attempt": 2,                                                    │ │                                      │
│                                          │ │       "sequence": 9007199254740993                                     │ │                                      │
│                                          │ │     }                                                                  │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
│                                          │ │                                                                        │ │                                      │
└──────────────────────────────────────────┘ └────────────────────────────────────────────────────────────────────────┘ └──────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                                                             
↑/k/↓/j fields · Enter fold JSON · r raw · z fold pane
```

Layout sidebar · truecolor · 120 × 44 terminal cells

```
Request req-82                                                                                                          
┌ ▾ HTTP requests ·  ‖ PAUSED  ────────┐ POST /v1/labels                                                                
│     GET       200  /v1/parcels       │ 502 Bad Gateway · 842ms                                                        
│ →   POST      502  /v1/labels        │ 2026-11-08T12:04:01Z                                                           
│     PATCH     204  /v1/parcels/p_42  │ Trace ID   trace-c7                                                            
│     GET        ◷   /health           │ ┌ ▾ Request ──────────────────────────────────────────────────────────────────┐
│                                      │ │ ▾ Parameters                                                              ┃ │
│                                      │ │ Query                                                                     │ │
│                                      │ │ expand: "rate"                                                            │ │
│                                      │ │ expand: "tracking"                                                        │ │
│                                      │ └─────────────────────────────────────────────────────────────────────────────┘
│                                      │ ┌ ▾ Response ─────────────────────────────────────────────────────────────────┐
│                                      │ │ Content-Type: application/problem+json; charset=utf-8                       │
│                                      │ │ Retry-After: 1                                                              │
│                                      │ │   ▾ {                                                                       │
│                                      │ │       "error": "upstream_timeout",                                          │
│                                      │ │       "retryable": true,                                                    │
│                                      │ │       "attempt": 2                                                          │
│                                      │ │     }                                                                       │
│                                      │ │                                                                             │
│                                      │ │                                                                             │
│                                      │ │                                                                             │
│                                      │ │                                                                             │
│                                      │ │                                                                             │
│                                      │ └─────────────────────────────────────────────────────────────────────────────┘
│                                      │ ┌ ▾ Related logs ·  ‖ PAUSED  ────────────────────────────────────────────────┐
│                                      │ │ request req-82 · 3 matching logs                                            │
│                                      │ │   ▾  Logs · 3 retained                                                      │
│                                      │ │  └── ▾  ⋈ Trace trace-c7                                                    │
│                                      │ │     └── ▾  ─ Span http-label                                                │
│                                      │ │        ├── ▾  ─ Span carrier                                                │
│                                      │ │ ┌ Log · log-3 ───────────────────────────────────────────────────────────┐  │
│                                      │ │ │ 12:04:01.200 ≡ WARN  Carrier response delayed                          │  │
│                                      │ │ │ ├──time: 2026-11-08T12:04:01.2Z                                        │  │
│                                      │ │ │ └──upstream {2}                                                        │  │
│                                      │ │ │    ├──host: "carrier.example"                                          │  │
│                                      │ │ │    └──timeout_ms: 800                                                  │  │
│                                      │ │ └────────────────────────────────────────────────────────────────────────┘  │
│                                      │ │        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                      │
│                                      │ │        └── ▸  12:04:01.000 ≡ INFO  Create label                             │
│                                      │ │                                                                             │
└──────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope · z fold pane
```

Layout timeline · truecolor · 120 × 44 terminal cells

```
Request req-82                                                                                                          
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   api · api.example.test:443 ─────────────────────────────────────────────────────────────────────────────────────── │
│   GET /v1/parcels                       ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤           ┊              ┊   200     2.2s │
│ > POST /v1/labels                       ┊    ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┤            ┊   502     2.8s │
│   GET /health                           ┊              ┊        ├━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━▷   ◷       2.4s │
│   carrier · carrier.example:8443 ─────────────────────────────────────────────────────────────────────────────────── │
│   PATCH /v1/shipments/p_42              ┊            ├━━━━━━━━┤       ┊              ┊              ┊   204    600ms │
│   QUERY /v1/rates                       ┊              ┊              ┊      ├━━━━━━━━━━━━━┤        ┊   200    900ms │
│   events · events.internal:9090 ──────────────────────────────────────────────────────────────────────────────────── │
│   POST /events                          ┊              ┊           ├━━━━━━━━┤        ┊              ┊   202    600ms │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v1/labels                                                                                                         
502 Bad Gateway · 2.8s                                                                                                  
2026-11-08T12:04:00.35Z                                                                                                 
Trace ID   trace-c7                                                                                                     
┌ ▾ Request ───────────────────────────────┐ ┌ ▾ Response ─────────────────────────────────────────────────────────────┐
│ ▾ Parameters                             │ │ Content-Type: application/problem+json; charset=utf-8                   │
│ Query                                    │ │ Retry-After: 1                                                          │
│ expand: "rate"                           │ │   ▾ {                                                                   │
│ expand: "tracking"                       │ │       "error": "upstream_timeout",                                      │
│ locale: "en_US"                          │ │       "retryable": true,                                                │
│                                          │ │       "attempt": 2                                                      │
│ Content-Type: application/json           │ │     }                                                                   │
│ X-Request-Id: req-82                     │ │                                                                         │
│   ▾ {                                    │ │                                                                         │
│       "parcel_id": "p_42",               │ │                                                                         │
│       "service": "express"               │ │                                                                         │
│     }                                    │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│ ┌ Log · log-3 ────────────────────────────────────────────────────────────────────────────────────────────────────┐│ │
│ │ 12:04:01.200 ≡ WARN  Carrier response delayed                                                                 ┃ ││ │
│ └─────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘┃ │
│        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                                                             ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope · z fold pane
```

Layout columns · truecolor · 160 × 44 terminal cells

```
Request req-82                                                                                                                                                  
┌ ▾ HTTP requests ·  ‖ PAUSED  ──────────┐ POST /v1/labels                                                      ┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────┐
│     GET       200  /v1/parcels         │ 502 Bad Gateway · 842ms                                              │ request req-82 · 3 matching logs             │
│ →   POST      502  /v1/labels          │ 2026-11-08T12:04:01Z                                                 │   ▾  Logs · 3 retained                       │
│     PATCH     204  /v1/parcels/p_42    │ Trace ID   trace-c7                                                  │  └── ▾  ⋈ Trace trace-c7                     │
│     GET        ◷   /health             │ ┌ ▾ Request ───────────────────────────────────────────────────────┐ │     └── ▾  ─ Span http-label                 │
│                                        │ │ ▾ Parameters                                                   ┃ │ │        ├── ▾  ─ Span carrier                 │
│                                        │ │ Query                                                          ┃ │ │ ┌ Log · log-3 ────────────────────────────┐  │
│                                        │ │ expand: "rate"                                                 ┃ │ │ │ ≡ WARN  Carrier response delayed        │  │
│                                        │ │ expand: "tracking"                                             ┃ │ │ │ ├──time: 2026-11-08T12:04:01.2Z         │  │
│                                        │ │ locale: "en_US"                                                ┃ │ │ │ └──upstream {2}                         │  │
│                                        │ │                                                                ┃ │ │ │    ├──host: "carrier.example"           │  │
│                                        │ │ Content-Type: application/json                                 ┃ │ │ │    └──timeout_ms: 800                   │  │
│                                        │ │ X-Request-Id: req-82                                           ┃ │ │ └─────────────────────────────────────────┘  │
│                                        │ │   ▾ {                                                          │ │ │        │  └── ▸  ≡ ERROR Upstream timeout    │
│                                        │ │       "parcel_id": "p_42",                                     │ │ │        └── ▸  ≡ INFO  Create label           │
│                                        │ └──────────────────────────────────────────────────────────────────┘ │                                              │
│                                        │ ┌ ▾ Response ──────────────────────────────────────────────────────┐ │                                              │
│                                        │ │ Content-Type: application/problem+json; charset=utf-8            │ │                                              │
│                                        │ │ Retry-After: 1                                                   │ │                                              │
│                                        │ │   ▾ {                                                            │ │                                              │
│                                        │ │       "error": "upstream_timeout",                               │ │                                              │
│                                        │ │       "retryable": true,                                         │ │                                              │
│                                        │ │       "attempt": 2                                               │ │                                              │
│                                        │ │     }                                                            │ │                                              │
│                                        │ │                                                                  │ │                                              │
│                                        │ │                                                                  │ │                                              │
│                                        │ │                                                                  │ │                                              │
│                                        │ │                                                                  │ │                                              │
│                                        │ │                                                                  │ │                                              │
│                                        │ │                                                                  │ │                                              │
│                                        │ │                                                                  │ │                                              │
│                                        │ │                                                                  │ │                                              │
│                                        │ │                                                                  │ │                                              │
│                                        │ │                                                                  │ │                                              │
│                                        │ │                                                                  │ │                                              │
│                                        │ │                                                                  │ │                                              │
│                                        │ │                                                                  │ │                                              │
│                                        │ │                                                                  │ │                                              │
│                                        │ │                                                                  │ │                                              │
│                                        │ │                                                                  │ │                                              │
└────────────────────────────────────────┘ └──────────────────────────────────────────────────────────────────┘ └──────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                                                             
↑/k/↓/j select · F follow · t scope · z fold pane
```

Layout focused · truecolor · 60 × 30 terminal cells

```
Request req-82                                              
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────┐
│ request req-82 · 3 matching logs                         │
│   ▾  Logs · 3 retained                                   │
│  └── ▾  ⋈ Trace trace-c7                                 │
│     └── ▾  ─ Span http-label                             │
│        ├── ▾  ─ Span carrier                             │
│ ┌ Log · log-3 ────────────────────────────────────────┐  │
│ │ ≡ WARN  Carrier response delayed                    │  │
│ │ ├──time: 2026-11-08T12:04:01.2Z                     │  │
│ │ └──upstream {2}                                     │  │
│ │    ├──host: "carrier.example"                       │  │
│ │    └──timeout_ms: 800                               │  │
│ └─────────────────────────────────────────────────────┘  │
│        │  └── ▸  ≡ ERROR Upstream timeout                │
│        └── ▸  ≡ INFO  Create label                       │
│                                                          │
│                                                          │
│                                                          │
│                                                          │
│                                                          │
│                                                          │
│                                                          │
│                                                          │
│                                                          │
│                                                          │
│                                                          │
└──────────────────────────────────────────────────────────┘
Tab panes · ? help · Ctrl+K commands                        
D details · g group · F follow · z fold pane
```

Focused narrow · truecolor · 40 × 24 terminal cells

```
Request req-82                          
┌ ▾ Related logs ·  ‖ PAUSED  ─────────┐
│ request req-82 · 3 matching logs     │
│   ▾  Logs · 3 retained               │
│  └── ▾  ⋈ Trace trace-c7             │
│     └── ▾  ─ Span http-label         │
│        ├── ▾  ─ Span carrier         │
│ ┌ Log · log-3 ────────────────────┐  │
│ │ ≡ WARN  Carrier response        │  │
│ │         delayed                 │  │
│ │ ├──time: 2026-11-               │  │
│ │ │        08T12:04:01.2Z         │  │
│ │ └──upstream {2}                 │  │
│ │    ├──host: "carrier.example"   │  │
│ │    └──timeout_ms: 800           │  │
│ └─────────────────────────────────┘  │
│        │  └── ▸  ≡ ERROR Upstream…   │
│        └── ▸  ≡ INFO  Create label   │
│                                      │
│                                      │
│                                      │
└──────────────────────────────────────┘
Tab panes · ? help · Ctrl+K commands    
D details · g group · F follow · z fold 
```

Narrow details · truecolor · 40 × 24 terminal cells

```
Request req-82                          
POST /v1/labels                         
502 Bad Gateway · 842ms                 
2026-11-08T12:04:01Z                    
Trace ID   trace-c7                     
┌ ▾ Request ───────────────────────────┐
│ ▾ Parameters                       ┃ │
│ Query                              │ │
│ expand: "rate"                     │ │
└──────────────────────────────────────┘
┌ ▾ Response ──────────────────────────┐
│ Content-Type: application/problem+j  │
│               son; charset=utf-8     │
│ Retry-After: 1                       │
│ → ▾ {                                │
│       "error": "upstream_timeout",   │
│       "retryable": true,             │
│       "attempt": 2                   │
│     }                                │
│                                      │
│                                      │
└──────────────────────────────────────┘
Tab panes · ? help · Ctrl+K commands    
↑/k/↓/j fields · Enter fold JSON · r raw
```

Narrow logs · truecolor · 40 × 24 terminal cells

```
Request req-82                          
┌ ▾ Related logs ·  ‖ PAUSED  ─────────┐
│ request req-82 · 3 matching logs     │
│    ▾  Logs · 3 retained              │
│  └── ▾  ⋈ Trace trace-c7             │
│     └── ▾  ─ Span http-label         │
│        ├── ▾  ─ Span carrier         │
│        │  ├── ▸  ≡ WARN  Carrier     │
│        │  │              response    │
│        │  │              delayed     │
│        │  └── ▸  ≡ ERROR Upstream    │
│        │                 timeout     │
│ →      └── ▾  ≡ INFO  Create label   │
│           ├── ▾  time: 2026-11-      │
│           │            08T12:04:01   │
│           │            Z             │
│           ├── ▸  parcel {2}          │
│           └── ▾  attempt: 2          │
│                                      │
│                                      │
│                                      │
└──────────────────────────────────────┘
Tab panes · ? help · Ctrl+K commands    
Enter expand · g group · F follow · z fo
```

Trace context · truecolor · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌ ▾ Request timeline ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   GET /v1/parcels                       ┤              ┊              ┊              ┊              ┊   200     12ms │
│ > POST /v1/labels                       ┊              ├━━━━━━━━━━━┤  ┊              ┊              ┊   502    842ms │
│   PATCH /v1/parcels/p_42                ┊              ┊              ┤              ┊              ┊   204      8ms │
│   GET /health                           ┊              ┊              ┊              ├━━━━━━━━━━━━━━▷   ◷         1s │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
Trace ID   trace-c7                                                                                                     
┌ ▾ Request ───────────────────────────────┐ ┌ ▾ Response ─────────────────────────────────────────────────────────────┐
│ ▾ Parameters                             │ │ Content-Type: application/problem+json; charset=utf-8                   │
│ Query                                    │ │ Retry-After: 1                                                          │
│ expand: "rate"                           │ │   ▾ {                                                                   │
│ expand: "tracking"                       │ │       "error": "upstream_timeout",                                      │
│ locale: "en_US"                          │ │       "retryable": true,                                                │
│                                          │ │       "attempt": 2                                                      │
│ Content-Type: application/json           │ │     }                                                                   │
│ X-Request-Id: req-82                     │ │                                                                         │
│   ▾ {                                    │ │                                                                         │
│       "parcel_id": "p_42",               │ │                                                                         │
│       "service": "express"               │ │                                                                         │
│     }                                    │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ ▾ Related logs ·  ‖ PAUSED  ─────────────────────────────────────────────────────────────────────────────────────────┐
│ trace trace-c7 · 4 matching logs                                                                                     │
│        │  └── ▸  12:04:01.842 ≡ ERROR Upstream timeout                                                             │ │
│ →      └── ▾  12:04:01.000 ≡ INFO  Create label                                                                    │ │
│           ├── ▾  time: 2026-11-08T12:04:01Z                                                                        │ │
│           ├── ▸  parcel {2}                                                                                        ┃ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope · z fold pane
```

No folds · truecolor · 120 × 38 terminal cells

```
Request req-82                                                                                                          
┌ Request timeline ·  ‖ PAUSED  ───────────────────────────────────────────────────────────────────────────────────────┐
│ 12:04:00.000Z                           0              +1s            +2s            +3s          +4s                │
│   GET /v1/parcels                       ┤              ┊              ┊              ┊              ┊   200     12ms │
│ > POST /v1/labels                       ┊              ├━━━━━━━━━━━┤  ┊              ┊              ┊   502    842ms │
│   PATCH /v1/parcels/p_42                ┊              ┊              ┤              ┊              ┊   204      8ms │
│   GET /health                           ┊              ┊              ┊              ├━━━━━━━━━━━━━━▷   ◷         1s │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
POST /v1/labels                                                                                                         
502 Bad Gateway · 842ms                                                                                                 
2026-11-08T12:04:01Z                                                                                                    
Trace ID   trace-c7                                                                                                     
┌ Request ─────────────────────────────────┐ ┌ Response ───────────────────────────────────────────────────────────────┐
│ Parameters                               │ │ Content-Type: application/problem+json; charset=utf-8                   │
│ Query                                    │ │ Retry-After: 1                                                          │
│ expand: "rate"                           │ │ {                                                                       │
│ expand: "tracking"                       │ │   "error": "upstream_timeout",                                          │
│ locale: "en_US"                          │ │   "retryable": true,                                                    │
│                                          │ │   "attempt": 2                                                          │
│ Content-Type: application/json           │ │ }                                                                       │
│ X-Request-Id: req-82                     │ │                                                                         │
│ {                                        │ │                                                                         │
│   "parcel_id": "p_42",                   │ │                                                                         │
│   "service": "express"                   │ │                                                                         │
│ }                                        │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
│                                          │ │                                                                         │
└──────────────────────────────────────────┘ └─────────────────────────────────────────────────────────────────────────┘
┌ Related logs ·  ‖ PAUSED  ───────────────────────────────────────────────────────────────────────────────────────────┐
│ request req-82 · 3 matching logs                                                                                     │
│        │        └──  [1]: "retry exhausted"                                                                        │ │
│ →      └──  12:04:01.000 ≡ INFO  Create label                                                                      │ │
│           ├──  time: 2026-11-08T12:04:01Z                                                                          ┃ │
│           ├──  parcel {2}                                                                                          │ │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Tab panes · L layout · Alt+1/2/3/4 views · Ctrl+K commands · ? help                                                     
↑/k/↓/j select · F follow · t scope
```

### [Go usage](https://lanej.io/charmed/organisms-logs.html#pattern%2flog-inspector-code)

Imports: [`time`](https://pkg.go.dev/time "time") · [`log`](https://pkg.go.dev/charm.land/log/v2 "charm.land/log/v2") · [`collection`](https://pkg.go.dev/github.com/lanej/charmed/collection "github.com/lanej/charmed/collection") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`pattern`](https://pkg.go.dev/github.com/lanej/charmed/pattern "github.com/lanej/charmed/pattern") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
at := time.Date(2026, 11, 8, 12, 4, 0, 0, time.UTC)
records := []collection.LogRecord{{
	ID: "log-1", RequestID: "req-82", TraceID: "trace-c7", SpanID: "carrier",
	LogEntry: collection.LogEntry{
		Time: at, Level: log.WarnLevel,
		Message: "Carrier response delayed", Fields: []any{"timeout_ms", 800},
	},
}}

requests := []collection.HTTPExchange{{
	ID: "http-2", RequestID: "req-82", TraceID: "trace-c7", Time: at,
	Method: "POST", URL: "https://api.example.test/v1/labels",
	StatusCode: 502, Completed: true, Duration: 842 * time.Millisecond, DurationKnown: true,
	ResponseBody: collection.HTTPBody{
		State: collection.Captured, ContentType: "application/json",
		Bytes: []byte(`{"error":"upstream_timeout","retryable":true}`),
	},
}}

inspector, err := pattern.NewLogInspector(pattern.LogInspectorOptions{
	Theme: theme.New(theme.Dark), Requests: requests, Logs: records,
	NoHelp: true, DetachedLogDetails: true, // The default layout adapts to available space.
})
if err != nil {
	return "", err
}

inspector.Resize(layout.Size{Width: 100, Height: 30})
inspector.SelectRequest("http-2")
// Optional: inspector.SetPaneCollapsed("request", true) gives the response both columns.
// inspector.SetHelpVisible(true) restores keyboard hints; FocusPane("response") changes focus.
// Customize payload styles on inspector.Request and inspector.Response.

return inspector.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-logs.html#pattern%2flog-inspector-integration)

Return inspector.Init from the root model and forward messages through inspector.Update. Submit observations on that same update loop. Call Resize when the terminal changes size; the automatic layout moves logs below the independent payloads before stacking them. FocusPane changes the active region, and native viewport styles survive reflow.

### [Run this example](https://lanej.io/charmed/organisms-logs.html#pattern%2flog-inspector-run)

```sh
go run ./examples/catalog pattern/log-inspector
```

[Example source](https://github.com/lanej/charmed/blob/ebf9018696944bf7fff7b999487964b17085a58d/docs/snippets/pattern_log_inspector.go) · [Implementation](https://github.com/lanej/charmed/blob/ebf9018696944bf7fff7b999487964b17085a58d/story/logs.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [collection/http-log](https://lanej.io/charmed/molecules-logs.html#collection/http-log) [collection/json-viewer](https://lanej.io/charmed/molecules-data.html#collection/json-viewer) [collection/log-tree](https://lanej.io/charmed/molecules-logs.html#collection/log-tree) [collection/command-palette](https://lanej.io/charmed/molecules-commands.html#collection/command-palette) [feedback/follow](https://lanej.io/charmed/atoms-feedback.html#feedback/follow) [text/presentation](https://lanej.io/charmed/atoms-text.html#text/presentation) [text/key-hint](https://lanej.io/charmed/atoms-text.html#text/key-hint) [layout/section](https://lanej.io/charmed/molecules-layout.html#layout/section) [layout/identity-header](https://lanej.io/charmed/molecules-layout.html#layout/identity-header) [foundations/layout](https://lanej.io/charmed/foundations-layout.html#foundations/layout)
