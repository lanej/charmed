# Upstream APIs are part of Charmed's contract

Status: accepted.

Charmed deliberately exposes Bubble Tea and the underlying component libraries. Bubble Tea messages, commands and lifecycle methods, concrete upstream models, options, styling types and chart data types may appear in public APIs. Consumers can use native capabilities directly; dependency independence is not a goal.

Thin constructors return native models where shared defaults are sufficient. Wrappers add Charmed behavior and composition while retaining native access. Do not introduce parallel message systems, generic widget interfaces or replacement data types solely to conceal an upstream dependency.

Charmed owns semantic themes, cell geometry, focus conventions, status meaning and composed behavior. Application services and business decisions stay outside the UI library. Pure analysis and policy packages such as `cohort` and `attention` remain independent of terminal rendering.

Exposed upstream types are part of our compatibility surface. Review dependency upgrades for changes to public types, methods, options and behavior, and document migration requirements when consumers are affected. Supporting a different terminal framework would require a separate design decision.

Native access carries lifecycle responsibilities. A wrapper may rebuild its underlying model during data updates or resize; direct mutations are not automatically persistent. Document what is retained or reset and how callers should apply customization. Preserve documented guarantees and verify them through the wrapper's public operations. Existing preservation gaps remain defects to address; this decision does not establish guarantees the implementation lacks.

Shared visual geometry is part of composition: controls reserve one horizontal cell on each side, and panels include those content gutters inside their declared allocation. Use `layout.PanelContentSize` or `chart.PanelPlotSize` when allocating native children. Theme constructors apply named glyphs and matching focus insets as defaults; concrete native styles and glyph setters remain available to callers. The guide exports completed static cells with exact RGB rather than emulating live terminal control sequences.
