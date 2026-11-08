# Foundations

Themes, cell geometry, focus and motion are shared across every component. Choose semantic roles rather than embedding color decisions in application widgets. Layout parents allocate rectangles in terminal columns and rows; focus scopes route keys to the active child.

The default saved frame is the truecolor reference. Separate tabs show narrow, 256-color, 16-color and monochrome presentations; fallback color limitations do not define the reference palette. Run with `--theme light`, `--columns 40` or `--ascii --profile none` to inspect those presentations in a terminal. The HTML renderer reads completed static cells, preserves RGB and escapes text; it does not emulate cursor movement or other live terminal effects.

## Colors and emphasis

Foreground and muted text are neutral. Accent identifies interaction; success, information, warning and danger identify semantic meaning. UNKNOWN uses a cool slate-blue tone for uncertainty; OFF uses achromatic gray for inactivity. Their normal and dim tokens remain separate, while labels and glyphs retain meaning without color. Categorical series use a separate palette and never alias those interface or status tokens. Series order remains stable across views; labels and matching legend keys retain identity when fallback colors converge.

Bright, normal and dim emphasis preserve the label, glyph and cell allocation. Dim uses explicit shades of each semantic hue instead of blending every tone toward gray or relying on the terminal's faint attribute. The built-in dim text tokens meet a 4.5:1 contrast floor against their theme background; native fonts, color profiles and terminal settings still affect appearance. Monochrome keeps labels, glyphs and bright weight; normal and dim may coincide without color.

## Glyphs and spacing

`theme.Glyph` names status, selection, expansion, navigation, loading and unavailable marks. `Theme.Glyph` selects a one-cell Unicode or ASCII representation. A glyph supplements a semantic text label; `UNKNOWN` and `OFF` remain different states. Animated spinners retain native frame sequences, with a named static loading mark for reduced motion. Application-specific chart markers remain caller-owned.

`theme.InlinePadding` reserves one horizontal cell on each side of controls, pills and panel content. `Theme.Control` and `Theme.Focus` allocate the same space. Status always composes its glyph, one separating cell and its label inside the same padded pill; emphasis changes color and weight without changing geometry. Monochrome focus retains inverse highlighting. `Theme.Selection` uses a quieter selection surface so nested table text and semantic pills remain readable; tables reuse their existing cell insets instead of adding a second focus inset.

Panel dimensions include borders and content gutters. Allocate children with `layout.PanelContentSize`; `chart.PanelPlotSize` also reserves an optional legend row. At widths below five cells, panels reduce the gutter to preserve their allocation. Consumers previously allocating only two border cells must reserve the new content gutters as well. Theme literals without explicit dim tokens retain their supplied normal tone until the caller provides dim shades; literals without a selection surface use bold and underline to retain a visible selection without assuming fill contrast.
