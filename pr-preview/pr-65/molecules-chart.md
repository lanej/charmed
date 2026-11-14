[Molecules](https://lanej.io/charmed/molecules.html)Charts

# [Charts](https://lanej.io/charmed/molecules-chart.html#level-molecules-chart)

6 components · [All molecules](https://lanej.io/charmed/molecules.html)

- [Waffle status header](https://lanej.io/charmed/molecules-chart.html#chart%2fwaffle-header)
- [Band matrix](https://lanej.io/charmed/molecules-chart.html#chart%2fband-matrix)
- [Metric trend](https://lanej.io/charmed/molecules-chart.html#chart%2fmetric)
- [Chart panel](https://lanej.io/charmed/molecules-chart.html#chart%2fpanel)
- [Cohort distributions](https://lanej.io/charmed/molecules-chart.html#chart%2fdistributions)
- [Merged delta property](https://lanej.io/charmed/molecules-chart.html#chart%2fdelta)

[Molecules](https://lanej.io/charmed/molecules.html) / Charts

## [Waffle status header](https://lanej.io/charmed/molecules-chart.html#chart%2fwaffle-header)

A fleet micro-grid above a service list with shared-scale instance bars and semantic status atoms. A one-row option supports tighter headers.

### [Native preview](https://lanej.io/charmed/molecules-chart.html#chart%2fwaffle-header-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
■ ■ ■ ■ ■ ■ ■ ■ ■ ■   PRODUCTION  /  100 total                  
■ ■ ■ ■ ■ ■ ■ ■ ▧ ▨   ■ ready 92 / ▧ pending 3 / ▨ failed 5
```

Demo · truecolor · 80 × 16 terminal cells

```
■ ■ ■ ■ ■ ■ ■ ■ ■ ■   PRODUCTION  /  100 total                                  
■ ■ ■ ■ ■ ■ ■ ■ ▧ ▨   ■ ready 92 / ▧ pending 3 / ▨ failed 5                     
                                                                                
SERVICE                                    INSTANCES                STATE       
checkout                                    24 ━━━━━━━━━━━━────────  ✓ OK       
labels                                      28 ━━━━━━━━━━━━━━──────  ✓ OK       
rating                                      40 ━━━━━━━━━━━━━━━━━━━━  ✓ OK       
tracking                                     3 ━━──────────────────  ◷ PENDING  
events                                       5 ━━━─────────────────  × ERROR    
                                                                                
Instance scale: 0–40
```

Narrow · truecolor · 40 × 16 terminal cells

```
PRODUCTION  /  100 total                
■ ready 92 / ▧ pending 3 / ▨ failed 5   
                                        
SERVICE       INSTANCES     STATE       
checkout       24 ━━━━━────  ✓ OK       
labels         28 ━━━━━━───  ✓ OK       
rating         40 ━━━━━━━━━  ✓ OK       
tracking        3 ━────────  ◷ PENDING  
events          5 ━────────  × ERROR    
                                        
Instance scale: 0–40
```

256 colors · ansi256 · 80 × 16 terminal cells

```
■ ■ ■ ■ ■ ■ ■ ■ ■ ■   PRODUCTION  /  100 total                                  
■ ■ ■ ■ ■ ■ ■ ■ ▧ ▨   ■ ready 92 / ▧ pending 3 / ▨ failed 5                     
                                                                                
SERVICE                                    INSTANCES                STATE       
checkout                                    24 ━━━━━━━━━━━━────────  ✓ OK       
labels                                      28 ━━━━━━━━━━━━━━──────  ✓ OK       
rating                                      40 ━━━━━━━━━━━━━━━━━━━━  ✓ OK       
tracking                                     3 ━━──────────────────  ◷ PENDING  
events                                       5 ━━━─────────────────  × ERROR    
                                                                                
Instance scale: 0–40
```

16 colors · ansi16 · 80 × 16 terminal cells

```
■ ■ ■ ■ ■ ■ ■ ■ ■ ■   PRODUCTION  /  100 total                                  
■ ■ ■ ■ ■ ■ ■ ■ ▧ ▨   ■ ready 92 / ▧ pending 3 / ▨ failed 5                     
                                                                                
SERVICE                                    INSTANCES                STATE       
checkout                                    24 ━━━━━━━━━━━━────────  ✓ OK       
labels                                      28 ━━━━━━━━━━━━━━──────  ✓ OK       
rating                                      40 ━━━━━━━━━━━━━━━━━━━━  ✓ OK       
tracking                                     3 ━━──────────────────  ◷ PENDING  
events                                       5 ━━━─────────────────  × ERROR    
                                                                                
Instance scale: 0–40
```

Monochrome · none · 80 × 16 terminal cells

```
A A A A A A A A A A   PRODUCTION  /  100 total                                  
A A A A A A A A B C   A ready 92 / B pending 3 / C failed 5                     
                                                                                
SERVICE                                    INSTANCES                STATE       
checkout                                    24 ############-------- [+ OK]      
labels                                      28 ##############------ [+ OK]      
rating                                      40 #################### [+ OK]      
tracking                                     3 ##------------------ [p PENDING] 
events                                       5 ###----------------- [x ERROR]   
                                                                                
Instance scale: 0-40
```

### [Go usage](https://lanej.io/charmed/molecules-chart.html#chart%2fwaffle-header-code)

Imports: [`chart`](https://pkg.go.dev/github.com/lanej/charmed/chart "github.com/lanej/charmed/chart") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
fleet := chart.Waffle{Parts: []chart.WafflePart{
	{Label: "ready", Count: 92, Tone: text.Success},
	{Label: "pending", Count: 3, Tone: text.Info},
	{Label: "failed", Count: 5, Tone: text.Danger},
}}

return fleet.StatusHeader("PRODUCTION", layout.Size{Width: 64, Height: 2}, palette)
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-chart.html#chart%2fwaffle-header-integration)

Keep observations and their units in the application, and pass changes through the chart's public data API. Allocate terminal columns and rows from the parent layout. Forward native model commands when enabling interaction, and check whether resizing rebuilds plotted data before applying native overrides.

### [Run this example](https://lanej.io/charmed/molecules-chart.html#chart%2fwaffle-header-run)

```sh
go run ./examples/catalog chart/waffle-header
```

[Example source](https://github.com/lanej/charmed/blob/682a17afe4ccff26744d2ac202a2bcda8d784cf8/docs/snippets/chart_waffle_header.go) · [Implementation](https://github.com/lanej/charmed/blob/682a17afe4ccff26744d2ac202a2bcda8d784cf8/chart/waffle.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [chart/waffle](https://lanej.io/charmed/atoms-chart.html#chart/waffle) [collection/table](https://lanej.io/charmed/atoms-collection.html#collection/table) [feedback/status](https://lanej.io/charmed/atoms-feedback.html#feedback/status) [feedback/progress](https://lanej.io/charmed/atoms-feedback.html#feedback/progress)

[Molecules](https://lanej.io/charmed/molecules.html) / Charts

## [Band matrix](https://lanej.io/charmed/molecules-chart.html#chart%2fband-matrix)

Directly labeled, disconnected distribution rows composed in a rectangular frame.

### [Native preview](https://lanej.io/charmed/molecules-chart.html#chart%2fband-matrix-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
┌ Host distributions ──────────────────────────────────────────┐
│ Success                                                      │
│ Reference    96.000%··░░░░▒▒▒▓▓▓▓|▓▓▓▒▒▒▒░░░················ │
│ Candidate    97.000%······░░░░░▒▒▒▒▒▓▓▓▓▓|▓▓▓▓▓▒▒▒▒▒░░░░░░·· │
│                     94.520%                          99.080% │
│                                                              │
│ Server error                                                 │
│ Reference     0.200%····░░░░▒▒▒▒▓▓▓▓|▓▓▓▓▒▒▒░░░░············ │
│ Candidate     0.230%··░░░░░░▒▒▒▒▒▒▓▓▓▓▓▓|▓▓▓▓▓▒▒▒▒▒▒░░░░░░·· │
│                     0.076%                            0.384% │
│                                                              │
└──────────────────────────────────────────────────────────────┘
```

Demo · truecolor · 100 × 30 terminal cells

```
┌ Host mean +/- 3s ────────────────────────────────────────────────────────────────────────────────┐
│ Success                                                                                          │
│ a13c9f2               96.000%·······░░░░░░░▒▒▒▒▒▒▓▓▓▓▓▓▓|▓▓▓▓▓▓▒▒▒▒▒▒░░░░░░░···················· │
│ f4d7e81               96.400%····░░░░░░░░░░▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓|▓▓▓▓▓▓▓▓▓▓▒▒▒▒▒▒▒▒▒░░░░░░X░░░···· │
│                              94.348%                                                     98.452% │
│                                                                                                  │
│ Server error                                                                                     │
│ a13c9f2                0.200%·······░░░░░░░▒▒▒▒▒▒▓▓▓▓▓▓▓|▓▓▓▓▓▓▒▒▒▒▒▒░░░░░░░···················· │
│ f4d7e81                0.230%····░░░░░░░░░░▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓|▓▓▓▓▓▓▓▓▓▓▒▒▒▒▒▒▒▒▒░░░░░░X░░░···· │
│                              0.076%                                                       0.384% │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
```

Narrow · truecolor · 40 × 16 terminal cells

```
┌ Host mean +/- 3s ────────────────────┐
│ Success                              │
│ a13~ 96.000%···░░▒▒▓▓|▓▓▓▒▒░░······· │
│ f4d~ 96.400%·░░░░▒▒▒▓▓▓▓|▓▓▓▒▒▒░X░░· │
│             94.348%          98.452% │
│                                      │
│ Server error                         │
│ a13c~ 0.200%···░░▒▒▓▓|▓▓▓▒▒░░······· │
│ f4d7~ 0.230%·░░░░▒▒▒▓▓▓▓|▓▓▓▒▒▒░X░░· │
│             0.076%            0.384% │
│                                      │
│                                      │
│                                      │
│                                      │
│                                      │
└──────────────────────────────────────┘
```

256 colors · ansi256 · 100 × 30 terminal cells

```
┌ Host mean +/- 3s ────────────────────────────────────────────────────────────────────────────────┐
│ Success                                                                                          │
│ a13c9f2               96.000%·······░░░░░░░▒▒▒▒▒▒▓▓▓▓▓▓▓|▓▓▓▓▓▓▒▒▒▒▒▒░░░░░░░···················· │
│ f4d7e81               96.400%····░░░░░░░░░░▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓|▓▓▓▓▓▓▓▓▓▓▒▒▒▒▒▒▒▒▒░░░░░░X░░░···· │
│                              94.348%                                                     98.452% │
│                                                                                                  │
│ Server error                                                                                     │
│ a13c9f2                0.200%·······░░░░░░░▒▒▒▒▒▒▓▓▓▓▓▓▓|▓▓▓▓▓▓▒▒▒▒▒▒░░░░░░░···················· │
│ f4d7e81                0.230%····░░░░░░░░░░▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓|▓▓▓▓▓▓▓▓▓▓▒▒▒▒▒▒▒▒▒░░░░░░X░░░···· │
│                              0.076%                                                       0.384% │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
```

16 colors · ansi16 · 100 × 30 terminal cells

```
┌ Host mean +/- 3s ────────────────────────────────────────────────────────────────────────────────┐
│ Success                                                                                          │
│ a13c9f2               96.000%·······░░░░░░░▒▒▒▒▒▒▓▓▓▓▓▓▓|▓▓▓▓▓▓▒▒▒▒▒▒░░░░░░░···················· │
│ f4d7e81               96.400%····░░░░░░░░░░▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓|▓▓▓▓▓▓▓▓▓▓▒▒▒▒▒▒▒▒▒░░░░░░X░░░···· │
│                              94.348%                                                     98.452% │
│                                                                                                  │
│ Server error                                                                                     │
│ a13c9f2                0.200%·······░░░░░░░▒▒▒▒▒▒▓▓▓▓▓▓▓|▓▓▓▓▓▓▒▒▒▒▒▒░░░░░░░···················· │
│ f4d7e81                0.230%····░░░░░░░░░░▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓|▓▓▓▓▓▓▓▓▓▓▒▒▒▒▒▒▒▒▒░░░░░░X░░░···· │
│                              0.076%                                                       0.384% │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
│                                                                                                  │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
```

Monochrome · none · 100 × 30 terminal cells

```
+ Host mean +/- 3s --------------------------------------------------------------------------------+
| Success                                                                                          |
| a13c9f2               96.000%..............======#######|######======........................... |
| f4d7e81               96.400%..............=========##########|##########=========......X....... |
|                              94.348%                                                     98.452% |
|                                                                                                  |
| Server error                                                                                     |
| a13c9f2                0.200%..............======#######|######======........................... |
| f4d7e81                0.230%..............=========##########|##########=========......X....... |
|                              0.076%                                                       0.384% |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
+--------------------------------------------------------------------------------------------------+
```

### [Go usage](https://lanej.io/charmed/molecules-chart.html#chart%2fband-matrix-code)

Imports: [`chart`](https://pkg.go.dev/github.com/lanej/charmed/chart "github.com/lanej/charmed/chart") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
matrix := chart.BandMatrix{
	Title: "Host distributions",
	Rows: []chart.BandRow{
		{Title: "Success", Distribution: chart.DistributionComparison{
			Series: []chart.DistributionSeries{
				{Label: "Reference", Mean: 96, StdDev: .4, Members: 64, HasSpread: true},
				{Label: "Candidate", Mean: 97, StdDev: .6, Members: 32, HasSpread: true},
			},
			Sigma: 3, Min: 0, Max: 100, Unit: "%", SpreadUnit: "pp",
		}},
		{Title: "Server error", Distribution: chart.DistributionComparison{
			Series: []chart.DistributionSeries{
				{Label: "Reference", Mean: .2, StdDev: .03, Members: 64, HasSpread: true},
				{Label: "Candidate", Mean: .23, StdDev: .045, Members: 32, HasSpread: true},
			},
			Sigma: 3, Min: 0, Max: 100, Unit: "%", SpreadUnit: "pp",
		}},
	},
}

return matrix.View(layout.Size{Width: 64, Height: 12}, palette, false)
// Add one BandRow per metric; each row owns its axis shared by both versions.
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-chart.html#chart%2fband-matrix-integration)

Keep observations and their units in the application, and pass changes through the chart's public data API. Allocate terminal columns and rows from the parent layout. Forward native model commands when enabling interaction, and check whether resizing rebuilds plotted data before applying native overrides.

### [Run this example](https://lanej.io/charmed/molecules-chart.html#chart%2fband-matrix-run)

```sh
go run ./examples/catalog chart/band-matrix
```

[Example source](https://github.com/lanej/charmed/blob/682a17afe4ccff26744d2ac202a2bcda8d784cf8/docs/snippets/chart_band_matrix.go) · [Implementation](https://github.com/lanej/charmed/blob/682a17afe4ccff26744d2ac202a2bcda8d784cf8/chart/comparison.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [chart/distributions](https://lanej.io/charmed/molecules-chart.html#chart/distributions) [foundations/layout](https://lanej.io/charmed/foundations-layout.html#foundations/layout)

[Molecules](https://lanej.io/charmed/molecules.html) / Charts

## [Metric trend](https://lanej.io/charmed/molecules-chart.html#chart%2fmetric)

A metric value, unit and sparkline share one bounded allocation; unavailable values and missing history remain distinct.

### [Native preview](https://lanej.io/charmed/molecules-chart.html#chart%2fmetric-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Throughput                      
1,240 req/min                   
                                
                         ▁▂▂▄▅▅█
                         ███████
                         ███████
```

Demo · truecolor · 100 × 30 terminal cells

```
Throughput                                                                                          
1360 req/min                                                                                        
                                                                                                    
                                                    ▂▃▃▃▃▄▅▅▆▄▂            ▁▃▆██▇▄▁            ▂▅▇▇▆
                                                    ███████████▇▃▁  ▁▃▅▆▆▆▇████████▆▄▃▃▄▃▃▂▁▂▄▇█████
                                                    ████████████████████████████████████████████████
                                                    ████████████████████████████████████████████████
                                                    ████████████████████████████████████████████████
```

Narrow · truecolor · 40 × 16 terminal cells

```
Throughput                              
1360 req/min                            
                                        
▆▄▂            ▁▃▆██▇▄▁            ▂▅▇▇▆
███▇▃▁  ▁▃▅▆▆▆▇████████▆▄▃▃▄▃▃▂▁▂▄▇█████
████████████████████████████████████████
████████████████████████████████████████
████████████████████████████████████████
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Throughput                                                                                          
1360 req/min                                                                                        
                                                                                                    
                                                    ▂▃▃▃▃▄▅▅▆▄▂            ▁▃▆██▇▄▁            ▂▅▇▇▆
                                                    ███████████▇▃▁  ▁▃▅▆▆▆▇████████▆▄▃▃▄▃▃▂▁▂▄▇█████
                                                    ████████████████████████████████████████████████
                                                    ████████████████████████████████████████████████
                                                    ████████████████████████████████████████████████
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Throughput                                                                                          
1360 req/min                                                                                        
                                                                                                    
                                                    ▂▃▃▃▃▄▅▅▆▄▂            ▁▃▆██▇▄▁            ▂▅▇▇▆
                                                    ███████████▇▃▁  ▁▃▅▆▆▆▇████████▆▄▃▃▄▃▃▂▁▂▄▇█████
                                                    ████████████████████████████████████████████████
                                                    ████████████████████████████████████████████████
                                                    ████████████████████████████████████████████████
```

Monochrome · none · 100 × 30 terminal cells

```
Throughput                                                                                          
1360 req/min                                                                                        
                                                                                                    
latest 82.4
```

### [Go usage](https://lanej.io/charmed/molecules-chart.html#chart%2fmetric-code)

Imports: [`chart`](https://pkg.go.dev/github.com/lanej/charmed/chart "github.com/lanej/charmed/chart") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
history := []float64{900, 950, 920, 1040, 1100, 1080, 1240}

return chart.MetricTrend("Throughput", "1,240", "req/min", history,
	layout.Size{Width: 32, Height: 6}, theme.New(theme.Dark))
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-chart.html#chart%2fmetric-integration)

Keep observations and their units in the application, and pass changes through the chart's public data API. Allocate terminal columns and rows from the parent layout. Forward native model commands when enabling interaction, and check whether resizing rebuilds plotted data before applying native overrides.

### [Run this example](https://lanej.io/charmed/molecules-chart.html#chart%2fmetric-run)

```sh
go run ./examples/catalog chart/metric
```

[Example source](https://github.com/lanej/charmed/blob/682a17afe4ccff26744d2ac202a2bcda8d784cf8/docs/snippets/chart_metric.go) · [Implementation](https://github.com/lanej/charmed/blob/682a17afe4ccff26744d2ac202a2bcda8d784cf8/story/molecules.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [chart/sparkline](https://lanej.io/charmed/atoms-chart.html#chart/sparkline)

[Molecules](https://lanej.io/charmed/molecules.html) / Charts

## [Chart panel](https://lanej.io/charmed/molecules-chart.html#chart%2fpanel)

A chart and title compose into a reusable region with an optional color-matched legend.

### [Native preview](https://lanej.io/charmed/molecules-chart.html#chart%2fpanel-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
┌ Throughput ──────────────────────────────────────────────────┐
│ 99│                                                          │
│ 94│                                        ⡠⠤⠤⣀⣀⣀⣀           │
│   │                                       ⡔⠁      ⠉⠉⠢⣀       │
│ 85│                                     ⢠⠊         ⢀  ⠑⢄     │
│   │     ⢀⡠⠔⠢⢄⡀                         ⡰⠁        ⣀⠔⠁⠉⠢⣀ ⠑⠤⡀  │
│ 77│ ⣀⠤⠒⠊⠁    ⠈⠒⢄⡀                    ⢀⠜        ⡠⠊      ⠑⠤⡀⠈⠢ │
│   │⠉            ⠈⠑⠤⡀                ⡠⠃      ⢀⠤⠊          ⠈⠒⢄ │
│ 68│    ⢀⣀⠤⠔⠒⠢⠤⢄⣀⡀  ⠈⠒⢄            ⡠⠔⠁    ⢀⡠⠔⠁                │
│   │⠤⠒⠊⠉⠁        ⠈⠉⠑⠢⢄ ⠑⠢⡀     ⢀⡠⠔⠉    ⣀⠤⠊⠁                   │
│ 59│                  ⠑⠢⡀⠈⠢⢄⢀⡠⠒⠁   ⣀⠤⠒⠉                       │
│   │                    ⠈⠑⢄⡀⠁ ⣀⡠⠔⠊⠉                           │
│ 50│                       ⠈⠒⠉                                │
│   │                                                          │
│ 41└───────────────────────────────────────────────────────── │
│   12:00:00 12:01:06 12:02:12 12:03:18 12:04:25 12:05:31      │
│ 1 API  2 Worker                                              │
└──────────────────────────────────────────────────────────────┘
```

Demo · truecolor · 100 × 30 terminal cells

```
┌ Request throughput ──────────────────────────────────────────────────────────────────────────────┐
│ 105│                                                                                             │
│    │                                                                                             │
│ 100│                                                                                             │
│    │⠉⠉⠉⠑⢄                                                                                        │
│  96│     ⠑⢄                                                                                      │
│    │      ⠈⠢⡀                                              ⢀⠤⠤⠤⠤⣀                                │
│  91│        ⠑⢄                                           ⡠⠒⠁     ⠉⠒⠤⢄⣀⣀⣀⣀⣀⠤⠤⠤⠒⠒⠢⠤⡀               │
│    │         ⠈⠢⡀                                       ⢀⠜                        ⠈⢑⣄             │
│  86│           ⠑⢄                                     ⡠⠃                       ⢀⠎⠉⠁ ⠳⡢⡀          │
│    │             ⠑⢄                                  ⡰⠁                       ⡰⠁     ⠑⡌⢆         │
│  82│               ⠑⠢⢄⣀⡠⠔⠒⠊⠉⠒⢄⡀                     ⡰⠁                      ⢀⠜        ⠘⡌⢢        │
│    │              ⢀⡠⠔⠊⠉⠉⠒⠒⠢⠤⣀ ⠈⢆                   ⡰⠁                      ⢠⠊          ⠘⡄⢣       │
│  77│  ⡠⠒⠒⠉⠉⠉⠉⠉⠉⠒⠒⠉⠁          ⠑⠢⡀⠣⡀                ⡰⠁                      ⡰⠁            ⠘⡄⢣      │
│    │⠔⠉                         ⠈⢆⠱⡀              ⢠⠃                     ⢀⠜               ⠘⡄⢣     │
│  72│                             ⠑⢵⡀            ⢠⠃                     ⡠⠊                 ⠘⡄⠣⡀   │
│    │                               ⠳⡀          ⢠⠃                    ⡠⠊                    ⠘⡄⠱⡀  │
│  68│                                ⠹⡢⡀       ⡰⠁                 ⣀⡠⠔⠉                       ⠈⠢⡘⡄ │
│    │                                 ⠱⡈⠢⢄⣀ ⢀⡠⠔⠁           ⣀⠤⠒⠒⠉⠉⠉                             ⠑⢌ │
│  63│                                  ⠱⡀  ⠉⠁           ⢀⠔⠉                                       │
│    │                                   ⠑⡄            ⢀⠔⠁                                         │
│  58│                                    ⠘⢄         ⢀⠔⠁                                           │
│    │                                     ⠈⢆       ⡔⠁                                             │
│  53│                                       ⠉⠢⢄⣀⡠⠤⠊                                               │
│    │                                                                                             │
│  49│                                                                                             │
│    │                                                                                             │
│  44└──────────────────────────────────────────────────────────────────────────────────────────── │
│    12:00:00 12:03:01 12:06:03 12:09:05 12:12:07 12:15:09 12:18:11 12:21:13 12:24:15 12:27:17     │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
```

Narrow · truecolor · 40 × 16 terminal cells

```
┌ Request throughput ──────────────────┐
│ 105│                                 │
│    │⠒⢄                               │
│  95│  ⢣                ⡠⠤⢄⡀  ⢀⡀      │
│    │   ⢣              ⡜   ⠈⠉⠉⠁⢈⣢⡀    │
│  85│    ⠑⢄ ⣀⢄        ⡜       ⢠⠃ ⢳⡀   │
│    │⢀⠤⠤⠤⠤⠒⠉⠉⠢⣣      ⡰⠁      ⢠⠃   ⢷⡀  │
│  75│⠊        ⠘⢧    ⡰⠁      ⡠⠃    ⠘⣵  │
│    │          ⠈⣦  ⡰⠁    ⢀⡠⠔⠁      ⠸⣣ │
│  64│           ⠘⡍⠊    ⡔⠊⠁          ⠑ │
│    │            ⠱⡀  ⢀⠎               │
│  54│             ⠑⠤⠊⠁                │
│    │                                 │
│  44└──────────────────────────────── │
│    12:00:00 12:08:43 12:17:26        │
└──────────────────────────────────────┘
```

256 colors · ansi256 · 100 × 30 terminal cells

```
┌ Request throughput ──────────────────────────────────────────────────────────────────────────────┐
│ 105│                                                                                             │
│    │                                                                                             │
│ 100│                                                                                             │
│    │⠉⠉⠉⠑⢄                                                                                        │
│  96│     ⠑⢄                                                                                      │
│    │      ⠈⠢⡀                                              ⢀⠤⠤⠤⠤⣀                                │
│  91│        ⠑⢄                                           ⡠⠒⠁     ⠉⠒⠤⢄⣀⣀⣀⣀⣀⠤⠤⠤⠒⠒⠢⠤⡀               │
│    │         ⠈⠢⡀                                       ⢀⠜                        ⠈⢑⣄             │
│  86│           ⠑⢄                                     ⡠⠃                       ⢀⠎⠉⠁ ⠳⡢⡀          │
│    │             ⠑⢄                                  ⡰⠁                       ⡰⠁     ⠑⡌⢆         │
│  82│               ⠑⠢⢄⣀⡠⠔⠒⠊⠉⠒⢄⡀                     ⡰⠁                      ⢀⠜        ⠘⡌⢢        │
│    │              ⢀⡠⠔⠊⠉⠉⠒⠒⠢⠤⣀ ⠈⢆                   ⡰⠁                      ⢠⠊          ⠘⡄⢣       │
│  77│  ⡠⠒⠒⠉⠉⠉⠉⠉⠉⠒⠒⠉⠁          ⠑⠢⡀⠣⡀                ⡰⠁                      ⡰⠁            ⠘⡄⢣      │
│    │⠔⠉                         ⠈⢆⠱⡀              ⢠⠃                     ⢀⠜               ⠘⡄⢣     │
│  72│                             ⠑⢵⡀            ⢠⠃                     ⡠⠊                 ⠘⡄⠣⡀   │
│    │                               ⠳⡀          ⢠⠃                    ⡠⠊                    ⠘⡄⠱⡀  │
│  68│                                ⠹⡢⡀       ⡰⠁                 ⣀⡠⠔⠉                       ⠈⠢⡘⡄ │
│    │                                 ⠱⡈⠢⢄⣀ ⢀⡠⠔⠁           ⣀⠤⠒⠒⠉⠉⠉                             ⠑⢌ │
│  63│                                  ⠱⡀  ⠉⠁           ⢀⠔⠉                                       │
│    │                                   ⠑⡄            ⢀⠔⠁                                         │
│  58│                                    ⠘⢄         ⢀⠔⠁                                           │
│    │                                     ⠈⢆       ⡔⠁                                             │
│  53│                                       ⠉⠢⢄⣀⡠⠤⠊                                               │
│    │                                                                                             │
│  49│                                                                                             │
│    │                                                                                             │
│  44└──────────────────────────────────────────────────────────────────────────────────────────── │
│    12:00:00 12:03:01 12:06:03 12:09:05 12:12:07 12:15:09 12:18:11 12:21:13 12:24:15 12:27:17     │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
```

16 colors · ansi16 · 100 × 30 terminal cells

```
┌ Request throughput ──────────────────────────────────────────────────────────────────────────────┐
│ 105│                                                                                             │
│    │                                                                                             │
│ 100│                                                                                             │
│    │⠉⠉⠉⠑⢄                                                                                        │
│  96│     ⠑⢄                                                                                      │
│    │      ⠈⠢⡀                                              ⢀⠤⠤⠤⠤⣀                                │
│  91│        ⠑⢄                                           ⡠⠒⠁     ⠉⠒⠤⢄⣀⣀⣀⣀⣀⠤⠤⠤⠒⠒⠢⠤⡀               │
│    │         ⠈⠢⡀                                       ⢀⠜                        ⠈⢑⣄             │
│  86│           ⠑⢄                                     ⡠⠃                       ⢀⠎⠉⠁ ⠳⡢⡀          │
│    │             ⠑⢄                                  ⡰⠁                       ⡰⠁     ⠑⡌⢆         │
│  82│               ⠑⠢⢄⣀⡠⠔⠒⠊⠉⠒⢄⡀                     ⡰⠁                      ⢀⠜        ⠘⡌⢢        │
│    │              ⢀⡠⠔⠊⠉⠉⠒⠒⠢⠤⣀ ⠈⢆                   ⡰⠁                      ⢠⠊          ⠘⡄⢣       │
│  77│  ⡠⠒⠒⠉⠉⠉⠉⠉⠉⠒⠒⠉⠁          ⠑⠢⡀⠣⡀                ⡰⠁                      ⡰⠁            ⠘⡄⢣      │
│    │⠔⠉                         ⠈⢆⠱⡀              ⢠⠃                     ⢀⠜               ⠘⡄⢣     │
│  72│                             ⠑⢵⡀            ⢠⠃                     ⡠⠊                 ⠘⡄⠣⡀   │
│    │                               ⠳⡀          ⢠⠃                    ⡠⠊                    ⠘⡄⠱⡀  │
│  68│                                ⠹⡢⡀       ⡰⠁                 ⣀⡠⠔⠉                       ⠈⠢⡘⡄ │
│    │                                 ⠱⡈⠢⢄⣀ ⢀⡠⠔⠁           ⣀⠤⠒⠒⠉⠉⠉                             ⠑⢌ │
│  63│                                  ⠱⡀  ⠉⠁           ⢀⠔⠉                                       │
│    │                                   ⠑⡄            ⢀⠔⠁                                         │
│  58│                                    ⠘⢄         ⢀⠔⠁                                           │
│    │                                     ⠈⢆       ⡔⠁                                             │
│  53│                                       ⠉⠢⢄⣀⡠⠤⠊                                               │
│    │                                                                                             │
│  49│                                                                                             │
│    │                                                                                             │
│  44└──────────────────────────────────────────────────────────────────────────────────────────── │
│    12:00:00 12:03:01 12:06:03 12:09:05 12:12:07 12:15:09 12:18:11 12:21:13 12:24:15 12:27:17     │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
```

Monochrome · none · 100 × 30 terminal cells

```
+ Request throughput ------------------------------------------------------------------------------+
| Series 1: 65.0 req/min                                                                           |
| Series 2: 62.9 req/min                                                                           |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
|                                                                                                  |
+--------------------------------------------------------------------------------------------------+
```

With legend · truecolor · 100 × 30 terminal cells

```
┌ Request throughput ──────────────────────────────────────────────────────────────────────────────┐
│ 105│                                                                                             │
│ 103│                                                                                             │
│    │⣀⣀                                                                                           │
│  98│  ⠉⠑⠤⡀                                                                                       │
│    │     ⠈⢢                                                                                      │
│  93│       ⠑⡄                                              ⣀⠤⠤⠤⠤⣀⡀                               │
│    │        ⠈⢢                                          ⢀⠔⠊      ⠈⠉⠒⠢⠤⠤⠤⠤⠤⠒⠒⠒⠉⠉⠑⠒⢄⡀              │
│  88│          ⠑⡄                                       ⡠⠊                       ⢀⣀⡨⠦⡀            │
│    │           ⠈⠢⡀                                   ⢀⠔⠁                       ⡰⠁   ⠙⡕⢄          │
│  83│             ⠈⠢⣀    ⢀⣀⣀⣀⡀                       ⢀⠎                       ⢠⠊      ⠈⢆⠣⡀        │
│    │                ⢉⡲⠶⠮⣁⣀⡀ ⠈⠑⠢⡀                   ⢀⠎                       ⡰⠁         ⢣⠑⡄       │
│  78│   ⣀⣀⠤⠤⠤⣀⣀⣀⣀⣀⠤⠔⠊⠁     ⠈⠉⠒⢄⡀⠑⡄                 ⢀⠎                      ⢀⠜            ⢣⠈⢆      │
│    │⢀⠤⠊                       ⠈⠢⡈⢆                ⡎                      ⢠⠊              ⠣⡈⢆     │
│  73│⠁                           ⠈⠪⣆              ⡜                      ⡰⠁                ⠱⡈⢆    │
│    │                              ⠙⢆            ⡜                     ⡠⠊                   ⠱⡀⢣   │
│  68│                               ⠈⢦⡀        ⢀⠎                   ⣀⠤⠊                      ⠑⡄⠱⡀ │
│    │                                 ⢫⠢⡀    ⢀⡠⠃             ⣀⠤⠔⠒⠒⠉⠉                          ⠈⢢⠑ │
│  64│                                  ⢣⠈⠉⠑⠒⠊⠁           ⢀⠤⠒⠉                                   ⠑ │
│    │                                   ⠣⡀             ⡠⠒⠁                                        │
│  59│                                    ⠱⡀          ⡠⠊                                           │
│    │                                     ⠘⡄       ⡠⠊                                             │
│  54│                                      ⠈⠒⢄⡀⢀⣀⠤⠊                                               │
│    │                                         ⠈⠁                                                  │
│  49│                                                                                             │
│    │                                                                                             │
│  44└──────────────────────────────────────────────────────────────────────────────────────────── │
│    12:00:00 12:03:01 12:06:03 12:09:05 12:12:07 12:15:09 12:18:11 12:21:13 12:24:15 12:27:17     │
│ 1 Series 1  2 Series 2                                                                           │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
```

Narrow with legend · truecolor · 40 × 16 terminal cells

```
┌ Request throughput ──────────────────┐
│ 127│                                 │
│    │          ⣀⠤⣀⣀⠤⠒⠒⠊⢢⡀             │
│ 110│⡢⡀      ⡰⠉      ⡠⠊⠁⠈⢦⡀         ⢀ │
│    │⠚⢳⡀   ⢀⠜      ⢠⠊    ⠈⢗⢄   ⢀⠔⠊⠉⠉⠁ │
│  93│  ⠙⢦⣀⡠⠊   ⣀⣀⣀⠔⠁   ⡔⠊⠉⠑⠓⠳⠶⠾⢗⣢⡀    │
│    │ ⣀⣀⣈⣚⣶⣤⠶⠶⣍      ⢀⠎      ⢀⠎⠁ ⠙⡦⣄⣀ │
│  76│⠊        ⠈⢣⡀  ⢀⠔⠁     ⢀⠔⠁    ⠈⢗⡄ │
│    │           ⠹⡒⠒⠁  ⢀⠔⠒⠊⠉⠁        ⠙ │
│  58│            ⠈⠒⠤⠤⠒⠁               │
│    │                                 │
│  41└──────────────────────────────── │
│    12:00:00 12:08:43 12:17:26        │
│ 1 Series 1  2 Series 2  3 Series 3   │
│ 4 Series 4                           │
└──────────────────────────────────────┘
```

### [Go usage](https://lanej.io/charmed/molecules-chart.html#chart%2fpanel-code)

Imports: [`time`](https://pkg.go.dev/time "time") · [`timeserieslinechart`](https://pkg.go.dev/github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart "github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart") · [`chart`](https://pkg.go.dev/github.com/lanej/charmed/chart "github.com/lanej/charmed/chart") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
start := time.Date(2026, 11, 8, 12, 0, 0, 0, time.UTC)
series := []chart.TimeSeriesData{
	{ID: "api", Label: "API", Unit: "req/min"},
	{ID: "worker", Label: "Worker", Unit: "req/min"},
}

samples := [][]float64{{72, 80, 68, 54, 66, 92, 90, 74}, {60, 66, 62, 48, 56, 66, 82, 68}}
for i, values := range samples {
	for minute, value := range values {
		series[i].Points = append(series[i].Points, timeserieslinechart.TimePoint{
			Time: start.Add(time.Duration(minute) * time.Minute), Value: value,
		})
	}
}

size := layout.Size{Width: 64, Height: 18}
options := chart.PanelOptions{ShowLegend: true}
plot := chart.NewTimeSeries(series, options.PlotSize(size, series, palette), palette)

return chart.Panel("Throughput", plot, series, size, palette, false, options)
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-chart.html#chart%2fpanel-integration)

Keep observations and their units in the application, and pass changes through the chart's public data API. Allocate terminal columns and rows from the parent layout. Forward native model commands when enabling interaction, and check whether resizing rebuilds plotted data before applying native overrides.

### [Run this example](https://lanej.io/charmed/molecules-chart.html#chart%2fpanel-run)

```sh
go run ./examples/catalog chart/panel
```

[Example source](https://github.com/lanej/charmed/blob/682a17afe4ccff26744d2ac202a2bcda8d784cf8/docs/snippets/chart_panel.go) · [Implementation](https://github.com/lanej/charmed/blob/682a17afe4ccff26744d2ac202a2bcda8d784cf8/story/molecules.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [chart/timeseries](https://lanej.io/charmed/atoms-chart.html#chart/timeseries) [foundations/layout](https://lanej.io/charmed/foundations-layout.html#foundations/layout)

[Molecules](https://lanej.io/charmed/molecules.html) / Charts

## [Cohort distributions](https://lanej.io/charmed/molecules-chart.html#chart%2fdistributions)

Mean/spread labels, shared-axis interval bands and optional boundary markers compose into a reusable distribution comparison.

### [Native preview](https://lanej.io/charmed/molecules-chart.html#chart%2fdistributions-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Reference mean 20.000% s=3.000pp n=64
··········░░░▒▒▒▒▓▓▓|▓▓▓▒▒▒W░░E·································
Candidate mean 30.000% s=8.000pp n=32
····░░░░░░░░░▒▒▒▒▒▒▒▒▒▓▓▓▓▓W▓▓E▓|▓▓▓▓▓▓▓▓▓▒X▒▒▒▒▒▒▒░░░░░░░░░····
2.640%                                                   57.360%
| Reference mean  W warning boundary  E error boundary
| Candidate mean  X selected observation
```

Demo · truecolor · 100 × 30 terminal cells

```
Reference mean 20.000% s=3.000pp n=64                                                               
···············░░░░░░▒▒▒▒▒▓▓▓▓▓|▓▓▓▓▓▓▒▒▒▒W░░░░░E···················································
Candidate mean 30.000% s=8.000pp n=32                                                               
······░░░░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓W▓▓▓▓▓E▓|▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒▒X▒▒▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░░······
2.640%                                                                                       57.360%
```

Narrow · truecolor · 40 × 16 terminal cells

```
Reference mean 20.000% s=3.000pp n=64   
······░░▒▒▓▓|▓▓▓▒W░E····················
Candidate mean 30.000% s=8.000pp n=32   
··░░░░░░▒▒▒▒▒▒▓▓▓W▓E|▓▓▓▓▓▒X▒▒▒▒░░░░░░··
2.640%                           57.360%
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Reference mean 20.000% s=3.000pp n=64                                                               
···············░░░░░░▒▒▒▒▒▓▓▓▓▓|▓▓▓▓▓▓▒▒▒▒W░░░░░E···················································
Candidate mean 30.000% s=8.000pp n=32                                                               
······░░░░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓W▓▓▓▓▓E▓|▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒▒X▒▒▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░░······
2.640%                                                                                       57.360%
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Reference mean 20.000% s=3.000pp n=64                                                               
···············░░░░░░▒▒▒▒▒▓▓▓▓▓|▓▓▓▓▓▓▒▒▒▒W░░░░░E···················································
Candidate mean 30.000% s=8.000pp n=32                                                               
······░░░░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓W▓▓▓▓▓E▓|▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒▒X▒▒▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░░······
2.640%                                                                                       57.360%
```

Monochrome · none · 100 × 30 terminal cells

```
Reference mean 20.000% s=3.000pp n=64                                                               
.....................=====#####|######====W.....E...................................................
Candidate mean 30.000% s=8.000pp n=32                                                               
.....................==============#######W#####E#|##############===X==========.....................
2.640%                                                                                       57.360%
```

With legend · truecolor · 100 × 30 terminal cells

```
Reference mean 20.000% s=3.000pp n=64                                                               
···············░░░░░░▒▒▒▒▒▓▓▓▓▓|▓▓▓▓▓▓▒▒▒▒W░░░░░E···················································
Candidate mean 30.000% s=8.000pp n=32                                                               
······░░░░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓W▓▓▓▓▓E▓|▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒▒X▒▒▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░░······
2.640%                                                                                       57.360%
| Reference mean  W warning boundary  E error boundary  | Candidate mean  X selected observation
```

No boundaries · truecolor · 100 × 30 terminal cells

```
Reference mean 20.000% s=3.000pp n=64                                                               
···············░░░░░░▒▒▒▒▒▓▓▓▓▓|▓▓▓▓▓▓▒▒▒▒▒░░░░░░···················································
Candidate mean 30.000% s=8.000pp n=32                                                               
······░░░░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓|▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒▒X▒▒▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░░······
2.640%                                                                                       57.360%
```

### [Go usage](https://lanej.io/charmed/molecules-chart.html#chart%2fdistributions-code)

Imports: [`attention`](https://pkg.go.dev/github.com/lanej/charmed/attention "github.com/lanej/charmed/attention") · [`chart`](https://pkg.go.dev/github.com/lanej/charmed/chart "github.com/lanej/charmed/chart") · [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
selected := 40.0
comparison := chart.DistributionComparison{
	Series: []chart.DistributionSeries{
		{Label: "Reference", Mean: 20, StdDev: 3, Members: 64, HasSpread: true},
		{Label: "Candidate", Mean: 30, StdDev: 8, Members: 32, HasSpread: true, Selected: &selected},
	},
	Sigma: 3, Min: 0, Max: 60, Unit: "%", SpreadUnit: "pp",
	ShowLegend: true,
	Markers: []chart.BandMarker{
		{Value: 26, Glyph: 'W', Label: "warning boundary", Style: feedback.AttentionStyle(attention.Warn, palette)},
		{Value: 29, Glyph: 'E', Label: "error boundary", Style: feedback.AttentionStyle(attention.Error, palette)},
	},
}

return comparison.View(64, palette)
// Markers are optional caller-supplied boundaries; no rule is inferred from spread.
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-chart.html#chart%2fdistributions-integration)

Keep observations and their units in the application, and pass changes through the chart's public data API. Allocate terminal columns and rows from the parent layout. Forward native model commands when enabling interaction, and check whether resizing rebuilds plotted data before applying native overrides.

### [Run this example](https://lanej.io/charmed/molecules-chart.html#chart%2fdistributions-run)

```sh
go run ./examples/catalog chart/distributions
```

[Example source](https://github.com/lanej/charmed/blob/682a17afe4ccff26744d2ac202a2bcda8d784cf8/docs/snippets/chart_distributions.go) · [Implementation](https://github.com/lanej/charmed/blob/682a17afe4ccff26744d2ac202a2bcda8d784cf8/chart/comparison.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [chart/interval-band](https://lanej.io/charmed/atoms-chart.html#chart/interval-band)

[Molecules](https://lanej.io/charmed/molecules.html) / Charts

## [Merged delta property](https://lanej.io/charmed/molecules-chart.html#chart%2fdelta)

A single zero-centred bar cancels the shared reference/candidate amount and emphasizes the signed percentage-point difference.

### [Native preview](https://lanej.io/charmed/molecules-chart.html#chart%2fdelta-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
··············|▓▓▓▓▓·········   +0.180pp
```

Demo · truecolor · 100 × 30 terminal cells

```
Merged delta property                                                                               
                                                                                                    
············································|▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓····························   +0.180pp
```

Narrow · truecolor · 40 × 16 terminal cells

```
Merged delta property                   
                                        
··············|▓▓▓▓▓·········   +0.180pp
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Merged delta property                                                                               
                                                                                                    
············································|▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓····························   +0.180pp
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Merged delta property                                                                               
                                                                                                    
············································|▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓····························   +0.180pp
```

Monochrome · none · 100 × 30 terminal cells

```
Merged delta property                                                                               
                                                                                                    
............................................|################............................   +0.180pp
```

### [Go usage](https://lanej.io/charmed/molecules-chart.html#chart%2fdelta-code)

Imports: [`chart`](https://pkg.go.dev/github.com/lanej/charmed/chart "github.com/lanej/charmed/chart") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
reference, candidate := 20.0, 20.18 // Percentage points, not fractions.
property := chart.DeltaProperty{
	Value: candidate - reference, Extent: .5,
	Available: true, Style: palette.Text(),
}

return property.View(40, palette)
// Reuse Extent across peer properties to keep their zero-centred scales comparable.
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-chart.html#chart%2fdelta-integration)

Keep observations and their units in the application, and pass changes through the chart's public data API. Allocate terminal columns and rows from the parent layout. Forward native model commands when enabling interaction, and check whether resizing rebuilds plotted data before applying native overrides.

### [Run this example](https://lanej.io/charmed/molecules-chart.html#chart%2fdelta-run)

```sh
go run ./examples/catalog chart/delta
```

[Example source](https://github.com/lanej/charmed/blob/682a17afe4ccff26744d2ac202a2bcda8d784cf8/docs/snippets/chart_delta.go) · [Implementation](https://github.com/lanej/charmed/blob/682a17afe4ccff26744d2ac202a2bcda8d784cf8/chart/delta.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [chart/interval-band](https://lanej.io/charmed/atoms-chart.html#chart/interval-band) [text/signed-number](https://lanej.io/charmed/atoms-text.html#text/signed-number)
