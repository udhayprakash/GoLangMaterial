package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Employee struct {
	Name string
	Age  int
	Job  string
}

func main() {
	// read data from file

	jsondatafromfile, err := ioutil.ReadFile("./data.json")

	if err != nil {
		fmt.Println(err)
	}

	// Unmarshal JSON data

	var jsondata []Employee

	err = json.Unmarshal([]byte(jsondatafromfile), &jsondata)

	if err != nil {
		fmt.Println(err)
	}

	csvdatafile, err := os.Create("./data.csv")

	if err != nil {
		fmt.Println(err)
	}
	defer csvdatafile.Close()

	writer := csv.NewWriter(csvdatafile)

	for _, worker := range jsondata {
		var record []string
		record = append(record, worker.Name)
		record = append(record, strconv.Itoa(worker.Age))
		record = append(record, worker.Job)
		writer.Write(record)
	}

	// remember to flush!
	writer.Flush()

}
