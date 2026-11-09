package story

import (
	"fmt"
	"math"
	"strings"

	"charm.land/bubbles/v2/table"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/text"
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

		var rows []table.Row
		if e.Config.State == "empty" {
			rows = nil
		} else {
			ready := w.Parts[0].Count
			rows = []table.Row{{"checkout", fmt.Sprint(min(24, ready)), "ready"}, {"labels", fmt.Sprint(min(28, max(0, ready-24))), "ready"}, {"rating", fmt.Sprint(max(0, ready-52)), "ready"}, {"tracking", fmt.Sprint(w.Parts[1].Count), "pending"}, {"events", fmt.Sprint(w.Parts[2].Count), "failed"}}
		}

		list := collection.NewTable([]table.Column{{Title: "SERVICE", Width: max(10, size.Width-29)}, {Title: "INSTANCES", Width: 10}, {Title: "STATE", Width: 10}}, rows, size.Width, 7, t)

		return w.StatusHeader("PRODUCTION", layout.Size{Width: size.Width, Height: 2}, t) + "\n\n" + list.View()
	}), nil
}

func waffleConfig() Config { c := atomConfig(); c.Rows = 16; c.Amount = 92; return c }
