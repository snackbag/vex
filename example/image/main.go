package main

import (
	"github.com/snackbag/vex"
	"path"
)

func main() {
	process := vex.Init("image", 400, 400)

	image := vex.NewImage(path.Join("example", "assets", "image.png"))
	image.SetSize(362, 211)
	process.AddWidget(image)

	process.Run()
}
