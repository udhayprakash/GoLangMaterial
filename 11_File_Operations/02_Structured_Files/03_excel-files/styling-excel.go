package main

import (
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {

	f := excelize.NewFile()

	f.SetCellValue("Sheet1", "A1", "an old falcon")
	f.SetColWidth("Sheet1", "A", "A", 20)

	style, _ := f.NewStyle(`{"alignment":{"horizontal":"center"}, 
        "font":{"bold":true,"italic":true}}`)

	f.SetCellStyle("Sheet1", "A1", "A1", style)

	if err := f.SaveAs("styled.xlsx"); err != nil {
		log.Fatal(err)
	}
}
