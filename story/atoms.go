package story

import (
	"fmt"
	"image"
	"image/color"
	"strings"
	"time"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/input"
	"github.com/lanej/charmed/interaction"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/text"
)

func atomsDefinitions() []Definition {
	d := []Definition{
		widget("text/code", "Atoms/Text/Code", "Atoms", "text", "Literal inline code and syntax-highlighted source blocks share Markdown rendering defaults.", "codeStory", nil, codeStory),
		widget("text/primitives", "Atoms/Text/Primitives", "Atoms", "text", "Labels, badges and keyboard hints share semantic roles.", "textStory", []string{"value"}, textStory),
		widget("input/text", "Atoms/Input/Text", "Atoms", "input", "A focused input with a native virtual cursor and shared theme.", "inputStory", []string{"value", "state"}, inputStory),
		widget("input/textarea", "Atoms/Input/Textarea", "Atoms", "input", "Multiline input with native editing and bounded dimensions.", "areaStory", []string{"value"}, areaStory),
		widget("feedback/spinner", "Atoms/Feedback/Spinner", "Atoms", "feedback", "Native spinner; motion can run live or stay paused for evidence.", "spinnerStory", nil, spinnerStory),
		widget("feedback/progress", "Atoms/Feedback/Progress", "Atoms", "feedback", "A themed progress meter with a textual percentage.", "progressStory", []string{"amount"}, progressStory),
		widget("feedback/time", "Atoms/Feedback/Timers", "Atoms", "feedback", "Timer and stopwatch messages continue independently of keyboard focus.", "timeStory", nil, timeStory),
		widget("collection/viewport", "Atoms/Collections/Viewport", "Atoms", "collection", "A scrollable content region with its own dimensions.", "viewportStory", nil, viewportStory),
	}

	for _, kind := range []struct{ id, name string }{
		{"sparkline", "Sparkline"}, {"timeseries", "Time series"}, {"line", "Line"}, {"scatter", "Scatter"}, {"bars", "Bars"}, {"heatmap", "Heatmap"}, {"stream", "Streamline"}, {"wave", "Waveline"}, {"ohlc", "OHLC candles"}, {"canvas", "Canvas"},
	} {
		kind := kind
		controls := []string{"state"}

		switch kind.id {
		case "timeseries", "line", "scatter":
			controls = append(controls, "series", "amount")
			if kind.id == "timeseries" {
				controls = append(controls, "legend")
			}
		case "sparkline", "stream", "wave", "ohlc":
			controls = append(controls, "amount")
		case "bars":
			controls = append(controls, "series", "horizontal")
		}

		d = append(d, Definition{ID: "chart/" + kind.id, Title: "Atoms/Charts/" + kind.name, Level: "Atoms", Family: "chart", Description: "Real ntcharts " + kind.name + " rendering with semantic styling and deterministic fixture data.", Source: "story/atoms_charts.go", Symbol: "chartStory", Upstream: "https://github.com/NimbleMarkets/ntcharts", Controls: controls, Clock: true, Step: true, Defaults: DefaultConfig(), Factory: func(e Environment) (runtime.Scene, error) { return chartStory(kind.id, e) }})
	}

	d = append(d,
		Definition{ID: "text/signed-number", Title: "Atoms/Text/Signed number", Level: "Atoms", Family: "text", Description: "Exact signed values, configurable positive/negative tones, neutral zero, tiny-value preservation and explicit unavailable state.", Source: "text/signed.go", Controls: []string{"amount", "state", "emphasis"}, States: []string{"ready", "empty"}, Defaults: atomConfig(), Factory: signedNumberStory},
		Definition{ID: "text/pill", Title: "Atoms/Text/Color pill", Level: "Atoms", Family: "text", Description: "Semantic tone and bright, normal or dim emphasis are independent. ASCII output retains the label.", Source: "text/text.go", Symbol: "Pill", Controls: []string{"value", "emphasis"}, Defaults: atomConfig(), Factory: pillsStory},
		Definition{ID: "feedback/status", Title: "Atoms/Feedback/Status", Level: "Atoms", Family: "feedback", Description: "Health and lifecycle states share semantic colors, named glyphs and emphasis-independent spacing.", Source: "feedback/status.go", Controls: []string{"emphasis"}, Composes: []string{"text/pill"}, Defaults: atomConfig(), Factory: statusStory},
		Definition{ID: "chart/picture", Title: "Atoms/Graphics/Picture", Level: "Atoms", Family: "graphics", Description: "Portable glyph rendering of an embedded image; no external image service.", Source: "story/atoms.go", Symbol: "pictureStory", Upstream: "https://github.com/NimbleMarkets/ntcharts", Defaults: DefaultConfig(), Factory: pictureStory},
		Definition{ID: "chart/interval-band", Title: "Atoms/Charts/Interval band", Level: "Atoms", Family: "chart", Description: "Nested caller-supplied intervals, a mean and selected observation rendered on an ntcharts canvas.", Source: "chart/band.go", Symbol: "IntervalBand", Upstream: "https://github.com/NimbleMarkets/ntcharts", Controls: []string{"state", "sigma", "legend"}, States: []string{"ready", "empty"}, Defaults: DefaultConfig(), Factory: bandStory},
	)
	for i := range d {
		switch d[i].ID {
		case "text/code":
			d[i].ExampleSource = "docs/snippets/code.go.txt"
		case "input/text":
			d[i].States = []string{"ready", "disabled"}
		case "feedback/spinner":
			d[i].Clock = true
			d[i].Step = true
		case "feedback/time":
			d[i].Clock = true
		case "feedback/progress":
			d[i].Step = true
		case "chart/sparkline":
			d[i].ExampleSource = "docs/snippets/sparkline.go.txt"
		case "chart/timeseries":
			d[i].Source = "examples/timeseries/main.go"
			d[i].Symbol = ""
		}
	}

	return d
}

