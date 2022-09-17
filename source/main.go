package main

import (
	"fmt"

	excelparse "github.com/Sagar848/excel-reader/excel"
)

func main() {
	fmt.Println("here we will read excel files")

	header, rows := excelparse.Parse("D:/files/DummyFile.xlsx", false)
	fmt.Println("header =", header)
	fmt.Println("rows =", rows)

}
