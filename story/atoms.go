package story

import (
	"fmt"
	"strings"

	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/text"
)

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
	return static(func(layout.Size) string {
		t := e.Config.Palette()
		rows := []string{"Pills / tone and emphasis", ""}
		for _, emphasis := range []text.Emphasis{text.Bright, text.Normal, text.Dim} {
			line := []string{string(emphasis)}
			for _, tone := range []text.Tone{text.Neutral, text.Success, text.Info, text.Warning, text.Danger} {
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
		for _, status := range []feedback.Status{feedback.OK, feedback.Warn, feedback.Error, feedback.Unknown, feedback.Off} {
			rows = append(rows, status.View(text.Emphasis(e.Config.Emphasis), t))
		}
		return strings.Join(rows, "\n")
	}), nil
}

func atomConfig() Config { c := DefaultConfig(); c.Columns = 80; c.Rows = 12; return c }
