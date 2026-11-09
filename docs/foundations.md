# Foundations

Themes, cell geometry, focus and motion are shared across every component. Choose semantic roles rather than embedding color decisions in application widgets. Layout parents allocate rectangles in terminal columns and rows; focus scopes route keys to the active child.

The default saved frame is the truecolor reference. Separate tabs show narrow, 256-color, 16-color and monochrome presentations; fallback color limitations do not define the reference palette. Run with `--theme light`, `--columns 40` or `--ascii --profile none` to inspect those presentations in a terminal. The HTML renderer reads completed static cells, preserves RGB and escapes text; it does not emulate cursor movement or other live terminal effects.

## Colors and emphasis

Foreground and muted text are neutral. Accent identifies interaction; success, information, warning and danger identify semantic meaning. UNKNOWN uses a cool slate-blue tone for uncertainty; OFF uses achromatic gray for inactivity. Their normal and dim tokens remain separate, while labels and glyphs retain meaning without color. Continuous heatmaps use `Theme.Sequential`, an ordered intensity scale independent of category identity. Built-in scales become lighter with larger values in dark/contrast themes and darker in light themes. Custom themes without a sequential scale use a neutral black-to-white ramp. Categorical series use a separate palette and never alias those interface or status tokens. Series order remains stable across views; labels and matching legend keys retain identity when fallback colors converge.

Bright, normal and dim emphasis preserve the label, glyph and cell allocation. Dim uses explicit shades of each semantic hue instead of blending every tone toward gray or relying on the terminal's faint attribute. The built-in dim text tokens meet a 4.5:1 contrast floor against their theme background; native fonts, color profiles and terminal settings still affect appearance. Monochrome keeps labels, glyphs and bright weight; normal and dim may coincide without color.

## Display states

Health, lifecycle and connection states share one catalog through `feedback.Statuses()`. The application decides when each state applies; rendering a state never starts, pauses or disables work. Allocate fixed status columns with `feedback.MaxStatusWidth` so longer labels retain their glyphs and padding. Attention-only tables use `feedback.MaxAttentionWidth`, which measures the attention catalog separately from lifecycle states.

| State | Meaning | Color family |
| --- | --- | --- |
| OK | Healthy or successful | Green |
| WARN | Attention needed | Amber |
| ERROR | Failed or unhealthy | Red |
| UNKNOWN | State cannot be established | Slate blue |
| OFF | Switched off or inactive | Gray |
| PENDING | Waiting to start | Blue |
| RUNNING | Actively executing | Cyan |
| PAUSED | Temporarily suspended and resumable | Sand |
| DISABLED | Unavailable to start or select | Gray, with a slashed-circle glyph |
| CONNECTED | An established connection | Green, with a filled-circle glyph |
| DISCONNECTED | No established connection | Gray, with a hollow-circle glyph |

Pending, running and paused have independent normal/dim tokens. Disabled shares the inactivity palette with off; its label and slashed-circle glyph distinguish unavailability from being switched off. These states do not imply warning/error severity and are independent of attention-rule evaluation.

Connected and disconnected describe connectivity, independently of service health or severity. Disconnected can be intentional; use UNKNOWN when connection state cannot be established. Rendering these atoms does not probe or manage a connection. They reuse success and inactivity colors, including the dim variants; ASCII uses `+` and `-` with the full labels. `MaxStatusWidth` now reserves 16 cells for DISCONNECTED; consumers with hard-coded 12-cell status columns should use the helper or increase their allocation. Attention-only allocations remain unchanged.

## Glyphs and spacing

`theme.Glyphs()` lists the named status, action, selection, expansion, navigation, loading and unavailable marks. Search, filter, sort, refresh, link and expand actions have their own glyphs; expand changes the available viewing area, while expanded/collapsed indicate disclosure state. `Theme.Glyph` selects a one-cell Unicode or ASCII representation. A glyph supplements a semantic text label; `UNKNOWN` and `OFF` remain different states. Animated spinners retain native frame sequences, with a named static loading mark for reduced motion. Application-specific chart markers remain caller-owned.

## Permission combinations

