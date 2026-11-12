## The terminal talisman

The Charmed mark combines a terminal with a lucky charm. The terminal window becomes the body of a small talisman: an attachment loop above, a faceted outline below, and a shell prompt inside.

| Form | Meaning |
| --- | --- |
| Attachment loop | A charm that can be carried or attached |
| Faceted frame | The shape of a talisman, enclosing one compact object |
| Window rail and `>_` prompt | A native terminal interface |

The complete silhouette carries the idea in one color. The mark uses the website's accent through `currentColor`; the reversed specimen uses its existing canvas and text tokens. Terminal data keeps its own semantic colors.

## Use the identity

The primary lockup pairs the terminal talisman with the lowercase **charmed** wordmark. Use it for documentation, the website, and repository art. The mark can stand alone when the product name is already visible.

| Identity constraint | Requirement |
| --- | --- |
| Source | Use the supplied SVG and retain the loop and faceted outline |
| Clearance | At least one cursor-width of clear space around the mark |
| Minimum size | At least 24 pixels high; use the plain product name where vectors are unavailable |
| Semantic role | Product identity only; preserve explicit status labels and their own colors |
| Surface | Check light and dark legibility; the documentation identity defaults to dark |

## Tokens and roles

The swatches below are **derived directly from Charmed's native dark theme**. Generated CSS and this page draw from the same semantic roles used in terminal fixtures, including surfaces, borders, text, statuses, and data-series colors. Syntax highlighting maps to these roles instead of defining another hex palette.

For native widgets, source semantic colors from [`theme.New(theme.Dark)`](https://github.com/lanej/charmed/blob/main/theme/theme.go). Documentation consumes those values during static generation rather than maintaining a second palette. Success, information, warning, danger, and data-series colors remain distinct.

## Typography and code

Use system sans-serif for documentation and labels. Use bundled JuliaMono for terminal snapshots, source listings, and short inline references. Source examples should use lexical highlighting, while terminal snapshots should retain the ANSI colors emitted by the native fixtures.

## Composition

Lead with the terminal output. Show alternate states in adjacent variant tabs and a short consumer-API snippet immediately afterward. Display the consumer body without its package or outer function, link required imports and complete source, and include the exact catalog-render command. The same example produces the native output; interactive workbench fixture commands belong in the contributor guide.
