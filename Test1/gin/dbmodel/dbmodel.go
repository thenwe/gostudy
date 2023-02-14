package dbmodel

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"time"
)

type SocketUser struct{
	Id string
	Name string
	Password string
	CreateAt time.Time
	UpdateAt time.Time
}

func DbSocket(){
	dsn:="root:236332@tcp(localhost:3306)/first_test?charset=utf8mb4&loc=Local"
	db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{})

	if err!=nil{
		fmt.Println("连接失败:",err)
	}
	err=db.AutoMigrate(&SocketUser{})
	if err != nil{
		fmt.Println("创建失败:",err)
	}
	fmt.Println("创建成功")
}
