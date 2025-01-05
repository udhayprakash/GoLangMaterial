package main

import (
	"fmt"
	"image/color"

	"github.com/fogleman/gg"
)

func main() {
	// Create a new context with a blank image (800x600)
	dc := gg.NewContext(800, 600)

	// Set a background color
	dc.SetColor(color.White)
	dc.Clear()

	// Load a font (replace with the path to a TTF font file)
	if err := dc.LoadFontFace("arial.ttf", 48); err != nil {
		panic(err)
	}

	// Set text color
	dc.SetColor(color.Black)

	// Draw text on the image
	dc.DrawStringAnchored("Hello, World!", 400, 300, 0.5, 0.5)

	// Save the image to a file
	if err := dc.SavePNG("text_image.png"); err != nil {
		panic(err)
	}

	fmt.Println("Image with text created successfully!")
}
