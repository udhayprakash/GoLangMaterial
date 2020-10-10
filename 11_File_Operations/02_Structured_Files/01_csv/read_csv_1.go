package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Reading whole file at once")
	csvReaderAll()
	fmt.Println("===================================")
	fmt.Println("Reading single row at a time")
	csvReaderRow()
}

func csvReaderAll() {
	// Open the file
	recordFile, err := os.Open("records.csv")
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return
	}

	// Setup the reader
	reader := csv.NewReader(recordFile)

	// Read the records
	allRecords, err := reader.ReadAll()
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return
	}
	fmt.Println(allRecords)

	err = recordFile.Close()
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return
	}
}

func csvReaderRow() {
	// Open the file
	recordFile, err := os.Open("records.csv")
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return
	}

	// Setup the reader
	reader := csv.NewReader(recordFile)

	// Read the records
	header, err := reader.Read()
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return
	}
	fmt.Printf("Headers : %v \n", header)

	for i:= 0 ;; i = i + 1 {
		record, err := reader.Read()
		if err == io.EOF {
			break // reached end of the file
		} else if err != nil {
			fmt.Println("An error encountered ::", err)
			return
		}

		fmt.Printf("Row %d : %v \n", i, record)
	}

	// Note: Each time Read() is called, it reads the next line from the file
	// r1, _ := reader.Read() // Reads the first row, useful for headers
	// r2, _ := reader.Read() // Reads the second row
}