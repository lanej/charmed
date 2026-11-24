[Organisms](https://lanej.io/charmed/organisms.html)Forms

# [Forms](https://lanej.io/charmed/organisms-form.html#level-organisms-form)

1 component · [All organisms](https://lanej.io/charmed/organisms.html)

- [Settings](https://lanej.io/charmed/organisms-form.html#form%2fsettings)

[Organisms](https://lanej.io/charmed/organisms.html) / Forms

## [Settings](https://lanej.io/charmed/organisms-form.html#form%2fsettings)

A Huh form owns its nested focus and validation.

### [Native preview](https://lanej.io/charmed/organisms-form.html#form%2fsettings-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
  Environment                                   
  > production                                  
                                                
  Region                                        
  > US West                                     
    US East                                     
                                                
  Features                                      
    [•] Metrics                                 
                                                
                                                

enter next
```

Demo · truecolor · 100 × 30 terminal cells

```
  Environment                                                                                       
  > production                                                                                      
                                                                                                    
  Region                                                                                            
  > US West                                                                                         
    US East                                                                                         
                                                                                                    
  Features                                                                                          
    [•] Metrics                                                                                     
                                                                                                    
  Enable environment?                                                                               
                                                                                                    
      Yes     No  
```

Narrow · truecolor · 40 × 16 terminal cells

```
  Environment                           
  > production                          
                                        
  Region                                
  > US West                             
    US East                             
                                        
  Features                              
    [•] Metrics                         
                                        
  Enable environment?                   
                                        
      Yes     No  
```

256 colors · ansi256 · 100 × 30 terminal cells

```
  Environment                                                                                       
  > production                                                                                      
                                                                                                    
  Region                                                                                            
  > US West                                                                                         
    US East                                                                                         
                                                                                                    
  Features                                                                                          
    [•] Metrics                                                                                     
                                                                                                    
  Enable environment?                                                                               
                                                                                                    
      Yes     No  
```

16 colors · ansi16 · 100 × 30 terminal cells

```
  Environment                                                                                       
  > production                                                                                      
                                                                                                    
  Region                                                                                            
  > US West                                                                                         
    US East                                                                                         
                                                                                                    
  Features                                                                                          
    [•] Metrics                                                                                     
                                                                                                    
  Enable environment?                                                                               
                                                                                                    
      Yes     No  
```

Monochrome · none · 100 × 30 terminal cells

```
  Environment                                                                                       
  > production                                                                                      
                                                                                                    
  Region                                                                                            
  > US West                                                                                         
    US East                                                                                         
                                                                                                    
  Features                                                                                          
    [*] Metrics                                                                                     
                                                                                                    
  Enable environment?                                                                               
                                                                                                    
      Yes     No  
```

### [Go usage](https://lanej.io/charmed/organisms-form.html#form%2fsettings-code)

Imports: [`tea`](https://pkg.go.dev/charm.land/bubbletea/v2 "charm.land/bubbletea/v2") · [`huh`](https://pkg.go.dev/charm.land/huh/v2 "charm.land/huh/v2") · [`form`](https://pkg.go.dev/github.com/lanej/charmed/form "github.com/lanej/charmed/form") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
name, region := "production", "west"
features := []string{"metrics"}
settings := form.New(theme.New(theme.Dark), huh.NewGroup(
	huh.NewInput().Title("Environment").Value(&name),
	huh.NewSelect[string]().Title("Region").Options(
		huh.NewOption("US West", "west"), huh.NewOption("US East", "east"),
	).Value(&region),
	huh.NewMultiSelect[string]().Title("Features").Options(
		huh.NewOption("Metrics", "metrics"), huh.NewOption("Logs", "logs"),
	).Value(&features),
))
settings.WithWidth(48).WithHeight(12)
_, _ = settings.Update(tea.WindowSizeMsg{Width: 48, Height: 12})
// For a standalone interactive form, call settings.Run() and handle its error.
return settings.View()
```

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-form.html#form%2fsettings-integration)

The parent model supplies observations, coordinates focus and owns application decisions. Forward initialization and update commands rather than executing them during rendering. Resize the existing composition from the terminal allocation so selection remains stable, and apply native customization at the exposed child models after any rebuilding lifecycle step.

### [Run this example](https://lanej.io/charmed/organisms-form.html#form%2fsettings-run)

```sh
go run ./examples/catalog form/settings
```

[Example source](https://github.com/lanej/charmed/blob/27e05a1997640606888402c2223835e8d5e8e451/docs/snippets/form_settings.go) · [Implementation](https://github.com/lanej/charmed/blob/27e05a1997640606888402c2223835e8d5e8e451/story/organisms.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.
