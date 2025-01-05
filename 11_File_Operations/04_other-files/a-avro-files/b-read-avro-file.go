package main

import (
	"fmt"
	"os"

	"github.com/linkedin/goavro/v2"
)

func main() {
	// Open the Avro file
	file, err := os.Open("users.avro")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create an Avro OCF reader
	ocfReader, err := goavro.NewOCFReader(file)
	if err != nil {
		fmt.Println("Error creating OCF reader:", err)
		return
	}

	// Read and print each record
	for ocfReader.Scan() {
		record, err := ocfReader.Read()
		if err != nil {
			fmt.Println("Error reading record:", err)
			return
		}
		fmt.Println("Record:", record)
	}

	if err := ocfReader.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Data read from users.avro successfully!")
}
