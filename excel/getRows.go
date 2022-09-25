package excel

import "log"

func getRows(filename string) [][]string {

	f := openFile(filename)

	activeSheetNumber := f.GetActiveSheetIndex()
	activeSheet := f.GetSheetName(activeSheetNumber)
	// fmt.Println("first sheet =", firstSheet)

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows(activeSheet)
	if err != nil {
		log.Println("Error in getting rows: ", err)
		return nil
	}
	return rows
}
