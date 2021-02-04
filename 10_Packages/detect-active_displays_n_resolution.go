package main

import (
	"fmt"

	"github.com/kbinani/screenshot"
)

func main() {
	n := screenshot.NumActiveDisplays()

	fmt.Println("Number of active monitor(s) : ", n)

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		x := bounds.Dx()
		y := bounds.Dy()

		fmt.Printf("Display #%d resolution is %d x %d\n", i, x, y)
	}
	res, err := screenshot.CaptureDisplay(100)
	fmt.Println(res, err)
}
