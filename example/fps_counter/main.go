package main

import (
	"fmt"
	"github.com/snackbag/vex"
)

func main() {
	process := vex.Init("FPS Counter", 400, 400)

	label := vex.NewLabel("FPS: ?")
	process.AddWidget(label)

	vex.DoEvery(1000, func(iteration int) {
		label.Text = fmt.Sprintf("FPS: %d", vex.GetFPS())
	})

	process.Run()
}
