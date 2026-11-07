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

Use the supplied SVG so the header, lockups, and small-size specimens share the same geometry. Keep the attachment loop and faceted outline intact. Leave at least one cursor-width of clear space around the mark and keep it at least 24 pixels high. Use the plain product name where vector graphics are unavailable.

Do not use the mark as a generic status badge or replace a status label with brand green. Preserve legibility on both light and dark backgrounds; the default documentation identity is dark.

## Tokens and roles

The swatches below are **website identity tokens**, not the terminal theme. Documentation uses the canvas, surface, border, text, muted, and accent values to maintain its fixed night mode. Code highlighting uses separate syntax tokens, so keywords, strings, comments, and literals remain independently distinguishable.

For native widgets, source the semantic colors from [`theme.New(theme.Dark)`](https://github.com/lanej/charmed/blob/main/theme/theme.go), not from CSS. Success, information, warning, danger, and comparison series should remain separable without conflating their roles.

## Typography and code

Use system sans-serif for documentation and labels. Use bundled JuliaMono for terminal snapshots, source listings, and short inline references. Source examples should use lexical highlighting, while terminal snapshots should retain the ANSI colors emitted by the native fixtures.

## Composition

Lead with the terminal output. Show alternate states in adjacent variant tabs and source immediately afterward. Include the command that reproduces the example and a link to its implementation. The brand should support the component content rather than compete with it.
