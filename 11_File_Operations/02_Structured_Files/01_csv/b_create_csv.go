package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fileHandler, err := os.Create("personalDetails2.csv")
	if err != nil {
		fmt.Println("Error while creating the file::", err)
	}
	defer fileHandler.Close()

	// Initialize the csv writer
	writer := csv.NewWriter(fileHandler)

	var content2File = [][]string{
		{"Person", "age", "Location", "country"},
		{"Udhay", "30", "Hyderabad", "India"},
		{"Prakash", "99", "Secunderabad", "India"},
	}

	err = writer.WriteAll(content2File)
	if err != nil {
		fmt.Println("Error while writing to the file ::", err)
		return
	}

}
