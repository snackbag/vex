package vex

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

type VRect struct {
	VBaseWidget
}

func (r *VRect) Render() {
	rl.DrawRectangle(r.X(), r.Y(), r.Width(), r.Height(), r.GetStyleAsColor("background-color"))
}

func NewRect(color color.RGBA) *VRect {
	rect := &VRect{}
	rect.width = 50
	rect.height = 50

	rect.SetStyle("background-color", ColorAll(90))

	return rect
}
