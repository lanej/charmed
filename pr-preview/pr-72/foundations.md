Atomic design

# [Foundations](https://lanej.io/charmed/foundations.html#level-foundations)

Shared rules for color, type, spacing, geometry and interaction. Start here to keep every component consistent.

7 components · 4 groups

[**Layout**  
\
Layout1 component](https://lanej.io/charmed/foundations-layout.html)

[**Motion**  
\
Motion1 component](https://lanej.io/charmed/foundations-motion.html)

[**Text**  
\
Permissions1 component](https://lanej.io/charmed/foundations-text.html)

[**Theme**  
\
Theme4 components](https://lanej.io/charmed/foundations-theme.html)

## [Foundation contracts](https://lanej.io/charmed/foundations.html#foundation-contracts)

Themes, cell geometry, focus and motion are shared across components. Use semantic roles rather than embedding color decisions in widgets. Parents allocate rectangles in terminal columns and rows, while focus scopes route keyboard messages to the active child.

| Shared concern   | Contract                                                                                                            |
|------------------|---------------------------------------------------------------------------------------------------------------------|
| Reference output | The default saved frame is truecolor; narrow, reduced-color and monochrome tabs demonstrate alternate presentations |
| Native behavior  | Run locally to inspect animation, focus, updates and terminal settings                                              |
| HTML capture     | Preserve completed static cells, RGB colors and escaped text; do not emulate a live terminal                        |

## [Colors and emphasis](https://lanej.io/charmed/foundations.html#colors-and-emphasis)

The default dark palette, Sage Midnight, uses the named roles and state definitions from the [dotfiles design system](https://github.com/lanej/dotfiles/tree/master/design-system). Its sage accent identifies interaction, lavender outlines keyboard focus, and a deep green fill marks selected rows. Semantic colors express meaning, while ordered intensity and categorical series use separate palettes. Keep series order stable across views, and preserve labels when fallback colors converge.

| Role                                  | Meaning and customization                                                                                                  |
|---------------------------------------|----------------------------------------------------------------------------------------------------------------------------|
| Success, information, warning, danger | Semantic state, independent of interaction and series identity                                                             |
| UNKNOWN                               | The `health.unknown` role uses muted text. Configuration off uses subtle text and a power glyph                            |
| Sequential scale                      | Ordered intensity, lighter with larger values in dark themes and darker in light themes                                    |
| Custom sequential fallback            | Neutral black-to-white ramp when the theme has no scale                                                                    |
| Bright, normal, dim                   | Same label, glyph and allocation; dim uses explicit semantic shades rather than terminal faint                             |
| Contrast                              | Built-in dim tokens target 4.5:1 against their background; profiles, fonts and terminal settings still affect presentation |
| Monochrome                            | Keep labels, glyphs and bright weight; normal and dim may coincide                                                         |

## [Display states](https://lanej.io/charmed/foundations.html#display-states)

Applications decide when a state applies; rendering never starts, pauses or disables work. Use feedback.MaxStatusWidth for lifecycle columns and feedback.MaxAttentionWidth for attention-only tables, rather than assuming every label fits a fixed twelve-cell allocation.

| State        | Meaning                        | Sage Midnight role          |
|--------------|--------------------------------|-----------------------------|
| OK           | Healthy or successful          | `status.success`            |
| WARN         | Degraded health                | `status.warning`            |
| ERROR        | Failed or unhealthy            | `status.error`              |
| UNKNOWN      | Cannot establish state         | `health.unknown`            |
| OFF          | Switched off                   | `configuration.off`         |
| PENDING      | Waiting to start               | `execution.queued`          |
| RUNNING      | Actively executing             | `execution.running`         |
| PAUSED       | Suspended and resumable        | `execution.paused`          |
| DISABLED     | Unavailable to start or select | `interaction.disabled`      |
| CONNECTED    | Established connection         | `connectivity.connected`    |
| DISCONNECTED | No established connection      | `connectivity.disconnected` |

The theme resolves each state to its shared color role. Labels and glyphs distinguish states that share a hue, including pending and paused. Charmed supplies its terminal emphasis levels, categorical chart colors and ordered intensity ramp.

## [Glyphs and spacing](https://lanej.io/charmed/foundations.html#glyphs-and-spacing)

Theme.Glyph selects one-cell Unicode or ASCII marks from the named status, action, selection, expansion, navigation and loading catalog. Glyphs supplement labels. Expand changes the viewing area; expanded and collapsed describe disclosure. Reduced motion uses a static loading mark instead of native spinner animation. Infrastructure glyphs distinguish entity labels without changing stable IDs. Unicode is primary; fallback forms are selected only by explicit ASCII policy.

| Entity        | Named glyph         | Unicode   | ASCII |
|---------------|---------------------|-----------|-------|
| Load balancer | `GlyphLoadBalancer` | ⇄ routing | L     |
| Host          | `GlyphHost`         | ▤ host    | H     |

```sh
go run ./examples/catalog foundations/glyphs
go run ./examples/catalog foundations/tokens
```

## [Permission combinations](https://lanej.io/charmed/foundations.html#permission-combinations)

Permissions display supplied io/fs.FileMode bits without reading files or checking effective access. The zero value is unknown; known mode 000 is distinct. Callers identify whether a supplied observation describes a link or its target, and use Go's special-bit constants rather than raw Unix octal flags.

| Permission API       | Contract                                                                                                                                                             |
|----------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Groups(theme)        | Owner, group, other order; one Unicode cell per group                                                                                                                |
| Group shapes         | Read is upper left, write upper right, execute lower left; exact bits remain visible without color                                                                   |
| Capability colors    | Sage Midnight uses the read-only and read/write access colors, executable-file green for execute, and dusty rose for write-plus-execute. The glyph records each bit. |
| Compact(theme)       | Separate groups with spaces and append explicit setuid, setgid and sticky modifiers                                                                                  |
| Label()              | Preserve exact nine-character rwx spelling and s/S/t/T special markers                                                                                               |
| ASCII                | Exact three-character groups, never invented single-letter abbreviations                                                                                             |
| Unknown versus none  | Question mark for unknown, empty-set glyph for no bits                                                                                                               |
| Native customization | Theme.Permissions owns capability colors; absent custom colors retain ordinary foreground                                                                            |

Keep the exact label beside compact permission displays or in their details. Built-in permission colors target 4.5:1 contrast against background, surface and selection fills. Terminal-profile quantization can reduce visible distinctions, so color does not replace the preserved bit shapes and explicit labels.

## [Spacing, borders and selection tokens](https://lanej.io/charmed/foundations.html#spacing-borders-and-selection-tokens)

Spacing is measured in terminal cells. Apply horizontal and vertical gaps explicitly because character cells are taller than they are wide. Panel dimensions include borders and content gutters; allocate children after subtracting those measured regions.

| Token or helper                                   | Contract                                                                                                |
|---------------------------------------------------|---------------------------------------------------------------------------------------------------------|
| SpaceNone / SpaceSmall / SpaceMedium / SpaceLarge | Zero, one, two and four cells                                                                           |
| InlinePadding                                     | One horizontal cell on each side of controls, pills and panel content                                   |
| BorderCells                                       | One cell per frame side; FrameBorder exposes a native border with ASCII fallback                        |
| PanelContentSize                                  | Subtract panel borders and gutters; narrow panels reduce gutters to retain allocation                   |
| PanelPlotSize                                     | Also reserve the optional legacy legend row                                                             |
| FocusBorder                                       | Customize interaction boundaries independently of accent, with accent as fallback                       |
| TablePaddingWidth / DenseTablePaddingWidth        | Budget native cell insets outside declared column-content widths; recalculate after customizing padding |

Focus identifies the active control; selection identifies the chosen item on a quieter surface. Their native styles retain horizontal insets without double-padding tables. Changing a theme may reapply shared defaults, so the component lifecycle documents when native overrides need to be restored.

| Theme customization          | Fallback and ownership                                                                                   |
|------------------------------|----------------------------------------------------------------------------------------------------------|
| Dim tokens omitted           | Retain the supplied normal tone until explicit dim shades exist                                          |
| Selection background omitted | Use bold and underline rather than assume a fill                                                         |
| Selection foreground omitted | Retain the theme's ordinary text color                                                                   |
| Nested native foreground     | Takes precedence; customize native table cells directly and check contrast against the selection surface |
| Monochrome focus             | Preserve inverse highlighting and labels                                                                 |
