[Molecules](https://lanej.io/charmed/molecules.html)Commands

# [Commands](https://lanej.io/charmed/molecules-commands.html#level-molecules-commands)

2 components · [All molecules](https://lanej.io/charmed/molecules.html)

- [Command palette](https://lanej.io/charmed/molecules-commands.html#collection%2fcommand-palette)
- [CLI help](https://lanej.io/charmed/molecules-commands.html#cli%2fhelp)

[Molecules](https://lanej.io/charmed/molecules.html) / Commands

## [Command palette](https://lanej.io/charmed/molecules-commands.html#collection%2fcommand-palette)

Fuzzy-searchable commands with shortcuts, current settings and unavailable reasons; applications own execution.

### [Native preview](https://lanej.io/charmed/molecules-commands.html#collection%2fcommand-palette-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
┌ Commands ────────────────────────────────────────────────────┐
│ > Search commands…                                           │
│                                                              │
│ → Refresh observations                                     r │
│                                                              │
│   Scope: trace                                   unavailable │
│   No trace selected                                          │
│   Layout: columns                                    columns │
│                                                              │
│ 1/3 · 3 matches                                              │
│ ↑/↓ select · Enter run · Esc cancel                          │
└──────────────────────────────────────────────────────────────┘
```

Demo · truecolor · 88 × 28 terminal cells

```
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Applicat┌ Commands ────────────────────────────────────────────────────────────┐        
Applicat│ > Search commands…                                                   │        
Applicat│                                                                      │        
Applicat│ → Layout: timeline                                           current │        
Applicat│   Requests above details and related logs                            │        
Applicat│   Layout: columns                                                    │        
Applicat│   Requests, details and logs side by side                            │        
Applicat│   Focus: HTTP details                                          Alt+2 │        
Applicat│   Inspect parameters, headers and response bodies                    │        
Applicat│   Follow requests                                                  F │        
Applicat│   Keep arriving observations in view                                 │        
Applicat│   Scope: trace                                           unavailable │        
Applicat│   Selected request has no trace ID                                   │        
Applicat│   Show raw body                                                    r │        
Applicat│   Inspect captured bytes without formatting                          │        
Applicat│   Expand parameters                                                3 │        
Applicat│   Show decoded query and form values                                 │        
Applicat│   URL: full                                                          │        
Applicat│   Display the original captured URL                                  │        
Applicat│ 1/9 · 9 matches                                                      │        
Applicat│ ↑/↓ select · Enter run · Esc cancel                                  │        
Applicat└──────────────────────────────────────────────────────────────────────┘        
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
```

Narrow · truecolor · 40 × 16 terminal cells

```
Application content remains visible bene
┌ Commands ────────────────────────────┐
│ > Search commands…                   │
│                                      │
│ → Layout: timeline           current │
│   Layout: columns                    │
│   Focus: HTTP details          Alt+2 │
│   Follow requests                  F │
│   Scope: trace           unavailable │
│   Show raw body                    r │
│   Expand parameters                3 │
│   URL: full                          │
│ 1/9 · 9 matches                      │
│ ↑/↓ select · Enter run · Esc cancel  │
└──────────────────────────────────────┘
Application content remains visible bene
```

256 colors · ansi256 · 88 × 28 terminal cells

```
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Applicat┌ Commands ────────────────────────────────────────────────────────────┐        
Applicat│ > Search commands…                                                   │        
Applicat│                                                                      │        
Applicat│ → Layout: timeline                                           current │        
Applicat│   Requests above details and related logs                            │        
Applicat│   Layout: columns                                                    │        
Applicat│   Requests, details and logs side by side                            │        
Applicat│   Focus: HTTP details                                          Alt+2 │        
Applicat│   Inspect parameters, headers and response bodies                    │        
Applicat│   Follow requests                                                  F │        
Applicat│   Keep arriving observations in view                                 │        
Applicat│   Scope: trace                                           unavailable │        
Applicat│   Selected request has no trace ID                                   │        
Applicat│   Show raw body                                                    r │        
Applicat│   Inspect captured bytes without formatting                          │        
Applicat│   Expand parameters                                                3 │        
Applicat│   Show decoded query and form values                                 │        
Applicat│   URL: full                                                          │        
Applicat│   Display the original captured URL                                  │        
Applicat│ 1/9 · 9 matches                                                      │        
Applicat│ ↑/↓ select · Enter run · Esc cancel                                  │        
Applicat└──────────────────────────────────────────────────────────────────────┘        
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
```

16 colors · ansi16 · 88 × 28 terminal cells

```
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Applicat┌ Commands ────────────────────────────────────────────────────────────┐        
Applicat│ > Search commands…                                                   │        
Applicat│                                                                      │        
Applicat│ → Layout: timeline                                           current │        
Applicat│   Requests above details and related logs                            │        
Applicat│   Layout: columns                                                    │        
Applicat│   Requests, details and logs side by side                            │        
Applicat│   Focus: HTTP details                                          Alt+2 │        
Applicat│   Inspect parameters, headers and response bodies                    │        
Applicat│   Follow requests                                                  F │        
Applicat│   Keep arriving observations in view                                 │        
Applicat│   Scope: trace                                           unavailable │        
Applicat│   Selected request has no trace ID                                   │        
Applicat│   Show raw body                                                    r │        
Applicat│   Inspect captured bytes without formatting                          │        
Applicat│   Expand parameters                                                3 │        
Applicat│   Show decoded query and form values                                 │        
Applicat│   URL: full                                                          │        
Applicat│   Display the original captured URL                                  │        
Applicat│ 1/9 · 9 matches                                                      │        
Applicat│ ↑/↓ select · Enter run · Esc cancel                                  │        
Applicat└──────────────────────────────────────────────────────────────────────┘        
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
```

Monochrome · none · 88 × 28 terminal cells

```
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Applicat+ Commands ------------------------------------------------------------+        
Applicat| > Search commands.                                                   |        
Applicat|                                                                      |        
Applicat| > Layout: timeline                                           current |        
Applicat|   Requests above details and related logs                            |        
Applicat|   Layout: columns                                                    |        
Applicat|   Requests, details and logs side by side                            |        
Applicat|   Focus: HTTP details                                          Alt+2 |        
Applicat|   Inspect parameters, headers and response bodies                    |        
Applicat|   Follow requests                                                  F |        
Applicat|   Keep arriving observations in view                                 |        
Applicat|   Scope: trace                                           unavailable |        
Applicat|   Selected request has no trace ID                                   |        
Applicat|   Show raw body                                                    r |        
Applicat|   Inspect captured bytes without formatting                          |        
Applicat|   Expand parameters                                                3 |        
Applicat|   Show decoded query and form values                                 |        
Applicat|   URL: full                                                          |        
Applicat|   Display the original captured URL                                  |        
Applicat| 1/9 * 9 matches                                                      |        
Applicat| ^/v select * Enter run * Esc cancel                                  |        
Applicat+----------------------------------------------------------------------+        
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
```

Backdrop dim · truecolor · 88 × 28 terminal cells

```
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Applicat┌ Commands ────────────────────────────────────────────────────────────┐        
Applicat│ > laytm                                                              │        
Applicat│                                                                      │        
Applicat│ → Layout: timeline                                           current │        
Applicat│   Requests above details and related logs                            │        
Applicat│   Layout: columns                                                    │        
Applicat│   Requests, details and logs side by side                            │        
Applicat│ 1/2 · 2 matches                                                      │        
Applicat│ ↑/↓ select · Enter run · Esc cancel                                  │        
Applicat└──────────────────────────────────────────────────────────────────────┘        
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
```

Backdrop monochrome · truecolor · 88 × 28 terminal cells

```
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Applicat┌ Commands ────────────────────────────────────────────────────────────┐        
Applicat│ > laytm                                                              │        
Applicat│                                                                      │        
Applicat│ → Layout: timeline                                           current │        
Applicat│   Requests above details and related logs                            │        
Applicat│   Layout: columns                                                    │        
Applicat│   Requests, details and logs side by side                            │        
Applicat│ 1/2 · 2 matches                                                      │        
Applicat│ ↑/↓ select · Enter run · Esc cancel                                  │        
Applicat└──────────────────────────────────────────────────────────────────────┘        
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
```

Backdrop none · truecolor · 88 × 28 terminal cells

```
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Applicat┌ Commands ────────────────────────────────────────────────────────────┐        
Applicat│ > laytm                                                              │        
Applicat│                                                                      │        
Applicat│ → Layout: timeline                                           current │        
Applicat│   Requests above details and related logs                            │        
Applicat│   Layout: columns                                                    │        
Applicat│   Requests, details and logs side by side                            │        
Applicat│ 1/2 · 2 matches                                                      │        
Applicat│ ↑/↓ select · Enter run · Esc cancel                                  │        
Applicat└──────────────────────────────────────────────────────────────────────┘        
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette
```

Vim search · truecolor · 88 × 28 terminal cells

```
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Applicat┌ Commands ────────────────────────────────────────────────────────────┐        
Applicat│ > laytm                                                              │        
Applicat│                                                                      │        
Applicat│ → Layout: timeline                                           current │        
Applicat│   Requests above details and related logs                            │        
Applicat│   Layout: columns                                                    │        
Applicat│   Requests, details and logs side by side                            │        
Applicat│ SEARCH · 1/2 · 2 matches                                             │        
Applicat│ ↑↓ select · Esc normal · Enter run                                   │        
Applicat└──────────────────────────────────────────────────────────────────────┘        
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
```

Vim normal · truecolor · 88 × 28 terminal cells

```
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Applicat┌ Commands ────────────────────────────────────────────────────────────┐        
Applicat│ > laytm                                                              │        
Applicat│                                                                      │        
Applicat│ → Layout: timeline                                           current │        
Applicat│   Requests above details and related logs                            │        
Applicat│   Layout: columns                                                    │        
Applicat│   Requests, details and logs side by side                            │        
Applicat│ NORMAL · 1/2 · 2 matches                                             │        
Applicat│ k/↑/j/↓ move · i search · Esc close                                  │        
Applicat└──────────────────────────────────────────────────────────────────────┘        
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
```

Vim narrow · truecolor · 40 × 20 terminal cells

```
Application content remains visible bene
Application content remains visible bene
Application content remains visible bene
┌ Commands ────────────────────────────┐
│ > Search commands…                   │
│                                      │
│ → Layout: timeline           current │
│   Layout: columns                    │
│   Focus: HTTP details          Alt+2 │
│   Follow requests                  F │
│   Scope: trace           unavailable │
│   Show raw body                    r │
│   Expand parameters                3 │
│   URL: full                          │
│ NORMAL · 1/9 · 9 matches             │
│ k/↑/j/↓ move · i search · Esc close  │
└──────────────────────────────────────┘
Application content remains visible bene
Application content remains visible bene
Application content remains visible bene
```

Commands · truecolor · 88 × 28 terminal cells

```
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Applicat┌ Commands ────────────────────────────────────────────────────────────┐        
Applicat│ > Search commands…                                                   │        
Applicat│                                                                      │        
Applicat│ → Layout: timeline                                           current │        
Applicat│   Requests above details and related logs                            │        
Applicat│   Layout: columns                                                    │        
Applicat│   Requests, details and logs side by side                            │        
Applicat│   Focus: HTTP details                                          Alt+2 │        
Applicat│   Inspect parameters, headers and response bodies                    │        
Applicat│   Follow requests                                                  F │        
Applicat│   Keep arriving observations in view                                 │        
Applicat│   Scope: trace                                           unavailable │        
Applicat│   Selected request has no trace ID                                   │        
Applicat│   Show raw body                                                    r │        
Applicat│   Inspect captured bytes without formatting                          │        
Applicat│   Expand parameters                                                3 │        
Applicat│   Show decoded query and form values                                 │        
Applicat│   URL: full                                                          │        
Applicat│   Display the original captured URL                                  │        
Applicat│ 1/9 · 9 matches                                                      │        
Applicat│ ↑/↓ select · Enter run · Esc cancel                                  │        
Applicat└──────────────────────────────────────────────────────────────────────┘        
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
```

Commands filtered · truecolor · 88 × 28 terminal cells

```
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Applicat┌ Commands ────────────────────────────────────────────────────────────┐        
Applicat│ > laytm                                                              │        
Applicat│                                                                      │        
Applicat│ → Layout: timeline                                           current │        
Applicat│   Requests above details and related logs                            │        
Applicat│   Layout: columns                                                    │        
Applicat│   Requests, details and logs side by side                            │        
Applicat│ 1/2 · 2 matches                                                      │        
Applicat│ ↑/↓ select · Enter run · Esc cancel                                  │        
Applicat└──────────────────────────────────────────────────────────────────────┘        
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
```

Commands highlighted · truecolor · 88 × 28 terminal cells

```
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Applicat┌ Commands ────────────────────────────────────────────────────────────┐        
Applicat│ > laytm                                                              │        
Applicat│                                                                      │        
Applicat│ → Layout: timeline                                           current │        
Applicat│   Requests above details and related logs                            │        
Applicat│   Layout: columns                                                    │        
Applicat│   Requests, details and logs side by side                            │        
Applicat│ 1/2 · 2 matches                                                      │        
Applicat│ ↑/↓ select · Enter run · Esc cancel                                  │        
Applicat└──────────────────────────────────────────────────────────────────────┘        
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
```

Commands unavailable · truecolor · 88 × 28 terminal cells

```
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Applicat┌ Commands ────────────────────────────────────────────────────────────┐        
Applicat│ > Scope: trace                                                       │        
Applicat│                                                                      │        
Applicat│ → Scope: trace                                           unavailable │        
Applicat│   Selected request has no trace ID                                   │        
Applicat│ 1/1 · 1 matches                                                      │        
Applicat│ ↑/↓ select · Enter run · Esc cancel                                  │        
Applicat└──────────────────────────────────────────────────────────────────────┘        
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
```

Commands empty · truecolor · 88 × 28 terminal cells

```
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Applicat┌ Commands ────────────────────────────────────────────────────────────┐        
Applicat│ > zzzzzz                                                             │        
Applicat│                                                                      │        
Applicat│ No matching commands                                                 │        
Applicat│                                                                      │        
Applicat│ 0/0 · 0 matches                                                      │        
Applicat│ ↑/↓ select · Enter run · Esc cancel                                  │        
Applicat└──────────────────────────────────────────────────────────────────────┘        
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
```

Commands narrow · truecolor · 40 × 20 terminal cells

```
Application content remains visible bene
Application content remains visible bene
Application content remains visible bene
┌ Commands ────────────────────────────┐
│ > Search commands…                   │
│                                      │
│ → Layout: timeline           current │
│   Layout: columns                    │
│   Focus: HTTP details          Alt+2 │
│   Follow requests                  F │
│   Scope: trace           unavailable │
│   Show raw body                    r │
│   Expand parameters                3 │
│   URL: full                          │
│ 1/9 · 9 matches                      │
│ ↑/↓ select · Enter run · Esc cancel  │
└──────────────────────────────────────┘
Application content remains visible bene
Application content remains visible bene
Application content remains visible bene
```

light · truecolor · 88 × 28 terminal cells

```
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Applicat┌ Commands ────────────────────────────────────────────────────────────┐        
Applicat│ > Search commands…                                                   │        
Applicat│                                                                      │        
Applicat│ → Layout: timeline                                           current │        
Applicat│   Requests above details and related logs                            │        
Applicat│   Layout: columns                                                    │        
Applicat│   Requests, details and logs side by side                            │        
Applicat│   Focus: HTTP details                                          Alt+2 │        
Applicat│   Inspect parameters, headers and response bodies                    │        
Applicat│   Follow requests                                                  F │        
Applicat│   Keep arriving observations in view                                 │        
Applicat│   Scope: trace                                           unavailable │        
Applicat│   Selected request has no trace ID                                   │        
Applicat│   Show raw body                                                    r │        
Applicat│   Inspect captured bytes without formatting                          │        
Applicat│   Expand parameters                                                3 │        
Applicat│   Show decoded query and form values                                 │        
Applicat│   URL: full                                                          │        
Applicat│   Display the original captured URL                                  │        
Applicat│ 1/9 · 9 matches                                                      │        
Applicat│ ↑/↓ select · Enter run · Esc cancel                                  │        
Applicat└──────────────────────────────────────────────────────────────────────┘        
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
```

contrast · truecolor · 88 × 28 terminal cells

```
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Applicat┌ Commands ────────────────────────────────────────────────────────────┐        
Applicat│ > Search commands…                                                   │        
Applicat│                                                                      │        
Applicat│ → Layout: timeline                                           current │        
Applicat│   Requests above details and related logs                            │        
Applicat│   Layout: columns                                                    │        
Applicat│   Requests, details and logs side by side                            │        
Applicat│   Focus: HTTP details                                          Alt+2 │        
Applicat│   Inspect parameters, headers and response bodies                    │        
Applicat│   Follow requests                                                  F │        
Applicat│   Keep arriving observations in view                                 │        
Applicat│   Scope: trace                                           unavailable │        
Applicat│   Selected request has no trace ID                                   │        
Applicat│   Show raw body                                                    r │        
Applicat│   Inspect captured bytes without formatting                          │        
Applicat│   Expand parameters                                                3 │        
Applicat│   Show decoded query and form values                                 │        
Applicat│   URL: full                                                          │        
Applicat│   Display the original captured URL                                  │        
Applicat│ 1/9 · 9 matches                                                      │        
Applicat│ ↑/↓ select · Enter run · Esc cancel                                  │        
Applicat└──────────────────────────────────────────────────────────────────────┘        
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
Application content remains visible beneath the palette                                 
```

### [Go usage](https://lanej.io/charmed/molecules-commands.html#collection%2fcommand-palette-code)

Imports: [`collection`](https://pkg.go.dev/github.com/lanej/charmed/collection "github.com/lanej/charmed/collection") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
commands, err := collection.NewCommandPalette([]collection.Command{
	{ID: "refresh", Label: "Refresh observations", Shortcut: "r"},
	{ID: "trace", Label: "Scope: trace", DisabledReason: "No trace selected"},
	{ID: "layout", Label: "Layout: columns", State: "columns"},
}, 64, 12, theme.New(theme.Dark))
if err != nil {
	return "", err
}

commands.Open() // Return this command when opening the palette in your application.
// Forward commands.Update(msg); handle collection.CommandSelectedMsg.ID to run an action.
return commands.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-commands.html#collection%2fcommand-palette-integration)

Return Open's focus command, then route keyboard and paste messages exclusively to the open palette. Handle CommandSelectedMsg in the application and keep processing observations in the background. Resize preserves native styles and key maps. Restore the previous pane's focus when the palette closes without an action.

### [Run this example](https://lanej.io/charmed/molecules-commands.html#collection%2fcommand-palette-run)

```sh
go run ./examples/catalog collection/command-palette
```

[Example source](https://github.com/lanej/charmed/blob/49f755da5bd25a044043cf98bc303185c2ebce6d/docs/snippets/collection_command_palette.go) · [Implementation](https://github.com/lanej/charmed/blob/49f755da5bd25a044043cf98bc303185c2ebce6d/story/commands.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [input/text](https://lanej.io/charmed/atoms-input.html#input/text) [text/key-hint](https://lanej.io/charmed/atoms-text.html#text/key-hint) [text/match](https://lanej.io/charmed/atoms-text.html#text/match) [foundations/layout](https://lanej.io/charmed/foundations-layout.html#foundations/layout)

[Molecules](https://lanej.io/charmed/molecules.html) / Commands

## [CLI help](https://lanej.io/charmed/molecules-commands.html#cli%2fhelp)

Render help from a Cobra command or install the formatter on its root. Flag names are bold and value types use normal weight. Descriptions move below long flags while shorter rows stay aligned.

### [Native preview](https://lanej.io/charmed/molecules-commands.html#cli%2fhelp-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Show the services in a release.

Usage
  release VERSION [flags]

Examples
  release v1.4 --format json

Flags
  -f, --format string  Output format (default "table")
  -h, --help           help for release
```

Demo · truecolor · 88 × 28 terminal cells

```
Show the services in a release.                                                         
                                                                                        
Usage                                                                                   
  release show VERSION [flags]                                                          
                                                                                        
Aliases                                                                                 
  show, inspect                                                                         
                                                                                        
Examples                                                                                
  release show v1.4 --format json                                                       
                                                                                        
Flags                                                                                   
  --changed            Show only services that changed in this release                  
  -f, --format string  Output format (default "table")                                  
  -h, --help           help for show                                                    
                                                                                        
Global flags                                                                            
  -p, --profile string  Configuration profile to use (default "local")
```

Narrow · truecolor · 40 × 28 terminal cells

```
Show the services in a release.         
                                        
Usage                                   
  release show VERSION [flags]          
                                        
Aliases                                 
  show, inspect                         
                                        
Examples                                
  release show v1.4 --format json       
                                        
Flags                                   
  --changed   Show only services that   
              changed in this release   
  -f, --format string                   
    Output format (default "table")     
  -h, --help  help for show             
                                        
Global flags                            
  -p, --profile string                  
    Configuration profile to use        
    (default "local")
```

256 colors · ansi256 · 88 × 28 terminal cells

```
Show the services in a release.                                                         
                                                                                        
Usage                                                                                   
  release show VERSION [flags]                                                          
                                                                                        
Aliases                                                                                 
  show, inspect                                                                         
                                                                                        
Examples                                                                                
  release show v1.4 --format json                                                       
                                                                                        
Flags                                                                                   
  --changed            Show only services that changed in this release                  
  -f, --format string  Output format (default "table")                                  
  -h, --help           help for show                                                    
                                                                                        
Global flags                                                                            
  -p, --profile string  Configuration profile to use (default "local")
```

16 colors · ansi16 · 88 × 28 terminal cells

```
Show the services in a release.                                                         
                                                                                        
Usage                                                                                   
  release show VERSION [flags]                                                          
                                                                                        
Aliases                                                                                 
  show, inspect                                                                         
                                                                                        
Examples                                                                                
  release show v1.4 --format json                                                       
                                                                                        
Flags                                                                                   
  --changed            Show only services that changed in this release                  
  -f, --format string  Output format (default "table")                                  
  -h, --help           help for show                                                    
                                                                                        
Global flags                                                                            
  -p, --profile string  Configuration profile to use (default "local")
```

Monochrome · none · 88 × 28 terminal cells

```
Show the services in a release.                                                         
                                                                                        
Usage                                                                                   
  release show VERSION [flags]                                                          
                                                                                        
Aliases                                                                                 
  show, inspect                                                                         
                                                                                        
Examples                                                                                
  release show v1.4 --format json                                                       
                                                                                        
Flags                                                                                   
  --changed            Show only services that changed in this release                  
  -f, --format string  Output format (default "table")                                  
  -h, --help           help for show                                                    
                                                                                        
Global flags                                                                            
  -p, --profile string  Configuration profile to use (default "local")
```

light · truecolor · 88 × 28 terminal cells

```
Show the services in a release.                                                         
                                                                                        
Usage                                                                                   
  release show VERSION [flags]                                                          
                                                                                        
Aliases                                                                                 
  show, inspect                                                                         
                                                                                        
Examples                                                                                
  release show v1.4 --format json                                                       
                                                                                        
Flags                                                                                   
  --changed            Show only services that changed in this release                  
  -f, --format string  Output format (default "table")                                  
  -h, --help           help for show                                                    
                                                                                        
Global flags                                                                            
  -p, --profile string  Configuration profile to use (default "local")
```

contrast · truecolor · 88 × 28 terminal cells

```
Show the services in a release.                                                         
                                                                                        
Usage                                                                                   
  release show VERSION [flags]                                                          
                                                                                        
Aliases                                                                                 
  show, inspect                                                                         
                                                                                        
Examples                                                                                
  release show v1.4 --format json                                                       
                                                                                        
Flags                                                                                   
  --changed            Show only services that changed in this release                  
  -f, --format string  Output format (default "table")                                  
  -h, --help           help for show                                                    
                                                                                        
Global flags                                                                            
  -p, --profile string  Configuration profile to use (default "local")
```

### [Go usage](https://lanej.io/charmed/molecules-commands.html#cli%2fhelp-code)

Imports: [`cobra`](https://pkg.go.dev/github.com/spf13/cobra "github.com/spf13/cobra") · [`cli`](https://pkg.go.dev/github.com/lanej/charmed/cli "github.com/lanej/charmed/cli") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
show := &cobra.Command{
	Use: "release VERSION", Short: "Show the services in a release.",
	Example: "release v1.4 --format json",
}
show.Flags().StringP("format", "f", "table", "Output format")
show.InitDefaultHelpFlag()

formatter := cli.HelpFormatter{Width: 56, Theme: theme.New(theme.Dark)}

return formatter.Format(show)
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-commands.html#cli%2fhelp-integration)

Install HelpFormatter on the Cobra root for inherited help, or call Format to render a command directly. The formatter reads the current command tree and flags on each request. Set Width for reproducible examples or let installed help detect its terminal.

### [Run this example](https://lanej.io/charmed/molecules-commands.html#cli%2fhelp-run)

```sh
go run ./examples/catalog cli/help
```

[Example source](https://github.com/lanej/charmed/blob/49f755da5bd25a044043cf98bc303185c2ebce6d/docs/snippets/cli_help.go) · [Implementation](https://github.com/lanej/charmed/blob/49f755da5bd25a044043cf98bc303185c2ebce6d/cli/help.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [foundations/theme](https://lanej.io/charmed/foundations-theme.html#foundations/theme)
