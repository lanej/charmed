# Brand sheet

Primary lockup

charmed

terminal design system

The terminal is the talisman.

Standalone mark

Attachment loop · terminal window · prompt

One color · dark surface

charmed

One color · light surface

charmed

24px

32px

48px

[Download vector mark ↗](https://lanej.io/charmed/assets/charmed-mark.svg)SVG · currentColor · 64 × 72 grid

## [The terminal talisman](https://lanej.io/charmed/brand.html#the-terminal-talisman)

The Charmed mark combines a terminal with a lucky charm. The terminal window becomes the body of a small talisman: an attachment loop above, a faceted outline below, and a shell prompt inside.

| Form                        | Meaning                                               |
|-----------------------------|-------------------------------------------------------|
| Attachment loop             | A charm that can be carried or attached               |
| Faceted frame               | The shape of a talisman, enclosing one compact object |
| Window rail and `>_` prompt | A native terminal interface                           |

The complete silhouette carries the idea in one color. The mark uses the website's accent through `currentColor`; the reversed specimen uses its existing canvas and text tokens. Terminal data keeps its own semantic colors.

## [Use the identity](https://lanej.io/charmed/brand.html#use-the-identity)

The primary lockup pairs the terminal talisman with the lowercase **charmed** wordmark. Use it for documentation, the website, and repository art. The mark can stand alone when the product name is already visible.

| Identity constraint | Requirement                                                                       |
|---------------------|-----------------------------------------------------------------------------------|
| Source              | Use the supplied SVG and retain the loop and faceted outline                      |
| Clearance           | At least one cursor-width of clear space around the mark                          |
| Minimum size        | At least 24 pixels high; use the plain product name where vectors are unavailable |
| Semantic role       | Product identity only; preserve explicit status labels and their own colors       |
| Surface             | Check light and dark legibility; the documentation identity defaults to dark      |

## [Tokens and roles](https://lanej.io/charmed/brand.html#tokens-and-roles)

Charmed uses [**Sage Midnight**](https://lanej.io/design-system/) for its native dark theme and this guide. The talisman takes the sage accent, keyboard focus uses lavender, and selected items use a deeper green surface. The swatches show the values used in the terminal.

Use [`theme.New(theme.Dark)`](https://github.com/lanej/charmed/blob/cc0823efe6b1be9d1911e813477912332826207b/theme/theme.go) for the same palette in a native widget. The guide takes its syntax colors from `Theme.Syntax` and `Theme.Lexical`, its active and inactive tab colors from `Theme.Tabs`, and its chart colors from `Theme.Series`.

## [Typography and code](https://lanej.io/charmed/brand.html#typography-and-code)

Use system sans-serif for documentation and labels. Use bundled JuliaMono for terminal snapshots, source listings, and short inline references. Source examples should use lexical highlighting, while terminal snapshots should retain the ANSI colors emitted by the native fixtures.

## [Composition](https://lanej.io/charmed/brand.html#composition)

Lead with the terminal output. Show alternate states in adjacent variant tabs and a short consumer-API snippet immediately afterward. Display the consumer body without its package or outer function, link required imports and complete source, and include the exact catalog-render command. The same example produces the native output; interactive workbench fixture commands belong in the contributor guide.

## Web palette

**Canvas**`#0d1513`

**Surface**`#17231e`

**Border**`#6b8575`

**Text**`#d7dfd9`

**Muted**`#a3b1a7`

**Accent**`#acc8b7`

## Typography

Display · system sans / 700**Observe. Compare. Act.**

Body · system sans / 400

Semantic meaning remains legible without color. Keep annotations close to native output.

Code · JuliaMono / 400

```go
theme.New(theme.Dark) // semantic colors
```

## Native semantic colors

The browser identity is intentionally quieter than terminal data. Native colors belong to the Go `theme` package and must not be substituted with web brand accents.

OK Success `#a6cd82`

INFO Information `#8cc5cc`

WARN Warning `#e8c88b`

ERROR Danger `#f18c96`

SERIES Comparison `#a798c9`
