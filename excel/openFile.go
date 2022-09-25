package excel

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func openFile(filename string) *excelize.File {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		log.Println("Error in opening file: ", err)
		return nil
	} else {
		log.Println("Opened new file", filename)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	return f
}
