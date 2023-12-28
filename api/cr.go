package api

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

// Create a blank xlsx file cause this excelize library is weird
func Create() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	index, err := f.NewSheet("sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	f.SetActiveSheet(index)
	if err := f.SaveAs("data.xlsx"); err != nil {
		fmt.Println(err)
	}
}
