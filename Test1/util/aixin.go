package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	// 打开 Excel 文件
	f, err := excelize.OpenFile("D:/2022DOC/新建文件夹/HY00001221.xlsx")
	f1 := excelize.NewFile()
	_ = f1.NewSheet("sheet1")
	for i := 0; i < 10; i++ {

	}
	f1.SetCellValue("Sheet1", "A1", "Hello")
	f1.SetCellValue("Sheet1", "B12", "Hello2")
	f1.SetCellValue("Sheet1", "C13", "Hello2")
	f1.SetCellValue("Sheet1", "D14", "Hello3")
	f1.SetCellValue("Sheet1", "E15", "Hello")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err1 := f1.SaveAs("D:/Go_Study/excel/example.xlsx"); err != nil {
		panic(err1)
	}

	// 读取单元格数据
	//cell := f.GetCellValue("Sheet1", "A1")

	//fmt.Println("A1:", cell)

	//cell = f.GetCellValue("Sheet1", "B2")

	//fmt.Println("B2:", cell)

	// 读取整个工作表的数据
	rows := f.GetRows("Sheet1")
	fmt.Println(rows[2])
	/*for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}*/
}
