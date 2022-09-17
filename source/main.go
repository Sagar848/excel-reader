package main

import (
	"fmt"

	excel "github.com/Sagar848/excel-reader/excel"
)

func main() {
	fmt.Println("here we will read excel files")

	header, rows := excel.Parse("D:/files/DummyFile.xlsx", false)
	fmt.Println("header =", header)
	fmt.Println("rows =", rows)

}
