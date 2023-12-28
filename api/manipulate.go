package api

import (
	"fmt"
	"log"

	"github.com/sigurn2/WorldHeritage_GO/data"
	"github.com/xuri/excelize/v2"
)

const sheet = "sheet1"
const newColPos = "AL"

func Open() (*excelize.File, error) {
	f, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil, excelize.ErrSheetNotExist{
			SheetName: "file not exist",
		}
	}

	return f, nil
}

func getColNames(f *excelize.File) []string {
	cols, err := f.GetCols(sheet)
	if err != nil {
		log.Fatal(err)
	}
	colNames := []string{}
	for _, col := range cols {
		colNames = append(colNames, col[0])
	}
	return colNames
}

func isExistedCol(f *excelize.File, rowName string) bool {
	c := getColNames(f)
	for _, v := range c {
		if rowName == v {
			return true
		}
	}
	return false
}

func AddCol(f *excelize.File, row data.Attribute) bool {

	// if row not in existed rows
	if !isExistedCol(f, row.Name) {
		err := f.InsertCols(sheet, "AL", 1)
		if err != nil {
			return false
		}
	} else {
		return false
	}
	err := f.SetCellValue(sheet, "AL1", row.Name)
	if err != nil {
		return false
	}
	if err := f.Save(); err != nil {
		fmt.Println(err)
	}
	return true
}

func GetTypes(f *excelize.File) ([]string, error) {
	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, err
	}
	res := make([]string, 0)
	for _, row := range rows {
		res = append(res, row[28])
	}
	res = res[1:]
	return res, nil
}

func GetHeritages(f *excelize.File) ([]string, error) {
	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, err
	}
	res := make([]string, 0)
	for _, row := range rows {
		res = append(res, row[3])
	}
	res = res[1:]
	return res, nil
}

func WriteValue(f *excelize.File, col int, value string) {
	newPos := fmt.Sprintf(newColPos+"%v", col)
	err := f.SetCellValue(sheet, newPos, value)
	if err != nil {
		return
	}
	if err := f.Save(); err != nil {
		fmt.Println(err)
	}
}
