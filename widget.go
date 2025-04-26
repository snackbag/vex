package vex

import "image/color"

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

	GetClasses() []string
	SetStyle(key string, value interface{})
	GetStyle(key string) interface{}
	GetStyleAsColor(key string) color.RGBA
	GetStyleAsString(key string) string
	GetStyleAsInt(key string) int
	GetStyleAsFloat32(key string) float32
	GetStyleAsFloat64(key string) float64
}

type VBaseWidget struct {
	VWidget

	x      int32
	y      int32
	width  int32
	height int32

	classes []string
}

func (w *VBaseWidget) Render() {
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

func (w *VBaseWidget) GetClasses() []string {
	return w.classes
}

func (w *VBaseWidget) SetStyle(key string, value interface{}) {
	Process.StyleSheet.SetKey(&w.VWidget, key, value)
}

func (w *VBaseWidget) GetStyle(key string) interface{} {
	return Process.StyleSheet.GetKeyRaw(&w.VWidget, key)
}

func (w *VBaseWidget) GetStyleAsColor(key string) color.RGBA {
	return Process.StyleSheet.GetKeyAsColor(&w.VWidget, key)
}

func (w *VBaseWidget) GetStyleAsString(key string) string {
	return Process.StyleSheet.GetKeyAsString(&w.VWidget, key)
}

func (w *VBaseWidget) GetStyleAsInt(key string) int {
	return Process.StyleSheet.GetKeyAsInt(&w.VWidget, key)
}

func (w *VBaseWidget) GetStyleAsFloat32(key string) float32 {
	return Process.StyleSheet.GetKeyAsFloat32(&w.VWidget, key)
}

func (w *VBaseWidget) GetStyleAsFloat64(key string) float64 {
	return Process.StyleSheet.GetKeyAsFloat64(&w.VWidget, key)
}
