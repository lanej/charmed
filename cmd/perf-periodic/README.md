# Periodic observability CPU probe

This developer fixture measures four time-series charts, an HTTP/log inspector and a health grid receiving synthetic snapshots every five seconds. It is an application-owned producer and measurement tool, not a polling service or new production organism. Run it from the repository root. Tab and the inspector's native keys remain available; H hides or restores the inspector. `--history` inspects the oldest request instead of following the latest one.

```sh
go run ./cmd/perf-periodic --scenario mixed --report /tmp/periodic.json
uv run cmd/perf-periodic/measure.py --out .cache/perf/periodic-run
```

The collector builds with the repository's pinned Go toolchain, runs replacement and application modes serially in real 120×40 PTYs, and retains source identity, binary hash, patches, new source files, raw terminal bytes, output timestamps and JSON reports. The output directory must be new or empty and ignored by Git, or outside the repository. Default measurements last two minutes after six seconds of warmup, avoiding a warmup boundary exactly on a five-second source tick. Builds occur before measurement; run without competing builds, tests or workloads.

| Control | Meaning |
| --- | --- |
| `--scenario unchanged` | Resubmit identical observations |
| `--scenario metric` | Change one point in one of four charts |
| `--scenario request` | Change one request and one health cell |
| `--scenario append` | Append one bounded chart point and one log |
| `--scenario mixed` | Combine metric, request, health and append changes |
| `--scenario hidden` | Update hidden inspector observations without preparing them |
| `--records 100\|1000` | Bound retained request/log snapshots; ten initial logs correlate with the latest request |
| `--points 60\|300` | Bound each chart's history |
| `--mode replace\|apply` | Force existing setters, or use opt-in data application |
| `--headless` | Disable terminal rendering; lower-bound event-loop measurements only |

## Apply snapshots without forcing refresh

`TimeSeries.ApplyData` and `Sparkline.ApplyData` report whether data preparation ran. Equal observations preserve native customization; changed data follows the existing `SetData` lifecycle. `HTTPLog.ApplyEntries`, `LogTree.ApplyEntries` and `LogInspector.ApplyEntries` return `(applied, error)`. A true result includes initial empty projection. Complete snapshots matching previously validated data need no repeated validation; changed input and evicted prefixes are validated before mutation. Errors leave the inspector's request/log observations unchanged.

Use the result to advance the owning scene's presentation revision only when needed. The fixture's compiled example is `internal/periodic/dashboard.go`: it forwards native command results, retains unrelated region output, and supplies `PresentationRevision` to the shared runtime scheduler. Hidden observations are retained and applied when reopened. Existing `SetData`, `SetEntries` and refresh methods remain explicit full-refresh operations after direct native presentation edits; the new methods do not infer those edits. Log nested fields remain caller-owned immutable observations. HTTP captured bytes remain copied.

## Interpret results

Reports separate unchanged decisions, completed frames and cancelled pending updates. `processing_ns` is update acceptance to decision/frame completion, not terminal painting. The external collector adds first-output latency only when bytes were observed before the next update or measurement end. A completed frame may produce no bytes when its cells are unchanged; missing output has a null latency, not zero and not an assumed dropped observation. Percentiles include sample counts and use nearest rank; approximately 24 observations cannot establish universal tail-latency bounds.

Process CPU includes the native runtime, garbage collection and probe overhead, expressed as percent of one core over the measured window. Per-update counters include source snapshot generation and preparation through frame completion. Warmup, forced boundary GC and terminal shutdown are excluded from summary CPU/bytes. Start/end heaps are sampled after GC and include retained fixture data, current buffers and report samples. Compare equivalent geometry, history mode, retained sizes, toolchain and host. Interactive controls or resize alter work; actual initial/final and per-update sizes are recorded.

No idle frame timer is added. Bubble Tea still owns its native terminal flush timer; zero terminal bytes do not establish zero CPU. A headless report is explicitly marked and must not be presented as native terminal performance. Source observations are synthetic and do not include network polling or terminal-emulator painting. The collector always terminates only its own process on failure and closes its PTYs.

## Repeatable checks and profiles

```sh
gotestsum -- -race ./chart ./collection ./pattern ./internal/periodic ./cmd/perf-periodic
uv run python -m unittest discover -s cmd/perf-periodic -p 'test_*.py'
go test -c -o /tmp/periodic.test ./internal/periodic
/tmp/periodic.test -test.run '^$' -test.bench BenchmarkPeriodicSnapshot -test.benchmem -test.count=5
/tmp/periodic.test -test.run '^$' -test.bench 'BenchmarkPeriodicSnapshot/replace/unchanged' -test.benchtime=2s -test.cpuprofile=/tmp/periodic.cpu -test.memprofile=/tmp/periodic.mem
```

The microbenchmarks alternate two prepared observations so changed workloads remain changed on every iteration. They include synchronous frame generation without waiting for an artificial FPS interval; native probes exercise real scheduling and idle time. Output-equivalence tests compare application with full replacement, including hidden-pane reopening. Regression checks cover native customization, copied bytes, empty initialization, every HTTP observation field, invalid input, no-op frame suppression, missing health states, error recovery, command forwarding and completed measurement windows.
