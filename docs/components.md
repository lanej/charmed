# Component reference

Browse [Atoms](atoms.html), [Molecules](molecules.html) and [Organisms](organisms.html) by capability. Each entry explains its purpose, displays the output of its consumer example, and gives the same usage code with integration guidance and an exact run command.

| Need | Reference |
| --- | --- |
| Color, geometry and shared interaction rules | [Foundation contracts](foundations.html#foundation-contracts) |
| An editable value with validation | [Input field](molecules-input.html#input/field) |
| Coordinated HTTP requests and logs | [Inspector lifecycle](log-inspector.html) |
| Source-preserving structured JSON | [JSON viewer](json-viewer.html) |

## Charts and labeled values

Chart allocations use terminal cells, including borders and gutters. Time-series traces use native braille dots; ASCII retains numeric fallbacks. Legends are opt-in, and callers must reserve their actual wrapped height before assigning space to the plot.

| Component | Sizing and state contract |
| --- | --- |
| `chart.Panel` | Use `options.PlotSize(size, series, palette)` to reserve borders and wrapped legend rows; `PanelPlotSize` retains the legacy single-row allocation |
| Distribution comparisons | Wrap labels and metadata at the supplied width; `ShowLegend` adds matching series keys |
| Labeled fields | Resize `field.Input` to the available columns; help and validation wrap to that width |
| Metric trends | Keep a known current value when history is empty; an absent current value remains unavailable |
| Percentage bars | Preserve tiny nonzero values in scientific notation; a truncation mark identifies an allocation that cannot show the label |

Rendering a panel preserves the child's native state and styles. Data and resize operations may rebuild individual chart models: the time-series entry documents when to reapply direct native customization. The parent decides focus and forwards commands from the active child.

## Inputs

Text inputs reserve their actual prompt width and native cursor cell. Disabling an input suppresses wrapper focus and editing while retaining its native state; re-enabling it restores that state. Validation feedback belongs to a labeled field, rather than the primitive input.

| Change | Native customization behavior |
| --- | --- |
| Disabled view | Applies inactivity presentation to a copy; a caller-owned prompt remains unchanged |
| `SetTheme` | Reapplies shared native styles; apply direct overrides afterward |
| Textarea styling | Uses theme surfaces for active lines, selection and line numbers |
| Saved example | Prepares static display without running cursor or clock commands |

## State diagrams

`diagram.StateDiagram` displays application state and explicit transitions. It does not execute transitions. Declaration order makes placement deterministic, and a complete graph is returned or an error reports the minimum width required across its horizontal and vertical layouts.

| Contract | Behavior |
| --- | --- |
| States | Unique IDs, optional labels, at most one initial state, optional final flags and explicit outcomes |
| Transitions | Valid endpoints; self-loops, cycles and parallel edges remain separate; final states cannot have outgoing edges |
| Geometry | Horizontal when it fits, vertical otherwise; height depends on orientation and belongs to the parent allocation |
| Labels | Single-line plain text, with wide and combining characters measured in terminal cells |
| Appearance | Initial states are rounded, ordinary states square and final states double-bordered; ASCII preserves those distinctions |
| Meaning | Active state is bright and bold; success/error outcomes remain independent of finality and retain text markers |

Transition history references exact indices into the declared transitions. The most recent edge receives stronger emphasis; prior visited states and edges retain their semantic colors. Highlighting preserves geometry and does not claim to display traversal counts or complete chronological order.

| Trace option | Contract |
| --- | --- |
| `Steps` | A continuous chronological path through existing transitions, ending at Active when supplied |
| `Last: 0` | Show the most recent transition |
| Positive `Last` | Show that many recent steps |
| `All: true` | Show the complete supplied path; incompatible with positive Last |
| Invalid history | Reject even an invalid step outside the visible history window; negative Last is an error |

## Connection states

[Connection atoms](atoms-feedback.html#feedback/connection) describe supplied connectivity without probing or managing a connection. Connected uses success styling; disconnected uses inactivity styling. Use UNKNOWN when the application cannot establish connection state, and retain the full label when colors are unavailable.

## Application chrome

Layout parents own event routing and allocation. Actions displayed in headers are text; the application owns their bindings. Reserve wrapped identity-header height before sizing content, and use the section's reported height when returning folded space to siblings.

| Component | Contract |
| --- | --- |
| [Tabs](molecules-navigation.html#navigation/tabs) | Left/Right wrap; Home/End choose endpoints; only the focused bar handles keys; return Update's selection command |
| Tab items | Stable unique IDs; SetItems copies items and retains selection when possible; resize retains selection, focus and customized theme |
| [Header](molecules-layout.html#layout/header) | Reserve one row, or two with context; actions disappear before the title clips |
| [Footer](molecules-layout.html#layout/footer) | One row; status disappears before help clips; styled strings use only their first line |
| [Identity header](molecules-layout.html#layout/identity-header) | Escape Kind, ID and Context; wrap the complete identifier without abbreviation; measure rendered height |
| [Foldable section](molecules-layout.html#layout/section) | Collapsed is one heading row; NoFolds expands without markers; the caller owns state, keys and focus |

The [application shell](templates-pattern.html#template/application) composes these pieces around an active content panel. Sidebar navigation, breadcrumbs and application-specific toolbars remain separate compositions, with their own focus and event ownership.

## Log levels

[Log level atoms](atoms-feedback.html#feedback/log-level) keep severity independent of emphasis. Rendering a fatal label does not log or exit. Unknown levels retain their numeric value, and equal-width labels keep columns aligned across Unicode, ASCII and monochrome output.

| Level | Semantic tone |
| --- | --- |
| TRACE | Subdued neutral; `feedback.TraceLevel` is the display-only typed extension -8 |
| DEBUG | Cyan |
| INFO | Blue |
| WARN | Amber |
| ERROR | Red |
| FATAL | Magenta |

Customize Trace, Debug, Fatal and their Dim theme counterparts independently. Charm Log's native String and ParseLevel do not recognize the TRACE extension. The log viewer composes this atom with timestamps and messages while retaining native structured-field formatting, bounded history and scrolling.
