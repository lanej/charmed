# Static guide design

## Purpose and task

Show Go developers the actual components first, then help them understand composition and run native examples. The front page is a showcase, with full screens, shared status atoms and a curated chart gallery. Related components share capability sections so readers can compare them without opening a page for each widget.

## Principles and tradeoffs

Use a compact guide navigation and readable prose. Keep examples beside their name, description and native command. Preserve complete terminal cell geometry with horizontal scrolling inside frames; the document itself must fit mobile widths. Saved terminal frames retain ANSI colors as static HTML and inherit the guide’s night surfaces; native commands provide truecolor and full behavior.

## Visual system and ownership

`workbench/site/style.css` owns the guide's palette and typography. Preserve the existing green-black background, pale text and muted green links. Prose uses system sans-serif; source and terminal cells use monospace. Theme semantics inside native applications remain owned by the Go theme package.

## Behavior and resilience

The generated guide uses HTML, CSS and local images with relative links. Navigation and content work without JavaScript or a preview service. Native details elements expand saved variants and source with keyboard access. A skip link reaches the main content; focused links and summaries have visible outlines. Missing inputs fail generation with an error. No loading or live connection states are needed for these static pages.

## Requirements and verification

### GUIDE-001

Every registered fixture must appear once in foundations, components or examples, with a native command and saved frame. Go generator tests check fixture presence, local links, anchors and repeatability. No client scripts or form controls belong in published output.

### GUIDE-002

Navigation and content must fit desktop and mobile viewports. Terminal frames and wide tables may scroll internally to preserve cell geometry. Viewrule checks main content and screenshots support manual layout review. Static snapshots use the ANSI 256-color profile; native truecolor output can differ slightly.

## Evidence and unresolved decisions

The user requested removal of interactive documentation and consolidation of per-component pages. Six guide pages and capability grouping implement that direction; they are proposed organization, without recorded user visual approval. The user requested consistent night mode across the guide. Embedded examples use saved native text on the shared night palette; historical browser screenshots are not embedded. Live deployment verification requires publishing the resulting artifact.

### GUIDE-003

Night mode is fixed across page chrome, prose, code, disclosures and saved output, independently of the operating system appearance. Embedded screenshots must not introduce light browser chrome or a different surface palette. Viewrule checks dark color-scheme and surface colors; visual review checks the examples. Native CLI theme options remain available independently.

### GUIDE-004

The home page leads with a composed native service screen and representative atoms/charts. Guide navigation stays in the sidebar; the coverage inventory lives on the development page. Terminal colors, pill fills and labels survive static export. Service lists use status atoms without redundant service/state headers; selection and keyboard behavior remain native.

Base terminal text uses neutral off-white and gray, including after ANSI 256-color conversion. Lavender is reserved for data-series differentiation, not ordinary service names or neutral status labels.

Legends are optional and absent by default. When enabled, each marker or series key reuses the style of its plotted counterpart; neutral prose does not replace the visual key. Warning and error boundaries have separate colored keys. Mean keys retain each series color, and selection retains its neutral marker.
