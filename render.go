package vex

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var RenderThreadQueue = make(chan func(), 1024)

func (process *VProcess) startRenderLoop() {
	isFirstStart := true

	for !rl.WindowShouldClose() {
		select {
		case task := <-RenderThreadQueue:
			task()
		default:

		}

		if rl.IsWindowResized() || isFirstStart {
			isFirstStart = false

			Process.width = int32(rl.GetScreenWidth())
			Process.height = int32(rl.GetScreenHeight())

			Process.EventHandler.FireEvent("update")
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