`text.Permissions{Mode: observedMode, Available: true}` displays supplied `io/fs.FileMode` permission bits without reading the filesystem or checking effective access. The zero value is unknown; known mode `000` is distinct. The caller identifies whether the observation describes a link or its target. Pass Go's `fs.ModeSetuid`, `fs.ModeSetgid` and `fs.ModeSticky` flags for special bits; raw Unix `04000`, `02000` and `01000` are not Go's special-bit constants.

`Groups(theme)` returns owner, group and other in that order. Unicode uses one terminal cell per group: `∅ ---`, `▖ --x`, `▝ -w-`, `▞ -wx`, `▘ r--`, `▌ r-x`, `▀ rw-`, `▛ rwx`. Read occupies the upper left, write the upper right, and execute the lower left. `Theme.Permissions` owns independent capability colors: cyan for read only, amber for write, violet for execute, and rose for write plus execute; the latter three include combinations with read. These are capability categories, not safety or severity judgments. Shapes retain the exact bits when color is unavailable. None is `∅`, while unknown is `?`.

`Compact(theme)` separates the three groups with spaces and appends explicit `setuid`, `setgid` and `sticky` modifiers. `Label()` retains the exact nine-character `rwx` spelling, including `s/S/t/T` to distinguish special bits with and without execute. Keep this label available beside compact permission displays or in details. ASCII themes use exact three-character groups, preserving all bits and special markers, rather than inventing single-letter abbreviations. Custom themes without permission colors retain their ordinary foreground. Built-in permission colors meet 4.5:1 contrast against background, surface and selection fills; terminal profile quantization can reduce the visible distinctions.

`theme.InlinePadding` reserves one horizontal cell on each side of controls, pills and panel content. `Theme.Control` and `Theme.Focus` allocate the same space. Status always composes its glyph, one separating cell and its label inside the same padded pill; emphasis changes color and weight without changing geometry. Monochrome focus retains inverse highlighting. `Theme.Selection` uses a quieter selection surface so nested table text and semantic pills remain readable; tables reuse their existing cell insets instead of adding a second focus inset.

Panel dimensions include borders and content gutters. Allocate children with `layout.PanelContentSize`; `chart.PanelPlotSize` also reserves an optional legend row. At widths below five cells, panels reduce the gutter to preserve their allocation. Consumers previously allocating only two border cells must reserve the new content gutters as well. Theme literals without explicit dim tokens retain their supplied normal tone until the caller provides dim shades; literals without a selection surface use bold and underline to retain a visible selection without assuming fill contrast.

## Spacing, borders and selection tokens

Spacing is measured in terminal cells. Apply horizontal and vertical gaps explicitly because character cells are taller than they are wide.

| Token | Cells | Typical use |
| --- | --- | --- |
| `theme.SpaceNone` | 0 | Flush composition |
| `theme.SpaceSmall` | 1 | Inline inset or icon-to-label gap |
| `theme.SpaceMedium` | 2 | Related groups |
| `theme.SpaceLarge` | 4 | Separate sections |

`theme.InlinePadding` aliases `SpaceSmall`. Each frame side occupies `theme.BorderCells` (one cell). `Theme.FrameBorder()` exposes a native Lip Gloss border with an explicit ASCII fallback, and `Theme.BorderStyle(focused)` uses `Border` or `FocusBorder` without changing the frame geometry. `layout.Panel` consumes those defaults. Set `FocusBorder` independently when interaction chrome needs a different color from the accent; custom themes without it retain their accent fallback.

`SelectionForeground` and `SelectionBackground` style the outer surface returned by `Theme.Selection()`. They are independent of inverse keyboard focus: `Theme.Focus()` identifies the active control, while `Theme.Selection()` keeps the selected item's text readable on a quieter surface. Both retain the control's horizontal insets. Built-in selection colors meet a 4.5:1 text contrast floor; monochrome keeps inverse selection, and custom themes without a selection background keep the bold/underline fallback. Omitting only the selection foreground retains the theme's ordinary text color. Nested foreground styles take precedence: native table cells retain their ordinary text or semantic pill colors instead of inheriting `SelectionForeground`. Customize native table cell styles directly when changing those colors, and check them against the selection background. Native styles remain directly customizable.

Table column widths describe content, excluding native cell insets. Budget those insets with `collection.TablePaddingWidth` or `collection.DenseTablePaddingWidth`, then subtract the other columns from the available panel content width. These helpers describe the constructor defaults; recalculate the budget if you customize native cell padding.
