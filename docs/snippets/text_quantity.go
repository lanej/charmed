package snippets

import (
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextQuantity() string {
	quantity := text.Quantity{
		Value: 1258291, Unit: "B", Scale: text.BinaryScale,
		Precision: 1, Width: 16, Emphasis: text.Normal,
	}

	return quantity.View(theme.New(theme.Dark))
}
