package vex

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

type VRect struct {
	VBaseWidget

	Color color.RGBA
}

func (r *VRect) Render() {
	rl.DrawRectangle(r.X(), r.Y(), r.Width(), r.Height(), r.Color)
}

func NewRect(color color.RGBA) *VRect {
	rect := &VRect{Color: color}
	rect.width = 50
	rect.height = 50
	return rect
}
