package main

import (
	"github.com/snackbag/vex"
	"path"
)

func main() {
	process := vex.Init("image", 400, 400)
	process.SetAllowResize(true)

	image := vex.NewImage(path.Join("example", "assets", "image.png"))
	process.AddUpdateListener(func() { // also gets called on start
		image.SetSize(process.GetWidth(), 211)
	})
	process.AddWidget(image)

	process.Run()
}
