package excelparse

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// Send the file you want to parse
// File has to be provided with full path
func Parse(filename string, hasHeader bool) ([]string, [][]string) {

	var header = make([]string, 0, 10)
	var data = make([][]string, 0, 10)

	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return header, data
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	activeSheetNumber := f.GetActiveSheetIndex()
	activeSheet := f.GetSheetName(activeSheetNumber)
	// fmt.Println("first sheet =", firstSheet)

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows(activeSheet)
	if err != nil {
		fmt.Println(err)
		return header, data
	}
	for rowNum, row := range rows {

		if rowNum == 0 {
			if hasHeader {
				header = row
			} else {
				for cellNum := range row {
					header = append(header, fmt.Sprint("column", cellNum))
				}
				data = append(data, row)
			}
		} else {
			data = append(data, row)
		}

	}
	return header, data
}
