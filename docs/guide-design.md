# Static guide design

This document applies the [Charmed design philosophy](../DESIGN.md) to the reference site. Product and component decisions start with that philosophy.

## Purpose and task

Show Go developers the actual components first, then help them understand composition and run native examples. The front page is a showcase, with full screens, shared status atoms and a curated chart gallery. Related components share capability sections so readers can compare them without opening a page for each widget.

The website describes the current supported product: available APIs, ownership, lifecycle, constraints and runnable usage. Write directly for someone using it today. Keep migration narratives, implementation history, PR chronology, rollout plans and before/after audit reports in Git, PRs or developer benchmark artifacts rather than the website.

## Principles and tradeoffs

Use a compact guide navigation and readable prose. Each fixture has a sidebar link. Lead with the saved native visual, put variants in keyboard-accessible tabs, and immediately follow with highlighted Go source. Link each component to its runnable consumer example. Keep fixture setup and workbench flags in the development guide. Preserve complete terminal cell geometry with horizontal scrolling inside frames; the document itself must fit mobile widths. Saved terminal frames preserve exact RGB as static HTML on the guide’s night surfaces; native commands provide keyboard interaction and animation.

| Reader task | Guide entry |
| --- | --- |
| See what a component renders | Open its sidebar anchor and inspect the default Example tab. |
| Compare states, palettes or widths | Switch between saved native frames. |
| Try the example | Use the displayed Go source and the catalog command beneath its integration guidance. |

## Visual system and ownership

The Go `theme` package owns the canonical palette for native applications and the guide. The site generator derives CSS variables and brand swatches from `theme.New(theme.Dark)`; `workbench/site/style.css` owns layout and typography without duplicating hex values. Prose uses system sans-serif; source and terminal cells use JuliaMono. Truecolor previews preserve the native foreground and fill values.

## Behavior and resilience

The generated guide uses HTML, CSS and local images with relative links. Navigation and source remain readable without JavaScript or a preview service. A small local script enhances semantic variant tabs with roving keyboard focus; without it, all variant snapshots stay visible. Source is never hidden behind a disclosure. A skip link reaches the main content; focused links and summaries have visible outlines. Missing inputs fail generation with an error. No loading or live connection states are needed for these static pages.

| Surface | Owner and fallback |
| --- | --- |
| Terminal colors | The native theme supplies exact RGB. Saved variants show reduced color profiles. |
| Page layout | Site CSS sets typography and spacing. Frames and tables can scroll within their allocation. |
| Navigation | Ordinary links work without scripts. Keyboard-accessible tabs enhance the saved variant list. |

## Requirements and verification

### GUIDE-001

Every registered fixture must appear once in foundations, components or examples, with a saved frame and code example. Go generator tests check fixture presence, local links, anchors and repeatability. Client scripts enhance variant tabs, mobile navigation and legacy bookmarks; all documentation, navigation, code highlighting and snapshot rendering happen at build time.

### GUIDE-002

Navigation and content must fit desktop and mobile viewports. Terminal frames and wide tables may scroll internally to preserve cell geometry. Viewrule checks main content and screenshots support manual layout review. Static snapshots use truecolor as the reference. Separate 256-color, 16-color and monochrome variants expose fallback behavior.

| Verification | Evidence |
| --- | --- |
| Generated guide | Registered fixtures, local links, anchors and example output. |
| Browser layout checks | Page overflow, visible code, native geometry and navigation at mobile and desktop widths. |
| Native runner | Keyboard input, focus transitions, resizing and animation. |

## Evidence and unresolved decisions

Native fixtures and compiled consumer examples provide the guide's evidence. Saved frames demonstrate specific states and allocations; native interaction checks and rendered review provide separate evidence. Embedded examples use saved native text on the shared night palette. Live deployment verification requires publishing and checking the artifact's revision.

### GUIDE-003

Night mode is fixed across page chrome, prose, code, disclosures and saved output, independently of the operating system appearance. Embedded screenshots must not introduce light browser chrome or a different surface palette. Viewrule checks dark color-scheme and surface colors; visual review checks the examples. Native CLI theme options remain available independently.

### GUIDE-004

The home page leads with a composed native service screen and representative atoms/charts. Guide navigation stays in the sidebar; the coverage inventory lives on the development page. Terminal colors, pill fills and labels survive static export. Service lists use status atoms without redundant service/state headers; selection and keyboard behavior remain native.

| Visual element | Required presentation |
| --- | --- |
| Base terminal text | Neutral off-white or gray, including after ANSI 256-color conversion. Reserve lavender for data series. |
| Optional legend | Hidden by default. Each enabled series or marker uses its plotted style. |
| Boundary and mean keys | Warning and error boundaries use separate colors, mean keys retain each series color, and selection keeps its neutral marker. |


### GUIDE-005

Every fixture has a sidebar entry linking to its exact section anchor. A component's first visible content is its saved terminal preview, with all snapshots available as semantic tabs. Its highlighted Go source follows the preview directly, not inside a disclosure. Follow the code with lifecycle guidance and one command that runs that consumer example. Tabs support arrow keys, Home and End; the original static frames remain visible if JavaScript is unavailable.

### GUIDE-006

Charmed has a documented identity with a lowercase lockup and code typography. Brand swatches and page CSS variables derive from the canonical native theme; component previews preserve those native colors. Build-time code highlighting escapes source HTML and maintains distinguishable keyword, string, comment and numeric roles.

### GUIDE-007

Documentation PRs publish an isolated Pages snapshot at `/charmed/pr-preview/pr-<number>/`. The PR description links to actual preview pages, not to itself, only after the published revision marker matches the current PR head. Reviewers should be able to inspect the rendered component catalog without merging first.

| Publication | Evidence to inspect |
| --- | --- |
| Pull request | Preview links resolve to pages built from the PR head. |
| Main branch | The published catalog contains the merged component and its runnable example. |

### GUIDE-008

Catalog Go examples demonstrate the consumer API and need not be standalone programs. Prefer statements or expressions without `package main`, import blocks, or a surrounding `func main()`. Keep curated snippets in compiled `docs/snippets/*.go` consumer examples and render their actual output beside the highlighted usage. Full fixture code remains available through the implementation link.

### FOUNDATION-001

Glyphs have names and explicit one-cell ASCII equivalents. Semantic states retain labels alongside icons. Controls, pills and panel content reserve one horizontal cell on both sides; focus and emphasis preserve the same allocation. Panel children use the shared content-size calculation. Dim semantic tokens preserve hue and readable contrast in the truecolor reference; fallback limitations are reviewed separately. Categorical series colors do not alias interaction or status colors. These requirements incorporate the user feedback on missing icons and padding, unpadded inverse focus, shared accent/series color and gray dimming.
