## Use Charmed

Start with [Getting started](getting-started.html) to install Charmed and run an editable field. This contributor guide covers adding capabilities, checking their documentation and publishing the verified site. The [documentation standard](documentation-standard.html) defines the content and rendered bounds used throughout the guide.

| Work | Entry point |
| --- | --- |
| Use a public component | [Component catalog](atoms.html) and its consumer example |
| Coordinate a complete screen | [Recipes](recipes.html) and the composition guide |
| Change library behavior | Native Go packages, registry stories and retained-behavior tests |
| Change the guide | Markdown, consumer examples and the shared site template |

## Inspect a native fixture

The workbench uses Charmed's [CLI help formatter](molecules-commands.html#cli/help). Compare `go run ./cmd/workbench --help` with `go run ./cmd/workbench help capture` to see the root's commands and a subcommand's flags.

The documentation snapshots are deterministic captures of real Go components, not browser recreations. Use the terminal runner when you need to inspect keyboard behavior, animation, focus, or resizing.

```sh
go run ./cmd/workbench list
go run ./cmd/workbench run screen/services
```

Run from the repository root with **Go 1.26.11+**. Use Ctrl+C to exit. You can pass fixture-specific options such as `--theme light`, `--legend`, or `--columns 40`; see the [component registry](https://github.com/lanej/charmed/blob/main/story/catalog.go) for available fixtures.

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

## Git fixtures

Run `go run ./cmd/workbench run git/browser` for the composed repository browser, or use `git/history`, `git/changes` and `git/worktrees` for individual organisms. All data is synthetic and supplied in memory; no Git executable or repository access is needed. Ctrl+Left/Right switches browser sections, Tab switches panes, Enter inspects a selected worktree, and Ctrl+E/Ctrl+Y scrolls history continuation rows independently of commit selection.

Saved native previews: [branching history](assets/git-history.png), [status and diff inspector](assets/git-changes.png), [worktree list](assets/git-worktrees.png), and [repository browser](assets/git-browser.png).

## Cohort measurements

```sh
go run ./examples/cohorts
go run ./examples/catalog screen/cohorts
```

The regional example supplies temperature, energy and yield for three cohorts. Read the [cohort contract](cohorts.html) for scalar aggregation, missing observations, optional references and migration from two-cohort comparisons.
