[Organisms](https://lanej.io/charmed/organisms.html)Patterns

# [Patterns](https://lanej.io/charmed/organisms-pattern.html#level-organisms-pattern)

1 component · [All organisms](https://lanej.io/charmed/organisms.html)

- [Measurement history](https://lanej.io/charmed/organisms-pattern.html#pattern%2fmeasurement-history)

[Organisms](https://lanej.io/charmed/organisms.html) / Patterns

## [Measurement history](https://lanej.io/charmed/organisms-pattern.html#pattern%2fmeasurement-history)

Inspect one member across supplied measurement windows and metric rows. Exact intervals and raw observations remain accessible independently of the compact value shown in each cell.

### [Native preview](https://lanej.io/charmed/organisms-pattern.html#pattern%2fmeasurement-history-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 Metric / Window      02:00:00       03:00:00       04:00:00       05:00:00     
 Temperature         Metric absent  Metric absent  Metric absent   Metric absent
 Energy              25kWh          26kWh          27kWh          28kWh         
 Yield               85%            86%            87%            88%           
 Humidity            39%            40%            41%            42%           
Rows 1-4/4  Columns 3-6/6                                                       
                                                                                
                                                                                
west / sensor-4                                                                 
Temperature                                                                     
2026-01-01T05:00:00Z – 2026-01-01T06:00:00Z                                     
Metric absent
```

Demo · truecolor · 100 × 30 terminal cells

```
 Metric / Window      04:00:00       05:00:00               west / sensor-4                         
 Temperature         Metric absent   Metric absent          Temperature                             
 Energy              27kWh          28kWh                   2026-01-01T05:00:00Z – 2026-01-01T06:00 
 Yield               87%            88%                     :00Z                                    
 Humidity            41%            42%                     Metric absent                           
Rows 1-4/4  Columns 5-6/6
```

Narrow · truecolor · 40 × 16 terminal cells

```
 Metric / Window      05:00:00          
 Temperature          Metric absent     
 Energy              28kWh              
 Yield               88%                
 Humidity            42%                
Rows 1-4/4  Columns 6-6/6               
                                        
                                        
west / sensor-4                         
Temperature                             
2026-01-01T05:00:00Z – 2026-01-01T06:00 
:00Z                                    
Metric absent
```

256 colors · ansi256 · 100 × 30 terminal cells

```
 Metric / Window      04:00:00       05:00:00               west / sensor-4                         
 Temperature         Metric absent   Metric absent          Temperature                             
 Energy              27kWh          28kWh                   2026-01-01T05:00:00Z – 2026-01-01T06:00 
 Yield               87%            88%                     :00Z                                    
 Humidity            41%            42%                     Metric absent                           
Rows 1-4/4  Columns 5-6/6
```

16 colors · ansi16 · 100 × 30 terminal cells

```
 Metric / Window      04:00:00       05:00:00               west / sensor-4                         
 Temperature         Metric absent   Metric absent          Temperature                             
 Energy              27kWh          28kWh                   2026-01-01T05:00:00Z – 2026-01-01T06:00 
 Yield               87%            88%                     :00Z                                    
 Humidity            41%            42%                     Metric absent                           
Rows 1-4/4  Columns 5-6/6
```

Monochrome · none · 100 × 30 terminal cells

```
 Metric / Window      04:00:00       05:00:00               west / sensor-4                         
 Temperature         Metric absent   Metric absent          Temperature                             
 Energy              27kWh          28kWh                   2026-01-01T05:00:00Z - 2026-01-01T06:00 
 Yield               87%            88%                     :00Z                                    
 Humidity            41%            42%                     Metric absent                           
Rows 1-4/4  Columns 5-6/6
```

### [Go usage](https://lanej.io/charmed/organisms-pattern.html#pattern%2fmeasurement-history-code)

Imports: [`cohort`](https://pkg.go.dev/github.com/lanej/charmed/cohort "github.com/lanej/charmed/cohort") · [`fixture`](https://pkg.go.dev/github.com/lanej/charmed/examples/cohorts/fixture "github.com/lanej/charmed/examples/cohorts/fixture") · [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`pattern`](https://pkg.go.dev/github.com/lanej/charmed/pattern "github.com/lanej/charmed/pattern") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
history, err := pattern.NewMeasurementHistory(fixture.Data(), cohort.MemberKey{CohortID: "west", MemberID: "sensor-4"}, layout.Size{Width: 80, Height: 16}, theme.New(theme.Dark))
if err != nil {
	return "", err
}
// Temperature is missing for this sensor; other measurements remain observable.
// SetData copies and validates. SetMember resets detail scroll, preserving the coordinate.
// Parent owns Tab through history.Inspector.FocusDetail and forwards Update commands.
return history.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-pattern.html#pattern%2fmeasurement-history-integration)

Supply cohort data and a member key; construction and SetData validate and copy the observations. Surviving metric and window identities retain selection; removing a selected window chooses the latest supplied window. SetMember resets detail scrolling. Route focus through Inspector.FocusDetail and forward Update commands. Native grid customization survives refreshes; missing members, missing metrics and unavailable ratios remain distinct from scalar zero.

### [Run this example](https://lanej.io/charmed/organisms-pattern.html#pattern%2fmeasurement-history-run)

```sh
go run ./examples/catalog pattern/measurement-history
```

[Example source](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/pattern_measurement_history.go) · [Implementation](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/pattern/measurement_history.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

### Reusable component

[Measurement history dependency graph](https://lanej.io/charmed/dependencies.html#component-pattern%2Fmeasurement-history) · Partial map · [Source declaration](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/pattern/measurement_history.go#L36)

- [Grid inspector](https://lanej.io/charmed/dependencies.html#component-collection%2Fgrid-inspector) · composes · inspection

Consumer usage uses

- [Measurement history](https://lanej.io/charmed/dependencies.html#component-pattern%2Fmeasurement-history) · primary demonstration

[Example source](https://github.com/lanej/charmed/blob/a32ca8934068617468e6e1f29305319108b33f92/docs/snippets/pattern_measurement_history.go#L11)
