# Brand sheet

Charmed is a terminal-first design system. Its identity should feel like the interfaces it creates: precise, compact, and legible. Brand color stays restrained; semantic colors in terminal fixtures carry the meaning.

## Use the identity

The primary lockup pairs a compact prompt mark (`>_`) with the lowercase **charmed** wordmark. Keep the mark and name together for the website, documentation, and repository art. The small prompt mark can stand alone when the surrounding product name is visible.

Do not use the mark as a generic status badge or replace a status label with brand green. Preserve legibility on both light and dark backgrounds; the default documentation identity is dark.

## Tokens and roles

The swatches below are **website identity tokens**, not the terminal theme. Documentation uses the canvas, surface, border, text, muted, and accent values to maintain its fixed night mode. Code highlighting uses separate syntax tokens, so keywords, strings, comments, and literals remain independently distinguishable.

For native widgets, source the semantic colors from [`theme.New(theme.Dark)`](https://github.com/lanej/charmed/blob/main/theme/theme.go), not from CSS. Success, information, warning, danger, and comparison series should remain separable without conflating their roles.

## Typography and code

Use system sans-serif for documentation and labels. Use bundled JuliaMono for terminal snapshots, source listings, and short inline references. Source examples should use lexical highlighting, while terminal snapshots should retain the ANSI colors emitted by the native fixtures.

## Composition

Lead with the terminal output. Show alternate states in adjacent variant tabs and source immediately afterward. Include the command that reproduces the example and a link to its implementation. The brand should support the component content rather than compete with it.
