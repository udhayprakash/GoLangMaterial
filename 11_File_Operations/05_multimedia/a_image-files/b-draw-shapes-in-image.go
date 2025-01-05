package main

import (
	"fmt"
	"image/color"

	"github.com/fogleman/gg"
)

// go get github.com/fogleman/gg

func main() {
	// Create a new context with a blank image (800x600)
	dc := gg.NewContext(800, 600)

	// Set a background color
	dc.SetColor(color.White)
	dc.Clear()

	// Draw a rectangle
	dc.SetColor(color.RGBA{255, 0, 0, 255}) // Red
	dc.DrawRectangle(100, 100, 200, 150)
	dc.Fill()

	// Draw a circle
	dc.SetColor(color.RGBA{0, 0, 255, 255}) // Blue
	dc.DrawCircle(400, 300, 100)
	dc.Fill()

	// Save the image to a file
	if err := dc.SavePNG("b-shapes.png"); err != nil {
		panic(err)
	}

	fmt.Println("Image with shapes created successfully!")
}
