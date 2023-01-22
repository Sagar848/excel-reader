package main

import (
	"fmt"

	"github.com/Sagar848/excel-reader/excel"
)

func main() {
	fmt.Println("here we will read excel files")

	dataMap := excel.Parse("D:/files/DummyFile.xlsx", true)
	// dataMap := excel.ParseWithSpeceficHeaders("D:/files/DummyFile.xlsx", true, []string{"A", "C"})
	fmt.Println("dataMap =", dataMap)

	// f, err := excelize.OpenFile("D:/files/DummyFile.xlsx")
	// if err != nil {
	// 	log.Println("Error in opening file: ", err)
	// } else {
	// 	log.Println("Opened new file")
	// }
	// sheetMap := f.GetSheetMap()
	// fmt.Println("sheetMap =", sheetMap)
	// defer func() {
	// 	if err := f.Close(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()

}