func codeStory(e Environment) (runtime.Scene, error) {
	return static(func(s layout.Size) string {
		t := e.Config.Palette()

		block, err := text.CodeBlock("// Shared terminal defaults\nfield := input.NewText(options)\ncmd := field.Focus()", "go", s.Width, t)
		if err != nil {
			return err.Error()
		}

		return title("Code", e) + "Call " + text.InlineCode("field.Focus()", t) + " to focus the input.\n\n" + block
	}), nil
}

func textStory(e Environment) (runtime.Scene, error) {
	return static(func(layout.Size) string {
		t := e.Config.Palette()

		return title("Text primitives", e) + text.Label(e.Config.Value, t) + "\n\n" + text.Badge("READY", t) + " " + text.Badge("3 SELECTED", t) + "\n\n" +
			text.Pill("Success", text.Success, t) + " " + text.Pill("Info", text.Info, t) + "\n" + text.Pill("Warning", text.Warning, t) + " " + text.Pill("Danger", text.Danger, t) + "\n\n" +
			text.KeyHint("Enter", "select", t) + "   " + text.KeyHint("Esc", "dismiss", t)
	}), nil
}

func inputStory(e Environment) (runtime.Scene, error) {
	m := input.NewText(input.TextOptions{Placeholder: "Type a value", Value: e.Config.Value, Width: layout.PanelContentSize(e.Config.Size()).Width, Disabled: e.Config.State == "disabled", Theme: e.inputPalette()})

	if e.capture {
		_ = m.Focus()
	}

	return &scene{init: m.Focus, resize: func(s layout.Size) { m.Resize(layout.PanelContentSize(s).Width) }, update: func(msg tea.Msg) tea.Cmd {
		var cmd tea.Cmd

		m, cmd = m.Update(msg)
		if k, ok := msg.(tea.KeyPressMsg); ok && k.String() == "enter" {
			emit(e, interaction.Submitted{Component: "input", Value: m.Value()})
		}

		return cmd
	}, view: func(s layout.Size) string {
		return title("Text input", e) + layout.Panel("Value", m.View(), layout.Size{Width: s.Width, Height: 3}, e.Config.Palette(), m.Focused()) + footer("Type to edit / Enter emits Submitted / Ctrl+C exits", e)
	}}, nil
}

