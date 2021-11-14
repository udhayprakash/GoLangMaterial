package main

import (
	"compress/flate"
	"fmt"
	"io"
	"os"
)

func CheckError(e error, msg string) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

}

func main() {
	inputFile, err := os.Open("file.txt.compressed")
	CheckError(err, "Unable to Open file.txt.compressed")
	defer inputFile.Close()

	outputFile, err := os.Create("file.txt.decompressed")
	CheckError(err, "Unable to Create file.txt.decompressed")
	defer outputFile.Close()

	flateReader := flate.NewReader(inputFile)

	defer flateReader.Close()
	io.Copy(outputFile, flateReader)

}
