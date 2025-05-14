package main

import (
	"fmt"
	"github.com/snackbag/vex"
	"path"
)

func main() {
	process := vex.Init("Button", 400, 400)
	process.LoadFont("arial", path.Join("example", "assets", "Arial.ttf"), 16)

	// the actual button
	button := vex.NewLabel("Click me!")
	button.SetStyle("font-name", "arial")
	button.SetStyle("background-color", vex.Color(255, 0, 0))

	// Clicking functionality
	button.RegisterOnLeftClick(func() {
		fmt.Println("Clicked!")

		button.Text = "I was clicked"
	})

	// Also add
	process.AddWidget(button)

	process.Run()
}
