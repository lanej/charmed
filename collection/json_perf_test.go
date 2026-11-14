package collection

import (
	"testing"

	"github.com/lanej/charmed/theme"
)

func TestJSONGeometryKeepsOneSyntaxProjection(t *testing.T) {
	m := NewJSONViewer(80, 20, theme.New(theme.Dark))
	if err := m.SetSource([]byte(`{"value":"a long wrapped value","nested":{"id":42}}`)); err != nil {
		t.Fatal(err)
	}

	original := m.Model.GetContent()
	for _, width := range []int{20, 100, 40, 80} {
		m.ResizeViewport(width, 20)
	}

	if m.Model.GetContent() != original || len(m.preparedSyntax) == 0 {
		t.Fatal("width reflow changed prepared content")
	}

	m.Theme = theme.New(theme.Mono)
	m.Refresh()

	fresh := NewJSONViewer(80, 20, theme.New(theme.Mono))
	if err := fresh.SetSource(m.source); err != nil {
		t.Fatal(err)
	}

	if m.Model.GetContent() != fresh.Model.GetContent() {
		t.Fatal("explicit refresh retained old syntax theme")
	}

	m.Clear()

	if m.preparedSyntax != nil || m.fallbackSyntax != "" {
		t.Fatal("cleared document retained syntax")
	}
}
