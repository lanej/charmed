package gitui

import (
	"fmt"
	"strings"
	"testing"

	"github.com/lanej/charmed/theme"
)

func TestListBurstRetainsEverySelectionAndBuildsRowsOnce(t *testing.T) {
	items := make([]Item, 200)
	for i := range items {
		items[i] = Item{ID: fmt.Sprint(i), Text: fmt.Sprintf("item-%03d", i)}
	}

	m, err := NewList("burst", items, 40, 8, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	initial := m.View()

	before := m.rowBuilds
	for range 100 {
		m.Update(press("j"))
	}

	if m.SelectedID() != "100" {
		t.Fatal("selection updates lost")
	}

	if m.rowBuilds != before {
		t.Fatal("input burst repeatedly projected native rows")
	}

	frame := m.View()
	if frame == initial || !strings.Contains(frame, "item-100") || m.rowBuilds != before+1 {
		t.Fatal("latest selection not presented in one row build")
	}

	fresh, err := NewList("burst", items, 40, 8, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	fresh.Select("100")

	if frame != fresh.View() {
		t.Fatal("retained selection differs from fresh construction")
	}

	if allocations := testing.AllocsPerRun(10, func() { m.View() }); allocations != 0 {
		t.Fatalf("warm list read allocated %.0f times", allocations)
	}

	m.Invalidate()
	m.View()

	if m.rowBuilds != before+2 {
		t.Fatal("explicit invalidation did not refresh native presentation")
	}
}

func TestNativeStyleInvalidationPreservesIndependentScroll(t *testing.T) {
	items := make([]Item, 100)
	for i := range items {
		items[i] = Item{ID: fmt.Sprint(i), Text: fmt.Sprint(i)}
	}

	m, err := NewList("native", items, 40, 8, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	m.View()
	m.Viewport.SetYOffset(50)
	m.Invalidate()
	m.View()

	if m.Viewport.YOffset() != 50 {
		t.Fatal("style invalidation followed selection instead of preserving native scroll")
	}

	m.Update(press("j"))
	m.View()

	if m.Viewport.YOffset() >= 50 {
		t.Fatal("real selection no longer revealed after style invalidation")
	}
}

func TestReplacementRevealsSelectionAfterStyleInvalidation(t *testing.T) {
	items := make([]Item, 100)
	for i := range items {
		items[i] = Item{ID: fmt.Sprint(i), Text: fmt.Sprint(i)}
	}

	m, err := NewList("replace", items, 40, 8, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	m.View()
	m.Invalidate()

	items[0], items[99] = items[99], items[0]
	if err := m.SetItems(items); err != nil {
		t.Fatal(err)
	}

	m.View()

	if m.SelectedID() != "0" || m.Viewport.YOffset() == 0 {
		t.Fatal("source replacement failed to reveal the retained selected ID")
	}
}
