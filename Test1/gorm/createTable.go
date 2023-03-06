package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type YourStruct struct {
	Column1 string `gorm:"column:column1"`
	Column2 int    `gorm:"column:column2"`
	Column3 bool   `gorm:"column:column3"`
}

func main() {
	var url = "root:236332@tcp(162.14.64.254:3306)/first_test?charset=utf8mb4&loc=Local"

	var db1, _ = gorm.Open(mysql.Open(url), &gorm.Config{})
	db1.Table("socket_users").First(&YourStruct{})

}
