package story

import (
	"fmt"
	"strings"
	"testing"

	"github.com/lanej/charmed/layout"
)

func TestStaticFixtureReusesOnlyTheCurrentAllocation(t *testing.T) {
	calls := 0

	m := static(func(size layout.Size) string {
		calls++
		return fmt.Sprintf("%dx%d", size.Width, size.Height)
	})
	for _, width := range []int{10, 20, 10} {
		m.Resize(layout.Size{Width: width, Height: 2})

		before := calls

		first, second := m.View(), m.View()
		if first != second || calls != before+1 || !strings.HasPrefix(first, fmt.Sprintf("%dx2", width)) {
			t.Fatal("static fixture did not retain exactly its current allocation")
		}

		m.Update(struct{}{})

		if m.View() != first || calls != before+1 {
			t.Fatal("unhandled message rebuilt a static fixture")
		}
	}
}

func TestInteractiveSceneStillReadsCurrentNativeState(t *testing.T) {
	value := "before"
	m := &scene{view: func(layout.Size) string { return value }}
	m.Resize(layout.Size{Width: 20, Height: 2})
	m.View()

	value = "after"

	if !strings.HasPrefix(m.View(), "after") {
		t.Fatal("interactive scene view was cached")
	}
}
