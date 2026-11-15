[Molecules](https://lanej.io/charmed/molecules.html)Input

# [Input](https://lanej.io/charmed/molecules-input.html#level-molecules-input)

2 components · [All molecules](https://lanej.io/charmed/molecules.html)

- [Labeled field](https://lanej.io/charmed/molecules-input.html#input%2ffield)
- [Radio group](https://lanej.io/charmed/molecules-input.html#input%2fradio-group)

[Molecules](https://lanej.io/charmed/molecules.html) / Input

## [Labeled field](https://lanej.io/charmed/molecules-input.html#input%2ffield)

Label, help and validation composed around the same input atom.

### [Native preview](https://lanej.io/charmed/molecules-input.html#input%2ffield-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Environment
> production                            
Use at least three characters.
```

Demo · truecolor · 100 × 30 terminal cells

```
Environment name                                                                                    
> production                                                                                        
Use at least three characters.
```

Narrow · truecolor · 40 × 16 terminal cells

```
Environment name                        
> production                            
Use at least three characters.
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Environment name                                                                                    
> production                                                                                        
Use at least three characters.
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Environment name                                                                                    
> production                                                                                        
Use at least three characters.
```

Monochrome · none · 100 × 30 terminal cells

```
Environment name                                                                                    
> production                                                                                        
Use at least three characters.
```

### [Go usage](https://lanej.io/charmed/molecules-input.html#input%2ffield-code)

Imports: [`errors`](https://pkg.go.dev/errors "errors") · [`input`](https://pkg.go.dev/github.com/lanej/charmed/input "github.com/lanej/charmed/input") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
field := input.Field{
	Label: "Environment", Help: "Use at least three characters.", Theme: palette,
	Input: input.NewText(input.TextOptions{
		Value: "production", Width: 40, Theme: palette,
		Validate: func(value string) error {
			if len(value) < 3 {
				return errors.New("use at least three characters")
			}

			return nil
		},
	}),
}
// In Init, return field.Input.Focus(); in Update, forward messages to field.Input.
return field.View()
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-input.html#input%2ffield-integration)

Return Focus's command when the input becomes active. In Update, retain both the returned model and command; blur inactive controls before routing keys. Call Resize with the allocated column count. SetTheme reapplies shared styles, so apply native style overrides after a theme change.

### [Run this example](https://lanej.io/charmed/molecules-input.html#input%2ffield-run)

```sh
go run ./examples/catalog input/field
```

[Example source](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/docs/snippets/input_field.go) · [Implementation](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/story/molecules.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [input/text](https://lanej.io/charmed/atoms-input.html#input/text)

[Molecules](https://lanej.io/charmed/molecules.html) / Input

## [Radio group](https://lanej.io/charmed/molecules-input.html#input%2fradio-group)

Choose exactly one option while preserving a separate keyboard focus position. Disabled options remain visible with their explanation and are skipped during internal navigation.

### [Native preview](https://lanej.io/charmed/molecules-input.html#input%2fradio-group-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
  ◉  Last transition 
  ○  Last three transitions 
  ○  Complete trace 
  ○  Live execution ⊘ disabled Replay is inspection only
```

Demo · truecolor · 100 × 30 terminal cells

```
  ◉  Last transition                                                                                
  ○  Last three transitions                                                                         
  ○  Complete trace                                                                                 
  ○  Live execution ⊘ disabled Replay is inspection only
```

Narrow · truecolor · 40 × 16 terminal cells

```
  ◉  Last transition                    
  ○  Last three transitions             
  ○  Complete trace                     
  ○  Live execution ⊘ disabled Replay is
```

256 colors · ansi256 · 100 × 30 terminal cells

```
  ◉  Last transition                                                                                
  ○  Last three transitions                                                                         
  ○  Complete trace                                                                                 
  ○  Live execution ⊘ disabled Replay is inspection only
```

16 colors · ansi16 · 100 × 30 terminal cells

```
  ◉  Last transition                                                                                
  ○  Last three transitions                                                                         
  ○  Complete trace                                                                                 
  ○  Live execution ⊘ disabled Replay is inspection only
```

Monochrome · none · 100 × 30 terminal cells

```
 (*) Last transition                                                                                
 ( ) Last three transitions                                                                         
 ( ) Complete trace                                                                                 
 ( ) Live execution [DISABLED] Replay is inspection only
```

### [Go usage](https://lanej.io/charmed/molecules-input.html#input%2fradio-group-code)

Imports: [`input`](https://pkg.go.dev/github.com/lanej/charmed/input "github.com/lanej/charmed/input") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
group, err := input.NewRadioGroup([]input.RadioOption{
	{ID: "last", Label: "Last transition"},
	{ID: "three", Label: "Last three transitions"},
	{ID: "all", Label: "Complete trace"},
	{ID: "live", Label: "Live execution", DisabledReason: "Replay is inspection only"},
}, "last", theme.New(theme.Dark))
if err != nil {
	return "", err
}

group.Focus()
// Arrows move focus only; Enter/Space commits and emits RadioSelected.
// Select changes selection programmatically without emitting a command.
return group.View(), nil
```

### [Updates, focus & sizing](https://lanej.io/charmed/molecules-input.html#input%2fradio-group-integration)

Construct ordered options with unique nonempty IDs and an optional selected ID. Focus or blur the group from the parent, which owns Tab traversal. Arrows move focus only; Enter or Space commits selection and returns a RadioSelected command. Return that command to the application. Programmatic Select changes selection without activation; individual Radio atoms remain available independently.

### [Run this example](https://lanej.io/charmed/molecules-input.html#input%2fradio-group-run)

```sh
go run ./examples/catalog input/radio-group
```

[Example source](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/docs/snippets/input_radio_group.go) · [Implementation](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/input/radio_group.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

### Reusable component

[Radio group dependency graph](https://lanej.io/charmed/dependencies.html#component-input%2Fradio-group) · Partial map · [Source declaration](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/input/radio_group.go#L20)

- [Radio](https://lanej.io/charmed/dependencies.html#component-input%2Fradio) · composes · options

Consumer usage uses

- [Radio group](https://lanej.io/charmed/dependencies.html#component-input%2Fradio-group) · primary demonstration

[Example source](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/docs/snippets/input_radio_group.go#L8)
