package vex

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
	"os"
)

var Process *VProcess = nil

type VProcess struct {
	Title  string
	Width  int32
	Height int32

	BackgroundColor color.RGBA
	StyleSheet      *VStyleSheet

	widgets []VWidget
	fonts   map[string]*rl.Font
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

func Init(title string, width int32, height int32) *VProcess {
	if Process != nil {
		panic("Cannot create multiple Vex processes")
	}

	val := &VProcess{title, width, height, ColorAll(255), newStyleSheet(), make([]VWidget, 0), make(map[string]*rl.Font)}
	Process = val
	Process.StyleSheet.widgetSpecificStyles = make(map[*VWidget]map[string]interface{})

	rl.SetTraceLogLevel(rl.LogError)
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(width, height, title)
	rl.SetTargetFPS(60)
	return val
}
