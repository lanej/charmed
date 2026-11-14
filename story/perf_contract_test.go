package story_test

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/story"
	"github.com/lanej/charmed/theme"
)

func TestCatalogPerformanceWorkloadInputs(t *testing.T) {
	seen := map[string]bool{}
	for _, definition := range story.Definitions() {
		if seen[definition.ID] {
			t.Fatalf("duplicate fixture %s", definition.ID)
		}

		seen[definition.ID] = true
		t.Run(definition.ID, func(t *testing.T) {
			for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
				config := definition.Defaults
				config.Theme = mode

				model, err := story.New(definition.ID, config, nil)
				if err != nil {
					t.Fatal(err)
				}

				for _, size := range [][2]int{{40, 12}, {120, 40}, {200, 60}, {120, 40}} {
					_, _ = model.Update(tea.WindowSizeMsg{Width: size[0], Height: size[1]})

					frame := model.View().Content
					if lipgloss.Width(frame) != size[0] || lipgloss.Height(frame) != size[1] {
						t.Fatalf("unbounded native fixture at %v/%s", size, mode)
					}
				}

				if err = model.Close(); err != nil {
					t.Fatal(err)
				}
			}
		})
	}

	if len(seen) == 0 {
		t.Fatal("empty performance inventory")
	}
}

func TestPreparedMarkdownAllocationBudget(t *testing.T) {
	// The observed warm frame uses about 32,000 allocations at this fixed size.
	// Keep a small toolchain margin; timing is reported separately.
	definition, err := story.Find("text/markdown")
	if err != nil {
		t.Fatal(err)
	}

	config := definition.Defaults
	config.Columns, config.Rows = 120, 40

	model, err := story.New(definition.ID, config, nil)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if closeErr := model.Close(); closeErr != nil {
			t.Error(closeErr)
		}
	}()

	model.View()

	allocations := testing.AllocsPerRun(5, func() { renderedFrame = model.View().Content })
	t.Logf("warm Markdown allocations: %.0f", allocations)

	if allocations > 33600 {
		t.Fatalf("warm Markdown allocation budget exceeded: %.0f", allocations)
	}
}

func TestChartResizeMatchesFreshFixture(t *testing.T) {
	for _, kind := range []string{"sparkline", "timeseries", "line", "scatter", "bars", "heatmap", "stream", "wave", "ohlc", "canvas"} {
		id := "chart/" + kind

		definition, err := story.Find(id)
		if err != nil {
			t.Fatal(err)
		}

		model, err := story.New(id, definition.Defaults, nil)
		if err != nil {
			t.Fatal(err)
		}

		for _, size := range [][2]int{{120, 40}, {80, 24}, {200, 60}, {120, 40}} {
			config := definition.Defaults
			config.Columns, config.Rows = size[0], size[1]

			fresh, constructErr := story.New(id, config, nil)
			if constructErr != nil {
				t.Fatal(constructErr)
			}

			_, _ = model.Update(tea.WindowSizeMsg{Width: size[0], Height: size[1]})
			if model.View().Content != fresh.View().Content {
				gotRows, wantRows := strings.Split(model.View().Content, "\n"), strings.Split(fresh.View().Content, "\n")
				for row := range gotRows {
					if gotRows[row] != wantRows[row] {
						t.Fatalf("resized %s at %v row %d got %q want %q", id, size, row, gotRows[row], wantRows[row])
					}
				}
			}

			if err = fresh.Close(); err != nil {
				t.Fatal(err)
			}
		}

		if err = model.Close(); err != nil {
			t.Fatal(err)
		}
	}
}
