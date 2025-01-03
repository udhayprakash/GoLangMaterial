package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CsvLine struct {
	Column1 string
	Column2 string
}

func main() {

	lines, err := ReadCsv("example.csv")
	if err != nil {
		panic(err)
	}

	// Loop through lines & turn into object
	for _, line := range lines {
		data := CsvLine{
			Column1: line[0],
			Column2: line[1],
		}
		fmt.Println(data.Column1 + " " + data.Column2)
	}
}

// ReadCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
