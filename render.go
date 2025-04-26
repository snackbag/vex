package vex

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var RenderThreadQueue = make(chan func(), 1024)

func (process *VProcess) startRenderLoop() {
	isFirstStart := true
	hasMouseMoved := false
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
			Process.width = int32(rl.GetScreenWidth())
			Process.height = int32(rl.GetScreenHeight())

			Process.EventHandler.FireEvent("update")
		}

		mouseX := rl.GetMouseX()
		mouseY := rl.GetMouseY()

		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && !prevMouseStateLeft {
			prevMouseStateLeft = true
			fireEventIfHovered(mouseX, mouseY, "left-click")
		}

		if rl.IsMouseButtonReleased(rl.MouseButtonLeft) && prevMouseStateLeft {
			prevMouseStateLeft = false
			fireEventIfHovered(mouseX, mouseY, "left-release")
		}

		if rl.IsMouseButtonPressed(rl.MouseButtonMiddle) && !prevMouseStateMiddle {
			prevMouseStateMiddle = true
			fireEventIfHovered(mouseX, mouseY, "middle-click")
		}

		if rl.IsMouseButtonReleased(rl.MouseButtonMiddle) && prevMouseStateMiddle {
			prevMouseStateMiddle = false
			fireEventIfHovered(mouseX, mouseY, "middle-release")
		}

		if rl.IsMouseButtonPressed(rl.MouseButtonRight) && !prevMouseStateRight {
			prevMouseStateRight = true
			fireEventIfHovered(mouseX, mouseY, "right-click")
		}

		if rl.IsMouseButtonReleased(rl.MouseButtonRight) && prevMouseStateRight {
			prevMouseStateRight = false
			fireEventIfHovered(mouseX, mouseY, "right-release")
		}

		if hasMouseMoved {
			Process.EventHandler.FireEvent("mouse-move")
		}

		if !hasMouseMoved && mouseX != 0 && mouseY != 0 {
			hasMouseMoved = true
		}

		rl.BeginDrawing()
		rl.ClearBackground(process.BackgroundColor)

		for _, widget := range process.widgets {
			widget.Render()
		}

		rl.EndDrawing()
		isFirstStart = false
	}
}

func fireEventIfHovered(x int32, y int32, event string) {
	for _, widget := range Process.widgets {
		if rl.CheckCollisionPointRec(rl.NewVector2(float32(x), float32(y)), widget.GenerateHitbox()) {
			widget.FireEvent(event)
		}
	}
}

func doIfHovered(x int32, y int32, runnable func(widget VWidget)) {
	for _, widget := range Process.widgets {
		if rl.CheckCollisionPointRec(rl.NewVector2(float32(x), float32(y)), widget.GenerateHitbox()) {
			runnable(widget)
		}
	}
}

func DoOnRenderThread(task func()) {
	RenderThreadQueue <- task
}
