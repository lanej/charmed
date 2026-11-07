# Develop & publish

## Run native fixtures

Requires Go 1.26.8 or later.

```sh
go run ./cmd/workbench list
go run ./cmd/workbench run screen/services
go run ./cmd/workbench run chart/timeseries --theme light --series 3
go run ./cmd/workbench run chart/distributions --legend
go run ./cmd/workbench capture screen/services --plain
```

Use Ctrl+C to exit a native story. The library and native runner retain their terminal interactions independently of this static guide.

## Build the guide

From the repository root:

```sh
go run ./cmd/workbench generate --out docs
```

Publish `docs/dist/` to GitHub Pages or any static host. The Pages workflow builds with Go and uploads that directory. No Node, npm or running preview service is required. A small, local JavaScript file enhances accessible variant tabs; without JavaScript, all snapshots remain visible as ordinary static sections. All navigation and assets use relative URLs for repository subpaths. To preview, serve `docs/dist/` with any static HTTP server.

Edit the Markdown guides in `docs/` and the shared HTML/CSS in `workbench/site/`. The generator reads the compiled story registry, validates composition and captures deterministic frames. Foundations share one page, component capabilities share another, and full screens sit together under examples. The sidebar links directly to every fixture, with visual previews preceding syntax-highlighted source. Stable story IDs are section anchors.

## Add a fixture

Add a wrapper where shared defaults or composition help. Put its deterministic fixture and registration in the matching `story/foundations.go`, `atoms.go`, `atoms_charts.go`, `molecules.go`, `organisms.go` or `screens.go` file. Go packages remain organized by capability.

Declare actual building blocks in `Definition.Composes`. The generator rejects missing references, duplicates, cycles and references to higher atomic levels. It renders links to the corresponding sections automatically. Empty levels need no placeholders.

Keep application services outside the model. Use typed events, controlled steps and bounded sample data. Browser WASM remains an optional integration because the current Booba route requires a patched Bubble Tea replacement.

## Pull request previews

Each same-repository PR receives a public, revision-verified snapshot of the
complete static guide at:

- Components: `https://lanej.io/charmed/pr-preview/pr-<number>/components.html`
- Brand sheet: `https://lanej.io/charmed/pr-preview/pr-<number>/brand.html`
- Catalog: `https://lanej.io/charmed/pr-preview/pr-<number>/`

The unprivileged `pr-preview.yml` workflow generates and stamps the PR's
actual head SHA. A separate trusted publishing job stores the generated static
files in `pr-preview-cache`, then dispatches the existing main-branch
Pages workflow. That workflow overlays the snapshots onto the production catalog
without replacing its root pages. The preview is updated on each push and
removed when the PR closes.

The PR description receives direct browser links **only after** the published
`revision.txt` matches the PR's current head. The preview section is
maintained between `<!-- charmed-preview:start -->` and
`<!-- charmed-preview:end -->` comments; edits elsewhere are retained.
Do not add a link to the PR's own GitHub URL inside its description.
Previews are marked `noindex,nofollow` and use relative static assets.

Fork PRs intentionally do not receive privileged preview publishing. They
still receive regular CI checks. If the publisher workflow has just been
installed, run `pr-preview.yml` manually with the PR number once it is on
`main` to seed previews for already-open PRs.

`go run ./cmd/workbench generate --out docs` remains the same local
build command.

## Verify

```sh
go vet ./...
go test -race ./...
go run ./cmd/workbench generate --out docs
```

The generator tests check static output, component sidebar links, variant tab semantics, syntax highlighting, composition links, source escaping, native fixture coverage and repeatable builds. The native suite covers terminal bounds, focus, selection, resize and process cleanup.

See the [brand sheet](brand.html) for the website palette, typography, lockup, and native status color separation.

## Coverage

Implemented building blocks and optional integrations are listed below.

Terminal snapshots use a bundled [JuliaMono webfont](https://github.com/cormullion/juliamono/blob/0dfc749503778730d6af79183a17fcde668ae626/webfonts/JuliaMono-Regular.woff2) so braille dots and box-drawing glyphs retain a consistent cell width. It is distributed under its [included license](assets/JuliaMono-LICENSE.txt).
