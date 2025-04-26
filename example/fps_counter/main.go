package main

import (
	"fmt"
	"github.com/snackbag/vex"
	"path"
)

func main() {
	process := vex.Init("FPS Counter", 400, 400)
	process.LoadFont("arial", path.Join("example", "assets", "Arial.ttf"), 16)

	label := vex.NewLabel("FPS: ?")
	label.SetStyle("font-name", "arial")
	process.AddWidget(label)

	vex.DoEvery(1000, func(iteration int) {
		label.Text = fmt.Sprintf("FPS: %d", vex.GetFPS())
	})

	process.Run()
}
