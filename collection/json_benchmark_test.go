package collection_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/theme"
)

var jsonFrame string

func BenchmarkJSON(b *testing.B) {
	for _, size := range []int{1024, 16384, collection.DefaultJSONLimits().Bytes} {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			source := []byte(`{"payload":"` + strings.Repeat("a", size-14) + `"}`)

			m := collection.NewJSONViewer(80, 24, theme.New(theme.Dark))
			if err := m.SetSource(source); err != nil {
				b.Fatal(err)
			}

			for _, operation := range []string{"height", "width", "replace"} {
				b.Run(operation, func(b *testing.B) {
					b.ReportAllocs()

					i := 0

					for b.Loop() {
						switch operation {
						case "height":
							m.ResizeViewport(80, 20+i%2*5)
						case "width":
							m.ResizeViewport(60+i%2*20, 24)
						case "replace":
							source[len(source)-4] = byte('a' + i%2)
							if err := m.SetSource(source); err != nil {
								b.Fatal(err)
							}
						}

						jsonFrame = m.View()
						i++
					}
				})
			}
		})
	}
}
