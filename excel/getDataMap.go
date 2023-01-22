package excel

import "fmt"

func getDataMap(filename string, hasHeader bool, createMap func([]string, []string) map[string]interface{}) map[string][]map[string]interface{} {

	var header = make([]string, 0, 10)
	// var data = make([][]string, 0, 10)
	var dataMap map[string][]map[string]interface{} = map[string][]map[string]interface{}{}

	f := openFile(filename)

	sheetMap := f.GetSheetMap()
	for _, sheetName := range sheetMap {
		// activeSheet := f.GetSheetName(activeSheetNumber)
		rows := getRows(f, sheetName)
		var currentDataMap []map[string]interface{} = []map[string]interface{}{}
		for rowNum, row := range rows {

			if rowNum == 0 {
				if hasHeader {
					header = row
				} else {
					for cellNum := range row {
						header = append(header, fmt.Sprint("column", cellNum))
					}

					currentDataMap = append(currentDataMap, createMap(header, row))
				}
			} else {
				currentDataMap = append(currentDataMap, createMap(header, row))
			}
		}
		dataMap[sheetName] = currentDataMap
	}

	return dataMap
}
