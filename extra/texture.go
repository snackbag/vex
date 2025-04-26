package extra

import rl "github.com/gen2brain/raylib-go/raylib"

type Texture struct {
	Link   *rl.Texture2D
	Usages int
}

// PotentiallyOptimize Returns whether the texture should also be removed from the registry
func (t *Texture) PotentiallyOptimize() bool {
	if t.Usages <= 0 {
		rl.UnloadTexture(*t.Link)
		return true
	}

	return false
}
