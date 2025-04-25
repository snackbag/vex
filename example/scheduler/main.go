package main

import (
	"fmt"
	"github.com/snackbag/vex"
)

func main() {
	vex.DoEvery(1000, func(iteration int) {
		fmt.Println("Test from 1")
	})

	vex.DoEveryWithDelay(500, 2000, func(iteration int) {
		fmt.Println("Test from 2")
	})

	vex.DoEveryWithDelayAndIterations(2000, 5000, 5, func(iteration int) {
		fmt.Println("Test from 3")
	})

	for {
	}
}
