## Use Charmed

Charmed wraps native [Charm](https://charm.sh/) and ntcharts components with shared themes, terminal-cell sizing, and composition conventions. Application code owns its state and business decisions; widgets remain ordinary Go models.

Start with a component and pass the same theme to related elements:

```go
palette := theme.New(theme.Dark)
field := input.NewText(input.TextOptions{
    Placeholder: "Environment name",
    Width:       40,
    Theme:       palette,
})
cmd := field.Focus()
```

Explore the [input components](atoms-input.html#input/text), [charts](atoms-chart.html#chart/sparkline), and [complete screens](screens-pattern.html#screen/services). The catalog shows the visual first, then a focused usage snippet. Full fixture implementations are linked separately.

## Inspect a native fixture

The workbench uses Charmed's [CLI help formatter](molecules-commands.html#cli/help). Compare `go run ./cmd/workbench --help` with `go run ./cmd/workbench help capture` to see the root's commands and a subcommand's flags.

The documentation snapshots are deterministic captures of real Go components, not browser recreations. Use the terminal runner when you need to inspect keyboard behavior, animation, focus, or resizing.

```sh
go run ./cmd/workbench list
go run ./cmd/workbench run screen/services
```

Run from the repository root with **Go 1.26.11+**. Use Ctrl+C to exit. You can pass fixture-specific options such as `--theme light`, `--legend`, or `--columns 40`; see the [component registry](https://github.com/lanej/charmed/blob/main/story/catalog.go) for available fixtures.

## Add a component

Organize Go packages by capability rather than atomic level. Use atoms, molecules, and organisms to describe **composition depth in the catalog**, not to create extra package layers.

1. Implement a wrapper only where it establishes useful defaults, semantics, or reusable composition.
2. Register a deterministic story in the matching `story/` file. Set `Definition.Level` and `Definition.Family` to place it on a category and subgroup page. Set its actual building blocks in `Definition.Composes` so the catalog can link them.
3. Add a consumer example in `docs/snippets/<story_id>.go` (replace `/` and `-` with `_`) and register its renderer in `docs/snippets/render.go`. Include sample data, construction and rendering using public APIs. These are ordinary Go functions, compiled and linted with the library. The generator displays the function body without package or outer-function scaffolding and renders that same function in the default **Example** tab. Keep fixture variants in the other tabs; implementation source is linked separately. Missing or empty examples fail the guide build.

Keep services and decisions outside UI models. The registry rejects duplicate IDs, composition cycles, and references to higher atomic levels. See [DESIGN.md](https://github.com/lanej/charmed/blob/main/DESIGN.md) for the design constraints.

## Verify changes

```sh
just check
just docs
```

Install [just](https://github.com/casey/just#installation) and [ShellCheck](https://github.com/koalaman/shellcheck#installing), then run `just` to list targets. The targets select the module's Go version unless explicitly overridden. `just check` runs ShellCheck and checks formatting, vet, module tidiness, strict Go lint, uncached race tests, and reachable vulnerabilities on Linux and macOS. CI pins ShellCheck to 0.11.0. For a quicker iteration, run `just lint`, which includes ShellCheck; use `just shellcheck`, `just test` or `just vuln` for individual gates and `just fmt path/to/file.go` or `just fix ./path/to/package` for scoped fixes. Existing Go lint debt is in a shrinking baseline: new findings must be fixed, and resolved findings removed with `just ratchet`. Never enlarge the baseline to accept new findings.

The generator writes the static guide to `docs/dist/`, validates source and composition, and preserves terminal output colors. Each atomic level has a visual overview; its capability groups have separate pages generated from the registry. Overview cards link directly to those pages and show excerpts of saved native snapshots, selected in `workbench/specimen.go`. Counts, navigation and composition links update with the stories. Existing component, code and variant bookmarks redirect to their new group pages. The responsive browser check covers mobile and desktop layouts, navigation, and fragment links.

## Publish documentation

Edit `docs/*.md`, catalog stories in `story/`, and the shared site template/styles in `workbench/site/`. Generate locally before opening a pull request.

Every same-repository documentation PR gets an isolated preview at `lanej.io/charmed/pr-preview/pr-<number>/`. The PR description is updated with **direct links to rendered pages** after the publisher verifies that the live revision matches the PR head. Previews refresh on push, are excluded from search indexing, and are removed when the PR closes. Forks receive CI checks but no privileged preview publication.

Merging to `main` publishes the tested static bundle through GitHub Pages. No separate documentation server or JavaScript build tool is required; the small client script enhances variant tabs and mobile navigation, and forwards published bookmarks to their current pages.

## Coverage

The inventory below is generated from the registry and distinguishes implemented components from optional integrations. The bundled JuliaMono font preserves braille and box-drawing glyph geometry; see its [license](assets/JuliaMono-LICENSE.txt).

## Git fixtures

Run `go run ./cmd/workbench run git/browser` for the composed repository browser, or use `git/history`, `git/changes` and `git/worktrees` for individual organisms. All data is synthetic and supplied in memory; no Git executable or repository access is needed. Ctrl+Left/Right switches browser sections, Tab switches panes, Enter inspects a selected worktree, and Ctrl+E/Ctrl+Y scrolls history continuation rows independently of commit selection.

Saved native previews: [branching history](assets/git-history.png), [status and diff inspector](assets/git-changes.png), [worktree list](assets/git-worktrees.png), and [repository browser](assets/git-browser.png).
