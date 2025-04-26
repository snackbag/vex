package vex

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/snackbag/vex/extra"
	"os"
)

type VImage struct {
	VBaseWidget

	path    string
	texture *rl.Texture2D
}

func (i *VImage) Render() {
	rl.DrawTexturePro(
		*i.texture,
		extra.GenRec(0, 0, i.texture.Width, i.texture.Height),
		extra.GenRec(i.X(), i.Y(), i.Width(), i.Height()),
		rl.NewVector2(0, 0),
		0,
		i.GetStyleAsColor("image-tint"),
	)
}

func (i *VImage) SetPath(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("couldn't set path of image, because path does not exist (%s)", path)
	}

	if _, ok := Process.textures[path]; ok {
		Process.SafeUnloadTexture(i.path)
	}

	i.path = path
	definite := Process.RegisterOrGetTexture(path, func() rl.Texture2D {
		return rl.LoadTexture(path)
	})

	i.texture = definite
	return nil
}

func (i *VImage) GetPath() string {
	return i.path
}

func NewImage(path string) *VImage {
	img := &VImage{VBaseWidget: *NewBaseWidget()}

	if err := img.SetPath(path); err != nil {
		rl.TraceLog(rl.LogError, err.Error())
		return img
	}

	img.SetStyle("image-tint", ColorAll(255))
	img.width = img.texture.Width
	img.height = img.texture.Height

	return img
}
