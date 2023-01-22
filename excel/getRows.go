package excel

import (
	"log"

	"github.com/xuri/excelize/v2"
)

func getRows(f *excelize.File, sheetName string) [][]string {

	// fmt.Println("first sheet =", firstSheet)

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Println("Error in getting rows: ", err)
		return nil
	}
	return rows
}
