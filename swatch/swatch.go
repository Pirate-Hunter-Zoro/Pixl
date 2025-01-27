package swatch

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"zerotomastery.io/pixl/apptype"
)

type Swatch struct {
	widget.BaseWidget // extension of this struct
	Selected bool
	Color color.Color
	SwatchIndex int
	clickHandler func(s *Swatch) // a method of this swatch object
}

func (s *Swatch) SetColor(c color.Color) {
	s.Color = c
	s.Refresh() // that way screen is always up to date
}

func NewSwatch(state *apptype.State, color color.Color, swatchIndex int, clickHandler func(s *Swatch)) *Swatch {
	swatch := &Swatch{
		Selected: false,
		Color: color,
		clickHandler: clickHandler,
		SwatchIndex: swatchIndex,
	}
	swatch.ExtendBaseWidget(swatch) // supplies state information - part of BaseWidget

	return swatch
}

func (swatch *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(swatch.Color)
	objects := []fyne.CanvasObject{square}
	return &SwatchRenderer{
		square: *square,
		objects: objects,
		parent: swatch,
	}
}