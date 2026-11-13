package gitui

import (
	"fmt"
	"testing"

	tea "charm.land/bubbletea/v2"
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
