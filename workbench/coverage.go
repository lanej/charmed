package workbench

type Coverage struct {
	Upstream string   `json:"upstream"`
	Packages []string `json:"packages"`
	Status   string   `json:"status"`
	Notes    string   `json:"notes"`
	Stories  []string `json:"stories"`
}

func CoverageInventory() []Coverage {
	return []Coverage{
		{Upstream: "Bubble Tea", Packages: []string{"runtime"}, Status: "implemented", Notes: "One root runtime and isolated native PTY previews.", Stories: []string{"screen/services"}},
		{Upstream: "Lip Gloss", Packages: []string{"theme", "layout", "text"}, Status: "implemented", Notes: "Semantic styles, text, cell sizing and responsive panels.", Stories: []string{"foundations/theme", "foundations/layout", "text/primitives"}},
		{Upstream: "Bubbles inputs", Packages: []string{"input"}, Status: "implemented", Notes: "Text input, textarea and native cursor behavior; concrete models remain accessible.", Stories: []string{"input/text", "input/textarea", "input/field"}},
		{Upstream: "Bubbles collections", Packages: []string{"collection"}, Status: "implemented", Notes: "Table, list, tree, file picker and viewport.", Stories: []string{"collection/table", "collection/list", "collection/tree", "collection/files", "collection/viewport"}},
		{Upstream: "Bubbles feedback", Packages: []string{"feedback"}, Status: "implemented", Notes: "Spinner, progress, paginator, timer, stopwatch, help and native key bindings.", Stories: []string{"feedback/spinner", "feedback/progress", "feedback/paginator", "feedback/time", "feedback/help"}},
		{Upstream: "Huh", Packages: []string{"form"}, Status: "implemented", Notes: "Themed forms, input, select, multiselect, confirm, text and note fixtures; Huh owns nested focus.", Stories: []string{"form/settings"}},
		{Upstream: "Glamour", Packages: []string{"text"}, Status: "implemented", Notes: "Markdown rendered into a bounded viewport.", Stories: []string{"text/markdown"}},
		{Upstream: "Harmonica", Packages: []string{"feedback"}, Status: "implemented", Notes: "Explicit spring state with reduced-motion behavior.", Stories: []string{"feedback/motion"}},
		{Upstream: "ntcharts core", Packages: []string{"chart"}, Status: "implemented", Notes: "Canvas, bars, heatmap, line, scatter, streamline, waveline, time series, OHLC and sparklines; native models remain accessible.", Stories: []string{"chart/canvas", "chart/bars", "chart/heatmap", "chart/line", "chart/scatter", "chart/stream", "chart/wave", "chart/timeseries", "chart/ohlc", "chart/sparkline"}},
		{Upstream: "Fang / Cobra", Packages: []string{"cmd/workbench"}, Status: "implemented", Notes: "Command help, completions and error presentation.", Stories: []string{}},
		{Upstream: "Ultraviolet / ANSI / colorprofile", Packages: []string{"layout", "runtime", "story"}, Status: "implemented", Notes: "Reuse Charm's renderer; cell-aware clipping and explicit truecolor/256/16/no-color output.", Stories: []string{"foundations/theme"}},
		{Upstream: "ntcharts pictures", Packages: []string{"chart"}, Status: "partial", Notes: "Portable glyph picture rendering is implemented. URL fetching, heat pictures and high-resolution graphics integrations remain optional.", Stories: []string{"chart/picture"}},
		{Upstream: "ntcharts companion viewers", Packages: []string{}, Status: "planned", Notes: "PDF, SVG, maps and 3D are separate integrations, outside the core dependency graph.", Stories: []string{}},
		{Upstream: "Wish", Packages: []string{}, Status: "planned", Notes: "Optional SSH host adapter.", Stories: []string{}},
		{Upstream: "Log", Packages: []string{"collection"}, Status: "implemented", Notes: "Structured log entries rendered with Charm Log into a bounded viewport; producers remain outside the UI thread.", Stories: []string{"collection/logs"}},
		{Upstream: "Gum / VHS / Freeze / CharmTone", Packages: []string{}, Status: "reference", Notes: "Shell workflows, recording/export and palette integrations; not dependencies of the Go components.", Stories: []string{}},
		{Upstream: "Booba / WASM", Packages: []string{}, Status: "planned", Notes: "Native live previews and static snapshots currently avoid requiring a Bubble Tea fork.", Stories: []string{}},
	}
}
