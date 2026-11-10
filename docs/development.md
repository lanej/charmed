## Use Charmed

Start with [Getting started](getting-started.html) to install Charmed and run an editable field. This contributor guide covers adding capabilities, checking their documentation and publishing the verified site. The [documentation standard](documentation-standard.html) defines the content and rendered bounds used throughout the guide.

| Work | Entry point |
| --- | --- |
| Use a public component | [Component catalog](atoms.html) and its consumer example |
| Coordinate a complete screen | [Recipes](recipes.html) and the composition guide |
| Change library behavior | Native Go packages, registry stories and retained-behavior tests |
| Change the guide | Markdown, consumer examples and the shared site template |

## Add a component

Organize Go packages by capability. Atomic categories describe composition depth rather than package layers. Implement a wrapper where useful defaults, semantics or reusable composition justify one, then document the native API and any lifecycle operation that reapplies customization.

| Required piece | Purpose |
| --- | --- |
| Native implementation | Preserve deliberate upstream capabilities described in ARCHITECTURE.md |
| Registry story | Set Level, Family and actual Composes dependencies; duplicate IDs, cycles and upward-level dependencies are rejected |
| Consumer example | Add docs/snippets/story_id.go with explicit sample data, construction and rendering through public APIs |
| Renderer registration | Add the example to docs/snippets/render.go; the displayed body produces the default Example frame |
| Integration guidance | Document data updates, focus, sizing and customization in docs/component-guidance.json |
| Page profile | Classify new Markdown in docs/standards.json and use the shared template |

Keep fixtures independent of production services and application decisions. Consumer examples are compiled and linted with the library. Missing examples fail generation, and the output must be visually useful after initialization. Fixture variants explore states beyond the minimal example; implementation source stays separately linked.

## Verify changes

Install just and ShellCheck, then use the checked-in targets for local work. The fast source gate checks complete documentation sources without compiling Go; the optional staged-tree hook runs that same gate before a commit, and CI enforces it independently.

```sh
just hooks
just docs-check
just lint
just ratchet
just check
just docs
```

| Gate | What it establishes |
| --- | --- |
| docs-check | Paragraph and run budgets, example length, required task sections and declared profiles |
| lint | Formatting, vet, module tidiness, strict Go lint and ShellCheck |
| ratchet | Remove resolved findings; never grow the existing baseline |
| check | Full lint, uncached race tests and reachable-vulnerability checks |
| docs | Source checks, native generation, component structure and local file/anchor validation |
| Responsive CI | Browser overflow, content measure, artifact placement, visible examples, touch, keyboard and no-JavaScript navigation |

The full Go quality gate runs on Linux and macOS. Bounds and diagnostics support review, rather than proving the example is understandable. Inspect the chart, input and inspector together at desktop and mobile sizes, then retain representative screenshots when changing the shared template.

## Publish documentation

Every same-repository documentation PR receives an isolated preview under lanej.io/charmed/pr-preview/pr-number. The publisher verifies the live revision before updating direct links in the PR description. Review those pages before merge; previews refresh on push and disappear after closure.

| Publication boundary | Behavior |
| --- | --- |
| Pull request | Build, validate and measure the static guide; retain browser evidence |
| Same-repository preview | Publish an isolated no-index snapshot and verify its revision |
| Fork pull request | Run checks without privileged preview publication |
| Main branch | Publish the tested bundle through GitHub Pages |

The small browser script enhances variant tabs, mobile navigation and legacy bookmarks. Ordinary document links and navigation remain usable without JavaScript. No separate documentation server or browser build tool is required to publish the generated guide.

## Coverage

The inventory below distinguishes implemented native components from optional integrations. The bundled JuliaMono font preserves braille and box-drawing geometry; its [license](assets/JuliaMono-LICENSE.txt) travels with the generated assets and remains available from the published site.
