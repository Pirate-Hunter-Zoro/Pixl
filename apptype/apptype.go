package apptype

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type BrushType = int

type PxCanvasConfig struct { // A struct with different elements
	DrawingArea fyne.Size
	CanvasOffset fyne.Position
	PxRows, PxCols int
	PxSize int
}

type State struct {
	BrushColor color.Color
	BrushType int
	SwatchSelected int
	FilePath string
}

func (state *State) SetFilePath(path string) {
	// struct method to change the file path
	state.FilePath = path
}