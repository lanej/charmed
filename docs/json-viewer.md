# JSON viewer

The [JSON viewer molecule](components.html#collection/json-viewer) keeps syntax-highlighted JSON readable while letting you fold objects and arrays. It starts expanded. Collapsed containers display their field or item count. Member order, duplicate keys, exact number text, string escapes and original source bytes are retained. Formatting never decodes objects into maps or converts numbers to floating point.

Create a viewer with `collection.NewJSONViewer(width, height, palette)`, pass JSON bytes to `SetSource`, forward Bubble Tea messages through `Update`, and render `View`. `SetSource` copies its input. Invalid or over-budget input clears stale structure, returns a diagnostic and displays a bounded source fallback. The [consumer example](components.html#collection/json-viewer) shows a complete payload and setup.

## Interaction

Up/down navigate visible fields. Enter toggles the selected container. Right expands a closed container or enters its first child; left collapses an open container or selects its parent. Page, half-page and wheel input retain native line scrolling so the middle of a long wrapped string remains readable. The next structural action reveals the selected field again. Folding an ancestor selects that ancestor; descendants keep their own folds until reopened.

The inspector uses this viewer inside both payload panes for complete declared JSON in supported encodings, including `+json` media types and header content-type fallback. `z` folds the entire component, `r` switches raw/formatted bodies, and the command palette exposes the selected JSON container's toggle. `?` and `NoHelp` control all footer instructions. The standalone viewer emits no keyboard footer; applications own help visibility and raw-mode bindings.

## Native models and state

`Model` is a native `*viewport.Model`; it can point at an existing viewport before setting source. Structural navigation reads its up/down/left/right bindings, paging retains native behavior, and `Toggle` is a configurable binding. Refresh owns content while retaining native geometry, styles and key maps. Configure `Theme`, `Symbols`, `Prefix`, `Focused`, `Raw` and `Folding` before `Refresh` or `SetSource`. Change `Limits` before calling `SetSource` again to recheck parsing budgets. `Prefix` accepts caller-owned styled header text; the caller sizes and escapes it. Payload terminal controls remain escaped data.

`LogInspector.RequestJSON` and `ResponseJSON` add structure to the existing `Request` and `Response` viewports without replacing those models. Each payload has its own cursor and folds. Resize, whole-component folding and raw-mode round trips preserve structural state. Raw and formatted views have separate scroll positions, which clamp to the allocation. Changed body bytes reset only that direction. Returning to an exchange restores snapshots only for identical source bytes; discarded observations evict their snapshots. Header changes do not reset the body.

`Snapshot` returns an opaque `JSONState` containing no captured bytes or native customization. `RestoreState` rejects different source. Standalone hosts own snapshot retention. `Selected`, `Select`, `Collapsed` and `SetCollapsed` use preorder node indices, stable only for identical bytes; duplicate members therefore have distinct identities. Disabling `Folding` expands everything, clears folds and routes navigation to the viewport. Re-enabling it keeps the expanded state. The inspector's `SetFolding(false)` also clears retained snapshots.

## Limits and fallbacks

Defaults are 64 KiB of source, 64 container levels, 4,096 value nodes and 256 KiB of styled, escaped, wrapped body output. Nonpositive limits use defaults; depth has an absolute ceiling of 256 to bound recursive traversal. The output budget excludes the caller-owned prefix and a short diagnostic. The HTTP preview byte budget additionally gates parsing without changing the caller's JSON limits.

Invalid, truncated, redacted, unsupported-charset, encoded and over-budget HTTP bodies retain their existing source or summary fallback with explicit notices. A syntactically valid captured prefix is still incomplete and cannot become interactive JSON. Captured bytes and metadata remain untouched. Editing, query languages and large-array virtualization are outside this first version.
