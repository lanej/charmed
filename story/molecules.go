package story

import (
	"errors"
	"fmt"
	"slices"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/attention"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/input"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
)

func moleculesDefinitions() []Definition {
	d := []Definition{
		{ID: "diagram/state-transitions", Title: "Molecules/Diagrams/State transitions", Level: "Molecules", Family: "diagram", Description: "Labeled arrows connect state boxes directly, with local self-loops and vertical layout for narrow regions.", Source: "story/diagrams.go", Symbol: "stateDiagramStory", Defaults: DefaultConfig(), Factory: stateDiagramStory},
		widget("input/field", "Molecules/Input/Labeled field", "Molecules", "input", "Label, help and validation composed around the same input atom.", "fieldStory", []string{"value", "state"}, fieldStory),
		widget("feedback/paginator", "Molecules/Feedback/Pagination", "Molecules", "feedback", "Native pagination logic with shared active and inactive styles.", "paginatorStory", nil, paginatorStory),
		widget("feedback/help", "Molecules/Feedback/Help", "Molecules", "feedback", "Help is generated from the same native key bindings that handle input.", "helpStory", nil, helpStory),
	}

	d = append(d,
		Definition{ID: "chart/band-matrix", Title: "Molecules/Charts/Band matrix", Level: "Molecules", Family: "chart", Description: "Directly labeled, disconnected distribution rows composed in a rectangular frame.", Source: "chart/comparison.go", Controls: []string{"sigma"}, Composes: []string{"chart/distributions", "foundations/layout"}, Defaults: DefaultConfig(), Factory: bandMatrixStory},
		Definition{ID: "chart/metric", Title: "Molecules/Charts/Metric trend", Level: "Molecules", Family: "chart", Description: "A metric value, unit and sparkline share one bounded allocation; unavailable values and missing history remain distinct.", Source: "story/molecules.go", Symbol: "metricStory", Upstream: "https://github.com/NimbleMarkets/ntcharts", Controls: []string{"amount", "state"}, States: []string{"ready", "empty"}, Composes: []string{"chart/sparkline"}, Defaults: DefaultConfig(), Factory: metricStory},
		Definition{ID: "chart/panel", Title: "Molecules/Charts/Chart panel", Level: "Molecules", Family: "chart", Description: "A chart and title compose into a reusable region with an optional color-matched legend.", Source: "story/molecules.go", Symbol: "panelStory", Upstream: "https://github.com/NimbleMarkets/ntcharts", Controls: []string{"series", "amount", "legend"}, Composes: []string{"chart/timeseries", "foundations/layout"}, Defaults: DefaultConfig(), Factory: panelStory},
		Definition{ID: "feedback/attention", Title: "Molecules/Feedback/Attention", Level: "Molecules", Family: "feedback", Description: "Optional warn/error boundaries compose a labeled attention indicator from the shared semantic pill atom.", Source: "feedback/attention.go", Symbol: "Attention", Upstream: "https://github.com/charmbracelet/lipgloss", Controls: []string{"state", "amount", "boundaries"}, States: []string{"ready", "empty"}, Composes: []string{"feedback/status"}, Defaults: DefaultConfig(), Factory: attentionStory},
		Definition{ID: "chart/distributions", Title: "Molecules/Charts/Cohort distributions", Level: "Molecules", Family: "chart", Description: "Mean/spread labels, shared-axis interval bands and optional boundary markers compose into a reusable distribution comparison.", Source: "chart/comparison.go", Upstream: "https://github.com/NimbleMarkets/ntcharts", Controls: []string{"sigma", "boundaries", "legend", "state"}, States: []string{"ready", "empty"}, Composes: []string{"chart/interval-band"}, Defaults: DefaultConfig(), Factory: distributionStory},
		Definition{ID: "chart/delta", Title: "Molecules/Charts/Merged delta property", Level: "Molecules", Family: "chart", Description: "A single zero-centred bar cancels the shared reference/candidate amount and emphasizes the signed percentage-point difference.", Source: "chart/delta.go", Upstream: "https://github.com/NimbleMarkets/ntcharts", Controls: []string{"amount", "state"}, States: []string{"ready", "empty"}, Composes: []string{"chart/interval-band", "text/signed-number"}, Defaults: DefaultConfig(), Factory: deltaStory},
		Definition{ID: "feedback/percent", Title: "Molecules/Feedback/Percentage bar", Level: "Molecules", Family: "feedback", Description: "A native progress bar and exact percentage compose a fixed-width graphical table value on a 0–100% scale.", Source: "feedback/percent.go", Symbol: "PercentBar", Upstream: "https://github.com/charmbracelet/bubbles", Controls: []string{"amount", "state"}, States: []string{"ready", "empty"}, Composes: []string{"feedback/progress"}, Defaults: DefaultConfig(), Factory: percentStory},
	)

	snippets := map[string]string{
		"diagram/state-transitions": "state-transitions",
		"input/field":               "field", "feedback/paginator": "paginator",
		"feedback/help": "help", "chart/band-matrix": "band-matrix",
		"chart/metric": "metric-trend", "chart/panel": "chart-panel",
		"feedback/attention": "attention", "chart/distributions": "distributions",
		"chart/delta": "delta", "feedback/percent": "percent",
	}

	for i := range d {
		d[i].ExampleSource = "docs/snippets/" + snippets[d[i].ID] + ".go.txt"

		if d[i].ID == "input/field" {
			d[i].States = []string{"ready", "empty", "error"}
			d[i].Composes = []string{"input/text"}
		}

		if slices.Contains(d[i].Controls, "boundaries") {
			d[i].Defaults.Boundaries = true
		}
	}

	return d
}

