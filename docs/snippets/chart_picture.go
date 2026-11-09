package snippets

import (
	"image"
	"image/color"

	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
)

func ExampleChartPicture() string {
	gradient := image.NewRGBA(image.Rect(0, 0, 120, 60))

	for y := range 60 {
		for x := range 120 {
			gradient.SetRGBA(x, y, color.RGBA{
				R: uint8(x * 2), G: uint8(70 + y*2), B: uint8(180 - x), A: 255,
			})
		}
	}

	picture, init := chart.NewPicture(gradient, layout.Size{Width: 40, Height: 10})
	_ = init // Return this command from your Bubble Tea Init method.

	return picture.View().Content
}
