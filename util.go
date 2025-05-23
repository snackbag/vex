package vex

import rl "github.com/gen2brain/raylib-go/raylib"

func setWindowState(flag uint32, allow bool) {
	DoOnRenderThread(func() {
		if allow {
			rl.SetWindowState(flag)
		} else {
			rl.ClearWindowState(flag)
		}
	})
}

func GetFPS() int {
	return int(rl.GetFPS())
}
