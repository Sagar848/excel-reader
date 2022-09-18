package excelparse

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

// Send the file you want to parse
// File has to be provided with full path
func Parse(filename string, hasHeader bool) []map[string]interface{} {

	var header = make([]string, 0, 10)
	// var data = make([][]string, 0, 10)
	var dataMap []map[string]interface{} = []map[string]interface{}{}

	f, err := excelize.OpenFile(filename)
	if err != nil {
		log.Println("Error in opening file: ", err)
		return nil
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
		log.Println("Error in getting rows: ", err)
		return nil
	}
	for rowNum, row := range rows {

		if rowNum == 0 {
			if hasHeader {
				header = row
			} else {
				for cellNum := range row {
					header = append(header, fmt.Sprint("column", cellNum))
				}

				dataMap = append(dataMap, createDataMap(header, row))
			}
		} else {
			dataMap = append(dataMap, createDataMap(header, row))
		}

	}
	return dataMap
}

func createDataMap(header []string, row []string) map[string]interface{} {
	currentData := map[string]interface{}{}
	if len(header) == len(row) {
		for i, cell := range row {
			currentData[header[i]] = cell
		}
	}
	return currentData
}