func fieldStory(e Environment) (runtime.Scene, error) {
	f := input.Field{Label: "Environment name", Help: "Use at least three characters.", Theme: e.inputPalette(), Input: input.NewText(input.TextOptions{Theme: e.inputPalette(), Value: e.Config.Value, Width: e.Config.Columns - 2, Validate: func(s string) error {
		if len(s) < 3 {
			return errors.New("enter at least three characters")
		}

		return nil
	}})}
	if e.Config.State == "error" {
		f.Input.SetValue("x")
		f.Input.Err = errors.New("enter at least three characters")
	}

	if e.Config.State == "empty" {
		f.Input.SetValue("")
		f.Input.Placeholder = "production"
		f.Input.Err = nil
	}

	if e.capture {
		_ = f.Input.Focus()
	}

	return &scene{init: f.Input.Focus, resize: func(s layout.Size) { f.Input.Resize(s.Width) }, update: func(msg tea.Msg) tea.Cmd { var cmd tea.Cmd; f.Input, cmd = f.Input.Update(msg); return cmd }, view: func(layout.Size) string { return f.View() }}, nil
}

func paginatorStory(e Environment) (runtime.Scene, error) {
	m := feedback.NewPaginator(50, 10, e.Config.Palette())

	return &scene{update: func(msg tea.Msg) tea.Cmd { var cmd tea.Cmd; m, cmd = m.Update(msg); return cmd }, view: func(layout.Size) string {
		return title("Pagination", e) + m.View() + fmt.Sprintf("\n\nPage %d of %d", m.Page+1, m.TotalPages) + footer("Left / right changes page", e)
	}}, nil
}

type helpKeys struct{ next, selectKey key.Binding }

func (k helpKeys) ShortHelp() []key.Binding { return []key.Binding{k.next, k.selectKey} }

func (k helpKeys) FullHelp() [][]key.Binding { return [][]key.Binding{k.ShortHelp()} }

func helpStory(e Environment) (runtime.Scene, error) {
	m := feedback.NewHelp(e.Config.Columns, e.Config.Palette())
	keys := helpKeys{key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "next field")), key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "select"))}

	return &scene{resize: func(s layout.Size) { m.SetWidth(s.Width) }, view: func(layout.Size) string { return title("Help from key bindings", e) + m.View(keys) }}, nil
}

func metricStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		value := fmt.Sprintf("%.0f", e.Config.Amount*20)
		values := sampleValues(48, 0, e.Config.Amount)

		if e.Config.State == "empty" {
			value, values = "", nil
		}

		return chart.MetricTrend("Throughput", value, "req/min", values, layout.Size{Width: size.Width, Height: min(size.Height, 8)}, e.Config.Palette())
	}), nil
}

