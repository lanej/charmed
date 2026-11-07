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
		case "sparkline", "stream", "wave", "ohlc":
			controls = append(controls, "amount")
		case "bars":
			controls = append(controls, "series", "horizontal")
		}
		d = append(d, Definition{ID: "chart/" + kind.id, Title: "Atoms/Charts/" + kind.name, Level: "Atoms", Family: "chart", Description: "Real ntcharts " + kind.name + " rendering with semantic styling and deterministic fixture data.", Source: "story/atoms_charts.go", Symbol: "chartStory", Upstream: "https://github.com/NimbleMarkets/ntcharts", Controls: controls, Clock: true, Step: true, Defaults: DefaultConfig(), Factory: func(e Environment) (runtime.Scene, error) { return chartStory(kind.id, e) }})
	}
	d = append(d,
		Definition{ID: "chart/picture", Title: "Atoms/Graphics/Picture", Level: "Atoms", Family: "graphics", Description: "Portable glyph rendering of an embedded image; no external image service.", Source: "story/atoms.go", Symbol: "pictureStory", Upstream: "https://github.com/NimbleMarkets/ntcharts", Defaults: DefaultConfig(), Factory: pictureStory},
		Definition{ID: "chart/interval-band", Title: "Atoms/Charts/Interval band", Level: "Atoms", Family: "chart", Description: "Nested caller-supplied intervals, a mean and selected observation rendered on an ntcharts canvas.", Source: "chart/band.go", Symbol: "IntervalBand", Upstream: "https://github.com/NimbleMarkets/ntcharts", Controls: []string{"state", "sigma"}, States: []string{"ready", "empty"}, Defaults: DefaultConfig(), Factory: bandStory},
	)
	for i := range d {
		switch d[i].ID {
		case "input/text":
			d[i].States = []string{"ready", "disabled"}
		case "feedback/spinner":
			d[i].Clock = true
			d[i].Step = true
		case "feedback/time":
			d[i].Clock = true
		case "feedback/progress":
			d[i].Step = true
		case "chart/timeseries":
			d[i].Source = "examples/timeseries/main.go"
			d[i].Symbol = ""
		}
	}
	return d
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
	m := input.NewText(input.TextOptions{Placeholder: "Type a value", Value: e.Config.Value, Width: e.Config.Columns - 4, Disabled: e.Config.State == "disabled", Theme: e.Config.Palette()})
	return &scene{init: m.Focus, resize: func(s layout.Size) { m.Resize(s.Width - 4) }, update: func(msg tea.Msg) tea.Cmd {
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
	m := input.NewArea(e.Config.Columns, e.Config.Rows-5, e.Config.Palette())
	m.Placeholder = "Write a note"
	m.SetValue(e.Config.Value)
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
		return fmt.Sprintf("Nested intervals / +/- %ds\n\n", e.Config.Sigma) + chart.IntervalBand(chart.BandOptions{Min: 0, Max: 100, Center: 50, Intervals: intervals, Selected: &selected}, size.Width, e.Config.Palette()) +
			"\n0                                              100\n\n| center / X selected observation\nOuter to inner: 3s / 2s / 1s\nIntervals are supplied by the caller; this atom makes no statistical claim."
	}), nil
}
