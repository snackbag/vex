package vex

import rl "github.com/gen2brain/raylib-go/raylib"

func (process *VProcess) startRenderLoop() {
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(process.BackgroundColor)

		for _, widget := range process.widgets {
			widget.Render(process)
		}

		rl.EndDrawing()
	}
}
