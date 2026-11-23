---
name: charmed-contributing
description: Modify and review the Charmed Go terminal UI library, its native examples and reference website. Use for contributions to Charmed itself, rather than applications consuming the library.
---

# Contribute to Charmed

Use this skill in a Charmed source checkout. Read that checkout's `AGENTS.md` and explicitly selected project guidance before editing; the checked-out revision owns contributor policy. For application development, use the [Charmed usage skill](https://lanej.io/charmed/skills/charmed/SKILL.md). This package supplies guidance only and does not install tools, replace repository policy or authorize publishing changes.

## Establish the contract

Read `ARCHITECTURE.md` for native API ownership and `DESIGN.md` before component, visual or interaction changes. Their canonical sources are the [architecture contract](https://raw.githubusercontent.com/lanej/charmed/main/ARCHITECTURE.md) and [design philosophy](https://raw.githubusercontent.com/lanej/charmed/main/DESIGN.md). Preserve concrete upstream messages, commands, models, styling and chart capabilities. Charmed owns UI semantics and composition; applications own services, Git operations and business decisions.

Keep changes within the relevant capability. Make intended behavior and failures observable through public operations. Document wrapper lifecycle effects on native customization, including what data, theme and resize operations retain or reset. Use existing contracts and independent acceptance examples to protect retained behavior.

## Build components and examples

Use the native theme, shared glyphs, semantic state families and cell geometry. Lead with expressive Unicode, with ASCII selected explicitly. Keep focus, selection and semantic states distinct; inspect real native output at useful narrow and wide sizes, including keyboard interaction and resizing.

Follow [the contributor guide](https://lanej.io/charmed/development.md#add-a-component) for the implementation, registry story, compiled consumer example, renderer registration and lifecycle guidance. Prefer existing atoms before adding a composition. For screen layouts, validate declared regions and responsive branches; keep exceptions local with an explicit reason.

## Protect rendering efficiency

Follow [retained rendering](https://lanej.io/charmed/rendering.md) for ordered updates, scheduled presentation, explicit snapshots and native invalidation. Keep component caches bounded and account for real content, geometry and interaction dependencies. Preserve selection, independent scrolling and native customization when optimizing.

Use the [performance measurement instructions](https://lanej.io/charmed/development.md#measure-rendering) and existing regression tests. Separate startup, warm views, data replacement, width changes and height-only resizing. Run comparative timing serially on a consistent host; distinguish generated frames, terminal bytes and actual painting. Timing reports do not replace deterministic behavior checks.

## Verify and document

Use the checkout's `justfile`: run `just lint` while editing, `just ratchet` after fixes, `just check` before handoff, and `just docs` for guide or fixture changes. Run `just shellcheck` for shell changes. Preserve existing gates and shrinking lint allowances; fix findings at their cause without enlarging baselines or adding blanket suppressions. Report checks that cannot run and remaining failures accurately.

Follow `docs/guide-design.md` and the [documentation standard](https://lanej.io/charmed/documentation-standard.md). Describe current behavior and usage; retain implementation history and audit reports outside the website. Preserve complete consumer examples, meaningful native output and lifecycle guidance. Every guide page has matching HTML and Markdown; each advertised skill includes all referenced local resources and must work when downloaded in isolation.

Run `just docs-check` for fast source feedback. Do not raise content bounds to accommodate violations. Review chart, input and inspector examples together at mobile and desktop widths when changing shared templates. Retain useful evidence and verify the exact published preview revision when delivering documentation; follow the checkout's authorized review and publication workflow.
