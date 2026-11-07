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

Publish `docs/dist/` to GitHub Pages or any static host. The Pages workflow builds with Go and uploads that directory. No Node, npm, browser JavaScript or running preview service is required. All navigation and assets use relative URLs for repository subpaths. To preview, serve `docs/dist/` with any static HTTP server.

Edit the Markdown guides in `docs/` and the shared HTML/CSS in `workbench/site/`. The generator reads the compiled story registry, validates composition and captures deterministic frames. Foundations share one page, component capabilities share another, and full screens sit together under examples. Stable story IDs are section anchors.

## Add a fixture

Add a wrapper where shared defaults or composition help. Put its deterministic fixture and registration in the matching `story/foundations.go`, `atoms.go`, `atoms_charts.go`, `molecules.go`, `organisms.go` or `screens.go` file. Go packages remain organized by capability.

Declare actual building blocks in `Definition.Composes`. The generator rejects missing references, duplicates, cycles and references to higher atomic levels. It renders links to the corresponding sections automatically. Empty levels need no placeholders.

Keep application services outside the model. Use typed events, controlled steps and bounded sample data. Browser WASM remains an optional integration because the current Booba route requires a patched Bubble Tea replacement.

## Verify

```sh
go vet ./...
go test -race ./...
go run ./cmd/workbench generate --out docs
```

The generator tests check static output, composition links, source escaping, native fixture coverage and repeatable builds. The native suite covers terminal bounds, focus, selection, resize and process cleanup.

## Coverage

Implemented building blocks and optional integrations are listed below.

Terminal snapshots use a bundled [JuliaMono webfont](https://github.com/cormullion/juliamono/blob/0dfc749503778730d6af79183a17fcde668ae626/webfonts/JuliaMono-Regular.woff2) so braille dots and box-drawing glyphs retain a consistent cell width. It is distributed under its [included license](assets/JuliaMono-LICENSE.txt).
