package layout_test

import (
	"strings"
	"testing"

	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

var fittedFrame string

func BenchmarkFit(b *testing.B) {
	palette := theme.New(theme.Dark)
	for _, tc := range []struct{ name, content string }{
		{"plain", strings.Repeat(strings.Repeat("a", 120)+"\n", 39) + strings.Repeat("a", 120)},
		{"styled", strings.Repeat(palette.Heading().Render(strings.Repeat("a", 120))+"\n", 39) + palette.Heading().Render(strings.Repeat("a", 120))},
		{"overflow", strings.Repeat(strings.Repeat("a", 240)+"\n", 10000)},
	} {
		b.Run(tc.name, func(b *testing.B) {
			b.ReportAllocs()

			for b.Loop() {
				fittedFrame = layout.Fit(tc.content, layout.Size{Width: 120, Height: 40})
			}
		})
	}
}