func panelStory(e Environment) (runtime.Scene, error) {
	data := sampleSeries(e.Config.Series, 0, e.Config.Amount)
	plot := chart.NewTimeSeries(data, e.Config.Size(), e.Config.Palette())
	plot.Focus()

	options := chart.PanelOptions{ShowLegend: e.Config.Legend}

	return &scene{resize: func(size layout.Size) { plot.Resize(options.PlotSize(size, data, e.Config.Palette())) }, update: func(msg tea.Msg) tea.Cmd { var cmd tea.Cmd; plot, cmd = plot.Update(msg); return cmd }, view: func(size layout.Size) string {
		return chart.Panel("Request throughput", plot, data, size, e.Config.Palette(), true, options)
	}}, nil
}

func attentionStory(e Environment) (runtime.Scene, error) {
	return static(func(layout.Size) string {
		rule := attention.Rule{}
		if e.Config.Boundaries {
			rule = attention.Rule{Warn: &attention.Bounds{Upper: attention.Limit(60)}, Error: &attention.Bounds{Upper: attention.Limit(80)}}
		}

		result := rule.Evaluate(e.Config.Amount, e.Config.State != "empty")

		indicator := feedback.Attention(result.Level, e.Config.Palette())
		if result.Level == attention.Off {
			indicator = "OFF / no boundaries"
		}

		return title("Composed attention indicator", e) + indicator + fmt.Sprintf("\n\nValue %.0f\n%s\n\n", e.Config.Amount, rule.Summary("units")) + "Rules are optional. Missing measurements are UNKNOWN.\nEquality triggers; ERROR takes precedence over WARN.\nThe indicator composes the shared semantic pill atom."
	}), nil
}

func distributionStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		selected := 40.0
		d := chart.DistributionComparison{Series: []chart.DistributionSeries{{Label: "Reference", Mean: 20, StdDev: 3, Members: 64, HasSpread: true}, {Label: "Candidate", Mean: 30, StdDev: 8, Members: 32, HasSpread: true, Selected: &selected}}, ShowLegend: e.Config.Legend, Sigma: e.Config.Sigma, Min: 0, Max: 60, Unit: "%", SpreadUnit: "pp"}

		if e.Config.State == "empty" {
			d.Series = nil
		}

		if e.Config.Boundaries {
			d.Markers = []chart.BandMarker{{Value: 26, Glyph: 'W', Label: "warning boundary", Style: feedback.AttentionStyle(attention.Warn, e.Config.Palette())}, {Value: 29, Glyph: 'E', Label: "error boundary", Style: feedback.AttentionStyle(attention.Error, e.Config.Palette())}}
		}

		return d.View(size.Width, e.Config.Palette())
	}), nil
}

func deltaStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		property := chart.DeltaProperty{Value: (e.Config.Amount - 50) / 100, Extent: 0.5, Available: e.Config.State != "empty", Style: e.Config.Palette().Text()}
		return title("Merged delta property", e) + property.View(size.Width, e.Config.Palette())
	}), nil
}

func percentStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		return title("Compact percentage bar", e) + feedback.PercentBar(e.Config.Amount, e.Config.State != "empty", min(size.Width, 30), e.Config.Palette())
	}), nil
}

func bandMatrixStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		matrix := chart.BandMatrix{Title: fmt.Sprintf("Host mean +/- %ds", e.Config.Sigma)}

		for _, metric := range []struct {
			label    string
			mean, sd float64
		}{{"Success", 96, .4}, {"Server error", .2, .03}} {
			selected := metric.mean + 5*metric.sd
			d := chart.DistributionComparison{Sigma: e.Config.Sigma, Min: 0, Max: 100, Unit: "%", SpreadUnit: "pp", Series: []chart.DistributionSeries{
				{Label: "a13c9f2", Mean: metric.mean, StdDev: metric.sd, Members: 64, HasSpread: true},
				{Label: "f4d7e81", Mean: metric.mean + metric.sd, StdDev: metric.sd * 1.5, Members: 32, HasSpread: true, Selected: &selected},
			}}
			matrix.Rows = append(matrix.Rows, chart.BandRow{Title: metric.label, Distribution: d})
		}

		return matrix.View(size, e.Config.Palette(), false)
	}), nil
}
