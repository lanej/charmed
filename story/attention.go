package story

import (
	"fmt"

	"github.com/lanej/charmed/attention"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
)

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
		d := chart.DistributionComparison{Series: []chart.DistributionSeries{{Label: "R", Mean: 20, StdDev: 3, Members: 64, HasSpread: true}, {Label: "C", Mean: 30, StdDev: 8, Members: 32, HasSpread: true, Selected: &selected}}, Sigma: e.Config.Sigma, Min: 0, Max: 60, Unit: "%", SpreadUnit: "pp"}
		if e.Config.Boundaries {
			d.Markers = []chart.BandMarker{{Value: 26, Glyph: 'W', Style: feedback.AttentionStyle(attention.Warn, e.Config.Palette())}, {Value: 29, Glyph: 'E', Style: feedback.AttentionStyle(attention.Error, e.Config.Palette())}}
		}
		return title("Composed cohort distributions", e) + d.View(size.Width, e.Config.Palette()) + "\n\nW/E = optional caller-supplied warn/error boundaries\nX = selected observation; | = mean\nThis molecule composes interval bands, labels and a shared axis."
	}), nil
}
