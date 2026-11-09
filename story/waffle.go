package story

import (
	"fmt"
	"math"
	"strings"

	"charm.land/bubbles/v2/table"
	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func waffleFixture(e Environment) chart.Waffle {
	if e.Config.State == "empty" {
		return chart.Waffle{}
	}

	ready := int(math.Round(e.Config.Amount))
	failed := min(5, 100-ready)

	return chart.Waffle{Parts: []chart.WafflePart{
		{Label: "ready", Count: ready, Tone: text.Success},
		{Label: "pending", Count: 100 - ready - failed, Tone: text.Info},
		{Label: "failed", Count: failed, Tone: text.Danger},
	}}
}

func waffleStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		w := waffleFixture(e)
		t := e.Config.Palette()

		if len(w.Parts) == 0 {
			return title("Fleet health", e) + "No observations"
		}

		if size.Width < 48 {
			grid := w.View(layout.Size{Width: 19, Height: 5}, t)
			labels := strings.ReplaceAll(w.Legend(t), " / ", "\n")

			return title("Fleet health / 100 instances", e) + grid + "\n\n" + labels + "\n\n" + t.Subtle().Render("50 cells · 2 instances each")
		}

		grid := w.View(layout.Size{Width: 19, Height: 10}, t)
		details := []string{t.Heading().Render("100 instances"), t.Subtle().Render("Production fleet"), "", ""}

		for i, label := range strings.Split(w.Legend(t), " / ") {
			details = append(details, label+t.Subtle().Render(fmt.Sprintf("  %d%%", w.Parts[i].Count)), "")
		}

		return title("Fleet health", e) + layout.Row(grid, "    "+strings.Join(details, "\n    ")) + "\n\n" + t.Subtle().Render("100 cells · one instance per cell")
	}), nil
}

func waffleHeaderStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		w := waffleFixture(e)
		t := e.Config.Palette()

		statusMode := text.GlyphAndLabel
		if size.Width < 28 {
			statusMode = text.GlyphOnly
		}

		statusWidth := 5
		for _, status := range []feedback.Status{feedback.OK, feedback.Pending, feedback.Error} {
			statusWidth = max(statusWidth, lipgloss.Width(status.View(text.Normal, t, statusMode)))
		}

		magnitudeWidth := min(24, max(9, size.Width/3))
		serviceWidth := max(3, size.Width-statusWidth-magnitudeWidth-3)
		magnitudeWidth = max(3, size.Width-statusWidth-serviceWidth-3)
		scale := 1

		var rows []table.Row

		if len(w.Parts) > 0 {
			ready := w.Parts[0].Count
			counts := []int{min(24, ready), min(28, max(0, ready-24)), max(0, ready-52), w.Parts[1].Count, w.Parts[2].Count}
			names := []string{"checkout", "labels", "rating", "tracking", "events"}
			states := []feedback.Status{feedback.OK, feedback.OK, feedback.OK, feedback.Pending, feedback.Error}

			for _, count := range counts {
				scale = max(scale, count)
			}

			for i, count := range counts {
				rows = append(rows, table.Row{names[i], waffleMagnitude(count, scale, magnitudeWidth, t), states[i].View(text.Normal, t, statusMode)})
			}
		}

		list := collection.NewDenseTable([]table.Column{{Title: "SERVICE", Width: serviceWidth}, {Title: "INSTANCES", Width: magnitudeWidth}, {Title: "STATE", Width: statusWidth}}, rows, size.Width, 7, t)

		caption := ""
		if len(rows) > 0 {
			caption = t.Subtle().Render(fmt.Sprintf("Instance scale: 0–%d", scale))
		}

		return w.StatusHeader("PRODUCTION", layout.Size{Width: size.Width, Height: 2}, t) + "\n\n" + list.View() + "\n" + caption
	}), nil
}

func waffleMagnitude(count, scale, width int, t theme.Theme) string {
	label := fmt.Sprintf("%3d", count)
	if width < 8 {
		return label
	}

	bar := feedback.NewProgress(width-4, t)
	bar.EmptyColor = t.Border

	if !t.ASCII {
		bar.Full, bar.Empty = '━', '─'
	}

	return label + " " + bar.ViewAs(float64(count)/float64(scale))
}

func waffleConfig() Config { c := atomConfig(); c.Rows = 16; c.Amount = 92; return c }
