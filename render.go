package vex

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var RenderThreadQueue = make(chan func(), 1024)

func (process *VProcess) startRenderLoop() {
	isFirstStart := true
	prevMouseStateLeft := false
	prevMouseStateMiddle := false
	prevMouseStateRight := false

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

		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && !prevMouseStateLeft {
			prevMouseStateLeft = true
			fireEventIfHovered(rl.GetMouseX(), rl.GetMouseY(), "left-click")
		}

		if rl.IsMouseButtonReleased(rl.MouseButtonLeft) && prevMouseStateLeft {
			prevMouseStateLeft = false
			fireEventIfHovered(rl.GetMouseX(), rl.GetMouseY(), "left-release")
		}

		if rl.IsMouseButtonPressed(rl.MouseButtonMiddle) && !prevMouseStateMiddle {
			prevMouseStateMiddle = true
			fireEventIfHovered(rl.GetMouseX(), rl.GetMouseY(), "middle-click")
		}

		if rl.IsMouseButtonReleased(rl.MouseButtonMiddle) && prevMouseStateMiddle {
			prevMouseStateMiddle = false
			fireEventIfHovered(rl.GetMouseX(), rl.GetMouseY(), "middle-release")
		}

		if rl.IsMouseButtonPressed(rl.MouseButtonRight) && !prevMouseStateRight {
			prevMouseStateRight = true
			fireEventIfHovered(rl.GetMouseX(), rl.GetMouseY(), "right-click")
		}

		if rl.IsMouseButtonReleased(rl.MouseButtonRight) && prevMouseStateRight {
			prevMouseStateRight = false
			fireEventIfHovered(rl.GetMouseX(), rl.GetMouseY(), "right-release")
		}

		rl.BeginDrawing()
		rl.ClearBackground(process.BackgroundColor)

		for _, widget := range process.widgets {
			widget.Render()
		}

		rl.EndDrawing()
	}
}

func fireEventIfHovered(x int32, y int32, event string) {
	for _, widget := range Process.widgets {
		if rl.CheckCollisionPointRec(rl.NewVector2(float32(x), float32(y)), widget.GenerateHitbox()) {
			widget.FireEvent(event)
		}
	}
}

func DoOnRenderThread(task func()) {
	RenderThreadQueue <- task
}
