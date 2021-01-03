package main

import (
	"compress/flate"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, err := os.Open("file.txt.compressed")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer inputFile.Close()

	outputFile, err := os.Create("file.txt.decompressed")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer outputFile.Close()

	flateReader := flate.NewReader(inputFile)

	defer flateReader.Close()
	io.Copy(outputFile, flateReader)

}