func areaStory(e Environment) (runtime.Scene, error) {
	m := input.NewArea(e.Config.Columns, e.Config.Rows-5, e.inputPalette())
	m.Placeholder = "Write a note"
	m.SetValue(e.Config.Value)

	if e.capture {
		_ = m.Focus()
	}

	return &scene{init: m.Focus, resize: func(s layout.Size) { m.SetWidth(s.Width); m.SetHeight(max(1, s.Height-5)) }, update: func(msg tea.Msg) tea.Cmd { var cmd tea.Cmd; m, cmd = m.Update(msg); return cmd }, view: func(layout.Size) string { return title("Multiline input", e) + m.View() }}, nil
}

func spinnerStory(e Environment) (runtime.Scene, error) {
	m := feedback.NewSpinner(e.Config.Palette())

	return &scene{init: func() tea.Cmd {
		if e.Config.Live && !e.Config.Palette().ReducedMotion {
			return m.Tick
		}

		return nil
	}, update: func(msg tea.Msg) tea.Cmd {
		if _, ok := msg.(runtime.StepMsg); ok {
			msg = spinner.TickMsg{ID: m.ID()}
		}

		var cmd tea.Cmd

		m, cmd = m.Update(msg)

		if !e.Config.Live {
			return nil
		}

		return cmd
	}, view: func(layout.Size) string {
		return title("Loading indicator", e) + m.View() + "  Processing fixture data" + footer("Live runs the native clock / Ctrl+N advances a frame", e)
	}}, nil
}

func progressStory(e Environment) (runtime.Scene, error) {
	amount := e.Config.Amount
	m := feedback.NewProgress(e.Config.Columns-10, e.Config.Palette())

	return &scene{resize: func(s layout.Size) { m.SetWidth(max(1, s.Width-10)) }, update: func(msg tea.Msg) tea.Cmd {
		if _, ok := msg.(runtime.StepMsg); ok {
			amount = min(100, amount+5)
			emit(e, map[string]any{"type": "progress", "amount": amount})
		}

		return nil
	}, view: func(layout.Size) string {
		return title("Progress", e) + m.ViewAs(amount/100) + fmt.Sprintf("  %.0f%%", amount) + footer("Ctrl+N advances by five percent.", e)
	}}, nil
}

func timeStory(e Environment) (runtime.Scene, error) {
	timer := feedback.NewTimer(time.Minute)
	watch := feedback.NewStopwatch()

	return &scene{init: func() tea.Cmd {
		if e.Config.Live {
			return tea.Batch(timer.Init(), watch.Init())
		}

		return nil
	}, update: func(msg tea.Msg) tea.Cmd {
		var a, b tea.Cmd

		timer, a = timer.Update(msg)
		watch, b = watch.Update(msg)

		if !e.Config.Live {
			return nil
		}

		return tea.Batch(a, b)
	}, view: func(layout.Size) string {
		return title("Time components", e) + "Countdown  " + timer.View() + "\nStopwatch  " + watch.View() + footer("Enable Live to run the native timers.", e)
	}}, nil
}

func viewportStory(e Environment) (runtime.Scene, error) {
	var lines []string
	for i := range 80 {
		lines = append(lines, fmt.Sprintf("%02d  A scrollable line of fixture content.", i+1))
	}

	m := collection.NewViewport(strings.Join(lines, "\n"), e.Config.Columns, e.Config.Rows-4)

	return &scene{resize: func(s layout.Size) { m.SetWidth(s.Width); m.SetHeight(max(1, s.Height-4)) }, update: func(msg tea.Msg) tea.Cmd { var cmd tea.Cmd; m, cmd = m.Update(msg); return cmd }, view: func(layout.Size) string { return title("Viewport", e) + m.View() }}, nil
}

