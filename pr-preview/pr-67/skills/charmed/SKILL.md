---
name: charmed
description: Build Go terminal interfaces with Charmed components, declared layouts and native Bubble Tea lifecycle contracts. Use when an application adopts or composes Charmed UI components.
---

# Charmed

Use this skill for applications consuming Charmed. For changes to the library or its website, use [Charmed contributor guidance](https://lanej.io/charmed/skills/charmed-contributing/SKILL.md). Follow the consuming application's engineering and documentation conventions.

Use the current [guide](https://lanej.io/charmed/index.md), [getting started](https://lanej.io/charmed/getting-started.md) and [component catalog](https://lanej.io/charmed/atoms.md) to choose public APIs. Every website page has a matching `.md` URL, including component-family pages with complete consumer examples and native output. [The Markdown index](https://lanej.io/charmed/llms.txt) lists them. Fetch relevant pages rather than assuming APIs from memory; check the application's installed Charmed revision when it differs from the published guide.

## Compose native components

Choose the smallest existing component that fits the task. Atomic levels describe composition depth; Go packages group capabilities. Preserve Bubble Tea messages and commands, native model access, styling and chart types. Applications own service access, Git operations, telemetry retrieval and business decisions; supply observations and handle emitted actions outside the UI library.

Use a shared `theme.Theme`, named glyphs and independent semantic state families. Unicode is the primary presentation; choose ASCII explicitly when needed. Keep health, availability, severity, rollout, focus and selection distinct. Preserve missing observations separately from zero. Read the [composition guide](https://lanej.io/charmed/examples.md) and relevant component-family page for consumer examples and lifecycle guidance.

## Allocate and update

Use declared `layout` regions and `tui.Screen` for complete screens. Applications own component sizing, focus routing and message dispatch. Cover every visible region, validate responsive branches and use only local design exceptions with an explicit reason. Preserve active inspection and selection through narrow layouts; inspect real output at small and large allocations.

Return initialization, focus and update commands to Bubble Tea, retain returned models and serialize changes on the update thread. Read each wrapper's lifecycle contract before native customization: data, theme or full-refresh resizing can reapply presentation. Geometry-only resizing reuses prepared content where documented. Avoid parallel widget interfaces that hide native APIs.

## Retain presentation

Use `runtime.Root` when its scheduling and scene coordination fit the application; composing components in your own Bubble Tea root is also supported. Follow [retained rendering](https://lanej.io/charmed/rendering.md). `runtime.Root` processes updates in order and schedules presentation on demand; `Root.View` reads the completed frame. Forward all messages and returned commands when wrapping a root, and close owned resources. For synchronous snapshots, call `PresentNow` and handle its error before reading `View`; it does not execute application commands.

After direct edits to cached Git or JSON presentation, invalidate the component and request a root frame. Advance application-owned region or inspector-pane revisions after every visible change. An optional scene `PresentationRevision` must include focus, scrolling, cursor, theme and styling as well as data; omit it when that cannot be guaranteed. Keep caches bounded to documented content and allocation.

## Verify and document

Run the consuming application's established checks. Exercise keyboard input, resizing, independent scrolling, native customization and snapshots through public APIs. Inspect native output to judge visual quality. Use [measurement instructions](https://lanej.io/charmed/development.md#measure-rendering) for resource-sensitive changes; separate generated-frame costs, terminal output and painting.

When describing an integration, explain the APIs and lifecycle responsibilities relevant to the consuming application and cite the matching library documentation. Installing this skill supplies guidance; it does not install Go dependencies, require contributor tooling, change project policy or authorize publication.
