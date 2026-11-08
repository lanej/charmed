# Components by capability

Browse related building blocks together. Atomic levels describe composition depth; Go packages follow capabilities. Every component has a direct link in the sidebar. Each entry leads with a native terminal preview, groups alternate snapshots into tabs, and shows highlighted Go usage snippets immediately afterward. Snippets focus on the consumer API rather than a compilable program or the fixture machinery.

Use the sidebar to navigate to a component or your browser's Find command to search the entire page. Saved frames preserve terminal geometry and semantic color; use the native runner to inspect colors, focus and animation.

Time-series traces use native braille dots for finer terminal resolution. Shared axes and series colors retain their meaning; ASCII mode keeps its numeric fallback.

Legends are opt-in. Set `DistributionComparison.ShowLegend`, compose `BandOptions.Legend`, or pass `chart.PanelOptions{ShowLegend: true}` to both `chart.PanelPlotSize` and `chart.Panel`. Legend symbols reuse their plotted styles; label custom boundary markers with `BandMarker.Label`. Native fixtures accept `--legend`, and relevant saved variants include **With legend**.
