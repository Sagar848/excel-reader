package excel

import (
	"log"

	"github.com/xuri/excelize/v2"
)

func GetColumnNumbers(names []string) (map[int]int, error) {
	columnNumbers := map[int]int{}
	var err error

	for _, name := range names {
		columnNumber, err := excelize.ColumnNameToNumber(name)
		if err != nil {
			log.Fatalf("Error in getting column name for %v", name)
		}
		columnNumbers[columnNumber] = columnNumber
	}
	return columnNumbers, err
}
