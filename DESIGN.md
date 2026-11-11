# Charmed design philosophy

## Purpose and task

Charmed makes complex terminal workflows feel legible, tactile and expressive. It gives developers a coherent visual language for building interfaces that help people understand what is happening, inspect the evidence and act with confidence.

Beauty, personality and careful details are product goals alongside legibility and correctness. The terminal's cell grid, keyboard interaction, typography, color and glyphs are our design materials. Use them with ambition. A Charmed interface should feel intentionally composed from its smallest indicator to its complete screens.

This philosophy governs native components, their compositions and the reference guide. [ARCHITECTURE.md](ARCHITECTURE.md) defines API and implementation ownership; [the guide design](docs/guide-design.md) carries the documentation site's detailed requirements.

## Principles and tradeoffs

### Give the interface a recognizable character

Choose expressive controls, distinctive state marks, considered surfaces and coherent iconography. A switch should have a track and thumb; a selection should have a recognizable selected shape; an action can carry an appropriate icon. Reuse a visual vocabulary across components so that people learn it once. Simple atoms deserve the same care as a complete screen.

Unicode glyphs are the default visual language. Design the primary experience with the best supported terminal presentation, then provide explicit fallbacks. Hardcoded `[x]`, `(*)` and `[ON]` forms belong to the ASCII variant. A new component that only implements those forms has an unfinished primary presentation.

Visual richness needs hierarchy. Keep ordinary content calm so that icons, active controls, selected observations and important changes have room to stand out. Use repeated decoration deliberately; a badge, border or icon should contribute to the composition rather than appear everywhere by habit.

### Give different meanings different appearances

Health, severity, lifecycle, connectivity, availability, selection and keyboard focus are independent concepts. Preserve those distinctions in the model and in the rendered interface. Power-off, inactive, disconnected and disabled deserve distinct treatments. ERROR and FATAL deserve distinct severity presentations. A focused control can be off; a selected observation can contain a failure.

Color, shape and text work together. Use semantic colors and named glyphs consistently, and retain the labels needed to interpret them. Bright, normal and dim emphasis control prominence independently of meaning. Brand accent identifies interaction and hierarchy; health and outcome colors describe supplied state. Categorical chart colors identify series. Quantity alone does not justify a warning color.

### Make scanning fast and inspection trustworthy

Put the current task and its evidence first. Align comparable values, retain useful units and precision, group related information and remove redundant headings. Use space to establish relationships while keeping working screens compact enough to compare related observations without constant navigation.

A compact view can summarize, abbreviate or fold detail, provided that it signals the omission and preserves access to the original observation. Keep missing, invalid and redacted values distinct from valid empty strings, zero and false. Preserve small nonzero values. Distinguish content hidden by the display from source data that was never captured completely.

Never infer completion from intention. Busy, pending, canceled, timed out and succeeded communicate different facts. Disabled actions explain their unavailability. Freshness is separate from health. Display the evidence the application supplies without inventing a stronger conclusion.

### Make interaction feel stable and direct

Keyboard operation is a first-class experience. Focus must be visible, selection must remain understandable, and navigation must be predictable. Changes to focus, emphasis or control state should preserve geometry within the same presentation profile. Resizing and background updates should retain the user's place wherever the component's contract permits it.

Controls should look operable and reveal their state before activation. Feedback should appear beside the action or observation it explains. Motion can convey ongoing work and transitions; it must respect pause and reduced-motion preferences and must not become the only carrier of meaning.

### Compose a system people can recognize

Settle foundations and atoms before building larger compositions. Shared colors, glyphs, spacing, insets and focus behavior should make independently implemented components feel related. A table, inspector, form and chart should carry the same meaning for the same state.

Compose existing atoms where they fit. Add a primitive when it has a distinct responsibility or can remove a repeated visual decision. Atomic categories organize the catalog; capability boundaries organize the Go packages. Components expose native Charm capabilities and leave application policy and effects with their owners.

## Visual system and ownership

The `theme` package owns the semantic palette, named glyphs and shared geometry. Components own their visual structure and state presentation. Applications supply observations, policies and actions. The guide renders the actual Go components and derives its palette from the native theme. Keep implementation detail in the relevant component and architecture documentation.

Use neutral off-white and gray for ordinary terminal content, the accent for interaction, semantic hues for meaning and separate categorical colors for chart series. Preserve hue and readable contrast when reducing emphasis. Theme and profile variations should feel like the same system.

The current [control specimens](docs/assets/control-glyphs.png) and [atom and state specimens](docs/assets/data-atoms.png) illustrate this direction. They are implementation examples for comparison, not a freeze on exact glyph choices or a claim of visual approval for every detail.

## Behavior and resilience

Treat glyph repertoire, color depth and motion as separate presentation choices. A limited color profile does not imply ASCII. Select ASCII for an explicit preference or an established capability constraint; generic portability concerns must not silently replace the Unicode design. Preserve meaning and usability in each fallback while allowing the primary presentation its full character.

Terminal-cell geometry is a design constraint. Measure rendered text, including wide and combining characters, styled content and adornments. Allocate enough space for the real control. Narrow layouts should preserve the most useful content and signal what cannot fit; use deliberate reflow, disclosure and internal scrolling as appropriate.

## Requirements and verification

Review actual rendered components when changing visual behavior. Show primary Unicode specimens first, then the relevant ASCII, monochrome, color-depth, emphasis, focus and disabled variants. Check the affected components at useful narrow and wide allocations. Tests protect semantics, geometry and interaction contracts; screenshots expose visual choices that those tests cannot judge.

For each new or materially changed component, check that its primary presentation expresses the intended character, its visual states remain distinct, its shared atoms and tokens are reused, and its fallbacks were chosen deliberately. Do not hand off an ASCII-only primary implementation as finished design work.

The reference guide must demonstrate the real result with small consumer examples and faithful native output. Keep it useful for comparing states, understanding composition and finding ownership. Detailed site requirements live in [the guide design](docs/guide-design.md).

## Evidence and unresolved decisions

The established direction is expressive Unicode-first terminal design with semantic precision, stable interaction and explicit fallbacks. The repeated ASCII-only implementations show that fallback support was being mistaken for the product's visual ambition; future work must use the principles above when choosing its primary presentation.

Exact glyph families, visual treatments and compositions can evolve through rendered review. Distinguish an explored candidate from an accepted reference. Record concrete feedback against the specimen being reviewed, and update these principles when that feedback changes the product direction.
