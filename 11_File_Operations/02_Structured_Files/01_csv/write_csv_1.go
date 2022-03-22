package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

var csvData = [][]string{
	{"SuperHero Name", "Power", "Weakness"},
	{"Batman", "Wealth", "Human"},
	{"Superman", "Strength", "Kryptonite"},
}

func main() {
	csvWriter()
	csvLineWriter()
	// Above functions will create superheroes.csv and superheroes2.csv files
}

func csvWriter() {
	// Open the file
	recordFile, err := os.Create("./superheroes.csv")
	if err != nil {
		fmt.Println("Error while creating the file::", err)
		return
	}

	// Initialize the writer
	writer := csv.NewWriter(recordFile)

	// Write all the records
	err = writer.WriteAll(csvData)
	if err != nil {
		fmt.Println("Error while writing to the file ::", err)
		return
	}

	err = recordFile.Close()
	if err != nil {
		fmt.Println("Error while closing the file ::", err)
		return
	}
}

func csvLineWriter() {
	// Open the file
	recordFile, err := os.Create("./superheroes2.csv")
	if err != nil {
		fmt.Println("Error while creating the file ::", err)
		return
	}

	// Initialize the writer
	writer := csv.NewWriter(recordFile)

	// Write all the records
	for i, csvLineData := range csvData {
		// You can manipulate the single slice at a time here
		// I'm adding a index on the last column
		if i == 0 {
			csvLineData = append(csvLineData, "Name Length")
		} else {
			nameLength := strconv.Itoa(len(csvLineData[0]))
			csvLineData = append(csvLineData, nameLength)
		}
		err = writer.Write(append(csvLineData, string(i)))
		if err != nil {
			fmt.Println("Error while writing to the file ::", err)
			return
		}

		writer.Flush()       // Writes the buffered data to the writer
		err = writer.Error() // Checks if any error occurred while writing
		if err != nil {
			fmt.Println("Error while writing to the file ::", err)
			return
		}
	}

	err = recordFile.Close()
	if err != nil {
		fmt.Println("Error while closing the file ::", err)
		return
	}
}