func pictureStory(e Environment) (runtime.Scene, error) {
	img := image.NewRGBA(image.Rect(0, 0, 120, 60))

	for y := range 60 {
		for x := range 120 {
			img.SetRGBA(x, y, color.RGBA{R: uint8(x * 2), G: uint8(70 + y*2), B: uint8(180 - x), A: 255})
		}
	}

	plot, init := chart.NewPicture(img, layout.Size{Width: e.Config.Columns, Height: e.Config.Rows - 5})

	return &scene{init: func() tea.Cmd { return init }, resize: func(s layout.Size) { plot.SetSize(s.Width, max(1, s.Height-5)) }, update: plot.Update, view: func(layout.Size) string {
		if e.Config.Palette().ASCII {
			return title("Picture / text fallback", e) + "Embedded gradient fixture: 120 x 60 pixels"
		}

		return title("Picture / portable glyph rendering", e) + plot.View().Content + footer("This preview uses half-block glyphs. Kitty graphics are capability-dependent.", e)
	}}, nil
}

func bandStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		if e.Config.State == "empty" {
			return "No interval estimates"
		}

		var intervals []chart.Interval
		for k := e.Config.Sigma; k >= 1; k-- {
			intervals = append(intervals, chart.Interval{Low: 50 - float64(k)*8, High: 50 + float64(k)*8})
		}

		selected := 78.0
		options := chart.BandOptions{Min: 0, Max: 100, Center: 50, Intervals: intervals, Selected: &selected}

		view := fmt.Sprintf("Nested intervals / +/- %ds\n\n", e.Config.Sigma) + chart.IntervalBand(options, size.Width, e.Config.Palette()) + "\n" + "0" + strings.Repeat(" ", max(0, size.Width-len("0")-len("100"))) + "100"
		if e.Config.Legend {
			view += "\n\n" + options.Legend(size.Width, e.Config.Palette())
		}

		return view
	}), nil
}

func signedNumberStory(e Environment) (runtime.Scene, error) {
	return static(func(layout.Size) string {
		t := e.Config.Palette()
		emphasis := text.Emphasis(e.Config.Emphasis)
		value := (e.Config.Amount - 50) / 100
		standard := text.SignedNumber{Value: value, Available: e.Config.State != "empty", Precision: 3, Unit: "pp", Emphasis: emphasis}
		serverErrors := standard
		serverErrors.Tones = text.SignTones{Positive: text.Danger, Negative: text.Success}
		zero := text.SignedNumber{Available: true, Precision: 3, Unit: "pp", Emphasis: emphasis}
		tiny := text.SignedNumber{Value: -.000001, Available: true, Precision: 3, Unit: "pp", Emphasis: emphasis}
		missing := text.SignedNumber{Available: false}

		return strings.Join([]string{"Signed number / " + e.Config.Emphasis, "", "Positive means improvement   " + standard.View(t), "Higher server-error rate     " + serverErrors.View(t), "Zero                        " + zero.View(t), "Small nonzero               " + tiny.View(t), "Unavailable                 " + missing.View(t)}, "\n")
	}), nil
}

func pillsStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		t := e.Config.Palette()
		rows := []string{"Pills / tone and emphasis", ""}

		for _, emphasis := range []text.Emphasis{text.Bright, text.Normal, text.Dim} {
			line := []string{fmt.Sprintf("%-6s", emphasis)}
			if size.Width < 70 {
				rows = append(rows, string(emphasis))
				line = nil
			}

			for _, tone := range []text.Tone{text.Neutral, text.Success, text.Info, text.Warning, text.Danger} {
				if size.Width < 70 && len(line) == 3 {
					rows = append(rows, strings.Join(line, "  "))
					line = nil
				}

				line = append(line, text.Pill(string(tone), tone, t, emphasis))
			}

			rows = append(rows, strings.Join(line, "  "))
		}

		rows = append(rows, "", "Selected emphasis  "+text.Pill(e.Config.Value, text.Info, t, text.Emphasis(e.Config.Emphasis)))

		return strings.Join(rows, "\n")
	}), nil
}

func statusStory(e Environment) (runtime.Scene, error) {
	return static(func(layout.Size) string {
		t := e.Config.Palette()

		rows := []string{"Status atoms / " + e.Config.Emphasis, ""}
		for _, status := range feedback.Statuses() {
			rows = append(rows, status.View(text.Emphasis(e.Config.Emphasis), t))
		}

		return strings.Join(rows, "\n")
	}), nil
}

func atomConfig() Config { c := DefaultConfig(); c.Columns = 80; c.Rows = 12; return c }
