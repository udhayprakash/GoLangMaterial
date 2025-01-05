package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	// Create a blank image (800x600) with a white background
	width, height := 800, 600
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Fill the image with a white background
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.White)
		}
	}

	// Save the image to a file
	file, err := os.Create("a1-blank_image.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		panic(err)
	}

	fmt.Println("Blank image created successfully!")
}
