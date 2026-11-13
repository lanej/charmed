package story_test

import (
	"testing"

	"github.com/lanej/charmed/story"
)

var renderedFrame string

func BenchmarkRender(b *testing.B) {
	for _, id := range []string{"git/browser", "pattern/log-inspector", "screen/services", "collection/data-grid", "text/markdown", "chart/comparison-panel"} {
		b.Run(id, func(b *testing.B) {
			definition, err := story.Find(id)
			if err != nil {
				b.Fatal(err)
			}

			config := definition.Defaults
			config.Columns, config.Rows = 120, 40

			model, err := story.New(id, config, nil)
			if err != nil {
				b.Fatal(err)
			}

			b.Cleanup(func() {
				if err := model.Close(); err != nil {
					b.Error(err)
				}
			})
			b.ReportAllocs()

			for b.Loop() {
				renderedFrame = model.View().Content
			}
		})
	}
}
