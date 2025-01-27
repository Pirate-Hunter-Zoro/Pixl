package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type PxCanvasRenderer struct {
	pxCanvas *PxCanvas
	canvasImage *canvas.Image
	canvasBorder []canvas.Line // array of lines
	canvasCursor []fyne.CanvasObject
}

func (renderer *PxCanvasRenderer) SetCursor(objects []fyne.CanvasObject) {
	renderer.canvasCursor = objects
}

// WidgetRenderer interface implementation.
func (renderer *PxCanvasRenderer) MinSize() fyne.Size {
	return renderer.pxCanvas.DrawingArea
}

// WidgetRenderer interface implementation
func (renderer *PxCanvasRenderer) Objects() []fyne.CanvasObject {
	objects := make([]fyne.CanvasObject, 0, 5)
	for i := 0; i < len(renderer.canvasBorder); i++ { // four canvas borders
		objects = append(objects, &renderer.canvasBorder[i])
	}
	objects = append(objects, renderer.canvasImage) // actual image itself
	objects = append(objects, renderer.canvasCursor...)
	return objects
}

// WidgetRenderer interface implementation
func (renderer *PxCanvasRenderer) Destroy() {}

// WidgetRenderer interface implementation
func (renderer *PxCanvasRenderer) Layout(size fyne.Size) {
	renderer.LayoutCanvas(size) // call first because that resizes the image which musts be accurate before the border
	renderer.LayoutBorder(size)
}

// WidgetRenderer interface implementation
func (renderer *PxCanvasRenderer) Refresh() {
	if renderer.pxCanvas.reloadImage {
		renderer.canvasImage = canvas.NewImageFromImage(renderer.pxCanvas.PixelData)
		renderer.canvasImage.ScaleMode = canvas.ImageScalePixels // pixel-perfect scaling
		renderer.canvasImage.FillMode = canvas.ImageFillContain
		renderer.pxCanvas.reloadImage = false
	}
	renderer.Layout(renderer.pxCanvas.Size()) // consume TOTAL drawing area - plenty of room to move our pixel canvas around
	canvas.Refresh(renderer.canvasImage) // update everything on screen
}

func (renderer *PxCanvasRenderer) LayoutCanvas(size fyne.Size) {
	imgPxWidth := renderer.pxCanvas.PxCols
	imgPxHeight := renderer.pxCanvas.PxRows
	pxSize := renderer.pxCanvas.PxSize
	// account for left and top offest
	renderer.canvasImage.Move(fyne.NewPos(renderer.pxCanvas.CanvasOffset.X, renderer.pxCanvas.CanvasOffset.Y))
	renderer.canvasImage.Resize(fyne.NewSize(float32(imgPxWidth*pxSize), float32(imgPxHeight*pxSize))) // resize to make pixels bigger
}

func (renderer *PxCanvasRenderer) LayoutBorder(size fyne.Size) {
	offset := renderer.pxCanvas.CanvasOffset
	imgHeight := renderer.canvasImage.Size().Height
	imgWidth := renderer.canvasImage.Size().Width

	// left border of inner painting box
	left := &renderer.canvasBorder[0]
	left.Position1 = fyne.NewPos(offset.X, offset.Y)
	left.Position2 = fyne.NewPos(offset.X, offset.Y+imgHeight)

	// top border of inner painting box
	top := &renderer.canvasBorder[1]
	top.Position1 = fyne.NewPos(offset.X, offset.Y)
	top.Position2 = fyne.NewPos(offset.X + imgWidth, offset.Y)

	// right border of inner painting box
	right := &renderer.canvasBorder[2]
	right.Position1 = fyne.NewPos(offset.X + imgWidth, offset.Y)
	right.Position2 = fyne.NewPos(offset.X + imgWidth, offset.Y + imgHeight)

	// bottom border of inner painting box
	bottom := &renderer.canvasBorder[3]
	bottom.Position1 = fyne.NewPos(offset.X, offset.Y + imgHeight)
	bottom.Position2 = fyne.NewPos(offset.X + imgWidth, offset.Y + imgHeight)
}