package collection_test

import (
	"fmt"
	"testing"

	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/theme"
)

func BenchmarkLogIngestion(b *testing.B) {
	for _, count := range []int{100, 1000, 10000} {
		b.Run(fmt.Sprint(count), func(b *testing.B) {
			records := make([]collection.HTTPExchange, count)
			for i := range records {
				records[i] = collection.HTTPExchange{ID: fmt.Sprint(i), Method: "GET", URL: "https://example.test/requests", StatusCode: 200, Completed: true}
			}

			for _, operation := range []string{"replace", "update-batch"} {
				b.Run(operation, func(b *testing.B) {
					m := collection.NewHTTPLog(100, 30, theme.New(theme.Dark))

					m.Limit = count
					if err := m.SetEntries(records); err != nil {
						b.Fatal(err)
					}

					b.ReportAllocs()

					i := 0

					for b.Loop() {
						if operation == "replace" {
							if err := m.SetEntries(records); err != nil {
								b.Fatal(err)
							}
						} else {
							batch := append([]collection.HTTPExchange(nil), records[:min(10, len(records))]...)
							for j := range batch {
								batch[j].StatusCode = 200 + i%2*300
							}

							if err := m.Upsert(batch...); err != nil {
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
