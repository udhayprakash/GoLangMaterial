package main

import (
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {

	f := excelize.NewFile()

	f.SetCellValue("Sheet1", "A1", "Sunny Day")
	f.MergeCell("Sheet1", "A1", "B2")

	style, _ := f.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"}, 
        "font":{"bold":true,"italic":true}}`)

	f.SetCellStyle("Sheet1", "A1", "B2", style)

	if err := f.SaveAs("merging.xlsx"); err != nil {
		log.Fatal(err)
	}
}
