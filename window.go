package vex

var hasProcess = false

type VProcess struct {
	Title  string
	Width  int
	Height int

	widgets []VWidget
}

func Boot(title string, width int, height int) *VProcess {
	if hasProcess {
		panic("Cannot create multiple Vex processes")
	}

	val := &VProcess{title, width, height, make([]VWidget, 0)}
	hasProcess = true
	return val
}
