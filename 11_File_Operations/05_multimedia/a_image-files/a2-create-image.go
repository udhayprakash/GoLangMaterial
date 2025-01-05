package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
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

	// Save the image as a PNG file
	pngFile, err := os.Create("a2-image.png")
	if err != nil {
		panic(err)
	}
	defer pngFile.Close()

	if err := png.Encode(pngFile, img); err != nil {
		panic(err)
	}

	// Save the image as a JPEG file
	jpegFile, err := os.Create("a2-image.jpg")
	if err != nil {
		panic(err)
	}
	defer jpegFile.Close()

	if err := jpeg.Encode(jpegFile, img, &jpeg.Options{Quality: 90}); err != nil {
		panic(err)
	}

	fmt.Println("Image saved in PNG and JPEG formats successfully!")
}
