package vex

import rl "github.com/gen2brain/raylib-go/raylib"

type VLabel struct {
	VBaseWidget

	Text string
}

func (l *VLabel) Render() {
	var font rl.Font

	if f := Process.GetLoadedFont(l.GetStyleAsString("font-name")); f != nil {
		font = *f
	} else {
		font = rl.GetFontDefault()
	}

	size := int32(l.GetStyleAsInt("font-size"))
	color := l.GetStyleAsColor("color")
	spacing := float32(l.GetStyleAsInt("letter-spacing"))

	rl.DrawTextEx(font, l.Text, rl.NewVector2(float32(l.X()), float32(l.Y())), float32(size), spacing, color)
}

func NewLabel(text string) *VLabel {
	label := &VLabel{VBaseWidget: *NewBaseWidget(), Text: text}
	label.width = 100
	label.height = 16

	label.SetStyle("color", ColorAll(0))
	label.SetStyle("font-size", 16)
	label.SetStyle("font-name", "")
	label.SetStyle("letter-spacing", 1)
	return label
}
