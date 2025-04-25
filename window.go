package vex

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

var hasProcess = false

type VProcess struct {
	Title  string
	Width  int32
	Height int32

	BackgroundColor color.RGBA

	widgets []VWidget
}

func (process *VProcess) Show() {
	process.SetVisibility(true)
}

func (process *VProcess) Hide() {
	process.SetVisibility(false)
}

func (process *VProcess) SetVisibility(visibility bool) {
	setWindowState(rl.FlagWindowHidden, visibility)
}

func (process *VProcess) SetAllowResize(allow bool) {
	setWindowState(rl.FlagWindowResizable, allow)
}

func (process *VProcess) Quit() {
	rl.CloseWindow()
}

func (process *VProcess) Run() {
	process.startRenderLoop()
}

func (process *VProcess) AddWidget(widget VWidget) {
	process.widgets = append(process.widgets, widget)
}

func Init(title string, width int32, height int32) *VProcess {
	if hasProcess {
		panic("Cannot create multiple Vex processes")
	}

	val := &VProcess{title, width, height, ColorAll(255), make([]VWidget, 0)}
	hasProcess = true

	rl.SetTraceLogLevel(rl.LogError)
	rl.InitWindow(width, height, title)
	rl.SetTargetFPS(60)
	return val
}
