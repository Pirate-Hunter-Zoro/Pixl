package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"zerotomastery.io/pixl/pxcanvas/brush"
)

// Implement Scrollable interface
func (pxCanvas *PxCanvas) Scrolled(ev *fyne.ScrollEvent) {
	pxCanvas.scale(int(ev.Scrolled.DY))
	pxCanvas.Refresh()
}

// Implement Hoverable interface
func (pxCanvas *PxCanvas) MouseOut() {}
func (pxCanvas *PxCanvas) MouseIn(ev *desktop.MouseEvent) {}
func (pxCanvas *PxCanvas) MouseMoved(ev *desktop.MouseEvent) {
	// first check if mouse is over canvas
	if x, y := pxCanvas.MouseToCanvasXY(ev); x != nil && y != nil {
		// enables click & drag since this is in the 'MouseMoved' function
		brush.TryBrush(pxCanvas.appState, pxCanvas, ev)
		// recall x and y are pointers to integers
		cursor := brush.Cursor(pxCanvas.PxCanvasConfig, pxCanvas.appState.BrushType, ev, *x, *y)
		pxCanvas.renderer.SetCursor(cursor)
	} else {
		// hide the cursor if out of bounds
		pxCanvas.renderer.SetCursor(make([]fyne.CanvasObject, 0))
	}

	pxCanvas.TryPan(pxCanvas.mouseState.previousCoord, ev)
	pxCanvas.Refresh() // update ui
	pxCanvas.mouseState.previousCoord = &ev.PointEvent // update previous coordinate to be this new coordinate
}

// Implement the Mouseable interface
func (pxCanvas *PxCanvas) MouseUp(ev *desktop.MouseEvent) {}
func (pxCanvas *PxCanvas) MouseDown(ev *desktop.MouseEvent) {
	brush.TryBrush(pxCanvas.appState, pxCanvas, ev)
}
