package excel

func createDataMap(header []string, row []string) map[string]interface{} {
	currentData := map[string]interface{}{}
	if len(header) == len(row) {
		for i, cell := range row {
			currentData[header[i]] = cell
		}
	}
	return currentData
}

func createDataMapWithSpeceficheaders(header []string, row []string, speceficHeaders map[int]int) map[string]interface{} {
	currentData := map[string]interface{}{}
	if len(header) == len(row) {
		for i, cell := range row {
			if _, exists := speceficHeaders[i+1]; exists {
				currentData[header[i]] = cell
			}
		}
	}
	return currentData
}
