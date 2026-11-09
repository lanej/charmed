package story

import (
	"math"

	"github.com/lanej/charmed/chart"
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
		return title("Waffle / 100 equal-area cells", e) + w.View(layout.Size{Width: min(10, size.Width), Height: 10}, e.Config.Palette()) + "\n\n" + w.Legend(e.Config.Palette())
	}), nil
}

func waffleHeaderStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		w := waffleFixture(e)
		return w.Header("Fleet", size.Width, e.Config.Palette()) + "\n\n" + "Exact counts accompany the rounded 20-cell summary.\nNarrow headers drop the grid before clipping the counts."
	}), nil
}

func waffleConfig() Config { c := atomConfig(); c.Rows = 16; c.Amount = 92; return c }
