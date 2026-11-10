# Command palette

Use `collection.CommandPalette` to search application commands without moving execution policy into a widget. Commands carry stable IDs, labels, descriptions, keywords, shortcut hints and optional disabled reasons. Disabled commands remain searchable but cannot execute; invalid replacement data leaves existing commands intact.

```sh
go run ./examples/catalog collection/command-palette
go run ./cmd/workbench run pattern/log-inspector --command-palette
```

## Ownership and lifecycle

The palette emits a selected command ID; the application performs the action. While the palette is open, route keyboard and paste input exclusively to it. Continue processing observations normally, and retain the previously focused pane until an accepted action explicitly changes it.

| Operation | Application contract |
| --- | --- |
| Open | Call Open and return its focus command; query resets and the search input focuses |
| Update | Process configurable open, close and execute keys; activation closes and blurs before emitting CommandSelectedMsg |
| Cancel | Dismiss without an action and restore the prior focus region |
| Replace | Supply unique nonempty IDs and labels; caller-owned data is copied atomically |
| Resize | Retain native Search and Results customization, including prompts, styles and key maps |
| Selection | Retain a still-matching ID on command replacement; a changed query selects the best fuzzy match |

## Search and presentation

Match highlighting is opt-in through SetMatchHighlighting. It underlines matching label characters without changing results or selection; keyword-only matches need not underline a label. Rendering and searching never execute application commands, regardless of whether a result is enabled.

| Presentation | Contract |
| --- | --- |
| Selection | Directional marker without a row background; action and description remain distinct |
| Availability | Disabled actions are muted, with the selected reason retained in the status line |
| Size | At most 72 columns and eight visible results; descriptions disappear before essential search content |
| Backdrop | Dim by default; monochrome or unchanged backdrop can be selected explicitly |
| Dim | Blend existing colors halfway toward a known background, retaining text, attributes and links |
| Monochrome | Use one semantic foreground while retaining focus and selection markers |
| Instructions | SetHelpVisible(false) reclaims the keyboard-hint row without removing counts or disabled reasons |

SetBackdropMode rejects unknown values without mutation. An explicit mode takes precedence over the legacy DimBackdrop opt-out; themes without a known background preserve existing colors when dimming. CenterOverlayWithBackdrop composes the opaque foreground within the parent's actual cell allocation.

## Keyboard layouts

Standard navigation uses arrows, paging and endpoint keys, leaving printable characters for search. The Vim preset adds a visible SEARCH/NORMAL mode distinction. Apply presets before native key-map overrides; later resize, filtering and reopening preserve those overrides.

| Mode or key | Behavior |
| --- | --- |
| SEARCH | Printable characters, including j/k, edit the query |
| Esc in SEARCH | Blur the cursor and enter NORMAL |
| NORMAL | j/k navigate, Ctrl+B/F page, G selects the last result, Home the first |
| i or / in NORMAL | Resume search without inserting the transition key |
| Paste in NORMAL | Resume search and accept the paste |
| Esc in NORMAL | Close the palette |
| Enter | Execute an available result and close |

SetKeyLayout returns a command and error. Invalid presets change nothing; changing an open palette's preset returns it to SEARCH while preserving query and selection. Forward that command. Fine-grained movement stays on Results.KeyMap; key.WithHelp keeps displayed hints aligned with custom bindings.

## Inspector integration

The inspector exposes the same molecule as Palette. OpenCommands refreshes current actions and returns its focus command. Availability is checked again before accepting a selection, so an action uses current observations and state rather than trusting an earlier enabled result.

| Inspector control | Configuration |
| --- | --- |
| Layout, focus, follow and correlation | Commands are rebuilt from current inspector state |
| Raw/pretty bodies and folds | Commands operate on the selected request and focused payload |
| URL format | Choosing a built-in format deliberately replaces a custom formatter |
| Keyboard preset | KeyLayout applies across palette, request table, log tree and payload viewports |
| Backdrop | BackdropMode chooses dim, monochrome or unchanged presentation independently |
| Instructions | Inspector NoHelp and SetHelpVisible synchronize the palette's keyboard-hint row |

Use workbench flags `--key-layout standard|vim` and `--palette-backdrop dim|monochrome|none` to inspect those choices. Global application exit bindings may take precedence over the palette. The [inspector guide](log-inspector.html) covers ingestion, correlation and the lifetime of native child customization.
