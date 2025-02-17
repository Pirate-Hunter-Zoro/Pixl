package swatch

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type SwatchRenderer struct {
	square canvas.Rectangle
	objects []fyne.CanvasObject
	parent *Swatch
}

func (renderer *SwatchRenderer) MinSize() fyne.Size {
	return renderer.square.MinSize()
}

func (renderer *SwatchRenderer) Layout(size fyne.Size) {
	renderer.objects[0].Resize(size)
}

func (renderer *SwatchRenderer) Refresh() {
	renderer.Layout(fyne.NewSize(20,20))
	renderer.square.FillColor = renderer.parent.Color
	if renderer.parent.Selected {
		renderer.square.StrokeWidth = 3
		renderer.square.StrokeColor = color.NRGBA{255,255,255,255}
		renderer.objects[0] = &renderer.square // refresh objects[0]
	} else { // remove the stroke
		renderer.square.StrokeWidth = 0
		renderer.objects[0] = &renderer.square // refresh objects[0]
	}
	canvas.Refresh(renderer.parent) // redraw this widget
}

func (renderer *SwatchRenderer) Objects() []fyne.CanvasObject {
	return renderer.objects
}

func (renderer *SwatchRenderer) Destroy() {
	// Can be empty - just need for implementation
}