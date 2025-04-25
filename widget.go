package vex

type VWidget interface {
	Render()

	X() int32
	Y() int32
	move(x int32, y int32)
	SetWidth(width int32)
	SetHeight(width int32)
	SetSize(width int32, height int32)
	SetSizeAll(all int32)
	Width() int32
	Height() int32
}

type VBaseWidget struct {
	x      int32
	y      int32
	width  int32
	height int32
}

func (w *VBaseWidget) X() int32 {
	return w.x
}

func (w *VBaseWidget) Y() int32 {
	return w.y
}

func (w *VBaseWidget) move(x int32, y int32) {
	w.x = x
	w.y = y
}

func (w *VBaseWidget) SetWidth(width int32) {
	w.width = width
}

func (w *VBaseWidget) SetHeight(height int32) {
	w.height = height
}

func (w *VBaseWidget) SetSize(width int32, height int32) {
	w.SetWidth(width)
	w.SetHeight(height)
}

func (w *VBaseWidget) SetSizeAll(all int32) {
	w.SetSize(all, all)
}

func (w *VBaseWidget) Width() int32 {
	return w.width
}

func (w *VBaseWidget) Height() int32 {
	return w.height
}
