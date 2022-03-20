package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	cwd, err := os.Getwd()
	fmt.Println("cwd:", cwd)

	//fileHandler, err := os.Open("fileNotPresent.csv")
	fileHandler, err := os.Open("personDetails.csv")
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println("fileHandler:", fileHandler)
	defer fileHandler.Close()

	// Parse the csv file
	reader := csv.NewReader(fileHandler)

	names := []string{}
	// Iterate through the records
	for {
		recordRow, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		//fmt.Println("recordRow:", recordRow, len(recordRow))
		fmt.Println("Personal Name:", recordRow[0])
		names = append(names, recordRow[0])
	}
	fmt.Println("All names:", names)
}
