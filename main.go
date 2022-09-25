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

}
