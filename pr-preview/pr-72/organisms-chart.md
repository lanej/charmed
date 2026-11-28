[Organisms](https://lanej.io/charmed/organisms.html)Charts

# [Charts](https://lanej.io/charmed/organisms-chart.html#level-organisms-chart)

1 component · [All organisms](https://lanej.io/charmed/organisms.html)

- [Comparison panel](https://lanej.io/charmed/organisms-chart.html#chart%2fcomparison-panel)

[Organisms](https://lanej.io/charmed/organisms.html) / Charts

## [Comparison panel](https://lanej.io/charmed/organisms-chart.html#chart%2fcomparison-panel)

A distribution comparison and native time-series history compose one reusable metric panel.

### [Native preview](https://lanej.io/charmed/organisms-chart.html#chart%2fcomparison-panel-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
┌ Response share ──────────────────────────────────────────────┐
│ API mean 72% s=6pp n=64                                      │
│ ····················░░░░░░▒▒▒▒▒▓▓▓▓▓▓|▓▓▓▓▓▓▒▒▒▒▒▒░░░░░░···· │
│ Worker mean 62% s=8pp n=32                                   │
│ ····░░░░░░░░▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓|▓▓▓▓▓▓▓▒▒▒▒▒▒▒▒░░░░░░░░········ │
│ 34.36%                                                93.64% │
│ 99│                                                          │
│   │                                       ⡠⠊⠉⠉⠒⠒⠒⠒⠤⠤⣀        │
│ 87│                                     ⢀⠜         ⢀ ⠑⠢⣀     │
│   │    ⢀⣀⠤⠔⠒⠤⣀                        ⢀⠔⠁       ⢀⠤⠊⠁⠉⠒⢄⡀⠑⠤⡀  │
│ 76│⠤⠒⠊⠉⠁      ⠉⠒⠤⣀                   ⡰⠁      ⢀⡠⠒⠁      ⠈⠒⠤⡈⠑ │
│   │     ⢀⣀⡠⠤⢄⣀⣀⡀  ⠉⠢⢄⡀            ⣀⠤⠊     ⣀⠤⠒⠁            ⠈⠑ │
│ 64│⠤⠔⠒⠊⠉⠁      ⠈⠉⠉⠒⠢⢄⡈⠑⠢⣀     ⣀⠤⠒⠉   ⢀⡠⠤⠒⠉                   │
│   │                  ⠈⠑⠢⣀⠉⠒⠤⠒⠉ ⢀⣀⠤⠤⠒⠊⠁                       │
│ 53│                      ⠉⠒⠤⠒⠒⠉⠁                             │
│   │                                                          │
│ 41└───────────────────────────────────────────────────────── │
│   12:00:00 12:01:06 12:02:12 12:03:18 12:04:25 12:05:31      │
│                                                              │
└──────────────────────────────────────────────────────────────┘
```

Demo · truecolor · 100 × 30 terminal cells

```
┌ Response share ──────────────────────────────────────────────────────────────────────────────────┐
│ a13c9f2 mean 20% s=3pp n=64                                                                      │
│ ···············░░░░░▒▒▒▒▒▓▓▓▓▓|▓▓▓▓▓▒▒▒▒▒▒░░░░░················································· │
│ f4d7e81 mean 30% s=8pp n=32                                                                      │
│ ······░░░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓|▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒▒X▒▒▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░······ │
│ 2.64%                                                                                     57.36% │
│ 57│                                                                                              │
│   │                                                                                              │
│ 51│⠒⠒⠒⠢⣀                                                                                         │
│   │     ⠑⢄                                                                                       │
│ 45│       ⠑⢄                                             ⢀⡠⠔⠒⠒⠒⠢⠤⣀⡀         ⢀⣀⣀⡀                 │
│   │         ⠑⢄                                         ⡠⠊⠁        ⠈⠉⠒⠒⠒⠒⠒⠊⠉⠉⠁  ⠈⠉⠒⢄⡀             │
│ 39│           ⠑⢄                                     ⢀⠔⠁                       ⢀⠔⠒⠊⠉⠲⣄⡀          │
│   │             ⠉⠢⣀     ⣀⣀⣀⣀                        ⡠⠃                       ⡠⠊⠁     ⠑⢌⠢⡀        │
│ 33│                ⣉⡱⠶⠒⠭⠤⣀⡀ ⠉⠒⢄                    ⡔⠁                      ⢀⠎          ⠣⡑⢄       │
│   │  ⣀⠤⠤⠒⠒⠒⠒⠒⠒⠒⠒⠒⠉⠉       ⠈⠉⠒⠢⢄⠱⡀                 ⡜                      ⢀⠔⠁            ⠑⡄⠣⡀     │
│ 27│⠔⠊                          ⠉⠪⣢               ⡜                      ⡠⠃               ⠈⢆⠑⡄    │
│   │                              ⠈⠳⡀           ⢀⠜                    ⢀⡠⠊                   ⠣⡈⢆   │
│ 20│                                ⠘⢦⣀        ⡰⠁                ⢀⣀⡠⠤⠒⠁                      ⠑⢄⠱⡀ │
│   │                                 ⠈⢢⠑⠤⠤⣀⣀⠤⠔⠊           ⢀⡠⠔⠒⠊⠉⠉⠁                             ⠑⢌ │
│ 14│                                   ⠱⡀              ⣀⠔⠊⠁                                       │
│   │                                    ⠘⢄          ⢀⡠⠊                                           │
│  8│                                      ⠣⣀      ⣀⠔⠁                                             │
│   │                                        ⠉⠒⠤⠤⠒⠊                                                │
│  2│                                                                                              │
│   │                                                                                              │
│ -4└───────────────────────────────────────────────────────────────────────────────────────────── │
│   12:00:00 12:03:00 12:06:00 12:09:00 12:12:00 12:15:00 12:18:00 12:21:00 12:24:00 12:27:00      │
│                                                                                                  │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
```

Narrow · truecolor · 40 × 16 terminal cells

```
┌ Response share ──────────────────────┐
│ a13c9f2 mean 20% s=3pp n=64          │
│ ·····░░▒▒▓▓|▓▓▒▒░░·················· │
│ f4d7e81 mean 30% s=8pp n=32          │
│ ··░░░░░▒▒▒▒▒▓▓▓▓▓▓|▓▓▓▓▓X▒▒▒▒░░░░░·· │
│ 2.64%                         57.36% │
│ 57│⣀⣀                                │
│   │  ⠑⢄              ⢀⠤⠒⠒⠤⠤⠤⠤⠤⢤⣄⡀    │
│ 37│ ⣀⣀⣀⣉⠶⠤⠴⠶⢦⡀      ⡰⠁      ⡠⠊⠁ ⠙⡦⡀  │
│   │⠉         ⠙⢦⡀ ⢀⠔⠉   ⣀⣀⣀⠤⠊     ⠈⠳⣄ │
│ 16│            ⠙⢍⡁ ⢀⡠⠒⠉            ⠈ │
│   │              ⠈⠉⠁                 │
│ -4└───────────────────────────────── │
│   12:00:00 12:08:27 12:16:54         │
│                                      │
└──────────────────────────────────────┘
```

256 colors · ansi256 · 100 × 30 terminal cells

```
┌ Response share ──────────────────────────────────────────────────────────────────────────────────┐
│ a13c9f2 mean 20% s=3pp n=64                                                                      │
│ ···············░░░░░▒▒▒▒▒▓▓▓▓▓|▓▓▓▓▓▒▒▒▒▒▒░░░░░················································· │
│ f4d7e81 mean 30% s=8pp n=32                                                                      │
│ ······░░░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓|▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒▒X▒▒▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░······ │
│ 2.64%                                                                                     57.36% │
│ 57│                                                                                              │
│   │                                                                                              │
│ 51│⠒⠒⠒⠢⣀                                                                                         │
│   │     ⠑⢄                                                                                       │
│ 45│       ⠑⢄                                             ⢀⡠⠔⠒⠒⠒⠢⠤⣀⡀         ⢀⣀⣀⡀                 │
│   │         ⠑⢄                                         ⡠⠊⠁        ⠈⠉⠒⠒⠒⠒⠒⠊⠉⠉⠁  ⠈⠉⠒⢄⡀             │
│ 39│           ⠑⢄                                     ⢀⠔⠁                       ⢀⠔⠒⠊⠉⠲⣄⡀          │
│   │             ⠉⠢⣀     ⣀⣀⣀⣀                        ⡠⠃                       ⡠⠊⠁     ⠑⢌⠢⡀        │
│ 33│                ⣉⡱⠶⠒⠭⠤⣀⡀ ⠉⠒⢄                    ⡔⠁                      ⢀⠎          ⠣⡑⢄       │
│   │  ⣀⠤⠤⠒⠒⠒⠒⠒⠒⠒⠒⠒⠉⠉       ⠈⠉⠒⠢⢄⠱⡀                 ⡜                      ⢀⠔⠁            ⠑⡄⠣⡀     │
│ 27│⠔⠊                          ⠉⠪⣢               ⡜                      ⡠⠃               ⠈⢆⠑⡄    │
│   │                              ⠈⠳⡀           ⢀⠜                    ⢀⡠⠊                   ⠣⡈⢆   │
│ 20│                                ⠘⢦⣀        ⡰⠁                ⢀⣀⡠⠤⠒⠁                      ⠑⢄⠱⡀ │
│   │                                 ⠈⢢⠑⠤⠤⣀⣀⠤⠔⠊           ⢀⡠⠔⠒⠊⠉⠉⠁                             ⠑⢌ │
│ 14│                                   ⠱⡀              ⣀⠔⠊⠁                                       │
│   │                                    ⠘⢄          ⢀⡠⠊                                           │
│  8│                                      ⠣⣀      ⣀⠔⠁                                             │
│   │                                        ⠉⠒⠤⠤⠒⠊                                                │
│  2│                                                                                              │
│   │                                                                                              │
│ -4└───────────────────────────────────────────────────────────────────────────────────────────── │
│   12:00:00 12:03:00 12:06:00 12:09:00 12:12:00 12:15:00 12:18:00 12:21:00 12:24:00 12:27:00      │
│                                                                                                  │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
```

16 colors · ansi16 · 100 × 30 terminal cells

```
┌ Response share ──────────────────────────────────────────────────────────────────────────────────┐
│ a13c9f2 mean 20% s=3pp n=64                                                                      │
│ ···············░░░░░▒▒▒▒▒▓▓▓▓▓|▓▓▓▓▓▒▒▒▒▒▒░░░░░················································· │
│ f4d7e81 mean 30% s=8pp n=32                                                                      │
│ ······░░░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓|▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒▒X▒▒▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░······ │
│ 2.64%                                                                                     57.36% │
│ 57│                                                                                              │
│   │                                                                                              │
│ 51│⠒⠒⠒⠢⣀                                                                                         │
│   │     ⠑⢄                                                                                       │
│ 45│       ⠑⢄                                             ⢀⡠⠔⠒⠒⠒⠢⠤⣀⡀         ⢀⣀⣀⡀                 │
│   │         ⠑⢄                                         ⡠⠊⠁        ⠈⠉⠒⠒⠒⠒⠒⠊⠉⠉⠁  ⠈⠉⠒⢄⡀             │
│ 39│           ⠑⢄                                     ⢀⠔⠁                       ⢀⠔⠒⠊⠉⠲⣄⡀          │
│   │             ⠉⠢⣀     ⣀⣀⣀⣀                        ⡠⠃                       ⡠⠊⠁     ⠑⢌⠢⡀        │
│ 33│                ⣉⡱⠶⠒⠭⠤⣀⡀ ⠉⠒⢄                    ⡔⠁                      ⢀⠎          ⠣⡑⢄       │
│   │  ⣀⠤⠤⠒⠒⠒⠒⠒⠒⠒⠒⠒⠉⠉       ⠈⠉⠒⠢⢄⠱⡀                 ⡜                      ⢀⠔⠁            ⠑⡄⠣⡀     │
│ 27│⠔⠊                          ⠉⠪⣢               ⡜                      ⡠⠃               ⠈⢆⠑⡄    │
│   │                              ⠈⠳⡀           ⢀⠜                    ⢀⡠⠊                   ⠣⡈⢆   │
│ 20│                                ⠘⢦⣀        ⡰⠁                ⢀⣀⡠⠤⠒⠁                      ⠑⢄⠱⡀ │
│   │                                 ⠈⢢⠑⠤⠤⣀⣀⠤⠔⠊           ⢀⡠⠔⠒⠊⠉⠉⠁                             ⠑⢌ │
│ 14│                                   ⠱⡀              ⣀⠔⠊⠁                                       │
│   │                                    ⠘⢄          ⢀⡠⠊                                           │
│  8│                                      ⠣⣀      ⣀⠔⠁                                             │
│   │                                        ⠉⠒⠤⠤⠒⠊                                                │
│  2│                                                                                              │
│   │                                                                                              │
│ -4└───────────────────────────────────────────────────────────────────────────────────────────── │
│   12:00:00 12:03:00 12:06:00 12:09:00 12:12:00 12:15:00 12:18:00 12:21:00 12:24:00 12:27:00      │
│                                                                                                  │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘
```

Monochrome · none · 100 × 30 terminal cells

```
+ Response share ----------------------------------------------------------------------------------+
| a13c9f2 mean 20% s=3pp n=64                                                                      |
| ....................=====#####|#####======...................................................... |
| f4d7e81 mean 30% s=8pp n=32                                                                      |
| ....................==============##############|#############===X==========.................... |
| 2.64%                                                                                     57.36% |
| a13c9f2: 17 %                                                                                    |
| f4d7e81: 14.9 %                                                                                  |
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

### [Go usage](https://lanej.io/charmed/organisms-chart.html#chart%2fcomparison-panel-code)

Imports: [`time`](https://pkg.go.dev/time "time") · [`timeserieslinechart`](https://pkg.go.dev/github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart "github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart") · [`chart`](https://pkg.go.dev/github.com/lanej/charmed/chart "github.com/lanej/charmed/chart") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
start := time.Date(2026, 11, 8, 12, 0, 0, 0, time.UTC)
series := []chart.TimeSeriesData{
	{ID: "api", Label: "API", Unit: "%"},
	{ID: "worker", Label: "Worker", Unit: "%"},
}

samples := [][]float64{{72, 80, 68, 54, 66, 92, 90, 74}, {60, 66, 62, 48, 56, 66, 82, 68}}
for i, values := range samples {
	for minute, value := range values {
		series[i].Points = append(series[i].Points, timeserieslinechart.TimePoint{
			Time: start.Add(time.Duration(minute) * time.Minute), Value: value,
		})
	}
}

distribution := chart.DistributionComparison{
	Sigma: 3, Min: 0, Max: 100, Unit: "%", SpreadUnit: "pp",
	Series: []chart.DistributionSeries{
		{Label: "API", Mean: 72, StdDev: 6, Members: 64, HasSpread: true},
		{Label: "Worker", Mean: 62, StdDev: 8, Members: 32, HasSpread: true},
	},
}
history := chart.NewTimeSeries(series, layout.Size{Width: 60, Height: 12}, palette)
panel := chart.ComparisonPanel{Title: "Response share", Distribution: distribution, History: history}

return panel.View(layout.Size{Width: 64, Height: 20}, palette, false)
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-chart.html#chart%2fcomparison-panel-integration)

Keep observations and their units in the application, and pass changes through the chart's public data API. Allocate terminal columns and rows from the parent layout. Forward native model commands when enabling interaction, and check whether resizing rebuilds plotted data before applying native overrides.

### [Run this example](https://lanej.io/charmed/organisms-chart.html#chart%2fcomparison-panel-run)

```sh
go run ./examples/catalog chart/comparison-panel
```

[Example source](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/docs/snippets/chart_comparison_panel.go) · [Implementation](https://github.com/lanej/charmed/blob/2cad1ed1ef75b255703cce4145709deb9858e3f4/story/organisms.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [chart/distributions](https://lanej.io/charmed/molecules-chart.html#chart/distributions) [chart/timeseries](https://lanej.io/charmed/atoms-chart.html#chart/timeseries) [foundations/layout](https://lanej.io/charmed/foundations-layout.html#foundations/layout)
