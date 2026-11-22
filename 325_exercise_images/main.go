package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	h, w int
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (i Image) At(x int, y int) color.Color {
	v := uint8(x ^ y) // это не степень?
	return color.RGBA{v, v, 255, 255}
}
func main() {
	m := Image{256, 1000}
	pic.ShowImage(m)
}
