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

## Measure rendering

Measure repeated views separately from startup and idle behavior. The retained benchmarks construct models before timing, exercise native output and report allocations. Run comparisons sequentially on the same machine, toolchain, terminal color profile and allocation; concurrent builds and tests distort timing. Include `./layout`, `./gitui` and `./collection` to measure clipping, ignored keys and visible grid windows.

```sh
go test -c -o /tmp/charmed-render.test ./story
env -u NO_COLOR TERM=xterm-256color COLORTERM=truecolor /tmp/charmed-render.test \
  -test.run '^$' -test.bench '^BenchmarkRender$' \
  -test.benchmem -test.benchtime=200ms -test.count=3
```

Native runtime updates remain event driven. Screens reuse one cleared canvas at their current allocation, and only fixed workbench fixtures retain their source rendering. Interactive views read current model state on every render. Resizing replaces retained buffers rather than accumulating allocations for previous sizes; folded inspector bodies and unchanged Git searches avoid rebuilding content.

The November 2026 resize audit compared merged main at `eb2dc615` with these changes on an Apple M6, darwin/arm64, Go 1.27.4, truecolor. These are medians of five sequential 200 ms runs measuring resize plus output construction. Screen height changes alternate 120×40 and 120×30 cells; the 1,000-item Git pane alternates 80×20 and 80×35. Timing varies with the host; these measurements do not establish terminal-emulator painting latency.

| Resize workload | Time before → after | Allocated bytes before → after |
| --- | --- | --- |
| Git screen, repeated size | 1.32 → 0.12 ms | 949,342 → 82,544 |
| Git screen, height | 1.22 → 0.78 ms | 872,993 → 812,473 |
| Log inspector, height | 2.55 → 1.12 ms | 1,996,249 → 716,489 |
| Markdown viewport, height | 1.65 → 0.68 ms | 1,363,236 → 779,176 |
| Git pane, height | 0.89 → 0.05 ms | 409,435 → 33,359 |

The Git pane benchmark compares the existing full-refresh resize with the new geometry-only path. Its width-resize time remains approximately 0.84 → 0.82 ms because width changes still rebuild row presentation. The screen benchmark also covers width drags and breakpoint crossings; component-level gains do not imply that every screen resize avoids reflow.

Separate native PTY probes warmed each screen for 1.5 seconds, then sampled process CPU and terminal output for five seconds. Static Git and Markdown screens emitted no additional bytes and consumed roughly 0.2–0.4% of one core; the focused service screen emitted 250 bytes for cursor blinking and consumed about 2.6%. CPU counters have 0.01-second resolution, so these short samples cannot establish zero CPU. Bubble Tea's native flush ticker and default frame rate remain unchanged.

## Reuse prepared content

Use `layout.Compile` for a stable declaration and `Screen.ArrangeCompiled` for successive allocations. Compilation copies caller-owned nodes and validates their branches; resolving each new size still checks geometry and design rules. Recompile after changing the declaration. The existing `Arrange` remains suitable for mutable declarations.

`Screen.RenderRetained` accepts a map of region IDs to `tui.RegionView{Revision, Render}`. Each callback receives its allocated size. Advance the region revision after changing anything visible, including native styles, scroll, focus, data, theme or symbols. Zero is a valid revision; replacing a callback alone does not invalidate its result. Same-size moves reuse content, hidden regions release output, and switching to immediate `Render` clears retained state. Calls remain serialized by the owning model.

| Work | API and responsibility |
| --- | --- |
| Markdown preparation | `text.PrepareMarkdown(source, palette, options...)` owns source and presentation; `document.Render(width)` reflows without parsing or highlighting again |
| Git list geometry | `ResizeViewport(size)` follows selection and updates visibility; `Resize(size)` also reapplies presentation |
| JSON viewport geometry | `ResizeViewport(width, height)` avoids reprojection on height changes; call `Refresh` after direct presentation changes |
| Inspector geometry | `ResizeViewport(size)` reuses unchanged-width payloads and request rows; `Resize(size)` refreshes native presentation |
| Inspector pane caching | `ViewRetained(pattern.InspectorRevisions{Requests, Details, Logs})` retains panes independently; advance the affected revisions after mutations |
| Native access | Component views and existing full-refresh methods remain available; use `Invalidate` after direct edits to cached Git/JSON native styling |

