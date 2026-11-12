# Composition guide

A complete screen owns one Go model and coordinates its child widgets. The service workbench links an input, table, time-series chart and details viewport. The deployment example coordinates host distributions, histories and inspection; [read its telemetry contract and interpretation](deployment.html).

## Ownership rules

| Owner | Responsibility |
| --- | --- |
| Theme | Semantic colors, contrast, ASCII and motion policy |
| Layout parent | Allocate rectangles in columns and rows, including borders |
| Focus scope | Route keyboard input to the active child |
| Component | Own its native widget state and return commands |
| Root runtime | Own the terminal session and forward commands |
| Application | Supply data and handle typed UI events |

```sh
go run ./examples/catalog screen/services
go run ./examples/catalog template/application
```

| Native screen | Coordinated interaction |
| --- | --- |
| Service workbench | Tab changes focus, arrows select rows, Ctrl+N advances samples; selection updates chart and inspector together, and resize retains selection |
| Deployment comparison | Arrows choose hosts and windows, brackets choose response class, 1/2/3 select spread and b changes presentation; saved frames do not execute keys |

Switch between the [deployment fixture](screens-pattern.html#screen/deployment)'s stacked and disconnected-band snapshots using the variant tabs, then inspect the corresponding Go implementation on its group page.

## Explore the catalog

[Templates](templates.html) · [Screens](screens.html)

## Application shell

The application shell fixture composes `layout.Header`, `navigation.Tabs`, `layout.Panel` and `layout.Footer`. The header reserves two rows, the tab bar and footer one each, and the panel receives the remaining height. Left/Right or Home/End changes the active section and its content; resize keeps the selected tab. The fixture contains synthetic display data. In an application with focusable content, use the parent focus scope to focus or blur the tab bar before forwarding keys.
