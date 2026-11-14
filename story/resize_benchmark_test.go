package story_test

import (
	"strconv"
	"testing"

	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/theme"

	tea "charm.land/bubbletea/v2"

	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/story"
)

func BenchmarkResizeRender(b *testing.B) {
	for _, id := range []string{"git/browser", "pattern/log-inspector", "text/markdown", "chart/comparison-panel"} {
		b.Run(id, func(b *testing.B) {
			definition, err := story.Find(id)
			if err != nil {
				b.Fatal(err)
			}

			model, err := story.New(id, definition.Defaults, nil)
			if err != nil {
				b.Fatal(err)
			}

			b.Cleanup(func() {
				if closeErr := model.Close(); closeErr != nil {
					b.Error(closeErr)
				}
			})

			for _, scenario := range []struct {
				name  string
				sizes []layout.Size
			}{
				{"same", []layout.Size{{Width: 120, Height: 40}}},
				{"height", []layout.Size{{Width: 120, Height: 40}, {Width: 120, Height: 30}}},
				{"width", []layout.Size{{Width: 120, Height: 40}, {Width: 110, Height: 40}, {Width: 100, Height: 40}}},
				{"breakpoint", []layout.Size{{Width: 120, Height: 40}, {Width: 40, Height: 12}, {Width: 200, Height: 60}}},
			} {
				b.Run(scenario.name, func(b *testing.B) {
					b.ReportAllocs()

					i := 0
					for b.Loop() {
						size := scenario.sizes[i%len(scenario.sizes)]
						_, _ = model.Update(tea.WindowSizeMsg{Width: size.Width, Height: size.Height})
						presentCatalog(b, model)
						renderedFrame = model.View().Content
						i++
					}
				})
			}
		})
	}
}

func BenchmarkPaneResize(b *testing.B) {
	items := make([]gitui.Item, 1000)
	for i := range items {
		items[i] = gitui.Item{ID: strconv.Itoa(i), Text: "A supplied Git observation with a stable identity"}
	}

	for _, scenario := range []struct {
		name  string
		sizes []layout.Size
	}{
		{"height", []layout.Size{{Width: 80, Height: 20}, {Width: 80, Height: 35}}},
		{"width", []layout.Size{{Width: 80, Height: 20}, {Width: 60, Height: 20}}},
	} {
		b.Run(scenario.name, func(b *testing.B) {
			list, err := gitui.NewList("benchmark", items, 80, 20, theme.New(theme.Dark))
			if err != nil {
				b.Fatal(err)
			}

			resize := list.Resize
			if geometry, ok := any(list).(interface{ ResizeViewport(layout.Size) }); ok {
				resize = geometry.ResizeViewport
			}

			b.ReportAllocs()

			i := 0
			for b.Loop() {
				resize(scenario.sizes[i%len(scenario.sizes)])

				renderedFrame = list.View()
				i++
			}
		})
	}
}
