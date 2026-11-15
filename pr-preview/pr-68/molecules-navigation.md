[Molecules](https://lanej.io/charmed/molecules.html)Navigation

# [Navigation](https://lanej.io/charmed/molecules-navigation.html#level-molecules-navigation)

1 component · [All molecules](https://lanej.io/charmed/molecules.html)

- [Tabs](https://lanej.io/charmed/molecules-navigation.html#navigation%2ftabs)

[Molecules](https://lanej.io/charmed/molecules.html) / Navigation

## [Tabs](https://lanej.io/charmed/molecules-navigation.html#navigation%2ftabs)

Focused arrow-key navigation with stable selection, visible overflow and shared selection styling.

### [Native preview](https://lanej.io/charmed/molecules-navigation.html#navigation%2ftabs-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
  Overview  Services  Activity
```

Demo · truecolor · 64 × 8 terminal cells

```
Tabs                                                            
                                                                
  Overview  Services  Activity  Settings                        
                                                                
Left / Right select / Home / End jump
```

Narrow · truecolor · 40 × 16 terminal cells

```
Tabs                                    
                                        
  Overview  Services  Activity         >
                                        
Left / Right select / Home / End jump
```

256 colors · ansi256 · 64 × 8 terminal cells

```
Tabs                                                            
                                                                
  Overview  Services  Activity  Settings                        
                                                                
Left / Right select / Home / End jump
```

16 colors · ansi16 · 64 × 8 terminal cells

```
Tabs                                                            
                                                                
  Overview  Services  Activity  Settings                        
                                                                
Left / Right select / Home / End jump
```

Monochrome · none · 64 × 8 terminal cells

```
Tabs                                                            
                                                                
  Overview  Services  Activity  Settings                        
                                                                
Left / Right select / Home / End jump
```

### [Go usage](https://lanej.io/charmed/molecules-navigation.html#navigation%2ftabs-code)

Imports: [`navigation`](https://pkg.go.dev/github.com/lanej/charmed/navigation "github.com/lanej/charmed/navigation") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
tabs := navigation.NewTabs([]navigation.Tab{
	{ID: "overview", Label: "Overview"},
	{ID: "services", Label: "Services"},
	{ID: "activity", Label: "Activity"},
}, theme.New(theme.Dark))
tabs.SetWidth(64)
tabs.Focus()
// Return tabs.Update(msg)'s command from Update; tabs.Selected() gives the active ID.
return tabs.View()
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-navigation.html#navigation%2ftabs-integration)

Route arrow keys only to the focused tab bar and return the command from Update so selection events reach the application. Use stable unique item IDs. Resize retains selection and native theme customization; the parent decides which content and focus region follow the chosen tab.

### [Run this example](https://lanej.io/charmed/molecules-navigation.html#navigation%2ftabs-run)

```sh
go run ./examples/catalog navigation/tabs
```

[Example source](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/docs/snippets/navigation_tabs.go) · [Implementation](https://github.com/lanej/charmed/blob/91c558bb8aca62050bbfc3f0781f1e9ffd13e33d/story/chrome.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [foundations/layout](https://lanej.io/charmed/foundations-layout.html#foundations/layout)
