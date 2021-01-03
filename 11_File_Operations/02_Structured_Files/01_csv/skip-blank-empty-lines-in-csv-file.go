package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	csvDataWithEmptyLines := `first_name,last_name,username
 "Adam","Sandler",     adam


 Steve,     McQueen,steve
 "Robert","Spacey","robspacy"
 `
	csvReader := csv.NewReader(strings.NewReader(csvDataWithEmptyLines))

	// add these
	csvReader.FieldsPerRecord = -1
	csvReader.TrimLeadingSpace = true

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}
