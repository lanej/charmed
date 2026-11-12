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

## Evidence and key-value layouts

[Evidence states](atoms-feedback.html#feedback/evidence) describe how much trust to place in an observation. Parse the state once at the boundary with `feedback.ParseEvidence`; an unrecognized input becomes UNKNOWN. Every state retains its own label and glyph in each theme.

| Component | Contract |
| --- | --- |
| Evidence atom | Healthy maps to OK and down to ERROR. Degraded, stale and truncated map to WARN. Unknown, insufficient and attribution gap map to UNKNOWN. |
| [Summary line](molecules-feedback.html#feedback/evidence-summary) | Count by state, worst first. An empty tally reads “no evidence.” |
| [Evidence footer](molecules-feedback.html#feedback/evidence-footer) | Supply source, scope, state, observation time and current time. Missing parts stay explicit. `text.Age` formats one compact unit. |
| [Key-value layout](molecules-layout.html#layout/key-value) | Stack at 55 columns or fewer. `layout.WidthClassOf` uses 40, 55, 80, 100, 120 and 160 columns. Missing values stay distinct from empty strings. |

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

## Selection and action atoms

Construct each control with a stable application ID, a plain-text label and a shared theme. The embedded `input.Control` exposes native key bindings, focus methods and disabled state. The parent owns traversal and forwards Space or Enter to the focused control. Retain the model returned by `Update` and return its command to Bubble Tea.

| Control | Activation and message |
| --- | --- |
| [Checkbox](atoms-input.html#input/checkbox) | Unchecked and mixed become checked. Checked becomes unchecked. Invalid states display a question mark and suppress activation. Emits `input.CheckChanged{ID, State}`. |
| [Radio](atoms-input.html#input/radio) | Emits `input.RadioSelected{ID}` on a new selection. The parent owns the group and deselects siblings. |
| [Toggle](atoms-input.html#input/toggle) | Accepts an initial boolean, flips it on activation and emits `input.Toggled{ID, On}`. |
| [Button](atoms-input.html#input/button) | Emits `input.Activated{ID}`. The parent executes the operation, sets `Busy` while it runs, handles its outcome and clears the flag. |

A nonempty `DisabledReason` suppresses activation and appears with the disabled label. Busy buttons also suppress activation. Labels and reasons are escaped before rendering. Focus and selection preserve cell allocation, and native binding and theme changes survive updates.

## Data and evidence atoms

The application retains the underlying observation and supplies its meaning. A missing measurement, a known zero and an empty string carry different information. Pass the current time explicitly when rendering freshness so a saved frame can reproduce the original decision.

| Atom | Input and behavior |
| --- | --- |
| [Value state](atoms-text.html#text/value-state) | Present is the zero-value state and accepts zero, false and empty text. Missing, redacted and invalid hide the supplied value. An unrecognized state displays unknown. |
| [Freshness](atoms-feedback.html#feedback/freshness) | An age equal to or greater than a positive `StaleAfter` is stale. Nonpositive thresholds disable classification. Missing and future times stay explicit. The parent schedules refreshes. |
| [Quantity](atoms-text.html#text/quantity) | Supply units, precision and minimum alignment width. Decimal or binary scaling is opt-in. Precision spans 0–9 decimals. Tiny nonzero values use scientific notation, nonfinite values are invalid, and wide values stay complete. |
| [Identifier](atoms-text.html#text/identifier) | `Label(width)` abbreviates the middle in terminal cells and marks omission. An empty identity displays unavailable when space permits. Use the full `Value` for actions and identity. |
| [Match highlighting](atoms-text.html#text/match) | Supply native styles and half-open rune offsets into the source text. Overlaps combine, out-of-range offsets clip and reversed ranges are ignored. Any matching rune highlights its grapheme. Source controls are escaped. |
| [Sort indicator](atoms-text.html#text/sort) | Display supplied ascending, descending or unsorted order, optionally with a positive one-based priority. The caller sorts the records. |
| [Change marker](atoms-text.html#text/change) | Display unchanged, added, modified or removed using both a symbol and a label. The caller judges the effect. |
| [Omission marker](atoms-text.html#text/omission) | Distinguish hidden display content from truncated source. Known zero renders nothing, unknown count remains explicit and a negative known count displays invalid. |
| [Operation outcome](atoms-feedback.html#feedback/outcome) | Display succeeded, failed, canceled, timed out or unknown. The application owns cancellation, reconciliation and final classification. A cancellation or timeout alone cannot establish whether an external effect occurred. |

## Unicode control language

Each control reserves equal space across states. Unicode is the default, while `Theme.ASCII` selects the text fallback independently of color depth. Measure the rendered control when sizing its parent because a Unicode switch and its ASCII equivalent have different widths.

| Control or state | Unicode presentation |
| --- | --- |
| Checkbox and radio | `☐ / ☑ / ▣` for unchecked, checked and mixed. `○ / ◉` for radio selection. |
| Toggle | `○── OFF / ━━● ON` conveys the setting through position, fill and text. |
| Button | Angle enclosures and a named `Button.Icon`, defaulting to the action arrow. Busy uses ⟳ and disabled uses ⊘ with visible words. |
| Power and connection | Off ⏻, inactive ◌ and disconnected ○ remain distinct. The caller supplies each semantic state. |

Active controls use the accent and inactive controls use muted text. Focus uses an inverse surface with a reserved inset. Outcome, change, value-state, omission and freshness retain their labels alongside their named glyphs, including in monochrome output.

## Compact links

[Links](atoms-text.html#text/link) store a plain-text label separately from an exact absolute URI. Use `text.NewWebLink` to require HTTP or HTTPS and a host. Use `text.Link` for a configured custom scheme, file path URI, email address or URN. Callers own later field changes, handler policy and actions such as opening or copying.

| Operation | Output and validation |
| --- | --- |
| `link.View(palette)` | An underlined label and an OSC 8 link with the exact target. Each link closes before adjacent content. The Unicode glyph is ↗ and the ASCII glyph is @. |
| Presentation | `GlyphAndLabel` is the default. `LabelOnly` and `GlyphOnly` change the visible content. Normal, bright and dim emphasis leave the destination intact. |
| Theme | `LinkStyle` draws from `Link`, then `Accent`, then ordinary text in partial palettes. A missing `DimLink` retains the normal style. |
| `link.Plain()` | Returns the label and full URI for logs or redirected output. Removing ANSI from the styled view loses its hidden URI. |
| Invalid input | Empty labels, relative or malformed URIs, whitespace, terminal controls, format controls or invalid UTF-8 in the target return an error and no output. Labels are escaped. |

Kitty can select explicit links through [hyperlink hints](https://sw.kovidgoyal.net/kitty/kittens/hints/) and show targets through [target inspection](https://sw.kovidgoyal.net/kitty/conf/#opt-kitty.show_hyperlink_targets). Opening a custom URI depends on an installed handler. Static previews show the styling. Actual Kitty, multiplexer and SSH behavior is tracked in the [terminal links issue](https://github.com/lanej/charmed/issues/37).

## Related destinations

[Link groups](molecules-layout.html#layout/link-group) take links in declaration order. They preserve each presentation and exact destination, place two cells between items, and wrap whole items before wrapping oversized labels. Every wrapped line retains its target. Measure the resulting height or use a viewport when placing the group in a header.

| Case | Result |
| --- | --- |
| Empty group | Returns empty output. A nonempty group requires positive width. |
| Invalid link or oversized grapheme | Returns a zero-based item index and an error, with no partial output. |
| Plain output | Prints each full label and URI on a separate line, including links styled as glyph-only. |
| Selection or opening | The application handles interaction and retains the original link values. Rendering leaves the input slice intact. |

## Git presentation

The `gitui` components display caller-supplied observations. Applications own discovery, parsing, refresh schedules and Git operations. Compose refs, HEAD, file-change, divergence, diff and worktree atoms into rows, status groups, lists, history or the repository browser. Construct a group directly when you only need one section.

| Component | Data and interaction contract |
| --- | --- |
| Commit history | Supply child-before-parent commits with full IDs and ordered parents. Set `History.Truncated` for missing parents. Pending lanes stay stable, and missing history ends in a continuation marker. |
| Graph and search | Crossings use ╪ or ASCII colon. Junctions use joined connectors. Narrow graphs report insufficient width. Searching history moves between matches while preserving all rows. Other lists filter. |
| File and section indicators | `FileIndicator` shows what happened to a file. `SectionIndicator` identifies staged, unstaged, untracked or conflicting observations. Files with both index and working-tree changes appear in both matching sections. |
| Change groups | `NewStagedChanges`, `NewUnstagedChanges`, `NewUntrackedChanges` and `NewConflictChanges` expose count, rows and empty state. `View` renders a group and `Items` supplies stable rows to a parent. |
| Selection | `interaction.Selected` carries the full commit, ref or supplied file, stash and worktree identity. Status includes the section. Shortened ref labels preserve their full event identity. |
| Missing observations | Unknown divergence and unavailable worktree status remain explicit. `Repository.Unavailable` distinguishes unavailable evidence from zero recorded changes. |

Constructors copy collection inputs, and setters validate replacements. Refresh preserves identity or selects the nearest remaining item. History, branch, stash, worktree and status setters retain native viewport style and input customization. The component owns content, dimensions and offsets. Replace diffs after status refresh with `SetDiffs`: the diff viewport returns to its beginning and a missing file/section pair shows unavailable content. Rebuild styled rows after a theme change.

| Focused region | Keys |
| --- | --- |
| Lists | Arrows or j/k move, Home/End or g/G choose endpoints, paging moves through results, slash searches, Escape clears, and Enter selects. Search accepts pasted text. |
| History | Ctrl+E/Ctrl+Y scroll independently of commit selection so continuation rows remain reachable. |
| Changes inspector | The parent uses `SetFocus` to route messages to the status or diff pane. |
| Repository browser | Ctrl+Left/Right changes sections. Tab changes panes. Enter selects the supplied worktree context without filesystem effects. |

## Markdown and source code

Use [inline code and source blocks](atoms-text.html#text/code) for literal input. `text.InlineCode` preserves Markdown delimiters, and `text.CodeBlock` highlights a supplied language or falls back to plain text. [Rich Markdown](organisms-text.html#text/markdown) supports headings, links, quotes, lists, tables and code fences. Prose wraps to the requested width. Long source lines can extend beyond it, so place them in a viewport when readers need the whole line.

`text.MarkdownStyles` returns an owned Glamour style configuration. `text.MarkdownSyntax` returns a local Chroma style derived from the same palette. Pass either through `text.MarkdownOptions` to `Markdown` or `CodeBlock`. Use the `Syntax` option for highlighting, including named Chroma palettes. Reapply the options when rendering after a resize. Monochrome output strips ANSI styling.
