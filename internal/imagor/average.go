package imagor

import "image/color"

// AverageColor represents one table cell data.
type AverageColor struct {
	R     uint
	G     uint
	B     uint
	A     uint
	Count uint
}

func NewAverageColor() AverageColor {
	return AverageColor{}
}

func (a *AverageColor) Add(r, g, b, al uint32) {
	a.R += uint(r)
	a.G += uint(g)
	a.B += uint(b)
	a.A += uint(al)
	a.Count++
}

// Color generates average color of pixels for cell.
func (a *AverageColor) Color() color.RGBA {
	if a.Count == 0 {
		return color.RGBA{}
	}

	return color.RGBA{
		R: uint8(a.R >> 8 / a.Count),
		G: uint8(a.G >> 8 / a.Count),
		B: uint8(a.B >> 8 / a.Count),
		A: uint8(a.A >> 8 / a.Count),
	}
}
