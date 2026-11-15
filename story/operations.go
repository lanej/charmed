package story

import (
	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/examples/operations/fixture"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/runtime"
)

func operationalDefinitions() []Definition {
	return []Definition{
		{ID: "chart/buckets", Title: "Molecules/Charts/Bucket counts", Level: "Molecules", Family: "chart", Description: "Compare complete interval counts with discrete bars. Partial intervals retain exact counts without a comparable bar; missing and unknown intervals never plot as zero.", Source: "chart/buckets.go", Symbol: "BucketBars", Composes: []string{"foundations/theme"}, Defaults: func() Config { c := DefaultConfig(); c.Columns = 64; c.Rows = 8; return c }(), Factory: bucketsStory},
		{ID: "template/operations", Title: "Templates/Operational dashboard", Level: "Templates", Family: "pattern", Description: "Start a monitoring screen with summary facts, primary evidence and secondary inspection. One scrolling region, contextual bindings and unbordered sections preserve hierarchy at narrow widths.", Source: "pattern/operational.go", Symbol: "OperationalDashboard", Controls: []string{"state"}, States: []string{"ready", "loading", "empty", "error"}, Composes: []string{"layout/text-row", "layout/key-value", "collection/viewport", "text/key-hint"}, Defaults: func() Config { c := DefaultConfig(); c.Columns = 100; c.Rows = 26; return c }(), Factory: operationsStory},
	}
}

func operationsStory(e Environment) (runtime.Scene, error) {
	t := e.Config.Palette()
	m := pattern.NewOperationalDashboard(fixture.Snapshot(e.Config.State, t), e.Config.Size(), t)

	if e.Config.VerboseDetails {
		m.Update(tea.KeyPressMsg{Code: 'i'})
	}

	return &scene{update: m.Update, resize: m.Resize, view: func(layout.Size) string { return m.View() }}, nil
}

func bucketsStory(e Environment) (runtime.Scene, error) {
	bars := chart.BucketBars{Buckets: []chart.CountBucket{
		{Label: "22:00 UTC", Count: 160, State: chart.BucketComplete},
		{Label: "23:00 UTC", Count: 173, State: chart.BucketComplete},
		{Label: "00:00 UTC", Count: 0, State: chart.BucketPartial},
		{Label: "01:00 UTC", State: chart.BucketMissing},
		{Label: "02:00 UTC"},
	}}

	return static(func(size layout.Size) string { return bars.View(size.Width, e.Config.Palette()) }), nil
}
