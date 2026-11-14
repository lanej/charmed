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

func TestJSONSelectionReusesSyntaxAndViewportCache(t *testing.T) {
	m := NewJSONViewer(80, 20, theme.New(theme.Dark))
	if err := m.SetSource([]byte(`{"a":1,"b":2,"c":3}`)); err != nil {
		t.Fatal(err)
	}

	builds := m.syntaxBuilds
	for i := range 30 {
		m.Select(1 + i%3)
		m.View()
	}

	if m.syntaxBuilds != builds {
		t.Fatal("selection repeated syntax preparation")
	}

	if allocations := testing.AllocsPerRun(10, func() { m.View() }); allocations != 0 {
		t.Fatalf("warm JSON view allocated %.0f times", allocations)
	}

	m.Theme = theme.New(theme.Mono)
	m.Invalidate()

	if m.syntaxBuilds != builds+1 {
		t.Fatal("appearance invalidation did not replace syntax")
	}
}
