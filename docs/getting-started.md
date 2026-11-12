# Build your first terminal app

Charmed supplies themed Go components for native terminal interfaces. Start with an editable environment field, then compose charts, collections and inspectors around the same application model and theme.

## Install

Use Go 1.26.11 or newer. Create a module for your application, then add Charmed as a dependency. Bubble Tea owns the terminal session and delivers messages to your model.

```sh
mkdir environment-app
cd environment-app
go mod init example.com/environment-app
go get github.com/lanej/charmed
```

## Build your first app

Create the field with a shared palette and an explicit width. The catalog pairs every minimal usage example with its actual terminal output, required imports and a command that renders it.

```go
field := input.NewText(input.TextOptions{
    Value: "production", Width: 36,
    Theme: theme.New(theme.Dark),
})
```

The complete [quickstart program](https://github.com/lanej/charmed/blob/main/examples/quickstart/main.go) includes imports, the outer model and terminal startup. Its event loop has three responsibilities: return the focus command, retain input updates, and resize the existing field when the terminal changes.

| Lifecycle | Application responsibility |
| --- | --- |
| Init | Return `field.Focus()` so Bubble Tea starts the cursor command |
| Update | Retain both values from `field.Update(msg)` and return its command |
| Resize | Call `field.Resize(columns)` with the available terminal cells |
| View | Return `tea.NewView(field.View())`; rendering performs no input or network work |
| Exit | Handle Ctrl+C in the root model and return `tea.Quit` |

## Run and change it

To try the complete program immediately, clone Charmed and run the quickstart from its root. Edit the field, resize the terminal, and exit with Ctrl+C. Copy the linked program into your own module when ready.

```sh
git clone https://github.com/lanej/charmed.git
cd charmed
go run ./examples/quickstart
```

For a saved example without an interactive terminal, run the catalog renderer. It calls the same consumer function used by the documentation build, including the supplied sample data and final View call.

```sh
go run ./examples/catalog input/text
go run ./examples/catalog chart/timeseries
go run ./examples/catalog pattern/log-inspector
```

Choose [components](atoms.html) when assembling individual widgets, or [recipes](recipes.html) when coordinating a complete view. The [composition guide](examples.html) explains ownership of focus, terminal geometry and application observations.
