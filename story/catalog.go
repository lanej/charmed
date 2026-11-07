package story

import (
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/runtime"
	"strings"
)

func definitions() []Definition {
	widget := func(id, title, level, family, description, symbol string, controls []string, factory func(Environment) (runtime.Scene, error)) Definition {
		return Definition{ID: id, Title: title, Level: level, Family: family, Description: description, Source: "story/widgets.go", Symbol: symbol,
			Upstream: "https://github.com/charmbracelet/bubbles", Controls: controls, Defaults: DefaultConfig(), Factory: factory}
	}
	d := []Definition{
		widget("foundations/theme", "Foundations/Theme", "Foundations", "theme", "Semantic roles, terminal palettes, contrast and no-color output.", "themeStory", nil, themeStory),
		widget("foundations/layout", "Foundations/Layout", "Foundations", "layout", "Cell-based panels and responsive allocation.", "layoutStory", nil, layoutStory),
		widget("text/primitives", "Atoms/Text/Primitives", "Atoms", "text", "Labels, badges and keyboard hints share semantic roles.", "textStory", []string{"value"}, textStory),
		widget("input/text", "Atoms/Input/Text", "Atoms", "input", "A focused input with a native virtual cursor and shared theme.", "inputStory", []string{"value", "state"}, inputStory),
		widget("input/textarea", "Atoms/Input/Textarea", "Atoms", "input", "Multiline input with native editing and bounded dimensions.", "areaStory", []string{"value"}, areaStory),
		widget("input/field", "Molecules/Input/Labeled field", "Molecules", "input", "Label, help and validation composed around the same input atom.", "fieldStory", []string{"value", "state"}, fieldStory),
		widget("feedback/spinner", "Atoms/Feedback/Spinner", "Atoms", "feedback", "Native spinner; motion can run live or stay paused for evidence.", "spinnerStory", nil, spinnerStory),
		widget("feedback/progress", "Atoms/Feedback/Progress", "Atoms", "feedback", "A themed progress meter with a textual percentage.", "progressStory", []string{"amount"}, progressStory),
		widget("feedback/paginator", "Molecules/Feedback/Pagination", "Molecules", "feedback", "Native pagination logic with shared active and inactive styles.", "paginatorStory", nil, paginatorStory),
		widget("feedback/time", "Atoms/Feedback/Timers", "Atoms", "feedback", "Timer and stopwatch messages continue independently of keyboard focus.", "timeStory", nil, timeStory),
		widget("feedback/help", "Molecules/Feedback/Help", "Molecules", "feedback", "Help is generated from the same native key bindings that handle input.", "helpStory", nil, helpStory),
		widget("feedback/motion", "Foundations/Motion", "Foundations", "motion", "A Harmonica spring stepped explicitly by the owner; reduced motion resolves immediately.", "motionStory", []string{"amount"}, motionStory),
		widget("collection/table", "Organisms/Collections/Table", "Organisms", "collection", "Native row navigation with themed headers, cells and selection.", "tableStory", []string{"state"}, tableStory),
		widget("collection/list", "Organisms/Collections/List", "Organisms", "collection", "A native list with filtering, pagination and selection.", "listStory", nil, listStory),
		widget("collection/tree", "Organisms/Collections/Tree", "Organisms", "collection", "Expand and collapse nested data without losing native tree capabilities.", "treeStory", nil, treeStory),
		widget("collection/files", "Organisms/Collections/File picker", "Organisms", "collection", "A file picker confined to a disposable fixture directory.", "fileStory", nil, fileStory),
		widget("collection/viewport", "Atoms/Collections/Viewport", "Atoms", "collection", "A scrollable content region with its own dimensions.", "viewportStory", nil, viewportStory),
		widget("text/markdown", "Organisms/Text/Markdown", "Organisms", "text", "Glamour renders Markdown within a native scrollable viewport.", "markdownStory", nil, markdownStory),
		widget("form/settings", "Organisms/Forms/Settings", "Organisms", "form", "A Huh form owns its nested focus and validation.", "formStory", nil, formStory),
	}
	for i := range d {
		switch d[i].ID {
		case "text/markdown":
			d[i].Upstream = "https://github.com/charmbracelet/glamour"
		case "form/settings":
			d[i].Upstream = "https://github.com/charmbracelet/huh"
		case "feedback/motion":
			d[i].Upstream = "https://github.com/charmbracelet/harmonica"
		}
	}
	for _, item := range []struct {
		id, title, level, description, symbol, upstream string
		factory                                         func(Environment) (runtime.Scene, error)
	}{
		{"chart/metric", "Molecules/Charts/Metric trend", "Molecules", "A metric value, unit and sparkline share one bounded allocation.", "metricStory", "https://github.com/NimbleMarkets/ntcharts", metricStory},
		{"chart/panel", "Molecules/Charts/Chart panel", "Molecules", "A chart, title and series legend compose into a reusable region.", "panelStory", "https://github.com/NimbleMarkets/ntcharts", panelStory},
		{"chart/picture", "Atoms/Graphics/Picture", "Atoms", "Portable glyph rendering of an embedded image; no external image service.", "pictureStory", "https://github.com/NimbleMarkets/ntcharts", pictureStory},
		{"collection/logs", "Organisms/Collections/Log viewer", "Organisms", "Structured Charm Log entries in a bounded, scrollable viewport.", "logStory", "https://github.com/charmbracelet/log", logStory},
	} {
		d = append(d, Definition{ID: item.id, Title: item.title, Level: item.level, Family: "composition", Description: item.description, Source: "story/compositions.go", Symbol: item.symbol, Upstream: item.upstream, Defaults: DefaultConfig(), Factory: item.factory})
	}
	for _, kind := range []struct{ id, name string }{
		{"sparkline", "Sparkline"}, {"timeseries", "Time series"}, {"line", "Line"}, {"scatter", "Scatter"}, {"bars", "Bars"}, {"heatmap", "Heatmap"}, {"stream", "Streamline"}, {"wave", "Waveline"}, {"ohlc", "OHLC candles"}, {"canvas", "Canvas"},
	} {
		kind := kind
		controls := []string{"state"}
		switch kind.id {
		case "timeseries", "line", "scatter":
			controls = append(controls, "series", "amount")
		case "sparkline", "stream", "wave", "ohlc":
			controls = append(controls, "amount")
		case "bars":
			controls = append(controls, "series", "horizontal")
		}
		d = append(d, Definition{ID: "chart/" + kind.id, Title: "Atoms/Charts/" + kind.name, Level: "Atoms", Family: "chart", Description: "Real ntcharts " + kind.name + " rendering with semantic styling and deterministic fixture data.",
			Source: "story/charts.go", Symbol: "chartStory", Upstream: "https://github.com/NimbleMarkets/ntcharts", Controls: controls, Defaults: DefaultConfig(),
			Factory: func(e Environment) (runtime.Scene, error) { return chartStory(kind.id, e) }})
	}
	d = append(d, Definition{ID: "screen/services", Title: "Screens/Service workbench", Level: "Screens", Family: "pattern", Description: "Filter a service list and inspect linked throughput and details. Resize preserves the selected service.", Source: "pattern/dashboard.go", Symbol: "NewDashboard", Upstream: "https://github.com/charmbracelet/bubbletea", Controls: []string{"state"}, Defaults: func() Config { c := DefaultConfig(); c.Columns = 120; c.Rows = 34; return c }(), Factory: func(e Environment) (runtime.Scene, error) {
		services := pattern.SampleServices()
		if e.Config.State == "empty" {
			services = nil
		}
		return pattern.NewDashboard(pattern.DashboardOptions{Theme: e.Config.Palette(), Services: services, Live: e.Config.Live, Emit: e.Emit}), nil
	}})
	d = append(d,
		Definition{ID: "chart/band-matrix", Title: "Molecules/Charts/Band matrix", Level: "Molecules", Family: "chart", Description: "Disconnected version bands in one rectangular frame. Every metric has its own labeled percentage axis and directly labeled version means.", Source: "story/cohort.go", Symbol: "bandMatrixStory", Upstream: "https://github.com/NimbleMarkets/ntcharts", Controls: []string{"state", "sigma"}, Defaults: DefaultConfig(), Factory: bandMatrixStory},
		Definition{ID: "screen/deployment-bands", Title: "Screens/Deployment bands", Level: "Screens", Family: "pattern", Description: "Four aligned HTTP status rows compose version distributions, optional attention boundaries and a linked class history. No curves or series legends.", Source: "pattern/cohort_bands.go", Controls: []string{"state", "scenario", "sigma", "referenceID", "candidateID"}, Scenarios: []string{"isolated-hosts", "uniform-regression", "unequal-traffic", "insufficient-samples"}, Defaults: func() Config { c := DefaultConfig(); c.Columns = 132; c.Rows = 39; return c }(), Factory: deploymentBandsStory},
		Definition{ID: "chart/interval-band", Title: "Atoms/Charts/Interval band", Level: "Atoms", Family: "chart", Description: "Nested caller-supplied intervals, a mean and selected observation rendered on an ntcharts canvas.", Source: "chart/band.go", Symbol: "IntervalBand", Upstream: "https://github.com/NimbleMarkets/ntcharts", Controls: []string{"state", "sigma"}, Defaults: DefaultConfig(), Factory: bandStory},
		Definition{ID: "screen/deployment", Title: "Screens/Deployment comparison", Level: "Screens", Family: "pattern", Description: "Compare 2xx, 3xx, 4xx and 5xx by version identifier during a blue/green deployment. Reference/candidate roles, request-weighted histories, 1/2/3s host-spread bands and linked inspection use synthetic interval counts.", Source: "examples/deployment/main.go", Upstream: "https://github.com/NimbleMarkets/ntcharts", Controls: []string{"state", "scenario", "sigma", "referenceID", "candidateID"}, Scenarios: []string{"isolated-hosts", "uniform-regression", "unequal-traffic", "insufficient-samples"}, Defaults: func() Config { c := DefaultConfig(); c.Columns = 132; c.Rows = 40; return c }(), Factory: deploymentStory},
	)
	d = append(d,
		Definition{ID: "feedback/attention", Title: "Molecules/Feedback/Attention", Level: "Molecules", Family: "feedback", Description: "Optional warn/error boundaries compose a labeled attention indicator from the shared semantic pill atom.", Source: "feedback/attention.go", Symbol: "Attention", Upstream: "https://github.com/charmbracelet/lipgloss", Controls: []string{"state", "amount", "boundaries"}, Defaults: DefaultConfig(), Factory: attentionStory},
		Definition{ID: "chart/distributions", Title: "Molecules/Charts/Cohort distributions", Level: "Molecules", Family: "chart", Description: "Mean/spread labels, shared-axis interval bands and optional boundary markers compose into a reusable distribution comparison.", Source: "chart/comparison.go", Upstream: "https://github.com/NimbleMarkets/ntcharts", Controls: []string{"sigma", "boundaries"}, Defaults: DefaultConfig(), Factory: distributionStory},
	)
	for i := range d {
		id := d[i].ID
		for _, control := range d[i].Controls {
			if control == "state" {
				d[i].States = []string{"ready", "empty", "loading", "error"}
			}
		}
		switch id {
		case "input/text":
			d[i].States = []string{"ready", "disabled"}
		case "input/field":
			d[i].States = []string{"ready", "error"}
		case "screen/services":
			d[i].States = []string{"ready", "empty"}
		case "screen/deployment", "screen/deployment-bands":
			d[i].States = []string{"ready", "empty"}
			d[i].Controls = append(d[i].Controls, "boundaries")
			d[i].Defaults.Boundaries = true
		case "feedback/attention":
			d[i].States = []string{"ready", "empty"}
			d[i].Defaults.Boundaries = true
		case "chart/distributions":
			d[i].Defaults.Boundaries = true
		case "chart/interval-band", "chart/band-matrix":
			d[i].States = []string{"ready", "empty"}
		case "chart/metric":
			d[i].Controls = []string{"amount"}
			d[i].Family = "chart"
		case "chart/panel":
			d[i].Controls = []string{"series", "amount"}
			d[i].Family = "chart"
		case "chart/picture":
			d[i].Family = "graphics"
		case "collection/logs":
			d[i].Family = "collection"
		}
		coreChart := strings.HasPrefix(id, "chart/") && id != "chart/panel" && id != "chart/metric" && id != "chart/picture" && id != "chart/interval-band" && id != "chart/distributions" && id != "chart/band-matrix"
		d[i].Clock = coreChart || id == "feedback/spinner" || id == "feedback/time" || id == "screen/services"
		d[i].Step = coreChart || id == "feedback/spinner" || id == "feedback/progress" || id == "feedback/motion" || id == "collection/logs" || id == "screen/services" || id == "screen/deployment" || id == "screen/deployment-bands"
		if id == "chart/timeseries" {
			d[i].Source = "examples/timeseries/main.go"
			d[i].Symbol = ""
		}
	}
	return d
}
