package ui

import "fyne.io/fyne/v2/container"

func Setup(app *AppInit) {
	swatchesContainer := BuildSwatches(app)
	colorPicker := SetupColoPicker(app)

	appLayout := container.NewBorder(nil, swatchesContainer, nil, colorPicker) // top, bottom, left, right

	app.PixlWindow.SetContent(appLayout)
}