package vex

import "image/color"

func Color(red uint8, green uint8, blue uint8) color.RGBA {
	return color.RGBA{R: red, G: green, B: blue, A: 255}
}

func ColorO(red uint8, green uint8, blue uint8, opacity float32) color.RGBA {
	return color.RGBA{R: red, G: green, B: blue, A: uint8(opacity * 255)}
}

func ColorAll(all uint8) color.RGBA {
	return Color(all, all, all)
}
