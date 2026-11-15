[Atoms](https://lanej.io/charmed/atoms.html)Charts

# [Charts](https://lanej.io/charmed/atoms-chart.html#level-atoms-chart)

12 components · [All atoms](https://lanej.io/charmed/atoms.html)

- [Sparkline](https://lanej.io/charmed/atoms-chart.html#chart%2fsparkline)
- [Time series](https://lanej.io/charmed/atoms-chart.html#chart%2ftimeseries)
- [Line](https://lanej.io/charmed/atoms-chart.html#chart%2fline)
- [Scatter](https://lanej.io/charmed/atoms-chart.html#chart%2fscatter)
- [Bars](https://lanej.io/charmed/atoms-chart.html#chart%2fbars)
- [Heatmap](https://lanej.io/charmed/atoms-chart.html#chart%2fheatmap)
- [Streamline](https://lanej.io/charmed/atoms-chart.html#chart%2fstream)
- [Waveline](https://lanej.io/charmed/atoms-chart.html#chart%2fwave)
- [OHLC candles](https://lanej.io/charmed/atoms-chart.html#chart%2fohlc)
- [Canvas](https://lanej.io/charmed/atoms-chart.html#chart%2fcanvas)
- [Waffle](https://lanej.io/charmed/atoms-chart.html#chart%2fwaffle)
- [Interval band](https://lanej.io/charmed/atoms-chart.html#chart%2finterval-band)

[Atoms](https://lanej.io/charmed/atoms.html) / Charts

## [Sparkline](https://lanej.io/charmed/atoms-chart.html#chart%2fsparkline)

Show a compact numeric trend beside a current value or inside a table cell. Use a full time-series chart when readers need labeled axes and timestamps.

### [Native preview](https://lanej.io/charmed/atoms-chart.html#chart%2fsparkline-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
                                          ▂ ▄▃█▆
                                      ▁▄▂▇██████
                                      ██████████
                                      ██████████
                                      ██████████
```

Demo · truecolor · 100 × 30 terminal cells

```
Chart / sparkline                                                                                   
                                                                                                    
                                                    ▂▃▃▃▃▄▅▅▆▄▂            ▁▃▆██▇▄▁            ▂▅▇▇▆
                                                    ███████████▇▃▁  ▁▃▅▆▆▆▇████████▆▄▃▃▄▃▃▂▁▂▄▇█████
                                                    ████████████████████████████████████████████████
                                                    ████████████████████████████████████████████████
                                                    ████████████████████████████████████████████████
                                                                                                    
Latest 82.4 req/min                                                                                 
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

Narrow · truecolor · 40 × 16 terminal cells

```
Chart / sparkline                       
                                        
▆▄▂            ▁▃▆██▇▄▁            ▂▅▇▇▆
███▇▃▁  ▁▃▅▆▆▆▇████████▆▄▃▃▄▃▃▂▁▂▄▇█████
████████████████████████████████████████
████████████████████████████████████████
████████████████████████████████████████
                                        
Latest 82.4 req/min                     
                                        
Ctrl+N advances fixture data / arrows pa
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Chart / sparkline                                                                                   
                                                                                                    
                                                    ▂▃▃▃▃▄▅▅▆▄▂            ▁▃▆██▇▄▁            ▂▅▇▇▆
                                                    ███████████▇▃▁  ▁▃▅▆▆▆▇████████▆▄▃▃▄▃▃▂▁▂▄▇█████
                                                    ████████████████████████████████████████████████
                                                    ████████████████████████████████████████████████
                                                    ████████████████████████████████████████████████
                                                                                                    
Latest 82.4 req/min                                                                                 
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Chart / sparkline                                                                                   
                                                                                                    
                                                    ▂▃▃▃▃▄▅▅▆▄▂            ▁▃▆██▇▄▁            ▂▅▇▇▆
                                                    ███████████▇▃▁  ▁▃▅▆▆▆▇████████▆▄▃▃▄▃▃▂▁▂▄▇█████
                                                    ████████████████████████████████████████████████
                                                    ████████████████████████████████████████████████
                                                    ████████████████████████████████████████████████
                                                                                                    
Latest 82.4 req/min                                                                                 
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

Monochrome · none · 100 × 30 terminal cells

```
sparkline / text fallback                                                                           
                                                                                                    
2 series / latest 68.0                                                                              
                                                                                                    
ASCII mode presents values without chart glyphs.
```

### [Go usage](https://lanej.io/charmed/atoms-chart.html#chart%2fsparkline-code)

Imports: [`chart`](https://pkg.go.dev/github.com/lanej/charmed/chart "github.com/lanej/charmed/chart") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
values := []float64{32, 37, 34, 40, 44, 41, 47, 45, 52, 49}
spark := chart.NewSparkline(values,
	layout.Size{Width: 48, Height: 5}, theme.New(theme.Dark))

return spark.View()
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-chart.html#chart%2fsparkline-integration)

Keep observations and their units in the application, and pass changes through the chart's public data API. Allocate terminal columns and rows from the parent layout. Forward native model commands when enabling interaction, and check whether resizing rebuilds plotted data before applying native overrides.

### [Run this example](https://lanej.io/charmed/atoms-chart.html#chart%2fsparkline-run)

```sh
go run ./examples/catalog chart/sparkline
```

[Example source](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/docs/snippets/chart_sparkline.go) · [Implementation](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/story/atoms_charts.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Charts

## [Time series](https://lanej.io/charmed/atoms-chart.html#chart%2ftimeseries)

Compare timestamped measurements over the same interval. Supply labeled series with explicit units; the shared axes make changes in API and worker throughput directly comparable.

### [Native preview](https://lanej.io/charmed/atoms-chart.html#chart%2ftimeseries-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
99│                                                             
  │                                          ⡠⠒⠒⠢⠤⠤⠤⠤⣀⣀⡀        
89│                                        ⢀⠔⠁         ⠈⠒⢄      
  │        ⣀                              ⡠⠊          ⣀⢄⡀ ⠉⠢⣀   
80│   ⢀⡠⠔⠒⠉ ⠉⠒⠤⣀                        ⢀⠔⠁        ⢀⠤⠊  ⠈⠒⢄⡀ ⠑⠤⡀
  │⠤⠒⠉⠁         ⠑⠢⢄⡀                   ⡠⠃       ⢀⡠⠒⠁       ⠈⠑⠤⡀⠈
70│        ⣀⣀⡀     ⠈⠑⠤⡀              ⢀⠜      ⢀⡠⠔⠁             ⠈⠑
  │ ⣀⣀⠤⠔⠒⠉⠉  ⠈⠉⠉⠑⠒⠒⠢⠤⣀⠈⠒⢄⡀        ⣀⠔⠊⠁   ⢀⡠⠔⠊⠁                  
60│⠉                  ⠑⠤⡀⠈⠒⢄  ⢀⡠⠒⠉   ⣀⡠⠔⠊⠁                      
  │                     ⠈⠑⠤⡀⠉⠊⠁ ⢀⡠⠔⠒⠉                           
51│                        ⠈⠒⠤⠒⠉⠁                               
  │                                                             
41└─────────────────────────────────────────────────────────────
  12:00:00 12:01:01 12:02:03 12:03:05 12:04:07 12:05:09 12:06:11
```

Demo · truecolor · 100 × 30 terminal cells

```
Chart / timeseries                                                                                  
                                                                                                    
105│                                                                                                
   │                                                                                                
100│⠤⠤⠤⢄                                                                                            
   │    ⠉⠢⡀                                                                                         
 94│      ⠈⢢                                                  ⢀⣀⣀⣀⡀                                 
   │        ⠑⢄                                             ⡠⠔⠊⠁   ⠈⠉⠒⠢⠤⢄⣀⣀⣀⣀⣀⠤⠤⠤⠒⠒⠒⠤⢄⡀              
 88│          ⠑⢄                                         ⢀⠔⠁                        ⣀⣈⠶⢄⡀           
   │            ⠑⢄⡀                                     ⡠⠊                        ⡠⠊    ⠹⡢⢄         
 83│              ⠈⠢⢄⣀  ⢀⡠⠤⠤⠤⢄⡀                        ⡰⠁                       ⡠⠊       ⠈⠢⡣⡀       
   │                ⢀⡠⠝⠛⠓⠤⠤⢄⣀⡀⠈⠑⢄                     ⡜                       ⢀⠜           ⠱⡘⢄      
 77│  ⣀⠤⠒⠊⠉⠉⠒⠒⠒⠒⠒⠒⠊⠉⠁        ⠈⠑⠢⡀⠣⡀                 ⢀⠎                       ⡰⠁             ⠘⡄⢣     
   │⠔⠊                          ⠈⠒⢌⢢               ⡠⠃                      ⢀⠎                ⠈⢆⠱⡀   
 72│                               ⠑⢧⡀           ⢀⠜                      ⣀⠔⠁                  ⠈⢢⠑⡄  
   │                                 ⠱⣄         ⡠⠊                   ⢀⡠⠔⠊                       ⠣⡈⢢ 
 66│                                  ⠘⡍⠢⣀    ⣀⠔⠁             ⣀⣀⡠⠤⠔⠒⠉⠁                           ⠈⠢⡑
   │                                   ⠈⢆ ⠉⠑⠒⠉            ⢀⠔⠊⠉                                     ⠈
 61│                                    ⠈⢢              ⡠⠊⠁                                         
   │                                      ⠑⡄         ⢀⠔⠊                                            
 55│                                       ⠈⠢⢄    ⢀⡠⠊⠁                                              
   │                                          ⠉⠒⠒⠉⠁                                                 
 50│                                                                                                
   │                                                                                                
 44└────────────────────────────────────────────────────────────────────────────────────────────────
   12:00:00 12:02:54 12:05:48 12:08:43 12:11:37 12:14:31 12:17:26 12:20:20 12:23:15 12:26:09        
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

Narrow · truecolor · 40 × 16 terminal cells

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
                                        
Ctrl+N advances fixture data / arrows pa
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Chart / timeseries                                                                                  
                                                                                                    
105│                                                                                                
   │                                                                                                
100│⠤⠤⠤⢄                                                                                            
   │    ⠉⠢⡀                                                                                         
 94│      ⠈⢢                                                  ⢀⣀⣀⣀⡀                                 
   │        ⠑⢄                                             ⡠⠔⠊⠁   ⠈⠉⠒⠢⠤⢄⣀⣀⣀⣀⣀⠤⠤⠤⠒⠒⠒⠤⢄⡀              
 88│          ⠑⢄                                         ⢀⠔⠁                        ⣀⣈⠶⢄⡀           
   │            ⠑⢄⡀                                     ⡠⠊                        ⡠⠊    ⠹⡢⢄         
 83│              ⠈⠢⢄⣀  ⢀⡠⠤⠤⠤⢄⡀                        ⡰⠁                       ⡠⠊       ⠈⠢⡣⡀       
   │                ⢀⡠⠝⠛⠓⠤⠤⢄⣀⡀⠈⠑⢄                     ⡜                       ⢀⠜           ⠱⡘⢄      
 77│  ⣀⠤⠒⠊⠉⠉⠒⠒⠒⠒⠒⠒⠊⠉⠁        ⠈⠑⠢⡀⠣⡀                 ⢀⠎                       ⡰⠁             ⠘⡄⢣     
   │⠔⠊                          ⠈⠒⢌⢢               ⡠⠃                      ⢀⠎                ⠈⢆⠱⡀   
 72│                               ⠑⢧⡀           ⢀⠜                      ⣀⠔⠁                  ⠈⢢⠑⡄  
   │                                 ⠱⣄         ⡠⠊                   ⢀⡠⠔⠊                       ⠣⡈⢢ 
 66│                                  ⠘⡍⠢⣀    ⣀⠔⠁             ⣀⣀⡠⠤⠔⠒⠉⠁                           ⠈⠢⡑
   │                                   ⠈⢆ ⠉⠑⠒⠉            ⢀⠔⠊⠉                                     ⠈
 61│                                    ⠈⢢              ⡠⠊⠁                                         
   │                                      ⠑⡄         ⢀⠔⠊                                            
 55│                                       ⠈⠢⢄    ⢀⡠⠊⠁                                              
   │                                          ⠉⠒⠒⠉⠁                                                 
 50│                                                                                                
   │                                                                                                
 44└────────────────────────────────────────────────────────────────────────────────────────────────
   12:00:00 12:02:54 12:05:48 12:08:43 12:11:37 12:14:31 12:17:26 12:20:20 12:23:15 12:26:09        
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Chart / timeseries                                                                                  
                                                                                                    
105│                                                                                                
   │                                                                                                
100│⠤⠤⠤⢄                                                                                            
   │    ⠉⠢⡀                                                                                         
 94│      ⠈⢢                                                  ⢀⣀⣀⣀⡀                                 
   │        ⠑⢄                                             ⡠⠔⠊⠁   ⠈⠉⠒⠢⠤⢄⣀⣀⣀⣀⣀⠤⠤⠤⠒⠒⠒⠤⢄⡀              
 88│          ⠑⢄                                         ⢀⠔⠁                        ⣀⣈⠶⢄⡀           
   │            ⠑⢄⡀                                     ⡠⠊                        ⡠⠊    ⠹⡢⢄         
 83│              ⠈⠢⢄⣀  ⢀⡠⠤⠤⠤⢄⡀                        ⡰⠁                       ⡠⠊       ⠈⠢⡣⡀       
   │                ⢀⡠⠝⠛⠓⠤⠤⢄⣀⡀⠈⠑⢄                     ⡜                       ⢀⠜           ⠱⡘⢄      
 77│  ⣀⠤⠒⠊⠉⠉⠒⠒⠒⠒⠒⠒⠊⠉⠁        ⠈⠑⠢⡀⠣⡀                 ⢀⠎                       ⡰⠁             ⠘⡄⢣     
   │⠔⠊                          ⠈⠒⢌⢢               ⡠⠃                      ⢀⠎                ⠈⢆⠱⡀   
 72│                               ⠑⢧⡀           ⢀⠜                      ⣀⠔⠁                  ⠈⢢⠑⡄  
   │                                 ⠱⣄         ⡠⠊                   ⢀⡠⠔⠊                       ⠣⡈⢢ 
 66│                                  ⠘⡍⠢⣀    ⣀⠔⠁             ⣀⣀⡠⠤⠔⠒⠉⠁                           ⠈⠢⡑
   │                                   ⠈⢆ ⠉⠑⠒⠉            ⢀⠔⠊⠉                                     ⠈
 61│                                    ⠈⢢              ⡠⠊⠁                                         
   │                                      ⠑⡄         ⢀⠔⠊                                            
 55│                                       ⠈⠢⢄    ⢀⡠⠊⠁                                              
   │                                          ⠉⠒⠒⠉⠁                                                 
 50│                                                                                                
   │                                                                                                
 44└────────────────────────────────────────────────────────────────────────────────────────────────
   12:00:00 12:02:54 12:05:48 12:08:43 12:11:37 12:14:31 12:17:26 12:20:20 12:23:15 12:26:09        
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

Monochrome · none · 100 × 30 terminal cells

```
timeseries / text fallback                                                                          
                                                                                                    
2 series / latest 68.0                                                                              
                                                                                                    
ASCII mode presents values without chart glyphs.
```

With legend · truecolor · 100 × 30 terminal cells

```
Chart / timeseries                                                                                  
                                                                                                    
105│                                                                                                
   │                                                                                                
100│⠤⠤⠤⢄                                                                                            
   │    ⠉⠢⡀                                                                                         
 94│      ⠈⢢                                                  ⢀⣀⣀⣀⡀                                 
   │        ⠑⢄                                             ⡠⠔⠊⠁   ⠈⠉⠒⠢⠤⢄⣀⣀⣀⣀⣀⠤⠤⠤⠒⠒⠒⠤⢄⡀              
 88│          ⠑⢄                                         ⢀⠔⠁                        ⣀⣈⠶⢄⡀           
   │            ⠑⢄⡀                                     ⡠⠊                        ⡠⠊    ⠹⡢⢄         
 83│              ⠈⠢⢄⣀  ⢀⡠⠤⠤⠤⢄⡀                        ⡰⠁                       ⡠⠊       ⠈⠢⡣⡀       
   │                ⢀⡠⠝⠛⠓⠤⠤⢄⣀⡀⠈⠑⢄                     ⡜                       ⢀⠜           ⠱⡘⢄      
 77│  ⣀⠤⠒⠊⠉⠉⠒⠒⠒⠒⠒⠒⠊⠉⠁        ⠈⠑⠢⡀⠣⡀                 ⢀⠎                       ⡰⠁             ⠘⡄⢣     
   │⠔⠊                          ⠈⠒⢌⢢               ⡠⠃                      ⢀⠎                ⠈⢆⠱⡀   
 72│                               ⠑⢧⡀           ⢀⠜                      ⣀⠔⠁                  ⠈⢢⠑⡄  
   │                                 ⠱⣄         ⡠⠊                   ⢀⡠⠔⠊                       ⠣⡈⢢ 
 66│                                  ⠘⡍⠢⣀    ⣀⠔⠁             ⣀⣀⡠⠤⠔⠒⠉⠁                           ⠈⠢⡑
   │                                   ⠈⢆ ⠉⠑⠒⠉            ⢀⠔⠊⠉                                     ⠈
 61│                                    ⠈⢢              ⡠⠊⠁                                         
   │                                      ⠑⡄         ⢀⠔⠊                                            
 55│                                       ⠈⠢⢄    ⢀⡠⠊⠁                                              
   │                                          ⠉⠒⠒⠉⠁                                                 
 50│                                                                                                
   │                                                                                                
 44└────────────────────────────────────────────────────────────────────────────────────────────────
   12:00:00 12:02:54 12:05:48 12:08:43 12:11:37 12:14:31 12:17:26 12:20:20 12:23:15 12:26:09        
1 Series 1  2 Series 2                                                                              
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

### [Go usage](https://lanej.io/charmed/atoms-chart.html#chart%2ftimeseries-code)

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

plot := chart.NewTimeSeries(series, layout.Size{Width: 64, Height: 14}, palette)

return plot.View()
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-chart.html#chart%2ftimeseries-integration)

Call SetData when observations change and Resize with the new cell allocation. Both rebuild the native Bubble plot, so reapply direct Bubble customization afterward. Route keys to Update only while focused, retain its returned model and command, and use Focus or Blur when the active region changes.

### [Run this example](https://lanej.io/charmed/atoms-chart.html#chart%2ftimeseries-run)

```sh
go run ./examples/catalog chart/timeseries
```

[Example source](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/docs/snippets/chart_timeseries.go) · [Implementation](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/examples/timeseries/main.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Charts

## [Line](https://lanej.io/charmed/atoms-chart.html#chart%2fline)

Plot numeric observations on explicit horizontal and vertical axes. Choose this primitive when you need direct control over the native line renderer and its data range.

### [Native preview](https://lanej.io/charmed/atoms-chart.html#chart%2fline-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
100│                                                            
   │                                        ⢀⠔⠊⠉⠉⠉⠉⠒⠒⠒⠒⠒⠤⣀⡀     
 83│   ⢀⣀⡠⠤⠔⠒⠢⠤⣀⡀                        ⢀⠤⠊⠁             ⠈⠑⠢⠤⣀ 
   │⠒⠊⠉⠁        ⠈⠑⠒⠤⢄⣀                ⢀⡠⠊⠁                     ⠉
 67│                  ⠉⠒⠢⢄⡀      ⢀⡠⠤⠒⠊⠁                         
   │                      ⠈⠉⠒⠤⠔⠒⠉⠁                              
 50│                                                            
   │                                                            
 33│                                                            
   │                                                            
 17│                                                            
   │                                                            
  0└────────────────────────────────────────────────────────────
   0     1       2       3       4         5       6       7
```

Demo · truecolor · 100 × 30 terminal cells

```
Chart / line                                                                                        
                                                                                                    
120│                                                                                                
   │                                                                                                
109│                                                                                                
   │                                                                                                
 98│                                                                                                
   │                                                                                                
 87│                                                                             ⣀⠤⠔⠒⠒⠉⠉⠒⠒⠤⡀        
   │     ⣀⣀⣀      ⢀⣀⣀⡠⠤⠔⠒⠉⠉⠉⠉⠑⠒⠤⣀                                             ⡠⠔⠉          ⠈⠑⢄      
 76│⠶⠒⠚⠛⠛⠤⠤⣀⠉⠉⠉⠉⠉⠉⠁              ⠉⠒⢄                                       ⡠⠔⠉  ⢀⣀⣀⣀⣀         ⠑⠢⡀   
   │        ⠉⠒⠤⣀                    ⠉⠢⡀                       ⣀⣀⣀⣀⣀⣀⣀⣀⣠⣤⣤⣒⠭⠤⠤⠒⠒⠉⠁    ⠉⠑⠢⢄⡀      ⠈⠒⢄⡀
 65│            ⠉⠉⠑⠒⠒⠢⠤⠤⢄⣀⣀⡀          ⠈⠒⢄                ⢀⣠⠴⠮⠝⠒⠒⠊⠉⠉⠁                     ⠈⠑⢄⡀      ⠈
   │                       ⠈⠉⠑⠒⠤⢄⡀       ⠑⠢⣀         ⣀⡠⢔⠞⠉                                  ⠈⠢⡀     
 55│                             ⠈⠑⠢⢄       ⠉⠒⠢⠤⠤⠒⠒⠉⢉⡠⠊⠁                                      ⠈⠢⡀   
   │                                 ⠉⠢⢄          ⡠⠔⠁                                           ⠈⠒⢄⡀
 44│                                    ⠉⠒⠒⠤⠤⠤⠤⠤⠒⠉                                                 ⠈
   │                                                                                                
 33│                                                                                                
   │                                                                                                
 22│                                                                                                
   │                                                                                                
 11│                                                                                                
   │                                                                                                
  0└────────────────────────────────────────────────────────────────────────────────────────────────
   0 1   2 3   4 5   6   7 8   9 10  11  12  14  15  16  17  19  20  21  23  24  25  26  28  29  30 
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

Narrow · truecolor · 40 × 16 terminal cells

```
Chart / line                            
                                        
120│                                    
   │                                    
 90│ ⣀⣀⣀⣀⣀⣀⣀⠤⠤⣀⡀               ⢀⡠⠔⠒⠒⠢⡀  
   │⠉⠉⠉⠑⠢⢄⣀⣀⣀  ⠈⠢⡀      ⢀⡠⠤⠤⠤⠤⠶⠕⠒⠒⠒⠢⣀⠈⠑⠢
 60│         ⠉⠉⠢⢄⡈⠑⠒⠤⣤⠶⠊⠁            ⠑⠢⡀
   │             ⠈⠑⠒⠊                  ⠈
 30│                                    
   │                                    
  0└────────────────────────────────────
   0 2 3 5 7 9 10  14  17  21  24  28 31
                                        
Ctrl+N advances fixture data / arrows pa
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Chart / line                                                                                        
                                                                                                    
120│                                                                                                
   │                                                                                                
109│                                                                                                
   │                                                                                                
 98│                                                                                                
   │                                                                                                
 87│                                                                             ⣀⠤⠔⠒⠒⠉⠉⠒⠒⠤⡀        
   │     ⣀⣀⣀      ⢀⣀⣀⡠⠤⠔⠒⠉⠉⠉⠉⠑⠒⠤⣀                                             ⡠⠔⠉          ⠈⠑⢄      
 76│⠶⠒⠚⠛⠛⠤⠤⣀⠉⠉⠉⠉⠉⠉⠁              ⠉⠒⢄                                       ⡠⠔⠉  ⢀⣀⣀⣀⣀         ⠑⠢⡀   
   │        ⠉⠒⠤⣀                    ⠉⠢⡀                       ⣀⣀⣀⣀⣀⣀⣀⣀⣠⣤⣤⣒⠭⠤⠤⠒⠒⠉⠁    ⠉⠑⠢⢄⡀      ⠈⠒⢄⡀
 65│            ⠉⠉⠑⠒⠒⠢⠤⠤⢄⣀⣀⡀          ⠈⠒⢄                ⢀⣠⠴⠮⠝⠒⠒⠊⠉⠉⠁                     ⠈⠑⢄⡀      ⠈
   │                       ⠈⠉⠑⠒⠤⢄⡀       ⠑⠢⣀         ⣀⡠⢔⠞⠉                                  ⠈⠢⡀     
 55│                             ⠈⠑⠢⢄       ⠉⠒⠢⠤⠤⠒⠒⠉⢉⡠⠊⠁                                      ⠈⠢⡀   
   │                                 ⠉⠢⢄          ⡠⠔⠁                                           ⠈⠒⢄⡀
 44│                                    ⠉⠒⠒⠤⠤⠤⠤⠤⠒⠉                                                 ⠈
   │                                                                                                
 33│                                                                                                
   │                                                                                                
 22│                                                                                                
   │                                                                                                
 11│                                                                                                
   │                                                                                                
  0└────────────────────────────────────────────────────────────────────────────────────────────────
   0 1   2 3   4 5   6   7 8   9 10  11  12  14  15  16  17  19  20  21  23  24  25  26  28  29  30 
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Chart / line                                                                                        
                                                                                                    
120│                                                                                                
   │                                                                                                
109│                                                                                                
   │                                                                                                
 98│                                                                                                
   │                                                                                                
 87│                                                                             ⣀⠤⠔⠒⠒⠉⠉⠒⠒⠤⡀        
   │     ⣀⣀⣀      ⢀⣀⣀⡠⠤⠔⠒⠉⠉⠉⠉⠑⠒⠤⣀                                             ⡠⠔⠉          ⠈⠑⢄      
 76│⠶⠒⠚⠛⠛⠤⠤⣀⠉⠉⠉⠉⠉⠉⠁              ⠉⠒⢄                                       ⡠⠔⠉  ⢀⣀⣀⣀⣀         ⠑⠢⡀   
   │        ⠉⠒⠤⣀                    ⠉⠢⡀                       ⣀⣀⣀⣀⣀⣀⣀⣀⣠⣤⣤⣒⠭⠤⠤⠒⠒⠉⠁    ⠉⠑⠢⢄⡀      ⠈⠒⢄⡀
 65│            ⠉⠉⠑⠒⠒⠢⠤⠤⢄⣀⣀⡀          ⠈⠒⢄                ⢀⣠⠴⠮⠝⠒⠒⠊⠉⠉⠁                     ⠈⠑⢄⡀      ⠈
   │                       ⠈⠉⠑⠒⠤⢄⡀       ⠑⠢⣀         ⣀⡠⢔⠞⠉                                  ⠈⠢⡀     
 55│                             ⠈⠑⠢⢄       ⠉⠒⠢⠤⠤⠒⠒⠉⢉⡠⠊⠁                                      ⠈⠢⡀   
   │                                 ⠉⠢⢄          ⡠⠔⠁                                           ⠈⠒⢄⡀
 44│                                    ⠉⠒⠒⠤⠤⠤⠤⠤⠒⠉                                                 ⠈
   │                                                                                                
 33│                                                                                                
   │                                                                                                
 22│                                                                                                
   │                                                                                                
 11│                                                                                                
   │                                                                                                
  0└────────────────────────────────────────────────────────────────────────────────────────────────
   0 1   2 3   4 5   6   7 8   9 10  11  12  14  15  16  17  19  20  21  23  24  25  26  28  29  30 
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

Monochrome · none · 100 × 30 terminal cells

```
line / text fallback                                                                                
                                                                                                    
2 series / latest 68.0                                                                              
                                                                                                    
ASCII mode presents values without chart glyphs.
```

### [Go usage](https://lanej.io/charmed/atoms-chart.html#chart%2fline-code)

Imports: [`canvas`](https://pkg.go.dev/github.com/NimbleMarkets/ntcharts/v2/canvas "github.com/NimbleMarkets/ntcharts/v2/canvas") · [`chart`](https://pkg.go.dev/github.com/lanej/charmed/chart "github.com/lanej/charmed/chart") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
values := []float64{72, 80, 68, 54, 66, 92, 90, 74}
plot := chart.NewLine(layout.Size{Width: 64, Height: 14}, 0, 7, 0, 100, palette)
plot.DrawXYAxisAndLabel()

var previous canvas.Float64Point

for i, value := range values {
	point := canvas.Float64Point{X: float64(i), Y: value}
	if i > 0 {
		plot.DrawBrailleLine(previous, point)
	}

	previous = point
}

return plot.View()
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-chart.html#chart%2fline-integration)

Keep observations and their units in the application, and pass changes through the chart's public data API. Allocate terminal columns and rows from the parent layout. Forward native model commands when enabling interaction, and check whether resizing rebuilds plotted data before applying native overrides.

### [Run this example](https://lanej.io/charmed/atoms-chart.html#chart%2fline-run)

```sh
go run ./examples/catalog chart/line
```

[Example source](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/docs/snippets/chart_line.go) · [Implementation](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/story/atoms_charts.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Charts

## [Scatter](https://lanej.io/charmed/atoms-chart.html#chart%2fscatter)

Compare the position of individual observations on two numeric axes. Use distinct markers and labels when multiple groups need to remain recognizable without color.

### [Native preview](https://lanej.io/charmed/atoms-chart.html#chart%2fscatter-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
100│                                                            
   │                                          ●        ●        
 83│        ●                                                   
   │●                                                          ●
 67│                 ●                ●                         
   │                         ●                                  
 50│                                                            
   │                                                            
 33│                                                            
   │                                                            
 17│                                                            
   │                                                            
  0└────────────────────────────────────────────────────────────
   0     1       2       3       4         5       6       7
```

Demo · truecolor · 100 × 30 terminal cells

```
Chart / scatter                                                                                     
                                                                                                    
120│                                                                                                
   │                                                                                                
109│                                                                                                
   │                                                                                                
 98│                                                                                                
   │                                                                                                
 87│                                                                                ●  ●  ●         
   │               ●  ●  ●   ●  ●                                                ●           ●      
 76│●  ●  ●  ●  ●                  ●                                      ●   ●                     
   │         ●                        ●                       ●  ●  ●  ●  ●   ●  ●  ●  ●        ●   
 65│            ●  ●  ●  ●                                 ●  ●  ●  ●                     ●        ●
   │                         ●  ●        ●           ●  ●                                    ●      
 55│                               ●        ●  ●  ●  ●                                          ●   
   │                                  ●           ●                                                ●
 44│                                     ●  ●  ●                                                    
   │                                                                                                
 33│                                                                                                
   │                                                                                                
 22│                                                                                                
   │                                                                                                
 11│                                                                                                
   │                                                                                                
  0└────────────────────────────────────────────────────────────────────────────────────────────────
   0 1   2 3   4 5   6   7 8   9 10  11  12  14  15  16  17  19  20  21  23  24  25  26  28  29  30 
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

Narrow · truecolor · 40 × 16 terminal cells

```
Chart / scatter                         
                                        
120│                                    
   │                                    
 90│      ●●●●●                 ●●● ●●  
   │●●●● ●●●●  ●●       ●● ●●●●●●●● ● ●●
 60│         ●●●● ●●●●●●             ●●●
   │              ●●●                   
 30│                                    
   │                                    
  0└────────────────────────────────────
   0 2 3 5 7 9 10  14  17  21  24  28 31
                                        
Ctrl+N advances fixture data / arrows pa
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Chart / scatter                                                                                     
                                                                                                    
120│                                                                                                
   │                                                                                                
109│                                                                                                
   │                                                                                                
 98│                                                                                                
   │                                                                                                
 87│                                                                                ●  ●  ●         
   │               ●  ●  ●   ●  ●                                                ●           ●      
 76│●  ●  ●  ●  ●                  ●                                      ●   ●                     
   │         ●                        ●                       ●  ●  ●  ●  ●   ●  ●  ●  ●        ●   
 65│            ●  ●  ●  ●                                 ●  ●  ●  ●                     ●        ●
   │                         ●  ●        ●           ●  ●                                    ●      
 55│                               ●        ●  ●  ●  ●                                          ●   
   │                                  ●           ●                                                ●
 44│                                     ●  ●  ●                                                    
   │                                                                                                
 33│                                                                                                
   │                                                                                                
 22│                                                                                                
   │                                                                                                
 11│                                                                                                
   │                                                                                                
  0└────────────────────────────────────────────────────────────────────────────────────────────────
   0 1   2 3   4 5   6   7 8   9 10  11  12  14  15  16  17  19  20  21  23  24  25  26  28  29  30 
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Chart / scatter                                                                                     
                                                                                                    
120│                                                                                                
   │                                                                                                
109│                                                                                                
   │                                                                                                
 98│                                                                                                
   │                                                                                                
 87│                                                                                ●  ●  ●         
   │               ●  ●  ●   ●  ●                                                ●           ●      
 76│●  ●  ●  ●  ●                  ●                                      ●   ●                     
   │         ●                        ●                       ●  ●  ●  ●  ●   ●  ●  ●  ●        ●   
 65│            ●  ●  ●  ●                                 ●  ●  ●  ●                     ●        ●
   │                         ●  ●        ●           ●  ●                                    ●      
 55│                               ●        ●  ●  ●  ●                                          ●   
   │                                  ●           ●                                                ●
 44│                                     ●  ●  ●                                                    
   │                                                                                                
 33│                                                                                                
   │                                                                                                
 22│                                                                                                
   │                                                                                                
 11│                                                                                                
   │                                                                                                
  0└────────────────────────────────────────────────────────────────────────────────────────────────
   0 1   2 3   4 5   6   7 8   9 10  11  12  14  15  16  17  19  20  21  23  24  25  26  28  29  30 
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

Monochrome · none · 100 × 30 terminal cells

```
scatter / text fallback                                                                             
                                                                                                    
2 series / latest 68.0                                                                              
                                                                                                    
ASCII mode presents values without chart glyphs.
```

### [Go usage](https://lanej.io/charmed/atoms-chart.html#chart%2fscatter-code)

Imports: [`canvas`](https://pkg.go.dev/github.com/NimbleMarkets/ntcharts/v2/canvas "github.com/NimbleMarkets/ntcharts/v2/canvas") · [`chart`](https://pkg.go.dev/github.com/lanej/charmed/chart "github.com/lanej/charmed/chart") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
values := []float64{72, 80, 68, 54, 66, 92, 90, 74}
plot := chart.NewLine(layout.Size{Width: 64, Height: 14}, 0, 7, 0, 100, palette)
plot.DrawXYAxisAndLabel()

for i, value := range values {
	point := canvas.Float64Point{X: float64(i), Y: value}
	plot.DrawRune(point, '●')
}

return plot.View()
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-chart.html#chart%2fscatter-integration)

Keep observations and their units in the application, and pass changes through the chart's public data API. Allocate terminal columns and rows from the parent layout. Forward native model commands when enabling interaction, and check whether resizing rebuilds plotted data before applying native overrides.

### [Run this example](https://lanej.io/charmed/atoms-chart.html#chart%2fscatter-run)

```sh
go run ./examples/catalog chart/scatter
```

[Example source](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/docs/snippets/chart_scatter.go) · [Implementation](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/story/atoms_charts.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Charts

## [Bars](https://lanej.io/charmed/atoms-chart.html#chart%2fbars)

Compare category magnitudes in a fixed terminal allocation. Supply meaningful labels and a consistent scale so the lengths represent comparable quantities.

### [Native preview](https://lanej.io/charmed/atoms-chart.html#chart%2fbars-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
                                ███████████████ 
                ▆▆▆▆▆▆▆▆▆▆▆▆▆▆▆ ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 
▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄ ▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁ ███████████████ 
███████████████ ███████████████ ███████████████ 
▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃ ███████████████ ███████████████ 
███████████████ ███████████████ ███████████████ 
███████████████ ███████████████ ███████████████ 
███████████████ ███████████████ ███████████████ 
███████████████ ███████████████ ███████████████ 
███████████████ ███████████████ ███████████████ 
────────────────────────────────────────────────
API             Events          Jobs
```

Demo · truecolor · 100 × 30 terminal cells

```
Chart / bars                                                                                        
                                                                                                    
                                        ███████████████████                                         
                                        ███████████████████                                         
                                        ███████████████████                                         
                                        ███████████████████                                         
                                        ███████████████████                                         
                    ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄ ███████████████████ ███████████████████                     
                    ███████████████████ ███████████████████ ███████████████████                     
                    ███████████████████ ███████████████████ ███████████████████                     
                    ███████████████████ ███████████████████ ███████████████████                     
                    ███████████████████ ███████████████████ ▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃                     
▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂ ███████████████████ ███████████████████ ███████████████████ ▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
────────────────────────────────────────────────────────────────────────────────────────────────────
API                 Events              Jobs                Storage             Schedule            
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

Narrow · truecolor · 40 × 16 terminal cells

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

256 colors · ansi256 · 100 × 30 terminal cells

```
Chart / bars                                                                                        
                                                                                                    
                                        ███████████████████                                         
                                        ███████████████████                                         
                                        ███████████████████                                         
                                        ███████████████████                                         
                                        ███████████████████                                         
                    ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄ ███████████████████ ███████████████████                     
                    ███████████████████ ███████████████████ ███████████████████                     
                    ███████████████████ ███████████████████ ███████████████████                     
                    ███████████████████ ███████████████████ ███████████████████                     
                    ███████████████████ ███████████████████ ▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃                     
▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂ ███████████████████ ███████████████████ ███████████████████ ▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
────────────────────────────────────────────────────────────────────────────────────────────────────
API                 Events              Jobs                Storage             Schedule            
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Chart / bars                                                                                        
                                                                                                    
                                        ███████████████████                                         
                                        ███████████████████                                         
                                        ███████████████████                                         
                                        ███████████████████                                         
                                        ███████████████████                                         
                    ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄ ███████████████████ ███████████████████                     
                    ███████████████████ ███████████████████ ███████████████████                     
                    ███████████████████ ███████████████████ ███████████████████                     
                    ███████████████████ ███████████████████ ███████████████████                     
                    ███████████████████ ███████████████████ ▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃▃                     
▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂ ███████████████████ ███████████████████ ███████████████████ ▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
███████████████████ ███████████████████ ███████████████████ ███████████████████ ███████████████████ 
────────────────────────────────────────────────────────────────────────────────────────────────────
API                 Events              Jobs                Storage             Schedule            
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

Monochrome · none · 100 × 30 terminal cells

```
bars / text fallback                                                                                
                                                                                                    
2 series / latest 68.0                                                                              
                                                                                                    
ASCII mode presents values without chart glyphs.
```

### [Go usage](https://lanej.io/charmed/atoms-chart.html#chart%2fbars-code)

Imports: [`barchart`](https://pkg.go.dev/github.com/NimbleMarkets/ntcharts/v2/barchart "github.com/NimbleMarkets/ntcharts/v2/barchart") · [`chart`](https://pkg.go.dev/github.com/lanej/charmed/chart "github.com/lanej/charmed/chart") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

plot := chart.NewBars(layout.Size{Width: 48, Height: 12}, palette)
for i, name := range []string{"API", "Events", "Jobs"} {
	plot.Push(barchart.BarData{Label: name, Values: []barchart.BarValue{
		{Name: "Successful", Value: float64(30 + i*10), Style: palette.SeriesStyle(0)},
		{Name: "Retried", Value: float64(12 - i*3), Style: palette.SeriesStyle(1)},
	}})
}

plot.Draw()

return plot.View()
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-chart.html#chart%2fbars-integration)

Keep observations and their units in the application, and pass changes through the chart's public data API. Allocate terminal columns and rows from the parent layout. Forward native model commands when enabling interaction, and check whether resizing rebuilds plotted data before applying native overrides.

### [Run this example](https://lanej.io/charmed/atoms-chart.html#chart%2fbars-run)

```sh
go run ./examples/catalog chart/bars
```

[Example source](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/docs/snippets/chart_bars.go) · [Implementation](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/story/atoms_charts.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Charts

## [Heatmap](https://lanej.io/charmed/atoms-chart.html#chart%2fheatmap)

Show ordered intensity across a grid of observations. The sequential palette encodes magnitude independently of category identity; include numeric context when exact values matter.

### [Native preview](https://lanej.io/charmed/atoms-chart.html#chart%2fheatmap-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
                                                
                                                
                                                
                                                
                                                
                                                
                                                
                                                
                                                
                                              
```

Demo · truecolor · 100 × 30 terminal cells

```
Chart / heatmap                                                                                     
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

Narrow · truecolor · 40 × 16 terminal cells

```
Chart / heatmap                         
                                        
                                        
                                        
                                        
                                        
                                        
                                        
                                        
                                        
                                        
                                        
                                        
Ctrl+N advances fixture data / arrows pa
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Chart / heatmap                                                                                     
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Chart / heatmap                                                                                     
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

Monochrome · none · 100 × 30 terminal cells

```
heatmap / text fallback                                                                             
                                                                                                    
2 series / latest 68.0                                                                              
                                                                                                    
ASCII mode presents values without chart glyphs.
```

### [Go usage](https://lanej.io/charmed/atoms-chart.html#chart%2fheatmap-code)

Imports: [`math`](https://pkg.go.dev/math "math") · [`heatmap`](https://pkg.go.dev/github.com/NimbleMarkets/ntcharts/v2/heatmap "github.com/NimbleMarkets/ntcharts/v2/heatmap") · [`chart`](https://pkg.go.dev/github.com/lanej/charmed/chart "github.com/lanej/charmed/chart") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
plot := chart.NewHeatmap(layout.Size{Width: 48, Height: 12}, theme.New(theme.Dark),
	heatmap.WithValueRange(0, 100), heatmap.WithCellSize(1, 1))
plot.SetXYRange(0, 12, 0, 8)
plot.SetViewXYRange(0, 12, 0, 8)

for x := range 12 {
	for y := range 8 {
		value := 50 + 45*math.Sin(float64(x+y)*.4)
		plot.Push(heatmap.NewHeatPoint(float64(x), float64(y), value))
	}
}

plot.Draw()

return plot.View()
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-chart.html#chart%2fheatmap-integration)

Keep observations and their units in the application, and pass changes through the chart's public data API. Allocate terminal columns and rows from the parent layout. Forward native model commands when enabling interaction, and check whether resizing rebuilds plotted data before applying native overrides.

### [Run this example](https://lanej.io/charmed/atoms-chart.html#chart%2fheatmap-run)

```sh
go run ./examples/catalog chart/heatmap
```

[Example source](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/docs/snippets/chart_heatmap.go) · [Implementation](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/story/atoms_charts.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Charts

## [Streamline](https://lanej.io/charmed/atoms-chart.html#chart%2fstream)

Show the most recent numeric samples as a scrolling trace. Use this for a live trend whose recent shape matters more than a complete historical time axis.

### [Native preview](https://lanej.io/charmed/atoms-chart.html#chart%2fstream-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
   │                                            
 83│   ┌──┐                            ┌─┐      
   │┌──┘  └─┐                       ┌──┘ └──┐   
 67├┘       └─┐                   ┌─┘       └┐  
   │          └─┐                ┌┘          └─┐
 50│            └┐             ┌─┘             └
   │             └┐           ┌┘                
 33│              └─┐       ┌─┘                 
   │                └──┐ ┌──┘                   
 17│                   └─┘                      
   │                                            
  0│
```

Demo · truecolor · 100 × 30 terminal cells

```
Chart / stream                                                                                      
                                                                                                    
   │                                                                                                
110│                                                                                                
   │                                                                                                
100│                                                                                                
   │                                                                                                
 90│                                                                                                
   │                      ┌──┐               ┌─┐                                        ┌───┐       
 80│  ┌───┐              ┌┘  └┐             ┌┘ └─┐                  ┌────┐             ┌┘   │       
   ├──┘   └┐            ┌┘    │            ┌┘    └─┐           ┌────┘    └┐            │    └┐      
 70│       │          ┌─┘     └┐           │       └──┐       ┌┘          │          ┌─┘     └┐     
   │       └┐      ┌──┘        └┐         ┌┘          └┐      │           └┐     ┌───┘        └┐    
 60│        └┐   ┌─┘            └─────┐  ┌┘            └┐    ┌┘            └┐   ┌┘             └────
   │         └┐ ┌┘                    └──┘              └┐  ┌┘              └┐ ┌┘                   
 50│          └─┘                                        └──┘                └─┘                    
   │                                                                                                
 40│                                                                                                
   │                                                                                                
 30│                                                                                                
   │                                                                                                
 20│                                                                                                
   │                                                                                                
 10│                                                                                                
   │                                                                                                
  0│                                                                                                
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

Narrow · truecolor · 40 × 16 terminal cells

```
Chart / stream                          
                                        
   │                                    
 96│                                    
   │  ┌───┐              ┌───┐          
 72├──┘   └─┐         ┌──┘   └─┐        
   │        └┐  ┌─────┘        └────────
 48│         └──┘                       
   │                                    
 24│                                    
   │                                    
  0│                                    
                                        
Ctrl+N advances fixture data / arrows pa
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Chart / stream                                                                                      
                                                                                                    
   │                                                                                                
110│                                                                                                
   │                                                                                                
100│                                                                                                
   │                                                                                                
 90│                                                                                                
   │                      ┌──┐               ┌─┐                                        ┌───┐       
 80│  ┌───┐              ┌┘  └┐             ┌┘ └─┐                  ┌────┐             ┌┘   │       
   ├──┘   └┐            ┌┘    │            ┌┘    └─┐           ┌────┘    └┐            │    └┐      
 70│       │          ┌─┘     └┐           │       └──┐       ┌┘          │          ┌─┘     └┐     
   │       └┐      ┌──┘        └┐         ┌┘          └┐      │           └┐     ┌───┘        └┐    
 60│        └┐   ┌─┘            └─────┐  ┌┘            └┐    ┌┘            └┐   ┌┘             └────
   │         └┐ ┌┘                    └──┘              └┐  ┌┘              └┐ ┌┘                   
 50│          └─┘                                        └──┘                └─┘                    
   │                                                                                                
 40│                                                                                                
   │                                                                                                
 30│                                                                                                
   │                                                                                                
 20│                                                                                                
   │                                                                                                
 10│                                                                                                
   │                                                                                                
  0│                                                                                                
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Chart / stream                                                                                      
                                                                                                    
   │                                                                                                
110│                                                                                                
   │                                                                                                
100│                                                                                                
   │                                                                                                
 90│                                                                                                
   │                      ┌──┐               ┌─┐                                        ┌───┐       
 80│  ┌───┐              ┌┘  └┐             ┌┘ └─┐                  ┌────┐             ┌┘   │       
   ├──┘   └┐            ┌┘    │            ┌┘    └─┐           ┌────┘    └┐            │    └┐      
 70│       │          ┌─┘     └┐           │       └──┐       ┌┘          │          ┌─┘     └┐     
   │       └┐      ┌──┘        └┐         ┌┘          └┐      │           └┐     ┌───┘        └┐    
 60│        └┐   ┌─┘            └─────┐  ┌┘            └┐    ┌┘            └┐   ┌┘             └────
   │         └┐ ┌┘                    └──┘              └┐  ┌┘              └┐ ┌┘                   
 50│          └─┘                                        └──┘                └─┘                    
   │                                                                                                
 40│                                                                                                
   │                                                                                                
 30│                                                                                                
   │                                                                                                
 20│                                                                                                
   │                                                                                                
 10│                                                                                                
   │                                                                                                
  0│                                                                                                
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

Monochrome · none · 100 × 30 terminal cells

```
stream / text fallback                                                                              
                                                                                                    
2 series / latest 68.0                                                                              
                                                                                                    
ASCII mode presents values without chart glyphs.
```

### [Go usage](https://lanej.io/charmed/atoms-chart.html#chart%2fstream-code)

Imports: [`math`](https://pkg.go.dev/math "math") · [`streamlinechart`](https://pkg.go.dev/github.com/NimbleMarkets/ntcharts/v2/linechart/streamlinechart "github.com/NimbleMarkets/ntcharts/v2/linechart/streamlinechart") · [`chart`](https://pkg.go.dev/github.com/lanej/charmed/chart "github.com/lanej/charmed/chart") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
plot := chart.NewStream(layout.Size{Width: 48, Height: 12}, theme.New(theme.Dark),
	streamlinechart.WithYRange(0, 100))
for i := range 48 {
	plot.Push(50 + 30*math.Sin(float64(i)*.2))
}

plot.DrawAll()

return plot.View()
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-chart.html#chart%2fstream-integration)

Keep observations and their units in the application, and pass changes through the chart's public data API. Allocate terminal columns and rows from the parent layout. Forward native model commands when enabling interaction, and check whether resizing rebuilds plotted data before applying native overrides.

### [Run this example](https://lanej.io/charmed/atoms-chart.html#chart%2fstream-run)

```sh
go run ./examples/catalog chart/stream
```

[Example source](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/docs/snippets/chart_stream.go) · [Implementation](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/story/atoms_charts.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Charts

## [Waveline](https://lanej.io/charmed/atoms-chart.html#chart%2fwave)

Draw a waveform over a numeric range with native line geometry. Supply the observed points explicitly and keep the range stable when comparing successive frames.

### [Native preview](https://lanej.io/charmed/atoms-chart.html#chart%2fwave-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
100│                                            
   │                                            
 80│      ┌┐┌──┐┌─┐                             
   │   ┌┐┌┘││  ││ │┌──┐                         
 60│┌┐┌┘││ ││  ││ ││  │┌┐                       
   ├┘││ ││ ││  ││ ││  ││└┐┌┐                   ┌
 40│ ││ ││ ││  ││ ││  ││ ││└┐┌┐             ┌┐┌┘
   │ ││ ││ ││  ││ ││  ││ ││ ││└─┐        ┌┐┌┘││ 
 20│ ││ ││ ││  ││ ││  ││ ││ ││  │┌─┐┌──┐┌┘││ ││ 
   │ ││ ││ ││  ││ ││  ││ ││ ││  ││ ││  ││ ││ ││ 
  0└─┴┴─┴┴─┴┴──┴┴─┴┴──┴┴─┴┴─┴┴──┴┴─┴┴──┴┴─┴┴─┴┴─
   0 1 3 4 6 7 8 10  13  16  18  21  24  27  30
```

Demo · truecolor · 100 × 30 terminal cells

```
Chart / wave                                                                                        
                                                                                                    
120│                                                                                                
   │                                                                                                
109│                                                                                                
   │                                                                                                
 98│                                                                                                
   │                                                                                                
 87│                                                                                ┌┐ ┌┐           
   │                  ┌┐ ┌┐ ┌┐                                                  ┌┐  ││ ││ ┌┐        
 76│  ┌┐ ┌┐ ┌┐ ┌┐ ┌┐  ││ ││ ││ ┌┐ ┌┐                                         ┌┐ ││  ││ ││ ││ ┌┐     
   ├┐ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││                                      ┌┐ ││ ││  ││ ││ ││ ││ ┌┐  
 65││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ┌┐                       ┌┐ ┌┐ ┌┐ ┌┐ ││ ││ ││  ││ ││ ││ ││ ││ ┌
   ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ┌┐              ┌┐ ┌┐ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
 55││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ┌┐    ┌┐  ┌┐ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
   ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ┌┐ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
 44││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
   ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
 33││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
   ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
 22││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
   ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
 11││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
   ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
  0└┴─┴┴─┴┴─┴┴─┴┴─┴┴──┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴──┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴──┴┴─┴┴─┴┴─┴┴─┴┴─┴
   0 1   2 3   4 5   6   7 8   9 10  11  12  14  15  16  17  19  20  21  23  24  25  26  28  29  30 
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

Narrow · truecolor · 40 × 16 terminal cells

```
Chart / wave                            
                                        
120│                                    
   │                                    
 90│                             ┌─┐┌┐  
   ├───┐┌─────┐┌┐            ┌───┘ ││└─┐
 60│   ││     ││└──┐  ┌─────┐│     ││  └
   │   ││     ││   └─┐│     ││     ││   
 30│   ││     ││     ││     ││     ││   
   │   ││     ││     ││     ││     ││   
  0└───┴┴─────┴┴─────┴┴─────┴┴─────┴┴───
   0 2 3 5 7 9 10  14  17  21  24  28 31
                                        
Ctrl+N advances fixture data / arrows pa
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Chart / wave                                                                                        
                                                                                                    
120│                                                                                                
   │                                                                                                
109│                                                                                                
   │                                                                                                
 98│                                                                                                
   │                                                                                                
 87│                                                                                ┌┐ ┌┐           
   │                  ┌┐ ┌┐ ┌┐                                                  ┌┐  ││ ││ ┌┐        
 76│  ┌┐ ┌┐ ┌┐ ┌┐ ┌┐  ││ ││ ││ ┌┐ ┌┐                                         ┌┐ ││  ││ ││ ││ ┌┐     
   ├┐ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││                                      ┌┐ ││ ││  ││ ││ ││ ││ ┌┐  
 65││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ┌┐                       ┌┐ ┌┐ ┌┐ ┌┐ ││ ││ ││  ││ ││ ││ ││ ││ ┌
   ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ┌┐              ┌┐ ┌┐ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
 55││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ┌┐    ┌┐  ┌┐ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
   ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ┌┐ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
 44││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
   ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
 33││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
   ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
 22││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
   ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
 11││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
   ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
  0└┴─┴┴─┴┴─┴┴─┴┴─┴┴──┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴──┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴──┴┴─┴┴─┴┴─┴┴─┴┴─┴
   0 1   2 3   4 5   6   7 8   9 10  11  12  14  15  16  17  19  20  21  23  24  25  26  28  29  30 
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Chart / wave                                                                                        
                                                                                                    
120│                                                                                                
   │                                                                                                
109│                                                                                                
   │                                                                                                
 98│                                                                                                
   │                                                                                                
 87│                                                                                ┌┐ ┌┐           
   │                  ┌┐ ┌┐ ┌┐                                                  ┌┐  ││ ││ ┌┐        
 76│  ┌┐ ┌┐ ┌┐ ┌┐ ┌┐  ││ ││ ││ ┌┐ ┌┐                                         ┌┐ ││  ││ ││ ││ ┌┐     
   ├┐ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││                                      ┌┐ ││ ││  ││ ││ ││ ││ ┌┐  
 65││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ┌┐                       ┌┐ ┌┐ ┌┐ ┌┐ ││ ││ ││  ││ ││ ││ ││ ││ ┌
   ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ┌┐              ┌┐ ┌┐ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
 55││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ┌┐    ┌┐  ┌┐ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
   ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ┌┐ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
 44││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
   ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
 33││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
   ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
 22││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
   ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
 11││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
   ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ ││ ││ ││ ││ ││  ││ ││ ││ ││ ││ │
  0└┴─┴┴─┴┴─┴┴─┴┴─┴┴──┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴──┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴─┴┴──┴┴─┴┴─┴┴─┴┴─┴┴─┴
   0 1   2 3   4 5   6   7 8   9 10  11  12  14  15  16  17  19  20  21  23  24  25  26  28  29  30 
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

Monochrome · none · 100 × 30 terminal cells

```
wave / text fallback                                                                                
                                                                                                    
2 series / latest 68.0                                                                              
                                                                                                    
ASCII mode presents values without chart glyphs.
```

### [Go usage](https://lanej.io/charmed/atoms-chart.html#chart%2fwave-code)

Imports: [`math`](https://pkg.go.dev/math "math") · [`canvas`](https://pkg.go.dev/github.com/NimbleMarkets/ntcharts/v2/canvas "github.com/NimbleMarkets/ntcharts/v2/canvas") · [`wavelinechart`](https://pkg.go.dev/github.com/NimbleMarkets/ntcharts/v2/linechart/wavelinechart "github.com/NimbleMarkets/ntcharts/v2/linechart/wavelinechart") · [`chart`](https://pkg.go.dev/github.com/lanej/charmed/chart "github.com/lanej/charmed/chart") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
plot := chart.NewWave(layout.Size{Width: 48, Height: 12}, theme.New(theme.Dark),
	wavelinechart.WithXYRange(0, 31, 0, 100))
for x := range 32 {
	plot.PlotDataSet("observed", canvas.Float64Point{
		X: float64(x), Y: 50 + 30*math.Sin(float64(x)*.2),
	})
}

plot.DrawAll()

return plot.View()
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-chart.html#chart%2fwave-integration)

Keep observations and their units in the application, and pass changes through the chart's public data API. Allocate terminal columns and rows from the parent layout. Forward native model commands when enabling interaction, and check whether resizing rebuilds plotted data before applying native overrides.

### [Run this example](https://lanej.io/charmed/atoms-chart.html#chart%2fwave-run)

```sh
go run ./examples/catalog chart/wave
```

[Example source](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/docs/snippets/chart_wave.go) · [Implementation](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/story/atoms_charts.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Charts

## [OHLC candles](https://lanej.io/charmed/atoms-chart.html#chart%2fohlc)

Compare open, high, low and close observations over an ordered series. The native candle renderer preserves the supplied measurements; the application owns their units and interpretation.

### [Native preview](https://lanej.io/charmed/atoms-chart.html#chart%2fohlc-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
84│                                                             
  │                    │                                        
80│                    │                   │                    
  ││                   ╽                   ╽                    
75││                   ┃                   ┃                   ╷
  │╽                   ┃                   ┃                   ╽
71│┃                   │                   ┃                   ┃
  │╿                   │                   ┃                   ┃
67││                   │                   │                   │
  ││                                       ╵                   │
62││                                                           ╵
  │                                                             
58└─────────────────────────────────────────────────────────────
  12:00:00 12:00:26 12:00:53 12:01:19 12:01:46 12:02:12 12:02:39
```

Demo · truecolor · 100 × 30 terminal cells

```
Chart / ohlc                                                                                        
                                                                                                    
103│                                                                                                
   │                                                                                                
 97│                                                                                ╷  ╷            
   │                                                                             ╷  │  │  ╷         
 91│                  ╷  │  │                                                    │  ┃  ┃  │         
   │      ╷  ╷  ╷  ╷  │  │  ╽  │                                              ╷  ╽  ┃  ┃  ┃  │      
 85│   │  │  │  │  │  ╽  ┃  ┃  ╽  ╷                                           │  ┃  │  │  ╿  │      
   ││  ╽  ┃  ╽  ╽  ┃  ┃  ╿  ╿  ┃  │                                        ╷  ┃  │  │  │  │  ┃  ╷   
 79│╽  ┃  ┃  ┃  ┃  ┃  │  │  │  │  ┃                                     ╷  │  ┃  │  ╵     │  ╿  │   
   │┃  │  │  │  │  │  │  │  │  │  ┃   │                       ╷  ╷   │  │  ┃  │  ╵           │  ┃  ╷
 73││  │  │  │  │  │  ╵        ╵  │   ╽                    ╷  │  │   │  ┃  ╿  │              │  ┃  │
   ││  ╵  ╵  ╵  ╵                 │   ┃  ╷              ╷  │  ╽  ┃   ┃  ┃  │  ╵                 │  ┃
 67│╵                             ╵   ╿  │           ╷  │  ┃  ┃  ┃   ╿  │  │                    │  ┃
   │                                  │  ┃  │     ╷  │  ┃  ┃  │  │   │  │                          │
 61│                                  ╵  ┃  │  │  │  ╽  ┃  │  │  │   │  ╵                          │
   │                                     │  ┃  ╽  ╽  ┃  │  │  ╵                                     
 55│                                     │  ╿  ┃  ┃  │  │                                           
   │                                        │  ╿  │  │                                              
 49│                                        │  │  │  ╵                                              
   │                                           ╵  ╵                                                 
 43│                                                                                                
   │                                                                                                
 36└────────────────────────────────────────────────────────────────────────────────────────────────
   12:00:00 12:02:54 12:05:48 12:08:43 12:11:37 12:14:31 12:17:26 12:20:20 12:23:15 12:26:09        
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

Narrow · truecolor · 40 × 16 terminal cells

```
Chart / ohlc                            
                                        
103│                              ╷╷    
   │      ╷ ╷╷╷                  │╽╽╽╷  
 87│╷╽╽╽╽╽┃ ┃┃┃│              ╷╽ ┃╿╿╿╽╷ 
   │┃╿╿╿╿╿│ │╵│┃╽       ╷ ╷╷│╽┃╿ ╵  ╵╿┃╷
 70│╵╵╵╵╵      ╵╿╽ ╷  ╷╽┃ ┃┃┃╿│╵      │┃
   │            ╵╿ ┃╽╽┃╿│ │╵╵╵         ╵
 53│             ╵ │╿╿│╵                
   │                ╵                   
 36└────────────────────────────────────
   12:00:00 12:07:45 12:15:30 12:23:15  
                                        
Ctrl+N advances fixture data / arrows pa
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Chart / ohlc                                                                                        
                                                                                                    
103│                                                                                                
   │                                                                                                
 97│                                                                                ╷  ╷            
   │                                                                             ╷  │  │  ╷         
 91│                  ╷  │  │                                                    │  ┃  ┃  │         
   │      ╷  ╷  ╷  ╷  │  │  ╽  │                                              ╷  ╽  ┃  ┃  ┃  │      
 85│   │  │  │  │  │  ╽  ┃  ┃  ╽  ╷                                           │  ┃  │  │  ╿  │      
   ││  ╽  ┃  ╽  ╽  ┃  ┃  ╿  ╿  ┃  │                                        ╷  ┃  │  │  │  │  ┃  ╷   
 79│╽  ┃  ┃  ┃  ┃  ┃  │  │  │  │  ┃                                     ╷  │  ┃  │  ╵     │  ╿  │   
   │┃  │  │  │  │  │  │  │  │  │  ┃   │                       ╷  ╷   │  │  ┃  │  ╵           │  ┃  ╷
 73││  │  │  │  │  │  ╵        ╵  │   ╽                    ╷  │  │   │  ┃  ╿  │              │  ┃  │
   ││  ╵  ╵  ╵  ╵                 │   ┃  ╷              ╷  │  ╽  ┃   ┃  ┃  │  ╵                 │  ┃
 67│╵                             ╵   ╿  │           ╷  │  ┃  ┃  ┃   ╿  │  │                    │  ┃
   │                                  │  ┃  │     ╷  │  ┃  ┃  │  │   │  │                          │
 61│                                  ╵  ┃  │  │  │  ╽  ┃  │  │  │   │  ╵                          │
   │                                     │  ┃  ╽  ╽  ┃  │  │  ╵                                     
 55│                                     │  ╿  ┃  ┃  │  │                                           
   │                                        │  ╿  │  │                                              
 49│                                        │  │  │  ╵                                              
   │                                           ╵  ╵                                                 
 43│                                                                                                
   │                                                                                                
 36└────────────────────────────────────────────────────────────────────────────────────────────────
   12:00:00 12:02:54 12:05:48 12:08:43 12:11:37 12:14:31 12:17:26 12:20:20 12:23:15 12:26:09        
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Chart / ohlc                                                                                        
                                                                                                    
103│                                                                                                
   │                                                                                                
 97│                                                                                ╷  ╷            
   │                                                                             ╷  │  │  ╷         
 91│                  ╷  │  │                                                    │  ┃  ┃  │         
   │      ╷  ╷  ╷  ╷  │  │  ╽  │                                              ╷  ╽  ┃  ┃  ┃  │      
 85│   │  │  │  │  │  ╽  ┃  ┃  ╽  ╷                                           │  ┃  │  │  ╿  │      
   ││  ╽  ┃  ╽  ╽  ┃  ┃  ╿  ╿  ┃  │                                        ╷  ┃  │  │  │  │  ┃  ╷   
 79│╽  ┃  ┃  ┃  ┃  ┃  │  │  │  │  ┃                                     ╷  │  ┃  │  ╵     │  ╿  │   
   │┃  │  │  │  │  │  │  │  │  │  ┃   │                       ╷  ╷   │  │  ┃  │  ╵           │  ┃  ╷
 73││  │  │  │  │  │  ╵        ╵  │   ╽                    ╷  │  │   │  ┃  ╿  │              │  ┃  │
   ││  ╵  ╵  ╵  ╵                 │   ┃  ╷              ╷  │  ╽  ┃   ┃  ┃  │  ╵                 │  ┃
 67│╵                             ╵   ╿  │           ╷  │  ┃  ┃  ┃   ╿  │  │                    │  ┃
   │                                  │  ┃  │     ╷  │  ┃  ┃  │  │   │  │                          │
 61│                                  ╵  ┃  │  │  │  ╽  ┃  │  │  │   │  ╵                          │
   │                                     │  ┃  ╽  ╽  ┃  │  │  ╵                                     
 55│                                     │  ╿  ┃  ┃  │  │                                           
   │                                        │  ╿  │  │                                              
 49│                                        │  │  │  ╵                                              
   │                                           ╵  ╵                                                 
 43│                                                                                                
   │                                                                                                
 36└────────────────────────────────────────────────────────────────────────────────────────────────
   12:00:00 12:02:54 12:05:48 12:08:43 12:11:37 12:14:31 12:17:26 12:20:20 12:23:15 12:26:09        
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

Monochrome · none · 100 × 30 terminal cells

```
ohlc / text fallback                                                                                
                                                                                                    
2 series / latest 68.0                                                                              
                                                                                                    
ASCII mode presents values without chart glyphs.
```

### [Go usage](https://lanej.io/charmed/atoms-chart.html#chart%2fohlc-code)

Imports: [`time`](https://pkg.go.dev/time "time") · [`chart`](https://pkg.go.dev/github.com/lanej/charmed/chart "github.com/lanej/charmed/chart") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
start := time.Date(2026, 11, 8, 12, 0, 0, 0, time.UTC)
data := []chart.Candle{
	{Time: start, Open: 68, High: 77, Low: 61, Close: 72},
	{Time: start.Add(time.Minute), Open: 72, High: 81, Low: 65, Close: 76},
	{Time: start.Add(2 * time.Minute), Open: 76, High: 79, Low: 64, Close: 67},
	{Time: start.Add(3 * time.Minute), Open: 67, High: 74, Low: 62, Close: 71},
}

plot, err := chart.NewCandles(data, layout.Size{Width: 64, Height: 14}, theme.New(theme.Dark))
if err != nil {
	return "", err
}

return plot.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-chart.html#chart%2fohlc-integration)

Keep observations and their units in the application, and pass changes through the chart's public data API. Allocate terminal columns and rows from the parent layout. Forward native model commands when enabling interaction, and check whether resizing rebuilds plotted data before applying native overrides.

### [Run this example](https://lanej.io/charmed/atoms-chart.html#chart%2fohlc-run)

```sh
go run ./examples/catalog chart/ohlc
```

[Example source](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/docs/snippets/chart_ohlc.go) · [Implementation](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/story/atoms_charts.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Charts

## [Canvas](https://lanej.io/charmed/atoms-chart.html#chart%2fcanvas)

Place styled text and individual marks at exact terminal-cell coordinates. Use higher-level charts for measured data unless custom geometry requires direct canvas access.

### [Native preview](https://lanej.io/charmed/atoms-chart.html#chart%2fcanvas-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
                                                
                                                
  CHARMED / ntcharts canvas                     
                                                
  *
```

Demo · truecolor · 100 × 30 terminal cells

```
Chart / canvas                                                                                      
                                                                                                    
·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      · 
      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·  
  CHARMED / ntcharts canvas      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·   
    ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·    
   ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·     
  ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      
 ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·
·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      · 
      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·  
     ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·   
    ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·    
   ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·     
  ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      
 ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·
·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      · 
      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·  
     ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·   
    ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·    
   ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·     
  ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      
 ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·
·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      · 
      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·  
     ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·   
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

Narrow · truecolor · 40 × 16 terminal cells

```
Chart / canvas                          
                                        
·      ·      ·      ·      ·      ·    
      ·      ·      ·      ·      ·     
  CHARMED / ntcharts canvas      ·      
    ·      ·      ·      ·      ·      ·
   ·      ·      ·      ·      ·      · 
  ·      ·      ·      ·      ·      ·  
 ·      ·      ·      ·      ·      ·   
·      ·      ·      ·      ·      ·    
      ·      ·      ·      ·      ·     
     ·      ·      ·      ·      ·      
                                        
Ctrl+N advances fixture data / arrows pa
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Chart / canvas                                                                                      
                                                                                                    
·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      · 
      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·  
  CHARMED / ntcharts canvas      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·   
    ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·    
   ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·     
  ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      
 ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·
·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      · 
      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·  
     ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·   
    ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·    
   ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·     
  ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      
 ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·
·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      · 
      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·  
     ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·   
    ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·    
   ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·     
  ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      
 ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·
·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      · 
      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·  
     ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·   
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Chart / canvas                                                                                      
                                                                                                    
·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      · 
      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·  
  CHARMED / ntcharts canvas      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·   
    ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·    
   ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·     
  ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      
 ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·
·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      · 
      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·  
     ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·   
    ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·    
   ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·     
  ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      
 ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·
·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      · 
      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·  
     ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·   
    ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·    
   ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·     
  ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      
 ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·
·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      · 
      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·  
     ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·      ·   
                                                                                                    
Ctrl+N advances fixture data / arrows pan supported plots
```

Monochrome · none · 100 × 30 terminal cells

```
canvas / text fallback                                                                              
                                                                                                    
2 series / latest 68.0                                                                              
                                                                                                    
ASCII mode presents values without chart glyphs.
```

### [Go usage](https://lanej.io/charmed/atoms-chart.html#chart%2fcanvas-code)

Imports: [`canvas`](https://pkg.go.dev/github.com/NimbleMarkets/ntcharts/v2/canvas "github.com/NimbleMarkets/ntcharts/v2/canvas") · [`chart`](https://pkg.go.dev/github.com/lanej/charmed/chart "github.com/lanej/charmed/chart") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
canvasView := chart.NewCanvas(layout.Size{Width: 48, Height: 8}, palette)
canvasView.SetStringWithStyle(canvas.Point{X: 2, Y: 2}, "CHARMED / ntcharts canvas", palette.Heading())
canvasView.SetRuneWithStyle(canvas.Point{X: 2, Y: 4}, '*', palette.SeriesStyle(0))

return canvasView.View()
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-chart.html#chart%2fcanvas-integration)

Keep observations and their units in the application, and pass changes through the chart's public data API. Allocate terminal columns and rows from the parent layout. Forward native model commands when enabling interaction, and check whether resizing rebuilds plotted data before applying native overrides.

### [Run this example](https://lanej.io/charmed/atoms-chart.html#chart%2fcanvas-run)

```sh
go run ./examples/catalog chart/canvas
```

[Example source](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/docs/snippets/chart_canvas.go) · [Implementation](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/story/atoms_charts.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Charts

## [Waffle](https://lanej.io/charmed/atoms-chart.html#chart%2fwaffle)

Equal-area category cells with deterministic rounding and exact labeled counts. Patterned squares and an ASCII fallback retain category identity.

### [Native preview](https://lanej.io/charmed/atoms-chart.html#chart%2fwaffle-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                  
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                  
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                  
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                  
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                  
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                  
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                  
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                  
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                  
■ ■ ▧ ▧ ▧ ▨ ▨ ▨ ▨ ▨                  
■ ready 92 / ▧ pending 3 / ▨ failed 5
```

Demo · truecolor · 80 × 16 terminal cells

```
Fleet health                                                                    
                                                                                
■ ■ ■ ■ ■ ■ ■ ■ ■ ■    100 instances                                            
■ ■ ■ ■ ■ ■ ■ ■ ■ ■    Production fleet                                         
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                                                             
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                                                             
■ ■ ■ ■ ■ ■ ■ ■ ■ ■    ■ ready 92  92%                                          
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                                                             
■ ■ ■ ■ ■ ■ ■ ■ ■ ■    ▧ pending 3  3%                                          
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                                                             
■ ■ ■ ■ ■ ■ ■ ■ ■ ■    ▨ failed 5  5%                                           
■ ■ ▧ ▧ ▧ ▨ ▨ ▨ ▨ ▨                                                             
                                                                                
100 cells · one instance per cell
```

Narrow · truecolor · 40 × 16 terminal cells

```
Fleet health / 100 instances            
                                        
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                     
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                     
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                     
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                     
■ ■ ■ ■ ■ ■ ▧ ▧ ▨ ▨                     
                                        
■ ready 92                              
▧ pending 3                             
▨ failed 5                              
                                        
50 cells · 2 instances each
```

256 colors · ansi256 · 80 × 16 terminal cells

```
Fleet health                                                                    
                                                                                
■ ■ ■ ■ ■ ■ ■ ■ ■ ■    100 instances                                            
■ ■ ■ ■ ■ ■ ■ ■ ■ ■    Production fleet                                         
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                                                             
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                                                             
■ ■ ■ ■ ■ ■ ■ ■ ■ ■    ■ ready 92  92%                                          
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                                                             
■ ■ ■ ■ ■ ■ ■ ■ ■ ■    ▧ pending 3  3%                                          
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                                                             
■ ■ ■ ■ ■ ■ ■ ■ ■ ■    ▨ failed 5  5%                                           
■ ■ ▧ ▧ ▧ ▨ ▨ ▨ ▨ ▨                                                             
                                                                                
100 cells · one instance per cell
```

16 colors · ansi16 · 80 × 16 terminal cells

```
Fleet health                                                                    
                                                                                
■ ■ ■ ■ ■ ■ ■ ■ ■ ■    100 instances                                            
■ ■ ■ ■ ■ ■ ■ ■ ■ ■    Production fleet                                         
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                                                             
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                                                             
■ ■ ■ ■ ■ ■ ■ ■ ■ ■    ■ ready 92  92%                                          
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                                                             
■ ■ ■ ■ ■ ■ ■ ■ ■ ■    ▧ pending 3  3%                                          
■ ■ ■ ■ ■ ■ ■ ■ ■ ■                                                             
■ ■ ■ ■ ■ ■ ■ ■ ■ ■    ▨ failed 5  5%                                           
■ ■ ▧ ▧ ▧ ▨ ▨ ▨ ▨ ▨                                                             
                                                                                
100 cells · one instance per cell
```

Monochrome · none · 80 × 16 terminal cells

```
Fleet health                                                                    
                                                                                
A A A A A A A A A A    100 instances                                            
A A A A A A A A A A    Production fleet                                         
A A A A A A A A A A                                                             
A A A A A A A A A A                                                             
A A A A A A A A A A    A ready 92  92%                                          
A A A A A A A A A A                                                             
A A A A A A A A A A    B pending 3  3%                                          
A A A A A A A A A A                                                             
A A A A A A A A A A    C failed 5  5%                                           
A A B B B C C C C C                                                             
                                                                                
100 cells * one instance per cell
```

### [Go usage](https://lanej.io/charmed/atoms-chart.html#chart%2fwaffle-code)

Imports: [`chart`](https://pkg.go.dev/github.com/lanej/charmed/chart "github.com/lanej/charmed/chart") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
fleet := chart.Waffle{Parts: []chart.WafflePart{
	{Label: "ready", Count: 92, Tone: text.Success},
	{Label: "pending", Count: 3, Tone: text.Info},
	{Label: "failed", Count: 5, Tone: text.Danger},
}}

return layout.Column(fleet.View(layout.Size{Width: 19, Height: 10}, palette), fleet.Legend(palette))
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-chart.html#chart%2fwaffle-integration)

Keep observations and their units in the application, and pass changes through the chart's public data API. Allocate terminal columns and rows from the parent layout. Forward native model commands when enabling interaction, and check whether resizing rebuilds plotted data before applying native overrides.

### [Run this example](https://lanej.io/charmed/atoms-chart.html#chart%2fwaffle-run)

```sh
go run ./examples/catalog chart/waffle
```

[Example source](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/docs/snippets/chart_waffle.go) · [Implementation](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/chart/waffle.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / Charts

## [Interval band](https://lanej.io/charmed/atoms-chart.html#chart%2finterval-band)

Nested caller-supplied intervals, a mean and selected observation rendered on an ntcharts canvas.

### [Native preview](https://lanej.io/charmed/atoms-chart.html#chart%2finterval-band-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
················░░░░░▒▒▒▒▒▓▓▓▓▓▓|▓▓▓▓▓▒▒▒▒▒░░░░░·X··············
```

Demo · truecolor · 100 × 30 terminal cells

```
Nested intervals / +/- 3s                                                                           
                                                                                                    
··························░░░░░░░░▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓|▓▓▓▓▓▓▓▒▒▒▒▒▒▒▒░░░░░░░░···X······················
0                                                                                                100
```

Narrow · truecolor · 40 × 16 terminal cells

```
Nested intervals / +/- 3s               
                                        
··········░░░▒▒▒▓▓▓▓|▓▓▓▒▒▒░░░X·········
0                                    100
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Nested intervals / +/- 3s                                                                           
                                                                                                    
··························░░░░░░░░▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓|▓▓▓▓▓▓▓▒▒▒▒▒▒▒▒░░░░░░░░···X······················
0                                                                                                100
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Nested intervals / +/- 3s                                                                           
                                                                                                    
··························░░░░░░░░▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓|▓▓▓▓▓▓▓▒▒▒▒▒▒▒▒░░░░░░░░···X······················
0                                                                                                100
```

Monochrome · none · 100 × 30 terminal cells

```
Nested intervals / +/- 3s                                                                           
                                                                                                    
..................................========########|#######========...........X......................
0                                                                                                100
```

With legend · truecolor · 100 × 30 terminal cells

```
Nested intervals / +/- 3s                                                                           
                                                                                                    
··························░░░░░░░░▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓|▓▓▓▓▓▓▓▒▒▒▒▒▒▒▒░░░░░░░░···X······················
0                                                                                                100
                                                                                                    
| mean  X selected observation
```

### [Go usage](https://lanej.io/charmed/atoms-chart.html#chart%2finterval-band-code)

Imports: [`chart`](https://pkg.go.dev/github.com/lanej/charmed/chart "github.com/lanej/charmed/chart") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
selected := 78.0
band := chart.BandOptions{
	Min: 0, Max: 100, Center: 50, Selected: &selected,
	Intervals: []chart.Interval{
		{Low: 26, High: 74}, {Low: 34, High: 66}, {Low: 42, High: 58},
	},
}

return chart.IntervalBand(band, 64, theme.New(theme.Dark))
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-chart.html#chart%2finterval-band-integration)

Keep observations and their units in the application, and pass changes through the chart's public data API. Allocate terminal columns and rows from the parent layout. Forward native model commands when enabling interaction, and check whether resizing rebuilds plotted data before applying native overrides.

### [Run this example](https://lanej.io/charmed/atoms-chart.html#chart%2finterval-band-run)

```sh
go run ./examples/catalog chart/interval-band
```

[Example source](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/docs/snippets/chart_interval_band.go) · [Implementation](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/chart/band.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.