Run `BenchmarkResizeRender` alongside `BenchmarkRender` to measure the complete resize-to-output path. It covers repeated allocations, height changes, width changes and responsive breakpoint crossings. Compare against the same benchmark source on the base revision, running each executable sequentially after other tests and builds finish. These benchmarks measure generated frames; terminal emulator painting and interactive latency need separate native measurements. For consumer upgrade steps, follow the [migration guide](migration.html).

## Batch native presentation

`runtime.Root` processes every update in order and batches presentation at 60 fps by default. `FPS` selects another positive rate; zero selects the default. One frame command remains pending while dirty, and none runs while idle. Sparse updates enqueue immediately; bursts share the remaining frame budget. Bubble Tea retains its own terminal-flush timer. Resize bursts use their final bounds; the next input sees those bounds before handling paging or selection. Scenes may expose `PresentationRevision() uint64` to identify unchanged updates. Without it, scheduling is conservative.

`Root.View` reads the completed frame; `PresentNow` and `Render` provide synchronous snapshots. Return commands from `Update` so frames run. After direct native edits, invalidate the affected component, then return `Root.Invalidate()`. Git/JSON caches retain one current output. Git input bursts defer row projection; source replacement still installs complete native content immediately. JSON selection reuses syntax until structure or presentation changes. Regression checks count preparations, bound allocations and compare fresh output, alongside ordered-burst and lifecycle tests.

## Audit catalog performance

The performance suite discovers every registered native fixture and inventories exported types, functions and view methods. Public view methods must link to a catalog workload or an explicit direct benchmark. Consumer-example references are source links, not proof of every implementation branch. Support functions without an example remain labeled separately, so the inventory does not silently equate a source reference with a measurement.

```sh
just perf --suite catalog --count 1 --benchtime 20ms
just perf --suite catalog --base origin/main --count 1 --benchtime 20ms
just perf --suite hotspots
just perf --suite profiles --count 1 --benchtime 1s
just perf --suite heap
just perf-native --samples 100 --idle-seconds 5
```

Install uv and use Python 3.11+ for these developer commands. The catalog scan is a short ranking pass; the default five 500 ms repetitions are intended for focused comparisons on a consistent machine. `--match` accepts Go's slash-separated benchmark selection expressions. Reports include raw samples, source hashes, changed-source patches, new benchmark sources, toolchain and host metadata. Choose an ignored or external `--out` directory. Baseline comparisons archive the committed library and overlay identical current benchmark sources without changing another checkout.

| Report | Coverage and interpretation |
| --- | --- |
| Catalog | Construction, first frame, warm views, height and width resizing, update dispatch, declared states, light/contrast/monochrome and ASCII variants |
| Hotspots | Collections at 100/1,000/10,000 items, payloads at 1 KB/16 KB/application limit, chart datasets at 100/1,000/10,000 points, filtering and log update batches |
| Profiles | CPU and allocation profiles for the selected benchmark packages; inspect them with Go's pprof tooling |
| Heap | Process-wide live Go heap after GC across repeated resize cycles and after close; includes runtime pools and shared initialization |
| Native PTY | Instrumented update-to-completed-frame and input-to-first-output latency, idle CPU, terminal bytes and RSS; includes static screens, cursors, spinners and timers |

Native traces exclude startup and exit. Completed-frame records identify accepted-update counts, oldest/latest input latency and frame revision; cached reads and unchanged messages remain distinct. Aggregation retains three input-kind samples during bursts. Linux CPU samples use process ticks; other hosts report counter precision. First output is not terminal painting, and RSS is not retained Go heap. The harness verifies input counts, reports absent output separately and closes owned resources on failure. Timing reports are informational; deterministic regression checks run in the quality gate.

The manually dispatched **Native performance reports** workflow retains measurements as artifacts on Linux or macOS. Interpret the short catalog scan as a way to locate costs, then repeat the selected workloads and compare profiles before attributing gains. Full native viewport contents and scroll-offset semantics remain intact; collection updates may still require document-sized work. The guide website and terminal-emulator painting are outside these native measurements.

