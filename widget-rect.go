package vex

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/snackbag/vex/extra"
	"image/color"
)

type VRect struct {
	VBaseWidget
}

func (r *VRect) Render() {
	rect := extra.GenerateFloat32Rectangle(r.X(), r.Y(), r.Width(), r.Height())
	bgColor := r.GetStyleAsColor("background-color")
	borderRoundness := r.GetStyleAsFloat32("border-roundness")
	borderWidth := r.GetStyleAsInt("border-width")
	borderColor := r.GetStyleAsColor("border-color")
	borderSegments := int32(r.GetStyleAsInt("border-segments"))

	if borderRoundness <= 0 {
		rl.DrawRectangle(r.X(), r.Y(), r.Width(), r.Height(), bgColor)
	} else {
		rl.DrawRectangleRounded(rect, borderRoundness, borderSegments, bgColor)

		if borderWidth > 0 {
			rl.DrawRectangleRoundedLinesEx(rect, borderRoundness, borderSegments, float32(borderWidth), borderColor)
		}
	}
}

func NewRect(color color.RGBA) *VRect {
	rect := &VRect{}
	rect.width = 50
	rect.height = 50

	rect.SetStyle("background-color", color)
	rect.SetStyle("border-roundness", 0)
	rect.SetStyle("border-width", 0)
	rect.SetStyle("border-color", ColorO(0, 0, 0, 0))
	rect.SetStyle("border-segments", 1)

	return rect
}
