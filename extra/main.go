package extra

import rl "github.com/gen2brain/raylib-go/raylib"

// GenRec previously known as GenerateFloat32Rectangle
func GenRec(x int32, y int32, width int32, height int32) rl.Rectangle {
	return rl.NewRectangle(float32(x), float32(y), float32(width), float32(height))
}
