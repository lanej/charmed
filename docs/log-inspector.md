# Correlated HTTP and application logs

Use the [correlated inspector](organisms-logs.html#pattern/log-inspector) to investigate an exchange with its exact request or trace logs. Supply observations explicitly; the UI neither collects telemetry nor issues network requests. Request and response have independent payload viewports, coordinated selection and responsive allocation.

```sh
go run ./examples/catalog pattern/log-inspector
go run ./cmd/workbench run pattern/log-inspector --single-request --no-help
```

## Supply observations

Each HTTPExchange and LogRecord needs a unique nonempty ID. Correlation IDs are explicit optional metadata. Deliver replacements and upserts through the owning Bubble Tea update loop; producer goroutines should send messages rather than mutate UI models directly.

| Operation or observation | Contract |
| --- | --- |
| SetEntries | Validate before replacing a dataset; failures retain the existing dataset |
| UpsertRequests / UpsertLogs | Replace matching IDs in place; append new IDs in arrival order, independently of timestamps |
| Retention | Default 500 records; the Limit option applies to each collection |
| Eviction | Repair a removed selection using the first retained record and discard evicted payload snapshots |
| Ownership | Headers and body bytes are copied; nested LogEntry.Fields maps and slices remain caller-owned immutable observations |
| Missing values | Distinguish absent identifiers, no matches, unavailable duration, pending exchanges and empty captured bodies |

Requests initially follow the latest arrival. Navigating older records pauses follow; F resumes it for the focused request or log pane. Selection, body modes and retained observations survive layout changes, while resizing visible content may clamp native scroll offsets.

## Correlation and hierarchy

Exact request-ID matching is the default. Trace scope requires an observed trace ID on the selected request; timestamps and similar messages never establish a relationship. The application supplies IDs and their meaning, while the inspector keeps unavailable metadata distinct from an available ID with no matching logs.

| Structure | Behavior |
| --- | --- |
| Request scope | Match the selected request ID exactly |
| Trace scope | Press t to match its trace ID exactly when available |
| Log tree | Group by trace and span, or press g for flat entries |
| Unresolved hierarchy | Missing, cyclic and conflicting span parents stay visible as unresolved |
| Structured fields | Preserve duplicate top-level names by occurrence and array positions by index |
| Nested values | Traverse map[string]any and []any within depth, node and preview limits; label unsupported types without invoking arbitrary formatters |

## Payloads and source fidelity

HTTP observations preserve original targets, headers, metadata and captured bytes. Display transformations never convert a missing response into success or an incomplete body into an empty one. Terminal controls are escaped as data, and explicit capture notices remain visible beside the payload.

| Body state or format | Presentation |
| --- | --- |
| Complete declared JSON | Source-preserving [JSON viewer](json-viewer.html), including supported +json media types and header content-type fallback |
| JSON navigation | Up/down choose fields; Enter toggles containers; left/right choose parents or expand; paging and wheel input scroll wrapped values |
| Raw mode | Press r; structural state and raw/formatted scroll positions are retained for identical bytes |
| Changed bytes | Reset only the changed payload direction; header-only changes retain body state |
| Invalid or over-budget JSON | Clear stale structure and display a bounded explicit fallback |
| Truncated, redacted or encoded bodies | Keep capture notices and the existing summary/source fallback |
| URL or form parameters | Preserve repeated and empty values; reject malformed or partial decoding rather than presenting a partial parse |

Request parameters decode URL query strings and complete captured application/x-www-form-urlencoded bodies within the preview byte budget. Charset, encoding and capture state gate decoding. Route parameters require application routing knowledge and are never inferred. Disabling folds makes parameter sections visible without changing retained observations.

## Targets and summaries

The concise summary keeps the method and path, exact response code and description, duration and timestamp together. Optional correlation IDs retain labels because otherwise similar identifiers are ambiguous. Missing optional IDs and timestamps are omitted; unknown codes and unavailable elapsed time remain explicit.

| Setting | Behavior |
| --- | --- |
| Default HTTP target | Path only; query values remain in the Parameters section |
| PathAndQuery | Preserve captured query ordering and escapes |
| HostAndPath | Omit scheme and user information |
| FullURL | Display the captured target without reconstruction |
| Malformed or opaque target | Retain captured spelling |
| Custom formatter | SetURLFormatter is pure, inexpensive and escaped as data; original observations remain unchanged |
| VerboseDetails / ColumnLabels | Restore labeled summary fields or table headings; resize after changing column visibility |

Symbol presentation supports glyph-and-label, glyph-only and label-only modes. IDs and observed HTTP values stay visible in every mode. Native tree cursors and fold markers are structural controls; pending code cells retain their compact slot. Call SetSymbols to refresh a molecule or inspector after changing presentation.

## Responsive layout and native state

Automatic layout uses a full-width request timeline when space permits, then falls back to a sidebar and finally one focused pane. Logs move below the payload pair before request and response stack. Both payloads receive their region's full height; long values wrap inside their own viewports.

| Layout contract | Behavior |
| --- | --- |
| Presets | InspectorSidebar, InspectorTimeline, InspectorColumns and InspectorFocused; SetLayout changes the preference without replacing datasets |
| Preferred versus active | Layout reports the preference; ActiveLayout reports the arrangement that currently fits |
| Timeline height | Measure axis, headings, request rows and optional legend, capped near 30% of body height; a single request takes four framed rows |
| Payload pair | Request starts at 44 framed cells; response gets the remainder; below 95 cells the pair stacks |
| Native payload customization | Use Request and Response; HTTP.Details remains the standalone combined view |
| Resize | Keep styles, key maps, selection, folds and body modes; visible offsets may clamp |
| Standalone HTTP details | Use measured content widths, with minima of 24 and 35 cells plus a three-cell gap, before choosing columns |

Long header values that fit the pane move to their own line rather than fragmenting under a long name. Standalone whole-side folding returns its column to the remaining content. The inspector's independently framed payloads retain their own scrolling and structural JSON state.

## Timelines and grouping

Timeline bars use a shared explicit UTC clock. Each bar starts at the observed request time, and completed requests end only when a nonnegative duration is known. Pending bars extend to the supplied observation time with an open arrow; unknown timing stays visibly unknown.

| Timeline option | Contract |
| --- | --- |
| TimelineNow | Supply the observation time; the UI does not invent current time during rendering |
| Resolution | Sub-cell intervals become endpoint symbols; exact timestamps and duration remain in details |
| Width | At least 36 columns; smaller allocations show a size hint |
| TimelineLegend | Optional endpoint key, hidden by default |
| Grouping | Compose pure HTTPGroup callbacks; built-ins include service, endpoint, host and port |
| Group order | First seen, with arrival order inside each group; navigation follows display order while IDs remain stable |
| Destination semantics | URL authority and inferred HTTP/HTTPS ports describe targets, never observed socket peers |

Port grouping normalizes explicit numeric ports and infers 80 or 443 for absolute HTTP or HTTPS targets. Missing or invalid ports stay unavailable. Clearing grouping restores ungrouped order; short viewports retain selection and repeat its group heading when space allows.

## Focused logs and scrollbars

Detached log details replace the selected leaf at its position in the hierarchy, retaining siblings before and after it. The full-width focused container includes message, timestamp and expanded fields under the same escaping and bounded traversal rules as inline fields.

| Control or customization | Behavior |
| --- | --- |
| D | Switch keyboard scrolling between tree and detail viewport; detail scrolling pauses follow |
| Selection | Reset details offset for a different record; preserve it when updating the same record where possible |
| Trace/span selection | Clear the record inspector |
| Disable detached details | Restore inline fields, keep the selected record and reset field expansion |
| FocusedLeafGeometry | Supply rendered leaf row bounds for custom padded/wrapped nodes; invalid geometry retains the compact tree instead of guessing |
| ScrollbarGeometry | Supply total rows, visible rows, offset and track height when changing default tree framing or native help |
| ShowScrollbar=false | Disable the indicator and resize to reclaim its column; one-cell allocations prioritize content |

Default tree scrollbar geometry assumes an unframed tree with native help hidden. Request tables retain native scrolling without claiming an offset their API does not expose. Native Details and payload viewports remain public for styling and key maps; the caller owns any geometry changes those customizations require.

## Folding, instructions and the request headline

The request UUID is the default headline, using [IdentityHeader](molecules-layout.html#layout/identity-header). It wraps the complete identifier after moving the kind label above it when needed. Missing IDs stay explicitly unavailable; NoHeadline removes the heading when an application supplies its own header.

| Setting or key | Behavior |
| --- | --- |
| z / SetPaneCollapsed | Fold requests, request, response or logs to a focusable heading row |
| Fold requests | Return timeline rows to lower regions |
| Fold logs | Return the payload pair's full available width |
| Fold either payload | Return the pair's width to the other payload |
| Legacy http pane ID | Alias response; FocusedPane retains http for compatibility |
| NoFolds / SetFolding(false) | Expand component, HTTP and log folds together, clearing retained JSON folds |
| Tab / Alt+1/2/3/4 | Move focus, including collapsed headings; z reopens the focused heading |
| NoHelp / SetHelpVisible(false) / ? | Hide all bottom key-action instructions and reclaim those rows; keep shortcuts and command counts active |

Hidden panes ignore content keys while continuing to receive observations. Their native viewport allocation and scroll position remain retained until reopening; a new visible allocation may clamp offsets. Applications with their own footer should gate key hints on HelpVisible while leaving application status and coverage notices visible.

## Commands and presets

Ctrl+K opens the [command palette](command-palette.html) for layouts, focus, follow, correlation, body display and folds. It owns keyboard and paste input while observations continue updating. Esc restores the previous pane; Enter executes one currently available action and closes.

```sh
go run ./cmd/workbench run pattern/log-inspector \
  --log-layout timeline --timeline-group-by service,endpoint \
  --detached-log-details --key-layout vim --palette-backdrop dim
```

The Vim preset applies across palette, request table, log tree and both payload viewports; native overrides applied afterward survive resize and updates. Saved variants include single requests, folded payloads, no instructions, automatic layout, narrow allocations and JSON containers. Use the catalog renderer for exact minimal-example output and the workbench for interactive fixture exploration.
