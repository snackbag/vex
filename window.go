package vex

import rl "github.com/gen2brain/raylib-go/raylib"

var hasProcess = false

type VProcess struct {
	Title  string
	Width  int32
	Height int32

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

func (process *VProcess) startRenderLoop() {
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.EndDrawing()
	}
}

func (process *VProcess) Run() {
	process.startRenderLoop()
}

func Init(title string, width int32, height int32) *VProcess {
	if hasProcess {
		panic("Cannot create multiple Vex processes")
	}

	val := &VProcess{title, width, height, make([]VWidget, 0)}
	hasProcess = true

	rl.InitWindow(width, height, title)
	rl.SetTargetFPS(60)
	return val
}
