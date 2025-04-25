package main

import (
	"fmt"
	"github.com/snackbag/vex"
)

func main() {
	process := vex.Init("FPS Counter", 400, 400)

	vex.DoEvery(1000, func(iteration int) {
		fmt.Printf("Current FPS: %d\n", vex.GetFPS())

		if iteration == 3 {
			process.Hide()
		} else {
			process.Show()
		}
	})

	process.Run()
}
