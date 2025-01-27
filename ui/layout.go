package ui

import "fyne.io/fyne/v2/container"

func Setup(app *AppInit) {
	swatchesContainer := BuildSwatches(app)
	colorPicker := SetupColoPicker(app)

	appLayout := container.NewBorder(nil, swatchesContainer, nil, colorPicker, app.PixlCanvas) // top, bottom, left, right, all else is in center

	app.PixlWindow.SetContent(appLayout)
}