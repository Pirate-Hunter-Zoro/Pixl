package util

import (
	"image"
	"image/color"
)

func GetImageColors(img image.Image) map[color.Color]struct{} {
	colors := make(map[color.Color]struct{}) // emulate a set using maps
	var empty struct{}
	for y:=0; y<img.Bounds().Dy(); y++ {
		for x:=0; x<img.Bounds().Dx(); x++ {
			colors[img.At(x, y)] = empty // doesn't matter what the value is for our map - we only care to emulate a set
		}
	}
	return colors
}