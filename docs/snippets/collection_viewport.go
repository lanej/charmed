package snippets

import (
	"github.com/lanej/charmed/collection"
)

func ExampleCollectionViewport() string {
	content := "API gateway: healthy\nEvent stream: healthy\nJob workers: degraded"
	viewport := collection.NewViewport(content, 48, 6)
	// Retain the model and command returned by viewport.Update(msg) to enable scrolling.
	return viewport.View()
}
