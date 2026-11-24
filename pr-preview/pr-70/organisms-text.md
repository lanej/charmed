[Organisms](https://lanej.io/charmed/organisms.html)Text

# [Text](https://lanej.io/charmed/organisms-text.html#level-organisms-text)

1 component · [All organisms](https://lanej.io/charmed/organisms.html)

- [Markdown](https://lanej.io/charmed/organisms-text.html#text%2fmarkdown)

[Organisms](https://lanej.io/charmed/organisms.html) / Text

## [Markdown](https://lanej.io/charmed/organisms-text.html#text%2fmarkdown)

Glamour renders Markdown within a native scrollable viewport.

### [Native preview](https://lanej.io/charmed/organisms-text.html#text%2fmarkdown-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 Help                                           
                                                
Call  field.Focus()  to focus the input.        
                                                
  cmd := field.Focus()                          
  
```

Demo · truecolor · 100 × 30 terminal cells

```
   Terminal documentation                                                                           
                                                                                                    
  Use  field.Focus()  with Markdown, emphasis, and Glamour                                          
  https://github.com/charmbracelet/glamour.                                                         
                                                                                                    
  • Theme follows the terminal palette                                                              
  • Prose wraps to its allocated width                                                              
                                                                                                    
    // Focus a native input                                                                         
    field := input.NewText(options)                                                                 
    cmd := field.Focus()                                                                            
                                                                                                    
  │ Keep content, focus and layout independent.                                                     
                                                                                                    
   Content                                        │ Rendering                                       
  ────────────────────────────────────────────────┼───────────────────────────────────────────────  
   Inline code                                    │ Literal spans                                   
   Fenced source                                  │ Syntax highlighting
```

Narrow · truecolor · 40 × 16 terminal cells

```
   Terminal documentation               
                                        
  Use  field.Focus()  with Markdown,    
  emphasis, and Glamour                 
  https://github.com/charmbracelet/gla  
  mour.                                 
                                        
  • Theme follows the terminal palette  
  • Prose wraps to its allocated width  
                                        
    // Focus a native input             
    field := input.NewText(options)     
    cmd := field.Focus()                
                                        
  │ Keep content, focus and layout      
  │ independent.                      
```

256 colors · ansi256 · 100 × 30 terminal cells

```
   Terminal documentation                                                                           
                                                                                                    
  Use  field.Focus()  with Markdown, emphasis, and Glamour                                          
  https://github.com/charmbracelet/glamour.                                                         
                                                                                                    
  • Theme follows the terminal palette                                                              
  • Prose wraps to its allocated width                                                              
                                                                                                    
    // Focus a native input                                                                         
    field := input.NewText(options)                                                                 
    cmd := field.Focus()                                                                            
                                                                                                    
  │ Keep content, focus and layout independent.                                                     
                                                                                                    
   Content                                        │ Rendering                                       
  ────────────────────────────────────────────────┼───────────────────────────────────────────────  
   Inline code                                    │ Literal spans                                   
   Fenced source                                  │ Syntax highlighting
```

16 colors · ansi16 · 100 × 30 terminal cells

```
   Terminal documentation                                                                           
                                                                                                    
  Use  field.Focus()  with Markdown, emphasis, and Glamour                                          
  https://github.com/charmbracelet/glamour.                                                         
                                                                                                    
  • Theme follows the terminal palette                                                              
  • Prose wraps to its allocated width                                                              
                                                                                                    
    // Focus a native input                                                                         
    field := input.NewText(options)                                                                 
    cmd := field.Focus()                                                                            
                                                                                                    
  │ Keep content, focus and layout independent.                                                     
                                                                                                    
   Content                                        │ Rendering                                       
  ────────────────────────────────────────────────┼───────────────────────────────────────────────  
   Inline code                                    │ Literal spans                                   
   Fenced source                                  │ Syntax highlighting
```

Monochrome · none · 100 × 30 terminal cells

```
  # Terminal documentation                                                                          
                                                                                                    
  Use field.Focus() with **Markdown**, *emphasis*, and Glamour                                      
  https://github.com/charmbracelet/glamour.                                                         
                                                                                                    
  * Theme follows the terminal palette                                                              
  * Prose wraps to its allocated width                                                              
                                                                                                    
    // Focus a native input                                                                         
    field := input.NewText(options)                                                                 
    cmd := field.Focus()                                                                            
                                                                                                    
  | Keep content, focus and layout independent.                                                     
                                                                                                    
   Content                                        | Rendering                                       
  ------------------------------------------------|-----------------------------------------------  
   Inline code                                    | Literal spans                                   
   Fenced source                                  | Syntax highlighting
```

### [Go usage](https://lanej.io/charmed/organisms-text.html#text%2fmarkdown-code)

Imports: [`collection`](https://pkg.go.dev/github.com/lanej/charmed/collection "github.com/lanej/charmed/collection") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

````go
palette := theme.New(theme.Dark)
content := "# Help\n\nCall `field.Focus()` to **focus** the input.\n\n```go\ncmd := field.Focus()\n```"
styles := text.MarkdownStyles(palette)
styles.Document.Margin = new(uint(0))

rendered, err := text.Markdown(content, 48, palette, text.MarkdownOptions{Styles: &styles})
if err != nil {
	return "", err
}

viewport := collection.NewViewport(rendered, 48, 12)

return viewport.View(), nil
````

### [Updates, focus & sizing](https://lanej.io/charmed/organisms-text.html#text%2fmarkdown-integration)

Render trusted application content with the shared theme, then place long output in a native viewport. Route scrolling only to the active viewport and retain its returned model and command. Re-render at the new column allocation on resize; the application owns content replacement and link handling.

### [Run this example](https://lanej.io/charmed/organisms-text.html#text%2fmarkdown-run)

```sh
go run ./examples/catalog text/markdown
```

[Example source](https://github.com/lanej/charmed/blob/27e05a1997640606888402c2223835e8d5e8e451/docs/snippets/text_markdown.go) · [Implementation](https://github.com/lanej/charmed/blob/27e05a1997640606888402c2223835e8d5e8e451/story/organisms.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Declared example links (unaudited): [collection/viewport](https://lanej.io/charmed/atoms-collection.html#collection/viewport)
