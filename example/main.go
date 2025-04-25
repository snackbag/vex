package main

import "github.com/snackbag/vex"

func main() {
	process := vex.Init("Test", 400, 400)
	process.SetAllowResize(true)

	rect := vex.NewRect(vex.ColorAll(0))
	process.AddWidget(rect)

	process.Run()
}
