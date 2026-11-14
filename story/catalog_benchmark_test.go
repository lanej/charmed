package story_test

import (
	"fmt"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/story"
	"github.com/lanej/charmed/theme"
)

// BenchmarkCatalog follows the registry so new native fixtures cannot silently
// miss construction, warm rendering and resize coverage.
func BenchmarkCatalog(b *testing.B) {
	for _, definition := range story.Definitions() {
		b.Run(definition.ID, func(b *testing.B) {
			for _, geometry := range [][2]int{{40, 12}, {120, 40}, {200, 60}} {
				b.Run(fmt.Sprintf("%dx%d", geometry[0], geometry[1]), func(b *testing.B) {
					config := definition.Defaults
					config.Columns, config.Rows = geometry[0], geometry[1]

					b.Run("construct", func(b *testing.B) {
						b.ReportAllocs()

						for b.Loop() {
							model, err := story.New(definition.ID, config, nil)
							if err != nil {
								b.Fatal(err)
							}

							renderedFrame = model.View().Content
							if err = model.Close(); err != nil {
								b.Fatal(err)
							}
						}
					})

					model, err := story.New(definition.ID, config, nil)
					if err != nil {
						b.Fatal(err)
					}

					b.Cleanup(func() {
						if closeErr := model.Close(); closeErr != nil {
							b.Error(closeErr)
						}
					})

					renderedFrame = model.View().Content

					b.Run("warm", func(b *testing.B) {
						b.ReportAllocs()

						for b.Loop() {
							renderedFrame = model.View().Content
						}
					})
					b.Run("height", func(b *testing.B) {
						b.ReportAllocs()

						i := 0
						for b.Loop() {
							_, _ = model.Update(tea.WindowSizeMsg{Width: config.Columns, Height: config.Rows + i%2})
							presentCatalog(b, model)
							renderedFrame = model.View().Content
							i++
						}
					})
					b.Run("width", func(b *testing.B) {
						b.ReportAllocs()

						i := 0
						for b.Loop() {
							_, _ = model.Update(tea.WindowSizeMsg{Width: config.Columns + i%2, Height: config.Rows})
							presentCatalog(b, model)
							renderedFrame = model.View().Content
							i++
						}
					})
					b.Run("interaction", func(b *testing.B) {
						b.ReportAllocs()

						i := 0

						for b.Loop() {
							code := tea.KeyDown
							if i%2 != 0 {
								code = tea.KeyUp
							}

							_, _ = model.Update(tea.KeyPressMsg{Code: code})
							presentCatalog(b, model)
							renderedFrame = model.View().Content
							i++
						}
					})
				})
			}

			for _, state := range definition.States {
				if state == "ready" {
					continue
				}

				b.Run("state-"+state, func(b *testing.B) {
					config := definition.Defaults
					config.State = state

					model, err := story.New(definition.ID, config, nil)
					if err != nil {
						b.Fatal(err)
					}

					b.Cleanup(func() {
						if closeErr := model.Close(); closeErr != nil {
							b.Error(closeErr)
						}
					})
					b.ReportAllocs()

					for b.Loop() {
						renderedFrame = model.View().Content
					}
				})
			}

			b.Run("ascii", func(b *testing.B) {
				config := definition.Defaults
				config.ASCII = true

				model, err := story.New(definition.ID, config, nil)
				if err != nil {
					b.Fatal(err)
				}

				b.Cleanup(func() {
					if closeErr := model.Close(); closeErr != nil {
						b.Error(closeErr)
					}
				})
				b.ReportAllocs()

				for b.Loop() {
					renderedFrame = model.View().Content
				}
			})

			for _, mode := range []theme.Mode{theme.Light, theme.Contrast, theme.Mono} {
				b.Run(string(mode), func(b *testing.B) {
					config := definition.Defaults
					config.Theme = mode

					model, err := story.New(definition.ID, config, nil)
					if err != nil {
						b.Fatal(err)
					}

					b.Cleanup(func() {
						if closeErr := model.Close(); closeErr != nil {
							b.Error(closeErr)
						}
					})
					b.ReportAllocs()

					for b.Loop() {
						renderedFrame = model.View().Content
					}
				})
			}
		})
	}
}

// The baseline runtime renders immediately; scheduled runtimes expose a snapshot
// operation. Both paths measure a completed frame after each workload update.
func presentCatalog(tb testing.TB, model interface{ View() tea.View }) {
	tb.Helper()

	if presenter, ok := model.(interface{ PresentNow() error }); ok {
		if err := presenter.PresentNow(); err != nil {
			tb.Fatal(err)
		}
	}
}
