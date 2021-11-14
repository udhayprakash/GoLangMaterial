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
	inputFile, err := os.Open("file.txt")
	CheckError(err, "Unable to Open file.txt")
	defer inputFile.Close()

	outputFile, err := os.Create("file.txt.compressed")
	CheckError(err, "Unable to Create file.txt.compressed")
	defer outputFile.Close()

	flateWriter, err := flate.NewWriter(outputFile, flate.BestCompression)
	CheckError(err, "Unable to write to new file")
	defer flateWriter.Close()
	io.Copy(flateWriter, inputFile)

	flateWriter.Flush()
}
