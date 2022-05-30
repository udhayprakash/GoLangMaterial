package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("simple.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	c1, err := f.GetCellValue("Sheet1", "A1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c1)

	c2, err := f.GetCellValue("Sheet1", "A4")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c2)

	c3, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c3)
}
