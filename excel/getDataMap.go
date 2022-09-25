package excel

import "fmt"

func getDataMap(filename string, hasHeader bool, createMap func([]string, []string) map[string]interface{}) []map[string]interface{} {

	var header = make([]string, 0, 10)
	// var data = make([][]string, 0, 10)
	var dataMap []map[string]interface{} = []map[string]interface{}{}

	rows := getRows(filename)

	for rowNum, row := range rows {

		if rowNum == 0 {
			if hasHeader {
				header = row
			} else {
				for cellNum := range row {
					header = append(header, fmt.Sprint("column", cellNum))
				}

				dataMap = append(dataMap, createMap(header, row))
			}
		} else {
			dataMap = append(dataMap, createMap(header, row))
		}
	}
	return dataMap
}
