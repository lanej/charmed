# Build operational screens

Start from the user's task and the closest runnable composition. For operational monitoring, use `pattern.NewOperationalDashboard`: summary first, primary evidence next, secondary diagnostics behind inspection. The [complete example](https://github.com/lanej/charmed/tree/main/examples/operations) owns the Bubble Tea loop; the [consumer snippet](https://github.com/lanej/charmed/blob/main/docs/snippets/template_operations.go) shows the public API. Use another pattern when the task needs selection, editing or independent panes; do not force every TUI into a dashboard.

Read Charmed's [design philosophy](https://github.com/lanej/charmed/blob/main/DESIGN.md) and the chosen component's lifecycle guidance before modifying composition. Use the project's pinned dependency API; links to main describe current source and may be newer. When offline, read those files from an available Charmed checkout or module cache rather than guessing APIs.

## Compose before decorating

Decide what the reader needs to understand or do first. Give that evidence the primary region. Reuse theme, layout, state and navigation components; keep ordinary content neutral, reserve accent for interaction, and avoid repeating a title, summary or help area inside child panels. Borders should establish independent regions, not surround every heading. Keep the title and subtitle together in the panel header, with panel-wide observation status at the right when space permits. Use `DashboardSnapshot.Meters` for exact count/capacity summaries; label what the fill measures, such as remaining requests. Put scheduler configuration, exact accounting and long explanations in accessible inspection when they are secondary to the task.

Choose the visualization from the observation's meaning. Use `chart.BucketBars` for counts in equal-duration intervals; supply complete, partial, missing or unknown coverage explicitly. Keep partial counts distinct from complete intervals. Use time-series lines for appropriate timestamped measurements, and preserve gaps instead of inventing zeroes. Charmed presents supplied facts; the application owns fetching, freshness policy, thresholds and actions.

## Verify the consumer

Build one runnable path with representative observations before expanding the screen. Use the real application's rendering and key routing, with deterministic inputs where practical. Cover populated, empty, waiting, failed-refresh and incomplete observations when they apply. A failed refresh should retain and label any last observation rather than imply current success.

Run the consumer's checks and capture its real output at its normal allocation, a narrow allocation and a short allocation. For an operational dashboard, start with 132x40, 80x24 and 40x18; adapt these to the supported terminals. Inspect Unicode first, then explicit ASCII and monochrome fallbacks. Confirm that primary evidence remains legible, omissions are explicit, and hidden content is reachable by keyboard. Exercise scrolling, inspection/back, resize and data replacement while retaining the user's place within the documented lifecycle limits.

When working in a Charmed checkout, the workbench provides reference captures:

```sh
go run ./cmd/workbench capture template/operations --columns 80 --rows 24 --plain
go run ./cmd/workbench run template/operations
```

These commands validate the reference, not a consumer that merely imports Charmed. Add regression protection around the consumer's consequential geometry, state distinctions and keyboard paths. Preserve rendered captures for visual review: passing bounds checks cannot establish hierarchy, readability or aesthetic quality. Report what was exercised and any remaining gaps; do not claim user approval or publish changes merely because this skill was invoked.
