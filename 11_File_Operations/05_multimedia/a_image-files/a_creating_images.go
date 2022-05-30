package main

import (
	log "github.com/sirupsen/logrus"
	"image"
	"image/png"
	"os"
)

func main() {
	myImg := image.NewRGBA(image.Rect(0, 0, 12, 6))
	out, err := os.Create("image1.png")
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(out, myImg)
	out.Close()

	myImg = image.RGBA{
		Pix:    [3]int8{200, 255, 100},
		Stride: 0,
		Rect:   image.Rectangle{5, 10},
	}
	out, err = os.Create("image2.png")
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(out, myImg)
	out.Close()
}
