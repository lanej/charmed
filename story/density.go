package story

import (
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
)

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
