# Components by capability

Browse related building blocks together. Atomic levels describe composition depth; Go packages follow capabilities. Every component has a direct link in the sidebar. Each entry leads with a native terminal preview, groups alternate snapshots into tabs, and shows highlighted Go usage snippets immediately afterward. Snippets focus on the consumer API rather than a compilable program or the fixture machinery.

Use the sidebar to navigate to a component or your browser's Find command to search the entire page. Saved frames preserve terminal geometry and semantic color; use the native runner to inspect colors, focus and animation.

Time-series traces use native braille dots for finer terminal resolution. Shared axes and series colors retain their meaning; ASCII mode keeps its numeric fallback.

Molecules keep related information together within the child's cell allocation. A labeled field wraps help or validation feedback to its native input width; resize through `field.Input.Resize(columns)`. Metric trends keep the current value and unit visible in short allocations. An empty value is unavailable, while empty history does not discard a known current value. Distribution labels and metadata wrap at the supplied width; allow room for their resulting rows.

Legends are opt-in. Set `DistributionComparison.ShowLegend`, compose `BandOptions.Legend` or `chart.SeriesLegend`, or use `chart.PanelOptions{ShowLegend: true}` with `chart.Panel`. Allocate its child with `options.PlotSize(size, series, palette)` so wrapped keys get their actual number of rows. `chart.PanelPlotSize` remains available for single-row legacy allocations. Legend symbols reuse their plotted styles; label custom boundary markers with `BandMarker.Label`. Native fixtures accept `--legend`, and relevant saved variants include **With legend**. Rendering a panel preserves the native child's state and custom styles; callers own its resize lifecycle.
