package vex

import rl "github.com/gen2brain/raylib-go/raylib"

type VLabel struct {
	VBaseWidget

	Text string
}

func (l *VLabel) Render() {
	rl.DrawText(l.Text, l.X(), l.Y(), int32(l.GetStyleAsInt("font-size")), l.GetStyleAsColor("color"))
}

func NewLabel(text string) *VLabel {
	label := &VLabel{Text: text}
	label.width = 100
	label.height = 16

	label.SetStyle("color", ColorAll(0))
	label.SetStyle("font-size", 16)
	return label
}
