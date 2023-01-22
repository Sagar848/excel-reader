package excel

import (
	"log"
)

// Send the file you want to parse
// File has to be provided with full path
func ParseWithSpeceficHeaders(filename string, hasHeader bool, speceficHeaders []string) map[string][]map[string]interface{} {

	speceficColumnNumbers, err := GetColumnNumbers(speceficHeaders)
	if err != nil {
		log.Println("Error in getting specefic columns: ", err)
		return nil
	}

	return getDataMap(filename, hasHeader, func(header []string, row []string) map[string]interface{} {
		return createDataMapWithSpeceficheaders(header, row, speceficColumnNumbers)
	})
}
