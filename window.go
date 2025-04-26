package vex

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/snackbag/vex/extra"
	"image/color"
	"os"
)

var Process *VProcess = nil

type VProcess struct {
	Title  string
	width  int32
	height int32

	BackgroundColor color.RGBA
	StyleSheet      *VStyleSheet

	widgets  []VWidget
	fonts    map[string]*rl.Font
	textures map[string]*extra.Texture

	updateListeners []func()
}

func (process *VProcess) GetWidth() int32 {
	return process.width
}

func (process *VProcess) GetHeight() int32 {
	return process.height
}

func (process *VProcess) SetWidth(width int32) {
	process.SetSize(width, process.height)
}

func (process *VProcess) SetHeight(height int32) {
	process.SetSize(process.width, height)
}

func (process *VProcess) SetSize(width int32, height int32) {
	rl.SetWindowSize(int(width), int(height))
}

func (process *VProcess) Show() {
	process.SetVisibility(true)
}

func (process *VProcess) Hide() {
	process.SetVisibility(false)
}

func (process *VProcess) SetVisibility(visibility bool) {
	setWindowState(rl.FlagWindowHidden, !visibility)
}

func (process *VProcess) SetAllowResize(allow bool) {
	setWindowState(rl.FlagWindowResizable, allow)
}

func (process *VProcess) Quit() {
	rl.CloseWindow()
}

func (process *VProcess) Run() {
	process.startRenderLoop()
}

func (process *VProcess) AddWidget(widget VWidget) {
	process.widgets = append(process.widgets, widget)
}

func (process *VProcess) LoadFont(name string, path string, resolution int) {
	if _, exists := process.fonts[name]; exists {
		panic(fmt.Sprintf("cannot load font named %s, because there is already a font loaded with the same name", name))
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		rl.TraceLog(rl.LogError, fmt.Sprintf("Couldn't find path '%s' for font '%s' (maybe check working directory?)", name, path))
	}

	font := rl.LoadFontEx(path, int32(resolution), nil, 0)
	process.fonts[name] = &font
}

func (process *VProcess) GetLoadedFont(name string) *rl.Font {
	if val, ok := process.fonts[name]; ok {
		return val
	}
	return nil
}

// RegisterOrGetTexture gets the texture from cache if path is already registered
func (process *VProcess) RegisterOrGetTexture(path string, tex func() rl.Texture2D) *rl.Texture2D {
	if reg, ok := process.textures[path]; ok {
		reg.Usages++
		return reg.Link
	}

	texture := tex()
	reg := &extra.Texture{Link: &texture, Usages: 1}
	return reg.Link
}

// SafeUnloadTexture only unloads texture if not used anywhere anymore
func (process *VProcess) SafeUnloadTexture(path string) {
	if reg, ok := process.textures[path]; ok {
		reg.Usages--
		reg.PotentiallyOptimize()
		return
	}

	rl.TraceLog(rl.LogError, fmt.Sprintf("Failed to unload texture %s, because it isn't cached", path))
}

func (process *VProcess) AddUpdateListener(listener func()) {
	process.updateListeners = append(process.updateListeners, listener)
}

func (process *VProcess) FireUpdateListeners() {
	for _, listener := range process.updateListeners {
		listener()
	}
}

func Init(title string, width int32, height int32) *VProcess {
	if Process != nil {
		panic("Cannot create multiple Vex processes")
	}

	val := &VProcess{title, width, height, ColorAll(255), newStyleSheet(), make([]VWidget, 0), make(map[string]*rl.Font), make(map[string]*extra.Texture), make([]func(), 0)}
	Process = val
	Process.StyleSheet.widgetSpecificStyles = make(map[*VWidget]map[string]interface{})

	rl.SetTraceLogLevel(rl.LogError)
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(width, height, title)
	rl.SetTargetFPS(60)
	return val
}
