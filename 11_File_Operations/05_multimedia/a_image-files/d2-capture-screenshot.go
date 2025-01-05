package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)
// go get github.com/kbinani/screenshot

func main() {
	// Get the number of displays
	n := screenshot.NumActiveDisplays()
	if n <= 0 {
		panic("No active displays found")
	}

	// Capture the primary display (display 0)
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}

	// Save the screenshot to a file
	file, err := os.Create("screenshot.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		panic(err)
	}

	fmt.Println("Screenshot captured successfully!")
}
