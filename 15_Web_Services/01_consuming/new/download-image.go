package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// image link, open at browser
	image := "https://drive.google.com/open?id=1WPWH9zkQ25CASuh3uXs7T3i5PvsEfIKU"
	id := image[33:]
	url := "https://docs.google.com/uc?export=download&id=" + id

	// url := "https://drive.google.com/open?id=1WPWH9zkQ25CASuh3uXs7T3i5PvsEfIKU"

	fileName := "file.jpeg"
	fmt.Println("Downloading file...")

	output, err := os.Create(fileName)
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)

	fmt.Println(n, "bytes downloaded")
}
