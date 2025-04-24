package vex

type VWidget interface {
	Render(process *VProcess)

	X() int
	Y() int
	Width() int
	Height() int
}
