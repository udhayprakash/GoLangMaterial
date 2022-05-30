package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", 50)
	idx := f.NewSheet("Sheet2")
	fmt.Println("idx=", idx)

	f.SetCellValue("Sheet2", "A1", 50)
	f.SetActiveSheet(idx)
	if err := f.SaveAs("new_sheet.xlsx"); err != nil {
		log.Fatal(err)
	}
}
