package vex

import rl "github.com/gen2brain/raylib-go/raylib"

var RenderThreadQueue = make(chan func())

func (process *VProcess) startRenderLoop() {
	for !rl.WindowShouldClose() {
		select {
		case task := <-RenderThreadQueue:
			task()
		default:

		}

		rl.BeginDrawing()
		rl.ClearBackground(process.BackgroundColor)

		for _, widget := range process.widgets {
			widget.Render()
		}

		rl.EndDrawing()
	}
}

func DoOnRenderThread(task func()) {
	RenderThreadQueue <- task
}
