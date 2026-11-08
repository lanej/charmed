# Charmed

A Go design system for Charm and ntcharts, with a static reference guide and native terminal fixtures.

Components share semantic themes, terminal-cell geometry, focus conventions and state handling. The guide groups related components, saved Go output, source and composition examples by capability.

![Charmed static guide in night mode](.github/pr-assets/static-guide.png)

## Run it

Requires Go **1.26.10+**. Native terminal sessions support Linux and macOS.

```sh
go run ./cmd/workbench list
go run ./cmd/workbench run screen/services
go run ./cmd/workbench run screen/deployment --scenario isolated-hosts --sigma 3
go run ./cmd/workbench run screen/deployment --comparison-variant bands
go run ./cmd/workbench run chart/timeseries --theme light --series 3
go run ./cmd/workbench capture screen/services --plain
```

## Components

Go packages follow capabilities. Atomic levels organize the catalog.

| Package | Includes |
| --- | --- |
| `theme` | Dark, light, high-contrast and monochrome themes; semantic color roles; ASCII and motion policies |
| `layout` | ANSI-aware fitting, cell rectangles, responsive splits, stacks and panels |
| `interaction` | Focus scopes and typed selection/submission events |
| `input` | Bubbles text input, textarea, virtual cursor and labeled field composition |
| `collection` | Native tables, lists, trees, file pickers, viewports and bounded structured logs |
| `feedback` | Attention indicators, spinner, progress, pagination, help, timers, stopwatch and Harmonica motion |
| `attention` | Optional one-sided or two-sided warn/error boundaries, explicit unknown/off states and typed transitions |
| `form` | Themed Huh forms with native nested focus and validation |
| `text` | Labels, badges, semantic pills, key hints and Glamour Markdown rendering |
| `chart` | Sparklines, time series, line/scatter, bars, heatmaps, streamline/waveline, OHLC, canvas, interval bands, cohort distributions, comparison panels and portable pictures |
| `cohort` | Validated interval counts, request-weighted rates and unweighted member sample spread |
| `pattern` | Service workbench and cohort comparison with linked plots, member selection and inspection |
| `runtime` | One root Bubble Tea terminal session |
| `story` | Deterministic Go fixtures shared by terminal execution, browser previews and snapshots |

Stateful wrappers expose their concrete upstream models or return them directly. Applications retain access to native APIs. Components return Bubble Tea commands; their parent forwards them. Application services and business decisions remain outside the UI library.

```go
// A field remains an ordinary concrete Go model.
field := input.NewText(input.TextOptions{
    Placeholder: "Environment name",
    Width:       40,
    Theme:       theme.New(theme.Dark),
})

// Forward this command from the owning model.
cmd := field.Focus()
```

The core uses upstream Charm v2 modules. It has no JavaScript runtime requirement and no Bubble Tea fork.

## Static guide

[Browse the guide](https://lanej.github.io/charmed/) for foundations, components grouped by capability, composition examples and the deployment telemetry contract. Components share sections rather than individual pages, but every fixture has a sidebar link. Each fixture leads with a saved native preview and variant tabs, followed by highlighted source and a native command. See the [brand sheet](https://lanej.github.io/charmed/brand.html) for visual tokens and identity.

```sh
go run ./cmd/workbench generate --out docs
```

The generator writes `docs/dist/` using Go alone. GitHub Actions uploads this directory to Pages; the published guide uses ordinary HTML, CSS, images and links. It needs no npm or preview service. A tiny local script enhances keyboard-accessible variant tabs; without JavaScript all snapshots remain visible. Saved terminal frames retain ANSI colors as static HTML; run a native fixture for keyboard interaction and animation.

Edit the guides in `docs/*.md` and layout in `workbench/site/`. The registry supplies snapshots and source and validates declared component composition. See [development](docs/development.md) and the [deployment guide](docs/deployment.md).

## Coverage

The catalog covers the core Charm UI libraries and all ten core ntcharts chart families. Portable picture rendering is included. The generated coverage inventory distinguishes implemented components from optional extensions: SSH hosting, URL/high-resolution image integrations, companion PDF/SVG/map/3D viewers, and browser WASM execution.

WASM remains optional because the current Booba route requires a patched Bubble Tea replacement. The native runner and static export work independently of it.

## Verify

```sh
bash scripts/check-go.sh
go run ./cmd/workbench generate --out docs
```

The Go quality command uses the module's Go version by default and runs module tidiness, strict lint and formatting, uncached race tests, and reachable-vulnerability analysis. The same command runs on Linux and macOS in CI. Existing lint and style findings are recorded in a shrinking baseline; new findings fail, and resolved allowances must be removed with `bash scripts/check-go.sh ratchet`. Run `bash scripts/check-go.sh lint` for the lint gate alone.

Checks also exercise cell bounds, deterministic frames, focus and selection, resize, background updates, native process cleanup, static guide generation and component links.

MIT licensed. Upstream components retain their respective licenses.
