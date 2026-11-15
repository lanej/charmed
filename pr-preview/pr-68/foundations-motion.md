[Foundations](https://lanej.io/charmed/foundations.html)Motion

# [Motion](https://lanej.io/charmed/foundations-motion.html#level-foundations-motion)

1 component · [All foundations](https://lanej.io/charmed/foundations.html)

- [Motion](https://lanej.io/charmed/foundations-motion.html#feedback%2fmotion)

[Foundations](https://lanej.io/charmed/foundations.html) / Motion

## [Motion](https://lanej.io/charmed/foundations-motion.html#feedback%2fmotion)

A Harmonica spring stepped explicitly by the owner; reduced motion resolves immediately.

### [Native preview](https://lanej.io/charmed/foundations-motion.html#feedback%2fmotion-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
Position 61.39 / target 68.00
```

Demo · truecolor · 100 × 30 terminal cells

```
Spring motion                                                                                       
                                                                                                    
Position 0.00  /  Target 68                                                                         
                                                                                                    
Velocity 0.00                                                                                       
                                                                                                    
Ctrl+N steps the spring. Reduced motion resolves immediately.
```

Narrow · truecolor · 40 × 16 terminal cells

```
Spring motion                           
                                        
Position 0.00  /  Target 68             
                                        
Velocity 0.00                           
                                        
Ctrl+N steps the spring. Reduced motion 
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Spring motion                                                                                       
                                                                                                    
Position 0.00  /  Target 68                                                                         
                                                                                                    
Velocity 0.00                                                                                       
                                                                                                    
Ctrl+N steps the spring. Reduced motion resolves immediately.
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Spring motion                                                                                       
                                                                                                    
Position 0.00  /  Target 68                                                                         
                                                                                                    
Velocity 0.00                                                                                       
                                                                                                    
Ctrl+N steps the spring. Reduced motion resolves immediately.
```

Monochrome · none · 100 × 30 terminal cells

```
Spring motion                                                                                       
                                                                                                    
Position 0.00  /  Target 68                                                                         
                                                                                                    
Velocity 0.00                                                                                       
                                                                                                    
Ctrl+N steps the spring. Reduced motion resolves immediately.
```

### [Go usage](https://lanej.io/charmed/foundations-motion.html#feedback%2fmotion-code)

Imports: [`fmt`](https://pkg.go.dev/fmt "fmt") · [`feedback`](https://pkg.go.dev/github.com/lanej/charmed/feedback "github.com/lanej/charmed/feedback") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
motion := feedback.NewMotion(theme.New(theme.Dark))

target := 68.0
for range 30 {
	motion.Step(target) // In an application, step once per animation tick.
}

return fmt.Sprintf("Position %.2f / target %.2f", motion.Position, target)
```

### [Updates, focus & sizing](https://lanej.io/charmed/foundations-motion.html#feedback%2fmotion-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/foundations-motion.html#feedback%2fmotion-run)

```sh
go run ./examples/catalog feedback/motion
```

[Example source](https://github.com/lanej/charmed/blob/a8a29e65b33913f8df2f703e4af6f8443876d0b9/docs/snippets/feedback_motion.go) · [Implementation](https://github.com/lanej/charmed/blob/a8a29e65b33913f8df2f703e4af6f8443876d0b9/story/foundations.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.
