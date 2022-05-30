package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	fileHandler, err := os.Create("personDetails.csv")
	checkError("Cannot create file", err)
	defer fileHandler.Close()

	writer := csv.NewWriter(fileHandler)
	defer writer.Flush()

	var content2File = [][]string{
		{"Person", "age", "Location", "country"},
		{"Udhay", "30", "Hyderabad", "India"},
		{"Prakash", "99", "Secunderabad", "India"},
	}

	for _, value := range content2File {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}

}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
