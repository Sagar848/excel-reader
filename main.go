package main

import (
	"fmt"

	excel "github.com/Sagar848/excel-reader/excel"
)

func main() {
	fmt.Println("here we will read excel files")

	dataMap := excel.Parse("D:/files/DummyFile.xlsx", false)
	fmt.Println("dataMap =", dataMap)

}
