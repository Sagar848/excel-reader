package excel

// Send the file you want to parse
// File has to be provided with full path
func Parse(filename string, hasHeader bool) []map[string]interface{} {

	return getDataMap(filename, hasHeader, func(header []string, row []string) map[string]interface{} {
		return createDataMap(header, row)
	})
}