## Catalog audit findings

The November 2026 catalog audit covered 153 fixtures, 58 public view methods and 3,008 matched workloads against `f92eac93`. The broad scan used one 20 ms sample to rank costs. The scaling figures below are medians of five 500 ms runs on an Apple M6 with Go 1.26.11 and truecolor output. Baseline libraries used identical benchmark sources. These are generated-frame and update costs, not terminal painting latency.

| Repeated workload | Before → after | Allocated bytes before → after |
| --- | --- | --- |
| 10,000-item Git selection | 4.30 → 1.95 ms | 2,249,161 → 1,107,134 |
| 10,000-item Git filtering | 1.94 → 0.72 ms | 1,304,351 → 325,395 |
| 64 KB JSON width reflow | 7.78 → 2.31 ms | 13,026,641 → 6,754,482 |
| 10,000 HTTP records replaced | 55.24 → 9.65 ms | 158,756,820 → 80,149,187 |
| Ten HTTP updates in 10,000 records | 56.64 → 10.13 ms | 162,478,417 → 83,754,556 |
| Cohort height resize, 200×60 | 17.89 → 5.15 ms | 14,400,562 → 4,621,124 |
| Deployment height resize, 200×60 | 18.91 → 5.97 ms | 15,120,577 → 4,441,394 |

Git preparation now retains normalized search labels and one current-width row projection, including continuation rows. Filtering and selection preserve the full native viewport. JSON geometry-only resizing reuses a bounded syntax projection; Refresh and Clear invalidate it. Cohort resizing reuses plots, prepares card content once, and exposes `ResizeViewport` to retain unchanged-width stacked content. Existing Resize still refreshes native presentation. HTTP label caches last only one rebuild and do not cache caller-supplied URL formatters.

Native probes processed 100 resizes each for Git and Markdown and 100 Git navigation keys without missing output. First-output resize p95 was 20.1/19.5 ms; instrumented static idle CPU was about 0.4% of one core. Neither measures completed painting. Heap samples after GC covered every fixture across four phases of twenty resize cycles. Nonmonotonic changes include shared pools and initialization; they neither establish leaks nor prove leak freedom. Raw reports retain samples and after-close values.

The large HTTP workload explicitly raises the retention limit to 10,000. It remains document-sized work and is not representative of the smaller default limit. Native chart rasterization and per-cell styling remain substantial costs, especially at large allocations; preserving native APIs does not establish a universal frame-time guarantee. Startup, width changes and data replacement remain distinct from warm or height-only rendering. Source snapshots and raw profiles accompany reports so these conclusions can be challenged and repeated.

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

Run `go run ./cmd/workbench run git/browser` for the composed repository browser, or use `git/history`, `git/changes` and `git/worktrees` for individual organisms. All data is synthetic and supplied in memory; no Git executable or repository access is needed. Ctrl+Left/Right switches browser sections, Tab switches panes, Enter inspects a worktree, Alt+L changes layout preference, and Ctrl+L displays declared region bounds. Ctrl+E/Ctrl+Y scrolls history independently of selection.

Saved native previews: [branching history](assets/git-history.png), [status and diff inspector](assets/git-changes.png), [worktree list](assets/git-worktrees.png), and [repository browser](assets/git-browser.png).

## Cohort measurements

```sh
go run ./examples/cohorts
go run ./examples/catalog screen/cohorts
```

The regional example supplies temperature, energy, yield and humidity for three cohorts. Read the [cohort contract](cohorts.html) for scalar aggregation, missing observations, optional references and migration from two-cohort comparisons.

## Nested layouts

Run these examples to inspect columns containing independently allocated sections and a grid nested inside an inspection column. Tab and Shift+Tab move focus; Ctrl+L toggles the region map so each leaf rectangle can be inspected.

```sh
go run ./cmd/workbench run template/nested-columns
go run ./cmd/workbench run template/nested-grid
```

Both examples reflow into stacked panes when narrow and tall. Shorter allocations retain the selected host inspection and explicitly omit the host list. The [nested layout examples](templates-layout.html) show complete consumer declarations and saved wide, stacked and compact frames.
