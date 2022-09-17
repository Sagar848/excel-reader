package excelparse

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// Send the file you want to parse
// File has to be provided with full path
func Parse(filename string) {

	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	activeSheetNumber := f.GetActiveSheetIndex()
	activeSheet := f.GetSheetName(activeSheetNumber)
	// fmt.Println("first sheet =", firstSheet)

	// Get value from the cell by given worksheet name and axis.
	cell, err := f.GetCellValue(activeSheet, "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows(activeSheet)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
