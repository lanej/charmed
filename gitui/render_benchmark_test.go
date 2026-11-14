package gitui

import (
	"fmt"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func BenchmarkListIgnoredKey(b *testing.B) {
	palette := theme.New(theme.Dark)

	items := make([]Item, 1000)
	for i := range items {
		items[i] = Item{ID: fmt.Sprint(i), Text: palette.Text().Render(fmt.Sprintf("file-%d.go modified", i))}
	}

	model, err := NewList("benchmark", items, 80, 24, palette)
	if err != nil {
		b.Fatal(err)
	}

	message := tea.KeyPressMsg{Code: 'x', Text: "x"}

	b.ReportAllocs()

	for b.Loop() {
		model.Update(message)
	}
}

func BenchmarkListLifecycle(b *testing.B) {
	for _, count := range []int{100, 1000, 10000} {
		b.Run(fmt.Sprint(count), func(b *testing.B) {
			items := make([]Item, count)
			for i := range items {
				items[i] = Item{ID: fmt.Sprint(i), Text: fmt.Sprintf("file-%d.go modified", i)}
			}

			for _, operation := range []string{"selection", "scroll", "width", "filter", "replace"} {
				b.Run(operation, func(b *testing.B) {
					m, err := NewList("benchmark", items, 80, 24, theme.New(theme.Dark))
					if err != nil {
						b.Fatal(err)
					}

					if operation == "filter" {
						m.Update(press("/"))
					}

					b.ReportAllocs()

					i := 0

					for b.Loop() {
						switch operation {
						case "selection":
							if i%2 == 0 {
								m.Update(press("j"))
							} else {
								m.Update(press("k"))
							}
						case "scroll":
							if i%2 == 0 {
								m.Update(tea.KeyPressMsg{Code: 'e', Mod: tea.ModCtrl})
							} else {
								m.Update(tea.KeyPressMsg{Code: 'y', Mod: tea.ModCtrl})
							}
						case "width":
							m.ResizeViewport(layout.Size{Width: 60 + i%2*20, Height: 24})
						case "filter":
							m.Input.SetValue(fmt.Sprint(i % 10))
							m.Update(tea.KeyPressMsg{Code: tea.KeyLeft})
						case "replace":
							if err = m.SetItems(items); err != nil {
								b.Fatal(err)
							}
						}

						listFrame = m.View()
						i++
					}
				})
			}
		})
	}
}

var listFrame string
