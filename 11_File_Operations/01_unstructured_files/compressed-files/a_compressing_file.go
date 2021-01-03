package main

import (
	"compress/flate"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, err := os.Open("file.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer inputFile.Close()

	outputFile, err := os.Create("file.txt.compressed")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer outputFile.Close()

	flateWriter, err := flate.NewWriter(outputFile, flate.BestCompression)

	if err != nil {
		fmt.Println("NewWriter error ", err)
		os.Exit(1)
	}

	defer flateWriter.Close()
	io.Copy(flateWriter, inputFile)

	flateWriter.Flush()
}
