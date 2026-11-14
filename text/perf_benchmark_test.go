package text_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

var documentFrame string

func BenchmarkMarkdown(b *testing.B) {
	for _, size := range []int{1024, 16384, 262144} {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			source := "# Prepared document\n\n" + strings.Repeat("Paragraph with **emphasis** and `code`.\n\n", size/40)

			document, err := text.PrepareMarkdown(source, theme.New(theme.Dark))
			if err != nil {
				b.Fatal(err)
			}

			b.Run("width", func(b *testing.B) {
				b.ReportAllocs()

				i := 0
				for b.Loop() {
					documentFrame, err = document.Render(60 + i%2*20)
					if err != nil {
						b.Fatal(err)
					}

					i++
				}
			})
		})
	}
}
