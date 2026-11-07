# Static guide design

## Purpose and task

Show Go developers the actual components first, then help them understand composition and run native examples. The front page is a showcase, with full screens, shared status atoms and a curated chart gallery. Related components share capability sections so readers can compare them without opening a page for each widget.

## Principles and tradeoffs

Use a compact guide navigation and readable prose. Each fixture has a sidebar link. Lead with the saved native visual, put variants in keyboard-accessible tabs, and immediately follow with highlighted Go source and the native command. Preserve complete terminal cell geometry with horizontal scrolling inside frames; the document itself must fit mobile widths. Saved terminal frames retain ANSI colors as static HTML and inherit the guide’s night surfaces; native commands provide truecolor and full behavior.

## Visual system and ownership

`workbench/site/style.css` owns the guide's palette and typography. Preserve the existing green-black background, pale text and muted green links. Prose uses system sans-serif; source and terminal cells use monospace. Theme semantics inside native applications remain owned by the Go theme package.

## Behavior and resilience

The generated guide uses HTML, CSS and local images with relative links. Navigation and source remain readable without JavaScript or a preview service. A small local script enhances semantic variant tabs with roving keyboard focus; without it, all variant snapshots stay visible. Source is never hidden behind a disclosure. A skip link reaches the main content; focused links and summaries have visible outlines. Missing inputs fail generation with an error. No loading or live connection states are needed for these static pages.

## Requirements and verification

### GUIDE-001

Every registered fixture must appear once in foundations, components or examples, with a native command and saved frame. Go generator tests check fixture presence, local links, anchors and repeatability. The only client script is the static variant-tab enhancer; all documentation, navigation, code highlighting and snapshot rendering happen at build time.

### GUIDE-002

Navigation and content must fit desktop and mobile viewports. Terminal frames and wide tables may scroll internally to preserve cell geometry. Viewrule checks main content and screenshots support manual layout review. Static snapshots use the ANSI 256-color profile; native truecolor output can differ slightly.

## Evidence and unresolved decisions

The user requested removal of interactive documentation and consolidation of per-component pages. Seven guide pages, a brandsheet and capability grouping implement that direction; they are proposed organization, without recorded user visual approval. The user requested consistent night mode across the guide. Embedded examples use saved native text on the shared night palette; historical browser screenshots are not embedded. Live deployment verification requires publishing the resulting artifact.

### GUIDE-003

Night mode is fixed across page chrome, prose, code, disclosures and saved output, independently of the operating system appearance. Embedded screenshots must not introduce light browser chrome or a different surface palette. Viewrule checks dark color-scheme and surface colors; visual review checks the examples. Native CLI theme options remain available independently.

### GUIDE-004

The home page leads with a composed native service screen and representative atoms/charts. Guide navigation stays in the sidebar; the coverage inventory lives on the development page. Terminal colors, pill fills and labels survive static export. Service lists use status atoms without redundant service/state headers; selection and keyboard behavior remain native.

Base terminal text uses neutral off-white and gray, including after ANSI 256-color conversion. Lavender is reserved for data-series differentiation, not ordinary service names or neutral status labels.

Legends are optional and absent by default. When enabled, each marker or series key reuses the style of its plotted counterpart; neutral prose does not replace the visual key. Warning and error boundaries have separate colored keys. Mean keys retain each series color, and selection retains its neutral marker.


### GUIDE-005

Every fixture has a sidebar entry linking to its exact section anchor. A component's first visible content is its saved terminal preview, with all snapshots available as semantic tabs. Its highlighted Go source follows the preview directly, not inside a disclosure. Tabs support arrow keys, Home and End; the original static frames remain visible if JavaScript is unavailable.

### GUIDE-006

Charmed has a documented identity with its own green-black web palette, lowercase lockup and code typography. Brand colors never replace the native semantic terminal palette. Build-time code highlighting escapes source HTML and maintains distinguishable keyword, string, comment and numeric roles.
